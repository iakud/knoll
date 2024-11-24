package tree

import (
	"fmt"
	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Component struct {
	Name string
	Fields []*Field
}

func newComponent(ctx parser.IComponentDefContext) *Component {
	component := new(Component)
	component.Name = ctx.ComponentName().GetText()
	for _, element := range ctx.ComponentBody().AllComponentElement() {
		switch {
		case element.Field() != nil:
			component.Fields = append(component.Fields, newField(element.Field()))
		case element.MapField() != nil:
			component.Fields = append(component.Fields, newMapField(element.MapField()))
		}
	}
	return component
}

func (c *Component) String() string {
	return c.Name + fmt.Sprintf("%v", c.Fields)
}
