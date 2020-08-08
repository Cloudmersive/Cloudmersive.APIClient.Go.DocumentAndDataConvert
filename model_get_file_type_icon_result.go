/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of getting a file type icon from a file extension
type GetFileTypeIconResult struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// PNG icon as a byte array
	Icon string `json:"Icon,omitempty"`
	// Extension used for the icon
	Extension string `json:"Extension,omitempty"`
}
