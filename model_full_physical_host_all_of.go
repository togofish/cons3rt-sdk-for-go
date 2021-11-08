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

// FullPhysicalHostAllOf struct for FullPhysicalHostAllOf
type FullPhysicalHostAllOf struct {
	Components      *[]MinimalAbstractComponent `json:"components,omitempty"`
	PhysicalMachine *GeneralPhysicalMachine     `json:"physicalMachine,omitempty"`
}

// NewFullPhysicalHostAllOf instantiates a new FullPhysicalHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullPhysicalHostAllOf() *FullPhysicalHostAllOf {
	this := FullPhysicalHostAllOf{}
	return &this
}

// NewFullPhysicalHostAllOfWithDefaults instantiates a new FullPhysicalHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullPhysicalHostAllOfWithDefaults() *FullPhysicalHostAllOf {
	this := FullPhysicalHostAllOf{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *FullPhysicalHostAllOf) GetComponents() []MinimalAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []MinimalAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalHostAllOf) GetComponentsOk() (*[]MinimalAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *FullPhysicalHostAllOf) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []MinimalAbstractComponent and assigns it to the Components field.
func (o *FullPhysicalHostAllOf) SetComponents(v []MinimalAbstractComponent) {
	o.Components = &v
}

// GetPhysicalMachine returns the PhysicalMachine field value if set, zero value otherwise.
func (o *FullPhysicalHostAllOf) GetPhysicalMachine() GeneralPhysicalMachine {
	if o == nil || o.PhysicalMachine == nil {
		var ret GeneralPhysicalMachine
		return ret
	}
	return *o.PhysicalMachine
}

// GetPhysicalMachineOk returns a tuple with the PhysicalMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalHostAllOf) GetPhysicalMachineOk() (*GeneralPhysicalMachine, bool) {
	if o == nil || o.PhysicalMachine == nil {
		return nil, false
	}
	return o.PhysicalMachine, true
}

// HasPhysicalMachine returns a boolean if a field has been set.
func (o *FullPhysicalHostAllOf) HasPhysicalMachine() bool {
	if o != nil && o.PhysicalMachine != nil {
		return true
	}

	return false
}

// SetPhysicalMachine gets a reference to the given GeneralPhysicalMachine and assigns it to the PhysicalMachine field.
func (o *FullPhysicalHostAllOf) SetPhysicalMachine(v GeneralPhysicalMachine) {
	o.PhysicalMachine = &v
}

func (o FullPhysicalHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.PhysicalMachine != nil {
		toSerialize["physicalMachine"] = o.PhysicalMachine
	}
	return json.Marshal(toSerialize)
}

type NullableFullPhysicalHostAllOf struct {
	value *FullPhysicalHostAllOf
	isSet bool
}

func (v NullableFullPhysicalHostAllOf) Get() *FullPhysicalHostAllOf {
	return v.value
}

func (v *NullableFullPhysicalHostAllOf) Set(val *FullPhysicalHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullPhysicalHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullPhysicalHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullPhysicalHostAllOf(val *FullPhysicalHostAllOf) *NullableFullPhysicalHostAllOf {
	return &NullableFullPhysicalHostAllOf{value: val, isSet: true}
}

func (v NullableFullPhysicalHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullPhysicalHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
