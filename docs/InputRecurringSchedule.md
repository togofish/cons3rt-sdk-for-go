# InputRecurringSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** |  | 
**Schedule** | **string** |  | 
**MaxIterations** | Pointer to **int32** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**DeploymentRunOptions** | [**InputDeploymentRunOptions**](InputDeploymentRunOptions.md) |  | 

## Methods

### NewInputRecurringSchedule

`func NewInputRecurringSchedule(timezone string, schedule string, deploymentRunOptions InputDeploymentRunOptions, ) *InputRecurringSchedule`

NewInputRecurringSchedule instantiates a new InputRecurringSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRecurringScheduleWithDefaults

`func NewInputRecurringScheduleWithDefaults() *InputRecurringSchedule`

NewInputRecurringScheduleWithDefaults instantiates a new InputRecurringSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *InputRecurringSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InputRecurringSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InputRecurringSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetSchedule

`func (o *InputRecurringSchedule) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InputRecurringSchedule) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InputRecurringSchedule) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetMaxIterations

`func (o *InputRecurringSchedule) GetMaxIterations() int32`

GetMaxIterations returns the MaxIterations field if non-nil, zero value otherwise.

### GetMaxIterationsOk

`func (o *InputRecurringSchedule) GetMaxIterationsOk() (*int32, bool)`

GetMaxIterationsOk returns a tuple with the MaxIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIterations

`func (o *InputRecurringSchedule) SetMaxIterations(v int32)`

SetMaxIterations sets MaxIterations field to given value.

### HasMaxIterations

`func (o *InputRecurringSchedule) HasMaxIterations() bool`

HasMaxIterations returns a boolean if a field has been set.

### GetEndDate

`func (o *InputRecurringSchedule) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputRecurringSchedule) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputRecurringSchedule) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputRecurringSchedule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDeploymentRunOptions

`func (o *InputRecurringSchedule) GetDeploymentRunOptions() InputDeploymentRunOptions`

GetDeploymentRunOptions returns the DeploymentRunOptions field if non-nil, zero value otherwise.

### GetDeploymentRunOptionsOk

`func (o *InputRecurringSchedule) GetDeploymentRunOptionsOk() (*InputDeploymentRunOptions, bool)`

GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunOptions

`func (o *InputRecurringSchedule) SetDeploymentRunOptions(v InputDeploymentRunOptions)`

SetDeploymentRunOptions sets DeploymentRunOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


