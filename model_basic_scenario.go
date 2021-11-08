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

// BasicScenario struct for BasicScenario
type BasicScenario struct {
	Subtype string
}

// NewBasicScenario instantiates a new BasicScenario object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicScenario(subtype string) *BasicScenario {
	this := BasicScenario{}
	this.Subtype = subtype
	return &this
}

// NewBasicScenarioWithDefaults instantiates a new BasicScenario object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicScenarioWithDefaults() *BasicScenario {
	this := BasicScenario{}
	return &this
}

func (o BasicScenario) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableBasicScenario struct {
	value *BasicScenario
	isSet bool
}

func (v NullableBasicScenario) Get() *BasicScenario {
	return v.value
}

func (v *NullableBasicScenario) Set(val *BasicScenario) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicScenario) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicScenario) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicScenario(val *BasicScenario) *NullableBasicScenario {
	return &NullableBasicScenario{value: val, isSet: true}
}

func (v NullableBasicScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicScenario) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
