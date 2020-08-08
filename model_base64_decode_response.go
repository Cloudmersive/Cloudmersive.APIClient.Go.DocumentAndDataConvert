/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of base 64 decoding
type Base64DecodeResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Result of performing a base 64 decode operation, binary file content
	ContentResult string `json:"ContentResult,omitempty"`
}
