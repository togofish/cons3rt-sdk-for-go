# FullPhysicalMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**CpuCount** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**Bits** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**RemoteAccessCapable** | Pointer to **bool** |  | [optional] 
**RemoteAccessTemplates** | Pointer to [**[]MinimalRemoteAccessTemplate**](MinimalRemoteAccessTemplate.md) |  | [optional] 
**BaseDisks** | Pointer to [**[]FullDisk**](FullDisk.md) |  | [optional] 

## Methods

### NewFullPhysicalMachine

`func NewFullPhysicalMachine() *FullPhysicalMachine`

NewFullPhysicalMachine instantiates a new FullPhysicalMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullPhysicalMachineWithDefaults

`func NewFullPhysicalMachineWithDefaults() *FullPhysicalMachine`

NewFullPhysicalMachineWithDefaults instantiates a new FullPhysicalMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *FullPhysicalMachine) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *FullPhysicalMachine) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *FullPhysicalMachine) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *FullPhysicalMachine) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpAddress

`func (o *FullPhysicalMachine) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *FullPhysicalMachine) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *FullPhysicalMachine) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *FullPhysicalMachine) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *FullPhysicalMachine) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *FullPhysicalMachine) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *FullPhysicalMachine) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *FullPhysicalMachine) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetCpuCount

`func (o *FullPhysicalMachine) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *FullPhysicalMachine) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *FullPhysicalMachine) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *FullPhysicalMachine) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetRam

`func (o *FullPhysicalMachine) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *FullPhysicalMachine) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *FullPhysicalMachine) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *FullPhysicalMachine) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetArchitecture

`func (o *FullPhysicalMachine) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *FullPhysicalMachine) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *FullPhysicalMachine) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *FullPhysicalMachine) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBits

`func (o *FullPhysicalMachine) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *FullPhysicalMachine) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *FullPhysicalMachine) SetBits(v string)`

SetBits sets Bits field to given value.

### HasBits

`func (o *FullPhysicalMachine) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *FullPhysicalMachine) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *FullPhysicalMachine) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *FullPhysicalMachine) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *FullPhysicalMachine) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetRemoteAccessCapable

`func (o *FullPhysicalMachine) GetRemoteAccessCapable() bool`

GetRemoteAccessCapable returns the RemoteAccessCapable field if non-nil, zero value otherwise.

### GetRemoteAccessCapableOk

`func (o *FullPhysicalMachine) GetRemoteAccessCapableOk() (*bool, bool)`

GetRemoteAccessCapableOk returns a tuple with the RemoteAccessCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessCapable

`func (o *FullPhysicalMachine) SetRemoteAccessCapable(v bool)`

SetRemoteAccessCapable sets RemoteAccessCapable field to given value.

### HasRemoteAccessCapable

`func (o *FullPhysicalMachine) HasRemoteAccessCapable() bool`

HasRemoteAccessCapable returns a boolean if a field has been set.

### GetRemoteAccessTemplates

`func (o *FullPhysicalMachine) GetRemoteAccessTemplates() []MinimalRemoteAccessTemplate`

GetRemoteAccessTemplates returns the RemoteAccessTemplates field if non-nil, zero value otherwise.

### GetRemoteAccessTemplatesOk

`func (o *FullPhysicalMachine) GetRemoteAccessTemplatesOk() (*[]MinimalRemoteAccessTemplate, bool)`

GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessTemplates

`func (o *FullPhysicalMachine) SetRemoteAccessTemplates(v []MinimalRemoteAccessTemplate)`

SetRemoteAccessTemplates sets RemoteAccessTemplates field to given value.

### HasRemoteAccessTemplates

`func (o *FullPhysicalMachine) HasRemoteAccessTemplates() bool`

HasRemoteAccessTemplates returns a boolean if a field has been set.

### GetBaseDisks

`func (o *FullPhysicalMachine) GetBaseDisks() []FullDisk`

GetBaseDisks returns the BaseDisks field if non-nil, zero value otherwise.

### GetBaseDisksOk

`func (o *FullPhysicalMachine) GetBaseDisksOk() (*[]FullDisk, bool)`

GetBaseDisksOk returns a tuple with the BaseDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDisks

`func (o *FullPhysicalMachine) SetBaseDisks(v []FullDisk)`

SetBaseDisks sets BaseDisks field to given value.

### HasBaseDisks

`func (o *FullPhysicalMachine) HasBaseDisks() bool`

HasBaseDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


