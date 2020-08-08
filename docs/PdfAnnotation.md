# PdfAnnotation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the annotation; this is often the author of the annotation in Acrobat-created PDF files | [optional] [default to null]
**AnnotationType** | **string** | Type of the annotation; possible values are Text | [optional] [default to null]
**PageNumber** | **int32** | The 1-based index of the page containing the annotation | [optional] [default to null]
**AnnotationIndex** | **int32** | The 0-based index of the annotation in the document | [optional] [default to null]
**Subject** | **string** | Subject of the annotation | [optional] [default to null]
**TextContents** | **string** | Text contents of the annotation | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Date that the annotation was created | [optional] [default to null]
**ModifiedDate** | [**time.Time**](time.Time.md) | Date that the annotation was last modified | [optional] [default to null]
**LeftX** | **float64** | Left X coordinate for the location of the annotation | [optional] [default to null]
**TopY** | **float64** | Top Y coordination of the location of the annotation | [optional] [default to null]
**Width** | **float64** | Width of the annotation | [optional] [default to null]
**Height** | **float64** | Height of the annotation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


