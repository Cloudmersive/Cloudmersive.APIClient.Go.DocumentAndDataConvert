# DocxComment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Path to the comment in the document | [optional] [default to null]
**Author** | **string** | Author name of the comment | [optional] [default to null]
**AuthorInitials** | **string** | Initials of the author of the comment | [optional] [default to null]
**CommentText** | **string** | Text content of the comment | [optional] [default to null]
**CommentDate** | [**time.Time**](time.Time.md) | Date timestamp of the comment | [optional] [default to null]
**IsTopLevel** | **bool** | True if the comment is at the top level, false if this comment is a child reply of another comment | [optional] [default to null]
**IsReply** | **bool** | True if this comment is a reply to another comment, false otherwise | [optional] [default to null]
**ParentCommentPath** | **string** | Path to the parent of this comment, if this comment is a reply, otherwise this value will be null | [optional] [default to null]
**Done** | **bool** | True if this comment is marked as Done in Word, otherwise it is false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


