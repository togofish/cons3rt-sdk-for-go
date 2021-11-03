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

// BasicPhysicalHost struct for BasicPhysicalHost
type BasicPhysicalHost struct {
}

// NewBasicPhysicalHost instantiates a new BasicPhysicalHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicPhysicalHost(subtype string) *BasicPhysicalHost {
	this := BasicPhysicalHost{}
	this.Subtype = subtype
	return &this
}

// NewBasicPhysicalHostWithDefaults instantiates a new BasicPhysicalHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicPhysicalHostWithDefaults() *BasicPhysicalHost {
	this := BasicPhysicalHost{}
	return &this
}

func (o BasicPhysicalHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableBasicPhysicalHost struct {
	value *BasicPhysicalHost
	isSet bool
}

func (v NullableBasicPhysicalHost) Get() *BasicPhysicalHost {
	return v.value
}

func (v *NullableBasicPhysicalHost) Set(val *BasicPhysicalHost) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicPhysicalHost) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicPhysicalHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicPhysicalHost(val *BasicPhysicalHost) *NullableBasicPhysicalHost {
	return &NullableBasicPhysicalHost{value: val, isSet: true}
}

func (v NullableBasicPhysicalHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicPhysicalHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
