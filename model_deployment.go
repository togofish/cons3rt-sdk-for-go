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

// Deployment struct for Deployment
type Deployment struct {
	RecurringSchedules  *[]RecurringSchedule  `json:"recurringSchedules,omitempty"`
	Id                  *string               `json:"id,omitempty"`
	TrustedProjects     *[]Project            `json:"trustedProjects,omitempty"`
	Creator             *User                 `json:"creator,omitempty"`
	DependentAssetCount *int32                `json:"dependentAssetCount,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Metadata            *Metadata             `json:"metadata,omitempty"`
	Name                *string               `json:"name,omitempty"`
	Offline             *bool                 `json:"offline,omitempty"`
	OwningProject       *Project              `json:"owningProject,omitempty"`
	State               *string               `json:"state,omitempty"`
	Visibility          *string               `json:"visibility,omitempty"`
	ImpactLevel         *string               `json:"impactLevel,omitempty"`
	Categories          *[]Category           `json:"categories,omitempty"`
	Scenarios           *[]DeploymentScenario `json:"scenarios,omitempty"`
	TestBundles         *[]TestBundle         `json:"testBundles,omitempty"`
	DeploymentHosts     *[]DeploymentHost     `json:"deploymentHosts,omitempty"`
}

// NewDeployment instantiates a new Deployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployment() *Deployment {
	this := Deployment{}
	return &this
}

// NewDeploymentWithDefaults instantiates a new Deployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentWithDefaults() *Deployment {
	this := Deployment{}
	return &this
}

// GetRecurringSchedules returns the RecurringSchedules field value if set, zero value otherwise.
func (o *Deployment) GetRecurringSchedules() []RecurringSchedule {
	if o == nil || o.RecurringSchedules == nil {
		var ret []RecurringSchedule
		return ret
	}
	return *o.RecurringSchedules
}

// GetRecurringSchedulesOk returns a tuple with the RecurringSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetRecurringSchedulesOk() (*[]RecurringSchedule, bool) {
	if o == nil || o.RecurringSchedules == nil {
		return nil, false
	}
	return o.RecurringSchedules, true
}

// HasRecurringSchedules returns a boolean if a field has been set.
func (o *Deployment) HasRecurringSchedules() bool {
	if o != nil && o.RecurringSchedules != nil {
		return true
	}

	return false
}

// SetRecurringSchedules gets a reference to the given []RecurringSchedule and assigns it to the RecurringSchedules field.
func (o *Deployment) SetRecurringSchedules(v []RecurringSchedule) {
	o.RecurringSchedules = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Deployment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Deployment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Deployment) SetId(v string) {
	o.Id = &v
}

// GetTrustedProjects returns the TrustedProjects field value if set, zero value otherwise.
func (o *Deployment) GetTrustedProjects() []Project {
	if o == nil || o.TrustedProjects == nil {
		var ret []Project
		return ret
	}
	return *o.TrustedProjects
}

// GetTrustedProjectsOk returns a tuple with the TrustedProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetTrustedProjectsOk() (*[]Project, bool) {
	if o == nil || o.TrustedProjects == nil {
		return nil, false
	}
	return o.TrustedProjects, true
}

// HasTrustedProjects returns a boolean if a field has been set.
func (o *Deployment) HasTrustedProjects() bool {
	if o != nil && o.TrustedProjects != nil {
		return true
	}

	return false
}

// SetTrustedProjects gets a reference to the given []Project and assigns it to the TrustedProjects field.
func (o *Deployment) SetTrustedProjects(v []Project) {
	o.TrustedProjects = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Deployment) GetCreator() User {
	if o == nil || o.Creator == nil {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetCreatorOk() (*User, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Deployment) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *Deployment) SetCreator(v User) {
	o.Creator = &v
}

// GetDependentAssetCount returns the DependentAssetCount field value if set, zero value otherwise.
func (o *Deployment) GetDependentAssetCount() int32 {
	if o == nil || o.DependentAssetCount == nil {
		var ret int32
		return ret
	}
	return *o.DependentAssetCount
}

// GetDependentAssetCountOk returns a tuple with the DependentAssetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDependentAssetCountOk() (*int32, bool) {
	if o == nil || o.DependentAssetCount == nil {
		return nil, false
	}
	return o.DependentAssetCount, true
}

// HasDependentAssetCount returns a boolean if a field has been set.
func (o *Deployment) HasDependentAssetCount() bool {
	if o != nil && o.DependentAssetCount != nil {
		return true
	}

	return false
}

// SetDependentAssetCount gets a reference to the given int32 and assigns it to the DependentAssetCount field.
func (o *Deployment) SetDependentAssetCount(v int32) {
	o.DependentAssetCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Deployment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Deployment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Deployment) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Deployment) GetMetadata() Metadata {
	if o == nil || o.Metadata == nil {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetMetadataOk() (*Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Deployment) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *Deployment) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Deployment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Deployment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Deployment) SetName(v string) {
	o.Name = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *Deployment) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *Deployment) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *Deployment) SetOffline(v bool) {
	o.Offline = &v
}

// GetOwningProject returns the OwningProject field value if set, zero value otherwise.
func (o *Deployment) GetOwningProject() Project {
	if o == nil || o.OwningProject == nil {
		var ret Project
		return ret
	}
	return *o.OwningProject
}

// GetOwningProjectOk returns a tuple with the OwningProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetOwningProjectOk() (*Project, bool) {
	if o == nil || o.OwningProject == nil {
		return nil, false
	}
	return o.OwningProject, true
}

// HasOwningProject returns a boolean if a field has been set.
func (o *Deployment) HasOwningProject() bool {
	if o != nil && o.OwningProject != nil {
		return true
	}

	return false
}

// SetOwningProject gets a reference to the given Project and assigns it to the OwningProject field.
func (o *Deployment) SetOwningProject(v Project) {
	o.OwningProject = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Deployment) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Deployment) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Deployment) SetState(v string) {
	o.State = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Deployment) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Deployment) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *Deployment) SetVisibility(v string) {
	o.Visibility = &v
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise.
func (o *Deployment) GetImpactLevel() string {
	if o == nil || o.ImpactLevel == nil {
		var ret string
		return ret
	}
	return *o.ImpactLevel
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetImpactLevelOk() (*string, bool) {
	if o == nil || o.ImpactLevel == nil {
		return nil, false
	}
	return o.ImpactLevel, true
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *Deployment) HasImpactLevel() bool {
	if o != nil && o.ImpactLevel != nil {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given string and assigns it to the ImpactLevel field.
func (o *Deployment) SetImpactLevel(v string) {
	o.ImpactLevel = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Deployment) GetCategories() []Category {
	if o == nil || o.Categories == nil {
		var ret []Category
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetCategoriesOk() (*[]Category, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Deployment) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *Deployment) SetCategories(v []Category) {
	o.Categories = &v
}

// GetScenarios returns the Scenarios field value if set, zero value otherwise.
func (o *Deployment) GetScenarios() []DeploymentScenario {
	if o == nil || o.Scenarios == nil {
		var ret []DeploymentScenario
		return ret
	}
	return *o.Scenarios
}

// GetScenariosOk returns a tuple with the Scenarios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetScenariosOk() (*[]DeploymentScenario, bool) {
	if o == nil || o.Scenarios == nil {
		return nil, false
	}
	return o.Scenarios, true
}

// HasScenarios returns a boolean if a field has been set.
func (o *Deployment) HasScenarios() bool {
	if o != nil && o.Scenarios != nil {
		return true
	}

	return false
}

// SetScenarios gets a reference to the given []DeploymentScenario and assigns it to the Scenarios field.
func (o *Deployment) SetScenarios(v []DeploymentScenario) {
	o.Scenarios = &v
}

// GetTestBundles returns the TestBundles field value if set, zero value otherwise.
func (o *Deployment) GetTestBundles() []TestBundle {
	if o == nil || o.TestBundles == nil {
		var ret []TestBundle
		return ret
	}
	return *o.TestBundles
}

// GetTestBundlesOk returns a tuple with the TestBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetTestBundlesOk() (*[]TestBundle, bool) {
	if o == nil || o.TestBundles == nil {
		return nil, false
	}
	return o.TestBundles, true
}

// HasTestBundles returns a boolean if a field has been set.
func (o *Deployment) HasTestBundles() bool {
	if o != nil && o.TestBundles != nil {
		return true
	}

	return false
}

// SetTestBundles gets a reference to the given []TestBundle and assigns it to the TestBundles field.
func (o *Deployment) SetTestBundles(v []TestBundle) {
	o.TestBundles = &v
}

// GetDeploymentHosts returns the DeploymentHosts field value if set, zero value otherwise.
func (o *Deployment) GetDeploymentHosts() []DeploymentHost {
	if o == nil || o.DeploymentHosts == nil {
		var ret []DeploymentHost
		return ret
	}
	return *o.DeploymentHosts
}

// GetDeploymentHostsOk returns a tuple with the DeploymentHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDeploymentHostsOk() (*[]DeploymentHost, bool) {
	if o == nil || o.DeploymentHosts == nil {
		return nil, false
	}
	return o.DeploymentHosts, true
}

// HasDeploymentHosts returns a boolean if a field has been set.
func (o *Deployment) HasDeploymentHosts() bool {
	if o != nil && o.DeploymentHosts != nil {
		return true
	}

	return false
}

// SetDeploymentHosts gets a reference to the given []DeploymentHost and assigns it to the DeploymentHosts field.
func (o *Deployment) SetDeploymentHosts(v []DeploymentHost) {
	o.DeploymentHosts = &v
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecurringSchedules != nil {
		toSerialize["recurringSchedules"] = o.RecurringSchedules
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TrustedProjects != nil {
		toSerialize["trustedProjects"] = o.TrustedProjects
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.DependentAssetCount != nil {
		toSerialize["dependentAssetCount"] = o.DependentAssetCount
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	if o.OwningProject != nil {
		toSerialize["owningProject"] = o.OwningProject
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.ImpactLevel != nil {
		toSerialize["impactLevel"] = o.ImpactLevel
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Scenarios != nil {
		toSerialize["scenarios"] = o.Scenarios
	}
	if o.TestBundles != nil {
		toSerialize["testBundles"] = o.TestBundles
	}
	if o.DeploymentHosts != nil {
		toSerialize["deploymentHosts"] = o.DeploymentHosts
	}
	return json.Marshal(toSerialize)
}

type NullableDeployment struct {
	value *Deployment
	isSet bool
}

func (v NullableDeployment) Get() *Deployment {
	return v.value
}

func (v *NullableDeployment) Set(val *Deployment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployment(val *Deployment) *NullableDeployment {
	return &NullableDeployment{value: val, isSet: true}
}

func (v NullableDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
