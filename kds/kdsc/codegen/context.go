package codegen

type Context struct {
	KdsContexts []*Kds

	// type
	TypeArray map[string]struct{}
	// type -> keys
	TypeMap map[string][]string

	Common *Kds
	Imports map[string]*Kds
	Defs map[string]interface{}
}

func newContext() *Context {
	return &Context{
		TypeArray: make(map[string]struct{}),
		TypeMap: make(map[string][]string),
		Imports: make(map[string]*Kds),
		Defs: make(map[string]interface{}),
	}
}

func (ctx *Context) AddArray(name string) {
	ctx.TypeArray[name] = struct{}{}
}

func (ctx *Context) AddMap(name string, key string) {
	existKeys, ok := ctx.TypeMap[name]
	if !ok {
		ctx.TypeMap[name] = append([]string(nil), key)
		return
	}
	for _, existKey := range existKeys {
		if existKey == key {
			return
		}
	}
	ctx.TypeMap[name] = append(existKeys, key)
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

func (ctx *Context) FindArray(name string) bool {
	if _, ok := ctx.TypeArray[name]; ok {
		return true
	}
	return false
}

func (ctx *Context) FindMap(name string) []string {
	if keys, ok := ctx.TypeMap[name]; ok {
		return keys
	}
	return nil
}