# InputUnregisterTemplateObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveSubscriptions** | **bool** |  | 
**Client** | Pointer to [**InputVirtualizationRealmClient**](InputVirtualizationRealmClient.md) |  | [optional] 

## Methods

### NewInputUnregisterTemplateObject

`func NewInputUnregisterTemplateObject(removeSubscriptions bool, ) *InputUnregisterTemplateObject`

NewInputUnregisterTemplateObject instantiates a new InputUnregisterTemplateObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUnregisterTemplateObjectWithDefaults

`func NewInputUnregisterTemplateObjectWithDefaults() *InputUnregisterTemplateObject`

NewInputUnregisterTemplateObjectWithDefaults instantiates a new InputUnregisterTemplateObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveSubscriptions

`func (o *InputUnregisterTemplateObject) GetRemoveSubscriptions() bool`

GetRemoveSubscriptions returns the RemoveSubscriptions field if non-nil, zero value otherwise.

### GetRemoveSubscriptionsOk

`func (o *InputUnregisterTemplateObject) GetRemoveSubscriptionsOk() (*bool, bool)`

GetRemoveSubscriptionsOk returns a tuple with the RemoveSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSubscriptions

`func (o *InputUnregisterTemplateObject) SetRemoveSubscriptions(v bool)`

SetRemoveSubscriptions sets RemoveSubscriptions field to given value.


### GetClient

`func (o *InputUnregisterTemplateObject) GetClient() InputVirtualizationRealmClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *InputUnregisterTemplateObject) GetClientOk() (*InputVirtualizationRealmClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *InputUnregisterTemplateObject) SetClient(v InputVirtualizationRealmClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *InputUnregisterTemplateObject) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


