package main

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/client.dylib
#include <stdlib.h>

extern int32_t client_init(int64_t player_id);
extern int32_t client_apply_sync(const char* data, int32_t length);
extern int32_t client_get_info(void);
extern char* client_get_output(void);
extern void client_free_output(void* ptr);
extern void client_clear_output(void);
*/
import "C"

import (
	"fmt"
	"log/slog"

	"github.com/iakud/knoll/kds/kdsc/example/server/kds"
)

func main() {
	// ========== 生成同步数据 ==========
	slog.Info("=== 生成同步数据 ===")

	serverPlayer := kds.NewPlayerSync(1001)
	serverPlayer.GetInfo().SetName("Player1")
	serverPlayer.GetInfo().SetLevel(10)
	serverPlayer.GetInfo().SetIsNew(false)

	serverPlayer.GetBag().GetCurrencies().Set(1, 1000)
	serverPlayer.GetBag().GetCurrencies().Set(2, 100)
	serverPlayer.GetBag().GetCurrencies().Set(3, 50)

	serverPlayer.GetBag().GetItems().Set(1001, 10)
	serverPlayer.GetBag().GetItems().Set(1002, 5)

	hero1 := kds.NewSyncHero()
	hero1.SetHeroId(1)
	hero1.SetLevel(5)
	hero1.SetStar(3)
	hero1.SetExp(1000)
	serverPlayer.GetHero().GetHeroes().Set(1, hero1)

	hero2 := kds.NewSyncHero()
	hero2.SetHeroId(2)
	hero2.SetLevel(3)
	hero2.SetStar(2)
	hero2.SetExp(500)
	serverPlayer.GetHero().GetHeroes().Set(2, hero2)

	fullData, err := serverPlayer.Marshal(nil)
	if err != nil {
		panic(err)
	}

	C.client_init(1001)
	slog.Info(C.GoString(C.client_get_output()))
	C.client_clear_output()

	C.client_apply_sync(C.CString(string(fullData)), C.int32_t(len(fullData)))
	slog.Info(C.GoString(C.client_get_output()))
	C.client_clear_output()

	// 获取信息
	C.client_get_info()
	slog.Info(C.GoString(C.client_get_output()))
	C.client_clear_output()

	serverPlayer.ClearDirty()

	serverPlayer.GetInfo().SetLevel(11)
	serverPlayer.GetBag().GetCurrencies().Set(1, 1500)
	serverPlayer.GetBag().GetCurrencies().Set(4, 10)
	serverPlayer.GetBag().GetItems().Delete(1002)

	hero1.SetLevel(6)
	hero1.SetExp(1500)

	hero3 := kds.NewSyncHero()
	hero3.SetHeroId(3)
	hero3.SetLevel(1)
	hero3.SetStar(1)
	hero3.SetExp(0)
	serverPlayer.GetHero().GetHeroes().Set(3, hero3)

	dirtyData, err := serverPlayer.MarshalDirty(nil)
	if err != nil {
		panic(err)
	}

	C.client_apply_sync(C.CString(string(dirtyData)), C.int32_t(len(dirtyData)))
	slog.Info(C.GoString(C.client_get_output()))
	C.client_clear_output()

	// 获取信息
	C.client_get_info()
	fmt.Print(C.GoString(C.client_get_output()))
}
