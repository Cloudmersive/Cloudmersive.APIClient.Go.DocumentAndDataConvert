/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Row in an Excel spreadsheet worksheet
type XlsxSpreadsheetRow struct {
	// The Path of the location of this object; leave blank for new rows
	Path string `json:"Path,omitempty"`
	// Spreadsheet Cells in the spreadsheet row
	Cells []XlsxSpreadsheetCell `json:"Cells,omitempty"`
}