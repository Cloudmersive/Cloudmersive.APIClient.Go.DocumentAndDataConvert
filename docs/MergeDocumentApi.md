# \MergeDocumentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MergeDocumentDocx**](MergeDocumentApi.md#MergeDocumentDocx) | **Post** /convert/merge/docx | Merge Two Word DOCX Together
[**MergeDocumentDocxMulti**](MergeDocumentApi.md#MergeDocumentDocxMulti) | **Post** /convert/merge/docx/multi | Merge Multple Word DOCX Together
[**MergeDocumentHtml**](MergeDocumentApi.md#MergeDocumentHtml) | **Post** /convert/merge/html | Merge Two HTML (HTM) Files Together
[**MergeDocumentHtmlMulti**](MergeDocumentApi.md#MergeDocumentHtmlMulti) | **Post** /convert/merge/html/multi | Merge Multple HTML (HTM) Files Together
[**MergeDocumentPdf**](MergeDocumentApi.md#MergeDocumentPdf) | **Post** /convert/merge/pdf | Merge Two PDF Files Together
[**MergeDocumentPdfMulti**](MergeDocumentApi.md#MergeDocumentPdfMulti) | **Post** /convert/merge/pdf/multi | Merge Multple PDF Files Together
[**MergeDocumentPng**](MergeDocumentApi.md#MergeDocumentPng) | **Post** /convert/merge/png/vertical | Merge Two PNG Files Together
[**MergeDocumentPngMulti**](MergeDocumentApi.md#MergeDocumentPngMulti) | **Post** /convert/merge/png/vertical/multi | Merge Multple PNG Files Together
[**MergeDocumentPptx**](MergeDocumentApi.md#MergeDocumentPptx) | **Post** /convert/merge/pptx | Merge Two PowerPoint PPTX Together
[**MergeDocumentPptxMulti**](MergeDocumentApi.md#MergeDocumentPptxMulti) | **Post** /convert/merge/pptx/multi | Merge Multple PowerPoint PPTX Together
[**MergeDocumentTxt**](MergeDocumentApi.md#MergeDocumentTxt) | **Post** /convert/merge/txt | Merge Two Text (TXT) Files Together
[**MergeDocumentTxtMulti**](MergeDocumentApi.md#MergeDocumentTxtMulti) | **Post** /convert/merge/txt/multi | Merge Multple Text (TXT) Files Together
[**MergeDocumentXlsx**](MergeDocumentApi.md#MergeDocumentXlsx) | **Post** /convert/merge/xlsx | Merge Two Excel XLSX Together
[**MergeDocumentXlsxMulti**](MergeDocumentApi.md#MergeDocumentXlsxMulti) | **Post** /convert/merge/xlsx/multi | Merge Multple Excel XLSX Together


# **MergeDocumentDocx**
> string MergeDocumentDocx(ctx, inputFile1, inputFile2)
Merge Two Word DOCX Together

Combine two Office Word Documents (docx) into one single Office Word document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentDocxMulti**
> string MergeDocumentDocxMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple Word DOCX Together

Combine multiple Office Word Documents (docx) into one single Office Word document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentDocxMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentDocxMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentHtml**
> string MergeDocumentHtml(ctx, inputFile1, inputFile2)
Merge Two HTML (HTM) Files Together

Combine two HTML (.HTM) files into a single text document, preserving the order of the input documents in the combined document by stacking them vertically.  The title will be taken from the first document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentHtmlMulti**
> string MergeDocumentHtmlMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple HTML (HTM) Files Together

Combine multiple HTML (.HTM) files into a single text document, preserving the order of the input documents in the combined document by stacking them vertically.  The title will be taken from the first document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentHtmlMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentHtmlMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentPdf**
> string MergeDocumentPdf(ctx, inputFile1, inputFile2)
Merge Two PDF Files Together

Combine two PDF files (pdf) into a single PDF document, preserving the order of the input documents in the combined document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentPdfMulti**
> string MergeDocumentPdfMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple PDF Files Together

Combine multiple PDF files (pdf) into a single PDF document, preserving the order of the input documents in the combined document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentPdfMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentPdfMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentPng**
> string MergeDocumentPng(ctx, inputFile1, inputFile2)
Merge Two PNG Files Together

Combine two PNG files into a single PNG document, preserving the order of the input documents in the combined document by stacking them vertically

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentPngMulti**
> string MergeDocumentPngMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple PNG Files Together

Combine multiple PNG files into a single PNG document, preserving the order of the input documents in the combined document by stacking them vertically

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentPngMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentPngMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentPptx**
> string MergeDocumentPptx(ctx, inputFile1, inputFile2)
Merge Two PowerPoint PPTX Together

Combine two Office PowerPoint presentations (pptx) into one single Office PowerPoint presentation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentPptxMulti**
> string MergeDocumentPptxMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple PowerPoint PPTX Together

Combine multiple Office PowerPoint presentations (pptx) into one single Office PowerPoint presentation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentPptxMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentPptxMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentTxt**
> interface{} MergeDocumentTxt(ctx, inputFile1, inputFile2)
Merge Two Text (TXT) Files Together

Combine two Text (.TXT) files into a single text document, preserving the order of the input documents in the combined document by stacking them vertically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentTxtMulti**
> string MergeDocumentTxtMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple Text (TXT) Files Together

Combine multiple Text (.TXT) files into a single text document, preserving the order of the input documents in the combined document by stacking them vertically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentTxtMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentTxtMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentXlsx**
> string MergeDocumentXlsx(ctx, inputFile1, inputFile2)
Merge Two Excel XLSX Together

Combine two Office Excel spreadsheets (xlsx) into a single Office Excel spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on (more than 2 can be supplied). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeDocumentXlsxMulti**
> string MergeDocumentXlsxMulti(ctx, inputFile1, inputFile2, optional)
Merge Multple Excel XLSX Together

Combine multiple Office Excel spreadsheets (xlsx) into a single Office Excel spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
  **inputFile2** | ***os.File**| Second input file to perform the operation on. | 
 **optional** | ***MergeDocumentXlsxMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeDocumentXlsxMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile3** | **optional.Interface of *os.File**| Third input file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

