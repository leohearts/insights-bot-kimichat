// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/leohearts/insights-bot-kimichat/ent/telegramchatrecapsoptions"
)

// TelegramChatRecapsOptions is the model entity for the TelegramChatRecapsOptions schema.
type TelegramChatRecapsOptions struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ChatID holds the value of the "chat_id" field.
	ChatID int64 `json:"chat_id,omitempty"`
	// AutoRecapSendMode holds the value of the "auto_recap_send_mode" field.
	AutoRecapSendMode int `json:"auto_recap_send_mode,omitempty"`
	// ManualRecapRatePerSeconds holds the value of the "manual_recap_rate_per_seconds" field.
	ManualRecapRatePerSeconds int64 `json:"manual_recap_rate_per_seconds,omitempty"`
	// AutoRecapRatesPerDay holds the value of the "auto_recap_rates_per_day" field.
	AutoRecapRatesPerDay int `json:"auto_recap_rates_per_day,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    int64 `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TelegramChatRecapsOptions) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case telegramchatrecapsoptions.FieldChatID, telegramchatrecapsoptions.FieldAutoRecapSendMode, telegramchatrecapsoptions.FieldManualRecapRatePerSeconds, telegramchatrecapsoptions.FieldAutoRecapRatesPerDay, telegramchatrecapsoptions.FieldCreatedAt, telegramchatrecapsoptions.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case telegramchatrecapsoptions.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TelegramChatRecapsOptions fields.
func (tcro *TelegramChatRecapsOptions) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case telegramchatrecapsoptions.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tcro.ID = *value
			}
		case telegramchatrecapsoptions.FieldChatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value.Valid {
				tcro.ChatID = value.Int64
			}
		case telegramchatrecapsoptions.FieldAutoRecapSendMode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auto_recap_send_mode", values[i])
			} else if value.Valid {
				tcro.AutoRecapSendMode = int(value.Int64)
			}
		case telegramchatrecapsoptions.FieldManualRecapRatePerSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field manual_recap_rate_per_seconds", values[i])
			} else if value.Valid {
				tcro.ManualRecapRatePerSeconds = value.Int64
			}
		case telegramchatrecapsoptions.FieldAutoRecapRatesPerDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auto_recap_rates_per_day", values[i])
			} else if value.Valid {
				tcro.AutoRecapRatesPerDay = int(value.Int64)
			}
		case telegramchatrecapsoptions.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tcro.CreatedAt = value.Int64
			}
		case telegramchatrecapsoptions.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tcro.UpdatedAt = value.Int64
			}
		default:
			tcro.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TelegramChatRecapsOptions.
// This includes values selected through modifiers, order, etc.
func (tcro *TelegramChatRecapsOptions) Value(name string) (ent.Value, error) {
	return tcro.selectValues.Get(name)
}

// Update returns a builder for updating this TelegramChatRecapsOptions.
// Note that you need to call TelegramChatRecapsOptions.Unwrap() before calling this method if this TelegramChatRecapsOptions
// was returned from a transaction, and the transaction was committed or rolled back.
func (tcro *TelegramChatRecapsOptions) Update() *TelegramChatRecapsOptionsUpdateOne {
	return NewTelegramChatRecapsOptionsClient(tcro.config).UpdateOne(tcro)
}

// Unwrap unwraps the TelegramChatRecapsOptions entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tcro *TelegramChatRecapsOptions) Unwrap() *TelegramChatRecapsOptions {
	_tx, ok := tcro.config.driver.(*txDriver)
	if !ok {
		panic("ent: TelegramChatRecapsOptions is not a transactional entity")
	}
	tcro.config.driver = _tx.drv
	return tcro
}

// String implements the fmt.Stringer.
func (tcro *TelegramChatRecapsOptions) String() string {
	var builder strings.Builder
	builder.WriteString("TelegramChatRecapsOptions(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tcro.ID))
	builder.WriteString("chat_id=")
	builder.WriteString(fmt.Sprintf("%v", tcro.ChatID))
	builder.WriteString(", ")
	builder.WriteString("auto_recap_send_mode=")
	builder.WriteString(fmt.Sprintf("%v", tcro.AutoRecapSendMode))
	builder.WriteString(", ")
	builder.WriteString("manual_recap_rate_per_seconds=")
	builder.WriteString(fmt.Sprintf("%v", tcro.ManualRecapRatePerSeconds))
	builder.WriteString(", ")
	builder.WriteString("auto_recap_rates_per_day=")
	builder.WriteString(fmt.Sprintf("%v", tcro.AutoRecapRatesPerDay))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tcro.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tcro.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// TelegramChatRecapsOptionsSlice is a parsable slice of TelegramChatRecapsOptions.
type TelegramChatRecapsOptionsSlice []*TelegramChatRecapsOptions
