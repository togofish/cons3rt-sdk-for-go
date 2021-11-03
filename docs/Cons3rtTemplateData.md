# Cons3rtTemplateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**OperatingSystem** | **string** |  | 
**VirtRealmTemplateName** | **string** |  | 
**Cons3rtAgentInstalled** | Pointer to **bool** |  | [optional] 
**ContainerCapable** | Pointer to **bool** |  | [optional] 
**Disks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**FailCount** | Pointer to **int32** |  | [optional] 
**HasGpu** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**MaxNumCpus** | Pointer to **int32** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int64** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**PackageManagementType** | Pointer to **string** |  | [optional] 
**PowerOnDelayOverride** | Pointer to **int32** |  | [optional] 
**PowerShellVersion** | Pointer to **string** |  | [optional] 
**TemplateRegistration** | Pointer to [**TemplateRegistration**](TemplateRegistration.md) |  | [optional] 
**RemoteAccessTemplates** | Pointer to [**[]RemoteAccessTemplate**](RemoteAccessTemplate.md) |  | [optional] 
**ServiceManagementType** | Pointer to **string** |  | [optional] 
**UserCount** | Pointer to **int32** |  | [optional] 
**VirtRealmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCons3rtTemplateData

`func NewCons3rtTemplateData(operatingSystem string, virtRealmTemplateName string, ) *Cons3rtTemplateData`

NewCons3rtTemplateData instantiates a new Cons3rtTemplateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCons3rtTemplateDataWithDefaults

`func NewCons3rtTemplateDataWithDefaults() *Cons3rtTemplateData`

NewCons3rtTemplateDataWithDefaults instantiates a new Cons3rtTemplateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Cons3rtTemplateData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Cons3rtTemplateData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Cons3rtTemplateData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Cons3rtTemplateData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Cons3rtTemplateData) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Cons3rtTemplateData) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Cons3rtTemplateData) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetVirtRealmTemplateName

`func (o *Cons3rtTemplateData) GetVirtRealmTemplateName() string`

GetVirtRealmTemplateName returns the VirtRealmTemplateName field if non-nil, zero value otherwise.

### GetVirtRealmTemplateNameOk

`func (o *Cons3rtTemplateData) GetVirtRealmTemplateNameOk() (*string, bool)`

GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmTemplateName

`func (o *Cons3rtTemplateData) SetVirtRealmTemplateName(v string)`

SetVirtRealmTemplateName sets VirtRealmTemplateName field to given value.


### GetCons3rtAgentInstalled

`func (o *Cons3rtTemplateData) GetCons3rtAgentInstalled() bool`

GetCons3rtAgentInstalled returns the Cons3rtAgentInstalled field if non-nil, zero value otherwise.

### GetCons3rtAgentInstalledOk

`func (o *Cons3rtTemplateData) GetCons3rtAgentInstalledOk() (*bool, bool)`

GetCons3rtAgentInstalledOk returns a tuple with the Cons3rtAgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtAgentInstalled

`func (o *Cons3rtTemplateData) SetCons3rtAgentInstalled(v bool)`

SetCons3rtAgentInstalled sets Cons3rtAgentInstalled field to given value.

### HasCons3rtAgentInstalled

`func (o *Cons3rtTemplateData) HasCons3rtAgentInstalled() bool`

HasCons3rtAgentInstalled returns a boolean if a field has been set.

### GetContainerCapable

`func (o *Cons3rtTemplateData) GetContainerCapable() bool`

GetContainerCapable returns the ContainerCapable field if non-nil, zero value otherwise.

### GetContainerCapableOk

`func (o *Cons3rtTemplateData) GetContainerCapableOk() (*bool, bool)`

GetContainerCapableOk returns a tuple with the ContainerCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCapable

`func (o *Cons3rtTemplateData) SetContainerCapable(v bool)`

SetContainerCapable sets ContainerCapable field to given value.

### HasContainerCapable

`func (o *Cons3rtTemplateData) HasContainerCapable() bool`

HasContainerCapable returns a boolean if a field has been set.

### GetDisks

`func (o *Cons3rtTemplateData) GetDisks() []Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Cons3rtTemplateData) GetDisksOk() (*[]Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Cons3rtTemplateData) SetDisks(v []Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Cons3rtTemplateData) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetFailCount

`func (o *Cons3rtTemplateData) GetFailCount() int32`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *Cons3rtTemplateData) GetFailCountOk() (*int32, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *Cons3rtTemplateData) SetFailCount(v int32)`

SetFailCount sets FailCount field to given value.

### HasFailCount

`func (o *Cons3rtTemplateData) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### GetHasGpu

`func (o *Cons3rtTemplateData) GetHasGpu() bool`

GetHasGpu returns the HasGpu field if non-nil, zero value otherwise.

### GetHasGpuOk

`func (o *Cons3rtTemplateData) GetHasGpuOk() (*bool, bool)`

GetHasGpuOk returns a tuple with the HasGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpu

`func (o *Cons3rtTemplateData) SetHasGpu(v bool)`

SetHasGpu sets HasGpu field to given value.

### HasHasGpu

`func (o *Cons3rtTemplateData) HasHasGpu() bool`

HasHasGpu returns a boolean if a field has been set.

### GetId

`func (o *Cons3rtTemplateData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cons3rtTemplateData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cons3rtTemplateData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Cons3rtTemplateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicense

`func (o *Cons3rtTemplateData) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Cons3rtTemplateData) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Cons3rtTemplateData) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Cons3rtTemplateData) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *Cons3rtTemplateData) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *Cons3rtTemplateData) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *Cons3rtTemplateData) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *Cons3rtTemplateData) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *Cons3rtTemplateData) GetMaxRamInMegabytes() int64`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *Cons3rtTemplateData) GetMaxRamInMegabytesOk() (*int64, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *Cons3rtTemplateData) SetMaxRamInMegabytes(v int64)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *Cons3rtTemplateData) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetNote

`func (o *Cons3rtTemplateData) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Cons3rtTemplateData) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Cons3rtTemplateData) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Cons3rtTemplateData) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPackageManagementType

`func (o *Cons3rtTemplateData) GetPackageManagementType() string`

GetPackageManagementType returns the PackageManagementType field if non-nil, zero value otherwise.

### GetPackageManagementTypeOk

`func (o *Cons3rtTemplateData) GetPackageManagementTypeOk() (*string, bool)`

GetPackageManagementTypeOk returns a tuple with the PackageManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagementType

`func (o *Cons3rtTemplateData) SetPackageManagementType(v string)`

SetPackageManagementType sets PackageManagementType field to given value.

### HasPackageManagementType

`func (o *Cons3rtTemplateData) HasPackageManagementType() bool`

HasPackageManagementType returns a boolean if a field has been set.

### GetPowerOnDelayOverride

`func (o *Cons3rtTemplateData) GetPowerOnDelayOverride() int32`

GetPowerOnDelayOverride returns the PowerOnDelayOverride field if non-nil, zero value otherwise.

### GetPowerOnDelayOverrideOk

`func (o *Cons3rtTemplateData) GetPowerOnDelayOverrideOk() (*int32, bool)`

GetPowerOnDelayOverrideOk returns a tuple with the PowerOnDelayOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayOverride

`func (o *Cons3rtTemplateData) SetPowerOnDelayOverride(v int32)`

SetPowerOnDelayOverride sets PowerOnDelayOverride field to given value.

### HasPowerOnDelayOverride

`func (o *Cons3rtTemplateData) HasPowerOnDelayOverride() bool`

HasPowerOnDelayOverride returns a boolean if a field has been set.

### GetPowerShellVersion

`func (o *Cons3rtTemplateData) GetPowerShellVersion() string`

GetPowerShellVersion returns the PowerShellVersion field if non-nil, zero value otherwise.

### GetPowerShellVersionOk

`func (o *Cons3rtTemplateData) GetPowerShellVersionOk() (*string, bool)`

GetPowerShellVersionOk returns a tuple with the PowerShellVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShellVersion

`func (o *Cons3rtTemplateData) SetPowerShellVersion(v string)`

SetPowerShellVersion sets PowerShellVersion field to given value.

### HasPowerShellVersion

`func (o *Cons3rtTemplateData) HasPowerShellVersion() bool`

HasPowerShellVersion returns a boolean if a field has been set.

### GetTemplateRegistration

`func (o *Cons3rtTemplateData) GetTemplateRegistration() TemplateRegistration`

GetTemplateRegistration returns the TemplateRegistration field if non-nil, zero value otherwise.

### GetTemplateRegistrationOk

`func (o *Cons3rtTemplateData) GetTemplateRegistrationOk() (*TemplateRegistration, bool)`

GetTemplateRegistrationOk returns a tuple with the TemplateRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegistration

`func (o *Cons3rtTemplateData) SetTemplateRegistration(v TemplateRegistration)`

SetTemplateRegistration sets TemplateRegistration field to given value.

### HasTemplateRegistration

`func (o *Cons3rtTemplateData) HasTemplateRegistration() bool`

HasTemplateRegistration returns a boolean if a field has been set.

### GetRemoteAccessTemplates

`func (o *Cons3rtTemplateData) GetRemoteAccessTemplates() []RemoteAccessTemplate`

GetRemoteAccessTemplates returns the RemoteAccessTemplates field if non-nil, zero value otherwise.

### GetRemoteAccessTemplatesOk

`func (o *Cons3rtTemplateData) GetRemoteAccessTemplatesOk() (*[]RemoteAccessTemplate, bool)`

GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessTemplates

`func (o *Cons3rtTemplateData) SetRemoteAccessTemplates(v []RemoteAccessTemplate)`

SetRemoteAccessTemplates sets RemoteAccessTemplates field to given value.

### HasRemoteAccessTemplates

`func (o *Cons3rtTemplateData) HasRemoteAccessTemplates() bool`

HasRemoteAccessTemplates returns a boolean if a field has been set.

### GetServiceManagementType

`func (o *Cons3rtTemplateData) GetServiceManagementType() string`

GetServiceManagementType returns the ServiceManagementType field if non-nil, zero value otherwise.

### GetServiceManagementTypeOk

`func (o *Cons3rtTemplateData) GetServiceManagementTypeOk() (*string, bool)`

GetServiceManagementTypeOk returns a tuple with the ServiceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManagementType

`func (o *Cons3rtTemplateData) SetServiceManagementType(v string)`

SetServiceManagementType sets ServiceManagementType field to given value.

### HasServiceManagementType

`func (o *Cons3rtTemplateData) HasServiceManagementType() bool`

HasServiceManagementType returns a boolean if a field has been set.

### GetUserCount

`func (o *Cons3rtTemplateData) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *Cons3rtTemplateData) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *Cons3rtTemplateData) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *Cons3rtTemplateData) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetVirtRealmId

`func (o *Cons3rtTemplateData) GetVirtRealmId() int32`

GetVirtRealmId returns the VirtRealmId field if non-nil, zero value otherwise.

### GetVirtRealmIdOk

`func (o *Cons3rtTemplateData) GetVirtRealmIdOk() (*int32, bool)`

GetVirtRealmIdOk returns a tuple with the VirtRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmId

`func (o *Cons3rtTemplateData) SetVirtRealmId(v int32)`

SetVirtRealmId sets VirtRealmId field to given value.

### HasVirtRealmId

`func (o *Cons3rtTemplateData) HasVirtRealmId() bool`

HasVirtRealmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


