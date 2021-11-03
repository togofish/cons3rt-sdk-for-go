# MinimalSoftwareComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**MinimalAsset**](MinimalAsset.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 
**RebootDelay** | Pointer to **int32** |  | [optional] 
**RebootRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewMinimalSoftwareComponent

`func NewMinimalSoftwareComponent(asset MinimalAsset, ) *MinimalSoftwareComponent`

NewMinimalSoftwareComponent instantiates a new MinimalSoftwareComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalSoftwareComponentWithDefaults

`func NewMinimalSoftwareComponentWithDefaults() *MinimalSoftwareComponent`

NewMinimalSoftwareComponentWithDefaults instantiates a new MinimalSoftwareComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *MinimalSoftwareComponent) GetAsset() MinimalAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *MinimalSoftwareComponent) GetAssetOk() (*MinimalAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *MinimalSoftwareComponent) SetAsset(v MinimalAsset)`

SetAsset sets Asset field to given value.


### GetId

`func (o *MinimalSoftwareComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalSoftwareComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalSoftwareComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalSoftwareComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *MinimalSoftwareComponent) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *MinimalSoftwareComponent) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *MinimalSoftwareComponent) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *MinimalSoftwareComponent) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.

### GetRebootDelay

`func (o *MinimalSoftwareComponent) GetRebootDelay() int32`

GetRebootDelay returns the RebootDelay field if non-nil, zero value otherwise.

### GetRebootDelayOk

`func (o *MinimalSoftwareComponent) GetRebootDelayOk() (*int32, bool)`

GetRebootDelayOk returns a tuple with the RebootDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDelay

`func (o *MinimalSoftwareComponent) SetRebootDelay(v int32)`

SetRebootDelay sets RebootDelay field to given value.

### HasRebootDelay

`func (o *MinimalSoftwareComponent) HasRebootDelay() bool`

HasRebootDelay returns a boolean if a field has been set.

### GetRebootRequired

`func (o *MinimalSoftwareComponent) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *MinimalSoftwareComponent) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *MinimalSoftwareComponent) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.

### HasRebootRequired

`func (o *MinimalSoftwareComponent) HasRebootRequired() bool`

HasRebootRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


