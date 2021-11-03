# InputAbstractComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**InputAsset**](InputAsset.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewInputAbstractComponent

`func NewInputAbstractComponent(asset InputAsset, subtype string, ) *InputAbstractComponent`

NewInputAbstractComponent instantiates a new InputAbstractComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputAbstractComponentWithDefaults

`func NewInputAbstractComponentWithDefaults() *InputAbstractComponent`

NewInputAbstractComponentWithDefaults instantiates a new InputAbstractComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *InputAbstractComponent) GetAsset() InputAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *InputAbstractComponent) GetAssetOk() (*InputAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *InputAbstractComponent) SetAsset(v InputAsset)`

SetAsset sets Asset field to given value.


### GetId

`func (o *InputAbstractComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputAbstractComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputAbstractComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputAbstractComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *InputAbstractComponent) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *InputAbstractComponent) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *InputAbstractComponent) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *InputAbstractComponent) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.

### GetSubtype

`func (o *InputAbstractComponent) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputAbstractComponent) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputAbstractComponent) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


