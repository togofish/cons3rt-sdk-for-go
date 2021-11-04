/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gocons3rt

import (
	"encoding/json"
)

// MinimalVirtualHost struct for MinimalVirtualHost
type MinimalVirtualHost struct {
	Subtype string
}

// NewMinimalVirtualHost instantiates a new MinimalVirtualHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalVirtualHost(subtype string) *MinimalVirtualHost {
	this := MinimalVirtualHost{}
	this.Subtype = subtype
	return &this
}

// NewMinimalVirtualHostWithDefaults instantiates a new MinimalVirtualHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalVirtualHostWithDefaults() *MinimalVirtualHost {
	this := MinimalVirtualHost{}
	return &this
}

func (o MinimalVirtualHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableMinimalVirtualHost struct {
	value *MinimalVirtualHost
	isSet bool
}

func (v NullableMinimalVirtualHost) Get() *MinimalVirtualHost {
	return v.value
}

func (v *NullableMinimalVirtualHost) Set(val *MinimalVirtualHost) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalVirtualHost) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalVirtualHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalVirtualHost(val *MinimalVirtualHost) *NullableMinimalVirtualHost {
	return &NullableMinimalVirtualHost{value: val, isSet: true}
}

func (v NullableMinimalVirtualHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalVirtualHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
