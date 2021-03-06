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

// ContainerComponentAllOf struct for ContainerComponentAllOf
type ContainerComponentAllOf struct {
	Configuration *ContainerConfiguration `json:"configuration,omitempty"`
}

// NewContainerComponentAllOf instantiates a new ContainerComponentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerComponentAllOf() *ContainerComponentAllOf {
	this := ContainerComponentAllOf{}
	return &this
}

// NewContainerComponentAllOfWithDefaults instantiates a new ContainerComponentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerComponentAllOfWithDefaults() *ContainerComponentAllOf {
	this := ContainerComponentAllOf{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ContainerComponentAllOf) GetConfiguration() ContainerConfiguration {
	if o == nil || o.Configuration == nil {
		var ret ContainerConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerComponentAllOf) GetConfigurationOk() (*ContainerConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ContainerComponentAllOf) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given ContainerConfiguration and assigns it to the Configuration field.
func (o *ContainerComponentAllOf) SetConfiguration(v ContainerConfiguration) {
	o.Configuration = &v
}

func (o ContainerComponentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableContainerComponentAllOf struct {
	value *ContainerComponentAllOf
	isSet bool
}

func (v NullableContainerComponentAllOf) Get() *ContainerComponentAllOf {
	return v.value
}

func (v *NullableContainerComponentAllOf) Set(val *ContainerComponentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerComponentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerComponentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerComponentAllOf(val *ContainerComponentAllOf) *NullableContainerComponentAllOf {
	return &NullableContainerComponentAllOf{value: val, isSet: true}
}

func (v NullableContainerComponentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerComponentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
