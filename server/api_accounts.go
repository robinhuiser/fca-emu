/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abstract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// A AccountsApiController binds http requests to an api service and writes the service results to the http response
type AccountsApiController struct {
	service AccountsApiServicer
}

// NewAccountsApiController creates a default api controller
func NewAccountsApiController(s AccountsApiServicer) Router {
	return &AccountsApiController{service: s}
}

// Routes returns all of the api route for the AccountsApiController
func (c *AccountsApiController) Routes() Routes {
	return Routes{
		{
			"GetAccount",
			strings.ToUpper("Get"),
			"/cloud/v1/account/{accountId}",
			c.GetAccount,
		},
		{
			"GetAccountBalances",
			strings.ToUpper("Get"),
			"/cloud/v1/account/{accountId}/balances",
			c.GetAccountBalances,
		},
		{
			"GetAccountDetails",
			strings.ToUpper("Get"),
			"/cloud/v1/account/{accountId}/details",
			c.GetAccountDetails,
		},
		{
			"GetEntityAccountsList",
			strings.ToUpper("Get"),
			"/cloud/v1/entity/{entityId}/accounts",
			c.GetEntityAccountsList,
		},
		{
			"PostEntityAccountsList",
			strings.ToUpper("Post"),
			"/cloud/v1/entities/accounts",
			c.PostEntityAccountsList,
		},
		{
			"PutAccount",
			strings.ToUpper("Put"),
			"/cloud/v1/account/{accountId}",
			c.PutAccount,
		},
		{
			"SearchAccounts",
			strings.ToUpper("Post"),
			"/cloud/v1/accounts/search",
			c.SearchAccounts,
		},
	}
}

// GetAccount - Return a account
func (c *AccountsApiController) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	accountId := params["accountId"]
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	result, err := c.service.GetAccount(r.Context(), accountId, mask, enhance, xTRACEID, xTOKEN)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAccountBalances - Return a accounts balances
func (c *AccountsApiController) GetAccountBalances(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	accountId := params["accountId"]
	mask, _ := strconv.ParseBool(query.Get("mask"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	result, err := c.service.GetAccountBalances(r.Context(), accountId, mask, xTRACEID, xTOKEN)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAccountDetails - Return a accounts details
func (c *AccountsApiController) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	accountId := params["accountId"]
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	result, err := c.service.GetAccountDetails(r.Context(), accountId, mask, enhance, xTRACEID, xTOKEN)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetEntityAccountsList - Return list a of accounts by entity
func (c *AccountsApiController) GetEntityAccountsList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	entityId := params["entityId"]
	fields := query.Get("fields")
	limit, err := parseInt32Parameter(query.Get("limit"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cursor := query.Get("cursor")
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	result, err := c.service.GetEntityAccountsList(r.Context(), entityId, fields, limit, cursor, mask, enhance, xTRACEID, xTOKEN)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PostEntityAccountsList - Return list a of accounts based on a entity search
func (c *AccountsApiController) PostEntityAccountsList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, err := parseInt32Parameter(query.Get("limit"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cursor := query.Get("cursor")
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	searchFilter := &[]SearchFilter{}
	if err := json.NewDecoder(r.Body).Decode(&searchFilter); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.PostEntityAccountsList(r.Context(), limit, cursor, mask, enhance, xTRACEID, xTOKEN, *searchFilter)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PutAccount - Update a account
func (c *AccountsApiController) PutAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	accountId := params["accountId"]
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	account := &Account{}
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.PutAccount(r.Context(), accountId, mask, enhance, xTRACEID, xTOKEN, *account)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// SearchAccounts - Search for accounts
func (c *AccountsApiController) SearchAccounts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fields := query.Get("fields")
	limit, err := parseInt32Parameter(query.Get("limit"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cursor := query.Get("cursor")
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	searchFilter := &[]SearchFilter{}
	if err := json.NewDecoder(r.Body).Decode(&searchFilter); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.SearchAccounts(r.Context(), fields, limit, cursor, mask, enhance, xTRACEID, xTOKEN, *searchFilter)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
