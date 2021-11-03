# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Isolated** | Pointer to **bool** |  | [optional] 
**BoundaryIpAddress** | Pointer to **string** |  | [optional] 
**Cidr** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DnatRules** | [**[]DnatRule**](DnatRule.md) |  | 
**FirewallRules** | [**[]FirewallRule**](FirewallRule.md) |  | 
**NetworkFunction** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewNetwork

`func NewNetwork(cidr string, dnatRules []DnatRule, firewallRules []FirewallRule, name string, ) *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Network) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Network) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Network) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsolated

`func (o *Network) GetIsolated() bool`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *Network) GetIsolatedOk() (*bool, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *Network) SetIsolated(v bool)`

SetIsolated sets Isolated field to given value.

### HasIsolated

`func (o *Network) HasIsolated() bool`

HasIsolated returns a boolean if a field has been set.

### GetBoundaryIpAddress

`func (o *Network) GetBoundaryIpAddress() string`

GetBoundaryIpAddress returns the BoundaryIpAddress field if non-nil, zero value otherwise.

### GetBoundaryIpAddressOk

`func (o *Network) GetBoundaryIpAddressOk() (*string, bool)`

GetBoundaryIpAddressOk returns a tuple with the BoundaryIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryIpAddress

`func (o *Network) SetBoundaryIpAddress(v string)`

SetBoundaryIpAddress sets BoundaryIpAddress field to given value.

### HasBoundaryIpAddress

`func (o *Network) HasBoundaryIpAddress() bool`

HasBoundaryIpAddress returns a boolean if a field has been set.

### GetCidr

`func (o *Network) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Network) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Network) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *Network) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Network) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Network) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Network) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnatRules

`func (o *Network) GetDnatRules() []DnatRule`

GetDnatRules returns the DnatRules field if non-nil, zero value otherwise.

### GetDnatRulesOk

`func (o *Network) GetDnatRulesOk() (*[]DnatRule, bool)`

GetDnatRulesOk returns a tuple with the DnatRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatRules

`func (o *Network) SetDnatRules(v []DnatRule)`

SetDnatRules sets DnatRules field to given value.


### GetFirewallRules

`func (o *Network) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *Network) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *Network) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetNetworkFunction

`func (o *Network) GetNetworkFunction() string`

GetNetworkFunction returns the NetworkFunction field if non-nil, zero value otherwise.

### GetNetworkFunctionOk

`func (o *Network) GetNetworkFunctionOk() (*string, bool)`

GetNetworkFunctionOk returns a tuple with the NetworkFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunction

`func (o *Network) SetNetworkFunction(v string)`

SetNetworkFunction sets NetworkFunction field to given value.

### HasNetworkFunction

`func (o *Network) HasNetworkFunction() bool`

HasNetworkFunction returns a boolean if a field has been set.

### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


