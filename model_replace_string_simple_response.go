/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of replacing a string
type ReplaceStringSimpleResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Result of performing a replace string operation
	TextContentResult string `json:"TextContentResult,omitempty"`
}
