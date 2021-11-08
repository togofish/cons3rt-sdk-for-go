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

// AddAzureNetworkRequestAllOf struct for AddAzureNetworkRequestAllOf
type AddAzureNetworkRequestAllOf struct {
	NatInstanceType   *string            `json:"natInstanceType,omitempty"`
	NatImageReference *ImageReferenceDTO `json:"natImageReference,omitempty"`
}

// NewAddAzureNetworkRequestAllOf instantiates a new AddAzureNetworkRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAzureNetworkRequestAllOf() *AddAzureNetworkRequestAllOf {
	this := AddAzureNetworkRequestAllOf{}
	return &this
}

// NewAddAzureNetworkRequestAllOfWithDefaults instantiates a new AddAzureNetworkRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAzureNetworkRequestAllOfWithDefaults() *AddAzureNetworkRequestAllOf {
	this := AddAzureNetworkRequestAllOf{}
	return &this
}

// GetNatInstanceType returns the NatInstanceType field value if set, zero value otherwise.
func (o *AddAzureNetworkRequestAllOf) GetNatInstanceType() string {
	if o == nil || o.NatInstanceType == nil {
		var ret string
		return ret
	}
	return *o.NatInstanceType
}

// GetNatInstanceTypeOk returns a tuple with the NatInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAzureNetworkRequestAllOf) GetNatInstanceTypeOk() (*string, bool) {
	if o == nil || o.NatInstanceType == nil {
		return nil, false
	}
	return o.NatInstanceType, true
}

// HasNatInstanceType returns a boolean if a field has been set.
func (o *AddAzureNetworkRequestAllOf) HasNatInstanceType() bool {
	if o != nil && o.NatInstanceType != nil {
		return true
	}

	return false
}

// SetNatInstanceType gets a reference to the given string and assigns it to the NatInstanceType field.
func (o *AddAzureNetworkRequestAllOf) SetNatInstanceType(v string) {
	o.NatInstanceType = &v
}

// GetNatImageReference returns the NatImageReference field value if set, zero value otherwise.
func (o *AddAzureNetworkRequestAllOf) GetNatImageReference() ImageReferenceDTO {
	if o == nil || o.NatImageReference == nil {
		var ret ImageReferenceDTO
		return ret
	}
	return *o.NatImageReference
}

// GetNatImageReferenceOk returns a tuple with the NatImageReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAzureNetworkRequestAllOf) GetNatImageReferenceOk() (*ImageReferenceDTO, bool) {
	if o == nil || o.NatImageReference == nil {
		return nil, false
	}
	return o.NatImageReference, true
}

// HasNatImageReference returns a boolean if a field has been set.
func (o *AddAzureNetworkRequestAllOf) HasNatImageReference() bool {
	if o != nil && o.NatImageReference != nil {
		return true
	}

	return false
}

// SetNatImageReference gets a reference to the given ImageReferenceDTO and assigns it to the NatImageReference field.
func (o *AddAzureNetworkRequestAllOf) SetNatImageReference(v ImageReferenceDTO) {
	o.NatImageReference = &v
}

func (o AddAzureNetworkRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NatInstanceType != nil {
		toSerialize["natInstanceType"] = o.NatInstanceType
	}
	if o.NatImageReference != nil {
		toSerialize["natImageReference"] = o.NatImageReference
	}
	return json.Marshal(toSerialize)
}

type NullableAddAzureNetworkRequestAllOf struct {
	value *AddAzureNetworkRequestAllOf
	isSet bool
}

func (v NullableAddAzureNetworkRequestAllOf) Get() *AddAzureNetworkRequestAllOf {
	return v.value
}

func (v *NullableAddAzureNetworkRequestAllOf) Set(val *AddAzureNetworkRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAzureNetworkRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAzureNetworkRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAzureNetworkRequestAllOf(val *AddAzureNetworkRequestAllOf) *NullableAddAzureNetworkRequestAllOf {
	return &NullableAddAzureNetworkRequestAllOf{value: val, isSet: true}
}

func (v NullableAddAzureNetworkRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAzureNetworkRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
