# \ConvertTemplateApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertTemplateApplyDocxTemplate**](ConvertTemplateApi.md#ConvertTemplateApplyDocxTemplate) | **Post** /convert/template/docx/apply | Apply Word DOCX template
[**ConvertTemplateApplyHtmlTemplate**](ConvertTemplateApi.md#ConvertTemplateApplyHtmlTemplate) | **Post** /convert/template/html/apply | Apply HTML template


# **ConvertTemplateApplyDocxTemplate**
> string ConvertTemplateApplyDocxTemplate(ctx, inputFile, optional)
Apply Word DOCX template

Apply operations to fill in a Word DOCX template by replacing target template/placeholder strings in the DOCX with values, generating a final Word DOCX result.  For example, you could create a Word Document invoice containing strings such as \"{FirstName}\" and \"{LastName}\" and then replace these values with \"John\" and \"Smith\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***ConvertTemplateApplyDocxTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConvertTemplateApplyDocxTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateDefinition** | [**optional.Interface of interface{}**](.md)| Template definition for the document, including what values to replace | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertTemplateApplyHtmlTemplate**
> HtmlTemplateApplicationResponse ConvertTemplateApplyHtmlTemplate(ctx, value)
Apply HTML template

Apply operations to fill in an HTML template, generating a final HTML result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | [**HtmlTemplateApplicationRequest**](HtmlTemplateApplicationRequest.md)| Operations to apply to template | 

### Return type

[**HtmlTemplateApplicationResponse**](HtmlTemplateApplicationResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

