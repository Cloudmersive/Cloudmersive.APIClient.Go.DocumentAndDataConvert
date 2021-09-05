# \EditDocumentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditDocumentBeginEditing**](EditDocumentApi.md#EditDocumentBeginEditing) | **Post** /convert/edit/begin-editing | Begin editing a document
[**EditDocumentDocxAcceptAllTrackChanges**](EditDocumentApi.md#EditDocumentDocxAcceptAllTrackChanges) | **Post** /convert/edit/docx/track-changes/accept-all | Accept all tracked changes, revisions in a Word DOCX document
[**EditDocumentDocxBody**](EditDocumentApi.md#EditDocumentDocxBody) | **Post** /convert/edit/docx/get-body | Get body from a Word DOCX document
[**EditDocumentDocxCreateBlankDocument**](EditDocumentApi.md#EditDocumentDocxCreateBlankDocument) | **Post** /convert/edit/docx/create/blank | Create a blank Word DOCX document
[**EditDocumentDocxDeletePages**](EditDocumentApi.md#EditDocumentDocxDeletePages) | **Post** /convert/edit/docx/delete-pages | Delete, remove pages from a Word DOCX document
[**EditDocumentDocxDeleteTableRow**](EditDocumentApi.md#EditDocumentDocxDeleteTableRow) | **Post** /convert/edit/docx/delete-table-row | Deletes a table row in an existing table in a Word DOCX document
[**EditDocumentDocxDeleteTableRowRange**](EditDocumentApi.md#EditDocumentDocxDeleteTableRowRange) | **Post** /convert/edit/docx/delete-table-row/range | Deletes a range of multiple table rows in an existing table in a Word DOCX document
[**EditDocumentDocxDisableTrackChanges**](EditDocumentApi.md#EditDocumentDocxDisableTrackChanges) | **Post** /convert/edit/docx/track-changes/disable | Disable track changes, revisions in a Word DOCX document
[**EditDocumentDocxEnableTrackChanges**](EditDocumentApi.md#EditDocumentDocxEnableTrackChanges) | **Post** /convert/edit/docx/track-changes/enable | Enable track changes, revisions in a Word DOCX document
[**EditDocumentDocxFindParagraph**](EditDocumentApi.md#EditDocumentDocxFindParagraph) | **Post** /convert/edit/docx/find/paragraph | Find matching paragraphs in a Word DOCX document
[**EditDocumentDocxGetComments**](EditDocumentApi.md#EditDocumentDocxGetComments) | **Post** /convert/edit/docx/get-comments/flat-list | Get comments from a Word DOCX document as a flat list
[**EditDocumentDocxGetCommentsHierarchical**](EditDocumentApi.md#EditDocumentDocxGetCommentsHierarchical) | **Post** /convert/edit/docx/get-comments/hierarchical | Get comments from a Word DOCX document hierarchically
[**EditDocumentDocxGetContentControls**](EditDocumentApi.md#EditDocumentDocxGetContentControls) | **Post** /convert/edit/docx/get-content-controls | Get all content controls (form fields) and values in a Word DOCX document
[**EditDocumentDocxGetFormFields**](EditDocumentApi.md#EditDocumentDocxGetFormFields) | **Post** /convert/edit/docx/get-form-fields | Get all form fields in a Word DOCX document
[**EditDocumentDocxGetHeadersAndFooters**](EditDocumentApi.md#EditDocumentDocxGetHeadersAndFooters) | **Post** /convert/edit/docx/get-headers-and-footers | Get content of a footer from a Word DOCX document
[**EditDocumentDocxGetImages**](EditDocumentApi.md#EditDocumentDocxGetImages) | **Post** /convert/edit/docx/get-images | Get images from a Word DOCX document
[**EditDocumentDocxGetMacroInformation**](EditDocumentApi.md#EditDocumentDocxGetMacroInformation) | **Post** /convert/edit/docx/get-macros | Get macro information from a Word DOCX/DOCM document
[**EditDocumentDocxGetMetadataProperties**](EditDocumentApi.md#EditDocumentDocxGetMetadataProperties) | **Post** /convert/edit/docx/get-metadata | Get all metadata properties in Word DOCX document
[**EditDocumentDocxGetSections**](EditDocumentApi.md#EditDocumentDocxGetSections) | **Post** /convert/edit/docx/get-sections | Get sections from a Word DOCX document
[**EditDocumentDocxGetStyles**](EditDocumentApi.md#EditDocumentDocxGetStyles) | **Post** /convert/edit/docx/get-styles | Get styles from a Word DOCX document
[**EditDocumentDocxGetTableByIndex**](EditDocumentApi.md#EditDocumentDocxGetTableByIndex) | **Post** /convert/edit/docx/get-table/by-index | Get a specific table by index in a Word DOCX document
[**EditDocumentDocxGetTableRow**](EditDocumentApi.md#EditDocumentDocxGetTableRow) | **Post** /convert/edit/docx/get-table-row | Gets the contents of an existing table row in an existing table in a Word DOCX document
[**EditDocumentDocxGetTables**](EditDocumentApi.md#EditDocumentDocxGetTables) | **Post** /convert/edit/docx/get-tables | Get all tables in Word DOCX document
[**EditDocumentDocxInsertCommentOnParagraph**](EditDocumentApi.md#EditDocumentDocxInsertCommentOnParagraph) | **Post** /convert/edit/docx/insert-comment/on/paragraph | Insert a new comment into a Word DOCX document attached to a paragraph
[**EditDocumentDocxInsertImage**](EditDocumentApi.md#EditDocumentDocxInsertImage) | **Post** /convert/edit/docx/insert-image | Insert image into a Word DOCX document
[**EditDocumentDocxInsertParagraph**](EditDocumentApi.md#EditDocumentDocxInsertParagraph) | **Post** /convert/edit/docx/insert-paragraph | Insert a new paragraph into a Word DOCX document
[**EditDocumentDocxInsertTable**](EditDocumentApi.md#EditDocumentDocxInsertTable) | **Post** /convert/edit/docx/insert-table | Insert a new table into a Word DOCX document
[**EditDocumentDocxInsertTableRow**](EditDocumentApi.md#EditDocumentDocxInsertTableRow) | **Post** /convert/edit/docx/insert-table-row | Insert a new row into an existing table in a Word DOCX document
[**EditDocumentDocxPages**](EditDocumentApi.md#EditDocumentDocxPages) | **Post** /convert/edit/docx/get-pages | Get pages and content from a Word DOCX document
[**EditDocumentDocxRemoveAllComments**](EditDocumentApi.md#EditDocumentDocxRemoveAllComments) | **Post** /convert/edit/docx/comments/remove-all | Remove all comments from a Word DOCX document
[**EditDocumentDocxRemoveHeadersAndFooters**](EditDocumentApi.md#EditDocumentDocxRemoveHeadersAndFooters) | **Post** /convert/edit/docx/remove-headers-and-footers | Remove headers and footers from Word DOCX document
[**EditDocumentDocxRemoveObject**](EditDocumentApi.md#EditDocumentDocxRemoveObject) | **Post** /convert/edit/docx/remove-object | Delete any object in a Word DOCX document
[**EditDocumentDocxReplace**](EditDocumentApi.md#EditDocumentDocxReplace) | **Post** /convert/edit/docx/replace-all | Replace string in Word DOCX document
[**EditDocumentDocxReplaceMulti**](EditDocumentApi.md#EditDocumentDocxReplaceMulti) | **Post** /convert/edit/docx/replace-all/multi | Replace multiple strings in Word DOCX document, return result
[**EditDocumentDocxReplaceMultiEditSession**](EditDocumentApi.md#EditDocumentDocxReplaceMultiEditSession) | **Post** /convert/edit/docx/replace-all/multi/edit-session | Replace multiple strings in Word DOCX document, return edit session
[**EditDocumentDocxReplaceParagraph**](EditDocumentApi.md#EditDocumentDocxReplaceParagraph) | **Post** /convert/edit/docx/replace/paragraph | Replace matching paragraphs in a Word DOCX document
[**EditDocumentDocxSetCustomMetadataProperties**](EditDocumentApi.md#EditDocumentDocxSetCustomMetadataProperties) | **Post** /convert/edit/docx/set-metadata/custom-property | Set custom property metadata properties in Word DOCX document
[**EditDocumentDocxSetFooter**](EditDocumentApi.md#EditDocumentDocxSetFooter) | **Post** /convert/edit/docx/set-footer | Set the footer in a Word DOCX document
[**EditDocumentDocxSetFooterAddPageNumber**](EditDocumentApi.md#EditDocumentDocxSetFooterAddPageNumber) | **Post** /convert/edit/docx/set-footer/add-page-number | Add page number to footer in a Word DOCX document
[**EditDocumentDocxSetFormFields**](EditDocumentApi.md#EditDocumentDocxSetFormFields) | **Post** /convert/edit/docx/set-form-fields | Set and fill values for form fields in a Word DOCX document
[**EditDocumentDocxSetHeader**](EditDocumentApi.md#EditDocumentDocxSetHeader) | **Post** /convert/edit/docx/set-header | Set the header in a Word DOCX document
[**EditDocumentDocxUpdateTableCell**](EditDocumentApi.md#EditDocumentDocxUpdateTableCell) | **Post** /convert/edit/docx/update-table-cell | Update, set contents of a table cell in an existing table in a Word DOCX document
[**EditDocumentDocxUpdateTableRow**](EditDocumentApi.md#EditDocumentDocxUpdateTableRow) | **Post** /convert/edit/docx/update-table-row | Update, set contents of a table row in an existing table in a Word DOCX document
[**EditDocumentFinishEditing**](EditDocumentApi.md#EditDocumentFinishEditing) | **Post** /convert/edit/finish-editing | Finish editing document, and download result from document editing
[**EditDocumentPptxDeleteSlides**](EditDocumentApi.md#EditDocumentPptxDeleteSlides) | **Post** /convert/edit/pptx/delete-slides | Delete, remove slides from a PowerPoint PPTX presentation document
[**EditDocumentPptxGetMacroInformation**](EditDocumentApi.md#EditDocumentPptxGetMacroInformation) | **Post** /convert/edit/pptx/get-macros | Get macro information from a PowerPoint PPTX/PPTM presentation document
[**EditDocumentPptxReplace**](EditDocumentApi.md#EditDocumentPptxReplace) | **Post** /convert/edit/pptx/replace-all | Replace string in PowerPoint PPTX presentation
[**EditDocumentXlsxAppendRow**](EditDocumentApi.md#EditDocumentXlsxAppendRow) | **Post** /convert/edit/xlsx/append-row | Append row to a Excel XLSX spreadsheet, worksheet
[**EditDocumentXlsxClearCellByIndex**](EditDocumentApi.md#EditDocumentXlsxClearCellByIndex) | **Post** /convert/edit/xlsx/clear-cell/by-index | Clear cell contents in an Excel XLSX spreadsheet, worksheet by index
[**EditDocumentXlsxClearRow**](EditDocumentApi.md#EditDocumentXlsxClearRow) | **Post** /convert/edit/xlsx/clear-row | Clear row from a Excel XLSX spreadsheet, worksheet
[**EditDocumentXlsxCreateBlankSpreadsheet**](EditDocumentApi.md#EditDocumentXlsxCreateBlankSpreadsheet) | **Post** /convert/edit/xlsx/create/blank | Create a blank Excel XLSX spreadsheet
[**EditDocumentXlsxCreateSpreadsheetFromData**](EditDocumentApi.md#EditDocumentXlsxCreateSpreadsheetFromData) | **Post** /convert/edit/xlsx/create/from/data | Create a new Excel XLSX spreadsheet from column and row data
[**EditDocumentXlsxDeleteWorksheet**](EditDocumentApi.md#EditDocumentXlsxDeleteWorksheet) | **Post** /convert/edit/xlsx/delete-worksheet | Delete, remove worksheet from an Excel XLSX spreadsheet document
[**EditDocumentXlsxDisableSharedWorkbook**](EditDocumentApi.md#EditDocumentXlsxDisableSharedWorkbook) | **Post** /convert/edit/xlsx/configuration/disable-shared-workbook | Disable Shared Workbook (legacy) in Excel XLSX spreadsheet
[**EditDocumentXlsxEnableSharedWorkbook**](EditDocumentApi.md#EditDocumentXlsxEnableSharedWorkbook) | **Post** /convert/edit/xlsx/configuration/enable-shared-workbook | Enable Shared Workbook (legacy) in Excel XLSX spreadsheet
[**EditDocumentXlsxGetCellByIdentifier**](EditDocumentApi.md#EditDocumentXlsxGetCellByIdentifier) | **Post** /convert/edit/xlsx/get-cell/by-identifier | Get cell from an Excel XLSX spreadsheet, worksheet by cell identifier
[**EditDocumentXlsxGetCellByIndex**](EditDocumentApi.md#EditDocumentXlsxGetCellByIndex) | **Post** /convert/edit/xlsx/get-cell/by-index | Get cell from an Excel XLSX spreadsheet, worksheet by index
[**EditDocumentXlsxGetColumns**](EditDocumentApi.md#EditDocumentXlsxGetColumns) | **Post** /convert/edit/xlsx/get-columns | Get columns from a Excel XLSX spreadsheet, worksheet
[**EditDocumentXlsxGetImages**](EditDocumentApi.md#EditDocumentXlsxGetImages) | **Post** /convert/edit/xlsx/get-images | Get images from a Excel XLSX spreadsheet, worksheet
[**EditDocumentXlsxGetMacroInformation**](EditDocumentApi.md#EditDocumentXlsxGetMacroInformation) | **Post** /convert/edit/xlsx/get-macros | Get macro information from a Excel XLSX/XLSM spreadsheet, worksheet
[**EditDocumentXlsxGetRowsAndCells**](EditDocumentApi.md#EditDocumentXlsxGetRowsAndCells) | **Post** /convert/edit/xlsx/get-rows-and-cells | Get rows and cells from a Excel XLSX spreadsheet, worksheet
[**EditDocumentXlsxGetSpecificRow**](EditDocumentApi.md#EditDocumentXlsxGetSpecificRow) | **Post** /convert/edit/xlsx/get-specific-row | Get a specific row from a Excel XLSX spreadsheet, worksheet by path
[**EditDocumentXlsxGetStyles**](EditDocumentApi.md#EditDocumentXlsxGetStyles) | **Post** /convert/edit/xlsx/get-styles | Get styles from a Excel XLSX spreadsheet, worksheet
[**EditDocumentXlsxGetWorksheets**](EditDocumentApi.md#EditDocumentXlsxGetWorksheets) | **Post** /convert/edit/xlsx/get-worksheets | Get worksheets from a Excel XLSX spreadsheet
[**EditDocumentXlsxInsertWorksheet**](EditDocumentApi.md#EditDocumentXlsxInsertWorksheet) | **Post** /convert/edit/xlsx/insert-worksheet | Insert a new worksheet into an Excel XLSX spreadsheet
[**EditDocumentXlsxRenameWorksheet**](EditDocumentApi.md#EditDocumentXlsxRenameWorksheet) | **Post** /convert/edit/xlsx/rename-worksheet | Rename a specific worksheet in a Excel XLSX spreadsheet
[**EditDocumentXlsxSetCellByIdentifier**](EditDocumentApi.md#EditDocumentXlsxSetCellByIdentifier) | **Post** /convert/edit/xlsx/set-cell/by-identifier | Set, update cell contents in an Excel XLSX spreadsheet, worksheet by cell identifier
[**EditDocumentXlsxSetCellByIndex**](EditDocumentApi.md#EditDocumentXlsxSetCellByIndex) | **Post** /convert/edit/xlsx/set-cell/by-index | Set, update cell contents in an Excel XLSX spreadsheet, worksheet by index


# **EditDocumentBeginEditing**
> string EditDocumentBeginEditing(ctx, inputFile)
Begin editing a document

Uploads a document to Cloudmersive to begin a series of one or more editing operations.  To edit a document, first call Begin Editing on the document.  Then perform operations on the document using the secure URL returned from BeginEditing, such as Word DOCX Delete Pages and Insert Table.  Finally, perform finish editing on the URL to return the resulting edited document.  The editing URL is temporary and only stored in-memory cache, and will automatically expire from the cache after 30 minutes, and cannot be directly accessed.

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
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxAcceptAllTrackChanges**
> string EditDocumentDocxAcceptAllTrackChanges(ctx, inputFile)
Accept all tracked changes, revisions in a Word DOCX document

Accepts all tracked changes and revisions in a Word DOCX document.  This will accept all pending changes in the document when tracked changes is turned on.  Track changes will remain on (if it is on) after this oepration is completed.

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
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxBody**
> GetDocxBodyResponse EditDocumentDocxBody(ctx, reqConfig)
Get body from a Word DOCX document

Returns the body defined in the Word Document (DOCX) format file; this is the main content part of a DOCX document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxBodyRequest**](GetDocxBodyRequest.md)| Document input request | 

### Return type

[**GetDocxBodyResponse**](GetDocxBodyResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxCreateBlankDocument**
> CreateBlankDocxResponse EditDocumentDocxCreateBlankDocument(ctx, input)
Create a blank Word DOCX document

Returns a blank Word DOCX Document format file.  The file is blank, with no contents.  Use additional editing commands such as Insert Paragraph or Insert Table or Insert Image to populate the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**CreateBlankDocxRequest**](CreateBlankDocxRequest.md)| Document input request | 

### Return type

[**CreateBlankDocxResponse**](CreateBlankDocxResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxDeletePages**
> string EditDocumentDocxDeletePages(ctx, reqConfig)
Delete, remove pages from a Word DOCX document

Returns the edited Word Document in the Word Document (DOCX) format file with the specified pages removed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**RemoveDocxPagesRequest**](RemoveDocxPagesRequest.md)| Document input request | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxDeleteTableRow**
> DeleteDocxTableRowResponse EditDocumentDocxDeleteTableRow(ctx, reqConfig)
Deletes a table row in an existing table in a Word DOCX document

Deletes an existing table row in a Word DOCX Document and returns the result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DeleteDocxTableRowRequest**](DeleteDocxTableRowRequest.md)| Document input request | 

### Return type

[**DeleteDocxTableRowResponse**](DeleteDocxTableRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxDeleteTableRowRange**
> DeleteDocxTableRowRangeResponse EditDocumentDocxDeleteTableRowRange(ctx, reqConfig)
Deletes a range of multiple table rows in an existing table in a Word DOCX document

Deletes a range of 1 or more existing table rows in a Word DOCX Document and returns the result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DeleteDocxTableRowRangeRequest**](DeleteDocxTableRowRangeRequest.md)| Document input request | 

### Return type

[**DeleteDocxTableRowRangeResponse**](DeleteDocxTableRowRangeResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxDisableTrackChanges**
> string EditDocumentDocxDisableTrackChanges(ctx, inputFile)
Disable track changes, revisions in a Word DOCX document

Diables tracking of changes and revisions in a Word DOCX document, and accepts any pending changes.  Users editing the document will no longer see changes tracked automatically.

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
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxEnableTrackChanges**
> string EditDocumentDocxEnableTrackChanges(ctx, inputFile)
Enable track changes, revisions in a Word DOCX document

Enables tracking of changes and revisions in a Word DOCX document.  Users editing the document will see changes tracked automatically, with edits highlighted, and the ability to accept or reject changes made to the document.

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
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxFindParagraph**
> FindDocxParagraphResponse EditDocumentDocxFindParagraph(ctx, reqConfig)
Find matching paragraphs in a Word DOCX document

Returns the paragraphs defined in the Word Document (DOCX) format file that match the input criteria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**FindDocxParagraphRequest**](FindDocxParagraphRequest.md)| Document input request | 

### Return type

[**FindDocxParagraphResponse**](FindDocxParagraphResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetComments**
> GetDocxCommentsResponse EditDocumentDocxGetComments(ctx, reqConfig)
Get comments from a Word DOCX document as a flat list

Returns the comments and review annotations stored in the Word Document (DOCX) format file as a flattened list (not as a hierarchy of comments and replies).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxGetCommentsRequest**](GetDocxGetCommentsRequest.md)| Document input request | 

### Return type

[**GetDocxCommentsResponse**](GetDocxCommentsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetCommentsHierarchical**
> GetDocxCommentsHierarchicalResponse EditDocumentDocxGetCommentsHierarchical(ctx, reqConfig)
Get comments from a Word DOCX document hierarchically

Returns the comments and review annotations stored in the Word Document (DOCX) format file hierarchically, where reply comments are nested as children under top-level comments in the results returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxGetCommentsHierarchicalRequest**](GetDocxGetCommentsHierarchicalRequest.md)| Document input request | 

### Return type

[**GetDocxCommentsHierarchicalResponse**](GetDocxCommentsHierarchicalResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetContentControls**
> GetDocxContentControlsResponse EditDocumentDocxGetContentControls(ctx, inputFile)
Get all content controls (form fields) and values in a Word DOCX document

Returns all the content controls, used for creating form fields, in a Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetDocxContentControlsResponse**](GetDocxContentControlsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetFormFields**
> GetDocxGetFormFieldsResponse EditDocumentDocxGetFormFields(ctx, inputFile)
Get all form fields in a Word DOCX document

Returns all the content controls, used for creating form fields, as well as handlebar style text-based form fields such as \"{{FieldName}}\", in a Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetDocxGetFormFieldsResponse**](GetDocxGetFormFieldsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetHeadersAndFooters**
> GetDocxHeadersAndFootersResponse EditDocumentDocxGetHeadersAndFooters(ctx, reqConfig)
Get content of a footer from a Word DOCX document

Returns the footer content from a Word Document (DOCX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxHeadersAndFootersRequest**](GetDocxHeadersAndFootersRequest.md)| Document input request | 

### Return type

[**GetDocxHeadersAndFootersResponse**](GetDocxHeadersAndFootersResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetImages**
> GetDocxImagesResponse EditDocumentDocxGetImages(ctx, reqConfig)
Get images from a Word DOCX document

Returns the images defined in the Word Document (DOCX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxImagesRequest**](GetDocxImagesRequest.md)| Document input request | 

### Return type

[**GetDocxImagesResponse**](GetDocxImagesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetMacroInformation**
> GetMacrosResponse EditDocumentDocxGetMacroInformation(ctx, inputFile)
Get macro information from a Word DOCX/DOCM document

Returns information about the Macros (e.g. VBA) defined in the Word Document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetMacrosResponse**](GetMacrosResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetMetadataProperties**
> GetDocxMetadataPropertiesResponse EditDocumentDocxGetMetadataProperties(ctx, inputFile)
Get all metadata properties in Word DOCX document

Returns all the metadata properties in an Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetDocxMetadataPropertiesResponse**](GetDocxMetadataPropertiesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetSections**
> GetDocxSectionsResponse EditDocumentDocxGetSections(ctx, reqConfig)
Get sections from a Word DOCX document

Returns the sections defined in the Word Document (DOCX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxSectionsRequest**](GetDocxSectionsRequest.md)| Document input request | 

### Return type

[**GetDocxSectionsResponse**](GetDocxSectionsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetStyles**
> GetDocxStylesResponse EditDocumentDocxGetStyles(ctx, reqConfig)
Get styles from a Word DOCX document

Returns the styles defined in the Word Document (DOCX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxStylesRequest**](GetDocxStylesRequest.md)| Document input request | 

### Return type

[**GetDocxStylesResponse**](GetDocxStylesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetTableByIndex**
> GetDocxTableByIndexResponse EditDocumentDocxGetTableByIndex(ctx, reqConfig)
Get a specific table by index in a Word DOCX document

Returns the specific table object by its 0-based index in an Office Word Document (DOCX)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxTableByIndexRequest**](GetDocxTableByIndexRequest.md)| Document input request | 

### Return type

[**GetDocxTableByIndexResponse**](GetDocxTableByIndexResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetTableRow**
> GetDocxTableRowResponse EditDocumentDocxGetTableRow(ctx, reqConfig)
Gets the contents of an existing table row in an existing table in a Word DOCX document

Gets the contents of an existing table row in a Word DOCX Document and returns the result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxTableRowRequest**](GetDocxTableRowRequest.md)| Document input request | 

### Return type

[**GetDocxTableRowResponse**](GetDocxTableRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxGetTables**
> GetDocxTablesResponse EditDocumentDocxGetTables(ctx, reqConfig)
Get all tables in Word DOCX document

Returns all the table objects in an Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxTablesRequest**](GetDocxTablesRequest.md)| Document input request | 

### Return type

[**GetDocxTablesResponse**](GetDocxTablesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxInsertCommentOnParagraph**
> InsertDocxCommentOnParagraphResponse EditDocumentDocxInsertCommentOnParagraph(ctx, reqConfig)
Insert a new comment into a Word DOCX document attached to a paragraph

Adds a new comment into a Word DOCX document attached to a paragraph and returns the result.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxInsertCommentOnParagraphRequest**](DocxInsertCommentOnParagraphRequest.md)| Document input request | 

### Return type

[**InsertDocxCommentOnParagraphResponse**](InsertDocxCommentOnParagraphResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxInsertImage**
> DocxInsertImageResponse EditDocumentDocxInsertImage(ctx, reqConfig)
Insert image into a Word DOCX document

Set the footer in a Word Document (DOCX).  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxInsertImageRequest**](DocxInsertImageRequest.md)| Document input request | 

### Return type

[**DocxInsertImageResponse**](DocxInsertImageResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxInsertParagraph**
> InsertDocxInsertParagraphResponse EditDocumentDocxInsertParagraph(ctx, reqConfig)
Insert a new paragraph into a Word DOCX document

Adds a new paragraph into a DOCX and returns the result.  You can insert at the beginning/end of a document, or before/after an existing object using its Path (location within the document).  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**InsertDocxInsertParagraphRequest**](InsertDocxInsertParagraphRequest.md)| Document input request | 

### Return type

[**InsertDocxInsertParagraphResponse**](InsertDocxInsertParagraphResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxInsertTable**
> InsertDocxTablesResponse EditDocumentDocxInsertTable(ctx, reqConfig)
Insert a new table into a Word DOCX document

Adds a new table into a DOCX and returns the result.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**InsertDocxTablesRequest**](InsertDocxTablesRequest.md)| Document input request | 

### Return type

[**InsertDocxTablesResponse**](InsertDocxTablesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxInsertTableRow**
> InsertDocxTableRowResponse EditDocumentDocxInsertTableRow(ctx, reqConfig)
Insert a new row into an existing table in a Word DOCX document

Adds a new table row into a DOCX Document and returns the result.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**InsertDocxTableRowRequest**](InsertDocxTableRowRequest.md)| Document input request | 

### Return type

[**InsertDocxTableRowResponse**](InsertDocxTableRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxPages**
> GetDocxPagesResponse EditDocumentDocxPages(ctx, reqConfig)
Get pages and content from a Word DOCX document

Returns the pages and contents of each page defined in the Word Document (DOCX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**GetDocxPagesRequest**](GetDocxPagesRequest.md)| Document input request | 

### Return type

[**GetDocxPagesResponse**](GetDocxPagesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxRemoveAllComments**
> string EditDocumentDocxRemoveAllComments(ctx, inputFile)
Remove all comments from a Word DOCX document

Removes all of the comments from a Word Document.

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
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxRemoveHeadersAndFooters**
> RemoveDocxHeadersAndFootersResponse EditDocumentDocxRemoveHeadersAndFooters(ctx, reqConfig)
Remove headers and footers from Word DOCX document

Remove all headers, or footers, or both from a Word Document (DOCX).  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**RemoveDocxHeadersAndFootersRequest**](RemoveDocxHeadersAndFootersRequest.md)| Document input request | 

### Return type

[**RemoveDocxHeadersAndFootersResponse**](RemoveDocxHeadersAndFootersResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxRemoveObject**
> DocxRemoveObjectResponse EditDocumentDocxRemoveObject(ctx, reqConfig)
Delete any object in a Word DOCX document

Delete any object, such as a paragraph, table, image, etc. from a Word Document (DOCX).  Pass in the Path of the object you would like to delete.  You can call other functions such as Get-Tables, Get-Images, Get-Body, etc. to get the paths of the objects in the document.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxRemoveObjectRequest**](DocxRemoveObjectRequest.md)| Document input request | 

### Return type

[**DocxRemoveObjectResponse**](DocxRemoveObjectResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxReplace**
> string EditDocumentDocxReplace(ctx, reqConfig)
Replace string in Word DOCX document

Replace all instances of a string in an Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**ReplaceStringRequest**](ReplaceStringRequest.md)| Document string replacement configuration input | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxReplaceMulti**
> string EditDocumentDocxReplaceMulti(ctx, reqConfig)
Replace multiple strings in Word DOCX document, return result

Replace all instances of multiple strings in an Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**MultiReplaceStringRequest**](MultiReplaceStringRequest.md)| Document string replacement configuration input | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxReplaceMultiEditSession**
> DocumentEditingEditSession EditDocumentDocxReplaceMultiEditSession(ctx, reqConfig)
Replace multiple strings in Word DOCX document, return edit session

Replace all instances of multiple strings in an Office Word Document (docx).  Returns an edit session URL so that you can chain together multiple edit operations without having to send the entire document contents back and forth multiple times.  Call the Finish Editing API to retrieve the final document once editing is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**MultiReplaceStringRequest**](MultiReplaceStringRequest.md)| Document string replacement configuration input | 

### Return type

[**DocumentEditingEditSession**](DocumentEditingEditSession.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxReplaceParagraph**
> ReplaceDocxParagraphResponse EditDocumentDocxReplaceParagraph(ctx, reqConfig)
Replace matching paragraphs in a Word DOCX document

Returns the edited Word Document (DOCX) format file with the matching paragraphs replaced as requested.  Replace a paragraph with another object such as an image.  Useful for performing templating operations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**ReplaceDocxParagraphRequest**](ReplaceDocxParagraphRequest.md)| Document input request | 

### Return type

[**ReplaceDocxParagraphResponse**](ReplaceDocxParagraphResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxSetCustomMetadataProperties**
> string EditDocumentDocxSetCustomMetadataProperties(ctx, input)
Set custom property metadata properties in Word DOCX document

Sets the custom property metadata for the metadata properties in an Office Word Document (docx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**DocxSetCustomMetadataPropertiesRequest**](DocxSetCustomMetadataPropertiesRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxSetFooter**
> DocxSetFooterResponse EditDocumentDocxSetFooter(ctx, reqConfig)
Set the footer in a Word DOCX document

Set the footer in a Word Document (DOCX).  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxSetFooterRequest**](DocxSetFooterRequest.md)| Document input request | 

### Return type

[**DocxSetFooterResponse**](DocxSetFooterResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxSetFooterAddPageNumber**
> DocxSetFooterResponse EditDocumentDocxSetFooterAddPageNumber(ctx, reqConfig)
Add page number to footer in a Word DOCX document

Set the footer in a Word Document (DOCX) to contain a page number.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxSetFooterAddPageNumberRequest**](DocxSetFooterAddPageNumberRequest.md)| Document input request | 

### Return type

[**DocxSetFooterResponse**](DocxSetFooterResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxSetFormFields**
> string EditDocumentDocxSetFormFields(ctx, reqConfig)
Set and fill values for form fields in a Word DOCX document

Modifies a Office Word Document (docx) by filling in form fields using the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxSetFormFieldsRequest**](DocxSetFormFieldsRequest.md)|  | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxSetHeader**
> DocxSetHeaderResponse EditDocumentDocxSetHeader(ctx, reqConfig)
Set the header in a Word DOCX document

Set the header in a Word Document (DOCX).  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**DocxSetHeaderRequest**](DocxSetHeaderRequest.md)| Document input request | 

### Return type

[**DocxSetHeaderResponse**](DocxSetHeaderResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxUpdateTableCell**
> UpdateDocxTableCellResponse EditDocumentDocxUpdateTableCell(ctx, reqConfig)
Update, set contents of a table cell in an existing table in a Word DOCX document

Sets the contents of a table cell into a DOCX Document and returns the result.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**UpdateDocxTableCellRequest**](UpdateDocxTableCellRequest.md)| Document input request | 

### Return type

[**UpdateDocxTableCellResponse**](UpdateDocxTableCellResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentDocxUpdateTableRow**
> UpdateDocxTableRowResponse EditDocumentDocxUpdateTableRow(ctx, reqConfig)
Update, set contents of a table row in an existing table in a Word DOCX document

Sets the contents of a table row into a DOCX Document and returns the result.  Call Finish Editing on the output URL to complete the operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**UpdateDocxTableRowRequest**](UpdateDocxTableRowRequest.md)| Document input request | 

### Return type

[**UpdateDocxTableRowResponse**](UpdateDocxTableRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentFinishEditing**
> string EditDocumentFinishEditing(ctx, reqConfig)
Finish editing document, and download result from document editing

Once done editing a document, download the result.  Begin editing a document by calling begin-editing, then perform operations, then call finish-editing to get the result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**FinishEditingRequest**](FinishEditingRequest.md)| Cloudmersive Document URL to complete editing on | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentPptxDeleteSlides**
> string EditDocumentPptxDeleteSlides(ctx, reqConfig)
Delete, remove slides from a PowerPoint PPTX presentation document

Edits the input PowerPoint PPTX presentation document to remove the specified slides

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**RemovePptxSlidesRequest**](RemovePptxSlidesRequest.md)| Presentation input request | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentPptxGetMacroInformation**
> GetMacrosResponse EditDocumentPptxGetMacroInformation(ctx, inputFile)
Get macro information from a PowerPoint PPTX/PPTM presentation document

Returns information about the Macros (e.g. VBA) defined in the PowerPoint Document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetMacrosResponse**](GetMacrosResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentPptxReplace**
> string EditDocumentPptxReplace(ctx, reqConfig)
Replace string in PowerPoint PPTX presentation

Replace all instances of a string in an Office PowerPoint Document (pptx)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**ReplaceStringRequest**](ReplaceStringRequest.md)| Replacement document configuration input | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxAppendRow**
> AppendXlsxRowResponse EditDocumentXlsxAppendRow(ctx, input)
Append row to a Excel XLSX spreadsheet, worksheet

Appends a row to the end of an Excel Spreadsheet worksheet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**AppendXlsxRowRequest**](AppendXlsxRowRequest.md)| Document input request | 

### Return type

[**AppendXlsxRowResponse**](AppendXlsxRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxClearCellByIndex**
> ClearXlsxCellResponse EditDocumentXlsxClearCellByIndex(ctx, input)
Clear cell contents in an Excel XLSX spreadsheet, worksheet by index

Clears, sets to blank, the contents of a specific cell in an Excel XLSX spreadsheet, worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ClearXlsxCellRequest**](ClearXlsxCellRequest.md)| Document input request | 

### Return type

[**ClearXlsxCellResponse**](ClearXlsxCellResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxClearRow**
> ClearXlsxRowResponse EditDocumentXlsxClearRow(ctx, input)
Clear row from a Excel XLSX spreadsheet, worksheet

Clears data from a specific row in the Excel Spreadsheet worksheet, leaving a blank row. Use the Get Rows And Cells API to enumerate available rows in a spreadsheet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ClearXlsxRowRequest**](ClearXlsxRowRequest.md)| Document input request | 

### Return type

[**ClearXlsxRowResponse**](ClearXlsxRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxCreateBlankSpreadsheet**
> CreateBlankSpreadsheetResponse EditDocumentXlsxCreateBlankSpreadsheet(ctx, input)
Create a blank Excel XLSX spreadsheet

Returns a blank Excel XLSX Spreadsheet (XLSX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**CreateBlankSpreadsheetRequest**](CreateBlankSpreadsheetRequest.md)| Document input request | 

### Return type

[**CreateBlankSpreadsheetResponse**](CreateBlankSpreadsheetResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxCreateSpreadsheetFromData**
> CreateSpreadsheetFromDataResponse EditDocumentXlsxCreateSpreadsheetFromData(ctx, input)
Create a new Excel XLSX spreadsheet from column and row data

Returns a new Excel XLSX Spreadsheet (XLSX) format file populated with column and row data specified as input

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**CreateSpreadsheetFromDataRequest**](CreateSpreadsheetFromDataRequest.md)| Document input request | 

### Return type

[**CreateSpreadsheetFromDataResponse**](CreateSpreadsheetFromDataResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxDeleteWorksheet**
> string EditDocumentXlsxDeleteWorksheet(ctx, reqConfig)
Delete, remove worksheet from an Excel XLSX spreadsheet document

Edits the input Excel XLSX spreadsheet document to remove the specified worksheet (tab).  Use the Get Worksheets API to enumerate available worksheets in a spreadsheet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reqConfig** | [**RemoveXlsxWorksheetRequest**](RemoveXlsxWorksheetRequest.md)| Spreadsheet input request | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxDisableSharedWorkbook**
> DisableSharedWorkbookResponse EditDocumentXlsxDisableSharedWorkbook(ctx, input)
Disable Shared Workbook (legacy) in Excel XLSX spreadsheet

Disable the Shared Workbook (legacy) mode in an Excel XLSX spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**DisableSharedWorkbookRequest**](DisableSharedWorkbookRequest.md)| Document input request | 

### Return type

[**DisableSharedWorkbookResponse**](DisableSharedWorkbookResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxEnableSharedWorkbook**
> EnableSharedWorkbookResponse EditDocumentXlsxEnableSharedWorkbook(ctx, input)
Enable Shared Workbook (legacy) in Excel XLSX spreadsheet

Enables the Shared Workbook (legacy) mode in an Excel XLSX spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**EnableSharedWorkbookRequest**](EnableSharedWorkbookRequest.md)| Document input request | 

### Return type

[**EnableSharedWorkbookResponse**](EnableSharedWorkbookResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetCellByIdentifier**
> GetXlsxCellByIdentifierResponse EditDocumentXlsxGetCellByIdentifier(ctx, input)
Get cell from an Excel XLSX spreadsheet, worksheet by cell identifier

Returns the value of a specific cell based on its identifier (e.g. A1, B22, C33, etc.) in the Excel Spreadsheet worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxCellByIdentifierRequest**](GetXlsxCellByIdentifierRequest.md)| Document input request | 

### Return type

[**GetXlsxCellByIdentifierResponse**](GetXlsxCellByIdentifierResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetCellByIndex**
> GetXlsxCellResponse EditDocumentXlsxGetCellByIndex(ctx, input)
Get cell from an Excel XLSX spreadsheet, worksheet by index

Returns the value and definition of a specific cell in a specific row in the Excel Spreadsheet worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxCellRequest**](GetXlsxCellRequest.md)| Document input request | 

### Return type

[**GetXlsxCellResponse**](GetXlsxCellResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetColumns**
> GetXlsxColumnsResponse EditDocumentXlsxGetColumns(ctx, input)
Get columns from a Excel XLSX spreadsheet, worksheet

Returns the columns defined in the Excel Spreadsheet worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxColumnsRequest**](GetXlsxColumnsRequest.md)| Document input request | 

### Return type

[**GetXlsxColumnsResponse**](GetXlsxColumnsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetImages**
> GetXlsxImagesResponse EditDocumentXlsxGetImages(ctx, input)
Get images from a Excel XLSX spreadsheet, worksheet

Returns the images defined in the Excel Spreadsheet worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxImagesRequest**](GetXlsxImagesRequest.md)| Document input request | 

### Return type

[**GetXlsxImagesResponse**](GetXlsxImagesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetMacroInformation**
> GetMacrosResponse EditDocumentXlsxGetMacroInformation(ctx, inputFile)
Get macro information from a Excel XLSX/XLSM spreadsheet, worksheet

Returns information about the Macros (e.g. VBA) defined in the Excel Spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**GetMacrosResponse**](GetMacrosResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetRowsAndCells**
> GetXlsxRowsAndCellsResponse EditDocumentXlsxGetRowsAndCells(ctx, input)
Get rows and cells from a Excel XLSX spreadsheet, worksheet

Returns the rows and cells defined in the Excel Spreadsheet worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxRowsAndCellsRequest**](GetXlsxRowsAndCellsRequest.md)| Document input request | 

### Return type

[**GetXlsxRowsAndCellsResponse**](GetXlsxRowsAndCellsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetSpecificRow**
> GetXlsxSpecificRowResponse EditDocumentXlsxGetSpecificRow(ctx, input)
Get a specific row from a Excel XLSX spreadsheet, worksheet by path

Returns the specific row and its cells defined in the Excel Spreadsheet worksheet based on the specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxSpecificRowRequest**](GetXlsxSpecificRowRequest.md)| Document input request | 

### Return type

[**GetXlsxSpecificRowResponse**](GetXlsxSpecificRowResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetStyles**
> GetXlsxStylesResponse EditDocumentXlsxGetStyles(ctx, input)
Get styles from a Excel XLSX spreadsheet, worksheet

Returns the style defined in the Excel Spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxStylesRequest**](GetXlsxStylesRequest.md)| Document input request | 

### Return type

[**GetXlsxStylesResponse**](GetXlsxStylesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxGetWorksheets**
> GetXlsxWorksheetsResponse EditDocumentXlsxGetWorksheets(ctx, input)
Get worksheets from a Excel XLSX spreadsheet

Returns the worksheets (tabs) defined in the Excel Spreadsheet (XLSX) format file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetXlsxWorksheetsRequest**](GetXlsxWorksheetsRequest.md)| Document input request | 

### Return type

[**GetXlsxWorksheetsResponse**](GetXlsxWorksheetsResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxInsertWorksheet**
> InsertXlsxWorksheetResponse EditDocumentXlsxInsertWorksheet(ctx, input)
Insert a new worksheet into an Excel XLSX spreadsheet

Inserts a new worksheet into an Excel Spreadsheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**InsertXlsxWorksheetRequest**](InsertXlsxWorksheetRequest.md)| Document input request | 

### Return type

[**InsertXlsxWorksheetResponse**](InsertXlsxWorksheetResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxRenameWorksheet**
> RenameXlsxWorksheetResponse EditDocumentXlsxRenameWorksheet(ctx, input)
Rename a specific worksheet in a Excel XLSX spreadsheet

Edits the input Excel XLSX spreadsheet document to rename a specified worksheet (tab).  Use the Get Worksheets API to enumerate available worksheets in a spreadsheet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**RenameXlsxWorksheetRequest**](RenameXlsxWorksheetRequest.md)| Document input request | 

### Return type

[**RenameXlsxWorksheetResponse**](RenameXlsxWorksheetResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxSetCellByIdentifier**
> SetXlsxCellByIdentifierResponse EditDocumentXlsxSetCellByIdentifier(ctx, input)
Set, update cell contents in an Excel XLSX spreadsheet, worksheet by cell identifier

Sets, updates the contents of a specific cell in an Excel XLSX spreadsheet, worksheet using its cell identifier (e.g. A1, B22, C33) in the worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**SetXlsxCellByIdentifierRequest**](SetXlsxCellByIdentifierRequest.md)| Document input request | 

### Return type

[**SetXlsxCellByIdentifierResponse**](SetXlsxCellByIdentifierResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDocumentXlsxSetCellByIndex**
> SetXlsxCellResponse EditDocumentXlsxSetCellByIndex(ctx, input)
Set, update cell contents in an Excel XLSX spreadsheet, worksheet by index

Sets, updates the contents of a specific cell in an Excel XLSX spreadsheet, worksheet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**SetXlsxCellRequest**](SetXlsxCellRequest.md)| Document input request | 

### Return type

[**SetXlsxCellResponse**](SetXlsxCellResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

