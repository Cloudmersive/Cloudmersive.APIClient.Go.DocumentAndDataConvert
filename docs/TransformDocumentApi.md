# \TransformDocumentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransformDocumentDocxReplace**](TransformDocumentApi.md#TransformDocumentDocxReplace) | **Post** /convert/transform/docx/replace-all | Replace string in Word DOCX document
[**TransformDocumentPptxReplace**](TransformDocumentApi.md#TransformDocumentPptxReplace) | **Post** /convert/transform/pptx/replace-all | Replace string in PowerPoint PPTX presentation


# **TransformDocumentDocxReplace**
> string TransformDocumentDocxReplace(ctx, matchString, replaceString, optional)
Replace string in Word DOCX document

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

# **TransformDocumentPptxReplace**
> string TransformDocumentPptxReplace(ctx, matchString, replaceString, optional)
Replace string in PowerPoint PPTX presentation

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

