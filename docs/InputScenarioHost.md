# InputScenarioHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildOrder** | Pointer to **int32** |  | [optional] 
**SystemRole** | **string** |  | 
**SystemModule** | Pointer to [**InputSystemModuleFull**](InputSystemModuleFull.md) |  | [optional] 
**Master** | Pointer to **bool** |  | [optional] 
**ConfigureScenarioConfiguration** | Pointer to [**InputConfiguration**](InputConfiguration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**InputConfiguration**](InputConfiguration.md) |  | [optional] 

## Methods

### NewInputScenarioHost

`func NewInputScenarioHost(systemRole string, ) *InputScenarioHost`

NewInputScenarioHost instantiates a new InputScenarioHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputScenarioHostWithDefaults

`func NewInputScenarioHostWithDefaults() *InputScenarioHost`

NewInputScenarioHostWithDefaults instantiates a new InputScenarioHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildOrder

`func (o *InputScenarioHost) GetBuildOrder() int32`

GetBuildOrder returns the BuildOrder field if non-nil, zero value otherwise.

### GetBuildOrderOk

`func (o *InputScenarioHost) GetBuildOrderOk() (*int32, bool)`

GetBuildOrderOk returns a tuple with the BuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOrder

`func (o *InputScenarioHost) SetBuildOrder(v int32)`

SetBuildOrder sets BuildOrder field to given value.

### HasBuildOrder

`func (o *InputScenarioHost) HasBuildOrder() bool`

HasBuildOrder returns a boolean if a field has been set.

### GetSystemRole

`func (o *InputScenarioHost) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *InputScenarioHost) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *InputScenarioHost) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.


### GetSystemModule

`func (o *InputScenarioHost) GetSystemModule() InputSystemModuleFull`

GetSystemModule returns the SystemModule field if non-nil, zero value otherwise.

### GetSystemModuleOk

`func (o *InputScenarioHost) GetSystemModuleOk() (*InputSystemModuleFull, bool)`

GetSystemModuleOk returns a tuple with the SystemModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModule

`func (o *InputScenarioHost) SetSystemModule(v InputSystemModuleFull)`

SetSystemModule sets SystemModule field to given value.

### HasSystemModule

`func (o *InputScenarioHost) HasSystemModule() bool`

HasSystemModule returns a boolean if a field has been set.

### GetMaster

`func (o *InputScenarioHost) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *InputScenarioHost) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *InputScenarioHost) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *InputScenarioHost) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetConfigureScenarioConfiguration

`func (o *InputScenarioHost) GetConfigureScenarioConfiguration() InputConfiguration`

GetConfigureScenarioConfiguration returns the ConfigureScenarioConfiguration field if non-nil, zero value otherwise.

### GetConfigureScenarioConfigurationOk

`func (o *InputScenarioHost) GetConfigureScenarioConfigurationOk() (*InputConfiguration, bool)`

GetConfigureScenarioConfigurationOk returns a tuple with the ConfigureScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureScenarioConfiguration

`func (o *InputScenarioHost) SetConfigureScenarioConfiguration(v InputConfiguration)`

SetConfigureScenarioConfiguration sets ConfigureScenarioConfiguration field to given value.

### HasConfigureScenarioConfiguration

`func (o *InputScenarioHost) HasConfigureScenarioConfiguration() bool`

HasConfigureScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *InputScenarioHost) GetTeardownScenarioConfiguration() InputConfiguration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *InputScenarioHost) GetTeardownScenarioConfigurationOk() (*InputConfiguration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *InputScenarioHost) SetTeardownScenarioConfiguration(v InputConfiguration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *InputScenarioHost) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


