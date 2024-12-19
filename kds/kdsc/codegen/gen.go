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

func Parse(kdsFiles []string, tplPath string, out string) error {
	ctx := newContext()
	var kdsList []*Kds
	for _, kdsFile := range kdsFiles {
		kds := parseKds(ctx, kdsFile)
		
		kdsList = append(kdsList, kds)
	}
	var tpls []*template.Template

	// templates
	tplFiles, err := filepath.Glob(filepath.Join(tplPath, "*.tpl"))
	if err != nil {
		return err
	}
	for _, tplFile := range tplFiles {
		b, err := os.ReadFile(tplFile)
		if err != nil {
			return err
		}
		name := filepath.Base(tplFile)
		tpl, err := template.New(name).Funcs(Funcs(ctx)).Parse(string(b))
		if err != nil {
			return err
		}
		tpls = append(tpls, tpl)
	}
	// format
	ctx.format()

	for _, kds := range kdsList {
		formatKds(ctx, kds)
		for _, tpl := range tpls {
			buf := bytes.NewBuffer(nil)
			tpl.Execute(buf, kds)
			outFile := filepath.Join(out, kds.Name + "." + strings.TrimSuffix(filepath.Base(tpl.Name()), filepath.Ext(tpl.Name())))
			os.WriteFile(outFile, buf.Bytes(), os.ModePerm)
		}
	}
	return nil
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