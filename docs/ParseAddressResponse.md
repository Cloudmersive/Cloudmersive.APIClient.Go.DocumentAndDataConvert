# ParseAddressResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | **bool** | True if the parsing operation was successful, false otherwise | [optional] [default to null]
**Building** | **string** | The name of the building, house or structure if applicable, such as \&quot;Cloudmersive Building 2\&quot;.  This will often by null. | [optional] [default to null]
**StreetNumber** | **string** | The street number or house number of the address.  For example, in the address \&quot;1600 Pennsylvania Avenue NW\&quot; the street number would be \&quot;1600\&quot;.  This value will typically be populated for most addresses. | [optional] [default to null]
**Street** | **string** | The name of the street or road of the address.  For example, in the address \&quot;1600 Pennsylvania Avenue NW\&quot; the street number would be \&quot;Pennsylvania Avenue NW\&quot;. | [optional] [default to null]
**City** | **string** | The city of the address. | [optional] [default to null]
**StateOrProvince** | **string** | The state or province of the address. | [optional] [default to null]
**PostalCode** | **string** | The postal code or zip code of the address. | [optional] [default to null]
**CountryFullName** | **string** | Country of the address, if present in the address.  If not included in the address it will be null. | [optional] [default to null]
**ISOTwoLetterCode** | **string** | Two-letter ISO 3166-1 country code | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


