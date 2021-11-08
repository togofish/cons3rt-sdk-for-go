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

// InputSubmissionEndpointForProject struct for InputSubmissionEndpointForProject
type InputSubmissionEndpointForProject struct {
	Host    string `json:"host"`
	Port    *int32 `json:"port,omitempty"`
	Subtype string `json:"subtype"`
}

// NewInputSubmissionEndpointForProject instantiates a new InputSubmissionEndpointForProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputSubmissionEndpointForProject(host string, subtype string) *InputSubmissionEndpointForProject {
	this := InputSubmissionEndpointForProject{}
	this.Host = host
	this.Subtype = subtype
	return &this
}

// NewInputSubmissionEndpointForProjectWithDefaults instantiates a new InputSubmissionEndpointForProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputSubmissionEndpointForProjectWithDefaults() *InputSubmissionEndpointForProject {
	this := InputSubmissionEndpointForProject{}
	return &this
}

// GetHost returns the Host field value
func (o *InputSubmissionEndpointForProject) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *InputSubmissionEndpointForProject) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *InputSubmissionEndpointForProject) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InputSubmissionEndpointForProject) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSubmissionEndpointForProject) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InputSubmissionEndpointForProject) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InputSubmissionEndpointForProject) SetPort(v int32) {
	o.Port = &v
}

// GetSubtype returns the Subtype field value
func (o *InputSubmissionEndpointForProject) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *InputSubmissionEndpointForProject) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *InputSubmissionEndpointForProject) SetSubtype(v string) {
	o.Subtype = v
}

func (o InputSubmissionEndpointForProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableInputSubmissionEndpointForProject struct {
	value *InputSubmissionEndpointForProject
	isSet bool
}

func (v NullableInputSubmissionEndpointForProject) Get() *InputSubmissionEndpointForProject {
	return v.value
}

func (v *NullableInputSubmissionEndpointForProject) Set(val *InputSubmissionEndpointForProject) {
	v.value = val
	v.isSet = true
}

func (v NullableInputSubmissionEndpointForProject) IsSet() bool {
	return v.isSet
}

func (v *NullableInputSubmissionEndpointForProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputSubmissionEndpointForProject(val *InputSubmissionEndpointForProject) *NullableInputSubmissionEndpointForProject {
	return &NullableInputSubmissionEndpointForProject{value: val, isSet: true}
}

func (v NullableInputSubmissionEndpointForProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputSubmissionEndpointForProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
