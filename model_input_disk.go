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

// InputDisk struct for InputDisk
type InputDisk struct {
	CapacityInMegabytes int32 `json:"capacityInMegabytes"`
	CreateOrder         int32 `json:"createOrder"`
}

// NewInputDisk instantiates a new InputDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputDisk(capacityInMegabytes int32, createOrder int32) *InputDisk {
	this := InputDisk{}
	this.CapacityInMegabytes = capacityInMegabytes
	this.CreateOrder = createOrder
	return &this
}

// NewInputDiskWithDefaults instantiates a new InputDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputDiskWithDefaults() *InputDisk {
	this := InputDisk{}
	return &this
}

// GetCapacityInMegabytes returns the CapacityInMegabytes field value
func (o *InputDisk) GetCapacityInMegabytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CapacityInMegabytes
}

// GetCapacityInMegabytesOk returns a tuple with the CapacityInMegabytes field value
// and a boolean to check if the value has been set.
func (o *InputDisk) GetCapacityInMegabytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityInMegabytes, true
}

// SetCapacityInMegabytes sets field value
func (o *InputDisk) SetCapacityInMegabytes(v int32) {
	o.CapacityInMegabytes = v
}

// GetCreateOrder returns the CreateOrder field value
func (o *InputDisk) GetCreateOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreateOrder
}

// GetCreateOrderOk returns a tuple with the CreateOrder field value
// and a boolean to check if the value has been set.
func (o *InputDisk) GetCreateOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateOrder, true
}

// SetCreateOrder sets field value
func (o *InputDisk) SetCreateOrder(v int32) {
	o.CreateOrder = v
}

func (o InputDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["capacityInMegabytes"] = o.CapacityInMegabytes
	}
	if true {
		toSerialize["createOrder"] = o.CreateOrder
	}
	return json.Marshal(toSerialize)
}

type NullableInputDisk struct {
	value *InputDisk
	isSet bool
}

func (v NullableInputDisk) Get() *InputDisk {
	return v.value
}

func (v *NullableInputDisk) Set(val *InputDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableInputDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableInputDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputDisk(val *InputDisk) *NullableInputDisk {
	return &NullableInputDisk{value: val, isSet: true}
}

func (v NullableInputDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
