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

// ScenarioHost struct for ScenarioHost
type ScenarioHost struct {
	BuildOrder                     *int32         `json:"buildOrder,omitempty"`
	ConfigureScenarioConfiguration *Configuration `json:"configureScenarioConfiguration,omitempty"`
	TeardownScenarioConfiguration  *Configuration `json:"teardownScenarioConfiguration,omitempty"`
	Id                             *int32         `json:"id,omitempty"`
	Master                         *bool          `json:"master,omitempty"`
	SystemModule                   SystemModule   `json:"systemModule"`
	SystemRole                     string         `json:"systemRole"`
}

// NewScenarioHost instantiates a new ScenarioHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioHost(systemModule SystemModule, systemRole string) *ScenarioHost {
	this := ScenarioHost{}
	this.SystemModule = systemModule
	this.SystemRole = systemRole
	return &this
}

// NewScenarioHostWithDefaults instantiates a new ScenarioHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioHostWithDefaults() *ScenarioHost {
	this := ScenarioHost{}
	return &this
}

// GetBuildOrder returns the BuildOrder field value if set, zero value otherwise.
func (o *ScenarioHost) GetBuildOrder() int32 {
	if o == nil || o.BuildOrder == nil {
		var ret int32
		return ret
	}
	return *o.BuildOrder
}

// GetBuildOrderOk returns a tuple with the BuildOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetBuildOrderOk() (*int32, bool) {
	if o == nil || o.BuildOrder == nil {
		return nil, false
	}
	return o.BuildOrder, true
}

// HasBuildOrder returns a boolean if a field has been set.
func (o *ScenarioHost) HasBuildOrder() bool {
	if o != nil && o.BuildOrder != nil {
		return true
	}

	return false
}

// SetBuildOrder gets a reference to the given int32 and assigns it to the BuildOrder field.
func (o *ScenarioHost) SetBuildOrder(v int32) {
	o.BuildOrder = &v
}

// GetConfigureScenarioConfiguration returns the ConfigureScenarioConfiguration field value if set, zero value otherwise.
func (o *ScenarioHost) GetConfigureScenarioConfiguration() Configuration {
	if o == nil || o.ConfigureScenarioConfiguration == nil {
		var ret Configuration
		return ret
	}
	return *o.ConfigureScenarioConfiguration
}

// GetConfigureScenarioConfigurationOk returns a tuple with the ConfigureScenarioConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetConfigureScenarioConfigurationOk() (*Configuration, bool) {
	if o == nil || o.ConfigureScenarioConfiguration == nil {
		return nil, false
	}
	return o.ConfigureScenarioConfiguration, true
}

// HasConfigureScenarioConfiguration returns a boolean if a field has been set.
func (o *ScenarioHost) HasConfigureScenarioConfiguration() bool {
	if o != nil && o.ConfigureScenarioConfiguration != nil {
		return true
	}

	return false
}

// SetConfigureScenarioConfiguration gets a reference to the given Configuration and assigns it to the ConfigureScenarioConfiguration field.
func (o *ScenarioHost) SetConfigureScenarioConfiguration(v Configuration) {
	o.ConfigureScenarioConfiguration = &v
}

// GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field value if set, zero value otherwise.
func (o *ScenarioHost) GetTeardownScenarioConfiguration() Configuration {
	if o == nil || o.TeardownScenarioConfiguration == nil {
		var ret Configuration
		return ret
	}
	return *o.TeardownScenarioConfiguration
}

// GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetTeardownScenarioConfigurationOk() (*Configuration, bool) {
	if o == nil || o.TeardownScenarioConfiguration == nil {
		return nil, false
	}
	return o.TeardownScenarioConfiguration, true
}

// HasTeardownScenarioConfiguration returns a boolean if a field has been set.
func (o *ScenarioHost) HasTeardownScenarioConfiguration() bool {
	if o != nil && o.TeardownScenarioConfiguration != nil {
		return true
	}

	return false
}

// SetTeardownScenarioConfiguration gets a reference to the given Configuration and assigns it to the TeardownScenarioConfiguration field.
func (o *ScenarioHost) SetTeardownScenarioConfiguration(v Configuration) {
	o.TeardownScenarioConfiguration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScenarioHost) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScenarioHost) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ScenarioHost) SetId(v int32) {
	o.Id = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *ScenarioHost) GetMaster() bool {
	if o == nil || o.Master == nil {
		var ret bool
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetMasterOk() (*bool, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *ScenarioHost) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given bool and assigns it to the Master field.
func (o *ScenarioHost) SetMaster(v bool) {
	o.Master = &v
}

// GetSystemModule returns the SystemModule field value
func (o *ScenarioHost) GetSystemModule() SystemModule {
	if o == nil {
		var ret SystemModule
		return ret
	}

	return o.SystemModule
}

// GetSystemModuleOk returns a tuple with the SystemModule field value
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetSystemModuleOk() (*SystemModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemModule, true
}

// SetSystemModule sets field value
func (o *ScenarioHost) SetSystemModule(v SystemModule) {
	o.SystemModule = v
}

// GetSystemRole returns the SystemRole field value
func (o *ScenarioHost) GetSystemRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value
// and a boolean to check if the value has been set.
func (o *ScenarioHost) GetSystemRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemRole, true
}

// SetSystemRole sets field value
func (o *ScenarioHost) SetSystemRole(v string) {
	o.SystemRole = v
}

func (o ScenarioHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildOrder != nil {
		toSerialize["buildOrder"] = o.BuildOrder
	}
	if o.ConfigureScenarioConfiguration != nil {
		toSerialize["configureScenarioConfiguration"] = o.ConfigureScenarioConfiguration
	}
	if o.TeardownScenarioConfiguration != nil {
		toSerialize["teardownScenarioConfiguration"] = o.TeardownScenarioConfiguration
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if true {
		toSerialize["systemModule"] = o.SystemModule
	}
	if true {
		toSerialize["systemRole"] = o.SystemRole
	}
	return json.Marshal(toSerialize)
}

type NullableScenarioHost struct {
	value *ScenarioHost
	isSet bool
}

func (v NullableScenarioHost) Get() *ScenarioHost {
	return v.value
}

func (v *NullableScenarioHost) Set(val *ScenarioHost) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioHost) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioHost(val *ScenarioHost) *NullableScenarioHost {
	return &NullableScenarioHost{value: val, isSet: true}
}

func (v NullableScenarioHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
