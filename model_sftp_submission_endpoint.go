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

// SFTPSubmissionEndpoint struct for SFTPSubmissionEndpoint
type SFTPSubmissionEndpoint struct {
}

// NewSFTPSubmissionEndpoint instantiates a new SFTPSubmissionEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSFTPSubmissionEndpoint(host string, subtype string) *SFTPSubmissionEndpoint {
	this := SFTPSubmissionEndpoint{}
	this.Host = host
	this.Subtype = subtype
	return &this
}

// NewSFTPSubmissionEndpointWithDefaults instantiates a new SFTPSubmissionEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSFTPSubmissionEndpointWithDefaults() *SFTPSubmissionEndpoint {
	this := SFTPSubmissionEndpoint{}
	return &this
}

func (o SFTPSubmissionEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableSFTPSubmissionEndpoint struct {
	value *SFTPSubmissionEndpoint
	isSet bool
}

func (v NullableSFTPSubmissionEndpoint) Get() *SFTPSubmissionEndpoint {
	return v.value
}

func (v *NullableSFTPSubmissionEndpoint) Set(val *SFTPSubmissionEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSFTPSubmissionEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSFTPSubmissionEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSFTPSubmissionEndpoint(val *SFTPSubmissionEndpoint) *NullableSFTPSubmissionEndpoint {
	return &NullableSFTPSubmissionEndpoint{value: val, isSet: true}
}

func (v NullableSFTPSubmissionEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSFTPSubmissionEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


