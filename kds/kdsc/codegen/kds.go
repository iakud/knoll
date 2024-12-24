package codegen

import "slices"

type Kds struct {
	ctx            *Context
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
	GoProtoPackage() string
	Kind() string
}

type Enum struct {
	ctx        *Context
	kds        *Kds
	Name       string
	EnumFields []*EnumField
}

func (m *Enum) format(ctx *Context) {
	for _, field := range m.EnumFields {
		field.format(ctx)
	}
}

func (e *Enum) GetName() string {
	return e.Name
}

func (e *Enum) GoProtoPackage() string {
	return e.kds.ProtoPackage
}

func (e *Enum) GoType() string {
	return e.Name
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

func (f *EnumField) format(ctx *Context) {

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

func (m *Message) GoProtoPackage() string {
	return m.kds.ProtoPackage
}

func (m *Message) GoType() string {
	return "*" + m.Name
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

func (f *Field) GoType() string {
	if _, ok := f.ctx.Defs[f.Type]; ok {
		return f.Type
	}
	return GoType(f.Type)
}

func (f *Field) GoProtoType() string {
	if _, ok := f.ctx.Defs[f.Type]; ok {
		return f.kds.ProtoPackage + "." + f.Type
	}
	return GoProtoType(f.Type)
}

func (f *Field) TypeKind() string {
	if def, ok := f.ctx.Defs[f.Type]; ok {
		return def.Kind()
	}
	return Kind(f.Type)
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
	return Kind(l.Type)
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
	return Kind(m.Type)
}
