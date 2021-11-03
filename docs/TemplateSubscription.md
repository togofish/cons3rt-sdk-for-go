# TemplateSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MaxNumCpus** | Pointer to **int32** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubscribingVirtualizationRealm** | Pointer to [**VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**TemplateRegistration** | Pointer to [**TemplateRegistration**](TemplateRegistration.md) |  | [optional] 
**TemplateUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateSubscription

`func NewTemplateSubscription() *TemplateSubscription`

NewTemplateSubscription instantiates a new TemplateSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSubscriptionWithDefaults

`func NewTemplateSubscriptionWithDefaults() *TemplateSubscription`

NewTemplateSubscriptionWithDefaults instantiates a new TemplateSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateSubscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *TemplateSubscription) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *TemplateSubscription) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *TemplateSubscription) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *TemplateSubscription) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *TemplateSubscription) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *TemplateSubscription) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *TemplateSubscription) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *TemplateSubscription) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.

### GetState

`func (o *TemplateSubscription) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateSubscription) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateSubscription) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TemplateSubscription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscribingVirtualizationRealm

`func (o *TemplateSubscription) GetSubscribingVirtualizationRealm() VirtualizationRealm`

GetSubscribingVirtualizationRealm returns the SubscribingVirtualizationRealm field if non-nil, zero value otherwise.

### GetSubscribingVirtualizationRealmOk

`func (o *TemplateSubscription) GetSubscribingVirtualizationRealmOk() (*VirtualizationRealm, bool)`

GetSubscribingVirtualizationRealmOk returns a tuple with the SubscribingVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribingVirtualizationRealm

`func (o *TemplateSubscription) SetSubscribingVirtualizationRealm(v VirtualizationRealm)`

SetSubscribingVirtualizationRealm sets SubscribingVirtualizationRealm field to given value.

### HasSubscribingVirtualizationRealm

`func (o *TemplateSubscription) HasSubscribingVirtualizationRealm() bool`

HasSubscribingVirtualizationRealm returns a boolean if a field has been set.

### GetTemplateRegistration

`func (o *TemplateSubscription) GetTemplateRegistration() TemplateRegistration`

GetTemplateRegistration returns the TemplateRegistration field if non-nil, zero value otherwise.

### GetTemplateRegistrationOk

`func (o *TemplateSubscription) GetTemplateRegistrationOk() (*TemplateRegistration, bool)`

GetTemplateRegistrationOk returns a tuple with the TemplateRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegistration

`func (o *TemplateSubscription) SetTemplateRegistration(v TemplateRegistration)`

SetTemplateRegistration sets TemplateRegistration field to given value.

### HasTemplateRegistration

`func (o *TemplateSubscription) HasTemplateRegistration() bool`

HasTemplateRegistration returns a boolean if a field has been set.

### GetTemplateUuid

`func (o *TemplateSubscription) GetTemplateUuid() string`

GetTemplateUuid returns the TemplateUuid field if non-nil, zero value otherwise.

### GetTemplateUuidOk

`func (o *TemplateSubscription) GetTemplateUuidOk() (*string, bool)`

GetTemplateUuidOk returns a tuple with the TemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUuid

`func (o *TemplateSubscription) SetTemplateUuid(v string)`

SetTemplateUuid sets TemplateUuid field to given value.

### HasTemplateUuid

`func (o *TemplateSubscription) HasTemplateUuid() bool`

HasTemplateUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


