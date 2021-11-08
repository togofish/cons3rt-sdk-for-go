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

// SoftwareInstallationAllOf struct for SoftwareInstallationAllOf
type SoftwareInstallationAllOf struct {
	AssetType          *string `json:"assetType,omitempty"`
	RebootDelaySeconds *int32  `json:"rebootDelaySeconds,omitempty"`
	RebootRequired     *bool   `json:"rebootRequired,omitempty"`
}

// NewSoftwareInstallationAllOf instantiates a new SoftwareInstallationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareInstallationAllOf() *SoftwareInstallationAllOf {
	this := SoftwareInstallationAllOf{}
	return &this
}

// NewSoftwareInstallationAllOfWithDefaults instantiates a new SoftwareInstallationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareInstallationAllOfWithDefaults() *SoftwareInstallationAllOf {
	this := SoftwareInstallationAllOf{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *SoftwareInstallationAllOf) GetAssetType() string {
	if o == nil || o.AssetType == nil {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareInstallationAllOf) GetAssetTypeOk() (*string, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *SoftwareInstallationAllOf) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *SoftwareInstallationAllOf) SetAssetType(v string) {
	o.AssetType = &v
}

// GetRebootDelaySeconds returns the RebootDelaySeconds field value if set, zero value otherwise.
func (o *SoftwareInstallationAllOf) GetRebootDelaySeconds() int32 {
	if o == nil || o.RebootDelaySeconds == nil {
		var ret int32
		return ret
	}
	return *o.RebootDelaySeconds
}

// GetRebootDelaySecondsOk returns a tuple with the RebootDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareInstallationAllOf) GetRebootDelaySecondsOk() (*int32, bool) {
	if o == nil || o.RebootDelaySeconds == nil {
		return nil, false
	}
	return o.RebootDelaySeconds, true
}

// HasRebootDelaySeconds returns a boolean if a field has been set.
func (o *SoftwareInstallationAllOf) HasRebootDelaySeconds() bool {
	if o != nil && o.RebootDelaySeconds != nil {
		return true
	}

	return false
}

// SetRebootDelaySeconds gets a reference to the given int32 and assigns it to the RebootDelaySeconds field.
func (o *SoftwareInstallationAllOf) SetRebootDelaySeconds(v int32) {
	o.RebootDelaySeconds = &v
}

// GetRebootRequired returns the RebootRequired field value if set, zero value otherwise.
func (o *SoftwareInstallationAllOf) GetRebootRequired() bool {
	if o == nil || o.RebootRequired == nil {
		var ret bool
		return ret
	}
	return *o.RebootRequired
}

// GetRebootRequiredOk returns a tuple with the RebootRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareInstallationAllOf) GetRebootRequiredOk() (*bool, bool) {
	if o == nil || o.RebootRequired == nil {
		return nil, false
	}
	return o.RebootRequired, true
}

// HasRebootRequired returns a boolean if a field has been set.
func (o *SoftwareInstallationAllOf) HasRebootRequired() bool {
	if o != nil && o.RebootRequired != nil {
		return true
	}

	return false
}

// SetRebootRequired gets a reference to the given bool and assigns it to the RebootRequired field.
func (o *SoftwareInstallationAllOf) SetRebootRequired(v bool) {
	o.RebootRequired = &v
}

func (o SoftwareInstallationAllOf) MarshalJSON() ([]byte, error) {
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

type NullableSoftwareInstallationAllOf struct {
	value *SoftwareInstallationAllOf
	isSet bool
}

func (v NullableSoftwareInstallationAllOf) Get() *SoftwareInstallationAllOf {
	return v.value
}

func (v *NullableSoftwareInstallationAllOf) Set(val *SoftwareInstallationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareInstallationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareInstallationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareInstallationAllOf(val *SoftwareInstallationAllOf) *NullableSoftwareInstallationAllOf {
	return &NullableSoftwareInstallationAllOf{value: val, isSet: true}
}

func (v NullableSoftwareInstallationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareInstallationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
