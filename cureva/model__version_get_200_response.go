/*
Clinical Match API

A simple API to match patients to clinical trials.

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VersionGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionGet200Response{}

// VersionGet200Response struct for VersionGet200Response
type VersionGet200Response struct {
	Version *string `json:"version,omitempty"`
}

// NewVersionGet200Response instantiates a new VersionGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionGet200Response() *VersionGet200Response {
	this := VersionGet200Response{}
	return &this
}

// NewVersionGet200ResponseWithDefaults instantiates a new VersionGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionGet200ResponseWithDefaults() *VersionGet200Response {
	this := VersionGet200Response{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VersionGet200Response) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionGet200Response) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionGet200Response) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VersionGet200Response) SetVersion(v string) {
	o.Version = &v
}

func (o VersionGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableVersionGet200Response struct {
	value *VersionGet200Response
	isSet bool
}

func (v NullableVersionGet200Response) Get() *VersionGet200Response {
	return v.value
}

func (v *NullableVersionGet200Response) Set(val *VersionGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionGet200Response(val *VersionGet200Response) *NullableVersionGet200Response {
	return &NullableVersionGet200Response{value: val, isSet: true}
}

func (v NullableVersionGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


