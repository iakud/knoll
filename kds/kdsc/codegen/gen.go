package codegen

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iakud/keeper/kds/kdsc/parser"
)

func Parse(kdsFiles []string, tplPath string, out string) error {
	ctx := new(Context)
	var kdsList []*Kds
	for _, kdsFile := range kdsFiles {
		input, err := antlr.NewFileStream(kdsFile)
		if err != nil {
			return err
		}

		lexer := parser.NewkdsLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		kdsParser := parser.NewkdsParser(stream)
		kds := ctx.VisitKds(kdsParser.Kds())
		kds.Filename = kdsFile
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

	for _, kds := range kdsList {
		for _, tpl := range tpls {
			buf := bytes.NewBuffer(nil)
			tpl.Execute(buf, kds)
			outFile := filepath.Join(out, strings.TrimSuffix(filepath.Base(kds.Filename), filepath.Ext(kds.Filename)) + "." + strings.TrimSuffix(filepath.Base(tpl.Name()), filepath.Ext(tpl.Name())))
			os.WriteFile(outFile, buf.Bytes(), os.ModePerm)
		}
	}
	return nil
}

func Funcs(ctx *Context) template.FuncMap {
	return template.FuncMap{
		"isEnum": ctx.IsEnum,
		"IsComponent": ctx.IsEntity,
		"isComponent": ctx.IsComponent,
	}
}