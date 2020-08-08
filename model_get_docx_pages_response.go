/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of getting pages from a Word Document DOCX
type GetDocxPagesResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Pages in the document
	Pages []DocxPage `json:"Pages,omitempty"`
	// Count of pages
	PageCount int32 `json:"PageCount,omitempty"`
}
