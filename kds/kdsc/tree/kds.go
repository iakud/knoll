package tree

import (
	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Kds struct {
	Package string
	Entities []*Entity
	Components []*Component
}

func New(ctx parser.IKdsContext) *Kds {
	kds := new(Kds)
	kds.Package = ctx.PackageStatement().FullIdent().GetText()
	for _, topLevel := range ctx.AllTopLevelDef() {
		switch {
		case topLevel.EntityDef() != nil:
			kds.Entities = append(kds.Entities, newEntity(topLevel.EntityDef()))
		case topLevel.ComponentDef() != nil:
			kds.Components = append(kds.Components, newComponent(topLevel.ComponentDef()))
		}
	}
	return kds
}