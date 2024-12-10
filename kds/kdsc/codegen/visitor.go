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
			entity := visitEntity(ctx, kds, topLevel.EntityDef())
			ctx.Defs[entity.Name] = entity
			kds.Defs = append(kds.Defs, entity)
		case topLevel.ComponentDef() != nil:
			component := visitComponent(ctx, kds, topLevel.ComponentDef())
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
	enum.Name = GoCamelCase(enumCtx.EnumName().GetText())
	for _, element := range enumCtx.EnumBody().AllEnumElement() {
		enum.EnumFields = append(enum.EnumFields, visitEnumField(element.EnumField()))
	}
	return enum
}

func visitEnumField(enumFieldCtx parser.IEnumFieldContext) *EnumField {
	enumField := new(EnumField)
	enumField.Name = GoCamelCase(enumFieldCtx.Ident().GetText())
	if enumFieldCtx.MINUS() != nil {
		enumField.Value, _ = strconv.Atoi(enumFieldCtx.MINUS().GetText() + enumFieldCtx.IntLit().GetText())
	} else {
		enumField.Value, _ = strconv.Atoi(enumFieldCtx.IntLit().GetText())
	}
	return enumField
}

func visitEntity(ctx *Context, kds *Kds, entityCtx parser.IEntityDefContext) *Entity {
	entity := new(Entity)
	entity.Name = GoCamelCase(entityCtx.EntityName().GetText())
	for _, element := range entityCtx.EntityBody().AllEntityElement() {
		switch {
		case element.Field() != nil:
			field := visitField(ctx, kds, element.Field())
			entity.Fields = append(entity.Fields, field)
			if field.Repeated {
				ctx.AddArray(field.Type)
			}
		case element.MapField() != nil:
			field := visitMapField(ctx, kds, element.MapField())
			entity.Fields = append(entity.Fields, field)
			ctx.AddMap(field.Type, field.KeyType)
		}
	}
	return entity
}

func visitComponent(ctx *Context, kds *Kds, componentCtx parser.IComponentDefContext) *Component {
	component := new(Component)
	component.Name = componentCtx.ComponentName().GetText()
	for _, element := range componentCtx.ComponentBody().AllComponentElement() {
		switch {
		case element.Field() != nil:
			field := visitField(ctx, kds, element.Field())
			component.Fields = append(component.Fields, field)
			if field.Repeated {
				ctx.AddArray(field.Type)
			}
		case element.MapField() != nil:
			field := visitMapField(ctx, kds, element.MapField())
			component.Fields = append(component.Fields, field)
			ctx.AddMap(field.Type, field.KeyType)
		}
	}
	return component
}

func visitField(ctx *Context, kds *Kds, fieldCtx parser.IFieldContext) *Field {
	field := new(Field)
	field.Repeated = fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil
	if fieldCtx.Type_().MessageType() != nil || fieldCtx.Type_().EnumType() != nil {
		field.Type = GoCamelCase(fieldCtx.Type_().GetText())
	} else {
		field.Type = fieldCtx.Type_().GetText()
	}
	field.Name = GoCamelCase(fieldCtx.FieldName().GetText())
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))
	field.GoType = GoType(field.Type)
	
	kds.ImportTimestamp = kds.ImportTimestamp || fieldCtx.Type_().TIMESTAMP() != nil
	kds.ImportDuration = kds.ImportDuration || fieldCtx.Type_().DURATION() != nil
	kds.ImportEmpty = kds.ImportEmpty || fieldCtx.Type_().EMPTY() != nil
	return field
}

func visitMapField(ctx *Context, kds *Kds, mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	field.KeyType = mapFieldCtx.KeyType().GetText()
	if mapFieldCtx.Type_().MessageType() != nil || mapFieldCtx.Type_().EnumType() != nil {
		field.Type = GoCamelCase(mapFieldCtx.Type_().GetText())
	} else {
		field.Type = mapFieldCtx.Type_().GetText()
	}
	field.Name = GoCamelCase(mapFieldCtx.MapName().GetText())
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))
	field.GoType = GoType(field.Type)

	kds.ImportTimestamp = kds.ImportTimestamp || mapFieldCtx.Type_().TIMESTAMP() != nil
	kds.ImportDuration = kds.ImportDuration || mapFieldCtx.Type_().DURATION() != nil
	kds.ImportEmpty = kds.ImportEmpty || mapFieldCtx.Type_().EMPTY() != nil
	return field
}