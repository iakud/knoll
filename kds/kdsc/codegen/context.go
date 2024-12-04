package codegen

import (
	"strconv"
	"strings"

	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Context struct {
	KdsContexts []*Kds

	Imports map[string]*Kds
	Enums map[string]*Enum
	Entities map[string]*Entity
	Components map[string]*Component
}

type Kds struct {
	Filename string
	Package string
	Imports []string
	Enums []*Enum
	Entities []*Entity
	Components []*Component

	Defs []TopLevelDef
}

type TopLevelDef interface{
	Enum() *Enum
	Entity() *Entity
	Component() *Component
}

type Enum struct {
	Name string
	EnumFields []*EnumField
}

type EnumField struct {
	Name string
	Value int
}

type Entity struct {
	Name string
	Fields []*Field
}

type Component struct {
	Name string
	Fields []*Field
}

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

func (ctx *Context) VisitKds(kdsCtx parser.IKdsContext) *Kds {
	kds := new(Kds)
	kds.Package = kdsCtx.PackageStatement().FullIdent().GetText()
	for _, importStatement := range kdsCtx.AllImportStatement() {
		kds.Imports = append(kds.Imports, ctx.VisitImport(importStatement))
	}
	for i := 0; i < len(kds.Imports); i++ {
		kds.Imports[i] = strings.TrimSuffix(kds.Imports[i], ".kds")
	}
	
	for _, topLevel := range kdsCtx.AllTopLevelDef() {
		switch {
		case topLevel.EnumDef() != nil:
			kds.Enums = append(kds.Enums, ctx.VisitEnum(topLevel.EnumDef()))
		case topLevel.EntityDef() != nil:
			kds.Entities = append(kds.Entities, ctx.VisitEntity(topLevel.EntityDef()))
		case topLevel.ComponentDef() != nil:
			kds.Components = append(kds.Components, ctx.VisitComponent(topLevel.ComponentDef()))
		}
	}
	return kds
}

func (ctx *Context) VisitImport(importCtx parser.IImportStatementContext) string {
	importText := importCtx.STR_LIT().GetText()
	switch {
	case strings.HasPrefix(importText, "\"") && strings.HasSuffix(importText, "\"") :
		return strings.TrimSuffix(strings.TrimPrefix(importText, "\""), "\"")
	case strings.HasPrefix(importText, "'") && strings.HasSuffix(importText, "'") :
		return strings.TrimSuffix(strings.TrimPrefix(importText, "'"), "'")
	}
	return importText
}

func (ctx *Context) VisitEnum(enumCtx parser.IEnumDefContext) *Enum {
	enum := new(Enum)
	enum.Name = enumCtx.EnumName().GetText()
	for _, element := range enumCtx.EnumBody().AllEnumElement() {
		enum.EnumFields = append(enum.EnumFields, ctx.VisitEnumField(element.EnumField()))
	}
	return enum
}

func (ctx *Context) VisitEnumField(enumFieldCtx parser.IEnumFieldContext) *EnumField {
	enumField := new(EnumField)
	enumField.Name = enumFieldCtx.Ident().GetText()
	if enumFieldCtx.MINUS() != nil {
		enumField.Value, _ = strconv.Atoi(enumFieldCtx.MINUS().GetText() + enumFieldCtx.IntLit().GetText())
	} else {
		enumField.Value, _ = strconv.Atoi(enumFieldCtx.IntLit().GetText())
	}
	return enumField
}

func (ctx *Context) VisitEntity(entityCtx parser.IEntityDefContext) *Entity {
	entity := new(Entity)
	entity.Name = entityCtx.EntityName().GetText()
	for _, element := range entityCtx.EntityBody().AllEntityElement() {
		switch {
		case element.Field() != nil:
			entity.Fields = append(entity.Fields, ctx.VisitField(element.Field()))
		case element.MapField() != nil:
			entity.Fields = append(entity.Fields, ctx.VisitMapField(element.MapField()))
		}
	}
	return entity
}

func (ctx *Context)  VisitComponent(componentCtx parser.IComponentDefContext) *Component {
	component := new(Component)
	component.Name = componentCtx.ComponentName().GetText()
	for _, element := range componentCtx.ComponentBody().AllComponentElement() {
		switch {
		case element.Field() != nil:
			component.Fields = append(component.Fields, ctx.VisitField(element.Field()))
		case element.MapField() != nil:
			component.Fields = append(component.Fields, ctx.VisitMapField(element.MapField()))
		}
	}
	return component
}

func (ctx *Context) VisitField(fieldCtx parser.IFieldContext) *Field {
	field := new(Field)
	field.Repeated = fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil
	field.Type = fieldCtx.Type_().GetText()
	field.Name = fieldCtx.FieldName().GetText()
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())
	return field
}

func (ctx *Context) VisitMapField(mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	field.KeyType = mapFieldCtx.KeyType().GetText()
	field.Type = mapFieldCtx.Type_().GetText()
	field.Name = mapFieldCtx.MapName().GetText()
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())
	return field
}