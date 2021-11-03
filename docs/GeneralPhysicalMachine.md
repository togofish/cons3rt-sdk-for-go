# GeneralPhysicalMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDisks** | Pointer to [**[]FullDisk**](FullDisk.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**CpuCount** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**Bits** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 

## Methods

### NewGeneralPhysicalMachine

`func NewGeneralPhysicalMachine() *GeneralPhysicalMachine`

NewGeneralPhysicalMachine instantiates a new GeneralPhysicalMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralPhysicalMachineWithDefaults

`func NewGeneralPhysicalMachineWithDefaults() *GeneralPhysicalMachine`

NewGeneralPhysicalMachineWithDefaults instantiates a new GeneralPhysicalMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDisks

`func (o *GeneralPhysicalMachine) GetBaseDisks() []FullDisk`

GetBaseDisks returns the BaseDisks field if non-nil, zero value otherwise.

### GetBaseDisksOk

`func (o *GeneralPhysicalMachine) GetBaseDisksOk() (*[]FullDisk, bool)`

GetBaseDisksOk returns a tuple with the BaseDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDisks

`func (o *GeneralPhysicalMachine) SetBaseDisks(v []FullDisk)`

SetBaseDisks sets BaseDisks field to given value.

### HasBaseDisks

`func (o *GeneralPhysicalMachine) HasBaseDisks() bool`

HasBaseDisks returns a boolean if a field has been set.

### GetId

`func (o *GeneralPhysicalMachine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneralPhysicalMachine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneralPhysicalMachine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GeneralPhysicalMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GeneralPhysicalMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeneralPhysicalMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeneralPhysicalMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GeneralPhysicalMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GeneralPhysicalMachine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeneralPhysicalMachine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeneralPhysicalMachine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeneralPhysicalMachine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *GeneralPhysicalMachine) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *GeneralPhysicalMachine) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *GeneralPhysicalMachine) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *GeneralPhysicalMachine) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *GeneralPhysicalMachine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeneralPhysicalMachine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeneralPhysicalMachine) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GeneralPhysicalMachine) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *GeneralPhysicalMachine) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GeneralPhysicalMachine) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GeneralPhysicalMachine) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GeneralPhysicalMachine) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetHostName

`func (o *GeneralPhysicalMachine) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GeneralPhysicalMachine) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GeneralPhysicalMachine) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GeneralPhysicalMachine) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpAddress

`func (o *GeneralPhysicalMachine) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GeneralPhysicalMachine) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GeneralPhysicalMachine) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GeneralPhysicalMachine) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *GeneralPhysicalMachine) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GeneralPhysicalMachine) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GeneralPhysicalMachine) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GeneralPhysicalMachine) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetCpuCount

`func (o *GeneralPhysicalMachine) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *GeneralPhysicalMachine) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *GeneralPhysicalMachine) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *GeneralPhysicalMachine) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetRam

`func (o *GeneralPhysicalMachine) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *GeneralPhysicalMachine) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *GeneralPhysicalMachine) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *GeneralPhysicalMachine) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetArchitecture

`func (o *GeneralPhysicalMachine) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *GeneralPhysicalMachine) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *GeneralPhysicalMachine) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *GeneralPhysicalMachine) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBits

`func (o *GeneralPhysicalMachine) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GeneralPhysicalMachine) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GeneralPhysicalMachine) SetBits(v string)`

SetBits sets Bits field to given value.

### HasBits

`func (o *GeneralPhysicalMachine) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *GeneralPhysicalMachine) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *GeneralPhysicalMachine) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *GeneralPhysicalMachine) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *GeneralPhysicalMachine) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


