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

// FullApplianceAllOf struct for FullApplianceAllOf
type FullApplianceAllOf struct {
	Components      *[]MinimalAbstractComponent `json:"components,omitempty"`
	TemplateProfile *MinimalTemplateProfile     `json:"templateProfile,omitempty"`
}

// NewFullApplianceAllOf instantiates a new FullApplianceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullApplianceAllOf() *FullApplianceAllOf {
	this := FullApplianceAllOf{}
	return &this
}

// NewFullApplianceAllOfWithDefaults instantiates a new FullApplianceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullApplianceAllOfWithDefaults() *FullApplianceAllOf {
	this := FullApplianceAllOf{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *FullApplianceAllOf) GetComponents() []MinimalAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []MinimalAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplianceAllOf) GetComponentsOk() (*[]MinimalAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *FullApplianceAllOf) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []MinimalAbstractComponent and assigns it to the Components field.
func (o *FullApplianceAllOf) SetComponents(v []MinimalAbstractComponent) {
	o.Components = &v
}

// GetTemplateProfile returns the TemplateProfile field value if set, zero value otherwise.
func (o *FullApplianceAllOf) GetTemplateProfile() MinimalTemplateProfile {
	if o == nil || o.TemplateProfile == nil {
		var ret MinimalTemplateProfile
		return ret
	}
	return *o.TemplateProfile
}

// GetTemplateProfileOk returns a tuple with the TemplateProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplianceAllOf) GetTemplateProfileOk() (*MinimalTemplateProfile, bool) {
	if o == nil || o.TemplateProfile == nil {
		return nil, false
	}
	return o.TemplateProfile, true
}

// HasTemplateProfile returns a boolean if a field has been set.
func (o *FullApplianceAllOf) HasTemplateProfile() bool {
	if o != nil && o.TemplateProfile != nil {
		return true
	}

	return false
}

// SetTemplateProfile gets a reference to the given MinimalTemplateProfile and assigns it to the TemplateProfile field.
func (o *FullApplianceAllOf) SetTemplateProfile(v MinimalTemplateProfile) {
	o.TemplateProfile = &v
}

func (o FullApplianceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.TemplateProfile != nil {
		toSerialize["templateProfile"] = o.TemplateProfile
	}
	return json.Marshal(toSerialize)
}

type NullableFullApplianceAllOf struct {
	value *FullApplianceAllOf
	isSet bool
}

func (v NullableFullApplianceAllOf) Get() *FullApplianceAllOf {
	return v.value
}

func (v *NullableFullApplianceAllOf) Set(val *FullApplianceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullApplianceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullApplianceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullApplianceAllOf(val *FullApplianceAllOf) *NullableFullApplianceAllOf {
	return &NullableFullApplianceAllOf{value: val, isSet: true}
}

func (v NullableFullApplianceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullApplianceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
