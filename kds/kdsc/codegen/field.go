package codegen

import (
	"github.com/iakud/keeper/kds/kdsc/parser"
	"strconv"
)

type FieldKind int32
const (
	FieldKind_Primitive FieldKind = iota
	FieldKind_Enum
	FieldKind_Entity
	FieldKind_Component
)

type Field struct {
	Repeated bool
	Type string
	KeyType string
	ValueType string
	Name string
	Number int
	Kind string
}

func (f *Field) IsComponent() bool {
	switch f.Type {
	case "double", "float", "int32", "int64", "sint32", "sint64", "fixed32", "fixed64", "bool", "string", "bytes":
		break
	default:
		return true
	}
	return false
}

func (f *Field) IsEnum() bool {
	return false
}

func newField(ctx parser.IFieldContext) *Field {
	field := new(Field)
	labelCtx := ctx.FieldLabel();
	field.Repeated = labelCtx != nil && labelCtx.REPEATED() != nil
	field.Type = ctx.Type_().GetText()
	field.Name = ctx.FieldName().GetText()
	field.Number, _ = strconv.Atoi(ctx.FieldNumber().GetText())
	return field
}

func newMapField(ctx parser.IMapFieldContext) *Field {
	field := new(Field)
	field.KeyType = ctx.KeyType().GetText()
	field.Type = ctx.Type_().GetText()
	field.Name = ctx.MapName().GetText()
	field.Number, _ = strconv.Atoi(ctx.FieldNumber().GetText())
	return field
}

func (v *Field) String() string {
	return v.Name
}