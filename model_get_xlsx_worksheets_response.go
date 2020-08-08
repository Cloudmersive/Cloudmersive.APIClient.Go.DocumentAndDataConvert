/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of running a Get-Worksheets command
type GetXlsxWorksheetsResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Worksheets in the Excel XLSX spreadsheet
	Worksheets []XlsxWorksheet `json:"Worksheets,omitempty"`
}