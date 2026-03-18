package example

import (
	"fmt"
	"testing"
	"time"

	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

var all = kds.NewAll(0)

func TestExample(t *testing.T) {
	testInit()

	testTypesUpdate()

	testListsAdd()
	testListsUpdate()
	testListsDelete()

	testMapsAdd()
	testMapsUpdate()
	testMapsDelete()
}

func testInit() {
	types := all.GetTypes()
	types.SetInt32Val(32)
	types.SetInt64Val(64)
	types.SetUint32Val(32)
	types.SetUint64Val(64)
	types.SetSint32Val(-32)
	types.SetSint64Val(-64)
	types.SetFixed32Val(32)
	types.SetFixed64Val(64)
	types.SetSfixed32Val(-32)
	types.SetSfixed64Val(-64)
	types.SetFloatVal(3.14)
	types.SetDoubleVal(3.14159)
	types.SetBoolVal(true)
	types.SetStringVal("hello")
	types.SetBytesVal([]byte("bytes"))
	types.SetTimestampVal(time.Unix(1234567890, 0))
	types.SetDurationVal(time.Second * 30)
	types.SetEnumVal(kds.ItemType_ItemTypeWeapon)
	types.GetItemData().SetId(1)
	types.GetItemData().SetName("sword")
	types.GetItemData().SetCount(10)

	// Lists
	lists := all.GetLists()
	lists.GetInt32List().Append(1, 2, 3)
	lists.GetInt64List().Append(100, 200, 300)
	lists.GetFloatList().Append(1.1, 2.2, 3.3)
	lists.GetDoubleList().Append(1.11, 2.22, 3.33)
	lists.GetBoolList().Append(true, false)
	lists.GetStringList().Append("a", "b", "c")
	lists.GetTimestampList().Append(time.Unix(1000, 0), time.Unix(2000, 0))
	lists.GetDurationList().Append(time.Second, time.Minute)
	lists.GetEnumList().Append(kds.ItemType_ItemTypeWeapon, kds.ItemType_ItemTypeArmor)
	itemList := lists.GetItemList()
	item1 := kds.NewItemData()
	item1.SetId(1)
	item1.SetName("potion")
	item1.SetCount(5)
	itemList.Append(item1)

	// Maps
	maps := all.GetMaps()
	maps.GetInt32Int32().Set(1, 100)
	maps.GetInt32Int32().Set(2, 200)
	maps.GetInt32String().Set(1, "one")
	maps.GetInt32String().Set(2, "two")
	maps.GetInt32Timestamp().Set(1, time.Unix(1000, 0))
	maps.GetInt32Duration().Set(1, time.Second)
	maps.GetInt32Enum().Set(1, kds.ItemType_ItemTypeWeapon)
	item2 := kds.NewItemData()
	item2.SetId(2)
	item2.SetName("sword")
	item2.SetCount(10)
	maps.GetInt32ItemData().Set(1, item2)

	maps.GetInt64Int64().Set(1, 1000)
	maps.GetInt64Int64().Set(2, 2000)
	maps.GetInt64String().Set(1, "one")
	maps.GetInt64String().Set(2, "two")
	maps.GetInt64Timestamp().Set(1, time.Unix(1000, 0))
	maps.GetInt64Duration().Set(1, time.Second)
	maps.GetInt64Enum().Set(1, kds.ItemType_ItemTypeWeapon)
	item3 := kds.NewItemData()
	item3.SetId(3)
	item3.SetName("shield")
	item3.SetCount(3)
	maps.GetInt64ItemData().Set(1, item3)

	maps.GetStringInt32().Set("a", 1)
	maps.GetStringInt32().Set("b", 2)
	maps.GetStringString().Set("a", "value_a")
	maps.GetStringString().Set("b", "value_b")
	maps.GetStringTimestamp().Set("a", time.Unix(1000, 0))
	maps.GetStringDuration().Set("a", time.Second)
	maps.GetStringEnum().Set("a", kds.ItemType_ItemTypeWeapon)
	item4 := kds.NewItemData()
	item4.SetId(4)
	item4.SetName("armor")
	item4.SetCount(7)
	maps.GetStringItemData().Set("a", item4)

	maps.GetBoolInt32().Set(true, 1)
	maps.GetBoolInt32().Set(false, 0)
	maps.GetBoolString().Set(true, "true")
	maps.GetBoolString().Set(false, "false")
	maps.GetBoolTimestamp().Set(true, time.Unix(1000, 0))
	maps.GetBoolDuration().Set(true, time.Second)
	maps.GetBoolEnum().Set(true, kds.ItemType_ItemTypeWeapon)
	item5 := kds.NewItemData()
	item5.SetId(5)
	item5.SetName("helm")
	item5.SetCount(2)
	maps.GetBoolItemData().Set(true, item5)
	// sync
	sync()
}

func sync() {
	fullData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	all.ClearDirty()
	ApplySync(fullData)
	// check
	check()
}

func syncUpdate() {
	dirtyData, err := all.MarshalDirty(nil)
	if err != nil {
		panic(err)
	}
	all.ClearDirty()
	ApplySync(dirtyData)
	// check
	check()
}

func testTypesUpdate() {
	// Types 修改
	types := all.GetTypes()
	types.SetInt32Val(33)
	syncUpdate()

	types.SetInt64Val(65)
	syncUpdate()

	types.SetUint32Val(33)
	syncUpdate()

	types.SetUint64Val(65)
	syncUpdate()

	types.SetSint32Val(-33)
	syncUpdate()

	types.SetSint64Val(-65)
	syncUpdate()

	types.SetFixed32Val(33)
	syncUpdate()

	types.SetFixed64Val(65)
	syncUpdate()

	types.SetSfixed32Val(-33)
	syncUpdate()

	types.SetSfixed64Val(-65)
	syncUpdate()

	types.SetFloatVal(3.15)
	syncUpdate()

	types.SetDoubleVal(3.1415)
	syncUpdate()

	types.SetBoolVal(false)
	syncUpdate()

	types.SetStringVal("world")
	syncUpdate()

	types.SetBytesVal([]byte("hello"))
	syncUpdate()

	types.SetTimestampVal(time.Unix(9876543210, 0))
	syncUpdate()

	types.SetDurationVal(time.Minute)
	syncUpdate()

	types.SetEnumVal(kds.ItemType_ItemTypeArmor)
	syncUpdate()

	types.GetItemData().SetId(2)
	types.GetItemData().SetName("shield")
	types.GetItemData().SetCount(20)
	syncUpdate()
}

func testListsAdd() {
	// Lists 新增
	lists := all.GetLists()
	lists.GetInt32List().Append(4)
	syncUpdate()

	lists.GetInt64List().Append(400)
	syncUpdate()

	lists.GetFloatList().Append(4.4)
	syncUpdate()

	lists.GetDoubleList().Append(4.44)
	syncUpdate()

	lists.GetBoolList().Append(true)
	syncUpdate()

	lists.GetStringList().Append("d")
	syncUpdate()

	lists.GetTimestampList().Append(time.Unix(3000, 0))
	syncUpdate()

	lists.GetDurationList().Append(time.Hour)
	syncUpdate()

	lists.GetEnumList().Append(kds.ItemType_ItemTypePotion)
	syncUpdate()

	itemNew := kds.NewItemData()
	itemNew.SetId(6)
	itemNew.SetName("ring")
	itemNew.SetCount(3)
	lists.GetItemList().Append(itemNew)
	syncUpdate()
}

func testListsUpdate() {
	// Lists 修改
	lists := all.GetLists()
	lists.GetInt32List().Set(0, 10)
	syncUpdate()

	lists.GetInt64List().Set(0, 1000)
	syncUpdate()

	lists.GetFloatList().Set(0, 10.5)
	syncUpdate()

	lists.GetDoubleList().Set(0, 10.55)
	syncUpdate()

	lists.GetBoolList().Set(0, false)
	syncUpdate()

	lists.GetStringList().Set(0, "modified")
	syncUpdate()

	lists.GetTimestampList().Set(0, time.Unix(5000, 0))
	syncUpdate()

	lists.GetDurationList().Set(0, time.Hour*2)
	syncUpdate()

	lists.GetEnumList().Set(0, kds.ItemType_ItemTypePotion)
	syncUpdate()

	lists.GetItemList().Get(0).SetId(100)
	lists.GetItemList().Get(0).SetName("modified_item")
	lists.GetItemList().Get(0).SetCount(50)
	syncUpdate()
}

func testListsDelete() {
	// Lists 删除
	lists := all.GetLists()
	lists.GetInt32List().Delete(0, 1)
	syncUpdate()

	lists.GetInt64List().Delete(0, 1)
	syncUpdate()

	lists.GetFloatList().Delete(0, 1)
	syncUpdate()

	lists.GetDoubleList().Delete(0, 1)
	syncUpdate()

	lists.GetBoolList().Delete(0, 1)
	syncUpdate()

	lists.GetStringList().Delete(0, 1)
	syncUpdate()

	lists.GetTimestampList().Delete(0, 1)
	syncUpdate()

	lists.GetDurationList().Delete(0, 1)
	syncUpdate()

	lists.GetEnumList().Delete(0, 1)
	syncUpdate()

	lists.GetItemList().Delete(0, 1)
	syncUpdate()
}

func testMapsAdd() {
	// Maps 新增
	maps := all.GetMaps()
	maps.GetInt32Int32().Set(3, 300)
	syncUpdate()

	maps.GetInt32String().Set(3, "three")
	syncUpdate()

	maps.GetInt32Timestamp().Set(2, time.Unix(2000, 0))
	syncUpdate()

	maps.GetInt32Duration().Set(2, time.Minute)
	syncUpdate()

	maps.GetInt32Enum().Set(2, kds.ItemType_ItemTypeArmor)
	syncUpdate()

	itemMap2 := kds.NewItemData()
	itemMap2.SetId(7)
	itemMap2.SetName("boots")
	itemMap2.SetCount(1)
	maps.GetInt32ItemData().Set(2, itemMap2)
	syncUpdate()

	maps.GetInt64Int64().Set(3, 3000)
	syncUpdate()

	maps.GetInt64String().Set(3, "three")
	syncUpdate()

	maps.GetInt64Timestamp().Set(2, time.Unix(2000, 0))
	syncUpdate()

	maps.GetInt64Duration().Set(2, time.Minute)
	syncUpdate()

	maps.GetInt64Enum().Set(2, kds.ItemType_ItemTypeArmor)
	syncUpdate()

	itemMap3 := kds.NewItemData()
	itemMap3.SetId(8)
	itemMap3.SetName("gloves")
	itemMap3.SetCount(2)
	maps.GetInt64ItemData().Set(2, itemMap3)
	syncUpdate()

	maps.GetStringInt32().Set("c", 3)
	syncUpdate()

	maps.GetStringString().Set("c", "value_c")
	syncUpdate()

	maps.GetStringTimestamp().Set("b", time.Unix(2000, 0))
	syncUpdate()

	maps.GetStringDuration().Set("b", time.Minute)
	syncUpdate()

	maps.GetStringEnum().Set("b", kds.ItemType_ItemTypeArmor)
	syncUpdate()

	itemMap4 := kds.NewItemData()
	itemMap4.SetId(9)
	itemMap4.SetName("belt")
	itemMap4.SetCount(4)
	maps.GetStringItemData().Set("b", itemMap4)
	syncUpdate()

	maps.GetBoolInt32().Set(false, 2)
	syncUpdate()

	maps.GetBoolString().Set(false, "no")
	syncUpdate()

	maps.GetBoolTimestamp().Set(false, time.Unix(2000, 0))
	syncUpdate()

	maps.GetBoolDuration().Set(false, time.Minute)
	syncUpdate()

	maps.GetBoolEnum().Set(false, kds.ItemType_ItemTypeArmor)
	syncUpdate()

	itemMap5 := kds.NewItemData()
	itemMap5.SetId(10)
	itemMap5.SetName("amulet")
	itemMap5.SetCount(1)
	maps.GetBoolItemData().Set(false, itemMap5)
	syncUpdate()
}

func testMapsUpdate() {
	// Maps 修改
	maps := all.GetMaps()
	maps.GetInt32Int32().Set(1, 999)
	syncUpdate()

	maps.GetInt32String().Set(1, "modified_one")
	syncUpdate()

	maps.GetInt32Timestamp().Set(1, time.Unix(8888, 0))
	syncUpdate()

	maps.GetInt32Duration().Set(1, time.Hour)
	syncUpdate()

	maps.GetInt32Enum().Set(1, kds.ItemType_ItemTypePotion)
	syncUpdate()

	itemInt32, _ := maps.GetInt32ItemData().Get(1)
	itemInt32.SetId(111)
	itemInt32.SetName("modified_itemdata")
	itemInt32.SetCount(55)
	syncUpdate()

	maps.GetInt64Int64().Set(1, 9999)
	syncUpdate()

	maps.GetInt64String().Set(1, "modified_one")
	syncUpdate()

	maps.GetInt64Timestamp().Set(1, time.Unix(8888, 0))
	syncUpdate()

	maps.GetInt64Duration().Set(1, time.Hour)
	syncUpdate()

	maps.GetInt64Enum().Set(1, kds.ItemType_ItemTypePotion)
	syncUpdate()

	itemInt64, _ := maps.GetInt64ItemData().Get(1)
	itemInt64.SetId(112)
	itemInt64.SetName("modified_itemdata2")
	itemInt64.SetCount(56)
	syncUpdate()

	maps.GetStringInt32().Set("a", 999)
	syncUpdate()

	maps.GetStringString().Set("a", "modified_value_a")
	syncUpdate()

	maps.GetStringTimestamp().Set("a", time.Unix(9999, 0))
	syncUpdate()

	maps.GetStringDuration().Set("a", time.Hour*3)
	syncUpdate()

	maps.GetStringEnum().Set("a", kds.ItemType_ItemTypePotion)
	syncUpdate()

	itemString, _ := maps.GetStringItemData().Get("a")
	itemString.SetId(113)
	itemString.SetName("modified_itemdata3")
	itemString.SetCount(57)
	syncUpdate()

	maps.GetBoolInt32().Set(true, 888)
	syncUpdate()

	maps.GetBoolString().Set(true, "modified_true")
	syncUpdate()

	maps.GetBoolTimestamp().Set(true, time.Unix(7777, 0))
	syncUpdate()

	maps.GetBoolDuration().Set(true, time.Hour*4)
	syncUpdate()

	maps.GetBoolEnum().Set(true, kds.ItemType_ItemTypePotion)
	syncUpdate()

	itemBool, _ := maps.GetBoolItemData().Get(true)
	itemBool.SetId(114)
	itemBool.SetName("modified_itemdata4")
	itemBool.SetCount(58)
	syncUpdate()
}

func testMapsDelete() {
	// 删除 Map 元素
	maps := all.GetMaps()
	maps.GetInt32Int32().Delete(1)
	syncUpdate()

	maps.GetInt32String().Delete(1)
	syncUpdate()

	maps.GetInt32Timestamp().Delete(1)
	syncUpdate()

	maps.GetInt32Duration().Delete(1)
	syncUpdate()

	maps.GetInt32Enum().Delete(1)
	syncUpdate()

	maps.GetInt32ItemData().Delete(1)
	syncUpdate()

	maps.GetInt64Int64().Delete(1)
	syncUpdate()

	maps.GetInt64String().Delete(1)
	syncUpdate()

	maps.GetInt64Timestamp().Delete(1)
	syncUpdate()

	maps.GetInt64Duration().Delete(1)
	syncUpdate()

	maps.GetInt64Enum().Delete(1)
	syncUpdate()

	maps.GetInt64ItemData().Delete(1)
	syncUpdate()

	maps.GetStringInt32().Delete("a")
	syncUpdate()

	maps.GetStringString().Delete("a")
	syncUpdate()

	maps.GetStringTimestamp().Delete("a")
	syncUpdate()

	maps.GetStringDuration().Delete("a")
	syncUpdate()

	maps.GetStringEnum().Delete("a")
	syncUpdate()

	maps.GetStringItemData().Delete("a")
	syncUpdate()

	maps.GetBoolInt32().Delete(true)
	syncUpdate()

	maps.GetBoolString().Delete(true)
	syncUpdate()

	maps.GetBoolTimestamp().Delete(true)
	syncUpdate()

	maps.GetBoolDuration().Delete(true)
	syncUpdate()

	maps.GetBoolEnum().Delete(true)
	syncUpdate()

	maps.GetBoolItemData().Delete(true)
	syncUpdate()
}

func check() {
	goDump := all.String("")
	csharpDump := ToString()
	if goDump != csharpDump {
		fmt.Printf("=== Go ===\n%s\n", goDump)
		fmt.Printf("=== C# ===\n%s\n", csharpDump)
		// Find first difference
		for i := 0; i < len(goDump) && i < len(csharpDump); i++ {
			if goDump[i] != csharpDump[i] {
				fmt.Printf("First diff at char %d: go='%c' (0x%x), cs='%c' (0x%x), ", i, goDump[i], goDump[i], csharpDump[i], csharpDump[i])
				break
			}
		}
		panic("Dump mismatch!")
	}
	fmt.Println("Dump match!")
}
