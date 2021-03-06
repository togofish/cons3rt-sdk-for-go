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

// InputMetadata struct for InputMetadata
type InputMetadata struct {
	Properties []InputProperty `json:"properties"`
}

// NewInputMetadata instantiates a new InputMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputMetadata(properties []InputProperty) *InputMetadata {
	this := InputMetadata{}
	this.Properties = properties
	return &this
}

// NewInputMetadataWithDefaults instantiates a new InputMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputMetadataWithDefaults() *InputMetadata {
	this := InputMetadata{}
	return &this
}

// GetProperties returns the Properties field value
func (o *InputMetadata) GetProperties() []InputProperty {
	if o == nil {
		var ret []InputProperty
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *InputMetadata) GetPropertiesOk() (*[]InputProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *InputMetadata) SetProperties(v []InputProperty) {
	o.Properties = v
}

func (o InputMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableInputMetadata struct {
	value *InputMetadata
	isSet bool
}

func (v NullableInputMetadata) Get() *InputMetadata {
	return v.value
}

func (v *NullableInputMetadata) Set(val *InputMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableInputMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableInputMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputMetadata(val *InputMetadata) *NullableInputMetadata {
	return &NullableInputMetadata{value: val, isSet: true}
}

func (v NullableInputMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
