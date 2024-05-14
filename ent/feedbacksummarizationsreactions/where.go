// Code generated by ent, DO NOT EDIT.

package feedbacksummarizationsreactions

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/leohearts/insights-bot-kimichat/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLTE(FieldID, id))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldChatID, v))
}

// LogID applies equality check predicate on the "log_id" field. It's identical to LogIDEQ.
func LogID(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldLogID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldUserID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldUpdatedAt, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLTE(FieldChatID, v))
}

// LogIDEQ applies the EQ predicate on the "log_id" field.
func LogIDEQ(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldLogID, v))
}

// LogIDNEQ applies the NEQ predicate on the "log_id" field.
func LogIDNEQ(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldLogID, v))
}

// LogIDIn applies the In predicate on the "log_id" field.
func LogIDIn(vs ...uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldLogID, vs...))
}

// LogIDNotIn applies the NotIn predicate on the "log_id" field.
func LogIDNotIn(vs ...uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldLogID, vs...))
}

// LogIDGT applies the GT predicate on the "log_id" field.
func LogIDGT(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGT(FieldLogID, v))
}

// LogIDGTE applies the GTE predicate on the "log_id" field.
func LogIDGTE(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGTE(FieldLogID, v))
}

// LogIDLT applies the LT predicate on the "log_id" field.
func LogIDLT(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLT(FieldLogID, v))
}

// LogIDLTE applies the LTE predicate on the "log_id" field.
func LogIDLTE(v uuid.UUID) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLTE(FieldLogID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLTE(FieldUserID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldType, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeedbackSummarizationsReactions) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeedbackSummarizationsReactions) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FeedbackSummarizationsReactions) predicate.FeedbackSummarizationsReactions {
	return predicate.FeedbackSummarizationsReactions(sql.NotPredicates(p))
}
