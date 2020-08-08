# \ViewerToolsApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewerToolsCreateSimple**](ViewerToolsApi.md#ViewerToolsCreateSimple) | **Post** /convert/viewer/create/web/simple | Create a web-based viewer


# **ViewerToolsCreateSimple**
> ViewerResponse ViewerToolsCreateSimple(ctx, inputFile)
Create a web-based viewer

Creates an HTML embed code for a simple web-based viewer of a document; supports Office document types and PDF.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**ViewerResponse**](ViewerResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

