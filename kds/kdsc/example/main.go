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

	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

var all = kds.NewAll(0)

func main() {
	// FIXME: init datas
	fullData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(fullData)), C.int32_t(len(fullData)))
	Check()
	// FIXME: change datas
	all.GetTypes().SetInt32Val(10)

	dirtyData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(dirtyData)), C.int32_t(len(dirtyData)))
	Check()
}

func Dump() string {
	types := all.GetTypes()
	lists := all.GetLists()
	maps := all.GetMaps()
	var sb strings.Builder

	// Types
	fmt.Fprintf(&sb, "Int32Val=%d\n", types.GetInt32Val())
	fmt.Fprintf(&sb, "Int64Val=%d\n", types.GetInt64Val())
	fmt.Fprintf(&sb, "Uint32Val=%d\n", types.GetUint32Val())
	fmt.Fprintf(&sb, "Uint64Val=%d\n", types.GetUint64Val())
	fmt.Fprintf(&sb, "Sint32Val=%d\n", types.GetSint32Val())
	fmt.Fprintf(&sb, "Sint64Val=%d\n", types.GetSint64Val())
	fmt.Fprintf(&sb, "Fixed32Val=%d\n", types.GetFixed32Val())
	fmt.Fprintf(&sb, "Fixed64Val=%d\n", types.GetFixed64Val())
	fmt.Fprintf(&sb, "Sfixed32Val=%d\n", types.GetSfixed32Val())
	fmt.Fprintf(&sb, "Sfixed64Val=%d\n", types.GetSfixed64Val())
	fmt.Fprintf(&sb, "FloatVal=%f\n", types.GetFloatVal())
	fmt.Fprintf(&sb, "DoubleVal=%f\n", types.GetDoubleVal())
	fmt.Fprintf(&sb, "BoolVal=%v\n", types.GetBoolVal())
	fmt.Fprintf(&sb, "StringVal=%s\n", types.GetStringVal())
	fmt.Fprintf(&sb, "BytesVal=%s\n", base64.RawStdEncoding.EncodeToString(types.GetBytesVal()))
	fmt.Fprintf(&sb, "TimestampVal=%d\n", types.GetTimestampVal().UnixNano())
	fmt.Fprintf(&sb, "DurationVal=%d\n", types.GetDurationVal().Nanoseconds())
	fmt.Fprintf(&sb, "EnumVal=%d\n", types.GetEnumVal())
	fmt.Fprintf(&sb, "ItemData=(%d,%s,%d)\n", types.GetItemData().GetId(), types.GetItemData().GetName(), types.GetItemData().GetCount())

	// Lists
	fmt.Fprintf(&sb, "Int32List=%v\n", dumpInt32List(lists.GetInt32List()))
	fmt.Fprintf(&sb, "Int64List=%v\n", dumpInt64List(lists.GetInt64List()))
	fmt.Fprintf(&sb, "FloatList=%v\n", dumpFloatList(lists.GetFloatList()))
	fmt.Fprintf(&sb, "DoubleList=%v\n", dumpDoubleList(lists.GetDoubleList()))
	fmt.Fprintf(&sb, "BoolList=%v\n", dumpBoolList(lists.GetBoolList()))
	fmt.Fprintf(&sb, "StringList=%v\n", dumpStringList(lists.GetStringList()))
	fmt.Fprintf(&sb, "TimestampList=%v\n", dumpTimestampList(lists.GetTimestampList()))
	fmt.Fprintf(&sb, "DurationList=%v\n", dumpDurationList(lists.GetDurationList()))
	fmt.Fprintf(&sb, "EmptyList=%v\n", lists.GetEmptyList().Len())
	fmt.Fprintf(&sb, "EnumList=%v\n", dumpEnumList(lists.GetEnumList()))
	fmt.Fprintf(&sb, "ItemList=%v\n", dumpItemDataList(lists.GetItemList()))

	// Maps
	fmt.Fprintf(&sb, "Int32Int32=%v\n", dumpInt32Int32Map(maps.GetInt32Int32()))
	fmt.Fprintf(&sb, "Int32String=%v\n", dumpInt32StringMap(maps.GetInt32String()))
	fmt.Fprintf(&sb, "Int32Timestamp=%v\n", dumpInt32TimestampMap(maps.GetInt32Timestamp()))
	fmt.Fprintf(&sb, "Int32Duration=%v\n", dumpInt32DurationMap(maps.GetInt32Duration()))
	fmt.Fprintf(&sb, "Int32Empty=%v\n", maps.GetInt32Empty().Len())
	fmt.Fprintf(&sb, "Int32Enum=%v\n", dumpInt32EnumMap(maps.GetInt32Enum()))
	fmt.Fprintf(&sb, "Int32ItemData=%v\n", dumpInt32ItemDataMap(maps.GetInt32ItemData()))
	fmt.Fprintf(&sb, "Int64Int64=%v\n", dumpInt64Int64Map(maps.GetInt64Int64()))
	fmt.Fprintf(&sb, "Int64String=%v\n", dumpInt64StringMap(maps.GetInt64String()))
	fmt.Fprintf(&sb, "Int64Timestamp=%v\n", dumpInt64TimestampMap(maps.GetInt64Timestamp()))
	fmt.Fprintf(&sb, "Int64Duration=%v\n", dumpInt64DurationMap(maps.GetInt64Duration()))
	fmt.Fprintf(&sb, "Int64Empty=%v\n", maps.GetInt64Empty().Len())
	fmt.Fprintf(&sb, "Int64Enum=%v\n", dumpInt64EnumMap(maps.GetInt64Enum()))
	fmt.Fprintf(&sb, "Int64ItemData=%v\n", dumpInt64ItemDataMap(maps.GetInt64ItemData()))
	fmt.Fprintf(&sb, "StringInt32=%v\n", dumpStringInt32Map(maps.GetStringInt32()))
	fmt.Fprintf(&sb, "StringString=%v\n", dumpStringStringMap(maps.GetStringString()))
	fmt.Fprintf(&sb, "StringTimestamp=%v\n", dumpStringTimestampMap(maps.GetStringTimestamp()))
	fmt.Fprintf(&sb, "StringDuration=%v\n", dumpStringDurationMap(maps.GetStringDuration()))
	fmt.Fprintf(&sb, "StringEmpty=%v\n", maps.GetStringEmpty().Len())
	fmt.Fprintf(&sb, "StringEnum=%v\n", dumpStringEnumMap(maps.GetStringEnum()))
	fmt.Fprintf(&sb, "StringItemData=%v\n", dumpStringItemDataMap(maps.GetStringItemData()))
	fmt.Fprintf(&sb, "BoolInt32=%v\n", dumpBoolInt32Map(maps.GetBoolInt32()))
	fmt.Fprintf(&sb, "BoolString=%v\n", dumpBoolStringMap(maps.GetBoolString()))
	fmt.Fprintf(&sb, "BoolTimestamp=%v\n", dumpBoolTimestampMap(maps.GetBoolTimestamp()))
	fmt.Fprintf(&sb, "BoolDuration=%v\n", dumpBoolDurationMap(maps.GetBoolDuration()))
	fmt.Fprintf(&sb, "BoolEmpty=%v\n", maps.GetBoolEmpty().Len())
	fmt.Fprintf(&sb, "BoolEnum=%v\n", dumpBoolEnumMap(maps.GetBoolEnum()))
	fmt.Fprintf(&sb, "BoolItemData=%v\n", dumpBoolItemDataMap(maps.GetBoolItemData()))

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

func Check() {
	goDump := Dump()
	csharpDump := C.GoString(C.dump())
	if goDump != csharpDump {
		fmt.Printf("=== Go Dump ===\n%s\n", goDump)
		fmt.Printf("=== C# Dump ===\n%s\n", csharpDump)
		// Find first difference
		for i := 0; i < len(goDump) && i < len(csharpDump); i++ {
			if goDump[i] != csharpDump[i] {
				fmt.Printf("First diff at char %d: go='%c' (0x%x), cs='%c' (0x%x)\n", i, goDump[i], goDump[i], csharpDump[i], csharpDump[i])
				break
			}
		}
		panic("Dump mismatch!")
	}
	fmt.Println("Dump match!")
}
