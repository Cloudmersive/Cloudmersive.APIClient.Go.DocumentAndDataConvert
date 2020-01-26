/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of the GetGender operation
type GetGenderResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Gender for this name; possible values are Male, Female, and Neutral (can be applied to Male or Female)
	Gender string `json:"Gender,omitempty"`
}