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

// Appliance struct for Appliance
type Appliance struct {
}

// NewAppliance instantiates a new Appliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliance(subtype string) *Appliance {
	this := Appliance{}
	this.Subtype = subtype
	return &this
}

// NewApplianceWithDefaults instantiates a new Appliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceWithDefaults() *Appliance {
	this := Appliance{}
	return &this
}

func (o Appliance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableAppliance struct {
	value *Appliance
	isSet bool
}

func (v NullableAppliance) Get() *Appliance {
	return v.value
}

func (v *NullableAppliance) Set(val *Appliance) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliance) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliance(val *Appliance) *NullableAppliance {
	return &NullableAppliance{value: val, isSet: true}
}

func (v NullableAppliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
