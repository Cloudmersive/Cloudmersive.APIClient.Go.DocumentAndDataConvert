# \EditTextApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditTextBase64Decode**](EditTextApi.md#EditTextBase64Decode) | **Post** /convert/edit/text/encoding/base64/decode | Base 64 decode, convert base 64 string to binary content
[**EditTextBase64Detect**](EditTextApi.md#EditTextBase64Detect) | **Post** /convert/edit/text/encoding/base64/detect | Detect, check if text string is base 64 encoded
[**EditTextBase64Encode**](EditTextApi.md#EditTextBase64Encode) | **Post** /convert/edit/text/encoding/base64/encode | Base 64 encode, convert binary or file data to a text string
[**EditTextChangeLineEndings**](EditTextApi.md#EditTextChangeLineEndings) | **Post** /convert/edit/text/line-endings/change | Set, change line endings of a text file
[**EditTextDetectLineEndings**](EditTextApi.md#EditTextDetectLineEndings) | **Post** /convert/edit/text/line-endings/detect | Detect line endings of a text file
[**EditTextFindRegex**](EditTextApi.md#EditTextFindRegex) | **Post** /convert/edit/text/find/regex | Find a regular expression regex in text input
[**EditTextFindSimple**](EditTextApi.md#EditTextFindSimple) | **Post** /convert/edit/text/find/string | Find a string in text input
[**EditTextRemoveAllWhitespace**](EditTextApi.md#EditTextRemoveAllWhitespace) | **Post** /convert/edit/text/remove/whitespace/all | Remove whitespace from text string
[**EditTextRemoveHtml**](EditTextApi.md#EditTextRemoveHtml) | **Post** /convert/edit/text/remove/html | Remove HTML from text string
[**EditTextReplaceRegex**](EditTextApi.md#EditTextReplaceRegex) | **Post** /convert/edit/text/replace/regex | Replace a string in text with a regex regular expression string
[**EditTextReplaceSimple**](EditTextApi.md#EditTextReplaceSimple) | **Post** /convert/edit/text/replace/string | Replace a string in text with another string value
[**EditTextTextEncodingDetect**](EditTextApi.md#EditTextTextEncodingDetect) | **Post** /convert/edit/text/encoding/detect | Detect text encoding of file
[**EditTextTrimWhitespace**](EditTextApi.md#EditTextTrimWhitespace) | **Post** /convert/edit/text/remove/whitespace/trim | Trim leading and trailing whitespace from text string


# **EditTextBase64Decode**
> Base64DecodeResponse EditTextBase64Decode(ctx, request)
Base 64 decode, convert base 64 string to binary content

Decodes / converts base 64 UTF-8 text string to binary content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**Base64DecodeRequest**](Base64DecodeRequest.md)| Input request | 

### Return type

[**Base64DecodeResponse**](Base64DecodeResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextBase64Detect**
> Base64DetectResponse EditTextBase64Detect(ctx, request)
Detect, check if text string is base 64 encoded

Checks, detects if input string is base 64 encoded

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**Base64DetectRequest**](Base64DetectRequest.md)| Input request | 

### Return type

[**Base64DetectResponse**](Base64DetectResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextBase64Encode**
> Base64EncodeResponse EditTextBase64Encode(ctx, request)
Base 64 encode, convert binary or file data to a text string

Encodes / converts binary or file data to a text string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**Base64EncodeRequest**](Base64EncodeRequest.md)| Input request | 

### Return type

[**Base64EncodeResponse**](Base64EncodeResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextChangeLineEndings**
> ChangeLineEndingResponse EditTextChangeLineEndings(ctx, lineEndingType, inputFile)
Set, change line endings of a text file

Sets the line ending type of a text file; set to Windows, Unix or Mac.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lineEndingType** | **string**| Required; &#39;Windows&#39; will use carriage return and line feed, &#39;Unix&#39; will use newline, and &#39;Mac&#39; will use carriage return | 
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**ChangeLineEndingResponse**](ChangeLineEndingResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextDetectLineEndings**
> DetectLineEndingsResponse EditTextDetectLineEndings(ctx, inputFile)
Detect line endings of a text file

Detect line ending type (Windows, Unix or Mac) of an input file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**DetectLineEndingsResponse**](DetectLineEndingsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextFindRegex**
> FindStringRegexResponse EditTextFindRegex(ctx, request)
Find a regular expression regex in text input

Find all occurrences of the input regular expression in the input content, and returns the matches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**FindStringRegexRequest**](FindStringRegexRequest.md)| Input request | 

### Return type

[**FindStringRegexResponse**](FindStringRegexResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextFindSimple**
> FindStringSimpleResponse EditTextFindSimple(ctx, request)
Find a string in text input

Finds all occurrences of the input string in the input content, and returns the matches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**FindStringSimpleRequest**](FindStringSimpleRequest.md)| Input request | 

### Return type

[**FindStringSimpleResponse**](FindStringSimpleResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextRemoveAllWhitespace**
> RemoveWhitespaceFromTextResponse EditTextRemoveAllWhitespace(ctx, request)
Remove whitespace from text string

Removes all whitespace from text, leaving behind only non-whitespace characters.  Whitespace includes newlines, spaces and other whitespace characters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**RemoveWhitespaceFromTextRequest**](RemoveWhitespaceFromTextRequest.md)| Input request | 

### Return type

[**RemoveWhitespaceFromTextResponse**](RemoveWhitespaceFromTextResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextRemoveHtml**
> RemoveHtmlFromTextResponse EditTextRemoveHtml(ctx, request)
Remove HTML from text string

Removes HTML from text, leaving behind only text.  Formatted text will become plain text.  Important for protecting against HTML and Cross-Site-Scripting attacks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**RemoveHtmlFromTextRequest**](RemoveHtmlFromTextRequest.md)| Input request | 

### Return type

[**RemoveHtmlFromTextResponse**](RemoveHtmlFromTextResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextReplaceRegex**
> ReplaceStringRegexResponse EditTextReplaceRegex(ctx, request)
Replace a string in text with a regex regular expression string

Replaces all occurrences of the input regular expression regex string in the input content, and returns the result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ReplaceStringRegexRequest**](ReplaceStringRegexRequest.md)| Input request | 

### Return type

[**ReplaceStringRegexResponse**](ReplaceStringRegexResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextReplaceSimple**
> ReplaceStringSimpleResponse EditTextReplaceSimple(ctx, request)
Replace a string in text with another string value

Replaces all occurrences of the input string in the input content, and returns the result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ReplaceStringSimpleRequest**](ReplaceStringSimpleRequest.md)| Input request | 

### Return type

[**ReplaceStringSimpleResponse**](ReplaceStringSimpleResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextTextEncodingDetect**
> TextEncodingDetectResponse EditTextTextEncodingDetect(ctx, inputFile)
Detect text encoding of file

Checks text encoding of file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**TextEncodingDetectResponse**](TextEncodingDetectResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTextTrimWhitespace**
> RemoveWhitespaceFromTextResponse EditTextTrimWhitespace(ctx, request)
Trim leading and trailing whitespace from text string

Trim leading and trailing whitespace from text, leaving behind a trimmed string.  Whitespace includes newlines, spaces and other whitespace characters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**RemoveWhitespaceFromTextRequest**](RemoveWhitespaceFromTextRequest.md)| Input request | 

### Return type

[**RemoveWhitespaceFromTextResponse**](RemoveWhitespaceFromTextResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

