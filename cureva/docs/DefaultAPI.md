# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClinicalmatchPost**](DefaultAPI.md#V1ClinicalmatchPost) | **Post** /v1/clinicalmatch | Submit patient notes for clinical trial matching
[**VersionGet**](DefaultAPI.md#VersionGet) | **Get** /version | Get API version



## V1ClinicalmatchPost

> V1ClinicalmatchPost200Response V1ClinicalmatchPost(ctx).PatientNotesFile(patientNotesFile).FileSizeBytes(fileSizeBytes).Location(location).Execute()

Submit patient notes for clinical trial matching



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	patientNotesFile := os.NewFile(1234, "some_file") // *os.File | The file containing notes about the patient from the doctor visit in plain text format.
	fileSizeBytes := string(BYTE_ARRAY_DATA_HERE) // string | The number of bytes in the patient notes file.
	location := *openapiclient.NewV1ClinicalmatchPostRequestLocation(float32(42.349806), float32(-71.06149)) // V1ClinicalmatchPostRequestLocation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1ClinicalmatchPost(context.Background()).PatientNotesFile(patientNotesFile).FileSizeBytes(fileSizeBytes).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1ClinicalmatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ClinicalmatchPost`: V1ClinicalmatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1ClinicalmatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ClinicalmatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patientNotesFile** | ***os.File** | The file containing notes about the patient from the doctor visit in plain text format. | 
 **fileSizeBytes** | **string** | The number of bytes in the patient notes file. | 
 **location** | [**V1ClinicalmatchPostRequestLocation**](V1ClinicalmatchPostRequestLocation.md) |  | 

### Return type

[**V1ClinicalmatchPost200Response**](V1ClinicalmatchPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionGet

> VersionGet200Response VersionGet(ctx).Execute()

Get API version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionGet`: VersionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionGetRequest struct via the builder pattern


### Return type

[**VersionGet200Response**](VersionGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

