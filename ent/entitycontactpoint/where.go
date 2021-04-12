// Code generated by entc, DO NOT EDIT.

package entitycontactpoint

import (
	"entgo.io/ent/dialect/sql"
	"github.com/robinhuiser/finite-mock-server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
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
func IDGT(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Prefix applies equality check predicate on the "prefix" field. It's identical to PrefixEQ.
func Prefix(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrefix), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Suffix applies equality check predicate on the "suffix" field. It's identical to SuffixEQ.
func Suffix(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSuffix), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// PrefixEQ applies the EQ predicate on the "prefix" field.
func PrefixEQ(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrefix), v))
	})
}

// PrefixNEQ applies the NEQ predicate on the "prefix" field.
func PrefixNEQ(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrefix), v))
	})
}

// PrefixIn applies the In predicate on the "prefix" field.
func PrefixIn(vs ...int) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrefix), v...))
	})
}

// PrefixNotIn applies the NotIn predicate on the "prefix" field.
func PrefixNotIn(vs ...int) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrefix), v...))
	})
}

// PrefixGT applies the GT predicate on the "prefix" field.
func PrefixGT(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrefix), v))
	})
}

// PrefixGTE applies the GTE predicate on the "prefix" field.
func PrefixGTE(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrefix), v))
	})
}

// PrefixLT applies the LT predicate on the "prefix" field.
func PrefixLT(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrefix), v))
	})
}

// PrefixLTE applies the LTE predicate on the "prefix" field.
func PrefixLTE(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrefix), v))
	})
}

// PrefixIsNil applies the IsNil predicate on the "prefix" field.
func PrefixIsNil() predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrefix)))
	})
}

// PrefixNotNil applies the NotNil predicate on the "prefix" field.
func PrefixNotNil() predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrefix)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
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
func TypeNotIn(vs ...Type) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// SuffixEQ applies the EQ predicate on the "suffix" field.
func SuffixEQ(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSuffix), v))
	})
}

// SuffixNEQ applies the NEQ predicate on the "suffix" field.
func SuffixNEQ(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSuffix), v))
	})
}

// SuffixIn applies the In predicate on the "suffix" field.
func SuffixIn(vs ...int) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSuffix), v...))
	})
}

// SuffixNotIn applies the NotIn predicate on the "suffix" field.
func SuffixNotIn(vs ...int) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSuffix), v...))
	})
}

// SuffixGT applies the GT predicate on the "suffix" field.
func SuffixGT(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSuffix), v))
	})
}

// SuffixGTE applies the GTE predicate on the "suffix" field.
func SuffixGTE(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSuffix), v))
	})
}

// SuffixLT applies the LT predicate on the "suffix" field.
func SuffixLT(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSuffix), v))
	})
}

// SuffixLTE applies the LTE predicate on the "suffix" field.
func SuffixLTE(v int) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSuffix), v))
	})
}

// SuffixIsNil applies the IsNil predicate on the "suffix" field.
func SuffixIsNil() predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSuffix)))
	})
}

// SuffixNotNil applies the NotNil predicate on the "suffix" field.
func SuffixNotNil() predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSuffix)))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.EntityContactPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EntityContactPoint) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EntityContactPoint) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
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
func Not(p predicate.EntityContactPoint) predicate.EntityContactPoint {
	return predicate.EntityContactPoint(func(s *sql.Selector) {
		p(s.Not())
	})
}
