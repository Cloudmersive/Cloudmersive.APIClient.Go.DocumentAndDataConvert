/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Active document editing session result.  To retrieve the document, use the Finish Editing API.
type DocumentEditingEditSession struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Document editing session URL; in-memory temporary cache link with TTL (Time-to-Live expiration) of 30 minutes.  To retrieve the document, use the Finish Editing API.
	EditSessionURL string `json:"EditSessionURL,omitempty"`
}
