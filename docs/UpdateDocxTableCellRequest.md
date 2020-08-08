# UpdateDocxTableCellRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFileBytes** | **string** | Optional: Bytes of the input file to operate on | [optional] [default to null]
**InputFileUrl** | **string** | Optional: URL of a file to operate on as input.  This can be a public URL, or you can also use the begin-editing API to upload a document and pass in the secure URL result from that operation as the URL here (this URL is not public). | [optional] [default to null]
**CellToUpdate** | [***DocxTableCell**](DocxTableCell.md) | Table cell contents you would like to update the cell with | [optional] [default to null]
**TableRowIndex** | **int32** | 0-based index of the Table Row to update | [optional] [default to null]
**TableCellIndex** | **int32** | 0-based index of the Table Cell (within the row) to update | [optional] [default to null]
**ExistingTablePath** | **string** | Required; the path to the existing table to modify | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


