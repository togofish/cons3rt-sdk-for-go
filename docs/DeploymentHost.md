# DeploymentHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | Pointer to [**Deployment**](Deployment.md) |  | [optional] 
**Scenario** | Pointer to [**Scenario**](Scenario.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**BuildOrder** | Pointer to **int32** |  | [optional] 
**SystemModule** | Pointer to [**SystemModule**](SystemModule.md) |  | [optional] 
**SystemRole** | **string** |  | 

## Methods

### NewDeploymentHost

`func NewDeploymentHost(systemRole string, ) *DeploymentHost`

NewDeploymentHost instantiates a new DeploymentHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHostWithDefaults

`func NewDeploymentHostWithDefaults() *DeploymentHost`

NewDeploymentHostWithDefaults instantiates a new DeploymentHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *DeploymentHost) GetDeployment() Deployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *DeploymentHost) GetDeploymentOk() (*Deployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *DeploymentHost) SetDeployment(v Deployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *DeploymentHost) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetScenario

`func (o *DeploymentHost) GetScenario() Scenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *DeploymentHost) GetScenarioOk() (*Scenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *DeploymentHost) SetScenario(v Scenario)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *DeploymentHost) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetId

`func (o *DeploymentHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBuildOrder

`func (o *DeploymentHost) GetBuildOrder() int32`

GetBuildOrder returns the BuildOrder field if non-nil, zero value otherwise.

### GetBuildOrderOk

`func (o *DeploymentHost) GetBuildOrderOk() (*int32, bool)`

GetBuildOrderOk returns a tuple with the BuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOrder

`func (o *DeploymentHost) SetBuildOrder(v int32)`

SetBuildOrder sets BuildOrder field to given value.

### HasBuildOrder

`func (o *DeploymentHost) HasBuildOrder() bool`

HasBuildOrder returns a boolean if a field has been set.

### GetSystemModule

`func (o *DeploymentHost) GetSystemModule() SystemModule`

GetSystemModule returns the SystemModule field if non-nil, zero value otherwise.

### GetSystemModuleOk

`func (o *DeploymentHost) GetSystemModuleOk() (*SystemModule, bool)`

GetSystemModuleOk returns a tuple with the SystemModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModule

`func (o *DeploymentHost) SetSystemModule(v SystemModule)`

SetSystemModule sets SystemModule field to given value.

### HasSystemModule

`func (o *DeploymentHost) HasSystemModule() bool`

HasSystemModule returns a boolean if a field has been set.

### GetSystemRole

`func (o *DeploymentHost) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *DeploymentHost) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *DeploymentHost) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


