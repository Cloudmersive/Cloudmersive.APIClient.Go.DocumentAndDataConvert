/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of extracting links from an HTML file
type HtmlGetLinksResponse struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// All hyperlinks in the HTML document
	Links []HtmlHyperlink `json:"Links,omitempty"`
}