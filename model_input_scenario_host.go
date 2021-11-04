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

// InputScenarioHost struct for InputScenarioHost
type InputScenarioHost struct {
	BuildOrder *int32 `json:"buildOrder,omitempty"`
	SystemRole string `json:"systemRole"`
	SystemModule *InputSystemModuleFull `json:"systemModule,omitempty"`
	Master *bool `json:"master,omitempty"`
	ConfigureScenarioConfiguration *InputConfiguration `json:"configureScenarioConfiguration,omitempty"`
	TeardownScenarioConfiguration *InputConfiguration `json:"teardownScenarioConfiguration,omitempty"`
}

// NewInputScenarioHost instantiates a new InputScenarioHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputScenarioHost(systemRole string) *InputScenarioHost {
	this := InputScenarioHost{}
	this.SystemRole = systemRole
	return &this
}

// NewInputScenarioHostWithDefaults instantiates a new InputScenarioHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputScenarioHostWithDefaults() *InputScenarioHost {
	this := InputScenarioHost{}
	return &this
}

// GetBuildOrder returns the BuildOrder field value if set, zero value otherwise.
func (o *InputScenarioHost) GetBuildOrder() int32 {
	if o == nil || o.BuildOrder == nil {
		var ret int32
		return ret
	}
	return *o.BuildOrder
}

// GetBuildOrderOk returns a tuple with the BuildOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputScenarioHost) GetBuildOrderOk() (*int32, bool) {
	if o == nil || o.BuildOrder == nil {
		return nil, false
	}
	return o.BuildOrder, true
}

// HasBuildOrder returns a boolean if a field has been set.
func (o *InputScenarioHost) HasBuildOrder() bool {
	if o != nil && o.BuildOrder != nil {
		return true
	}

	return false
}

// SetBuildOrder gets a reference to the given int32 and assigns it to the BuildOrder field.
func (o *InputScenarioHost) SetBuildOrder(v int32) {
	o.BuildOrder = &v
}

// GetSystemRole returns the SystemRole field value
func (o *InputScenarioHost) GetSystemRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value
// and a boolean to check if the value has been set.
func (o *InputScenarioHost) GetSystemRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemRole, true
}

// SetSystemRole sets field value
func (o *InputScenarioHost) SetSystemRole(v string) {
	o.SystemRole = v
}

// GetSystemModule returns the SystemModule field value if set, zero value otherwise.
func (o *InputScenarioHost) GetSystemModule() InputSystemModuleFull {
	if o == nil || o.SystemModule == nil {
		var ret InputSystemModuleFull
		return ret
	}
	return *o.SystemModule
}

// GetSystemModuleOk returns a tuple with the SystemModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputScenarioHost) GetSystemModuleOk() (*InputSystemModuleFull, bool) {
	if o == nil || o.SystemModule == nil {
		return nil, false
	}
	return o.SystemModule, true
}

// HasSystemModule returns a boolean if a field has been set.
func (o *InputScenarioHost) HasSystemModule() bool {
	if o != nil && o.SystemModule != nil {
		return true
	}

	return false
}

// SetSystemModule gets a reference to the given InputSystemModuleFull and assigns it to the SystemModule field.
func (o *InputScenarioHost) SetSystemModule(v InputSystemModuleFull) {
	o.SystemModule = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *InputScenarioHost) GetMaster() bool {
	if o == nil || o.Master == nil {
		var ret bool
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputScenarioHost) GetMasterOk() (*bool, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *InputScenarioHost) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given bool and assigns it to the Master field.
func (o *InputScenarioHost) SetMaster(v bool) {
	o.Master = &v
}

// GetConfigureScenarioConfiguration returns the ConfigureScenarioConfiguration field value if set, zero value otherwise.
func (o *InputScenarioHost) GetConfigureScenarioConfiguration() InputConfiguration {
	if o == nil || o.ConfigureScenarioConfiguration == nil {
		var ret InputConfiguration
		return ret
	}
	return *o.ConfigureScenarioConfiguration
}

// GetConfigureScenarioConfigurationOk returns a tuple with the ConfigureScenarioConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputScenarioHost) GetConfigureScenarioConfigurationOk() (*InputConfiguration, bool) {
	if o == nil || o.ConfigureScenarioConfiguration == nil {
		return nil, false
	}
	return o.ConfigureScenarioConfiguration, true
}

// HasConfigureScenarioConfiguration returns a boolean if a field has been set.
func (o *InputScenarioHost) HasConfigureScenarioConfiguration() bool {
	if o != nil && o.ConfigureScenarioConfiguration != nil {
		return true
	}

	return false
}

// SetConfigureScenarioConfiguration gets a reference to the given InputConfiguration and assigns it to the ConfigureScenarioConfiguration field.
func (o *InputScenarioHost) SetConfigureScenarioConfiguration(v InputConfiguration) {
	o.ConfigureScenarioConfiguration = &v
}

// GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field value if set, zero value otherwise.
func (o *InputScenarioHost) GetTeardownScenarioConfiguration() InputConfiguration {
	if o == nil || o.TeardownScenarioConfiguration == nil {
		var ret InputConfiguration
		return ret
	}
	return *o.TeardownScenarioConfiguration
}

// GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputScenarioHost) GetTeardownScenarioConfigurationOk() (*InputConfiguration, bool) {
	if o == nil || o.TeardownScenarioConfiguration == nil {
		return nil, false
	}
	return o.TeardownScenarioConfiguration, true
}

// HasTeardownScenarioConfiguration returns a boolean if a field has been set.
func (o *InputScenarioHost) HasTeardownScenarioConfiguration() bool {
	if o != nil && o.TeardownScenarioConfiguration != nil {
		return true
	}

	return false
}

// SetTeardownScenarioConfiguration gets a reference to the given InputConfiguration and assigns it to the TeardownScenarioConfiguration field.
func (o *InputScenarioHost) SetTeardownScenarioConfiguration(v InputConfiguration) {
	o.TeardownScenarioConfiguration = &v
}

func (o InputScenarioHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildOrder != nil {
		toSerialize["buildOrder"] = o.BuildOrder
	}
	if true {
		toSerialize["systemRole"] = o.SystemRole
	}
	if o.SystemModule != nil {
		toSerialize["systemModule"] = o.SystemModule
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if o.ConfigureScenarioConfiguration != nil {
		toSerialize["configureScenarioConfiguration"] = o.ConfigureScenarioConfiguration
	}
	if o.TeardownScenarioConfiguration != nil {
		toSerialize["teardownScenarioConfiguration"] = o.TeardownScenarioConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableInputScenarioHost struct {
	value *InputScenarioHost
	isSet bool
}

func (v NullableInputScenarioHost) Get() *InputScenarioHost {
	return v.value
}

func (v *NullableInputScenarioHost) Set(val *InputScenarioHost) {
	v.value = val
	v.isSet = true
}

func (v NullableInputScenarioHost) IsSet() bool {
	return v.isSet
}

func (v *NullableInputScenarioHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputScenarioHost(val *InputScenarioHost) *NullableInputScenarioHost {
	return &NullableInputScenarioHost{value: val, isSet: true}
}

func (v NullableInputScenarioHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputScenarioHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


