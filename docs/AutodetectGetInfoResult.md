# AutodetectGetInfoResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | **bool** | True if the operation was successful, false otherwise | [optional] [default to null]
**DetectedFileExtension** | **string** | Detected file extension of the file format, with a leading period | [optional] [default to null]
**DetectedMimeType** | **string** | MIME type of this file extension | [optional] [default to null]
**PageCount** | **int64** | Number of pages in a page-based document; for presentations, this is the number of slides and for a spreadsheet this is the number of worksheets.  Contains 0 when the page count cannot be determined, or if the concept of page count does not apply (e.g. for an image) | [optional] [default to null]
**Author** | **string** | User name of the creator/author of the document, if available, null if not available | [optional] [default to null]
**DateModified** | [**time.Time**](time.Time.md) | The timestamp that the document was last modified, if available, null if not available | [optional] [default to null]
**AlternateFileTypeCandidates** | [**[]AlternateFileFormatCandidate**](AlternateFileFormatCandidate.md) | Alternate file type options and their probability | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


