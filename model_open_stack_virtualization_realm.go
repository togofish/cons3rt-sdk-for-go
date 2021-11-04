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

// OpenStackVirtualizationRealm struct for OpenStackVirtualizationRealm
type OpenStackVirtualizationRealm struct {
}

// NewOpenStackVirtualizationRealm instantiates a new OpenStackVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackVirtualizationRealm(accountId string, cidr string, description string, name string, password string, username string) *OpenStackVirtualizationRealm {
	this := OpenStackVirtualizationRealm{}
	this.AccountId = accountId
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	this.Password = password
	this.Username = username
	return &this
}

// NewOpenStackVirtualizationRealmWithDefaults instantiates a new OpenStackVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackVirtualizationRealmWithDefaults() *OpenStackVirtualizationRealm {
	this := OpenStackVirtualizationRealm{}
	return &this
}

func (o OpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableOpenStackVirtualizationRealm struct {
	value *OpenStackVirtualizationRealm
	isSet bool
}

func (v NullableOpenStackVirtualizationRealm) Get() *OpenStackVirtualizationRealm {
	return v.value
}

func (v *NullableOpenStackVirtualizationRealm) Set(val *OpenStackVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackVirtualizationRealm(val *OpenStackVirtualizationRealm) *NullableOpenStackVirtualizationRealm {
	return &NullableOpenStackVirtualizationRealm{value: val, isSet: true}
}

func (v NullableOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


