package codegen

import "slices"

type Context struct {
	KdsContexts []*Kds

	CommonTypes []string
	// type
	TypeList map[string]struct{}
	// type -> keys
	TypeMap map[string][]string

	Common *Kds
	Imports map[string]*Kds
	Defs map[string]interface{}
}

func newContext() *Context {
	return &Context{
		TypeList: make(map[string]struct{}),
		TypeMap: make(map[string][]string),
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
	ctx.TypeList[name] = struct{}{}

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

func (ctx *Context) AddMapType(name string, key string, customType bool) {
	keys, _ := ctx.TypeMap[name]
	if slices.Contains(keys, key) {
		return
	}
	keys = append(keys, key)
	ctx.TypeMap[name] = keys
	slices.Sort(keys)
	
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

func (ctx *Context) FindList(name string) bool {
	if _, ok := ctx.TypeList[name]; ok {
		return true
	}
	return false
}

func (ctx *Context) FindMap(name string) []string {
	if keys, ok := ctx.TypeMap[name]; ok {
		slices.Sort(keys)
		return keys
	}
	return nil
}