package codegen

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iakud/krocher/kds/kdsc/parser"
)

func visitKds(ctx *Context, filePath string, kdsCtx parser.IKdsContext) *Kds {
	kds := new(Kds)
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

func visitEnum(kds *Kds, enumCtx parser.IEnumDefContext) *Enum {
	enum := new(Enum)
	enum.Name = GoCamelCase(enumCtx.EnumName().GetText())
	for _, element := range enumCtx.EnumBody().AllEnumElement() {
		enum.EnumFields = append(enum.EnumFields, visitEnumField(element.EnumField()))
	}
	enum.ProtoPackage = kds.ProtoPackage
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
		case element.MapField() != nil:
			field := visitMapField(ctx, kds, element.MapField())
			entity.Fields = append(entity.Fields, field)
		}
	}
	entity.ProtoPackage = kds.ProtoPackage
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
		case element.MapField() != nil:
			field := visitMapField(ctx, kds, element.MapField())
			component.Fields = append(component.Fields, field)
		}
	}
	component.ProtoPackage = kds.ProtoPackage
	return component
}

func visitField(ctx *Context, kds *Kds, fieldCtx parser.IFieldContext) *Field {
	field := new(Field)
	field.Type = visitType(ctx, kds, fieldCtx.Type_())

	field.Name = GoCamelCase(fieldCtx.FieldName().GetText())
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))

	if fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil {
		field.Repeated = true
		field.ListType = ctx.addListType(field.Type)
	} else {
		switch {
		case fieldCtx.Type_().TIMESTAMP() != nil:
			kds.addGoTypes(fieldCtx.Type_().TIMESTAMP().GetText())
		case fieldCtx.Type_().DURATION() != nil:
			kds.addGoTypes(fieldCtx.Type_().DURATION().GetText())
		case fieldCtx.Type_().EMPTY() != nil:
			kds.addGoTypes(fieldCtx.Type_().EMPTY().GetText())
		}
	}
	return field
}

func visitMapField(ctx *Context, kds *Kds, mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	field.Type = visitType(ctx, kds, mapFieldCtx.Type_())

	field.KeyType = mapFieldCtx.KeyType().GetText()
	field.Name = GoCamelCase(mapFieldCtx.MapName().GetText())
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))

	field.MapType = ctx.addMapType(field.Type, field.KeyType)
	return field
}

func visitType(ctx *Context, kds *Kds, typeCtx parser.IType_Context) string {
	var type_ string
	customType := typeCtx.MessageType() != nil || typeCtx.EnumType() != nil
	if customType {
		type_ = GoCamelCase(typeCtx.GetText())
	} else {
		type_ = typeCtx.GetText()
	}
	switch {
	case typeCtx.TIMESTAMP() != nil:
		kds.addProtoTypes(typeCtx.TIMESTAMP().GetText())
	case typeCtx.DURATION() != nil:
		kds.addProtoTypes(typeCtx.DURATION().GetText())
	case typeCtx.EMPTY() != nil:
		kds.addProtoTypes(typeCtx.EMPTY().GetText())
	}

	return type_
}
