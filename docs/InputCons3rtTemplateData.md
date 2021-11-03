# InputCons3rtTemplateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**VirtRealmTemplateName** | **string** |  | 
**OperatingSystem** | **string** |  | 
**Cons3rtAgentInstalled** | Pointer to **bool** |  | [optional] 
**ContainerCapable** | Pointer to **bool** |  | [optional] 
**Disks** | Pointer to [**[]InputDiskForTemplate**](InputDiskForTemplate.md) |  | [optional] 
**HasGpu** | Pointer to **bool** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**MaxNumCpus** | Pointer to **int32** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int64** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**PackageManagementType** | Pointer to **string** |  | [optional] 
**PowerOnDelayOverride** | Pointer to **int32** |  | [optional] 
**PowerShellVersion** | Pointer to **string** |  | [optional] 
**ServiceManagementType** | Pointer to **string** |  | [optional] 
**RemoteAccessTemplates** | Pointer to [**[]InputRemoteAccessTemplate**](InputRemoteAccessTemplate.md) |  | [optional] 

## Methods

### NewInputCons3rtTemplateData

`func NewInputCons3rtTemplateData(virtRealmTemplateName string, operatingSystem string, ) *InputCons3rtTemplateData`

NewInputCons3rtTemplateData instantiates a new InputCons3rtTemplateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCons3rtTemplateDataWithDefaults

`func NewInputCons3rtTemplateDataWithDefaults() *InputCons3rtTemplateData`

NewInputCons3rtTemplateDataWithDefaults instantiates a new InputCons3rtTemplateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *InputCons3rtTemplateData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InputCons3rtTemplateData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InputCons3rtTemplateData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InputCons3rtTemplateData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVirtRealmTemplateName

`func (o *InputCons3rtTemplateData) GetVirtRealmTemplateName() string`

GetVirtRealmTemplateName returns the VirtRealmTemplateName field if non-nil, zero value otherwise.

### GetVirtRealmTemplateNameOk

`func (o *InputCons3rtTemplateData) GetVirtRealmTemplateNameOk() (*string, bool)`

GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmTemplateName

`func (o *InputCons3rtTemplateData) SetVirtRealmTemplateName(v string)`

SetVirtRealmTemplateName sets VirtRealmTemplateName field to given value.


### GetOperatingSystem

`func (o *InputCons3rtTemplateData) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *InputCons3rtTemplateData) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *InputCons3rtTemplateData) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetCons3rtAgentInstalled

`func (o *InputCons3rtTemplateData) GetCons3rtAgentInstalled() bool`

GetCons3rtAgentInstalled returns the Cons3rtAgentInstalled field if non-nil, zero value otherwise.

### GetCons3rtAgentInstalledOk

`func (o *InputCons3rtTemplateData) GetCons3rtAgentInstalledOk() (*bool, bool)`

GetCons3rtAgentInstalledOk returns a tuple with the Cons3rtAgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtAgentInstalled

`func (o *InputCons3rtTemplateData) SetCons3rtAgentInstalled(v bool)`

SetCons3rtAgentInstalled sets Cons3rtAgentInstalled field to given value.

### HasCons3rtAgentInstalled

`func (o *InputCons3rtTemplateData) HasCons3rtAgentInstalled() bool`

HasCons3rtAgentInstalled returns a boolean if a field has been set.

### GetContainerCapable

`func (o *InputCons3rtTemplateData) GetContainerCapable() bool`

GetContainerCapable returns the ContainerCapable field if non-nil, zero value otherwise.

### GetContainerCapableOk

`func (o *InputCons3rtTemplateData) GetContainerCapableOk() (*bool, bool)`

GetContainerCapableOk returns a tuple with the ContainerCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCapable

`func (o *InputCons3rtTemplateData) SetContainerCapable(v bool)`

SetContainerCapable sets ContainerCapable field to given value.

### HasContainerCapable

`func (o *InputCons3rtTemplateData) HasContainerCapable() bool`

HasContainerCapable returns a boolean if a field has been set.

### GetDisks

`func (o *InputCons3rtTemplateData) GetDisks() []InputDiskForTemplate`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *InputCons3rtTemplateData) GetDisksOk() (*[]InputDiskForTemplate, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *InputCons3rtTemplateData) SetDisks(v []InputDiskForTemplate)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *InputCons3rtTemplateData) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetHasGpu

`func (o *InputCons3rtTemplateData) GetHasGpu() bool`

GetHasGpu returns the HasGpu field if non-nil, zero value otherwise.

### GetHasGpuOk

`func (o *InputCons3rtTemplateData) GetHasGpuOk() (*bool, bool)`

GetHasGpuOk returns a tuple with the HasGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpu

`func (o *InputCons3rtTemplateData) SetHasGpu(v bool)`

SetHasGpu sets HasGpu field to given value.

### HasHasGpu

`func (o *InputCons3rtTemplateData) HasHasGpu() bool`

HasHasGpu returns a boolean if a field has been set.

### GetLicense

`func (o *InputCons3rtTemplateData) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *InputCons3rtTemplateData) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *InputCons3rtTemplateData) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *InputCons3rtTemplateData) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *InputCons3rtTemplateData) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *InputCons3rtTemplateData) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *InputCons3rtTemplateData) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *InputCons3rtTemplateData) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *InputCons3rtTemplateData) GetMaxRamInMegabytes() int64`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *InputCons3rtTemplateData) GetMaxRamInMegabytesOk() (*int64, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *InputCons3rtTemplateData) SetMaxRamInMegabytes(v int64)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *InputCons3rtTemplateData) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetNote

`func (o *InputCons3rtTemplateData) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *InputCons3rtTemplateData) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *InputCons3rtTemplateData) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *InputCons3rtTemplateData) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPackageManagementType

`func (o *InputCons3rtTemplateData) GetPackageManagementType() string`

GetPackageManagementType returns the PackageManagementType field if non-nil, zero value otherwise.

### GetPackageManagementTypeOk

`func (o *InputCons3rtTemplateData) GetPackageManagementTypeOk() (*string, bool)`

GetPackageManagementTypeOk returns a tuple with the PackageManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagementType

`func (o *InputCons3rtTemplateData) SetPackageManagementType(v string)`

SetPackageManagementType sets PackageManagementType field to given value.

### HasPackageManagementType

`func (o *InputCons3rtTemplateData) HasPackageManagementType() bool`

HasPackageManagementType returns a boolean if a field has been set.

### GetPowerOnDelayOverride

`func (o *InputCons3rtTemplateData) GetPowerOnDelayOverride() int32`

GetPowerOnDelayOverride returns the PowerOnDelayOverride field if non-nil, zero value otherwise.

### GetPowerOnDelayOverrideOk

`func (o *InputCons3rtTemplateData) GetPowerOnDelayOverrideOk() (*int32, bool)`

GetPowerOnDelayOverrideOk returns a tuple with the PowerOnDelayOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayOverride

`func (o *InputCons3rtTemplateData) SetPowerOnDelayOverride(v int32)`

SetPowerOnDelayOverride sets PowerOnDelayOverride field to given value.

### HasPowerOnDelayOverride

`func (o *InputCons3rtTemplateData) HasPowerOnDelayOverride() bool`

HasPowerOnDelayOverride returns a boolean if a field has been set.

### GetPowerShellVersion

`func (o *InputCons3rtTemplateData) GetPowerShellVersion() string`

GetPowerShellVersion returns the PowerShellVersion field if non-nil, zero value otherwise.

### GetPowerShellVersionOk

`func (o *InputCons3rtTemplateData) GetPowerShellVersionOk() (*string, bool)`

GetPowerShellVersionOk returns a tuple with the PowerShellVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShellVersion

`func (o *InputCons3rtTemplateData) SetPowerShellVersion(v string)`

SetPowerShellVersion sets PowerShellVersion field to given value.

### HasPowerShellVersion

`func (o *InputCons3rtTemplateData) HasPowerShellVersion() bool`

HasPowerShellVersion returns a boolean if a field has been set.

### GetServiceManagementType

`func (o *InputCons3rtTemplateData) GetServiceManagementType() string`

GetServiceManagementType returns the ServiceManagementType field if non-nil, zero value otherwise.

### GetServiceManagementTypeOk

`func (o *InputCons3rtTemplateData) GetServiceManagementTypeOk() (*string, bool)`

GetServiceManagementTypeOk returns a tuple with the ServiceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManagementType

`func (o *InputCons3rtTemplateData) SetServiceManagementType(v string)`

SetServiceManagementType sets ServiceManagementType field to given value.

### HasServiceManagementType

`func (o *InputCons3rtTemplateData) HasServiceManagementType() bool`

HasServiceManagementType returns a boolean if a field has been set.

### GetRemoteAccessTemplates

`func (o *InputCons3rtTemplateData) GetRemoteAccessTemplates() []InputRemoteAccessTemplate`

GetRemoteAccessTemplates returns the RemoteAccessTemplates field if non-nil, zero value otherwise.

### GetRemoteAccessTemplatesOk

`func (o *InputCons3rtTemplateData) GetRemoteAccessTemplatesOk() (*[]InputRemoteAccessTemplate, bool)`

GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessTemplates

`func (o *InputCons3rtTemplateData) SetRemoteAccessTemplates(v []InputRemoteAccessTemplate)`

SetRemoteAccessTemplates sets RemoteAccessTemplates field to given value.

### HasRemoteAccessTemplates

`func (o *InputCons3rtTemplateData) HasRemoteAccessTemplates() bool`

HasRemoteAccessTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


