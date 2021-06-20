# \EditPdfApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditPdfAddAnnotations**](EditPdfApi.md#EditPdfAddAnnotations) | **Post** /convert/edit/pdf/annotations/add-item | Add one or more PDF annotations, comments in the PDF document
[**EditPdfConvertToPdfA**](EditPdfApi.md#EditPdfConvertToPdfA) | **Post** /convert/edit/pdf/optimize/pdf-a | Convert a PDF file to PDF/A
[**EditPdfDecrypt**](EditPdfApi.md#EditPdfDecrypt) | **Post** /convert/edit/pdf/decrypt | Decrypt and password-protect a PDF
[**EditPdfDeletePages**](EditPdfApi.md#EditPdfDeletePages) | **Post** /convert/edit/pdf/pages/delete | Remove, delete pages from a PDF document
[**EditPdfEncrypt**](EditPdfApi.md#EditPdfEncrypt) | **Post** /convert/edit/pdf/encrypt | Encrypt and password-protect a PDF
[**EditPdfGetAnnotations**](EditPdfApi.md#EditPdfGetAnnotations) | **Post** /convert/edit/pdf/annotations/list | Get PDF annotations, including comments in the document
[**EditPdfGetFormFields**](EditPdfApi.md#EditPdfGetFormFields) | **Post** /convert/edit/pdf/form/get-fields | Gets PDF Form fields and values
[**EditPdfGetMetadata**](EditPdfApi.md#EditPdfGetMetadata) | **Post** /convert/edit/pdf/get-metadata | Get PDF document metadata
[**EditPdfGetPdfTextByPages**](EditPdfApi.md#EditPdfGetPdfTextByPages) | **Post** /convert/edit/pdf/pages/get-text | Get text in a PDF document by page
[**EditPdfInsertPages**](EditPdfApi.md#EditPdfInsertPages) | **Post** /convert/edit/pdf/pages/insert | Insert, copy pages from one PDF document into another
[**EditPdfLinearize**](EditPdfApi.md#EditPdfLinearize) | **Post** /convert/edit/pdf/optimize/linearize | Linearize and optimize a PDF for streaming download
[**EditPdfRasterize**](EditPdfApi.md#EditPdfRasterize) | **Post** /convert/edit/pdf/rasterize | Rasterize a PDF to an image-based PDF
[**EditPdfReduceFileSize**](EditPdfApi.md#EditPdfReduceFileSize) | **Post** /convert/edit/pdf/optimize/reduce-file-size | Reduce the file size and optimize a PDF
[**EditPdfRemoveAllAnnotations**](EditPdfApi.md#EditPdfRemoveAllAnnotations) | **Post** /convert/edit/pdf/annotations/remove-all | Remove all PDF annotations, including comments in the document
[**EditPdfRemoveAnnotationItem**](EditPdfApi.md#EditPdfRemoveAnnotationItem) | **Post** /convert/edit/pdf/annotations/remove-item | Remove a specific PDF annotation, comment in the document
[**EditPdfResize**](EditPdfApi.md#EditPdfResize) | **Post** /convert/edit/pdf/resize | Change PDF Document&#39;s Paper Size
[**EditPdfRotateAllPages**](EditPdfApi.md#EditPdfRotateAllPages) | **Post** /convert/edit/pdf/pages/rotate/all | Rotate all pages in a PDF document
[**EditPdfRotatePageRange**](EditPdfApi.md#EditPdfRotatePageRange) | **Post** /convert/edit/pdf/pages/rotate/page-range | Rotate a range, subset of pages in a PDF document
[**EditPdfSetFormFields**](EditPdfApi.md#EditPdfSetFormFields) | **Post** /convert/edit/pdf/form/set-fields | Sets ands fills PDF Form field values
[**EditPdfSetMetadata**](EditPdfApi.md#EditPdfSetMetadata) | **Post** /convert/edit/pdf/set-metadata | Sets PDF document metadata
[**EditPdfSetPermissions**](EditPdfApi.md#EditPdfSetPermissions) | **Post** /convert/edit/pdf/encrypt/set-permissions | Encrypt, password-protect and set restricted permissions on a PDF
[**EditPdfWatermarkText**](EditPdfApi.md#EditPdfWatermarkText) | **Post** /convert/edit/pdf/watermark/text | Add a text watermark to a PDF


# **EditPdfAddAnnotations**
> string EditPdfAddAnnotations(ctx, request)
Add one or more PDF annotations, comments in the PDF document

Adds one or more annotations, comments to a PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**AddPdfAnnotationRequest**](AddPdfAnnotationRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfConvertToPdfA**
> string EditPdfConvertToPdfA(ctx, inputFile, optional)
Convert a PDF file to PDF/A

Converts the input PDF file to a PDF/A-1b or PDF/A-2b standardized PDF.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***EditPdfConvertToPdfAOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditPdfConvertToPdfAOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conformanceLevel** | **optional.String**| Optional: Select the conformance level for PDF/A - specify &#39;1b&#39; for PDF/A-1b or specify &#39;2b&#39; for PDF/A-2b; default is PDF/A-1b | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfDecrypt**
> string EditPdfDecrypt(ctx, password, inputFile)
Decrypt and password-protect a PDF

Decrypt a PDF document with a password.  Decrypted PDF will no longer require a password to open.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **password** | **string**| Valid password for the PDF file | 
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfDeletePages**
> string EditPdfDeletePages(ctx, inputFile, pageStart, pageEnd)
Remove, delete pages from a PDF document

Remove one or more pages from a PDF document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **pageStart** | **int32**| Page number (1 based) to start deleting pages from (inclusive). | 
  **pageEnd** | **int32**| Page number (1 based) to stop deleting pages from (inclusive). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfEncrypt**
> string EditPdfEncrypt(ctx, inputFile, optional)
Encrypt and password-protect a PDF

Encrypt a PDF document with a password.  Set an owner password to control owner (editor/creator) permissions, and set a user (reader) password to control the viewer of the PDF.  Set the password fields null to omit the given password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***EditPdfEncryptOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditPdfEncryptOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPassword** | **optional.String**| Password of a user (reader) of the PDF file | 
 **ownerPassword** | **optional.String**| Password of a owner (creator/editor) of the PDF file | 
 **encryptionKeyLength** | **optional.String**| Possible values are \&quot;128\&quot; (128-bit RC4 encryption) and \&quot;256\&quot; (256-bit AES encryption).  Default is 256. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfGetAnnotations**
> GetPdfAnnotationsResult EditPdfGetAnnotations(ctx, inputFile)
Get PDF annotations, including comments in the document

Enumerates the annotations, including comments and notes, in a PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetPdfAnnotationsResult**](GetPdfAnnotationsResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfGetFormFields**
> PdfFormFields EditPdfGetFormFields(ctx, inputFile)
Gets PDF Form fields and values

Encrypt a PDF document with a password.  Set an owner password to control owner (editor/creator) permissions, and set a user (reader) password to control the viewer of the PDF.  Set the password fields null to omit the given password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**PdfFormFields**](PdfFormFields.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfGetMetadata**
> PdfMetadata EditPdfGetMetadata(ctx, inputFile)
Get PDF document metadata

Returns the metadata from the PDF document, including Title, Author, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**PdfMetadata**](PdfMetadata.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfGetPdfTextByPages**
> PdfTextByPageResult EditPdfGetPdfTextByPages(ctx, inputFile, optional)
Get text in a PDF document by page

Gets the text in a PDF by page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***EditPdfGetPdfTextByPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditPdfGetPdfTextByPagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **textFormattingMode** | **optional.String**| Optional; specify how whitespace should be handled when converting the document to text.  Possible values are &#39;preserveWhitespace&#39; which will attempt to preserve whitespace in the document and relative positioning of text within the document, and &#39;minimizeWhitespace&#39; which will not insert additional spaces into the document in most cases.  Default is &#39;preserveWhitespace&#39;. | 

### Return type

[**PdfTextByPageResult**](PdfTextByPageResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfInsertPages**
> string EditPdfInsertPages(ctx, sourceFile, destinationFile, pageStartSource, pageEndSource, pageInsertBeforeDesitnation)
Insert, copy pages from one PDF document into another

Copy one or more pages from one PDF document (source document) and insert them into a second PDF document (destination document).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceFile** | ***os.File**| Source PDF file to copy pages from. | 
  **destinationFile** | ***os.File**| Destination PDF file to copy pages into. | 
  **pageStartSource** | **int32**| Page number (1 based) to start copying pages from (inclusive) in the Source file. | 
  **pageEndSource** | **int32**| Page number (1 based) to stop copying pages pages from (inclusive) in the Source file. | 
  **pageInsertBeforeDesitnation** | **int32**| Page number (1 based) to insert the pages before in the Destination file. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfLinearize**
> string EditPdfLinearize(ctx, inputFile)
Linearize and optimize a PDF for streaming download

Linearizes the content of a PDF to optimize it for streaming download, particularly over web streaming.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfRasterize**
> string EditPdfRasterize(ctx, inputFile)
Rasterize a PDF to an image-based PDF

Rasterize a PDF into an image-based PDF.  The output is a PDF where each page is comprised of a high-resolution image, with all text, figures and other components removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfReduceFileSize**
> string EditPdfReduceFileSize(ctx, inputFile, optional)
Reduce the file size and optimize a PDF

Reduces the file size and optimizes the content of a PDF to minimize its file size.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***EditPdfReduceFileSizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditPdfReduceFileSizeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quality** | **optional.Float32**| Quality level for the images in the PDF, ranging from 0.0 (low quality) to 1.0 (high quality); default is 0.3 | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfRemoveAllAnnotations**
> string EditPdfRemoveAllAnnotations(ctx, inputFile)
Remove all PDF annotations, including comments in the document

Removes all of the annotations, including comments and notes, in a PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfRemoveAnnotationItem**
> string EditPdfRemoveAnnotationItem(ctx, inputFile, annotationIndex)
Remove a specific PDF annotation, comment in the document

Removes a specific annotation in a PDF document, using the AnnotationIndex.  To enumerate AnnotationIndex for all of the annotations in the PDF document, use the /edit/pdf/annotations/list API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **annotationIndex** | **int32**| The 0-based index of the annotation in the document | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfResize**
> string EditPdfResize(ctx, inputFile, paperSize)
Change PDF Document's Paper Size

Resizes a PDF document's paper size.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **paperSize** | **string**| The desired paper size for the resized PDF document. Size ranges from A7 (smallest) to A0 (largest). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfRotateAllPages**
> string EditPdfRotateAllPages(ctx, inputFile, rotationAngle)
Rotate all pages in a PDF document

Rotate all of the pages in a PDF document by a multiple of 90 degrees

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **rotationAngle** | **int32**| The angle to rotate the page in degrees, must be a multiple of 90 degrees, e.g. 90, 180, 270, or -90, -180, -270, etc. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfRotatePageRange**
> string EditPdfRotatePageRange(ctx, inputFile, rotationAngle, pageStart, pageEnd)
Rotate a range, subset of pages in a PDF document

Rotate a range of specific pages in a PDF document by a multiple of 90 degrees

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
  **rotationAngle** | **int32**| The angle to rotate the page in degrees, must be a multiple of 90 degrees, e.g. 90, 180, 270, or -90, -180, -270, etc. | 
  **pageStart** | **int32**| Page number (1 based) to start rotating pages from (inclusive). | 
  **pageEnd** | **int32**| Page number (1 based) to stop rotating pages from (inclusive). | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfSetFormFields**
> string EditPdfSetFormFields(ctx, fieldValues)
Sets ands fills PDF Form field values

Fill in the form fields in a PDF form with specific values.  Use form/get-fields to enumerate the available fields and their data types in an input form.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fieldValues** | [**SetPdfFormFieldsRequest**](SetPdfFormFieldsRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfSetMetadata**
> string EditPdfSetMetadata(ctx, request)
Sets PDF document metadata

Sets (writes) metadata into the input PDF document, including Title, Author, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**SetPdfMetadataRequest**](SetPdfMetadataRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfSetPermissions**
> string EditPdfSetPermissions(ctx, ownerPassword, userPassword, inputFile, optional)
Encrypt, password-protect and set restricted permissions on a PDF

Encrypt a PDF document with a password, and set permissions on the PDF.  Set an owner password to control owner (editor/creator) permissions [required], and set a user (reader) password to control the viewer of the PDF [optional].  Set the reader password to null to omit the password.  Restrict or allow printing, copying content, document assembly, editing (read-only), form filling, modification of annotations, and degraded printing through document Digital Rights Management (DRM).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerPassword** | **string**| Password of a owner (creator/editor) of the PDF file (required) | 
  **userPassword** | **string**| Password of a user (reader) of the PDF file (optional) | 
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***EditPdfSetPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditPdfSetPermissionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **encryptionKeyLength** | **optional.String**| Possible values are \&quot;128\&quot; (128-bit RC4 encryption) and \&quot;256\&quot; (256-bit AES encryption).  Default is 256. | 
 **allowPrinting** | **optional.Bool**| Set to false to disable printing through DRM.  Default is true. | 
 **allowDocumentAssembly** | **optional.Bool**| Set to false to disable document assembly through DRM.  Default is true. | 
 **allowContentExtraction** | **optional.Bool**| Set to false to disable copying/extracting content out of the PDF through DRM.  Default is true. | 
 **allowFormFilling** | **optional.Bool**| Set to false to disable filling out form fields in the PDF through DRM.  Default is true. | 
 **allowEditing** | **optional.Bool**| Set to false to disable editing in the PDF through DRM (making the PDF read-only).  Default is true. | 
 **allowAnnotations** | **optional.Bool**| Set to false to disable annotations and editing of annotations in the PDF through DRM.  Default is true. | 
 **allowDegradedPrinting** | **optional.Bool**| Set to false to disable degraded printing of the PDF through DRM.  Default is true. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPdfWatermarkText**
> string EditPdfWatermarkText(ctx, watermarkText, inputFile, optional)
Add a text watermark to a PDF

Adds a text watermark to a PDF

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watermarkText** | **string**| Watermark text to add to the PDF (required) | 
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***EditPdfWatermarkTextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditPdfWatermarkTextOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fontName** | **optional.String**| Font Family Name for the watermark text; default is Times New Roman | 
 **fontSize** | **optional.Float32**| Font Size in points of the text; default is 150 | 
 **fontColor** | **optional.String**| Font color in hexadecimal or HTML color name; default is Red | 
 **fontTransparency** | **optional.Float32**| Font transparency between 0.0 (completely transparent) to 1.0 (fully opaque); default is 0.5 | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

