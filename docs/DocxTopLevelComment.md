# DocxTopLevelComment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Path to the comment in the document | [optional] [default to null]
**Author** | **string** | Author name of the comment | [optional] [default to null]
**AuthorInitials** | **string** | Initials of the author of the comment | [optional] [default to null]
**CommentText** | **string** | Text content of the comment | [optional] [default to null]
**CommentDate** | [**time.Time**](time.Time.md) | Date timestamp of the comment | [optional] [default to null]
**ReplyChildComments** | [**[]DocxComment**](DocxComment.md) | Child comments, that are replies to this one | [optional] [default to null]
**Done** | **bool** | True if this comment is marked as Done in Word, otherwise it is false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


