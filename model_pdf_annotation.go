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

// PDF Annotation details
type PdfAnnotation struct {
	// Title of the annotation; this is often the author of the annotation in Acrobat-created PDF files
	Title string `json:"Title,omitempty"`
	// Type of the annotation; possible values are Text
	AnnotationType string `json:"AnnotationType,omitempty"`
	// The 1-based index of the page containing the annotation
	PageNumber int32 `json:"PageNumber,omitempty"`
	// The 0-based index of the annotation in the document
	AnnotationIndex int32 `json:"AnnotationIndex,omitempty"`
	// Subject of the annotation
	Subject string `json:"Subject,omitempty"`
	// Text contents of the annotation
	TextContents string `json:"TextContents,omitempty"`
	// Date that the annotation was created
	CreationDate time.Time `json:"CreationDate,omitempty"`
	// Date that the annotation was last modified
	ModifiedDate time.Time `json:"ModifiedDate,omitempty"`
	// Left X coordinate for the location of the annotation
	LeftX float64 `json:"LeftX,omitempty"`
	// Top Y coordination of the location of the annotation
	TopY float64 `json:"TopY,omitempty"`
	// Width of the annotation
	Width float64 `json:"Width,omitempty"`
	// Height of the annotation
	Height float64 `json:"Height,omitempty"`
}
