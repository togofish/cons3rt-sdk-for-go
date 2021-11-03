# InputCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disruptive** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Parent** | Pointer to [**InputParentCategory**](InputParentCategory.md) |  | [optional] 

## Methods

### NewInputCategory

`func NewInputCategory(name string, ) *InputCategory`

NewInputCategory instantiates a new InputCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCategoryWithDefaults

`func NewInputCategoryWithDefaults() *InputCategory`

NewInputCategoryWithDefaults instantiates a new InputCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisruptive

`func (o *InputCategory) GetDisruptive() bool`

GetDisruptive returns the Disruptive field if non-nil, zero value otherwise.

### GetDisruptiveOk

`func (o *InputCategory) GetDisruptiveOk() (*bool, bool)`

GetDisruptiveOk returns a tuple with the Disruptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptive

`func (o *InputCategory) SetDisruptive(v bool)`

SetDisruptive sets Disruptive field to given value.

### HasDisruptive

`func (o *InputCategory) HasDisruptive() bool`

HasDisruptive returns a boolean if a field has been set.

### GetHidden

`func (o *InputCategory) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *InputCategory) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *InputCategory) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *InputCategory) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *InputCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InputCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCategory) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *InputCategory) GetParent() InputParentCategory`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *InputCategory) GetParentOk() (*InputParentCategory, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *InputCategory) SetParent(v InputParentCategory)`

SetParent sets Parent field to given value.

### HasParent

`func (o *InputCategory) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


