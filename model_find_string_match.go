/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Individual match result of finding a target string in a longer text string
type FindStringMatch struct {
	// 0-based index of the start of the match
	CharacterOffsetStart int32 `json:"CharacterOffsetStart,omitempty"`
	// 0-based index of the end of the match
	CharacterOffsetEnd int32 `json:"CharacterOffsetEnd,omitempty"`
	// Text content of the line containing the match
	ContainingLine string `json:"ContainingLine,omitempty"`
}
