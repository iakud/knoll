package example

import (
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

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
