/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of converting an EML input to a PNG array
type EmlToPngResult struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Array of converted pages
	PngResultPages []ConvertedPngPage `json:"PngResultPages,omitempty"`
}
