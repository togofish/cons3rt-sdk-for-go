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

// BasicPhysicalMachineAllOf struct for BasicPhysicalMachineAllOf
type BasicPhysicalMachineAllOf struct {
	Architecture    *string `json:"architecture,omitempty"`
	Bits            *string `json:"bits,omitempty"`
	OperatingSystem *string `json:"operatingSystem,omitempty"`
}

// NewBasicPhysicalMachineAllOf instantiates a new BasicPhysicalMachineAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicPhysicalMachineAllOf() *BasicPhysicalMachineAllOf {
	this := BasicPhysicalMachineAllOf{}
	return &this
}

// NewBasicPhysicalMachineAllOfWithDefaults instantiates a new BasicPhysicalMachineAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicPhysicalMachineAllOfWithDefaults() *BasicPhysicalMachineAllOf {
	this := BasicPhysicalMachineAllOf{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *BasicPhysicalMachineAllOf) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicPhysicalMachineAllOf) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *BasicPhysicalMachineAllOf) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *BasicPhysicalMachineAllOf) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *BasicPhysicalMachineAllOf) GetBits() string {
	if o == nil || o.Bits == nil {
		var ret string
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicPhysicalMachineAllOf) GetBitsOk() (*string, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *BasicPhysicalMachineAllOf) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given string and assigns it to the Bits field.
func (o *BasicPhysicalMachineAllOf) SetBits(v string) {
	o.Bits = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *BasicPhysicalMachineAllOf) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicPhysicalMachineAllOf) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *BasicPhysicalMachineAllOf) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *BasicPhysicalMachineAllOf) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

func (o BasicPhysicalMachineAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Architecture != nil {
		toSerialize["architecture"] = o.Architecture
	}
	if o.Bits != nil {
		toSerialize["bits"] = o.Bits
	}
	if o.OperatingSystem != nil {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	return json.Marshal(toSerialize)
}

type NullableBasicPhysicalMachineAllOf struct {
	value *BasicPhysicalMachineAllOf
	isSet bool
}

func (v NullableBasicPhysicalMachineAllOf) Get() *BasicPhysicalMachineAllOf {
	return v.value
}

func (v *NullableBasicPhysicalMachineAllOf) Set(val *BasicPhysicalMachineAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicPhysicalMachineAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicPhysicalMachineAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicPhysicalMachineAllOf(val *BasicPhysicalMachineAllOf) *NullableBasicPhysicalMachineAllOf {
	return &NullableBasicPhysicalMachineAllOf{value: val, isSet: true}
}

func (v NullableBasicPhysicalMachineAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicPhysicalMachineAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
