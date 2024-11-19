# V1ClinicalmatchPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Matches** | Pointer to [**[]V1ClinicalmatchPost200ResponseMatchesInner**](V1ClinicalmatchPost200ResponseMatchesInner.md) |  | [optional] 

## Methods

### NewV1ClinicalmatchPost200Response

`func NewV1ClinicalmatchPost200Response() *V1ClinicalmatchPost200Response`

NewV1ClinicalmatchPost200Response instantiates a new V1ClinicalmatchPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClinicalmatchPost200ResponseWithDefaults

`func NewV1ClinicalmatchPost200ResponseWithDefaults() *V1ClinicalmatchPost200Response`

NewV1ClinicalmatchPost200ResponseWithDefaults instantiates a new V1ClinicalmatchPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V1ClinicalmatchPost200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ClinicalmatchPost200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ClinicalmatchPost200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ClinicalmatchPost200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMatches

`func (o *V1ClinicalmatchPost200Response) GetMatches() []V1ClinicalmatchPost200ResponseMatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *V1ClinicalmatchPost200Response) GetMatchesOk() (*[]V1ClinicalmatchPost200ResponseMatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *V1ClinicalmatchPost200Response) SetMatches(v []V1ClinicalmatchPost200ResponseMatchesInner)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *V1ClinicalmatchPost200Response) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


