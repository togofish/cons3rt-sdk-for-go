# MinimalTestRunTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**MinimalProject**](MinimalProject.md) |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 
**ReportUri** | Pointer to **string** |  | [optional] 
**TestBundle** | Pointer to [**MinimalTestBundle**](MinimalTestBundle.md) |  | [optional] 
**TestManagerStatus** | Pointer to **string** |  | [optional] 
**TestResult** | Pointer to **string** |  | [optional] 

## Methods

### NewMinimalTestRunTask

`func NewMinimalTestRunTask() *MinimalTestRunTask`

NewMinimalTestRunTask instantiates a new MinimalTestRunTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalTestRunTaskWithDefaults

`func NewMinimalTestRunTaskWithDefaults() *MinimalTestRunTask`

NewMinimalTestRunTaskWithDefaults instantiates a new MinimalTestRunTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *MinimalTestRunTask) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MinimalTestRunTask) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MinimalTestRunTask) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MinimalTestRunTask) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetId

`func (o *MinimalTestRunTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalTestRunTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalTestRunTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalTestRunTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *MinimalTestRunTask) GetProject() MinimalProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MinimalTestRunTask) GetProjectOk() (*MinimalProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MinimalTestRunTask) SetProject(v MinimalProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *MinimalTestRunTask) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetResult

`func (o *MinimalTestRunTask) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MinimalTestRunTask) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MinimalTestRunTask) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *MinimalTestRunTask) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *MinimalTestRunTask) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MinimalTestRunTask) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MinimalTestRunTask) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MinimalTestRunTask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetReportUri

`func (o *MinimalTestRunTask) GetReportUri() string`

GetReportUri returns the ReportUri field if non-nil, zero value otherwise.

### GetReportUriOk

`func (o *MinimalTestRunTask) GetReportUriOk() (*string, bool)`

GetReportUriOk returns a tuple with the ReportUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUri

`func (o *MinimalTestRunTask) SetReportUri(v string)`

SetReportUri sets ReportUri field to given value.

### HasReportUri

`func (o *MinimalTestRunTask) HasReportUri() bool`

HasReportUri returns a boolean if a field has been set.

### GetTestBundle

`func (o *MinimalTestRunTask) GetTestBundle() MinimalTestBundle`

GetTestBundle returns the TestBundle field if non-nil, zero value otherwise.

### GetTestBundleOk

`func (o *MinimalTestRunTask) GetTestBundleOk() (*MinimalTestBundle, bool)`

GetTestBundleOk returns a tuple with the TestBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBundle

`func (o *MinimalTestRunTask) SetTestBundle(v MinimalTestBundle)`

SetTestBundle sets TestBundle field to given value.

### HasTestBundle

`func (o *MinimalTestRunTask) HasTestBundle() bool`

HasTestBundle returns a boolean if a field has been set.

### GetTestManagerStatus

`func (o *MinimalTestRunTask) GetTestManagerStatus() string`

GetTestManagerStatus returns the TestManagerStatus field if non-nil, zero value otherwise.

### GetTestManagerStatusOk

`func (o *MinimalTestRunTask) GetTestManagerStatusOk() (*string, bool)`

GetTestManagerStatusOk returns a tuple with the TestManagerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestManagerStatus

`func (o *MinimalTestRunTask) SetTestManagerStatus(v string)`

SetTestManagerStatus sets TestManagerStatus field to given value.

### HasTestManagerStatus

`func (o *MinimalTestRunTask) HasTestManagerStatus() bool`

HasTestManagerStatus returns a boolean if a field has been set.

### GetTestResult

`func (o *MinimalTestRunTask) GetTestResult() string`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *MinimalTestRunTask) GetTestResultOk() (*string, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *MinimalTestRunTask) SetTestResult(v string)`

SetTestResult sets TestResult field to given value.

### HasTestResult

`func (o *MinimalTestRunTask) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


