/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of performing a country validation operation
type ValidateCountryResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Full name of the country
	CountryFullName string `json:"CountryFullName,omitempty"`
	// Two-letter ISO 3166-1 country code
	ISOTwoLetterCode string `json:"ISOTwoLetterCode,omitempty"`
	// Two-letter FIPS 10-4 country code
	FIPSTwoLetterCode string `json:"FIPSTwoLetterCode,omitempty"`
	// Three-letter ISO 3166-1 country code
	ThreeLetterCode string `json:"ThreeLetterCode,omitempty"`
	// Time zones (IANA/Olsen) in the country
	Timezones []Timezone `json:"Timezones,omitempty"`
}
