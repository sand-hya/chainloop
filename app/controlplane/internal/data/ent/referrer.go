// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/referrer"
	"github.com/google/uuid"
)

// Referrer is the model entity for the Referrer schema.
type Referrer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Digest holds the value of the "digest" field.
	Digest string `json:"digest,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// Downloadable holds the value of the "downloadable" field.
	Downloadable bool `json:"downloadable,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReferrerQuery when eager-loading is set.
	Edges        ReferrerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ReferrerEdges holds the relations/edges for other nodes in the graph.
type ReferrerEdges struct {
	// ReferredBy holds the value of the referred_by edge.
	ReferredBy []*Referrer `json:"referred_by,omitempty"`
	// References holds the value of the references edge.
	References []*Referrer `json:"references,omitempty"`
	// Organizations holds the value of the organizations edge.
	Organizations []*Organization `json:"organizations,omitempty"`
	// Workflows holds the value of the workflows edge.
	Workflows []*Workflow `json:"workflows,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ReferredByOrErr returns the ReferredBy value or an error if the edge
// was not loaded in eager-loading.
func (e ReferrerEdges) ReferredByOrErr() ([]*Referrer, error) {
	if e.loadedTypes[0] {
		return e.ReferredBy, nil
	}
	return nil, &NotLoadedError{edge: "referred_by"}
}

// ReferencesOrErr returns the References value or an error if the edge
// was not loaded in eager-loading.
func (e ReferrerEdges) ReferencesOrErr() ([]*Referrer, error) {
	if e.loadedTypes[1] {
		return e.References, nil
	}
	return nil, &NotLoadedError{edge: "references"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading.
func (e ReferrerEdges) OrganizationsOrErr() ([]*Organization, error) {
	if e.loadedTypes[2] {
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// WorkflowsOrErr returns the Workflows value or an error if the edge
// was not loaded in eager-loading.
func (e ReferrerEdges) WorkflowsOrErr() ([]*Workflow, error) {
	if e.loadedTypes[3] {
		return e.Workflows, nil
	}
	return nil, &NotLoadedError{edge: "workflows"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Referrer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case referrer.FieldDownloadable:
			values[i] = new(sql.NullBool)
		case referrer.FieldDigest, referrer.FieldKind:
			values[i] = new(sql.NullString)
		case referrer.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case referrer.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Referrer fields.
func (r *Referrer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case referrer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case referrer.FieldDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field digest", values[i])
			} else if value.Valid {
				r.Digest = value.String
			}
		case referrer.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				r.Kind = value.String
			}
		case referrer.FieldDownloadable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field downloadable", values[i])
			} else if value.Valid {
				r.Downloadable = value.Bool
			}
		case referrer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Referrer.
// This includes values selected through modifiers, order, etc.
func (r *Referrer) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryReferredBy queries the "referred_by" edge of the Referrer entity.
func (r *Referrer) QueryReferredBy() *ReferrerQuery {
	return NewReferrerClient(r.config).QueryReferredBy(r)
}

// QueryReferences queries the "references" edge of the Referrer entity.
func (r *Referrer) QueryReferences() *ReferrerQuery {
	return NewReferrerClient(r.config).QueryReferences(r)
}

// QueryOrganizations queries the "organizations" edge of the Referrer entity.
func (r *Referrer) QueryOrganizations() *OrganizationQuery {
	return NewReferrerClient(r.config).QueryOrganizations(r)
}

// QueryWorkflows queries the "workflows" edge of the Referrer entity.
func (r *Referrer) QueryWorkflows() *WorkflowQuery {
	return NewReferrerClient(r.config).QueryWorkflows(r)
}

// Update returns a builder for updating this Referrer.
// Note that you need to call Referrer.Unwrap() before calling this method if this Referrer
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Referrer) Update() *ReferrerUpdateOne {
	return NewReferrerClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Referrer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Referrer) Unwrap() *Referrer {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Referrer is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Referrer) String() string {
	var builder strings.Builder
	builder.WriteString("Referrer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("digest=")
	builder.WriteString(r.Digest)
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(r.Kind)
	builder.WriteString(", ")
	builder.WriteString("downloadable=")
	builder.WriteString(fmt.Sprintf("%v", r.Downloadable))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Referrers is a parsable slice of Referrer.
type Referrers []*Referrer
