// Code generated by entc, DO NOT EDIT.

package card

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/robinhuiser/fca-emu/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// StartDate applies equality check predicate on the "startDate" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// ExpiryDate applies equality check predicate on the "expiryDate" field. It's identical to ExpiryDateEQ.
func ExpiryDate(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiryDate), v))
	})
}

// HolderName applies equality check predicate on the "holderName" field. It's identical to HolderNameEQ.
func HolderName(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHolderName), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumber), v))
	})
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumber), v...))
	})
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumber), v...))
	})
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumber), v))
	})
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumber), v))
	})
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumber), v))
	})
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumber), v))
	})
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNumber), v))
	})
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNumber), v))
	})
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNumber), v))
	})
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNumber), v))
	})
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNumber), v))
	})
}

// StartDateEQ applies the EQ predicate on the "startDate" field.
func StartDateEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateNEQ applies the NEQ predicate on the "startDate" field.
func StartDateNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartDate), v))
	})
}

// StartDateIn applies the In predicate on the "startDate" field.
func StartDateIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartDate), v...))
	})
}

// StartDateNotIn applies the NotIn predicate on the "startDate" field.
func StartDateNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartDate), v...))
	})
}

// StartDateGT applies the GT predicate on the "startDate" field.
func StartDateGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartDate), v))
	})
}

// StartDateGTE applies the GTE predicate on the "startDate" field.
func StartDateGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartDate), v))
	})
}

// StartDateLT applies the LT predicate on the "startDate" field.
func StartDateLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartDate), v))
	})
}

// StartDateLTE applies the LTE predicate on the "startDate" field.
func StartDateLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartDate), v))
	})
}

// ExpiryDateEQ applies the EQ predicate on the "expiryDate" field.
func ExpiryDateEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiryDate), v))
	})
}

// ExpiryDateNEQ applies the NEQ predicate on the "expiryDate" field.
func ExpiryDateNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiryDate), v))
	})
}

// ExpiryDateIn applies the In predicate on the "expiryDate" field.
func ExpiryDateIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiryDate), v...))
	})
}

// ExpiryDateNotIn applies the NotIn predicate on the "expiryDate" field.
func ExpiryDateNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiryDate), v...))
	})
}

// ExpiryDateGT applies the GT predicate on the "expiryDate" field.
func ExpiryDateGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiryDate), v))
	})
}

// ExpiryDateGTE applies the GTE predicate on the "expiryDate" field.
func ExpiryDateGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiryDate), v))
	})
}

// ExpiryDateLT applies the LT predicate on the "expiryDate" field.
func ExpiryDateLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiryDate), v))
	})
}

// ExpiryDateLTE applies the LTE predicate on the "expiryDate" field.
func ExpiryDateLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiryDate), v))
	})
}

// HolderNameEQ applies the EQ predicate on the "holderName" field.
func HolderNameEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHolderName), v))
	})
}

// HolderNameNEQ applies the NEQ predicate on the "holderName" field.
func HolderNameNEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHolderName), v))
	})
}

// HolderNameIn applies the In predicate on the "holderName" field.
func HolderNameIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHolderName), v...))
	})
}

// HolderNameNotIn applies the NotIn predicate on the "holderName" field.
func HolderNameNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHolderName), v...))
	})
}

// HolderNameGT applies the GT predicate on the "holderName" field.
func HolderNameGT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHolderName), v))
	})
}

// HolderNameGTE applies the GTE predicate on the "holderName" field.
func HolderNameGTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHolderName), v))
	})
}

// HolderNameLT applies the LT predicate on the "holderName" field.
func HolderNameLT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHolderName), v))
	})
}

// HolderNameLTE applies the LTE predicate on the "holderName" field.
func HolderNameLTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHolderName), v))
	})
}

// HolderNameContains applies the Contains predicate on the "holderName" field.
func HolderNameContains(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHolderName), v))
	})
}

// HolderNameHasPrefix applies the HasPrefix predicate on the "holderName" field.
func HolderNameHasPrefix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHolderName), v))
	})
}

// HolderNameHasSuffix applies the HasSuffix predicate on the "holderName" field.
func HolderNameHasSuffix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHolderName), v))
	})
}

// HolderNameEqualFold applies the EqualFold predicate on the "holderName" field.
func HolderNameEqualFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHolderName), v))
	})
}

// HolderNameContainsFold applies the ContainsFold predicate on the "holderName" field.
func HolderNameContainsFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHolderName), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// HasNetwork applies the HasEdge predicate on the "network" edge.
func HasNetwork() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NetworkTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NetworkTable, NetworkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNetworkWith applies the HasEdge predicate on the "network" edge with a given conditions (other predicates).
func HasNetworkWith(preds ...predicate.CardNetwork) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NetworkInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NetworkTable, NetworkColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		p(s.Not())
	})
}
