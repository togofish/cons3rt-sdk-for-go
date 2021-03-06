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

// CompositionLaunchOptions struct for CompositionLaunchOptions
type CompositionLaunchOptions struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	TimeToLive *int32 `json:"timeToLive,omitempty"`
}

// NewCompositionLaunchOptions instantiates a new CompositionLaunchOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositionLaunchOptions(username string, password string) *CompositionLaunchOptions {
	this := CompositionLaunchOptions{}
	this.Username = username
	this.Password = password
	return &this
}

// NewCompositionLaunchOptionsWithDefaults instantiates a new CompositionLaunchOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositionLaunchOptionsWithDefaults() *CompositionLaunchOptions {
	this := CompositionLaunchOptions{}
	return &this
}

// GetUsername returns the Username field value
func (o *CompositionLaunchOptions) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CompositionLaunchOptions) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CompositionLaunchOptions) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CompositionLaunchOptions) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CompositionLaunchOptions) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CompositionLaunchOptions) SetPassword(v string) {
	o.Password = v
}

// GetTimeToLive returns the TimeToLive field value if set, zero value otherwise.
func (o *CompositionLaunchOptions) GetTimeToLive() int32 {
	if o == nil || o.TimeToLive == nil {
		var ret int32
		return ret
	}
	return *o.TimeToLive
}

// GetTimeToLiveOk returns a tuple with the TimeToLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositionLaunchOptions) GetTimeToLiveOk() (*int32, bool) {
	if o == nil || o.TimeToLive == nil {
		return nil, false
	}
	return o.TimeToLive, true
}

// HasTimeToLive returns a boolean if a field has been set.
func (o *CompositionLaunchOptions) HasTimeToLive() bool {
	if o != nil && o.TimeToLive != nil {
		return true
	}

	return false
}

// SetTimeToLive gets a reference to the given int32 and assigns it to the TimeToLive field.
func (o *CompositionLaunchOptions) SetTimeToLive(v int32) {
	o.TimeToLive = &v
}

func (o CompositionLaunchOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.TimeToLive != nil {
		toSerialize["timeToLive"] = o.TimeToLive
	}
	return json.Marshal(toSerialize)
}

type NullableCompositionLaunchOptions struct {
	value *CompositionLaunchOptions
	isSet bool
}

func (v NullableCompositionLaunchOptions) Get() *CompositionLaunchOptions {
	return v.value
}

func (v *NullableCompositionLaunchOptions) Set(val *CompositionLaunchOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositionLaunchOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositionLaunchOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositionLaunchOptions(val *CompositionLaunchOptions) *NullableCompositionLaunchOptions {
	return &NullableCompositionLaunchOptions{value: val, isSet: true}
}

func (v NullableCompositionLaunchOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositionLaunchOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
