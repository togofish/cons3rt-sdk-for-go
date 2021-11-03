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

// FullSystemModule struct for FullSystemModule
type FullSystemModule struct {
	Components *[]MinimalAbstractComponent `json:"components,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Offline *bool `json:"offline,omitempty"`
	State *string `json:"state,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Creator *MinimalUser `json:"creator,omitempty"`
	OwningProject *MinimalProject `json:"owningProject,omitempty"`
	TrustedProjects *[]MinimalProject `json:"trustedProjects,omitempty"`
	DependentAssetCount *int32 `json:"dependentAssetCount,omitempty"`
	Metadata *FullMetadata `json:"metadata,omitempty"`
	ImpactLevel *string `json:"impactLevel,omitempty"`
	Categories *[]MinimalCategory `json:"categories,omitempty"`
	Subtype string `json:"subtype"`
}

// NewFullSystemModule instantiates a new FullSystemModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullSystemModule(subtype string) *FullSystemModule {
	this := FullSystemModule{}
	this.Subtype = subtype
	return &this
}

// NewFullSystemModuleWithDefaults instantiates a new FullSystemModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullSystemModuleWithDefaults() *FullSystemModule {
	this := FullSystemModule{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *FullSystemModule) GetComponents() []MinimalAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []MinimalAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetComponentsOk() (*[]MinimalAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *FullSystemModule) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []MinimalAbstractComponent and assigns it to the Components field.
func (o *FullSystemModule) SetComponents(v []MinimalAbstractComponent) {
	o.Components = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullSystemModule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullSystemModule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullSystemModule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullSystemModule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullSystemModule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullSystemModule) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullSystemModule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullSystemModule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullSystemModule) SetDescription(v string) {
	o.Description = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *FullSystemModule) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *FullSystemModule) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *FullSystemModule) SetOffline(v bool) {
	o.Offline = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FullSystemModule) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FullSystemModule) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FullSystemModule) SetState(v string) {
	o.State = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *FullSystemModule) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *FullSystemModule) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *FullSystemModule) SetVisibility(v string) {
	o.Visibility = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *FullSystemModule) GetCreator() MinimalUser {
	if o == nil || o.Creator == nil {
		var ret MinimalUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetCreatorOk() (*MinimalUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *FullSystemModule) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given MinimalUser and assigns it to the Creator field.
func (o *FullSystemModule) SetCreator(v MinimalUser) {
	o.Creator = &v
}

// GetOwningProject returns the OwningProject field value if set, zero value otherwise.
func (o *FullSystemModule) GetOwningProject() MinimalProject {
	if o == nil || o.OwningProject == nil {
		var ret MinimalProject
		return ret
	}
	return *o.OwningProject
}

// GetOwningProjectOk returns a tuple with the OwningProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetOwningProjectOk() (*MinimalProject, bool) {
	if o == nil || o.OwningProject == nil {
		return nil, false
	}
	return o.OwningProject, true
}

// HasOwningProject returns a boolean if a field has been set.
func (o *FullSystemModule) HasOwningProject() bool {
	if o != nil && o.OwningProject != nil {
		return true
	}

	return false
}

// SetOwningProject gets a reference to the given MinimalProject and assigns it to the OwningProject field.
func (o *FullSystemModule) SetOwningProject(v MinimalProject) {
	o.OwningProject = &v
}

// GetTrustedProjects returns the TrustedProjects field value if set, zero value otherwise.
func (o *FullSystemModule) GetTrustedProjects() []MinimalProject {
	if o == nil || o.TrustedProjects == nil {
		var ret []MinimalProject
		return ret
	}
	return *o.TrustedProjects
}

// GetTrustedProjectsOk returns a tuple with the TrustedProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetTrustedProjectsOk() (*[]MinimalProject, bool) {
	if o == nil || o.TrustedProjects == nil {
		return nil, false
	}
	return o.TrustedProjects, true
}

// HasTrustedProjects returns a boolean if a field has been set.
func (o *FullSystemModule) HasTrustedProjects() bool {
	if o != nil && o.TrustedProjects != nil {
		return true
	}

	return false
}

// SetTrustedProjects gets a reference to the given []MinimalProject and assigns it to the TrustedProjects field.
func (o *FullSystemModule) SetTrustedProjects(v []MinimalProject) {
	o.TrustedProjects = &v
}

// GetDependentAssetCount returns the DependentAssetCount field value if set, zero value otherwise.
func (o *FullSystemModule) GetDependentAssetCount() int32 {
	if o == nil || o.DependentAssetCount == nil {
		var ret int32
		return ret
	}
	return *o.DependentAssetCount
}

// GetDependentAssetCountOk returns a tuple with the DependentAssetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetDependentAssetCountOk() (*int32, bool) {
	if o == nil || o.DependentAssetCount == nil {
		return nil, false
	}
	return o.DependentAssetCount, true
}

// HasDependentAssetCount returns a boolean if a field has been set.
func (o *FullSystemModule) HasDependentAssetCount() bool {
	if o != nil && o.DependentAssetCount != nil {
		return true
	}

	return false
}

// SetDependentAssetCount gets a reference to the given int32 and assigns it to the DependentAssetCount field.
func (o *FullSystemModule) SetDependentAssetCount(v int32) {
	o.DependentAssetCount = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FullSystemModule) GetMetadata() FullMetadata {
	if o == nil || o.Metadata == nil {
		var ret FullMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetMetadataOk() (*FullMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FullSystemModule) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given FullMetadata and assigns it to the Metadata field.
func (o *FullSystemModule) SetMetadata(v FullMetadata) {
	o.Metadata = &v
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise.
func (o *FullSystemModule) GetImpactLevel() string {
	if o == nil || o.ImpactLevel == nil {
		var ret string
		return ret
	}
	return *o.ImpactLevel
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetImpactLevelOk() (*string, bool) {
	if o == nil || o.ImpactLevel == nil {
		return nil, false
	}
	return o.ImpactLevel, true
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *FullSystemModule) HasImpactLevel() bool {
	if o != nil && o.ImpactLevel != nil {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given string and assigns it to the ImpactLevel field.
func (o *FullSystemModule) SetImpactLevel(v string) {
	o.ImpactLevel = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *FullSystemModule) GetCategories() []MinimalCategory {
	if o == nil || o.Categories == nil {
		var ret []MinimalCategory
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetCategoriesOk() (*[]MinimalCategory, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *FullSystemModule) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []MinimalCategory and assigns it to the Categories field.
func (o *FullSystemModule) SetCategories(v []MinimalCategory) {
	o.Categories = &v
}

// GetSubtype returns the Subtype field value
func (o *FullSystemModule) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *FullSystemModule) GetSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *FullSystemModule) SetSubtype(v string) {
	o.Subtype = v
}

func (o FullSystemModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.OwningProject != nil {
		toSerialize["owningProject"] = o.OwningProject
	}
	if o.TrustedProjects != nil {
		toSerialize["trustedProjects"] = o.TrustedProjects
	}
	if o.DependentAssetCount != nil {
		toSerialize["dependentAssetCount"] = o.DependentAssetCount
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ImpactLevel != nil {
		toSerialize["impactLevel"] = o.ImpactLevel
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableFullSystemModule struct {
	value *FullSystemModule
	isSet bool
}

func (v NullableFullSystemModule) Get() *FullSystemModule {
	return v.value
}

func (v *NullableFullSystemModule) Set(val *FullSystemModule) {
	v.value = val
	v.isSet = true
}

func (v NullableFullSystemModule) IsSet() bool {
	return v.isSet
}

func (v *NullableFullSystemModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullSystemModule(val *FullSystemModule) *NullableFullSystemModule {
	return &NullableFullSystemModule{value: val, isSet: true}
}

func (v NullableFullSystemModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullSystemModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


