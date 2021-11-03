# DeploymentScenario

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
**ScenarioHosts** | Pointer to [**[]ScenarioHost**](ScenarioHost.md) |  | [optional] 
**PrepareScenarioConfiguration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] 
**ScenarioBuildOrder** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeploymentScenario

`func NewDeploymentScenario() *DeploymentScenario`

NewDeploymentScenario instantiates a new DeploymentScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentScenarioWithDefaults

`func NewDeploymentScenarioWithDefaults() *DeploymentScenario`

NewDeploymentScenarioWithDefaults instantiates a new DeploymentScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentScenario) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentScenario) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentScenario) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentScenario) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *DeploymentScenario) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *DeploymentScenario) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *DeploymentScenario) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *DeploymentScenario) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *DeploymentScenario) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *DeploymentScenario) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *DeploymentScenario) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *DeploymentScenario) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *DeploymentScenario) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *DeploymentScenario) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *DeploymentScenario) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *DeploymentScenario) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentScenario) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentScenario) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentScenario) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentScenario) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *DeploymentScenario) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeploymentScenario) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeploymentScenario) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeploymentScenario) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DeploymentScenario) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentScenario) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentScenario) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentScenario) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *DeploymentScenario) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *DeploymentScenario) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *DeploymentScenario) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *DeploymentScenario) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *DeploymentScenario) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *DeploymentScenario) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *DeploymentScenario) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *DeploymentScenario) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *DeploymentScenario) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentScenario) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentScenario) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentScenario) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *DeploymentScenario) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *DeploymentScenario) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *DeploymentScenario) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *DeploymentScenario) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *DeploymentScenario) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *DeploymentScenario) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *DeploymentScenario) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *DeploymentScenario) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *DeploymentScenario) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *DeploymentScenario) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *DeploymentScenario) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *DeploymentScenario) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetScenarioHosts

`func (o *DeploymentScenario) GetScenarioHosts() []ScenarioHost`

GetScenarioHosts returns the ScenarioHosts field if non-nil, zero value otherwise.

### GetScenarioHostsOk

`func (o *DeploymentScenario) GetScenarioHostsOk() (*[]ScenarioHost, bool)`

GetScenarioHostsOk returns a tuple with the ScenarioHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioHosts

`func (o *DeploymentScenario) SetScenarioHosts(v []ScenarioHost)`

SetScenarioHosts sets ScenarioHosts field to given value.

### HasScenarioHosts

`func (o *DeploymentScenario) HasScenarioHosts() bool`

HasScenarioHosts returns a boolean if a field has been set.

### GetPrepareScenarioConfiguration

`func (o *DeploymentScenario) GetPrepareScenarioConfiguration() Configuration`

GetPrepareScenarioConfiguration returns the PrepareScenarioConfiguration field if non-nil, zero value otherwise.

### GetPrepareScenarioConfigurationOk

`func (o *DeploymentScenario) GetPrepareScenarioConfigurationOk() (*Configuration, bool)`

GetPrepareScenarioConfigurationOk returns a tuple with the PrepareScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareScenarioConfiguration

`func (o *DeploymentScenario) SetPrepareScenarioConfiguration(v Configuration)`

SetPrepareScenarioConfiguration sets PrepareScenarioConfiguration field to given value.

### HasPrepareScenarioConfiguration

`func (o *DeploymentScenario) HasPrepareScenarioConfiguration() bool`

HasPrepareScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *DeploymentScenario) GetTeardownScenarioConfiguration() Configuration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *DeploymentScenario) GetTeardownScenarioConfigurationOk() (*Configuration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *DeploymentScenario) SetTeardownScenarioConfiguration(v Configuration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *DeploymentScenario) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.

### GetScenarioBuildOrder

`func (o *DeploymentScenario) GetScenarioBuildOrder() int32`

GetScenarioBuildOrder returns the ScenarioBuildOrder field if non-nil, zero value otherwise.

### GetScenarioBuildOrderOk

`func (o *DeploymentScenario) GetScenarioBuildOrderOk() (*int32, bool)`

GetScenarioBuildOrderOk returns a tuple with the ScenarioBuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioBuildOrder

`func (o *DeploymentScenario) SetScenarioBuildOrder(v int32)`

SetScenarioBuildOrder sets ScenarioBuildOrder field to given value.

### HasScenarioBuildOrder

`func (o *DeploymentScenario) HasScenarioBuildOrder() bool`

HasScenarioBuildOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


