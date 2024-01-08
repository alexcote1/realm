// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"realm.pub/tavern/internal/ent/file"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Timestamp of when this ent was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Timestamp of when this ent was last updated
	LastModifiedAt time.Time `json:"last_modified_at,omitempty"`
	// The name of the file, used to reference it for downloads
	Name string `json:"name,omitempty"`
	// The size of the file in bytes
	Size int `json:"size,omitempty"`
	// A SHA3 digest of the content field
	Hash string `json:"hash,omitempty"`
	// The content of the file
	Content      []byte `json:"content,omitempty"`
	tome_files   *int
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldContent:
			values[i] = new([]byte)
		case file.FieldID, file.FieldSize:
			values[i] = new(sql.NullInt64)
		case file.FieldName, file.FieldHash:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldLastModifiedAt:
			values[i] = new(sql.NullTime)
		case file.ForeignKeys[0]: // tome_files
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldLastModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_modified_at", values[i])
			} else if value.Valid {
				f.LastModifiedAt = value.Time
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = int(value.Int64)
			}
		case file.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				f.Hash = value.String
			}
		case file.FieldContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value != nil {
				f.Content = *value
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tome_files", value)
			} else if value.Valid {
				f.tome_files = new(int)
				*f.tome_files = int(value.Int64)
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_modified_at=")
	builder.WriteString(f.LastModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(f.Hash)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(fmt.Sprintf("%v", f.Content))
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File
