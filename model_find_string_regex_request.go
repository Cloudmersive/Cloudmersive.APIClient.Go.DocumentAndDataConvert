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
type FindStringRegexRequest struct {
	// Input text content
	TextContent string `json:"TextContent,omitempty"`
	// Target input regular expression (regex) to find
	TargetRegex string `json:"TargetRegex,omitempty"`
	// Set to True to match case, False to ignore case
	MatchCase bool `json:"MatchCase,omitempty"`
}