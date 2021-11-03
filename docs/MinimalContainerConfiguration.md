# MinimalContainerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PortMapping** | Pointer to [**ContainerPortMapping**](ContainerPortMapping.md) |  | [optional] 
**ContainerName** | **string** |  | 
**RunArguments** | Pointer to **string** |  | [optional] 

## Methods

### NewMinimalContainerConfiguration

`func NewMinimalContainerConfiguration(containerName string, ) *MinimalContainerConfiguration`

NewMinimalContainerConfiguration instantiates a new MinimalContainerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalContainerConfigurationWithDefaults

`func NewMinimalContainerConfigurationWithDefaults() *MinimalContainerConfiguration`

NewMinimalContainerConfigurationWithDefaults instantiates a new MinimalContainerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalContainerConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalContainerConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalContainerConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalContainerConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortMapping

`func (o *MinimalContainerConfiguration) GetPortMapping() ContainerPortMapping`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *MinimalContainerConfiguration) GetPortMappingOk() (*ContainerPortMapping, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *MinimalContainerConfiguration) SetPortMapping(v ContainerPortMapping)`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *MinimalContainerConfiguration) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetContainerName

`func (o *MinimalContainerConfiguration) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *MinimalContainerConfiguration) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *MinimalContainerConfiguration) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetRunArguments

`func (o *MinimalContainerConfiguration) GetRunArguments() string`

GetRunArguments returns the RunArguments field if non-nil, zero value otherwise.

### GetRunArgumentsOk

`func (o *MinimalContainerConfiguration) GetRunArgumentsOk() (*string, bool)`

GetRunArgumentsOk returns a tuple with the RunArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunArguments

`func (o *MinimalContainerConfiguration) SetRunArguments(v string)`

SetRunArguments sets RunArguments field to given value.

### HasRunArguments

`func (o *MinimalContainerConfiguration) HasRunArguments() bool`

HasRunArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


