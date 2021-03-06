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

// FullSoftwareAssetBundleAllOf struct for FullSoftwareAssetBundleAllOf
type FullSoftwareAssetBundleAllOf struct {
	SoftwareComponents *[]MinimalSoftwareComponent `json:"softwareComponents,omitempty"`
}

// NewFullSoftwareAssetBundleAllOf instantiates a new FullSoftwareAssetBundleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullSoftwareAssetBundleAllOf() *FullSoftwareAssetBundleAllOf {
	this := FullSoftwareAssetBundleAllOf{}
	return &this
}

// NewFullSoftwareAssetBundleAllOfWithDefaults instantiates a new FullSoftwareAssetBundleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullSoftwareAssetBundleAllOfWithDefaults() *FullSoftwareAssetBundleAllOf {
	this := FullSoftwareAssetBundleAllOf{}
	return &this
}

// GetSoftwareComponents returns the SoftwareComponents field value if set, zero value otherwise.
func (o *FullSoftwareAssetBundleAllOf) GetSoftwareComponents() []MinimalSoftwareComponent {
	if o == nil || o.SoftwareComponents == nil {
		var ret []MinimalSoftwareComponent
		return ret
	}
	return *o.SoftwareComponents
}

// GetSoftwareComponentsOk returns a tuple with the SoftwareComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetBundleAllOf) GetSoftwareComponentsOk() (*[]MinimalSoftwareComponent, bool) {
	if o == nil || o.SoftwareComponents == nil {
		return nil, false
	}
	return o.SoftwareComponents, true
}

// HasSoftwareComponents returns a boolean if a field has been set.
func (o *FullSoftwareAssetBundleAllOf) HasSoftwareComponents() bool {
	if o != nil && o.SoftwareComponents != nil {
		return true
	}

	return false
}

// SetSoftwareComponents gets a reference to the given []MinimalSoftwareComponent and assigns it to the SoftwareComponents field.
func (o *FullSoftwareAssetBundleAllOf) SetSoftwareComponents(v []MinimalSoftwareComponent) {
	o.SoftwareComponents = &v
}

func (o FullSoftwareAssetBundleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SoftwareComponents != nil {
		toSerialize["softwareComponents"] = o.SoftwareComponents
	}
	return json.Marshal(toSerialize)
}

type NullableFullSoftwareAssetBundleAllOf struct {
	value *FullSoftwareAssetBundleAllOf
	isSet bool
}

func (v NullableFullSoftwareAssetBundleAllOf) Get() *FullSoftwareAssetBundleAllOf {
	return v.value
}

func (v *NullableFullSoftwareAssetBundleAllOf) Set(val *FullSoftwareAssetBundleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullSoftwareAssetBundleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullSoftwareAssetBundleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullSoftwareAssetBundleAllOf(val *FullSoftwareAssetBundleAllOf) *NullableFullSoftwareAssetBundleAllOf {
	return &NullableFullSoftwareAssetBundleAllOf{value: val, isSet: true}
}

func (v NullableFullSoftwareAssetBundleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullSoftwareAssetBundleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
