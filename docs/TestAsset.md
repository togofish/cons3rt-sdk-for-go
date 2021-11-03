# TestAsset

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
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTestAsset

`func NewTestAsset() *TestAsset`

NewTestAsset instantiates a new TestAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestAssetWithDefaults

`func NewTestAssetWithDefaults() *TestAsset`

NewTestAssetWithDefaults instantiates a new TestAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *TestAsset) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *TestAsset) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *TestAsset) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *TestAsset) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *TestAsset) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *TestAsset) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *TestAsset) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *TestAsset) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *TestAsset) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *TestAsset) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *TestAsset) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *TestAsset) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *TestAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *TestAsset) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TestAsset) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TestAsset) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TestAsset) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *TestAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *TestAsset) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *TestAsset) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *TestAsset) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *TestAsset) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *TestAsset) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *TestAsset) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *TestAsset) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *TestAsset) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *TestAsset) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestAsset) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestAsset) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TestAsset) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *TestAsset) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TestAsset) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TestAsset) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TestAsset) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *TestAsset) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *TestAsset) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *TestAsset) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *TestAsset) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *TestAsset) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TestAsset) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TestAsset) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TestAsset) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetType

`func (o *TestAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestAsset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TestAsset) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


