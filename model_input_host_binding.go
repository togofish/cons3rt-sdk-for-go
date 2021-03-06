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

// InputHostBinding struct for InputHostBinding
type InputHostBinding struct {
	SystemRole   string `json:"systemRole"`
	TemplateName string `json:"templateName"`
	// Required for Azure and AWS EC2 Cloudspaces
	InstanceType *string `json:"instanceType,omitempty"`
}

// NewInputHostBinding instantiates a new InputHostBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputHostBinding(systemRole string, templateName string) *InputHostBinding {
	this := InputHostBinding{}
	this.SystemRole = systemRole
	this.TemplateName = templateName
	return &this
}

// NewInputHostBindingWithDefaults instantiates a new InputHostBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputHostBindingWithDefaults() *InputHostBinding {
	this := InputHostBinding{}
	return &this
}

// GetSystemRole returns the SystemRole field value
func (o *InputHostBinding) GetSystemRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value
// and a boolean to check if the value has been set.
func (o *InputHostBinding) GetSystemRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemRole, true
}

// SetSystemRole sets field value
func (o *InputHostBinding) SetSystemRole(v string) {
	o.SystemRole = v
}

// GetTemplateName returns the TemplateName field value
func (o *InputHostBinding) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *InputHostBinding) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *InputHostBinding) SetTemplateName(v string) {
	o.TemplateName = v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *InputHostBinding) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputHostBinding) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *InputHostBinding) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *InputHostBinding) SetInstanceType(v string) {
	o.InstanceType = &v
}

func (o InputHostBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["systemRole"] = o.SystemRole
	}
	if true {
		toSerialize["templateName"] = o.TemplateName
	}
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}
	return json.Marshal(toSerialize)
}

type NullableInputHostBinding struct {
	value *InputHostBinding
	isSet bool
}

func (v NullableInputHostBinding) Get() *InputHostBinding {
	return v.value
}

func (v *NullableInputHostBinding) Set(val *InputHostBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableInputHostBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableInputHostBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputHostBinding(val *InputHostBinding) *NullableInputHostBinding {
	return &NullableInputHostBinding{value: val, isSet: true}
}

func (v NullableInputHostBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputHostBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
