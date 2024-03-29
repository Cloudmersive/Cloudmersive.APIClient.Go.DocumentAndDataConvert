/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Response from an HTML template application
type HtmlTemplateApplicationResponse struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Final HTML result of all operations on input
	FinalHtml string `json:"FinalHtml,omitempty"`
}
