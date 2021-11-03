# Composition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRunOptions** | Pointer to [**CompositionRunOptions**](CompositionRunOptions.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Scenario** | Pointer to [**Scenario**](Scenario.md) |  | [optional] 

## Methods

### NewComposition

`func NewComposition(name string, ) *Composition`

NewComposition instantiates a new Composition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositionWithDefaults

`func NewCompositionWithDefaults() *Composition`

NewCompositionWithDefaults instantiates a new Composition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRunOptions

`func (o *Composition) GetDeploymentRunOptions() CompositionRunOptions`

GetDeploymentRunOptions returns the DeploymentRunOptions field if non-nil, zero value otherwise.

### GetDeploymentRunOptionsOk

`func (o *Composition) GetDeploymentRunOptionsOk() (*CompositionRunOptions, bool)`

GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunOptions

`func (o *Composition) SetDeploymentRunOptions(v CompositionRunOptions)`

SetDeploymentRunOptions sets DeploymentRunOptions field to given value.

### HasDeploymentRunOptions

`func (o *Composition) HasDeploymentRunOptions() bool`

HasDeploymentRunOptions returns a boolean if a field has been set.

### GetDescription

`func (o *Composition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Composition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Composition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Composition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Composition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Composition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Composition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Composition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Composition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Composition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Composition) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *Composition) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Composition) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Composition) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Composition) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetScenario

`func (o *Composition) GetScenario() Scenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *Composition) GetScenarioOk() (*Scenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *Composition) SetScenario(v Scenario)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *Composition) HasScenario() bool`

HasScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


