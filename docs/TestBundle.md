# TestBundle

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
**TestAsset** | [**TestAsset**](TestAsset.md) |  | 
**TestToolPoolType** | **string** |  | 

## Methods

### NewTestBundle

`func NewTestBundle(testAsset TestAsset, testToolPoolType string, ) *TestBundle`

NewTestBundle instantiates a new TestBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestBundleWithDefaults

`func NewTestBundleWithDefaults() *TestBundle`

NewTestBundleWithDefaults instantiates a new TestBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestBundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestBundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestBundle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestBundle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *TestBundle) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *TestBundle) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *TestBundle) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *TestBundle) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *TestBundle) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *TestBundle) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *TestBundle) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *TestBundle) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *TestBundle) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *TestBundle) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *TestBundle) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *TestBundle) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *TestBundle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestBundle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestBundle) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestBundle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *TestBundle) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TestBundle) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TestBundle) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TestBundle) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *TestBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *TestBundle) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *TestBundle) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *TestBundle) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *TestBundle) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *TestBundle) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *TestBundle) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *TestBundle) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *TestBundle) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *TestBundle) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestBundle) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestBundle) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TestBundle) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *TestBundle) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TestBundle) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TestBundle) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TestBundle) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *TestBundle) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *TestBundle) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *TestBundle) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *TestBundle) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *TestBundle) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TestBundle) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TestBundle) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TestBundle) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTestAsset

`func (o *TestBundle) GetTestAsset() TestAsset`

GetTestAsset returns the TestAsset field if non-nil, zero value otherwise.

### GetTestAssetOk

`func (o *TestBundle) GetTestAssetOk() (*TestAsset, bool)`

GetTestAssetOk returns a tuple with the TestAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAsset

`func (o *TestBundle) SetTestAsset(v TestAsset)`

SetTestAsset sets TestAsset field to given value.


### GetTestToolPoolType

`func (o *TestBundle) GetTestToolPoolType() string`

GetTestToolPoolType returns the TestToolPoolType field if non-nil, zero value otherwise.

### GetTestToolPoolTypeOk

`func (o *TestBundle) GetTestToolPoolTypeOk() (*string, bool)`

GetTestToolPoolTypeOk returns a tuple with the TestToolPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestToolPoolType

`func (o *TestBundle) SetTestToolPoolType(v string)`

SetTestToolPoolType sets TestToolPoolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


