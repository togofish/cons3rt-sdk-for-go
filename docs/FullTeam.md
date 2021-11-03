# FullTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetBundleInstallerEnabled** | Pointer to **bool** |  | [optional] 
**AssetBypassScanningEnabled** | Pointer to **bool** |  | [optional] 
**AvailabilityZoneEnabled** | Pointer to **bool** |  | [optional] 
**BypassScanningEnabled** | Pointer to **bool** |  | [optional] 
**LeadUser** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 
**GpuTypeMaximums** | Pointer to **map[string]int32** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ManagedVirtualizationRealms** | Pointer to [**[]MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**TeamManagers** | Pointer to [**[]MinimalUser**](MinimalUser.md) |  | [optional] 
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
**Name** | Pointer to **string** |  | [optional] 
**OrderNumber** | Pointer to **string** |  | [optional] 
**OwnedClouds** | Pointer to [**[]MinimalCloud**](MinimalCloud.md) |  | [optional] 
**OwnedProjects** | Pointer to [**[]MinimalProject**](MinimalProject.md) |  | [optional] 
**ContactInfo** | [**PocInfo**](PocInfo.md) |  | 
**PowerScheduleEnabled** | Pointer to **bool** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**RdpClientProxyEnabled** | Pointer to **bool** |  | [optional] 
**RdpClientSessionDuration** | Pointer to **int32** |  | [optional] 
**SnapshotEnabled** | Pointer to **bool** |  | [optional] 
**State** | **string** |  | 
**ValidUtil** | Pointer to **int32** |  | [optional] 

## Methods

### NewFullTeam

`func NewFullTeam(contactInfo PocInfo, state string, ) *FullTeam`

NewFullTeam instantiates a new FullTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullTeamWithDefaults

`func NewFullTeamWithDefaults() *FullTeam`

NewFullTeamWithDefaults instantiates a new FullTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetBundleInstallerEnabled

`func (o *FullTeam) GetAssetBundleInstallerEnabled() bool`

GetAssetBundleInstallerEnabled returns the AssetBundleInstallerEnabled field if non-nil, zero value otherwise.

### GetAssetBundleInstallerEnabledOk

`func (o *FullTeam) GetAssetBundleInstallerEnabledOk() (*bool, bool)`

GetAssetBundleInstallerEnabledOk returns a tuple with the AssetBundleInstallerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetBundleInstallerEnabled

`func (o *FullTeam) SetAssetBundleInstallerEnabled(v bool)`

SetAssetBundleInstallerEnabled sets AssetBundleInstallerEnabled field to given value.

### HasAssetBundleInstallerEnabled

`func (o *FullTeam) HasAssetBundleInstallerEnabled() bool`

HasAssetBundleInstallerEnabled returns a boolean if a field has been set.

### GetAssetBypassScanningEnabled

`func (o *FullTeam) GetAssetBypassScanningEnabled() bool`

GetAssetBypassScanningEnabled returns the AssetBypassScanningEnabled field if non-nil, zero value otherwise.

### GetAssetBypassScanningEnabledOk

`func (o *FullTeam) GetAssetBypassScanningEnabledOk() (*bool, bool)`

GetAssetBypassScanningEnabledOk returns a tuple with the AssetBypassScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetBypassScanningEnabled

`func (o *FullTeam) SetAssetBypassScanningEnabled(v bool)`

SetAssetBypassScanningEnabled sets AssetBypassScanningEnabled field to given value.

### HasAssetBypassScanningEnabled

`func (o *FullTeam) HasAssetBypassScanningEnabled() bool`

HasAssetBypassScanningEnabled returns a boolean if a field has been set.

### GetAvailabilityZoneEnabled

`func (o *FullTeam) GetAvailabilityZoneEnabled() bool`

GetAvailabilityZoneEnabled returns the AvailabilityZoneEnabled field if non-nil, zero value otherwise.

### GetAvailabilityZoneEnabledOk

`func (o *FullTeam) GetAvailabilityZoneEnabledOk() (*bool, bool)`

GetAvailabilityZoneEnabledOk returns a tuple with the AvailabilityZoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneEnabled

`func (o *FullTeam) SetAvailabilityZoneEnabled(v bool)`

SetAvailabilityZoneEnabled sets AvailabilityZoneEnabled field to given value.

### HasAvailabilityZoneEnabled

`func (o *FullTeam) HasAvailabilityZoneEnabled() bool`

HasAvailabilityZoneEnabled returns a boolean if a field has been set.

### GetBypassScanningEnabled

`func (o *FullTeam) GetBypassScanningEnabled() bool`

GetBypassScanningEnabled returns the BypassScanningEnabled field if non-nil, zero value otherwise.

### GetBypassScanningEnabledOk

`func (o *FullTeam) GetBypassScanningEnabledOk() (*bool, bool)`

GetBypassScanningEnabledOk returns a tuple with the BypassScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassScanningEnabled

`func (o *FullTeam) SetBypassScanningEnabled(v bool)`

SetBypassScanningEnabled sets BypassScanningEnabled field to given value.

### HasBypassScanningEnabled

`func (o *FullTeam) HasBypassScanningEnabled() bool`

HasBypassScanningEnabled returns a boolean if a field has been set.

### GetLeadUser

`func (o *FullTeam) GetLeadUser() MinimalUser`

GetLeadUser returns the LeadUser field if non-nil, zero value otherwise.

### GetLeadUserOk

`func (o *FullTeam) GetLeadUserOk() (*MinimalUser, bool)`

GetLeadUserOk returns a tuple with the LeadUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadUser

`func (o *FullTeam) SetLeadUser(v MinimalUser)`

SetLeadUser sets LeadUser field to given value.

### HasLeadUser

`func (o *FullTeam) HasLeadUser() bool`

HasLeadUser returns a boolean if a field has been set.

### GetGpuTypeMaximums

`func (o *FullTeam) GetGpuTypeMaximums() map[string]int32`

GetGpuTypeMaximums returns the GpuTypeMaximums field if non-nil, zero value otherwise.

### GetGpuTypeMaximumsOk

`func (o *FullTeam) GetGpuTypeMaximumsOk() (*map[string]int32, bool)`

GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypeMaximums

`func (o *FullTeam) SetGpuTypeMaximums(v map[string]int32)`

SetGpuTypeMaximums sets GpuTypeMaximums field to given value.

### HasGpuTypeMaximums

`func (o *FullTeam) HasGpuTypeMaximums() bool`

HasGpuTypeMaximums returns a boolean if a field has been set.

### GetIcon

`func (o *FullTeam) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FullTeam) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FullTeam) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FullTeam) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *FullTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedVirtualizationRealms

`func (o *FullTeam) GetManagedVirtualizationRealms() []MinimalVirtualizationRealm`

GetManagedVirtualizationRealms returns the ManagedVirtualizationRealms field if non-nil, zero value otherwise.

### GetManagedVirtualizationRealmsOk

`func (o *FullTeam) GetManagedVirtualizationRealmsOk() (*[]MinimalVirtualizationRealm, bool)`

GetManagedVirtualizationRealmsOk returns a tuple with the ManagedVirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedVirtualizationRealms

`func (o *FullTeam) SetManagedVirtualizationRealms(v []MinimalVirtualizationRealm)`

SetManagedVirtualizationRealms sets ManagedVirtualizationRealms field to given value.

### HasManagedVirtualizationRealms

`func (o *FullTeam) HasManagedVirtualizationRealms() bool`

HasManagedVirtualizationRealms returns a boolean if a field has been set.

### GetTeamManagers

`func (o *FullTeam) GetTeamManagers() []MinimalUser`

GetTeamManagers returns the TeamManagers field if non-nil, zero value otherwise.

### GetTeamManagersOk

`func (o *FullTeam) GetTeamManagersOk() (*[]MinimalUser, bool)`

GetTeamManagersOk returns a tuple with the TeamManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManagers

`func (o *FullTeam) SetTeamManagers(v []MinimalUser)`

SetTeamManagers sets TeamManagers field to given value.

### HasTeamManagers

`func (o *FullTeam) HasTeamManagers() bool`

HasTeamManagers returns a boolean if a field has been set.

### GetMaxAssets

`func (o *FullTeam) GetMaxAssets() int32`

GetMaxAssets returns the MaxAssets field if non-nil, zero value otherwise.

### GetMaxAssetsOk

`func (o *FullTeam) GetMaxAssetsOk() (*int32, bool)`

GetMaxAssetsOk returns a tuple with the MaxAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAssets

`func (o *FullTeam) SetMaxAssets(v int32)`

SetMaxAssets sets MaxAssets field to given value.

### HasMaxAssets

`func (o *FullTeam) HasMaxAssets() bool`

HasMaxAssets returns a boolean if a field has been set.

### GetMaxManagedVirtualizationRealms

`func (o *FullTeam) GetMaxManagedVirtualizationRealms() int32`

GetMaxManagedVirtualizationRealms returns the MaxManagedVirtualizationRealms field if non-nil, zero value otherwise.

### GetMaxManagedVirtualizationRealmsOk

`func (o *FullTeam) GetMaxManagedVirtualizationRealmsOk() (*int32, bool)`

GetMaxManagedVirtualizationRealmsOk returns a tuple with the MaxManagedVirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManagedVirtualizationRealms

`func (o *FullTeam) SetMaxManagedVirtualizationRealms(v int32)`

SetMaxManagedVirtualizationRealms sets MaxManagedVirtualizationRealms field to given value.

### HasMaxManagedVirtualizationRealms

`func (o *FullTeam) HasMaxManagedVirtualizationRealms() bool`

HasMaxManagedVirtualizationRealms returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *FullTeam) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *FullTeam) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *FullTeam) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *FullTeam) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxNumGpus

`func (o *FullTeam) GetMaxNumGpus() int32`

GetMaxNumGpus returns the MaxNumGpus field if non-nil, zero value otherwise.

### GetMaxNumGpusOk

`func (o *FullTeam) GetMaxNumGpusOk() (*int32, bool)`

GetMaxNumGpusOk returns a tuple with the MaxNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumGpus

`func (o *FullTeam) SetMaxNumGpus(v int32)`

SetMaxNumGpus sets MaxNumGpus field to given value.

### HasMaxNumGpus

`func (o *FullTeam) HasMaxNumGpus() bool`

HasMaxNumGpus returns a boolean if a field has been set.

### GetMaxProjects

`func (o *FullTeam) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *FullTeam) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *FullTeam) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *FullTeam) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *FullTeam) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *FullTeam) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *FullTeam) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *FullTeam) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetMaxSharedRemoteAccessSessions

`func (o *FullTeam) GetMaxSharedRemoteAccessSessions() int32`

GetMaxSharedRemoteAccessSessions returns the MaxSharedRemoteAccessSessions field if non-nil, zero value otherwise.

### GetMaxSharedRemoteAccessSessionsOk

`func (o *FullTeam) GetMaxSharedRemoteAccessSessionsOk() (*int32, bool)`

GetMaxSharedRemoteAccessSessionsOk returns a tuple with the MaxSharedRemoteAccessSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSharedRemoteAccessSessions

`func (o *FullTeam) SetMaxSharedRemoteAccessSessions(v int32)`

SetMaxSharedRemoteAccessSessions sets MaxSharedRemoteAccessSessions field to given value.

### HasMaxSharedRemoteAccessSessions

`func (o *FullTeam) HasMaxSharedRemoteAccessSessions() bool`

HasMaxSharedRemoteAccessSessions returns a boolean if a field has been set.

### GetMaxStorageInMegabytes

`func (o *FullTeam) GetMaxStorageInMegabytes() int32`

GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaxStorageInMegabytesOk

`func (o *FullTeam) GetMaxStorageInMegabytesOk() (*int32, bool)`

GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorageInMegabytes

`func (o *FullTeam) SetMaxStorageInMegabytes(v int32)`

SetMaxStorageInMegabytes sets MaxStorageInMegabytes field to given value.

### HasMaxStorageInMegabytes

`func (o *FullTeam) HasMaxStorageInMegabytes() bool`

HasMaxStorageInMegabytes returns a boolean if a field has been set.

### GetMaxUsers

`func (o *FullTeam) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *FullTeam) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *FullTeam) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *FullTeam) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetMaxVirtualMachines

`func (o *FullTeam) GetMaxVirtualMachines() int32`

GetMaxVirtualMachines returns the MaxVirtualMachines field if non-nil, zero value otherwise.

### GetMaxVirtualMachinesOk

`func (o *FullTeam) GetMaxVirtualMachinesOk() (*int32, bool)`

GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualMachines

`func (o *FullTeam) SetMaxVirtualMachines(v int32)`

SetMaxVirtualMachines sets MaxVirtualMachines field to given value.

### HasMaxVirtualMachines

`func (o *FullTeam) HasMaxVirtualMachines() bool`

HasMaxVirtualMachines returns a boolean if a field has been set.

### GetName

`func (o *FullTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderNumber

`func (o *FullTeam) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *FullTeam) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *FullTeam) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *FullTeam) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetOwnedClouds

`func (o *FullTeam) GetOwnedClouds() []MinimalCloud`

GetOwnedClouds returns the OwnedClouds field if non-nil, zero value otherwise.

### GetOwnedCloudsOk

`func (o *FullTeam) GetOwnedCloudsOk() (*[]MinimalCloud, bool)`

GetOwnedCloudsOk returns a tuple with the OwnedClouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedClouds

`func (o *FullTeam) SetOwnedClouds(v []MinimalCloud)`

SetOwnedClouds sets OwnedClouds field to given value.

### HasOwnedClouds

`func (o *FullTeam) HasOwnedClouds() bool`

HasOwnedClouds returns a boolean if a field has been set.

### GetOwnedProjects

`func (o *FullTeam) GetOwnedProjects() []MinimalProject`

GetOwnedProjects returns the OwnedProjects field if non-nil, zero value otherwise.

### GetOwnedProjectsOk

`func (o *FullTeam) GetOwnedProjectsOk() (*[]MinimalProject, bool)`

GetOwnedProjectsOk returns a tuple with the OwnedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedProjects

`func (o *FullTeam) SetOwnedProjects(v []MinimalProject)`

SetOwnedProjects sets OwnedProjects field to given value.

### HasOwnedProjects

`func (o *FullTeam) HasOwnedProjects() bool`

HasOwnedProjects returns a boolean if a field has been set.

### GetContactInfo

`func (o *FullTeam) GetContactInfo() PocInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *FullTeam) GetContactInfoOk() (*PocInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *FullTeam) SetContactInfo(v PocInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetPowerScheduleEnabled

`func (o *FullTeam) GetPowerScheduleEnabled() bool`

GetPowerScheduleEnabled returns the PowerScheduleEnabled field if non-nil, zero value otherwise.

### GetPowerScheduleEnabledOk

`func (o *FullTeam) GetPowerScheduleEnabledOk() (*bool, bool)`

GetPowerScheduleEnabledOk returns a tuple with the PowerScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleEnabled

`func (o *FullTeam) SetPowerScheduleEnabled(v bool)`

SetPowerScheduleEnabled sets PowerScheduleEnabled field to given value.

### HasPowerScheduleEnabled

`func (o *FullTeam) HasPowerScheduleEnabled() bool`

HasPowerScheduleEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *FullTeam) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FullTeam) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FullTeam) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *FullTeam) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetRdpClientProxyEnabled

`func (o *FullTeam) GetRdpClientProxyEnabled() bool`

GetRdpClientProxyEnabled returns the RdpClientProxyEnabled field if non-nil, zero value otherwise.

### GetRdpClientProxyEnabledOk

`func (o *FullTeam) GetRdpClientProxyEnabledOk() (*bool, bool)`

GetRdpClientProxyEnabledOk returns a tuple with the RdpClientProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpClientProxyEnabled

`func (o *FullTeam) SetRdpClientProxyEnabled(v bool)`

SetRdpClientProxyEnabled sets RdpClientProxyEnabled field to given value.

### HasRdpClientProxyEnabled

`func (o *FullTeam) HasRdpClientProxyEnabled() bool`

HasRdpClientProxyEnabled returns a boolean if a field has been set.

### GetRdpClientSessionDuration

`func (o *FullTeam) GetRdpClientSessionDuration() int32`

GetRdpClientSessionDuration returns the RdpClientSessionDuration field if non-nil, zero value otherwise.

### GetRdpClientSessionDurationOk

`func (o *FullTeam) GetRdpClientSessionDurationOk() (*int32, bool)`

GetRdpClientSessionDurationOk returns a tuple with the RdpClientSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpClientSessionDuration

`func (o *FullTeam) SetRdpClientSessionDuration(v int32)`

SetRdpClientSessionDuration sets RdpClientSessionDuration field to given value.

### HasRdpClientSessionDuration

`func (o *FullTeam) HasRdpClientSessionDuration() bool`

HasRdpClientSessionDuration returns a boolean if a field has been set.

### GetSnapshotEnabled

`func (o *FullTeam) GetSnapshotEnabled() bool`

GetSnapshotEnabled returns the SnapshotEnabled field if non-nil, zero value otherwise.

### GetSnapshotEnabledOk

`func (o *FullTeam) GetSnapshotEnabledOk() (*bool, bool)`

GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnabled

`func (o *FullTeam) SetSnapshotEnabled(v bool)`

SetSnapshotEnabled sets SnapshotEnabled field to given value.

### HasSnapshotEnabled

`func (o *FullTeam) HasSnapshotEnabled() bool`

HasSnapshotEnabled returns a boolean if a field has been set.

### GetState

`func (o *FullTeam) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullTeam) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullTeam) SetState(v string)`

SetState sets State field to given value.


### GetValidUtil

`func (o *FullTeam) GetValidUtil() int32`

GetValidUtil returns the ValidUtil field if non-nil, zero value otherwise.

### GetValidUtilOk

`func (o *FullTeam) GetValidUtilOk() (*int32, bool)`

GetValidUtilOk returns a tuple with the ValidUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUtil

`func (o *FullTeam) SetValidUtil(v int32)`

SetValidUtil sets ValidUtil field to given value.

### HasValidUtil

`func (o *FullTeam) HasValidUtil() bool`

HasValidUtil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


