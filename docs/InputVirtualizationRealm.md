# InputVirtualizationRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualizationRealmType** | Pointer to **string** |  | [optional] 
**AccessPoint** | Pointer to **string** |  | [optional] 
**AccountId** | **string** |  | 
**Cidr** | **string** |  | 
**DefaultWindowsDomainName** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LocalStorageName** | Pointer to **string** |  | [optional] 
**MaximumNumCpus** | Pointer to **int32** |  | [optional] 
**MaximumNumGpus** | Pointer to **int32** |  | [optional] 
**MaximumRamInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumStorageInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumVirtualMachines** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Password** | **string** |  | 
**PowerOnDelayBase** | Pointer to **int64** |  | [optional] 
**PowerOnInitialDelayBase** | Pointer to **int32** |  | [optional] 
**PowerOnMaximumDelay** | Pointer to **int32** |  | [optional] 
**PowerOnMinimumDelay** | Pointer to **int32** |  | [optional] 
**RemoteAccessConfig** | Pointer to [**RemoteAccessConfig**](RemoteAccessConfig.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 
**ZoneCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputVirtualizationRealm

`func NewInputVirtualizationRealm(accountId string, cidr string, description string, name string, password string, username string, ) *InputVirtualizationRealm`

NewInputVirtualizationRealm instantiates a new InputVirtualizationRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVirtualizationRealmWithDefaults

`func NewInputVirtualizationRealmWithDefaults() *InputVirtualizationRealm`

NewInputVirtualizationRealmWithDefaults instantiates a new InputVirtualizationRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualizationRealmType

`func (o *InputVirtualizationRealm) GetVirtualizationRealmType() string`

GetVirtualizationRealmType returns the VirtualizationRealmType field if non-nil, zero value otherwise.

### GetVirtualizationRealmTypeOk

`func (o *InputVirtualizationRealm) GetVirtualizationRealmTypeOk() (*string, bool)`

GetVirtualizationRealmTypeOk returns a tuple with the VirtualizationRealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmType

`func (o *InputVirtualizationRealm) SetVirtualizationRealmType(v string)`

SetVirtualizationRealmType sets VirtualizationRealmType field to given value.

### HasVirtualizationRealmType

`func (o *InputVirtualizationRealm) HasVirtualizationRealmType() bool`

HasVirtualizationRealmType returns a boolean if a field has been set.

### GetAccessPoint

`func (o *InputVirtualizationRealm) GetAccessPoint() string`

GetAccessPoint returns the AccessPoint field if non-nil, zero value otherwise.

### GetAccessPointOk

`func (o *InputVirtualizationRealm) GetAccessPointOk() (*string, bool)`

GetAccessPointOk returns a tuple with the AccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoint

`func (o *InputVirtualizationRealm) SetAccessPoint(v string)`

SetAccessPoint sets AccessPoint field to given value.

### HasAccessPoint

`func (o *InputVirtualizationRealm) HasAccessPoint() bool`

HasAccessPoint returns a boolean if a field has been set.

### GetAccountId

`func (o *InputVirtualizationRealm) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InputVirtualizationRealm) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InputVirtualizationRealm) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCidr

`func (o *InputVirtualizationRealm) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InputVirtualizationRealm) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InputVirtualizationRealm) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDefaultWindowsDomainName

`func (o *InputVirtualizationRealm) GetDefaultWindowsDomainName() string`

GetDefaultWindowsDomainName returns the DefaultWindowsDomainName field if non-nil, zero value otherwise.

### GetDefaultWindowsDomainNameOk

`func (o *InputVirtualizationRealm) GetDefaultWindowsDomainNameOk() (*string, bool)`

GetDefaultWindowsDomainNameOk returns a tuple with the DefaultWindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWindowsDomainName

`func (o *InputVirtualizationRealm) SetDefaultWindowsDomainName(v string)`

SetDefaultWindowsDomainName sets DefaultWindowsDomainName field to given value.

### HasDefaultWindowsDomainName

`func (o *InputVirtualizationRealm) HasDefaultWindowsDomainName() bool`

HasDefaultWindowsDomainName returns a boolean if a field has been set.

### GetDescription

`func (o *InputVirtualizationRealm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputVirtualizationRealm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputVirtualizationRealm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *InputVirtualizationRealm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputVirtualizationRealm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputVirtualizationRealm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputVirtualizationRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalStorageName

`func (o *InputVirtualizationRealm) GetLocalStorageName() string`

GetLocalStorageName returns the LocalStorageName field if non-nil, zero value otherwise.

### GetLocalStorageNameOk

`func (o *InputVirtualizationRealm) GetLocalStorageNameOk() (*string, bool)`

GetLocalStorageNameOk returns a tuple with the LocalStorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageName

`func (o *InputVirtualizationRealm) SetLocalStorageName(v string)`

SetLocalStorageName sets LocalStorageName field to given value.

### HasLocalStorageName

`func (o *InputVirtualizationRealm) HasLocalStorageName() bool`

HasLocalStorageName returns a boolean if a field has been set.

### GetMaximumNumCpus

`func (o *InputVirtualizationRealm) GetMaximumNumCpus() int32`

GetMaximumNumCpus returns the MaximumNumCpus field if non-nil, zero value otherwise.

### GetMaximumNumCpusOk

`func (o *InputVirtualizationRealm) GetMaximumNumCpusOk() (*int32, bool)`

GetMaximumNumCpusOk returns a tuple with the MaximumNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumCpus

`func (o *InputVirtualizationRealm) SetMaximumNumCpus(v int32)`

SetMaximumNumCpus sets MaximumNumCpus field to given value.

### HasMaximumNumCpus

`func (o *InputVirtualizationRealm) HasMaximumNumCpus() bool`

HasMaximumNumCpus returns a boolean if a field has been set.

### GetMaximumNumGpus

`func (o *InputVirtualizationRealm) GetMaximumNumGpus() int32`

GetMaximumNumGpus returns the MaximumNumGpus field if non-nil, zero value otherwise.

### GetMaximumNumGpusOk

`func (o *InputVirtualizationRealm) GetMaximumNumGpusOk() (*int32, bool)`

GetMaximumNumGpusOk returns a tuple with the MaximumNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumGpus

`func (o *InputVirtualizationRealm) SetMaximumNumGpus(v int32)`

SetMaximumNumGpus sets MaximumNumGpus field to given value.

### HasMaximumNumGpus

`func (o *InputVirtualizationRealm) HasMaximumNumGpus() bool`

HasMaximumNumGpus returns a boolean if a field has been set.

### GetMaximumRamInMegabytes

`func (o *InputVirtualizationRealm) GetMaximumRamInMegabytes() int32`

GetMaximumRamInMegabytes returns the MaximumRamInMegabytes field if non-nil, zero value otherwise.

### GetMaximumRamInMegabytesOk

`func (o *InputVirtualizationRealm) GetMaximumRamInMegabytesOk() (*int32, bool)`

GetMaximumRamInMegabytesOk returns a tuple with the MaximumRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRamInMegabytes

`func (o *InputVirtualizationRealm) SetMaximumRamInMegabytes(v int32)`

SetMaximumRamInMegabytes sets MaximumRamInMegabytes field to given value.

### HasMaximumRamInMegabytes

`func (o *InputVirtualizationRealm) HasMaximumRamInMegabytes() bool`

HasMaximumRamInMegabytes returns a boolean if a field has been set.

### GetMaximumStorageInMegabytes

`func (o *InputVirtualizationRealm) GetMaximumStorageInMegabytes() int32`

GetMaximumStorageInMegabytes returns the MaximumStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaximumStorageInMegabytesOk

`func (o *InputVirtualizationRealm) GetMaximumStorageInMegabytesOk() (*int32, bool)`

GetMaximumStorageInMegabytesOk returns a tuple with the MaximumStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumStorageInMegabytes

`func (o *InputVirtualizationRealm) SetMaximumStorageInMegabytes(v int32)`

SetMaximumStorageInMegabytes sets MaximumStorageInMegabytes field to given value.

### HasMaximumStorageInMegabytes

`func (o *InputVirtualizationRealm) HasMaximumStorageInMegabytes() bool`

HasMaximumStorageInMegabytes returns a boolean if a field has been set.

### GetMaximumVirtualMachines

`func (o *InputVirtualizationRealm) GetMaximumVirtualMachines() int32`

GetMaximumVirtualMachines returns the MaximumVirtualMachines field if non-nil, zero value otherwise.

### GetMaximumVirtualMachinesOk

`func (o *InputVirtualizationRealm) GetMaximumVirtualMachinesOk() (*int32, bool)`

GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualMachines

`func (o *InputVirtualizationRealm) SetMaximumVirtualMachines(v int32)`

SetMaximumVirtualMachines sets MaximumVirtualMachines field to given value.

### HasMaximumVirtualMachines

`func (o *InputVirtualizationRealm) HasMaximumVirtualMachines() bool`

HasMaximumVirtualMachines returns a boolean if a field has been set.

### GetName

`func (o *InputVirtualizationRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputVirtualizationRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputVirtualizationRealm) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *InputVirtualizationRealm) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputVirtualizationRealm) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputVirtualizationRealm) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPowerOnDelayBase

`func (o *InputVirtualizationRealm) GetPowerOnDelayBase() int64`

GetPowerOnDelayBase returns the PowerOnDelayBase field if non-nil, zero value otherwise.

### GetPowerOnDelayBaseOk

`func (o *InputVirtualizationRealm) GetPowerOnDelayBaseOk() (*int64, bool)`

GetPowerOnDelayBaseOk returns a tuple with the PowerOnDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayBase

`func (o *InputVirtualizationRealm) SetPowerOnDelayBase(v int64)`

SetPowerOnDelayBase sets PowerOnDelayBase field to given value.

### HasPowerOnDelayBase

`func (o *InputVirtualizationRealm) HasPowerOnDelayBase() bool`

HasPowerOnDelayBase returns a boolean if a field has been set.

### GetPowerOnInitialDelayBase

`func (o *InputVirtualizationRealm) GetPowerOnInitialDelayBase() int32`

GetPowerOnInitialDelayBase returns the PowerOnInitialDelayBase field if non-nil, zero value otherwise.

### GetPowerOnInitialDelayBaseOk

`func (o *InputVirtualizationRealm) GetPowerOnInitialDelayBaseOk() (*int32, bool)`

GetPowerOnInitialDelayBaseOk returns a tuple with the PowerOnInitialDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnInitialDelayBase

`func (o *InputVirtualizationRealm) SetPowerOnInitialDelayBase(v int32)`

SetPowerOnInitialDelayBase sets PowerOnInitialDelayBase field to given value.

### HasPowerOnInitialDelayBase

`func (o *InputVirtualizationRealm) HasPowerOnInitialDelayBase() bool`

HasPowerOnInitialDelayBase returns a boolean if a field has been set.

### GetPowerOnMaximumDelay

`func (o *InputVirtualizationRealm) GetPowerOnMaximumDelay() int32`

GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field if non-nil, zero value otherwise.

### GetPowerOnMaximumDelayOk

`func (o *InputVirtualizationRealm) GetPowerOnMaximumDelayOk() (*int32, bool)`

GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMaximumDelay

`func (o *InputVirtualizationRealm) SetPowerOnMaximumDelay(v int32)`

SetPowerOnMaximumDelay sets PowerOnMaximumDelay field to given value.

### HasPowerOnMaximumDelay

`func (o *InputVirtualizationRealm) HasPowerOnMaximumDelay() bool`

HasPowerOnMaximumDelay returns a boolean if a field has been set.

### GetPowerOnMinimumDelay

`func (o *InputVirtualizationRealm) GetPowerOnMinimumDelay() int32`

GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field if non-nil, zero value otherwise.

### GetPowerOnMinimumDelayOk

`func (o *InputVirtualizationRealm) GetPowerOnMinimumDelayOk() (*int32, bool)`

GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMinimumDelay

`func (o *InputVirtualizationRealm) SetPowerOnMinimumDelay(v int32)`

SetPowerOnMinimumDelay sets PowerOnMinimumDelay field to given value.

### HasPowerOnMinimumDelay

`func (o *InputVirtualizationRealm) HasPowerOnMinimumDelay() bool`

HasPowerOnMinimumDelay returns a boolean if a field has been set.

### GetRemoteAccessConfig

`func (o *InputVirtualizationRealm) GetRemoteAccessConfig() RemoteAccessConfig`

GetRemoteAccessConfig returns the RemoteAccessConfig field if non-nil, zero value otherwise.

### GetRemoteAccessConfigOk

`func (o *InputVirtualizationRealm) GetRemoteAccessConfigOk() (*RemoteAccessConfig, bool)`

GetRemoteAccessConfigOk returns a tuple with the RemoteAccessConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessConfig

`func (o *InputVirtualizationRealm) SetRemoteAccessConfig(v RemoteAccessConfig)`

SetRemoteAccessConfig sets RemoteAccessConfig field to given value.

### HasRemoteAccessConfig

`func (o *InputVirtualizationRealm) HasRemoteAccessConfig() bool`

HasRemoteAccessConfig returns a boolean if a field has been set.

### GetState

`func (o *InputVirtualizationRealm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InputVirtualizationRealm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InputVirtualizationRealm) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InputVirtualizationRealm) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *InputVirtualizationRealm) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InputVirtualizationRealm) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InputVirtualizationRealm) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetZoneCount

`func (o *InputVirtualizationRealm) GetZoneCount() int32`

GetZoneCount returns the ZoneCount field if non-nil, zero value otherwise.

### GetZoneCountOk

`func (o *InputVirtualizationRealm) GetZoneCountOk() (*int32, bool)`

GetZoneCountOk returns a tuple with the ZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCount

`func (o *InputVirtualizationRealm) SetZoneCount(v int32)`

SetZoneCount sets ZoneCount field to given value.

### HasZoneCount

`func (o *InputVirtualizationRealm) HasZoneCount() bool`

HasZoneCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


