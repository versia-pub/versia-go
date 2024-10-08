// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/ent/instancemetadata"
	"github.com/versia-pub/versia-go/pkg/versia"
)

// InstanceMetadata is the model entity for the InstanceMetadata schema.
type InstanceMetadata struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// IsRemote holds the value of the "isRemote" field.
	IsRemote bool `json:"isRemote,omitempty"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri,omitempty"`
	// Extensions holds the value of the "extensions" field.
	Extensions versia.Extensions `json:"extensions,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Host holds the value of the "host" field.
	Host string `json:"host,omitempty"`
	// PublicKey holds the value of the "publicKey" field.
	PublicKey []byte `json:"publicKey,omitempty"`
	// PublicKeyAlgorithm holds the value of the "publicKeyAlgorithm" field.
	PublicKeyAlgorithm string `json:"publicKeyAlgorithm,omitempty"`
	// PrivateKey holds the value of the "privateKey" field.
	PrivateKey []byte `json:"privateKey,omitempty"`
	// SoftwareName holds the value of the "softwareName" field.
	SoftwareName string `json:"softwareName,omitempty"`
	// SoftwareVersion holds the value of the "softwareVersion" field.
	SoftwareVersion string `json:"softwareVersion,omitempty"`
	// SharedInboxURI holds the value of the "sharedInboxURI" field.
	SharedInboxURI string `json:"sharedInboxURI,omitempty"`
	// ModeratorsURI holds the value of the "moderatorsURI" field.
	ModeratorsURI *string `json:"moderatorsURI,omitempty"`
	// AdminsURI holds the value of the "adminsURI" field.
	AdminsURI *string `json:"adminsURI,omitempty"`
	// LogoEndpoint holds the value of the "logoEndpoint" field.
	LogoEndpoint *string `json:"logoEndpoint,omitempty"`
	// LogoMimeType holds the value of the "logoMimeType" field.
	LogoMimeType *string `json:"logoMimeType,omitempty"`
	// BannerEndpoint holds the value of the "bannerEndpoint" field.
	BannerEndpoint *string `json:"bannerEndpoint,omitempty"`
	// BannerMimeType holds the value of the "bannerMimeType" field.
	BannerMimeType *string `json:"bannerMimeType,omitempty"`
	// SupportedVersions holds the value of the "supportedVersions" field.
	SupportedVersions []string `json:"supportedVersions,omitempty"`
	// SupportedExtensions holds the value of the "supportedExtensions" field.
	SupportedExtensions []string `json:"supportedExtensions,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstanceMetadataQuery when eager-loading is set.
	Edges        InstanceMetadataEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InstanceMetadataEdges holds the relations/edges for other nodes in the graph.
type InstanceMetadataEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Moderators holds the value of the moderators edge.
	Moderators []*User `json:"moderators,omitempty"`
	// Admins holds the value of the admins edge.
	Admins []*User `json:"admins,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceMetadataEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ModeratorsOrErr returns the Moderators value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceMetadataEdges) ModeratorsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Moderators, nil
	}
	return nil, &NotLoadedError{edge: "moderators"}
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceMetadataEdges) AdminsOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InstanceMetadata) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case instancemetadata.FieldExtensions, instancemetadata.FieldPublicKey, instancemetadata.FieldPrivateKey, instancemetadata.FieldSupportedVersions, instancemetadata.FieldSupportedExtensions:
			values[i] = new([]byte)
		case instancemetadata.FieldIsRemote:
			values[i] = new(sql.NullBool)
		case instancemetadata.FieldURI, instancemetadata.FieldName, instancemetadata.FieldDescription, instancemetadata.FieldHost, instancemetadata.FieldPublicKeyAlgorithm, instancemetadata.FieldSoftwareName, instancemetadata.FieldSoftwareVersion, instancemetadata.FieldSharedInboxURI, instancemetadata.FieldModeratorsURI, instancemetadata.FieldAdminsURI, instancemetadata.FieldLogoEndpoint, instancemetadata.FieldLogoMimeType, instancemetadata.FieldBannerEndpoint, instancemetadata.FieldBannerMimeType:
			values[i] = new(sql.NullString)
		case instancemetadata.FieldCreatedAt, instancemetadata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case instancemetadata.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InstanceMetadata fields.
func (im *InstanceMetadata) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case instancemetadata.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				im.ID = *value
			}
		case instancemetadata.FieldIsRemote:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isRemote", values[i])
			} else if value.Valid {
				im.IsRemote = value.Bool
			}
		case instancemetadata.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				im.URI = value.String
			}
		case instancemetadata.FieldExtensions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field extensions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &im.Extensions); err != nil {
					return fmt.Errorf("unmarshal field extensions: %w", err)
				}
			}
		case instancemetadata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				im.CreatedAt = value.Time
			}
		case instancemetadata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				im.UpdatedAt = value.Time
			}
		case instancemetadata.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				im.Name = value.String
			}
		case instancemetadata.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				im.Description = new(string)
				*im.Description = value.String
			}
		case instancemetadata.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				im.Host = value.String
			}
		case instancemetadata.FieldPublicKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field publicKey", values[i])
			} else if value != nil {
				im.PublicKey = *value
			}
		case instancemetadata.FieldPublicKeyAlgorithm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publicKeyAlgorithm", values[i])
			} else if value.Valid {
				im.PublicKeyAlgorithm = value.String
			}
		case instancemetadata.FieldPrivateKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field privateKey", values[i])
			} else if value != nil {
				im.PrivateKey = *value
			}
		case instancemetadata.FieldSoftwareName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field softwareName", values[i])
			} else if value.Valid {
				im.SoftwareName = value.String
			}
		case instancemetadata.FieldSoftwareVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field softwareVersion", values[i])
			} else if value.Valid {
				im.SoftwareVersion = value.String
			}
		case instancemetadata.FieldSharedInboxURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sharedInboxURI", values[i])
			} else if value.Valid {
				im.SharedInboxURI = value.String
			}
		case instancemetadata.FieldModeratorsURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field moderatorsURI", values[i])
			} else if value.Valid {
				im.ModeratorsURI = new(string)
				*im.ModeratorsURI = value.String
			}
		case instancemetadata.FieldAdminsURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field adminsURI", values[i])
			} else if value.Valid {
				im.AdminsURI = new(string)
				*im.AdminsURI = value.String
			}
		case instancemetadata.FieldLogoEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logoEndpoint", values[i])
			} else if value.Valid {
				im.LogoEndpoint = new(string)
				*im.LogoEndpoint = value.String
			}
		case instancemetadata.FieldLogoMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logoMimeType", values[i])
			} else if value.Valid {
				im.LogoMimeType = new(string)
				*im.LogoMimeType = value.String
			}
		case instancemetadata.FieldBannerEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bannerEndpoint", values[i])
			} else if value.Valid {
				im.BannerEndpoint = new(string)
				*im.BannerEndpoint = value.String
			}
		case instancemetadata.FieldBannerMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bannerMimeType", values[i])
			} else if value.Valid {
				im.BannerMimeType = new(string)
				*im.BannerMimeType = value.String
			}
		case instancemetadata.FieldSupportedVersions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field supportedVersions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &im.SupportedVersions); err != nil {
					return fmt.Errorf("unmarshal field supportedVersions: %w", err)
				}
			}
		case instancemetadata.FieldSupportedExtensions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field supportedExtensions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &im.SupportedExtensions); err != nil {
					return fmt.Errorf("unmarshal field supportedExtensions: %w", err)
				}
			}
		default:
			im.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InstanceMetadata.
// This includes values selected through modifiers, order, etc.
func (im *InstanceMetadata) Value(name string) (ent.Value, error) {
	return im.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the InstanceMetadata entity.
func (im *InstanceMetadata) QueryUsers() *UserQuery {
	return NewInstanceMetadataClient(im.config).QueryUsers(im)
}

// QueryModerators queries the "moderators" edge of the InstanceMetadata entity.
func (im *InstanceMetadata) QueryModerators() *UserQuery {
	return NewInstanceMetadataClient(im.config).QueryModerators(im)
}

// QueryAdmins queries the "admins" edge of the InstanceMetadata entity.
func (im *InstanceMetadata) QueryAdmins() *UserQuery {
	return NewInstanceMetadataClient(im.config).QueryAdmins(im)
}

// Update returns a builder for updating this InstanceMetadata.
// Note that you need to call InstanceMetadata.Unwrap() before calling this method if this InstanceMetadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (im *InstanceMetadata) Update() *InstanceMetadataUpdateOne {
	return NewInstanceMetadataClient(im.config).UpdateOne(im)
}

// Unwrap unwraps the InstanceMetadata entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (im *InstanceMetadata) Unwrap() *InstanceMetadata {
	_tx, ok := im.config.driver.(*txDriver)
	if !ok {
		panic("ent: InstanceMetadata is not a transactional entity")
	}
	im.config.driver = _tx.drv
	return im
}

// String implements the fmt.Stringer.
func (im *InstanceMetadata) String() string {
	var builder strings.Builder
	builder.WriteString("InstanceMetadata(")
	builder.WriteString(fmt.Sprintf("id=%v, ", im.ID))
	builder.WriteString("isRemote=")
	builder.WriteString(fmt.Sprintf("%v", im.IsRemote))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(im.URI)
	builder.WriteString(", ")
	builder.WriteString("extensions=")
	builder.WriteString(fmt.Sprintf("%v", im.Extensions))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(im.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(im.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(im.Name)
	builder.WriteString(", ")
	if v := im.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("host=")
	builder.WriteString(im.Host)
	builder.WriteString(", ")
	builder.WriteString("publicKey=")
	builder.WriteString(fmt.Sprintf("%v", im.PublicKey))
	builder.WriteString(", ")
	builder.WriteString("publicKeyAlgorithm=")
	builder.WriteString(im.PublicKeyAlgorithm)
	builder.WriteString(", ")
	builder.WriteString("privateKey=")
	builder.WriteString(fmt.Sprintf("%v", im.PrivateKey))
	builder.WriteString(", ")
	builder.WriteString("softwareName=")
	builder.WriteString(im.SoftwareName)
	builder.WriteString(", ")
	builder.WriteString("softwareVersion=")
	builder.WriteString(im.SoftwareVersion)
	builder.WriteString(", ")
	builder.WriteString("sharedInboxURI=")
	builder.WriteString(im.SharedInboxURI)
	builder.WriteString(", ")
	if v := im.ModeratorsURI; v != nil {
		builder.WriteString("moderatorsURI=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := im.AdminsURI; v != nil {
		builder.WriteString("adminsURI=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := im.LogoEndpoint; v != nil {
		builder.WriteString("logoEndpoint=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := im.LogoMimeType; v != nil {
		builder.WriteString("logoMimeType=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := im.BannerEndpoint; v != nil {
		builder.WriteString("bannerEndpoint=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := im.BannerMimeType; v != nil {
		builder.WriteString("bannerMimeType=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("supportedVersions=")
	builder.WriteString(fmt.Sprintf("%v", im.SupportedVersions))
	builder.WriteString(", ")
	builder.WriteString("supportedExtensions=")
	builder.WriteString(fmt.Sprintf("%v", im.SupportedExtensions))
	builder.WriteByte(')')
	return builder.String()
}

// InstanceMetadataSlice is a parsable slice of InstanceMetadata.
type InstanceMetadataSlice []*InstanceMetadata
