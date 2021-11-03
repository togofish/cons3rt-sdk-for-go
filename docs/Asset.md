# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TrustedProjects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] 
**DependentAssetCount** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**OwningProject** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ImpactLevel** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *Asset) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *Asset) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *Asset) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *Asset) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *Asset) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Asset) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Asset) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Asset) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *Asset) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *Asset) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *Asset) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *Asset) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Asset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Asset) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Asset) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Asset) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Asset) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *Asset) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *Asset) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *Asset) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *Asset) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *Asset) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *Asset) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *Asset) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *Asset) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *Asset) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Asset) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Asset) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Asset) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *Asset) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Asset) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Asset) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Asset) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *Asset) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *Asset) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *Asset) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *Asset) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *Asset) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Asset) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Asset) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Asset) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


