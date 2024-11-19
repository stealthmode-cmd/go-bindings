# ClinicalTrialMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | **string** | The unique identifier of the clinical trial | 
**TrialName** | **string** | The name of the clinical trial | 
**Locations** | [**[]Location**](Location.md) |  | 
**Summary** | **string** | Detailed description of the clinical trial | 
**Status** | Pointer to **string** | Status about the clinical trial | [optional] 

## Methods

### NewClinicalTrialMatch

`func NewClinicalTrialMatch(trialId string, trialName string, locations []Location, summary string, ) *ClinicalTrialMatch`

NewClinicalTrialMatch instantiates a new ClinicalTrialMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClinicalTrialMatchWithDefaults

`func NewClinicalTrialMatchWithDefaults() *ClinicalTrialMatch`

NewClinicalTrialMatchWithDefaults instantiates a new ClinicalTrialMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *ClinicalTrialMatch) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *ClinicalTrialMatch) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *ClinicalTrialMatch) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.


### GetTrialName

`func (o *ClinicalTrialMatch) GetTrialName() string`

GetTrialName returns the TrialName field if non-nil, zero value otherwise.

### GetTrialNameOk

`func (o *ClinicalTrialMatch) GetTrialNameOk() (*string, bool)`

GetTrialNameOk returns a tuple with the TrialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialName

`func (o *ClinicalTrialMatch) SetTrialName(v string)`

SetTrialName sets TrialName field to given value.


### GetLocations

`func (o *ClinicalTrialMatch) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ClinicalTrialMatch) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ClinicalTrialMatch) SetLocations(v []Location)`

SetLocations sets Locations field to given value.


### GetSummary

`func (o *ClinicalTrialMatch) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ClinicalTrialMatch) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ClinicalTrialMatch) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetStatus

`func (o *ClinicalTrialMatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClinicalTrialMatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClinicalTrialMatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClinicalTrialMatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


