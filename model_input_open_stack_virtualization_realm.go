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

// InputOpenStackVirtualizationRealm struct for InputOpenStackVirtualizationRealm
type InputOpenStackVirtualizationRealm struct {
	AccountId   string
	Cidr        string
	Description string
	Name        string
	Password    string
	Username    string
}

// NewInputOpenStackVirtualizationRealm instantiates a new InputOpenStackVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputOpenStackVirtualizationRealm(accountId string, cidr string, description string, name string, password string, username string) *InputOpenStackVirtualizationRealm {
	this := InputOpenStackVirtualizationRealm{}
	this.AccountId = accountId
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	this.Password = password
	this.Username = username
	return &this
}

// NewInputOpenStackVirtualizationRealmWithDefaults instantiates a new InputOpenStackVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputOpenStackVirtualizationRealmWithDefaults() *InputOpenStackVirtualizationRealm {
	this := InputOpenStackVirtualizationRealm{}
	return &this
}

func (o InputOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableInputOpenStackVirtualizationRealm struct {
	value *InputOpenStackVirtualizationRealm
	isSet bool
}

func (v NullableInputOpenStackVirtualizationRealm) Get() *InputOpenStackVirtualizationRealm {
	return v.value
}

func (v *NullableInputOpenStackVirtualizationRealm) Set(val *InputOpenStackVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableInputOpenStackVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableInputOpenStackVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputOpenStackVirtualizationRealm(val *InputOpenStackVirtualizationRealm) *NullableInputOpenStackVirtualizationRealm {
	return &NullableInputOpenStackVirtualizationRealm{value: val, isSet: true}
}

func (v NullableInputOpenStackVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputOpenStackVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
