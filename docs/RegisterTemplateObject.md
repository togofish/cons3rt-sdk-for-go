# RegisterTemplateObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**VirtualizationRealmClient**](VirtualizationRealmClient.md) |  | [optional] 
**TemplateData** | [**Cons3rtTemplateData**](Cons3rtTemplateData.md) |  | 

## Methods

### NewRegisterTemplateObject

`func NewRegisterTemplateObject(templateData Cons3rtTemplateData, ) *RegisterTemplateObject`

NewRegisterTemplateObject instantiates a new RegisterTemplateObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterTemplateObjectWithDefaults

`func NewRegisterTemplateObjectWithDefaults() *RegisterTemplateObject`

NewRegisterTemplateObjectWithDefaults instantiates a new RegisterTemplateObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *RegisterTemplateObject) GetClient() VirtualizationRealmClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *RegisterTemplateObject) GetClientOk() (*VirtualizationRealmClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *RegisterTemplateObject) SetClient(v VirtualizationRealmClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *RegisterTemplateObject) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetTemplateData

`func (o *RegisterTemplateObject) GetTemplateData() Cons3rtTemplateData`

GetTemplateData returns the TemplateData field if non-nil, zero value otherwise.

### GetTemplateDataOk

`func (o *RegisterTemplateObject) GetTemplateDataOk() (*Cons3rtTemplateData, bool)`

GetTemplateDataOk returns a tuple with the TemplateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateData

`func (o *RegisterTemplateObject) SetTemplateData(v Cons3rtTemplateData)`

SetTemplateData sets TemplateData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


