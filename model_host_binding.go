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

// HostBinding struct for HostBinding
type HostBinding struct {
	VirtualizationRealmTemplates []TemplateBinding `json:"virtualizationRealmTemplates"`
	SystemRole                   string            `json:"systemRole"`
}

// NewHostBinding instantiates a new HostBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostBinding(virtualizationRealmTemplates []TemplateBinding, systemRole string) *HostBinding {
	this := HostBinding{}
	this.VirtualizationRealmTemplates = virtualizationRealmTemplates
	this.SystemRole = systemRole
	return &this
}

// NewHostBindingWithDefaults instantiates a new HostBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostBindingWithDefaults() *HostBinding {
	this := HostBinding{}
	return &this
}

// GetVirtualizationRealmTemplates returns the VirtualizationRealmTemplates field value
func (o *HostBinding) GetVirtualizationRealmTemplates() []TemplateBinding {
	if o == nil {
		var ret []TemplateBinding
		return ret
	}

	return o.VirtualizationRealmTemplates
}

// GetVirtualizationRealmTemplatesOk returns a tuple with the VirtualizationRealmTemplates field value
// and a boolean to check if the value has been set.
func (o *HostBinding) GetVirtualizationRealmTemplatesOk() (*[]TemplateBinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualizationRealmTemplates, true
}

// SetVirtualizationRealmTemplates sets field value
func (o *HostBinding) SetVirtualizationRealmTemplates(v []TemplateBinding) {
	o.VirtualizationRealmTemplates = v
}

// GetSystemRole returns the SystemRole field value
func (o *HostBinding) GetSystemRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value
// and a boolean to check if the value has been set.
func (o *HostBinding) GetSystemRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemRole, true
}

// SetSystemRole sets field value
func (o *HostBinding) SetSystemRole(v string) {
	o.SystemRole = v
}

func (o HostBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["virtualizationRealmTemplates"] = o.VirtualizationRealmTemplates
	}
	if true {
		toSerialize["systemRole"] = o.SystemRole
	}
	return json.Marshal(toSerialize)
}

type NullableHostBinding struct {
	value *HostBinding
	isSet bool
}

func (v NullableHostBinding) Get() *HostBinding {
	return v.value
}

func (v *NullableHostBinding) Set(val *HostBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableHostBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableHostBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostBinding(val *HostBinding) *NullableHostBinding {
	return &NullableHostBinding{value: val, isSet: true}
}

func (v NullableHostBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
