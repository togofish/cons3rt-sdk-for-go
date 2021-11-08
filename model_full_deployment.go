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

// FullDeployment struct for FullDeployment
type FullDeployment struct {
	RecurringSchedules *[]MinimalRecurringSchedule `json:"recurringSchedules,omitempty"`
	Scenario           *[]GeneralScenario          `json:"scenario,omitempty"`
	TestBundles        *[]MinimalTestBundle        `json:"testBundles,omitempty"`
	DeploymentHosts    *[]MinimalDeploymentHost    `json:"deploymentHosts,omitempty"`
	Subtype            string
}

// NewFullDeployment instantiates a new FullDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDeployment(subtype string) *FullDeployment {
	this := FullDeployment{}
	this.Subtype = subtype
	return &this
}

// NewFullDeploymentWithDefaults instantiates a new FullDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDeploymentWithDefaults() *FullDeployment {
	this := FullDeployment{}
	return &this
}

// GetRecurringSchedules returns the RecurringSchedules field value if set, zero value otherwise.
func (o *FullDeployment) GetRecurringSchedules() []MinimalRecurringSchedule {
	if o == nil || o.RecurringSchedules == nil {
		var ret []MinimalRecurringSchedule
		return ret
	}
	return *o.RecurringSchedules
}

// GetRecurringSchedulesOk returns a tuple with the RecurringSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeployment) GetRecurringSchedulesOk() (*[]MinimalRecurringSchedule, bool) {
	if o == nil || o.RecurringSchedules == nil {
		return nil, false
	}
	return o.RecurringSchedules, true
}

// HasRecurringSchedules returns a boolean if a field has been set.
func (o *FullDeployment) HasRecurringSchedules() bool {
	if o != nil && o.RecurringSchedules != nil {
		return true
	}

	return false
}

// SetRecurringSchedules gets a reference to the given []MinimalRecurringSchedule and assigns it to the RecurringSchedules field.
func (o *FullDeployment) SetRecurringSchedules(v []MinimalRecurringSchedule) {
	o.RecurringSchedules = &v
}

// GetScenario returns the Scenario field value if set, zero value otherwise.
func (o *FullDeployment) GetScenario() []GeneralScenario {
	if o == nil || o.Scenario == nil {
		var ret []GeneralScenario
		return ret
	}
	return *o.Scenario
}

// GetScenarioOk returns a tuple with the Scenario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeployment) GetScenarioOk() (*[]GeneralScenario, bool) {
	if o == nil || o.Scenario == nil {
		return nil, false
	}
	return o.Scenario, true
}

// HasScenario returns a boolean if a field has been set.
func (o *FullDeployment) HasScenario() bool {
	if o != nil && o.Scenario != nil {
		return true
	}

	return false
}

// SetScenario gets a reference to the given []GeneralScenario and assigns it to the Scenario field.
func (o *FullDeployment) SetScenario(v []GeneralScenario) {
	o.Scenario = &v
}

// GetTestBundles returns the TestBundles field value if set, zero value otherwise.
func (o *FullDeployment) GetTestBundles() []MinimalTestBundle {
	if o == nil || o.TestBundles == nil {
		var ret []MinimalTestBundle
		return ret
	}
	return *o.TestBundles
}

// GetTestBundlesOk returns a tuple with the TestBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeployment) GetTestBundlesOk() (*[]MinimalTestBundle, bool) {
	if o == nil || o.TestBundles == nil {
		return nil, false
	}
	return o.TestBundles, true
}

// HasTestBundles returns a boolean if a field has been set.
func (o *FullDeployment) HasTestBundles() bool {
	if o != nil && o.TestBundles != nil {
		return true
	}

	return false
}

// SetTestBundles gets a reference to the given []MinimalTestBundle and assigns it to the TestBundles field.
func (o *FullDeployment) SetTestBundles(v []MinimalTestBundle) {
	o.TestBundles = &v
}

// GetDeploymentHosts returns the DeploymentHosts field value if set, zero value otherwise.
func (o *FullDeployment) GetDeploymentHosts() []MinimalDeploymentHost {
	if o == nil || o.DeploymentHosts == nil {
		var ret []MinimalDeploymentHost
		return ret
	}
	return *o.DeploymentHosts
}

// GetDeploymentHostsOk returns a tuple with the DeploymentHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeployment) GetDeploymentHostsOk() (*[]MinimalDeploymentHost, bool) {
	if o == nil || o.DeploymentHosts == nil {
		return nil, false
	}
	return o.DeploymentHosts, true
}

// HasDeploymentHosts returns a boolean if a field has been set.
func (o *FullDeployment) HasDeploymentHosts() bool {
	if o != nil && o.DeploymentHosts != nil {
		return true
	}

	return false
}

// SetDeploymentHosts gets a reference to the given []MinimalDeploymentHost and assigns it to the DeploymentHosts field.
func (o *FullDeployment) SetDeploymentHosts(v []MinimalDeploymentHost) {
	o.DeploymentHosts = &v
}

func (o FullDeployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecurringSchedules != nil {
		toSerialize["recurringSchedules"] = o.RecurringSchedules
	}
	if o.Scenario != nil {
		toSerialize["scenario"] = o.Scenario
	}
	if o.TestBundles != nil {
		toSerialize["testBundles"] = o.TestBundles
	}
	if o.DeploymentHosts != nil {
		toSerialize["deploymentHosts"] = o.DeploymentHosts
	}
	return json.Marshal(toSerialize)
}

type NullableFullDeployment struct {
	value *FullDeployment
	isSet bool
}

func (v NullableFullDeployment) Get() *FullDeployment {
	return v.value
}

func (v *NullableFullDeployment) Set(val *FullDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDeployment(val *FullDeployment) *NullableFullDeployment {
	return &NullableFullDeployment{value: val, isSet: true}
}

func (v NullableFullDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
