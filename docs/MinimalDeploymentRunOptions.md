# MinimalDeploymentRunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Debug** | Pointer to **bool** |  | [optional] 
**DeploymentRunName** | Pointer to **string** |  | [optional] 
**DeploymentToSubmit** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndState** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**EarliestStartTime** | Pointer to **int32** |  | [optional] 
**EndExisting** | Pointer to **bool** |  | [optional] 
**HostOptions** | Pointer to [**[]HostOption**](HostOption.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**VirtualizationRealmId** | Pointer to **int32** |  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**RetainOnError** | Pointer to **bool** |  | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 

## Methods

### NewMinimalDeploymentRunOptions

`func NewMinimalDeploymentRunOptions() *MinimalDeploymentRunOptions`

NewMinimalDeploymentRunOptions instantiates a new MinimalDeploymentRunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalDeploymentRunOptionsWithDefaults

`func NewMinimalDeploymentRunOptionsWithDefaults() *MinimalDeploymentRunOptions`

NewMinimalDeploymentRunOptionsWithDefaults instantiates a new MinimalDeploymentRunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebug

`func (o *MinimalDeploymentRunOptions) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *MinimalDeploymentRunOptions) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *MinimalDeploymentRunOptions) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *MinimalDeploymentRunOptions) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDeploymentRunName

`func (o *MinimalDeploymentRunOptions) GetDeploymentRunName() string`

GetDeploymentRunName returns the DeploymentRunName field if non-nil, zero value otherwise.

### GetDeploymentRunNameOk

`func (o *MinimalDeploymentRunOptions) GetDeploymentRunNameOk() (*string, bool)`

GetDeploymentRunNameOk returns a tuple with the DeploymentRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRunName

`func (o *MinimalDeploymentRunOptions) SetDeploymentRunName(v string)`

SetDeploymentRunName sets DeploymentRunName field to given value.

### HasDeploymentRunName

`func (o *MinimalDeploymentRunOptions) HasDeploymentRunName() bool`

HasDeploymentRunName returns a boolean if a field has been set.

### GetDeploymentToSubmit

`func (o *MinimalDeploymentRunOptions) GetDeploymentToSubmit() int32`

GetDeploymentToSubmit returns the DeploymentToSubmit field if non-nil, zero value otherwise.

### GetDeploymentToSubmitOk

`func (o *MinimalDeploymentRunOptions) GetDeploymentToSubmitOk() (*int32, bool)`

GetDeploymentToSubmitOk returns a tuple with the DeploymentToSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentToSubmit

`func (o *MinimalDeploymentRunOptions) SetDeploymentToSubmit(v int32)`

SetDeploymentToSubmit sets DeploymentToSubmit field to given value.

### HasDeploymentToSubmit

`func (o *MinimalDeploymentRunOptions) HasDeploymentToSubmit() bool`

HasDeploymentToSubmit returns a boolean if a field has been set.

### GetDescription

`func (o *MinimalDeploymentRunOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MinimalDeploymentRunOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MinimalDeploymentRunOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MinimalDeploymentRunOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndState

`func (o *MinimalDeploymentRunOptions) GetEndState() string`

GetEndState returns the EndState field if non-nil, zero value otherwise.

### GetEndStateOk

`func (o *MinimalDeploymentRunOptions) GetEndStateOk() (*string, bool)`

GetEndStateOk returns a tuple with the EndState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndState

`func (o *MinimalDeploymentRunOptions) SetEndState(v string)`

SetEndState sets EndState field to given value.

### HasEndState

`func (o *MinimalDeploymentRunOptions) HasEndState() bool`

HasEndState returns a boolean if a field has been set.

### GetDuration

`func (o *MinimalDeploymentRunOptions) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MinimalDeploymentRunOptions) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MinimalDeploymentRunOptions) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MinimalDeploymentRunOptions) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEarliestStartTime

`func (o *MinimalDeploymentRunOptions) GetEarliestStartTime() int32`

GetEarliestStartTime returns the EarliestStartTime field if non-nil, zero value otherwise.

### GetEarliestStartTimeOk

`func (o *MinimalDeploymentRunOptions) GetEarliestStartTimeOk() (*int32, bool)`

GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStartTime

`func (o *MinimalDeploymentRunOptions) SetEarliestStartTime(v int32)`

SetEarliestStartTime sets EarliestStartTime field to given value.

### HasEarliestStartTime

`func (o *MinimalDeploymentRunOptions) HasEarliestStartTime() bool`

HasEarliestStartTime returns a boolean if a field has been set.

### GetEndExisting

`func (o *MinimalDeploymentRunOptions) GetEndExisting() bool`

GetEndExisting returns the EndExisting field if non-nil, zero value otherwise.

### GetEndExistingOk

`func (o *MinimalDeploymentRunOptions) GetEndExistingOk() (*bool, bool)`

GetEndExistingOk returns a tuple with the EndExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndExisting

`func (o *MinimalDeploymentRunOptions) SetEndExisting(v bool)`

SetEndExisting sets EndExisting field to given value.

### HasEndExisting

`func (o *MinimalDeploymentRunOptions) HasEndExisting() bool`

HasEndExisting returns a boolean if a field has been set.

### GetHostOptions

`func (o *MinimalDeploymentRunOptions) GetHostOptions() []HostOption`

GetHostOptions returns the HostOptions field if non-nil, zero value otherwise.

### GetHostOptionsOk

`func (o *MinimalDeploymentRunOptions) GetHostOptionsOk() (*[]HostOption, bool)`

GetHostOptionsOk returns a tuple with the HostOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOptions

`func (o *MinimalDeploymentRunOptions) SetHostOptions(v []HostOption)`

SetHostOptions sets HostOptions field to given value.

### HasHostOptions

`func (o *MinimalDeploymentRunOptions) HasHostOptions() bool`

HasHostOptions returns a boolean if a field has been set.

### GetId

`func (o *MinimalDeploymentRunOptions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalDeploymentRunOptions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalDeploymentRunOptions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalDeploymentRunOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *MinimalDeploymentRunOptions) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MinimalDeploymentRunOptions) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MinimalDeploymentRunOptions) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MinimalDeploymentRunOptions) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetVirtualizationRealmId

`func (o *MinimalDeploymentRunOptions) GetVirtualizationRealmId() int32`

GetVirtualizationRealmId returns the VirtualizationRealmId field if non-nil, zero value otherwise.

### GetVirtualizationRealmIdOk

`func (o *MinimalDeploymentRunOptions) GetVirtualizationRealmIdOk() (*int32, bool)`

GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmId

`func (o *MinimalDeploymentRunOptions) SetVirtualizationRealmId(v int32)`

SetVirtualizationRealmId sets VirtualizationRealmId field to given value.

### HasVirtualizationRealmId

`func (o *MinimalDeploymentRunOptions) HasVirtualizationRealmId() bool`

HasVirtualizationRealmId returns a boolean if a field has been set.

### GetProperties

`func (o *MinimalDeploymentRunOptions) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MinimalDeploymentRunOptions) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MinimalDeploymentRunOptions) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MinimalDeploymentRunOptions) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRetainOnError

`func (o *MinimalDeploymentRunOptions) GetRetainOnError() bool`

GetRetainOnError returns the RetainOnError field if non-nil, zero value otherwise.

### GetRetainOnErrorOk

`func (o *MinimalDeploymentRunOptions) GetRetainOnErrorOk() (*bool, bool)`

GetRetainOnErrorOk returns a tuple with the RetainOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainOnError

`func (o *MinimalDeploymentRunOptions) SetRetainOnError(v bool)`

SetRetainOnError sets RetainOnError field to given value.

### HasRetainOnError

`func (o *MinimalDeploymentRunOptions) HasRetainOnError() bool`

HasRetainOnError returns a boolean if a field has been set.

### GetRootPassword

`func (o *MinimalDeploymentRunOptions) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *MinimalDeploymentRunOptions) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *MinimalDeploymentRunOptions) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *MinimalDeploymentRunOptions) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetCreator

`func (o *MinimalDeploymentRunOptions) GetCreator() MinimalUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MinimalDeploymentRunOptions) GetCreatorOk() (*MinimalUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MinimalDeploymentRunOptions) SetCreator(v MinimalUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MinimalDeploymentRunOptions) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


