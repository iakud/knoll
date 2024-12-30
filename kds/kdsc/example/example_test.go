package example

import (
	// "github.com/iakud/knoll/kds/kdsc/example/pb"
	"github.com/iakud/knoll/kds/kdspb"
	"testing"

	"google.golang.org/protobuf/proto"
)

type FieldMask struct {
	Number int
	FieldMasks []FieldMask
}

type MapFieldMask[T interface{}] struct {
	Key T
	FieldMasks []FieldMask
}

func TestExample(t *testing.T) {
	// player := new(pb.PlayerTest)

	// protoreflect.

	// player.Info = new(pb.BasicInfo)
	buf, err := Marshal(t)
	if err != nil {
		panic(err)
	}
	err = Unmashal(t, buf)
	if err != nil {
		panic(err)
	}
}

func Unmashal(t *testing.T, buf []byte) error {
	mask := new(kdspb.FieldMask)
	proto.Unmarshal(buf, mask)
	t.Log("buf:", string(buf), ",len:", len(buf))
	t.Log("mask:", mask)
	return nil
}

func Marshal(t *testing.T) ([]byte, error) {
	mask1 := new(kdspb.FieldMask)
	mask1.Number = 1
	// mask.Field = &kdspb.FieldMask_Number{Number: 1}

	mapMask1 := new(kdspb.MapMask)
	mapMask1.Clear = true
	mask1.MapMask = mapMask1

	mask2 := new(kdspb.FieldMask)
	mask2.Number = 3

	mapMask2 := new(kdspb.MapMask)
	clearKeys := new(kdspb.Int32Array)
	clearKeys.Values = append(clearKeys.Values, 3, 5, 9 ,111)

	mapMask2.DeleteKeys = &kdspb.MapMask_Int32DeleteKeys{Int32DeleteKeys: clearKeys}
	mask2.MapMask = mapMask2

	mask := new(kdspb.FieldMask)
	mask.FieldMasks = append(mask.FieldMasks, mask1, mask2)

	return proto.Marshal(mask)
}