// Code generated by entc, DO NOT EDIT.

package branch

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/robinhuiser/fca-emu/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BranchCode applies equality check predicate on the "branchCode" field. It's identical to BranchCodeEQ.
func BranchCode(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBranchCode), v))
	})
}

// StreetNumber applies equality check predicate on the "streetNumber" field. It's identical to StreetNumberEQ.
func StreetNumber(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStreetNumber), v))
	})
}

// StreetName applies equality check predicate on the "streetName" field. It's identical to StreetNameEQ.
func StreetName(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStreetName), v))
	})
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// Zip applies equality check predicate on the "zip" field. It's identical to ZipEQ.
func Zip(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldZip), v))
	})
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// BranchCodeEQ applies the EQ predicate on the "branchCode" field.
func BranchCodeEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBranchCode), v))
	})
}

// BranchCodeNEQ applies the NEQ predicate on the "branchCode" field.
func BranchCodeNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBranchCode), v))
	})
}

// BranchCodeIn applies the In predicate on the "branchCode" field.
func BranchCodeIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBranchCode), v...))
	})
}

// BranchCodeNotIn applies the NotIn predicate on the "branchCode" field.
func BranchCodeNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBranchCode), v...))
	})
}

// BranchCodeGT applies the GT predicate on the "branchCode" field.
func BranchCodeGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBranchCode), v))
	})
}

// BranchCodeGTE applies the GTE predicate on the "branchCode" field.
func BranchCodeGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBranchCode), v))
	})
}

// BranchCodeLT applies the LT predicate on the "branchCode" field.
func BranchCodeLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBranchCode), v))
	})
}

// BranchCodeLTE applies the LTE predicate on the "branchCode" field.
func BranchCodeLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBranchCode), v))
	})
}

// BranchCodeContains applies the Contains predicate on the "branchCode" field.
func BranchCodeContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBranchCode), v))
	})
}

// BranchCodeHasPrefix applies the HasPrefix predicate on the "branchCode" field.
func BranchCodeHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBranchCode), v))
	})
}

// BranchCodeHasSuffix applies the HasSuffix predicate on the "branchCode" field.
func BranchCodeHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBranchCode), v))
	})
}

// BranchCodeEqualFold applies the EqualFold predicate on the "branchCode" field.
func BranchCodeEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBranchCode), v))
	})
}

// BranchCodeContainsFold applies the ContainsFold predicate on the "branchCode" field.
func BranchCodeContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBranchCode), v))
	})
}

// StreetNumberEQ applies the EQ predicate on the "streetNumber" field.
func StreetNumberEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberNEQ applies the NEQ predicate on the "streetNumber" field.
func StreetNumberNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberIn applies the In predicate on the "streetNumber" field.
func StreetNumberIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStreetNumber), v...))
	})
}

// StreetNumberNotIn applies the NotIn predicate on the "streetNumber" field.
func StreetNumberNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStreetNumber), v...))
	})
}

// StreetNumberGT applies the GT predicate on the "streetNumber" field.
func StreetNumberGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberGTE applies the GTE predicate on the "streetNumber" field.
func StreetNumberGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberLT applies the LT predicate on the "streetNumber" field.
func StreetNumberLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberLTE applies the LTE predicate on the "streetNumber" field.
func StreetNumberLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberContains applies the Contains predicate on the "streetNumber" field.
func StreetNumberContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberHasPrefix applies the HasPrefix predicate on the "streetNumber" field.
func StreetNumberHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberHasSuffix applies the HasSuffix predicate on the "streetNumber" field.
func StreetNumberHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberEqualFold applies the EqualFold predicate on the "streetNumber" field.
func StreetNumberEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStreetNumber), v))
	})
}

// StreetNumberContainsFold applies the ContainsFold predicate on the "streetNumber" field.
func StreetNumberContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStreetNumber), v))
	})
}

// StreetNameEQ applies the EQ predicate on the "streetName" field.
func StreetNameEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStreetName), v))
	})
}

// StreetNameNEQ applies the NEQ predicate on the "streetName" field.
func StreetNameNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStreetName), v))
	})
}

// StreetNameIn applies the In predicate on the "streetName" field.
func StreetNameIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStreetName), v...))
	})
}

// StreetNameNotIn applies the NotIn predicate on the "streetName" field.
func StreetNameNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStreetName), v...))
	})
}

// StreetNameGT applies the GT predicate on the "streetName" field.
func StreetNameGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStreetName), v))
	})
}

// StreetNameGTE applies the GTE predicate on the "streetName" field.
func StreetNameGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStreetName), v))
	})
}

// StreetNameLT applies the LT predicate on the "streetName" field.
func StreetNameLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStreetName), v))
	})
}

// StreetNameLTE applies the LTE predicate on the "streetName" field.
func StreetNameLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStreetName), v))
	})
}

// StreetNameContains applies the Contains predicate on the "streetName" field.
func StreetNameContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStreetName), v))
	})
}

// StreetNameHasPrefix applies the HasPrefix predicate on the "streetName" field.
func StreetNameHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStreetName), v))
	})
}

// StreetNameHasSuffix applies the HasSuffix predicate on the "streetName" field.
func StreetNameHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStreetName), v))
	})
}

// StreetNameEqualFold applies the EqualFold predicate on the "streetName" field.
func StreetNameEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStreetName), v))
	})
}

// StreetNameContainsFold applies the ContainsFold predicate on the "streetName" field.
func StreetNameContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStreetName), v))
	})
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCity), v))
	})
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCity), v...))
	})
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCity), v...))
	})
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCity), v))
	})
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCity), v))
	})
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCity), v))
	})
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCity), v))
	})
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCity), v))
	})
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCity), v))
	})
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCity), v))
	})
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCity), v))
	})
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCity), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// ZipEQ applies the EQ predicate on the "zip" field.
func ZipEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldZip), v))
	})
}

// ZipNEQ applies the NEQ predicate on the "zip" field.
func ZipNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldZip), v))
	})
}

// ZipIn applies the In predicate on the "zip" field.
func ZipIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldZip), v...))
	})
}

// ZipNotIn applies the NotIn predicate on the "zip" field.
func ZipNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldZip), v...))
	})
}

// ZipGT applies the GT predicate on the "zip" field.
func ZipGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldZip), v))
	})
}

// ZipGTE applies the GTE predicate on the "zip" field.
func ZipGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldZip), v))
	})
}

// ZipLT applies the LT predicate on the "zip" field.
func ZipLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldZip), v))
	})
}

// ZipLTE applies the LTE predicate on the "zip" field.
func ZipLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldZip), v))
	})
}

// ZipContains applies the Contains predicate on the "zip" field.
func ZipContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldZip), v))
	})
}

// ZipHasPrefix applies the HasPrefix predicate on the "zip" field.
func ZipHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldZip), v))
	})
}

// ZipHasSuffix applies the HasSuffix predicate on the "zip" field.
func ZipHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldZip), v))
	})
}

// ZipEqualFold applies the EqualFold predicate on the "zip" field.
func ZipEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldZip), v))
	})
}

// ZipContainsFold applies the ContainsFold predicate on the "zip" field.
func ZipContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldZip), v))
	})
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatitude), v))
	})
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatitude), v...))
	})
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatitude), v...))
	})
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatitude), v))
	})
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatitude), v))
	})
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatitude), v))
	})
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatitude), v))
	})
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongitude), v))
	})
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLongitude), v...))
	})
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLongitude), v...))
	})
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongitude), v))
	})
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongitude), v))
	})
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongitude), v))
	})
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongitude), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Bank) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Branch) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Branch) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func Not(p predicate.Branch) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		p(s.Not())
	})
}
