/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RecurringSchedule struct for RecurringSchedule
type RecurringSchedule struct {
	Id                   *int32               `json:"id,omitempty"`
	Complete             *bool                `json:"complete,omitempty"`
	EndDate              *int32               `json:"endDate,omitempty"`
	MaxIterations        *int32               `json:"maxIterations,omitempty"`
	DeploymentRunOptions DeploymentRunOptions `json:"deploymentRunOptions"`
	RemainingIterations  *int32               `json:"remainingIterations,omitempty"`
	Schedule             string               `json:"schedule"`
	Timezone             string               `json:"timezone"`
}

// NewRecurringSchedule instantiates a new RecurringSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringSchedule(deploymentRunOptions DeploymentRunOptions, schedule string, timezone string) *RecurringSchedule {
	this := RecurringSchedule{}
	this.DeploymentRunOptions = deploymentRunOptions
	this.Schedule = schedule
	this.Timezone = timezone
	return &this
}

// NewRecurringScheduleWithDefaults instantiates a new RecurringSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringScheduleWithDefaults() *RecurringSchedule {
	this := RecurringSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecurringSchedule) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecurringSchedule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RecurringSchedule) SetId(v int32) {
	o.Id = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *RecurringSchedule) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *RecurringSchedule) HasComplete() bool {
	if o != nil && o.Complete != nil {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *RecurringSchedule) SetComplete(v bool) {
	o.Complete = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *RecurringSchedule) GetEndDate() int32 {
	if o == nil || o.EndDate == nil {
		var ret int32
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetEndDateOk() (*int32, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *RecurringSchedule) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int32 and assigns it to the EndDate field.
func (o *RecurringSchedule) SetEndDate(v int32) {
	o.EndDate = &v
}

// GetMaxIterations returns the MaxIterations field value if set, zero value otherwise.
func (o *RecurringSchedule) GetMaxIterations() int32 {
	if o == nil || o.MaxIterations == nil {
		var ret int32
		return ret
	}
	return *o.MaxIterations
}

// GetMaxIterationsOk returns a tuple with the MaxIterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetMaxIterationsOk() (*int32, bool) {
	if o == nil || o.MaxIterations == nil {
		return nil, false
	}
	return o.MaxIterations, true
}

// HasMaxIterations returns a boolean if a field has been set.
func (o *RecurringSchedule) HasMaxIterations() bool {
	if o != nil && o.MaxIterations != nil {
		return true
	}

	return false
}

// SetMaxIterations gets a reference to the given int32 and assigns it to the MaxIterations field.
func (o *RecurringSchedule) SetMaxIterations(v int32) {
	o.MaxIterations = &v
}

// GetDeploymentRunOptions returns the DeploymentRunOptions field value
func (o *RecurringSchedule) GetDeploymentRunOptions() DeploymentRunOptions {
	if o == nil {
		var ret DeploymentRunOptions
		return ret
	}

	return o.DeploymentRunOptions
}

// GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field value
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetDeploymentRunOptionsOk() (*DeploymentRunOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentRunOptions, true
}

// SetDeploymentRunOptions sets field value
func (o *RecurringSchedule) SetDeploymentRunOptions(v DeploymentRunOptions) {
	o.DeploymentRunOptions = v
}

// GetRemainingIterations returns the RemainingIterations field value if set, zero value otherwise.
func (o *RecurringSchedule) GetRemainingIterations() int32 {
	if o == nil || o.RemainingIterations == nil {
		var ret int32
		return ret
	}
	return *o.RemainingIterations
}

// GetRemainingIterationsOk returns a tuple with the RemainingIterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetRemainingIterationsOk() (*int32, bool) {
	if o == nil || o.RemainingIterations == nil {
		return nil, false
	}
	return o.RemainingIterations, true
}

// HasRemainingIterations returns a boolean if a field has been set.
func (o *RecurringSchedule) HasRemainingIterations() bool {
	if o != nil && o.RemainingIterations != nil {
		return true
	}

	return false
}

// SetRemainingIterations gets a reference to the given int32 and assigns it to the RemainingIterations field.
func (o *RecurringSchedule) SetRemainingIterations(v int32) {
	o.RemainingIterations = &v
}

// GetSchedule returns the Schedule field value
func (o *RecurringSchedule) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *RecurringSchedule) SetSchedule(v string) {
	o.Schedule = v
}

// GetTimezone returns the Timezone field value
func (o *RecurringSchedule) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *RecurringSchedule) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *RecurringSchedule) SetTimezone(v string) {
	o.Timezone = v
}

func (o RecurringSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.MaxIterations != nil {
		toSerialize["maxIterations"] = o.MaxIterations
	}
	if true {
		toSerialize["deploymentRunOptions"] = o.DeploymentRunOptions
	}
	if o.RemainingIterations != nil {
		toSerialize["remainingIterations"] = o.RemainingIterations
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableRecurringSchedule struct {
	value *RecurringSchedule
	isSet bool
}

func (v NullableRecurringSchedule) Get() *RecurringSchedule {
	return v.value
}

func (v *NullableRecurringSchedule) Set(val *RecurringSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringSchedule(val *RecurringSchedule) *NullableRecurringSchedule {
	return &NullableRecurringSchedule{value: val, isSet: true}
}

func (v NullableRecurringSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
