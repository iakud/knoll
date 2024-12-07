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
		"IsEnum": IsEnum,
		"IsEntity": IsEntity,
		"IsComponent": IsComponent,
		"FindEnum": ctx.FindEnum,
		"FindEntity": ctx.FindEntity,
		"FindComponent": ctx.FindComponent,
		"GoType": GoType,
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
	r[0] = unicode.ToLower(r[0])
	return string(r)
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