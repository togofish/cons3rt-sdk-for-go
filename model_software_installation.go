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

// SoftwareInstallation struct for SoftwareInstallation
type SoftwareInstallation struct {
	AssetType          *string `json:"assetType,omitempty"`
	RebootDelaySeconds *int32  `json:"rebootDelaySeconds,omitempty"`
	RebootRequired     *bool   `json:"rebootRequired,omitempty"`
	Subtype            string
}

// NewSoftwareInstallation instantiates a new SoftwareInstallation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareInstallation(subtype string) *SoftwareInstallation {
	this := SoftwareInstallation{}
	this.Subtype = subtype
	return &this
}

// NewSoftwareInstallationWithDefaults instantiates a new SoftwareInstallation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareInstallationWithDefaults() *SoftwareInstallation {
	this := SoftwareInstallation{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *SoftwareInstallation) GetAssetType() string {
	if o == nil || o.AssetType == nil {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareInstallation) GetAssetTypeOk() (*string, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *SoftwareInstallation) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *SoftwareInstallation) SetAssetType(v string) {
	o.AssetType = &v
}

// GetRebootDelaySeconds returns the RebootDelaySeconds field value if set, zero value otherwise.
func (o *SoftwareInstallation) GetRebootDelaySeconds() int32 {
	if o == nil || o.RebootDelaySeconds == nil {
		var ret int32
		return ret
	}
	return *o.RebootDelaySeconds
}

// GetRebootDelaySecondsOk returns a tuple with the RebootDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareInstallation) GetRebootDelaySecondsOk() (*int32, bool) {
	if o == nil || o.RebootDelaySeconds == nil {
		return nil, false
	}
	return o.RebootDelaySeconds, true
}

// HasRebootDelaySeconds returns a boolean if a field has been set.
func (o *SoftwareInstallation) HasRebootDelaySeconds() bool {
	if o != nil && o.RebootDelaySeconds != nil {
		return true
	}

	return false
}

// SetRebootDelaySeconds gets a reference to the given int32 and assigns it to the RebootDelaySeconds field.
func (o *SoftwareInstallation) SetRebootDelaySeconds(v int32) {
	o.RebootDelaySeconds = &v
}

// GetRebootRequired returns the RebootRequired field value if set, zero value otherwise.
func (o *SoftwareInstallation) GetRebootRequired() bool {
	if o == nil || o.RebootRequired == nil {
		var ret bool
		return ret
	}
	return *o.RebootRequired
}

// GetRebootRequiredOk returns a tuple with the RebootRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareInstallation) GetRebootRequiredOk() (*bool, bool) {
	if o == nil || o.RebootRequired == nil {
		return nil, false
	}
	return o.RebootRequired, true
}

// HasRebootRequired returns a boolean if a field has been set.
func (o *SoftwareInstallation) HasRebootRequired() bool {
	if o != nil && o.RebootRequired != nil {
		return true
	}

	return false
}

// SetRebootRequired gets a reference to the given bool and assigns it to the RebootRequired field.
func (o *SoftwareInstallation) SetRebootRequired(v bool) {
	o.RebootRequired = &v
}

func (o SoftwareInstallation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.RebootDelaySeconds != nil {
		toSerialize["rebootDelaySeconds"] = o.RebootDelaySeconds
	}
	if o.RebootRequired != nil {
		toSerialize["rebootRequired"] = o.RebootRequired
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwareInstallation struct {
	value *SoftwareInstallation
	isSet bool
}

func (v NullableSoftwareInstallation) Get() *SoftwareInstallation {
	return v.value
}

func (v *NullableSoftwareInstallation) Set(val *SoftwareInstallation) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareInstallation) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareInstallation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareInstallation(val *SoftwareInstallation) *NullableSoftwareInstallation {
	return &NullableSoftwareInstallation{value: val, isSet: true}
}

func (v NullableSoftwareInstallation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareInstallation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
