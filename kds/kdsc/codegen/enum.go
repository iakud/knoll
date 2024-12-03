package codegen

import (
	"strconv"

	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Enum struct {
	Name string
	EnumFields []*EnumField
}

func newEnum(ctx parser.IEnumDefContext) *Enum {
	enum := new(Enum)
	enum.Name = ctx.EnumName().GetText()
	for _, element := range ctx.EnumBody().AllEnumElement() {
		enum.EnumFields = append(enum.EnumFields, newEnumField(element.EnumField()))
	}
	return enum
}

type EnumField struct {
	Name string
	Value int
}

func newEnumField(ctx parser.IEnumFieldContext) *EnumField {
	enumField := new(EnumField)
	enumField.Name = ctx.Ident().GetText()
	if ctx.MINUS() != nil {
		enumField.Value, _ = strconv.Atoi(ctx.MINUS().GetText() + ctx.IntLit().GetText())
	} else {
		enumField.Value, _ = strconv.Atoi(ctx.IntLit().GetText())
	}	
	return enumField
}