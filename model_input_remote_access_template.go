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

// InputRemoteAccessTemplate struct for InputRemoteAccessTemplate
type InputRemoteAccessTemplate struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Port     int32   `json:"port"`
	Password *string `json:"password,omitempty"`
}

// NewInputRemoteAccessTemplate instantiates a new InputRemoteAccessTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputRemoteAccessTemplate(name string, type_ string, port int32) *InputRemoteAccessTemplate {
	this := InputRemoteAccessTemplate{}
	this.Name = name
	this.Type = type_
	this.Port = port
	return &this
}

// NewInputRemoteAccessTemplateWithDefaults instantiates a new InputRemoteAccessTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputRemoteAccessTemplateWithDefaults() *InputRemoteAccessTemplate {
	this := InputRemoteAccessTemplate{}
	return &this
}

// GetName returns the Name field value
func (o *InputRemoteAccessTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputRemoteAccessTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputRemoteAccessTemplate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *InputRemoteAccessTemplate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputRemoteAccessTemplate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputRemoteAccessTemplate) SetType(v string) {
	o.Type = v
}

// GetPort returns the Port field value
func (o *InputRemoteAccessTemplate) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *InputRemoteAccessTemplate) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *InputRemoteAccessTemplate) SetPort(v int32) {
	o.Port = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InputRemoteAccessTemplate) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputRemoteAccessTemplate) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InputRemoteAccessTemplate) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InputRemoteAccessTemplate) SetPassword(v string) {
	o.Password = &v
}

func (o InputRemoteAccessTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableInputRemoteAccessTemplate struct {
	value *InputRemoteAccessTemplate
	isSet bool
}

func (v NullableInputRemoteAccessTemplate) Get() *InputRemoteAccessTemplate {
	return v.value
}

func (v *NullableInputRemoteAccessTemplate) Set(val *InputRemoteAccessTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInputRemoteAccessTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInputRemoteAccessTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputRemoteAccessTemplate(val *InputRemoteAccessTemplate) *NullableInputRemoteAccessTemplate {
	return &NullableInputRemoteAccessTemplate{value: val, isSet: true}
}

func (v NullableInputRemoteAccessTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputRemoteAccessTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
