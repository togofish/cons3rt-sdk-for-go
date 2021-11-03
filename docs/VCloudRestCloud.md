# VCloudRestCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuProfileTypes** | Pointer to **[]string** |  | [optional] 
**Hostname** | **string** |  | 
**Password** | **string** |  | 
**Port** | **int32** |  | 
**Protocol** | **string** |  | 
**RemoteAccessInternalIp** | **string** |  | 
**RemoteAccessPort** | **int32** |  | 
**StorageServiceHostname** | Pointer to **string** |  | [optional] 
**StorageServicePassword** | Pointer to **string** |  | [optional] 
**StorageServicePort** | Pointer to **int32** |  | [optional] 
**StorageServiceProtocol** | Pointer to **string** |  | [optional] 
**StorageServiceUsername** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 
**VsphereApiVersion** | Pointer to **string** |  | [optional] 
**VsphereHost** | Pointer to **string** |  | [optional] 
**VspherePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewVCloudRestCloud

`func NewVCloudRestCloud(hostname string, password string, port int32, protocol string, remoteAccessInternalIp string, remoteAccessPort int32, username string, ) *VCloudRestCloud`

NewVCloudRestCloud instantiates a new VCloudRestCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCloudRestCloudWithDefaults

`func NewVCloudRestCloudWithDefaults() *VCloudRestCloud`

NewVCloudRestCloudWithDefaults instantiates a new VCloudRestCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuProfileTypes

`func (o *VCloudRestCloud) GetGpuProfileTypes() []string`

GetGpuProfileTypes returns the GpuProfileTypes field if non-nil, zero value otherwise.

### GetGpuProfileTypesOk

`func (o *VCloudRestCloud) GetGpuProfileTypesOk() (*[]string, bool)`

GetGpuProfileTypesOk returns a tuple with the GpuProfileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuProfileTypes

`func (o *VCloudRestCloud) SetGpuProfileTypes(v []string)`

SetGpuProfileTypes sets GpuProfileTypes field to given value.

### HasGpuProfileTypes

`func (o *VCloudRestCloud) HasGpuProfileTypes() bool`

HasGpuProfileTypes returns a boolean if a field has been set.

### GetHostname

`func (o *VCloudRestCloud) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VCloudRestCloud) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VCloudRestCloud) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPassword

`func (o *VCloudRestCloud) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VCloudRestCloud) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VCloudRestCloud) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPort

`func (o *VCloudRestCloud) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VCloudRestCloud) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VCloudRestCloud) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *VCloudRestCloud) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VCloudRestCloud) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VCloudRestCloud) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRemoteAccessInternalIp

`func (o *VCloudRestCloud) GetRemoteAccessInternalIp() string`

GetRemoteAccessInternalIp returns the RemoteAccessInternalIp field if non-nil, zero value otherwise.

### GetRemoteAccessInternalIpOk

`func (o *VCloudRestCloud) GetRemoteAccessInternalIpOk() (*string, bool)`

GetRemoteAccessInternalIpOk returns a tuple with the RemoteAccessInternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessInternalIp

`func (o *VCloudRestCloud) SetRemoteAccessInternalIp(v string)`

SetRemoteAccessInternalIp sets RemoteAccessInternalIp field to given value.


### GetRemoteAccessPort

`func (o *VCloudRestCloud) GetRemoteAccessPort() int32`

GetRemoteAccessPort returns the RemoteAccessPort field if non-nil, zero value otherwise.

### GetRemoteAccessPortOk

`func (o *VCloudRestCloud) GetRemoteAccessPortOk() (*int32, bool)`

GetRemoteAccessPortOk returns a tuple with the RemoteAccessPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessPort

`func (o *VCloudRestCloud) SetRemoteAccessPort(v int32)`

SetRemoteAccessPort sets RemoteAccessPort field to given value.


### GetStorageServiceHostname

`func (o *VCloudRestCloud) GetStorageServiceHostname() string`

GetStorageServiceHostname returns the StorageServiceHostname field if non-nil, zero value otherwise.

### GetStorageServiceHostnameOk

`func (o *VCloudRestCloud) GetStorageServiceHostnameOk() (*string, bool)`

GetStorageServiceHostnameOk returns a tuple with the StorageServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceHostname

`func (o *VCloudRestCloud) SetStorageServiceHostname(v string)`

SetStorageServiceHostname sets StorageServiceHostname field to given value.

### HasStorageServiceHostname

`func (o *VCloudRestCloud) HasStorageServiceHostname() bool`

HasStorageServiceHostname returns a boolean if a field has been set.

### GetStorageServicePassword

`func (o *VCloudRestCloud) GetStorageServicePassword() string`

GetStorageServicePassword returns the StorageServicePassword field if non-nil, zero value otherwise.

### GetStorageServicePasswordOk

`func (o *VCloudRestCloud) GetStorageServicePasswordOk() (*string, bool)`

GetStorageServicePasswordOk returns a tuple with the StorageServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServicePassword

`func (o *VCloudRestCloud) SetStorageServicePassword(v string)`

SetStorageServicePassword sets StorageServicePassword field to given value.

### HasStorageServicePassword

`func (o *VCloudRestCloud) HasStorageServicePassword() bool`

HasStorageServicePassword returns a boolean if a field has been set.

### GetStorageServicePort

`func (o *VCloudRestCloud) GetStorageServicePort() int32`

GetStorageServicePort returns the StorageServicePort field if non-nil, zero value otherwise.

### GetStorageServicePortOk

`func (o *VCloudRestCloud) GetStorageServicePortOk() (*int32, bool)`

GetStorageServicePortOk returns a tuple with the StorageServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServicePort

`func (o *VCloudRestCloud) SetStorageServicePort(v int32)`

SetStorageServicePort sets StorageServicePort field to given value.

### HasStorageServicePort

`func (o *VCloudRestCloud) HasStorageServicePort() bool`

HasStorageServicePort returns a boolean if a field has been set.

### GetStorageServiceProtocol

`func (o *VCloudRestCloud) GetStorageServiceProtocol() string`

GetStorageServiceProtocol returns the StorageServiceProtocol field if non-nil, zero value otherwise.

### GetStorageServiceProtocolOk

`func (o *VCloudRestCloud) GetStorageServiceProtocolOk() (*string, bool)`

GetStorageServiceProtocolOk returns a tuple with the StorageServiceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceProtocol

`func (o *VCloudRestCloud) SetStorageServiceProtocol(v string)`

SetStorageServiceProtocol sets StorageServiceProtocol field to given value.

### HasStorageServiceProtocol

`func (o *VCloudRestCloud) HasStorageServiceProtocol() bool`

HasStorageServiceProtocol returns a boolean if a field has been set.

### GetStorageServiceUsername

`func (o *VCloudRestCloud) GetStorageServiceUsername() string`

GetStorageServiceUsername returns the StorageServiceUsername field if non-nil, zero value otherwise.

### GetStorageServiceUsernameOk

`func (o *VCloudRestCloud) GetStorageServiceUsernameOk() (*string, bool)`

GetStorageServiceUsernameOk returns a tuple with the StorageServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceUsername

`func (o *VCloudRestCloud) SetStorageServiceUsername(v string)`

SetStorageServiceUsername sets StorageServiceUsername field to given value.

### HasStorageServiceUsername

`func (o *VCloudRestCloud) HasStorageServiceUsername() bool`

HasStorageServiceUsername returns a boolean if a field has been set.

### GetUsername

`func (o *VCloudRestCloud) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VCloudRestCloud) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VCloudRestCloud) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVsphereApiVersion

`func (o *VCloudRestCloud) GetVsphereApiVersion() string`

GetVsphereApiVersion returns the VsphereApiVersion field if non-nil, zero value otherwise.

### GetVsphereApiVersionOk

`func (o *VCloudRestCloud) GetVsphereApiVersionOk() (*string, bool)`

GetVsphereApiVersionOk returns a tuple with the VsphereApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereApiVersion

`func (o *VCloudRestCloud) SetVsphereApiVersion(v string)`

SetVsphereApiVersion sets VsphereApiVersion field to given value.

### HasVsphereApiVersion

`func (o *VCloudRestCloud) HasVsphereApiVersion() bool`

HasVsphereApiVersion returns a boolean if a field has been set.

### GetVsphereHost

`func (o *VCloudRestCloud) GetVsphereHost() string`

GetVsphereHost returns the VsphereHost field if non-nil, zero value otherwise.

### GetVsphereHostOk

`func (o *VCloudRestCloud) GetVsphereHostOk() (*string, bool)`

GetVsphereHostOk returns a tuple with the VsphereHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHost

`func (o *VCloudRestCloud) SetVsphereHost(v string)`

SetVsphereHost sets VsphereHost field to given value.

### HasVsphereHost

`func (o *VCloudRestCloud) HasVsphereHost() bool`

HasVsphereHost returns a boolean if a field has been set.

### GetVspherePort

`func (o *VCloudRestCloud) GetVspherePort() int32`

GetVspherePort returns the VspherePort field if non-nil, zero value otherwise.

### GetVspherePortOk

`func (o *VCloudRestCloud) GetVspherePortOk() (*int32, bool)`

GetVspherePortOk returns a tuple with the VspherePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVspherePort

`func (o *VCloudRestCloud) SetVspherePort(v int32)`

SetVspherePort sets VspherePort field to given value.

### HasVspherePort

`func (o *VCloudRestCloud) HasVspherePort() bool`

HasVspherePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


