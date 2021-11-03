# InputVRAdminVirtualizationRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualizationRealmType** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**AdditionalNetworks** | Pointer to [**[]InputVirtualizationRealmNetwork**](InputVirtualizationRealmNetwork.md) |  | [optional] 
**Cidr** | **string** |  | 
**Cons3rtNetwork** | Pointer to [**InputVirtualizationRealmNetwork**](InputVirtualizationRealmNetwork.md) |  | [optional] 
**DefaultWindowsDomainName** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Name** | **string** |  | 
**PowerOnDelayBase** | Pointer to **int64** |  | [optional] 
**PowerOnInitialDelayBase** | Pointer to **int32** |  | [optional] 
**PowerOnMaximumDelay** | Pointer to **int32** |  | [optional] 
**PowerOnMinimumDelay** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputVRAdminVirtualizationRealm

`func NewInputVRAdminVirtualizationRealm(virtualizationRealmType string, cidr string, description string, name string, ) *InputVRAdminVirtualizationRealm`

NewInputVRAdminVirtualizationRealm instantiates a new InputVRAdminVirtualizationRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVRAdminVirtualizationRealmWithDefaults

`func NewInputVRAdminVirtualizationRealmWithDefaults() *InputVRAdminVirtualizationRealm`

NewInputVRAdminVirtualizationRealmWithDefaults instantiates a new InputVRAdminVirtualizationRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualizationRealmType

`func (o *InputVRAdminVirtualizationRealm) GetVirtualizationRealmType() string`

GetVirtualizationRealmType returns the VirtualizationRealmType field if non-nil, zero value otherwise.

### GetVirtualizationRealmTypeOk

`func (o *InputVRAdminVirtualizationRealm) GetVirtualizationRealmTypeOk() (*string, bool)`

GetVirtualizationRealmTypeOk returns a tuple with the VirtualizationRealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmType

`func (o *InputVRAdminVirtualizationRealm) SetVirtualizationRealmType(v string)`

SetVirtualizationRealmType sets VirtualizationRealmType field to given value.


### GetId

`func (o *InputVRAdminVirtualizationRealm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputVRAdminVirtualizationRealm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputVRAdminVirtualizationRealm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputVRAdminVirtualizationRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdditionalNetworks

`func (o *InputVRAdminVirtualizationRealm) GetAdditionalNetworks() []InputVirtualizationRealmNetwork`

GetAdditionalNetworks returns the AdditionalNetworks field if non-nil, zero value otherwise.

### GetAdditionalNetworksOk

`func (o *InputVRAdminVirtualizationRealm) GetAdditionalNetworksOk() (*[]InputVirtualizationRealmNetwork, bool)`

GetAdditionalNetworksOk returns a tuple with the AdditionalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNetworks

`func (o *InputVRAdminVirtualizationRealm) SetAdditionalNetworks(v []InputVirtualizationRealmNetwork)`

SetAdditionalNetworks sets AdditionalNetworks field to given value.

### HasAdditionalNetworks

`func (o *InputVRAdminVirtualizationRealm) HasAdditionalNetworks() bool`

HasAdditionalNetworks returns a boolean if a field has been set.

### GetCidr

`func (o *InputVRAdminVirtualizationRealm) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InputVRAdminVirtualizationRealm) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InputVRAdminVirtualizationRealm) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCons3rtNetwork

`func (o *InputVRAdminVirtualizationRealm) GetCons3rtNetwork() InputVirtualizationRealmNetwork`

GetCons3rtNetwork returns the Cons3rtNetwork field if non-nil, zero value otherwise.

### GetCons3rtNetworkOk

`func (o *InputVRAdminVirtualizationRealm) GetCons3rtNetworkOk() (*InputVirtualizationRealmNetwork, bool)`

GetCons3rtNetworkOk returns a tuple with the Cons3rtNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtNetwork

`func (o *InputVRAdminVirtualizationRealm) SetCons3rtNetwork(v InputVirtualizationRealmNetwork)`

SetCons3rtNetwork sets Cons3rtNetwork field to given value.

### HasCons3rtNetwork

`func (o *InputVRAdminVirtualizationRealm) HasCons3rtNetwork() bool`

HasCons3rtNetwork returns a boolean if a field has been set.

### GetDefaultWindowsDomainName

`func (o *InputVRAdminVirtualizationRealm) GetDefaultWindowsDomainName() string`

GetDefaultWindowsDomainName returns the DefaultWindowsDomainName field if non-nil, zero value otherwise.

### GetDefaultWindowsDomainNameOk

`func (o *InputVRAdminVirtualizationRealm) GetDefaultWindowsDomainNameOk() (*string, bool)`

GetDefaultWindowsDomainNameOk returns a tuple with the DefaultWindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWindowsDomainName

`func (o *InputVRAdminVirtualizationRealm) SetDefaultWindowsDomainName(v string)`

SetDefaultWindowsDomainName sets DefaultWindowsDomainName field to given value.

### HasDefaultWindowsDomainName

`func (o *InputVRAdminVirtualizationRealm) HasDefaultWindowsDomainName() bool`

HasDefaultWindowsDomainName returns a boolean if a field has been set.

### GetDescription

`func (o *InputVRAdminVirtualizationRealm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputVRAdminVirtualizationRealm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputVRAdminVirtualizationRealm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *InputVRAdminVirtualizationRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputVRAdminVirtualizationRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputVRAdminVirtualizationRealm) SetName(v string)`

SetName sets Name field to given value.


### GetPowerOnDelayBase

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnDelayBase() int64`

GetPowerOnDelayBase returns the PowerOnDelayBase field if non-nil, zero value otherwise.

### GetPowerOnDelayBaseOk

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnDelayBaseOk() (*int64, bool)`

GetPowerOnDelayBaseOk returns a tuple with the PowerOnDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnDelayBase

`func (o *InputVRAdminVirtualizationRealm) SetPowerOnDelayBase(v int64)`

SetPowerOnDelayBase sets PowerOnDelayBase field to given value.

### HasPowerOnDelayBase

`func (o *InputVRAdminVirtualizationRealm) HasPowerOnDelayBase() bool`

HasPowerOnDelayBase returns a boolean if a field has been set.

### GetPowerOnInitialDelayBase

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnInitialDelayBase() int32`

GetPowerOnInitialDelayBase returns the PowerOnInitialDelayBase field if non-nil, zero value otherwise.

### GetPowerOnInitialDelayBaseOk

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnInitialDelayBaseOk() (*int32, bool)`

GetPowerOnInitialDelayBaseOk returns a tuple with the PowerOnInitialDelayBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnInitialDelayBase

`func (o *InputVRAdminVirtualizationRealm) SetPowerOnInitialDelayBase(v int32)`

SetPowerOnInitialDelayBase sets PowerOnInitialDelayBase field to given value.

### HasPowerOnInitialDelayBase

`func (o *InputVRAdminVirtualizationRealm) HasPowerOnInitialDelayBase() bool`

HasPowerOnInitialDelayBase returns a boolean if a field has been set.

### GetPowerOnMaximumDelay

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnMaximumDelay() int32`

GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field if non-nil, zero value otherwise.

### GetPowerOnMaximumDelayOk

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnMaximumDelayOk() (*int32, bool)`

GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMaximumDelay

`func (o *InputVRAdminVirtualizationRealm) SetPowerOnMaximumDelay(v int32)`

SetPowerOnMaximumDelay sets PowerOnMaximumDelay field to given value.

### HasPowerOnMaximumDelay

`func (o *InputVRAdminVirtualizationRealm) HasPowerOnMaximumDelay() bool`

HasPowerOnMaximumDelay returns a boolean if a field has been set.

### GetPowerOnMinimumDelay

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnMinimumDelay() int32`

GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field if non-nil, zero value otherwise.

### GetPowerOnMinimumDelayOk

`func (o *InputVRAdminVirtualizationRealm) GetPowerOnMinimumDelayOk() (*int32, bool)`

GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMinimumDelay

`func (o *InputVRAdminVirtualizationRealm) SetPowerOnMinimumDelay(v int32)`

SetPowerOnMinimumDelay sets PowerOnMinimumDelay field to given value.

### HasPowerOnMinimumDelay

`func (o *InputVRAdminVirtualizationRealm) HasPowerOnMinimumDelay() bool`

HasPowerOnMinimumDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


