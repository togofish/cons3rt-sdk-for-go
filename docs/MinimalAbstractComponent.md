# MinimalAbstractComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**MinimalAsset**](MinimalAsset.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewMinimalAbstractComponent

`func NewMinimalAbstractComponent(asset MinimalAsset, subtype string, ) *MinimalAbstractComponent`

NewMinimalAbstractComponent instantiates a new MinimalAbstractComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalAbstractComponentWithDefaults

`func NewMinimalAbstractComponentWithDefaults() *MinimalAbstractComponent`

NewMinimalAbstractComponentWithDefaults instantiates a new MinimalAbstractComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *MinimalAbstractComponent) GetAsset() MinimalAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *MinimalAbstractComponent) GetAssetOk() (*MinimalAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *MinimalAbstractComponent) SetAsset(v MinimalAsset)`

SetAsset sets Asset field to given value.


### GetId

`func (o *MinimalAbstractComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalAbstractComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalAbstractComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalAbstractComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *MinimalAbstractComponent) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *MinimalAbstractComponent) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *MinimalAbstractComponent) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *MinimalAbstractComponent) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.

### GetSubtype

`func (o *MinimalAbstractComponent) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MinimalAbstractComponent) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MinimalAbstractComponent) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


