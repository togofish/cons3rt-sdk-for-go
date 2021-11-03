# FullSystemModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]MinimalAbstractComponent**](MinimalAbstractComponent.md) |  | [optional] 
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

### NewFullSystemModule

`func NewFullSystemModule(subtype string, ) *FullSystemModule`

NewFullSystemModule instantiates a new FullSystemModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullSystemModuleWithDefaults

`func NewFullSystemModuleWithDefaults() *FullSystemModule`

NewFullSystemModuleWithDefaults instantiates a new FullSystemModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *FullSystemModule) GetComponents() []MinimalAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FullSystemModule) GetComponentsOk() (*[]MinimalAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FullSystemModule) SetComponents(v []MinimalAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FullSystemModule) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetId

`func (o *FullSystemModule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullSystemModule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullSystemModule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullSystemModule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullSystemModule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullSystemModule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullSystemModule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullSystemModule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FullSystemModule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullSystemModule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullSystemModule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullSystemModule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *FullSystemModule) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *FullSystemModule) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *FullSystemModule) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *FullSystemModule) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *FullSystemModule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullSystemModule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullSystemModule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullSystemModule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *FullSystemModule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FullSystemModule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FullSystemModule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FullSystemModule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCreator

`func (o *FullSystemModule) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FullSystemModule) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FullSystemModule) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *FullSystemModule) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOwningProject

`func (o *FullSystemModule) GetOwningProject() MinimalProject`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *FullSystemModule) GetOwningProjectOk() (*MinimalProject, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *FullSystemModule) SetOwningProject(v MinimalProject)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *FullSystemModule) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *FullSystemModule) GetTrustedProjects() []MinimalProject`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *FullSystemModule) GetTrustedProjectsOk() (*[]MinimalProject, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *FullSystemModule) SetTrustedProjects(v []MinimalProject)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *FullSystemModule) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *FullSystemModule) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *FullSystemModule) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *FullSystemModule) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *FullSystemModule) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetMetadata

`func (o *FullSystemModule) GetMetadata() FullMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FullSystemModule) GetMetadataOk() (*FullMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FullSystemModule) SetMetadata(v FullMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FullSystemModule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImpactLevel

`func (o *FullSystemModule) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *FullSystemModule) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *FullSystemModule) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *FullSystemModule) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *FullSystemModule) GetCategories() []MinimalCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FullSystemModule) GetCategoriesOk() (*[]MinimalCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FullSystemModule) SetCategories(v []MinimalCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *FullSystemModule) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSubtype

`func (o *FullSystemModule) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FullSystemModule) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FullSystemModule) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


