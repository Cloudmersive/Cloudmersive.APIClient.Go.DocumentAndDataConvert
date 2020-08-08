# \CompareDocumentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareDocumentDocx**](CompareDocumentApi.md#CompareDocumentDocx) | **Post** /convert/compare/docx | Compare Two Word DOCX


# **CompareDocumentDocx**
> string CompareDocumentDocx(ctx, inputFile1, inputFile2)
Compare Two Word DOCX

Compare two Office Word Documents (docx) files and highlight the differences

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

