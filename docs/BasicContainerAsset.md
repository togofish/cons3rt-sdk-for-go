# BasicContainerAsset

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
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBasicContainerAsset

`func NewBasicContainerAsset() *BasicContainerAsset`

NewBasicContainerAsset instantiates a new BasicContainerAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicContainerAssetWithDefaults

`func NewBasicContainerAssetWithDefaults() *BasicContainerAsset`

NewBasicContainerAssetWithDefaults instantiates a new BasicContainerAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicContainerAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicContainerAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicContainerAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasicContainerAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BasicContainerAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicContainerAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicContainerAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicContainerAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BasicContainerAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicContainerAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicContainerAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BasicContainerAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *BasicContainerAsset) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *BasicContainerAsset) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *BasicContainerAsset) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *BasicContainerAsset) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *BasicContainerAsset) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BasicContainerAsset) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BasicContainerAsset) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BasicContainerAsset) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *BasicContainerAsset) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BasicContainerAsset) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BasicContainerAsset) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BasicContainerAsset) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreator

`func (o *BasicContainerAsset) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BasicContainerAsset) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BasicContainerAsset) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BasicContainerAsset) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOwningProject

`func (o *BasicContainerAsset) GetOwningProject() MinimalProject`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *BasicContainerAsset) GetOwningProjectOk() (*MinimalProject, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *BasicContainerAsset) SetOwningProject(v MinimalProject)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *BasicContainerAsset) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetType

`func (o *BasicContainerAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicContainerAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicContainerAsset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BasicContainerAsset) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


