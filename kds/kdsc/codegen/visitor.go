package codegen

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iakud/knoll/kds/kdsc/parser"
)

func visitKds(ctx *Context, filePath string, kdsCtx parser.IKdsContext) *Kds {
	kds := new(Kds)
	kds.ctx = ctx
	kds.Name = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	kds.SourceFile = filePath
	kds.Package = kdsCtx.PackageStatement().FullIdent().GetText()
	for _, importStatement := range kdsCtx.AllImportStatement() {
		kds.Imports = append(kds.Imports, visitImport(importStatement))
	}
	for i := 0; i < len(kds.Imports); i++ {
		kds.Imports[i] = strings.TrimSuffix(kds.Imports[i], ".kds")
	}
	kds.ProtoGoPackage = visitProtoGoPackage(kdsCtx.ProtoGoPackageStatement())
	kds.ProtoPackage = filepath.Base(kds.ProtoGoPackage)

	for _, topLevel := range kdsCtx.AllTopLevelDef() {
		switch {
		case topLevel.EnumDef() != nil:
			enum := visitEnum(ctx, kds, topLevel.EnumDef())
			ctx.Defs[enum.Name] = enum
			kds.Enums = append(kds.Enums, enum)
			kds.Defs = append(kds.Defs, enum)
		case topLevel.EntityDef() != nil:
			entity := visitEntity(ctx, kds, topLevel.EntityDef())
			ctx.Defs[entity.Name] = entity
			kds.Entities = append(kds.Entities, entity)
			kds.Defs = append(kds.Defs, entity)
		case topLevel.ComponentDef() != nil:
			component := visitComponent(ctx, kds, topLevel.ComponentDef())
			ctx.Defs[component.Name] = component
			kds.Components = append(kds.Components, component)
			kds.Defs = append(kds.Defs, component)
		}
	}

	if len(kds.Defs) > 0 {
		kds.addGoImport(kds.ProtoGoPackage, "")
	}

	return kds
}

func visitProtoGoPackage(protoGoPackageCtx parser.IProtoGoPackageStatementContext) string {
	if protoGoPackageText, err := strconv.Unquote(protoGoPackageCtx.STR_LIT().GetText()); err == nil {
		return protoGoPackageText
	}
	return protoGoPackageCtx.STR_LIT().GetText()
}

func visitImport(importCtx parser.IImportStatementContext) string {
	if importText, err := strconv.Unquote(importCtx.STR_LIT().GetText()); err == nil {
		return importText
	}
	return importCtx.STR_LIT().GetText()
}

func visitEnum(ctx *Context, kds *Kds, enumCtx parser.IEnumDefContext) *Enum {
	enum := new(Enum)
	enum.ctx = ctx
	enum.kds = kds
	enum.Name = GoCamelCase(enumCtx.EnumName().GetText())
	for _, element := range enumCtx.EnumBody().AllEnumElement() {
		enum.EnumFields = append(enum.EnumFields, visitEnumField(ctx, kds, element.EnumField()))
	}
	return enum
}

func visitEnumField(ctx *Context, kds *Kds, enumFieldCtx parser.IEnumFieldContext) *EnumField {
	enumField := new(EnumField)
	enumField.ctx = ctx
	enumField.kds = kds
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
	entity.ctx = ctx
	entity.kds = kds
	entity.Name = GoCamelCase(entityCtx.EntityName().GetText())
	for _, element := range entityCtx.EntityBody().AllEntityElement() {
		switch {
		case element.Field() != nil:
			field := visitField(ctx, kds, element.Field())
			entity.Fields = append(entity.Fields, field)
		case element.MapField() != nil:
			field := visitMapField(ctx, kds, element.MapField())
			entity.Fields = append(entity.Fields, field)
		}
	}
	return entity
}

func visitComponent(ctx *Context, kds *Kds, componentCtx parser.IComponentDefContext) *Component {
	component := new(Component)
	component.ctx = ctx
	component.kds = kds
	component.Name = componentCtx.ComponentName().GetText()
	for _, element := range componentCtx.ComponentBody().AllComponentElement() {
		switch {
		case element.Field() != nil:
			field := visitField(ctx, kds, element.Field())
			component.Fields = append(component.Fields, field)
		case element.MapField() != nil:
			field := visitMapField(ctx, kds, element.MapField())
			component.Fields = append(component.Fields, field)
		}
	}
	return component
}

func visitField(ctx *Context, kds *Kds, fieldCtx parser.IFieldContext) *Field {
	field := new(Field)
	field.ctx = ctx
	field.kds = kds
	field.Repeated = fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil
	field.Type = visitType(kds, fieldCtx.Type_())

	field.Name = GoCamelCase(fieldCtx.FieldName().GetText())
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))

	if field.Repeated {
		field.ListType = ctx.addListType(field.Type)
	}
	return field
}

func visitMapField(ctx *Context, kds *Kds, mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	field.ctx = ctx
	field.kds = kds
	field.Map = true
	field.Type = visitType(kds, mapFieldCtx.Type_())

	field.KeyType = mapFieldCtx.KeyType().GetText()
	field.Name = GoCamelCase(mapFieldCtx.MapName().GetText())
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))

	field.MapType = ctx.addMapType(field.Type, field.KeyType)
	return field
}

func visitType(kds *Kds, typeCtx parser.IType_Context) string {
	if typeCtx.MessageType() != nil || typeCtx.EnumType() != nil {
		return GoCamelCase(typeCtx.GetText())
	}
	return typeCtx.GetText()
}
