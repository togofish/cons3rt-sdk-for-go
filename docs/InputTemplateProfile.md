# InputTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtRealmTemplateName** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**MinNumCpus** | **int32** |  | 
**MinRam** | **int32** |  | 
**VgpuRequired** | Pointer to **bool** |  | [optional] 
**RequiresNestedVirtualization** | Pointer to **bool** |  | [optional] 
**AdditionalDisks** | Pointer to [**[]InputDisk**](InputDisk.md) |  | [optional] 
**MinBootDiskCapacity** | Pointer to **int32** |  | [optional] 
**RemoteAccessRequired** | Pointer to **bool** |  | [optional] 
**VirtRealmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputTemplateProfile

`func NewInputTemplateProfile(minNumCpus int32, minRam int32, ) *InputTemplateProfile`

NewInputTemplateProfile instantiates a new InputTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTemplateProfileWithDefaults

`func NewInputTemplateProfileWithDefaults() *InputTemplateProfile`

NewInputTemplateProfileWithDefaults instantiates a new InputTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtRealmTemplateName

`func (o *InputTemplateProfile) GetVirtRealmTemplateName() string`

GetVirtRealmTemplateName returns the VirtRealmTemplateName field if non-nil, zero value otherwise.

### GetVirtRealmTemplateNameOk

`func (o *InputTemplateProfile) GetVirtRealmTemplateNameOk() (*string, bool)`

GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmTemplateName

`func (o *InputTemplateProfile) SetVirtRealmTemplateName(v string)`

SetVirtRealmTemplateName sets VirtRealmTemplateName field to given value.

### HasVirtRealmTemplateName

`func (o *InputTemplateProfile) HasVirtRealmTemplateName() bool`

HasVirtRealmTemplateName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *InputTemplateProfile) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *InputTemplateProfile) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *InputTemplateProfile) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *InputTemplateProfile) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetMinNumCpus

`func (o *InputTemplateProfile) GetMinNumCpus() int32`

GetMinNumCpus returns the MinNumCpus field if non-nil, zero value otherwise.

### GetMinNumCpusOk

`func (o *InputTemplateProfile) GetMinNumCpusOk() (*int32, bool)`

GetMinNumCpusOk returns a tuple with the MinNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumCpus

`func (o *InputTemplateProfile) SetMinNumCpus(v int32)`

SetMinNumCpus sets MinNumCpus field to given value.


### GetMinRam

`func (o *InputTemplateProfile) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *InputTemplateProfile) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *InputTemplateProfile) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.


### GetVgpuRequired

`func (o *InputTemplateProfile) GetVgpuRequired() bool`

GetVgpuRequired returns the VgpuRequired field if non-nil, zero value otherwise.

### GetVgpuRequiredOk

`func (o *InputTemplateProfile) GetVgpuRequiredOk() (*bool, bool)`

GetVgpuRequiredOk returns a tuple with the VgpuRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuRequired

`func (o *InputTemplateProfile) SetVgpuRequired(v bool)`

SetVgpuRequired sets VgpuRequired field to given value.

### HasVgpuRequired

`func (o *InputTemplateProfile) HasVgpuRequired() bool`

HasVgpuRequired returns a boolean if a field has been set.

### GetRequiresNestedVirtualization

`func (o *InputTemplateProfile) GetRequiresNestedVirtualization() bool`

GetRequiresNestedVirtualization returns the RequiresNestedVirtualization field if non-nil, zero value otherwise.

### GetRequiresNestedVirtualizationOk

`func (o *InputTemplateProfile) GetRequiresNestedVirtualizationOk() (*bool, bool)`

GetRequiresNestedVirtualizationOk returns a tuple with the RequiresNestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresNestedVirtualization

`func (o *InputTemplateProfile) SetRequiresNestedVirtualization(v bool)`

SetRequiresNestedVirtualization sets RequiresNestedVirtualization field to given value.

### HasRequiresNestedVirtualization

`func (o *InputTemplateProfile) HasRequiresNestedVirtualization() bool`

HasRequiresNestedVirtualization returns a boolean if a field has been set.

### GetAdditionalDisks

`func (o *InputTemplateProfile) GetAdditionalDisks() []InputDisk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *InputTemplateProfile) GetAdditionalDisksOk() (*[]InputDisk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *InputTemplateProfile) SetAdditionalDisks(v []InputDisk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *InputTemplateProfile) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetMinBootDiskCapacity

`func (o *InputTemplateProfile) GetMinBootDiskCapacity() int32`

GetMinBootDiskCapacity returns the MinBootDiskCapacity field if non-nil, zero value otherwise.

### GetMinBootDiskCapacityOk

`func (o *InputTemplateProfile) GetMinBootDiskCapacityOk() (*int32, bool)`

GetMinBootDiskCapacityOk returns a tuple with the MinBootDiskCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBootDiskCapacity

`func (o *InputTemplateProfile) SetMinBootDiskCapacity(v int32)`

SetMinBootDiskCapacity sets MinBootDiskCapacity field to given value.

### HasMinBootDiskCapacity

`func (o *InputTemplateProfile) HasMinBootDiskCapacity() bool`

HasMinBootDiskCapacity returns a boolean if a field has been set.

### GetRemoteAccessRequired

`func (o *InputTemplateProfile) GetRemoteAccessRequired() bool`

GetRemoteAccessRequired returns the RemoteAccessRequired field if non-nil, zero value otherwise.

### GetRemoteAccessRequiredOk

`func (o *InputTemplateProfile) GetRemoteAccessRequiredOk() (*bool, bool)`

GetRemoteAccessRequiredOk returns a tuple with the RemoteAccessRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessRequired

`func (o *InputTemplateProfile) SetRemoteAccessRequired(v bool)`

SetRemoteAccessRequired sets RemoteAccessRequired field to given value.

### HasRemoteAccessRequired

`func (o *InputTemplateProfile) HasRemoteAccessRequired() bool`

HasRemoteAccessRequired returns a boolean if a field has been set.

### GetVirtRealmId

`func (o *InputTemplateProfile) GetVirtRealmId() int32`

GetVirtRealmId returns the VirtRealmId field if non-nil, zero value otherwise.

### GetVirtRealmIdOk

`func (o *InputTemplateProfile) GetVirtRealmIdOk() (*int32, bool)`

GetVirtRealmIdOk returns a tuple with the VirtRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmId

`func (o *InputTemplateProfile) SetVirtRealmId(v int32)`

SetVirtRealmId sets VirtRealmId field to given value.

### HasVirtRealmId

`func (o *InputTemplateProfile) HasVirtRealmId() bool`

HasVirtRealmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


