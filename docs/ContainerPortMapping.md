# ContainerPortMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIp** | **string** |  | 
**ExternalPort** | **int32** |  | 
**InternalPort** | **int32** |  | 
**Protocol** | **string** |  | 

## Methods

### NewContainerPortMapping

`func NewContainerPortMapping(externalIp string, externalPort int32, internalPort int32, protocol string, ) *ContainerPortMapping`

NewContainerPortMapping instantiates a new ContainerPortMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerPortMappingWithDefaults

`func NewContainerPortMappingWithDefaults() *ContainerPortMapping`

NewContainerPortMappingWithDefaults instantiates a new ContainerPortMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIp

`func (o *ContainerPortMapping) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ContainerPortMapping) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ContainerPortMapping) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.


### GetExternalPort

`func (o *ContainerPortMapping) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ContainerPortMapping) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ContainerPortMapping) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.


### GetInternalPort

`func (o *ContainerPortMapping) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ContainerPortMapping) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ContainerPortMapping) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetProtocol

`func (o *ContainerPortMapping) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ContainerPortMapping) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ContainerPortMapping) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


