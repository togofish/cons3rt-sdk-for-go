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

// RestIdObject struct for RestIdObject
type RestIdObject struct {
	IdAsInteger *int32 `json:"idAsInteger,omitempty"`
	Id          string `json:"id"`
}

// NewRestIdObject instantiates a new RestIdObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestIdObject(id string) *RestIdObject {
	this := RestIdObject{}
	this.Id = id
	return &this
}

// NewRestIdObjectWithDefaults instantiates a new RestIdObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestIdObjectWithDefaults() *RestIdObject {
	this := RestIdObject{}
	return &this
}

// GetIdAsInteger returns the IdAsInteger field value if set, zero value otherwise.
func (o *RestIdObject) GetIdAsInteger() int32 {
	if o == nil || o.IdAsInteger == nil {
		var ret int32
		return ret
	}
	return *o.IdAsInteger
}

// GetIdAsIntegerOk returns a tuple with the IdAsInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestIdObject) GetIdAsIntegerOk() (*int32, bool) {
	if o == nil || o.IdAsInteger == nil {
		return nil, false
	}
	return o.IdAsInteger, true
}

// HasIdAsInteger returns a boolean if a field has been set.
func (o *RestIdObject) HasIdAsInteger() bool {
	if o != nil && o.IdAsInteger != nil {
		return true
	}

	return false
}

// SetIdAsInteger gets a reference to the given int32 and assigns it to the IdAsInteger field.
func (o *RestIdObject) SetIdAsInteger(v int32) {
	o.IdAsInteger = &v
}

// GetId returns the Id field value
func (o *RestIdObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RestIdObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RestIdObject) SetId(v string) {
	o.Id = v
}

func (o RestIdObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdAsInteger != nil {
		toSerialize["idAsInteger"] = o.IdAsInteger
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableRestIdObject struct {
	value *RestIdObject
	isSet bool
}

func (v NullableRestIdObject) Get() *RestIdObject {
	return v.value
}

func (v *NullableRestIdObject) Set(val *RestIdObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRestIdObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRestIdObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestIdObject(val *RestIdObject) *NullableRestIdObject {
	return &NullableRestIdObject{value: val, isSet: true}
}

func (v NullableRestIdObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestIdObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
