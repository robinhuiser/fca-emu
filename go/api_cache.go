/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A CacheApiController binds http requests to an api service and writes the service results to the http response
type CacheApiController struct {
	service CacheApiServicer
}

// NewCacheApiController creates a default api controller
func NewCacheApiController(s CacheApiServicer) Router {
	return &CacheApiController{ service: s }
}

// Routes returns all of the api route for the CacheApiController
func (c *CacheApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteAllCache",
			strings.ToUpper("Delete"),
			"/cloud/v1/admin/cache",
			c.DeleteAllCache,
		},
		{
			"DeleteCache",
			strings.ToUpper("Delete"),
			"/cloud/v1/admin/cache/{finiteType}",
			c.DeleteCache,
		},
		{
			"GetCache",
			strings.ToUpper("Get"),
			"/cloud/v1/admin/cache/{finiteType}",
			c.GetCache,
		},
		{
			"ListAllCache",
			strings.ToUpper("Get"),
			"/cloud/v1/admin/cache",
			c.ListAllCache,
		},
		{
			"RenewCache",
			strings.ToUpper("Post"),
			"/cloud/v1/admin/cache/{finiteType}",
			c.RenewCache,
		},
	}
}

// DeleteAllCache - Clear the cache for all Finite types
func (c *CacheApiController) DeleteAllCache(w http.ResponseWriter, r *http.Request) { 
	xTRACEID := r.Header.Get("X-TRACE-ID")
	result, err := c.service.DeleteAllCache(r.Context(), xTRACEID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DeleteCache - Clear the cache for a given type
func (c *CacheApiController) DeleteCache(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	finiteType := params["finiteType"]
	xTRACEID := r.Header.Get("X-TRACE-ID")
	result, err := c.service.DeleteCache(r.Context(), finiteType, xTRACEID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetCache - Get the cache statistics for a given type
func (c *CacheApiController) GetCache(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	finiteType := params["finiteType"]
	xTRACEID := r.Header.Get("X-TRACE-ID")
	result, err := c.service.GetCache(r.Context(), finiteType, xTRACEID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ListAllCache - List of all cache statistics
func (c *CacheApiController) ListAllCache(w http.ResponseWriter, r *http.Request) { 
	xTRACEID := r.Header.Get("X-TRACE-ID")
	result, err := c.service.ListAllCache(r.Context(), xTRACEID)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// RenewCache - Clear and renew the cache
func (c *CacheApiController) RenewCache(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	finiteType := params["finiteType"]
	xTRACEID := r.Header.Get("X-TRACE-ID")
	cacheReference := &CacheReference{}
	if err := json.NewDecoder(r.Body).Decode(&cacheReference); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.RenewCache(r.Context(), finiteType, xTRACEID, *cacheReference)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
