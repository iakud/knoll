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
	compareMapInt32Int32(t, cs["Int32Int32"].(map[string]any), goMaps.GetInt32Int32(), path+".Int32Int32")
	compareMapInt32String(t, cs["Int32String"].(map[string]any), goMaps.GetInt32String(), path+".Int32String")
	compareMapInt32Timestamp(t, cs["Int32Timestamp"].(map[string]any), goMaps.GetInt32Timestamp(), path+".Int32Timestamp")
	compareMapInt32Duration(t, cs["Int32Duration"].(map[string]any), goMaps.GetInt32Duration(), path+".Int32Duration")
	compareMapInt32Enum(t, cs["Int32Enum"].(map[string]any), goMaps.GetInt32Enum(), path+".Int32Enum")
	compareMapInt32ItemData(t, cs["Int32ItemData"].(map[string]any), goMaps.GetInt32ItemData(), path+".Int32ItemData")
	// skip Int32Empty - nosync

	compareMapInt64Int64(t, cs["Int64Int64"].(map[string]any), goMaps.GetInt64Int64(), path+".Int64Int64")
	compareMapInt64String(t, cs["Int64String"].(map[string]any), goMaps.GetInt64String(), path+".Int64String")
	compareMapInt64Timestamp(t, cs["Int64Timestamp"].(map[string]any), goMaps.GetInt64Timestamp(), path+".Int64Timestamp")
	compareMapInt64Duration(t, cs["Int64Duration"].(map[string]any), goMaps.GetInt64Duration(), path+".Int64Duration")
	compareMapInt64Enum(t, cs["Int64Enum"].(map[string]any), goMaps.GetInt64Enum(), path+".Int64Enum")
	compareMapInt64ItemData(t, cs["Int64ItemData"].(map[string]any), goMaps.GetInt64ItemData(), path+".Int64ItemData")
	// skip Int64Empty - nosync

	compareMapStringInt32(t, cs["StringInt32"].(map[string]any), goMaps.GetStringInt32(), path+".StringInt32")
	compareMapStringString(t, cs["StringString"].(map[string]any), goMaps.GetStringString(), path+".StringString")
	compareMapStringTimestamp(t, cs["StringTimestamp"].(map[string]any), goMaps.GetStringTimestamp(), path+".StringTimestamp")
	compareMapStringDuration(t, cs["StringDuration"].(map[string]any), goMaps.GetStringDuration(), path+".StringDuration")
	compareMapStringEnum(t, cs["StringEnum"].(map[string]any), goMaps.GetStringEnum(), path+".StringEnum")
	compareMapStringItemData(t, cs["StringItemData"].(map[string]any), goMaps.GetStringItemData(), path+".StringItemData")
	// skip StringEmpty - nosync

	compareMapBoolInt32(t, cs["BoolInt32"].(map[string]any), goMaps.GetBoolInt32(), path+".BoolInt32")
	compareMapBoolString(t, cs["BoolString"].(map[string]any), goMaps.GetBoolString(), path+".BoolString")
	compareMapBoolTimestamp(t, cs["BoolTimestamp"].(map[string]any), goMaps.GetBoolTimestamp(), path+".BoolTimestamp")
	compareMapBoolDuration(t, cs["BoolDuration"].(map[string]any), goMaps.GetBoolDuration(), path+".BoolDuration")
	// skip BoolEmpty - nosync
	// skip BoolEnum - nosync,nopersist
	compareMapBoolItemData(t, cs["BoolItemData"].(map[string]any), goMaps.GetBoolItemData(), path+".BoolItemData")
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

func compareMapInt32String(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, string], path string) {
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

func compareMapInt32Timestamp(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, time.Time], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := int32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareTimestampValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt32Duration(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, time.Duration], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := int32(mustParseInt(k))
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareDurationValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt32Enum(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, kds.ItemType], path string) {
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

func compareMapInt32ItemData(t *testing.T, cs map[string]any, goMap kdsync.Map[int32, *kds.ItemData], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := int32(mustParseInt(k))
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

func compareMapInt64Timestamp(t *testing.T, cs map[string]any, goMap kdsync.Map[int64, time.Time], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := mustParseInt(k)
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareTimestampValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt64Duration(t *testing.T, cs map[string]any, goMap kdsync.Map[int64, time.Duration], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := mustParseInt(k)
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareDurationValue(t, v, goVal, path+"."+k)
	}
}

func compareMapInt64Enum(t *testing.T, cs map[string]any, goMap kdsync.Map[int64, kds.ItemType], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		key := mustParseInt(k)
		goVal, ok := goMap.Get(key)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, int(goVal), path+"."+k)
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

func compareMapStringInt32(t *testing.T, cs map[string]any, goMap kdsync.Map[string, int32], path string) {
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

func compareMapStringString(t *testing.T, cs map[string]any, goMap kdsync.Map[string, string], path string) {
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

func compareMapStringTimestamp(t *testing.T, cs map[string]any, goMap kdsync.Map[string, time.Time], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareTimestampValue(t, v, goVal, path+"."+k)
	}
}

func compareMapStringDuration(t *testing.T, cs map[string]any, goMap kdsync.Map[string, time.Duration], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareDurationValue(t, v, goVal, path+"."+k)
	}
}

func compareMapStringEnum(t *testing.T, cs map[string]any, goMap kdsync.Map[string, kds.ItemType], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, int(goVal), path+"."+k)
	}
}

func compareMapStringItemData(t *testing.T, cs map[string]any, goMap kdsync.Map[string, *kds.ItemData], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k)
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		itemData := v.(map[string]any)
		compareValue(t, itemData["Id"], goVal.GetId(), path+"."+k+".id")
		compareValue(t, itemData["Name"], goVal.GetName(), path+"."+k+".name")
		compareValue(t, itemData["Count"], goVal.GetCount(), path+"."+k+".count")
	}
}

func compareMapBoolInt32(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, int32], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k == "true")
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapBoolString(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, string], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k == "true")
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, goVal, path+"."+k)
	}
}

func compareMapBoolTimestamp(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, time.Time], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k == "true")
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareTimestampValue(t, v, goVal, path+"."+k)
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

func compareMapBoolEnum(t *testing.T, cs map[string]any, goMap kdsync.Map[bool, kds.ItemType], path string) {
	if len(cs) != goMap.Len() {
		t.Fatalf("%s length mismatch: cs=%d, go=%d", path, len(cs), goMap.Len())
	}
	for k, v := range cs {
		goVal, ok := goMap.Get(k == "true")
		if !ok {
			t.Fatalf("%s missing key %s", path, k)
		}
		compareValue(t, v, int(goVal), path+"."+k)
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
	maps.GetInt32Int32().Set(3, 300)
	syncUpdate(t)

	maps.GetInt32String().Set(3, "three")
	syncUpdate(t)

	maps.GetInt32Timestamp().Set(2, time.Unix(2000, 0))
	syncUpdate(t)

	maps.GetInt32Duration().Set(2, time.Minute)
	syncUpdate(t)

	maps.GetInt32Enum().Set(2, kds.ItemType_ItemTypeArmor)
	syncUpdate(t)

	itemMap2 := kds.NewItemData()
	itemMap2.SetId(7)
	itemMap2.SetName("boots")
	itemMap2.SetCount(1)
	maps.GetInt32ItemData().Set(2, itemMap2)
	syncUpdate(t)

	maps.GetInt64Int64().Set(3, 3000)
	syncUpdate(t)

	maps.GetInt64String().Set(3, "three")
	syncUpdate(t)

	maps.GetInt64Timestamp().Set(2, time.Unix(2000, 0))
	syncUpdate(t)

	maps.GetInt64Duration().Set(2, time.Minute)
	syncUpdate(t)

	maps.GetInt64Enum().Set(2, kds.ItemType_ItemTypeArmor)
	syncUpdate(t)

	itemMap3 := kds.NewItemData()
	itemMap3.SetId(8)
	itemMap3.SetName("gloves")
	itemMap3.SetCount(2)
	maps.GetInt64ItemData().Set(2, itemMap3)
	syncUpdate(t)

	maps.GetStringInt32().Set("c", 3)
	syncUpdate(t)

	maps.GetStringString().Set("c", "value_c")
	syncUpdate(t)

	maps.GetStringTimestamp().Set("b", time.Unix(2000, 0))
	syncUpdate(t)

	maps.GetStringDuration().Set("b", time.Minute)
	syncUpdate(t)

	maps.GetStringEnum().Set("b", kds.ItemType_ItemTypeArmor)
	syncUpdate(t)

	itemMap4 := kds.NewItemData()
	itemMap4.SetId(9)
	itemMap4.SetName("belt")
	itemMap4.SetCount(4)
	maps.GetStringItemData().Set("b", itemMap4)
	syncUpdate(t)

	maps.GetBoolInt32().Set(false, 2)
	syncUpdate(t)

	maps.GetBoolString().Set(false, "no")
	syncUpdate(t)

	maps.GetBoolTimestamp().Set(false, time.Unix(2000, 0))
	syncUpdate(t)

	maps.GetBoolDuration().Set(false, time.Minute)
	syncUpdate(t)

	maps.GetBoolEnum().Set(false, kds.ItemType_ItemTypeArmor)
	syncUpdate(t)

	itemMap5 := kds.NewItemData()
	itemMap5.SetId(10)
	itemMap5.SetName("amulet")
	itemMap5.SetCount(1)
	maps.GetBoolItemData().Set(false, itemMap5)
	syncUpdate(t)
}

func testMapsUpdate(t *testing.T) {
	// Maps 修改
	maps := all.GetMaps()
	maps.GetInt32Int32().Set(1, 999)
	syncUpdate(t)

	maps.GetInt32String().Set(1, "modified_one")
	syncUpdate(t)

	maps.GetInt32Timestamp().Set(1, time.Unix(8888, 0))
	syncUpdate(t)

	maps.GetInt32Duration().Set(1, time.Hour)
	syncUpdate(t)

	maps.GetInt32Enum().Set(1, kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	itemInt32, _ := maps.GetInt32ItemData().Get(1)
	itemInt32.SetId(111)
	itemInt32.SetName("modified_itemdata")
	itemInt32.SetCount(55)
	syncUpdate(t)

	maps.GetInt64Int64().Set(1, 9999)
	syncUpdate(t)

	maps.GetInt64String().Set(1, "modified_one")
	syncUpdate(t)

	maps.GetInt64Timestamp().Set(1, time.Unix(8888, 0))
	syncUpdate(t)

	maps.GetInt64Duration().Set(1, time.Hour)
	syncUpdate(t)

	maps.GetInt64Enum().Set(1, kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	itemInt64, _ := maps.GetInt64ItemData().Get(1)
	itemInt64.SetId(112)
	itemInt64.SetName("modified_itemdata2")
	itemInt64.SetCount(56)
	syncUpdate(t)

	maps.GetStringInt32().Set("a", 999)
	syncUpdate(t)

	maps.GetStringString().Set("a", "modified_value_a")
	syncUpdate(t)

	maps.GetStringTimestamp().Set("a", time.Unix(9999, 0))
	syncUpdate(t)

	maps.GetStringDuration().Set("a", time.Hour*3)
	syncUpdate(t)

	maps.GetStringEnum().Set("a", kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	itemString, _ := maps.GetStringItemData().Get("a")
	itemString.SetId(113)
	itemString.SetName("modified_itemdata3")
	itemString.SetCount(57)
	syncUpdate(t)

	maps.GetBoolInt32().Set(true, 888)
	syncUpdate(t)

	maps.GetBoolString().Set(true, "modified_true")
	syncUpdate(t)

	maps.GetBoolTimestamp().Set(true, time.Unix(7777, 0))
	syncUpdate(t)

	maps.GetBoolDuration().Set(true, time.Hour*4)
	syncUpdate(t)

	maps.GetBoolEnum().Set(true, kds.ItemType_ItemTypePotion)
	syncUpdate(t)

	itemBool, _ := maps.GetBoolItemData().Get(true)
	itemBool.SetId(114)
	itemBool.SetName("modified_itemdata4")
	itemBool.SetCount(58)
	syncUpdate(t)
}

func testMapsDelete(t *testing.T) {
	// 删除 Map 元素
	maps := all.GetMaps()
	maps.GetInt32Int32().Delete(1)
	syncUpdate(t)

	maps.GetInt32String().Delete(1)
	syncUpdate(t)

	maps.GetInt32Timestamp().Delete(1)
	syncUpdate(t)

	maps.GetInt32Duration().Delete(1)
	syncUpdate(t)

	maps.GetInt32Enum().Delete(1)
	syncUpdate(t)

	maps.GetInt32ItemData().Delete(1)
	syncUpdate(t)

	maps.GetInt64Int64().Delete(1)
	syncUpdate(t)

	maps.GetInt64String().Delete(1)
	syncUpdate(t)

	maps.GetInt64Timestamp().Delete(1)
	syncUpdate(t)

	maps.GetInt64Duration().Delete(1)
	syncUpdate(t)

	maps.GetInt64Enum().Delete(1)
	syncUpdate(t)

	maps.GetInt64ItemData().Delete(1)
	syncUpdate(t)

	maps.GetStringInt32().Delete("a")
	syncUpdate(t)

	maps.GetStringString().Delete("a")
	syncUpdate(t)

	maps.GetStringTimestamp().Delete("a")
	syncUpdate(t)

	maps.GetStringDuration().Delete("a")
	syncUpdate(t)

	maps.GetStringEnum().Delete("a")
	syncUpdate(t)

	maps.GetStringItemData().Delete("a")
	syncUpdate(t)

	maps.GetBoolInt32().Delete(true)
	syncUpdate(t)

	maps.GetBoolString().Delete(true)
	syncUpdate(t)

	maps.GetBoolTimestamp().Delete(true)
	syncUpdate(t)

	maps.GetBoolDuration().Delete(true)
	syncUpdate(t)

	maps.GetBoolEnum().Delete(true)
	syncUpdate(t)

	maps.GetBoolItemData().Delete(true)
	syncUpdate(t)
}
