// Code generated by ent, DO NOT EDIT.

package beacon

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"realm.pub/tavern/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldName, v))
}

// Principal applies equality check predicate on the "principal" field. It's identical to PrincipalEQ.
func Principal(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldPrincipal, v))
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldIdentifier, v))
}

// AgentIdentifier applies equality check predicate on the "agent_identifier" field. It's identical to AgentIdentifierEQ.
func AgentIdentifier(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldAgentIdentifier, v))
}

// LastSeenAt applies equality check predicate on the "last_seen_at" field. It's identical to LastSeenAtEQ.
func LastSeenAt(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldLastSeenAt, v))
}

// Interval applies equality check predicate on the "interval" field. It's identical to IntervalEQ.
func Interval(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldInterval, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContainsFold(FieldName, v))
}

// PrincipalEQ applies the EQ predicate on the "principal" field.
func PrincipalEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldPrincipal, v))
}

// PrincipalNEQ applies the NEQ predicate on the "principal" field.
func PrincipalNEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldPrincipal, v))
}

// PrincipalIn applies the In predicate on the "principal" field.
func PrincipalIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldPrincipal, vs...))
}

// PrincipalNotIn applies the NotIn predicate on the "principal" field.
func PrincipalNotIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldPrincipal, vs...))
}

// PrincipalGT applies the GT predicate on the "principal" field.
func PrincipalGT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldPrincipal, v))
}

// PrincipalGTE applies the GTE predicate on the "principal" field.
func PrincipalGTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldPrincipal, v))
}

// PrincipalLT applies the LT predicate on the "principal" field.
func PrincipalLT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldPrincipal, v))
}

// PrincipalLTE applies the LTE predicate on the "principal" field.
func PrincipalLTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldPrincipal, v))
}

// PrincipalContains applies the Contains predicate on the "principal" field.
func PrincipalContains(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContains(FieldPrincipal, v))
}

// PrincipalHasPrefix applies the HasPrefix predicate on the "principal" field.
func PrincipalHasPrefix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasPrefix(FieldPrincipal, v))
}

// PrincipalHasSuffix applies the HasSuffix predicate on the "principal" field.
func PrincipalHasSuffix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasSuffix(FieldPrincipal, v))
}

// PrincipalIsNil applies the IsNil predicate on the "principal" field.
func PrincipalIsNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldIsNull(FieldPrincipal))
}

// PrincipalNotNil applies the NotNil predicate on the "principal" field.
func PrincipalNotNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldNotNull(FieldPrincipal))
}

// PrincipalEqualFold applies the EqualFold predicate on the "principal" field.
func PrincipalEqualFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEqualFold(FieldPrincipal, v))
}

// PrincipalContainsFold applies the ContainsFold predicate on the "principal" field.
func PrincipalContainsFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContainsFold(FieldPrincipal, v))
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldIdentifier, v))
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContains(FieldIdentifier, v))
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasPrefix(FieldIdentifier, v))
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasSuffix(FieldIdentifier, v))
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEqualFold(FieldIdentifier, v))
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContainsFold(FieldIdentifier, v))
}

// AgentIdentifierEQ applies the EQ predicate on the "agent_identifier" field.
func AgentIdentifierEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldAgentIdentifier, v))
}

// AgentIdentifierNEQ applies the NEQ predicate on the "agent_identifier" field.
func AgentIdentifierNEQ(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldAgentIdentifier, v))
}

// AgentIdentifierIn applies the In predicate on the "agent_identifier" field.
func AgentIdentifierIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldAgentIdentifier, vs...))
}

// AgentIdentifierNotIn applies the NotIn predicate on the "agent_identifier" field.
func AgentIdentifierNotIn(vs ...string) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldAgentIdentifier, vs...))
}

// AgentIdentifierGT applies the GT predicate on the "agent_identifier" field.
func AgentIdentifierGT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldAgentIdentifier, v))
}

// AgentIdentifierGTE applies the GTE predicate on the "agent_identifier" field.
func AgentIdentifierGTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldAgentIdentifier, v))
}

// AgentIdentifierLT applies the LT predicate on the "agent_identifier" field.
func AgentIdentifierLT(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldAgentIdentifier, v))
}

// AgentIdentifierLTE applies the LTE predicate on the "agent_identifier" field.
func AgentIdentifierLTE(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldAgentIdentifier, v))
}

// AgentIdentifierContains applies the Contains predicate on the "agent_identifier" field.
func AgentIdentifierContains(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContains(FieldAgentIdentifier, v))
}

// AgentIdentifierHasPrefix applies the HasPrefix predicate on the "agent_identifier" field.
func AgentIdentifierHasPrefix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasPrefix(FieldAgentIdentifier, v))
}

// AgentIdentifierHasSuffix applies the HasSuffix predicate on the "agent_identifier" field.
func AgentIdentifierHasSuffix(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldHasSuffix(FieldAgentIdentifier, v))
}

// AgentIdentifierIsNil applies the IsNil predicate on the "agent_identifier" field.
func AgentIdentifierIsNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldIsNull(FieldAgentIdentifier))
}

// AgentIdentifierNotNil applies the NotNil predicate on the "agent_identifier" field.
func AgentIdentifierNotNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldNotNull(FieldAgentIdentifier))
}

// AgentIdentifierEqualFold applies the EqualFold predicate on the "agent_identifier" field.
func AgentIdentifierEqualFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldEqualFold(FieldAgentIdentifier, v))
}

// AgentIdentifierContainsFold applies the ContainsFold predicate on the "agent_identifier" field.
func AgentIdentifierContainsFold(v string) predicate.Beacon {
	return predicate.Beacon(sql.FieldContainsFold(FieldAgentIdentifier, v))
}

// LastSeenAtEQ applies the EQ predicate on the "last_seen_at" field.
func LastSeenAtEQ(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldLastSeenAt, v))
}

// LastSeenAtNEQ applies the NEQ predicate on the "last_seen_at" field.
func LastSeenAtNEQ(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldLastSeenAt, v))
}

// LastSeenAtIn applies the In predicate on the "last_seen_at" field.
func LastSeenAtIn(vs ...time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldLastSeenAt, vs...))
}

// LastSeenAtNotIn applies the NotIn predicate on the "last_seen_at" field.
func LastSeenAtNotIn(vs ...time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldLastSeenAt, vs...))
}

// LastSeenAtGT applies the GT predicate on the "last_seen_at" field.
func LastSeenAtGT(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldLastSeenAt, v))
}

// LastSeenAtGTE applies the GTE predicate on the "last_seen_at" field.
func LastSeenAtGTE(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldLastSeenAt, v))
}

// LastSeenAtLT applies the LT predicate on the "last_seen_at" field.
func LastSeenAtLT(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldLastSeenAt, v))
}

// LastSeenAtLTE applies the LTE predicate on the "last_seen_at" field.
func LastSeenAtLTE(v time.Time) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldLastSeenAt, v))
}

// LastSeenAtIsNil applies the IsNil predicate on the "last_seen_at" field.
func LastSeenAtIsNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldIsNull(FieldLastSeenAt))
}

// LastSeenAtNotNil applies the NotNil predicate on the "last_seen_at" field.
func LastSeenAtNotNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldNotNull(FieldLastSeenAt))
}

// IntervalEQ applies the EQ predicate on the "interval" field.
func IntervalEQ(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldEQ(FieldInterval, v))
}

// IntervalNEQ applies the NEQ predicate on the "interval" field.
func IntervalNEQ(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldNEQ(FieldInterval, v))
}

// IntervalIn applies the In predicate on the "interval" field.
func IntervalIn(vs ...uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldIn(FieldInterval, vs...))
}

// IntervalNotIn applies the NotIn predicate on the "interval" field.
func IntervalNotIn(vs ...uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldNotIn(FieldInterval, vs...))
}

// IntervalGT applies the GT predicate on the "interval" field.
func IntervalGT(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldGT(FieldInterval, v))
}

// IntervalGTE applies the GTE predicate on the "interval" field.
func IntervalGTE(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldGTE(FieldInterval, v))
}

// IntervalLT applies the LT predicate on the "interval" field.
func IntervalLT(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldLT(FieldInterval, v))
}

// IntervalLTE applies the LTE predicate on the "interval" field.
func IntervalLTE(v uint64) predicate.Beacon {
	return predicate.Beacon(sql.FieldLTE(FieldInterval, v))
}

// IntervalIsNil applies the IsNil predicate on the "interval" field.
func IntervalIsNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldIsNull(FieldInterval))
}

// IntervalNotNil applies the NotNil predicate on the "interval" field.
func IntervalNotNil() predicate.Beacon {
	return predicate.Beacon(sql.FieldNotNull(FieldInterval))
}

// HasHost applies the HasEdge predicate on the "host" edge.
func HasHost() predicate.Beacon {
	return predicate.Beacon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostTable, HostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostWith applies the HasEdge predicate on the "host" edge with a given conditions (other predicates).
func HasHostWith(preds ...predicate.Host) predicate.Beacon {
	return predicate.Beacon(func(s *sql.Selector) {
		step := newHostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Beacon {
	return predicate.Beacon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Beacon {
	return predicate.Beacon(func(s *sql.Selector) {
		step := newTasksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Beacon) predicate.Beacon {
	return predicate.Beacon(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Beacon) predicate.Beacon {
	return predicate.Beacon(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Beacon) predicate.Beacon {
	return predicate.Beacon(sql.NotPredicates(p))
}
