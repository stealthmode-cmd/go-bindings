/*
Clinical Match API

A simple API to match patients to clinical trials.

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiV1ClinicalmatchPostRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	patientNotesFile *os.File
	fileSizeBytes *string
	location *V1ClinicalmatchPostRequestLocation
}

// The file containing notes about the patient from the doctor visit in plain text format.
func (r ApiV1ClinicalmatchPostRequest) PatientNotesFile(patientNotesFile *os.File) ApiV1ClinicalmatchPostRequest {
	r.patientNotesFile = patientNotesFile
	return r
}

// The number of bytes in the patient notes file.
func (r ApiV1ClinicalmatchPostRequest) FileSizeBytes(fileSizeBytes string) ApiV1ClinicalmatchPostRequest {
	r.fileSizeBytes = &fileSizeBytes
	return r
}

func (r ApiV1ClinicalmatchPostRequest) Location(location V1ClinicalmatchPostRequestLocation) ApiV1ClinicalmatchPostRequest {
	r.location = &location
	return r
}

func (r ApiV1ClinicalmatchPostRequest) Execute() (*V1ClinicalmatchPost200Response, *http.Response, error) {
	return r.ApiService.V1ClinicalmatchPostExecute(r)
}

/*
V1ClinicalmatchPost Submit patient notes for clinical trial matching

Submits a file containing doctor notes about a patient, and a byte count for the file, to find relevant clinical trials.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ClinicalmatchPostRequest
*/
func (a *DefaultAPIService) V1ClinicalmatchPost(ctx context.Context) ApiV1ClinicalmatchPostRequest {
	return ApiV1ClinicalmatchPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ClinicalmatchPost200Response
func (a *DefaultAPIService) V1ClinicalmatchPostExecute(r ApiV1ClinicalmatchPostRequest) (*V1ClinicalmatchPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ClinicalmatchPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.V1ClinicalmatchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clinicalmatch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patientNotesFile == nil {
		return localVarReturnValue, nil, reportError("patientNotesFile is required and must be specified")
	}
	if r.fileSizeBytes == nil {
		return localVarReturnValue, nil, reportError("fileSizeBytes is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var patientNotesFileLocalVarFormFileName string
	var patientNotesFileLocalVarFileName     string
	var patientNotesFileLocalVarFileBytes    []byte

	patientNotesFileLocalVarFormFileName = "patientNotesFile"
	patientNotesFileLocalVarFile := r.patientNotesFile

	if patientNotesFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(patientNotesFileLocalVarFile)

		patientNotesFileLocalVarFileBytes = fbs
		patientNotesFileLocalVarFileName = patientNotesFileLocalVarFile.Name()
		patientNotesFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: patientNotesFileLocalVarFileBytes, fileName: patientNotesFileLocalVarFileName, formFileName: patientNotesFileLocalVarFormFileName})
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "fileSizeBytes", r.fileSizeBytes, "", "")
	if r.location != nil {
		paramJson, err := parameterToJson(*r.location)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("location", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1ClinicalmatchPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVersionGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
}

func (r ApiVersionGetRequest) Execute() (*VersionGet200Response, *http.Response, error) {
	return r.ApiService.VersionGetExecute(r)
}

/*
VersionGet Get API version

Returns the version of the Clinical Match API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVersionGetRequest
*/
func (a *DefaultAPIService) VersionGet(ctx context.Context) ApiVersionGetRequest {
	return ApiVersionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VersionGet200Response
func (a *DefaultAPIService) VersionGetExecute(r ApiVersionGetRequest) (*VersionGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VersionGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.VersionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/version"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
