# RecurringSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Complete** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**MaxIterations** | Pointer to **int32** |  | [optional] 
**DeploymentRunOptions** | [**DeploymentRunOptions**](DeploymentRunOptions.md) |  | 
**RemainingIterations** | Pointer to **int32** |  | [optional] 
**Schedule** | **string** |  | 
**Timezone** | **string** |  | 

## Methods

### NewRecurringSchedule

`func NewRecurringSchedule(deploymentRunOptions DeploymentRunOptions, schedule string, timezone string, ) *RecurringSchedule`

NewRecurringSchedule instantiates a new RecurringSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringScheduleWithDefaults

`func NewRecurringScheduleWithDefaults() *RecurringSchedule`

NewRecurringScheduleWithDefaults instantiates a new RecurringSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurringSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurringSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurringSchedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RecurringSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComplete

`func (o *RecurringSchedule) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *RecurringSchedule) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *RecurringSchedule) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *RecurringSchedule) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetEndDate

`func (o *RecurringSchedule) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RecurringSchedule) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RecurringSchedule) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *RecurringSchedule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxIterations

`func (o *RecurringSchedule) GetMaxIterations() int32`

GetMaxIterations returns the MaxIterations field if non-nil, zero value otherwise.

### GetMaxIterationsOk

`func (o *RecurringSchedule) GetMaxIterationsOk() (*int32, bool)`

GetMaxIterationsOk returns a tuple with the MaxIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIterations

`func (o *RecurringSchedule) SetMaxIterations(v int32)`

SetMaxIterations sets MaxIterations field to given value.

### HasMaxIterations

`func (o *RecurringSchedule) HasMaxIterations() bool`

HasMaxIterations returns a boolean if a field has been set.

### GetDeploymentRunOptions

`func (o *RecurringSchedule) GetDeploymentRunOptions() DeploymentRunOptions`

GetDeploymentRunOptions returns the DeploymentRunOptions field if non-nil, zero value otherwise.

### GetDeploymentRunOptionsOk

`func (o *RecurringSchedule) GetDeploymentRunOptionsOk() (*DeploymentRunOptions, bool)`

GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunOptions

`func (o *RecurringSchedule) SetDeploymentRunOptions(v DeploymentRunOptions)`

SetDeploymentRunOptions sets DeploymentRunOptions field to given value.


### GetRemainingIterations

`func (o *RecurringSchedule) GetRemainingIterations() int32`

GetRemainingIterations returns the RemainingIterations field if non-nil, zero value otherwise.

### GetRemainingIterationsOk

`func (o *RecurringSchedule) GetRemainingIterationsOk() (*int32, bool)`

GetRemainingIterationsOk returns a tuple with the RemainingIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingIterations

`func (o *RecurringSchedule) SetRemainingIterations(v int32)`

SetRemainingIterations sets RemainingIterations field to given value.

### HasRemainingIterations

`func (o *RecurringSchedule) HasRemainingIterations() bool`

HasRemainingIterations returns a boolean if a field has been set.

### GetSchedule

`func (o *RecurringSchedule) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RecurringSchedule) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RecurringSchedule) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetTimezone

`func (o *RecurringSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RecurringSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RecurringSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


