{{- /* BEGIN DEFINE */ -}}

{{- define "Enum"}}
{{- $EnumType := .Name}}

type {{.Name}} = pb.{{.Name}}

const (
{{- range .EnumFields}}
	{{$EnumType}}_{{.Name}} {{$EnumType}} = {{.Value}}
{{- end}}
)
{{- end}}

{{- define "Message"}}
{{- $MessageName := .Name}}

{{- range .Fields}}
{{- if .Repeated}}

{{- else if len .KeyType}}

{{- else}}
{{- if findComponent .Type}}

func (x *{{$MessageName}}) Get{{.Name}}() *{{.Type}} {
	return x.{{.GoVarName}}
}

func (x *{{$MessageName}}) set{{.Name}}(v *{{.Type}}) {
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed or evicted from its original place first")
	}
	if v == x.{{.GoVarName}} {
		return
	}
	if x.{{.GoVarName}} != nil {
		x.{{.GoVarName}}.dirtyParent = nil
	}
	x.{{.GoVarName}} = v
	v.dirtyParent = func() {
		x.markDirty(uint64(0x01) << {{.Number}})
	}
	x.markDirty(uint64(0x01) << {{.Number}})
	if v != nil {
		v.dirty |= uint64(0x01)
	}
}
{{- else if findEnum .Type}}

func (x *{{$MessageName}}) Get{{.Name}}() {{.Type}} {
	return x.{{.GoVarName}}
}

func (x *{{$MessageName}}) Set{{.Name}}(v {{.Type}}) {
	if v == x.{{.GoVarName}} {
		return
	}
	x.{{.GoVarName}} = v
	x.markDirty(uint64(0x01) << {{.Number}})
}
{{- else}}

func (x *{{$MessageName}}) Get{{.Name}}() {{.GoType}} {
	return x.{{.GoVarName}}
}

func (x *{{$MessageName}}) Set{{.Name}}(v {{.GoType}}) {
	if v == x.{{.GoVarName}} {
		return
	}
	x.{{.GoVarName}} = v
	x.markDirty(uint64(0x01) << {{.Number}})
}
{{- end}}
{{- end}}
{{- end}}

func (x *{{$MessageName}}) DumpChange() *pb.{{.Name}} {
	m := new(pb.{{.Name}})
{{- range .Fields}}
	if x.checkDirty(uint64(0x01) << {{.Number}}) {
{{- if .Repeated}}
		for _, v := range x.{{.GoVarName}} {
{{- if findComponent .Type}}
			m.{{.Name}} = append(m.{{.Name}}, v.DumpChange())
{{- else if isTimestamp .Type}}
			m.{{.Name}} = append(m.{{.Name}}, timestamppb.New(v))
{{- else if isDuration .Type}}
			m.{{.Name}} = append(m.{{.Name}}, durationpb.New(v))
{{- else}}
			m.{{.Name}} = append(m.{{.Name}}, v)
{{- end}}
		}
{{- else if len .KeyType}}
		for k, v := range x.{{.GoVarName}} {
{{- if findComponent .Type}}
			m.{{.Name}}[k] = v.DumpChange()
{{- else if isTimestamp .Type}}
			m.{{.Name}}[k] = timestamppb.New(v)
{{- else if isDuration .Type}}
			m.{{.Name}}[k] = durationpb.New(v)
{{- else}}
			m.{{.Name}}[k] = v
{{- end}}
		}
{{- else}}
{{- if findComponent .Type}}
		m.{{.Name}} = x.{{.GoVarName}}.DumpChange()
{{- else if isTimestamp .Type}}
		m.{{.Name}} = timestamppb.New(x.{{.GoVarName}})
{{- else if isDuration .Type}}
		m.{{.Name}} = durationpb.New(x.{{.GoVarName}})
{{- else}}
		m.{{.Name}} = x.{{.GoVarName}}
{{- end}}
{{- end}}
	}
{{- end}}
	return m
}

func (x *{{$MessageName}}) DumpFull() *pb.{{.Name}} {
	m := new(pb.{{.Name}})
{{- range .Fields}}
{{- if .Repeated}}
	for _, v := range x.{{.GoVarName}} {
{{- if findComponent .Type}}
		m.{{.Name}} = append(m.{{.Name}}, v.DumpFull())
{{- else if isTimestamp .Type}}
		m.{{.Name}} = append(m.{{.Name}}, timestamppb.New(v))
{{- else if isDuration .Type}}
		m.{{.Name}} = append(m.{{.Name}}, durationpb.New(v))
{{- else}}
		m.{{.Name}} = append(m.{{.Name}}, v)
{{- end}}
	}
{{- else if len .KeyType}}
	for k, v := range x.{{.GoVarName}} {
{{- if findComponent .Type}}
		m.{{.Name}}[k] = v.DumpFull()
{{- else if isTimestamp .Type}}
		m.{{.Name}}[k] = timestamppb.New(v)
{{- else if isDuration .Type}}
		m.{{.Name}}[k] = durationpb.New(v)
{{- else}}
		m.{{.Name}}[k] = v
{{- end}}
	}
{{- else}}
{{- if findComponent .Type}}
	m.{{.Name}} = x.{{.GoVarName}}.DumpFull()
{{- else if isTimestamp .Type}}
	m.{{.Name}} = timestamppb.New(x.{{.GoVarName}})
{{- else if isDuration .Type}}
	m.{{.Name}} = durationpb.New(x.{{.GoVarName}})
{{- else}}
	m.{{.Name}} = x.{{.GoVarName}}
{{- end}}
{{- end}}
{{- end}}
	return m
}

{{- end}}

{{- define "Entity"}}
{{- $EntityName := (print .Name)}}

type {{$EntityName}} struct {
	id int64
{{- range .Fields}}
{{- if .Repeated}}{{/* Array */}}
{{- if findComponent .Type}}
	{{.GoVarName}} []*{{.Type}}
{{- else if findEnum .Type}}
	{{.GoVarName}} []{{.Type}}
{{- else}}
	{{.GoVarName}} []{{.GoType}}
{{- end}}
{{- else if len .KeyType}}{{/* Map */}}
{{- if findComponent .Type}}
	{{.GoVarName}} map[{{.KeyType}}]*{{.Type}}
{{- else if findEnum .Type}}
	{{.GoVarName}} map[{{.KeyType}}]{{.Type}}
{{- else}}
	{{.GoVarName}} map[{{.KeyType}}]{{.GoType}}
{{- end}}
{{- else}}{{/* Field */}}
{{- if findComponent .Type}}
	{{.GoVarName}} *{{.Type}}
{{- else if findEnum .Type}}
	{{.GoVarName}} {{.Type}}
{{- else}}
	{{.GoVarName}} {{.GoType}}
{{- end}}
{{- end}}
{{- end}}

	dirty uint64
}

func New{{$EntityName}}() *{{$EntityName}} {
	x := new({{$EntityName}})
	x.dirty = 1
	x.id = 0 // FIXME: gen nextId()
{{- range .Fields}}
{{- if .Repeated}}{{/* nothing to do*/}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	x.{{.GoVarName}} = make(map[{{.KeyType}}]*{{.Type}})
{{- else if findEnum .Type}}
	x.{{.GoVarName}} = make(map[{{.KeyType}}]{{.Type}})
{{- else}}
	x.{{.GoVarName}} = make(map[{{.KeyType}}]{{.GoType}})
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.set{{.Name}}(New{{.Type}}())
{{- end}}
{{- end}}
{{- end}}
	return x
}

func (x *{{$EntityName}}) GetId() int64 {
	return x.id
}
{{- template "Message" .}}

func (x *{{$EntityName}}) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
}

func (x *{{$EntityName}}) clearDirty() {
	if x.dirty == 0 {
		return
	}
	x.dirty = 0
{{- range .Fields}}
{{- if .Repeated}}
{{- if findComponent .Type}}
	for i := 0; i < len(x.{{.GoVarName}}); i++ {
		x.{{.GoVarName}}[i].clearDirty()
	}
{{- end}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	for _, v := range x.{{.GoVarName}} {
		v.clearDirty()
	}
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.{{.GoVarName}}.clearDirty()
{{- end}}
{{- end}}
{{- end}}
}

func (x *{{$EntityName}}) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}

{{- end}}

{{- define "Component"}}
{{- $ComponentName := .Name}}

type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{$ComponentName}} struct {
{{- range .Fields}}
{{- if .Repeated}}{{/* Array */}}
{{- if findComponent .Type}}
	{{.GoVarName}} []*{{.Type}}
{{- else if findEnum .Type}}
	{{.GoVarName}} []{{.Type}}
{{- else}}
	{{.GoVarName}} []{{.GoType}}
{{- end}}
{{- else if len .KeyType}}{{/* Map */}}
{{- if findComponent .Type}}
	{{.GoVarName}} map[{{.KeyType}}]*{{.Type}}
{{- else if findEnum .Type}}
	{{.GoVarName}} map[{{.KeyType}}]{{.Type}}
{{- else}}
	{{.GoVarName}} map[{{.KeyType}}]{{.GoType}}
{{- end}}
{{- else}}{{/* Field */}}
{{- if findComponent .Type}}
	{{.GoVarName}} *{{.Type}}
{{- else if findEnum .Type}}
	{{.GoVarName}} {{.Type}}
{{- else}}
	{{.GoVarName}} {{.GoType}}
{{- end}}
{{- end}}
{{- end}}

	dirty uint64
	dirtyParent dirtyParentFunc_{{.Name}}
}

func New{{$ComponentName}}() *{{$ComponentName}} {
	x := new({{$ComponentName}})
	x.dirty = 1
{{- range .Fields}}
{{- if .Repeated}}{{/* nothing to do*/}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	x.{{.GoVarName}} = make(map[{{.KeyType}}]*{{.Type}})
{{- else if findEnum .Type}}
	x.{{.GoVarName}} = make(map[{{.KeyType}}]{{.Type}})
{{- else}}
	x.{{.GoVarName}} = make(map[{{.KeyType}}]{{.GoType}})
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

func (x *{{$ComponentName}}) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
	x.dirtyParent.invoke()
}

func (x *{{$ComponentName}}) clearDirty() {
	if x.dirty == 0 {
		return
	}
	x.dirty = 0
{{- range .Fields}}
{{- if .Repeated}}
{{- if findComponent .Type}}
	for _, v := range x.{{.GoVarName}} {
		v.clearDirty()
	}
{{- end}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	for _, v := range x.{{.GoVarName}} {
		v.clearDirty()
	}
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.{{.GoVarName}}.clearDirty()
{{- end}}
{{- end}}
{{- end}}
}

func (x *{{$ComponentName}}) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}
{{- end}}

{{- /* END DEFINE */ -}}

// Code generated by kds. DO NOT EDIT.
// source: {{"TODO: Source File"}}

package {{.Package}};

import (
{{- if or .ImportTimestamp .ImportDuration}}
	"time"
{{""}}
{{- end}}
	"github.com/iakud/keeper/kds/kdsc/example/pb"
{{- if .ImportTimestamp}}
	"google.golang.org/protobuf/types/known/timestamppb"
{{- end}}
{{- if .ImportDuration}}
	"google.golang.org/protobuf/types/known/durationpb"
{{- end}}
)

{{- range .Defs}}
{{- if findEnum .Name}}
{{- template "Enum" .}}
{{- else if findEntity .Name}}
{{- template "Entity" .}}
{{- else if findComponent .Name}}
{{- template "Component" .}}
{{- end}}
{{- end}}