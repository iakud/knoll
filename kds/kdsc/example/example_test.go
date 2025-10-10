package example

import (
	// "github.com/iakud/knoll/kds/kdsc/example/pb"
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

/*
	type FieldMask struct {
		Number     int
		FieldMasks []FieldMask
	}

	type MapFieldMask[T interface{}] struct {
		Key        T
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
		clearKeys.Values = append(clearKeys.Values, 3, 5, 9, 111)

		mapMask2.DeleteKeys = &kdspb.MapMask_Int32DeleteKeys{Int32DeleteKeys: clearKeys}
		mask2.MapMask = mapMask2

		mask := new(kdspb.FieldMask)
		mask.FieldMasks = append(mask.FieldMasks, mask1, mask2)

		return proto.Marshal(mask)
	}
*/
func TestMarshalMessage(t *testing.T) {
	player := kds.NewPlayer(1)
	player.GetBag().GetResources().Set(1, 1001)
	player.GetBag().GetResources().Set(2, 1005)
	hero := kds.NewHero()
	hero.SetHeroId(99)
	hero.SetHeroLevel(1)
	hero.SetType(kds.HeroType_HeroType3)
	hero.SetNeedTime(time.Second * 1000)
	player.GetHero().GetHeroes().Set(int64(hero.GetHeroId()), hero)
	player.GetInfo().SetName("测试")
	player.GetInfo().SetIsNew(true)
	player.GetInfo().SetCreateTime(time.Now())

	buf, err := player.Marshal(nil)
	if err != nil {
		panic(err)
	}
	player2 := kds.NewPlayer(1)
	err = player2.Unmarshal(buf)
	if err != nil {
		panic(err)
	}
	slog.Info("player1", "bag", player.GetBag())
	slog.Info("player2", "bag", player2.GetBag())

	slog.Info("player1", "info", player.GetInfo(), "createtime", player.GetInfo().GetCreateTime(), "unix", player.GetInfo().GetCreateTime().Unix(), "nanosecond", player.GetInfo().GetCreateTime().Nanosecond())
	slog.Info("player2", "info", player2.GetInfo(), "createtime", player2.GetInfo().GetCreateTime(), "unix", player2.GetInfo().GetCreateTime().Unix(), "nanosecond", player2.GetInfo().GetCreateTime().Nanosecond())

	hero1, _ := player.GetHero().GetHeroes().Get(99)
	hero2, _ := player2.GetHero().GetHeroes().Get(99)
	slog.Info("player1", "hero", hero1)
	slog.Info("player2", "hero", hero2)

	buf1, err := player.Marshal(nil)
	if err != nil {
		panic(err)
	}
	buf2, err := player2.Marshal(nil)
	if err != nil {
		panic(err)
	}
	slog.Info("proto", "pb1", len(buf1), "pb2", len(buf2), "equal", bytes.Equal(buf1, buf2), "buf", buf)

	bufDirty, err := player.MarshalDirty(nil)
	if err != nil {
		panic(err)
	}
	player2 = kds.NewPlayer(1)
	err = player2.Unmarshal(bufDirty)
	if err != nil {
		panic(err)
	}
	slog.Info("player1", "bag", player.GetBag())
	slog.Info("player2", "bag", player2.GetBag())

	slog.Info("player1", "info", player.GetInfo(), "createtime", player.GetInfo().GetCreateTime(), "unix", player.GetInfo().GetCreateTime().Unix(), "nanosecond", player.GetInfo().GetCreateTime().Nanosecond())
	slog.Info("player2", "info", player2.GetInfo(), "createtime", player2.GetInfo().GetCreateTime(), "unix", player2.GetInfo().GetCreateTime().Unix(), "nanosecond", player2.GetInfo().GetCreateTime().Nanosecond())

	hero1, _ = player.GetHero().GetHeroes().Get(99)
	hero2, _ = player2.GetHero().GetHeroes().Get(99)
	slog.Info("player1", "hero", hero1)
	slog.Info("player2", "hero", hero2)

	buf1, err = player.Marshal(nil)
	if err != nil {
		panic(err)
	}
	buf2, err = player2.Marshal(nil)
	if err != nil {
		panic(err)
	}

	slog.Info("proto", "pb1", len(buf1), "pb2", len(buf2), "equal", bytes.Equal(buf1, buf2), "buf", bufDirty)
	player.ClearDirty()
	player.GetInfo().SetName("IF")
	bufDirty, err = player.MarshalDirty(nil)
	if err != nil {
		panic(err)
	}
	err = player2.Unmarshal(bufDirty)
	if err != nil {
		panic(err)
	}

	slog.Info("player1", "bag", player.GetBag())
	slog.Info("player2", "bag", player2.GetBag())

	slog.Info("player1", "info", player.GetInfo(), "createtime", player.GetInfo().GetCreateTime(), "unix", player.GetInfo().GetCreateTime().Unix(), "nanosecond", player.GetInfo().GetCreateTime().Nanosecond())
	slog.Info("player2", "info", player2.GetInfo(), "createtime", player2.GetInfo().GetCreateTime(), "unix", player2.GetInfo().GetCreateTime().Unix(), "nanosecond", player2.GetInfo().GetCreateTime().Nanosecond())

	hero1, _ = player.GetHero().GetHeroes().Get(99)
	hero2, _ = player2.GetHero().GetHeroes().Get(99)
	slog.Info("player1", "hero", hero1)
	slog.Info("player2", "hero", hero2)

	buf1, err = player.Marshal(nil)
	if err != nil {
		panic(err)
	}
	buf2, err = player2.Marshal(nil)
	if err != nil {
		panic(err)
	}

	slog.Info("proto", "pb1", len(buf1), "pb2", len(buf2), "equal", bytes.Equal(buf1, buf2), "buf", bufDirty)
}
