# FullTemplateRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RegisteringVirtualizationRealm** | Pointer to [**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**[]MinimalTemplateSubscription**](MinimalTemplateSubscription.md) |  | [optional] 
**TemplateData** | Pointer to [**FullCons3rtTemplateData**](FullCons3rtTemplateData.md) |  | [optional] 
**TemplateUuid** | Pointer to **string** |  | [optional] 
**VirtRealmsSharedTo** | Pointer to [**[]MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 

## Methods

### NewFullTemplateRegistration

`func NewFullTemplateRegistration() *FullTemplateRegistration`

NewFullTemplateRegistration instantiates a new FullTemplateRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullTemplateRegistrationWithDefaults

`func NewFullTemplateRegistrationWithDefaults() *FullTemplateRegistration`

NewFullTemplateRegistrationWithDefaults instantiates a new FullTemplateRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullTemplateRegistration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullTemplateRegistration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullTemplateRegistration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullTemplateRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegisteringVirtualizationRealm

`func (o *FullTemplateRegistration) GetRegisteringVirtualizationRealm() MinimalVirtualizationRealm`

GetRegisteringVirtualizationRealm returns the RegisteringVirtualizationRealm field if non-nil, zero value otherwise.

### GetRegisteringVirtualizationRealmOk

`func (o *FullTemplateRegistration) GetRegisteringVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool)`

GetRegisteringVirtualizationRealmOk returns a tuple with the RegisteringVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteringVirtualizationRealm

`func (o *FullTemplateRegistration) SetRegisteringVirtualizationRealm(v MinimalVirtualizationRealm)`

SetRegisteringVirtualizationRealm sets RegisteringVirtualizationRealm field to given value.

### HasRegisteringVirtualizationRealm

`func (o *FullTemplateRegistration) HasRegisteringVirtualizationRealm() bool`

HasRegisteringVirtualizationRealm returns a boolean if a field has been set.

### GetState

`func (o *FullTemplateRegistration) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullTemplateRegistration) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullTemplateRegistration) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullTemplateRegistration) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscriptions

`func (o *FullTemplateRegistration) GetSubscriptions() []MinimalTemplateSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FullTemplateRegistration) GetSubscriptionsOk() (*[]MinimalTemplateSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FullTemplateRegistration) SetSubscriptions(v []MinimalTemplateSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *FullTemplateRegistration) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTemplateData

`func (o *FullTemplateRegistration) GetTemplateData() FullCons3rtTemplateData`

GetTemplateData returns the TemplateData field if non-nil, zero value otherwise.

### GetTemplateDataOk

`func (o *FullTemplateRegistration) GetTemplateDataOk() (*FullCons3rtTemplateData, bool)`

GetTemplateDataOk returns a tuple with the TemplateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateData

`func (o *FullTemplateRegistration) SetTemplateData(v FullCons3rtTemplateData)`

SetTemplateData sets TemplateData field to given value.

### HasTemplateData

`func (o *FullTemplateRegistration) HasTemplateData() bool`

HasTemplateData returns a boolean if a field has been set.

### GetTemplateUuid

`func (o *FullTemplateRegistration) GetTemplateUuid() string`

GetTemplateUuid returns the TemplateUuid field if non-nil, zero value otherwise.

### GetTemplateUuidOk

`func (o *FullTemplateRegistration) GetTemplateUuidOk() (*string, bool)`

GetTemplateUuidOk returns a tuple with the TemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUuid

`func (o *FullTemplateRegistration) SetTemplateUuid(v string)`

SetTemplateUuid sets TemplateUuid field to given value.

### HasTemplateUuid

`func (o *FullTemplateRegistration) HasTemplateUuid() bool`

HasTemplateUuid returns a boolean if a field has been set.

### GetVirtRealmsSharedTo

`func (o *FullTemplateRegistration) GetVirtRealmsSharedTo() []MinimalVirtualizationRealm`

GetVirtRealmsSharedTo returns the VirtRealmsSharedTo field if non-nil, zero value otherwise.

### GetVirtRealmsSharedToOk

`func (o *FullTemplateRegistration) GetVirtRealmsSharedToOk() (*[]MinimalVirtualizationRealm, bool)`

GetVirtRealmsSharedToOk returns a tuple with the VirtRealmsSharedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmsSharedTo

`func (o *FullTemplateRegistration) SetVirtRealmsSharedTo(v []MinimalVirtualizationRealm)`

SetVirtRealmsSharedTo sets VirtRealmsSharedTo field to given value.

### HasVirtRealmsSharedTo

`func (o *FullTemplateRegistration) HasVirtRealmsSharedTo() bool`

HasVirtRealmsSharedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


