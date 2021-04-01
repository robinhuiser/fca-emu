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

// ExchangeTransactionResult - The resulting information from attempting to perform an Exchange Transaction
type ExchangeTransactionResult struct {

	// The status of the attempt
	Status string `json:"status,omitempty"`

	// A readable reason for the status provided.
	Reason string `json:"reason,omitempty"`

	// The newly created exchange transaction identifier
	ExchangeTransactionId string `json:"exchangeTransactionId,omitempty"`
}