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

// InputDeviceAllOf struct for InputDeviceAllOf
type InputDeviceAllOf struct {
	Id *int32 `json:"id,omitempty"`
}

// NewInputDeviceAllOf instantiates a new InputDeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputDeviceAllOf() *InputDeviceAllOf {
	this := InputDeviceAllOf{}
	return &this
}

// NewInputDeviceAllOfWithDefaults instantiates a new InputDeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputDeviceAllOfWithDefaults() *InputDeviceAllOf {
	this := InputDeviceAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InputDeviceAllOf) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDeviceAllOf) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InputDeviceAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InputDeviceAllOf) SetId(v int32) {
	o.Id = &v
}

func (o InputDeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInputDeviceAllOf struct {
	value *InputDeviceAllOf
	isSet bool
}

func (v NullableInputDeviceAllOf) Get() *InputDeviceAllOf {
	return v.value
}

func (v *NullableInputDeviceAllOf) Set(val *InputDeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInputDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInputDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputDeviceAllOf(val *InputDeviceAllOf) *NullableInputDeviceAllOf {
	return &NullableInputDeviceAllOf{value: val, isSet: true}
}

func (v NullableInputDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
