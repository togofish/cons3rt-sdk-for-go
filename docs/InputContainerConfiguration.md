# InputContainerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ContainerName** | **string** |  | 
**ContainerRuntime** | Pointer to **string** |  | [optional] 
**EnvironmentMap** | Pointer to **map[string]string** |  | [optional] 
**Mounts** | Pointer to [**[]ContainerMount**](ContainerMount.md) |  | [optional] 
**PortMappings** | Pointer to [**[]ContainerPortMapping**](ContainerPortMapping.md) |  | [optional] 
**RunArguments** | Pointer to **string** |  | [optional] 
**RunDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputContainerConfiguration

`func NewInputContainerConfiguration(containerName string, ) *InputContainerConfiguration`

NewInputContainerConfiguration instantiates a new InputContainerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputContainerConfigurationWithDefaults

`func NewInputContainerConfigurationWithDefaults() *InputContainerConfiguration`

NewInputContainerConfigurationWithDefaults instantiates a new InputContainerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputContainerConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputContainerConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputContainerConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputContainerConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContainerName

`func (o *InputContainerConfiguration) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *InputContainerConfiguration) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *InputContainerConfiguration) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetContainerRuntime

`func (o *InputContainerConfiguration) GetContainerRuntime() string`

GetContainerRuntime returns the ContainerRuntime field if non-nil, zero value otherwise.

### GetContainerRuntimeOk

`func (o *InputContainerConfiguration) GetContainerRuntimeOk() (*string, bool)`

GetContainerRuntimeOk returns a tuple with the ContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntime

`func (o *InputContainerConfiguration) SetContainerRuntime(v string)`

SetContainerRuntime sets ContainerRuntime field to given value.

### HasContainerRuntime

`func (o *InputContainerConfiguration) HasContainerRuntime() bool`

HasContainerRuntime returns a boolean if a field has been set.

### GetEnvironmentMap

`func (o *InputContainerConfiguration) GetEnvironmentMap() map[string]string`

GetEnvironmentMap returns the EnvironmentMap field if non-nil, zero value otherwise.

### GetEnvironmentMapOk

`func (o *InputContainerConfiguration) GetEnvironmentMapOk() (*map[string]string, bool)`

GetEnvironmentMapOk returns a tuple with the EnvironmentMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentMap

`func (o *InputContainerConfiguration) SetEnvironmentMap(v map[string]string)`

SetEnvironmentMap sets EnvironmentMap field to given value.

### HasEnvironmentMap

`func (o *InputContainerConfiguration) HasEnvironmentMap() bool`

HasEnvironmentMap returns a boolean if a field has been set.

### GetMounts

`func (o *InputContainerConfiguration) GetMounts() []ContainerMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *InputContainerConfiguration) GetMountsOk() (*[]ContainerMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *InputContainerConfiguration) SetMounts(v []ContainerMount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *InputContainerConfiguration) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetPortMappings

`func (o *InputContainerConfiguration) GetPortMappings() []ContainerPortMapping`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *InputContainerConfiguration) GetPortMappingsOk() (*[]ContainerPortMapping, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *InputContainerConfiguration) SetPortMappings(v []ContainerPortMapping)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *InputContainerConfiguration) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetRunArguments

`func (o *InputContainerConfiguration) GetRunArguments() string`

GetRunArguments returns the RunArguments field if non-nil, zero value otherwise.

### GetRunArgumentsOk

`func (o *InputContainerConfiguration) GetRunArgumentsOk() (*string, bool)`

GetRunArgumentsOk returns a tuple with the RunArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunArguments

`func (o *InputContainerConfiguration) SetRunArguments(v string)`

SetRunArguments sets RunArguments field to given value.

### HasRunArguments

`func (o *InputContainerConfiguration) HasRunArguments() bool`

HasRunArguments returns a boolean if a field has been set.

### GetRunDisabled

`func (o *InputContainerConfiguration) GetRunDisabled() bool`

GetRunDisabled returns the RunDisabled field if non-nil, zero value otherwise.

### GetRunDisabledOk

`func (o *InputContainerConfiguration) GetRunDisabledOk() (*bool, bool)`

GetRunDisabledOk returns a tuple with the RunDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDisabled

`func (o *InputContainerConfiguration) SetRunDisabled(v bool)`

SetRunDisabled sets RunDisabled field to given value.

### HasRunDisabled

`func (o *InputContainerConfiguration) HasRunDisabled() bool`

HasRunDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


