// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/leohearts/insights-bot-kimichat/ent/telegramchatautorecapssubscribers"
)

// TelegramChatAutoRecapsSubscribers is the model entity for the TelegramChatAutoRecapsSubscribers schema.
type TelegramChatAutoRecapsSubscribers struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ChatID holds the value of the "chat_id" field.
	ChatID int64 `json:"chat_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    int64 `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TelegramChatAutoRecapsSubscribers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case telegramchatautorecapssubscribers.FieldChatID, telegramchatautorecapssubscribers.FieldUserID, telegramchatautorecapssubscribers.FieldCreatedAt, telegramchatautorecapssubscribers.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case telegramchatautorecapssubscribers.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TelegramChatAutoRecapsSubscribers fields.
func (tcars *TelegramChatAutoRecapsSubscribers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case telegramchatautorecapssubscribers.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tcars.ID = *value
			}
		case telegramchatautorecapssubscribers.FieldChatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value.Valid {
				tcars.ChatID = value.Int64
			}
		case telegramchatautorecapssubscribers.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tcars.UserID = value.Int64
			}
		case telegramchatautorecapssubscribers.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tcars.CreatedAt = value.Int64
			}
		case telegramchatautorecapssubscribers.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tcars.UpdatedAt = value.Int64
			}
		default:
			tcars.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TelegramChatAutoRecapsSubscribers.
// This includes values selected through modifiers, order, etc.
func (tcars *TelegramChatAutoRecapsSubscribers) Value(name string) (ent.Value, error) {
	return tcars.selectValues.Get(name)
}

// Update returns a builder for updating this TelegramChatAutoRecapsSubscribers.
// Note that you need to call TelegramChatAutoRecapsSubscribers.Unwrap() before calling this method if this TelegramChatAutoRecapsSubscribers
// was returned from a transaction, and the transaction was committed or rolled back.
func (tcars *TelegramChatAutoRecapsSubscribers) Update() *TelegramChatAutoRecapsSubscribersUpdateOne {
	return NewTelegramChatAutoRecapsSubscribersClient(tcars.config).UpdateOne(tcars)
}

// Unwrap unwraps the TelegramChatAutoRecapsSubscribers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tcars *TelegramChatAutoRecapsSubscribers) Unwrap() *TelegramChatAutoRecapsSubscribers {
	_tx, ok := tcars.config.driver.(*txDriver)
	if !ok {
		panic("ent: TelegramChatAutoRecapsSubscribers is not a transactional entity")
	}
	tcars.config.driver = _tx.drv
	return tcars
}

// String implements the fmt.Stringer.
func (tcars *TelegramChatAutoRecapsSubscribers) String() string {
	var builder strings.Builder
	builder.WriteString("TelegramChatAutoRecapsSubscribers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tcars.ID))
	builder.WriteString("chat_id=")
	builder.WriteString(fmt.Sprintf("%v", tcars.ChatID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", tcars.UserID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tcars.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tcars.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// TelegramChatAutoRecapsSubscribersSlice is a parsable slice of TelegramChatAutoRecapsSubscribers.
type TelegramChatAutoRecapsSubscribersSlice []*TelegramChatAutoRecapsSubscribers
