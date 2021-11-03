# InputVCloudRestCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**Password** | **string** |  | 
**Port** | **int32** |  | 
**Protocol** | **string** |  | 
**RemoteAccessInternalIp** | **string** |  | 
**RemoteAccessPort** | **int32** |  | 
**Username** | **string** |  | 
**GpuProfileTypes** | Pointer to **[]string** |  | [optional] 
**VsphereApiVersion** | Pointer to **string** |  | [optional] 
**VsphereHost** | Pointer to **string** |  | [optional] 
**VspherePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputVCloudRestCloud

`func NewInputVCloudRestCloud(hostname string, password string, port int32, protocol string, remoteAccessInternalIp string, remoteAccessPort int32, username string, ) *InputVCloudRestCloud`

NewInputVCloudRestCloud instantiates a new InputVCloudRestCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVCloudRestCloudWithDefaults

`func NewInputVCloudRestCloudWithDefaults() *InputVCloudRestCloud`

NewInputVCloudRestCloudWithDefaults instantiates a new InputVCloudRestCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *InputVCloudRestCloud) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InputVCloudRestCloud) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InputVCloudRestCloud) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPassword

`func (o *InputVCloudRestCloud) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputVCloudRestCloud) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputVCloudRestCloud) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPort

`func (o *InputVCloudRestCloud) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InputVCloudRestCloud) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InputVCloudRestCloud) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *InputVCloudRestCloud) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InputVCloudRestCloud) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InputVCloudRestCloud) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRemoteAccessInternalIp

`func (o *InputVCloudRestCloud) GetRemoteAccessInternalIp() string`

GetRemoteAccessInternalIp returns the RemoteAccessInternalIp field if non-nil, zero value otherwise.

### GetRemoteAccessInternalIpOk

`func (o *InputVCloudRestCloud) GetRemoteAccessInternalIpOk() (*string, bool)`

GetRemoteAccessInternalIpOk returns a tuple with the RemoteAccessInternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessInternalIp

`func (o *InputVCloudRestCloud) SetRemoteAccessInternalIp(v string)`

SetRemoteAccessInternalIp sets RemoteAccessInternalIp field to given value.


### GetRemoteAccessPort

`func (o *InputVCloudRestCloud) GetRemoteAccessPort() int32`

GetRemoteAccessPort returns the RemoteAccessPort field if non-nil, zero value otherwise.

### GetRemoteAccessPortOk

`func (o *InputVCloudRestCloud) GetRemoteAccessPortOk() (*int32, bool)`

GetRemoteAccessPortOk returns a tuple with the RemoteAccessPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessPort

`func (o *InputVCloudRestCloud) SetRemoteAccessPort(v int32)`

SetRemoteAccessPort sets RemoteAccessPort field to given value.


### GetUsername

`func (o *InputVCloudRestCloud) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InputVCloudRestCloud) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InputVCloudRestCloud) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetGpuProfileTypes

`func (o *InputVCloudRestCloud) GetGpuProfileTypes() []string`

GetGpuProfileTypes returns the GpuProfileTypes field if non-nil, zero value otherwise.

### GetGpuProfileTypesOk

`func (o *InputVCloudRestCloud) GetGpuProfileTypesOk() (*[]string, bool)`

GetGpuProfileTypesOk returns a tuple with the GpuProfileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuProfileTypes

`func (o *InputVCloudRestCloud) SetGpuProfileTypes(v []string)`

SetGpuProfileTypes sets GpuProfileTypes field to given value.

### HasGpuProfileTypes

`func (o *InputVCloudRestCloud) HasGpuProfileTypes() bool`

HasGpuProfileTypes returns a boolean if a field has been set.

### GetVsphereApiVersion

`func (o *InputVCloudRestCloud) GetVsphereApiVersion() string`

GetVsphereApiVersion returns the VsphereApiVersion field if non-nil, zero value otherwise.

### GetVsphereApiVersionOk

`func (o *InputVCloudRestCloud) GetVsphereApiVersionOk() (*string, bool)`

GetVsphereApiVersionOk returns a tuple with the VsphereApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereApiVersion

`func (o *InputVCloudRestCloud) SetVsphereApiVersion(v string)`

SetVsphereApiVersion sets VsphereApiVersion field to given value.

### HasVsphereApiVersion

`func (o *InputVCloudRestCloud) HasVsphereApiVersion() bool`

HasVsphereApiVersion returns a boolean if a field has been set.

### GetVsphereHost

`func (o *InputVCloudRestCloud) GetVsphereHost() string`

GetVsphereHost returns the VsphereHost field if non-nil, zero value otherwise.

### GetVsphereHostOk

`func (o *InputVCloudRestCloud) GetVsphereHostOk() (*string, bool)`

GetVsphereHostOk returns a tuple with the VsphereHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHost

`func (o *InputVCloudRestCloud) SetVsphereHost(v string)`

SetVsphereHost sets VsphereHost field to given value.

### HasVsphereHost

`func (o *InputVCloudRestCloud) HasVsphereHost() bool`

HasVsphereHost returns a boolean if a field has been set.

### GetVspherePort

`func (o *InputVCloudRestCloud) GetVspherePort() int32`

GetVspherePort returns the VspherePort field if non-nil, zero value otherwise.

### GetVspherePortOk

`func (o *InputVCloudRestCloud) GetVspherePortOk() (*int32, bool)`

GetVspherePortOk returns a tuple with the VspherePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVspherePort

`func (o *InputVCloudRestCloud) SetVspherePort(v int32)`

SetVspherePort sets VspherePort field to given value.

### HasVspherePort

`func (o *InputVCloudRestCloud) HasVspherePort() bool`

HasVspherePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


