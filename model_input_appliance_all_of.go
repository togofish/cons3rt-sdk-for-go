/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InputApplianceAllOf struct for InputApplianceAllOf
type InputApplianceAllOf struct {
	Name            *string               `json:"name,omitempty"`
	TemplateProfile *InputTemplateProfile `json:"templateProfile,omitempty"`
}

// NewInputApplianceAllOf instantiates a new InputApplianceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputApplianceAllOf() *InputApplianceAllOf {
	this := InputApplianceAllOf{}
	return &this
}

// NewInputApplianceAllOfWithDefaults instantiates a new InputApplianceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputApplianceAllOfWithDefaults() *InputApplianceAllOf {
	this := InputApplianceAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputApplianceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplianceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputApplianceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputApplianceAllOf) SetName(v string) {
	o.Name = &v
}

// GetTemplateProfile returns the TemplateProfile field value if set, zero value otherwise.
func (o *InputApplianceAllOf) GetTemplateProfile() InputTemplateProfile {
	if o == nil || o.TemplateProfile == nil {
		var ret InputTemplateProfile
		return ret
	}
	return *o.TemplateProfile
}

// GetTemplateProfileOk returns a tuple with the TemplateProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplianceAllOf) GetTemplateProfileOk() (*InputTemplateProfile, bool) {
	if o == nil || o.TemplateProfile == nil {
		return nil, false
	}
	return o.TemplateProfile, true
}

// HasTemplateProfile returns a boolean if a field has been set.
func (o *InputApplianceAllOf) HasTemplateProfile() bool {
	if o != nil && o.TemplateProfile != nil {
		return true
	}

	return false
}

// SetTemplateProfile gets a reference to the given InputTemplateProfile and assigns it to the TemplateProfile field.
func (o *InputApplianceAllOf) SetTemplateProfile(v InputTemplateProfile) {
	o.TemplateProfile = &v
}

func (o InputApplianceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TemplateProfile != nil {
		toSerialize["templateProfile"] = o.TemplateProfile
	}
	return json.Marshal(toSerialize)
}

type NullableInputApplianceAllOf struct {
	value *InputApplianceAllOf
	isSet bool
}

func (v NullableInputApplianceAllOf) Get() *InputApplianceAllOf {
	return v.value
}

func (v *NullableInputApplianceAllOf) Set(val *InputApplianceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInputApplianceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInputApplianceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputApplianceAllOf(val *InputApplianceAllOf) *NullableInputApplianceAllOf {
	return &NullableInputApplianceAllOf{value: val, isSet: true}
}

func (v NullableInputApplianceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputApplianceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
