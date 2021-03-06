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

// FullProject struct for FullProject
type FullProject struct {
	Id                         int32                         `json:"id"`
	Name                       *string                       `json:"name,omitempty"`
	ItarRestricted             *bool                         `json:"itarRestricted,omitempty"`
	ActiveMemberCount          *int32                        `json:"activeMemberCount,omitempty"`
	CreatedAt                  *int32                        `json:"createdAt,omitempty"`
	DefaultPowerSchedule       *PowerSchedule                `json:"defaultPowerSchedule,omitempty"`
	DefaultRole                *string                       `json:"defaultRole,omitempty"`
	DefaultVirtualizationRealm *MinimalVirtualizationRealm   `json:"defaultVirtualizationRealm,omitempty"`
	Description                *string                       `json:"description,omitempty"`
	Features                   *ProjectFeatures              `json:"features,omitempty"`
	ItarComment                *string                       `json:"itarComment,omitempty"`
	Limits                     *ProjectLimits                `json:"limits,omitempty"`
	TotalMemberCount           *int32                        `json:"totalMemberCount,omitempty"`
	OwningTeam                 *MinimalTeam                  `json:"owningTeam,omitempty"`
	IsPrivate                  *bool                         `json:"isPrivate,omitempty"`
	TrustedProjects            *[]MinimalProject             `json:"trustedProjects,omitempty"`
	ResourceUsage              *ResourceUsage                `json:"resourceUsage,omitempty"`
	SubmissionServices         *[]SubmissionService          `json:"submissionServices,omitempty"`
	UpdatedAt                  *int32                        `json:"updatedAt,omitempty"`
	Members                    *[]MinimalUser                `json:"members,omitempty"`
	VirtualizationRealms       *[]MinimalVirtualizationRealm `json:"virtualizationRealms,omitempty"`
}

// NewFullProject instantiates a new FullProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullProject(id int32) *FullProject {
	this := FullProject{}
	this.Id = id
	return &this
}

// NewFullProjectWithDefaults instantiates a new FullProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullProjectWithDefaults() *FullProject {
	this := FullProject{}
	return &this
}

// GetId returns the Id field value
func (o *FullProject) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FullProject) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FullProject) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullProject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullProject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullProject) SetName(v string) {
	o.Name = &v
}

// GetItarRestricted returns the ItarRestricted field value if set, zero value otherwise.
func (o *FullProject) GetItarRestricted() bool {
	if o == nil || o.ItarRestricted == nil {
		var ret bool
		return ret
	}
	return *o.ItarRestricted
}

// GetItarRestrictedOk returns a tuple with the ItarRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetItarRestrictedOk() (*bool, bool) {
	if o == nil || o.ItarRestricted == nil {
		return nil, false
	}
	return o.ItarRestricted, true
}

// HasItarRestricted returns a boolean if a field has been set.
func (o *FullProject) HasItarRestricted() bool {
	if o != nil && o.ItarRestricted != nil {
		return true
	}

	return false
}

// SetItarRestricted gets a reference to the given bool and assigns it to the ItarRestricted field.
func (o *FullProject) SetItarRestricted(v bool) {
	o.ItarRestricted = &v
}

// GetActiveMemberCount returns the ActiveMemberCount field value if set, zero value otherwise.
func (o *FullProject) GetActiveMemberCount() int32 {
	if o == nil || o.ActiveMemberCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveMemberCount
}

// GetActiveMemberCountOk returns a tuple with the ActiveMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetActiveMemberCountOk() (*int32, bool) {
	if o == nil || o.ActiveMemberCount == nil {
		return nil, false
	}
	return o.ActiveMemberCount, true
}

// HasActiveMemberCount returns a boolean if a field has been set.
func (o *FullProject) HasActiveMemberCount() bool {
	if o != nil && o.ActiveMemberCount != nil {
		return true
	}

	return false
}

// SetActiveMemberCount gets a reference to the given int32 and assigns it to the ActiveMemberCount field.
func (o *FullProject) SetActiveMemberCount(v int32) {
	o.ActiveMemberCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FullProject) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FullProject) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *FullProject) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDefaultPowerSchedule returns the DefaultPowerSchedule field value if set, zero value otherwise.
func (o *FullProject) GetDefaultPowerSchedule() PowerSchedule {
	if o == nil || o.DefaultPowerSchedule == nil {
		var ret PowerSchedule
		return ret
	}
	return *o.DefaultPowerSchedule
}

// GetDefaultPowerScheduleOk returns a tuple with the DefaultPowerSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetDefaultPowerScheduleOk() (*PowerSchedule, bool) {
	if o == nil || o.DefaultPowerSchedule == nil {
		return nil, false
	}
	return o.DefaultPowerSchedule, true
}

// HasDefaultPowerSchedule returns a boolean if a field has been set.
func (o *FullProject) HasDefaultPowerSchedule() bool {
	if o != nil && o.DefaultPowerSchedule != nil {
		return true
	}

	return false
}

// SetDefaultPowerSchedule gets a reference to the given PowerSchedule and assigns it to the DefaultPowerSchedule field.
func (o *FullProject) SetDefaultPowerSchedule(v PowerSchedule) {
	o.DefaultPowerSchedule = &v
}

// GetDefaultRole returns the DefaultRole field value if set, zero value otherwise.
func (o *FullProject) GetDefaultRole() string {
	if o == nil || o.DefaultRole == nil {
		var ret string
		return ret
	}
	return *o.DefaultRole
}

// GetDefaultRoleOk returns a tuple with the DefaultRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetDefaultRoleOk() (*string, bool) {
	if o == nil || o.DefaultRole == nil {
		return nil, false
	}
	return o.DefaultRole, true
}

// HasDefaultRole returns a boolean if a field has been set.
func (o *FullProject) HasDefaultRole() bool {
	if o != nil && o.DefaultRole != nil {
		return true
	}

	return false
}

// SetDefaultRole gets a reference to the given string and assigns it to the DefaultRole field.
func (o *FullProject) SetDefaultRole(v string) {
	o.DefaultRole = &v
}

// GetDefaultVirtualizationRealm returns the DefaultVirtualizationRealm field value if set, zero value otherwise.
func (o *FullProject) GetDefaultVirtualizationRealm() MinimalVirtualizationRealm {
	if o == nil || o.DefaultVirtualizationRealm == nil {
		var ret MinimalVirtualizationRealm
		return ret
	}
	return *o.DefaultVirtualizationRealm
}

// GetDefaultVirtualizationRealmOk returns a tuple with the DefaultVirtualizationRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetDefaultVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool) {
	if o == nil || o.DefaultVirtualizationRealm == nil {
		return nil, false
	}
	return o.DefaultVirtualizationRealm, true
}

// HasDefaultVirtualizationRealm returns a boolean if a field has been set.
func (o *FullProject) HasDefaultVirtualizationRealm() bool {
	if o != nil && o.DefaultVirtualizationRealm != nil {
		return true
	}

	return false
}

// SetDefaultVirtualizationRealm gets a reference to the given MinimalVirtualizationRealm and assigns it to the DefaultVirtualizationRealm field.
func (o *FullProject) SetDefaultVirtualizationRealm(v MinimalVirtualizationRealm) {
	o.DefaultVirtualizationRealm = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FullProject) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FullProject) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FullProject) SetDescription(v string) {
	o.Description = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FullProject) GetFeatures() ProjectFeatures {
	if o == nil || o.Features == nil {
		var ret ProjectFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetFeaturesOk() (*ProjectFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FullProject) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given ProjectFeatures and assigns it to the Features field.
func (o *FullProject) SetFeatures(v ProjectFeatures) {
	o.Features = &v
}

// GetItarComment returns the ItarComment field value if set, zero value otherwise.
func (o *FullProject) GetItarComment() string {
	if o == nil || o.ItarComment == nil {
		var ret string
		return ret
	}
	return *o.ItarComment
}

// GetItarCommentOk returns a tuple with the ItarComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetItarCommentOk() (*string, bool) {
	if o == nil || o.ItarComment == nil {
		return nil, false
	}
	return o.ItarComment, true
}

// HasItarComment returns a boolean if a field has been set.
func (o *FullProject) HasItarComment() bool {
	if o != nil && o.ItarComment != nil {
		return true
	}

	return false
}

// SetItarComment gets a reference to the given string and assigns it to the ItarComment field.
func (o *FullProject) SetItarComment(v string) {
	o.ItarComment = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *FullProject) GetLimits() ProjectLimits {
	if o == nil || o.Limits == nil {
		var ret ProjectLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetLimitsOk() (*ProjectLimits, bool) {
	if o == nil || o.Limits == nil {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *FullProject) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given ProjectLimits and assigns it to the Limits field.
func (o *FullProject) SetLimits(v ProjectLimits) {
	o.Limits = &v
}

// GetTotalMemberCount returns the TotalMemberCount field value if set, zero value otherwise.
func (o *FullProject) GetTotalMemberCount() int32 {
	if o == nil || o.TotalMemberCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalMemberCount
}

// GetTotalMemberCountOk returns a tuple with the TotalMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetTotalMemberCountOk() (*int32, bool) {
	if o == nil || o.TotalMemberCount == nil {
		return nil, false
	}
	return o.TotalMemberCount, true
}

// HasTotalMemberCount returns a boolean if a field has been set.
func (o *FullProject) HasTotalMemberCount() bool {
	if o != nil && o.TotalMemberCount != nil {
		return true
	}

	return false
}

// SetTotalMemberCount gets a reference to the given int32 and assigns it to the TotalMemberCount field.
func (o *FullProject) SetTotalMemberCount(v int32) {
	o.TotalMemberCount = &v
}

// GetOwningTeam returns the OwningTeam field value if set, zero value otherwise.
func (o *FullProject) GetOwningTeam() MinimalTeam {
	if o == nil || o.OwningTeam == nil {
		var ret MinimalTeam
		return ret
	}
	return *o.OwningTeam
}

// GetOwningTeamOk returns a tuple with the OwningTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetOwningTeamOk() (*MinimalTeam, bool) {
	if o == nil || o.OwningTeam == nil {
		return nil, false
	}
	return o.OwningTeam, true
}

// HasOwningTeam returns a boolean if a field has been set.
func (o *FullProject) HasOwningTeam() bool {
	if o != nil && o.OwningTeam != nil {
		return true
	}

	return false
}

// SetOwningTeam gets a reference to the given MinimalTeam and assigns it to the OwningTeam field.
func (o *FullProject) SetOwningTeam(v MinimalTeam) {
	o.OwningTeam = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *FullProject) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *FullProject) HasIsPrivate() bool {
	if o != nil && o.IsPrivate != nil {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *FullProject) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetTrustedProjects returns the TrustedProjects field value if set, zero value otherwise.
func (o *FullProject) GetTrustedProjects() []MinimalProject {
	if o == nil || o.TrustedProjects == nil {
		var ret []MinimalProject
		return ret
	}
	return *o.TrustedProjects
}

// GetTrustedProjectsOk returns a tuple with the TrustedProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetTrustedProjectsOk() (*[]MinimalProject, bool) {
	if o == nil || o.TrustedProjects == nil {
		return nil, false
	}
	return o.TrustedProjects, true
}

// HasTrustedProjects returns a boolean if a field has been set.
func (o *FullProject) HasTrustedProjects() bool {
	if o != nil && o.TrustedProjects != nil {
		return true
	}

	return false
}

// SetTrustedProjects gets a reference to the given []MinimalProject and assigns it to the TrustedProjects field.
func (o *FullProject) SetTrustedProjects(v []MinimalProject) {
	o.TrustedProjects = &v
}

// GetResourceUsage returns the ResourceUsage field value if set, zero value otherwise.
func (o *FullProject) GetResourceUsage() ResourceUsage {
	if o == nil || o.ResourceUsage == nil {
		var ret ResourceUsage
		return ret
	}
	return *o.ResourceUsage
}

// GetResourceUsageOk returns a tuple with the ResourceUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetResourceUsageOk() (*ResourceUsage, bool) {
	if o == nil || o.ResourceUsage == nil {
		return nil, false
	}
	return o.ResourceUsage, true
}

// HasResourceUsage returns a boolean if a field has been set.
func (o *FullProject) HasResourceUsage() bool {
	if o != nil && o.ResourceUsage != nil {
		return true
	}

	return false
}

// SetResourceUsage gets a reference to the given ResourceUsage and assigns it to the ResourceUsage field.
func (o *FullProject) SetResourceUsage(v ResourceUsage) {
	o.ResourceUsage = &v
}

// GetSubmissionServices returns the SubmissionServices field value if set, zero value otherwise.
func (o *FullProject) GetSubmissionServices() []SubmissionService {
	if o == nil || o.SubmissionServices == nil {
		var ret []SubmissionService
		return ret
	}
	return *o.SubmissionServices
}

// GetSubmissionServicesOk returns a tuple with the SubmissionServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetSubmissionServicesOk() (*[]SubmissionService, bool) {
	if o == nil || o.SubmissionServices == nil {
		return nil, false
	}
	return o.SubmissionServices, true
}

// HasSubmissionServices returns a boolean if a field has been set.
func (o *FullProject) HasSubmissionServices() bool {
	if o != nil && o.SubmissionServices != nil {
		return true
	}

	return false
}

// SetSubmissionServices gets a reference to the given []SubmissionService and assigns it to the SubmissionServices field.
func (o *FullProject) SetSubmissionServices(v []SubmissionService) {
	o.SubmissionServices = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FullProject) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FullProject) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *FullProject) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *FullProject) GetMembers() []MinimalUser {
	if o == nil || o.Members == nil {
		var ret []MinimalUser
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetMembersOk() (*[]MinimalUser, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *FullProject) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MinimalUser and assigns it to the Members field.
func (o *FullProject) SetMembers(v []MinimalUser) {
	o.Members = &v
}

// GetVirtualizationRealms returns the VirtualizationRealms field value if set, zero value otherwise.
func (o *FullProject) GetVirtualizationRealms() []MinimalVirtualizationRealm {
	if o == nil || o.VirtualizationRealms == nil {
		var ret []MinimalVirtualizationRealm
		return ret
	}
	return *o.VirtualizationRealms
}

// GetVirtualizationRealmsOk returns a tuple with the VirtualizationRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullProject) GetVirtualizationRealmsOk() (*[]MinimalVirtualizationRealm, bool) {
	if o == nil || o.VirtualizationRealms == nil {
		return nil, false
	}
	return o.VirtualizationRealms, true
}

// HasVirtualizationRealms returns a boolean if a field has been set.
func (o *FullProject) HasVirtualizationRealms() bool {
	if o != nil && o.VirtualizationRealms != nil {
		return true
	}

	return false
}

// SetVirtualizationRealms gets a reference to the given []MinimalVirtualizationRealm and assigns it to the VirtualizationRealms field.
func (o *FullProject) SetVirtualizationRealms(v []MinimalVirtualizationRealm) {
	o.VirtualizationRealms = &v
}

func (o FullProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ItarRestricted != nil {
		toSerialize["itarRestricted"] = o.ItarRestricted
	}
	if o.ActiveMemberCount != nil {
		toSerialize["activeMemberCount"] = o.ActiveMemberCount
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DefaultPowerSchedule != nil {
		toSerialize["defaultPowerSchedule"] = o.DefaultPowerSchedule
	}
	if o.DefaultRole != nil {
		toSerialize["defaultRole"] = o.DefaultRole
	}
	if o.DefaultVirtualizationRealm != nil {
		toSerialize["defaultVirtualizationRealm"] = o.DefaultVirtualizationRealm
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.ItarComment != nil {
		toSerialize["itarComment"] = o.ItarComment
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.TotalMemberCount != nil {
		toSerialize["totalMemberCount"] = o.TotalMemberCount
	}
	if o.OwningTeam != nil {
		toSerialize["owningTeam"] = o.OwningTeam
	}
	if o.IsPrivate != nil {
		toSerialize["isPrivate"] = o.IsPrivate
	}
	if o.TrustedProjects != nil {
		toSerialize["trustedProjects"] = o.TrustedProjects
	}
	if o.ResourceUsage != nil {
		toSerialize["resourceUsage"] = o.ResourceUsage
	}
	if o.SubmissionServices != nil {
		toSerialize["submissionServices"] = o.SubmissionServices
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.VirtualizationRealms != nil {
		toSerialize["virtualizationRealms"] = o.VirtualizationRealms
	}
	return json.Marshal(toSerialize)
}

type NullableFullProject struct {
	value *FullProject
	isSet bool
}

func (v NullableFullProject) Get() *FullProject {
	return v.value
}

func (v *NullableFullProject) Set(val *FullProject) {
	v.value = val
	v.isSet = true
}

func (v NullableFullProject) IsSet() bool {
	return v.isSet
}

func (v *NullableFullProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullProject(val *FullProject) *NullableFullProject {
	return &NullableFullProject{value: val, isSet: true}
}

func (v NullableFullProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
