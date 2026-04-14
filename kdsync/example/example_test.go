package example

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/iakud/knoll/kdsync"
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
	checkKds2(t)
}

func syncUpdate(t *testing.T) {
	dirtyData, err := all.MarshalChange(nil)
	if err != nil {
		panic(err)
	}
	all.ClearDirty()
	mergeFrom(dirtyData)
	// check
	checkKds2(t)
}

func checkKds(t *testing.T) {
	goKds := all.String()
	csKds := getString()

	goLines := strings.Split(goKds, "\n")
	csLines := strings.Split(csKds, "\n")

	for i := 0; i < len(goLines) || i < len(csLines); i++ {
		goLine := ""
		if i < len(goLines) {
			goLine = goLines[i]
		}
		csLine := ""
		if i < len(csLines) {
			csLine = csLines[i]
		}
		if goLine != csLine {
			t.Logf("=== Go ===\n%s\n", goKds)
			t.Logf("=== C# ===\n%s\n", csKds)
			t.Logf("Line %d differ:\n  Go: %s\n  C#: %s\n", i+1, goLine, csLine)
			t.Fatal("Kds mismatch!")
		}
	}
}

func checkKds2(t *testing.T) {
	csJson := getString()
	var csData map[string]any
	if err := json.Unmarshal([]byte(csJson), &csData); err != nil {
		t.Fatal(err)
	}

	types := csData["Types"].(map[string]any)
	compareTypes(t, types, all.GetTypes(), "Types")

	lists := csData["Lists"].(map[string]any)
	compareLists(t, lists, all.GetLists(), "Lists")

	maps := csData["Maps"].(map[string]any)
	compareMaps2(t, maps, all.GetMaps(), "Maps")
}

func compareTypes(t *testing.T, cs map[string]any, goTypes *kds.AllType, path string) {
	compareValue(t, cs["Int32Val"], goTypes.GetInt32Val(), path+".Int32Val")
	compareInt64Value(t, cs["Int64Val"], goTypes.GetInt64Val(), path+".Int64Val")
	compareValue(t, cs["Uint32Val"], goTypes.GetUint32Val(), path+".Uint32Val")
	compareUint64Value(t, cs["Uint64Val"], goTypes.GetUint64Val(), path+".Uint64Val")
	compareValue(t, cs["Sint32Val"], goTypes.GetSint32Val(), path+".Sint32Val")
	compareInt64Value(t, cs["Sint64Val"], goTypes.GetSint64Val(), path+".Sint64Val")
	compareValue(t, cs["Fixed32Val"], goTypes.GetFixed32Val(), path+".Fixed32Val")
	compareUint64Value(t, cs["Fixed64Val"], goTypes.GetFixed64Val(), path+".Fixed64Val")
	compareValue(t, cs["Sfixed32Val"], goTypes.GetSfixed32Val(), path+".Sfixed32Val")
	compareInt64Value(t, cs["Sfixed64Val"], goTypes.GetSfixed64Val(), path+".Sfixed64Val")
	compareValue(t, cs["FloatVal"], goTypes.GetFloatVal(), path+".FloatVal")
	compareValue(t, cs["DoubleVal"], goTypes.GetDoubleVal(), path+".DoubleVal")
	compareValue(t, cs["BoolVal"], goTypes.GetBoolVal(), path+".BoolVal")
	compareValue(t, cs["StringVal"], goTypes.GetStringVal(), path+".StringVal")
	compareBytesValue(t, cs["BytesVal"], goTypes.GetBytesVal(), path+".BytesVal")
	compareTimestampValue(t, cs["TimestampVal"], goTypes.GetTimestampVal(), path+".TimestampVal")
	compareDurationValue(t, cs["DurationVal"], goTypes.GetDurationVal(), path+".DurationVal")
	compareValue(t, cs["EnumVal"], int(goTypes.GetEnumVal()), path+".EnumVal")

	itemData := cs["ItemData"].(map[string]any)
	goItemData := goTypes.GetItemData()
	compareValue(t, itemData["Id"], goItemData.GetId(), path+".ItemData.Id")
	compareValue(t, itemData["Name"], goItemData.GetName(), path+".ItemData.Name")
	compareValue(t, itemData["Count"], goItemData.GetCount(), path+".ItemData.Count")
}

func compareLists(t *testing.T, cs map[string]any, goLists *kds.AllList, path string) {
	compareRepeatedInt32(t, cs["Int32List"].([]any), goLists.GetInt32List(), path+".Int32List")
	compareRepeatedInt64(t, cs["Int64List"].([]any), goLists.GetInt64List(), path+".Int64List")
	compareRepeatedFloat(t, cs["FloatList"].([]any), goLists.GetFloatList(), path+".FloatList")
	compareRepeatedDouble(t, cs["DoubleList"].([]any), goLists.GetDoubleList(), path+".DoubleList")
	compareRepeatedBool(t, cs["BoolList"].([]any), goLists.GetBoolList(), path+".BoolList")
	compareRepeatedString(t, cs["StringList"].([]any), goLists.GetStringList(), path+".StringList")
	compareRepeatedTimestamp(t, cs["TimestampList"].([]any), goLists.GetTimestampList(), path+".TimestampList")
	compareRepeatedDuration(t, cs["DurationList"].([]any), goLists.GetDurationList(), path+".DurationList")
	compareRepeatedEnum(t, cs["EnumList"].([]any), goLists.GetEnumList(), path+".EnumList")
	// skip EmptyList - nosync

	csItemList := cs["ItemList"].([]any)
	goItemList := goLists.GetItemList()
	if len(csItemList) != goItemList.Len() {
		t.Fatalf("%s Length mismatch: cs=%d, go=%d", path+".ItemList", len(csItemList), goItemList.Len())
	}
	for i := 0; i < len(csItemList); i++ {
		itemData := csItemList[i].(map[string]any)
		goItemData := goItemList.Get(i)
		compareValue(t, itemData["Id"], goItemData.GetId(), fmt.Sprintf("%s[%d].Id", path+".ItemList", i))
		compareValue(t, itemData["Name"], goItemData.GetName(), fmt.Sprintf("%s[%d].Name", path+".ItemList", i))
		compareValue(t, itemData["Count"], goItemData.GetCount(), fmt.Sprintf("%s[%d].Count", path+".ItemList", i))
	}
}

func compareMaps2(t *testing.T, cs map[string]any, goMaps *kds.AllMap, path string) {
	compareMapInt32Int32(t, cs["Int32Int32Map"].(map[string]any), goMaps.GetInt32Int32Map(), path+".Int32Int32Map")
	compareMapInt64Int64(t, cs["Int64Int64Map"].(map[string]any), goMaps.GetInt64Int64Map(), path+".Int64Int64Map")
	compareMapUint32Uint32(t, cs["Uint32Uint32Map"].(map[string]any), goMaps.GetUint32Uint32Map(), path+".Uint32Uint32Map")
	compareMapUint64Uint64(t, cs["Uint64Uint64Map"].(map[string]any), goMaps.GetUint64Uint64Map(), path+".Uint64Uint64Map")
	compareMapBoolFloat(t, cs["BoolFloatMap"].(map[string]any), goMaps.GetBoolFloatMap(), path+".BoolFloatMap")
	compareMapStringDouble(t, cs["StringDoubleMap"].(map[string]any), goMaps.GetStringDoubleMap(), path+".StringDoubleMap")
	compareMapInt32Bool(t, cs["Int32BoolMap"].(map[string]any), goMaps.GetInt32BoolMap(), path+".Int32BoolMap")
	compareMapInt64String(t, cs["Int64StringMap"].(map[string]any), goMaps.GetInt64StringMap(), path+".Int64StringMap")
	compareMapUint32Bytes(t, cs["Uint32BytesMap"].(map[string]any), goMaps.GetUint32BytesMap(), path+".Uint32BytesMap")
	compareMapUint64Timestamp(t, cs["Uint64TimestampMap"].(map[string]any), goMaps.GetUint64TimestampMap(), path+".Uint64TimestampMap")
	compareMapBoolDuration(t, cs["BoolDurationMap"].(map[string]any), goMaps.GetBoolDurationMap(), path+".BoolDurationMap")
	compareMapStringEmpty(t, cs["StringEmptyMap"].(map[string]any), goMaps.GetStringEmptyMap(), path+".StringEmptyMap")
	compareMapInt32ItemType(t, cs["Int32ItemTypeMap"].(map[string]any), goMaps.GetInt32ItemTypeMap(), path+".Int32ItemTypeMap")
	compareMapInt64ItemData(t, cs["Int64ItemDataMap"].(map[string]any), goMaps.GetInt64ItemDataMap(), path+".Int64ItemDataMap")
}

func compareValue(t *testing.T, cs, goVal any, path string) {
	if fmt.Sprintf("%v", cs) != fmt.Sprintf("%v", goVal) {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, cs, goVal)
	}
}

func compareBytesValue(t *testing.T, cs, goVal any, path string) {
	csStr, ok := cs.(string)
	if !ok {
		t.Fatalf("%s cs is not string: %v", path, cs)
	}
	goBytes, ok := goVal.([]byte)
	if !ok {
		t.Fatalf("%s goVal is not []byte: %v", path, goVal)
	}
	decoded, err := base64.StdEncoding.DecodeString(csStr)
	if err != nil {
		t.Fatalf("%s base64 decode failed: %v", path, err)
	}
	if string(decoded) != string(goBytes) {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, string(decoded), string(goBytes))
	}
}

func compareInt64Value(t *testing.T, cs, goVal any, path string) {
	var csInt64 int64
	if s, ok := cs.(string); ok {
		csInt64 = mustParseInt(s)
	} else if f, ok := cs.(float64); ok {
		csInt64 = int64(f)
	}
	if csInt64 != goVal.(int64) {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, csInt64, goVal)
	}
}

func compareUint64Value(t *testing.T, cs, goVal any, path string) {
	var csUint64 uint64
	if s, ok := cs.(string); ok {
		csUint64 = uint64(mustParseInt(s))
	} else if f, ok := cs.(float64); ok {
		csUint64 = uint64(f)
	}
	if csUint64 != goVal.(uint64) {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, csUint64, goVal)
	}
}

func compareUint32Value(t *testing.T, cs, goVal any, path string) {
	var csUint32 uint32
	if s, ok := cs.(string); ok {
		csUint32 = uint32(mustParseInt(s))
	} else if f, ok := cs.(float64); ok {
		csUint32 = uint32(f)
	}
	if csUint32 != goVal.(uint32) {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, csUint32, goVal)
	}
}

func compareTimestampValue(t *testing.T, cs any, goVal time.Time, path string) {
	csMap := cs.(map[string]any)
	var csSeconds int64
	if s, ok := csMap["Seconds"].(string); ok {
		csSeconds = mustParseInt(s)
	} else if f, ok := csMap["Seconds"].(float64); ok {
		csSeconds = int64(f)
	}
	if csSeconds != goVal.Unix() {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, csSeconds, goVal.Unix())
	}
}

func compareDurationValue(t *testing.T, cs any, goVal time.Duration, path string) {
	csMap := cs.(map[string]any)
	var csSeconds int64
	if s, ok := csMap["Seconds"].(string); ok {
		csSeconds = mustParseInt(s)
	} else if f, ok := csMap["Seconds"].(float64); ok {
		csSeconds = int64(f)
	}
	goSeconds := goVal.Milliseconds() / 1000
	if csSeconds != goSeconds {
		t.Fatalf("%s mismatch: cs=%v, go=%v", path, csSeconds, goSeconds)
	}
}

func mustParseInt(s string) int64 {
	var n int64
	fmt.Sscanf(s, "%d", &n)
	return n
}

func compareRepeatedInt32(t *testing.T, cs []any, goList kdsync.Repeated[int32], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareValue(t, cs[i], goList.Get(i), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedInt64(t *testing.T, cs []any, goList kdsync.Repeated[int64], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareInt64Value(t, cs[i], goList.Get(i), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedFloat(t *testing.T, cs []any, goList kdsync.Repeated[float32], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareValue(t, cs[i], goList.Get(i), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedDouble(t *testing.T, cs []any, goList kdsync.Repeated[float64], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareValue(t, cs[i], goList.Get(i), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedBool(t *testing.T, cs []any, goList kdsync.Repeated[bool], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareValue(t, cs[i], goList.Get(i), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedString(t *testing.T, cs []any, goList kdsync.Repeated[string], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareValue(t, cs[i], goList.Get(i), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedTimestamp(t *testing.T, cs []any, goList kdsync.Repeated[time.Time], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		csMap := cs[i].(map[string]any)
		var csSeconds int64
		if s, ok := csMap["Seconds"].(string); ok {
			csSeconds = mustParseInt(s)
		} else if f, ok := csMap["Seconds"].(float64); ok {
			csSeconds = int64(f)
		}
		goTime := goList.Get(i).Unix()
		compareValue(t, csSeconds, goTime, fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedDuration(t *testing.T, cs []any, goList kdsync.Repeated[time.Duration], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		csMap := cs[i].(map[string]any)
		var csSeconds int64
		if s, ok := csMap["Seconds"].(string); ok {
			csSeconds = mustParseInt(s)
		} else if f, ok := csMap["Seconds"].(float64); ok {
			csSeconds = int64(f)
		}
		goDur := goList.Get(i).Milliseconds() / 1000
		compareValue(t, csSeconds, goDur, fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareRepeatedEnum(t *testing.T, cs []any, goList kdsync.Repeated[kds.ItemType], path string) {
	if len(cs) != goList.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goList.Len())
	}
	for i := 0; i < len(cs); i++ {
		compareValue(t, cs[i], int(goList.Get(i)), fmt.Sprintf("%s[%d]", path, i))
	}
}

func compareMapInt32Int32(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, int32], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := int32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt64Int64(t *testing.T, cs map[string]any, goMap kdsync.Map[int64, int64], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := mustParseInt(k)
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt64String(t *testing.T, cs map[string]any, goMap kdsync.Map[int64, string], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := mustParseInt(k)
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapBoolDuration(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, time.Duration], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k == "true")
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareDurationValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt64ItemData(t *testing.T, cs map[string]any, goMap kdsync.Map[int64, *kds.ItemData], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := mustParseInt(k)
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		itemData := v.(map[string]any)
		compareValue(t, itemData["Id"], goVal.GetId(), path+"."+k+".id")
		compareValue(t, itemData["Name"], goVal.GetName(), path+"."+k+".name")
		compareValue(t, itemData["Count"], goVal.GetCount(), path+"."+k+".count")
	}
}

func compareMapBoolItemData(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, *kds.ItemData], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k == "true")
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		itemData := v.(map[string]any)
		compareValue(t, itemData["Id"], goVal.GetId(), path+"."+k+".id")
		compareValue(t, itemData["Name"], goVal.GetName(), path+"."+k+".name")
		compareValue(t, itemData["Count"], goVal.GetCount(), path+"."+k+".count")
	}
}

func compareMapUint32Uint32(t *testing.T, cs map[string]any, goMap kdsync.Map[uint32, uint32], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := uint32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareUint32Value(t, v, goVal, path+"."+k)
	}
}

func compareMapUint64Uint64(t *testing.T, cs map[string]any, goMap kdsync.Map[uint64, uint64], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := uint64(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareUint64Value(t, v, goVal, path+"."+k)
	}
}

func compareMapBoolFloat(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, float32], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := k == "true"
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapStringDouble(t *testing.T, cs map[string]any, goMap kdsync.Map[string, float64], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt32Bool(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, bool], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := int32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapUint32Bytes(t *testing.T, cs map[string]any, goMap kdsync.Map[uint32, []byte], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := uint32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareBytesValue(t, v, goVal, path+"."+k)
	}
}

func compareMapUint64Timestamp(t *testing.T, cs map[string]any, goMap kdsync.Map[uint64, time.Time], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := uint64(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareTimestampValue(t, v, goVal, path+"."+k)
	}
}

func compareMapStringEmpty(t *testing.T, cs map[string]any, goMap kdsync.Map[string, struct{}], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k := range cs {
		_, ok := goMap.Get(k)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
	}
}

func compareMapInt32ItemType(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, kds.ItemType], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := int32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, int(goVal), path+"."+k)
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
