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

// DeploymentRunOptions struct for DeploymentRunOptions
type DeploymentRunOptions struct {
	Debug                       *bool                            `json:"debug,omitempty"`
	DeploymentRunName           *string                          `json:"deploymentRunName,omitempty"`
	DeploymentToSubmit          *int32                           `json:"deploymentToSubmit,omitempty"`
	Description                 *string                          `json:"description,omitempty"`
	EndState                    *string                          `json:"endState,omitempty"`
	EarliestStartTime           *int32                           `json:"earliestStartTime,omitempty"`
	EndExisting                 *bool                            `json:"endExisting,omitempty"`
	HostOptions                 *[]HostOption                    `json:"hostOptions,omitempty"`
	Id                          *int32                           `json:"id,omitempty"`
	Locked                      *bool                            `json:"locked,omitempty"`
	Password                    string                           `json:"password"`
	PowerSchedule               *PowerSchedule                   `json:"powerSchedule,omitempty"`
	VirtualizationRealmId       *int32                           `json:"virtualizationRealmId,omitempty"`
	Properties                  *[]Property                      `json:"properties,omitempty"`
	QuickBuildCleanupOverridden *bool                            `json:"quickBuildCleanupOverridden,omitempty"`
	Duration                    *int64                           `json:"duration,omitempty"`
	EndDate                     *int32                           `json:"endDate,omitempty"`
	RetainOnError               *bool                            `json:"retainOnError,omitempty"`
	TaskGroup                   *string                          `json:"taskGroup,omitempty"`
	Username                    string                           `json:"username"`
	VirtRealmBinding            *InputVirtualizationRealmBinding `json:"virtRealmBinding,omitempty"`
	WindowsDomainName           *string                          `json:"windowsDomainName,omitempty"`
}

// NewDeploymentRunOptions instantiates a new DeploymentRunOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentRunOptions(password string, username string) *DeploymentRunOptions {
	this := DeploymentRunOptions{}
	this.Password = password
	this.Username = username
	return &this
}

// NewDeploymentRunOptionsWithDefaults instantiates a new DeploymentRunOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentRunOptionsWithDefaults() *DeploymentRunOptions {
	this := DeploymentRunOptions{}
	return &this
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetDebug() bool {
	if o == nil || o.Debug == nil {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetDebugOk() (*bool, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *DeploymentRunOptions) SetDebug(v bool) {
	o.Debug = &v
}

// GetDeploymentRunName returns the DeploymentRunName field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetDeploymentRunName() string {
	if o == nil || o.DeploymentRunName == nil {
		var ret string
		return ret
	}
	return *o.DeploymentRunName
}

// GetDeploymentRunNameOk returns a tuple with the DeploymentRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetDeploymentRunNameOk() (*string, bool) {
	if o == nil || o.DeploymentRunName == nil {
		return nil, false
	}
	return o.DeploymentRunName, true
}

// HasDeploymentRunName returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasDeploymentRunName() bool {
	if o != nil && o.DeploymentRunName != nil {
		return true
	}

	return false
}

// SetDeploymentRunName gets a reference to the given string and assigns it to the DeploymentRunName field.
func (o *DeploymentRunOptions) SetDeploymentRunName(v string) {
	o.DeploymentRunName = &v
}

// GetDeploymentToSubmit returns the DeploymentToSubmit field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetDeploymentToSubmit() int32 {
	if o == nil || o.DeploymentToSubmit == nil {
		var ret int32
		return ret
	}
	return *o.DeploymentToSubmit
}

// GetDeploymentToSubmitOk returns a tuple with the DeploymentToSubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetDeploymentToSubmitOk() (*int32, bool) {
	if o == nil || o.DeploymentToSubmit == nil {
		return nil, false
	}
	return o.DeploymentToSubmit, true
}

// HasDeploymentToSubmit returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasDeploymentToSubmit() bool {
	if o != nil && o.DeploymentToSubmit != nil {
		return true
	}

	return false
}

// SetDeploymentToSubmit gets a reference to the given int32 and assigns it to the DeploymentToSubmit field.
func (o *DeploymentRunOptions) SetDeploymentToSubmit(v int32) {
	o.DeploymentToSubmit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploymentRunOptions) SetDescription(v string) {
	o.Description = &v
}

// GetEndState returns the EndState field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetEndState() string {
	if o == nil || o.EndState == nil {
		var ret string
		return ret
	}
	return *o.EndState
}

// GetEndStateOk returns a tuple with the EndState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetEndStateOk() (*string, bool) {
	if o == nil || o.EndState == nil {
		return nil, false
	}
	return o.EndState, true
}

// HasEndState returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasEndState() bool {
	if o != nil && o.EndState != nil {
		return true
	}

	return false
}

// SetEndState gets a reference to the given string and assigns it to the EndState field.
func (o *DeploymentRunOptions) SetEndState(v string) {
	o.EndState = &v
}

// GetEarliestStartTime returns the EarliestStartTime field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetEarliestStartTime() int32 {
	if o == nil || o.EarliestStartTime == nil {
		var ret int32
		return ret
	}
	return *o.EarliestStartTime
}

// GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetEarliestStartTimeOk() (*int32, bool) {
	if o == nil || o.EarliestStartTime == nil {
		return nil, false
	}
	return o.EarliestStartTime, true
}

// HasEarliestStartTime returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasEarliestStartTime() bool {
	if o != nil && o.EarliestStartTime != nil {
		return true
	}

	return false
}

// SetEarliestStartTime gets a reference to the given int32 and assigns it to the EarliestStartTime field.
func (o *DeploymentRunOptions) SetEarliestStartTime(v int32) {
	o.EarliestStartTime = &v
}

// GetEndExisting returns the EndExisting field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetEndExisting() bool {
	if o == nil || o.EndExisting == nil {
		var ret bool
		return ret
	}
	return *o.EndExisting
}

// GetEndExistingOk returns a tuple with the EndExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetEndExistingOk() (*bool, bool) {
	if o == nil || o.EndExisting == nil {
		return nil, false
	}
	return o.EndExisting, true
}

// HasEndExisting returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasEndExisting() bool {
	if o != nil && o.EndExisting != nil {
		return true
	}

	return false
}

// SetEndExisting gets a reference to the given bool and assigns it to the EndExisting field.
func (o *DeploymentRunOptions) SetEndExisting(v bool) {
	o.EndExisting = &v
}

// GetHostOptions returns the HostOptions field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetHostOptions() []HostOption {
	if o == nil || o.HostOptions == nil {
		var ret []HostOption
		return ret
	}
	return *o.HostOptions
}

// GetHostOptionsOk returns a tuple with the HostOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetHostOptionsOk() (*[]HostOption, bool) {
	if o == nil || o.HostOptions == nil {
		return nil, false
	}
	return o.HostOptions, true
}

// HasHostOptions returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasHostOptions() bool {
	if o != nil && o.HostOptions != nil {
		return true
	}

	return false
}

// SetHostOptions gets a reference to the given []HostOption and assigns it to the HostOptions field.
func (o *DeploymentRunOptions) SetHostOptions(v []HostOption) {
	o.HostOptions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeploymentRunOptions) SetId(v int32) {
	o.Id = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *DeploymentRunOptions) SetLocked(v bool) {
	o.Locked = &v
}

// GetPassword returns the Password field value
func (o *DeploymentRunOptions) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *DeploymentRunOptions) SetPassword(v string) {
	o.Password = v
}

// GetPowerSchedule returns the PowerSchedule field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetPowerSchedule() PowerSchedule {
	if o == nil || o.PowerSchedule == nil {
		var ret PowerSchedule
		return ret
	}
	return *o.PowerSchedule
}

// GetPowerScheduleOk returns a tuple with the PowerSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetPowerScheduleOk() (*PowerSchedule, bool) {
	if o == nil || o.PowerSchedule == nil {
		return nil, false
	}
	return o.PowerSchedule, true
}

// HasPowerSchedule returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasPowerSchedule() bool {
	if o != nil && o.PowerSchedule != nil {
		return true
	}

	return false
}

// SetPowerSchedule gets a reference to the given PowerSchedule and assigns it to the PowerSchedule field.
func (o *DeploymentRunOptions) SetPowerSchedule(v PowerSchedule) {
	o.PowerSchedule = &v
}

// GetVirtualizationRealmId returns the VirtualizationRealmId field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetVirtualizationRealmId() int32 {
	if o == nil || o.VirtualizationRealmId == nil {
		var ret int32
		return ret
	}
	return *o.VirtualizationRealmId
}

// GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetVirtualizationRealmIdOk() (*int32, bool) {
	if o == nil || o.VirtualizationRealmId == nil {
		return nil, false
	}
	return o.VirtualizationRealmId, true
}

// HasVirtualizationRealmId returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasVirtualizationRealmId() bool {
	if o != nil && o.VirtualizationRealmId != nil {
		return true
	}

	return false
}

// SetVirtualizationRealmId gets a reference to the given int32 and assigns it to the VirtualizationRealmId field.
func (o *DeploymentRunOptions) SetVirtualizationRealmId(v int32) {
	o.VirtualizationRealmId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetProperties() []Property {
	if o == nil || o.Properties == nil {
		var ret []Property
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetPropertiesOk() (*[]Property, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *DeploymentRunOptions) SetProperties(v []Property) {
	o.Properties = &v
}

// GetQuickBuildCleanupOverridden returns the QuickBuildCleanupOverridden field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetQuickBuildCleanupOverridden() bool {
	if o == nil || o.QuickBuildCleanupOverridden == nil {
		var ret bool
		return ret
	}
	return *o.QuickBuildCleanupOverridden
}

// GetQuickBuildCleanupOverriddenOk returns a tuple with the QuickBuildCleanupOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetQuickBuildCleanupOverriddenOk() (*bool, bool) {
	if o == nil || o.QuickBuildCleanupOverridden == nil {
		return nil, false
	}
	return o.QuickBuildCleanupOverridden, true
}

// HasQuickBuildCleanupOverridden returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasQuickBuildCleanupOverridden() bool {
	if o != nil && o.QuickBuildCleanupOverridden != nil {
		return true
	}

	return false
}

// SetQuickBuildCleanupOverridden gets a reference to the given bool and assigns it to the QuickBuildCleanupOverridden field.
func (o *DeploymentRunOptions) SetQuickBuildCleanupOverridden(v bool) {
	o.QuickBuildCleanupOverridden = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentRunOptions) SetDuration(v int64) {
	o.Duration = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetEndDate() int32 {
	if o == nil || o.EndDate == nil {
		var ret int32
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetEndDateOk() (*int32, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int32 and assigns it to the EndDate field.
func (o *DeploymentRunOptions) SetEndDate(v int32) {
	o.EndDate = &v
}

// GetRetainOnError returns the RetainOnError field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetRetainOnError() bool {
	if o == nil || o.RetainOnError == nil {
		var ret bool
		return ret
	}
	return *o.RetainOnError
}

// GetRetainOnErrorOk returns a tuple with the RetainOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetRetainOnErrorOk() (*bool, bool) {
	if o == nil || o.RetainOnError == nil {
		return nil, false
	}
	return o.RetainOnError, true
}

// HasRetainOnError returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasRetainOnError() bool {
	if o != nil && o.RetainOnError != nil {
		return true
	}

	return false
}

// SetRetainOnError gets a reference to the given bool and assigns it to the RetainOnError field.
func (o *DeploymentRunOptions) SetRetainOnError(v bool) {
	o.RetainOnError = &v
}

// GetTaskGroup returns the TaskGroup field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetTaskGroup() string {
	if o == nil || o.TaskGroup == nil {
		var ret string
		return ret
	}
	return *o.TaskGroup
}

// GetTaskGroupOk returns a tuple with the TaskGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetTaskGroupOk() (*string, bool) {
	if o == nil || o.TaskGroup == nil {
		return nil, false
	}
	return o.TaskGroup, true
}

// HasTaskGroup returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasTaskGroup() bool {
	if o != nil && o.TaskGroup != nil {
		return true
	}

	return false
}

// SetTaskGroup gets a reference to the given string and assigns it to the TaskGroup field.
func (o *DeploymentRunOptions) SetTaskGroup(v string) {
	o.TaskGroup = &v
}

// GetUsername returns the Username field value
func (o *DeploymentRunOptions) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *DeploymentRunOptions) SetUsername(v string) {
	o.Username = v
}

// GetVirtRealmBinding returns the VirtRealmBinding field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetVirtRealmBinding() InputVirtualizationRealmBinding {
	if o == nil || o.VirtRealmBinding == nil {
		var ret InputVirtualizationRealmBinding
		return ret
	}
	return *o.VirtRealmBinding
}

// GetVirtRealmBindingOk returns a tuple with the VirtRealmBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetVirtRealmBindingOk() (*InputVirtualizationRealmBinding, bool) {
	if o == nil || o.VirtRealmBinding == nil {
		return nil, false
	}
	return o.VirtRealmBinding, true
}

// HasVirtRealmBinding returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasVirtRealmBinding() bool {
	if o != nil && o.VirtRealmBinding != nil {
		return true
	}

	return false
}

// SetVirtRealmBinding gets a reference to the given InputVirtualizationRealmBinding and assigns it to the VirtRealmBinding field.
func (o *DeploymentRunOptions) SetVirtRealmBinding(v InputVirtualizationRealmBinding) {
	o.VirtRealmBinding = &v
}

// GetWindowsDomainName returns the WindowsDomainName field value if set, zero value otherwise.
func (o *DeploymentRunOptions) GetWindowsDomainName() string {
	if o == nil || o.WindowsDomainName == nil {
		var ret string
		return ret
	}
	return *o.WindowsDomainName
}

// GetWindowsDomainNameOk returns a tuple with the WindowsDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRunOptions) GetWindowsDomainNameOk() (*string, bool) {
	if o == nil || o.WindowsDomainName == nil {
		return nil, false
	}
	return o.WindowsDomainName, true
}

// HasWindowsDomainName returns a boolean if a field has been set.
func (o *DeploymentRunOptions) HasWindowsDomainName() bool {
	if o != nil && o.WindowsDomainName != nil {
		return true
	}

	return false
}

// SetWindowsDomainName gets a reference to the given string and assigns it to the WindowsDomainName field.
func (o *DeploymentRunOptions) SetWindowsDomainName(v string) {
	o.WindowsDomainName = &v
}

func (o DeploymentRunOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.DeploymentRunName != nil {
		toSerialize["deploymentRunName"] = o.DeploymentRunName
	}
	if o.DeploymentToSubmit != nil {
		toSerialize["deploymentToSubmit"] = o.DeploymentToSubmit
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EndState != nil {
		toSerialize["endState"] = o.EndState
	}
	if o.EarliestStartTime != nil {
		toSerialize["earliestStartTime"] = o.EarliestStartTime
	}
	if o.EndExisting != nil {
		toSerialize["endExisting"] = o.EndExisting
	}
	if o.HostOptions != nil {
		toSerialize["hostOptions"] = o.HostOptions
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.PowerSchedule != nil {
		toSerialize["powerSchedule"] = o.PowerSchedule
	}
	if o.VirtualizationRealmId != nil {
		toSerialize["virtualizationRealmId"] = o.VirtualizationRealmId
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.QuickBuildCleanupOverridden != nil {
		toSerialize["quickBuildCleanupOverridden"] = o.QuickBuildCleanupOverridden
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.RetainOnError != nil {
		toSerialize["retainOnError"] = o.RetainOnError
	}
	if o.TaskGroup != nil {
		toSerialize["taskGroup"] = o.TaskGroup
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.VirtRealmBinding != nil {
		toSerialize["virtRealmBinding"] = o.VirtRealmBinding
	}
	if o.WindowsDomainName != nil {
		toSerialize["windowsDomainName"] = o.WindowsDomainName
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentRunOptions struct {
	value *DeploymentRunOptions
	isSet bool
}

func (v NullableDeploymentRunOptions) Get() *DeploymentRunOptions {
	return v.value
}

func (v *NullableDeploymentRunOptions) Set(val *DeploymentRunOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentRunOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentRunOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentRunOptions(val *DeploymentRunOptions) *NullableDeploymentRunOptions {
	return &NullableDeploymentRunOptions{value: val, isSet: true}
}

func (v NullableDeploymentRunOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentRunOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
