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

// Result of an autodetect/get-info operation
type AutodetectGetInfoResult struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Detected file extension of the file format, with a leading period
	DetectedFileExtension string `json:"DetectedFileExtension,omitempty"`
	// MIME type of this file extension
	DetectedMimeType string `json:"DetectedMimeType,omitempty"`
	// Number of pages in a page-based document; for presentations, this is the number of slides and for a spreadsheet this is the number of worksheets.  Contains 0 when the page count cannot be determined, or if the concept of page count does not apply (e.g. for an image)
	PageCount int64 `json:"PageCount,omitempty"`
	// User name of the creator/author of the document, if available, null if not available
	Author string `json:"Author,omitempty"`
	// The timestamp that the document was last modified, if available, null if not available
	DateModified time.Time `json:"DateModified,omitempty"`
	// Alternate file type options and their probability
	AlternateFileTypeCandidates []AlternateFileFormatCandidate `json:"AlternateFileTypeCandidates,omitempty"`
}