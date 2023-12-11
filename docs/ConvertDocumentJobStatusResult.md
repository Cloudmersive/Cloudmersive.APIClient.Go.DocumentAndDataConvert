# ConvertDocumentJobStatusResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | **bool** | Tru eif the operation to check the status of the job was successful, false otherwise | [optional] [default to null]
**AsyncJobStatus** | **string** | Returns the job status of the Async Job, if applicable.  Possible states are STARTED and COMPLETED | [optional] [default to null]
**AsyncJobID** | **string** | When the job exceeds 25 pages, an Async Job ID is returned.  Use the CheckPdfOcrJobStatus API to check on the status of this job using the AsyncJobID and get the result when it finishes | [optional] [default to null]
**OutputFileResult** | **string** | Output file result (if applicable) | [optional] [default to null]
**ErrorMessage** | **string** | Error message (if any) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


