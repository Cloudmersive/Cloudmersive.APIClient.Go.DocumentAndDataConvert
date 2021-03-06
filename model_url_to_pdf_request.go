/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Request to convert a URL to a PDF file
type UrlToPdfRequest struct {
	// URL address of the website to screenshot.  HTTP and HTTPS are both supported, as are custom ports.
	Url string `json:"Url,omitempty"`
	// Optional: Additional number of milliseconds to wait once the web page has finished loading before taking the screenshot.  Can be helpful for highly asynchronous websites.  Provide a value of 0 for the default of 5000 milliseconds (5 seconds). Maximum is 20000 milliseconds (20 seconds).
	ExtraLoadingWait int32 `json:"ExtraLoadingWait,omitempty"`
	// Optional: Set to true to include background graphics in the PDF, or false to not include.  Default is true.
	IncludeBackgroundGraphics bool `json:"IncludeBackgroundGraphics,omitempty"`
	// Optional: Set to 100 to scale at 100%, set to 50% to scale down to 50% scale, set to 200% to scale up to 200% scale, etc.  Default is 100%. Maximum is 1000%
	ScaleFactor int32 `json:"ScaleFactor,omitempty"`
}
