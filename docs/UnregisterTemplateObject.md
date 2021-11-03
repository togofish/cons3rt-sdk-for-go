# UnregisterTemplateObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**VirtualizationRealmClient**](VirtualizationRealmClient.md) |  | [optional] 
**RemoveSubscriptions** | **bool** |  | 

## Methods

### NewUnregisterTemplateObject

`func NewUnregisterTemplateObject(removeSubscriptions bool, ) *UnregisterTemplateObject`

NewUnregisterTemplateObject instantiates a new UnregisterTemplateObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisterTemplateObjectWithDefaults

`func NewUnregisterTemplateObjectWithDefaults() *UnregisterTemplateObject`

NewUnregisterTemplateObjectWithDefaults instantiates a new UnregisterTemplateObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *UnregisterTemplateObject) GetClient() VirtualizationRealmClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *UnregisterTemplateObject) GetClientOk() (*VirtualizationRealmClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *UnregisterTemplateObject) SetClient(v VirtualizationRealmClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *UnregisterTemplateObject) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetRemoveSubscriptions

`func (o *UnregisterTemplateObject) GetRemoveSubscriptions() bool`

GetRemoveSubscriptions returns the RemoveSubscriptions field if non-nil, zero value otherwise.

### GetRemoveSubscriptionsOk

`func (o *UnregisterTemplateObject) GetRemoveSubscriptionsOk() (*bool, bool)`

GetRemoveSubscriptionsOk returns a tuple with the RemoveSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSubscriptions

`func (o *UnregisterTemplateObject) SetRemoveSubscriptions(v bool)`

SetRemoveSubscriptions sets RemoveSubscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


