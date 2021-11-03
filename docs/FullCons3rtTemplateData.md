# FullCons3rtTemplateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Cons3rtAgentInstalled** | Pointer to **bool** |  | [optional] 
**ContainerCapable** | Pointer to **bool** |  | [optional] 
**Disks** | Pointer to [**[]MinimalDisk**](MinimalDisk.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FailCount** | Pointer to **int32** |  | [optional] 
**HasGpu** | Pointer to **bool** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int64** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**OperatingSystem** | **string** |  | 
**PackageManagementType** | Pointer to **string** |  | [optional] 
**PowerOnDelayOverride** | Pointer to **int32** |  | [optional] 
**PowerShellVersion** | Pointer to **string** |  | [optional] 
**ServiceManagementType** | Pointer to **string** |  | [optional] 
**RemoteAccessTemplates** | Pointer to [**[]MinimalRemoteAccessTemplate**](MinimalRemoteAccessTemplate.md) |  | [optional] 
**UserCount** | Pointer to **int32** |  | [optional] 
**VirtRealmId** | Pointer to **int32** |  | [optional] 
**VirtRealmTemplateName** | **string** |  | 

## Methods

### NewFullCons3rtTemplateData

`func NewFullCons3rtTemplateData(operatingSystem string, virtRealmTemplateName string, ) *FullCons3rtTemplateData`

NewFullCons3rtTemplateData instantiates a new FullCons3rtTemplateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullCons3rtTemplateDataWithDefaults

`func NewFullCons3rtTemplateDataWithDefaults() *FullCons3rtTemplateData`

NewFullCons3rtTemplateDataWithDefaults instantiates a new FullCons3rtTemplateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullCons3rtTemplateData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullCons3rtTemplateData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullCons3rtTemplateData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullCons3rtTemplateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCons3rtAgentInstalled

`func (o *FullCons3rtTemplateData) GetCons3rtAgentInstalled() bool`

GetCons3rtAgentInstalled returns the Cons3rtAgentInstalled field if non-nil, zero value otherwise.

### GetCons3rtAgentInstalledOk

`func (o *FullCons3rtTemplateData) GetCons3rtAgentInstalledOk() (*bool, bool)`

GetCons3rtAgentInstalledOk returns a tuple with the Cons3rtAgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtAgentInstalled

`func (o *FullCons3rtTemplateData) SetCons3rtAgentInstalled(v bool)`

SetCons3rtAgentInstalled sets Cons3rtAgentInstalled field to given value.

### HasCons3rtAgentInstalled

`func (o *FullCons3rtTemplateData) HasCons3rtAgentInstalled() bool`

HasCons3rtAgentInstalled returns a boolean if a field has been set.

### GetContainerCapable

`func (o *FullCons3rtTemplateData) GetContainerCapable() bool`

GetContainerCapable returns the ContainerCapable field if non-nil, zero value otherwise.

### GetContainerCapableOk

`func (o *FullCons3rtTemplateData) GetContainerCapableOk() (*bool, bool)`

GetContainerCapableOk returns a tuple with the ContainerCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCapable

`func (o *FullCons3rtTemplateData) SetContainerCapable(v bool)`

SetContainerCapable sets ContainerCapable field to given value.

### HasContainerCapable

`func (o *FullCons3rtTemplateData) HasContainerCapable() bool`

HasContainerCapable returns a boolean if a field has been set.

### GetDisks

`func (o *FullCons3rtTemplateData) GetDisks() []MinimalDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *FullCons3rtTemplateData) GetDisksOk() (*[]MinimalDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *FullCons3rtTemplateData) SetDisks(v []MinimalDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *FullCons3rtTemplateData) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetDisplayName

`func (o *FullCons3rtTemplateData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FullCons3rtTemplateData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FullCons3rtTemplateData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FullCons3rtTemplateData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFailCount

`func (o *FullCons3rtTemplateData) GetFailCount() int32`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *FullCons3rtTemplateData) GetFailCountOk() (*int32, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *FullCons3rtTemplateData) SetFailCount(v int32)`

SetFailCount sets FailCount field to given value.

### HasFailCount

`func (o *FullCons3rtTemplateData) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### GetHasGpu

`func (o *FullCons3rtTemplateData) GetHasGpu() bool`

GetHasGpu returns the HasGpu field if non-nil, zero value otherwise.

### GetHasGpuOk

`func (o *FullCons3rtTemplateData) GetHasGpuOk() (*bool, bool)`

GetHasGpuOk returns a tuple with the HasGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpu

`func (o *FullCons3rtTemplateData) SetHasGpu(v bool)`

SetHasGpu sets HasGpu field to given value.

### HasHasGpu

`func (o *FullCons3rtTemplateData) HasHasGpu() bool`

HasHasGpu returns a boolean if a field has been set.

### GetLicense

`func (o *FullCons3rtTemplateData) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *FullCons3rtTemplateData) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *FullCons3rtTemplateData) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *FullCons3rtTemplateData) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *FullCons3rtTemplateData) GetMaxRamInMegabytes() int64`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *FullCons3rtTemplateData) GetMaxRamInMegabytesOk() (*int64, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *FullCons3rtTemplateData) SetMaxRamInMegabytes(v int64)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *FullCons3rtTemplateData) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetNote

`func (o *FullCons3rtTemplateData) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FullCons3rtTemplateData) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FullCons3rtTemplateData) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FullCons3rtTemplateData) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *FullCons3rtTemplateData) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *FullCons3rtTemplateData) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *FullCons3rtTemplateData) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetPackageManagementType

`func (o *FullCons3rtTemplateData) GetPackageManagementType() string`

GetPackageManagementType returns the PackageManagementType field if non-nil, zero value otherwise.

### GetPackageManagementTypeOk

`func (o *FullCons3rtTemplateData) GetPackageManagementTypeOk() (*string, bool)`

GetPackageManagementTypeOk returns a tuple with the PackageManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagementType

`func (o *FullCons3rtTemplateData) SetPackageManagementType(v string)`

SetPackageManagementType sets PackageManagementType field to given value.

### HasPackageManagementType

`func (o *FullCons3rtTemplateData) HasPackageManagementType() bool`

HasPackageManagementType returns a boolean if a field has been set.

### GetPowerOnDelayOverride

`func (o *FullCons3rtTemplateData) GetPowerOnDelayOverride() int32`

GetPowerOnDelayOverride returns the PowerOnDelayOverride field if non-nil, zero value otherwise.

### GetPowerOnDelayOverrideOk

`func (o *FullCons3rtTemplateData) GetPowerOnDelayOverrideOk() (*int32, bool)`

GetPowerOnDelayOverrideOk returns a tuple with the PowerOnDelayOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayOverride

`func (o *FullCons3rtTemplateData) SetPowerOnDelayOverride(v int32)`

SetPowerOnDelayOverride sets PowerOnDelayOverride field to given value.

### HasPowerOnDelayOverride

`func (o *FullCons3rtTemplateData) HasPowerOnDelayOverride() bool`

HasPowerOnDelayOverride returns a boolean if a field has been set.

### GetPowerShellVersion

`func (o *FullCons3rtTemplateData) GetPowerShellVersion() string`

GetPowerShellVersion returns the PowerShellVersion field if non-nil, zero value otherwise.

### GetPowerShellVersionOk

`func (o *FullCons3rtTemplateData) GetPowerShellVersionOk() (*string, bool)`

GetPowerShellVersionOk returns a tuple with the PowerShellVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShellVersion

`func (o *FullCons3rtTemplateData) SetPowerShellVersion(v string)`

SetPowerShellVersion sets PowerShellVersion field to given value.

### HasPowerShellVersion

`func (o *FullCons3rtTemplateData) HasPowerShellVersion() bool`

HasPowerShellVersion returns a boolean if a field has been set.

### GetServiceManagementType

`func (o *FullCons3rtTemplateData) GetServiceManagementType() string`

GetServiceManagementType returns the ServiceManagementType field if non-nil, zero value otherwise.

### GetServiceManagementTypeOk

`func (o *FullCons3rtTemplateData) GetServiceManagementTypeOk() (*string, bool)`

GetServiceManagementTypeOk returns a tuple with the ServiceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManagementType

`func (o *FullCons3rtTemplateData) SetServiceManagementType(v string)`

SetServiceManagementType sets ServiceManagementType field to given value.

### HasServiceManagementType

`func (o *FullCons3rtTemplateData) HasServiceManagementType() bool`

HasServiceManagementType returns a boolean if a field has been set.

### GetRemoteAccessTemplates

`func (o *FullCons3rtTemplateData) GetRemoteAccessTemplates() []MinimalRemoteAccessTemplate`

GetRemoteAccessTemplates returns the RemoteAccessTemplates field if non-nil, zero value otherwise.

### GetRemoteAccessTemplatesOk

`func (o *FullCons3rtTemplateData) GetRemoteAccessTemplatesOk() (*[]MinimalRemoteAccessTemplate, bool)`

GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessTemplates

`func (o *FullCons3rtTemplateData) SetRemoteAccessTemplates(v []MinimalRemoteAccessTemplate)`

SetRemoteAccessTemplates sets RemoteAccessTemplates field to given value.

### HasRemoteAccessTemplates

`func (o *FullCons3rtTemplateData) HasRemoteAccessTemplates() bool`

HasRemoteAccessTemplates returns a boolean if a field has been set.

### GetUserCount

`func (o *FullCons3rtTemplateData) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *FullCons3rtTemplateData) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *FullCons3rtTemplateData) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *FullCons3rtTemplateData) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetVirtRealmId

`func (o *FullCons3rtTemplateData) GetVirtRealmId() int32`

GetVirtRealmId returns the VirtRealmId field if non-nil, zero value otherwise.

### GetVirtRealmIdOk

`func (o *FullCons3rtTemplateData) GetVirtRealmIdOk() (*int32, bool)`

GetVirtRealmIdOk returns a tuple with the VirtRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmId

`func (o *FullCons3rtTemplateData) SetVirtRealmId(v int32)`

SetVirtRealmId sets VirtRealmId field to given value.

### HasVirtRealmId

`func (o *FullCons3rtTemplateData) HasVirtRealmId() bool`

HasVirtRealmId returns a boolean if a field has been set.

### GetVirtRealmTemplateName

`func (o *FullCons3rtTemplateData) GetVirtRealmTemplateName() string`

GetVirtRealmTemplateName returns the VirtRealmTemplateName field if non-nil, zero value otherwise.

### GetVirtRealmTemplateNameOk

`func (o *FullCons3rtTemplateData) GetVirtRealmTemplateNameOk() (*string, bool)`

GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmTemplateName

`func (o *FullCons3rtTemplateData) SetVirtRealmTemplateName(v string)`

SetVirtRealmTemplateName sets VirtRealmTemplateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


