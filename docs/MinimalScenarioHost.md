# MinimalScenarioHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BuildOrder** | Pointer to **int32** |  | [optional] 
**Master** | Pointer to **bool** |  | [optional] 
**SystemRole** | Pointer to **string** |  | [optional] 
**SystemModule** | Pointer to [**MinimalSystemModule**](MinimalSystemModule.md) |  | [optional] 
**ConfigureScenarioConfiguration** | Pointer to [**MinimalConfiguration**](MinimalConfiguration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**MinimalConfiguration**](MinimalConfiguration.md) |  | [optional] 

## Methods

### NewMinimalScenarioHost

`func NewMinimalScenarioHost() *MinimalScenarioHost`

NewMinimalScenarioHost instantiates a new MinimalScenarioHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalScenarioHostWithDefaults

`func NewMinimalScenarioHostWithDefaults() *MinimalScenarioHost`

NewMinimalScenarioHostWithDefaults instantiates a new MinimalScenarioHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalScenarioHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalScenarioHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalScenarioHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalScenarioHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBuildOrder

`func (o *MinimalScenarioHost) GetBuildOrder() int32`

GetBuildOrder returns the BuildOrder field if non-nil, zero value otherwise.

### GetBuildOrderOk

`func (o *MinimalScenarioHost) GetBuildOrderOk() (*int32, bool)`

GetBuildOrderOk returns a tuple with the BuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOrder

`func (o *MinimalScenarioHost) SetBuildOrder(v int32)`

SetBuildOrder sets BuildOrder field to given value.

### HasBuildOrder

`func (o *MinimalScenarioHost) HasBuildOrder() bool`

HasBuildOrder returns a boolean if a field has been set.

### GetMaster

`func (o *MinimalScenarioHost) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *MinimalScenarioHost) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *MinimalScenarioHost) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *MinimalScenarioHost) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetSystemRole

`func (o *MinimalScenarioHost) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *MinimalScenarioHost) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *MinimalScenarioHost) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.

### HasSystemRole

`func (o *MinimalScenarioHost) HasSystemRole() bool`

HasSystemRole returns a boolean if a field has been set.

### GetSystemModule

`func (o *MinimalScenarioHost) GetSystemModule() MinimalSystemModule`

GetSystemModule returns the SystemModule field if non-nil, zero value otherwise.

### GetSystemModuleOk

`func (o *MinimalScenarioHost) GetSystemModuleOk() (*MinimalSystemModule, bool)`

GetSystemModuleOk returns a tuple with the SystemModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModule

`func (o *MinimalScenarioHost) SetSystemModule(v MinimalSystemModule)`

SetSystemModule sets SystemModule field to given value.

### HasSystemModule

`func (o *MinimalScenarioHost) HasSystemModule() bool`

HasSystemModule returns a boolean if a field has been set.

### GetConfigureScenarioConfiguration

`func (o *MinimalScenarioHost) GetConfigureScenarioConfiguration() MinimalConfiguration`

GetConfigureScenarioConfiguration returns the ConfigureScenarioConfiguration field if non-nil, zero value otherwise.

### GetConfigureScenarioConfigurationOk

`func (o *MinimalScenarioHost) GetConfigureScenarioConfigurationOk() (*MinimalConfiguration, bool)`

GetConfigureScenarioConfigurationOk returns a tuple with the ConfigureScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureScenarioConfiguration

`func (o *MinimalScenarioHost) SetConfigureScenarioConfiguration(v MinimalConfiguration)`

SetConfigureScenarioConfiguration sets ConfigureScenarioConfiguration field to given value.

### HasConfigureScenarioConfiguration

`func (o *MinimalScenarioHost) HasConfigureScenarioConfiguration() bool`

HasConfigureScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *MinimalScenarioHost) GetTeardownScenarioConfiguration() MinimalConfiguration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *MinimalScenarioHost) GetTeardownScenarioConfigurationOk() (*MinimalConfiguration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *MinimalScenarioHost) SetTeardownScenarioConfiguration(v MinimalConfiguration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *MinimalScenarioHost) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


