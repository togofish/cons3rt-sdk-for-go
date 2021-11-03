# MinimalDeploymentHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDisks** | Pointer to [**[]FullDisk**](FullDisk.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**NumCpus** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**BuildOrder** | Pointer to **int32** |  | [optional] 
**SystemModule** | Pointer to [**MinimalSystemModule**](MinimalSystemModule.md) |  | [optional] 
**SystemRole** | Pointer to **string** |  | [optional] 

## Methods

### NewMinimalDeploymentHost

`func NewMinimalDeploymentHost() *MinimalDeploymentHost`

NewMinimalDeploymentHost instantiates a new MinimalDeploymentHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalDeploymentHostWithDefaults

`func NewMinimalDeploymentHostWithDefaults() *MinimalDeploymentHost`

NewMinimalDeploymentHostWithDefaults instantiates a new MinimalDeploymentHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDisks

`func (o *MinimalDeploymentHost) GetAdditionalDisks() []FullDisk`

GetAdditionalDisks returns the AdditionalDisks field if non-nil, zero value otherwise.

### GetAdditionalDisksOk

`func (o *MinimalDeploymentHost) GetAdditionalDisksOk() (*[]FullDisk, bool)`

GetAdditionalDisksOk returns a tuple with the AdditionalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisks

`func (o *MinimalDeploymentHost) SetAdditionalDisks(v []FullDisk)`

SetAdditionalDisks sets AdditionalDisks field to given value.

### HasAdditionalDisks

`func (o *MinimalDeploymentHost) HasAdditionalDisks() bool`

HasAdditionalDisks returns a boolean if a field has been set.

### GetId

`func (o *MinimalDeploymentHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalDeploymentHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalDeploymentHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalDeploymentHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumCpus

`func (o *MinimalDeploymentHost) GetNumCpus() int32`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *MinimalDeploymentHost) GetNumCpusOk() (*int32, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *MinimalDeploymentHost) SetNumCpus(v int32)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *MinimalDeploymentHost) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetRam

`func (o *MinimalDeploymentHost) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *MinimalDeploymentHost) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *MinimalDeploymentHost) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *MinimalDeploymentHost) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetBuildOrder

`func (o *MinimalDeploymentHost) GetBuildOrder() int32`

GetBuildOrder returns the BuildOrder field if non-nil, zero value otherwise.

### GetBuildOrderOk

`func (o *MinimalDeploymentHost) GetBuildOrderOk() (*int32, bool)`

GetBuildOrderOk returns a tuple with the BuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOrder

`func (o *MinimalDeploymentHost) SetBuildOrder(v int32)`

SetBuildOrder sets BuildOrder field to given value.

### HasBuildOrder

`func (o *MinimalDeploymentHost) HasBuildOrder() bool`

HasBuildOrder returns a boolean if a field has been set.

### GetSystemModule

`func (o *MinimalDeploymentHost) GetSystemModule() MinimalSystemModule`

GetSystemModule returns the SystemModule field if non-nil, zero value otherwise.

### GetSystemModuleOk

`func (o *MinimalDeploymentHost) GetSystemModuleOk() (*MinimalSystemModule, bool)`

GetSystemModuleOk returns a tuple with the SystemModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModule

`func (o *MinimalDeploymentHost) SetSystemModule(v MinimalSystemModule)`

SetSystemModule sets SystemModule field to given value.

### HasSystemModule

`func (o *MinimalDeploymentHost) HasSystemModule() bool`

HasSystemModule returns a boolean if a field has been set.

### GetSystemRole

`func (o *MinimalDeploymentHost) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *MinimalDeploymentHost) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *MinimalDeploymentHost) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.

### HasSystemRole

`func (o *MinimalDeploymentHost) HasSystemRole() bool`

HasSystemRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


