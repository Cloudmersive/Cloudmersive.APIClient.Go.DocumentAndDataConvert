/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

import (
	"time"
)

// Name and value pair for a custom-defined DOCX metadata property
type DocxMetadataCustomProperty struct {
	// Name of the property
	PropertyName string `json:"PropertyName,omitempty"`
	// Data type of the property; possible values are \"string\", \"integer\", \"double\" or \"date\"
	PropertyDataType string `json:"PropertyDataType,omitempty"`
	// If the property is of a string data type, then provides the string value if available
	StringValue string `json:"StringValue,omitempty"`
	// If the property is of a integer data type, then provides the integer value if available
	IntegerValue int64 `json:"IntegerValue,omitempty"`
	// If the property is of a double floating point data type, then provides the double value if available
	DoubleValue float64 `json:"DoubleValue,omitempty"`
	// If the property is of a date time data type, then provides the date time value if available
	DateValue time.Time `json:"DateValue,omitempty"`
}
