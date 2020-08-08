# DocxTable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableID** | **string** | The ID of the table; leave blank for new tables | [optional] [default to null]
**Path** | **string** | The Path of the location of this table object; leave blank for new tables | [optional] [default to null]
**Width** | **string** | The Width of the table, or 0 if not specified | [optional] [default to null]
**WidthType** | **string** | The Width configuration type of the table | [optional] [default to null]
**TableRows** | [**[]DocxTableRow**](DocxTableRow.md) | Rows in the table; this is where the contents is located | [optional] [default to null]
**TopBorderType** | **string** | Type for the top border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**TopBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**TopBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**TopBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**BottomBorderType** | **string** | Type for the bottom border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**BottomBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**BottomBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**BottomBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**LeftBorderType** | **string** | Type for the left border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**LeftBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**LeftBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**LeftBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**RightBorderType** | **string** | Type for the right border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**RightBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**RightBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**RightBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**CellHorizontalBorderType** | **string** | Type for the cell horizontal border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**CellHorizontalBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**CellHorizontalBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**CellHorizontalBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**CellVerticalBorderType** | **string** | Type for the cell vertical border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**CellVerticalBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**CellVerticalBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**CellVerticalBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**StartBorderType** | **string** | Type for the start border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**StartBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**StartBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**StartBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**EndBorderType** | **string** | Type for the end border - can be a Single, DashDotStroked, Dashed, DashSmallGap, DotDash, DotDotDash, Dotted, Double, DoubleWave, Inset, Nil, None, Outset, Thick, ThickThinLargeGap, ThickThinMediumGap, ThickThinSmallGap, ThinThickLargeGap, ThinThickMediumGap, ThinThickSmallGap, ThinThickThinLargeGap, ThinThickThinMediumGap, ThinThickThinSmallGap, ThreeDEmboss, ThreeDEngrave, Triple, Wave | [optional] [default to null]
**EndBorderSize** | **int32** | Width of the border in points (1/72nd of an inch) | [optional] [default to null]
**EndBorderSpace** | **int32** | Spacing around the border in points (1/72nd of an inch) | [optional] [default to null]
**EndBorderColor** | **string** | HTML-style color hex value (do not include a #) | [optional] [default to null]
**TableIndentationMode** | **string** | Table indentation type | [optional] [default to null]
**TableIndentationWidth** | **int32** | Table indentation width | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


