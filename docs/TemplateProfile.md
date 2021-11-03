# TemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDisks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**Cons3rtAgentRequired** | Pointer to **bool** |  | [optional] 
**VgpuRequired** | Pointer to **bool** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**RequiresNestedVirtualization** | Pointer to **bool** |  | [optional] 
**MinBootDiskCapacity** | Pointer to **int32** |  | [optional] 
**MinNumCpus** | **int32** |  | 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**RemoteAccessRequired** | Pointer to **bool** |  | [optional] 
**MinRam** | **int32** |  | 
**VirtRealmId** | Pointer to **int32** |  | [optional] 
**VirtRealmTemplateName** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateProfile

`func NewTemplateProfile(minNumCpus int32, minRam int32, ) *TemplateProfile`

NewTemplateProfile instantiates a new TemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateProfileWithDefaults

`func NewTemplateProfileWithDefaults() *TemplateProfile`

NewTemplateProfileWithDefaults instantiates a new TemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDisks

`func (o *TemplateProfile) GetAdditionalDisks() []Disk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *TemplateProfile) GetAdditionalDisksOk() (*[]Disk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *TemplateProfile) SetAdditionalDisks(v []Disk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *TemplateProfile) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetCons3rtAgentRequired

`func (o *TemplateProfile) GetCons3rtAgentRequired() bool`

GetCons3rtAgentRequired returns the Cons3rtAgentRequired field if non-nil, zero value otherwise.

### GetCons3rtAgentRequiredOk

`func (o *TemplateProfile) GetCons3rtAgentRequiredOk() (*bool, bool)`

GetCons3rtAgentRequiredOk returns a tuple with the Cons3rtAgentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtAgentRequired

`func (o *TemplateProfile) SetCons3rtAgentRequired(v bool)`

SetCons3rtAgentRequired sets Cons3rtAgentRequired field to given value.

### HasCons3rtAgentRequired

`func (o *TemplateProfile) HasCons3rtAgentRequired() bool`

HasCons3rtAgentRequired returns a boolean if a field has been set.

### GetVgpuRequired

`func (o *TemplateProfile) GetVgpuRequired() bool`

GetVgpuRequired returns the VgpuRequired field if non-nil, zero value otherwise.

### GetVgpuRequiredOk

`func (o *TemplateProfile) GetVgpuRequiredOk() (*bool, bool)`

GetVgpuRequiredOk returns a tuple with the VgpuRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuRequired

`func (o *TemplateProfile) SetVgpuRequired(v bool)`

SetVgpuRequired sets VgpuRequired field to given value.

### HasVgpuRequired

`func (o *TemplateProfile) HasVgpuRequired() bool`

HasVgpuRequired returns a boolean if a field has been set.

### GetGpuType

`func (o *TemplateProfile) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *TemplateProfile) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *TemplateProfile) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *TemplateProfile) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetRequiresNestedVirtualization

`func (o *TemplateProfile) GetRequiresNestedVirtualization() bool`

GetRequiresNestedVirtualization returns the RequiresNestedVirtualization field if non-nil, zero value otherwise.

### GetRequiresNestedVirtualizationOk

`func (o *TemplateProfile) GetRequiresNestedVirtualizationOk() (*bool, bool)`

GetRequiresNestedVirtualizationOk returns a tuple with the RequiresNestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresNestedVirtualization

`func (o *TemplateProfile) SetRequiresNestedVirtualization(v bool)`

SetRequiresNestedVirtualization sets RequiresNestedVirtualization field to given value.

### HasRequiresNestedVirtualization

`func (o *TemplateProfile) HasRequiresNestedVirtualization() bool`

HasRequiresNestedVirtualization returns a boolean if a field has been set.

### GetMinBootDiskCapacity

`func (o *TemplateProfile) GetMinBootDiskCapacity() int32`

GetMinBootDiskCapacity returns the MinBootDiskCapacity field if non-nil, zero value otherwise.

### GetMinBootDiskCapacityOk

`func (o *TemplateProfile) GetMinBootDiskCapacityOk() (*int32, bool)`

GetMinBootDiskCapacityOk returns a tuple with the MinBootDiskCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBootDiskCapacity

`func (o *TemplateProfile) SetMinBootDiskCapacity(v int32)`

SetMinBootDiskCapacity sets MinBootDiskCapacity field to given value.

### HasMinBootDiskCapacity

`func (o *TemplateProfile) HasMinBootDiskCapacity() bool`

HasMinBootDiskCapacity returns a boolean if a field has been set.

### GetMinNumCpus

`func (o *TemplateProfile) GetMinNumCpus() int32`

GetMinNumCpus returns the MinNumCpus field if non-nil, zero value otherwise.

### GetMinNumCpusOk

`func (o *TemplateProfile) GetMinNumCpusOk() (*int32, bool)`

GetMinNumCpusOk returns a tuple with the MinNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumCpus

`func (o *TemplateProfile) SetMinNumCpus(v int32)`

SetMinNumCpus sets MinNumCpus field to given value.


### GetOperatingSystem

`func (o *TemplateProfile) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *TemplateProfile) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *TemplateProfile) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *TemplateProfile) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetRemoteAccessRequired

`func (o *TemplateProfile) GetRemoteAccessRequired() bool`

GetRemoteAccessRequired returns the RemoteAccessRequired field if non-nil, zero value otherwise.

### GetRemoteAccessRequiredOk

`func (o *TemplateProfile) GetRemoteAccessRequiredOk() (*bool, bool)`

GetRemoteAccessRequiredOk returns a tuple with the RemoteAccessRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessRequired

`func (o *TemplateProfile) SetRemoteAccessRequired(v bool)`

SetRemoteAccessRequired sets RemoteAccessRequired field to given value.

### HasRemoteAccessRequired

`func (o *TemplateProfile) HasRemoteAccessRequired() bool`

HasRemoteAccessRequired returns a boolean if a field has been set.

### GetMinRam

`func (o *TemplateProfile) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *TemplateProfile) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *TemplateProfile) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.


### GetVirtRealmId

`func (o *TemplateProfile) GetVirtRealmId() int32`

GetVirtRealmId returns the VirtRealmId field if non-nil, zero value otherwise.

### GetVirtRealmIdOk

`func (o *TemplateProfile) GetVirtRealmIdOk() (*int32, bool)`

GetVirtRealmIdOk returns a tuple with the VirtRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmId

`func (o *TemplateProfile) SetVirtRealmId(v int32)`

SetVirtRealmId sets VirtRealmId field to given value.

### HasVirtRealmId

`func (o *TemplateProfile) HasVirtRealmId() bool`

HasVirtRealmId returns a boolean if a field has been set.

### GetVirtRealmTemplateName

`func (o *TemplateProfile) GetVirtRealmTemplateName() string`

GetVirtRealmTemplateName returns the VirtRealmTemplateName field if non-nil, zero value otherwise.

### GetVirtRealmTemplateNameOk

`func (o *TemplateProfile) GetVirtRealmTemplateNameOk() (*string, bool)`

GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmTemplateName

`func (o *TemplateProfile) SetVirtRealmTemplateName(v string)`

SetVirtRealmTemplateName sets VirtRealmTemplateName field to given value.

### HasVirtRealmTemplateName

`func (o *TemplateProfile) HasVirtRealmTemplateName() bool`

HasVirtRealmTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


