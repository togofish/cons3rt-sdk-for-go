# FullDeploymentRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]MinimalCategory**](MinimalCategory.md) |  | [optional] 
**Creator** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 
**EarliestStartTime** | Pointer to **int32** |  | [optional] 
**EndTime** | Pointer to **int32** |  | [optional] 
**LeaseTime** | Pointer to **int32** |  | [optional] 
**EstimatedReadyTime** | Pointer to **int32** |  | [optional] 
**EstimatedStartTime** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LogEntries** | Pointer to [**[]MinimalLogEntry**](MinimalLogEntry.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**MinimalProject**](MinimalProject.md) |  | [optional] 
**ReadyTime** | Pointer to **int32** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 
**TimeOfRequest** | Pointer to **int32** |  | [optional] 
**Canceled** | Pointer to **bool** |  | [optional] 
**Deployment** | Pointer to [**MinimalDeployment**](MinimalDeployment.md) |  | [optional] 
**DeploymentRunHosts** | Pointer to [**[]MinimalDeploymentRunHost**](MinimalDeploymentRunHost.md) |  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**DeploymentRunStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FapStatus** | Pointer to **string** |  | [optional] 
**HostSetName** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PowerSchedule** | Pointer to [**PowerSchedule**](PowerSchedule.md) |  | [optional] 
**RecurringSchedule** | Pointer to [**MinimalRecurringSchedule**](MinimalRecurringSchedule.md) |  | [optional] 
**SchedulerStatusMessage** | Pointer to **string** |  | [optional] 
**TargetState** | Pointer to **string** |  | [optional] 
**TestError** | Pointer to **bool** |  | [optional] 
**TestRuns** | Pointer to [**[]MinimalTestRunTask**](MinimalTestRunTask.md) |  | [optional] 
**RetainedOnError** | Pointer to **bool** |  | [optional] 
**VirtualizationRealm** | Pointer to [**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**DeploymentRunResultType** | Pointer to **string** |  | [optional] 

## Methods

### NewFullDeploymentRun

`func NewFullDeploymentRun() *FullDeploymentRun`

NewFullDeploymentRun instantiates a new FullDeploymentRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDeploymentRunWithDefaults

`func NewFullDeploymentRunWithDefaults() *FullDeploymentRun`

NewFullDeploymentRunWithDefaults instantiates a new FullDeploymentRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *FullDeploymentRun) GetCategories() []MinimalCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FullDeploymentRun) GetCategoriesOk() (*[]MinimalCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FullDeploymentRun) SetCategories(v []MinimalCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *FullDeploymentRun) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreator

`func (o *FullDeploymentRun) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FullDeploymentRun) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FullDeploymentRun) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *FullDeploymentRun) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEarliestStartTime

`func (o *FullDeploymentRun) GetEarliestStartTime() int32`

GetEarliestStartTime returns the EarliestStartTime field if non-nil, zero value otherwise.

### GetEarliestStartTimeOk

`func (o *FullDeploymentRun) GetEarliestStartTimeOk() (*int32, bool)`

GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStartTime

`func (o *FullDeploymentRun) SetEarliestStartTime(v int32)`

SetEarliestStartTime sets EarliestStartTime field to given value.

### HasEarliestStartTime

`func (o *FullDeploymentRun) HasEarliestStartTime() bool`

HasEarliestStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *FullDeploymentRun) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FullDeploymentRun) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FullDeploymentRun) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *FullDeploymentRun) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLeaseTime

`func (o *FullDeploymentRun) GetLeaseTime() int32`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *FullDeploymentRun) GetLeaseTimeOk() (*int32, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *FullDeploymentRun) SetLeaseTime(v int32)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *FullDeploymentRun) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetEstimatedReadyTime

`func (o *FullDeploymentRun) GetEstimatedReadyTime() int32`

GetEstimatedReadyTime returns the EstimatedReadyTime field if non-nil, zero value otherwise.

### GetEstimatedReadyTimeOk

`func (o *FullDeploymentRun) GetEstimatedReadyTimeOk() (*int32, bool)`

GetEstimatedReadyTimeOk returns a tuple with the EstimatedReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedReadyTime

`func (o *FullDeploymentRun) SetEstimatedReadyTime(v int32)`

SetEstimatedReadyTime sets EstimatedReadyTime field to given value.

### HasEstimatedReadyTime

`func (o *FullDeploymentRun) HasEstimatedReadyTime() bool`

HasEstimatedReadyTime returns a boolean if a field has been set.

### GetEstimatedStartTime

`func (o *FullDeploymentRun) GetEstimatedStartTime() int32`

GetEstimatedStartTime returns the EstimatedStartTime field if non-nil, zero value otherwise.

### GetEstimatedStartTimeOk

`func (o *FullDeploymentRun) GetEstimatedStartTimeOk() (*int32, bool)`

GetEstimatedStartTimeOk returns a tuple with the EstimatedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartTime

`func (o *FullDeploymentRun) SetEstimatedStartTime(v int32)`

SetEstimatedStartTime sets EstimatedStartTime field to given value.

### HasEstimatedStartTime

`func (o *FullDeploymentRun) HasEstimatedStartTime() bool`

HasEstimatedStartTime returns a boolean if a field has been set.

### GetId

`func (o *FullDeploymentRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullDeploymentRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullDeploymentRun) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullDeploymentRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogEntries

`func (o *FullDeploymentRun) GetLogEntries() []MinimalLogEntry`

GetLogEntries returns the LogEntries field if non-nil, zero value otherwise.

### GetLogEntriesOk

`func (o *FullDeploymentRun) GetLogEntriesOk() (*[]MinimalLogEntry, bool)`

GetLogEntriesOk returns a tuple with the LogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEntries

`func (o *FullDeploymentRun) SetLogEntries(v []MinimalLogEntry)`

SetLogEntries sets LogEntries field to given value.

### HasLogEntries

`func (o *FullDeploymentRun) HasLogEntries() bool`

HasLogEntries returns a boolean if a field has been set.

### GetMessage

`func (o *FullDeploymentRun) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FullDeploymentRun) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FullDeploymentRun) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FullDeploymentRun) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProject

`func (o *FullDeploymentRun) GetProject() MinimalProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FullDeploymentRun) GetProjectOk() (*MinimalProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FullDeploymentRun) SetProject(v MinimalProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *FullDeploymentRun) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReadyTime

`func (o *FullDeploymentRun) GetReadyTime() int32`

GetReadyTime returns the ReadyTime field if non-nil, zero value otherwise.

### GetReadyTimeOk

`func (o *FullDeploymentRun) GetReadyTimeOk() (*int32, bool)`

GetReadyTimeOk returns a tuple with the ReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyTime

`func (o *FullDeploymentRun) SetReadyTime(v int32)`

SetReadyTime sets ReadyTime field to given value.

### HasReadyTime

`func (o *FullDeploymentRun) HasReadyTime() bool`

HasReadyTime returns a boolean if a field has been set.

### GetResult

`func (o *FullDeploymentRun) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FullDeploymentRun) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FullDeploymentRun) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *FullDeploymentRun) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *FullDeploymentRun) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FullDeploymentRun) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FullDeploymentRun) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FullDeploymentRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimeOfRequest

`func (o *FullDeploymentRun) GetTimeOfRequest() int32`

GetTimeOfRequest returns the TimeOfRequest field if non-nil, zero value otherwise.

### GetTimeOfRequestOk

`func (o *FullDeploymentRun) GetTimeOfRequestOk() (*int32, bool)`

GetTimeOfRequestOk returns a tuple with the TimeOfRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfRequest

`func (o *FullDeploymentRun) SetTimeOfRequest(v int32)`

SetTimeOfRequest sets TimeOfRequest field to given value.

### HasTimeOfRequest

`func (o *FullDeploymentRun) HasTimeOfRequest() bool`

HasTimeOfRequest returns a boolean if a field has been set.

### GetCanceled

`func (o *FullDeploymentRun) GetCanceled() bool`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *FullDeploymentRun) GetCanceledOk() (*bool, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *FullDeploymentRun) SetCanceled(v bool)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *FullDeploymentRun) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetDeployment

`func (o *FullDeploymentRun) GetDeployment() MinimalDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *FullDeploymentRun) GetDeploymentOk() (*MinimalDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *FullDeploymentRun) SetDeployment(v MinimalDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *FullDeploymentRun) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeploymentRunHosts

`func (o *FullDeploymentRun) GetDeploymentRunHosts() []MinimalDeploymentRunHost`

GetDeploymentRunHosts returns the DeploymentRunHosts field if non-nil, zero value otherwise.

### GetDeploymentRunHostsOk

`func (o *FullDeploymentRun) GetDeploymentRunHostsOk() (*[]MinimalDeploymentRunHost, bool)`

GetDeploymentRunHostsOk returns a tuple with the DeploymentRunHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunHosts

`func (o *FullDeploymentRun) SetDeploymentRunHosts(v []MinimalDeploymentRunHost)`

SetDeploymentRunHosts sets DeploymentRunHosts field to given value.

### HasDeploymentRunHosts

`func (o *FullDeploymentRun) HasDeploymentRunHosts() bool`

HasDeploymentRunHosts returns a boolean if a field has been set.

### GetProperties

`func (o *FullDeploymentRun) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FullDeploymentRun) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FullDeploymentRun) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FullDeploymentRun) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDeploymentRunStatus

`func (o *FullDeploymentRun) GetDeploymentRunStatus() string`

GetDeploymentRunStatus returns the DeploymentRunStatus field if non-nil, zero value otherwise.

### GetDeploymentRunStatusOk

`func (o *FullDeploymentRun) GetDeploymentRunStatusOk() (*string, bool)`

GetDeploymentRunStatusOk returns a tuple with the DeploymentRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunStatus

`func (o *FullDeploymentRun) SetDeploymentRunStatus(v string)`

SetDeploymentRunStatus sets DeploymentRunStatus field to given value.

### HasDeploymentRunStatus

`func (o *FullDeploymentRun) HasDeploymentRunStatus() bool`

HasDeploymentRunStatus returns a boolean if a field has been set.

### GetDescription

`func (o *FullDeploymentRun) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullDeploymentRun) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullDeploymentRun) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullDeploymentRun) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFapStatus

`func (o *FullDeploymentRun) GetFapStatus() string`

GetFapStatus returns the FapStatus field if non-nil, zero value otherwise.

### GetFapStatusOk

`func (o *FullDeploymentRun) GetFapStatusOk() (*string, bool)`

GetFapStatusOk returns a tuple with the FapStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFapStatus

`func (o *FullDeploymentRun) SetFapStatus(v string)`

SetFapStatus sets FapStatus field to given value.

### HasFapStatus

`func (o *FullDeploymentRun) HasFapStatus() bool`

HasFapStatus returns a boolean if a field has been set.

### GetHostSetName

`func (o *FullDeploymentRun) GetHostSetName() string`

GetHostSetName returns the HostSetName field if non-nil, zero value otherwise.

### GetHostSetNameOk

`func (o *FullDeploymentRun) GetHostSetNameOk() (*string, bool)`

GetHostSetNameOk returns a tuple with the HostSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSetName

`func (o *FullDeploymentRun) SetHostSetName(v string)`

SetHostSetName sets HostSetName field to given value.

### HasHostSetName

`func (o *FullDeploymentRun) HasHostSetName() bool`

HasHostSetName returns a boolean if a field has been set.

### GetLocked

`func (o *FullDeploymentRun) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *FullDeploymentRun) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *FullDeploymentRun) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *FullDeploymentRun) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *FullDeploymentRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullDeploymentRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullDeploymentRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullDeploymentRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *FullDeploymentRun) GetPowerSchedule() PowerSchedule`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *FullDeploymentRun) GetPowerScheduleOk() (*PowerSchedule, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *FullDeploymentRun) SetPowerSchedule(v PowerSchedule)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *FullDeploymentRun) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetRecurringSchedule

`func (o *FullDeploymentRun) GetRecurringSchedule() MinimalRecurringSchedule`

GetRecurringSchedule returns the RecurringSchedule field if non-nil, zero value otherwise.

### GetRecurringScheduleOk

`func (o *FullDeploymentRun) GetRecurringScheduleOk() (*MinimalRecurringSchedule, bool)`

GetRecurringScheduleOk returns a tuple with the RecurringSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSchedule

`func (o *FullDeploymentRun) SetRecurringSchedule(v MinimalRecurringSchedule)`

SetRecurringSchedule sets RecurringSchedule field to given value.

### HasRecurringSchedule

`func (o *FullDeploymentRun) HasRecurringSchedule() bool`

HasRecurringSchedule returns a boolean if a field has been set.

### GetSchedulerStatusMessage

`func (o *FullDeploymentRun) GetSchedulerStatusMessage() string`

GetSchedulerStatusMessage returns the SchedulerStatusMessage field if non-nil, zero value otherwise.

### GetSchedulerStatusMessageOk

`func (o *FullDeploymentRun) GetSchedulerStatusMessageOk() (*string, bool)`

GetSchedulerStatusMessageOk returns a tuple with the SchedulerStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerStatusMessage

`func (o *FullDeploymentRun) SetSchedulerStatusMessage(v string)`

SetSchedulerStatusMessage sets SchedulerStatusMessage field to given value.

### HasSchedulerStatusMessage

`func (o *FullDeploymentRun) HasSchedulerStatusMessage() bool`

HasSchedulerStatusMessage returns a boolean if a field has been set.

### GetTargetState

`func (o *FullDeploymentRun) GetTargetState() string`

GetTargetState returns the TargetState field if non-nil, zero value otherwise.

### GetTargetStateOk

`func (o *FullDeploymentRun) GetTargetStateOk() (*string, bool)`

GetTargetStateOk returns a tuple with the TargetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetState

`func (o *FullDeploymentRun) SetTargetState(v string)`

SetTargetState sets TargetState field to given value.

### HasTargetState

`func (o *FullDeploymentRun) HasTargetState() bool`

HasTargetState returns a boolean if a field has been set.

### GetTestError

`func (o *FullDeploymentRun) GetTestError() bool`

GetTestError returns the TestError field if non-nil, zero value otherwise.

### GetTestErrorOk

`func (o *FullDeploymentRun) GetTestErrorOk() (*bool, bool)`

GetTestErrorOk returns a tuple with the TestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestError

`func (o *FullDeploymentRun) SetTestError(v bool)`

SetTestError sets TestError field to given value.

### HasTestError

`func (o *FullDeploymentRun) HasTestError() bool`

HasTestError returns a boolean if a field has been set.

### GetTestRuns

`func (o *FullDeploymentRun) GetTestRuns() []MinimalTestRunTask`

GetTestRuns returns the TestRuns field if non-nil, zero value otherwise.

### GetTestRunsOk

`func (o *FullDeploymentRun) GetTestRunsOk() (*[]MinimalTestRunTask, bool)`

GetTestRunsOk returns a tuple with the TestRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRuns

`func (o *FullDeploymentRun) SetTestRuns(v []MinimalTestRunTask)`

SetTestRuns sets TestRuns field to given value.

### HasTestRuns

`func (o *FullDeploymentRun) HasTestRuns() bool`

HasTestRuns returns a boolean if a field has been set.

### GetRetainedOnError

`func (o *FullDeploymentRun) GetRetainedOnError() bool`

GetRetainedOnError returns the RetainedOnError field if non-nil, zero value otherwise.

### GetRetainedOnErrorOk

`func (o *FullDeploymentRun) GetRetainedOnErrorOk() (*bool, bool)`

GetRetainedOnErrorOk returns a tuple with the RetainedOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedOnError

`func (o *FullDeploymentRun) SetRetainedOnError(v bool)`

SetRetainedOnError sets RetainedOnError field to given value.

### HasRetainedOnError

`func (o *FullDeploymentRun) HasRetainedOnError() bool`

HasRetainedOnError returns a boolean if a field has been set.

### GetVirtualizationRealm

`func (o *FullDeploymentRun) GetVirtualizationRealm() MinimalVirtualizationRealm`

GetVirtualizationRealm returns the VirtualizationRealm field if non-nil, zero value otherwise.

### GetVirtualizationRealmOk

`func (o *FullDeploymentRun) GetVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool)`

GetVirtualizationRealmOk returns a tuple with the VirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealm

`func (o *FullDeploymentRun) SetVirtualizationRealm(v MinimalVirtualizationRealm)`

SetVirtualizationRealm sets VirtualizationRealm field to given value.

### HasVirtualizationRealm

`func (o *FullDeploymentRun) HasVirtualizationRealm() bool`

HasVirtualizationRealm returns a boolean if a field has been set.

### GetDeploymentRunResultType

`func (o *FullDeploymentRun) GetDeploymentRunResultType() string`

GetDeploymentRunResultType returns the DeploymentRunResultType field if non-nil, zero value otherwise.

### GetDeploymentRunResultTypeOk

`func (o *FullDeploymentRun) GetDeploymentRunResultTypeOk() (*string, bool)`

GetDeploymentRunResultTypeOk returns a tuple with the DeploymentRunResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunResultType

`func (o *FullDeploymentRun) SetDeploymentRunResultType(v string)`

SetDeploymentRunResultType sets DeploymentRunResultType field to given value.

### HasDeploymentRunResultType

`func (o *FullDeploymentRun) HasDeploymentRunResultType() bool`

HasDeploymentRunResultType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


