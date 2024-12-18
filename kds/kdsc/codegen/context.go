package codegen

import (
	"slices"
	"strings"
)

type Context struct {
	KdsContexts []*Kds

	Common Kds

	// type
	TypeList map[string]*ListType
	// type -> keys
	TypeMap map[string][]*MapType

	Imports map[string]*Kds
	Defs map[string]interface{}
}

type ListType struct {
	Type string
}

type MapType struct {
	Type string
	KeyType string
}

func newContext() *Context {
	return &Context{
		TypeList: make(map[string]*ListType),
		TypeMap: make(map[string][]*MapType),
		Imports: make(map[string]*Kds),
		Defs: make(map[string]interface{}),
	}
}

func (ctx *Context) format() {
	for _, mapTypes := range ctx.TypeMap {
		slices.SortFunc(mapTypes, func(a, b *MapType) int {
			return strings.Compare(a.KeyType, b.KeyType)
		})
	}
}

func (ctx *Context) addListType(name string) {
	if _, ok := ctx.TypeList[name]; ok {
		return
	}
	listType := new(ListType)
	listType.Type = name
	ctx.TypeList[name] = listType
}

func (ctx *Context) addMapType(name string, keyType string) {
	mapTypes, _ := ctx.TypeMap[name]
	if slices.ContainsFunc(mapTypes, func(mapType *MapType) bool {
		return mapType.KeyType == keyType
	}) {
		return
	}

	mapType := new(MapType)
	mapType.Type = name
	mapType.KeyType = keyType
	ctx.TypeMap[name] = append(mapTypes, mapType)
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

func (ctx *Context) FindList(name string) *ListType {
	if listType, ok := ctx.TypeList[name]; ok {
		return listType
	}
	return nil
}

func (ctx *Context) FindMap(name string) []*MapType {
	if mapTypes, ok := ctx.TypeMap[name]; ok {
		return mapTypes
	}
	return nil
}