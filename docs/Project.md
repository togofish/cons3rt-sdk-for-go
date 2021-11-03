# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveMemberCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**DefaultPowerSchedule** | Pointer to [**PowerSchedule**](PowerSchedule.md) |  | [optional] 
**DefaultRole** | Pointer to **string** |  | [optional] 
**DefaultVirtualizationRealm** | Pointer to [**VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**ProjectFeatures**](ProjectFeatures.md) |  | [optional] 
**Id** | **int32** |  | 
**ItarComment** | Pointer to **string** |  | [optional] 
**ItarRestricted** | Pointer to **bool** |  | [optional] 
**Limits** | [**ProjectLimits**](ProjectLimits.md) |  | 
**TotalMemberCount** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwningTeam** | Pointer to [**Team**](Team.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**TrustedProjects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**ResourceUsage** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**SubmissionServices** | Pointer to [**[]SubmissionService**](SubmissionService.md) |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**Members** | Pointer to [**[]User**](User.md) |  | [optional] 
**VirtualizationRealms** | Pointer to [**[]VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 

## Methods

### NewProject

`func NewProject(id int32, limits ProjectLimits, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveMemberCount

`func (o *Project) GetActiveMemberCount() int32`

GetActiveMemberCount returns the ActiveMemberCount field if non-nil, zero value otherwise.

### GetActiveMemberCountOk

`func (o *Project) GetActiveMemberCountOk() (*int32, bool)`

GetActiveMemberCountOk returns a tuple with the ActiveMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMemberCount

`func (o *Project) SetActiveMemberCount(v int32)`

SetActiveMemberCount sets ActiveMemberCount field to given value.

### HasActiveMemberCount

`func (o *Project) HasActiveMemberCount() bool`

HasActiveMemberCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Project) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Project) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultPowerSchedule

`func (o *Project) GetDefaultPowerSchedule() PowerSchedule`

GetDefaultPowerSchedule returns the DefaultPowerSchedule field if non-nil, zero value otherwise.

### GetDefaultPowerScheduleOk

`func (o *Project) GetDefaultPowerScheduleOk() (*PowerSchedule, bool)`

GetDefaultPowerScheduleOk returns a tuple with the DefaultPowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPowerSchedule

`func (o *Project) SetDefaultPowerSchedule(v PowerSchedule)`

SetDefaultPowerSchedule sets DefaultPowerSchedule field to given value.

### HasDefaultPowerSchedule

`func (o *Project) HasDefaultPowerSchedule() bool`

HasDefaultPowerSchedule returns a boolean if a field has been set.

### GetDefaultRole

`func (o *Project) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *Project) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *Project) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *Project) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDefaultVirtualizationRealm

`func (o *Project) GetDefaultVirtualizationRealm() VirtualizationRealm`

GetDefaultVirtualizationRealm returns the DefaultVirtualizationRealm field if non-nil, zero value otherwise.

### GetDefaultVirtualizationRealmOk

`func (o *Project) GetDefaultVirtualizationRealmOk() (*VirtualizationRealm, bool)`

GetDefaultVirtualizationRealmOk returns a tuple with the DefaultVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualizationRealm

`func (o *Project) SetDefaultVirtualizationRealm(v VirtualizationRealm)`

SetDefaultVirtualizationRealm sets DefaultVirtualizationRealm field to given value.

### HasDefaultVirtualizationRealm

`func (o *Project) HasDefaultVirtualizationRealm() bool`

HasDefaultVirtualizationRealm returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *Project) GetFeatures() ProjectFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Project) GetFeaturesOk() (*ProjectFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Project) SetFeatures(v ProjectFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Project) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int32)`

SetId sets Id field to given value.


### GetItarComment

`func (o *Project) GetItarComment() string`

GetItarComment returns the ItarComment field if non-nil, zero value otherwise.

### GetItarCommentOk

`func (o *Project) GetItarCommentOk() (*string, bool)`

GetItarCommentOk returns a tuple with the ItarComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItarComment

`func (o *Project) SetItarComment(v string)`

SetItarComment sets ItarComment field to given value.

### HasItarComment

`func (o *Project) HasItarComment() bool`

HasItarComment returns a boolean if a field has been set.

### GetItarRestricted

`func (o *Project) GetItarRestricted() bool`

GetItarRestricted returns the ItarRestricted field if non-nil, zero value otherwise.

### GetItarRestrictedOk

`func (o *Project) GetItarRestrictedOk() (*bool, bool)`

GetItarRestrictedOk returns a tuple with the ItarRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItarRestricted

`func (o *Project) SetItarRestricted(v bool)`

SetItarRestricted sets ItarRestricted field to given value.

### HasItarRestricted

`func (o *Project) HasItarRestricted() bool`

HasItarRestricted returns a boolean if a field has been set.

### GetLimits

`func (o *Project) GetLimits() ProjectLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Project) GetLimitsOk() (*ProjectLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Project) SetLimits(v ProjectLimits)`

SetLimits sets Limits field to given value.


### GetTotalMemberCount

`func (o *Project) GetTotalMemberCount() int32`

GetTotalMemberCount returns the TotalMemberCount field if non-nil, zero value otherwise.

### GetTotalMemberCountOk

`func (o *Project) GetTotalMemberCountOk() (*int32, bool)`

GetTotalMemberCountOk returns a tuple with the TotalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemberCount

`func (o *Project) SetTotalMemberCount(v int32)`

SetTotalMemberCount sets TotalMemberCount field to given value.

### HasTotalMemberCount

`func (o *Project) HasTotalMemberCount() bool`

HasTotalMemberCount returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwningTeam

`func (o *Project) GetOwningTeam() Team`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *Project) GetOwningTeamOk() (*Team, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *Project) SetOwningTeam(v Team)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *Project) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Project) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Project) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Project) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Project) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *Project) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *Project) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *Project) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *Project) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetResourceUsage

`func (o *Project) GetResourceUsage() ResourceUsage`

GetResourceUsage returns the ResourceUsage field if non-nil, zero value otherwise.

### GetResourceUsageOk

`func (o *Project) GetResourceUsageOk() (*ResourceUsage, bool)`

GetResourceUsageOk returns a tuple with the ResourceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUsage

`func (o *Project) SetResourceUsage(v ResourceUsage)`

SetResourceUsage sets ResourceUsage field to given value.

### HasResourceUsage

`func (o *Project) HasResourceUsage() bool`

HasResourceUsage returns a boolean if a field has been set.

### GetSubmissionServices

`func (o *Project) GetSubmissionServices() []SubmissionService`

GetSubmissionServices returns the SubmissionServices field if non-nil, zero value otherwise.

### GetSubmissionServicesOk

`func (o *Project) GetSubmissionServicesOk() (*[]SubmissionService, bool)`

GetSubmissionServicesOk returns a tuple with the SubmissionServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionServices

`func (o *Project) SetSubmissionServices(v []SubmissionService)`

SetSubmissionServices sets SubmissionServices field to given value.

### HasSubmissionServices

`func (o *Project) HasSubmissionServices() bool`

HasSubmissionServices returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Project) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Project) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMembers

`func (o *Project) GetMembers() []User`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Project) GetMembersOk() (*[]User, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Project) SetMembers(v []User)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Project) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetVirtualizationRealms

`func (o *Project) GetVirtualizationRealms() []VirtualizationRealm`

GetVirtualizationRealms returns the VirtualizationRealms field if non-nil, zero value otherwise.

### GetVirtualizationRealmsOk

`func (o *Project) GetVirtualizationRealmsOk() (*[]VirtualizationRealm, bool)`

GetVirtualizationRealmsOk returns a tuple with the VirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealms

`func (o *Project) SetVirtualizationRealms(v []VirtualizationRealm)`

SetVirtualizationRealms sets VirtualizationRealms field to given value.

### HasVirtualizationRealms

`func (o *Project) HasVirtualizationRealms() bool`

HasVirtualizationRealms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


