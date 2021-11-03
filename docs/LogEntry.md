# LogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewLogEntry

`func NewLogEntry() *LogEntry`

NewLogEntry instantiates a new LogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEntryWithDefaults

`func NewLogEntryWithDefaults() *LogEntry`

NewLogEntryWithDefaults instantiates a new LogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *LogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProjectName

`func (o *LogEntry) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *LogEntry) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *LogEntry) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *LogEntry) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetSeverity

`func (o *LogEntry) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LogEntry) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LogEntry) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *LogEntry) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *LogEntry) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogEntry) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogEntry) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LogEntry) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTime

`func (o *LogEntry) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LogEntry) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LogEntry) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *LogEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUsername

`func (o *LogEntry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LogEntry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LogEntry) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LogEntry) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


