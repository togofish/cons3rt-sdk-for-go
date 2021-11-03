# FullContainerAsset

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
**Type** | Pointer to **string** |  | [optional] 
**ImageFile** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**ImageTag** | Pointer to **string** |  | [optional] 

## Methods

### NewFullContainerAsset

`func NewFullContainerAsset() *FullContainerAsset`

NewFullContainerAsset instantiates a new FullContainerAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullContainerAssetWithDefaults

`func NewFullContainerAssetWithDefaults() *FullContainerAsset`

NewFullContainerAssetWithDefaults instantiates a new FullContainerAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullContainerAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullContainerAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullContainerAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullContainerAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullContainerAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullContainerAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullContainerAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullContainerAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FullContainerAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullContainerAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullContainerAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullContainerAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *FullContainerAsset) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *FullContainerAsset) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *FullContainerAsset) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *FullContainerAsset) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *FullContainerAsset) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullContainerAsset) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullContainerAsset) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullContainerAsset) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *FullContainerAsset) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FullContainerAsset) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FullContainerAsset) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FullContainerAsset) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreator

`func (o *FullContainerAsset) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FullContainerAsset) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FullContainerAsset) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *FullContainerAsset) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOwningProject

`func (o *FullContainerAsset) GetOwningProject() MinimalProject`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *FullContainerAsset) GetOwningProjectOk() (*MinimalProject, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *FullContainerAsset) SetOwningProject(v MinimalProject)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *FullContainerAsset) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *FullContainerAsset) GetTrustedProjects() []MinimalProject`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *FullContainerAsset) GetTrustedProjectsOk() (*[]MinimalProject, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *FullContainerAsset) SetTrustedProjects(v []MinimalProject)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *FullContainerAsset) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *FullContainerAsset) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *FullContainerAsset) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *FullContainerAsset) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *FullContainerAsset) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetMetadata

`func (o *FullContainerAsset) GetMetadata() FullMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FullContainerAsset) GetMetadataOk() (*FullMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FullContainerAsset) SetMetadata(v FullMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FullContainerAsset) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImpactLevel

`func (o *FullContainerAsset) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *FullContainerAsset) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *FullContainerAsset) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *FullContainerAsset) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *FullContainerAsset) GetCategories() []MinimalCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FullContainerAsset) GetCategoriesOk() (*[]MinimalCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FullContainerAsset) SetCategories(v []MinimalCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *FullContainerAsset) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetType

`func (o *FullContainerAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullContainerAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullContainerAsset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FullContainerAsset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageFile

`func (o *FullContainerAsset) GetImageFile() string`

GetImageFile returns the ImageFile field if non-nil, zero value otherwise.

### GetImageFileOk

`func (o *FullContainerAsset) GetImageFileOk() (*string, bool)`

GetImageFileOk returns a tuple with the ImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFile

`func (o *FullContainerAsset) SetImageFile(v string)`

SetImageFile sets ImageFile field to given value.

### HasImageFile

`func (o *FullContainerAsset) HasImageFile() bool`

HasImageFile returns a boolean if a field has been set.

### GetImageName

`func (o *FullContainerAsset) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *FullContainerAsset) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *FullContainerAsset) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *FullContainerAsset) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageTag

`func (o *FullContainerAsset) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *FullContainerAsset) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *FullContainerAsset) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *FullContainerAsset) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


