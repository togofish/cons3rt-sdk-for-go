# MinimalDeploymentRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**MinimalProject**](MinimalProject.md) |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 
**Canceled** | Pointer to **bool** |  | [optional] 
**DeploymentRunStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FapStatus** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DeploymentRunResultType** | Pointer to **string** |  | [optional] 

## Methods

### NewMinimalDeploymentRun

`func NewMinimalDeploymentRun() *MinimalDeploymentRun`

NewMinimalDeploymentRun instantiates a new MinimalDeploymentRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalDeploymentRunWithDefaults

`func NewMinimalDeploymentRunWithDefaults() *MinimalDeploymentRun`

NewMinimalDeploymentRunWithDefaults instantiates a new MinimalDeploymentRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *MinimalDeploymentRun) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MinimalDeploymentRun) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MinimalDeploymentRun) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MinimalDeploymentRun) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetId

`func (o *MinimalDeploymentRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalDeploymentRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalDeploymentRun) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalDeploymentRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *MinimalDeploymentRun) GetProject() MinimalProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MinimalDeploymentRun) GetProjectOk() (*MinimalProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MinimalDeploymentRun) SetProject(v MinimalProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *MinimalDeploymentRun) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetResult

`func (o *MinimalDeploymentRun) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MinimalDeploymentRun) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MinimalDeploymentRun) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *MinimalDeploymentRun) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *MinimalDeploymentRun) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MinimalDeploymentRun) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MinimalDeploymentRun) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MinimalDeploymentRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCanceled

`func (o *MinimalDeploymentRun) GetCanceled() bool`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *MinimalDeploymentRun) GetCanceledOk() (*bool, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *MinimalDeploymentRun) SetCanceled(v bool)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *MinimalDeploymentRun) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetDeploymentRunStatus

`func (o *MinimalDeploymentRun) GetDeploymentRunStatus() string`

GetDeploymentRunStatus returns the DeploymentRunStatus field if non-nil, zero value otherwise.

### GetDeploymentRunStatusOk

`func (o *MinimalDeploymentRun) GetDeploymentRunStatusOk() (*string, bool)`

GetDeploymentRunStatusOk returns a tuple with the DeploymentRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunStatus

`func (o *MinimalDeploymentRun) SetDeploymentRunStatus(v string)`

SetDeploymentRunStatus sets DeploymentRunStatus field to given value.

### HasDeploymentRunStatus

`func (o *MinimalDeploymentRun) HasDeploymentRunStatus() bool`

HasDeploymentRunStatus returns a boolean if a field has been set.

### GetDescription

`func (o *MinimalDeploymentRun) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MinimalDeploymentRun) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MinimalDeploymentRun) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MinimalDeploymentRun) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFapStatus

`func (o *MinimalDeploymentRun) GetFapStatus() string`

GetFapStatus returns the FapStatus field if non-nil, zero value otherwise.

### GetFapStatusOk

`func (o *MinimalDeploymentRun) GetFapStatusOk() (*string, bool)`

GetFapStatusOk returns a tuple with the FapStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFapStatus

`func (o *MinimalDeploymentRun) SetFapStatus(v string)`

SetFapStatus sets FapStatus field to given value.

### HasFapStatus

`func (o *MinimalDeploymentRun) HasFapStatus() bool`

HasFapStatus returns a boolean if a field has been set.

### GetLocked

`func (o *MinimalDeploymentRun) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MinimalDeploymentRun) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MinimalDeploymentRun) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MinimalDeploymentRun) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *MinimalDeploymentRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalDeploymentRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalDeploymentRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MinimalDeploymentRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeploymentRunResultType

`func (o *MinimalDeploymentRun) GetDeploymentRunResultType() string`

GetDeploymentRunResultType returns the DeploymentRunResultType field if non-nil, zero value otherwise.

### GetDeploymentRunResultTypeOk

`func (o *MinimalDeploymentRun) GetDeploymentRunResultTypeOk() (*string, bool)`

GetDeploymentRunResultTypeOk returns a tuple with the DeploymentRunResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunResultType

`func (o *MinimalDeploymentRun) SetDeploymentRunResultType(v string)`

SetDeploymentRunResultType sets DeploymentRunResultType field to given value.

### HasDeploymentRunResultType

`func (o *MinimalDeploymentRun) HasDeploymentRunResultType() bool`

HasDeploymentRunResultType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


