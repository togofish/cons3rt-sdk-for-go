# FullAsset

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
**TrustedProjects** | Pointer to [**[]MinimalProject**](MinimalProject.md) |  | [optional] 
**DependentAssetCount** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**FullMetadata**](FullMetadata.md) |  | [optional] 
**ImpactLevel** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]MinimalCategory**](MinimalCategory.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewFullAsset

`func NewFullAsset(subtype string, ) *FullAsset`

NewFullAsset instantiates a new FullAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullAssetWithDefaults

`func NewFullAssetWithDefaults() *FullAsset`

NewFullAssetWithDefaults instantiates a new FullAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FullAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *FullAsset) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *FullAsset) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *FullAsset) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *FullAsset) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *FullAsset) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullAsset) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullAsset) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullAsset) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *FullAsset) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FullAsset) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FullAsset) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FullAsset) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreator

`func (o *FullAsset) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FullAsset) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FullAsset) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *FullAsset) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOwningProject

`func (o *FullAsset) GetOwningProject() MinimalProject`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *FullAsset) GetOwningProjectOk() (*MinimalProject, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *FullAsset) SetOwningProject(v MinimalProject)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *FullAsset) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *FullAsset) GetTrustedProjects() []MinimalProject`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *FullAsset) GetTrustedProjectsOk() (*[]MinimalProject, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *FullAsset) SetTrustedProjects(v []MinimalProject)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *FullAsset) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *FullAsset) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *FullAsset) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *FullAsset) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *FullAsset) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetMetadata

`func (o *FullAsset) GetMetadata() FullMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FullAsset) GetMetadataOk() (*FullMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FullAsset) SetMetadata(v FullMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FullAsset) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImpactLevel

`func (o *FullAsset) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *FullAsset) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *FullAsset) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *FullAsset) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *FullAsset) GetCategories() []MinimalCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FullAsset) GetCategoriesOk() (*[]MinimalCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FullAsset) SetCategories(v []MinimalCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *FullAsset) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSubtype

`func (o *FullAsset) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FullAsset) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FullAsset) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


