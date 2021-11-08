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

// MinimalDeploymentRun struct for MinimalDeploymentRun
type MinimalDeploymentRun struct {
	Creator                 *MinimalUser    `json:"creator,omitempty"`
	Id                      *int32          `json:"id,omitempty"`
	Project                 *MinimalProject `json:"project,omitempty"`
	Result                  *string         `json:"result,omitempty"`
	StartTime               *int32          `json:"startTime,omitempty"`
	Canceled                *bool           `json:"canceled,omitempty"`
	DeploymentRunStatus     *string         `json:"deploymentRunStatus,omitempty"`
	Description             *string         `json:"description,omitempty"`
	FapStatus               *string         `json:"fapStatus,omitempty"`
	Locked                  *bool           `json:"locked,omitempty"`
	Name                    *string         `json:"name,omitempty"`
	DeploymentRunResultType *string         `json:"deploymentRunResultType,omitempty"`
}

// NewMinimalDeploymentRun instantiates a new MinimalDeploymentRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalDeploymentRun() *MinimalDeploymentRun {
	this := MinimalDeploymentRun{}
	return &this
}

// NewMinimalDeploymentRunWithDefaults instantiates a new MinimalDeploymentRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalDeploymentRunWithDefaults() *MinimalDeploymentRun {
	this := MinimalDeploymentRun{}
	return &this
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetCreator() MinimalUser {
	if o == nil || o.Creator == nil {
		var ret MinimalUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetCreatorOk() (*MinimalUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given MinimalUser and assigns it to the Creator field.
func (o *MinimalDeploymentRun) SetCreator(v MinimalUser) {
	o.Creator = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MinimalDeploymentRun) SetId(v int32) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetProject() MinimalProject {
	if o == nil || o.Project == nil {
		var ret MinimalProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetProjectOk() (*MinimalProject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given MinimalProject and assigns it to the Project field.
func (o *MinimalDeploymentRun) SetProject(v MinimalProject) {
	o.Project = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *MinimalDeploymentRun) SetResult(v string) {
	o.Result = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetStartTime() int32 {
	if o == nil || o.StartTime == nil {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetStartTimeOk() (*int32, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *MinimalDeploymentRun) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetCanceled() bool {
	if o == nil || o.Canceled == nil {
		var ret bool
		return ret
	}
	return *o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetCanceledOk() (*bool, bool) {
	if o == nil || o.Canceled == nil {
		return nil, false
	}
	return o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasCanceled() bool {
	if o != nil && o.Canceled != nil {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given bool and assigns it to the Canceled field.
func (o *MinimalDeploymentRun) SetCanceled(v bool) {
	o.Canceled = &v
}

// GetDeploymentRunStatus returns the DeploymentRunStatus field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetDeploymentRunStatus() string {
	if o == nil || o.DeploymentRunStatus == nil {
		var ret string
		return ret
	}
	return *o.DeploymentRunStatus
}

// GetDeploymentRunStatusOk returns a tuple with the DeploymentRunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetDeploymentRunStatusOk() (*string, bool) {
	if o == nil || o.DeploymentRunStatus == nil {
		return nil, false
	}
	return o.DeploymentRunStatus, true
}

// HasDeploymentRunStatus returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasDeploymentRunStatus() bool {
	if o != nil && o.DeploymentRunStatus != nil {
		return true
	}

	return false
}

// SetDeploymentRunStatus gets a reference to the given string and assigns it to the DeploymentRunStatus field.
func (o *MinimalDeploymentRun) SetDeploymentRunStatus(v string) {
	o.DeploymentRunStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MinimalDeploymentRun) SetDescription(v string) {
	o.Description = &v
}

// GetFapStatus returns the FapStatus field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetFapStatus() string {
	if o == nil || o.FapStatus == nil {
		var ret string
		return ret
	}
	return *o.FapStatus
}

// GetFapStatusOk returns a tuple with the FapStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetFapStatusOk() (*string, bool) {
	if o == nil || o.FapStatus == nil {
		return nil, false
	}
	return o.FapStatus, true
}

// HasFapStatus returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasFapStatus() bool {
	if o != nil && o.FapStatus != nil {
		return true
	}

	return false
}

// SetFapStatus gets a reference to the given string and assigns it to the FapStatus field.
func (o *MinimalDeploymentRun) SetFapStatus(v string) {
	o.FapStatus = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *MinimalDeploymentRun) SetLocked(v bool) {
	o.Locked = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MinimalDeploymentRun) SetName(v string) {
	o.Name = &v
}

// GetDeploymentRunResultType returns the DeploymentRunResultType field value if set, zero value otherwise.
func (o *MinimalDeploymentRun) GetDeploymentRunResultType() string {
	if o == nil || o.DeploymentRunResultType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentRunResultType
}

// GetDeploymentRunResultTypeOk returns a tuple with the DeploymentRunResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalDeploymentRun) GetDeploymentRunResultTypeOk() (*string, bool) {
	if o == nil || o.DeploymentRunResultType == nil {
		return nil, false
	}
	return o.DeploymentRunResultType, true
}

// HasDeploymentRunResultType returns a boolean if a field has been set.
func (o *MinimalDeploymentRun) HasDeploymentRunResultType() bool {
	if o != nil && o.DeploymentRunResultType != nil {
		return true
	}

	return false
}

// SetDeploymentRunResultType gets a reference to the given string and assigns it to the DeploymentRunResultType field.
func (o *MinimalDeploymentRun) SetDeploymentRunResultType(v string) {
	o.DeploymentRunResultType = &v
}

func (o MinimalDeploymentRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.Canceled != nil {
		toSerialize["canceled"] = o.Canceled
	}
	if o.DeploymentRunStatus != nil {
		toSerialize["deploymentRunStatus"] = o.DeploymentRunStatus
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FapStatus != nil {
		toSerialize["fapStatus"] = o.FapStatus
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DeploymentRunResultType != nil {
		toSerialize["deploymentRunResultType"] = o.DeploymentRunResultType
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalDeploymentRun struct {
	value *MinimalDeploymentRun
	isSet bool
}

func (v NullableMinimalDeploymentRun) Get() *MinimalDeploymentRun {
	return v.value
}

func (v *NullableMinimalDeploymentRun) Set(val *MinimalDeploymentRun) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalDeploymentRun) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalDeploymentRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalDeploymentRun(val *MinimalDeploymentRun) *NullableMinimalDeploymentRun {
	return &NullableMinimalDeploymentRun{value: val, isSet: true}
}

func (v NullableMinimalDeploymentRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalDeploymentRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
