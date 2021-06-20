# \EditHtmlApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditHtmlHtmlAppendHeaderTag**](EditHtmlApi.md#EditHtmlHtmlAppendHeaderTag) | **Post** /convert/edit/html/head/append/tag | Append an HTML tag to the HEAD section of an HTML Document
[**EditHtmlHtmlAppendHeading**](EditHtmlApi.md#EditHtmlHtmlAppendHeading) | **Post** /convert/edit/html/append/heading | Append a Heading to an HTML Document
[**EditHtmlHtmlAppendImageFromUrl**](EditHtmlApi.md#EditHtmlHtmlAppendImageFromUrl) | **Post** /convert/edit/html/append/image/from-url | Append an Image to an HTML Document from a URL
[**EditHtmlHtmlAppendImageInline**](EditHtmlApi.md#EditHtmlHtmlAppendImageInline) | **Post** /convert/edit/html/append/image/inline | Append a Base64 Inline Image to an HTML Document
[**EditHtmlHtmlAppendParagraph**](EditHtmlApi.md#EditHtmlHtmlAppendParagraph) | **Post** /convert/edit/html/append/paragraph | Append a Paragraph to an HTML Document
[**EditHtmlHtmlCreateBlankDocument**](EditHtmlApi.md#EditHtmlHtmlCreateBlankDocument) | **Post** /convert/edit/html/create/blank | Create a Blank HTML Document
[**EditHtmlHtmlGetLanguage**](EditHtmlApi.md#EditHtmlHtmlGetLanguage) | **Post** /convert/edit/html/head/get/language | Gets the language for the HTML document
[**EditHtmlHtmlGetLinks**](EditHtmlApi.md#EditHtmlHtmlGetLinks) | **Post** /convert/edit/html/extract/links | Extract resolved link URLs from HTML File
[**EditHtmlHtmlGetRelCanonical**](EditHtmlApi.md#EditHtmlHtmlGetRelCanonical) | **Post** /convert/edit/html/head/get/rel-canonical-url | Gets the rel canonical URL for the HTML document
[**EditHtmlHtmlGetSitemap**](EditHtmlApi.md#EditHtmlHtmlGetSitemap) | **Post** /convert/edit/html/head/get/sitemap-url | Gets the sitemap URL for the HTML document
[**EditHtmlHtmlSetLanguage**](EditHtmlApi.md#EditHtmlHtmlSetLanguage) | **Post** /convert/edit/html/head/set/language | Sets the language for the HTML document
[**EditHtmlHtmlSetRelCanonical**](EditHtmlApi.md#EditHtmlHtmlSetRelCanonical) | **Post** /convert/edit/html/head/set/rel-canonical-url | Sets the rel canonical URL for the HTML document
[**EditHtmlHtmlSetSitemapUrl**](EditHtmlApi.md#EditHtmlHtmlSetSitemapUrl) | **Post** /convert/edit/html/head/set/sitemap-url | Sets the sitemap URL for the HTML document


# **EditHtmlHtmlAppendHeaderTag**
> string EditHtmlHtmlAppendHeaderTag(ctx, htmlTag, optional)
Append an HTML tag to the HEAD section of an HTML Document

Appends an HTML tag to the HEAD section of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **htmlTag** | **string**| The HTML tag to append. | 
 **optional** | ***EditHtmlHtmlAppendHeaderTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlAppendHeaderTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlAppendHeading**
> string EditHtmlHtmlAppendHeading(ctx, headingText, optional)
Append a Heading to an HTML Document

Appends a heading to the end of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **headingText** | **string**| The text content to be used in the header. | 
 **optional** | ***EditHtmlHtmlAppendHeadingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlAppendHeadingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 
 **headingSize** | **optional.Int32**| Optional: The heading size number. Default is 1. Accepts values between 1 and 6. | 
 **cssStyle** | **optional.String**| Optional: The CSS style for the heading. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlAppendImageFromUrl**
> string EditHtmlHtmlAppendImageFromUrl(ctx, imageUrl, optional)
Append an Image to an HTML Document from a URL

Appends an image to the end of an HTML document using a URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **imageUrl** | **string**| The URL for the image. | 
 **optional** | ***EditHtmlHtmlAppendImageFromUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlAppendImageFromUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 
 **cssStyle** | **optional.String**| Optional: CSS style for the image. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlAppendImageInline**
> string EditHtmlHtmlAppendImageInline(ctx, optional)
Append a Base64 Inline Image to an HTML Document

Appends a base64 inline image to the end of an HTML document from an input file or URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditHtmlHtmlAppendImageInlineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlAppendImageInlineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 
 **imageFile** | **optional.Interface of *os.File**| Optional: Image file to be appended as base64 inline image. | 
 **imageUrl** | **optional.String**| Optional: Image URL to be appended as base64 inline image. | 
 **cssStyle** | **optional.String**| Optional: CSS style for the image. | 
 **imageExtension** | **optional.String**| Optional: The extension (JPG, PNG, GIF, etc.) of the image file. Recommended if uploading an imageFile directly, instead of using imageUrl. If no extension can be determined, will default to JPG. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlAppendParagraph**
> string EditHtmlHtmlAppendParagraph(ctx, paragraphText, optional)
Append a Paragraph to an HTML Document

Appends a paragraph to the end of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paragraphText** | **string**| The text content to be used in the paragraph. | 
 **optional** | ***EditHtmlHtmlAppendParagraphOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlAppendParagraphOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 
 **cssStyle** | **optional.String**| Optional: The CSS style for the paragraph. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlCreateBlankDocument**
> string EditHtmlHtmlCreateBlankDocument(ctx, optional)
Create a Blank HTML Document

Returns a blank HTML Document format file.  The file is blank, with no contents by default.  Use the optional input parameters to add various starting elements.  Use additional editing commands such as Append Header, Append Paragraph or Append Image from URL to populate the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditHtmlHtmlCreateBlankDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlCreateBlankDocumentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.String**| Optional: The title of the HTML document | 
 **cssUrl** | **optional.String**| Optional: A CSS style URL to be added to the document. | 
 **cssInline** | **optional.String**| Optional: An inline CSS style to be added to the document. | 
 **javascriptUrl** | **optional.String**| Optional: Javascript URL to be added to the document. | 
 **javascriptInline** | **optional.String**| Optional: Inline Javascript to be added to the document. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlGetLanguage**
> HtmlGetLanguageResult EditHtmlHtmlGetLanguage(ctx, optional)
Gets the language for the HTML document

Retrieves the language code (e.g. \"en\" or \"de\") of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditHtmlHtmlGetLanguageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlGetLanguageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

[**HtmlGetLanguageResult**](HtmlGetLanguageResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlGetLinks**
> HtmlGetLinksResponse EditHtmlHtmlGetLinks(ctx, optional)
Extract resolved link URLs from HTML File

Extracts the resolved link URLs, fully-qualified if possible, from an input HTML file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditHtmlHtmlGetLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlGetLinksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 
 **baseUrl** | **optional.String**| Optional: Base URL of the page, such as https://mydomain.com | 

### Return type

[**HtmlGetLinksResponse**](HtmlGetLinksResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlGetRelCanonical**
> HtmlGetRelCanonicalUrlResult EditHtmlHtmlGetRelCanonical(ctx, optional)
Gets the rel canonical URL for the HTML document

Gets the rel canonical URL of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditHtmlHtmlGetRelCanonicalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlGetRelCanonicalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

[**HtmlGetRelCanonicalUrlResult**](HtmlGetRelCanonicalUrlResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlGetSitemap**
> HtmlGetSitemapUrlResult EditHtmlHtmlGetSitemap(ctx, optional)
Gets the sitemap URL for the HTML document

Gets the sitemap link URL of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditHtmlHtmlGetSitemapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlGetSitemapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

[**HtmlGetSitemapUrlResult**](HtmlGetSitemapUrlResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlSetLanguage**
> string EditHtmlHtmlSetLanguage(ctx, languageCode, optional)
Sets the language for the HTML document

Sets the language code of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **languageCode** | **string**| The HTML langauge code to set. | 
 **optional** | ***EditHtmlHtmlSetLanguageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlSetLanguageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlSetRelCanonical**
> string EditHtmlHtmlSetRelCanonical(ctx, canonicalUrl, optional)
Sets the rel canonical URL for the HTML document

Sets the rel canonical URL of an HTML document.  This is useful for telling search engines and other indexers which pages are duplicates of eachother; any pages with the rel=canonical tag will be treated as duplicates of the canonical URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **canonicalUrl** | **string**| The HTML canonical URL to set. | 
 **optional** | ***EditHtmlHtmlSetRelCanonicalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlSetRelCanonicalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditHtmlHtmlSetSitemapUrl**
> string EditHtmlHtmlSetSitemapUrl(ctx, sitemapUrl, optional)
Sets the sitemap URL for the HTML document

Sets the sitemap URL of an HTML document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sitemapUrl** | **string**| The HTML sitemap URL to set. | 
 **optional** | ***EditHtmlHtmlSetSitemapUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditHtmlHtmlSetSitemapUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

