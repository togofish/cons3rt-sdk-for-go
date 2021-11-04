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

// InputVirtualHostFull struct for InputVirtualHostFull
type InputVirtualHostFull struct {
	Id              *int32                    `json:"id,omitempty"`
	TemplateProfile *InputTemplateProfile     `json:"templateProfile,omitempty"`
	Components      *[]InputAbstractComponent `json:"components,omitempty"`
	Name            *string                   `json:"name,omitempty"`
	Subtype         string
}

// NewInputVirtualHostFull instantiates a new InputVirtualHostFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVirtualHostFull(subtype string) *InputVirtualHostFull {
	this := InputVirtualHostFull{}
	this.Subtype = subtype
	return &this
}

// NewInputVirtualHostFullWithDefaults instantiates a new InputVirtualHostFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVirtualHostFullWithDefaults() *InputVirtualHostFull {
	this := InputVirtualHostFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InputVirtualHostFull) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualHostFull) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InputVirtualHostFull) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InputVirtualHostFull) SetId(v int32) {
	o.Id = &v
}

// GetTemplateProfile returns the TemplateProfile field value if set, zero value otherwise.
func (o *InputVirtualHostFull) GetTemplateProfile() InputTemplateProfile {
	if o == nil || o.TemplateProfile == nil {
		var ret InputTemplateProfile
		return ret
	}
	return *o.TemplateProfile
}

// GetTemplateProfileOk returns a tuple with the TemplateProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualHostFull) GetTemplateProfileOk() (*InputTemplateProfile, bool) {
	if o == nil || o.TemplateProfile == nil {
		return nil, false
	}
	return o.TemplateProfile, true
}

// HasTemplateProfile returns a boolean if a field has been set.
func (o *InputVirtualHostFull) HasTemplateProfile() bool {
	if o != nil && o.TemplateProfile != nil {
		return true
	}

	return false
}

// SetTemplateProfile gets a reference to the given InputTemplateProfile and assigns it to the TemplateProfile field.
func (o *InputVirtualHostFull) SetTemplateProfile(v InputTemplateProfile) {
	o.TemplateProfile = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *InputVirtualHostFull) GetComponents() []InputAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []InputAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualHostFull) GetComponentsOk() (*[]InputAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *InputVirtualHostFull) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []InputAbstractComponent and assigns it to the Components field.
func (o *InputVirtualHostFull) SetComponents(v []InputAbstractComponent) {
	o.Components = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputVirtualHostFull) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualHostFull) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputVirtualHostFull) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputVirtualHostFull) SetName(v string) {
	o.Name = &v
}

func (o InputVirtualHostFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TemplateProfile != nil {
		toSerialize["templateProfile"] = o.TemplateProfile
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInputVirtualHostFull struct {
	value *InputVirtualHostFull
	isSet bool
}

func (v NullableInputVirtualHostFull) Get() *InputVirtualHostFull {
	return v.value
}

func (v *NullableInputVirtualHostFull) Set(val *InputVirtualHostFull) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVirtualHostFull) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVirtualHostFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVirtualHostFull(val *InputVirtualHostFull) *NullableInputVirtualHostFull {
	return &NullableInputVirtualHostFull{value: val, isSet: true}
}

func (v NullableInputVirtualHostFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVirtualHostFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
