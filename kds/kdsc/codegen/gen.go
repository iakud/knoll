package codegen

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iakud/knoll/kds/kdsc/parser"
)

func Parse(kdsFiles []string) *Context {
	ctx := newContext()
	for _, filePath := range kdsFiles {
		kds := parseKds(ctx, filePath)

		ctx.AllKds = append(ctx.AllKds, kds)
	}
	// format
	ctx.format()
	for _, kds := range ctx.AllKds {
		formatKds(ctx, kds)
	}
	return ctx
}

func WriteProtobuf(ctx *Context, out string) {
	tpl, err := template.New("protobuf").Funcs(Funcs(ctx)).Parse(TemplateProtobuf)
	if err != nil {
		panic(err)
	}
	for _, kds := range ctx.AllKds {
		buf := bytes.NewBuffer(nil)
		tpl.Execute(buf, kds)
		outFile := filepath.Join(out, kds.Name+".proto")
		os.WriteFile(outFile, buf.Bytes(), os.ModePerm)
	}
}

func WriteKdsGo(ctx *Context, out string) {
	tpl, err := template.New("kds").Funcs(Funcs(ctx)).Parse(TemplateKdsGo)
	if err != nil {
		panic(err)
	}
	for _, kds := range ctx.AllKds {
		buf := bytes.NewBuffer(nil)
		tpl.Execute(buf, kds)
		outFile := filepath.Join(out, kds.Name+".kds.go")
		os.WriteFile(outFile, buf.Bytes(), os.ModePerm)
	}
}

func WriteTemplate(ctx *Context, filename string, out string) {
	name := filepath.Base(filename)
	ext := strings.TrimSuffix(name, filepath.Ext(name))

	text, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	tpl, err := template.New("kds").Funcs(Funcs(ctx)).Parse(string(text))
	slog.Info("write", "file", filename)
	if err != nil {
		panic(err)
	}
	for _, kds := range ctx.AllKds {
		buf := bytes.NewBuffer(nil)
		tpl.Execute(buf, kds)
		outFile := filepath.Join(out, kds.Name+"."+ext)
		os.WriteFile(outFile, buf.Bytes(), os.ModePerm)
	}
}

func parseKds(ctx *Context, filePath string) *Kds {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("kds parse file %s error: %v\n%s", filePath, err, debug.Stack()))
		}
	}()

	input, err := antlr.NewFileStream(filePath)
	if err != nil {
		panic(err)
	}

	lexer := parser.NewkdsLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	kdsParser := parser.NewkdsParser(stream)
	kds := visitKds(ctx, filePath, kdsParser.Kds())
	return kds
}

func formatKds(ctx *Context, kds *Kds) {
	if kds.Name == "common" {
		for typ, listType := range ctx.TypeList {
			if _, ok := ctx.Defs[typ]; ok {
				continue
			}
			kds.addType(typ)
			kds.ListTypes = append(kds.ListTypes, listType)
		}
		for typ, mapTypes := range ctx.TypeMap {
			if _, ok := ctx.Defs[typ]; ok {
				continue
			}
			kds.addType(typ)
			kds.MapTypes = append(kds.MapTypes, mapTypes...)
		}
	}

	for _, def := range kds.Defs {
		if listType, ok := ctx.TypeList[def.GetName()]; ok {
			kds.ListTypes = append(kds.ListTypes, listType)
		}
		if mapTypes, ok := ctx.TypeMap[def.GetName()]; ok {
			kds.MapTypes = append(kds.MapTypes, mapTypes...)
		}
		for _, field := range def.GetFields() {
			if slices.Contains(kds.FieldTypes, field.Type) {
				continue
			}
			kds.FieldTypes = append(kds.FieldTypes, field.Type)
		}
	}

	kds.format()
}
