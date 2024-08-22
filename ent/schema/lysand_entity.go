package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/pkg/versia"
	"net/url"
	"time"
)

type LysandEntityMixin struct{ mixin.Schema }

var _ ent.Mixin = (*LysandEntityMixin)(nil)

func (LysandEntityMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.Bool("isRemote"),
		field.String("uri").Validate(ValidateURI),

		field.JSON("extensions", versia.Extensions{}).Default(versia.Extensions{}),

		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func ValidateURI(s string) error {
	_, err := url.Parse(s)
	return err
}
