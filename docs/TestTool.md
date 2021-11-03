# TestTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestToolType** | Pointer to **string** |  | [optional] 
**TestToolVendor** | Pointer to **string** |  | [optional] 
**AutoRegistered** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InstanceLimit** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**TestRunTasks** | Pointer to **[]int32** |  | [optional] 
**TestToolAgentInstalled** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**TmaEnabled** | Pointer to **bool** |  | [optional] 
**TmaOnline** | Pointer to **bool** |  | [optional] 
**TmaServiceName** | Pointer to **string** |  | [optional] 
**TmaVersion** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewTestTool

`func NewTestTool() *TestTool`

NewTestTool instantiates a new TestTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestToolWithDefaults

`func NewTestToolWithDefaults() *TestTool`

NewTestToolWithDefaults instantiates a new TestTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestToolType

`func (o *TestTool) GetTestToolType() string`

GetTestToolType returns the TestToolType field if non-nil, zero value otherwise.

### GetTestToolTypeOk

`func (o *TestTool) GetTestToolTypeOk() (*string, bool)`

GetTestToolTypeOk returns a tuple with the TestToolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestToolType

`func (o *TestTool) SetTestToolType(v string)`

SetTestToolType sets TestToolType field to given value.

### HasTestToolType

`func (o *TestTool) HasTestToolType() bool`

HasTestToolType returns a boolean if a field has been set.

### GetTestToolVendor

`func (o *TestTool) GetTestToolVendor() string`

GetTestToolVendor returns the TestToolVendor field if non-nil, zero value otherwise.

### GetTestToolVendorOk

`func (o *TestTool) GetTestToolVendorOk() (*string, bool)`

GetTestToolVendorOk returns a tuple with the TestToolVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestToolVendor

`func (o *TestTool) SetTestToolVendor(v string)`

SetTestToolVendor sets TestToolVendor field to given value.

### HasTestToolVendor

`func (o *TestTool) HasTestToolVendor() bool`

HasTestToolVendor returns a boolean if a field has been set.

### GetAutoRegistered

`func (o *TestTool) GetAutoRegistered() bool`

GetAutoRegistered returns the AutoRegistered field if non-nil, zero value otherwise.

### GetAutoRegisteredOk

`func (o *TestTool) GetAutoRegisteredOk() (*bool, bool)`

GetAutoRegisteredOk returns a tuple with the AutoRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegistered

`func (o *TestTool) SetAutoRegistered(v bool)`

SetAutoRegistered sets AutoRegistered field to given value.

### HasAutoRegistered

`func (o *TestTool) HasAutoRegistered() bool`

HasAutoRegistered returns a boolean if a field has been set.

### GetDescription

`func (o *TestTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostname

`func (o *TestTool) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TestTool) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TestTool) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *TestTool) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *TestTool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestTool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestTool) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TestTool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceLimit

`func (o *TestTool) GetInstanceLimit() int32`

GetInstanceLimit returns the InstanceLimit field if non-nil, zero value otherwise.

### GetInstanceLimitOk

`func (o *TestTool) GetInstanceLimitOk() (*int32, bool)`

GetInstanceLimitOk returns a tuple with the InstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLimit

`func (o *TestTool) SetInstanceLimit(v int32)`

SetInstanceLimit sets InstanceLimit field to given value.

### HasInstanceLimit

`func (o *TestTool) HasInstanceLimit() bool`

HasInstanceLimit returns a boolean if a field has been set.

### GetName

`func (o *TestTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *TestTool) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TestTool) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TestTool) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TestTool) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTestRunTasks

`func (o *TestTool) GetTestRunTasks() []int32`

GetTestRunTasks returns the TestRunTasks field if non-nil, zero value otherwise.

### GetTestRunTasksOk

`func (o *TestTool) GetTestRunTasksOk() (*[]int32, bool)`

GetTestRunTasksOk returns a tuple with the TestRunTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunTasks

`func (o *TestTool) SetTestRunTasks(v []int32)`

SetTestRunTasks sets TestRunTasks field to given value.

### HasTestRunTasks

`func (o *TestTool) HasTestRunTasks() bool`

HasTestRunTasks returns a boolean if a field has been set.

### GetTestToolAgentInstalled

`func (o *TestTool) GetTestToolAgentInstalled() bool`

GetTestToolAgentInstalled returns the TestToolAgentInstalled field if non-nil, zero value otherwise.

### GetTestToolAgentInstalledOk

`func (o *TestTool) GetTestToolAgentInstalledOk() (*bool, bool)`

GetTestToolAgentInstalledOk returns a tuple with the TestToolAgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestToolAgentInstalled

`func (o *TestTool) SetTestToolAgentInstalled(v bool)`

SetTestToolAgentInstalled sets TestToolAgentInstalled field to given value.

### HasTestToolAgentInstalled

`func (o *TestTool) HasTestToolAgentInstalled() bool`

HasTestToolAgentInstalled returns a boolean if a field has been set.

### GetVersion

`func (o *TestTool) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TestTool) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TestTool) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TestTool) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTmaEnabled

`func (o *TestTool) GetTmaEnabled() bool`

GetTmaEnabled returns the TmaEnabled field if non-nil, zero value otherwise.

### GetTmaEnabledOk

`func (o *TestTool) GetTmaEnabledOk() (*bool, bool)`

GetTmaEnabledOk returns a tuple with the TmaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmaEnabled

`func (o *TestTool) SetTmaEnabled(v bool)`

SetTmaEnabled sets TmaEnabled field to given value.

### HasTmaEnabled

`func (o *TestTool) HasTmaEnabled() bool`

HasTmaEnabled returns a boolean if a field has been set.

### GetTmaOnline

`func (o *TestTool) GetTmaOnline() bool`

GetTmaOnline returns the TmaOnline field if non-nil, zero value otherwise.

### GetTmaOnlineOk

`func (o *TestTool) GetTmaOnlineOk() (*bool, bool)`

GetTmaOnlineOk returns a tuple with the TmaOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmaOnline

`func (o *TestTool) SetTmaOnline(v bool)`

SetTmaOnline sets TmaOnline field to given value.

### HasTmaOnline

`func (o *TestTool) HasTmaOnline() bool`

HasTmaOnline returns a boolean if a field has been set.

### GetTmaServiceName

`func (o *TestTool) GetTmaServiceName() string`

GetTmaServiceName returns the TmaServiceName field if non-nil, zero value otherwise.

### GetTmaServiceNameOk

`func (o *TestTool) GetTmaServiceNameOk() (*string, bool)`

GetTmaServiceNameOk returns a tuple with the TmaServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmaServiceName

`func (o *TestTool) SetTmaServiceName(v string)`

SetTmaServiceName sets TmaServiceName field to given value.

### HasTmaServiceName

`func (o *TestTool) HasTmaServiceName() bool`

HasTmaServiceName returns a boolean if a field has been set.

### GetTmaVersion

`func (o *TestTool) GetTmaVersion() string`

GetTmaVersion returns the TmaVersion field if non-nil, zero value otherwise.

### GetTmaVersionOk

`func (o *TestTool) GetTmaVersionOk() (*string, bool)`

GetTmaVersionOk returns a tuple with the TmaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmaVersion

`func (o *TestTool) SetTmaVersion(v string)`

SetTmaVersion sets TmaVersion field to given value.

### HasTmaVersion

`func (o *TestTool) HasTmaVersion() bool`

HasTmaVersion returns a boolean if a field has been set.

### GetVisibility

`func (o *TestTool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TestTool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TestTool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TestTool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


