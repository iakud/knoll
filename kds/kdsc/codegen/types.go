package codegen

func Kind(type_ string) string {
	switch type_ {
	case "bool", "int32", "sint32", "uint32", "int64", "sint64", "uint64", "sfixed32", "fixed32", "float", "sfixed64", "fixed64", "double", "string", "bytes":
		return "primitive"
	case "timestamp", "duration", "empty":
		return "proto"
	default:
		return "unknow"
	}
}

func GoType(type_ string) string {
	switch type_ {
	case "bool":
		return "bool"
	case "int32":
		return "int32"
	case "sint32":
		return "int32"
	case "uint32":
		return "uint32"
	case "int64":
		return "int64"
	case "sint64":
		return "int64"
	case "uint64":
		return "uint64"
	case "sfixed32":
		return "int32"
	case "fixed32":
		return "uint32"
	case "float":
		return "float32"
	case "sfixed64":
		return "int64"
	case "fixed64":
		return "uint64"
	case "double":
		return "float64"
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
	case "bool":
		return "bool"
	case "int32":
		return "int32"
	case "sint32":
		return "int32"
	case "uint32":
		return "uint32"
	case "int64":
		return "int64"
	case "sint64":
		return "int64"
	case "uint64":
		return "uint64"
	case "sfixed32":
		return "int32"
	case "fixed32":
		return "uint32"
	case "float":
		return "float32"
	case "sfixed64":
		return "int64"
	case "fixed64":
		return "uint64"
	case "double":
		return "float64"
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
