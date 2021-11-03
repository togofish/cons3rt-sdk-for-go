# FullCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disruptive** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**MinimalCategory**](MinimalCategory.md) |  | [optional] 
**Subcategories** | Pointer to [**[]MinimalCategory**](MinimalCategory.md) |  | [optional] 

## Methods

### NewFullCategory

`func NewFullCategory() *FullCategory`

NewFullCategory instantiates a new FullCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullCategoryWithDefaults

`func NewFullCategoryWithDefaults() *FullCategory`

NewFullCategoryWithDefaults instantiates a new FullCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisruptive

`func (o *FullCategory) GetDisruptive() bool`

GetDisruptive returns the Disruptive field if non-nil, zero value otherwise.

### GetDisruptiveOk

`func (o *FullCategory) GetDisruptiveOk() (*bool, bool)`

GetDisruptiveOk returns a tuple with the Disruptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptive

`func (o *FullCategory) SetDisruptive(v bool)`

SetDisruptive sets Disruptive field to given value.

### HasDisruptive

`func (o *FullCategory) HasDisruptive() bool`

HasDisruptive returns a boolean if a field has been set.

### GetEditable

`func (o *FullCategory) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *FullCategory) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *FullCategory) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *FullCategory) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetHidden

`func (o *FullCategory) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FullCategory) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FullCategory) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FullCategory) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *FullCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *FullCategory) GetParent() MinimalCategory`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FullCategory) GetParentOk() (*MinimalCategory, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FullCategory) SetParent(v MinimalCategory)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FullCategory) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSubcategories

`func (o *FullCategory) GetSubcategories() []MinimalCategory`

GetSubcategories returns the Subcategories field if non-nil, zero value otherwise.

### GetSubcategoriesOk

`func (o *FullCategory) GetSubcategoriesOk() (*[]MinimalCategory, bool)`

GetSubcategoriesOk returns a tuple with the Subcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategories

`func (o *FullCategory) SetSubcategories(v []MinimalCategory)`

SetSubcategories sets Subcategories field to given value.

### HasSubcategories

`func (o *FullCategory) HasSubcategories() bool`

HasSubcategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


