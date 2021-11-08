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

// AddVcloudNetworkRequestAllOf struct for AddVcloudNetworkRequestAllOf
type AddVcloudNetworkRequestAllOf struct {
	ExternalNetworkName *string `json:"externalNetworkName,omitempty"`
}

// NewAddVcloudNetworkRequestAllOf instantiates a new AddVcloudNetworkRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVcloudNetworkRequestAllOf() *AddVcloudNetworkRequestAllOf {
	this := AddVcloudNetworkRequestAllOf{}
	return &this
}

// NewAddVcloudNetworkRequestAllOfWithDefaults instantiates a new AddVcloudNetworkRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVcloudNetworkRequestAllOfWithDefaults() *AddVcloudNetworkRequestAllOf {
	this := AddVcloudNetworkRequestAllOf{}
	return &this
}

// GetExternalNetworkName returns the ExternalNetworkName field value if set, zero value otherwise.
func (o *AddVcloudNetworkRequestAllOf) GetExternalNetworkName() string {
	if o == nil || o.ExternalNetworkName == nil {
		var ret string
		return ret
	}
	return *o.ExternalNetworkName
}

// GetExternalNetworkNameOk returns a tuple with the ExternalNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVcloudNetworkRequestAllOf) GetExternalNetworkNameOk() (*string, bool) {
	if o == nil || o.ExternalNetworkName == nil {
		return nil, false
	}
	return o.ExternalNetworkName, true
}

// HasExternalNetworkName returns a boolean if a field has been set.
func (o *AddVcloudNetworkRequestAllOf) HasExternalNetworkName() bool {
	if o != nil && o.ExternalNetworkName != nil {
		return true
	}

	return false
}

// SetExternalNetworkName gets a reference to the given string and assigns it to the ExternalNetworkName field.
func (o *AddVcloudNetworkRequestAllOf) SetExternalNetworkName(v string) {
	o.ExternalNetworkName = &v
}

func (o AddVcloudNetworkRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalNetworkName != nil {
		toSerialize["externalNetworkName"] = o.ExternalNetworkName
	}
	return json.Marshal(toSerialize)
}

type NullableAddVcloudNetworkRequestAllOf struct {
	value *AddVcloudNetworkRequestAllOf
	isSet bool
}

func (v NullableAddVcloudNetworkRequestAllOf) Get() *AddVcloudNetworkRequestAllOf {
	return v.value
}

func (v *NullableAddVcloudNetworkRequestAllOf) Set(val *AddVcloudNetworkRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVcloudNetworkRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVcloudNetworkRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVcloudNetworkRequestAllOf(val *AddVcloudNetworkRequestAllOf) *NullableAddVcloudNetworkRequestAllOf {
	return &NullableAddVcloudNetworkRequestAllOf{value: val, isSet: true}
}

func (v NullableAddVcloudNetworkRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVcloudNetworkRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
