# GetXlsxCellRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFileBytes** | **string** | Optional: Bytes of the input file to operate on | [optional] [default to null]
**InputFileUrl** | **string** | Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public). | [optional] [default to null]
**WorksheetToQuery** | [***XlsxWorksheet**](XlsxWorksheet.md) | Optional; Worksheet (tab) within the spreadsheet to get the rows and cells of; leave blank to default to the first worksheet | [optional] [default to null]
**RowIndex** | **int32** | 0-based index of the row, 0, 1, 2, ... to retrieve | [optional] [default to null]
**CellIndex** | **int32** | 0-based index of the cell, 0, 1, 2, ... in the row to retrieve | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


