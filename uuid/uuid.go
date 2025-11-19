package uuid

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	NodeBits uint8 = 16
	StepBits uint8 = 12

	nodeMax       = -1 ^ (-1 << NodeBits)
	nodeMask      = nodeMax << StepBits
	stepMask      = -1 ^ (-1 << StepBits)
	sequenceShift = NodeBits + StepBits
	nodeShift     = StepBits
)

type Node struct {
	mu       sync.Mutex
	sequence int64
	node     int64
	step     int64

	coll *mongo.Collection
	kind string
}

func NewNode(node int64, coll *mongo.Collection, kind string) (*Node, error) {
	n := Node{}
	n.node = node
	n.kind = kind
	n.coll = coll

	if n.node < 0 || n.node > nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(nodeMax, 10))
	}

	if err := n.update(); err != nil {
		return nil, err
	}

	return &n, nil
}

func (n *Node) Generate() (ID, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.step = (n.step + 1) & stepMask
	if n.step == 0 {
		if err := n.update(); err != nil {
			return 0, err
		}
	}

	id := ID((n.sequence << sequenceShift) |
		(n.node << nodeShift) |
		(n.step),
	)

	return id, nil
}

func (n *Node) update() error {
	var doc struct {
		N int64 `bson:"n"`
	}

	filter := bson.D{bson.E{Key: "_id", Value: n.kind}}
	update := bson.D{
		bson.E{
			Key: "$inc",
			Value: bson.D{
				bson.E{Key: "n", Value: 1},
			},
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	result := n.coll.FindOneAndUpdate(context.Background(), filter, update, opts)
	if err := result.Decode(&doc); err != nil {
		return err
	}

	n.sequence = doc.N

	return nil
}

type ID int64

func (id ID) Int64() int64 {
	return int64(id)
}

func ParseInt64(id int64) ID {
	return ID(id)
}

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err
}

func (id ID) Sequence() int64 {
	return (int64(id) >> sequenceShift)
}

func (id ID) Node() int64 {
	return int64(id) & nodeMask >> nodeShift
}

func (id ID) Step() int64 {
	return int64(id) & stepMask
}
