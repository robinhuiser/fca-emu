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

// CacheReference - Reference to items in cache used to refresh/evict/renew the cache
type CacheReference struct {

	// List of attributes to locate and reference items in cache
	Attributes []Attribute `json:"attributes,omitempty"`
}
