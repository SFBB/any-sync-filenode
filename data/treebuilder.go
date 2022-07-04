package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/data/pb"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/data/threadmodels"
	//"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/util/slice"
	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/common/log"
	"github.com/textileio/go-threads/core/thread"
	"time"
)

var (
	//log      = logging.Logger("anytype-data")
	ErrEmpty = errors.New("logs empty")
)

type TreeBuilder struct {
	cache                map[string]*Change
	logHeads             map[string]*Change
	identityKeys         map[string]threadmodels.SigningPubKey
	signingPubKeyDecoder threadmodels.SigningPubKeyDecoder
	tree                 *Tree
	thread               threadmodels.Thread
}

func NewTreeBuilder(t threadmodels.Thread, decoder threadmodels.SigningPubKeyDecoder) *TreeBuilder {
	return &TreeBuilder{
		cache:                make(map[string]*Change),
		logHeads:             make(map[string]*Change),
		identityKeys:         make(map[string]threadmodels.SigningPubKey),
		signingPubKeyDecoder: decoder,
		tree:                 &Tree{}, // TODO: add NewTree method
		thread:               t,
	}
}

func (tb *TreeBuilder) loadChange(id string) (ch *Change, err error) {
	if ch, ok := tb.cache[id]; ok {
		return ch, nil
	}

	// TODO: Add virtual changes logic
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	change, err := tb.thread.GetChange(ctx, id)
	if err != nil {
		return nil, err
	}

	aclChange := new(pb.ACLChange)

	// TODO: think what should we do with such cases, because this can be used by attacker to break our tree
	if err = proto.Unmarshal(change.Payload, aclChange); err != nil {
		return
	}
	var verified bool
	verified, err = tb.verify(aclChange.Identity, change.Payload, change.Signature)
	if err != nil {
		return
	}
	if !verified {
		err = fmt.Errorf("the signature of the payload cannot be verified")
		return
	}

	ch, err = NewChange(id, aclChange)
	tb.cache[id] = ch

	return ch, nil
}

func (tb *TreeBuilder) verify(identity string, payload, signature []byte) (isVerified bool, err error) {
	identityKey, exists := tb.identityKeys[identity]
	if !exists {
		identityKey, err = tb.signingPubKeyDecoder.DecodeFromString(identity)
		if err != nil {
			return
		}
		tb.identityKeys[identity] = identityKey
	}
	return identityKey.Verify(payload, signature)
}

func (tb *TreeBuilder) Build(fromStart bool) (*Tree, error) {
	heads := tb.thread.Heads()

	if fromStart {
		if err := tb.buildTreeFromStart(heads); err != nil {
			return nil, fmt.Errorf("buildTree error: %v", err)
		}
	} else {
		breakpoint, err := tb.findBreakpoint(heads)
		if err != nil {
			return nil, fmt.Errorf("findBreakpoint error: %v", err)
		}

		if err = tb.buildTree(heads, breakpoint); err != nil {
			return nil, fmt.Errorf("buildTree error: %v", err)
		}
	}

	tb.cache = nil

	return tb.tree, nil
}

func (tb *TreeBuilder) buildTreeFromStart(heads []string) (err error) {
	changes, possibleRoots, err := tb.dfsFromStart(heads)
	if len(possibleRoots) == 0 {
		return fmt.Errorf("cannot have tree without root")
	}
	root, err := tb.getRoot(possibleRoots)
	if err != nil {
		return err
	}

	tb.tree.AddFast(root)
	tb.tree.AddFast(changes...)
	return
}

func (tb *TreeBuilder) dfsFromStart(heads []string) (buf []*Change, possibleRoots []*Change, err error) {
	stack := make([]string, len(heads), len(heads)*2)
	copy(stack, heads)

	buf = make([]*Change, 0, len(stack)*2)
	uniqMap := make(map[string]struct{})
	for len(stack) > 0 {
		id := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, exists := uniqMap[id]; exists {
			continue
		}

		ch, err := tb.loadChange(id)
		if err != nil {
			continue
		}

		uniqMap[id] = struct{}{}
		buf = append(buf, ch)

		for _, prev := range ch.PreviousIds {
			stack = append(stack, prev)
		}
		if len(ch.PreviousIds) == 0 {
			possibleRoots = append(possibleRoots, ch)
		}
	}
	return buf, possibleRoots, nil
}

func (tb *TreeBuilder) buildTree(heads []string, breakpoint string) (err error) {
	ch, err := tb.loadChange(breakpoint)
	if err != nil {
		return
	}
	tb.tree.AddFast(ch)
	changes, err := tb.dfs(heads, breakpoint)

	tb.tree.AddFast(changes...)
	return
}

func (tb *TreeBuilder) dfs(heads []string, breakpoint string) (buf []*Change, err error) {
	stack := make([]string, len(heads), len(heads)*2)
	copy(stack, heads)

	buf = make([]*Change, 0, len(stack)*2)
	uniqMap := map[string]struct{}{breakpoint: {}}
	for len(stack) > 0 {
		id := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, exists := uniqMap[id]; exists {
			continue
		}

		ch, err := tb.loadChange(id)
		if err != nil {
			continue
		}

		uniqMap[id] = struct{}{}
		buf = append(buf, ch)

		for _, prev := range ch.PreviousIds {
			stack = append(stack, prev)
		}
	}
	return buf, nil
}

func (tb *TreeBuilder) findBreakpoint(heads []string) (breakpoint string, err error) {
	var (
		ch          *Change
		snapshotIds []string
	)
	for _, head := range heads {
		if ch, err = tb.loadChange(head); err != nil {
			return
		}
		shId := ch.SnapshotId
		if slice.FindPos(snapshotIds, shId) == -1 {
			snapshotIds = append(snapshotIds, shId)
		}
	}
	return tb.findCommonSnapshot(snapshotIds)
}

func (tb *TreeBuilder) findCommonSnapshot(snapshotIds []string) (snapshotId string, err error) {
	if len(snapshotIds) == 1 {
		return snapshotIds[0], nil
	} else if len(snapshotIds) == 0 {
		return "", fmt.Errorf("snapshots not found")
	}

	for len(snapshotIds) > 1 {
		l := len(snapshotIds)
		shId, e := tb.findCommonForTwoSnapshots(snapshotIds[l-2], snapshotIds[l-1])
		if e != nil {
			return "", e
		}
		snapshotIds[l-2] = shId
		snapshotIds = snapshotIds[:l-1]
	}
	return snapshotIds[0], nil
}

func (tb *TreeBuilder) findCommonForTwoSnapshots(s1, s2 string) (s string, err error) {
	// fast cases
	if s1 == s2 {
		return s1, nil
	}
	ch1, err := tb.loadChange(s1)
	if err != nil {
		return "", err
	}
	if ch1.SnapshotId == s2 {
		return s2, nil
	}
	ch2, err := tb.loadChange(s2)
	if err != nil {
		return "", err
	}
	if ch2.SnapshotId == s1 {
		return s1, nil
	}
	if ch1.SnapshotId == ch2.SnapshotId && ch1.SnapshotId != "" {
		return ch1.SnapshotId, nil
	}
	// traverse
	var t1 = make([]string, 0, 5)
	var t2 = make([]string, 0, 5)
	t1 = append(t1, ch1.Id, ch1.SnapshotId)
	t2 = append(t2, ch2.Id, ch2.SnapshotId)
	for {
		lid1 := t1[len(t1)-1]
		if lid1 != "" {
			l1, e := tb.loadChange(lid1)
			if e != nil {
				return "", e
			}
			if l1.SnapshotId != "" {
				if slice.FindPos(t2, l1.SnapshotId) != -1 {
					return l1.SnapshotId, nil
				}
			}
			t1 = append(t1, l1.SnapshotId)
		}
		lid2 := t2[len(t2)-1]
		if lid2 != "" {
			l2, e := tb.loadChange(t2[len(t2)-1])
			if e != nil {
				return "", e
			}
			if l2.SnapshotId != "" {
				if slice.FindPos(t1, l2.SnapshotId) != -1 {
					return l2.SnapshotId, nil
				}
			}
			t2 = append(t2, l2.SnapshotId)
		}
		if lid1 == "" && lid2 == "" {
			break
		}
	}

	log.Warnf("changes build tree: possible versions split")

	// prefer not first snapshot
	if len(ch1.PreviousIds) == 0 && len(ch2.PreviousIds) > 0 {
		log.Warnf("changes build tree: prefer %s(%d prevIds) over %s(%d prevIds)", s2, len(ch2.PreviousIds), s1, len(ch1.PreviousIds))
		return s2, nil
	} else if len(ch1.PreviousIds) > 0 && len(ch2.PreviousIds) == 0 {
		log.Warnf("changes build tree: prefer %s(%d prevIds) over %s(%d prevIds)", s1, len(ch1.PreviousIds), s2, len(ch2.PreviousIds))
		return s1, nil
	}

	isEmptySnapshot := func(ch *Change) bool {
		// TODO: add more sophisticated checks in Change for snapshots
		return !ch.IsSnapshot
	}

	// TODO: can we even have empty snapshots?
	// prefer not empty snapshot
	if isEmptySnapshot(ch1) && !isEmptySnapshot(ch2) {
		log.Warnf("changes build tree: prefer %s(not empty) over %s(empty)", s2, s1)
		return s2, nil
	} else if isEmptySnapshot(ch2) && !isEmptySnapshot(ch1) {
		log.Warnf("changes build tree: prefer %s(not empty) over %s(empty)", s1, s2)
		return s1, nil
	}

	// TODO: add virtual change mechanics
	// unexpected behavior - just return lesser id
	if s1 < s2 {
		log.Warnf("changes build tree: prefer %s (%s<%s)", s1, s1, s2)
		return s1, nil
	}
	log.Warnf("changes build tree: prefer %s (%s<%s)", s2, s2, s1)

	return s2, nil
}

func (tb *TreeBuilder) getRoot(possibleRoots []*Change) (*Change, error) {
	threadId, err := thread.Decode(tb.thread.ID())
	if err != nil {
		return nil, err
	}

	for _, r := range possibleRoots {
		id := r.Content.Identity
		sk, err := tb.signingPubKeyDecoder.DecodeFromString(id)
		if err != nil {
			continue
		}

		res, err := threadmodels.VerifyACLThreadID(sk, threadId)
		if err != nil {
			continue
		}

		if res {
			return r, nil
		}
	}
	return nil, fmt.Errorf("could not find any root")
}
