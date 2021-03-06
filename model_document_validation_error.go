/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Validation error found in document
type DocumentValidationError struct {
	// Description of the error
	Description string `json:"Description,omitempty"`
	// XPath to the error
	Path string `json:"Path,omitempty"`
	// URI of the part in question
	Uri string `json:"Uri,omitempty"`
	// True if this is an error, false otherwise
	IsError bool `json:"IsError,omitempty"`
}
