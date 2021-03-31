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

type ErrorResponse struct {

	// Indicates whether a operation was successful
	Status bool `json:"status,omitempty"`

	// Potential information or errror messages
	Messages []string `json:"messages,omitempty"`
}
