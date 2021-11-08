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

// SubmissionEndpoint struct for SubmissionEndpoint
type SubmissionEndpoint struct {
	Host    string  `json:"host"`
	Id      *int32  `json:"id,omitempty"`
	Port    *int32  `json:"port,omitempty"`
	Type    *string `json:"type,omitempty"`
	Subtype string  `json:"subtype"`
}

// NewSubmissionEndpoint instantiates a new SubmissionEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionEndpoint(host string, subtype string) *SubmissionEndpoint {
	this := SubmissionEndpoint{}
	this.Host = host
	this.Subtype = subtype
	return &this
}

// NewSubmissionEndpointWithDefaults instantiates a new SubmissionEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionEndpointWithDefaults() *SubmissionEndpoint {
	this := SubmissionEndpoint{}
	return &this
}

// GetHost returns the Host field value
func (o *SubmissionEndpoint) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SubmissionEndpoint) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *SubmissionEndpoint) SetHost(v string) {
	o.Host = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubmissionEndpoint) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionEndpoint) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubmissionEndpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SubmissionEndpoint) SetId(v int32) {
	o.Id = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SubmissionEndpoint) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionEndpoint) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SubmissionEndpoint) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SubmissionEndpoint) SetPort(v int32) {
	o.Port = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubmissionEndpoint) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubmissionEndpoint) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SubmissionEndpoint) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value
func (o *SubmissionEndpoint) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *SubmissionEndpoint) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *SubmissionEndpoint) SetSubtype(v string) {
	o.Subtype = v
}

func (o SubmissionEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableSubmissionEndpoint struct {
	value *SubmissionEndpoint
	isSet bool
}

func (v NullableSubmissionEndpoint) Get() *SubmissionEndpoint {
	return v.value
}

func (v *NullableSubmissionEndpoint) Set(val *SubmissionEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionEndpoint(val *SubmissionEndpoint) *NullableSubmissionEndpoint {
	return &NullableSubmissionEndpoint{value: val, isSet: true}
}

func (v NullableSubmissionEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
