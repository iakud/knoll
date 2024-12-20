package codegen

import "slices"

type Kds struct {
	Name           string
	SourceFile     string
	Package        string
	ProtoGoPackage string
	ProtoPackage   string
	Imports        []string
	ProtoImports   []string
	GoImportSpecs  []*ImportSpec

	Defs  []TopLevelDef
	Types []string

	GoTypes    []string
	ProtoTypes []string
}

func (k *Kds) addGoTypes(type_ string) {
	if slices.Contains(k.GoTypes, type_) {
		return
	}
	k.GoTypes = append(k.GoTypes, type_)
}

func (k *Kds) addProtoTypes(type_ string) {
	if slices.Contains(k.ProtoTypes, type_) {
		return
	}
	k.ProtoTypes = append(k.ProtoTypes, type_)
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

func (k *Kds) addGoImportByType(type_ string) {
	switch type_ {
	case "timestamp":
		k.addGoImport("time", "")
		k.addGoImport("google.golang.org/protobuf/types/known/timestamppb", "")
	case "duration":
		k.addGoImport("time", "")
		k.addGoImport("google.golang.org/protobuf/types/known/durationpb", "")
	case "empty":
		k.addGoImport("google.golang.org/protobuf/types/known/emptypb", "")
	}
}

func (k *Kds) addProtoImportByType(type_ string) {
	switch type_ {
	case "timestamp":
		k.addProtoImport("google/protobuf/timestamp.proto")
	case "duration":
		k.addProtoImport("google/protobuf/duration.proto")
	case "empty":
		k.addProtoImport("google/protobuf/empty.proto")
	}
}

func (k *Kds) addType(name string) {
	if slices.Contains(k.Types, name) {
		return
	}
	k.Types = append(k.Types, name)
}

func (k *Kds) format() {
	for _, type_ := range k.GoTypes {
		k.addGoImportByType(type_)
	}
	for _, type_ := range k.ProtoTypes {
		k.addProtoImportByType(type_)
	}

	for _, type_ := range k.Types {
		k.addGoImportByType(type_)
		k.addProtoImportByType(type_)
	}

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
	Path         string
	Name         string
	SpacesBefore bool
}

type TopLevelDef interface {
	GetName() string
	ProtoGoType() string
	GoType() string
}

type Enum struct {
	Name         string
	EnumFields   []*EnumField
	ProtoPackage string
}

func (e *Enum) GetName() string {
	return e.Name
}

func (e *Enum) ProtoGoType() string {
	return e.ProtoPackage + "." + e.Name
}

func (e *Enum) GoType() string {
	return e.Name
}

type EnumField struct {
	Name  string
	Value int
}

type Message struct {
	Name         string
	Fields       []*Field
	ProtoPackage string
}

func (m *Message) GetName() string {
	return m.Name
}

func (m *Message) ProtoGoType() string {
	return "*" + m.ProtoPackage + "." + m.Name
}

func (m *Message) GoType() string {
	return "*" + m.Name
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
	KeyType  string
	Type     string
	Name     string
	Number   int

	GoVarName string
	ListType  string
	MapType   string

	Kind string
}
