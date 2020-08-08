# \ZipArchiveApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZipArchiveZipCreate**](ZipArchiveApi.md#ZipArchiveZipCreate) | **Post** /convert/archive/zip/create | Compress files to create a new zip archive
[**ZipArchiveZipCreateAdvanced**](ZipArchiveApi.md#ZipArchiveZipCreateAdvanced) | **Post** /convert/archive/zip/create/advanced | Compress files and folders to create a new zip archive with advanced options
[**ZipArchiveZipDecrypt**](ZipArchiveApi.md#ZipArchiveZipDecrypt) | **Post** /convert/archive/zip/decrypt | Decrypt and remove password protection on a zip file
[**ZipArchiveZipEncryptAdvanced**](ZipArchiveApi.md#ZipArchiveZipEncryptAdvanced) | **Post** /convert/archive/zip/encrypt/advanced | Encrypt and password protect a zip file
[**ZipArchiveZipExtract**](ZipArchiveApi.md#ZipArchiveZipExtract) | **Post** /convert/archive/zip/extract | Extract, decompress files and folders from a zip archive


# **ZipArchiveZipCreate**
> string ZipArchiveZipCreate(ctx, inputFile1, optional)
Compress files to create a new zip archive

Create a new zip archive by compressing input files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input file to perform the operation on. | 
 **optional** | ***ZipArchiveZipCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ZipArchiveZipCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputFile2** | **optional.Interface of *os.File**| Second input file to perform the operation on. | 
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

# **ZipArchiveZipCreateAdvanced**
> interface{} ZipArchiveZipCreateAdvanced(ctx, request)
Compress files and folders to create a new zip archive with advanced options

Create a new zip archive by compressing input files, folders and leverage advanced options to control the structure of the resulting zip archive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**CreateZipArchiveRequest**](CreateZipArchiveRequest.md)| Input request | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZipArchiveZipDecrypt**
> interface{} ZipArchiveZipDecrypt(ctx, inputFile, zipPassword)
Decrypt and remove password protection on a zip file

Decrypts and removes password protection from an encrypted zip file with the specified password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **zipPassword** | **string**| Required; Password for the input archive | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZipArchiveZipEncryptAdvanced**
> interface{} ZipArchiveZipEncryptAdvanced(ctx, encryptionRequest)
Encrypt and password protect a zip file

Encrypts and password protects an existing zip file with the specified password and encryption algorithm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **encryptionRequest** | [**ZipEncryptionAdvancedRequest**](ZipEncryptionAdvancedRequest.md)| Encryption request | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZipArchiveZipExtract**
> ZipExtractResponse ZipArchiveZipExtract(ctx, inputFile)
Extract, decompress files and folders from a zip archive

Extracts a zip archive by decompressing files, and folders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**ZipExtractResponse**](ZipExtractResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

