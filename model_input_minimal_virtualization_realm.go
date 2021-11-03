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

// InputMinimalVirtualizationRealm struct for InputMinimalVirtualizationRealm
type InputMinimalVirtualizationRealm struct {
	Id int32 `json:"id"`
}

// NewInputMinimalVirtualizationRealm instantiates a new InputMinimalVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputMinimalVirtualizationRealm(id int32) *InputMinimalVirtualizationRealm {
	this := InputMinimalVirtualizationRealm{}
	this.Id = id
	return &this
}

// NewInputMinimalVirtualizationRealmWithDefaults instantiates a new InputMinimalVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputMinimalVirtualizationRealmWithDefaults() *InputMinimalVirtualizationRealm {
	this := InputMinimalVirtualizationRealm{}
	return &this
}

// GetId returns the Id field value
func (o *InputMinimalVirtualizationRealm) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InputMinimalVirtualizationRealm) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InputMinimalVirtualizationRealm) SetId(v int32) {
	o.Id = v
}

func (o InputMinimalVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInputMinimalVirtualizationRealm struct {
	value *InputMinimalVirtualizationRealm
	isSet bool
}

func (v NullableInputMinimalVirtualizationRealm) Get() *InputMinimalVirtualizationRealm {
	return v.value
}

func (v *NullableInputMinimalVirtualizationRealm) Set(val *InputMinimalVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableInputMinimalVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableInputMinimalVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputMinimalVirtualizationRealm(val *InputMinimalVirtualizationRealm) *NullableInputMinimalVirtualizationRealm {
	return &NullableInputMinimalVirtualizationRealm{value: val, isSet: true}
}

func (v NullableInputMinimalVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputMinimalVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
