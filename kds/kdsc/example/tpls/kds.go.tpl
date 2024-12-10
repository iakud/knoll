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

type syncable{{.Name}} struct {
{{- range .Fields}}
{{- if .Repeated}}{{/* Array */}}
{{- if findComponent .Type}}
	{{.Name}} []*{{.Type}}
{{- else if findEnum .Type}}
	{{.Name}} []{{.Type}}
{{- else}}
	{{.Name}} []{{.GoType}}
{{- end}}
{{- else if len .KeyType}}{{/* Map */}}
{{- if findComponent .Type}}
	{{.Name}} map[{{.KeyType}}]*{{.Type}}
{{- else if findEnum .Type}}
	{{.Name}} map[{{.KeyType}}]{{.Type}}
{{- else}}
	{{.Name}} map[{{.KeyType}}]{{.GoType}}
{{- end}}
{{- else}}{{/* Field */}}
{{- if findComponent .Type}}
	{{.Name}} *{{.Type}}
{{- else if findEnum .Type}}
	{{.Name}} {{.Type}}
{{- else}}
	{{.Name}} {{.GoType}}
{{- end}}
{{- end}}
{{- end}}
}

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

func (x *{{$MessageName}}) Get{{.Name}}() {{.GoType}} {
	return x.syncable.{{.Name}}
}

func (x *{{$MessageName}}) Set{{.Name}}(v {{.GoType}}) {
	if v == x.syncable.{{.Name}} {
		return
	}
	x.syncable.{{.Name}} = v
	x.markDirty(uint64(0x01) << {{.Number}})
}
{{- end}}
{{- end}}
{{- end}}

func (x *{{$MessageName}}) DumpChange() *pb.{{.Name}} {
	if x.checkDirty(uint64(0x01)) {
		return x.DumpFull()
	}
	m := new(pb.{{.Name}})
{{- range .Fields}}
	if x.checkDirty(uint64(0x01) << {{.Number}}) {
{{- if .Repeated}}
		
{{- if findComponent .Type}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, v.DumpChange())
		}
{{- else if isTimestamp .Type}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, timestamppb.New(v))
		}
{{- else if isDuration .Type}}
		for _, v := range x.syncable.{{.Name}} {
			m.{{.Name}} = append(m.{{.Name}}, durationpb.New(v))
		}
{{- else if isEmpty .Type}}
		for _, _ = range x.syncable.{{.Name}} {
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
{{- else if isTimestamp .Type}}
		for k, v := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = timestamppb.New(v)
		}
{{- else if isDuration .Type}}
		for k, v := range x.syncable.{{.Name}} {
			m.{{.Name}}[k] = durationpb.New(v)
		}
{{- else if isEmpty .Type}}
		for k, _ := range x.syncable.{{.Name}} {
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
{{- else if isTimestamp .Type}}
		m.{{.Name}} = timestamppb.New(x.syncable.{{.Name}})
{{- else if isDuration .Type}}
		m.{{.Name}} = durationpb.New(x.syncable.{{.Name}})
{{- else if isEmpty .Type}}
		m.{{.Name}} = new(emptypb.Empty)
{{- else}}
		m.{{.Name}} = x.syncable.{{.Name}}
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
{{- if findComponent .Type}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, v.DumpFull())
	}
{{- else if isTimestamp .Type}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, timestamppb.New(v))
	}
{{- else if isDuration .Type}}
	for _, v := range x.syncable.{{.Name}} {
		m.{{.Name}} = append(m.{{.Name}}, durationpb.New(v))
	}
{{- else if isEmpty .Type}}
	for _, _ = range x.syncable.{{.Name}} {
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
{{- else if isTimestamp .Type}}
	for k, v := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = timestamppb.New(v)
	}
{{- else if isDuration .Type}}
	for k, v := range x.syncable.{{.Name}} {
		m.{{.Name}}[k] = durationpb.New(v)
	}
{{- else if isEmpty .Type}}
	for k, _ := range x.syncable.{{.Name}} {
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
{{- else if isTimestamp .Type}}
	m.{{.Name}} = timestamppb.New(x.syncable.{{.Name}})
{{- else if isDuration .Type}}
	m.{{.Name}} = durationpb.New(x.syncable.{{.Name}})
{{- else if isEmpty .Type}}
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
{{- $EntityName := (print .Name)}}

type {{$EntityName}} struct {
	id int64
	syncable syncable{{.Name}}

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
	x.syncable.{{.Name}} = make(map[{{.KeyType}}]*{{.Type}})
{{- else if findEnum .Type}}
	x.syncable.{{.Name}} = make(map[{{.KeyType}}]{{.Type}})
{{- else}}
	x.syncable.{{.Name}} = make(map[{{.KeyType}}]{{.GoType}})
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.set{{.Name}}(New{{.Type}}())
{{- end}}
{{- end}}
{{- end}}
	return x
}

func (x *{{$EntityName}}) Id() int64 {
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
	for i := 0; i < len(x.syncable.{{.Name}}); i++ {
		x.syncable.{{.Name}}[i].clearDirty()
	}
{{- end}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	for _, v := range x.syncable.{{.Name}} {
		v.clearDirty()
	}
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.syncable.{{.Name}}.clearDirty()
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
	syncable syncable{{.Name}}

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
	x.syncable.{{.Name}} = make(map[{{.KeyType}}]*{{.Type}})
{{- else if findEnum .Type}}
	x.syncable.{{.Name}} = make(map[{{.KeyType}}]{{.Type}})
{{- else}}
	x.syncable.{{.Name}} = make(map[{{.KeyType}}]{{.GoType}})
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
	for _, v := range x.syncable.{{.Name}} {
		v.clearDirty()
	}
{{- end}}
{{- else if len .KeyType}}
{{- if findComponent .Type}}
	for _, v := range x.syncable.{{.Name}} {
		v.clearDirty()
	}
{{- end}}
{{- else}}
{{- if findComponent .Type}}
	x.syncable.{{.Name}}.clearDirty()
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
{{- if .ImportEmpty}}
	"google.golang.org/protobuf/types/known/emptypb"
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