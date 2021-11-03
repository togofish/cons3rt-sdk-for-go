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

// InputVirtualizationRealmBinding struct for InputVirtualizationRealmBinding
type InputVirtualizationRealmBinding struct {
	VirtualizationRealm InputMinimalVirtualizationRealm `json:"virtualizationRealm"`
	TemplateBindings []InputHostBinding `json:"templateBindings"`
}

// NewInputVirtualizationRealmBinding instantiates a new InputVirtualizationRealmBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVirtualizationRealmBinding(virtualizationRealm InputMinimalVirtualizationRealm, templateBindings []InputHostBinding) *InputVirtualizationRealmBinding {
	this := InputVirtualizationRealmBinding{}
	this.VirtualizationRealm = virtualizationRealm
	this.TemplateBindings = templateBindings
	return &this
}

// NewInputVirtualizationRealmBindingWithDefaults instantiates a new InputVirtualizationRealmBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVirtualizationRealmBindingWithDefaults() *InputVirtualizationRealmBinding {
	this := InputVirtualizationRealmBinding{}
	return &this
}

// GetVirtualizationRealm returns the VirtualizationRealm field value
func (o *InputVirtualizationRealmBinding) GetVirtualizationRealm() InputMinimalVirtualizationRealm {
	if o == nil {
		var ret InputMinimalVirtualizationRealm
		return ret
	}

	return o.VirtualizationRealm
}

// GetVirtualizationRealmOk returns a tuple with the VirtualizationRealm field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealmBinding) GetVirtualizationRealmOk() (*InputMinimalVirtualizationRealm, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualizationRealm, true
}

// SetVirtualizationRealm sets field value
func (o *InputVirtualizationRealmBinding) SetVirtualizationRealm(v InputMinimalVirtualizationRealm) {
	o.VirtualizationRealm = v
}

// GetTemplateBindings returns the TemplateBindings field value
func (o *InputVirtualizationRealmBinding) GetTemplateBindings() []InputHostBinding {
	if o == nil {
		var ret []InputHostBinding
		return ret
	}

	return o.TemplateBindings
}

// GetTemplateBindingsOk returns a tuple with the TemplateBindings field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealmBinding) GetTemplateBindingsOk() (*[]InputHostBinding, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateBindings, true
}

// SetTemplateBindings sets field value
func (o *InputVirtualizationRealmBinding) SetTemplateBindings(v []InputHostBinding) {
	o.TemplateBindings = v
}

func (o InputVirtualizationRealmBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["virtualizationRealm"] = o.VirtualizationRealm
	}
	if true {
		toSerialize["templateBindings"] = o.TemplateBindings
	}
	return json.Marshal(toSerialize)
}

type NullableInputVirtualizationRealmBinding struct {
	value *InputVirtualizationRealmBinding
	isSet bool
}

func (v NullableInputVirtualizationRealmBinding) Get() *InputVirtualizationRealmBinding {
	return v.value
}

func (v *NullableInputVirtualizationRealmBinding) Set(val *InputVirtualizationRealmBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVirtualizationRealmBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVirtualizationRealmBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVirtualizationRealmBinding(val *InputVirtualizationRealmBinding) *NullableInputVirtualizationRealmBinding {
	return &NullableInputVirtualizationRealmBinding{value: val, isSet: true}
}

func (v NullableInputVirtualizationRealmBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVirtualizationRealmBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


