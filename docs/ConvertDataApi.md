# \ConvertDataApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertDataCsvToJson**](ConvertDataApi.md#ConvertDataCsvToJson) | **Post** /convert/csv/to/json | Convert CSV to JSON conversion
[**ConvertDataCsvToXml**](ConvertDataApi.md#ConvertDataCsvToXml) | **Post** /convert/csv/to/xml | Convert CSV to XML conversion
[**ConvertDataJsonToXml**](ConvertDataApi.md#ConvertDataJsonToXml) | **Post** /convert/json/to/xml | Convert JSON to XML conversion
[**ConvertDataXlsToJson**](ConvertDataApi.md#ConvertDataXlsToJson) | **Post** /convert/xls/to/json | Convert Excel (97-2003) XLS to JSON conversion
[**ConvertDataXlsxToJson**](ConvertDataApi.md#ConvertDataXlsxToJson) | **Post** /convert/xlsx/to/json | Convert Excel XLSX to JSON conversion
[**ConvertDataXlsxToXml**](ConvertDataApi.md#ConvertDataXlsxToXml) | **Post** /convert/xlsx/to/xml | Convert Excel XLSX to XML conversion
[**ConvertDataXmlEditAddAttributeWithXPath**](ConvertDataApi.md#ConvertDataXmlEditAddAttributeWithXPath) | **Post** /convert/xml/edit/xpath/add-attribute | Adds an attribute to all XML nodes matching XPath expression
[**ConvertDataXmlEditAddChildWithXPath**](ConvertDataApi.md#ConvertDataXmlEditAddChildWithXPath) | **Post** /convert/xml/edit/xpath/add-child | Adds an XML node as a child to XML nodes matching XPath expression
[**ConvertDataXmlEditRemoveAllChildNodesWithXPath**](ConvertDataApi.md#ConvertDataXmlEditRemoveAllChildNodesWithXPath) | **Post** /convert/xml/edit/xpath/remove-all-children | Removes, deletes all children of nodes matching XPath expression, but does not remove the nodes
[**ConvertDataXmlEditReplaceWithXPath**](ConvertDataApi.md#ConvertDataXmlEditReplaceWithXPath) | **Post** /convert/xml/edit/xpath/replace | Replaces XML nodes matching XPath expression with new node
[**ConvertDataXmlEditSetValueWithXPath**](ConvertDataApi.md#ConvertDataXmlEditSetValueWithXPath) | **Post** /convert/xml/edit/xpath/set-value | Sets the value contents of XML nodes matching XPath expression
[**ConvertDataXmlFilterWithXPath**](ConvertDataApi.md#ConvertDataXmlFilterWithXPath) | **Post** /convert/xml/select/xpath | Filter, select XML nodes using XPath expression, get results
[**ConvertDataXmlQueryWithXQuery**](ConvertDataApi.md#ConvertDataXmlQueryWithXQuery) | **Post** /convert/xml/query/xquery | Query an XML file using XQuery query, get results
[**ConvertDataXmlQueryWithXQueryMulti**](ConvertDataApi.md#ConvertDataXmlQueryWithXQueryMulti) | **Post** /convert/xml/query/xquery/multi | Query multiple XML files using XQuery query, get results
[**ConvertDataXmlRemoveWithXPath**](ConvertDataApi.md#ConvertDataXmlRemoveWithXPath) | **Post** /convert/xml/edit/xpath/remove | Remove, delete XML nodes and items matching XPath expression
[**ConvertDataXmlToJson**](ConvertDataApi.md#ConvertDataXmlToJson) | **Post** /convert/xml/to/json | Convert XML to JSON conversion
[**ConvertDataXmlTransformWithXsltToXml**](ConvertDataApi.md#ConvertDataXmlTransformWithXsltToXml) | **Post** /convert/xml/transform/xslt/to/xml | Transform XML document file with XSLT into a new XML document


# **ConvertDataCsvToJson**
> interface{} ConvertDataCsvToJson(ctx, inputFile, optional)
Convert CSV to JSON conversion

Convert a CSV file to a JSON object array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***ConvertDataCsvToJsonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConvertDataCsvToJsonOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnNamesFromFirstRow** | **optional.Bool**| Optional; If true, the first row will be used as the labels for the columns; if false, columns will be named Column0, Column1, etc.  Default is true.  Set to false if you are not using column headings, or have an irregular column structure. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataCsvToXml**
> string ConvertDataCsvToXml(ctx, inputFile, optional)
Convert CSV to XML conversion

Convert a CSV file to a XML file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 
 **optional** | ***ConvertDataCsvToXmlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConvertDataCsvToXmlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnNamesFromFirstRow** | **optional.Bool**| Optional; If true, the first row will be used as the labels for the columns; if false, columns will be named Column0, Column1, etc.  Default is true.  Set to false if you are not using column headings, or have an irregular column structure. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataJsonToXml**
> string ConvertDataJsonToXml(ctx, jsonObject)
Convert JSON to XML conversion

Convert a JSON object into XML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jsonObject** | [**interface{}**](interface{}.md)| Input JSON to convert to XML | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXlsToJson**
> interface{} ConvertDataXlsToJson(ctx, inputFile)
Convert Excel (97-2003) XLS to JSON conversion

Convert an Excel (97-2003) XLS file to a JSON object array

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXlsxToJson**
> string ConvertDataXlsxToJson(ctx, inputFile)
Convert Excel XLSX to JSON conversion

Convert an Excel XLSX file to a JSON object array

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

# **ConvertDataXlsxToXml**
> string ConvertDataXlsxToXml(ctx, inputFile)
Convert Excel XLSX to XML conversion

Convert an Excel XLSX file to a XML file

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

# **ConvertDataXmlEditAddAttributeWithXPath**
> XmlAddAttributeWithXPathResult ConvertDataXmlEditAddAttributeWithXPath(ctx, inputFile, xPathExpression, xmlAttributeName, xmlAttributeValue)
Adds an attribute to all XML nodes matching XPath expression

Return the reuslts of editing an XML document by adding an attribute to all of the nodes that match an input XPath expression.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **xPathExpression** | **string**| Valid XML XPath query expression | 
  **xmlAttributeName** | **string**| Name of the XML attribute to add | 
  **xmlAttributeValue** | **string**| Value of the XML attribute to add | 

### Return type

[**XmlAddAttributeWithXPathResult**](XmlAddAttributeWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlEditAddChildWithXPath**
> XmlAddChildWithXPathResult ConvertDataXmlEditAddChildWithXPath(ctx, inputFile, xPathExpression, xmlNodeToAdd)
Adds an XML node as a child to XML nodes matching XPath expression

Return the reuslts of editing an XML document by adding an XML node as a child to all of the nodes that match an input XPath expression.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **xPathExpression** | **string**| Valid XML XPath query expression | 
  **xmlNodeToAdd** | **string**| XML Node to add as a child | 

### Return type

[**XmlAddChildWithXPathResult**](XmlAddChildWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlEditRemoveAllChildNodesWithXPath**
> XmlRemoveAllChildrenWithXPathResult ConvertDataXmlEditRemoveAllChildNodesWithXPath(ctx, inputFile, xPathExpression)
Removes, deletes all children of nodes matching XPath expression, but does not remove the nodes

Return the reuslts of editing an XML document by removing all child nodes of the nodes that match an input XPath expression.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **xPathExpression** | **string**| Valid XML XPath query expression | 

### Return type

[**XmlRemoveAllChildrenWithXPathResult**](XmlRemoveAllChildrenWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlEditReplaceWithXPath**
> XmlReplaceWithXPathResult ConvertDataXmlEditReplaceWithXPath(ctx, inputFile, xPathExpression, xmlNodeReplacement)
Replaces XML nodes matching XPath expression with new node

Return the reuslts of editing an XML document by replacing all of the nodes that match an input XPath expression with a new XML node expression.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **xPathExpression** | **string**| Valid XML XPath query expression | 
  **xmlNodeReplacement** | **string**| XML Node replacement content | 

### Return type

[**XmlReplaceWithXPathResult**](XmlReplaceWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlEditSetValueWithXPath**
> XmlSetValueWithXPathResult ConvertDataXmlEditSetValueWithXPath(ctx, inputFile, xPathExpression, xmlValue)
Sets the value contents of XML nodes matching XPath expression

Return the reuslts of editing an XML document by setting the contents of all of the nodes that match an input XPath expression.  Supports elements and attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **xPathExpression** | **string**| Valid XML XPath query expression | 
  **xmlValue** | **string**| XML Value to set into the matching XML nodes | 

### Return type

[**XmlSetValueWithXPathResult**](XmlSetValueWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlFilterWithXPath**
> XmlFilterWithXPathResult ConvertDataXmlFilterWithXPath(ctx, xPathExpression, inputFile)
Filter, select XML nodes using XPath expression, get results

Return the reuslts of filtering, selecting an XML document with an XPath expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xPathExpression** | **string**| Valid XML XPath query expression | 
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**XmlFilterWithXPathResult**](XmlFilterWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlQueryWithXQuery**
> XmlQueryWithXQueryResult ConvertDataXmlQueryWithXQuery(ctx, inputFile, xQuery)
Query an XML file using XQuery query, get results

Return the reuslts of querying a single XML document with an XQuery expression.  Supports XQuery 3.1 and earlier.  This API is optimized for a single XML document as input.  Provided XML document is automatically loaded as the default context; to access elements in the document, simply refer to them without a document reference, such as bookstore/book

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **xQuery** | **string**| Valid XML XQuery 3.1 or earlier query expression; multi-line expressions are supported | 

### Return type

[**XmlQueryWithXQueryResult**](XmlQueryWithXQueryResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlQueryWithXQueryMulti**
> XmlQueryWithXQueryMultiResult ConvertDataXmlQueryWithXQueryMulti(ctx, inputFile1, xQuery, optional)
Query multiple XML files using XQuery query, get results

Return the reuslts of querying an XML document with an XQuery expression.  Supports XQuery 3.1 and earlier.  This API is optimized for multiple XML documents as input.  You can refer to the contents of a given document by name, for example doc(\"books.xml\") or doc(\"restaurants.xml\") if you included two input files named books.xml and restaurants.xml.  If input files contain no file name, they will default to file names input1.xml, input2.xml and so on.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile1** | ***os.File**| First input XML file to perform the operation on. | 
  **xQuery** | **string**| Valid XML XQuery 3.1 or earlier query expression; multi-line expressions are supported | 
 **optional** | ***ConvertDataXmlQueryWithXQueryMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConvertDataXmlQueryWithXQueryMultiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputFile2** | **optional.Interface of *os.File**| Second input XML file to perform the operation on. | 
 **inputFile3** | **optional.Interface of *os.File**| Third input XML file to perform the operation on. | 
 **inputFile4** | **optional.Interface of *os.File**| Fourth input XML file to perform the operation on. | 
 **inputFile5** | **optional.Interface of *os.File**| Fifth input XML file to perform the operation on. | 
 **inputFile6** | **optional.Interface of *os.File**| Sixth input XML file to perform the operation on. | 
 **inputFile7** | **optional.Interface of *os.File**| Seventh input XML file to perform the operation on. | 
 **inputFile8** | **optional.Interface of *os.File**| Eighth input XML file to perform the operation on. | 
 **inputFile9** | **optional.Interface of *os.File**| Ninth input XML file to perform the operation on. | 
 **inputFile10** | **optional.Interface of *os.File**| Tenth input XML file to perform the operation on. | 

### Return type

[**XmlQueryWithXQueryMultiResult**](XmlQueryWithXQueryMultiResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlRemoveWithXPath**
> XmlRemoveWithXPathResult ConvertDataXmlRemoveWithXPath(ctx, xPathExpression, inputFile)
Remove, delete XML nodes and items matching XPath expression

Return the reuslts of editing an XML document by removing all of the nodes that match an input XPath expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xPathExpression** | **string**| Valid XML XPath query expression | 
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**XmlRemoveWithXPathResult**](XmlRemoveWithXPathResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlToJson**
> interface{} ConvertDataXmlToJson(ctx, inputFile)
Convert XML to JSON conversion

Convert an XML string or file into JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input file to perform the operation on. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDataXmlTransformWithXsltToXml**
> string ConvertDataXmlTransformWithXsltToXml(ctx, inputFile, transformFile)
Transform XML document file with XSLT into a new XML document

Convert an XML string or file into JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputFile** | ***os.File**| Input XML file to perform the operation on. | 
  **transformFile** | ***os.File**| Input XSLT file to use to transform the input XML file. | 

### Return type

**string**

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

