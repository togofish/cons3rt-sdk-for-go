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

// InputPhysicalHostAllOf struct for InputPhysicalHostAllOf
type InputPhysicalHostAllOf struct {
	Name            *string                   `json:"name,omitempty"`
	PhysicalMachine *InputPhysicalMachine     `json:"physicalMachine,omitempty"`
	Components      *[]InputAbstractComponent `json:"components,omitempty"`
}

// NewInputPhysicalHostAllOf instantiates a new InputPhysicalHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputPhysicalHostAllOf() *InputPhysicalHostAllOf {
	this := InputPhysicalHostAllOf{}
	return &this
}

// NewInputPhysicalHostAllOfWithDefaults instantiates a new InputPhysicalHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputPhysicalHostAllOfWithDefaults() *InputPhysicalHostAllOf {
	this := InputPhysicalHostAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputPhysicalHostAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputPhysicalHostAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputPhysicalHostAllOf) SetName(v string) {
	o.Name = &v
}

// GetPhysicalMachine returns the PhysicalMachine field value if set, zero value otherwise.
func (o *InputPhysicalHostAllOf) GetPhysicalMachine() InputPhysicalMachine {
	if o == nil || o.PhysicalMachine == nil {
		var ret InputPhysicalMachine
		return ret
	}
	return *o.PhysicalMachine
}

// GetPhysicalMachineOk returns a tuple with the PhysicalMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostAllOf) GetPhysicalMachineOk() (*InputPhysicalMachine, bool) {
	if o == nil || o.PhysicalMachine == nil {
		return nil, false
	}
	return o.PhysicalMachine, true
}

// HasPhysicalMachine returns a boolean if a field has been set.
func (o *InputPhysicalHostAllOf) HasPhysicalMachine() bool {
	if o != nil && o.PhysicalMachine != nil {
		return true
	}

	return false
}

// SetPhysicalMachine gets a reference to the given InputPhysicalMachine and assigns it to the PhysicalMachine field.
func (o *InputPhysicalHostAllOf) SetPhysicalMachine(v InputPhysicalMachine) {
	o.PhysicalMachine = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *InputPhysicalHostAllOf) GetComponents() []InputAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []InputAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostAllOf) GetComponentsOk() (*[]InputAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *InputPhysicalHostAllOf) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []InputAbstractComponent and assigns it to the Components field.
func (o *InputPhysicalHostAllOf) SetComponents(v []InputAbstractComponent) {
	o.Components = &v
}

func (o InputPhysicalHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PhysicalMachine != nil {
		toSerialize["physicalMachine"] = o.PhysicalMachine
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableInputPhysicalHostAllOf struct {
	value *InputPhysicalHostAllOf
	isSet bool
}

func (v NullableInputPhysicalHostAllOf) Get() *InputPhysicalHostAllOf {
	return v.value
}

func (v *NullableInputPhysicalHostAllOf) Set(val *InputPhysicalHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInputPhysicalHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInputPhysicalHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputPhysicalHostAllOf(val *InputPhysicalHostAllOf) *NullableInputPhysicalHostAllOf {
	return &NullableInputPhysicalHostAllOf{value: val, isSet: true}
}

func (v NullableInputPhysicalHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputPhysicalHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
