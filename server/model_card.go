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

// Card - Card object
type Card struct {

	// The type of card
	Type string `json:"type,omitempty"`

	// Unique card identifier in the system of record
	Id string `json:"id,omitempty"`

	// Card number (masked if masked=true)
	Number string `json:"number,omitempty"`

	// ISO 6801 date of when the card is started/activated
	StartDate string `json:"startDate,omitempty"`

	// ISO 6801 date of when the card expires
	ExpiryDate string `json:"expiryDate,omitempty"`

	// The name of the cardholder
	HolderName string `json:"holderName,omitempty"`

	// The card network program
	Network string `json:"network,omitempty"`

	// The status of the card
	Status string `json:"status,omitempty"`

	URI FiniteUri `json:"URI,omitempty"`
}
