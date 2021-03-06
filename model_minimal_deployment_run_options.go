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

// MinimalDeploymentRunOptions struct for MinimalDeploymentRunOptions
type MinimalDeploymentRunOptions struct {
	Debug                 *bool         `json:"debug,omitempty"`
	DeploymentRunName     *string       `json:"deploymentRunName,omitempty"`
	DeploymentToSubmit    *int32        `json:"deploymentToSubmit,omitempty"`
	Description           *string       `json:"description,omitempty"`
	EndState              *string       `json:"endState,omitempty"`
	Duration              *int64        `json:"duration,omitempty"`
	EarliestStartTime     *int32        `json:"earliestStartTime,omitempty"`
	EndExisting           *bool         `json:"endExisting,omitempty"`
	HostOptions           *[]HostOption `json:"hostOptions,omitempty"`
	Id                    *int32        `json:"id,omitempty"`
	Locked                *bool         `json:"locked,omitempty"`
	VirtualizationRealmId *int32        `json:"virtualizationRealmId,omitempty"`
	Properties            *[]Property   `json:"properties,omitempty"`
	RetainOnError         *bool         `json:"retainOnError,omitempty"`
	RootPassword          *string       `json:"rootPassword,omitempty"`
	Creator               *MinimalUser  `json:"creator,omitempty"`
}

// NewMinimalDeploymentRunOptions instantiates a new MinimalDeploymentRunOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalDeploymentRunOptions() *MinimalDeploymentRunOptions {
	this := MinimalDeploymentRunOptions{}
	return &this
}

// NewMinimalDeploymentRunOptionsWithDefaults instantiates a new MinimalDeploymentRunOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalDeploymentRunOptionsWithDefaults() *MinimalDeploymentRunOptions {
	this := MinimalDeploymentRunOptions{}
	return &this
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetDebug() bool {
	if o == nil || o.Debug == nil {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetDebugOk() (*bool, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *MinimalDeploymentRunOptions) SetDebug(v bool) {
	o.Debug = &v
}

// GetDeploymentRunName returns the DeploymentRunName field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetDeploymentRunName() string {
	if o == nil || o.DeploymentRunName == nil {
		var ret string
		return ret
	}
	return *o.DeploymentRunName
}

// GetDeploymentRunNameOk returns a tuple with the DeploymentRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetDeploymentRunNameOk() (*string, bool) {
	if o == nil || o.DeploymentRunName == nil {
		return nil, false
	}
	return o.DeploymentRunName, true
}

// HasDeploymentRunName returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasDeploymentRunName() bool {
	if o != nil && o.DeploymentRunName != nil {
		return true
	}

	return false
}

// SetDeploymentRunName gets a reference to the given string and assigns it to the DeploymentRunName field.
func (o *MinimalDeploymentRunOptions) SetDeploymentRunName(v string) {
	o.DeploymentRunName = &v
}

// GetDeploymentToSubmit returns the DeploymentToSubmit field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetDeploymentToSubmit() int32 {
	if o == nil || o.DeploymentToSubmit == nil {
		var ret int32
		return ret
	}
	return *o.DeploymentToSubmit
}

// GetDeploymentToSubmitOk returns a tuple with the DeploymentToSubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetDeploymentToSubmitOk() (*int32, bool) {
	if o == nil || o.DeploymentToSubmit == nil {
		return nil, false
	}
	return o.DeploymentToSubmit, true
}

// HasDeploymentToSubmit returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasDeploymentToSubmit() bool {
	if o != nil && o.DeploymentToSubmit != nil {
		return true
	}

	return false
}

// SetDeploymentToSubmit gets a reference to the given int32 and assigns it to the DeploymentToSubmit field.
func (o *MinimalDeploymentRunOptions) SetDeploymentToSubmit(v int32) {
	o.DeploymentToSubmit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MinimalDeploymentRunOptions) SetDescription(v string) {
	o.Description = &v
}

// GetEndState returns the EndState field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetEndState() string {
	if o == nil || o.EndState == nil {
		var ret string
		return ret
	}
	return *o.EndState
}

// GetEndStateOk returns a tuple with the EndState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetEndStateOk() (*string, bool) {
	if o == nil || o.EndState == nil {
		return nil, false
	}
	return o.EndState, true
}

// HasEndState returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasEndState() bool {
	if o != nil && o.EndState != nil {
		return true
	}

	return false
}

// SetEndState gets a reference to the given string and assigns it to the EndState field.
func (o *MinimalDeploymentRunOptions) SetEndState(v string) {
	o.EndState = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *MinimalDeploymentRunOptions) SetDuration(v int64) {
	o.Duration = &v
}

// GetEarliestStartTime returns the EarliestStartTime field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetEarliestStartTime() int32 {
	if o == nil || o.EarliestStartTime == nil {
		var ret int32
		return ret
	}
	return *o.EarliestStartTime
}

// GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetEarliestStartTimeOk() (*int32, bool) {
	if o == nil || o.EarliestStartTime == nil {
		return nil, false
	}
	return o.EarliestStartTime, true
}

// HasEarliestStartTime returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasEarliestStartTime() bool {
	if o != nil && o.EarliestStartTime != nil {
		return true
	}

	return false
}

// SetEarliestStartTime gets a reference to the given int32 and assigns it to the EarliestStartTime field.
func (o *MinimalDeploymentRunOptions) SetEarliestStartTime(v int32) {
	o.EarliestStartTime = &v
}

// GetEndExisting returns the EndExisting field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetEndExisting() bool {
	if o == nil || o.EndExisting == nil {
		var ret bool
		return ret
	}
	return *o.EndExisting
}

// GetEndExistingOk returns a tuple with the EndExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetEndExistingOk() (*bool, bool) {
	if o == nil || o.EndExisting == nil {
		return nil, false
	}
	return o.EndExisting, true
}

// HasEndExisting returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasEndExisting() bool {
	if o != nil && o.EndExisting != nil {
		return true
	}

	return false
}

// SetEndExisting gets a reference to the given bool and assigns it to the EndExisting field.
func (o *MinimalDeploymentRunOptions) SetEndExisting(v bool) {
	o.EndExisting = &v
}

// GetHostOptions returns the HostOptions field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetHostOptions() []HostOption {
	if o == nil || o.HostOptions == nil {
		var ret []HostOption
		return ret
	}
	return *o.HostOptions
}

// GetHostOptionsOk returns a tuple with the HostOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetHostOptionsOk() (*[]HostOption, bool) {
	if o == nil || o.HostOptions == nil {
		return nil, false
	}
	return o.HostOptions, true
}

// HasHostOptions returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasHostOptions() bool {
	if o != nil && o.HostOptions != nil {
		return true
	}

	return false
}

// SetHostOptions gets a reference to the given []HostOption and assigns it to the HostOptions field.
func (o *MinimalDeploymentRunOptions) SetHostOptions(v []HostOption) {
	o.HostOptions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MinimalDeploymentRunOptions) SetId(v int32) {
	o.Id = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *MinimalDeploymentRunOptions) SetLocked(v bool) {
	o.Locked = &v
}

// GetVirtualizationRealmId returns the VirtualizationRealmId field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetVirtualizationRealmId() int32 {
	if o == nil || o.VirtualizationRealmId == nil {
		var ret int32
		return ret
	}
	return *o.VirtualizationRealmId
}

// GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetVirtualizationRealmIdOk() (*int32, bool) {
	if o == nil || o.VirtualizationRealmId == nil {
		return nil, false
	}
	return o.VirtualizationRealmId, true
}

// HasVirtualizationRealmId returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasVirtualizationRealmId() bool {
	if o != nil && o.VirtualizationRealmId != nil {
		return true
	}

	return false
}

// SetVirtualizationRealmId gets a reference to the given int32 and assigns it to the VirtualizationRealmId field.
func (o *MinimalDeploymentRunOptions) SetVirtualizationRealmId(v int32) {
	o.VirtualizationRealmId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetProperties() []Property {
	if o == nil || o.Properties == nil {
		var ret []Property
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetPropertiesOk() (*[]Property, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *MinimalDeploymentRunOptions) SetProperties(v []Property) {
	o.Properties = &v
}

// GetRetainOnError returns the RetainOnError field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetRetainOnError() bool {
	if o == nil || o.RetainOnError == nil {
		var ret bool
		return ret
	}
	return *o.RetainOnError
}

// GetRetainOnErrorOk returns a tuple with the RetainOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetRetainOnErrorOk() (*bool, bool) {
	if o == nil || o.RetainOnError == nil {
		return nil, false
	}
	return o.RetainOnError, true
}

// HasRetainOnError returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasRetainOnError() bool {
	if o != nil && o.RetainOnError != nil {
		return true
	}

	return false
}

// SetRetainOnError gets a reference to the given bool and assigns it to the RetainOnError field.
func (o *MinimalDeploymentRunOptions) SetRetainOnError(v bool) {
	o.RetainOnError = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetRootPassword() string {
	if o == nil || o.RootPassword == nil {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetRootPasswordOk() (*string, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *MinimalDeploymentRunOptions) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *MinimalDeploymentRunOptions) GetCreator() MinimalUser {
	if o == nil || o.Creator == nil {
		var ret MinimalUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRunOptions) GetCreatorOk() (*MinimalUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *MinimalDeploymentRunOptions) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given MinimalUser and assigns it to the Creator field.
func (o *MinimalDeploymentRunOptions) SetCreator(v MinimalUser) {
	o.Creator = &v
}

func (o MinimalDeploymentRunOptions) MarshalJSON() ([]byte, error) {
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
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
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
	if o.VirtualizationRealmId != nil {
		toSerialize["virtualizationRealmId"] = o.VirtualizationRealmId
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.RetainOnError != nil {
		toSerialize["retainOnError"] = o.RetainOnError
	}
	if o.RootPassword != nil {
		toSerialize["rootPassword"] = o.RootPassword
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalDeploymentRunOptions struct {
	value *MinimalDeploymentRunOptions
	isSet bool
}

func (v NullableMinimalDeploymentRunOptions) Get() *MinimalDeploymentRunOptions {
	return v.value
}

func (v *NullableMinimalDeploymentRunOptions) Set(val *MinimalDeploymentRunOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalDeploymentRunOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalDeploymentRunOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalDeploymentRunOptions(val *MinimalDeploymentRunOptions) *NullableMinimalDeploymentRunOptions {
	return &NullableMinimalDeploymentRunOptions{value: val, isSet: true}
}

func (v NullableMinimalDeploymentRunOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalDeploymentRunOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
