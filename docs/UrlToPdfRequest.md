# UrlToPdfRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL address of the website to screenshot.  HTTP and HTTPS are both supported, as are custom ports. | [optional] [default to null]
**ExtraLoadingWait** | **int32** | Optional: Additional number of milliseconds to wait once the web page has finished loading before taking the screenshot.  Can be helpful for highly asynchronous websites.  Provide a value of 0 for the default of 5000 milliseconds (5 seconds). Maximum is 20000 milliseconds (20 seconds). | [optional] [default to null]
**IncludeBackgroundGraphics** | **bool** | Optional: Set to true to include background graphics in the PDF, or false to not include.  Default is true. | [optional] [default to null]
**ScaleFactor** | **int32** | Optional: Set to 100 to scale at 100%, set to 50% to scale down to 50% scale, set to 200% to scale up to 200% scale, etc.  Default is 100%. Maximum is 1000% | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


