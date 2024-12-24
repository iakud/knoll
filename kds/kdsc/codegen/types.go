package codegen

func Kind(type_ string) string {
	switch type_ {
	case "double", "float", "int32", "int64", "sint32", "sint64", "fixed32", "fixed64", "sfixed32", "sfixed64", "bool", "string", "bytes":
		return "primitive"
	case "timestamp", "duration", "empty":
		return "proto"
	default:
		return "unknow"
	}
}

func GoType(type_ string) string {
	switch type_ {
	case "double":
		return "float64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "sint32":
		return "int32"
	case "sint64":
		return "int64"
	case "fixed32":
		return "uint32"
	case "fixed64":
		return "uint64"
	case "sfixed32":
		return "int32"
	case "sfixed64":
		return "int64"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "bytes":
		return "[]byte"
	case "timestamp":
		return "time.Time"
	case "duration":
		return "time.Duration"
	case "empty":
		return "struct{}"
	default:
		return type_
	}
}

func GoProtoType(type_ string) string {
	switch type_ {
	case "double":
		return "float64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "sint32":
		return "int32"
	case "sint64":
		return "int64"
	case "fixed32":
		return "uint32"
	case "fixed64":
		return "uint64"
	case "sfixed32":
		return "int32"
	case "sfixed64":
		return "int64"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "bytes":
		return "[]byte"
	case "timestamp":
		return "timestamppb.Timestamp"
	case "duration":
		return "durationpb.Duration"
	case "empty":
		return "emptypb.Empty"
	default:
		return type_
	}
}

func ProtoType(type_ string) string {
	switch type_ {
	case "timestamp":
		return "google.protobuf.Timestamp"
	case "duration":
		return "google.protobuf.Duration"
	case "empty":
		return "google.protobuf.Empty"
	default:
		return type_
	}
}
