/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Threat assessment for a hyperlink URL
type HtmlThreatLink struct {
	// URL of the link
	LinkUrl string `json:"LinkUrl,omitempty"`
	// Threat assessment level; possible values are None, Low, Medium and High
	ThreatLevel string `json:"ThreatLevel,omitempty"`
}