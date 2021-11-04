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

// InputAwsVirtualizationRealmAllOf struct for InputAwsVirtualizationRealmAllOf
type InputAwsVirtualizationRealmAllOf struct {
	Region *string `json:"region,omitempty"`
	VpcId *string `json:"vpcId,omitempty"`
}

// NewInputAwsVirtualizationRealmAllOf instantiates a new InputAwsVirtualizationRealmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputAwsVirtualizationRealmAllOf() *InputAwsVirtualizationRealmAllOf {
	this := InputAwsVirtualizationRealmAllOf{}
	return &this
}

// NewInputAwsVirtualizationRealmAllOfWithDefaults instantiates a new InputAwsVirtualizationRealmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputAwsVirtualizationRealmAllOfWithDefaults() *InputAwsVirtualizationRealmAllOf {
	this := InputAwsVirtualizationRealmAllOf{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *InputAwsVirtualizationRealmAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAwsVirtualizationRealmAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *InputAwsVirtualizationRealmAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *InputAwsVirtualizationRealmAllOf) SetRegion(v string) {
	o.Region = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *InputAwsVirtualizationRealmAllOf) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAwsVirtualizationRealmAllOf) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *InputAwsVirtualizationRealmAllOf) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *InputAwsVirtualizationRealmAllOf) SetVpcId(v string) {
	o.VpcId = &v
}

func (o InputAwsVirtualizationRealmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.VpcId != nil {
		toSerialize["vpcId"] = o.VpcId
	}
	return json.Marshal(toSerialize)
}

type NullableInputAwsVirtualizationRealmAllOf struct {
	value *InputAwsVirtualizationRealmAllOf
	isSet bool
}

func (v NullableInputAwsVirtualizationRealmAllOf) Get() *InputAwsVirtualizationRealmAllOf {
	return v.value
}

func (v *NullableInputAwsVirtualizationRealmAllOf) Set(val *InputAwsVirtualizationRealmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInputAwsVirtualizationRealmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInputAwsVirtualizationRealmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputAwsVirtualizationRealmAllOf(val *InputAwsVirtualizationRealmAllOf) *NullableInputAwsVirtualizationRealmAllOf {
	return &NullableInputAwsVirtualizationRealmAllOf{value: val, isSet: true}
}

func (v NullableInputAwsVirtualizationRealmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputAwsVirtualizationRealmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


