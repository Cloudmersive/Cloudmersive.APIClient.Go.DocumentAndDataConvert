/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Style in an Excel spreadsheet
type DocxCellStyle struct {
	// The Path of the location of this object; leave blank for new rows
	Path string `json:"Path,omitempty"`
	// Name of the style
	Name string `json:"Name,omitempty"`
	// Format ID of the cell style
	FormatID int32 `json:"FormatID,omitempty"`
	// Built=in ID of the cell style
	BuiltInID int32 `json:"BuiltInID,omitempty"`
}