# HtmlSsrfThreatCheckResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | **bool** | True if the document is valid and has no errors, false otherwise | [optional] [default to null]
**IsThreat** | **bool** | True if the document contains an SSRF threat, false otherwise | [optional] [default to null]
**ThreatLinks** | [**[]HtmlThreatLink**](HtmlThreatLink.md) | Links found in the input HTML that contains threats | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


