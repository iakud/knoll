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

		"findEnum":      ctx.FindEnum,
		"findEntity":    ctx.FindEntity,
		"findComponent": ctx.FindComponent,

		"findList": ctx.FindList,
		"findMap":  ctx.FindMap,
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

var enumType = reflect.TypeOf((*Enum)(nil))
var entityType = reflect.TypeOf((*Entity)(nil))
var componentType = reflect.TypeOf((*Component)(nil))

func IsEnum(def interface{}) bool {
	return reflect.TypeOf(def) == enumType
}

func IsEntity(def interface{}) bool {
	return reflect.TypeOf(def) == entityType
}

func IsComponent(def interface{}) bool {
	return reflect.TypeOf(def) == componentType
}
