/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Details of the screenshot request
type ScreenshotRequest struct {
	// URL address of the website to screenshot.  HTTP and HTTPS are both supported, as are custom ports.
	Url string `json:"Url,omitempty"`
	// Optional: Additional number of milliseconds to wait once the web page has finished loading before taking the screenshot.  Can be helpful for highly asynchronous websites.  Provide a value of 0 for the default of 5000 milliseconds (5 seconds). Maximum is 20000 milliseconds (20 seconds).
	ExtraLoadingWait int32 `json:"ExtraLoadingWait,omitempty"`
	// Optional: Width of the screenshot in pixels; supply 0 to default to 1280 x 1024
	ScreenshotWidth int32 `json:"ScreenshotWidth,omitempty"`
	// Optional: Height of the screenshot in pixels; supply 0 to default to 1280 x 1024, supply -1 to measure the full screen height of the page and attempt to take a screen-height screenshot
	ScreenshotHeight int32 `json:"ScreenshotHeight,omitempty"`
}
