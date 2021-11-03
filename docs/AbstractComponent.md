# AbstractComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**Asset**](Asset.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewAbstractComponent

`func NewAbstractComponent(asset Asset, subtype string, ) *AbstractComponent`

NewAbstractComponent instantiates a new AbstractComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractComponentWithDefaults

`func NewAbstractComponentWithDefaults() *AbstractComponent`

NewAbstractComponentWithDefaults instantiates a new AbstractComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *AbstractComponent) GetAsset() Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AbstractComponent) GetAssetOk() (*Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AbstractComponent) SetAsset(v Asset)`

SetAsset sets Asset field to given value.


### GetId

`func (o *AbstractComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *AbstractComponent) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *AbstractComponent) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *AbstractComponent) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *AbstractComponent) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.

### GetSubtype

`func (o *AbstractComponent) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AbstractComponent) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AbstractComponent) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


