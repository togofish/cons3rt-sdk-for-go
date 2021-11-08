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

// InputTestBundle struct for InputTestBundle
type InputTestBundle struct {
	TestAsset        *InputTestAsset `json:"testAsset,omitempty"`
	TestToolPoolType string          `json:"testToolPoolType"`
}

// NewInputTestBundle instantiates a new InputTestBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputTestBundle(testToolPoolType string) *InputTestBundle {
	this := InputTestBundle{}
	this.TestToolPoolType = testToolPoolType
	return &this
}

// NewInputTestBundleWithDefaults instantiates a new InputTestBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputTestBundleWithDefaults() *InputTestBundle {
	this := InputTestBundle{}
	return &this
}

// GetTestAsset returns the TestAsset field value if set, zero value otherwise.
func (o *InputTestBundle) GetTestAsset() InputTestAsset {
	if o == nil || o.TestAsset == nil {
		var ret InputTestAsset
		return ret
	}
	return *o.TestAsset
}

// GetTestAssetOk returns a tuple with the TestAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestBundle) GetTestAssetOk() (*InputTestAsset, bool) {
	if o == nil || o.TestAsset == nil {
		return nil, false
	}
	return o.TestAsset, true
}

// HasTestAsset returns a boolean if a field has been set.
func (o *InputTestBundle) HasTestAsset() bool {
	if o != nil && o.TestAsset != nil {
		return true
	}

	return false
}

// SetTestAsset gets a reference to the given InputTestAsset and assigns it to the TestAsset field.
func (o *InputTestBundle) SetTestAsset(v InputTestAsset) {
	o.TestAsset = &v
}

// GetTestToolPoolType returns the TestToolPoolType field value
func (o *InputTestBundle) GetTestToolPoolType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestToolPoolType
}

// GetTestToolPoolTypeOk returns a tuple with the TestToolPoolType field value
// and a boolean to check if the value has been set.
func (o *InputTestBundle) GetTestToolPoolTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestToolPoolType, true
}

// SetTestToolPoolType sets field value
func (o *InputTestBundle) SetTestToolPoolType(v string) {
	o.TestToolPoolType = v
}

func (o InputTestBundle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TestAsset != nil {
		toSerialize["testAsset"] = o.TestAsset
	}
	if true {
		toSerialize["testToolPoolType"] = o.TestToolPoolType
	}
	return json.Marshal(toSerialize)
}

type NullableInputTestBundle struct {
	value *InputTestBundle
	isSet bool
}

func (v NullableInputTestBundle) Get() *InputTestBundle {
	return v.value
}

func (v *NullableInputTestBundle) Set(val *InputTestBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTestBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTestBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTestBundle(val *InputTestBundle) *NullableInputTestBundle {
	return &NullableInputTestBundle{value: val, isSet: true}
}

func (v NullableInputTestBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputTestBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
