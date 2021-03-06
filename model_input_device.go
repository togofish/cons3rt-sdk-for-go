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

// InputDevice struct for InputDevice
type InputDevice struct {
	Id      int32 `json:"id"`
	Subtype string
}

// NewInputDevice instantiates a new InputDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputDevice(id int32, subtype string) *InputDevice {
	this := InputDevice{}
	this.Subtype = subtype
	this.Id = id
	return &this
}

// NewInputDeviceWithDefaults instantiates a new InputDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputDeviceWithDefaults() *InputDevice {
	this := InputDevice{}
	return &this
}

// GetId returns the Id field value
func (o *InputDevice) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InputDevice) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InputDevice) SetId(v int32) {
	o.Id = v
}

func (o InputDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInputDevice struct {
	value *InputDevice
	isSet bool
}

func (v NullableInputDevice) Get() *InputDevice {
	return v.value
}

func (v *NullableInputDevice) Set(val *InputDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableInputDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableInputDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputDevice(val *InputDevice) *NullableInputDevice {
	return &NullableInputDevice{value: val, isSet: true}
}

func (v NullableInputDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
