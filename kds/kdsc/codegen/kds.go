package codegen

type Kds struct {
	Filename string
	Package string
	Imports []string
	Enums []*Enum
	Entities []*Entity
	Components []*Component

	Defs []interface{}
}

func (k *Kds) IsImportTimestamp() bool {
	return false
}

type Enum struct {
	Name string
	EnumFields []*EnumField
}

type EnumField struct {
	Name string
	Value int
}

type Entity struct {
	Name string
	Fields []*Field
}

type Component struct {
	Name string
	Fields []*Field
}

type FieldKind int32

const (
	FieldKind_Primitive FieldKind = iota
	FieldKind_Enum
	FieldKind_Entity
	FieldKind_Component
)

type Field struct {
	Repeated bool
	Type string
	KeyType string
	ValueType string
	Name string
	Number int
	Kind string
}

func (f *Field) IsComponent() bool {
	switch f.Type {
	case "double", "float", "int32", "int64", "sint32", "sint64", "fixed32", "fixed64", "bool", "string", "bytes":
		break
	default:
		return true
	}
	return false
}

func (f *Field) IsEnum() bool {
	return false
}