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

// RoutingNumber - Routing Number object
type RoutingNumber struct {

	// The type of routing number
	Type string `json:"type,omitempty"`

	// The routing number value
	Number string `json:"number,omitempty"`
}
