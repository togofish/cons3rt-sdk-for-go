# DeploymentRunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Debug** | Pointer to **bool** |  | [optional] 
**DeploymentRunName** | Pointer to **string** |  | [optional] 
**DeploymentToSubmit** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndState** | Pointer to **string** |  | [optional] 
**EarliestStartTime** | Pointer to **int32** |  | [optional] 
**EndExisting** | Pointer to **bool** |  | [optional] 
**HostOptions** | Pointer to [**[]HostOption**](HostOption.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Password** | **string** |  | 
**PowerSchedule** | Pointer to [**PowerSchedule**](PowerSchedule.md) |  | [optional] 
**VirtualizationRealmId** | Pointer to **int32** |  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**QuickBuildCleanupOverridden** | Pointer to **bool** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**RetainOnError** | Pointer to **bool** |  | [optional] 
**TaskGroup** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 
**VirtRealmBinding** | Pointer to [**InputVirtualizationRealmBinding**](InputVirtualizationRealmBinding.md) |  | [optional] 
**WindowsDomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentRunOptions

`func NewDeploymentRunOptions(password string, username string, ) *DeploymentRunOptions`

NewDeploymentRunOptions instantiates a new DeploymentRunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRunOptionsWithDefaults

`func NewDeploymentRunOptionsWithDefaults() *DeploymentRunOptions`

NewDeploymentRunOptionsWithDefaults instantiates a new DeploymentRunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebug

`func (o *DeploymentRunOptions) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *DeploymentRunOptions) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *DeploymentRunOptions) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *DeploymentRunOptions) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDeploymentRunName

`func (o *DeploymentRunOptions) GetDeploymentRunName() string`

GetDeploymentRunName returns the DeploymentRunName field if non-nil, zero value otherwise.

### GetDeploymentRunNameOk

`func (o *DeploymentRunOptions) GetDeploymentRunNameOk() (*string, bool)`

GetDeploymentRunNameOk returns a tuple with the DeploymentRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunName

`func (o *DeploymentRunOptions) SetDeploymentRunName(v string)`

SetDeploymentRunName sets DeploymentRunName field to given value.

### HasDeploymentRunName

`func (o *DeploymentRunOptions) HasDeploymentRunName() bool`

HasDeploymentRunName returns a boolean if a field has been set.

### GetDeploymentToSubmit

`func (o *DeploymentRunOptions) GetDeploymentToSubmit() int32`

GetDeploymentToSubmit returns the DeploymentToSubmit field if non-nil, zero value otherwise.

### GetDeploymentToSubmitOk

`func (o *DeploymentRunOptions) GetDeploymentToSubmitOk() (*int32, bool)`

GetDeploymentToSubmitOk returns a tuple with the DeploymentToSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentToSubmit

`func (o *DeploymentRunOptions) SetDeploymentToSubmit(v int32)`

SetDeploymentToSubmit sets DeploymentToSubmit field to given value.

### HasDeploymentToSubmit

`func (o *DeploymentRunOptions) HasDeploymentToSubmit() bool`

HasDeploymentToSubmit returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentRunOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentRunOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentRunOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentRunOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndState

`func (o *DeploymentRunOptions) GetEndState() string`

GetEndState returns the EndState field if non-nil, zero value otherwise.

### GetEndStateOk

`func (o *DeploymentRunOptions) GetEndStateOk() (*string, bool)`

GetEndStateOk returns a tuple with the EndState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndState

`func (o *DeploymentRunOptions) SetEndState(v string)`

SetEndState sets EndState field to given value.

### HasEndState

`func (o *DeploymentRunOptions) HasEndState() bool`

HasEndState returns a boolean if a field has been set.

### GetEarliestStartTime

`func (o *DeploymentRunOptions) GetEarliestStartTime() int32`

GetEarliestStartTime returns the EarliestStartTime field if non-nil, zero value otherwise.

### GetEarliestStartTimeOk

`func (o *DeploymentRunOptions) GetEarliestStartTimeOk() (*int32, bool)`

GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStartTime

`func (o *DeploymentRunOptions) SetEarliestStartTime(v int32)`

SetEarliestStartTime sets EarliestStartTime field to given value.

### HasEarliestStartTime

`func (o *DeploymentRunOptions) HasEarliestStartTime() bool`

HasEarliestStartTime returns a boolean if a field has been set.

### GetEndExisting

`func (o *DeploymentRunOptions) GetEndExisting() bool`

GetEndExisting returns the EndExisting field if non-nil, zero value otherwise.

### GetEndExistingOk

`func (o *DeploymentRunOptions) GetEndExistingOk() (*bool, bool)`

GetEndExistingOk returns a tuple with the EndExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndExisting

`func (o *DeploymentRunOptions) SetEndExisting(v bool)`

SetEndExisting sets EndExisting field to given value.

### HasEndExisting

`func (o *DeploymentRunOptions) HasEndExisting() bool`

HasEndExisting returns a boolean if a field has been set.

### GetHostOptions

`func (o *DeploymentRunOptions) GetHostOptions() []HostOption`

GetHostOptions returns the HostOptions field if non-nil, zero value otherwise.

### GetHostOptionsOk

`func (o *DeploymentRunOptions) GetHostOptionsOk() (*[]HostOption, bool)`

GetHostOptionsOk returns a tuple with the HostOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOptions

`func (o *DeploymentRunOptions) SetHostOptions(v []HostOption)`

SetHostOptions sets HostOptions field to given value.

### HasHostOptions

`func (o *DeploymentRunOptions) HasHostOptions() bool`

HasHostOptions returns a boolean if a field has been set.

### GetId

`func (o *DeploymentRunOptions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentRunOptions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentRunOptions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentRunOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *DeploymentRunOptions) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DeploymentRunOptions) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DeploymentRunOptions) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *DeploymentRunOptions) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPassword

`func (o *DeploymentRunOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeploymentRunOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeploymentRunOptions) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPowerSchedule

`func (o *DeploymentRunOptions) GetPowerSchedule() PowerSchedule`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *DeploymentRunOptions) GetPowerScheduleOk() (*PowerSchedule, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *DeploymentRunOptions) SetPowerSchedule(v PowerSchedule)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *DeploymentRunOptions) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetVirtualizationRealmId

`func (o *DeploymentRunOptions) GetVirtualizationRealmId() int32`

GetVirtualizationRealmId returns the VirtualizationRealmId field if non-nil, zero value otherwise.

### GetVirtualizationRealmIdOk

`func (o *DeploymentRunOptions) GetVirtualizationRealmIdOk() (*int32, bool)`

GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmId

`func (o *DeploymentRunOptions) SetVirtualizationRealmId(v int32)`

SetVirtualizationRealmId sets VirtualizationRealmId field to given value.

### HasVirtualizationRealmId

`func (o *DeploymentRunOptions) HasVirtualizationRealmId() bool`

HasVirtualizationRealmId returns a boolean if a field has been set.

### GetProperties

`func (o *DeploymentRunOptions) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeploymentRunOptions) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeploymentRunOptions) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeploymentRunOptions) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuickBuildCleanupOverridden

`func (o *DeploymentRunOptions) GetQuickBuildCleanupOverridden() bool`

GetQuickBuildCleanupOverridden returns the QuickBuildCleanupOverridden field if non-nil, zero value otherwise.

### GetQuickBuildCleanupOverriddenOk

`func (o *DeploymentRunOptions) GetQuickBuildCleanupOverriddenOk() (*bool, bool)`

GetQuickBuildCleanupOverriddenOk returns a tuple with the QuickBuildCleanupOverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBuildCleanupOverridden

`func (o *DeploymentRunOptions) SetQuickBuildCleanupOverridden(v bool)`

SetQuickBuildCleanupOverridden sets QuickBuildCleanupOverridden field to given value.

### HasQuickBuildCleanupOverridden

`func (o *DeploymentRunOptions) HasQuickBuildCleanupOverridden() bool`

HasQuickBuildCleanupOverridden returns a boolean if a field has been set.

### GetDuration

`func (o *DeploymentRunOptions) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeploymentRunOptions) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeploymentRunOptions) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DeploymentRunOptions) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndDate

`func (o *DeploymentRunOptions) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DeploymentRunOptions) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DeploymentRunOptions) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DeploymentRunOptions) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRetainOnError

`func (o *DeploymentRunOptions) GetRetainOnError() bool`

GetRetainOnError returns the RetainOnError field if non-nil, zero value otherwise.

### GetRetainOnErrorOk

`func (o *DeploymentRunOptions) GetRetainOnErrorOk() (*bool, bool)`

GetRetainOnErrorOk returns a tuple with the RetainOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainOnError

`func (o *DeploymentRunOptions) SetRetainOnError(v bool)`

SetRetainOnError sets RetainOnError field to given value.

### HasRetainOnError

`func (o *DeploymentRunOptions) HasRetainOnError() bool`

HasRetainOnError returns a boolean if a field has been set.

### GetTaskGroup

`func (o *DeploymentRunOptions) GetTaskGroup() string`

GetTaskGroup returns the TaskGroup field if non-nil, zero value otherwise.

### GetTaskGroupOk

`func (o *DeploymentRunOptions) GetTaskGroupOk() (*string, bool)`

GetTaskGroupOk returns a tuple with the TaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroup

`func (o *DeploymentRunOptions) SetTaskGroup(v string)`

SetTaskGroup sets TaskGroup field to given value.

### HasTaskGroup

`func (o *DeploymentRunOptions) HasTaskGroup() bool`

HasTaskGroup returns a boolean if a field has been set.

### GetUsername

`func (o *DeploymentRunOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeploymentRunOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeploymentRunOptions) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVirtRealmBinding

`func (o *DeploymentRunOptions) GetVirtRealmBinding() InputVirtualizationRealmBinding`

GetVirtRealmBinding returns the VirtRealmBinding field if non-nil, zero value otherwise.

### GetVirtRealmBindingOk

`func (o *DeploymentRunOptions) GetVirtRealmBindingOk() (*InputVirtualizationRealmBinding, bool)`

GetVirtRealmBindingOk returns a tuple with the VirtRealmBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtRealmBinding

`func (o *DeploymentRunOptions) SetVirtRealmBinding(v InputVirtualizationRealmBinding)`

SetVirtRealmBinding sets VirtRealmBinding field to given value.

### HasVirtRealmBinding

`func (o *DeploymentRunOptions) HasVirtRealmBinding() bool`

HasVirtRealmBinding returns a boolean if a field has been set.

### GetWindowsDomainName

`func (o *DeploymentRunOptions) GetWindowsDomainName() string`

GetWindowsDomainName returns the WindowsDomainName field if non-nil, zero value otherwise.

### GetWindowsDomainNameOk

`func (o *DeploymentRunOptions) GetWindowsDomainNameOk() (*string, bool)`

GetWindowsDomainNameOk returns a tuple with the WindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomainName

`func (o *DeploymentRunOptions) SetWindowsDomainName(v string)`

SetWindowsDomainName sets WindowsDomainName field to given value.

### HasWindowsDomainName

`func (o *DeploymentRunOptions) HasWindowsDomainName() bool`

HasWindowsDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


