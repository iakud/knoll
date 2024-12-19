package codegen

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iakud/krocher/kds/kdsc/parser"
)

func visitKds(ctx *Context, name string, kdsCtx parser.IKdsContext) *Kds {
	var kds *Kds
	if name == "common" {
		kds = &ctx.Common
	} else {
		kds = new(Kds)
	}

	kds.Name = name
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
	customType := fieldCtx.Type_().MessageType() != nil || fieldCtx.Type_().EnumType() != nil
	if customType {
		field.Type = GoCamelCase(fieldCtx.Type_().GetText())
	} else {
		field.Type = fieldCtx.Type_().GetText()	
	}
	switch {
	case fieldCtx.Type_().TIMESTAMP() != nil:
		kds.addGoImport("time", "")
		kds.addGoImport("google.golang.org/protobuf/types/known/timestamppb", "")
		kds.addProtoImport("google/protobuf/timestamp.proto")
	case fieldCtx.Type_().DURATION() != nil:
		kds.addGoImport("time", "")
		kds.addGoImport("google.golang.org/protobuf/types/known/durationpb", "")
		kds.addProtoImport("google/protobuf/duration.proto")
	case fieldCtx.Type_().EMPTY() != nil:
		kds.addGoImport("google.golang.org/protobuf/types/known/emptypb", "")
		kds.addProtoImport("google/protobuf/empty.proto")
	}

	field.Name = GoCamelCase(fieldCtx.FieldName().GetText())
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))

	if fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil {
		field.Repeated = true
		ctx.addListType(field.Type)
		
		if customType {
			kds.addGoImport("slices", "")
			kds.addGoImport("iter", "")
		} else {
			ctx.Common.addType(field.Type)
			ctx.Common.addGoImport("slices", "")
			ctx.Common.addGoImport("iter", "")

			switch {
			case fieldCtx.Type_().TIMESTAMP() != nil:
				ctx.Common.addGoImport("time", "")
				ctx.Common.addGoImport("google.golang.org/protobuf/types/known/timestamppb", "")
				// ctx.Common.addProtoImport("google/protobuf/timestamp.proto")
			case fieldCtx.Type_().DURATION() != nil:
				ctx.Common.addGoImport("time", "")
				ctx.Common.addGoImport("google.golang.org/protobuf/types/known/durationpb", "")
				// ctx.Common.addProtoImport("google/protobuf/duration.proto")
			case fieldCtx.Type_().EMPTY() != nil:
				ctx.Common.addGoImport("google.golang.org/protobuf/types/known/emptypb", "")
				// ctx.Common.addProtoImport("google/protobuf/empty.proto")
			}
		}
	}
	return field
}

func visitMapField(ctx *Context, kds *Kds, mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	customType := mapFieldCtx.Type_().MessageType() != nil || mapFieldCtx.Type_().EnumType() != nil
	if customType {
		field.Type = GoCamelCase(mapFieldCtx.Type_().GetText())
	} else {
		field.Type = mapFieldCtx.Type_().GetText()	
	}
	switch {
	case mapFieldCtx.Type_().TIMESTAMP() != nil:
		kds.addGoImport("time", "")
		kds.addGoImport("google.golang.org/protobuf/types/known/timestamppb", "")
		kds.addProtoImport("google/protobuf/timestamp.proto")
	case mapFieldCtx.Type_().DURATION() != nil:
		kds.addGoImport("time", "")
		kds.addGoImport("google.golang.org/protobuf/types/known/durationpb", "")
		kds.addProtoImport("google/protobuf/duration.proto")
	case mapFieldCtx.Type_().EMPTY() != nil:
		kds.addGoImport("google.golang.org/protobuf/types/known/emptypb", "")
		kds.addProtoImport("google/protobuf/empty.proto")
	}

	field.KeyType = mapFieldCtx.KeyType().GetText()
	field.Name = GoCamelCase(mapFieldCtx.MapName().GetText())
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())

	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))

	ctx.addMapType(field.Type, field.KeyType)

	if customType {
		kds.addGoImport("maps", "")
		kds.addGoImport("iter", "")
	} else {
		ctx.Common.addType(field.Type)
		ctx.Common.addGoImport("maps", "")
		ctx.Common.addGoImport("iter", "")

		switch {
		case mapFieldCtx.Type_().TIMESTAMP() != nil:
			ctx.Common.addGoImport("time", "")
			ctx.Common.addGoImport("google.golang.org/protobuf/types/known/timestamppb", "")
			// ctx.Common.addProtoImport("google/protobuf/timestamp.proto")
		case mapFieldCtx.Type_().DURATION() != nil:
			ctx.Common.addGoImport("time", "")
			ctx.Common.addGoImport("google.golang.org/protobuf/types/known/durationpb", "")
			// ctx.Common.addProtoImport("google/protobuf/duration.proto")
		case mapFieldCtx.Type_().EMPTY() != nil:
			ctx.Common.addGoImport("google.golang.org/protobuf/types/known/emptypb", "")
			// ctx.Common.addProtoImport("google/protobuf/empty.proto")
		}

	}
	return field
}