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

// CloudResourceObject struct for CloudResourceObject
type CloudResourceObject struct {
	Type       string  `json:"type"`
	Name       *string `json:"name,omitempty"`
	Identifier string  `json:"identifier"`
}

// NewCloudResourceObject instantiates a new CloudResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudResourceObject(type_ string, identifier string) *CloudResourceObject {
	this := CloudResourceObject{}
	this.Type = type_
	this.Identifier = identifier
	return &this
}

// NewCloudResourceObjectWithDefaults instantiates a new CloudResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudResourceObjectWithDefaults() *CloudResourceObject {
	this := CloudResourceObject{}
	return &this
}

// GetType returns the Type field value
func (o *CloudResourceObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudResourceObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudResourceObject) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudResourceObject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResourceObject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudResourceObject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudResourceObject) SetName(v string) {
	o.Name = &v
}

// GetIdentifier returns the Identifier field value
func (o *CloudResourceObject) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *CloudResourceObject) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *CloudResourceObject) SetIdentifier(v string) {
	o.Identifier = v
}

func (o CloudResourceObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableCloudResourceObject struct {
	value *CloudResourceObject
	isSet bool
}

func (v NullableCloudResourceObject) Get() *CloudResourceObject {
	return v.value
}

func (v *NullableCloudResourceObject) Set(val *CloudResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudResourceObject(val *CloudResourceObject) *NullableCloudResourceObject {
	return &NullableCloudResourceObject{value: val, isSet: true}
}

func (v NullableCloudResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
