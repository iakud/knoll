package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iakud/krocher/kds/kdsc/parser"
)

func Parse(kdsFiles []string) *Context {
	ctx := newContext()
	for _, kdsFile := range kdsFiles {
		kds := parseKds(ctx, kdsFile)
		
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
		outFile := filepath.Join(out, kds.Name + ".proto")
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
		outFile := filepath.Join(out, kds.Name + ".kds.go")
		os.WriteFile(outFile, buf.Bytes(), os.ModePerm)
	}
}

func parseKds(ctx *Context, kdsFile string) *Kds {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("kds parse file %s error: %v\n%s", kdsFile, err, debug.Stack()))
		}
	}()

	input, err := antlr.NewFileStream(kdsFile)
	if err != nil {
		panic(err)
	}

	lexer := parser.NewkdsLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	kdsParser := parser.NewkdsParser(stream)
	kdsName := strings.TrimSuffix(filepath.Base(kdsFile), filepath.Ext(kdsFile))
	kds := visitKds(ctx, kdsName, kdsParser.Kds())
	return kds
}

func formatKds(ctx *Context, kds *Kds) {
	var importSlices, importMaps bool
	if kds.Name == "common" {
		for type_ := range ctx.TypeList {
			if _, ok := ctx.Defs[type_]; ok {
				continue
			}
			kds.addType(type_)
			importSlices = true
		}
		for type_ := range ctx.TypeMap {
			if _, ok := ctx.Defs[type_]; ok {
				continue
			}
			kds.addType(type_)
			importMaps = true
		}
	}
	importSlices = importSlices || slices.ContainsFunc(kds.Defs, func(def TopLevelDef) bool {
		if _, ok := ctx.TypeList[def.GetName()]; ok {
			return true
		}
		return false
	})
	importMaps = importMaps || slices.ContainsFunc(kds.Defs, func(def TopLevelDef) bool {
		if _, ok := ctx.TypeMap[def.GetName()]; ok {
			return true
		}
		return false
	})

	if importSlices {
		kds.addGoImport("slices", "")
		kds.addGoImport("iter", "")
	}
	if importMaps {
		kds.addGoImport("maps", "")
		kds.addGoImport("iter", "")
	}

	kds.format()
}