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

// RemoteAccessSession struct for RemoteAccessSession
type RemoteAccessSession struct {
	EndDate *int32 `json:"endDate,omitempty"`
	Id *int32 `json:"id,omitempty"`
	LowBandwidth *bool `json:"lowBandwidth,omitempty"`
	StartDate *int32 `json:"startDate,omitempty"`
	Username *string `json:"username,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewRemoteAccessSession instantiates a new RemoteAccessSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteAccessSession() *RemoteAccessSession {
	this := RemoteAccessSession{}
	return &this
}

// NewRemoteAccessSessionWithDefaults instantiates a new RemoteAccessSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteAccessSessionWithDefaults() *RemoteAccessSession {
	this := RemoteAccessSession{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *RemoteAccessSession) GetEndDate() int32 {
	if o == nil || o.EndDate == nil {
		var ret int32
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessSession) GetEndDateOk() (*int32, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *RemoteAccessSession) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int32 and assigns it to the EndDate field.
func (o *RemoteAccessSession) SetEndDate(v int32) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteAccessSession) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessSession) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteAccessSession) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RemoteAccessSession) SetId(v int32) {
	o.Id = &v
}

// GetLowBandwidth returns the LowBandwidth field value if set, zero value otherwise.
func (o *RemoteAccessSession) GetLowBandwidth() bool {
	if o == nil || o.LowBandwidth == nil {
		var ret bool
		return ret
	}
	return *o.LowBandwidth
}

// GetLowBandwidthOk returns a tuple with the LowBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessSession) GetLowBandwidthOk() (*bool, bool) {
	if o == nil || o.LowBandwidth == nil {
		return nil, false
	}
	return o.LowBandwidth, true
}

// HasLowBandwidth returns a boolean if a field has been set.
func (o *RemoteAccessSession) HasLowBandwidth() bool {
	if o != nil && o.LowBandwidth != nil {
		return true
	}

	return false
}

// SetLowBandwidth gets a reference to the given bool and assigns it to the LowBandwidth field.
func (o *RemoteAccessSession) SetLowBandwidth(v bool) {
	o.LowBandwidth = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *RemoteAccessSession) GetStartDate() int32 {
	if o == nil || o.StartDate == nil {
		var ret int32
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessSession) GetStartDateOk() (*int32, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *RemoteAccessSession) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int32 and assigns it to the StartDate field.
func (o *RemoteAccessSession) SetStartDate(v int32) {
	o.StartDate = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RemoteAccessSession) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessSession) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RemoteAccessSession) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RemoteAccessSession) SetUsername(v string) {
	o.Username = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RemoteAccessSession) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessSession) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RemoteAccessSession) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RemoteAccessSession) SetType(v string) {
	o.Type = &v
}

func (o RemoteAccessSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LowBandwidth != nil {
		toSerialize["lowBandwidth"] = o.LowBandwidth
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteAccessSession struct {
	value *RemoteAccessSession
	isSet bool
}

func (v NullableRemoteAccessSession) Get() *RemoteAccessSession {
	return v.value
}

func (v *NullableRemoteAccessSession) Set(val *RemoteAccessSession) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteAccessSession) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteAccessSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteAccessSession(val *RemoteAccessSession) *NullableRemoteAccessSession {
	return &NullableRemoteAccessSession{value: val, isSet: true}
}

func (v NullableRemoteAccessSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteAccessSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


