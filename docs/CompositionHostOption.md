# CompositionHostOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDisks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**BatchSoftwareInstall** | Pointer to **bool** |  | [optional] 
**InstanceTypeName** | Pointer to **string** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) |  | [optional] 
**Cpus** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**SystemRole** | **string** |  | 
**TemplateName** | Pointer to **string** |  | [optional] 

## Methods

### NewCompositionHostOption

`func NewCompositionHostOption(systemRole string, ) *CompositionHostOption`

NewCompositionHostOption instantiates a new CompositionHostOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositionHostOptionWithDefaults

`func NewCompositionHostOptionWithDefaults() *CompositionHostOption`

NewCompositionHostOptionWithDefaults instantiates a new CompositionHostOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDisks

`func (o *CompositionHostOption) GetAdditionalDisks() []Disk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *CompositionHostOption) GetAdditionalDisksOk() (*[]Disk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *CompositionHostOption) SetAdditionalDisks(v []Disk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *CompositionHostOption) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetBatchSoftwareInstall

`func (o *CompositionHostOption) GetBatchSoftwareInstall() bool`

GetBatchSoftwareInstall returns the BatchSoftwareInstall field if non-nil, zero value otherwise.

### GetBatchSoftwareInstallOk

`func (o *CompositionHostOption) GetBatchSoftwareInstallOk() (*bool, bool)`

GetBatchSoftwareInstallOk returns a tuple with the BatchSoftwareInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSoftwareInstall

`func (o *CompositionHostOption) SetBatchSoftwareInstall(v bool)`

SetBatchSoftwareInstall sets BatchSoftwareInstall field to given value.

### HasBatchSoftwareInstall

`func (o *CompositionHostOption) HasBatchSoftwareInstall() bool`

HasBatchSoftwareInstall returns a boolean if a field has been set.

### GetInstanceTypeName

`func (o *CompositionHostOption) GetInstanceTypeName() string`

GetInstanceTypeName returns the InstanceTypeName field if non-nil, zero value otherwise.

### GetInstanceTypeNameOk

`func (o *CompositionHostOption) GetInstanceTypeNameOk() (*string, bool)`

GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeName

`func (o *CompositionHostOption) SetInstanceTypeName(v string)`

SetInstanceTypeName sets InstanceTypeName field to given value.

### HasInstanceTypeName

`func (o *CompositionHostOption) HasInstanceTypeName() bool`

HasInstanceTypeName returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *CompositionHostOption) GetNetworkInterfaces() []NetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *CompositionHostOption) GetNetworkInterfacesOk() (*[]NetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *CompositionHostOption) SetNetworkInterfaces(v []NetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *CompositionHostOption) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetCpus

`func (o *CompositionHostOption) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CompositionHostOption) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CompositionHostOption) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CompositionHostOption) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetRam

`func (o *CompositionHostOption) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CompositionHostOption) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CompositionHostOption) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *CompositionHostOption) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetSystemRole

`func (o *CompositionHostOption) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *CompositionHostOption) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *CompositionHostOption) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.


### GetTemplateName

`func (o *CompositionHostOption) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *CompositionHostOption) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *CompositionHostOption) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *CompositionHostOption) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


