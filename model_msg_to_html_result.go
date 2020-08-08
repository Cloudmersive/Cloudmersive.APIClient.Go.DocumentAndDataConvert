/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of converting an MSG file to HTML string
type MsgToHtmlResult struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// An HTML string version of the MSG file
	Content string `json:"Content,omitempty"`
	// The main body of the MSG file's email as an HTML string
	Body string `json:"Body,omitempty"`
	// The From sender of the MSG file's email
	From string `json:"From,omitempty"`
	// The To recipients of the MSG file's email
	To string `json:"To,omitempty"`
	// The CC recipients of the MSG file's email
	Cc string `json:"Cc,omitempty"`
	// The time that the MSG file's email was received
	ReceivedTime string `json:"ReceivedTime,omitempty"`
	// The subject of the MSG file's email
	Subject string `json:"Subject,omitempty"`
	// List of all attachments for the MSG file
	Attachments []MsgAttachment `json:"Attachments,omitempty"`
}