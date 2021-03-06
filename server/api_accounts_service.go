/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/account"
	"github.com/robinhuiser/fca-emu/ent/entity"
)

var clt *ent.Client

// AccountsApiService is a service that implents the logic for the AccountsApiServicer
// This service should implement the business logic for every endpoint for the AccountsApi API.
// Include any external packages or services that will be required by this service.
type AccountsApiService struct {
}

// NewAccountsApiService creates a default api service
func NewAccountsApiService(client *ent.Client) AccountsApiServicer {
	clt = client
	return &AccountsApiService{}
}

// GetAccount - Return a account
func (s *AccountsApiService) GetAccount(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Parse and verify UUID
	u, err := uuid.Parse(accountId)
	if err != nil {
		return Response(400, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Account.
		Query().
		Where(account.ID(u)).
		Only(ctx)
	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	a, err := mapAccountDetails(rs, mask, ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	return Response(200, a), nil
}

// GetAccountBalances - Return a accounts balances
func (s *AccountsApiService) GetAccountBalances(ctx context.Context, accountId string, mask bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Parse and verify UUID
	u, err := uuid.Parse(accountId)
	if err != nil {
		return Response(400, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Account.
		Query().
		Where(account.ID(u)).
		Only(ctx)

	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	b := Balances{
		AvailableBalance: rs.AvailableBalance,
		CurrentBalance:   rs.CurrentBalance,
	}

	return Response(200, b), nil
}

// GetAccountDetails - Return a accounts details
func (s *AccountsApiService) GetAccountDetails(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Parse and verify UUID
	u, err := uuid.Parse(accountId)
	if err != nil {
		return Response(400, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Account.
		Query().
		Where(account.ID(u)).
		Only(ctx)

	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	a, err := mapAccountDetails(rs, mask, ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	return Response(200, a), nil
}

// GetEntityAccountsList - Return list a of accounts by entity
func (s *AccountsApiService) GetEntityAccountsList(ctx context.Context, entityId string, fields string, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Parse and verify UUID
	u, err := uuid.Parse(entityId)
	if err != nil {
		return Response(400, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// 1: Find the entity
	rs, err := clt.Entity.
		Query().
		Where(entity.ID(u)).
		Only(ctx)
	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// 2: Find the accounts belonging to this entity
	accts, err := clt.Entity.
		QueryOwnsAccount(rs).
		Order(ent.Desc(account.FieldDateCreated)).
		All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	accounts := []Account{}
	if len(accts) > 0 {
		for _, acct := range accts {
			a, err := mapAccountDetails(acct, mask, ctx)
			if err != nil {
				return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
			}
			accounts = append(accounts, a)
		}
	}

	a := AccountsList{
		Status:     true,
		TotalItems: int32(len(accts)),
		Accounts:   accounts,
	}

	return Response(200, a), nil
}

// PostEntityAccountsList - Return list a of accounts based on a entity search
func (s *AccountsApiService) PostEntityAccountsList(ctx context.Context, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string, searchFilter []SearchFilter) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("PostEntityAccountsList method not implemented")
}

// PutAccount - Update a account
func (s *AccountsApiService) PutAccount(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string, account Account) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("PutAccount method not implemented")
}

// SearchAccounts - Search for accounts
func (s *AccountsApiService) SearchAccounts(ctx context.Context, fields string, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string, searchFilter []SearchFilter) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("SearchAccounts method not implemented")
}

// Helper functions for mapping
func mapPreferences(attributes []*ent.Preference) []Attribute {
	preferences := []Attribute{}
	for _, pref := range attributes {
		attr := Attribute{
			Name:  pref.Name,
			Value: pref.Value,
		}
		preferences = append(preferences, attr)
	}
	return preferences
}

func mapRoutingNumbers(routingnumbers []*ent.RoutingNumber) []RoutingNumber {
	rtns := []RoutingNumber{}
	for _, rtn := range routingnumbers {
		rn := RoutingNumber{
			Number: rtn.Number,
			Type:   string(rtn.Type),
		}
		rtns = append(rtns, rn)
	}
	return rtns
}

func mapAccountDetails(acc *ent.Account, mask bool, ctx context.Context) (Account, error) {
	// Retrieve the linked bank
	qba, err := acc.QueryBranch().QueryOwner().Only(ctx)
	if err != nil {
		return Account{}, fmt.Errorf("%v", err)
	}

	// Retrieve the linked branch
	qbr, err := acc.QueryBranch().Only(ctx)
	if err != nil {
		return Account{}, fmt.Errorf("%v", err)
	}

	// Retrieve the linked product
	qpr, err := acc.QueryProduct().Only(ctx)
	if err != nil {
		return Account{}, fmt.Errorf("%v", err)
	}

	// Retrieve the linked account attributes
	atrs, err := acc.QueryPreferences().All(ctx)
	if err != nil {
		return Account{}, fmt.Errorf("%v", err)
	}
	preferences := mapPreferences(atrs)

	// Retrieve routing number(s) of the account
	rtns, err := acc.QueryRoutingnumbers().All(ctx)
	if err != nil {
		return Account{}, fmt.Errorf("%v", err)
	}
	routingnumbers := mapRoutingNumbers(rtns)

	// Retrieve the owner(s) (entities) of the account
	ents, err := acc.QueryOwners().All(ctx)
	if err != nil {
		return Account{}, fmt.Errorf("%v", err)
	}
	entities := mapEntities(ents, ctx)

	a := Account{
		Id:                acc.ID.String(),
		Type:              acc.Type,
		Number:            isMasked(mask, acc.Number),
		ParentId:          isValidUUID(acc.ParentId.String()),
		Name:              ents[0].Fullname,
		Title:             acc.Title,
		DateCreated:       isValidBankDate(acc.DateCreated.Format(API_DATE_LAYOUT)),
		DateOpened:        isValidBankDate(acc.DateOpened.Format(API_DATE_LAYOUT)),
		DateLastUpdated:   isValidBankDate(acc.DateLastUpdated.Format(API_DATE_LAYOUT)),
		DateClosed:        isValidBankDate(acc.DateClosed.Format(API_DATE_LAYOUT)),
		CurrencyCode:      acc.CurrencyCode,
		Status:            acc.Status.String(),
		Source:            acc.Source,
		InterestReporting: acc.InterestReporting,
		Balances: Balances{
			AvailableBalance: acc.AvailableBalance,
			CurrentBalance:   acc.CurrentBalance,
		},
		URI: FiniteUri{
			URL: acc.URL,
		},
		Bank: Bank{
			BranchCode: qbr.BranchCode,
			BankName:   qba.BankName,
			Swift:      qba.Swift,
			URI: FiniteUri{
				URL: qba.URL,
			},
		},
		Entities:       entities,
		Preferences:    preferences,
		Routingnumbers: routingnumbers,
		Product: Product{
			Number:      qpr.Name,
			Type:        string(qpr.Type),
			TypeName:    qpr.TypeName,
			SubType:     qpr.SubType,
			SubTypeName: qpr.SubTypeName,
			URI: FiniteUri{
				URL: qpr.URL,
			},
		},
	}
	return a, nil
}
