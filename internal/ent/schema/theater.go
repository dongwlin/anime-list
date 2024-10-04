package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Theater holds the schema definition for the Theater entity.
type Theater struct {
	ent.Schema
}

// Fields of the Theater.
func (Theater) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("cover"),
		field.Time("released_at").
			Default(time.Now()),
		field.String("description").
			Default(""),
		field.Int("status").
			Default(0),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Theater.
func (Theater) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("anime", Anime.Type).
			Ref("theaters").
			Unique(),
	}
}
