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

// BasicDeployment struct for BasicDeployment
type BasicDeployment struct {
}

// NewBasicDeployment instantiates a new BasicDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicDeployment(subtype string) *BasicDeployment {
	this := BasicDeployment{}
	this.Subtype = subtype
	return &this
}

// NewBasicDeploymentWithDefaults instantiates a new BasicDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicDeploymentWithDefaults() *BasicDeployment {
	this := BasicDeployment{}
	return &this
}

func (o BasicDeployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableBasicDeployment struct {
	value *BasicDeployment
	isSet bool
}

func (v NullableBasicDeployment) Get() *BasicDeployment {
	return v.value
}

func (v *NullableBasicDeployment) Set(val *BasicDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicDeployment(val *BasicDeployment) *NullableBasicDeployment {
	return &NullableBasicDeployment{value: val, isSet: true}
}

func (v NullableBasicDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
