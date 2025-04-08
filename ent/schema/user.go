package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return append([]ent.Field{
		field.String("username"),
		field.String("password").
			Annotations(entgql.Skip(entgql.SkipType)),
	})
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
