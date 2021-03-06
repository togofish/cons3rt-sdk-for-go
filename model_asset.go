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

// Asset struct for Asset
type Asset struct {
	Id                  *string     `json:"id,omitempty"`
	TrustedProjects     *[]Project  `json:"trustedProjects,omitempty"`
	Creator             *User       `json:"creator,omitempty"`
	DependentAssetCount *int32      `json:"dependentAssetCount,omitempty"`
	Description         *string     `json:"description,omitempty"`
	Metadata            *Metadata   `json:"metadata,omitempty"`
	Name                *string     `json:"name,omitempty"`
	Offline             *bool       `json:"offline,omitempty"`
	OwningProject       *Project    `json:"owningProject,omitempty"`
	State               *string     `json:"state,omitempty"`
	Visibility          *string     `json:"visibility,omitempty"`
	ImpactLevel         *string     `json:"impactLevel,omitempty"`
	Categories          *[]Category `json:"categories,omitempty"`
}

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset() *Asset {
	this := Asset{}
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Asset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Asset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Asset) SetId(v string) {
	o.Id = &v
}

// GetTrustedProjects returns the TrustedProjects field value if set, zero value otherwise.
func (o *Asset) GetTrustedProjects() []Project {
	if o == nil || o.TrustedProjects == nil {
		var ret []Project
		return ret
	}
	return *o.TrustedProjects
}

// GetTrustedProjectsOk returns a tuple with the TrustedProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetTrustedProjectsOk() (*[]Project, bool) {
	if o == nil || o.TrustedProjects == nil {
		return nil, false
	}
	return o.TrustedProjects, true
}

// HasTrustedProjects returns a boolean if a field has been set.
func (o *Asset) HasTrustedProjects() bool {
	if o != nil && o.TrustedProjects != nil {
		return true
	}

	return false
}

// SetTrustedProjects gets a reference to the given []Project and assigns it to the TrustedProjects field.
func (o *Asset) SetTrustedProjects(v []Project) {
	o.TrustedProjects = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Asset) GetCreator() User {
	if o == nil || o.Creator == nil {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCreatorOk() (*User, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Asset) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *Asset) SetCreator(v User) {
	o.Creator = &v
}

// GetDependentAssetCount returns the DependentAssetCount field value if set, zero value otherwise.
func (o *Asset) GetDependentAssetCount() int32 {
	if o == nil || o.DependentAssetCount == nil {
		var ret int32
		return ret
	}
	return *o.DependentAssetCount
}

// GetDependentAssetCountOk returns a tuple with the DependentAssetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetDependentAssetCountOk() (*int32, bool) {
	if o == nil || o.DependentAssetCount == nil {
		return nil, false
	}
	return o.DependentAssetCount, true
}

// HasDependentAssetCount returns a boolean if a field has been set.
func (o *Asset) HasDependentAssetCount() bool {
	if o != nil && o.DependentAssetCount != nil {
		return true
	}

	return false
}

// SetDependentAssetCount gets a reference to the given int32 and assigns it to the DependentAssetCount field.
func (o *Asset) SetDependentAssetCount(v int32) {
	o.DependentAssetCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Asset) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Asset) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Asset) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Asset) GetMetadata() Metadata {
	if o == nil || o.Metadata == nil {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetMetadataOk() (*Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Asset) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *Asset) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Asset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Asset) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Asset) SetName(v string) {
	o.Name = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *Asset) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *Asset) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *Asset) SetOffline(v bool) {
	o.Offline = &v
}

// GetOwningProject returns the OwningProject field value if set, zero value otherwise.
func (o *Asset) GetOwningProject() Project {
	if o == nil || o.OwningProject == nil {
		var ret Project
		return ret
	}
	return *o.OwningProject
}

// GetOwningProjectOk returns a tuple with the OwningProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetOwningProjectOk() (*Project, bool) {
	if o == nil || o.OwningProject == nil {
		return nil, false
	}
	return o.OwningProject, true
}

// HasOwningProject returns a boolean if a field has been set.
func (o *Asset) HasOwningProject() bool {
	if o != nil && o.OwningProject != nil {
		return true
	}

	return false
}

// SetOwningProject gets a reference to the given Project and assigns it to the OwningProject field.
func (o *Asset) SetOwningProject(v Project) {
	o.OwningProject = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Asset) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Asset) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Asset) SetState(v string) {
	o.State = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Asset) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Asset) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *Asset) SetVisibility(v string) {
	o.Visibility = &v
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise.
func (o *Asset) GetImpactLevel() string {
	if o == nil || o.ImpactLevel == nil {
		var ret string
		return ret
	}
	return *o.ImpactLevel
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetImpactLevelOk() (*string, bool) {
	if o == nil || o.ImpactLevel == nil {
		return nil, false
	}
	return o.ImpactLevel, true
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *Asset) HasImpactLevel() bool {
	if o != nil && o.ImpactLevel != nil {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given string and assigns it to the ImpactLevel field.
func (o *Asset) SetImpactLevel(v string) {
	o.ImpactLevel = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Asset) GetCategories() []Category {
	if o == nil || o.Categories == nil {
		var ret []Category
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCategoriesOk() (*[]Category, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Asset) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *Asset) SetCategories(v []Category) {
	o.Categories = &v
}

func (o Asset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
