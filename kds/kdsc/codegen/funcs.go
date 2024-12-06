package codegen

import (
	"reflect"
	"text/template"
)

func Funcs(ctx *Context) template.FuncMap {
	return template.FuncMap{
		"IsEnum": IsEnum,
		"IsEntity": IsEntity,
		"IsComponent": IsComponent,
		"FindEnum": ctx.FindEnum,
		"FindEntity": ctx.FindEntity,
		"FindComponent": ctx.FindComponent,
		"GoType": GoType,
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
	default:
		return type_
	}
}