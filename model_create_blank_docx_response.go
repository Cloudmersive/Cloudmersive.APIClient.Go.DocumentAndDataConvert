/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of creating a blank Word document
type CreateBlankDocxResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// URL to the edited XLSX file; file is stored in an in-memory cache and will be deleted.  Call Finish-Editing to get the result document contents.
	EditedDocumentURL string `json:"EditedDocumentURL,omitempty"`
}
