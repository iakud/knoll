{{- /* BEGIN DEFINE */ -}}

{{- define "CommonList"}}

type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{.Name}} struct {
	syncable []{{toGoType .Type}}

	dirtyParent dirtyParentFunc_{{.Name}}
}

func (x *{{.Name}}) Len() int {
	return len(x.syncable)
}

func (x *{{.Name}}) Get(i int) {{toGoType .Type}} {
	return x.syncable[i]
}

func (x *{{.Name}}) Set(i int, v {{toGoType .Type}}) {
	x.syncable[i] = v
}

func (x *{{.Name}}) Append(v ...{{toGoType .Type}}) {
	x.syncable = append(x.syncable, v...)
}

func (x *{{.Name}}) Insert(i int, v ...{{toGoType .Type}}) {
	x.syncable = slices.Insert(x.syncable, i, v...)
}

func (x *{{.Name}}) Delete(i, j int) {
	x.syncable = slices.Delete(x.syncable, i, j)
}

func (x *{{.Name}}) Replace(i, j int, v ...{{toGoType .Type}}) {
	x.syncable = slices.Replace(x.syncable, i, j, v...)
}

func (x *{{.Name}}) Reverse() {
	slices.Reverse(x.syncable)
}

func (x *{{.Name}}) All() iter.Seq2[int, {{toGoType .Type}}] {
	return slices.All(x.syncable)
}

func (x *{{.Name}}) Backward() iter.Seq2[int, {{toGoType .Type}}] {
	return slices.Backward(x.syncable)
}

func (x *{{.Name}}) Values() iter.Seq[{{toGoType .Type}}] {
	return slices.Values(x.syncable)
}
{{- end}}

{{- define "CommonMap"}}

type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{.Name}} struct {
	syncable map[{{toGoType .KeyType}}]{{toGoType .Type}}

	dirtyParent dirtyParentFunc_{{.Name}}
}

func (x *{{.Name}}) Len() int {
	return len(x.syncable)
}

func (x *{{.Name}}) Clear() {
	clear(x.syncable)
}

func (x *{{.Name}}) Get(k {{toGoType .KeyType}}) ({{toGoType .Type}}, bool) {
	v, ok := x.syncable[k]
	return v, ok
}

func (x *{{.Name}}) Set(k {{toGoType .KeyType}}, v {{toGoType .Type}}) {
	x.syncable[k] = v
}

func (x *{{.Name}}) Delete(k {{toGoType .KeyType}}) {
	delete(x.syncable, k)
}

func (x *{{.Name}}) All() iter.Seq2[{{toGoType .KeyType}}, {{toGoType .Type}}] {
	return maps.All(x.syncable)
}

func (x *{{.Name}}) Keys() iter.Seq[{{toGoType .KeyType}}] {
	return maps.Keys(x.syncable)
}

func (x *{{.Name}}) Values() iter.Seq[{{toGoType .Type}}] {
	return maps.Values(x.syncable)
}
{{- end}}

{{- define "Enum"}}
{{- $EnumType := .Name}}

type {{.Name}} = {{.ProtoPackage}}.{{.Name}}

const (
{{- range .EnumFields}}
	{{$EnumType}}_{{.Name}} {{$EnumType}} = {{.Value}}
{{- end}}
)
{{- end}}

{{- define "Syncable"}}

type syncable{{.Name}} struct {
{{- range .Fields}}
{{- if .Repeated}}{{/* Array */}}
{{- if findComponent .Type}}
	{{.Name}} []*{{.Type}}
{{- else if findEnum .Type}}
	{{.Name}} []{{.Type}}
{{- else}}
	{{.Name}} []{{toGoType .Type}}
{{- end}}
{{- else if len .KeyType}}{{/* Map */}}
{{- if findComponent .Type}}
	{{.Name}} map[{{toGoType .KeyType}}]*{{.Type}}
{{- else if findEnum .Type}}
	{{.Name}} map[{{toGoType .KeyType}}]{{.Type}}
{{- else}}
	{{.Name}} map[{{toGoType .KeyType}}]{{toGoType .Type}}
{{- end}}
{{- else}}{{/* Field */}}
{{- if findComponent .Type}}
	{{.Name}} *{{.Type}}
{{- else if findEnum .Type}}
	{{.Name}} {{.Type}}
{{- else}}
	{{.Name}} {{toGoType .Type}}
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}

{{- define "Message"}}
{{- $MessageName := .Name}}

{{- range .Fields}}
{{- if .Repeated}}

{{- else if len .KeyType}}

{{- else}}
{{- if findComponent .Type}}

func (x *{{$MessageName}}) Get{{.Name}}() *{{.Type}} {
	return x.syncable.{{.Name}}
}

func (x *{{$MessageName}}) set{{.Name}}(v *{{.Type}}) {
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed or evicted from its original place first")
	}
	if v == x.syncable.{{.Name}} {
		return
	}
	if x.syncable.{{.Name}} != nil {
		x.syncable.{{.Name}}.dirtyParent = nil
	}
	x.syncable.{{.Name}} = v
	v.dirtyParent = func() {
		x.markDirty(uint64(0x01) << {{.Number}})
	}
	x.markDirty(uint64(0x01) << {{.Number}})
	if v != nil {
		v.markDirty(uint64(0x01))
	}
}
{{- else if findEnum .Type}}

func (x *{{$MessageName}}) Get{{.Name}}() {{.Type}} {
	return x.syncable.{{.Name}}
}

func (x *{{$MessageName}}) Set{{.Name}}(v {{.Type}}) {
	if v == x.syncable.{{.Name}} {
		return
	}
	x.syncable.{{.Name}} = v
	x.markDirty(uint64(0x01) << {{.Number}})
}
{{- else}}

func (x *{{$MessageName}}) Get{{.Name}}() {{toGoType .Type}} {
	return x.syncable.{{.Name}}
}

func (x *{{$MessageName}}) Set{{.Name}}(v {{toGoType .Type}}) {
{{- if eq .Type "bytes"}}
	if v != nil || x.syncable.{{.Name}} != nil {
		return
	}
{{- else}}
	if v == x.syncable.{{.Name}} {
		return
	}
{{- end}}
	x.syncable.{{.Name}} = v
	x.markDirty(uint64(0x01) << {{.Number}})
}
{{- end}}
{{- end}}
{{- end}}

func (x *{{$MessageName}}) DumpChange() *{{.ProtoPackage}}.{{.Name}} {
	if x.checkDirty(uint64(0x01)) {
		return x.DumpFull()
	}
	m := new({{.ProtoPackage}}.{{.Name}})
{{- range .Fields}}
	if x.checkDirty(uint64(0x01) << {{.Number}}) {
{{- if .Repeated}}
		
{{- if findComponent .Type}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, v.DumpChange())
		}
{{- else if eq .Type "timestamp"}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, timestamppb.New(v))
		}
{{- else if eq .Type "duration"}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, durationpb.New(v))
		}
{{- else if eq .Type "empty"}}
		for range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, new(emptypb.Empty))
		}
{{- else}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, v)
		}
{{- end}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
		for k, v := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = v.DumpChange()
		}
{{- else if eq .Type "timestamp"}}
		for k, v := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = timestamppb.New(v)
		}
{{- else if eq .Type "duration"}}
		for k, v := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = durationpb.New(v)
		}
{{- else if eq .Type "empty"}}
		for k := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = new(emptypb.Empty)
		}
{{- else}}
		for k, v := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = v
		}
{{- end}}
{{- else}}
{{- if findComponent .Type}}
		m.{{.Name}} = x.syncable.{{.Name}}.DumpChange()
{{- else if eq .Type "timestamp"}}
		m.{{.Name}} = timestamppb.New(x.syncable.{{.Name}})
{{- else if eq .Type "duration"}}
		m.{{.Name}} = durationpb.New(x.syncable.{{.Name}})
{{- else if eq .Type "empty"}}
		m.{{.Name}} = new(emptypb.Empty)
{{- else}}
		m.{{.Name}} = x.syncable.{{.Name}}
{{- end}}
{{- end}}
	}
{{- end}}
	return m
}

func (x *{{$MessageName}}) DumpFull() *{{.ProtoPackage}}.{{.Name}} {
	m := new({{.ProtoPackage}}.{{.Name}})
{{- range .Fields}}
{{- if .Repeated}}
{{- if findComponent .Type}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, v.DumpFull())
	}
{{- else if eq .Type "timestamp"}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, timestamppb.New(v))
	}
{{- else if eq .Type "duration"}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, durationpb.New(v))
	}
{{- else if eq .Type "empty"}}
	for range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, new(emptypb.Empty))
	}
{{- else}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, v)
	}
{{- end}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	for k, v := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = v.DumpFull()
	}
{{- else if eq .Type "timestamp"}}
	for k, v := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = timestamppb.New(v)
	}
{{- else if eq .Type "duration"}}
	for k, v := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = durationpb.New(v)
	}
{{- else if eq .Type "empty"}}
	for k := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = new(emptypb.Empty)
	}
{{- else}}
	for k, v := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = v
	}
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	m.{{.Name}} = x.syncable.{{.Name}}.DumpFull()
{{- else if eq .Type "timestamp"}}
	m.{{.Name}} = timestamppb.New(x.syncable.{{.Name}})
{{- else if eq .Type "duration"}}
	m.{{.Name}} = durationpb.New(x.syncable.{{.Name}})
{{- else if eq .Type "empty"}}
	m.{{.Name}} = new(emptypb.Empty)
{{- else}}
	m.{{.Name}} = x.syncable.{{.Name}}
{{- end}}
{{- end}}
{{- end}}
	return m
}

{{- end}}

{{- define "Entity"}}
{{- template "Syncable" .}}

type {{.Name}} struct {
	id int64
	syncable syncable{{.Name}}

	dirty uint64
}

func New{{.Name}}() *{{.Name}} {
	x := new({{.Name}})
	x.dirty = 1
	x.id = 0 // FIXME: gen nextId()
{{- range .Fields}}
{{- if .Repeated}}{{/* nothing to do*/}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	x.syncable.{{.Name}} = make(map[{{toGoType .KeyType}}]*{{.Type}})
{{- else if findEnum .Type}}
	x.syncable.{{.Name}} = make(map[{{toGoType .KeyType}}]{{.Type}})
{{- else}}
	x.syncable.{{.Name}} = make(map[{{toGoType .KeyType}}]{{toGoType .Type}})
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.set{{.Name}}(New{{.Type}}())
{{- end}}
{{- end}}
{{- end}}
	return x
}

func (x *{{.Name}}) Id() int64 {
	return x.id
}
{{- template "Message" .}}

func (x *{{.Name}}) markAll() {
	x.dirty = uint64(0x01)
}

func (x *{{.Name}}) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
}

func (x *{{.Name}}) clearAll() {
{{- range .Fields}}
{{- if findComponent .Type}}
{{- if .Repeated}}
	for i := 0; i < len(x.syncable.{{.Name}}); i++ {
		x.syncable.{{.Name}}[i].clearDirty()
	}
{{- else if len .KeyType}}
	for _, v := range x.syncable.{{.Name}} {
		v.clearDirty()
	}
{{- else}}
	x.syncable.{{.Name}}.clearDirty()
{{- end}}
{{- end}}
{{- end}}
	x.dirty = 0
}

func (x *{{.Name}}) clearDirty() {
	if x.dirty == 0 {
		return
	}
	if x.dirty & uint64(0x01) != 0 {
		x.clearAll()
		return
	}
{{- range .Fields}}
{{- if findComponent .Type}}
	if x.dirty & uint64(0x01) << {{.Number}} != 0 {
{{- if .Repeated}}
		for i := 0; i < len(x.syncable.{{.Name}}); i++ {
			x.syncable.{{.Name}}[i].clearDirty()
		}
{{- else if len .KeyType}}
		for _, v := range x.syncable.{{.Name}} {
			v.clearDirty()
		}
{{- else}}
		x.syncable.{{.Name}}.clearDirty()
{{- end}}
	}
{{- end}}
{{- end}}
	x.dirty = 0
}

func (x *{{.Name}}) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}

{{- end}}

{{- define "Component"}}
{{- template "Syncable" .}}

type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{.Name}} struct {
	syncable syncable{{.Name}}

	dirty uint64
	dirtyParent dirtyParentFunc_{{.Name}}
}

func New{{.Name}}() *{{.Name}} {
	x := new({{.Name}})
	x.dirty = 1
{{- range .Fields}}
{{- if .Repeated}}{{/* nothing to do*/}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	x.syncable.{{.Name}} = make(map[{{toGoType .KeyType}}]*{{.Type}})
{{- else if findEnum .Type}}
	x.syncable.{{.Name}} = make(map[{{toGoType .KeyType}}]{{.Type}})
{{- else}}
	x.syncable.{{.Name}} = make(map[{{toGoType .KeyType}}]{{toGoType .Type}})
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.set{{.Name}}(New{{.Type}}())
{{- end}}
{{- end}}
{{- end}}
	return x
}
{{- template "Message" .}}

func (x *{{.Name}}) markAll() {
	x.dirty = uint64(0x01)
}

func (x *{{.Name}}) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
	x.dirtyParent.invoke()
}

func (x *{{.Name}}) clearAll() {
{{- range .Fields}}
{{- if findComponent .Type}}
{{- if .Repeated}}
	for i := 0; i < len(x.syncable.{{.Name}}); i++ {
		x.syncable.{{.Name}}[i].clearDirty()
	}
{{- else if len .KeyType}}
	for _, v := range x.syncable.{{.Name}} {
		v.clearDirty()
	}
{{- else}}
	x.syncable.{{.Name}}.clearDirty()
{{- end}}
{{- end}}
{{- end}}
	x.dirty = 0
}

func (x *{{.Name}}) clearDirty() {
	if x.dirty == 0 {
		return
	}
	if x.dirty & uint64(0x01) != 0 {
		x.clearAll()
		return
	}
{{- range .Fields}}
{{- if findComponent .Type}}
	if x.dirty & uint64(0x01) << {{.Number}} != 0 {
{{- if .Repeated}}
		for i := 0; i < len(x.syncable.{{.Name}}); i++ {
			x.syncable.{{.Name}}[i].clearDirty()
		}
{{- else if len .KeyType}}
		for _, v := range x.syncable.{{.Name}} {
			v.clearDirty()
		}
{{- else}}
		x.syncable.{{.Name}}.clearDirty()
{{- end}}
	}
{{- end}}
{{- end}}
	x.dirty = 0
}

func (x *{{.Name}}) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}
{{- end}}

{{- /* END DEFINE */ -}}

// Code generated by kds. DO NOT EDIT.
// source: {{"TODO: Source File"}}

package {{.Package}};

{{- if len .GoImportSpecs}}
import (
{{- range .GoImportSpecs}}
{{- if .SpacesBefore}}
{{""}}
{{- end}}
	"{{.Path}}"
{{- end}}
)
{{- end}}

{{- range .Types}}

{{- with findList .}}
{{- if .}}
{{- template "CommonList" .}}
{{- end}}
{{- end}}

{{- range findMap .}}
{{- template "CommonMap" .}}
{{- end}}

{{- end}}

{{- range .Defs}}
{{- if findEnum .Name}}
{{- template "Enum" .}}
{{- else if findEntity .Name}}
{{- template "Entity" .}}
{{- else if findComponent .Name}}
{{- template "Component" .}}

{{- with findList .Name}}
{{- if .}}
{{- template "CommonList" .}}
{{- end}}
{{- end}}

{{- range findMap .Name}}
{{- template "CommonMap" .}}
{{- end}}

{{- end}}
{{- end}}