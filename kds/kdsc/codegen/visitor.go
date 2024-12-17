package codegen

import (
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/iakud/krocher/kds/kdsc/parser"
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
		kds.addGoImport(kds.ProtoGoPackage)
	}

	slices.Sort(kds.ProtoImports)
	slices.Sort(kds.GoStandardImports)
	slices.Sort(kds.GoImports)
	return kds
}

func visitProtoGoPackage(protoGoPackageCtx parser.IProtoGoPackageStatementContext) string {
	protoGoPackageText := protoGoPackageCtx.STR_LIT().GetText()
	switch {
	case strings.HasPrefix(protoGoPackageText, "\"") && strings.HasSuffix(protoGoPackageText, "\"") :
		return strings.TrimSuffix(strings.TrimPrefix(protoGoPackageText, "\""), "\"")
	case strings.HasPrefix(protoGoPackageText, "'") && strings.HasSuffix(protoGoPackageText, "'") :
		return strings.TrimSuffix(strings.TrimPrefix(protoGoPackageText, "'"), "'")
	}
	return protoGoPackageText
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
	customType := visitType(ctx, kds, field, fieldCtx.Type_())
	field.Name = GoCamelCase(fieldCtx.FieldName().GetText())
	field.Number, _ = strconv.Atoi(fieldCtx.FieldNumber().GetText())

	field.ProtoType = ProtoType(field.Type)
	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))
	field.GoType = GoType(field.Type)

	if fieldCtx.FieldLabel() != nil && fieldCtx.FieldLabel().REPEATED() != nil {
		field.Repeated = true
		ctx.AddListType(field.Type, customType)
	}
	return field
}

func visitMapField(ctx *Context, kds *Kds, mapFieldCtx parser.IMapFieldContext) *Field {
	field := new(Field)
	customType := visitType(ctx, kds, field, mapFieldCtx.Type_())
	field.KeyType = mapFieldCtx.KeyType().GetText()
	field.Name = GoCamelCase(mapFieldCtx.MapName().GetText())
	field.Number, _ = strconv.Atoi(mapFieldCtx.FieldNumber().GetText())

	field.ProtoType = ProtoType(field.Type)
	field.GoVarName = GoSanitized(ToLowerFirst(field.Name))
	field.GoType = GoType(field.Type)

	ctx.AddMapType(field.Type, field.KeyType, customType)
	return field
}

func visitType(ctx *Context, kds *Kds, field *Field, typeCtx parser.IType_Context) bool {
	customType := typeCtx.MessageType() != nil || typeCtx.EnumType() != nil
	if customType {
		field.Type = GoCamelCase(typeCtx.GetText())
	} else {
		field.Type = typeCtx.GetText()	
	}
	switch {
	case typeCtx.TIMESTAMP() != nil:
		kds.addGoStandardImport("time")
		kds.addGoImport("google.golang.org/protobuf/types/known/timestamppb")
		kds.addProtoImport("google/protobuf/timestamp.proto")
	case typeCtx.DURATION() != nil:
		kds.addGoStandardImport("time")
		kds.addGoImport("google.golang.org/protobuf/types/known/durationpb")
		kds.addProtoImport("google/protobuf/duration.proto")
	case typeCtx.EMPTY() != nil:
		kds.addGoImport("google.golang.org/protobuf/types/known/emptypb")
		kds.addProtoImport("google/protobuf/empty.proto")
	}
	return customType
}