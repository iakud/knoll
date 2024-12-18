package codegen

import "slices"

type Kds struct {
	Name string
	Filename string
	Package string
	ProtoGoPackage string
	ProtoPackage string
	Imports []string
	ProtoImports []string
	GoImportSpecs []*ImportSpec

	Defs []interface{}

	Types []string
}

func (k *Kds) addProtoImport(path string) {
	if slices.Contains(k.ProtoImports, path) {
		return
	}
	k.ProtoImports = append(k.ProtoImports, path)
}

func (k *Kds) addGoImport(path, name string) {
	if slices.ContainsFunc(k.GoImportSpecs, func(spec *ImportSpec) bool {
		return spec.Path == path && spec.Name == name
	}) {
		return
	}
	spec := &ImportSpec{Path: path, Name: name}
	k.GoImportSpecs = append(k.GoImportSpecs, spec)
}

func (k *Kds) addType(name string) {
	if slices.Contains(k.Types, name) {
		return
	}
	k.Types = append(k.Types, name)
}

func (k *Kds) format() {
	k.sortImports()
	slices.Sort(k.Types)
}

func (k *Kds) sortImports() {
	// sort proto imports
	slices.Sort(k.ProtoImports)
	// sort go imports
	localPrefix := ""
	sortImports(localPrefix, k.GoImportSpecs)

	lastGroup := -1
	for _, importSpec := range k.GoImportSpecs {
		groupNum := importGroup(localPrefix, importSpec.Path)
		if groupNum != lastGroup && lastGroup != -1 {
			importSpec.SpacesBefore = true
		}
		lastGroup = groupNum
	}
}

type ImportSpec struct {
	Path string
	Name string
	SpacesBefore bool
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

	Kind string
}