# InputHostOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRole** | **string** |  | 
**Cpus** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**AdditionalDisks** | Pointer to [**[]InputDisk**](InputDisk.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]InputNetworkInterface**](InputNetworkInterface.md) |  | [optional] 

## Methods

### NewInputHostOption

`func NewInputHostOption(systemRole string, ) *InputHostOption`

NewInputHostOption instantiates a new InputHostOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputHostOptionWithDefaults

`func NewInputHostOptionWithDefaults() *InputHostOption`

NewInputHostOptionWithDefaults instantiates a new InputHostOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRole

`func (o *InputHostOption) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *InputHostOption) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *InputHostOption) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.


### GetCpus

`func (o *InputHostOption) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *InputHostOption) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *InputHostOption) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *InputHostOption) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetRam

`func (o *InputHostOption) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InputHostOption) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InputHostOption) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InputHostOption) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetAdditionalDisks

`func (o *InputHostOption) GetAdditionalDisks() []InputDisk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *InputHostOption) GetAdditionalDisksOk() (*[]InputDisk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *InputHostOption) SetAdditionalDisks(v []InputDisk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *InputHostOption) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *InputHostOption) GetNetworkInterfaces() []InputNetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InputHostOption) GetNetworkInterfacesOk() (*[]InputNetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InputHostOption) SetNetworkInterfaces(v []InputNetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InputHostOption) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


