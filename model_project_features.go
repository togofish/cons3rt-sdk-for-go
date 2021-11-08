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

// ProjectFeatures struct for ProjectFeatures
type ProjectFeatures struct {
	AssetBypassScanningEnabled      *bool `json:"assetBypassScanningEnabled,omitempty"`
	RemoteAccessCopyEnabled         *bool `json:"remoteAccessCopyEnabled,omitempty"`
	RemoteAccessFileDownloadEnabled *bool `json:"remoteAccessFileDownloadEnabled,omitempty"`
	RemoteAccessFileUploadEnabled   *bool `json:"remoteAccessFileUploadEnabled,omitempty"`
	RemoteAccessPasteEnabled        *bool `json:"remoteAccessPasteEnabled,omitempty"`
}

// NewProjectFeatures instantiates a new ProjectFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectFeatures() *ProjectFeatures {
	this := ProjectFeatures{}
	return &this
}

// NewProjectFeaturesWithDefaults instantiates a new ProjectFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectFeaturesWithDefaults() *ProjectFeatures {
	this := ProjectFeatures{}
	return &this
}

// GetAssetBypassScanningEnabled returns the AssetBypassScanningEnabled field value if set, zero value otherwise.
func (o *ProjectFeatures) GetAssetBypassScanningEnabled() bool {
	if o == nil || o.AssetBypassScanningEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AssetBypassScanningEnabled
}

// GetAssetBypassScanningEnabledOk returns a tuple with the AssetBypassScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatures) GetAssetBypassScanningEnabledOk() (*bool, bool) {
	if o == nil || o.AssetBypassScanningEnabled == nil {
		return nil, false
	}
	return o.AssetBypassScanningEnabled, true
}

// HasAssetBypassScanningEnabled returns a boolean if a field has been set.
func (o *ProjectFeatures) HasAssetBypassScanningEnabled() bool {
	if o != nil && o.AssetBypassScanningEnabled != nil {
		return true
	}

	return false
}

// SetAssetBypassScanningEnabled gets a reference to the given bool and assigns it to the AssetBypassScanningEnabled field.
func (o *ProjectFeatures) SetAssetBypassScanningEnabled(v bool) {
	o.AssetBypassScanningEnabled = &v
}

// GetRemoteAccessCopyEnabled returns the RemoteAccessCopyEnabled field value if set, zero value otherwise.
func (o *ProjectFeatures) GetRemoteAccessCopyEnabled() bool {
	if o == nil || o.RemoteAccessCopyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessCopyEnabled
}

// GetRemoteAccessCopyEnabledOk returns a tuple with the RemoteAccessCopyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatures) GetRemoteAccessCopyEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteAccessCopyEnabled == nil {
		return nil, false
	}
	return o.RemoteAccessCopyEnabled, true
}

// HasRemoteAccessCopyEnabled returns a boolean if a field has been set.
func (o *ProjectFeatures) HasRemoteAccessCopyEnabled() bool {
	if o != nil && o.RemoteAccessCopyEnabled != nil {
		return true
	}

	return false
}

// SetRemoteAccessCopyEnabled gets a reference to the given bool and assigns it to the RemoteAccessCopyEnabled field.
func (o *ProjectFeatures) SetRemoteAccessCopyEnabled(v bool) {
	o.RemoteAccessCopyEnabled = &v
}

// GetRemoteAccessFileDownloadEnabled returns the RemoteAccessFileDownloadEnabled field value if set, zero value otherwise.
func (o *ProjectFeatures) GetRemoteAccessFileDownloadEnabled() bool {
	if o == nil || o.RemoteAccessFileDownloadEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessFileDownloadEnabled
}

// GetRemoteAccessFileDownloadEnabledOk returns a tuple with the RemoteAccessFileDownloadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatures) GetRemoteAccessFileDownloadEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteAccessFileDownloadEnabled == nil {
		return nil, false
	}
	return o.RemoteAccessFileDownloadEnabled, true
}

// HasRemoteAccessFileDownloadEnabled returns a boolean if a field has been set.
func (o *ProjectFeatures) HasRemoteAccessFileDownloadEnabled() bool {
	if o != nil && o.RemoteAccessFileDownloadEnabled != nil {
		return true
	}

	return false
}

// SetRemoteAccessFileDownloadEnabled gets a reference to the given bool and assigns it to the RemoteAccessFileDownloadEnabled field.
func (o *ProjectFeatures) SetRemoteAccessFileDownloadEnabled(v bool) {
	o.RemoteAccessFileDownloadEnabled = &v
}

// GetRemoteAccessFileUploadEnabled returns the RemoteAccessFileUploadEnabled field value if set, zero value otherwise.
func (o *ProjectFeatures) GetRemoteAccessFileUploadEnabled() bool {
	if o == nil || o.RemoteAccessFileUploadEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessFileUploadEnabled
}

// GetRemoteAccessFileUploadEnabledOk returns a tuple with the RemoteAccessFileUploadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatures) GetRemoteAccessFileUploadEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteAccessFileUploadEnabled == nil {
		return nil, false
	}
	return o.RemoteAccessFileUploadEnabled, true
}

// HasRemoteAccessFileUploadEnabled returns a boolean if a field has been set.
func (o *ProjectFeatures) HasRemoteAccessFileUploadEnabled() bool {
	if o != nil && o.RemoteAccessFileUploadEnabled != nil {
		return true
	}

	return false
}

// SetRemoteAccessFileUploadEnabled gets a reference to the given bool and assigns it to the RemoteAccessFileUploadEnabled field.
func (o *ProjectFeatures) SetRemoteAccessFileUploadEnabled(v bool) {
	o.RemoteAccessFileUploadEnabled = &v
}

// GetRemoteAccessPasteEnabled returns the RemoteAccessPasteEnabled field value if set, zero value otherwise.
func (o *ProjectFeatures) GetRemoteAccessPasteEnabled() bool {
	if o == nil || o.RemoteAccessPasteEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessPasteEnabled
}

// GetRemoteAccessPasteEnabledOk returns a tuple with the RemoteAccessPasteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatures) GetRemoteAccessPasteEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteAccessPasteEnabled == nil {
		return nil, false
	}
	return o.RemoteAccessPasteEnabled, true
}

// HasRemoteAccessPasteEnabled returns a boolean if a field has been set.
func (o *ProjectFeatures) HasRemoteAccessPasteEnabled() bool {
	if o != nil && o.RemoteAccessPasteEnabled != nil {
		return true
	}

	return false
}

// SetRemoteAccessPasteEnabled gets a reference to the given bool and assigns it to the RemoteAccessPasteEnabled field.
func (o *ProjectFeatures) SetRemoteAccessPasteEnabled(v bool) {
	o.RemoteAccessPasteEnabled = &v
}

func (o ProjectFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetBypassScanningEnabled != nil {
		toSerialize["assetBypassScanningEnabled"] = o.AssetBypassScanningEnabled
	}
	if o.RemoteAccessCopyEnabled != nil {
		toSerialize["remoteAccessCopyEnabled"] = o.RemoteAccessCopyEnabled
	}
	if o.RemoteAccessFileDownloadEnabled != nil {
		toSerialize["remoteAccessFileDownloadEnabled"] = o.RemoteAccessFileDownloadEnabled
	}
	if o.RemoteAccessFileUploadEnabled != nil {
		toSerialize["remoteAccessFileUploadEnabled"] = o.RemoteAccessFileUploadEnabled
	}
	if o.RemoteAccessPasteEnabled != nil {
		toSerialize["remoteAccessPasteEnabled"] = o.RemoteAccessPasteEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableProjectFeatures struct {
	value *ProjectFeatures
	isSet bool
}

func (v NullableProjectFeatures) Get() *ProjectFeatures {
	return v.value
}

func (v *NullableProjectFeatures) Set(val *ProjectFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectFeatures(val *ProjectFeatures) *NullableProjectFeatures {
	return &NullableProjectFeatures{value: val, isSet: true}
}

func (v NullableProjectFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
