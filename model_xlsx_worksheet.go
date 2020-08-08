/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// A worksheet (tab) in an Excel (XLSX) spreadsheet
type XlsxWorksheet struct {
	// The Path of the location of this object; leave blank for new worksheets
	Path string `json:"Path,omitempty"`
	// User-facing name of the worksheet tab
	WorksheetName string `json:"WorksheetName,omitempty"`
}