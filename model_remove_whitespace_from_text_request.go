/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Request to remove whitespace from a string
type RemoveWhitespaceFromTextRequest struct {
	// Input text string to remove the whitespace from
	TextContainingWhitespace string `json:"TextContainingWhitespace,omitempty"`
}
