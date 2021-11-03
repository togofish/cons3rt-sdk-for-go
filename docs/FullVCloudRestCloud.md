# FullVCloudRestCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuProfileTypes** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RemoteAccessInternalIp** | Pointer to **string** |  | [optional] 
**RemoteAccessPort** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**VsphereApiVersion** | Pointer to **string** |  | [optional] 
**VsphereHost** | Pointer to **string** |  | [optional] 
**VspherePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewFullVCloudRestCloud

`func NewFullVCloudRestCloud() *FullVCloudRestCloud`

NewFullVCloudRestCloud instantiates a new FullVCloudRestCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullVCloudRestCloudWithDefaults

`func NewFullVCloudRestCloudWithDefaults() *FullVCloudRestCloud`

NewFullVCloudRestCloudWithDefaults instantiates a new FullVCloudRestCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuProfileTypes

`func (o *FullVCloudRestCloud) GetGpuProfileTypes() []string`

GetGpuProfileTypes returns the GpuProfileTypes field if non-nil, zero value otherwise.

### GetGpuProfileTypesOk

`func (o *FullVCloudRestCloud) GetGpuProfileTypesOk() (*[]string, bool)`

GetGpuProfileTypesOk returns a tuple with the GpuProfileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuProfileTypes

`func (o *FullVCloudRestCloud) SetGpuProfileTypes(v []string)`

SetGpuProfileTypes sets GpuProfileTypes field to given value.

### HasGpuProfileTypes

`func (o *FullVCloudRestCloud) HasGpuProfileTypes() bool`

HasGpuProfileTypes returns a boolean if a field has been set.

### GetHostname

`func (o *FullVCloudRestCloud) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FullVCloudRestCloud) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FullVCloudRestCloud) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FullVCloudRestCloud) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *FullVCloudRestCloud) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FullVCloudRestCloud) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FullVCloudRestCloud) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FullVCloudRestCloud) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *FullVCloudRestCloud) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FullVCloudRestCloud) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FullVCloudRestCloud) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FullVCloudRestCloud) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteAccessInternalIp

`func (o *FullVCloudRestCloud) GetRemoteAccessInternalIp() string`

GetRemoteAccessInternalIp returns the RemoteAccessInternalIp field if non-nil, zero value otherwise.

### GetRemoteAccessInternalIpOk

`func (o *FullVCloudRestCloud) GetRemoteAccessInternalIpOk() (*string, bool)`

GetRemoteAccessInternalIpOk returns a tuple with the RemoteAccessInternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessInternalIp

`func (o *FullVCloudRestCloud) SetRemoteAccessInternalIp(v string)`

SetRemoteAccessInternalIp sets RemoteAccessInternalIp field to given value.

### HasRemoteAccessInternalIp

`func (o *FullVCloudRestCloud) HasRemoteAccessInternalIp() bool`

HasRemoteAccessInternalIp returns a boolean if a field has been set.

### GetRemoteAccessPort

`func (o *FullVCloudRestCloud) GetRemoteAccessPort() int32`

GetRemoteAccessPort returns the RemoteAccessPort field if non-nil, zero value otherwise.

### GetRemoteAccessPortOk

`func (o *FullVCloudRestCloud) GetRemoteAccessPortOk() (*int32, bool)`

GetRemoteAccessPortOk returns a tuple with the RemoteAccessPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessPort

`func (o *FullVCloudRestCloud) SetRemoteAccessPort(v int32)`

SetRemoteAccessPort sets RemoteAccessPort field to given value.

### HasRemoteAccessPort

`func (o *FullVCloudRestCloud) HasRemoteAccessPort() bool`

HasRemoteAccessPort returns a boolean if a field has been set.

### GetUsername

`func (o *FullVCloudRestCloud) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FullVCloudRestCloud) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FullVCloudRestCloud) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FullVCloudRestCloud) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVsphereApiVersion

`func (o *FullVCloudRestCloud) GetVsphereApiVersion() string`

GetVsphereApiVersion returns the VsphereApiVersion field if non-nil, zero value otherwise.

### GetVsphereApiVersionOk

`func (o *FullVCloudRestCloud) GetVsphereApiVersionOk() (*string, bool)`

GetVsphereApiVersionOk returns a tuple with the VsphereApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereApiVersion

`func (o *FullVCloudRestCloud) SetVsphereApiVersion(v string)`

SetVsphereApiVersion sets VsphereApiVersion field to given value.

### HasVsphereApiVersion

`func (o *FullVCloudRestCloud) HasVsphereApiVersion() bool`

HasVsphereApiVersion returns a boolean if a field has been set.

### GetVsphereHost

`func (o *FullVCloudRestCloud) GetVsphereHost() string`

GetVsphereHost returns the VsphereHost field if non-nil, zero value otherwise.

### GetVsphereHostOk

`func (o *FullVCloudRestCloud) GetVsphereHostOk() (*string, bool)`

GetVsphereHostOk returns a tuple with the VsphereHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHost

`func (o *FullVCloudRestCloud) SetVsphereHost(v string)`

SetVsphereHost sets VsphereHost field to given value.

### HasVsphereHost

`func (o *FullVCloudRestCloud) HasVsphereHost() bool`

HasVsphereHost returns a boolean if a field has been set.

### GetVspherePort

`func (o *FullVCloudRestCloud) GetVspherePort() int32`

GetVspherePort returns the VspherePort field if non-nil, zero value otherwise.

### GetVspherePortOk

`func (o *FullVCloudRestCloud) GetVspherePortOk() (*int32, bool)`

GetVspherePortOk returns a tuple with the VspherePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVspherePort

`func (o *FullVCloudRestCloud) SetVspherePort(v int32)`

SetVspherePort sets VspherePort field to given value.

### HasVspherePort

`func (o *FullVCloudRestCloud) HasVspherePort() bool`

HasVspherePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


