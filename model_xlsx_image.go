/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

type XlsxImage struct {
	// The Path of the location of this object; leave blank for new rows
	Path string `json:"Path,omitempty"`
	// Read-only; internal ID for the image contents
	ImageDataEmbedId string `json:"ImageDataEmbedId,omitempty"`
	// Read-only; image data MIME content-type
	ImageDataContentType string `json:"ImageDataContentType,omitempty"`
	// Read-only; internal file name/path for the image
	ImageInternalFileName string `json:"ImageInternalFileName,omitempty"`
	// URL to the image contents; file is stored in an in-memory cache and will be deleted.  Call Finish-Editing to get the contents.
	ImageContentsURL string `json:"ImageContentsURL,omitempty"`
}
