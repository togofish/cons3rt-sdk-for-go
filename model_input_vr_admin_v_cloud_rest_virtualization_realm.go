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

// InputVRAdminVCloudRestVirtualizationRealm struct for InputVRAdminVCloudRestVirtualizationRealm
type InputVRAdminVCloudRestVirtualizationRealm struct {
	VirtualizationRealmType string
	Cidr                    string
	Description             string
	Name                    string
}

// NewInputVRAdminVCloudRestVirtualizationRealm instantiates a new InputVRAdminVCloudRestVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVRAdminVCloudRestVirtualizationRealm(virtualizationRealmType string, cidr string, description string, name string) *InputVRAdminVCloudRestVirtualizationRealm {
	this := InputVRAdminVCloudRestVirtualizationRealm{}
	this.VirtualizationRealmType = virtualizationRealmType
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	return &this
}

// NewInputVRAdminVCloudRestVirtualizationRealmWithDefaults instantiates a new InputVRAdminVCloudRestVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVRAdminVCloudRestVirtualizationRealmWithDefaults() *InputVRAdminVCloudRestVirtualizationRealm {
	this := InputVRAdminVCloudRestVirtualizationRealm{}
	return &this
}

func (o InputVRAdminVCloudRestVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableInputVRAdminVCloudRestVirtualizationRealm struct {
	value *InputVRAdminVCloudRestVirtualizationRealm
	isSet bool
}

func (v NullableInputVRAdminVCloudRestVirtualizationRealm) Get() *InputVRAdminVCloudRestVirtualizationRealm {
	return v.value
}

func (v *NullableInputVRAdminVCloudRestVirtualizationRealm) Set(val *InputVRAdminVCloudRestVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVRAdminVCloudRestVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVRAdminVCloudRestVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVRAdminVCloudRestVirtualizationRealm(val *InputVRAdminVCloudRestVirtualizationRealm) *NullableInputVRAdminVCloudRestVirtualizationRealm {
	return &NullableInputVRAdminVCloudRestVirtualizationRealm{value: val, isSet: true}
}

func (v NullableInputVRAdminVCloudRestVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVRAdminVCloudRestVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
