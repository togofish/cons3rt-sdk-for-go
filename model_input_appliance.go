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

// InputAppliance struct for InputAppliance
type InputAppliance struct {
	Name *string `json:"name,omitempty"`
	TemplateProfile *InputTemplateProfile `json:"templateProfile,omitempty"`
}

// NewInputAppliance instantiates a new InputAppliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputAppliance(subtype string) *InputAppliance {
	this := InputAppliance{}
	this.Subtype = subtype
	return &this
}

// NewInputApplianceWithDefaults instantiates a new InputAppliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputApplianceWithDefaults() *InputAppliance {
	this := InputAppliance{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputAppliance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAppliance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputAppliance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputAppliance) SetName(v string) {
	o.Name = &v
}

// GetTemplateProfile returns the TemplateProfile field value if set, zero value otherwise.
func (o *InputAppliance) GetTemplateProfile() InputTemplateProfile {
	if o == nil || o.TemplateProfile == nil {
		var ret InputTemplateProfile
		return ret
	}
	return *o.TemplateProfile
}

// GetTemplateProfileOk returns a tuple with the TemplateProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAppliance) GetTemplateProfileOk() (*InputTemplateProfile, bool) {
	if o == nil || o.TemplateProfile == nil {
		return nil, false
	}
	return o.TemplateProfile, true
}

// HasTemplateProfile returns a boolean if a field has been set.
func (o *InputAppliance) HasTemplateProfile() bool {
	if o != nil && o.TemplateProfile != nil {
		return true
	}

	return false
}

// SetTemplateProfile gets a reference to the given InputTemplateProfile and assigns it to the TemplateProfile field.
func (o *InputAppliance) SetTemplateProfile(v InputTemplateProfile) {
	o.TemplateProfile = &v
}

func (o InputAppliance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TemplateProfile != nil {
		toSerialize["templateProfile"] = o.TemplateProfile
	}
	return json.Marshal(toSerialize)
}

type NullableInputAppliance struct {
	value *InputAppliance
	isSet bool
}

func (v NullableInputAppliance) Get() *InputAppliance {
	return v.value
}

func (v *NullableInputAppliance) Set(val *InputAppliance) {
	v.value = val
	v.isSet = true
}

func (v NullableInputAppliance) IsSet() bool {
	return v.isSet
}

func (v *NullableInputAppliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputAppliance(val *InputAppliance) *NullableInputAppliance {
	return &NullableInputAppliance{value: val, isSet: true}
}

func (v NullableInputAppliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputAppliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


