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

// MinimalCloud struct for MinimalCloud
type MinimalCloud struct {
	Id        *int32  `json:"id,omitempty"`
	Name      string  `json:"name"`
	State     *string `json:"state,omitempty"`
	CloudType *string `json:"cloudType,omitempty"`
}

// NewMinimalCloud instantiates a new MinimalCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalCloud(name string) *MinimalCloud {
	this := MinimalCloud{}
	this.Name = name
	return &this
}

// NewMinimalCloudWithDefaults instantiates a new MinimalCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalCloudWithDefaults() *MinimalCloud {
	this := MinimalCloud{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MinimalCloud) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalCloud) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MinimalCloud) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MinimalCloud) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *MinimalCloud) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MinimalCloud) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MinimalCloud) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MinimalCloud) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalCloud) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MinimalCloud) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MinimalCloud) SetState(v string) {
	o.State = &v
}

// GetCloudType returns the CloudType field value if set, zero value otherwise.
func (o *MinimalCloud) GetCloudType() string {
	if o == nil || o.CloudType == nil {
		var ret string
		return ret
	}
	return *o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalCloud) GetCloudTypeOk() (*string, bool) {
	if o == nil || o.CloudType == nil {
		return nil, false
	}
	return o.CloudType, true
}

// HasCloudType returns a boolean if a field has been set.
func (o *MinimalCloud) HasCloudType() bool {
	if o != nil && o.CloudType != nil {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given string and assigns it to the CloudType field.
func (o *MinimalCloud) SetCloudType(v string) {
	o.CloudType = &v
}

func (o MinimalCloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CloudType != nil {
		toSerialize["cloudType"] = o.CloudType
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalCloud struct {
	value *MinimalCloud
	isSet bool
}

func (v NullableMinimalCloud) Get() *MinimalCloud {
	return v.value
}

func (v *NullableMinimalCloud) Set(val *MinimalCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalCloud(val *MinimalCloud) *NullableMinimalCloud {
	return &NullableMinimalCloud{value: val, isSet: true}
}

func (v NullableMinimalCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
