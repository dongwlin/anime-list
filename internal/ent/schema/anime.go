package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Anime holds the schema definition for the Anime entity.
type Anime struct {
	ent.Schema
}

// Fields of the Anime.
func (Anime) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Default(""),
		field.Int("status").
			Default(0),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Anime.
func (Anime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("seasons", Season.Type),
		edge.To("theaters", Theater.Type),
	}
}
