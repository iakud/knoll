package tree

import (
	"fmt"
	"strings"

	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Kds struct {
	Package string
	Imports []string
	Entities []*Entity
	Components []*Component
}

func New(ctx parser.IKdsContext) *Kds {
	kds := new(Kds)
	kds.Package = ctx.PackageStatement().FullIdent().GetText()
	for _, importStatement := range ctx.AllImportStatement() {
		element := importStatement.STR_LIT().GetText()
		switch {
		case strings.HasPrefix(element, "\"") && strings.HasSuffix(element, "\"") :
			kds.Imports = append(kds.Imports, strings.TrimSuffix(strings.TrimPrefix(element, "\""), "\""))
		case strings.HasPrefix(element, "'") && strings.HasSuffix(element, "'") :
			kds.Imports = append(kds.Imports, strings.TrimSuffix(strings.TrimPrefix(element, "'"), "'"))
		}
		fmt.Println("import element:", element)
	}
	for i := 0; i < len(kds.Imports); i++ {
		kds.Imports[i] = strings.TrimSuffix(kds.Imports[i], ".kds")
	}
	
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