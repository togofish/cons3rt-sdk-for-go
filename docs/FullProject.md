# FullProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ItarRestricted** | Pointer to **bool** |  | [optional] 
**ActiveMemberCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**DefaultPowerSchedule** | Pointer to [**PowerSchedule**](PowerSchedule.md) |  | [optional] 
**DefaultRole** | Pointer to **string** |  | [optional] 
**DefaultVirtualizationRealm** | Pointer to [**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**ProjectFeatures**](ProjectFeatures.md) |  | [optional] 
**ItarComment** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**ProjectLimits**](ProjectLimits.md) |  | [optional] 
**TotalMemberCount** | Pointer to **int32** |  | [optional] 
**OwningTeam** | Pointer to [**MinimalTeam**](MinimalTeam.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**TrustedProjects** | Pointer to [**[]MinimalProject**](MinimalProject.md) |  | [optional] 
**ResourceUsage** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**SubmissionServices** | Pointer to [**[]SubmissionService**](SubmissionService.md) |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**Members** | Pointer to [**[]MinimalUser**](MinimalUser.md) |  | [optional] 
**VirtualizationRealms** | Pointer to [**[]MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 

## Methods

### NewFullProject

`func NewFullProject(id int32, ) *FullProject`

NewFullProject instantiates a new FullProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullProjectWithDefaults

`func NewFullProjectWithDefaults() *FullProject`

NewFullProjectWithDefaults instantiates a new FullProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullProject) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FullProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItarRestricted

`func (o *FullProject) GetItarRestricted() bool`

GetItarRestricted returns the ItarRestricted field if non-nil, zero value otherwise.

### GetItarRestrictedOk

`func (o *FullProject) GetItarRestrictedOk() (*bool, bool)`

GetItarRestrictedOk returns a tuple with the ItarRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItarRestricted

`func (o *FullProject) SetItarRestricted(v bool)`

SetItarRestricted sets ItarRestricted field to given value.

### HasItarRestricted

`func (o *FullProject) HasItarRestricted() bool`

HasItarRestricted returns a boolean if a field has been set.

### GetActiveMemberCount

`func (o *FullProject) GetActiveMemberCount() int32`

GetActiveMemberCount returns the ActiveMemberCount field if non-nil, zero value otherwise.

### GetActiveMemberCountOk

`func (o *FullProject) GetActiveMemberCountOk() (*int32, bool)`

GetActiveMemberCountOk returns a tuple with the ActiveMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMemberCount

`func (o *FullProject) SetActiveMemberCount(v int32)`

SetActiveMemberCount sets ActiveMemberCount field to given value.

### HasActiveMemberCount

`func (o *FullProject) HasActiveMemberCount() bool`

HasActiveMemberCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FullProject) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullProject) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullProject) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultPowerSchedule

`func (o *FullProject) GetDefaultPowerSchedule() PowerSchedule`

GetDefaultPowerSchedule returns the DefaultPowerSchedule field if non-nil, zero value otherwise.

### GetDefaultPowerScheduleOk

`func (o *FullProject) GetDefaultPowerScheduleOk() (*PowerSchedule, bool)`

GetDefaultPowerScheduleOk returns a tuple with the DefaultPowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPowerSchedule

`func (o *FullProject) SetDefaultPowerSchedule(v PowerSchedule)`

SetDefaultPowerSchedule sets DefaultPowerSchedule field to given value.

### HasDefaultPowerSchedule

`func (o *FullProject) HasDefaultPowerSchedule() bool`

HasDefaultPowerSchedule returns a boolean if a field has been set.

### GetDefaultRole

`func (o *FullProject) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *FullProject) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *FullProject) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *FullProject) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDefaultVirtualizationRealm

`func (o *FullProject) GetDefaultVirtualizationRealm() MinimalVirtualizationRealm`

GetDefaultVirtualizationRealm returns the DefaultVirtualizationRealm field if non-nil, zero value otherwise.

### GetDefaultVirtualizationRealmOk

`func (o *FullProject) GetDefaultVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool)`

GetDefaultVirtualizationRealmOk returns a tuple with the DefaultVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualizationRealm

`func (o *FullProject) SetDefaultVirtualizationRealm(v MinimalVirtualizationRealm)`

SetDefaultVirtualizationRealm sets DefaultVirtualizationRealm field to given value.

### HasDefaultVirtualizationRealm

`func (o *FullProject) HasDefaultVirtualizationRealm() bool`

HasDefaultVirtualizationRealm returns a boolean if a field has been set.

### GetDescription

`func (o *FullProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *FullProject) GetFeatures() ProjectFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FullProject) GetFeaturesOk() (*ProjectFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FullProject) SetFeatures(v ProjectFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FullProject) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetItarComment

`func (o *FullProject) GetItarComment() string`

GetItarComment returns the ItarComment field if non-nil, zero value otherwise.

### GetItarCommentOk

`func (o *FullProject) GetItarCommentOk() (*string, bool)`

GetItarCommentOk returns a tuple with the ItarComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItarComment

`func (o *FullProject) SetItarComment(v string)`

SetItarComment sets ItarComment field to given value.

### HasItarComment

`func (o *FullProject) HasItarComment() bool`

HasItarComment returns a boolean if a field has been set.

### GetLimits

`func (o *FullProject) GetLimits() ProjectLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *FullProject) GetLimitsOk() (*ProjectLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *FullProject) SetLimits(v ProjectLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *FullProject) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTotalMemberCount

`func (o *FullProject) GetTotalMemberCount() int32`

GetTotalMemberCount returns the TotalMemberCount field if non-nil, zero value otherwise.

### GetTotalMemberCountOk

`func (o *FullProject) GetTotalMemberCountOk() (*int32, bool)`

GetTotalMemberCountOk returns a tuple with the TotalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemberCount

`func (o *FullProject) SetTotalMemberCount(v int32)`

SetTotalMemberCount sets TotalMemberCount field to given value.

### HasTotalMemberCount

`func (o *FullProject) HasTotalMemberCount() bool`

HasTotalMemberCount returns a boolean if a field has been set.

### GetOwningTeam

`func (o *FullProject) GetOwningTeam() MinimalTeam`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *FullProject) GetOwningTeamOk() (*MinimalTeam, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *FullProject) SetOwningTeam(v MinimalTeam)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *FullProject) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetIsPrivate

`func (o *FullProject) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *FullProject) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *FullProject) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *FullProject) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *FullProject) GetTrustedProjects() []MinimalProject`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *FullProject) GetTrustedProjectsOk() (*[]MinimalProject, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *FullProject) SetTrustedProjects(v []MinimalProject)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *FullProject) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetResourceUsage

`func (o *FullProject) GetResourceUsage() ResourceUsage`

GetResourceUsage returns the ResourceUsage field if non-nil, zero value otherwise.

### GetResourceUsageOk

`func (o *FullProject) GetResourceUsageOk() (*ResourceUsage, bool)`

GetResourceUsageOk returns a tuple with the ResourceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUsage

`func (o *FullProject) SetResourceUsage(v ResourceUsage)`

SetResourceUsage sets ResourceUsage field to given value.

### HasResourceUsage

`func (o *FullProject) HasResourceUsage() bool`

HasResourceUsage returns a boolean if a field has been set.

### GetSubmissionServices

`func (o *FullProject) GetSubmissionServices() []SubmissionService`

GetSubmissionServices returns the SubmissionServices field if non-nil, zero value otherwise.

### GetSubmissionServicesOk

`func (o *FullProject) GetSubmissionServicesOk() (*[]SubmissionService, bool)`

GetSubmissionServicesOk returns a tuple with the SubmissionServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionServices

`func (o *FullProject) SetSubmissionServices(v []SubmissionService)`

SetSubmissionServices sets SubmissionServices field to given value.

### HasSubmissionServices

`func (o *FullProject) HasSubmissionServices() bool`

HasSubmissionServices returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FullProject) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FullProject) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FullProject) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FullProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMembers

`func (o *FullProject) GetMembers() []MinimalUser`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *FullProject) GetMembersOk() (*[]MinimalUser, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *FullProject) SetMembers(v []MinimalUser)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *FullProject) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetVirtualizationRealms

`func (o *FullProject) GetVirtualizationRealms() []MinimalVirtualizationRealm`

GetVirtualizationRealms returns the VirtualizationRealms field if non-nil, zero value otherwise.

### GetVirtualizationRealmsOk

`func (o *FullProject) GetVirtualizationRealmsOk() (*[]MinimalVirtualizationRealm, bool)`

GetVirtualizationRealmsOk returns a tuple with the VirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealms

`func (o *FullProject) SetVirtualizationRealms(v []MinimalVirtualizationRealm)`

SetVirtualizationRealms sets VirtualizationRealms field to given value.

### HasVirtualizationRealms

`func (o *FullProject) HasVirtualizationRealms() bool`

HasVirtualizationRealms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


