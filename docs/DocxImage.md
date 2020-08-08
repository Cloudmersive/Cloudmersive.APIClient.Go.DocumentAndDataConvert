# DocxImage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The Path of the location of this object; leave blank for new tables | [optional] [default to null]
**ImageName** | **string** | The Name of the image | [optional] [default to null]
**ImageId** | **int64** | The Id of the image | [optional] [default to null]
**ImageDescription** | **string** | The Description of the image | [optional] [default to null]
**ImageWidth** | **int64** | Width of the image in EMUs (English Metric Units); set to 0 to default to page width and aspect-ratio based height | [optional] [default to null]
**ImageHeight** | **int64** | Height of the image in EMUs (English Metric Units); set to 0 to default to page width and aspect-ratio based height | [optional] [default to null]
**XOffset** | **int64** | X (horizontal) offset of the image | [optional] [default to null]
**YOffset** | **int64** | Y (vertical) offset of the image | [optional] [default to null]
**ImageDataEmbedId** | **string** | Read-only; internal ID for the image contents | [optional] [default to null]
**ImageDataContentType** | **string** | Read-only; image data MIME content-type | [optional] [default to null]
**ImageInternalFileName** | **string** | Read-only; internal file name/path for the image | [optional] [default to null]
**ImageContentsURL** | **string** | URL to the image contents; file is stored in an in-memory cache and will be deleted.  Call Finish-Editing to get the contents. | [optional] [default to null]
**InlineWithText** | **bool** | True if the image is inline with the text; false if it is floating | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


