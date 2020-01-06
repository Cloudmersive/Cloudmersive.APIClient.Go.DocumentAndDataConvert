/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of validating a URL with full validation
type ValidateUrlResponseFull struct {
	// True if the URL has valid syntax, a valid domain, a valid endpoint, and passes virus scan checks; false otherwise
	ValidURL bool `json:"ValidURL,omitempty"`
	// True if the URL has valid syntax, false otherwise
	ValidSyntax bool `json:"Valid_Syntax,omitempty"`
	// True if the domain name is valid and exists, false otherwise
	ValidDomain bool `json:"Valid_Domain,omitempty"`
	// True if the endpoint is up and responsive and passes a virus scan check, false otherwise
	ValidEndpoint bool `json:"Valid_Endpoint,omitempty"`
	// Well-formed version of the URL
	WellFormedURL string `json:"WellFormedURL,omitempty"`
}
