package schema

import "entgo.io/ent"

// Shop holds the schema definition for the Shop entity.
type Shop struct {
	ent.Schema
}

// Fields of the Shop.
func (Shop) Fields() []ent.Field {
	return nil
}

// Edges of the Shop.
func (Shop) Edges() []ent.Edge {
	return nil
}
