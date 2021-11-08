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

// FullOpenStackVirtualizationRealm struct for FullOpenStackVirtualizationRealm
type FullOpenStackVirtualizationRealm struct {
	Name        string
	AccountId   string
	Networks    []Network
	Description string
	Username    string
}

// NewFullOpenStackVirtualizationRealm instantiates a new FullOpenStackVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullOpenStackVirtualizationRealm(name string, accountId string, networks []Network, description string, username string) *FullOpenStackVirtualizationRealm {
	this := FullOpenStackVirtualizationRealm{}
	this.Name = name
	this.AccountId = accountId
	this.Networks = networks
	this.Description = description
	this.Username = username
	return &this
}

// NewFullOpenStackVirtualizationRealmWithDefaults instantiates a new FullOpenStackVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullOpenStackVirtualizationRealmWithDefaults() *FullOpenStackVirtualizationRealm {
	this := FullOpenStackVirtualizationRealm{}
	return &this
}

func (o FullOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableFullOpenStackVirtualizationRealm struct {
	value *FullOpenStackVirtualizationRealm
	isSet bool
}

func (v NullableFullOpenStackVirtualizationRealm) Get() *FullOpenStackVirtualizationRealm {
	return v.value
}

func (v *NullableFullOpenStackVirtualizationRealm) Set(val *FullOpenStackVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableFullOpenStackVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableFullOpenStackVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullOpenStackVirtualizationRealm(val *FullOpenStackVirtualizationRealm) *NullableFullOpenStackVirtualizationRealm {
	return &NullableFullOpenStackVirtualizationRealm{value: val, isSet: true}
}

func (v NullableFullOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullOpenStackVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
