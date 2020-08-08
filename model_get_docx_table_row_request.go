/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Input to a get DOCX table row request
type GetDocxTableRowRequest struct {
	// Optional: Bytes of the input file to operate on
	InputFileBytes string `json:"InputFileBytes,omitempty"`
	// Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public).
	InputFileUrl string `json:"InputFileUrl,omitempty"`
	// Path to the table to retrievew the row from
	TablePath string `json:"TablePath,omitempty"`
	// 0-based index of the row to retrieve (e.g. 0, 1, 2, ...) in the table
	TableRowRowIndex int32 `json:"TableRowRowIndex,omitempty"`
}
