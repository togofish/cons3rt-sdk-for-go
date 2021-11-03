# PowerSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**WeekdayEndTimeHour** | Pointer to **int32** |  | [optional] 
**WeekdayEndTimeMinutes** | Pointer to **int32** |  | [optional] 
**WeekdayStartTimeHour** | Pointer to **int32** |  | [optional] 
**WeekdayStartTimeMinutes** | Pointer to **int32** |  | [optional] 
**WeekendEndTimeHour** | Pointer to **int32** |  | [optional] 
**WeekendEndTimeMinutes** | Pointer to **int32** |  | [optional] 
**WeekendStartTimeHour** | Pointer to **int32** |  | [optional] 
**WeekendStartTimeMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewPowerSchedule

`func NewPowerSchedule(mode string, ) *PowerSchedule`

NewPowerSchedule instantiates a new PowerSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerScheduleWithDefaults

`func NewPowerScheduleWithDefaults() *PowerSchedule`

NewPowerScheduleWithDefaults instantiates a new PowerSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PowerSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PowerSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PowerSchedule) SetMode(v string)`

SetMode sets Mode field to given value.


### GetWeekdayEndTimeHour

`func (o *PowerSchedule) GetWeekdayEndTimeHour() int32`

GetWeekdayEndTimeHour returns the WeekdayEndTimeHour field if non-nil, zero value otherwise.

### GetWeekdayEndTimeHourOk

`func (o *PowerSchedule) GetWeekdayEndTimeHourOk() (*int32, bool)`

GetWeekdayEndTimeHourOk returns a tuple with the WeekdayEndTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayEndTimeHour

`func (o *PowerSchedule) SetWeekdayEndTimeHour(v int32)`

SetWeekdayEndTimeHour sets WeekdayEndTimeHour field to given value.

### HasWeekdayEndTimeHour

`func (o *PowerSchedule) HasWeekdayEndTimeHour() bool`

HasWeekdayEndTimeHour returns a boolean if a field has been set.

### GetWeekdayEndTimeMinutes

`func (o *PowerSchedule) GetWeekdayEndTimeMinutes() int32`

GetWeekdayEndTimeMinutes returns the WeekdayEndTimeMinutes field if non-nil, zero value otherwise.

### GetWeekdayEndTimeMinutesOk

`func (o *PowerSchedule) GetWeekdayEndTimeMinutesOk() (*int32, bool)`

GetWeekdayEndTimeMinutesOk returns a tuple with the WeekdayEndTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayEndTimeMinutes

`func (o *PowerSchedule) SetWeekdayEndTimeMinutes(v int32)`

SetWeekdayEndTimeMinutes sets WeekdayEndTimeMinutes field to given value.

### HasWeekdayEndTimeMinutes

`func (o *PowerSchedule) HasWeekdayEndTimeMinutes() bool`

HasWeekdayEndTimeMinutes returns a boolean if a field has been set.

### GetWeekdayStartTimeHour

`func (o *PowerSchedule) GetWeekdayStartTimeHour() int32`

GetWeekdayStartTimeHour returns the WeekdayStartTimeHour field if non-nil, zero value otherwise.

### GetWeekdayStartTimeHourOk

`func (o *PowerSchedule) GetWeekdayStartTimeHourOk() (*int32, bool)`

GetWeekdayStartTimeHourOk returns a tuple with the WeekdayStartTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayStartTimeHour

`func (o *PowerSchedule) SetWeekdayStartTimeHour(v int32)`

SetWeekdayStartTimeHour sets WeekdayStartTimeHour field to given value.

### HasWeekdayStartTimeHour

`func (o *PowerSchedule) HasWeekdayStartTimeHour() bool`

HasWeekdayStartTimeHour returns a boolean if a field has been set.

### GetWeekdayStartTimeMinutes

`func (o *PowerSchedule) GetWeekdayStartTimeMinutes() int32`

GetWeekdayStartTimeMinutes returns the WeekdayStartTimeMinutes field if non-nil, zero value otherwise.

### GetWeekdayStartTimeMinutesOk

`func (o *PowerSchedule) GetWeekdayStartTimeMinutesOk() (*int32, bool)`

GetWeekdayStartTimeMinutesOk returns a tuple with the WeekdayStartTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayStartTimeMinutes

`func (o *PowerSchedule) SetWeekdayStartTimeMinutes(v int32)`

SetWeekdayStartTimeMinutes sets WeekdayStartTimeMinutes field to given value.

### HasWeekdayStartTimeMinutes

`func (o *PowerSchedule) HasWeekdayStartTimeMinutes() bool`

HasWeekdayStartTimeMinutes returns a boolean if a field has been set.

### GetWeekendEndTimeHour

`func (o *PowerSchedule) GetWeekendEndTimeHour() int32`

GetWeekendEndTimeHour returns the WeekendEndTimeHour field if non-nil, zero value otherwise.

### GetWeekendEndTimeHourOk

`func (o *PowerSchedule) GetWeekendEndTimeHourOk() (*int32, bool)`

GetWeekendEndTimeHourOk returns a tuple with the WeekendEndTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekendEndTimeHour

`func (o *PowerSchedule) SetWeekendEndTimeHour(v int32)`

SetWeekendEndTimeHour sets WeekendEndTimeHour field to given value.

### HasWeekendEndTimeHour

`func (o *PowerSchedule) HasWeekendEndTimeHour() bool`

HasWeekendEndTimeHour returns a boolean if a field has been set.

### GetWeekendEndTimeMinutes

`func (o *PowerSchedule) GetWeekendEndTimeMinutes() int32`

GetWeekendEndTimeMinutes returns the WeekendEndTimeMinutes field if non-nil, zero value otherwise.

### GetWeekendEndTimeMinutesOk

`func (o *PowerSchedule) GetWeekendEndTimeMinutesOk() (*int32, bool)`

GetWeekendEndTimeMinutesOk returns a tuple with the WeekendEndTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekendEndTimeMinutes

`func (o *PowerSchedule) SetWeekendEndTimeMinutes(v int32)`

SetWeekendEndTimeMinutes sets WeekendEndTimeMinutes field to given value.

### HasWeekendEndTimeMinutes

`func (o *PowerSchedule) HasWeekendEndTimeMinutes() bool`

HasWeekendEndTimeMinutes returns a boolean if a field has been set.

### GetWeekendStartTimeHour

`func (o *PowerSchedule) GetWeekendStartTimeHour() int32`

GetWeekendStartTimeHour returns the WeekendStartTimeHour field if non-nil, zero value otherwise.

### GetWeekendStartTimeHourOk

`func (o *PowerSchedule) GetWeekendStartTimeHourOk() (*int32, bool)`

GetWeekendStartTimeHourOk returns a tuple with the WeekendStartTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekendStartTimeHour

`func (o *PowerSchedule) SetWeekendStartTimeHour(v int32)`

SetWeekendStartTimeHour sets WeekendStartTimeHour field to given value.

### HasWeekendStartTimeHour

`func (o *PowerSchedule) HasWeekendStartTimeHour() bool`

HasWeekendStartTimeHour returns a boolean if a field has been set.

### GetWeekendStartTimeMinutes

`func (o *PowerSchedule) GetWeekendStartTimeMinutes() int32`

GetWeekendStartTimeMinutes returns the WeekendStartTimeMinutes field if non-nil, zero value otherwise.

### GetWeekendStartTimeMinutesOk

`func (o *PowerSchedule) GetWeekendStartTimeMinutesOk() (*int32, bool)`

GetWeekendStartTimeMinutesOk returns a tuple with the WeekendStartTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekendStartTimeMinutes

`func (o *PowerSchedule) SetWeekendStartTimeMinutes(v int32)`

SetWeekendStartTimeMinutes sets WeekendStartTimeMinutes field to given value.

### HasWeekendStartTimeMinutes

`func (o *PowerSchedule) HasWeekendStartTimeMinutes() bool`

HasWeekendStartTimeMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


