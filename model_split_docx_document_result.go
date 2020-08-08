/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// The result of splitting a Word document into individual Word DOCX pages
type SplitDocxDocumentResult struct {
	ResultDocuments []SplitDocumentResult `json:"ResultDocuments,omitempty"`
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
}
