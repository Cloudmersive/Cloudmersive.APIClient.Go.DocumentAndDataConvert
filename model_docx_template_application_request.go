/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Word DOCX template application request
type DocxTemplateApplicationRequest struct {
	// Operations to apply to this template
	Operations []DocxTemplateOperation `json:"Operations,omitempty"`
}