# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RuleOrder** | **int32** |  | 
**Protocol** | **string** |  | 
**RuleAction** | **string** |  | 
**RuleDestination** | **string** |  | 
**RuleDestPort** | Pointer to **string** |  | [optional] 
**RuleEnabled** | **bool** |  | 
**RuleSource** | **string** |  | 
**RuleSourcePort** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallRule

`func NewFirewallRule(ruleOrder int32, protocol string, ruleAction string, ruleDestination string, ruleEnabled bool, ruleSource string, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRuleOrder

`func (o *FirewallRule) GetRuleOrder() int32`

GetRuleOrder returns the RuleOrder field if non-nil, zero value otherwise.

### GetRuleOrderOk

`func (o *FirewallRule) GetRuleOrderOk() (*int32, bool)`

GetRuleOrderOk returns a tuple with the RuleOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOrder

`func (o *FirewallRule) SetRuleOrder(v int32)`

SetRuleOrder sets RuleOrder field to given value.


### GetProtocol

`func (o *FirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRuleAction

`func (o *FirewallRule) GetRuleAction() string`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *FirewallRule) GetRuleActionOk() (*string, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *FirewallRule) SetRuleAction(v string)`

SetRuleAction sets RuleAction field to given value.


### GetRuleDestination

`func (o *FirewallRule) GetRuleDestination() string`

GetRuleDestination returns the RuleDestination field if non-nil, zero value otherwise.

### GetRuleDestinationOk

`func (o *FirewallRule) GetRuleDestinationOk() (*string, bool)`

GetRuleDestinationOk returns a tuple with the RuleDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDestination

`func (o *FirewallRule) SetRuleDestination(v string)`

SetRuleDestination sets RuleDestination field to given value.


### GetRuleDestPort

`func (o *FirewallRule) GetRuleDestPort() string`

GetRuleDestPort returns the RuleDestPort field if non-nil, zero value otherwise.

### GetRuleDestPortOk

`func (o *FirewallRule) GetRuleDestPortOk() (*string, bool)`

GetRuleDestPortOk returns a tuple with the RuleDestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDestPort

`func (o *FirewallRule) SetRuleDestPort(v string)`

SetRuleDestPort sets RuleDestPort field to given value.

### HasRuleDestPort

`func (o *FirewallRule) HasRuleDestPort() bool`

HasRuleDestPort returns a boolean if a field has been set.

### GetRuleEnabled

`func (o *FirewallRule) GetRuleEnabled() bool`

GetRuleEnabled returns the RuleEnabled field if non-nil, zero value otherwise.

### GetRuleEnabledOk

`func (o *FirewallRule) GetRuleEnabledOk() (*bool, bool)`

GetRuleEnabledOk returns a tuple with the RuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleEnabled

`func (o *FirewallRule) SetRuleEnabled(v bool)`

SetRuleEnabled sets RuleEnabled field to given value.


### GetRuleSource

`func (o *FirewallRule) GetRuleSource() string`

GetRuleSource returns the RuleSource field if non-nil, zero value otherwise.

### GetRuleSourceOk

`func (o *FirewallRule) GetRuleSourceOk() (*string, bool)`

GetRuleSourceOk returns a tuple with the RuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSource

`func (o *FirewallRule) SetRuleSource(v string)`

SetRuleSource sets RuleSource field to given value.


### GetRuleSourcePort

`func (o *FirewallRule) GetRuleSourcePort() string`

GetRuleSourcePort returns the RuleSourcePort field if non-nil, zero value otherwise.

### GetRuleSourcePortOk

`func (o *FirewallRule) GetRuleSourcePortOk() (*string, bool)`

GetRuleSourcePortOk returns a tuple with the RuleSourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSourcePort

`func (o *FirewallRule) SetRuleSourcePort(v string)`

SetRuleSourcePort sets RuleSourcePort field to given value.

### HasRuleSourcePort

`func (o *FirewallRule) HasRuleSourcePort() bool`

HasRuleSourcePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


