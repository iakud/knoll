package example

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/iakud/knoll/kdsync/example/kds"
)

var all = kds.NewAll(0)

func TestExample(t *testing.T) {
	t.Run("Init", testInit)
	t.Run("TypesUpdate", testTypesUpdate)
	t.Run("ListsAdd", testListsAdd)
	t.Run("ListsUpdate", testListsUpdate)
	t.Run("ListsDelete", testListsDelete)
	t.Run("MapsAdd", testMapsAdd)
	t.Run("MapsUpdate", testMapsUpdate)
	t.Run("MapsDelete", testMapsDelete)
}

func sync(t *testing.T) {
	fullData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	all.ClearDirty()
	mergeFrom(fullData)
	// check
	checkKds(t)
}

func syncUpdate(t *testing.T) {
	dirtyData, err := all.MarshalChange(nil)
	if err != nil {
		panic(err)
	}
	all.ClearDirty()
	mergeFrom(dirtyData)
	// check
	checkKds(t)
}

func checkKds(t *testing.T) {
	csJson := toString()
	goJson := all.String()
	var csData, goData any
	if err := json.Unmarshal([]byte(csJson), &csData); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(goJson), &goData); err != nil {
		t.Fatal(err)
	}
	compareValue(t, goData, csData, "")
}

func compareValue(t *testing.T, goVal, csVal any, path string) {
	switch cs := csVal.(type) {
	case map[string]any:
		goMap, ok := goVal.(map[string]any)
		if !ok {
			t.Fatalf("%s type mismatch: cs is map, go is %T", path, goVal)
		}
		compareMap(t, goMap, cs, path)
	case []any:
		goList, ok := goVal.([]any)
		if !ok {
			t.Fatalf("%s type mismatch: cs is list, go is %T", path, goVal)
		}
		compareList(t, goList, cs, path)
	default:
		if fmt.Sprintf("%v", goVal) != fmt.Sprintf("%v", csVal) {
			t.Fatalf("%s mismatch: cs=%v, go=%v", path, csVal, goVal)
		}
	}
}

func compareMap(t *testing.T, goMap, csMap map[string]any, path string) {
	for key, csVal := range csMap {
		currentPath := path + "." + key
		goVal, ok := goMap[key]
		if !ok {
			t.Fatalf("%s missing in Go data", currentPath)
		}
		compareValue(t, goVal, csVal, currentPath)
	}
}

func compareList(t *testing.T, goList, csList []any, path string) {
	if len(csList) != len(goList) {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(csList), len(goList))
	}
	for i := 0; i < len(csList); i++ {
		compareValue(t, goList[i], csList[i], fmt.Sprintf("%s[%d]", path, i))
	}
}

func testInit(t *testing.T) {
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
	lists.GetUint32List().Append(10, 20, 30)
	lists.GetInt64List().Append(100, 200, 300)
	lists.GetUint64List().Append(1000, 2000, 3000)
	lists.GetFloatList().Append(1.1, 2.2, 3.3)
	lists.GetDoubleList().Append(1.11, 2.22, 3.33)
	lists.GetBoolList().Append(true, false)
	lists.GetStringList().Append("a", "b", "c")
	lists.GetTimestampList().Append(time.Unix(1000, 0), time.Unix(2000, 0))
	lists.GetDurationList().Append(time.Second, time.Minute)
	lists.GetEmptyList().Append(struct{}{}, struct{}{})
	lists.GetEnumList().Append(kds.ItemType_ItemTypeWeapon, kds.ItemType_ItemTypeArmor)
	itemList := lists.GetItemList()
	item1 := kds.NewItemData()
	item1.SetId(1)
	item1.SetName("potion")
	item1.SetCount(5)
	itemList.Append(item1)

	// Maps
	maps := all.GetMaps()
	maps.GetInt32Int32Map().Set(1, 100)
	maps.GetInt32Int32Map().Set(2, 200)
	maps.GetInt64Int64Map().Set(1, 1000)
	maps.GetInt64Int64Map().Set(2, 2000)
	maps.GetUint32Uint32Map().Set(3, 300)
	maps.GetUint64Uint64Map().Set(4, 4000)
	maps.GetBoolFloatMap().Set(true, 3.14)
	maps.GetStringDoubleMap().Set("key1", 2.718)
	maps.GetInt32BoolMap().Set(1, true)
	maps.GetInt32BoolMap().Set(2, false)
	maps.GetInt64StringMap().Set(1, "one")
	maps.GetInt64StringMap().Set(2, "two")
	maps.GetUint32BytesMap().Set(1, []byte("bytes1"))
	maps.GetUint32BytesMap().Set(2, []byte("bytes2"))
	maps.GetUint64TimestampMap().Set(1, time.Unix(1000, 0))
	maps.GetUint64TimestampMap().Set(2, time.Unix(2000, 0))
	maps.GetBoolDurationMap().Set(true, time.Second)
	maps.GetBoolDurationMap().Set(false, time.Minute)
	maps.GetStringEmptyMap().Set("key1", struct{}{})
	maps.GetStringEmptyMap().Set("key2", struct{}{})
	maps.GetInt32ItemTypeMap().Set(1, kds.ItemType_ItemTypeWeapon)
	maps.GetInt32ItemTypeMap().Set(2, kds.ItemType_ItemTypeArmor)
	item2 := kds.NewItemData()
	item2.SetId(2)
	item2.SetName("sword")
	item2.SetCount(10)
	maps.GetInt64ItemDataMap().Set(1, item2)
	// sync
	sync(t)
}

func testTypesUpdate(t *testing.T) {
	// Types 修改
	types := all.GetTypes()
	types.SetInt32Val(33)
	syncUpdate(t)

	types.SetInt64Val(65)
	syncUpdate(t)

	types.SetUint32Val(33)
	syncUpdate(t)

	types.SetUint64Val(65)
	syncUpdate(t)

	types.SetSint32Val(-33)
	syncUpdate(t)

	types.SetSint64Val(-65)
	syncUpdate(t)

	types.SetFixed32Val(33)
	syncUpdate(t)

	types.SetFixed64Val(65)
	syncUpdate(t)

	types.SetSfixed32Val(-33)
	syncUpdate(t)

	types.SetSfixed64Val(-65)
	syncUpdate(t)

	types.SetFloatVal(3.15)
	syncUpdate(t)

	types.SetDoubleVal(3.1415)
	syncUpdate(t)

	types.SetBoolVal(false)
	syncUpdate(t)

	types.SetStringVal("world")
	syncUpdate(t)

	types.SetBytesVal([]byte("hello"))
	syncUpdate(t)

	types.SetTimestampVal(time.Unix(9876543210, 0))
	syncUpdate(t)

	types.SetDurationVal(time.Minute)
	syncUpdate(t)

	types.SetEnumVal(kds.ItemType_ItemTypeArmor)
	syncUpdate(t)

	types.GetItemData().SetId(2)
	types.GetItemData().SetName("shield")
	types.GetItemData().SetCount(20)
	syncUpdate(t)
}

func testListsAdd(t *testing.T) {
	// Lists 新增
	lists := all.GetLists()
	lists.GetInt32List().Append(4)
	syncUpdate(t)

	lists.GetInt64List().Append(400)
	syncUpdate(t)

	lists.GetFloatList().Append(4.4)
	syncUpdate(t)

	lists.GetDoubleList().Append(4.44)
	syncUpdate(t)

	lists.GetBoolList().Append(true)
	syncUpdate(t)

	lists.GetStringList().Append("d")
	syncUpdate(t)

	lists.GetTimestampList().Append(time.Unix(3000, 0))
	syncUpdate(t)

	lists.GetDurationList().Append(time.Hour)
	syncUpdate(t)

	lists.GetEnumList().Append(kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	itemNew := kds.NewItemData()
	itemNew.SetId(6)
	itemNew.SetName("ring")
	itemNew.SetCount(3)
	lists.GetItemList().Append(itemNew)
	syncUpdate(t)
}

func testListsUpdate(t *testing.T) {
	// Lists 修改
	lists := all.GetLists()
	lists.GetInt32List().Set(0, 10)
	syncUpdate(t)

	lists.GetInt64List().Set(0, 1000)
	syncUpdate(t)

	lists.GetFloatList().Set(0, 10.5)
	syncUpdate(t)

	lists.GetDoubleList().Set(0, 10.55)
	syncUpdate(t)

	lists.GetBoolList().Set(0, false)
	syncUpdate(t)

	lists.GetStringList().Set(0, "modified")
	syncUpdate(t)

	lists.GetTimestampList().Set(0, time.Unix(5000, 0))
	syncUpdate(t)

	lists.GetDurationList().Set(0, time.Hour*2)
	syncUpdate(t)

	lists.GetEnumList().Set(0, kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	lists.GetItemList().Get(0).SetId(100)
	lists.GetItemList().Get(0).SetName("modified_item")
	lists.GetItemList().Get(0).SetCount(50)
	syncUpdate(t)
}

func testListsDelete(t *testing.T) {
	// Lists 删除
	lists := all.GetLists()
	lists.GetInt32List().Delete(0, 1)
	syncUpdate(t)

	lists.GetInt64List().Delete(0, 1)
	syncUpdate(t)

	lists.GetFloatList().Delete(0, 1)
	syncUpdate(t)

	lists.GetDoubleList().Delete(0, 1)
	syncUpdate(t)

	lists.GetBoolList().Delete(0, 1)
	syncUpdate(t)

	lists.GetStringList().Delete(0, 1)
	syncUpdate(t)

	lists.GetTimestampList().Delete(0, 1)
	syncUpdate(t)

	lists.GetDurationList().Delete(0, 1)
	syncUpdate(t)

	lists.GetEnumList().Delete(0, 1)
	syncUpdate(t)

	lists.GetItemList().Delete(0, 1)
	syncUpdate(t)
}

func testMapsAdd(t *testing.T) {
	// Maps 新增
	maps := all.GetMaps()

	maps.GetInt32Int32Map().Set(11, 1100)
	syncUpdate(t)

	maps.GetInt64Int64Map().Set(12, 12000)
	syncUpdate(t)

	maps.GetUint32Uint32Map().Set(13, 1300)
	syncUpdate(t)

	maps.GetUint64Uint64Map().Set(14, 14000)
	syncUpdate(t)

	maps.GetBoolFloatMap().Set(false, 31.4)
	syncUpdate(t)

	maps.GetStringDoubleMap().Set("key2", 27.18)
	syncUpdate(t)

	maps.GetInt32BoolMap().Set(31, true)
	syncUpdate(t)

	maps.GetInt64StringMap().Set(32, "hello")
	syncUpdate(t)

	maps.GetUint32BytesMap().Set(33, []byte("bytes"))
	syncUpdate(t)

	maps.GetUint64TimestampMap().Set(34, time.Unix(3000, 0))
	syncUpdate(t)

	maps.GetBoolDurationMap().Set(true, time.Hour)
	syncUpdate(t)

	maps.GetStringEmptyMap().Set("key3", struct{}{})
	syncUpdate(t)

	maps.GetInt32ItemTypeMap().Set(35, kds.ItemType_ItemTypeArmor)
	syncUpdate(t)

	itemMap1 := kds.NewItemData()
	itemMap1.SetId(11)
	itemMap1.SetName("sword")
	itemMap1.SetCount(1)
	maps.GetInt64ItemDataMap().Set(10, itemMap1)
	syncUpdate(t)
}

func testMapsUpdate(t *testing.T) {
	// Maps 修改
	maps := all.GetMaps()

	maps.GetInt32Int32Map().Set(11, 999)
	syncUpdate(t)

	maps.GetInt64Int64Map().Set(12, 9999)
	syncUpdate(t)

	maps.GetUint32Uint32Map().Set(13, 999)
	syncUpdate(t)

	maps.GetUint64Uint64Map().Set(14, 9999)
	syncUpdate(t)

	maps.GetBoolFloatMap().Set(false, 99.9)
	syncUpdate(t)

	maps.GetStringDoubleMap().Set("key2", 99.99)
	syncUpdate(t)

	maps.GetInt32BoolMap().Set(31, false)
	syncUpdate(t)

	maps.GetInt64StringMap().Set(32, "modified")
	syncUpdate(t)

	maps.GetUint32BytesMap().Set(33, []byte("modified_bytes"))
	syncUpdate(t)

	maps.GetUint64TimestampMap().Set(34, time.Unix(9999, 0))
	syncUpdate(t)

	maps.GetBoolDurationMap().Set(true, time.Hour*2)
	syncUpdate(t)

	maps.GetInt32ItemTypeMap().Set(35, kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	itemMap1, _ := maps.GetInt64ItemDataMap().Get(10)
	itemMap1.SetId(100)
	itemMap1.SetName("modified_sword")
	itemMap1.SetCount(99)
	syncUpdate(t)
}

func testMapsDelete(t *testing.T) {
	// 删除 Map 元素
	maps := all.GetMaps()

	maps.GetInt32Int32Map().Delete(11)
	syncUpdate(t)

	maps.GetInt64Int64Map().Delete(12)
	syncUpdate(t)

	maps.GetUint32Uint32Map().Delete(13)
	syncUpdate(t)

	maps.GetUint64Uint64Map().Delete(14)
	syncUpdate(t)

	maps.GetBoolFloatMap().Delete(false)
	syncUpdate(t)

	maps.GetStringDoubleMap().Delete("key2")
	syncUpdate(t)

	maps.GetInt32BoolMap().Delete(31)
	syncUpdate(t)

	maps.GetInt64StringMap().Delete(32)
	syncUpdate(t)

	maps.GetUint32BytesMap().Delete(33)
	syncUpdate(t)

	maps.GetUint64TimestampMap().Delete(34)
	syncUpdate(t)

	maps.GetBoolDurationMap().Delete(true)
	syncUpdate(t)

	maps.GetStringEmptyMap().Delete("key3")
	syncUpdate(t)

	maps.GetInt32ItemTypeMap().Delete(35)
	syncUpdate(t)

	maps.GetInt64ItemDataMap().Delete(10)
	syncUpdate(t)
}
