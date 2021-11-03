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

// InputDockerRegistrySubmissionEndpointForProject struct for InputDockerRegistrySubmissionEndpointForProject
type InputDockerRegistrySubmissionEndpointForProject struct {
}

// NewInputDockerRegistrySubmissionEndpointForProject instantiates a new InputDockerRegistrySubmissionEndpointForProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputDockerRegistrySubmissionEndpointForProject(host string, subtype string) *InputDockerRegistrySubmissionEndpointForProject {
	this := InputDockerRegistrySubmissionEndpointForProject{}
	this.Host = host
	this.Subtype = subtype
	return &this
}

// NewInputDockerRegistrySubmissionEndpointForProjectWithDefaults instantiates a new InputDockerRegistrySubmissionEndpointForProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputDockerRegistrySubmissionEndpointForProjectWithDefaults() *InputDockerRegistrySubmissionEndpointForProject {
	this := InputDockerRegistrySubmissionEndpointForProject{}
	return &this
}

func (o InputDockerRegistrySubmissionEndpointForProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableInputDockerRegistrySubmissionEndpointForProject struct {
	value *InputDockerRegistrySubmissionEndpointForProject
	isSet bool
}

func (v NullableInputDockerRegistrySubmissionEndpointForProject) Get() *InputDockerRegistrySubmissionEndpointForProject {
	return v.value
}

func (v *NullableInputDockerRegistrySubmissionEndpointForProject) Set(val *InputDockerRegistrySubmissionEndpointForProject) {
	v.value = val
	v.isSet = true
}

func (v NullableInputDockerRegistrySubmissionEndpointForProject) IsSet() bool {
	return v.isSet
}

func (v *NullableInputDockerRegistrySubmissionEndpointForProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputDockerRegistrySubmissionEndpointForProject(val *InputDockerRegistrySubmissionEndpointForProject) *NullableInputDockerRegistrySubmissionEndpointForProject {
	return &NullableInputDockerRegistrySubmissionEndpointForProject{value: val, isSet: true}
}

func (v NullableInputDockerRegistrySubmissionEndpointForProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputDockerRegistrySubmissionEndpointForProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


