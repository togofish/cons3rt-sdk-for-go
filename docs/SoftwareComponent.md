# SoftwareComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**Asset**](Asset.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 
**RebootDelay** | Pointer to **int32** |  | [optional] 
**RebootRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewSoftwareComponent

`func NewSoftwareComponent(asset Asset, ) *SoftwareComponent`

NewSoftwareComponent instantiates a new SoftwareComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareComponentWithDefaults

`func NewSoftwareComponentWithDefaults() *SoftwareComponent`

NewSoftwareComponentWithDefaults instantiates a new SoftwareComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *SoftwareComponent) GetAsset() Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *SoftwareComponent) GetAssetOk() (*Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *SoftwareComponent) SetAsset(v Asset)`

SetAsset sets Asset field to given value.


### GetId

`func (o *SoftwareComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftwareComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftwareComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SoftwareComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *SoftwareComponent) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *SoftwareComponent) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *SoftwareComponent) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *SoftwareComponent) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.

### GetRebootDelay

`func (o *SoftwareComponent) GetRebootDelay() int32`

GetRebootDelay returns the RebootDelay field if non-nil, zero value otherwise.

### GetRebootDelayOk

`func (o *SoftwareComponent) GetRebootDelayOk() (*int32, bool)`

GetRebootDelayOk returns a tuple with the RebootDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDelay

`func (o *SoftwareComponent) SetRebootDelay(v int32)`

SetRebootDelay sets RebootDelay field to given value.

### HasRebootDelay

`func (o *SoftwareComponent) HasRebootDelay() bool`

HasRebootDelay returns a boolean if a field has been set.

### GetRebootRequired

`func (o *SoftwareComponent) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *SoftwareComponent) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *SoftwareComponent) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.

### HasRebootRequired

`func (o *SoftwareComponent) HasRebootRequired() bool`

HasRebootRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


