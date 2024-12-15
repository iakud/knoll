package codegen

type Kds struct {
	Filename string
	Package string
	ProtoGoPackage string
	ProtoPackage string
	Imports []string

	ImportTimestamp bool
	ImportDuration bool
	ImportEmpty bool

	Defs []interface{}
}

type Enum struct {
	Name string
	EnumFields []*EnumField
	ProtoPackage string
}

type EnumField struct {
	Name string
	Value int
}

type Message struct {
	Name string
	Fields []*Field
	ProtoPackage string
}

type Entity struct {
	Message
}

type Component struct {
	Message
}

type FieldKind int32

const (
	FieldKind_Primitive FieldKind = iota
	FieldKind_Enum
	FieldKind_Entity
	FieldKind_Component
	FieldKind_Timestamp
	FieldKind_Duration
)

type Field struct {
	Repeated bool
	KeyType string
	Type string
	Name string
	Number int

	GoVarName string
	GoType string

	Kind string
}