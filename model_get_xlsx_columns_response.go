/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of running a Get-Columns command
type GetXlsxColumnsResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Spreadsheet Columns in the XLSX document
	Columns []XlsxSpreadsheetColumn `json:"Columns,omitempty"`
}
