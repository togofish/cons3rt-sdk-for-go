/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InputPhysicalHostFullAllOf struct for InputPhysicalHostFullAllOf
type InputPhysicalHostFullAllOf struct {
	Id              *int32                    `json:"id,omitempty"`
	Name            *string                   `json:"name,omitempty"`
	PhysicalMachine *InputPhysicalMachine     `json:"physicalMachine,omitempty"`
	Components      *[]InputAbstractComponent `json:"components,omitempty"`
}

// NewInputPhysicalHostFullAllOf instantiates a new InputPhysicalHostFullAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputPhysicalHostFullAllOf() *InputPhysicalHostFullAllOf {
	this := InputPhysicalHostFullAllOf{}
	return &this
}

// NewInputPhysicalHostFullAllOfWithDefaults instantiates a new InputPhysicalHostFullAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputPhysicalHostFullAllOfWithDefaults() *InputPhysicalHostFullAllOf {
	this := InputPhysicalHostFullAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InputPhysicalHostFullAllOf) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostFullAllOf) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InputPhysicalHostFullAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InputPhysicalHostFullAllOf) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputPhysicalHostFullAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostFullAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputPhysicalHostFullAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputPhysicalHostFullAllOf) SetName(v string) {
	o.Name = &v
}

// GetPhysicalMachine returns the PhysicalMachine field value if set, zero value otherwise.
func (o *InputPhysicalHostFullAllOf) GetPhysicalMachine() InputPhysicalMachine {
	if o == nil || o.PhysicalMachine == nil {
		var ret InputPhysicalMachine
		return ret
	}
	return *o.PhysicalMachine
}

// GetPhysicalMachineOk returns a tuple with the PhysicalMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostFullAllOf) GetPhysicalMachineOk() (*InputPhysicalMachine, bool) {
	if o == nil || o.PhysicalMachine == nil {
		return nil, false
	}
	return o.PhysicalMachine, true
}

// HasPhysicalMachine returns a boolean if a field has been set.
func (o *InputPhysicalHostFullAllOf) HasPhysicalMachine() bool {
	if o != nil && o.PhysicalMachine != nil {
		return true
	}

	return false
}

// SetPhysicalMachine gets a reference to the given InputPhysicalMachine and assigns it to the PhysicalMachine field.
func (o *InputPhysicalHostFullAllOf) SetPhysicalMachine(v InputPhysicalMachine) {
	o.PhysicalMachine = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *InputPhysicalHostFullAllOf) GetComponents() []InputAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []InputAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPhysicalHostFullAllOf) GetComponentsOk() (*[]InputAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *InputPhysicalHostFullAllOf) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []InputAbstractComponent and assigns it to the Components field.
func (o *InputPhysicalHostFullAllOf) SetComponents(v []InputAbstractComponent) {
	o.Components = &v
}

func (o InputPhysicalHostFullAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableInputPhysicalHostFullAllOf struct {
	value *InputPhysicalHostFullAllOf
	isSet bool
}

func (v NullableInputPhysicalHostFullAllOf) Get() *InputPhysicalHostFullAllOf {
	return v.value
}

func (v *NullableInputPhysicalHostFullAllOf) Set(val *InputPhysicalHostFullAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInputPhysicalHostFullAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInputPhysicalHostFullAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputPhysicalHostFullAllOf(val *InputPhysicalHostFullAllOf) *NullableInputPhysicalHostFullAllOf {
	return &NullableInputPhysicalHostFullAllOf{value: val, isSet: true}
}

func (v NullableInputPhysicalHostFullAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputPhysicalHostFullAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
