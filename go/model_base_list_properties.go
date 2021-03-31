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

// BaseListProperties - Base properties for all lists object
type BaseListProperties struct {

	// Indicates whether a operation was successful
	Status bool `json:"status,omitempty"`

	// Potential information or errror messages
	Messages []string `json:"messages,omitempty"`

	// The total number of list items that exist for the list or query
	TotalItems int32 `json:"totalItems,omitempty"`

	// The Base64 encoded value that indicates the next cursor. The cursor will be decoded to use as the starting point to fetch the next set of results.
	NextCursor string `json:"nextCursor,omitempty"`
}
