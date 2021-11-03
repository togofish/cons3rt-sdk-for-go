# BasicSoftwareAssetBundle

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

## Methods

### NewBasicSoftwareAssetBundle

`func NewBasicSoftwareAssetBundle() *BasicSoftwareAssetBundle`

NewBasicSoftwareAssetBundle instantiates a new BasicSoftwareAssetBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicSoftwareAssetBundleWithDefaults

`func NewBasicSoftwareAssetBundleWithDefaults() *BasicSoftwareAssetBundle`

NewBasicSoftwareAssetBundleWithDefaults instantiates a new BasicSoftwareAssetBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicSoftwareAssetBundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicSoftwareAssetBundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicSoftwareAssetBundle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasicSoftwareAssetBundle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BasicSoftwareAssetBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicSoftwareAssetBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicSoftwareAssetBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicSoftwareAssetBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BasicSoftwareAssetBundle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicSoftwareAssetBundle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicSoftwareAssetBundle) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BasicSoftwareAssetBundle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *BasicSoftwareAssetBundle) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *BasicSoftwareAssetBundle) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *BasicSoftwareAssetBundle) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *BasicSoftwareAssetBundle) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *BasicSoftwareAssetBundle) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BasicSoftwareAssetBundle) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BasicSoftwareAssetBundle) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BasicSoftwareAssetBundle) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *BasicSoftwareAssetBundle) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BasicSoftwareAssetBundle) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BasicSoftwareAssetBundle) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BasicSoftwareAssetBundle) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreator

`func (o *BasicSoftwareAssetBundle) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BasicSoftwareAssetBundle) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BasicSoftwareAssetBundle) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BasicSoftwareAssetBundle) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOwningProject

`func (o *BasicSoftwareAssetBundle) GetOwningProject() MinimalProject`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *BasicSoftwareAssetBundle) GetOwningProjectOk() (*MinimalProject, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *BasicSoftwareAssetBundle) SetOwningProject(v MinimalProject)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *BasicSoftwareAssetBundle) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


