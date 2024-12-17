package codegen

import (
	"slices"
	"strings"
)

type Context struct {
	KdsContexts []*Kds

	CommonTypes []string
	// type
	TypeList map[string]*CommonType
	// type -> keys
	TypeMap map[string][]*CommonMapType

	Imports map[string]*Kds
	Defs map[string]interface{}
}

type CommonType struct {
	Type string
	GoType string
}

type CommonMapType struct {
	Type string
	KeyType string
	GoType string
	GoKeyType string
}

func newContext() *Context {
	return &Context{
		TypeList: make(map[string]*CommonType),
		TypeMap: make(map[string][]*CommonMapType),
		Imports: make(map[string]*Kds),
		Defs: make(map[string]interface{}),
	}
}

func (ctx *Context) GetCommonTypes() []string {
	return ctx.CommonTypes
}

func (ctx *Context) AddListType(name string, customType bool) {
	if _, ok := ctx.TypeList[name]; ok {
		return
	}
	commonType := new(CommonType)
	commonType.Type = name
	commonType.GoType = GoType(name)
	ctx.TypeList[name] = commonType

	if customType {
		return
	}
	// common
	if slices.Contains(ctx.CommonTypes, name) {
		return
	}
	ctx.CommonTypes = append(ctx.CommonTypes, name)
	slices.Sort(ctx.CommonTypes)
	slices.SortFunc(ctx.CommonTypes, func(a, b string) int {
		return strings.Compare(a, b)
	})
}

func (ctx *Context) AddMapType(name string, keyType string, customType bool) {
	mapTypes, _ := ctx.TypeMap[name]
	if slices.ContainsFunc(mapTypes, func(mapType *CommonMapType) bool {
		return mapType.KeyType == keyType
	}) {
		return
	}
	
	mapType := new(CommonMapType)
	mapType.Type = name
	mapType.KeyType = keyType
	mapType.GoType = GoType(name)
	mapType.GoKeyType = GoType(keyType)

	mapTypes = append(mapTypes, mapType)
	slices.SortFunc(mapTypes, func(a, b *CommonMapType) int {
		return strings.Compare(a.KeyType, b.KeyType)
	})
	ctx.TypeMap[name] = mapTypes
	
	if customType {
		return
	}
	// common
	if slices.Contains(ctx.CommonTypes, name) {
		return
	}
	ctx.CommonTypes = append(ctx.CommonTypes, name)
	slices.Sort(ctx.CommonTypes)
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

func (ctx *Context) FindList(name string) *CommonType {
	if commonType, ok := ctx.TypeList[name]; ok {
		return commonType
	}
	return nil
}

func (ctx *Context) FindMap(name string) []*CommonMapType {
	if mapTypes, ok := ctx.TypeMap[name]; ok {
		return mapTypes
	}
	return nil
}