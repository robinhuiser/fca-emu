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

// TaxInformation - Tax information
type TaxInformation struct {
	Type string `json:"type,omitempty"`

	// Tax identification value
	TaxId string `json:"taxId,omitempty"`
}
