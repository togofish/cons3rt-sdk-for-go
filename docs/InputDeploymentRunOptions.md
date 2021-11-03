# InputDeploymentRunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRunName** | Pointer to **string** |  | [optional] 
**EndState** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**VirtualizationRealmId** | Pointer to **int32** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**RetainOnError** | Pointer to **bool** |  | [optional] 
**Username** | **string** |  | 
**Password** | **string** |  | 
**EarliestStartTime** | Pointer to **int32** |  | [optional] 
**EndExisting** | Pointer to **bool** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Properties** | Pointer to [**[]InputProperty**](InputProperty.md) |  | [optional] 
**HostOptions** | Pointer to [**[]InputHostOption**](InputHostOption.md) |  | [optional] 
**PowerSchedule** | Pointer to [**PowerSchedule**](PowerSchedule.md) |  | [optional] 
**VirtRealmBinding** | Pointer to [**InputVirtualizationRealmBinding**](InputVirtualizationRealmBinding.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**WindowsDomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewInputDeploymentRunOptions

`func NewInputDeploymentRunOptions(endState string, username string, password string, ) *InputDeploymentRunOptions`

NewInputDeploymentRunOptions instantiates a new InputDeploymentRunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDeploymentRunOptionsWithDefaults

`func NewInputDeploymentRunOptionsWithDefaults() *InputDeploymentRunOptions`

NewInputDeploymentRunOptionsWithDefaults instantiates a new InputDeploymentRunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRunName

`func (o *InputDeploymentRunOptions) GetDeploymentRunName() string`

GetDeploymentRunName returns the DeploymentRunName field if non-nil, zero value otherwise.

### GetDeploymentRunNameOk

`func (o *InputDeploymentRunOptions) GetDeploymentRunNameOk() (*string, bool)`

GetDeploymentRunNameOk returns a tuple with the DeploymentRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunName

`func (o *InputDeploymentRunOptions) SetDeploymentRunName(v string)`

SetDeploymentRunName sets DeploymentRunName field to given value.

### HasDeploymentRunName

`func (o *InputDeploymentRunOptions) HasDeploymentRunName() bool`

HasDeploymentRunName returns a boolean if a field has been set.

### GetEndState

`func (o *InputDeploymentRunOptions) GetEndState() string`

GetEndState returns the EndState field if non-nil, zero value otherwise.

### GetEndStateOk

`func (o *InputDeploymentRunOptions) GetEndStateOk() (*string, bool)`

GetEndStateOk returns a tuple with the EndState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndState

`func (o *InputDeploymentRunOptions) SetEndState(v string)`

SetEndState sets EndState field to given value.


### GetDescription

`func (o *InputDeploymentRunOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputDeploymentRunOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputDeploymentRunOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputDeploymentRunOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVirtualizationRealmId

`func (o *InputDeploymentRunOptions) GetVirtualizationRealmId() int32`

GetVirtualizationRealmId returns the VirtualizationRealmId field if non-nil, zero value otherwise.

### GetVirtualizationRealmIdOk

`func (o *InputDeploymentRunOptions) GetVirtualizationRealmIdOk() (*int32, bool)`

GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmId

`func (o *InputDeploymentRunOptions) SetVirtualizationRealmId(v int32)`

SetVirtualizationRealmId sets VirtualizationRealmId field to given value.

### HasVirtualizationRealmId

`func (o *InputDeploymentRunOptions) HasVirtualizationRealmId() bool`

HasVirtualizationRealmId returns a boolean if a field has been set.

### GetLocked

`func (o *InputDeploymentRunOptions) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InputDeploymentRunOptions) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InputDeploymentRunOptions) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InputDeploymentRunOptions) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetRetainOnError

`func (o *InputDeploymentRunOptions) GetRetainOnError() bool`

GetRetainOnError returns the RetainOnError field if non-nil, zero value otherwise.

### GetRetainOnErrorOk

`func (o *InputDeploymentRunOptions) GetRetainOnErrorOk() (*bool, bool)`

GetRetainOnErrorOk returns a tuple with the RetainOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainOnError

`func (o *InputDeploymentRunOptions) SetRetainOnError(v bool)`

SetRetainOnError sets RetainOnError field to given value.

### HasRetainOnError

`func (o *InputDeploymentRunOptions) HasRetainOnError() bool`

HasRetainOnError returns a boolean if a field has been set.

### GetUsername

`func (o *InputDeploymentRunOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InputDeploymentRunOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InputDeploymentRunOptions) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *InputDeploymentRunOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputDeploymentRunOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputDeploymentRunOptions) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEarliestStartTime

`func (o *InputDeploymentRunOptions) GetEarliestStartTime() int32`

GetEarliestStartTime returns the EarliestStartTime field if non-nil, zero value otherwise.

### GetEarliestStartTimeOk

`func (o *InputDeploymentRunOptions) GetEarliestStartTimeOk() (*int32, bool)`

GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStartTime

`func (o *InputDeploymentRunOptions) SetEarliestStartTime(v int32)`

SetEarliestStartTime sets EarliestStartTime field to given value.

### HasEarliestStartTime

`func (o *InputDeploymentRunOptions) HasEarliestStartTime() bool`

HasEarliestStartTime returns a boolean if a field has been set.

### GetEndExisting

`func (o *InputDeploymentRunOptions) GetEndExisting() bool`

GetEndExisting returns the EndExisting field if non-nil, zero value otherwise.

### GetEndExistingOk

`func (o *InputDeploymentRunOptions) GetEndExistingOk() (*bool, bool)`

GetEndExistingOk returns a tuple with the EndExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndExisting

`func (o *InputDeploymentRunOptions) SetEndExisting(v bool)`

SetEndExisting sets EndExisting field to given value.

### HasEndExisting

`func (o *InputDeploymentRunOptions) HasEndExisting() bool`

HasEndExisting returns a boolean if a field has been set.

### GetDuration

`func (o *InputDeploymentRunOptions) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InputDeploymentRunOptions) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InputDeploymentRunOptions) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InputDeploymentRunOptions) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetProperties

`func (o *InputDeploymentRunOptions) GetProperties() []InputProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InputDeploymentRunOptions) GetPropertiesOk() (*[]InputProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InputDeploymentRunOptions) SetProperties(v []InputProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InputDeploymentRunOptions) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetHostOptions

`func (o *InputDeploymentRunOptions) GetHostOptions() []InputHostOption`

GetHostOptions returns the HostOptions field if non-nil, zero value otherwise.

### GetHostOptionsOk

`func (o *InputDeploymentRunOptions) GetHostOptionsOk() (*[]InputHostOption, bool)`

GetHostOptionsOk returns a tuple with the HostOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOptions

`func (o *InputDeploymentRunOptions) SetHostOptions(v []InputHostOption)`

SetHostOptions sets HostOptions field to given value.

### HasHostOptions

`func (o *InputDeploymentRunOptions) HasHostOptions() bool`

HasHostOptions returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *InputDeploymentRunOptions) GetPowerSchedule() PowerSchedule`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *InputDeploymentRunOptions) GetPowerScheduleOk() (*PowerSchedule, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *InputDeploymentRunOptions) SetPowerSchedule(v PowerSchedule)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *InputDeploymentRunOptions) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetVirtRealmBinding

`func (o *InputDeploymentRunOptions) GetVirtRealmBinding() InputVirtualizationRealmBinding`

GetVirtRealmBinding returns the VirtRealmBinding field if non-nil, zero value otherwise.

### GetVirtRealmBindingOk

`func (o *InputDeploymentRunOptions) GetVirtRealmBindingOk() (*InputVirtualizationRealmBinding, bool)`

GetVirtRealmBindingOk returns a tuple with the VirtRealmBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmBinding

`func (o *InputDeploymentRunOptions) SetVirtRealmBinding(v InputVirtualizationRealmBinding)`

SetVirtRealmBinding sets VirtRealmBinding field to given value.

### HasVirtRealmBinding

`func (o *InputDeploymentRunOptions) HasVirtRealmBinding() bool`

HasVirtRealmBinding returns a boolean if a field has been set.

### GetId

`func (o *InputDeploymentRunOptions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputDeploymentRunOptions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputDeploymentRunOptions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputDeploymentRunOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDebug

`func (o *InputDeploymentRunOptions) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *InputDeploymentRunOptions) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *InputDeploymentRunOptions) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *InputDeploymentRunOptions) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetEndDate

`func (o *InputDeploymentRunOptions) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputDeploymentRunOptions) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputDeploymentRunOptions) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputDeploymentRunOptions) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetWindowsDomainName

`func (o *InputDeploymentRunOptions) GetWindowsDomainName() string`

GetWindowsDomainName returns the WindowsDomainName field if non-nil, zero value otherwise.

### GetWindowsDomainNameOk

`func (o *InputDeploymentRunOptions) GetWindowsDomainNameOk() (*string, bool)`

GetWindowsDomainNameOk returns a tuple with the WindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomainName

`func (o *InputDeploymentRunOptions) SetWindowsDomainName(v string)`

SetWindowsDomainName sets WindowsDomainName field to given value.

### HasWindowsDomainName

`func (o *InputDeploymentRunOptions) HasWindowsDomainName() bool`

HasWindowsDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


