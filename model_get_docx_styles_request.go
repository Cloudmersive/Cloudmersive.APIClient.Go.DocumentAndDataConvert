/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Input to a Get Tables request
type GetDocxStylesRequest struct {
	// Optional: Bytes of the input file to operate on
	InputFileBytes string `json:"InputFileBytes,omitempty"`
	// Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public).
	InputFileUrl string `json:"InputFileUrl,omitempty"`
}
