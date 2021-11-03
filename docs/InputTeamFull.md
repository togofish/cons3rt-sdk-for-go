# InputTeamFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetBundleInstallerEnabled** | Pointer to **bool** |  | [optional] 
**AssetBypassScanningEnabled** | Pointer to **bool** |  | [optional] 
**AvailabilityZoneEnabled** | Pointer to **bool** |  | [optional] 
**BypassScanningEnabled** | Pointer to **bool** |  | [optional] 
**LeadUser** | [**InputUser**](InputUser.md) |  | 
**GpuTypeMaximums** | Pointer to **map[string]int32** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**TeamManagers** | [**[]InputUser**](InputUser.md) |  | 
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
**ContactInfo** | [**PocInfo**](PocInfo.md) |  | 
**PowerScheduleEnabled** | Pointer to **bool** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**RdpClientProxyEnabled** | Pointer to **bool** |  | [optional] 
**RdpClientSessionDuration** | Pointer to **int32** |  | [optional] 
**SnapshotEnabled** | Pointer to **bool** |  | [optional] 
**State** | **string** |  | 
**ValidUntil** | **int32** |  | 

## Methods

### NewInputTeamFull

`func NewInputTeamFull(leadUser InputUser, teamManagers []InputUser, name string, contactInfo PocInfo, state string, validUntil int32, ) *InputTeamFull`

NewInputTeamFull instantiates a new InputTeamFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTeamFullWithDefaults

`func NewInputTeamFullWithDefaults() *InputTeamFull`

NewInputTeamFullWithDefaults instantiates a new InputTeamFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetBundleInstallerEnabled

`func (o *InputTeamFull) GetAssetBundleInstallerEnabled() bool`

GetAssetBundleInstallerEnabled returns the AssetBundleInstallerEnabled field if non-nil, zero value otherwise.

### GetAssetBundleInstallerEnabledOk

`func (o *InputTeamFull) GetAssetBundleInstallerEnabledOk() (*bool, bool)`

GetAssetBundleInstallerEnabledOk returns a tuple with the AssetBundleInstallerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetBundleInstallerEnabled

`func (o *InputTeamFull) SetAssetBundleInstallerEnabled(v bool)`

SetAssetBundleInstallerEnabled sets AssetBundleInstallerEnabled field to given value.

### HasAssetBundleInstallerEnabled

`func (o *InputTeamFull) HasAssetBundleInstallerEnabled() bool`

HasAssetBundleInstallerEnabled returns a boolean if a field has been set.

### GetAssetBypassScanningEnabled

`func (o *InputTeamFull) GetAssetBypassScanningEnabled() bool`

GetAssetBypassScanningEnabled returns the AssetBypassScanningEnabled field if non-nil, zero value otherwise.

### GetAssetBypassScanningEnabledOk

`func (o *InputTeamFull) GetAssetBypassScanningEnabledOk() (*bool, bool)`

GetAssetBypassScanningEnabledOk returns a tuple with the AssetBypassScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetBypassScanningEnabled

`func (o *InputTeamFull) SetAssetBypassScanningEnabled(v bool)`

SetAssetBypassScanningEnabled sets AssetBypassScanningEnabled field to given value.

### HasAssetBypassScanningEnabled

`func (o *InputTeamFull) HasAssetBypassScanningEnabled() bool`

HasAssetBypassScanningEnabled returns a boolean if a field has been set.

### GetAvailabilityZoneEnabled

`func (o *InputTeamFull) GetAvailabilityZoneEnabled() bool`

GetAvailabilityZoneEnabled returns the AvailabilityZoneEnabled field if non-nil, zero value otherwise.

### GetAvailabilityZoneEnabledOk

`func (o *InputTeamFull) GetAvailabilityZoneEnabledOk() (*bool, bool)`

GetAvailabilityZoneEnabledOk returns a tuple with the AvailabilityZoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneEnabled

`func (o *InputTeamFull) SetAvailabilityZoneEnabled(v bool)`

SetAvailabilityZoneEnabled sets AvailabilityZoneEnabled field to given value.

### HasAvailabilityZoneEnabled

`func (o *InputTeamFull) HasAvailabilityZoneEnabled() bool`

HasAvailabilityZoneEnabled returns a boolean if a field has been set.

### GetBypassScanningEnabled

`func (o *InputTeamFull) GetBypassScanningEnabled() bool`

GetBypassScanningEnabled returns the BypassScanningEnabled field if non-nil, zero value otherwise.

### GetBypassScanningEnabledOk

`func (o *InputTeamFull) GetBypassScanningEnabledOk() (*bool, bool)`

GetBypassScanningEnabledOk returns a tuple with the BypassScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassScanningEnabled

`func (o *InputTeamFull) SetBypassScanningEnabled(v bool)`

SetBypassScanningEnabled sets BypassScanningEnabled field to given value.

### HasBypassScanningEnabled

`func (o *InputTeamFull) HasBypassScanningEnabled() bool`

HasBypassScanningEnabled returns a boolean if a field has been set.

### GetLeadUser

`func (o *InputTeamFull) GetLeadUser() InputUser`

GetLeadUser returns the LeadUser field if non-nil, zero value otherwise.

### GetLeadUserOk

`func (o *InputTeamFull) GetLeadUserOk() (*InputUser, bool)`

GetLeadUserOk returns a tuple with the LeadUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadUser

`func (o *InputTeamFull) SetLeadUser(v InputUser)`

SetLeadUser sets LeadUser field to given value.


### GetGpuTypeMaximums

`func (o *InputTeamFull) GetGpuTypeMaximums() map[string]int32`

GetGpuTypeMaximums returns the GpuTypeMaximums field if non-nil, zero value otherwise.

### GetGpuTypeMaximumsOk

`func (o *InputTeamFull) GetGpuTypeMaximumsOk() (*map[string]int32, bool)`

GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypeMaximums

`func (o *InputTeamFull) SetGpuTypeMaximums(v map[string]int32)`

SetGpuTypeMaximums sets GpuTypeMaximums field to given value.

### HasGpuTypeMaximums

`func (o *InputTeamFull) HasGpuTypeMaximums() bool`

HasGpuTypeMaximums returns a boolean if a field has been set.

### GetIcon

`func (o *InputTeamFull) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *InputTeamFull) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *InputTeamFull) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *InputTeamFull) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *InputTeamFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputTeamFull) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputTeamFull) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputTeamFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamManagers

`func (o *InputTeamFull) GetTeamManagers() []InputUser`

GetTeamManagers returns the TeamManagers field if non-nil, zero value otherwise.

### GetTeamManagersOk

`func (o *InputTeamFull) GetTeamManagersOk() (*[]InputUser, bool)`

GetTeamManagersOk returns a tuple with the TeamManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManagers

`func (o *InputTeamFull) SetTeamManagers(v []InputUser)`

SetTeamManagers sets TeamManagers field to given value.


### GetMaxManagedVirtualizationRealms

`func (o *InputTeamFull) GetMaxManagedVirtualizationRealms() int32`

GetMaxManagedVirtualizationRealms returns the MaxManagedVirtualizationRealms field if non-nil, zero value otherwise.

### GetMaxManagedVirtualizationRealmsOk

`func (o *InputTeamFull) GetMaxManagedVirtualizationRealmsOk() (*int32, bool)`

GetMaxManagedVirtualizationRealmsOk returns a tuple with the MaxManagedVirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManagedVirtualizationRealms

`func (o *InputTeamFull) SetMaxManagedVirtualizationRealms(v int32)`

SetMaxManagedVirtualizationRealms sets MaxManagedVirtualizationRealms field to given value.

### HasMaxManagedVirtualizationRealms

`func (o *InputTeamFull) HasMaxManagedVirtualizationRealms() bool`

HasMaxManagedVirtualizationRealms returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *InputTeamFull) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *InputTeamFull) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *InputTeamFull) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *InputTeamFull) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxNumGpus

`func (o *InputTeamFull) GetMaxNumGpus() int32`

GetMaxNumGpus returns the MaxNumGpus field if non-nil, zero value otherwise.

### GetMaxNumGpusOk

`func (o *InputTeamFull) GetMaxNumGpusOk() (*int32, bool)`

GetMaxNumGpusOk returns a tuple with the MaxNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumGpus

`func (o *InputTeamFull) SetMaxNumGpus(v int32)`

SetMaxNumGpus sets MaxNumGpus field to given value.

### HasMaxNumGpus

`func (o *InputTeamFull) HasMaxNumGpus() bool`

HasMaxNumGpus returns a boolean if a field has been set.

### GetMaxProjects

`func (o *InputTeamFull) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *InputTeamFull) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *InputTeamFull) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *InputTeamFull) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *InputTeamFull) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *InputTeamFull) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *InputTeamFull) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *InputTeamFull) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetMaxSharedRemoteAccessSessions

`func (o *InputTeamFull) GetMaxSharedRemoteAccessSessions() int32`

GetMaxSharedRemoteAccessSessions returns the MaxSharedRemoteAccessSessions field if non-nil, zero value otherwise.

### GetMaxSharedRemoteAccessSessionsOk

`func (o *InputTeamFull) GetMaxSharedRemoteAccessSessionsOk() (*int32, bool)`

GetMaxSharedRemoteAccessSessionsOk returns a tuple with the MaxSharedRemoteAccessSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSharedRemoteAccessSessions

`func (o *InputTeamFull) SetMaxSharedRemoteAccessSessions(v int32)`

SetMaxSharedRemoteAccessSessions sets MaxSharedRemoteAccessSessions field to given value.

### HasMaxSharedRemoteAccessSessions

`func (o *InputTeamFull) HasMaxSharedRemoteAccessSessions() bool`

HasMaxSharedRemoteAccessSessions returns a boolean if a field has been set.

### GetMaxStorageInMegabytes

`func (o *InputTeamFull) GetMaxStorageInMegabytes() int32`

GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaxStorageInMegabytesOk

`func (o *InputTeamFull) GetMaxStorageInMegabytesOk() (*int32, bool)`

GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorageInMegabytes

`func (o *InputTeamFull) SetMaxStorageInMegabytes(v int32)`

SetMaxStorageInMegabytes sets MaxStorageInMegabytes field to given value.

### HasMaxStorageInMegabytes

`func (o *InputTeamFull) HasMaxStorageInMegabytes() bool`

HasMaxStorageInMegabytes returns a boolean if a field has been set.

### GetMaxUsers

`func (o *InputTeamFull) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *InputTeamFull) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *InputTeamFull) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *InputTeamFull) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetMaxVirtualMachines

`func (o *InputTeamFull) GetMaxVirtualMachines() int32`

GetMaxVirtualMachines returns the MaxVirtualMachines field if non-nil, zero value otherwise.

### GetMaxVirtualMachinesOk

`func (o *InputTeamFull) GetMaxVirtualMachinesOk() (*int32, bool)`

GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualMachines

`func (o *InputTeamFull) SetMaxVirtualMachines(v int32)`

SetMaxVirtualMachines sets MaxVirtualMachines field to given value.

### HasMaxVirtualMachines

`func (o *InputTeamFull) HasMaxVirtualMachines() bool`

HasMaxVirtualMachines returns a boolean if a field has been set.

### GetName

`func (o *InputTeamFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputTeamFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputTeamFull) SetName(v string)`

SetName sets Name field to given value.


### GetOrderNumber

`func (o *InputTeamFull) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InputTeamFull) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InputTeamFull) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *InputTeamFull) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetContactInfo

`func (o *InputTeamFull) GetContactInfo() PocInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *InputTeamFull) GetContactInfoOk() (*PocInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *InputTeamFull) SetContactInfo(v PocInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetPowerScheduleEnabled

`func (o *InputTeamFull) GetPowerScheduleEnabled() bool`

GetPowerScheduleEnabled returns the PowerScheduleEnabled field if non-nil, zero value otherwise.

### GetPowerScheduleEnabledOk

`func (o *InputTeamFull) GetPowerScheduleEnabledOk() (*bool, bool)`

GetPowerScheduleEnabledOk returns a tuple with the PowerScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleEnabled

`func (o *InputTeamFull) SetPowerScheduleEnabled(v bool)`

SetPowerScheduleEnabled sets PowerScheduleEnabled field to given value.

### HasPowerScheduleEnabled

`func (o *InputTeamFull) HasPowerScheduleEnabled() bool`

HasPowerScheduleEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *InputTeamFull) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *InputTeamFull) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *InputTeamFull) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *InputTeamFull) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetRdpClientProxyEnabled

`func (o *InputTeamFull) GetRdpClientProxyEnabled() bool`

GetRdpClientProxyEnabled returns the RdpClientProxyEnabled field if non-nil, zero value otherwise.

### GetRdpClientProxyEnabledOk

`func (o *InputTeamFull) GetRdpClientProxyEnabledOk() (*bool, bool)`

GetRdpClientProxyEnabledOk returns a tuple with the RdpClientProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpClientProxyEnabled

`func (o *InputTeamFull) SetRdpClientProxyEnabled(v bool)`

SetRdpClientProxyEnabled sets RdpClientProxyEnabled field to given value.

### HasRdpClientProxyEnabled

`func (o *InputTeamFull) HasRdpClientProxyEnabled() bool`

HasRdpClientProxyEnabled returns a boolean if a field has been set.

### GetRdpClientSessionDuration

`func (o *InputTeamFull) GetRdpClientSessionDuration() int32`

GetRdpClientSessionDuration returns the RdpClientSessionDuration field if non-nil, zero value otherwise.

### GetRdpClientSessionDurationOk

`func (o *InputTeamFull) GetRdpClientSessionDurationOk() (*int32, bool)`

GetRdpClientSessionDurationOk returns a tuple with the RdpClientSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpClientSessionDuration

`func (o *InputTeamFull) SetRdpClientSessionDuration(v int32)`

SetRdpClientSessionDuration sets RdpClientSessionDuration field to given value.

### HasRdpClientSessionDuration

`func (o *InputTeamFull) HasRdpClientSessionDuration() bool`

HasRdpClientSessionDuration returns a boolean if a field has been set.

### GetSnapshotEnabled

`func (o *InputTeamFull) GetSnapshotEnabled() bool`

GetSnapshotEnabled returns the SnapshotEnabled field if non-nil, zero value otherwise.

### GetSnapshotEnabledOk

`func (o *InputTeamFull) GetSnapshotEnabledOk() (*bool, bool)`

GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnabled

`func (o *InputTeamFull) SetSnapshotEnabled(v bool)`

SetSnapshotEnabled sets SnapshotEnabled field to given value.

### HasSnapshotEnabled

`func (o *InputTeamFull) HasSnapshotEnabled() bool`

HasSnapshotEnabled returns a boolean if a field has been set.

### GetState

`func (o *InputTeamFull) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InputTeamFull) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InputTeamFull) SetState(v string)`

SetState sets State field to given value.


### GetValidUntil

`func (o *InputTeamFull) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *InputTeamFull) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *InputTeamFull) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


