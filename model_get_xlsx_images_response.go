/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of running a Get-Images command
type GetXlsxImagesResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Spreadsheet Images in the XLSX document
	Images []XlsxImage `json:"Images,omitempty"`
}
