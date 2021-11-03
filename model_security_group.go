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

// SecurityGroup struct for SecurityGroup
type SecurityGroup struct {
	Identifier string `json:"identifier"`
}

// NewSecurityGroup instantiates a new SecurityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroup(identifier string) *SecurityGroup {
	this := SecurityGroup{}
	this.Identifier = identifier
	return &this
}

// NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupWithDefaults() *SecurityGroup {
	this := SecurityGroup{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *SecurityGroup) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *SecurityGroup) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *SecurityGroup) SetIdentifier(v string) {
	o.Identifier = v
}

func (o SecurityGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityGroup struct {
	value *SecurityGroup
	isSet bool
}

func (v NullableSecurityGroup) Get() *SecurityGroup {
	return v.value
}

func (v *NullableSecurityGroup) Set(val *SecurityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroup(val *SecurityGroup) *NullableSecurityGroup {
	return &NullableSecurityGroup{value: val, isSet: true}
}

func (v NullableSecurityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
