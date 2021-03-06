/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// A single line of a Text document
type TextDocumentLine struct {
	// The 1-based line index of the line
	LineNumber int32 `json:"LineNumber,omitempty"`
	// The text contents of a single line of a text file
	LineContents string `json:"LineContents,omitempty"`
}
