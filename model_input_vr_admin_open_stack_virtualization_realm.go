/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gocons3rt

import (
	"encoding/json"
)

// InputVRAdminOpenStackVirtualizationRealm struct for InputVRAdminOpenStackVirtualizationRealm
type InputVRAdminOpenStackVirtualizationRealm struct {
}

// NewInputVRAdminOpenStackVirtualizationRealm instantiates a new InputVRAdminOpenStackVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVRAdminOpenStackVirtualizationRealm(virtualizationRealmType string, cidr string, description string, name string) *InputVRAdminOpenStackVirtualizationRealm {
	this := InputVRAdminOpenStackVirtualizationRealm{}
	this.VirtualizationRealmType = virtualizationRealmType
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	return &this
}

// NewInputVRAdminOpenStackVirtualizationRealmWithDefaults instantiates a new InputVRAdminOpenStackVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVRAdminOpenStackVirtualizationRealmWithDefaults() *InputVRAdminOpenStackVirtualizationRealm {
	this := InputVRAdminOpenStackVirtualizationRealm{}
	return &this
}

func (o InputVRAdminOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableInputVRAdminOpenStackVirtualizationRealm struct {
	value *InputVRAdminOpenStackVirtualizationRealm
	isSet bool
}

func (v NullableInputVRAdminOpenStackVirtualizationRealm) Get() *InputVRAdminOpenStackVirtualizationRealm {
	return v.value
}

func (v *NullableInputVRAdminOpenStackVirtualizationRealm) Set(val *InputVRAdminOpenStackVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVRAdminOpenStackVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVRAdminOpenStackVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVRAdminOpenStackVirtualizationRealm(val *InputVRAdminOpenStackVirtualizationRealm) *NullableInputVRAdminOpenStackVirtualizationRealm {
	return &NullableInputVRAdminOpenStackVirtualizationRealm{value: val, isSet: true}
}

func (v NullableInputVRAdminOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVRAdminOpenStackVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


