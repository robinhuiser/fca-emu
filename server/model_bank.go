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

// Bank - Bank object
type Bank struct {

	// Standard format of Business Identifier Codes (bic), unique identification code
	Swift string `json:"swift,omitempty"`

	// A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.
	BankCode string `json:"bankCode,omitempty"`

	// The financial institution name
	BankName string `json:"bankName,omitempty"`

	// The branch code of a bank branch helps in distinguishing one branch from another
	BranchCode string `json:"branchCode,omitempty"`

	URI FiniteUri `json:"URI,omitempty"`
}
