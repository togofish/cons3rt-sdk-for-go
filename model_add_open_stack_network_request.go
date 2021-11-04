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

// AddOpenStackNetworkRequest struct for AddOpenStackNetworkRequest
type AddOpenStackNetworkRequest struct {
	NatImageId        string `json:"natImageId"`
	NatInstanceFlavor string `json:"natInstanceFlavor"`
	Network           Network
	Subtype           string
}

// NewAddOpenStackNetworkRequest instantiates a new AddOpenStackNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOpenStackNetworkRequest(natImageId string, natInstanceFlavor string, network Network, subtype string) *AddOpenStackNetworkRequest {
	this := AddOpenStackNetworkRequest{}
	this.Network = network
	this.Subtype = subtype
	this.NatImageId = natImageId
	this.NatInstanceFlavor = natInstanceFlavor
	return &this
}

// NewAddOpenStackNetworkRequestWithDefaults instantiates a new AddOpenStackNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOpenStackNetworkRequestWithDefaults() *AddOpenStackNetworkRequest {
	this := AddOpenStackNetworkRequest{}
	return &this
}

// GetNatImageId returns the NatImageId field value
func (o *AddOpenStackNetworkRequest) GetNatImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NatImageId
}

// GetNatImageIdOk returns a tuple with the NatImageId field value
// and a boolean to check if the value has been set.
func (o *AddOpenStackNetworkRequest) GetNatImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NatImageId, true
}

// SetNatImageId sets field value
func (o *AddOpenStackNetworkRequest) SetNatImageId(v string) {
	o.NatImageId = v
}

// GetNatInstanceFlavor returns the NatInstanceFlavor field value
func (o *AddOpenStackNetworkRequest) GetNatInstanceFlavor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NatInstanceFlavor
}

// GetNatInstanceFlavorOk returns a tuple with the NatInstanceFlavor field value
// and a boolean to check if the value has been set.
func (o *AddOpenStackNetworkRequest) GetNatInstanceFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NatInstanceFlavor, true
}

// SetNatInstanceFlavor sets field value
func (o *AddOpenStackNetworkRequest) SetNatInstanceFlavor(v string) {
	o.NatInstanceFlavor = v
}

func (o AddOpenStackNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["natImageId"] = o.NatImageId
	}
	if true {
		toSerialize["natInstanceFlavor"] = o.NatInstanceFlavor
	}
	return json.Marshal(toSerialize)
}

type NullableAddOpenStackNetworkRequest struct {
	value *AddOpenStackNetworkRequest
	isSet bool
}

func (v NullableAddOpenStackNetworkRequest) Get() *AddOpenStackNetworkRequest {
	return v.value
}

func (v *NullableAddOpenStackNetworkRequest) Set(val *AddOpenStackNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOpenStackNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOpenStackNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOpenStackNetworkRequest(val *AddOpenStackNetworkRequest) *NullableAddOpenStackNetworkRequest {
	return &NullableAddOpenStackNetworkRequest{value: val, isSet: true}
}

func (v NullableAddOpenStackNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOpenStackNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
