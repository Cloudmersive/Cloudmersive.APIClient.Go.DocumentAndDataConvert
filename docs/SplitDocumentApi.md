# \SplitDocumentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SplitDocumentDocx**](SplitDocumentApi.md#SplitDocumentDocx) | **Post** /convert/split/docx | Split a single Word Document DOCX into Separate Documents by Page
[**SplitDocumentPdfByPage**](SplitDocumentApi.md#SplitDocumentPdfByPage) | **Post** /convert/split/pdf | Split a PDF file into separate PDF files, one per page
[**SplitDocumentPptx**](SplitDocumentApi.md#SplitDocumentPptx) | **Post** /convert/split/pptx | Split a single PowerPoint Presentation PPTX into Separate Slides
[**SplitDocumentTxtByLine**](SplitDocumentApi.md#SplitDocumentTxtByLine) | **Post** /convert/split/txt/by-line | Split a single Text file (txt) into lines
[**SplitDocumentTxtByString**](SplitDocumentApi.md#SplitDocumentTxtByString) | **Post** /convert/split/txt/by-string | Split a single Text file (txt) by a string delimiter
[**SplitDocumentXlsx**](SplitDocumentApi.md#SplitDocumentXlsx) | **Post** /convert/split/xlsx | Split a single Excel XLSX into Separate Worksheets


# **SplitDocumentDocx**
> SplitDocxDocumentResult SplitDocumentDocx(ctx, inputFile, optional)
Split a single Word Document DOCX into Separate Documents by Page

Split a Word DOCX Document, comprised of multiple pages into separate Word DOCX document files, with each containing exactly one page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***SplitDocumentDocxOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitDocumentDocxOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnDocumentContents** | **optional.Bool**| Set to true to return the contents of each Worksheet directly, set to false to only return URLs to each resulting document.  Default is true. | 

### Return type

[**SplitDocxDocumentResult**](SplitDocxDocumentResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitDocumentPdfByPage**
> SplitPdfResult SplitDocumentPdfByPage(ctx, inputFile, optional)
Split a PDF file into separate PDF files, one per page

Split an input PDF file into separate pages, comprised of one PDF file per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***SplitDocumentPdfByPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitDocumentPdfByPageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnDocumentContents** | **optional.Bool**| Set to true to directly return all of the document contents in the DocumentContents field; set to false to return contents as temporary URLs (more efficient for large operations).  Default is false. | 

### Return type

[**SplitPdfResult**](SplitPdfResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitDocumentPptx**
> SplitPptxPresentationResult SplitDocumentPptx(ctx, inputFile, optional)
Split a single PowerPoint Presentation PPTX into Separate Slides

Split an PowerPoint PPTX Presentation, comprised of multiple slides into separate PowerPoint PPTX presentation files, with each containing exactly one slide.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***SplitDocumentPptxOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitDocumentPptxOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnDocumentContents** | **optional.Bool**| Set to true to return the contents of each presentation directly, set to false to only return URLs to each resulting presentation.  Default is true. | 

### Return type

[**SplitPptxPresentationResult**](SplitPptxPresentationResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitDocumentTxtByLine**
> SplitTextDocumentByLinesResult SplitDocumentTxtByLine(ctx, inputFile)
Split a single Text file (txt) into lines

Split a Text (txt) Document by line, returning each line separately in order.  Supports multiple types of newlines.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**SplitTextDocumentByLinesResult**](SplitTextDocumentByLinesResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitDocumentTxtByString**
> SplitTextDocumentByStringResult SplitDocumentTxtByString(ctx, inputFile, splitDelimiter, optional)
Split a single Text file (txt) by a string delimiter

Split a Text (txt) Document by a string delimiter, returning each component of the string as an array of strings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **splitDelimiter** | **string**| Required; String to split up the input file on | 
 **optional** | ***SplitDocumentTxtByStringOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitDocumentTxtByStringOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipEmptyElements** | **optional.Bool**| Optional; If true, empty elements will be skipped in the output | 

### Return type

[**SplitTextDocumentByStringResult**](SplitTextDocumentByStringResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitDocumentXlsx**
> SplitXlsxWorksheetResult SplitDocumentXlsx(ctx, inputFile, optional)
Split a single Excel XLSX into Separate Worksheets

Split an Excel XLSX Spreadsheet, comprised of multiple Worksheets (or Tabs) into separate Excel XLSX spreadsheet files, with each containing exactly one Worksheet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***SplitDocumentXlsxOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitDocumentXlsxOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnDocumentContents** | **optional.Bool**| Set to true to return the contents of each Worksheet directly, set to false to only return URLs to each resulting worksheet.  Default is true. | 

### Return type

[**SplitXlsxWorksheetResult**](SplitXlsxWorksheetResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

