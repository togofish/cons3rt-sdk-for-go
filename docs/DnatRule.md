# DnatRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnatEnabled** | **bool** |  | 
**DnatPort** | **string** |  | 
**DnatProtocol** | **string** |  | 
**DnatTargetIp** | **string** |  | 
**DnatTargetPort** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnatRule

`func NewDnatRule(dnatEnabled bool, dnatPort string, dnatProtocol string, dnatTargetIp string, dnatTargetPort string, ) *DnatRule`

NewDnatRule instantiates a new DnatRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnatRuleWithDefaults

`func NewDnatRuleWithDefaults() *DnatRule`

NewDnatRuleWithDefaults instantiates a new DnatRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnatEnabled

`func (o *DnatRule) GetDnatEnabled() bool`

GetDnatEnabled returns the DnatEnabled field if non-nil, zero value otherwise.

### GetDnatEnabledOk

`func (o *DnatRule) GetDnatEnabledOk() (*bool, bool)`

GetDnatEnabledOk returns a tuple with the DnatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatEnabled

`func (o *DnatRule) SetDnatEnabled(v bool)`

SetDnatEnabled sets DnatEnabled field to given value.


### GetDnatPort

`func (o *DnatRule) GetDnatPort() string`

GetDnatPort returns the DnatPort field if non-nil, zero value otherwise.

### GetDnatPortOk

`func (o *DnatRule) GetDnatPortOk() (*string, bool)`

GetDnatPortOk returns a tuple with the DnatPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatPort

`func (o *DnatRule) SetDnatPort(v string)`

SetDnatPort sets DnatPort field to given value.


### GetDnatProtocol

`func (o *DnatRule) GetDnatProtocol() string`

GetDnatProtocol returns the DnatProtocol field if non-nil, zero value otherwise.

### GetDnatProtocolOk

`func (o *DnatRule) GetDnatProtocolOk() (*string, bool)`

GetDnatProtocolOk returns a tuple with the DnatProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatProtocol

`func (o *DnatRule) SetDnatProtocol(v string)`

SetDnatProtocol sets DnatProtocol field to given value.


### GetDnatTargetIp

`func (o *DnatRule) GetDnatTargetIp() string`

GetDnatTargetIp returns the DnatTargetIp field if non-nil, zero value otherwise.

### GetDnatTargetIpOk

`func (o *DnatRule) GetDnatTargetIpOk() (*string, bool)`

GetDnatTargetIpOk returns a tuple with the DnatTargetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatTargetIp

`func (o *DnatRule) SetDnatTargetIp(v string)`

SetDnatTargetIp sets DnatTargetIp field to given value.


### GetDnatTargetPort

`func (o *DnatRule) GetDnatTargetPort() string`

GetDnatTargetPort returns the DnatTargetPort field if non-nil, zero value otherwise.

### GetDnatTargetPortOk

`func (o *DnatRule) GetDnatTargetPortOk() (*string, bool)`

GetDnatTargetPortOk returns a tuple with the DnatTargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnatTargetPort

`func (o *DnatRule) SetDnatTargetPort(v string)`

SetDnatTargetPort sets DnatTargetPort field to given value.


### GetId

`func (o *DnatRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnatRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnatRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DnatRule) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


