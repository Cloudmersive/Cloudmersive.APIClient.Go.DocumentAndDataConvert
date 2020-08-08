/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// An MSG file attachment
type MsgAttachment struct {
	// Name of the attachment, including file extension
	Name string `json:"Name,omitempty"`
	// The MSG attachment as a byte[]
	Content string `json:"Content,omitempty"`
}
