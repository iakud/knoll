package codegen

import (
	"strconv"
	"strings"

	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Context struct {
	KdsContexts []*Kds

	Imports map[string]*Kds
	Defs map[string]interface{}
}

func newContext() *Context {
	return &Context{
		Imports: make(map[string]*Kds),
		Defs: make(map[string]interface{}),
	}
}

func (ctx *Context) FindEnum(name string) *Enum {
	topLevelDef, ok := ctx.Defs[name]
	if !ok {
		return nil
	}
	enum, ok := topLevelDef.(*Enum)
	if !ok {
		return nil
	}
	return enum
}

func (ctx *Context) FindEntity(name string) *Entity {
	topLevelDef, ok := ctx.Defs[name]
	if !ok {
		return nil
	}
	entity, ok := topLevelDef.(*Entity)
	if !ok {
		return nil
	}
	return entity
}

func (ctx *Context) FindComponent(name string) *Component {
	topLevelDef, ok := ctx.Defs[name]
	if !ok {
		return nil
	}
	component, ok := topLevelDef.(*Component)
	if !ok {
		return nil
	}
	return component
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
			enum := ctx.VisitEnum(topLevel.EnumDef())
			kds.Enums = append(kds.Enums, enum)
			kds.Defs = append(kds.Defs, enum)
		case topLevel.EntityDef() != nil:
			entity := ctx.VisitEntity(topLevel.EntityDef())
			kds.Entities = append(kds.Entities, entity)
			kds.Defs = append(kds.Defs, entity)
		case topLevel.ComponentDef() != nil:
			component := ctx.VisitComponent(topLevel.ComponentDef())
			kds.Components = append(kds.Components, component)
			kds.Defs = append(kds.Defs, component)
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
	ctx.Defs[enum.Name] = enum
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
	ctx.Defs[entity.Name] = entity
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
	ctx.Defs[component.Name] = component
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