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

// BasicTestAsset struct for BasicTestAsset
type BasicTestAsset struct {
	Type    *string `json:"type,omitempty"`
	Subtype string
}

// NewBasicTestAsset instantiates a new BasicTestAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicTestAsset(subtype string) *BasicTestAsset {
	this := BasicTestAsset{}
	this.Subtype = subtype
	return &this
}

// NewBasicTestAssetWithDefaults instantiates a new BasicTestAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicTestAssetWithDefaults() *BasicTestAsset {
	this := BasicTestAsset{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BasicTestAsset) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicTestAsset) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BasicTestAsset) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BasicTestAsset) SetType(v string) {
	o.Type = &v
}

func (o BasicTestAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBasicTestAsset struct {
	value *BasicTestAsset
	isSet bool
}

func (v NullableBasicTestAsset) Get() *BasicTestAsset {
	return v.value
}

func (v *NullableBasicTestAsset) Set(val *BasicTestAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicTestAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicTestAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicTestAsset(val *BasicTestAsset) *NullableBasicTestAsset {
	return &NullableBasicTestAsset{value: val, isSet: true}
}

func (v NullableBasicTestAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicTestAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
