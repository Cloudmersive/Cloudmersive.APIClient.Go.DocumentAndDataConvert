/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Request to find a string in a string
type FindStringSimpleRequest struct {
	// Input text content
	TextContent string `json:"TextContent,omitempty"`
	// Target input string to find
	TargetString string `json:"TargetString,omitempty"`
}
