/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

import (
	"time"
)

// Result of a WHOIS operation
type WhoisResponse struct {
	// True if the domain is valid, false if it is not
	ValidDomain bool `json:"ValidDomain,omitempty"`
	// Name of the domain registrant
	RegistrantName string `json:"RegistrantName,omitempty"`
	// Organization name of the domain registrant
	RegistrantOrganization string `json:"RegistrantOrganization,omitempty"`
	// Email address of the domain registrant
	RegistrantEmail string `json:"RegistrantEmail,omitempty"`
	// Street number of the address of the domain registrant, if available
	RegistrantStreetNumber string `json:"RegistrantStreetNumber,omitempty"`
	// Street name of the address of the domain registrant, if available
	RegistrantStreet string `json:"RegistrantStreet,omitempty"`
	// City of the domain registrant, if available
	RegistrantCity string `json:"RegistrantCity,omitempty"`
	// State or Province of the address of the domain registrant, if available
	RegistrantStateOrProvince string `json:"RegistrantStateOrProvince,omitempty"`
	// Postal code of the address of the domain registrant, if available
	RegistrantPostalCode string `json:"RegistrantPostalCode,omitempty"`
	// Country of the address of the domain registrant, if available
	RegistrantCountry string `json:"RegistrantCountry,omitempty"`
	// Raw address string of the domain registrant, if available
	RegistrantRawAddress string `json:"RegistrantRawAddress,omitempty"`
	// Telephone number of the address of the domain registrant
	RegistrantTelephone string `json:"RegistrantTelephone,omitempty"`
	// Server used to lookup WHOIS information (may change based on lookup).
	WhoisServer string `json:"WhoisServer,omitempty"`
	// WHOIS raw text record
	RawTextRecord string `json:"RawTextRecord,omitempty"`
	// Creation date for the record
	CreatedDt time.Time `json:"CreatedDt,omitempty"`
}
