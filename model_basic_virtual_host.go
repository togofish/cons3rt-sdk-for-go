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

// BasicVirtualHost struct for BasicVirtualHost
type BasicVirtualHost struct {
}

// NewBasicVirtualHost instantiates a new BasicVirtualHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicVirtualHost(subtype string) *BasicVirtualHost {
	this := BasicVirtualHost{}
	this.Subtype = subtype
	return &this
}

// NewBasicVirtualHostWithDefaults instantiates a new BasicVirtualHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicVirtualHostWithDefaults() *BasicVirtualHost {
	this := BasicVirtualHost{}
	return &this
}

func (o BasicVirtualHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableBasicVirtualHost struct {
	value *BasicVirtualHost
	isSet bool
}

func (v NullableBasicVirtualHost) Get() *BasicVirtualHost {
	return v.value
}

func (v *NullableBasicVirtualHost) Set(val *BasicVirtualHost) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicVirtualHost) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicVirtualHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicVirtualHost(val *BasicVirtualHost) *NullableBasicVirtualHost {
	return &NullableBasicVirtualHost{value: val, isSet: true}
}

func (v NullableBasicVirtualHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicVirtualHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


