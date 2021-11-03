# SystemModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]AbstractComponent**](AbstractComponent.md) |  | [optional] 
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
**PhysicalMachineUuid** | Pointer to **string** |  | [optional] 
**TemplateProfile** | Pointer to [**TemplateProfile**](TemplateProfile.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewSystemModule

`func NewSystemModule(subtype string, ) *SystemModule`

NewSystemModule instantiates a new SystemModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemModuleWithDefaults

`func NewSystemModuleWithDefaults() *SystemModule`

NewSystemModuleWithDefaults instantiates a new SystemModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *SystemModule) GetComponents() []AbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SystemModule) GetComponentsOk() (*[]AbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SystemModule) SetComponents(v []AbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SystemModule) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetId

`func (o *SystemModule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemModule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemModule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SystemModule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *SystemModule) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *SystemModule) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *SystemModule) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *SystemModule) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *SystemModule) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SystemModule) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SystemModule) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SystemModule) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *SystemModule) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *SystemModule) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *SystemModule) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *SystemModule) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *SystemModule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemModule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemModule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemModule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SystemModule) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SystemModule) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SystemModule) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SystemModule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SystemModule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemModule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemModule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemModule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *SystemModule) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *SystemModule) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *SystemModule) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *SystemModule) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *SystemModule) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *SystemModule) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *SystemModule) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *SystemModule) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *SystemModule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SystemModule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SystemModule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SystemModule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *SystemModule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SystemModule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SystemModule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SystemModule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *SystemModule) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *SystemModule) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *SystemModule) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *SystemModule) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *SystemModule) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SystemModule) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SystemModule) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *SystemModule) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPhysicalMachineUuid

`func (o *SystemModule) GetPhysicalMachineUuid() string`

GetPhysicalMachineUuid returns the PhysicalMachineUuid field if non-nil, zero value otherwise.

### GetPhysicalMachineUuidOk

`func (o *SystemModule) GetPhysicalMachineUuidOk() (*string, bool)`

GetPhysicalMachineUuidOk returns a tuple with the PhysicalMachineUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachineUuid

`func (o *SystemModule) SetPhysicalMachineUuid(v string)`

SetPhysicalMachineUuid sets PhysicalMachineUuid field to given value.

### HasPhysicalMachineUuid

`func (o *SystemModule) HasPhysicalMachineUuid() bool`

HasPhysicalMachineUuid returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *SystemModule) GetTemplateProfile() TemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *SystemModule) GetTemplateProfileOk() (*TemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *SystemModule) SetTemplateProfile(v TemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *SystemModule) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.

### GetSubtype

`func (o *SystemModule) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SystemModule) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SystemModule) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


