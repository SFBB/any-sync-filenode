package index

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func (ri *redisIndex) FileUnbind(ctx context.Context, key Key, fileId string) (err error) {
	var (
		sk = spaceKey(key)
		gk = groupKey(key)
	)
	_, gRelease, err := ri.AcquireKey(ctx, gk)
	if err != nil {
		return
	}
	defer gRelease()
	_, sRelease, err := ri.AcquireKey(ctx, sk)
	if err != nil {
		return
	}
	defer sRelease()

	// get file entry
	fileInfo, isNewFile, err := ri.getFileEntry(ctx, key, fileId)
	if err != nil {
		return
	}
	if isNewFile {
		// means file doesn't exist
		return nil
	}

	// fetch cids
	cids, err := ri.CidEntriesByString(ctx, fileInfo.Cids)
	if err != nil {
		return err
	}
	defer cids.Release()

	// fetch cid refs in one pipeline
	var (
		groupCidRefs = make([]*redis.StringCmd, len(cids.entries))
		spaceCidRefs = make([]*redis.StringCmd, len(cids.entries))
	)
	_, err = ri.cl.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i, c := range cids.entries {
			groupCidRefs[i] = pipe.HGet(ctx, gk, cidKey(c.Cid))
			spaceCidRefs[i] = pipe.HGet(ctx, sk, cidKey(c.Cid))
		}
		return nil
	})
	if err != nil {
		return
	}

	// load group and space info
	spaceInfo, err := ri.getSpaceEntry(ctx, key)
	if err != nil {
		return
	}
	groupInfo, err := ri.getGroupEntry(ctx, key)
	if err != nil {
		return
	}

	// update info and calculate changes
	var (
		groupRemoveKeys = make([]string, 0, len(cids.entries))
		spaceRemoveKeys = make([]string, 0, len(cids.entries))
		groupDecrKeys   = make([]string, 0, len(cids.entries))
		spaceDecrKeys   = make([]string, 0, len(cids.entries))
		affectedCidIdx  = make([]int, 0, len(cids.entries))
	)

	spaceInfo.FileCount--
	for i, c := range cids.entries {
		res, err := groupCidRefs[i].Result()
		if err != nil {
			return err
		}
		ck := cidKey(c.Cid)
		if res == "1" {
			groupRemoveKeys = append(groupRemoveKeys, ck)
			groupInfo.Size_ -= c.Size_
			groupInfo.CidCount--
			affectedCidIdx = append(affectedCidIdx, i)
		} else {
			groupDecrKeys = append(groupDecrKeys, ck)
		}
		res, err = spaceCidRefs[i].Result()
		if err != nil {
			return err
		}
		if res == "1" {
			spaceRemoveKeys = append(spaceRemoveKeys, ck)
			spaceInfo.Size_ -= c.Size_
			spaceInfo.CidCount--
		} else {
			spaceDecrKeys = append(spaceDecrKeys, ck)
		}
	}

	// do updates in one tx
	_, err = ri.cl.TxPipelined(ctx, func(tx redis.Pipeliner) error {
		tx.HDel(ctx, sk, fileKey(fileId))
		if len(spaceRemoveKeys) != 0 {
			tx.HDel(ctx, sk, spaceRemoveKeys...)
		}
		if len(groupRemoveKeys) != 0 {
			tx.HDel(ctx, gk, groupRemoveKeys...)
		}
		if len(spaceDecrKeys) != 0 {
			for _, k := range spaceDecrKeys {
				tx.HIncrBy(ctx, sk, k, -1)
			}
		}
		if len(groupDecrKeys) != 0 {
			for _, k := range groupDecrKeys {
				tx.HIncrBy(ctx, gk, k, -1)
			}
		}
		spaceInfo.Save(ctx, key, tx)
		groupInfo.Save(ctx, key, tx)
		return nil
	})

	// update cids
	for _, idx := range affectedCidIdx {
		cids.entries[idx].RemoveGroupId(key.GroupId)
		if saveErr := cids.entries[idx].Save(ctx, ri.cl); saveErr != nil {
			log.WarnCtx(ctx, "unable to save cid info", zap.Error(saveErr), zap.String("cid", cids.entries[idx].Cid.String()))
		}
	}
	return
}
