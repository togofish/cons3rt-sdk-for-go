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

// FullScenarioAllOf struct for FullScenarioAllOf
type FullScenarioAllOf struct {
	ScenarioHosts *[]MinimalScenarioHost `json:"scenarioHosts,omitempty"`
	PrepareScenarioConfiguration *MinimalConfiguration `json:"prepareScenarioConfiguration,omitempty"`
	TeardownScenarioConfiguration *MinimalConfiguration `json:"teardownScenarioConfiguration,omitempty"`
}

// NewFullScenarioAllOf instantiates a new FullScenarioAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullScenarioAllOf() *FullScenarioAllOf {
	this := FullScenarioAllOf{}
	return &this
}

// NewFullScenarioAllOfWithDefaults instantiates a new FullScenarioAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullScenarioAllOfWithDefaults() *FullScenarioAllOf {
	this := FullScenarioAllOf{}
	return &this
}

// GetScenarioHosts returns the ScenarioHosts field value if set, zero value otherwise.
func (o *FullScenarioAllOf) GetScenarioHosts() []MinimalScenarioHost {
	if o == nil || o.ScenarioHosts == nil {
		var ret []MinimalScenarioHost
		return ret
	}
	return *o.ScenarioHosts
}

// GetScenarioHostsOk returns a tuple with the ScenarioHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullScenarioAllOf) GetScenarioHostsOk() (*[]MinimalScenarioHost, bool) {
	if o == nil || o.ScenarioHosts == nil {
		return nil, false
	}
	return o.ScenarioHosts, true
}

// HasScenarioHosts returns a boolean if a field has been set.
func (o *FullScenarioAllOf) HasScenarioHosts() bool {
	if o != nil && o.ScenarioHosts != nil {
		return true
	}

	return false
}

// SetScenarioHosts gets a reference to the given []MinimalScenarioHost and assigns it to the ScenarioHosts field.
func (o *FullScenarioAllOf) SetScenarioHosts(v []MinimalScenarioHost) {
	o.ScenarioHosts = &v
}

// GetPrepareScenarioConfiguration returns the PrepareScenarioConfiguration field value if set, zero value otherwise.
func (o *FullScenarioAllOf) GetPrepareScenarioConfiguration() MinimalConfiguration {
	if o == nil || o.PrepareScenarioConfiguration == nil {
		var ret MinimalConfiguration
		return ret
	}
	return *o.PrepareScenarioConfiguration
}

// GetPrepareScenarioConfigurationOk returns a tuple with the PrepareScenarioConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullScenarioAllOf) GetPrepareScenarioConfigurationOk() (*MinimalConfiguration, bool) {
	if o == nil || o.PrepareScenarioConfiguration == nil {
		return nil, false
	}
	return o.PrepareScenarioConfiguration, true
}

// HasPrepareScenarioConfiguration returns a boolean if a field has been set.
func (o *FullScenarioAllOf) HasPrepareScenarioConfiguration() bool {
	if o != nil && o.PrepareScenarioConfiguration != nil {
		return true
	}

	return false
}

// SetPrepareScenarioConfiguration gets a reference to the given MinimalConfiguration and assigns it to the PrepareScenarioConfiguration field.
func (o *FullScenarioAllOf) SetPrepareScenarioConfiguration(v MinimalConfiguration) {
	o.PrepareScenarioConfiguration = &v
}

// GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field value if set, zero value otherwise.
func (o *FullScenarioAllOf) GetTeardownScenarioConfiguration() MinimalConfiguration {
	if o == nil || o.TeardownScenarioConfiguration == nil {
		var ret MinimalConfiguration
		return ret
	}
	return *o.TeardownScenarioConfiguration
}

// GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullScenarioAllOf) GetTeardownScenarioConfigurationOk() (*MinimalConfiguration, bool) {
	if o == nil || o.TeardownScenarioConfiguration == nil {
		return nil, false
	}
	return o.TeardownScenarioConfiguration, true
}

// HasTeardownScenarioConfiguration returns a boolean if a field has been set.
func (o *FullScenarioAllOf) HasTeardownScenarioConfiguration() bool {
	if o != nil && o.TeardownScenarioConfiguration != nil {
		return true
	}

	return false
}

// SetTeardownScenarioConfiguration gets a reference to the given MinimalConfiguration and assigns it to the TeardownScenarioConfiguration field.
func (o *FullScenarioAllOf) SetTeardownScenarioConfiguration(v MinimalConfiguration) {
	o.TeardownScenarioConfiguration = &v
}

func (o FullScenarioAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScenarioHosts != nil {
		toSerialize["scenarioHosts"] = o.ScenarioHosts
	}
	if o.PrepareScenarioConfiguration != nil {
		toSerialize["prepareScenarioConfiguration"] = o.PrepareScenarioConfiguration
	}
	if o.TeardownScenarioConfiguration != nil {
		toSerialize["teardownScenarioConfiguration"] = o.TeardownScenarioConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableFullScenarioAllOf struct {
	value *FullScenarioAllOf
	isSet bool
}

func (v NullableFullScenarioAllOf) Get() *FullScenarioAllOf {
	return v.value
}

func (v *NullableFullScenarioAllOf) Set(val *FullScenarioAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullScenarioAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullScenarioAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullScenarioAllOf(val *FullScenarioAllOf) *NullableFullScenarioAllOf {
	return &NullableFullScenarioAllOf{value: val, isSet: true}
}

func (v NullableFullScenarioAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullScenarioAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


