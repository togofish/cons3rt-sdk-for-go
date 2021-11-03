# InputComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DeploymentRunOptions** | [**InputCompositionRunOptions**](InputCompositionRunOptions.md) |  | 

## Methods

### NewInputComposition

`func NewInputComposition(name string, deploymentRunOptions InputCompositionRunOptions, ) *InputComposition`

NewInputComposition instantiates a new InputComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCompositionWithDefaults

`func NewInputCompositionWithDefaults() *InputComposition`

NewInputCompositionWithDefaults instantiates a new InputComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputComposition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputComposition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputComposition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InputComposition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputComposition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputComposition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputComposition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentRunOptions

`func (o *InputComposition) GetDeploymentRunOptions() InputCompositionRunOptions`

GetDeploymentRunOptions returns the DeploymentRunOptions field if non-nil, zero value otherwise.

### GetDeploymentRunOptionsOk

`func (o *InputComposition) GetDeploymentRunOptionsOk() (*InputCompositionRunOptions, bool)`

GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunOptions

`func (o *InputComposition) SetDeploymentRunOptions(v InputCompositionRunOptions)`

SetDeploymentRunOptions sets DeploymentRunOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


