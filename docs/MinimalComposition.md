# MinimalComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeploymentRunOptions** | Pointer to [**MinimalDeploymentRunOptions**](MinimalDeploymentRunOptions.md) |  | [optional] 
**Project** | Pointer to [**MinimalProject**](MinimalProject.md) |  | [optional] 
**Scenario** | Pointer to [**MinimalScenario**](MinimalScenario.md) |  | [optional] 

## Methods

### NewMinimalComposition

`func NewMinimalComposition() *MinimalComposition`

NewMinimalComposition instantiates a new MinimalComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalCompositionWithDefaults

`func NewMinimalCompositionWithDefaults() *MinimalComposition`

NewMinimalCompositionWithDefaults instantiates a new MinimalComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalComposition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalComposition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalComposition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalComposition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MinimalComposition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalComposition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalComposition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MinimalComposition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MinimalComposition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MinimalComposition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MinimalComposition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MinimalComposition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentRunOptions

`func (o *MinimalComposition) GetDeploymentRunOptions() MinimalDeploymentRunOptions`

GetDeploymentRunOptions returns the DeploymentRunOptions field if non-nil, zero value otherwise.

### GetDeploymentRunOptionsOk

`func (o *MinimalComposition) GetDeploymentRunOptionsOk() (*MinimalDeploymentRunOptions, bool)`

GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunOptions

`func (o *MinimalComposition) SetDeploymentRunOptions(v MinimalDeploymentRunOptions)`

SetDeploymentRunOptions sets DeploymentRunOptions field to given value.

### HasDeploymentRunOptions

`func (o *MinimalComposition) HasDeploymentRunOptions() bool`

HasDeploymentRunOptions returns a boolean if a field has been set.

### GetProject

`func (o *MinimalComposition) GetProject() MinimalProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MinimalComposition) GetProjectOk() (*MinimalProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MinimalComposition) SetProject(v MinimalProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *MinimalComposition) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetScenario

`func (o *MinimalComposition) GetScenario() MinimalScenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *MinimalComposition) GetScenarioOk() (*MinimalScenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *MinimalComposition) SetScenario(v MinimalScenario)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *MinimalComposition) HasScenario() bool`

HasScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


