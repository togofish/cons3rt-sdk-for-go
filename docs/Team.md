# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetBundleInstallerEnabled** | Pointer to **bool** |  | [optional] 
**AssetBypassScanningEnabled** | Pointer to **bool** |  | [optional] 
**AvailabilityZoneEnabled** | Pointer to **bool** |  | [optional] 
**BypassScanningEnabled** | Pointer to **bool** |  | [optional] 
**LeadUser** | [**User**](User.md) |  | 
**GpuTypeMaximums** | Pointer to **map[string]int32** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ManagedVirtualizationRealms** | Pointer to [**[]VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**TeamManagers** | Pointer to [**[]User**](User.md) |  | [optional] 
**MaxAssets** | Pointer to **int32** |  | [optional] 
**MaxManagedVirtualizationRealms** | Pointer to **int32** |  | [optional] 
**MaxNumCpus** | Pointer to **int32** |  | [optional] 
**MaxNumGpus** | Pointer to **int32** |  | [optional] 
**MaxProjects** | Pointer to **int32** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int32** |  | [optional] 
**MaxSharedRemoteAccessSessions** | Pointer to **int32** |  | [optional] 
**MaxStorageInMegabytes** | Pointer to **int32** |  | [optional] 
**MaxUsers** | Pointer to **int32** |  | [optional] 
**MaxVirtualMachines** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**OrderNumber** | Pointer to **string** |  | [optional] 
**OwnedClouds** | Pointer to [**[]Cloud**](Cloud.md) |  | [optional] 
**OwnedProjects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**ContactInfo** | [**PocInfo**](PocInfo.md) |  | 
**PowerScheduleEnabled** | Pointer to **bool** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**RdpClientProxyEnabled** | Pointer to **bool** |  | [optional] 
**RdpClientSessionDuration** | Pointer to **int32** |  | [optional] 
**SnapshotEnabled** | Pointer to **bool** |  | [optional] 
**State** | **string** |  | 
**ValidUntil** | **int32** |  | 

## Methods

### NewTeam

`func NewTeam(leadUser User, name string, contactInfo PocInfo, state string, validUntil int32, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetBundleInstallerEnabled

`func (o *Team) GetAssetBundleInstallerEnabled() bool`

GetAssetBundleInstallerEnabled returns the AssetBundleInstallerEnabled field if non-nil, zero value otherwise.

### GetAssetBundleInstallerEnabledOk

`func (o *Team) GetAssetBundleInstallerEnabledOk() (*bool, bool)`

GetAssetBundleInstallerEnabledOk returns a tuple with the AssetBundleInstallerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetBundleInstallerEnabled

`func (o *Team) SetAssetBundleInstallerEnabled(v bool)`

SetAssetBundleInstallerEnabled sets AssetBundleInstallerEnabled field to given value.

### HasAssetBundleInstallerEnabled

`func (o *Team) HasAssetBundleInstallerEnabled() bool`

HasAssetBundleInstallerEnabled returns a boolean if a field has been set.

### GetAssetBypassScanningEnabled

`func (o *Team) GetAssetBypassScanningEnabled() bool`

GetAssetBypassScanningEnabled returns the AssetBypassScanningEnabled field if non-nil, zero value otherwise.

### GetAssetBypassScanningEnabledOk

`func (o *Team) GetAssetBypassScanningEnabledOk() (*bool, bool)`

GetAssetBypassScanningEnabledOk returns a tuple with the AssetBypassScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetBypassScanningEnabled

`func (o *Team) SetAssetBypassScanningEnabled(v bool)`

SetAssetBypassScanningEnabled sets AssetBypassScanningEnabled field to given value.

### HasAssetBypassScanningEnabled

`func (o *Team) HasAssetBypassScanningEnabled() bool`

HasAssetBypassScanningEnabled returns a boolean if a field has been set.

### GetAvailabilityZoneEnabled

`func (o *Team) GetAvailabilityZoneEnabled() bool`

GetAvailabilityZoneEnabled returns the AvailabilityZoneEnabled field if non-nil, zero value otherwise.

### GetAvailabilityZoneEnabledOk

`func (o *Team) GetAvailabilityZoneEnabledOk() (*bool, bool)`

GetAvailabilityZoneEnabledOk returns a tuple with the AvailabilityZoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneEnabled

`func (o *Team) SetAvailabilityZoneEnabled(v bool)`

SetAvailabilityZoneEnabled sets AvailabilityZoneEnabled field to given value.

### HasAvailabilityZoneEnabled

`func (o *Team) HasAvailabilityZoneEnabled() bool`

HasAvailabilityZoneEnabled returns a boolean if a field has been set.

### GetBypassScanningEnabled

`func (o *Team) GetBypassScanningEnabled() bool`

GetBypassScanningEnabled returns the BypassScanningEnabled field if non-nil, zero value otherwise.

### GetBypassScanningEnabledOk

`func (o *Team) GetBypassScanningEnabledOk() (*bool, bool)`

GetBypassScanningEnabledOk returns a tuple with the BypassScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassScanningEnabled

`func (o *Team) SetBypassScanningEnabled(v bool)`

SetBypassScanningEnabled sets BypassScanningEnabled field to given value.

### HasBypassScanningEnabled

`func (o *Team) HasBypassScanningEnabled() bool`

HasBypassScanningEnabled returns a boolean if a field has been set.

### GetLeadUser

`func (o *Team) GetLeadUser() User`

GetLeadUser returns the LeadUser field if non-nil, zero value otherwise.

### GetLeadUserOk

`func (o *Team) GetLeadUserOk() (*User, bool)`

GetLeadUserOk returns a tuple with the LeadUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadUser

`func (o *Team) SetLeadUser(v User)`

SetLeadUser sets LeadUser field to given value.


### GetGpuTypeMaximums

`func (o *Team) GetGpuTypeMaximums() map[string]int32`

GetGpuTypeMaximums returns the GpuTypeMaximums field if non-nil, zero value otherwise.

### GetGpuTypeMaximumsOk

`func (o *Team) GetGpuTypeMaximumsOk() (*map[string]int32, bool)`

GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypeMaximums

`func (o *Team) SetGpuTypeMaximums(v map[string]int32)`

SetGpuTypeMaximums sets GpuTypeMaximums field to given value.

### HasGpuTypeMaximums

`func (o *Team) HasGpuTypeMaximums() bool`

HasGpuTypeMaximums returns a boolean if a field has been set.

### GetIcon

`func (o *Team) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Team) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Team) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Team) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *Team) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedVirtualizationRealms

`func (o *Team) GetManagedVirtualizationRealms() []VirtualizationRealm`

GetManagedVirtualizationRealms returns the ManagedVirtualizationRealms field if non-nil, zero value otherwise.

### GetManagedVirtualizationRealmsOk

`func (o *Team) GetManagedVirtualizationRealmsOk() (*[]VirtualizationRealm, bool)`

GetManagedVirtualizationRealmsOk returns a tuple with the ManagedVirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedVirtualizationRealms

`func (o *Team) SetManagedVirtualizationRealms(v []VirtualizationRealm)`

SetManagedVirtualizationRealms sets ManagedVirtualizationRealms field to given value.

### HasManagedVirtualizationRealms

`func (o *Team) HasManagedVirtualizationRealms() bool`

HasManagedVirtualizationRealms returns a boolean if a field has been set.

### GetTeamManagers

`func (o *Team) GetTeamManagers() []User`

GetTeamManagers returns the TeamManagers field if non-nil, zero value otherwise.

### GetTeamManagersOk

`func (o *Team) GetTeamManagersOk() (*[]User, bool)`

GetTeamManagersOk returns a tuple with the TeamManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManagers

`func (o *Team) SetTeamManagers(v []User)`

SetTeamManagers sets TeamManagers field to given value.

### HasTeamManagers

`func (o *Team) HasTeamManagers() bool`

HasTeamManagers returns a boolean if a field has been set.

### GetMaxAssets

`func (o *Team) GetMaxAssets() int32`

GetMaxAssets returns the MaxAssets field if non-nil, zero value otherwise.

### GetMaxAssetsOk

`func (o *Team) GetMaxAssetsOk() (*int32, bool)`

GetMaxAssetsOk returns a tuple with the MaxAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAssets

`func (o *Team) SetMaxAssets(v int32)`

SetMaxAssets sets MaxAssets field to given value.

### HasMaxAssets

`func (o *Team) HasMaxAssets() bool`

HasMaxAssets returns a boolean if a field has been set.

### GetMaxManagedVirtualizationRealms

`func (o *Team) GetMaxManagedVirtualizationRealms() int32`

GetMaxManagedVirtualizationRealms returns the MaxManagedVirtualizationRealms field if non-nil, zero value otherwise.

### GetMaxManagedVirtualizationRealmsOk

`func (o *Team) GetMaxManagedVirtualizationRealmsOk() (*int32, bool)`

GetMaxManagedVirtualizationRealmsOk returns a tuple with the MaxManagedVirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManagedVirtualizationRealms

`func (o *Team) SetMaxManagedVirtualizationRealms(v int32)`

SetMaxManagedVirtualizationRealms sets MaxManagedVirtualizationRealms field to given value.

### HasMaxManagedVirtualizationRealms

`func (o *Team) HasMaxManagedVirtualizationRealms() bool`

HasMaxManagedVirtualizationRealms returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *Team) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *Team) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *Team) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *Team) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxNumGpus

`func (o *Team) GetMaxNumGpus() int32`

GetMaxNumGpus returns the MaxNumGpus field if non-nil, zero value otherwise.

### GetMaxNumGpusOk

`func (o *Team) GetMaxNumGpusOk() (*int32, bool)`

GetMaxNumGpusOk returns a tuple with the MaxNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumGpus

`func (o *Team) SetMaxNumGpus(v int32)`

SetMaxNumGpus sets MaxNumGpus field to given value.

### HasMaxNumGpus

`func (o *Team) HasMaxNumGpus() bool`

HasMaxNumGpus returns a boolean if a field has been set.

### GetMaxProjects

`func (o *Team) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *Team) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *Team) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *Team) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *Team) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *Team) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *Team) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *Team) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetMaxSharedRemoteAccessSessions

`func (o *Team) GetMaxSharedRemoteAccessSessions() int32`

GetMaxSharedRemoteAccessSessions returns the MaxSharedRemoteAccessSessions field if non-nil, zero value otherwise.

### GetMaxSharedRemoteAccessSessionsOk

`func (o *Team) GetMaxSharedRemoteAccessSessionsOk() (*int32, bool)`

GetMaxSharedRemoteAccessSessionsOk returns a tuple with the MaxSharedRemoteAccessSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSharedRemoteAccessSessions

`func (o *Team) SetMaxSharedRemoteAccessSessions(v int32)`

SetMaxSharedRemoteAccessSessions sets MaxSharedRemoteAccessSessions field to given value.

### HasMaxSharedRemoteAccessSessions

`func (o *Team) HasMaxSharedRemoteAccessSessions() bool`

HasMaxSharedRemoteAccessSessions returns a boolean if a field has been set.

### GetMaxStorageInMegabytes

`func (o *Team) GetMaxStorageInMegabytes() int32`

GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaxStorageInMegabytesOk

`func (o *Team) GetMaxStorageInMegabytesOk() (*int32, bool)`

GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorageInMegabytes

`func (o *Team) SetMaxStorageInMegabytes(v int32)`

SetMaxStorageInMegabytes sets MaxStorageInMegabytes field to given value.

### HasMaxStorageInMegabytes

`func (o *Team) HasMaxStorageInMegabytes() bool`

HasMaxStorageInMegabytes returns a boolean if a field has been set.

### GetMaxUsers

`func (o *Team) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *Team) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *Team) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *Team) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetMaxVirtualMachines

`func (o *Team) GetMaxVirtualMachines() int32`

GetMaxVirtualMachines returns the MaxVirtualMachines field if non-nil, zero value otherwise.

### GetMaxVirtualMachinesOk

`func (o *Team) GetMaxVirtualMachinesOk() (*int32, bool)`

GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualMachines

`func (o *Team) SetMaxVirtualMachines(v int32)`

SetMaxVirtualMachines sets MaxVirtualMachines field to given value.

### HasMaxVirtualMachines

`func (o *Team) HasMaxVirtualMachines() bool`

HasMaxVirtualMachines returns a boolean if a field has been set.

### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetOrderNumber

`func (o *Team) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *Team) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *Team) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *Team) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetOwnedClouds

`func (o *Team) GetOwnedClouds() []Cloud`

GetOwnedClouds returns the OwnedClouds field if non-nil, zero value otherwise.

### GetOwnedCloudsOk

`func (o *Team) GetOwnedCloudsOk() (*[]Cloud, bool)`

GetOwnedCloudsOk returns a tuple with the OwnedClouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedClouds

`func (o *Team) SetOwnedClouds(v []Cloud)`

SetOwnedClouds sets OwnedClouds field to given value.

### HasOwnedClouds

`func (o *Team) HasOwnedClouds() bool`

HasOwnedClouds returns a boolean if a field has been set.

### GetOwnedProjects

`func (o *Team) GetOwnedProjects() []Project`

GetOwnedProjects returns the OwnedProjects field if non-nil, zero value otherwise.

### GetOwnedProjectsOk

`func (o *Team) GetOwnedProjectsOk() (*[]Project, bool)`

GetOwnedProjectsOk returns a tuple with the OwnedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedProjects

`func (o *Team) SetOwnedProjects(v []Project)`

SetOwnedProjects sets OwnedProjects field to given value.

### HasOwnedProjects

`func (o *Team) HasOwnedProjects() bool`

HasOwnedProjects returns a boolean if a field has been set.

### GetContactInfo

`func (o *Team) GetContactInfo() PocInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *Team) GetContactInfoOk() (*PocInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *Team) SetContactInfo(v PocInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetPowerScheduleEnabled

`func (o *Team) GetPowerScheduleEnabled() bool`

GetPowerScheduleEnabled returns the PowerScheduleEnabled field if non-nil, zero value otherwise.

### GetPowerScheduleEnabledOk

`func (o *Team) GetPowerScheduleEnabledOk() (*bool, bool)`

GetPowerScheduleEnabledOk returns a tuple with the PowerScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleEnabled

`func (o *Team) SetPowerScheduleEnabled(v bool)`

SetPowerScheduleEnabled sets PowerScheduleEnabled field to given value.

### HasPowerScheduleEnabled

`func (o *Team) HasPowerScheduleEnabled() bool`

HasPowerScheduleEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *Team) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Team) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Team) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Team) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetRdpClientProxyEnabled

`func (o *Team) GetRdpClientProxyEnabled() bool`

GetRdpClientProxyEnabled returns the RdpClientProxyEnabled field if non-nil, zero value otherwise.

### GetRdpClientProxyEnabledOk

`func (o *Team) GetRdpClientProxyEnabledOk() (*bool, bool)`

GetRdpClientProxyEnabledOk returns a tuple with the RdpClientProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpClientProxyEnabled

`func (o *Team) SetRdpClientProxyEnabled(v bool)`

SetRdpClientProxyEnabled sets RdpClientProxyEnabled field to given value.

### HasRdpClientProxyEnabled

`func (o *Team) HasRdpClientProxyEnabled() bool`

HasRdpClientProxyEnabled returns a boolean if a field has been set.

### GetRdpClientSessionDuration

`func (o *Team) GetRdpClientSessionDuration() int32`

GetRdpClientSessionDuration returns the RdpClientSessionDuration field if non-nil, zero value otherwise.

### GetRdpClientSessionDurationOk

`func (o *Team) GetRdpClientSessionDurationOk() (*int32, bool)`

GetRdpClientSessionDurationOk returns a tuple with the RdpClientSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpClientSessionDuration

`func (o *Team) SetRdpClientSessionDuration(v int32)`

SetRdpClientSessionDuration sets RdpClientSessionDuration field to given value.

### HasRdpClientSessionDuration

`func (o *Team) HasRdpClientSessionDuration() bool`

HasRdpClientSessionDuration returns a boolean if a field has been set.

### GetSnapshotEnabled

`func (o *Team) GetSnapshotEnabled() bool`

GetSnapshotEnabled returns the SnapshotEnabled field if non-nil, zero value otherwise.

### GetSnapshotEnabledOk

`func (o *Team) GetSnapshotEnabledOk() (*bool, bool)`

GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnabled

`func (o *Team) SetSnapshotEnabled(v bool)`

SetSnapshotEnabled sets SnapshotEnabled field to given value.

### HasSnapshotEnabled

`func (o *Team) HasSnapshotEnabled() bool`

HasSnapshotEnabled returns a boolean if a field has been set.

### GetState

`func (o *Team) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Team) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Team) SetState(v string)`

SetState sets State field to given value.


### GetValidUntil

`func (o *Team) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *Team) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *Team) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


