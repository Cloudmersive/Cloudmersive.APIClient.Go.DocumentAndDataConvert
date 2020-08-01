/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of validating a postal code
type ValidatePostalCodeResponse struct {
	// True if the address is valid, false otherwise
	ValidPostalCode bool `json:"ValidPostalCode,omitempty"`
	// If valid, City corresponding to the input postal code, such as 'Walnut Creek'
	City string `json:"City,omitempty"`
	// If valid; State or province corresponding to the input postal code, such as 'CA' or 'California'
	StateOrProvince string `json:"StateOrProvince,omitempty"`
	// If the postal code is valid, the degrees latitude of the centroid of the postal code, null otherwise
	Latitude float64 `json:"Latitude,omitempty"`
	// If the postal code is valid, the degrees longitude of the centroid of the postal code, null otherwise
	Longitude float64 `json:"Longitude,omitempty"`
}
