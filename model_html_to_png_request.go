/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Details of the HTML to PNG request
type HtmlToPngRequest struct {
	// HTML to render to PNG (screenshot)
	Html string `json:"Html,omitempty"`
	// Optional: Additional number of milliseconds to wait once the web page has finished loading before taking the screenshot.  Can be helpful for highly asynchronous websites. Provide a value of 0 for the default of 5000 milliseconds (5 seconds). Maximum is 30000 milliseconds (30 seconds).
	ExtraLoadingWait int32 `json:"ExtraLoadingWait,omitempty"`
	// Optional: Width of the screenshot in pixels; supply 0 to default to 1280 x 1024, supply -1 to measure the full screen height of the page and attempt to take a screen-height screenshot
	ScreenshotWidth int32 `json:"ScreenshotWidth,omitempty"`
	// Optional: Height of the screenshot in pixels; supply 0 to default to 1280 x 1024, supply -1 to measure the full screen height of the page and attempt to take a screen-height screenshot
	ScreenshotHeight int32 `json:"ScreenshotHeight,omitempty"`
}