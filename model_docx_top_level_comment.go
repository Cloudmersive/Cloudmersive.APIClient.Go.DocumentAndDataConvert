/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

import (
	"time"
)

// Top-level Comment in a Word Document
type DocxTopLevelComment struct {
	// Path to the comment in the document
	Path string `json:"Path,omitempty"`
	// Author name of the comment
	Author string `json:"Author,omitempty"`
	// Initials of the author of the comment
	AuthorInitials string `json:"AuthorInitials,omitempty"`
	// Text content of the comment
	CommentText string `json:"CommentText,omitempty"`
	// Date timestamp of the comment
	CommentDate time.Time `json:"CommentDate,omitempty"`
	// Child comments, that are replies to this one
	ReplyChildComments []DocxComment `json:"ReplyChildComments,omitempty"`
	// True if this comment is marked as Done in Word, otherwise it is false
	Done bool `json:"Done,omitempty"`
}