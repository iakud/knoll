package codegen

import (
	"strconv"
	"strings"

	"github.com/iakud/keeper/kds/kdsc/parser"
)


func visitKds(ctx *Context, kdsCtx parser.IKdsContext) *Kds {
	kds := new(Kds)
	kds.Package = kdsCtx.PackageStatement().FullIdent().GetText()
	for _, importStatement := range kdsCtx.AllImportStatement() {
		kds.Imports = append(kds.Imports, visitImport(importStatement))
	}
	for i := 0; i < len(kds.Imports); i++ {
		kds.Imports[i] = strings.TrimSuffix(kds.Imports[i], ".kds")
	}
	
	for _, topLevel := range kdsCtx.AllTopLevelDef() {
		switch {
		case topLevel.EnumDef() != nil:
			enum := visitEnum(kds, topLevel.EnumDef())
			ctx.Defs[enum.Name] = enum
			kds.Defs = append(kds.Defs, enum)
		case topLevel.EntityDef() != nil:
			entity := visitEntity(kds, topLevel.EntityDef())
			ctx.Defs[entity.Name] = entity
			kds.Defs = append(kds.Defs, entity)
		case topLevel.ComponentDef() != nil:
			component := visitComponent(kds, topLevel.ComponentDef())
			ctx.Defs[component.Name] = component
			kds.Defs = append(kds.Defs, component)
		}
	}
	return kds
}

func visitImport(importCtx parser.IImportStatementContext) string {
	importText := importCtx.STR_LIT().GetText()
	switch {
	case strings.HasPrefix(importText, "\"") && strings.HasSuffix(importText, "\"") :
		return strings.TrimSuffix(strings.TrimPrefix(importText, "\""), "\"")
	case strings.HasPrefix(importText, "'") && strings.HasSuffix(importText, "'") :
		return strings.TrimSuffix(strings.TrimPrefix(importText, "'"), "'")
	}
	return importText
}

func visitEnum(kds *Kds, enumCtx parser.IEnumDefContext) *Enum {
	enum := new(Enum)
	enum.Name = enumCtx.EnumName().GetText()
	for _, element := range enumCtx.EnumBody().AllEnumElement() {
		enum.EnumFields = append(enum.EnumFields, visitEnumField(element.EnumField()))
	}
	return enum
}

func visitEnumField(enumFieldCtx parser.IEnumFieldContext) *EnumField {
	enumField := new(EnumField)
	enumField.Name = enumFieldCtx.Ident().GetText()
	if enumFieldCtx.MINUS() != nil {
		enumField.Value, _ = strconv.Atoi(enumFieldCtx.MINUS().GetText() + enumFieldCtx.IntLit().GetText())
	} else {
		enumField.Value, _ = strconv.Atoi(enumFieldCtx.IntLit().GetText())
	}
	return enumField
}

func visitEntity(kds *Kds, entityCtx parser.IEntityDefContext) *Entity {
	entity := new(Entity)
	entity.Name = entityCtx.EntityName().GetText()
	for _, element := range entityCtx.EntityBody().AllEntityElement() {
		switch {
		case element.Field() != nil:
			entity.Fields = append(entity.Fields, visitField(kds, element.Field()))
		case element.MapField() != nil:
			entity.Fields = append(entity.Fields, visitMapField(kds, element.MapField()))
		}
	}
	return entity
}

func visitComponent(kds *Kds, componentCtx parser.IComponentDefContext) *Component {
	component := new(Component)
	component.Name = componentCtx.ComponentName().GetText()
	for _, element := range componentCtx.ComponentBody().AllComponentElement() {
		switch {
		case element.Field() != nil:
			field := visitField(kds, element.Field())
			component.Fields = append(component.Fields, field)
			
			kds.ImportTimestamp = kds.ImportTimestamp || field.IsTimestamp
			kds.ImportDuration = kds.ImportDuration || field.IsDuration
		case element.MapField() != nil:
			field := visitMapField(kds, element.MapField())
			component.Fields = append(component.Fields, field)

			kds.ImportTimestamp = kds.ImportTimestamp || field.IsTimestamp
			kds.ImportDuration = kds.ImportDuration || field.IsDuration
		}
	}
	return component
}

func visitField(kds *Kds, fieldCtx parser.IFieldContext) *Field {
	field := new(Field)
	field.Repeated = fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil
	fieldCtx.Type_().MessageType()
	field.Type = fieldCtx.Type_().GetText()
	field.Name = fieldCtx.FieldName().GetText()
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())

	field.IsTimestamp = fieldCtx.Type_().TIMESTAMP() != nil
	field.IsDuration = fieldCtx.Type_().DURATION() != nil
	return field
}

func visitMapField(kds *Kds, mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	field.KeyType = mapFieldCtx.KeyType().GetText()
	field.Type = mapFieldCtx.Type_().GetText()
	field.Name = mapFieldCtx.MapName().GetText()
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())

	field.IsTimestamp = mapFieldCtx.Type_().TIMESTAMP() != nil
	field.IsDuration = mapFieldCtx.Type_().DURATION() != nil
	return field
}