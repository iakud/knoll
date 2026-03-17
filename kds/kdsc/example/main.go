package main

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/example.dylib
#include <stdlib.h>

extern int32_t apply_sync(const char* data, int32_t length);
extern char* dump(void);
*/
import "C"

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

var all = kds.NewAll(0)

func main() {
	// FIXME: init datas
	initAll()

	syncAdd()

	// FIXME: change datas
	syncUpdate()
}

func initAll() {
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
}

func syncAdd() {
	fullData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(fullData)), C.int32_t(len(fullData)))
	check()
}

func syncUpdate() {
	dirtyData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(dirtyData)), C.int32_t(len(dirtyData)))
	check()
}

func dump() string {
	types := all.GetTypes()
	lists := all.GetLists()
	maps := all.GetMaps()
	var sb strings.Builder

	// Types
	fmt.Fprintf(&sb, "Int32Val=%d", types.GetInt32Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64Val=%d", types.GetInt64Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Uint32Val=%d", types.GetUint32Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Uint64Val=%d", types.GetUint64Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Sint32Val=%d", types.GetSint32Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Sint64Val=%d", types.GetSint64Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Fixed32Val=%d", types.GetFixed32Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Fixed64Val=%d", types.GetFixed64Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Sfixed32Val=%d", types.GetSfixed32Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Sfixed64Val=%d", types.GetSfixed64Val())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "FloatVal=%f", types.GetFloatVal())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "DoubleVal=%f", types.GetDoubleVal())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolVal=%v", types.GetBoolVal())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringVal=%s", types.GetStringVal())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BytesVal=%s", base64.RawStdEncoding.EncodeToString(types.GetBytesVal()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "TimestampVal=%v", types.GetTimestampVal().UnixNano())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "DurationVal=%v", types.GetDurationVal().Nanoseconds())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "EnumVal=%d", types.GetEnumVal())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "ItemData=(%d,%s,%d)", types.GetItemData().GetId(), types.GetItemData().GetName(), types.GetItemData().GetCount())
	sb.WriteString(", ")

	// Lists
	fmt.Fprintf(&sb, "Int32List=%v", dumpInt32List(lists.GetInt32List()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64List=%v", dumpInt64List(lists.GetInt64List()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "FloatList=%v", dumpFloatList(lists.GetFloatList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "DoubleList=%v", dumpDoubleList(lists.GetDoubleList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolList=%v", dumpBoolList(lists.GetBoolList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringList=%v", dumpStringList(lists.GetStringList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "TimestampList=%v", dumpTimestampList(lists.GetTimestampList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "DurationList=%v", dumpDurationList(lists.GetDurationList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "EmptyList=%v", lists.GetEmptyList().Len())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "EnumList=%v", dumpEnumList(lists.GetEnumList()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "ItemList=%v", dumpItemDataList(lists.GetItemList()))
	sb.WriteString(", ")

	// Maps
	fmt.Fprintf(&sb, "Int32Int32=%v", dumpInt32Int32Map(maps.GetInt32Int32()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int32String=%v", dumpInt32StringMap(maps.GetInt32String()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int32Timestamp=%v", dumpInt32TimestampMap(maps.GetInt32Timestamp()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int32Duration=%v", dumpInt32DurationMap(maps.GetInt32Duration()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int32Empty=%v", maps.GetInt32Empty().Len())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int32Enum=%v", dumpInt32EnumMap(maps.GetInt32Enum()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int32ItemData=%v", dumpInt32ItemDataMap(maps.GetInt32ItemData()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64Int64=%v", dumpInt64Int64Map(maps.GetInt64Int64()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64String=%v", dumpInt64StringMap(maps.GetInt64String()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64Timestamp=%v", dumpInt64TimestampMap(maps.GetInt64Timestamp()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64Duration=%v", dumpInt64DurationMap(maps.GetInt64Duration()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64Empty=%v", maps.GetInt64Empty().Len())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64Enum=%v", dumpInt64EnumMap(maps.GetInt64Enum()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "Int64ItemData=%v", dumpInt64ItemDataMap(maps.GetInt64ItemData()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringInt32=%v", dumpStringInt32Map(maps.GetStringInt32()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringString=%v", dumpStringStringMap(maps.GetStringString()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringTimestamp=%v", dumpStringTimestampMap(maps.GetStringTimestamp()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringDuration=%v", dumpStringDurationMap(maps.GetStringDuration()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringEmpty=%v", maps.GetStringEmpty().Len())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringEnum=%v", dumpStringEnumMap(maps.GetStringEnum()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "StringItemData=%v", dumpStringItemDataMap(maps.GetStringItemData()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolInt32=%v", dumpBoolInt32Map(maps.GetBoolInt32()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolString=%v", dumpBoolStringMap(maps.GetBoolString()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolTimestamp=%v", dumpBoolTimestampMap(maps.GetBoolTimestamp()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolDuration=%v", dumpBoolDurationMap(maps.GetBoolDuration()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolEmpty=%v", maps.GetBoolEmpty().Len())
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolEnum=%v", dumpBoolEnumMap(maps.GetBoolEnum()))
	sb.WriteString(", ")
	fmt.Fprintf(&sb, "BoolItemData=%v", dumpBoolItemDataMap(maps.GetBoolItemData()))

	return sb.String()
}

func dumpInt32List(l *kds.Int32_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64List(l *kds.Int64_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpFloatList(l *kds.Float_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%f", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpDoubleList(l *kds.Double_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%f", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolList(l *kds.Bool_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpTimestampList(l *kds.Timestamp_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d", l.Get(i).UnixNano())
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpDurationList(l *kds.Duration_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d", l.Get(i).Nanoseconds())
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpEnumList(l *kds.ItemType_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpItemDataList(l *kds.ItemData_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		item := l.Get(i)
		fmt.Fprintf(&sb, "(%d,%s,%d)", item.GetId(), item.GetName(), item.GetCount())
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringList(l *kds.String_list) string {
	if l.Len() == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Len(); i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s", l.Get(i))
	}
	sb.WriteString("]")
	return sb.String()
}

// Map dump functions
func dumpInt32Int32Map(m *kds.Int32Int32_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt32StringMap(m *kds.Int32String_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%s", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt32TimestampMap(m *kds.Int32Timestamp_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v.UnixNano())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt32DurationMap(m *kds.Int32Duration_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v.Nanoseconds())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt32EnumMap(m *kds.Int32ItemType_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt32ItemDataMap(m *kds.Int32ItemData_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:(%d,%s,%d)", k, v.GetId(), v.GetName(), v.GetCount())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64Int64Map(m *kds.Int64Int64_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64StringMap(m *kds.Int64String_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%s", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64TimestampMap(m *kds.Int64Timestamp_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v.UnixNano())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64DurationMap(m *kds.Int64Duration_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v.Nanoseconds())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64EnumMap(m *kds.Int64ItemType_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpInt64ItemDataMap(m *kds.Int64ItemData_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%d:(%d,%s,%d)", k, v.GetId(), v.GetName(), v.GetCount())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringInt32Map(m *kds.StringInt32_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringStringMap(m *kds.StringString_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s:%s", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringTimestampMap(m *kds.StringTimestamp_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s:%d", k, v.UnixNano())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringDurationMap(m *kds.StringDuration_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s:%d", k, v.Nanoseconds())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringEnumMap(m *kds.StringItemType_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpStringItemDataMap(m *kds.StringItemData_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%s:(%d,%s,%d)", k, v.GetId(), v.GetName(), v.GetCount())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolInt32Map(m *kds.BoolInt32_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolStringMap(m *kds.BoolString_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v:%s", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolTimestampMap(m *kds.BoolTimestamp_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v:%d", k, v.UnixNano())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolDurationMap(m *kds.BoolDuration_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v:%d", k, v.Nanoseconds())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolEnumMap(m *kds.BoolItemType_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v:%d", k, v)
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func dumpBoolItemDataMap(m *kds.BoolItemData_map) string {
	if m.Len() == 0 {
		return "map[]"
	}
	var sb strings.Builder
	sb.WriteString("map[")
	first := true
	for k, v := range m.All() {
		if !first {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v:(%d,%s,%d)", k, v.GetId(), v.GetName(), v.GetCount())
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}

func check() {
	goDump := dump()
	csharpDump := C.GoString(C.dump())
	if goDump != csharpDump {
		fmt.Printf("=== Go Dump ===\n%s, ", goDump)
		fmt.Printf("=== C# Dump ===\n%s, ", csharpDump)
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
