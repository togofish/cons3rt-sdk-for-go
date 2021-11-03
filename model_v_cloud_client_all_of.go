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

// VCloudClientAllOf struct for VCloudClientAllOf
type VCloudClientAllOf struct {
	OrgName *string `json:"orgName,omitempty"`
	VdcName *string `json:"vdcName,omitempty"`
}

// NewVCloudClientAllOf instantiates a new VCloudClientAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCloudClientAllOf() *VCloudClientAllOf {
	this := VCloudClientAllOf{}
	return &this
}

// NewVCloudClientAllOfWithDefaults instantiates a new VCloudClientAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCloudClientAllOfWithDefaults() *VCloudClientAllOf {
	this := VCloudClientAllOf{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *VCloudClientAllOf) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudClientAllOf) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *VCloudClientAllOf) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *VCloudClientAllOf) SetOrgName(v string) {
	o.OrgName = &v
}

// GetVdcName returns the VdcName field value if set, zero value otherwise.
func (o *VCloudClientAllOf) GetVdcName() string {
	if o == nil || o.VdcName == nil {
		var ret string
		return ret
	}
	return *o.VdcName
}

// GetVdcNameOk returns a tuple with the VdcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudClientAllOf) GetVdcNameOk() (*string, bool) {
	if o == nil || o.VdcName == nil {
		return nil, false
	}
	return o.VdcName, true
}

// HasVdcName returns a boolean if a field has been set.
func (o *VCloudClientAllOf) HasVdcName() bool {
	if o != nil && o.VdcName != nil {
		return true
	}

	return false
}

// SetVdcName gets a reference to the given string and assigns it to the VdcName field.
func (o *VCloudClientAllOf) SetVdcName(v string) {
	o.VdcName = &v
}

func (o VCloudClientAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.VdcName != nil {
		toSerialize["vdcName"] = o.VdcName
	}
	return json.Marshal(toSerialize)
}

type NullableVCloudClientAllOf struct {
	value *VCloudClientAllOf
	isSet bool
}

func (v NullableVCloudClientAllOf) Get() *VCloudClientAllOf {
	return v.value
}

func (v *NullableVCloudClientAllOf) Set(val *VCloudClientAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVCloudClientAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVCloudClientAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCloudClientAllOf(val *VCloudClientAllOf) *NullableVCloudClientAllOf {
	return &NullableVCloudClientAllOf{value: val, isSet: true}
}

func (v NullableVCloudClientAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCloudClientAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


