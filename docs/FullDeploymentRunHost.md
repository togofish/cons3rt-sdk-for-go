# FullDeploymentRunHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildOrder** | Pointer to **int32** |  | [optional] 
**CreatedPassword** | Pointer to **string** |  | [optional] 
**CreatedUsername** | Pointer to **string** |  | [optional] 
**DefaultPassword** | Pointer to **string** |  | [optional] 
**DefaultUsername** | Pointer to **string** |  | [optional] 
**Disks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**FapStatus** | Pointer to **string** |  | [optional] 
**GpuProfile** | Pointer to **string** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**HasGpu** | Pointer to **bool** |  | [optional] 
**HostActionInProcess** | Pointer to **bool** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Installations** | Pointer to [**[]AbstractInstallation**](AbstractInstallation.md) |  | [optional] 
**InstanceTypeName** | Pointer to **string** |  | [optional] 
**NestedVirtualization** | Pointer to **bool** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) |  | [optional] 
**NumCpus** | Pointer to **int32** |  | [optional] 
**PhysicalMachineDataOrTemplateUuid** | Pointer to **string** |  | [optional] 
**PhysicalMachineOrTemplateName** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**SnapshotAvailable** | Pointer to **bool** |  | [optional] 
**SnapshotDate** | Pointer to **int32** |  | [optional] 
**SystemModuleId** | Pointer to **int32** |  | [optional] 
**SystemModuleType** | Pointer to **string** |  | [optional] 
**SystemRole** | Pointer to **string** |  | [optional] 
**VirtualizationRealmStatus** | Pointer to **string** |  | [optional] 
**Virtual** | Pointer to **bool** |  | [optional] 
**Provisionable** | Pointer to **bool** |  | [optional] 

## Methods

### NewFullDeploymentRunHost

`func NewFullDeploymentRunHost() *FullDeploymentRunHost`

NewFullDeploymentRunHost instantiates a new FullDeploymentRunHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDeploymentRunHostWithDefaults

`func NewFullDeploymentRunHostWithDefaults() *FullDeploymentRunHost`

NewFullDeploymentRunHostWithDefaults instantiates a new FullDeploymentRunHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildOrder

`func (o *FullDeploymentRunHost) GetBuildOrder() int32`

GetBuildOrder returns the BuildOrder field if non-nil, zero value otherwise.

### GetBuildOrderOk

`func (o *FullDeploymentRunHost) GetBuildOrderOk() (*int32, bool)`

GetBuildOrderOk returns a tuple with the BuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOrder

`func (o *FullDeploymentRunHost) SetBuildOrder(v int32)`

SetBuildOrder sets BuildOrder field to given value.

### HasBuildOrder

`func (o *FullDeploymentRunHost) HasBuildOrder() bool`

HasBuildOrder returns a boolean if a field has been set.

### GetCreatedPassword

`func (o *FullDeploymentRunHost) GetCreatedPassword() string`

GetCreatedPassword returns the CreatedPassword field if non-nil, zero value otherwise.

### GetCreatedPasswordOk

`func (o *FullDeploymentRunHost) GetCreatedPasswordOk() (*string, bool)`

GetCreatedPasswordOk returns a tuple with the CreatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedPassword

`func (o *FullDeploymentRunHost) SetCreatedPassword(v string)`

SetCreatedPassword sets CreatedPassword field to given value.

### HasCreatedPassword

`func (o *FullDeploymentRunHost) HasCreatedPassword() bool`

HasCreatedPassword returns a boolean if a field has been set.

### GetCreatedUsername

`func (o *FullDeploymentRunHost) GetCreatedUsername() string`

GetCreatedUsername returns the CreatedUsername field if non-nil, zero value otherwise.

### GetCreatedUsernameOk

`func (o *FullDeploymentRunHost) GetCreatedUsernameOk() (*string, bool)`

GetCreatedUsernameOk returns a tuple with the CreatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUsername

`func (o *FullDeploymentRunHost) SetCreatedUsername(v string)`

SetCreatedUsername sets CreatedUsername field to given value.

### HasCreatedUsername

`func (o *FullDeploymentRunHost) HasCreatedUsername() bool`

HasCreatedUsername returns a boolean if a field has been set.

### GetDefaultPassword

`func (o *FullDeploymentRunHost) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *FullDeploymentRunHost) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *FullDeploymentRunHost) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *FullDeploymentRunHost) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### GetDefaultUsername

`func (o *FullDeploymentRunHost) GetDefaultUsername() string`

GetDefaultUsername returns the DefaultUsername field if non-nil, zero value otherwise.

### GetDefaultUsernameOk

`func (o *FullDeploymentRunHost) GetDefaultUsernameOk() (*string, bool)`

GetDefaultUsernameOk returns a tuple with the DefaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsername

`func (o *FullDeploymentRunHost) SetDefaultUsername(v string)`

SetDefaultUsername sets DefaultUsername field to given value.

### HasDefaultUsername

`func (o *FullDeploymentRunHost) HasDefaultUsername() bool`

HasDefaultUsername returns a boolean if a field has been set.

### GetDisks

`func (o *FullDeploymentRunHost) GetDisks() []Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *FullDeploymentRunHost) GetDisksOk() (*[]Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *FullDeploymentRunHost) SetDisks(v []Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *FullDeploymentRunHost) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetFapStatus

`func (o *FullDeploymentRunHost) GetFapStatus() string`

GetFapStatus returns the FapStatus field if non-nil, zero value otherwise.

### GetFapStatusOk

`func (o *FullDeploymentRunHost) GetFapStatusOk() (*string, bool)`

GetFapStatusOk returns a tuple with the FapStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFapStatus

`func (o *FullDeploymentRunHost) SetFapStatus(v string)`

SetFapStatus sets FapStatus field to given value.

### HasFapStatus

`func (o *FullDeploymentRunHost) HasFapStatus() bool`

HasFapStatus returns a boolean if a field has been set.

### GetGpuProfile

`func (o *FullDeploymentRunHost) GetGpuProfile() string`

GetGpuProfile returns the GpuProfile field if non-nil, zero value otherwise.

### GetGpuProfileOk

`func (o *FullDeploymentRunHost) GetGpuProfileOk() (*string, bool)`

GetGpuProfileOk returns a tuple with the GpuProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuProfile

`func (o *FullDeploymentRunHost) SetGpuProfile(v string)`

SetGpuProfile sets GpuProfile field to given value.

### HasGpuProfile

`func (o *FullDeploymentRunHost) HasGpuProfile() bool`

HasGpuProfile returns a boolean if a field has been set.

### GetGpuType

`func (o *FullDeploymentRunHost) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *FullDeploymentRunHost) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *FullDeploymentRunHost) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *FullDeploymentRunHost) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetHasGpu

`func (o *FullDeploymentRunHost) GetHasGpu() bool`

GetHasGpu returns the HasGpu field if non-nil, zero value otherwise.

### GetHasGpuOk

`func (o *FullDeploymentRunHost) GetHasGpuOk() (*bool, bool)`

GetHasGpuOk returns a tuple with the HasGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpu

`func (o *FullDeploymentRunHost) SetHasGpu(v bool)`

SetHasGpu sets HasGpu field to given value.

### HasHasGpu

`func (o *FullDeploymentRunHost) HasHasGpu() bool`

HasHasGpu returns a boolean if a field has been set.

### GetHostActionInProcess

`func (o *FullDeploymentRunHost) GetHostActionInProcess() bool`

GetHostActionInProcess returns the HostActionInProcess field if non-nil, zero value otherwise.

### GetHostActionInProcessOk

`func (o *FullDeploymentRunHost) GetHostActionInProcessOk() (*bool, bool)`

GetHostActionInProcessOk returns a tuple with the HostActionInProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostActionInProcess

`func (o *FullDeploymentRunHost) SetHostActionInProcess(v bool)`

SetHostActionInProcess sets HostActionInProcess field to given value.

### HasHostActionInProcess

`func (o *FullDeploymentRunHost) HasHostActionInProcess() bool`

HasHostActionInProcess returns a boolean if a field has been set.

### GetHostname

`func (o *FullDeploymentRunHost) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FullDeploymentRunHost) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FullDeploymentRunHost) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FullDeploymentRunHost) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *FullDeploymentRunHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullDeploymentRunHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullDeploymentRunHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullDeploymentRunHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallations

`func (o *FullDeploymentRunHost) GetInstallations() []AbstractInstallation`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *FullDeploymentRunHost) GetInstallationsOk() (*[]AbstractInstallation, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *FullDeploymentRunHost) SetInstallations(v []AbstractInstallation)`

SetInstallations sets Installations field to given value.

### HasInstallations

`func (o *FullDeploymentRunHost) HasInstallations() bool`

HasInstallations returns a boolean if a field has been set.

### GetInstanceTypeName

`func (o *FullDeploymentRunHost) GetInstanceTypeName() string`

GetInstanceTypeName returns the InstanceTypeName field if non-nil, zero value otherwise.

### GetInstanceTypeNameOk

`func (o *FullDeploymentRunHost) GetInstanceTypeNameOk() (*string, bool)`

GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeName

`func (o *FullDeploymentRunHost) SetInstanceTypeName(v string)`

SetInstanceTypeName sets InstanceTypeName field to given value.

### HasInstanceTypeName

`func (o *FullDeploymentRunHost) HasInstanceTypeName() bool`

HasInstanceTypeName returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *FullDeploymentRunHost) GetNestedVirtualization() bool`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *FullDeploymentRunHost) GetNestedVirtualizationOk() (*bool, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *FullDeploymentRunHost) SetNestedVirtualization(v bool)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *FullDeploymentRunHost) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *FullDeploymentRunHost) GetNetworkInterfaces() []NetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *FullDeploymentRunHost) GetNetworkInterfacesOk() (*[]NetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *FullDeploymentRunHost) SetNetworkInterfaces(v []NetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *FullDeploymentRunHost) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetNumCpus

`func (o *FullDeploymentRunHost) GetNumCpus() int32`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *FullDeploymentRunHost) GetNumCpusOk() (*int32, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *FullDeploymentRunHost) SetNumCpus(v int32)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *FullDeploymentRunHost) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetPhysicalMachineDataOrTemplateUuid

`func (o *FullDeploymentRunHost) GetPhysicalMachineDataOrTemplateUuid() string`

GetPhysicalMachineDataOrTemplateUuid returns the PhysicalMachineDataOrTemplateUuid field if non-nil, zero value otherwise.

### GetPhysicalMachineDataOrTemplateUuidOk

`func (o *FullDeploymentRunHost) GetPhysicalMachineDataOrTemplateUuidOk() (*string, bool)`

GetPhysicalMachineDataOrTemplateUuidOk returns a tuple with the PhysicalMachineDataOrTemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachineDataOrTemplateUuid

`func (o *FullDeploymentRunHost) SetPhysicalMachineDataOrTemplateUuid(v string)`

SetPhysicalMachineDataOrTemplateUuid sets PhysicalMachineDataOrTemplateUuid field to given value.

### HasPhysicalMachineDataOrTemplateUuid

`func (o *FullDeploymentRunHost) HasPhysicalMachineDataOrTemplateUuid() bool`

HasPhysicalMachineDataOrTemplateUuid returns a boolean if a field has been set.

### GetPhysicalMachineOrTemplateName

`func (o *FullDeploymentRunHost) GetPhysicalMachineOrTemplateName() string`

GetPhysicalMachineOrTemplateName returns the PhysicalMachineOrTemplateName field if non-nil, zero value otherwise.

### GetPhysicalMachineOrTemplateNameOk

`func (o *FullDeploymentRunHost) GetPhysicalMachineOrTemplateNameOk() (*string, bool)`

GetPhysicalMachineOrTemplateNameOk returns a tuple with the PhysicalMachineOrTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachineOrTemplateName

`func (o *FullDeploymentRunHost) SetPhysicalMachineOrTemplateName(v string)`

SetPhysicalMachineOrTemplateName sets PhysicalMachineOrTemplateName field to given value.

### HasPhysicalMachineOrTemplateName

`func (o *FullDeploymentRunHost) HasPhysicalMachineOrTemplateName() bool`

HasPhysicalMachineOrTemplateName returns a boolean if a field has been set.

### GetPublished

`func (o *FullDeploymentRunHost) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *FullDeploymentRunHost) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *FullDeploymentRunHost) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *FullDeploymentRunHost) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRam

`func (o *FullDeploymentRunHost) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *FullDeploymentRunHost) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *FullDeploymentRunHost) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *FullDeploymentRunHost) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetSnapshotAvailable

`func (o *FullDeploymentRunHost) GetSnapshotAvailable() bool`

GetSnapshotAvailable returns the SnapshotAvailable field if non-nil, zero value otherwise.

### GetSnapshotAvailableOk

`func (o *FullDeploymentRunHost) GetSnapshotAvailableOk() (*bool, bool)`

GetSnapshotAvailableOk returns a tuple with the SnapshotAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotAvailable

`func (o *FullDeploymentRunHost) SetSnapshotAvailable(v bool)`

SetSnapshotAvailable sets SnapshotAvailable field to given value.

### HasSnapshotAvailable

`func (o *FullDeploymentRunHost) HasSnapshotAvailable() bool`

HasSnapshotAvailable returns a boolean if a field has been set.

### GetSnapshotDate

`func (o *FullDeploymentRunHost) GetSnapshotDate() int32`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *FullDeploymentRunHost) GetSnapshotDateOk() (*int32, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *FullDeploymentRunHost) SetSnapshotDate(v int32)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *FullDeploymentRunHost) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### GetSystemModuleId

`func (o *FullDeploymentRunHost) GetSystemModuleId() int32`

GetSystemModuleId returns the SystemModuleId field if non-nil, zero value otherwise.

### GetSystemModuleIdOk

`func (o *FullDeploymentRunHost) GetSystemModuleIdOk() (*int32, bool)`

GetSystemModuleIdOk returns a tuple with the SystemModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModuleId

`func (o *FullDeploymentRunHost) SetSystemModuleId(v int32)`

SetSystemModuleId sets SystemModuleId field to given value.

### HasSystemModuleId

`func (o *FullDeploymentRunHost) HasSystemModuleId() bool`

HasSystemModuleId returns a boolean if a field has been set.

### GetSystemModuleType

`func (o *FullDeploymentRunHost) GetSystemModuleType() string`

GetSystemModuleType returns the SystemModuleType field if non-nil, zero value otherwise.

### GetSystemModuleTypeOk

`func (o *FullDeploymentRunHost) GetSystemModuleTypeOk() (*string, bool)`

GetSystemModuleTypeOk returns a tuple with the SystemModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModuleType

`func (o *FullDeploymentRunHost) SetSystemModuleType(v string)`

SetSystemModuleType sets SystemModuleType field to given value.

### HasSystemModuleType

`func (o *FullDeploymentRunHost) HasSystemModuleType() bool`

HasSystemModuleType returns a boolean if a field has been set.

### GetSystemRole

`func (o *FullDeploymentRunHost) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *FullDeploymentRunHost) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *FullDeploymentRunHost) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.

### HasSystemRole

`func (o *FullDeploymentRunHost) HasSystemRole() bool`

HasSystemRole returns a boolean if a field has been set.

### GetVirtualizationRealmStatus

`func (o *FullDeploymentRunHost) GetVirtualizationRealmStatus() string`

GetVirtualizationRealmStatus returns the VirtualizationRealmStatus field if non-nil, zero value otherwise.

### GetVirtualizationRealmStatusOk

`func (o *FullDeploymentRunHost) GetVirtualizationRealmStatusOk() (*string, bool)`

GetVirtualizationRealmStatusOk returns a tuple with the VirtualizationRealmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmStatus

`func (o *FullDeploymentRunHost) SetVirtualizationRealmStatus(v string)`

SetVirtualizationRealmStatus sets VirtualizationRealmStatus field to given value.

### HasVirtualizationRealmStatus

`func (o *FullDeploymentRunHost) HasVirtualizationRealmStatus() bool`

HasVirtualizationRealmStatus returns a boolean if a field has been set.

### GetVirtual

`func (o *FullDeploymentRunHost) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *FullDeploymentRunHost) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *FullDeploymentRunHost) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *FullDeploymentRunHost) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### GetProvisionable

`func (o *FullDeploymentRunHost) GetProvisionable() bool`

GetProvisionable returns the Provisionable field if non-nil, zero value otherwise.

### GetProvisionableOk

`func (o *FullDeploymentRunHost) GetProvisionableOk() (*bool, bool)`

GetProvisionableOk returns a tuple with the Provisionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionable

`func (o *FullDeploymentRunHost) SetProvisionable(v bool)`

SetProvisionable sets Provisionable field to given value.

### HasProvisionable

`func (o *FullDeploymentRunHost) HasProvisionable() bool`

HasProvisionable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


