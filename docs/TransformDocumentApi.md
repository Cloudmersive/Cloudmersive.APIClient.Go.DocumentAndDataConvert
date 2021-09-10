# \TransformDocumentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransformDocumentDocxReplace**](TransformDocumentApi.md#TransformDocumentDocxReplace) | **Post** /convert/transform/docx/replace-all | Replace string in Word DOCX document, return result
[**TransformDocumentDocxReplaceEditSession**](TransformDocumentApi.md#TransformDocumentDocxReplaceEditSession) | **Post** /convert/transform/docx/replace-all/edit-session | Replace string in Word DOCX document, return edit session
[**TransformDocumentDocxTableFillIn**](TransformDocumentApi.md#TransformDocumentDocxTableFillIn) | **Post** /convert/transform/docx/table/fill/data | Fill in data in a table in a Word DOCX document, return result
[**TransformDocumentDocxTableFillInEditSession**](TransformDocumentApi.md#TransformDocumentDocxTableFillInEditSession) | **Post** /convert/transform/docx/table/fill/data/edit-session | Fill in data in a table in a Word DOCX document, return edit session
[**TransformDocumentDocxTableFillInMulti**](TransformDocumentApi.md#TransformDocumentDocxTableFillInMulti) | **Post** /convert/transform/docx/table/fill/data/multi | Fill in data in multiple tables in a Word DOCX document, return result
[**TransformDocumentPptxReplace**](TransformDocumentApi.md#TransformDocumentPptxReplace) | **Post** /convert/transform/pptx/replace-all | Replace string in PowerPoint PPTX presentation, return result


# **TransformDocumentDocxReplace**
> string TransformDocumentDocxReplace(ctx, matchString, replaceString, optional)
Replace string in Word DOCX document, return result

Replace all instances of a string in an Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchString** | **string**| String to search for and match against, to be replaced | 
  **replaceString** | **string**| String to replace the matched values with | 
 **optional** | ***TransformDocumentDocxReplaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransformDocumentDocxReplaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API (part of EditDocumentApi) to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public). | 
 **matchCase** | **optional.Bool**| Optional: True if the case should be matched, false for case insensitive match. Default is false. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransformDocumentDocxReplaceEditSession**
> DocumentTransformEditSession TransformDocumentDocxReplaceEditSession(ctx, matchString, replaceString, optional)
Replace string in Word DOCX document, return edit session

Replace all instances of a string in an Office Word Document (docx).  Returns an edit session URL so that you can chain together multiple edit operations without having to send the entire document contents back and forth multiple times.  Call the Finish Editing API to retrieve the final document once editing is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchString** | **string**| String to search for and match against, to be replaced | 
  **replaceString** | **string**| String to replace the matched values with | 
 **optional** | ***TransformDocumentDocxReplaceEditSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransformDocumentDocxReplaceEditSessionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API (part of EditDocumentApi) to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public). | 
 **matchCase** | **optional.Bool**| Optional: True if the case should be matched, false for case insensitive match. Default is false. | 

### Return type

[**DocumentTransformEditSession**](DocumentTransformEditSession.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransformDocumentDocxTableFillIn**
> string TransformDocumentDocxTableFillIn(ctx, request)
Fill in data in a table in a Word DOCX document, return result

Replace placeholder rows ina  table in an Office Word Document (docx) using one or more templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**DocxTableTableFillRequest**](DocxTableTableFillRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransformDocumentDocxTableFillInEditSession**
> DocumentTransformEditSession TransformDocumentDocxTableFillInEditSession(ctx, request)
Fill in data in a table in a Word DOCX document, return edit session

Replace placeholder rows ina  table in an Office Word Document (docx) using one or more templates.  Returns an edit session URL so that you can chain together multiple edit operations without having to send the entire document contents back and forth multiple times.  Call the Finish Editing API to retrieve the final document once editing is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**DocxTableTableFillRequest**](DocxTableTableFillRequest.md)|  | 

### Return type

[**DocumentTransformEditSession**](DocumentTransformEditSession.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransformDocumentDocxTableFillInMulti**
> string TransformDocumentDocxTableFillInMulti(ctx, request)
Fill in data in multiple tables in a Word DOCX document, return result

Replace placeholder rows in multiple tables in an Office Word Document (docx) using one or more templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**DocxTableTableFillMultiRequest**](DocxTableTableFillMultiRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransformDocumentPptxReplace**
> string TransformDocumentPptxReplace(ctx, matchString, replaceString, optional)
Replace string in PowerPoint PPTX presentation, return result

Replace all instances of a string in an Office PowerPoint Document (pptx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchString** | **string**| String to search for and match against, to be replaced | 
  **replaceString** | **string**| String to replace the matched values with | 
 **optional** | ***TransformDocumentPptxReplaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransformDocumentPptxReplaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile** | **optional.Interface of *os.File**| Optional: Input file to perform the operation on. | 
 **inputFileUrl** | **optional.String**| Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API (part of EditDocumentApi) to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public). | 
 **matchCase** | **optional.Bool**| Optional: True if the case should be matched, false for case insensitive match. Default is false. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

