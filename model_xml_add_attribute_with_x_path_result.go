/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of performing a add attribute node operation on XML input using XPath
type XmlAddAttributeWithXPathResult struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Resulting, modified XML document
	ResultingXmlDocument string `json:"ResultingXmlDocument,omitempty"`
	// Count of the matching results
	NodesEditedCount int32 `json:"NodesEditedCount,omitempty"`
}
