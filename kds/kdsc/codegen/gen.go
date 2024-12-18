package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
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
		kds.format()
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