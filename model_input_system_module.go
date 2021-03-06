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

// InputSystemModule struct for InputSystemModule
type InputSystemModule struct {
	Subtype string `json:"subtype"`
}

// NewInputSystemModule instantiates a new InputSystemModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputSystemModule(subtype string) *InputSystemModule {
	this := InputSystemModule{}
	this.Subtype = subtype
	return &this
}

// NewInputSystemModuleWithDefaults instantiates a new InputSystemModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputSystemModuleWithDefaults() *InputSystemModule {
	this := InputSystemModule{}
	return &this
}

// GetSubtype returns the Subtype field value
func (o *InputSystemModule) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *InputSystemModule) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *InputSystemModule) SetSubtype(v string) {
	o.Subtype = v
}

func (o InputSystemModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableInputSystemModule struct {
	value *InputSystemModule
	isSet bool
}

func (v NullableInputSystemModule) Get() *InputSystemModule {
	return v.value
}

func (v *NullableInputSystemModule) Set(val *InputSystemModule) {
	v.value = val
	v.isSet = true
}

func (v NullableInputSystemModule) IsSet() bool {
	return v.isSet
}

func (v *NullableInputSystemModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputSystemModule(val *InputSystemModule) *NullableInputSystemModule {
	return &NullableInputSystemModule{value: val, isSet: true}
}

func (v NullableInputSystemModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputSystemModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
