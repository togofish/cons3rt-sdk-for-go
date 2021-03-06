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

// AddVcloudRestNetworkRequest struct for AddVcloudRestNetworkRequest
type AddVcloudRestNetworkRequest struct {
	ExternalNetworkName string `json:"externalNetworkName"`
	Network             Network
	Subtype             string
}

// NewAddVcloudRestNetworkRequest instantiates a new AddVcloudRestNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVcloudRestNetworkRequest(externalNetworkName string, network Network, subtype string) *AddVcloudRestNetworkRequest {
	this := AddVcloudRestNetworkRequest{}
	this.Network = network
	this.Subtype = subtype
	this.ExternalNetworkName = externalNetworkName
	return &this
}

// NewAddVcloudRestNetworkRequestWithDefaults instantiates a new AddVcloudRestNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVcloudRestNetworkRequestWithDefaults() *AddVcloudRestNetworkRequest {
	this := AddVcloudRestNetworkRequest{}
	return &this
}

// GetExternalNetworkName returns the ExternalNetworkName field value
func (o *AddVcloudRestNetworkRequest) GetExternalNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalNetworkName
}

// GetExternalNetworkNameOk returns a tuple with the ExternalNetworkName field value
// and a boolean to check if the value has been set.
func (o *AddVcloudRestNetworkRequest) GetExternalNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalNetworkName, true
}

// SetExternalNetworkName sets field value
func (o *AddVcloudRestNetworkRequest) SetExternalNetworkName(v string) {
	o.ExternalNetworkName = v
}

func (o AddVcloudRestNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["externalNetworkName"] = o.ExternalNetworkName
	}
	return json.Marshal(toSerialize)
}

type NullableAddVcloudRestNetworkRequest struct {
	value *AddVcloudRestNetworkRequest
	isSet bool
}

func (v NullableAddVcloudRestNetworkRequest) Get() *AddVcloudRestNetworkRequest {
	return v.value
}

func (v *NullableAddVcloudRestNetworkRequest) Set(val *AddVcloudRestNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVcloudRestNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVcloudRestNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVcloudRestNetworkRequest(val *AddVcloudRestNetworkRequest) *NullableAddVcloudRestNetworkRequest {
	return &NullableAddVcloudRestNetworkRequest{value: val, isSet: true}
}

func (v NullableAddVcloudRestNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVcloudRestNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
