package example

import (
	"github.com/iakud/keeper/kds/kdsc/example/examplepb"
	"github.com/iakud/keeper/kds/kdspb"
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
	player := new(examplepb.PlayerTest)

	// protoreflect.

	player.Info = new(examplepb.BasicInfo)
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
	/*
	switch field := mask.GetField().(type) {
	case *kdspb.FieldMask_Number:
		t.Log("Number:", field.Number)
	}*/
	t.Log("Number:", mask.Number)
	for _, fieldMask := range mask.FieldMasks {
		switch field := fieldMask.GetField().(type) {
		//case *kdspb.FieldMask_Number:
		//	t.Log("Number:", field.Number)
		case *kdspb.FieldMask_DelInt32Key:
			t.Log("DelInt32Key:", field.DelInt32Key)
		case *kdspb.FieldMask_StringKey:
			t.Log("StringKey:", field.StringKey)
		}	
	}
	return nil
}

func Marshal(t *testing.T) ([]byte, error) {
	mask := new(kdspb.FieldMask)
	mask.Number = 1
	// mask.Field = &kdspb.FieldMask_Number{Number: 1}

	mapMask1 := new(kdspb.FieldMask)
	mapMask1.Field = &kdspb.FieldMask_DelInt32Key{DelInt32Key: 7}
	mask.FieldMasks = append(mask.FieldMasks, mapMask1)

	mapMask2 := new(kdspb.FieldMask)
	mapMask2.Field = &kdspb.FieldMask_StringKey{StringKey: "321"}
	mask.FieldMasks = append(mask.FieldMasks, mapMask2)

	return proto.Marshal(mask)
}