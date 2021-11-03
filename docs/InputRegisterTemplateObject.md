# InputRegisterTemplateObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**InputVirtualizationRealmClient**](InputVirtualizationRealmClient.md) |  | [optional] 
**TemplateData** | [**InputCons3rtTemplateData**](InputCons3rtTemplateData.md) |  | 

## Methods

### NewInputRegisterTemplateObject

`func NewInputRegisterTemplateObject(templateData InputCons3rtTemplateData, ) *InputRegisterTemplateObject`

NewInputRegisterTemplateObject instantiates a new InputRegisterTemplateObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRegisterTemplateObjectWithDefaults

`func NewInputRegisterTemplateObjectWithDefaults() *InputRegisterTemplateObject`

NewInputRegisterTemplateObjectWithDefaults instantiates a new InputRegisterTemplateObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *InputRegisterTemplateObject) GetClient() InputVirtualizationRealmClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *InputRegisterTemplateObject) GetClientOk() (*InputVirtualizationRealmClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *InputRegisterTemplateObject) SetClient(v InputVirtualizationRealmClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *InputRegisterTemplateObject) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetTemplateData

`func (o *InputRegisterTemplateObject) GetTemplateData() InputCons3rtTemplateData`

GetTemplateData returns the TemplateData field if non-nil, zero value otherwise.

### GetTemplateDataOk

`func (o *InputRegisterTemplateObject) GetTemplateDataOk() (*InputCons3rtTemplateData, bool)`

GetTemplateDataOk returns a tuple with the TemplateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateData

`func (o *InputRegisterTemplateObject) SetTemplateData(v InputCons3rtTemplateData)`

SetTemplateData sets TemplateData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


