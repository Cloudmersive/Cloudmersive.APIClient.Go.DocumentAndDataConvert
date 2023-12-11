/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// A single PowerPoint PPTX file corresponding to one slide in the original presentation
type PresentationResult struct {
	// Worksheet number of the converted page, starting with 1 for the left-most worksheet
	SlideNumber int32 `json:"SlideNumber,omitempty"`
	// URL to the PPTX file of this slide; file is stored in an in-memory cache and will be deleted
	URL string `json:"URL,omitempty"`
	// Contents of the presentation in bytes
	PresentationContents string `json:"PresentationContents,omitempty"`
}
