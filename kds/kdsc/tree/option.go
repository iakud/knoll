package tree

import (
	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Option struct {
	Name string
	Constant string
}

func newOption(ctx parser.IOptionStatementContext) *Option {
	option := new(Option)
	option.Name = ctx.OptionName().GetText()
	option.Constant = ctx.Constant().GetText()
	
	return option
}
