# FullTemplateSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MaxNumCpus** | Pointer to **int32** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubscribingVirtualizationRealm** | Pointer to [**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**TemplateRegistration** | Pointer to [**FullTemplateRegistrationForSubscription**](FullTemplateRegistrationForSubscription.md) |  | [optional] 
**TemplateUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewFullTemplateSubscription

`func NewFullTemplateSubscription() *FullTemplateSubscription`

NewFullTemplateSubscription instantiates a new FullTemplateSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullTemplateSubscriptionWithDefaults

`func NewFullTemplateSubscriptionWithDefaults() *FullTemplateSubscription`

NewFullTemplateSubscriptionWithDefaults instantiates a new FullTemplateSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullTemplateSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullTemplateSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullTemplateSubscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullTemplateSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *FullTemplateSubscription) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *FullTemplateSubscription) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *FullTemplateSubscription) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *FullTemplateSubscription) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *FullTemplateSubscription) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *FullTemplateSubscription) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *FullTemplateSubscription) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *FullTemplateSubscription) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetState

`func (o *FullTemplateSubscription) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullTemplateSubscription) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullTemplateSubscription) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullTemplateSubscription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscribingVirtualizationRealm

`func (o *FullTemplateSubscription) GetSubscribingVirtualizationRealm() MinimalVirtualizationRealm`

GetSubscribingVirtualizationRealm returns the SubscribingVirtualizationRealm field if non-nil, zero value otherwise.

### GetSubscribingVirtualizationRealmOk

`func (o *FullTemplateSubscription) GetSubscribingVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool)`

GetSubscribingVirtualizationRealmOk returns a tuple with the SubscribingVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribingVirtualizationRealm

`func (o *FullTemplateSubscription) SetSubscribingVirtualizationRealm(v MinimalVirtualizationRealm)`

SetSubscribingVirtualizationRealm sets SubscribingVirtualizationRealm field to given value.

### HasSubscribingVirtualizationRealm

`func (o *FullTemplateSubscription) HasSubscribingVirtualizationRealm() bool`

HasSubscribingVirtualizationRealm returns a boolean if a field has been set.

### GetTemplateRegistration

`func (o *FullTemplateSubscription) GetTemplateRegistration() FullTemplateRegistrationForSubscription`

GetTemplateRegistration returns the TemplateRegistration field if non-nil, zero value otherwise.

### GetTemplateRegistrationOk

`func (o *FullTemplateSubscription) GetTemplateRegistrationOk() (*FullTemplateRegistrationForSubscription, bool)`

GetTemplateRegistrationOk returns a tuple with the TemplateRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegistration

`func (o *FullTemplateSubscription) SetTemplateRegistration(v FullTemplateRegistrationForSubscription)`

SetTemplateRegistration sets TemplateRegistration field to given value.

### HasTemplateRegistration

`func (o *FullTemplateSubscription) HasTemplateRegistration() bool`

HasTemplateRegistration returns a boolean if a field has been set.

### GetTemplateUuid

`func (o *FullTemplateSubscription) GetTemplateUuid() string`

GetTemplateUuid returns the TemplateUuid field if non-nil, zero value otherwise.

### GetTemplateUuidOk

`func (o *FullTemplateSubscription) GetTemplateUuidOk() (*string, bool)`

GetTemplateUuidOk returns a tuple with the TemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUuid

`func (o *FullTemplateSubscription) SetTemplateUuid(v string)`

SetTemplateUuid sets TemplateUuid field to given value.

### HasTemplateUuid

`func (o *FullTemplateSubscription) HasTemplateUuid() bool`

HasTemplateUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


