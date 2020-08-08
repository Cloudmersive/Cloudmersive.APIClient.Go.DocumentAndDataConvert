/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Input to a Set Cell in XLSX Worksheets request
type SetXlsxCellRequest struct {
	// Optional: Bytes of the input file to operate on
	InputFileBytes string `json:"InputFileBytes,omitempty"`
	// Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public).
	InputFileUrl string `json:"InputFileUrl,omitempty"`
	// Optional; Worksheet (tab) within the spreadsheet to update; leave blank to default to the first worksheet
	WorksheetToUpdate *XlsxWorksheet `json:"WorksheetToUpdate,omitempty"`
	// 0-based index of the row, 0, 1, 2, ... to set
	RowIndex int32 `json:"RowIndex,omitempty"`
	// 0-based index of the cell, 0, 1, 2, ... in the row to set
	CellIndex int32 `json:"CellIndex,omitempty"`
	// New Cell value to update/overwrite into the Excel XLSX spreadsheet
	CellValue *XlsxSpreadsheetCell `json:"CellValue,omitempty"`
}
