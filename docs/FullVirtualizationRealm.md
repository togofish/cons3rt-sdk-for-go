# FullVirtualizationRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualizationRealmType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**AccessPoint** | Pointer to **string** |  | [optional] 
**AccountId** | **string** |  | 
**ActiveVirtualMachines** | Pointer to **int32** |  | [optional] 
**Networks** | [**[]Network**](Network.md) |  | 
**AdminUsers** | Pointer to [**[]MinimalUser**](MinimalUser.md) |  | [optional] 
**Allocated** | Pointer to **bool** |  | [optional] 
**Cloud** | Pointer to [**MinimalCloud**](MinimalCloud.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**DateLastReachable** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  | 
**DefaultWindowsDomainName** | Pointer to **string** |  | [optional] 
**LocalStorageName** | Pointer to **string** |  | [optional] 
**MaximumImpactLevel** | Pointer to **string** |  | [optional] 
**MaximumNumCpus** | Pointer to **int32** |  | [optional] 
**MaximumNumGpus** | Pointer to **int32** |  | [optional] 
**MaximumRamInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumStorageInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumVirtualMachines** | Pointer to **int32** |  | [optional] 
**PowerOnDelayBase** | Pointer to **int64** |  | [optional] 
**PowerOnInitialDelayBase** | Pointer to **int32** |  | [optional] 
**PowerOnMinimumDelay** | Pointer to **int32** |  | [optional] 
**Projects** | Pointer to [**[]MinimalProject**](MinimalProject.md) |  | [optional] 
**Reachable** | Pointer to **bool** |  | [optional] 
**RemoteAccessConfig** | Pointer to [**RemoteAccessConfig**](RemoteAccessConfig.md) |  | [optional] 
**RemoteAccessDeploymentId** | Pointer to **int32** |  | [optional] 
**RemoteAccessDeploymentRunStatus** | Pointer to **string** |  | [optional] 
**RemoteAccessStatus** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **[]string** |  | [optional] 
**TemplateRegistrations** | Pointer to [**[]MinimalTemplateRegistration**](MinimalTemplateRegistration.md) |  | [optional] 
**Templates** | Pointer to [**[]MinimalCons3rtTemplateData**](MinimalCons3rtTemplateData.md) |  | [optional] 
**TemplateSubscriptions** | Pointer to [**[]MinimalTemplateSubscription**](MinimalTemplateSubscription.md) |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  | 
**ZoneCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewFullVirtualizationRealm

`func NewFullVirtualizationRealm(name string, accountId string, networks []Network, description string, username string, ) *FullVirtualizationRealm`

NewFullVirtualizationRealm instantiates a new FullVirtualizationRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullVirtualizationRealmWithDefaults

`func NewFullVirtualizationRealmWithDefaults() *FullVirtualizationRealm`

NewFullVirtualizationRealmWithDefaults instantiates a new FullVirtualizationRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualizationRealmType

`func (o *FullVirtualizationRealm) GetVirtualizationRealmType() string`

GetVirtualizationRealmType returns the VirtualizationRealmType field if non-nil, zero value otherwise.

### GetVirtualizationRealmTypeOk

`func (o *FullVirtualizationRealm) GetVirtualizationRealmTypeOk() (*string, bool)`

GetVirtualizationRealmTypeOk returns a tuple with the VirtualizationRealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmType

`func (o *FullVirtualizationRealm) SetVirtualizationRealmType(v string)`

SetVirtualizationRealmType sets VirtualizationRealmType field to given value.

### HasVirtualizationRealmType

`func (o *FullVirtualizationRealm) HasVirtualizationRealmType() bool`

HasVirtualizationRealmType returns a boolean if a field has been set.

### GetId

`func (o *FullVirtualizationRealm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullVirtualizationRealm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullVirtualizationRealm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullVirtualizationRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullVirtualizationRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullVirtualizationRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullVirtualizationRealm) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *FullVirtualizationRealm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullVirtualizationRealm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullVirtualizationRealm) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullVirtualizationRealm) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAccessPoint

`func (o *FullVirtualizationRealm) GetAccessPoint() string`

GetAccessPoint returns the AccessPoint field if non-nil, zero value otherwise.

### GetAccessPointOk

`func (o *FullVirtualizationRealm) GetAccessPointOk() (*string, bool)`

GetAccessPointOk returns a tuple with the AccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoint

`func (o *FullVirtualizationRealm) SetAccessPoint(v string)`

SetAccessPoint sets AccessPoint field to given value.

### HasAccessPoint

`func (o *FullVirtualizationRealm) HasAccessPoint() bool`

HasAccessPoint returns a boolean if a field has been set.

### GetAccountId

`func (o *FullVirtualizationRealm) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FullVirtualizationRealm) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *FullVirtualizationRealm) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetActiveVirtualMachines

`func (o *FullVirtualizationRealm) GetActiveVirtualMachines() int32`

GetActiveVirtualMachines returns the ActiveVirtualMachines field if non-nil, zero value otherwise.

### GetActiveVirtualMachinesOk

`func (o *FullVirtualizationRealm) GetActiveVirtualMachinesOk() (*int32, bool)`

GetActiveVirtualMachinesOk returns a tuple with the ActiveVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVirtualMachines

`func (o *FullVirtualizationRealm) SetActiveVirtualMachines(v int32)`

SetActiveVirtualMachines sets ActiveVirtualMachines field to given value.

### HasActiveVirtualMachines

`func (o *FullVirtualizationRealm) HasActiveVirtualMachines() bool`

HasActiveVirtualMachines returns a boolean if a field has been set.

### GetNetworks

`func (o *FullVirtualizationRealm) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *FullVirtualizationRealm) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *FullVirtualizationRealm) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.


### GetAdminUsers

`func (o *FullVirtualizationRealm) GetAdminUsers() []MinimalUser`

GetAdminUsers returns the AdminUsers field if non-nil, zero value otherwise.

### GetAdminUsersOk

`func (o *FullVirtualizationRealm) GetAdminUsersOk() (*[]MinimalUser, bool)`

GetAdminUsersOk returns a tuple with the AdminUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsers

`func (o *FullVirtualizationRealm) SetAdminUsers(v []MinimalUser)`

SetAdminUsers sets AdminUsers field to given value.

### HasAdminUsers

`func (o *FullVirtualizationRealm) HasAdminUsers() bool`

HasAdminUsers returns a boolean if a field has been set.

### GetAllocated

`func (o *FullVirtualizationRealm) GetAllocated() bool`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *FullVirtualizationRealm) GetAllocatedOk() (*bool, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *FullVirtualizationRealm) SetAllocated(v bool)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *FullVirtualizationRealm) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetCloud

`func (o *FullVirtualizationRealm) GetCloud() MinimalCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FullVirtualizationRealm) GetCloudOk() (*MinimalCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FullVirtualizationRealm) SetCloud(v MinimalCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FullVirtualizationRealm) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FullVirtualizationRealm) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullVirtualizationRealm) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullVirtualizationRealm) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullVirtualizationRealm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDateLastReachable

`func (o *FullVirtualizationRealm) GetDateLastReachable() int32`

GetDateLastReachable returns the DateLastReachable field if non-nil, zero value otherwise.

### GetDateLastReachableOk

`func (o *FullVirtualizationRealm) GetDateLastReachableOk() (*int32, bool)`

GetDateLastReachableOk returns a tuple with the DateLastReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastReachable

`func (o *FullVirtualizationRealm) SetDateLastReachable(v int32)`

SetDateLastReachable sets DateLastReachable field to given value.

### HasDateLastReachable

`func (o *FullVirtualizationRealm) HasDateLastReachable() bool`

HasDateLastReachable returns a boolean if a field has been set.

### GetDescription

`func (o *FullVirtualizationRealm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullVirtualizationRealm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullVirtualizationRealm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefaultWindowsDomainName

`func (o *FullVirtualizationRealm) GetDefaultWindowsDomainName() string`

GetDefaultWindowsDomainName returns the DefaultWindowsDomainName field if non-nil, zero value otherwise.

### GetDefaultWindowsDomainNameOk

`func (o *FullVirtualizationRealm) GetDefaultWindowsDomainNameOk() (*string, bool)`

GetDefaultWindowsDomainNameOk returns a tuple with the DefaultWindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWindowsDomainName

`func (o *FullVirtualizationRealm) SetDefaultWindowsDomainName(v string)`

SetDefaultWindowsDomainName sets DefaultWindowsDomainName field to given value.

### HasDefaultWindowsDomainName

`func (o *FullVirtualizationRealm) HasDefaultWindowsDomainName() bool`

HasDefaultWindowsDomainName returns a boolean if a field has been set.

### GetLocalStorageName

`func (o *FullVirtualizationRealm) GetLocalStorageName() string`

GetLocalStorageName returns the LocalStorageName field if non-nil, zero value otherwise.

### GetLocalStorageNameOk

`func (o *FullVirtualizationRealm) GetLocalStorageNameOk() (*string, bool)`

GetLocalStorageNameOk returns a tuple with the LocalStorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageName

`func (o *FullVirtualizationRealm) SetLocalStorageName(v string)`

SetLocalStorageName sets LocalStorageName field to given value.

### HasLocalStorageName

`func (o *FullVirtualizationRealm) HasLocalStorageName() bool`

HasLocalStorageName returns a boolean if a field has been set.

### GetMaximumImpactLevel

`func (o *FullVirtualizationRealm) GetMaximumImpactLevel() string`

GetMaximumImpactLevel returns the MaximumImpactLevel field if non-nil, zero value otherwise.

### GetMaximumImpactLevelOk

`func (o *FullVirtualizationRealm) GetMaximumImpactLevelOk() (*string, bool)`

GetMaximumImpactLevelOk returns a tuple with the MaximumImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumImpactLevel

`func (o *FullVirtualizationRealm) SetMaximumImpactLevel(v string)`

SetMaximumImpactLevel sets MaximumImpactLevel field to given value.

### HasMaximumImpactLevel

`func (o *FullVirtualizationRealm) HasMaximumImpactLevel() bool`

HasMaximumImpactLevel returns a boolean if a field has been set.

### GetMaximumNumCpus

`func (o *FullVirtualizationRealm) GetMaximumNumCpus() int32`

GetMaximumNumCpus returns the MaximumNumCpus field if non-nil, zero value otherwise.

### GetMaximumNumCpusOk

`func (o *FullVirtualizationRealm) GetMaximumNumCpusOk() (*int32, bool)`

GetMaximumNumCpusOk returns a tuple with the MaximumNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumCpus

`func (o *FullVirtualizationRealm) SetMaximumNumCpus(v int32)`

SetMaximumNumCpus sets MaximumNumCpus field to given value.

### HasMaximumNumCpus

`func (o *FullVirtualizationRealm) HasMaximumNumCpus() bool`

HasMaximumNumCpus returns a boolean if a field has been set.

### GetMaximumNumGpus

`func (o *FullVirtualizationRealm) GetMaximumNumGpus() int32`

GetMaximumNumGpus returns the MaximumNumGpus field if non-nil, zero value otherwise.

### GetMaximumNumGpusOk

`func (o *FullVirtualizationRealm) GetMaximumNumGpusOk() (*int32, bool)`

GetMaximumNumGpusOk returns a tuple with the MaximumNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumGpus

`func (o *FullVirtualizationRealm) SetMaximumNumGpus(v int32)`

SetMaximumNumGpus sets MaximumNumGpus field to given value.

### HasMaximumNumGpus

`func (o *FullVirtualizationRealm) HasMaximumNumGpus() bool`

HasMaximumNumGpus returns a boolean if a field has been set.

### GetMaximumRamInMegabytes

`func (o *FullVirtualizationRealm) GetMaximumRamInMegabytes() int32`

GetMaximumRamInMegabytes returns the MaximumRamInMegabytes field if non-nil, zero value otherwise.

### GetMaximumRamInMegabytesOk

`func (o *FullVirtualizationRealm) GetMaximumRamInMegabytesOk() (*int32, bool)`

GetMaximumRamInMegabytesOk returns a tuple with the MaximumRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRamInMegabytes

`func (o *FullVirtualizationRealm) SetMaximumRamInMegabytes(v int32)`

SetMaximumRamInMegabytes sets MaximumRamInMegabytes field to given value.

### HasMaximumRamInMegabytes

`func (o *FullVirtualizationRealm) HasMaximumRamInMegabytes() bool`

HasMaximumRamInMegabytes returns a boolean if a field has been set.

### GetMaximumStorageInMegabytes

`func (o *FullVirtualizationRealm) GetMaximumStorageInMegabytes() int32`

GetMaximumStorageInMegabytes returns the MaximumStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaximumStorageInMegabytesOk

`func (o *FullVirtualizationRealm) GetMaximumStorageInMegabytesOk() (*int32, bool)`

GetMaximumStorageInMegabytesOk returns a tuple with the MaximumStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumStorageInMegabytes

`func (o *FullVirtualizationRealm) SetMaximumStorageInMegabytes(v int32)`

SetMaximumStorageInMegabytes sets MaximumStorageInMegabytes field to given value.

### HasMaximumStorageInMegabytes

`func (o *FullVirtualizationRealm) HasMaximumStorageInMegabytes() bool`

HasMaximumStorageInMegabytes returns a boolean if a field has been set.

### GetMaximumVirtualMachines

`func (o *FullVirtualizationRealm) GetMaximumVirtualMachines() int32`

GetMaximumVirtualMachines returns the MaximumVirtualMachines field if non-nil, zero value otherwise.

### GetMaximumVirtualMachinesOk

`func (o *FullVirtualizationRealm) GetMaximumVirtualMachinesOk() (*int32, bool)`

GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualMachines

`func (o *FullVirtualizationRealm) SetMaximumVirtualMachines(v int32)`

SetMaximumVirtualMachines sets MaximumVirtualMachines field to given value.

### HasMaximumVirtualMachines

`func (o *FullVirtualizationRealm) HasMaximumVirtualMachines() bool`

HasMaximumVirtualMachines returns a boolean if a field has been set.

### GetPowerOnDelayBase

`func (o *FullVirtualizationRealm) GetPowerOnDelayBase() int64`

GetPowerOnDelayBase returns the PowerOnDelayBase field if non-nil, zero value otherwise.

### GetPowerOnDelayBaseOk

`func (o *FullVirtualizationRealm) GetPowerOnDelayBaseOk() (*int64, bool)`

GetPowerOnDelayBaseOk returns a tuple with the PowerOnDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayBase

`func (o *FullVirtualizationRealm) SetPowerOnDelayBase(v int64)`

SetPowerOnDelayBase sets PowerOnDelayBase field to given value.

### HasPowerOnDelayBase

`func (o *FullVirtualizationRealm) HasPowerOnDelayBase() bool`

HasPowerOnDelayBase returns a boolean if a field has been set.

### GetPowerOnInitialDelayBase

`func (o *FullVirtualizationRealm) GetPowerOnInitialDelayBase() int32`

GetPowerOnInitialDelayBase returns the PowerOnInitialDelayBase field if non-nil, zero value otherwise.

### GetPowerOnInitialDelayBaseOk

`func (o *FullVirtualizationRealm) GetPowerOnInitialDelayBaseOk() (*int32, bool)`

GetPowerOnInitialDelayBaseOk returns a tuple with the PowerOnInitialDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnInitialDelayBase

`func (o *FullVirtualizationRealm) SetPowerOnInitialDelayBase(v int32)`

SetPowerOnInitialDelayBase sets PowerOnInitialDelayBase field to given value.

### HasPowerOnInitialDelayBase

`func (o *FullVirtualizationRealm) HasPowerOnInitialDelayBase() bool`

HasPowerOnInitialDelayBase returns a boolean if a field has been set.

### GetPowerOnMinimumDelay

`func (o *FullVirtualizationRealm) GetPowerOnMinimumDelay() int32`

GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field if non-nil, zero value otherwise.

### GetPowerOnMinimumDelayOk

`func (o *FullVirtualizationRealm) GetPowerOnMinimumDelayOk() (*int32, bool)`

GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMinimumDelay

`func (o *FullVirtualizationRealm) SetPowerOnMinimumDelay(v int32)`

SetPowerOnMinimumDelay sets PowerOnMinimumDelay field to given value.

### HasPowerOnMinimumDelay

`func (o *FullVirtualizationRealm) HasPowerOnMinimumDelay() bool`

HasPowerOnMinimumDelay returns a boolean if a field has been set.

### GetProjects

`func (o *FullVirtualizationRealm) GetProjects() []MinimalProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FullVirtualizationRealm) GetProjectsOk() (*[]MinimalProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FullVirtualizationRealm) SetProjects(v []MinimalProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FullVirtualizationRealm) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetReachable

`func (o *FullVirtualizationRealm) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *FullVirtualizationRealm) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *FullVirtualizationRealm) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *FullVirtualizationRealm) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetRemoteAccessConfig

`func (o *FullVirtualizationRealm) GetRemoteAccessConfig() RemoteAccessConfig`

GetRemoteAccessConfig returns the RemoteAccessConfig field if non-nil, zero value otherwise.

### GetRemoteAccessConfigOk

`func (o *FullVirtualizationRealm) GetRemoteAccessConfigOk() (*RemoteAccessConfig, bool)`

GetRemoteAccessConfigOk returns a tuple with the RemoteAccessConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessConfig

`func (o *FullVirtualizationRealm) SetRemoteAccessConfig(v RemoteAccessConfig)`

SetRemoteAccessConfig sets RemoteAccessConfig field to given value.

### HasRemoteAccessConfig

`func (o *FullVirtualizationRealm) HasRemoteAccessConfig() bool`

HasRemoteAccessConfig returns a boolean if a field has been set.

### GetRemoteAccessDeploymentId

`func (o *FullVirtualizationRealm) GetRemoteAccessDeploymentId() int32`

GetRemoteAccessDeploymentId returns the RemoteAccessDeploymentId field if non-nil, zero value otherwise.

### GetRemoteAccessDeploymentIdOk

`func (o *FullVirtualizationRealm) GetRemoteAccessDeploymentIdOk() (*int32, bool)`

GetRemoteAccessDeploymentIdOk returns a tuple with the RemoteAccessDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessDeploymentId

`func (o *FullVirtualizationRealm) SetRemoteAccessDeploymentId(v int32)`

SetRemoteAccessDeploymentId sets RemoteAccessDeploymentId field to given value.

### HasRemoteAccessDeploymentId

`func (o *FullVirtualizationRealm) HasRemoteAccessDeploymentId() bool`

HasRemoteAccessDeploymentId returns a boolean if a field has been set.

### GetRemoteAccessDeploymentRunStatus

`func (o *FullVirtualizationRealm) GetRemoteAccessDeploymentRunStatus() string`

GetRemoteAccessDeploymentRunStatus returns the RemoteAccessDeploymentRunStatus field if non-nil, zero value otherwise.

### GetRemoteAccessDeploymentRunStatusOk

`func (o *FullVirtualizationRealm) GetRemoteAccessDeploymentRunStatusOk() (*string, bool)`

GetRemoteAccessDeploymentRunStatusOk returns a tuple with the RemoteAccessDeploymentRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessDeploymentRunStatus

`func (o *FullVirtualizationRealm) SetRemoteAccessDeploymentRunStatus(v string)`

SetRemoteAccessDeploymentRunStatus sets RemoteAccessDeploymentRunStatus field to given value.

### HasRemoteAccessDeploymentRunStatus

`func (o *FullVirtualizationRealm) HasRemoteAccessDeploymentRunStatus() bool`

HasRemoteAccessDeploymentRunStatus returns a boolean if a field has been set.

### GetRemoteAccessStatus

`func (o *FullVirtualizationRealm) GetRemoteAccessStatus() string`

GetRemoteAccessStatus returns the RemoteAccessStatus field if non-nil, zero value otherwise.

### GetRemoteAccessStatusOk

`func (o *FullVirtualizationRealm) GetRemoteAccessStatusOk() (*string, bool)`

GetRemoteAccessStatusOk returns a tuple with the RemoteAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessStatus

`func (o *FullVirtualizationRealm) SetRemoteAccessStatus(v string)`

SetRemoteAccessStatus sets RemoteAccessStatus field to given value.

### HasRemoteAccessStatus

`func (o *FullVirtualizationRealm) HasRemoteAccessStatus() bool`

HasRemoteAccessStatus returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *FullVirtualizationRealm) GetSupportedFeatures() []string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *FullVirtualizationRealm) GetSupportedFeaturesOk() (*[]string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *FullVirtualizationRealm) SetSupportedFeatures(v []string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *FullVirtualizationRealm) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetTemplateRegistrations

`func (o *FullVirtualizationRealm) GetTemplateRegistrations() []MinimalTemplateRegistration`

GetTemplateRegistrations returns the TemplateRegistrations field if non-nil, zero value otherwise.

### GetTemplateRegistrationsOk

`func (o *FullVirtualizationRealm) GetTemplateRegistrationsOk() (*[]MinimalTemplateRegistration, bool)`

GetTemplateRegistrationsOk returns a tuple with the TemplateRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegistrations

`func (o *FullVirtualizationRealm) SetTemplateRegistrations(v []MinimalTemplateRegistration)`

SetTemplateRegistrations sets TemplateRegistrations field to given value.

### HasTemplateRegistrations

`func (o *FullVirtualizationRealm) HasTemplateRegistrations() bool`

HasTemplateRegistrations returns a boolean if a field has been set.

### GetTemplates

`func (o *FullVirtualizationRealm) GetTemplates() []MinimalCons3rtTemplateData`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *FullVirtualizationRealm) GetTemplatesOk() (*[]MinimalCons3rtTemplateData, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *FullVirtualizationRealm) SetTemplates(v []MinimalCons3rtTemplateData)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *FullVirtualizationRealm) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTemplateSubscriptions

`func (o *FullVirtualizationRealm) GetTemplateSubscriptions() []MinimalTemplateSubscription`

GetTemplateSubscriptions returns the TemplateSubscriptions field if non-nil, zero value otherwise.

### GetTemplateSubscriptionsOk

`func (o *FullVirtualizationRealm) GetTemplateSubscriptionsOk() (*[]MinimalTemplateSubscription, bool)`

GetTemplateSubscriptionsOk returns a tuple with the TemplateSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSubscriptions

`func (o *FullVirtualizationRealm) SetTemplateSubscriptions(v []MinimalTemplateSubscription)`

SetTemplateSubscriptions sets TemplateSubscriptions field to given value.

### HasTemplateSubscriptions

`func (o *FullVirtualizationRealm) HasTemplateSubscriptions() bool`

HasTemplateSubscriptions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FullVirtualizationRealm) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FullVirtualizationRealm) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FullVirtualizationRealm) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FullVirtualizationRealm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *FullVirtualizationRealm) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FullVirtualizationRealm) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FullVirtualizationRealm) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetZoneCount

`func (o *FullVirtualizationRealm) GetZoneCount() int32`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *FullVirtualizationRealm) GetZoneCountOk() (*int32, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *FullVirtualizationRealm) SetZoneCount(v int32)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *FullVirtualizationRealm) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


