# BasicSystemModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 
**OwningProject** | Pointer to [**MinimalProject**](MinimalProject.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewBasicSystemModule

`func NewBasicSystemModule(subtype string, ) *BasicSystemModule`

NewBasicSystemModule instantiates a new BasicSystemModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicSystemModuleWithDefaults

`func NewBasicSystemModuleWithDefaults() *BasicSystemModule`

NewBasicSystemModuleWithDefaults instantiates a new BasicSystemModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicSystemModule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicSystemModule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicSystemModule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasicSystemModule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BasicSystemModule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicSystemModule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicSystemModule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicSystemModule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BasicSystemModule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicSystemModule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicSystemModule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BasicSystemModule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *BasicSystemModule) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *BasicSystemModule) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *BasicSystemModule) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *BasicSystemModule) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *BasicSystemModule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BasicSystemModule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BasicSystemModule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BasicSystemModule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *BasicSystemModule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BasicSystemModule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BasicSystemModule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BasicSystemModule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreator

`func (o *BasicSystemModule) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BasicSystemModule) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BasicSystemModule) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BasicSystemModule) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOwningProject

`func (o *BasicSystemModule) GetOwningProject() MinimalProject`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *BasicSystemModule) GetOwningProjectOk() (*MinimalProject, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *BasicSystemModule) SetOwningProject(v MinimalProject)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *BasicSystemModule) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetSubtype

`func (o *BasicSystemModule) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BasicSystemModule) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BasicSystemModule) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


