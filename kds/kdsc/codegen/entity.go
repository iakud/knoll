package codegen

import (
	"fmt"
	"github.com/iakud/keeper/kds/kdsc/parser"
)

type Entity struct {
	Name string
	Fields []*Field
}

func newEntity(ctx parser.IEntityDefContext) *Entity {
	entity := new(Entity)
	entity.Name = ctx.EntityName().GetText()
	for _, element := range ctx.EntityBody().AllEntityElement() {
		switch {
		case element.Field() != nil:
			entity.Fields = append(entity.Fields, newField(element.Field()))
		case element.MapField() != nil:
			entity.Fields = append(entity.Fields, newMapField(element.MapField()))
		}
	}
	return entity
}

func (e *Entity) String() string {
	return e.Name + fmt.Sprintf("%v", e.Fields)
}