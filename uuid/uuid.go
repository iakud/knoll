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

var (
	NodeBits uint8 = 16
	StepBits uint8 = 12
)

type Node struct {
	mu       sync.Mutex
	sequence int64
	node     int64
	step     int64

	coll *mongo.Collection
	tag  string

	nodeMax       int64
	nodeMask      int64
	stepMask      int64
	sequenceShift uint8
	nodeShift     uint8
}

type ID int64

func NewNode(node int64, coll *mongo.Collection, tag string) (*Node, error) {
	n := Node{}
	n.node = node
	n.tag = tag
	n.coll = coll
	n.nodeMax = -1 ^ (-1 << NodeBits)
	n.nodeMask = n.nodeMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.sequenceShift = NodeBits + StepBits
	n.nodeShift = StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(n.nodeMax, 10))
	}

	if err := n.update(); err != nil {
		return nil, err
	}

	return &n, nil
}

func (n *Node) Generate() (ID, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.step = (n.step + 1) & n.stepMask
	if n.step == 0 {
		if err := n.update(); err != nil {
			return 0, err
		}
	}

	r := ID(n.sequence<<n.sequenceShift |
		(n.node << n.nodeShift) |
		(n.step),
	)

	return r, nil
}

func (n *Node) update() error {
	var doc struct {
		N int64 `bson:"n"`
	}

	filter := bson.D{bson.E{Key: "_id", Value: n.tag}}
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

func (f ID) Int64() int64 {
	return int64(f)
}

func ParseInt64(id int64) ID {
	return ID(id)
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err
}
