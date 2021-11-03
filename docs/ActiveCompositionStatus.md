# ActiveCompositionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRunId** | Pointer to **int32** |  | [optional] 
**DeploymentRunStatus** | Pointer to **string** |  | [optional] 
**DeploymentRunHost** | Pointer to [**[]MinimalDeploymentRunHost**](MinimalDeploymentRunHost.md) |  | [optional] 

## Methods

### NewActiveCompositionStatus

`func NewActiveCompositionStatus() *ActiveCompositionStatus`

NewActiveCompositionStatus instantiates a new ActiveCompositionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveCompositionStatusWithDefaults

`func NewActiveCompositionStatusWithDefaults() *ActiveCompositionStatus`

NewActiveCompositionStatusWithDefaults instantiates a new ActiveCompositionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRunId

`func (o *ActiveCompositionStatus) GetDeploymentRunId() int32`

GetDeploymentRunId returns the DeploymentRunId field if non-nil, zero value otherwise.

### GetDeploymentRunIdOk

`func (o *ActiveCompositionStatus) GetDeploymentRunIdOk() (*int32, bool)`

GetDeploymentRunIdOk returns a tuple with the DeploymentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunId

`func (o *ActiveCompositionStatus) SetDeploymentRunId(v int32)`

SetDeploymentRunId sets DeploymentRunId field to given value.

### HasDeploymentRunId

`func (o *ActiveCompositionStatus) HasDeploymentRunId() bool`

HasDeploymentRunId returns a boolean if a field has been set.

### GetDeploymentRunStatus

`func (o *ActiveCompositionStatus) GetDeploymentRunStatus() string`

GetDeploymentRunStatus returns the DeploymentRunStatus field if non-nil, zero value otherwise.

### GetDeploymentRunStatusOk

`func (o *ActiveCompositionStatus) GetDeploymentRunStatusOk() (*string, bool)`

GetDeploymentRunStatusOk returns a tuple with the DeploymentRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunStatus

`func (o *ActiveCompositionStatus) SetDeploymentRunStatus(v string)`

SetDeploymentRunStatus sets DeploymentRunStatus field to given value.

### HasDeploymentRunStatus

`func (o *ActiveCompositionStatus) HasDeploymentRunStatus() bool`

HasDeploymentRunStatus returns a boolean if a field has been set.

### GetDeploymentRunHost

`func (o *ActiveCompositionStatus) GetDeploymentRunHost() []MinimalDeploymentRunHost`

GetDeploymentRunHost returns the DeploymentRunHost field if non-nil, zero value otherwise.

### GetDeploymentRunHostOk

`func (o *ActiveCompositionStatus) GetDeploymentRunHostOk() (*[]MinimalDeploymentRunHost, bool)`

GetDeploymentRunHostOk returns a tuple with the DeploymentRunHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunHost

`func (o *ActiveCompositionStatus) SetDeploymentRunHost(v []MinimalDeploymentRunHost)`

SetDeploymentRunHost sets DeploymentRunHost field to given value.

### HasDeploymentRunHost

`func (o *ActiveCompositionStatus) HasDeploymentRunHost() bool`

HasDeploymentRunHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


