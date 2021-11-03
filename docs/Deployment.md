# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurringSchedules** | Pointer to [**[]RecurringSchedule**](RecurringSchedule.md) |  | [optional] 
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
**Scenarios** | Pointer to [**[]DeploymentScenario**](DeploymentScenario.md) |  | [optional] 
**TestBundles** | Pointer to [**[]TestBundle**](TestBundle.md) |  | [optional] 
**DeploymentHosts** | Pointer to [**[]DeploymentHost**](DeploymentHost.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurringSchedules

`func (o *Deployment) GetRecurringSchedules() []RecurringSchedule`

GetRecurringSchedules returns the RecurringSchedules field if non-nil, zero value otherwise.

### GetRecurringSchedulesOk

`func (o *Deployment) GetRecurringSchedulesOk() (*[]RecurringSchedule, bool)`

GetRecurringSchedulesOk returns a tuple with the RecurringSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSchedules

`func (o *Deployment) SetRecurringSchedules(v []RecurringSchedule)`

SetRecurringSchedules sets RecurringSchedules field to given value.

### HasRecurringSchedules

`func (o *Deployment) HasRecurringSchedules() bool`

HasRecurringSchedules returns a boolean if a field has been set.

### GetId

`func (o *Deployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *Deployment) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *Deployment) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *Deployment) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *Deployment) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *Deployment) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Deployment) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Deployment) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Deployment) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *Deployment) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *Deployment) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *Deployment) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *Deployment) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *Deployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Deployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Deployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Deployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Deployment) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Deployment) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Deployment) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Deployment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Deployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Deployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *Deployment) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *Deployment) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *Deployment) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *Deployment) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *Deployment) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *Deployment) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *Deployment) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *Deployment) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *Deployment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Deployment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Deployment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Deployment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *Deployment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Deployment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Deployment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Deployment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *Deployment) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *Deployment) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *Deployment) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *Deployment) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *Deployment) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Deployment) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Deployment) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Deployment) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetScenarios

`func (o *Deployment) GetScenarios() []DeploymentScenario`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *Deployment) GetScenariosOk() (*[]DeploymentScenario, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *Deployment) SetScenarios(v []DeploymentScenario)`

SetScenarios sets Scenarios field to given value.

### HasScenarios

`func (o *Deployment) HasScenarios() bool`

HasScenarios returns a boolean if a field has been set.

### GetTestBundles

`func (o *Deployment) GetTestBundles() []TestBundle`

GetTestBundles returns the TestBundles field if non-nil, zero value otherwise.

### GetTestBundlesOk

`func (o *Deployment) GetTestBundlesOk() (*[]TestBundle, bool)`

GetTestBundlesOk returns a tuple with the TestBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBundles

`func (o *Deployment) SetTestBundles(v []TestBundle)`

SetTestBundles sets TestBundles field to given value.

### HasTestBundles

`func (o *Deployment) HasTestBundles() bool`

HasTestBundles returns a boolean if a field has been set.

### GetDeploymentHosts

`func (o *Deployment) GetDeploymentHosts() []DeploymentHost`

GetDeploymentHosts returns the DeploymentHosts field if non-nil, zero value otherwise.

### GetDeploymentHostsOk

`func (o *Deployment) GetDeploymentHostsOk() (*[]DeploymentHost, bool)`

GetDeploymentHostsOk returns a tuple with the DeploymentHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentHosts

`func (o *Deployment) SetDeploymentHosts(v []DeploymentHost)`

SetDeploymentHosts sets DeploymentHosts field to given value.

### HasDeploymentHosts

`func (o *Deployment) HasDeploymentHosts() bool`

HasDeploymentHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


