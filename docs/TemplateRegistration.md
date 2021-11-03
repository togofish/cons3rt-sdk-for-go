# TemplateRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**RegisteringVirtualizationRealm** | Pointer to [**VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]TemplateSubscription**](TemplateSubscription.md) |  | [optional] 
**TemplateData** | Pointer to [**Cons3rtTemplateData**](Cons3rtTemplateData.md) |  | [optional] 
**TemplateUuid** | Pointer to **string** |  | [optional] 
**VirtRealmsSharedTo** | Pointer to [**[]VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 

## Methods

### NewTemplateRegistration

`func NewTemplateRegistration() *TemplateRegistration`

NewTemplateRegistration instantiates a new TemplateRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateRegistrationWithDefaults

`func NewTemplateRegistrationWithDefaults() *TemplateRegistration`

NewTemplateRegistrationWithDefaults instantiates a new TemplateRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateRegistration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateRegistration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateRegistration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOffline

`func (o *TemplateRegistration) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *TemplateRegistration) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *TemplateRegistration) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *TemplateRegistration) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetRegisteringVirtualizationRealm

`func (o *TemplateRegistration) GetRegisteringVirtualizationRealm() VirtualizationRealm`

GetRegisteringVirtualizationRealm returns the RegisteringVirtualizationRealm field if non-nil, zero value otherwise.

### GetRegisteringVirtualizationRealmOk

`func (o *TemplateRegistration) GetRegisteringVirtualizationRealmOk() (*VirtualizationRealm, bool)`

GetRegisteringVirtualizationRealmOk returns a tuple with the RegisteringVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteringVirtualizationRealm

`func (o *TemplateRegistration) SetRegisteringVirtualizationRealm(v VirtualizationRealm)`

SetRegisteringVirtualizationRealm sets RegisteringVirtualizationRealm field to given value.

### HasRegisteringVirtualizationRealm

`func (o *TemplateRegistration) HasRegisteringVirtualizationRealm() bool`

HasRegisteringVirtualizationRealm returns a boolean if a field has been set.

### GetSubscriptions

`func (o *TemplateRegistration) GetSubscriptions() []TemplateSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *TemplateRegistration) GetSubscriptionsOk() (*[]TemplateSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *TemplateRegistration) SetSubscriptions(v []TemplateSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *TemplateRegistration) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTemplateData

`func (o *TemplateRegistration) GetTemplateData() Cons3rtTemplateData`

GetTemplateData returns the TemplateData field if non-nil, zero value otherwise.

### GetTemplateDataOk

`func (o *TemplateRegistration) GetTemplateDataOk() (*Cons3rtTemplateData, bool)`

GetTemplateDataOk returns a tuple with the TemplateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateData

`func (o *TemplateRegistration) SetTemplateData(v Cons3rtTemplateData)`

SetTemplateData sets TemplateData field to given value.

### HasTemplateData

`func (o *TemplateRegistration) HasTemplateData() bool`

HasTemplateData returns a boolean if a field has been set.

### GetTemplateUuid

`func (o *TemplateRegistration) GetTemplateUuid() string`

GetTemplateUuid returns the TemplateUuid field if non-nil, zero value otherwise.

### GetTemplateUuidOk

`func (o *TemplateRegistration) GetTemplateUuidOk() (*string, bool)`

GetTemplateUuidOk returns a tuple with the TemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUuid

`func (o *TemplateRegistration) SetTemplateUuid(v string)`

SetTemplateUuid sets TemplateUuid field to given value.

### HasTemplateUuid

`func (o *TemplateRegistration) HasTemplateUuid() bool`

HasTemplateUuid returns a boolean if a field has been set.

### GetVirtRealmsSharedTo

`func (o *TemplateRegistration) GetVirtRealmsSharedTo() []VirtualizationRealm`

GetVirtRealmsSharedTo returns the VirtRealmsSharedTo field if non-nil, zero value otherwise.

### GetVirtRealmsSharedToOk

`func (o *TemplateRegistration) GetVirtRealmsSharedToOk() (*[]VirtualizationRealm, bool)`

GetVirtRealmsSharedToOk returns a tuple with the VirtRealmsSharedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmsSharedTo

`func (o *TemplateRegistration) SetVirtRealmsSharedTo(v []VirtualizationRealm)`

SetVirtRealmsSharedTo sets VirtRealmsSharedTo field to given value.

### HasVirtRealmsSharedTo

`func (o *TemplateRegistration) HasVirtRealmsSharedTo() bool`

HasVirtRealmsSharedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


