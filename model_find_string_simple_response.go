/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of finding a string
type FindStringSimpleResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Found matches
	Matches []FindStringMatch `json:"Matches,omitempty"`
	// The number of matches
	MatchCount int32 `json:"MatchCount,omitempty"`
}
