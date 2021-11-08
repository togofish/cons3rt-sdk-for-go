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

// DockerRegistrySubmissionEndpointAllOf struct for DockerRegistrySubmissionEndpointAllOf
type DockerRegistrySubmissionEndpointAllOf struct {
	ImageName *string `json:"imageName,omitempty"`
	ImageTag  *string `json:"imageTag,omitempty"`
}

// NewDockerRegistrySubmissionEndpointAllOf instantiates a new DockerRegistrySubmissionEndpointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerRegistrySubmissionEndpointAllOf() *DockerRegistrySubmissionEndpointAllOf {
	this := DockerRegistrySubmissionEndpointAllOf{}
	return &this
}

// NewDockerRegistrySubmissionEndpointAllOfWithDefaults instantiates a new DockerRegistrySubmissionEndpointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerRegistrySubmissionEndpointAllOfWithDefaults() *DockerRegistrySubmissionEndpointAllOf {
	this := DockerRegistrySubmissionEndpointAllOf{}
	return &this
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *DockerRegistrySubmissionEndpointAllOf) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerRegistrySubmissionEndpointAllOf) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *DockerRegistrySubmissionEndpointAllOf) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *DockerRegistrySubmissionEndpointAllOf) SetImageName(v string) {
	o.ImageName = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *DockerRegistrySubmissionEndpointAllOf) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerRegistrySubmissionEndpointAllOf) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *DockerRegistrySubmissionEndpointAllOf) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *DockerRegistrySubmissionEndpointAllOf) SetImageTag(v string) {
	o.ImageTag = &v
}

func (o DockerRegistrySubmissionEndpointAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageName != nil {
		toSerialize["imageName"] = o.ImageName
	}
	if o.ImageTag != nil {
		toSerialize["imageTag"] = o.ImageTag
	}
	return json.Marshal(toSerialize)
}

type NullableDockerRegistrySubmissionEndpointAllOf struct {
	value *DockerRegistrySubmissionEndpointAllOf
	isSet bool
}

func (v NullableDockerRegistrySubmissionEndpointAllOf) Get() *DockerRegistrySubmissionEndpointAllOf {
	return v.value
}

func (v *NullableDockerRegistrySubmissionEndpointAllOf) Set(val *DockerRegistrySubmissionEndpointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerRegistrySubmissionEndpointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerRegistrySubmissionEndpointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerRegistrySubmissionEndpointAllOf(val *DockerRegistrySubmissionEndpointAllOf) *NullableDockerRegistrySubmissionEndpointAllOf {
	return &NullableDockerRegistrySubmissionEndpointAllOf{value: val, isSet: true}
}

func (v NullableDockerRegistrySubmissionEndpointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerRegistrySubmissionEndpointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
