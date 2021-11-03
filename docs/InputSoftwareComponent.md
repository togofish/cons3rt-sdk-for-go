# InputSoftwareComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**InputAsset**](InputAsset.md) |  | 
**RebootDelay** | Pointer to **int32** |  | [optional] 
**RebootRequired** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputSoftwareComponent

`func NewInputSoftwareComponent(asset InputAsset, ) *InputSoftwareComponent`

NewInputSoftwareComponent instantiates a new InputSoftwareComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSoftwareComponentWithDefaults

`func NewInputSoftwareComponentWithDefaults() *InputSoftwareComponent`

NewInputSoftwareComponentWithDefaults instantiates a new InputSoftwareComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *InputSoftwareComponent) GetAsset() InputAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *InputSoftwareComponent) GetAssetOk() (*InputAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *InputSoftwareComponent) SetAsset(v InputAsset)`

SetAsset sets Asset field to given value.


### GetRebootDelay

`func (o *InputSoftwareComponent) GetRebootDelay() int32`

GetRebootDelay returns the RebootDelay field if non-nil, zero value otherwise.

### GetRebootDelayOk

`func (o *InputSoftwareComponent) GetRebootDelayOk() (*int32, bool)`

GetRebootDelayOk returns a tuple with the RebootDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDelay

`func (o *InputSoftwareComponent) SetRebootDelay(v int32)`

SetRebootDelay sets RebootDelay field to given value.

### HasRebootDelay

`func (o *InputSoftwareComponent) HasRebootDelay() bool`

HasRebootDelay returns a boolean if a field has been set.

### GetRebootRequired

`func (o *InputSoftwareComponent) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *InputSoftwareComponent) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *InputSoftwareComponent) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.

### HasRebootRequired

`func (o *InputSoftwareComponent) HasRebootRequired() bool`

HasRebootRequired returns a boolean if a field has been set.

### GetId

`func (o *InputSoftwareComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputSoftwareComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputSoftwareComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputSoftwareComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *InputSoftwareComponent) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *InputSoftwareComponent) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *InputSoftwareComponent) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *InputSoftwareComponent) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


