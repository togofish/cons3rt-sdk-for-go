# ActiveCompositionStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRunId** | Pointer to **int32** |  | [optional] 
**DeploymentRunStatus** | Pointer to **string** |  | [optional] 
**DeploymentRunHost** | Pointer to [**[]MinimalDeploymentRunHost**](MinimalDeploymentRunHost.md) |  | [optional] 

## Methods

### NewActiveCompositionStatusAllOf

`func NewActiveCompositionStatusAllOf() *ActiveCompositionStatusAllOf`

NewActiveCompositionStatusAllOf instantiates a new ActiveCompositionStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveCompositionStatusAllOfWithDefaults

`func NewActiveCompositionStatusAllOfWithDefaults() *ActiveCompositionStatusAllOf`

NewActiveCompositionStatusAllOfWithDefaults instantiates a new ActiveCompositionStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRunId

`func (o *ActiveCompositionStatusAllOf) GetDeploymentRunId() int32`

GetDeploymentRunId returns the DeploymentRunId field if non-nil, zero value otherwise.

### GetDeploymentRunIdOk

`func (o *ActiveCompositionStatusAllOf) GetDeploymentRunIdOk() (*int32, bool)`

GetDeploymentRunIdOk returns a tuple with the DeploymentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunId

`func (o *ActiveCompositionStatusAllOf) SetDeploymentRunId(v int32)`

SetDeploymentRunId sets DeploymentRunId field to given value.

### HasDeploymentRunId

`func (o *ActiveCompositionStatusAllOf) HasDeploymentRunId() bool`

HasDeploymentRunId returns a boolean if a field has been set.

### GetDeploymentRunStatus

`func (o *ActiveCompositionStatusAllOf) GetDeploymentRunStatus() string`

GetDeploymentRunStatus returns the DeploymentRunStatus field if non-nil, zero value otherwise.

### GetDeploymentRunStatusOk

`func (o *ActiveCompositionStatusAllOf) GetDeploymentRunStatusOk() (*string, bool)`

GetDeploymentRunStatusOk returns a tuple with the DeploymentRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunStatus

`func (o *ActiveCompositionStatusAllOf) SetDeploymentRunStatus(v string)`

SetDeploymentRunStatus sets DeploymentRunStatus field to given value.

### HasDeploymentRunStatus

`func (o *ActiveCompositionStatusAllOf) HasDeploymentRunStatus() bool`

HasDeploymentRunStatus returns a boolean if a field has been set.

### GetDeploymentRunHost

`func (o *ActiveCompositionStatusAllOf) GetDeploymentRunHost() []MinimalDeploymentRunHost`

GetDeploymentRunHost returns the DeploymentRunHost field if non-nil, zero value otherwise.

### GetDeploymentRunHostOk

`func (o *ActiveCompositionStatusAllOf) GetDeploymentRunHostOk() (*[]MinimalDeploymentRunHost, bool)`

GetDeploymentRunHostOk returns a tuple with the DeploymentRunHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunHost

`func (o *ActiveCompositionStatusAllOf) SetDeploymentRunHost(v []MinimalDeploymentRunHost)`

SetDeploymentRunHost sets DeploymentRunHost field to given value.

### HasDeploymentRunHost

`func (o *ActiveCompositionStatusAllOf) HasDeploymentRunHost() bool`

HasDeploymentRunHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


