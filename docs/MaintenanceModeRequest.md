# MaintenanceModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewMaintenanceModeRequest

`func NewMaintenanceModeRequest() *MaintenanceModeRequest`

NewMaintenanceModeRequest instantiates a new MaintenanceModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceModeRequestWithDefaults

`func NewMaintenanceModeRequestWithDefaults() *MaintenanceModeRequest`

NewMaintenanceModeRequestWithDefaults instantiates a new MaintenanceModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *MaintenanceModeRequest) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MaintenanceModeRequest) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MaintenanceModeRequest) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MaintenanceModeRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMessage

`func (o *MaintenanceModeRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MaintenanceModeRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MaintenanceModeRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MaintenanceModeRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimeout

`func (o *MaintenanceModeRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MaintenanceModeRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MaintenanceModeRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MaintenanceModeRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


