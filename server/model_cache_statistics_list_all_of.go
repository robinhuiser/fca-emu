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

type CacheStatisticsListAllOf struct {

	// The list of cache statistics
	Caches []CacheStatistics `json:"caches,omitempty"`
}
