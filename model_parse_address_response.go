/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of parsing an address into its component parts
type ParseAddressResponse struct {
	// True if the parsing operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// The name of the building, house or structure if applicable, such as \"Cloudmersive Building 2\".  This will often by null.
	Building string `json:"Building,omitempty"`
	// The street number or house number of the address.  For example, in the address \"1600 Pennsylvania Avenue NW\" the street number would be \"1600\".  This value will typically be populated for most addresses.
	StreetNumber string `json:"StreetNumber,omitempty"`
	// The name of the street or road of the address.  For example, in the address \"1600 Pennsylvania Avenue NW\" the street number would be \"Pennsylvania Avenue NW\".
	Street string `json:"Street,omitempty"`
	// The city of the address.
	City string `json:"City,omitempty"`
	// The state or province of the address.
	StateOrProvince string `json:"StateOrProvince,omitempty"`
	// The postal code or zip code of the address.
	PostalCode string `json:"PostalCode,omitempty"`
	// Country of the address, if present in the address.  If not included in the address it will be null.
	CountryFullName string `json:"CountryFullName,omitempty"`
	// Two-letter ISO 3166-1 country code
	ISOTwoLetterCode string `json:"ISOTwoLetterCode,omitempty"`
}