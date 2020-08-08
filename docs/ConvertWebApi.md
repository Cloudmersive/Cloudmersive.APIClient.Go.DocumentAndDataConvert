# \ConvertWebApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertWebHtmlToDocx**](ConvertWebApi.md#ConvertWebHtmlToDocx) | **Post** /convert/html/to/docx | Convert HTML to Word DOCX Document
[**ConvertWebHtmlToPdf**](ConvertWebApi.md#ConvertWebHtmlToPdf) | **Post** /convert/web/html/to/pdf | Convert HTML string to PDF
[**ConvertWebHtmlToPng**](ConvertWebApi.md#ConvertWebHtmlToPng) | **Post** /convert/web/html/to/png | Convert HTML string to PNG screenshot
[**ConvertWebHtmlToTxt**](ConvertWebApi.md#ConvertWebHtmlToTxt) | **Post** /convert/web/html/to/txt | Convert HTML string to text (txt)
[**ConvertWebMdToHtml**](ConvertWebApi.md#ConvertWebMdToHtml) | **Post** /convert/web/md/to/html | Convert Markdown to HTML
[**ConvertWebUrlToPdf**](ConvertWebApi.md#ConvertWebUrlToPdf) | **Post** /convert/web/url/to/pdf | Convert a URL to PDF
[**ConvertWebUrlToScreenshot**](ConvertWebApi.md#ConvertWebUrlToScreenshot) | **Post** /convert/web/url/to/screenshot | Take screenshot of URL
[**ConvertWebUrlToTxt**](ConvertWebApi.md#ConvertWebUrlToTxt) | **Post** /convert/web/url/to/txt | Convert website URL page to text (txt)


# **ConvertWebHtmlToDocx**
> string ConvertWebHtmlToDocx(ctx, inputRequest)
Convert HTML to Word DOCX Document

Convert HTML to Office Word Document (DOCX) format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputRequest** | [**HtmlToOfficeRequest**](HtmlToOfficeRequest.md)| HTML input to convert to DOCX | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebHtmlToPdf**
> string ConvertWebHtmlToPdf(ctx, input)
Convert HTML string to PDF

Fully renders a website and returns a PDF of the HTML.  Javascript, HTML5, CSS and other advanced features are all supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**HtmlToPdfRequest**](HtmlToPdfRequest.md)| HTML to PDF request parameters | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebHtmlToPng**
> string ConvertWebHtmlToPng(ctx, input)
Convert HTML string to PNG screenshot

Fully renders a website and returns a PNG (screenshot) of the HTML.  Javascript, HTML5, CSS and other advanced features are all supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**HtmlToPngRequest**](HtmlToPngRequest.md)| HTML to PNG request parameters | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebHtmlToTxt**
> HtmlToTextResponse ConvertWebHtmlToTxt(ctx, input)
Convert HTML string to text (txt)

Converts an HTML string input into text (txt); extracts text from HTML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**HtmlToTextRequest**](HtmlToTextRequest.md)| HTML to Text request parameters | 

### Return type

[**HtmlToTextResponse**](HtmlToTextResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebMdToHtml**
> HtmlMdResult ConvertWebMdToHtml(ctx, inputFile)
Convert Markdown to HTML

Convert a markdown file (.md) to HTML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**HtmlMdResult**](HtmlMdResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebUrlToPdf**
> string ConvertWebUrlToPdf(ctx, input)
Convert a URL to PDF

Fully renders a website and returns a PDF of the full page.  Javascript, HTML5, CSS and other advanced features are all supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**UrlToPdfRequest**](UrlToPdfRequest.md)| URL to PDF request parameters | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebUrlToScreenshot**
> string ConvertWebUrlToScreenshot(ctx, input)
Take screenshot of URL

Fully renders a website and returns a PNG screenshot of the full page image.  Javascript, HTML5, CSS and other advanced features are all supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ScreenshotRequest**](ScreenshotRequest.md)| Screenshot request parameters | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertWebUrlToTxt**
> UrlToTextResponse ConvertWebUrlToTxt(ctx, input)
Convert website URL page to text (txt)

Converts a website URL page into text (txt); extracts text from HTML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**UrlToTextRequest**](UrlToTextRequest.md)| HTML to Text request parameters | 

### Return type

[**UrlToTextResponse**](UrlToTextResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

