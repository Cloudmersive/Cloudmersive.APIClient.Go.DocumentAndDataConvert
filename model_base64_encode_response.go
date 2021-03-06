/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of base 64 encoding
type Base64EncodeResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Result of performing a base 64 encoding operation, a text string representing the encoded original file content
	Base64TextContentResult string `json:"Base64TextContentResult,omitempty"`
}
