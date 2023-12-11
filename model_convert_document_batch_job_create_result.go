/*
 * convertapi
 *
 * Convert API lets you effortlessly convert file formats and types.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package GoCloudmersiveDocumentConvertApiClient

// Result of performing a convert documentbatch job operation
type ConvertDocumentBatchJobCreateResult struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// When creating a job, an Async Job ID is returned.  Use the GetAsyncJobStatus API to check on the status of this job using the AsyncJobID and get the result when it finishes
	AsyncJobID string `json:"AsyncJobID,omitempty"`
}
