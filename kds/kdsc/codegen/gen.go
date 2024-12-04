package codegen

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iakud/keeper/kds/kdsc/parser"
)

func Parse(files []string, tplFiles []string, out string) error {
	ctx := new(Context)
	var kdsList []*Kds
	for _, file := range files {
		input, err := antlr.NewFileStream(file)
		if err != nil {
			return err
		}

		lexer := parser.NewkdsLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		kdsParser := parser.NewkdsParser(stream)
		kds := ctx.VisitKds(kdsParser.Kds())
		kds.Filename = file
		kdsList = append(kdsList, kds)
	}
	var tpls []*template.Template
	for _, tplFile := range tplFiles {
		b, err := os.ReadFile(tplFile)
		if err != nil {
			panic(err)
		}
		name := filepath.Base(tplFile)
		tpl, err := template.New(name).Funcs(Funcs(ctx)).Parse(string(b))
		if err != nil {
			panic(err)
		}
		tpls = append(tpls, tpl)
	}

	for _, kds := range kdsList {
		for _, tpl := range tpls {
			buf := bytes.NewBuffer(nil)
			tpl.Execute(buf, kds)
			outFile := filepath.Join(out, strings.TrimSuffix(filepath.Base(kds.Filename), filepath.Ext(kds.Filename)) + "." + strings.TrimSuffix(filepath.Base(tpl.Name()), filepath.Ext(tpl.Name())))
			ioutil.WriteFile(outFile, buf.Bytes(), 0777)
		}
	}

	return nil
}

func Funcs(ctx *Context) template.FuncMap {
	return template.FuncMap{
		"isEnum": func(name string) bool {
			for _, enum := range ctx.Enums {
				if name == enum.Name {
					return true
				}
			}
			return false
		},
		"isComponent": func(name string) bool {
			for _, component := range ctx.Components {
				if name == component.Name {
					return true
				}
			}
			return false
		},
	}
}