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

// MinimalCategory struct for MinimalCategory
type MinimalCategory struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Disruptive *bool `json:"disruptive,omitempty"`
}

// NewMinimalCategory instantiates a new MinimalCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalCategory() *MinimalCategory {
	this := MinimalCategory{}
	return &this
}

// NewMinimalCategoryWithDefaults instantiates a new MinimalCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalCategoryWithDefaults() *MinimalCategory {
	this := MinimalCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MinimalCategory) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalCategory) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MinimalCategory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MinimalCategory) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MinimalCategory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalCategory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MinimalCategory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MinimalCategory) SetName(v string) {
	o.Name = &v
}

// GetDisruptive returns the Disruptive field value if set, zero value otherwise.
func (o *MinimalCategory) GetDisruptive() bool {
	if o == nil || o.Disruptive == nil {
		var ret bool
		return ret
	}
	return *o.Disruptive
}

// GetDisruptiveOk returns a tuple with the Disruptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalCategory) GetDisruptiveOk() (*bool, bool) {
	if o == nil || o.Disruptive == nil {
		return nil, false
	}
	return o.Disruptive, true
}

// HasDisruptive returns a boolean if a field has been set.
func (o *MinimalCategory) HasDisruptive() bool {
	if o != nil && o.Disruptive != nil {
		return true
	}

	return false
}

// SetDisruptive gets a reference to the given bool and assigns it to the Disruptive field.
func (o *MinimalCategory) SetDisruptive(v bool) {
	o.Disruptive = &v
}

func (o MinimalCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Disruptive != nil {
		toSerialize["disruptive"] = o.Disruptive
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalCategory struct {
	value *MinimalCategory
	isSet bool
}

func (v NullableMinimalCategory) Get() *MinimalCategory {
	return v.value
}

func (v *NullableMinimalCategory) Set(val *MinimalCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalCategory(val *MinimalCategory) *NullableMinimalCategory {
	return &NullableMinimalCategory{value: val, isSet: true}
}

func (v NullableMinimalCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


