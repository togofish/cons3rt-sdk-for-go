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

// InputOpenStackClient struct for InputOpenStackClient
type InputOpenStackClient struct {
	Username string
	Password string
	Subtype  string
}

// NewInputOpenStackClient instantiates a new InputOpenStackClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputOpenStackClient(username string, password string, subtype string) *InputOpenStackClient {
	this := InputOpenStackClient{}
	this.Username = username
	this.Password = password
	this.Subtype = subtype
	return &this
}

// NewInputOpenStackClientWithDefaults instantiates a new InputOpenStackClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputOpenStackClientWithDefaults() *InputOpenStackClient {
	this := InputOpenStackClient{}
	return &this
}

func (o InputOpenStackClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableInputOpenStackClient struct {
	value *InputOpenStackClient
	isSet bool
}

func (v NullableInputOpenStackClient) Get() *InputOpenStackClient {
	return v.value
}

func (v *NullableInputOpenStackClient) Set(val *InputOpenStackClient) {
	v.value = val
	v.isSet = true
}

func (v NullableInputOpenStackClient) IsSet() bool {
	return v.isSet
}

func (v *NullableInputOpenStackClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputOpenStackClient(val *InputOpenStackClient) *NullableInputOpenStackClient {
	return &NullableInputOpenStackClient{value: val, isSet: true}
}

func (v NullableInputOpenStackClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputOpenStackClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
