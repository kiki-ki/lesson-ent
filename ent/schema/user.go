package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/guregu/null"
	"github.com/kiki-ki/lesson-ent/ent/schema/validation"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("company_id").
			Positive(),
		field.String("name").
			Validate(validation.BlackListString([]string{"hoge", "fuga"})),
		field.String("email").
			Unique().
			Match(regexp.MustCompile(validation.EmailRegex)),
		field.Enum("role").
			Values("admin", "normal"),
		field.Text("comment").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			GoType(null.String{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("users").
			Unique().
			Required().
			Field("company_id"),
	}
}
