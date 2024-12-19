package codegen

import (
	"reflect"
	"text/template"
	"unicode"
)

func Funcs(ctx *Context) template.FuncMap {
	return template.FuncMap{
		"lcFirst": lcFirst,
		"ucFirst": ucFirst,

		"toGoType": ctx.toGoType,
		"toProtoType": toProtoType,
		"toProtoGoType": ctx.toProtoGoType,

		"findEnum": ctx.FindEnum,
		"findEntity": ctx.FindEntity,
		"findComponent": ctx.FindComponent,
		"findProtoPackage": ctx.FindProtoPackage,

		"findList": ctx.FindList,
		"findMap": ctx.FindMap,
	}
}

func lcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func ucFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (ctx *Context) toGoType(type_ string) string {
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

func (ctx *Context) toProtoGoType(type_ string) string {
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
		if def, ok := ctx.Defs[type_]; ok {
			return def.GetProtoPackage() + "." + type_
		}
		return type_
	}
}

func toProtoType(type_ string) string {
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

var enumType = reflect.TypeOf((*Enum)(nil))
var entityType = reflect.TypeOf((*Entity)(nil))
var componentType = reflect.TypeOf((*Component)(nil))

// 
func IsEnum(def interface{}) bool {
	return reflect.TypeOf(def) == enumType
}

func IsEntity(def interface{}) bool {
	return reflect.TypeOf(def) == entityType
}

func IsComponent(def interface{}) bool {
	return reflect.TypeOf(def) == componentType
}

func isTimestamp(name string) bool {
	return "timestamp" == name
}

func isDuration(name string) bool {
	return "duration" == name
}

func isEmpty(name string) bool {
	return "empty" == name
}