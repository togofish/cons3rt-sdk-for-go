# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disruptive** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Parent** | Pointer to [**Category**](Category.md) |  | [optional] 
**Subcategories** | Pointer to [**[]Category**](Category.md) |  | [optional] 

## Methods

### NewCategory

`func NewCategory(id int32, name string, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisruptive

`func (o *Category) GetDisruptive() bool`

GetDisruptive returns the Disruptive field if non-nil, zero value otherwise.

### GetDisruptiveOk

`func (o *Category) GetDisruptiveOk() (*bool, bool)`

GetDisruptiveOk returns a tuple with the Disruptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptive

`func (o *Category) SetDisruptive(v bool)`

SetDisruptive sets Disruptive field to given value.

### HasDisruptive

`func (o *Category) HasDisruptive() bool`

HasDisruptive returns a boolean if a field has been set.

### GetHidden

`func (o *Category) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Category) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Category) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Category) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *Category) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *Category) GetParent() Category`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Category) GetParentOk() (*Category, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Category) SetParent(v Category)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Category) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSubcategories

`func (o *Category) GetSubcategories() []Category`

GetSubcategories returns the Subcategories field if non-nil, zero value otherwise.

### GetSubcategoriesOk

`func (o *Category) GetSubcategoriesOk() (*[]Category, bool)`

GetSubcategoriesOk returns a tuple with the Subcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategories

`func (o *Category) SetSubcategories(v []Category)`

SetSubcategories sets Subcategories field to given value.

### HasSubcategories

`func (o *Category) HasSubcategories() bool`

HasSubcategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


