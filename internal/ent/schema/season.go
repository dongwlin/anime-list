package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("value"),
		field.String("cover"),
		field.Time("released_at"),
		field.String("description").
			Default(""),
		field.Int("status").
			Default(0),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("episodes", Episode.Type),
		edge.From("anime", Anime.Type).
			Ref("seasons").
			Unique(),
	}
}
