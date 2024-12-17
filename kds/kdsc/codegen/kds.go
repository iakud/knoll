package codegen

import "slices"

type Kds struct {
	Filename string
	Package string
	ProtoGoPackage string
	ProtoPackage string
	Imports []string
	ProtoImports []string

	GoStandardImports []string
	GoImports []string

	Defs []interface{}
}

func (k *Kds) addProtoImport(protoImport string) {
	if slices.Contains(k.ProtoImports, protoImport) {
		return
	}
	k.ProtoImports = append(k.ProtoImports, protoImport)
}

func (k *Kds) addGoStandardImport(goStandardImport string) {
	if slices.Contains(k.GoStandardImports, goStandardImport) {
		return
	}
	k.GoStandardImports = append(k.GoStandardImports, goStandardImport)
}

func (k *Kds) addGoImport(goImport string) {
	if slices.Contains(k.GoImports, goImport) {
		return
	}
	k.GoImports = append(k.GoImports, goImport)
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

	ProtoType string

	GoVarName string
	GoType string

	Kind string
}