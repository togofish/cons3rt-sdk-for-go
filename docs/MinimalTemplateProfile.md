# MinimalTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtRealmTemplateName** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**MinBootDiskCapacity** | Pointer to **int32** |  | [optional] 
**MinNumCpus** | Pointer to **int32** |  | [optional] 
**MinRam** | Pointer to **int32** |  | [optional] 
**AdditionalDisks** | Pointer to [**[]MinimalDisk**](MinimalDisk.md) |  | [optional] 
**Cons3rtAgentRequired** | Pointer to **bool** |  | [optional] 
**VgpuRequired** | Pointer to **bool** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**RequiresNestedVirtualization** | Pointer to **bool** |  | [optional] 
**RemoteAccessRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewMinimalTemplateProfile

`func NewMinimalTemplateProfile() *MinimalTemplateProfile`

NewMinimalTemplateProfile instantiates a new MinimalTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalTemplateProfileWithDefaults

`func NewMinimalTemplateProfileWithDefaults() *MinimalTemplateProfile`

NewMinimalTemplateProfileWithDefaults instantiates a new MinimalTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtRealmTemplateName

`func (o *MinimalTemplateProfile) GetVirtRealmTemplateName() string`

GetVirtRealmTemplateName returns the VirtRealmTemplateName field if non-nil, zero value otherwise.

### GetVirtRealmTemplateNameOk

`func (o *MinimalTemplateProfile) GetVirtRealmTemplateNameOk() (*string, bool)`

GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmTemplateName

`func (o *MinimalTemplateProfile) SetVirtRealmTemplateName(v string)`

SetVirtRealmTemplateName sets VirtRealmTemplateName field to given value.

### HasVirtRealmTemplateName

`func (o *MinimalTemplateProfile) HasVirtRealmTemplateName() bool`

HasVirtRealmTemplateName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *MinimalTemplateProfile) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MinimalTemplateProfile) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *MinimalTemplateProfile) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *MinimalTemplateProfile) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetMinBootDiskCapacity

`func (o *MinimalTemplateProfile) GetMinBootDiskCapacity() int32`

GetMinBootDiskCapacity returns the MinBootDiskCapacity field if non-nil, zero value otherwise.

### GetMinBootDiskCapacityOk

`func (o *MinimalTemplateProfile) GetMinBootDiskCapacityOk() (*int32, bool)`

GetMinBootDiskCapacityOk returns a tuple with the MinBootDiskCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBootDiskCapacity

`func (o *MinimalTemplateProfile) SetMinBootDiskCapacity(v int32)`

SetMinBootDiskCapacity sets MinBootDiskCapacity field to given value.

### HasMinBootDiskCapacity

`func (o *MinimalTemplateProfile) HasMinBootDiskCapacity() bool`

HasMinBootDiskCapacity returns a boolean if a field has been set.

### GetMinNumCpus

`func (o *MinimalTemplateProfile) GetMinNumCpus() int32`

GetMinNumCpus returns the MinNumCpus field if non-nil, zero value otherwise.

### GetMinNumCpusOk

`func (o *MinimalTemplateProfile) GetMinNumCpusOk() (*int32, bool)`

GetMinNumCpusOk returns a tuple with the MinNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumCpus

`func (o *MinimalTemplateProfile) SetMinNumCpus(v int32)`

SetMinNumCpus sets MinNumCpus field to given value.

### HasMinNumCpus

`func (o *MinimalTemplateProfile) HasMinNumCpus() bool`

HasMinNumCpus returns a boolean if a field has been set.

### GetMinRam

`func (o *MinimalTemplateProfile) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *MinimalTemplateProfile) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *MinimalTemplateProfile) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *MinimalTemplateProfile) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### GetAdditionalDisks

`func (o *MinimalTemplateProfile) GetAdditionalDisks() []MinimalDisk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *MinimalTemplateProfile) GetAdditionalDisksOk() (*[]MinimalDisk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *MinimalTemplateProfile) SetAdditionalDisks(v []MinimalDisk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *MinimalTemplateProfile) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetCons3rtAgentRequired

`func (o *MinimalTemplateProfile) GetCons3rtAgentRequired() bool`

GetCons3rtAgentRequired returns the Cons3rtAgentRequired field if non-nil, zero value otherwise.

### GetCons3rtAgentRequiredOk

`func (o *MinimalTemplateProfile) GetCons3rtAgentRequiredOk() (*bool, bool)`

GetCons3rtAgentRequiredOk returns a tuple with the Cons3rtAgentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtAgentRequired

`func (o *MinimalTemplateProfile) SetCons3rtAgentRequired(v bool)`

SetCons3rtAgentRequired sets Cons3rtAgentRequired field to given value.

### HasCons3rtAgentRequired

`func (o *MinimalTemplateProfile) HasCons3rtAgentRequired() bool`

HasCons3rtAgentRequired returns a boolean if a field has been set.

### GetVgpuRequired

`func (o *MinimalTemplateProfile) GetVgpuRequired() bool`

GetVgpuRequired returns the VgpuRequired field if non-nil, zero value otherwise.

### GetVgpuRequiredOk

`func (o *MinimalTemplateProfile) GetVgpuRequiredOk() (*bool, bool)`

GetVgpuRequiredOk returns a tuple with the VgpuRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuRequired

`func (o *MinimalTemplateProfile) SetVgpuRequired(v bool)`

SetVgpuRequired sets VgpuRequired field to given value.

### HasVgpuRequired

`func (o *MinimalTemplateProfile) HasVgpuRequired() bool`

HasVgpuRequired returns a boolean if a field has been set.

### GetGpuType

`func (o *MinimalTemplateProfile) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *MinimalTemplateProfile) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *MinimalTemplateProfile) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *MinimalTemplateProfile) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetRequiresNestedVirtualization

`func (o *MinimalTemplateProfile) GetRequiresNestedVirtualization() bool`

GetRequiresNestedVirtualization returns the RequiresNestedVirtualization field if non-nil, zero value otherwise.

### GetRequiresNestedVirtualizationOk

`func (o *MinimalTemplateProfile) GetRequiresNestedVirtualizationOk() (*bool, bool)`

GetRequiresNestedVirtualizationOk returns a tuple with the RequiresNestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresNestedVirtualization

`func (o *MinimalTemplateProfile) SetRequiresNestedVirtualization(v bool)`

SetRequiresNestedVirtualization sets RequiresNestedVirtualization field to given value.

### HasRequiresNestedVirtualization

`func (o *MinimalTemplateProfile) HasRequiresNestedVirtualization() bool`

HasRequiresNestedVirtualization returns a boolean if a field has been set.

### GetRemoteAccessRequired

`func (o *MinimalTemplateProfile) GetRemoteAccessRequired() bool`

GetRemoteAccessRequired returns the RemoteAccessRequired field if non-nil, zero value otherwise.

### GetRemoteAccessRequiredOk

`func (o *MinimalTemplateProfile) GetRemoteAccessRequiredOk() (*bool, bool)`

GetRemoteAccessRequiredOk returns a tuple with the RemoteAccessRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessRequired

`func (o *MinimalTemplateProfile) SetRemoteAccessRequired(v bool)`

SetRemoteAccessRequired sets RemoteAccessRequired field to given value.

### HasRemoteAccessRequired

`func (o *MinimalTemplateProfile) HasRemoteAccessRequired() bool`

HasRemoteAccessRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


