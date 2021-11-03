# InputCompositionHostOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRole** | **string** |  | 
**Cpus** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**AdditionalDisks** | Pointer to [**[]InputDisk**](InputDisk.md) |  | [optional] 
**InstanceTypeName** | Pointer to **string** | Required for Azure and AWS EC2 Cloudspaces | [optional] 
**NetworkInterfaces** | Pointer to [**[]InputNetworkInterface**](InputNetworkInterface.md) |  | [optional] 
**TemplateName** | **string** |  | 

## Methods

### NewInputCompositionHostOption

`func NewInputCompositionHostOption(systemRole string, templateName string, ) *InputCompositionHostOption`

NewInputCompositionHostOption instantiates a new InputCompositionHostOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCompositionHostOptionWithDefaults

`func NewInputCompositionHostOptionWithDefaults() *InputCompositionHostOption`

NewInputCompositionHostOptionWithDefaults instantiates a new InputCompositionHostOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRole

`func (o *InputCompositionHostOption) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *InputCompositionHostOption) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *InputCompositionHostOption) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.


### GetCpus

`func (o *InputCompositionHostOption) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *InputCompositionHostOption) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *InputCompositionHostOption) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *InputCompositionHostOption) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetRam

`func (o *InputCompositionHostOption) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InputCompositionHostOption) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InputCompositionHostOption) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InputCompositionHostOption) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetAdditionalDisks

`func (o *InputCompositionHostOption) GetAdditionalDisks() []InputDisk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *InputCompositionHostOption) GetAdditionalDisksOk() (*[]InputDisk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *InputCompositionHostOption) SetAdditionalDisks(v []InputDisk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *InputCompositionHostOption) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetInstanceTypeName

`func (o *InputCompositionHostOption) GetInstanceTypeName() string`

GetInstanceTypeName returns the InstanceTypeName field if non-nil, zero value otherwise.

### GetInstanceTypeNameOk

`func (o *InputCompositionHostOption) GetInstanceTypeNameOk() (*string, bool)`

GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeName

`func (o *InputCompositionHostOption) SetInstanceTypeName(v string)`

SetInstanceTypeName sets InstanceTypeName field to given value.

### HasInstanceTypeName

`func (o *InputCompositionHostOption) HasInstanceTypeName() bool`

HasInstanceTypeName returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *InputCompositionHostOption) GetNetworkInterfaces() []InputNetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InputCompositionHostOption) GetNetworkInterfacesOk() (*[]InputNetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InputCompositionHostOption) SetNetworkInterfaces(v []InputNetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InputCompositionHostOption) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetTemplateName

`func (o *InputCompositionHostOption) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *InputCompositionHostOption) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *InputCompositionHostOption) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


