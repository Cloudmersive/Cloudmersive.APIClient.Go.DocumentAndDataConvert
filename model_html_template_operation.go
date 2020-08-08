/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

type HtmlTemplateOperation struct {
	// Operation action to take; possible values are \"Replace\"
	Action int32 `json:"Action,omitempty"`
	// For Replace operations, the string to match against (to be replaced with ReplaceWith string)
	MatchAgsint string `json:"MatchAgsint,omitempty"`
	// For Replace operations, the string to Replace the original string with
	ReplaceWith string `json:"ReplaceWith,omitempty"`
}