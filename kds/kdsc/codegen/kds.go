package codegen

import "slices"

type Kds struct {
	ctx        *Context
	Name       string
	SourceFile string
	Package    string
	Imports    []string

	Defs  []TopLevelDef
	Types []string

	Enums      []*Enum
	Entities   []*Entity
	Components []*Component

	FieldTypes []string
	ListTypes  []*ListType
	MapTypes   []*MapType
}

func (k *Kds) addType(name string) {
	if slices.Contains(k.Types, name) {
		return
	}
	k.Types = append(k.Types, name)
}

func (k *Kds) format() {
	slices.Sort(k.Types)
}

type TopLevelDef interface {
	GetName() string
	Kind() string
	GetFields() []*Field
}

type Enum struct {
	ctx        *Context
	kds        *Kds
	Name       string
	EnumFields []*EnumField
}

func (e *Enum) GetName() string {
	return e.Name
}

func (m *Enum) GetFields() []*Field {
	return nil
}

func (e *Enum) Kind() string {
	return "enum"
}

type EnumField struct {
	ctx   *Context
	kds   *Kds
	Name  string
	Value int
}

type Message struct {
	ctx    *Context
	kds    *Kds
	Name   string
	Fields []*Field
}

func (m *Message) GetName() string {
	return m.Name
}

func (m *Message) GetFields() []*Field {
	return m.Fields
}

type Entity struct {
	Message
}

func (e *Entity) Kind() string {
	return "entity"
}

type Component struct {
	Message
}

func (c *Component) Kind() string {
	return "component"
}

type Field struct {
	ctx      *Context
	kds      *Kds
	Repeated bool
	Map      bool
	KeyType  string
	Type     string
	Name     string
	Number   int

	GoVarName string
	ListType  string
	MapType   string
}

func (f *Field) TypeKind() string {
	if def, ok := f.ctx.Defs[f.Type]; ok {
		return def.Kind()
	}
	return f.Type
}

type ListType struct {
	ctx  *Context
	Name string
	Type string
}

func (l *ListType) TypeKind() string {
	if def, ok := l.ctx.Defs[l.Type]; ok {
		return def.Kind()
	}
	return l.Type
}

type MapType struct {
	ctx     *Context
	Name    string
	Type    string
	KeyType string
}

func (m *MapType) TypeKind() string {
	if def, ok := m.ctx.Defs[m.Type]; ok {
		return def.Kind()
	}
	return m.Type
}

func (m *MapType) KeyTypeKind() string {
	if def, ok := m.ctx.Defs[m.KeyType]; ok {
		return def.Kind()
	}
	return m.KeyType
}
