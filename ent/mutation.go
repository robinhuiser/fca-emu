// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/robinhuiser/finite-mock-server/ent/account"
	"github.com/robinhuiser/finite-mock-server/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount = "Account"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op                Op
	typ               string
	id                *uuid.UUID
	_type             *string
	number            *string
	parentId          *uuid.UUID
	name              *string
	title             *string
	iban              *string
	dateCreated       *time.Time
	dateOpened        *time.Time
	dateLastUpdated   *time.Time
	dateClosed        *time.Time
	currencyCode      *string
	status            *string
	source            *string
	interestReporting *bool
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Account, error)
	predicates        []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id uuid.UUID) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Account entities.
func (m *AccountMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *AccountMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetType sets the "type" field.
func (m *AccountMutation) SetType(s string) {
	m._type = &s
}

// GetType returns the value of the "type" field in the mutation.
func (m *AccountMutation) GetType() (r string, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *AccountMutation) ResetType() {
	m._type = nil
}

// SetNumber sets the "number" field.
func (m *AccountMutation) SetNumber(s string) {
	m.number = &s
}

// Number returns the value of the "number" field in the mutation.
func (m *AccountMutation) Number() (r string, exists bool) {
	v := m.number
	if v == nil {
		return
	}
	return *v, true
}

// OldNumber returns the old "number" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNumber: %w", err)
	}
	return oldValue.Number, nil
}

// ResetNumber resets all changes to the "number" field.
func (m *AccountMutation) ResetNumber() {
	m.number = nil
}

// SetParentId sets the "parentId" field.
func (m *AccountMutation) SetParentId(u uuid.UUID) {
	m.parentId = &u
}

// ParentId returns the value of the "parentId" field in the mutation.
func (m *AccountMutation) ParentId() (r uuid.UUID, exists bool) {
	v := m.parentId
	if v == nil {
		return
	}
	return *v, true
}

// OldParentId returns the old "parentId" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldParentId(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldParentId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldParentId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParentId: %w", err)
	}
	return oldValue.ParentId, nil
}

// ResetParentId resets all changes to the "parentId" field.
func (m *AccountMutation) ResetParentId() {
	m.parentId = nil
}

// SetName sets the "name" field.
func (m *AccountMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *AccountMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *AccountMutation) ResetName() {
	m.name = nil
}

// SetTitle sets the "title" field.
func (m *AccountMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *AccountMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *AccountMutation) ResetTitle() {
	m.title = nil
}

// SetIban sets the "iban" field.
func (m *AccountMutation) SetIban(s string) {
	m.iban = &s
}

// Iban returns the value of the "iban" field in the mutation.
func (m *AccountMutation) Iban() (r string, exists bool) {
	v := m.iban
	if v == nil {
		return
	}
	return *v, true
}

// OldIban returns the old "iban" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldIban(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldIban is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldIban requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIban: %w", err)
	}
	return oldValue.Iban, nil
}

// ResetIban resets all changes to the "iban" field.
func (m *AccountMutation) ResetIban() {
	m.iban = nil
}

// SetDateCreated sets the "dateCreated" field.
func (m *AccountMutation) SetDateCreated(t time.Time) {
	m.dateCreated = &t
}

// DateCreated returns the value of the "dateCreated" field in the mutation.
func (m *AccountMutation) DateCreated() (r time.Time, exists bool) {
	v := m.dateCreated
	if v == nil {
		return
	}
	return *v, true
}

// OldDateCreated returns the old "dateCreated" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldDateCreated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDateCreated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDateCreated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDateCreated: %w", err)
	}
	return oldValue.DateCreated, nil
}

// ResetDateCreated resets all changes to the "dateCreated" field.
func (m *AccountMutation) ResetDateCreated() {
	m.dateCreated = nil
}

// SetDateOpened sets the "dateOpened" field.
func (m *AccountMutation) SetDateOpened(t time.Time) {
	m.dateOpened = &t
}

// DateOpened returns the value of the "dateOpened" field in the mutation.
func (m *AccountMutation) DateOpened() (r time.Time, exists bool) {
	v := m.dateOpened
	if v == nil {
		return
	}
	return *v, true
}

// OldDateOpened returns the old "dateOpened" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldDateOpened(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDateOpened is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDateOpened requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDateOpened: %w", err)
	}
	return oldValue.DateOpened, nil
}

// ResetDateOpened resets all changes to the "dateOpened" field.
func (m *AccountMutation) ResetDateOpened() {
	m.dateOpened = nil
}

// SetDateLastUpdated sets the "dateLastUpdated" field.
func (m *AccountMutation) SetDateLastUpdated(t time.Time) {
	m.dateLastUpdated = &t
}

// DateLastUpdated returns the value of the "dateLastUpdated" field in the mutation.
func (m *AccountMutation) DateLastUpdated() (r time.Time, exists bool) {
	v := m.dateLastUpdated
	if v == nil {
		return
	}
	return *v, true
}

// OldDateLastUpdated returns the old "dateLastUpdated" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldDateLastUpdated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDateLastUpdated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDateLastUpdated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDateLastUpdated: %w", err)
	}
	return oldValue.DateLastUpdated, nil
}

// ResetDateLastUpdated resets all changes to the "dateLastUpdated" field.
func (m *AccountMutation) ResetDateLastUpdated() {
	m.dateLastUpdated = nil
}

// SetDateClosed sets the "dateClosed" field.
func (m *AccountMutation) SetDateClosed(t time.Time) {
	m.dateClosed = &t
}

// DateClosed returns the value of the "dateClosed" field in the mutation.
func (m *AccountMutation) DateClosed() (r time.Time, exists bool) {
	v := m.dateClosed
	if v == nil {
		return
	}
	return *v, true
}

// OldDateClosed returns the old "dateClosed" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldDateClosed(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDateClosed is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDateClosed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDateClosed: %w", err)
	}
	return oldValue.DateClosed, nil
}

// ResetDateClosed resets all changes to the "dateClosed" field.
func (m *AccountMutation) ResetDateClosed() {
	m.dateClosed = nil
}

// SetCurrencyCode sets the "currencyCode" field.
func (m *AccountMutation) SetCurrencyCode(s string) {
	m.currencyCode = &s
}

// CurrencyCode returns the value of the "currencyCode" field in the mutation.
func (m *AccountMutation) CurrencyCode() (r string, exists bool) {
	v := m.currencyCode
	if v == nil {
		return
	}
	return *v, true
}

// OldCurrencyCode returns the old "currencyCode" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldCurrencyCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCurrencyCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCurrencyCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCurrencyCode: %w", err)
	}
	return oldValue.CurrencyCode, nil
}

// ResetCurrencyCode resets all changes to the "currencyCode" field.
func (m *AccountMutation) ResetCurrencyCode() {
	m.currencyCode = nil
}

// SetStatus sets the "status" field.
func (m *AccountMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *AccountMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldStatus(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *AccountMutation) ResetStatus() {
	m.status = nil
}

// SetSource sets the "source" field.
func (m *AccountMutation) SetSource(s string) {
	m.source = &s
}

// Source returns the value of the "source" field in the mutation.
func (m *AccountMutation) Source() (r string, exists bool) {
	v := m.source
	if v == nil {
		return
	}
	return *v, true
}

// OldSource returns the old "source" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldSource(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSource is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSource requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSource: %w", err)
	}
	return oldValue.Source, nil
}

// ResetSource resets all changes to the "source" field.
func (m *AccountMutation) ResetSource() {
	m.source = nil
}

// SetInterestReporting sets the "interestReporting" field.
func (m *AccountMutation) SetInterestReporting(b bool) {
	m.interestReporting = &b
}

// InterestReporting returns the value of the "interestReporting" field in the mutation.
func (m *AccountMutation) InterestReporting() (r bool, exists bool) {
	v := m.interestReporting
	if v == nil {
		return
	}
	return *v, true
}

// OldInterestReporting returns the old "interestReporting" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldInterestReporting(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldInterestReporting is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldInterestReporting requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInterestReporting: %w", err)
	}
	return oldValue.InterestReporting, nil
}

// ResetInterestReporting resets all changes to the "interestReporting" field.
func (m *AccountMutation) ResetInterestReporting() {
	m.interestReporting = nil
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 14)
	if m._type != nil {
		fields = append(fields, account.FieldType)
	}
	if m.number != nil {
		fields = append(fields, account.FieldNumber)
	}
	if m.parentId != nil {
		fields = append(fields, account.FieldParentId)
	}
	if m.name != nil {
		fields = append(fields, account.FieldName)
	}
	if m.title != nil {
		fields = append(fields, account.FieldTitle)
	}
	if m.iban != nil {
		fields = append(fields, account.FieldIban)
	}
	if m.dateCreated != nil {
		fields = append(fields, account.FieldDateCreated)
	}
	if m.dateOpened != nil {
		fields = append(fields, account.FieldDateOpened)
	}
	if m.dateLastUpdated != nil {
		fields = append(fields, account.FieldDateLastUpdated)
	}
	if m.dateClosed != nil {
		fields = append(fields, account.FieldDateClosed)
	}
	if m.currencyCode != nil {
		fields = append(fields, account.FieldCurrencyCode)
	}
	if m.status != nil {
		fields = append(fields, account.FieldStatus)
	}
	if m.source != nil {
		fields = append(fields, account.FieldSource)
	}
	if m.interestReporting != nil {
		fields = append(fields, account.FieldInterestReporting)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldType:
		return m.GetType()
	case account.FieldNumber:
		return m.Number()
	case account.FieldParentId:
		return m.ParentId()
	case account.FieldName:
		return m.Name()
	case account.FieldTitle:
		return m.Title()
	case account.FieldIban:
		return m.Iban()
	case account.FieldDateCreated:
		return m.DateCreated()
	case account.FieldDateOpened:
		return m.DateOpened()
	case account.FieldDateLastUpdated:
		return m.DateLastUpdated()
	case account.FieldDateClosed:
		return m.DateClosed()
	case account.FieldCurrencyCode:
		return m.CurrencyCode()
	case account.FieldStatus:
		return m.Status()
	case account.FieldSource:
		return m.Source()
	case account.FieldInterestReporting:
		return m.InterestReporting()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldType:
		return m.OldType(ctx)
	case account.FieldNumber:
		return m.OldNumber(ctx)
	case account.FieldParentId:
		return m.OldParentId(ctx)
	case account.FieldName:
		return m.OldName(ctx)
	case account.FieldTitle:
		return m.OldTitle(ctx)
	case account.FieldIban:
		return m.OldIban(ctx)
	case account.FieldDateCreated:
		return m.OldDateCreated(ctx)
	case account.FieldDateOpened:
		return m.OldDateOpened(ctx)
	case account.FieldDateLastUpdated:
		return m.OldDateLastUpdated(ctx)
	case account.FieldDateClosed:
		return m.OldDateClosed(ctx)
	case account.FieldCurrencyCode:
		return m.OldCurrencyCode(ctx)
	case account.FieldStatus:
		return m.OldStatus(ctx)
	case account.FieldSource:
		return m.OldSource(ctx)
	case account.FieldInterestReporting:
		return m.OldInterestReporting(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case account.FieldNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNumber(v)
		return nil
	case account.FieldParentId:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParentId(v)
		return nil
	case account.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case account.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case account.FieldIban:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIban(v)
		return nil
	case account.FieldDateCreated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDateCreated(v)
		return nil
	case account.FieldDateOpened:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDateOpened(v)
		return nil
	case account.FieldDateLastUpdated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDateLastUpdated(v)
		return nil
	case account.FieldDateClosed:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDateClosed(v)
		return nil
	case account.FieldCurrencyCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCurrencyCode(v)
		return nil
	case account.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case account.FieldSource:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSource(v)
		return nil
	case account.FieldInterestReporting:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInterestReporting(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldType:
		m.ResetType()
		return nil
	case account.FieldNumber:
		m.ResetNumber()
		return nil
	case account.FieldParentId:
		m.ResetParentId()
		return nil
	case account.FieldName:
		m.ResetName()
		return nil
	case account.FieldTitle:
		m.ResetTitle()
		return nil
	case account.FieldIban:
		m.ResetIban()
		return nil
	case account.FieldDateCreated:
		m.ResetDateCreated()
		return nil
	case account.FieldDateOpened:
		m.ResetDateOpened()
		return nil
	case account.FieldDateLastUpdated:
		m.ResetDateLastUpdated()
		return nil
	case account.FieldDateClosed:
		m.ResetDateClosed()
		return nil
	case account.FieldCurrencyCode:
		m.ResetCurrencyCode()
		return nil
	case account.FieldStatus:
		m.ResetStatus()
		return nil
	case account.FieldSource:
		m.ResetSource()
		return nil
	case account.FieldInterestReporting:
		m.ResetInterestReporting()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Account edge %s", name)
}
