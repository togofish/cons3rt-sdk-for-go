/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cons3rt

import (
	"encoding/json"
)

// MinimalAppliance struct for MinimalAppliance
type MinimalAppliance struct {
	Subtype string
}

// NewMinimalAppliance instantiates a new MinimalAppliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalAppliance(subtype string) *MinimalAppliance {
	this := MinimalAppliance{}
	this.Subtype = subtype
	return &this
}

// NewMinimalApplianceWithDefaults instantiates a new MinimalAppliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalApplianceWithDefaults() *MinimalAppliance {
	this := MinimalAppliance{}
	return &this
}

func (o MinimalAppliance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableMinimalAppliance struct {
	value *MinimalAppliance
	isSet bool
}

func (v NullableMinimalAppliance) Get() *MinimalAppliance {
	return v.value
}

func (v *NullableMinimalAppliance) Set(val *MinimalAppliance) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalAppliance) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalAppliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalAppliance(val *MinimalAppliance) *NullableMinimalAppliance {
	return &NullableMinimalAppliance{value: val, isSet: true}
}

func (v NullableMinimalAppliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalAppliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
