# AbstractRegisterCloudSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualizationRealmType** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AccessPoint** | Pointer to **string** |  | [optional] 
**ActiveAfterRegistration** | Pointer to **bool** |  | [optional] 
**AdditionalNetworkNames** | Pointer to **[]string** |  | [optional] 
**Cons3rtNetworkName** | **string** |  | 
**MaximumNumCpus** | Pointer to **int32** |  | [optional] 
**MaximumNumGpus** | Pointer to **int32** |  | [optional] 
**MaximumRamInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumStorageInMegabytes** | Pointer to **int32** |  | [optional] 
**MaximumVirtualMachines** | Pointer to **int32** |  | [optional] 
**PowerOnMinimumDelay** | Pointer to **int32** |  | [optional] 
**PowerOnMaximumDelay** | Pointer to **int32** |  | [optional] 
**Password** | **string** |  | 
**PowerOnInitialDelayBase** | Pointer to **int32** |  | [optional] 
**PrimaryNetworkName** | **string** |  | 
**RemoteAccessConfig** | Pointer to [**RemoteAccessConfig**](RemoteAccessConfig.md) |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewAbstractRegisterCloudSpaceRequest

`func NewAbstractRegisterCloudSpaceRequest(virtualizationRealmType string, name string, cons3rtNetworkName string, password string, primaryNetworkName string, username string, ) *AbstractRegisterCloudSpaceRequest`

NewAbstractRegisterCloudSpaceRequest instantiates a new AbstractRegisterCloudSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractRegisterCloudSpaceRequestWithDefaults

`func NewAbstractRegisterCloudSpaceRequestWithDefaults() *AbstractRegisterCloudSpaceRequest`

NewAbstractRegisterCloudSpaceRequestWithDefaults instantiates a new AbstractRegisterCloudSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualizationRealmType

`func (o *AbstractRegisterCloudSpaceRequest) GetVirtualizationRealmType() string`

GetVirtualizationRealmType returns the VirtualizationRealmType field if non-nil, zero value otherwise.

### GetVirtualizationRealmTypeOk

`func (o *AbstractRegisterCloudSpaceRequest) GetVirtualizationRealmTypeOk() (*string, bool)`

GetVirtualizationRealmTypeOk returns a tuple with the VirtualizationRealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmType

`func (o *AbstractRegisterCloudSpaceRequest) SetVirtualizationRealmType(v string)`

SetVirtualizationRealmType sets VirtualizationRealmType field to given value.


### GetName

`func (o *AbstractRegisterCloudSpaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractRegisterCloudSpaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractRegisterCloudSpaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AbstractRegisterCloudSpaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AbstractRegisterCloudSpaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AbstractRegisterCloudSpaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AbstractRegisterCloudSpaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessPoint

`func (o *AbstractRegisterCloudSpaceRequest) GetAccessPoint() string`

GetAccessPoint returns the AccessPoint field if non-nil, zero value otherwise.

### GetAccessPointOk

`func (o *AbstractRegisterCloudSpaceRequest) GetAccessPointOk() (*string, bool)`

GetAccessPointOk returns a tuple with the AccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoint

`func (o *AbstractRegisterCloudSpaceRequest) SetAccessPoint(v string)`

SetAccessPoint sets AccessPoint field to given value.

### HasAccessPoint

`func (o *AbstractRegisterCloudSpaceRequest) HasAccessPoint() bool`

HasAccessPoint returns a boolean if a field has been set.

### GetActiveAfterRegistration

`func (o *AbstractRegisterCloudSpaceRequest) GetActiveAfterRegistration() bool`

GetActiveAfterRegistration returns the ActiveAfterRegistration field if non-nil, zero value otherwise.

### GetActiveAfterRegistrationOk

`func (o *AbstractRegisterCloudSpaceRequest) GetActiveAfterRegistrationOk() (*bool, bool)`

GetActiveAfterRegistrationOk returns a tuple with the ActiveAfterRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAfterRegistration

`func (o *AbstractRegisterCloudSpaceRequest) SetActiveAfterRegistration(v bool)`

SetActiveAfterRegistration sets ActiveAfterRegistration field to given value.

### HasActiveAfterRegistration

`func (o *AbstractRegisterCloudSpaceRequest) HasActiveAfterRegistration() bool`

HasActiveAfterRegistration returns a boolean if a field has been set.

### GetAdditionalNetworkNames

`func (o *AbstractRegisterCloudSpaceRequest) GetAdditionalNetworkNames() []string`

GetAdditionalNetworkNames returns the AdditionalNetworkNames field if non-nil, zero value otherwise.

### GetAdditionalNetworkNamesOk

`func (o *AbstractRegisterCloudSpaceRequest) GetAdditionalNetworkNamesOk() (*[]string, bool)`

GetAdditionalNetworkNamesOk returns a tuple with the AdditionalNetworkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNetworkNames

`func (o *AbstractRegisterCloudSpaceRequest) SetAdditionalNetworkNames(v []string)`

SetAdditionalNetworkNames sets AdditionalNetworkNames field to given value.

### HasAdditionalNetworkNames

`func (o *AbstractRegisterCloudSpaceRequest) HasAdditionalNetworkNames() bool`

HasAdditionalNetworkNames returns a boolean if a field has been set.

### GetCons3rtNetworkName

`func (o *AbstractRegisterCloudSpaceRequest) GetCons3rtNetworkName() string`

GetCons3rtNetworkName returns the Cons3rtNetworkName field if non-nil, zero value otherwise.

### GetCons3rtNetworkNameOk

`func (o *AbstractRegisterCloudSpaceRequest) GetCons3rtNetworkNameOk() (*string, bool)`

GetCons3rtNetworkNameOk returns a tuple with the Cons3rtNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtNetworkName

`func (o *AbstractRegisterCloudSpaceRequest) SetCons3rtNetworkName(v string)`

SetCons3rtNetworkName sets Cons3rtNetworkName field to given value.


### GetMaximumNumCpus

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumNumCpus() int32`

GetMaximumNumCpus returns the MaximumNumCpus field if non-nil, zero value otherwise.

### GetMaximumNumCpusOk

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumNumCpusOk() (*int32, bool)`

GetMaximumNumCpusOk returns a tuple with the MaximumNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumCpus

`func (o *AbstractRegisterCloudSpaceRequest) SetMaximumNumCpus(v int32)`

SetMaximumNumCpus sets MaximumNumCpus field to given value.

### HasMaximumNumCpus

`func (o *AbstractRegisterCloudSpaceRequest) HasMaximumNumCpus() bool`

HasMaximumNumCpus returns a boolean if a field has been set.

### GetMaximumNumGpus

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumNumGpus() int32`

GetMaximumNumGpus returns the MaximumNumGpus field if non-nil, zero value otherwise.

### GetMaximumNumGpusOk

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumNumGpusOk() (*int32, bool)`

GetMaximumNumGpusOk returns a tuple with the MaximumNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumGpus

`func (o *AbstractRegisterCloudSpaceRequest) SetMaximumNumGpus(v int32)`

SetMaximumNumGpus sets MaximumNumGpus field to given value.

### HasMaximumNumGpus

`func (o *AbstractRegisterCloudSpaceRequest) HasMaximumNumGpus() bool`

HasMaximumNumGpus returns a boolean if a field has been set.

### GetMaximumRamInMegabytes

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumRamInMegabytes() int32`

GetMaximumRamInMegabytes returns the MaximumRamInMegabytes field if non-nil, zero value otherwise.

### GetMaximumRamInMegabytesOk

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumRamInMegabytesOk() (*int32, bool)`

GetMaximumRamInMegabytesOk returns a tuple with the MaximumRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRamInMegabytes

`func (o *AbstractRegisterCloudSpaceRequest) SetMaximumRamInMegabytes(v int32)`

SetMaximumRamInMegabytes sets MaximumRamInMegabytes field to given value.

### HasMaximumRamInMegabytes

`func (o *AbstractRegisterCloudSpaceRequest) HasMaximumRamInMegabytes() bool`

HasMaximumRamInMegabytes returns a boolean if a field has been set.

### GetMaximumStorageInMegabytes

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumStorageInMegabytes() int32`

GetMaximumStorageInMegabytes returns the MaximumStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaximumStorageInMegabytesOk

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumStorageInMegabytesOk() (*int32, bool)`

GetMaximumStorageInMegabytesOk returns a tuple with the MaximumStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumStorageInMegabytes

`func (o *AbstractRegisterCloudSpaceRequest) SetMaximumStorageInMegabytes(v int32)`

SetMaximumStorageInMegabytes sets MaximumStorageInMegabytes field to given value.

### HasMaximumStorageInMegabytes

`func (o *AbstractRegisterCloudSpaceRequest) HasMaximumStorageInMegabytes() bool`

HasMaximumStorageInMegabytes returns a boolean if a field has been set.

### GetMaximumVirtualMachines

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumVirtualMachines() int32`

GetMaximumVirtualMachines returns the MaximumVirtualMachines field if non-nil, zero value otherwise.

### GetMaximumVirtualMachinesOk

`func (o *AbstractRegisterCloudSpaceRequest) GetMaximumVirtualMachinesOk() (*int32, bool)`

GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualMachines

`func (o *AbstractRegisterCloudSpaceRequest) SetMaximumVirtualMachines(v int32)`

SetMaximumVirtualMachines sets MaximumVirtualMachines field to given value.

### HasMaximumVirtualMachines

`func (o *AbstractRegisterCloudSpaceRequest) HasMaximumVirtualMachines() bool`

HasMaximumVirtualMachines returns a boolean if a field has been set.

### GetPowerOnMinimumDelay

`func (o *AbstractRegisterCloudSpaceRequest) GetPowerOnMinimumDelay() int32`

GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field if non-nil, zero value otherwise.

### GetPowerOnMinimumDelayOk

`func (o *AbstractRegisterCloudSpaceRequest) GetPowerOnMinimumDelayOk() (*int32, bool)`

GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMinimumDelay

`func (o *AbstractRegisterCloudSpaceRequest) SetPowerOnMinimumDelay(v int32)`

SetPowerOnMinimumDelay sets PowerOnMinimumDelay field to given value.

### HasPowerOnMinimumDelay

`func (o *AbstractRegisterCloudSpaceRequest) HasPowerOnMinimumDelay() bool`

HasPowerOnMinimumDelay returns a boolean if a field has been set.

### GetPowerOnMaximumDelay

`func (o *AbstractRegisterCloudSpaceRequest) GetPowerOnMaximumDelay() int32`

GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field if non-nil, zero value otherwise.

### GetPowerOnMaximumDelayOk

`func (o *AbstractRegisterCloudSpaceRequest) GetPowerOnMaximumDelayOk() (*int32, bool)`

GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMaximumDelay

`func (o *AbstractRegisterCloudSpaceRequest) SetPowerOnMaximumDelay(v int32)`

SetPowerOnMaximumDelay sets PowerOnMaximumDelay field to given value.

### HasPowerOnMaximumDelay

`func (o *AbstractRegisterCloudSpaceRequest) HasPowerOnMaximumDelay() bool`

HasPowerOnMaximumDelay returns a boolean if a field has been set.

### GetPassword

`func (o *AbstractRegisterCloudSpaceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AbstractRegisterCloudSpaceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AbstractRegisterCloudSpaceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPowerOnInitialDelayBase

`func (o *AbstractRegisterCloudSpaceRequest) GetPowerOnInitialDelayBase() int32`

GetPowerOnInitialDelayBase returns the PowerOnInitialDelayBase field if non-nil, zero value otherwise.

### GetPowerOnInitialDelayBaseOk

`func (o *AbstractRegisterCloudSpaceRequest) GetPowerOnInitialDelayBaseOk() (*int32, bool)`

GetPowerOnInitialDelayBaseOk returns a tuple with the PowerOnInitialDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnInitialDelayBase

`func (o *AbstractRegisterCloudSpaceRequest) SetPowerOnInitialDelayBase(v int32)`

SetPowerOnInitialDelayBase sets PowerOnInitialDelayBase field to given value.

### HasPowerOnInitialDelayBase

`func (o *AbstractRegisterCloudSpaceRequest) HasPowerOnInitialDelayBase() bool`

HasPowerOnInitialDelayBase returns a boolean if a field has been set.

### GetPrimaryNetworkName

`func (o *AbstractRegisterCloudSpaceRequest) GetPrimaryNetworkName() string`

GetPrimaryNetworkName returns the PrimaryNetworkName field if non-nil, zero value otherwise.

### GetPrimaryNetworkNameOk

`func (o *AbstractRegisterCloudSpaceRequest) GetPrimaryNetworkNameOk() (*string, bool)`

GetPrimaryNetworkNameOk returns a tuple with the PrimaryNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNetworkName

`func (o *AbstractRegisterCloudSpaceRequest) SetPrimaryNetworkName(v string)`

SetPrimaryNetworkName sets PrimaryNetworkName field to given value.


### GetRemoteAccessConfig

`func (o *AbstractRegisterCloudSpaceRequest) GetRemoteAccessConfig() RemoteAccessConfig`

GetRemoteAccessConfig returns the RemoteAccessConfig field if non-nil, zero value otherwise.

### GetRemoteAccessConfigOk

`func (o *AbstractRegisterCloudSpaceRequest) GetRemoteAccessConfigOk() (*RemoteAccessConfig, bool)`

GetRemoteAccessConfigOk returns a tuple with the RemoteAccessConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessConfig

`func (o *AbstractRegisterCloudSpaceRequest) SetRemoteAccessConfig(v RemoteAccessConfig)`

SetRemoteAccessConfig sets RemoteAccessConfig field to given value.

### HasRemoteAccessConfig

`func (o *AbstractRegisterCloudSpaceRequest) HasRemoteAccessConfig() bool`

HasRemoteAccessConfig returns a boolean if a field has been set.

### GetUsername

`func (o *AbstractRegisterCloudSpaceRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AbstractRegisterCloudSpaceRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AbstractRegisterCloudSpaceRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


