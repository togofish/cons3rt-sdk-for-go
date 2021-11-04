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

// MinimalTemplateRegistration struct for MinimalTemplateRegistration
type MinimalTemplateRegistration struct {
	Id *int32 `json:"id,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty"`
	TemplateData *MinimalCons3rtTemplateData `json:"templateData,omitempty"`
	Offline *bool `json:"offline,omitempty"`
}

// NewMinimalTemplateRegistration instantiates a new MinimalTemplateRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalTemplateRegistration() *MinimalTemplateRegistration {
	this := MinimalTemplateRegistration{}
	return &this
}

// NewMinimalTemplateRegistrationWithDefaults instantiates a new MinimalTemplateRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalTemplateRegistrationWithDefaults() *MinimalTemplateRegistration {
	this := MinimalTemplateRegistration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MinimalTemplateRegistration) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateRegistration) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MinimalTemplateRegistration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MinimalTemplateRegistration) SetId(v int32) {
	o.Id = &v
}

// GetTemplateUuid returns the TemplateUuid field value if set, zero value otherwise.
func (o *MinimalTemplateRegistration) GetTemplateUuid() string {
	if o == nil || o.TemplateUuid == nil {
		var ret string
		return ret
	}
	return *o.TemplateUuid
}

// GetTemplateUuidOk returns a tuple with the TemplateUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateRegistration) GetTemplateUuidOk() (*string, bool) {
	if o == nil || o.TemplateUuid == nil {
		return nil, false
	}
	return o.TemplateUuid, true
}

// HasTemplateUuid returns a boolean if a field has been set.
func (o *MinimalTemplateRegistration) HasTemplateUuid() bool {
	if o != nil && o.TemplateUuid != nil {
		return true
	}

	return false
}

// SetTemplateUuid gets a reference to the given string and assigns it to the TemplateUuid field.
func (o *MinimalTemplateRegistration) SetTemplateUuid(v string) {
	o.TemplateUuid = &v
}

// GetTemplateData returns the TemplateData field value if set, zero value otherwise.
func (o *MinimalTemplateRegistration) GetTemplateData() MinimalCons3rtTemplateData {
	if o == nil || o.TemplateData == nil {
		var ret MinimalCons3rtTemplateData
		return ret
	}
	return *o.TemplateData
}

// GetTemplateDataOk returns a tuple with the TemplateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateRegistration) GetTemplateDataOk() (*MinimalCons3rtTemplateData, bool) {
	if o == nil || o.TemplateData == nil {
		return nil, false
	}
	return o.TemplateData, true
}

// HasTemplateData returns a boolean if a field has been set.
func (o *MinimalTemplateRegistration) HasTemplateData() bool {
	if o != nil && o.TemplateData != nil {
		return true
	}

	return false
}

// SetTemplateData gets a reference to the given MinimalCons3rtTemplateData and assigns it to the TemplateData field.
func (o *MinimalTemplateRegistration) SetTemplateData(v MinimalCons3rtTemplateData) {
	o.TemplateData = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *MinimalTemplateRegistration) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateRegistration) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *MinimalTemplateRegistration) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *MinimalTemplateRegistration) SetOffline(v bool) {
	o.Offline = &v
}

func (o MinimalTemplateRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TemplateUuid != nil {
		toSerialize["templateUuid"] = o.TemplateUuid
	}
	if o.TemplateData != nil {
		toSerialize["templateData"] = o.TemplateData
	}
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalTemplateRegistration struct {
	value *MinimalTemplateRegistration
	isSet bool
}

func (v NullableMinimalTemplateRegistration) Get() *MinimalTemplateRegistration {
	return v.value
}

func (v *NullableMinimalTemplateRegistration) Set(val *MinimalTemplateRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalTemplateRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalTemplateRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalTemplateRegistration(val *MinimalTemplateRegistration) *NullableMinimalTemplateRegistration {
	return &NullableMinimalTemplateRegistration{value: val, isSet: true}
}

func (v NullableMinimalTemplateRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalTemplateRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


