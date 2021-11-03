# HostOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRole** | **string** |  | 
**Cpus** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**AdditionalDisks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) |  | [optional] 
**BatchSoftwareInstall** | Pointer to **bool** |  | [optional] 
**GpuProfile** | Pointer to **string** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 

## Methods

### NewHostOption

`func NewHostOption(systemRole string, ) *HostOption`

NewHostOption instantiates a new HostOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOptionWithDefaults

`func NewHostOptionWithDefaults() *HostOption`

NewHostOptionWithDefaults instantiates a new HostOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRole

`func (o *HostOption) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *HostOption) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *HostOption) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.


### GetCpus

`func (o *HostOption) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *HostOption) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *HostOption) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *HostOption) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetRam

`func (o *HostOption) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *HostOption) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *HostOption) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *HostOption) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetAdditionalDisks

`func (o *HostOption) GetAdditionalDisks() []Disk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *HostOption) GetAdditionalDisksOk() (*[]Disk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *HostOption) SetAdditionalDisks(v []Disk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *HostOption) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *HostOption) GetNetworkInterfaces() []NetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *HostOption) GetNetworkInterfacesOk() (*[]NetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *HostOption) SetNetworkInterfaces(v []NetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *HostOption) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetBatchSoftwareInstall

`func (o *HostOption) GetBatchSoftwareInstall() bool`

GetBatchSoftwareInstall returns the BatchSoftwareInstall field if non-nil, zero value otherwise.

### GetBatchSoftwareInstallOk

`func (o *HostOption) GetBatchSoftwareInstallOk() (*bool, bool)`

GetBatchSoftwareInstallOk returns a tuple with the BatchSoftwareInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSoftwareInstall

`func (o *HostOption) SetBatchSoftwareInstall(v bool)`

SetBatchSoftwareInstall sets BatchSoftwareInstall field to given value.

### HasBatchSoftwareInstall

`func (o *HostOption) HasBatchSoftwareInstall() bool`

HasBatchSoftwareInstall returns a boolean if a field has been set.

### GetGpuProfile

`func (o *HostOption) GetGpuProfile() string`

GetGpuProfile returns the GpuProfile field if non-nil, zero value otherwise.

### GetGpuProfileOk

`func (o *HostOption) GetGpuProfileOk() (*string, bool)`

GetGpuProfileOk returns a tuple with the GpuProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuProfile

`func (o *HostOption) SetGpuProfile(v string)`

SetGpuProfile sets GpuProfile field to given value.

### HasGpuProfile

`func (o *HostOption) HasGpuProfile() bool`

HasGpuProfile returns a boolean if a field has been set.

### GetGpuType

`func (o *HostOption) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *HostOption) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *HostOption) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *HostOption) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


