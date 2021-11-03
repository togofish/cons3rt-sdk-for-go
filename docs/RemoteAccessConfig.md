# RemoteAccessConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RemoteAccessIpAddress** | **string** |  | 
**RemoteAccessPort** | **int32** |  | 
**InstanceType** | **string** |  | 
**TemplateName** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoteAccessConfig

`func NewRemoteAccessConfig(remoteAccessIpAddress string, remoteAccessPort int32, instanceType string, ) *RemoteAccessConfig`

NewRemoteAccessConfig instantiates a new RemoteAccessConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAccessConfigWithDefaults

`func NewRemoteAccessConfigWithDefaults() *RemoteAccessConfig`

NewRemoteAccessConfigWithDefaults instantiates a new RemoteAccessConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteAccessConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteAccessConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteAccessConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteAccessConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteAccessIpAddress

`func (o *RemoteAccessConfig) GetRemoteAccessIpAddress() string`

GetRemoteAccessIpAddress returns the RemoteAccessIpAddress field if non-nil, zero value otherwise.

### GetRemoteAccessIpAddressOk

`func (o *RemoteAccessConfig) GetRemoteAccessIpAddressOk() (*string, bool)`

GetRemoteAccessIpAddressOk returns a tuple with the RemoteAccessIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessIpAddress

`func (o *RemoteAccessConfig) SetRemoteAccessIpAddress(v string)`

SetRemoteAccessIpAddress sets RemoteAccessIpAddress field to given value.


### GetRemoteAccessPort

`func (o *RemoteAccessConfig) GetRemoteAccessPort() int32`

GetRemoteAccessPort returns the RemoteAccessPort field if non-nil, zero value otherwise.

### GetRemoteAccessPortOk

`func (o *RemoteAccessConfig) GetRemoteAccessPortOk() (*int32, bool)`

GetRemoteAccessPortOk returns a tuple with the RemoteAccessPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessPort

`func (o *RemoteAccessConfig) SetRemoteAccessPort(v int32)`

SetRemoteAccessPort sets RemoteAccessPort field to given value.


### GetInstanceType

`func (o *RemoteAccessConfig) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RemoteAccessConfig) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RemoteAccessConfig) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetTemplateName

`func (o *RemoteAccessConfig) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *RemoteAccessConfig) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *RemoteAccessConfig) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *RemoteAccessConfig) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


