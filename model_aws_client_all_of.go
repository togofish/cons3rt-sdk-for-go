/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AwsClientAllOf struct for AwsClientAllOf
type AwsClientAllOf struct {
	Region *string `json:"region,omitempty"`
}

// NewAwsClientAllOf instantiates a new AwsClientAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsClientAllOf() *AwsClientAllOf {
	this := AwsClientAllOf{}
	return &this
}

// NewAwsClientAllOfWithDefaults instantiates a new AwsClientAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsClientAllOfWithDefaults() *AwsClientAllOf {
	this := AwsClientAllOf{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AwsClientAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsClientAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AwsClientAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AwsClientAllOf) SetRegion(v string) {
	o.Region = &v
}

func (o AwsClientAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableAwsClientAllOf struct {
	value *AwsClientAllOf
	isSet bool
}

func (v NullableAwsClientAllOf) Get() *AwsClientAllOf {
	return v.value
}

func (v *NullableAwsClientAllOf) Set(val *AwsClientAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsClientAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsClientAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsClientAllOf(val *AwsClientAllOf) *NullableAwsClientAllOf {
	return &NullableAwsClientAllOf{value: val, isSet: true}
}

func (v NullableAwsClientAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsClientAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
