/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of performing a filter operation on XML input using XPath
type XmlFilterWithXPathResult struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Matching selected XML nodes as strings
	XmlNodes []string `json:"XmlNodes,omitempty"`
	// Count of the matching results
	ResultCount int32 `json:"ResultCount,omitempty"`
}
