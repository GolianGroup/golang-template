package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(50).Unique(),
		field.String("password").MaxLen(50).Sensitive(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
