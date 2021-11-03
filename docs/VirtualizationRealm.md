# VirtualizationRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualizationRealmType** | Pointer to **string** |  | [optional] 
**AccessPoint** | Pointer to **string** |  | [optional] 
**AccountId** | **string** |  | 
**ActiveVirtualMachines** | Pointer to **int32** |  | [optional] 
**Networks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**AdminUsers** | Pointer to [**[]User**](User.md) |  | [optional] 
**Allocated** | Pointer to **bool** |  | [optional] 
**Cidr** | **string** |  | 
**ConnectedCloud** | Pointer to [**Cloud**](Cloud.md) |  | [optional] 
**Cloud** | Pointer to [**Cloud**](Cloud.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**DateLastReachable** | Pointer to **int32** |  | [optional] 
**DefaultWindowsDomainName** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LocalStorageName** | Pointer to **string** |  | [optional] 
**MaximumNumCpus** | Pointer to **int32** |  | [optional] 
**MaximumNumGpus** | Pointer to **int32** |  | [optional] 
**MaximumRamInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumStorageInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumVirtualMachines** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Password** | **string** |  | 
**PowerOnDelayBase** | Pointer to **int64** |  | [optional] 
**PowerOnInitialDelayBase** | Pointer to **int32** |  | [optional] 
**PowerOnMaximumDelay** | Pointer to **int32** |  | [optional] 
**PowerOnMinimumDelay** | Pointer to **int32** |  | [optional] 
**Projects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**Reachable** | Pointer to **bool** |  | [optional] 
**RemoteAccessConfig** | Pointer to [**RemoteAccessConfig**](RemoteAccessConfig.md) |  | [optional] 
**RemoteAccessDeploymentId** | Pointer to **int32** |  | [optional] 
**RemoteAccessDeploymentRunStatus** | Pointer to **string** |  | [optional] 
**RemoteAccessStatus** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **[]string** |  | [optional] 
**TemplateRegistrations** | Pointer to [**[]TemplateRegistration**](TemplateRegistration.md) |  | [optional] 
**Templates** | Pointer to [**[]Cons3rtTemplateData**](Cons3rtTemplateData.md) |  | [optional] 
**TemplateSubscriptions** | Pointer to [**[]TemplateSubscription**](TemplateSubscription.md) |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  | 
**ZoneCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewVirtualizationRealm

`func NewVirtualizationRealm(accountId string, cidr string, description string, name string, password string, username string, ) *VirtualizationRealm`

NewVirtualizationRealm instantiates a new VirtualizationRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationRealmWithDefaults

`func NewVirtualizationRealmWithDefaults() *VirtualizationRealm`

NewVirtualizationRealmWithDefaults instantiates a new VirtualizationRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualizationRealmType

`func (o *VirtualizationRealm) GetVirtualizationRealmType() string`

GetVirtualizationRealmType returns the VirtualizationRealmType field if non-nil, zero value otherwise.

### GetVirtualizationRealmTypeOk

`func (o *VirtualizationRealm) GetVirtualizationRealmTypeOk() (*string, bool)`

GetVirtualizationRealmTypeOk returns a tuple with the VirtualizationRealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmType

`func (o *VirtualizationRealm) SetVirtualizationRealmType(v string)`

SetVirtualizationRealmType sets VirtualizationRealmType field to given value.

### HasVirtualizationRealmType

`func (o *VirtualizationRealm) HasVirtualizationRealmType() bool`

HasVirtualizationRealmType returns a boolean if a field has been set.

### GetAccessPoint

`func (o *VirtualizationRealm) GetAccessPoint() string`

GetAccessPoint returns the AccessPoint field if non-nil, zero value otherwise.

### GetAccessPointOk

`func (o *VirtualizationRealm) GetAccessPointOk() (*string, bool)`

GetAccessPointOk returns a tuple with the AccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoint

`func (o *VirtualizationRealm) SetAccessPoint(v string)`

SetAccessPoint sets AccessPoint field to given value.

### HasAccessPoint

`func (o *VirtualizationRealm) HasAccessPoint() bool`

HasAccessPoint returns a boolean if a field has been set.

### GetAccountId

`func (o *VirtualizationRealm) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VirtualizationRealm) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VirtualizationRealm) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetActiveVirtualMachines

`func (o *VirtualizationRealm) GetActiveVirtualMachines() int32`

GetActiveVirtualMachines returns the ActiveVirtualMachines field if non-nil, zero value otherwise.

### GetActiveVirtualMachinesOk

`func (o *VirtualizationRealm) GetActiveVirtualMachinesOk() (*int32, bool)`

GetActiveVirtualMachinesOk returns a tuple with the ActiveVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVirtualMachines

`func (o *VirtualizationRealm) SetActiveVirtualMachines(v int32)`

SetActiveVirtualMachines sets ActiveVirtualMachines field to given value.

### HasActiveVirtualMachines

`func (o *VirtualizationRealm) HasActiveVirtualMachines() bool`

HasActiveVirtualMachines returns a boolean if a field has been set.

### GetNetworks

`func (o *VirtualizationRealm) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VirtualizationRealm) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VirtualizationRealm) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VirtualizationRealm) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetAdminUsers

`func (o *VirtualizationRealm) GetAdminUsers() []User`

GetAdminUsers returns the AdminUsers field if non-nil, zero value otherwise.

### GetAdminUsersOk

`func (o *VirtualizationRealm) GetAdminUsersOk() (*[]User, bool)`

GetAdminUsersOk returns a tuple with the AdminUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsers

`func (o *VirtualizationRealm) SetAdminUsers(v []User)`

SetAdminUsers sets AdminUsers field to given value.

### HasAdminUsers

`func (o *VirtualizationRealm) HasAdminUsers() bool`

HasAdminUsers returns a boolean if a field has been set.

### GetAllocated

`func (o *VirtualizationRealm) GetAllocated() bool`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *VirtualizationRealm) GetAllocatedOk() (*bool, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *VirtualizationRealm) SetAllocated(v bool)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *VirtualizationRealm) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetCidr

`func (o *VirtualizationRealm) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *VirtualizationRealm) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *VirtualizationRealm) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetConnectedCloud

`func (o *VirtualizationRealm) GetConnectedCloud() Cloud`

GetConnectedCloud returns the ConnectedCloud field if non-nil, zero value otherwise.

### GetConnectedCloudOk

`func (o *VirtualizationRealm) GetConnectedCloudOk() (*Cloud, bool)`

GetConnectedCloudOk returns a tuple with the ConnectedCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedCloud

`func (o *VirtualizationRealm) SetConnectedCloud(v Cloud)`

SetConnectedCloud sets ConnectedCloud field to given value.

### HasConnectedCloud

`func (o *VirtualizationRealm) HasConnectedCloud() bool`

HasConnectedCloud returns a boolean if a field has been set.

### GetCloud

`func (o *VirtualizationRealm) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *VirtualizationRealm) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *VirtualizationRealm) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *VirtualizationRealm) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VirtualizationRealm) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualizationRealm) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualizationRealm) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualizationRealm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDateLastReachable

`func (o *VirtualizationRealm) GetDateLastReachable() int32`

GetDateLastReachable returns the DateLastReachable field if non-nil, zero value otherwise.

### GetDateLastReachableOk

`func (o *VirtualizationRealm) GetDateLastReachableOk() (*int32, bool)`

GetDateLastReachableOk returns a tuple with the DateLastReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastReachable

`func (o *VirtualizationRealm) SetDateLastReachable(v int32)`

SetDateLastReachable sets DateLastReachable field to given value.

### HasDateLastReachable

`func (o *VirtualizationRealm) HasDateLastReachable() bool`

HasDateLastReachable returns a boolean if a field has been set.

### GetDefaultWindowsDomainName

`func (o *VirtualizationRealm) GetDefaultWindowsDomainName() string`

GetDefaultWindowsDomainName returns the DefaultWindowsDomainName field if non-nil, zero value otherwise.

### GetDefaultWindowsDomainNameOk

`func (o *VirtualizationRealm) GetDefaultWindowsDomainNameOk() (*string, bool)`

GetDefaultWindowsDomainNameOk returns a tuple with the DefaultWindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWindowsDomainName

`func (o *VirtualizationRealm) SetDefaultWindowsDomainName(v string)`

SetDefaultWindowsDomainName sets DefaultWindowsDomainName field to given value.

### HasDefaultWindowsDomainName

`func (o *VirtualizationRealm) HasDefaultWindowsDomainName() bool`

HasDefaultWindowsDomainName returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualizationRealm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationRealm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationRealm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *VirtualizationRealm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualizationRealm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualizationRealm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualizationRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalStorageName

`func (o *VirtualizationRealm) GetLocalStorageName() string`

GetLocalStorageName returns the LocalStorageName field if non-nil, zero value otherwise.

### GetLocalStorageNameOk

`func (o *VirtualizationRealm) GetLocalStorageNameOk() (*string, bool)`

GetLocalStorageNameOk returns a tuple with the LocalStorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageName

`func (o *VirtualizationRealm) SetLocalStorageName(v string)`

SetLocalStorageName sets LocalStorageName field to given value.

### HasLocalStorageName

`func (o *VirtualizationRealm) HasLocalStorageName() bool`

HasLocalStorageName returns a boolean if a field has been set.

### GetMaximumNumCpus

`func (o *VirtualizationRealm) GetMaximumNumCpus() int32`

GetMaximumNumCpus returns the MaximumNumCpus field if non-nil, zero value otherwise.

### GetMaximumNumCpusOk

`func (o *VirtualizationRealm) GetMaximumNumCpusOk() (*int32, bool)`

GetMaximumNumCpusOk returns a tuple with the MaximumNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumCpus

`func (o *VirtualizationRealm) SetMaximumNumCpus(v int32)`

SetMaximumNumCpus sets MaximumNumCpus field to given value.

### HasMaximumNumCpus

`func (o *VirtualizationRealm) HasMaximumNumCpus() bool`

HasMaximumNumCpus returns a boolean if a field has been set.

### GetMaximumNumGpus

`func (o *VirtualizationRealm) GetMaximumNumGpus() int32`

GetMaximumNumGpus returns the MaximumNumGpus field if non-nil, zero value otherwise.

### GetMaximumNumGpusOk

`func (o *VirtualizationRealm) GetMaximumNumGpusOk() (*int32, bool)`

GetMaximumNumGpusOk returns a tuple with the MaximumNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumGpus

`func (o *VirtualizationRealm) SetMaximumNumGpus(v int32)`

SetMaximumNumGpus sets MaximumNumGpus field to given value.

### HasMaximumNumGpus

`func (o *VirtualizationRealm) HasMaximumNumGpus() bool`

HasMaximumNumGpus returns a boolean if a field has been set.

### GetMaximumRamInMegabytes

`func (o *VirtualizationRealm) GetMaximumRamInMegabytes() int32`

GetMaximumRamInMegabytes returns the MaximumRamInMegabytes field if non-nil, zero value otherwise.

### GetMaximumRamInMegabytesOk

`func (o *VirtualizationRealm) GetMaximumRamInMegabytesOk() (*int32, bool)`

GetMaximumRamInMegabytesOk returns a tuple with the MaximumRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRamInMegabytes

`func (o *VirtualizationRealm) SetMaximumRamInMegabytes(v int32)`

SetMaximumRamInMegabytes sets MaximumRamInMegabytes field to given value.

### HasMaximumRamInMegabytes

`func (o *VirtualizationRealm) HasMaximumRamInMegabytes() bool`

HasMaximumRamInMegabytes returns a boolean if a field has been set.

### GetMaximumStorageInMegabytes

`func (o *VirtualizationRealm) GetMaximumStorageInMegabytes() int32`

GetMaximumStorageInMegabytes returns the MaximumStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaximumStorageInMegabytesOk

`func (o *VirtualizationRealm) GetMaximumStorageInMegabytesOk() (*int32, bool)`

GetMaximumStorageInMegabytesOk returns a tuple with the MaximumStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumStorageInMegabytes

`func (o *VirtualizationRealm) SetMaximumStorageInMegabytes(v int32)`

SetMaximumStorageInMegabytes sets MaximumStorageInMegabytes field to given value.

### HasMaximumStorageInMegabytes

`func (o *VirtualizationRealm) HasMaximumStorageInMegabytes() bool`

HasMaximumStorageInMegabytes returns a boolean if a field has been set.

### GetMaximumVirtualMachines

`func (o *VirtualizationRealm) GetMaximumVirtualMachines() int32`

GetMaximumVirtualMachines returns the MaximumVirtualMachines field if non-nil, zero value otherwise.

### GetMaximumVirtualMachinesOk

`func (o *VirtualizationRealm) GetMaximumVirtualMachinesOk() (*int32, bool)`

GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualMachines

`func (o *VirtualizationRealm) SetMaximumVirtualMachines(v int32)`

SetMaximumVirtualMachines sets MaximumVirtualMachines field to given value.

### HasMaximumVirtualMachines

`func (o *VirtualizationRealm) HasMaximumVirtualMachines() bool`

HasMaximumVirtualMachines returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationRealm) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *VirtualizationRealm) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VirtualizationRealm) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VirtualizationRealm) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPowerOnDelayBase

`func (o *VirtualizationRealm) GetPowerOnDelayBase() int64`

GetPowerOnDelayBase returns the PowerOnDelayBase field if non-nil, zero value otherwise.

### GetPowerOnDelayBaseOk

`func (o *VirtualizationRealm) GetPowerOnDelayBaseOk() (*int64, bool)`

GetPowerOnDelayBaseOk returns a tuple with the PowerOnDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayBase

`func (o *VirtualizationRealm) SetPowerOnDelayBase(v int64)`

SetPowerOnDelayBase sets PowerOnDelayBase field to given value.

### HasPowerOnDelayBase

`func (o *VirtualizationRealm) HasPowerOnDelayBase() bool`

HasPowerOnDelayBase returns a boolean if a field has been set.

### GetPowerOnInitialDelayBase

`func (o *VirtualizationRealm) GetPowerOnInitialDelayBase() int32`

GetPowerOnInitialDelayBase returns the PowerOnInitialDelayBase field if non-nil, zero value otherwise.

### GetPowerOnInitialDelayBaseOk

`func (o *VirtualizationRealm) GetPowerOnInitialDelayBaseOk() (*int32, bool)`

GetPowerOnInitialDelayBaseOk returns a tuple with the PowerOnInitialDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnInitialDelayBase

`func (o *VirtualizationRealm) SetPowerOnInitialDelayBase(v int32)`

SetPowerOnInitialDelayBase sets PowerOnInitialDelayBase field to given value.

### HasPowerOnInitialDelayBase

`func (o *VirtualizationRealm) HasPowerOnInitialDelayBase() bool`

HasPowerOnInitialDelayBase returns a boolean if a field has been set.

### GetPowerOnMaximumDelay

`func (o *VirtualizationRealm) GetPowerOnMaximumDelay() int32`

GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field if non-nil, zero value otherwise.

### GetPowerOnMaximumDelayOk

`func (o *VirtualizationRealm) GetPowerOnMaximumDelayOk() (*int32, bool)`

GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMaximumDelay

`func (o *VirtualizationRealm) SetPowerOnMaximumDelay(v int32)`

SetPowerOnMaximumDelay sets PowerOnMaximumDelay field to given value.

### HasPowerOnMaximumDelay

`func (o *VirtualizationRealm) HasPowerOnMaximumDelay() bool`

HasPowerOnMaximumDelay returns a boolean if a field has been set.

### GetPowerOnMinimumDelay

`func (o *VirtualizationRealm) GetPowerOnMinimumDelay() int32`

GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field if non-nil, zero value otherwise.

### GetPowerOnMinimumDelayOk

`func (o *VirtualizationRealm) GetPowerOnMinimumDelayOk() (*int32, bool)`

GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMinimumDelay

`func (o *VirtualizationRealm) SetPowerOnMinimumDelay(v int32)`

SetPowerOnMinimumDelay sets PowerOnMinimumDelay field to given value.

### HasPowerOnMinimumDelay

`func (o *VirtualizationRealm) HasPowerOnMinimumDelay() bool`

HasPowerOnMinimumDelay returns a boolean if a field has been set.

### GetProjects

`func (o *VirtualizationRealm) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *VirtualizationRealm) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *VirtualizationRealm) SetProjects(v []Project)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *VirtualizationRealm) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetReachable

`func (o *VirtualizationRealm) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *VirtualizationRealm) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *VirtualizationRealm) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *VirtualizationRealm) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetRemoteAccessConfig

`func (o *VirtualizationRealm) GetRemoteAccessConfig() RemoteAccessConfig`

GetRemoteAccessConfig returns the RemoteAccessConfig field if non-nil, zero value otherwise.

### GetRemoteAccessConfigOk

`func (o *VirtualizationRealm) GetRemoteAccessConfigOk() (*RemoteAccessConfig, bool)`

GetRemoteAccessConfigOk returns a tuple with the RemoteAccessConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessConfig

`func (o *VirtualizationRealm) SetRemoteAccessConfig(v RemoteAccessConfig)`

SetRemoteAccessConfig sets RemoteAccessConfig field to given value.

### HasRemoteAccessConfig

`func (o *VirtualizationRealm) HasRemoteAccessConfig() bool`

HasRemoteAccessConfig returns a boolean if a field has been set.

### GetRemoteAccessDeploymentId

`func (o *VirtualizationRealm) GetRemoteAccessDeploymentId() int32`

GetRemoteAccessDeploymentId returns the RemoteAccessDeploymentId field if non-nil, zero value otherwise.

### GetRemoteAccessDeploymentIdOk

`func (o *VirtualizationRealm) GetRemoteAccessDeploymentIdOk() (*int32, bool)`

GetRemoteAccessDeploymentIdOk returns a tuple with the RemoteAccessDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessDeploymentId

`func (o *VirtualizationRealm) SetRemoteAccessDeploymentId(v int32)`

SetRemoteAccessDeploymentId sets RemoteAccessDeploymentId field to given value.

### HasRemoteAccessDeploymentId

`func (o *VirtualizationRealm) HasRemoteAccessDeploymentId() bool`

HasRemoteAccessDeploymentId returns a boolean if a field has been set.

### GetRemoteAccessDeploymentRunStatus

`func (o *VirtualizationRealm) GetRemoteAccessDeploymentRunStatus() string`

GetRemoteAccessDeploymentRunStatus returns the RemoteAccessDeploymentRunStatus field if non-nil, zero value otherwise.

### GetRemoteAccessDeploymentRunStatusOk

`func (o *VirtualizationRealm) GetRemoteAccessDeploymentRunStatusOk() (*string, bool)`

GetRemoteAccessDeploymentRunStatusOk returns a tuple with the RemoteAccessDeploymentRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessDeploymentRunStatus

`func (o *VirtualizationRealm) SetRemoteAccessDeploymentRunStatus(v string)`

SetRemoteAccessDeploymentRunStatus sets RemoteAccessDeploymentRunStatus field to given value.

### HasRemoteAccessDeploymentRunStatus

`func (o *VirtualizationRealm) HasRemoteAccessDeploymentRunStatus() bool`

HasRemoteAccessDeploymentRunStatus returns a boolean if a field has been set.

### GetRemoteAccessStatus

`func (o *VirtualizationRealm) GetRemoteAccessStatus() string`

GetRemoteAccessStatus returns the RemoteAccessStatus field if non-nil, zero value otherwise.

### GetRemoteAccessStatusOk

`func (o *VirtualizationRealm) GetRemoteAccessStatusOk() (*string, bool)`

GetRemoteAccessStatusOk returns a tuple with the RemoteAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessStatus

`func (o *VirtualizationRealm) SetRemoteAccessStatus(v string)`

SetRemoteAccessStatus sets RemoteAccessStatus field to given value.

### HasRemoteAccessStatus

`func (o *VirtualizationRealm) HasRemoteAccessStatus() bool`

HasRemoteAccessStatus returns a boolean if a field has been set.

### GetState

`func (o *VirtualizationRealm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualizationRealm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualizationRealm) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualizationRealm) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *VirtualizationRealm) GetSupportedFeatures() []string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *VirtualizationRealm) GetSupportedFeaturesOk() (*[]string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *VirtualizationRealm) SetSupportedFeatures(v []string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *VirtualizationRealm) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetTemplateRegistrations

`func (o *VirtualizationRealm) GetTemplateRegistrations() []TemplateRegistration`

GetTemplateRegistrations returns the TemplateRegistrations field if non-nil, zero value otherwise.

### GetTemplateRegistrationsOk

`func (o *VirtualizationRealm) GetTemplateRegistrationsOk() (*[]TemplateRegistration, bool)`

GetTemplateRegistrationsOk returns a tuple with the TemplateRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegistrations

`func (o *VirtualizationRealm) SetTemplateRegistrations(v []TemplateRegistration)`

SetTemplateRegistrations sets TemplateRegistrations field to given value.

### HasTemplateRegistrations

`func (o *VirtualizationRealm) HasTemplateRegistrations() bool`

HasTemplateRegistrations returns a boolean if a field has been set.

### GetTemplates

`func (o *VirtualizationRealm) GetTemplates() []Cons3rtTemplateData`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *VirtualizationRealm) GetTemplatesOk() (*[]Cons3rtTemplateData, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *VirtualizationRealm) SetTemplates(v []Cons3rtTemplateData)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *VirtualizationRealm) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTemplateSubscriptions

`func (o *VirtualizationRealm) GetTemplateSubscriptions() []TemplateSubscription`

GetTemplateSubscriptions returns the TemplateSubscriptions field if non-nil, zero value otherwise.

### GetTemplateSubscriptionsOk

`func (o *VirtualizationRealm) GetTemplateSubscriptionsOk() (*[]TemplateSubscription, bool)`

GetTemplateSubscriptionsOk returns a tuple with the TemplateSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSubscriptions

`func (o *VirtualizationRealm) SetTemplateSubscriptions(v []TemplateSubscription)`

SetTemplateSubscriptions sets TemplateSubscriptions field to given value.

### HasTemplateSubscriptions

`func (o *VirtualizationRealm) HasTemplateSubscriptions() bool`

HasTemplateSubscriptions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VirtualizationRealm) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualizationRealm) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualizationRealm) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VirtualizationRealm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *VirtualizationRealm) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VirtualizationRealm) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VirtualizationRealm) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetZoneCount

`func (o *VirtualizationRealm) GetZoneCount() int32`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *VirtualizationRealm) GetZoneCountOk() (*int32, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *VirtualizationRealm) SetZoneCount(v int32)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *VirtualizationRealm) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


