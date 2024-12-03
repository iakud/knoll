package codegen

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/iakud/keeper/kds/kdsc/parser"
)

func Parse(file string) *Kds {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		panic(err)
	}

	lexer := parser.NewkdsLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	kdsParser := parser.NewkdsParser(stream)
	kds := New(kdsParser.Kds())
	return kds
}