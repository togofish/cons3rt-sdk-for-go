# ScenarioHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildOrder** | Pointer to **int32** |  | [optional] 
**ConfigureScenarioConfiguration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Master** | Pointer to **bool** |  | [optional] 
**SystemModule** | [**SystemModule**](SystemModule.md) |  | 
**SystemRole** | **string** |  | 

## Methods

### NewScenarioHost

`func NewScenarioHost(systemModule SystemModule, systemRole string, ) *ScenarioHost`

NewScenarioHost instantiates a new ScenarioHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioHostWithDefaults

`func NewScenarioHostWithDefaults() *ScenarioHost`

NewScenarioHostWithDefaults instantiates a new ScenarioHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildOrder

`func (o *ScenarioHost) GetBuildOrder() int32`

GetBuildOrder returns the BuildOrder field if non-nil, zero value otherwise.

### GetBuildOrderOk

`func (o *ScenarioHost) GetBuildOrderOk() (*int32, bool)`

GetBuildOrderOk returns a tuple with the BuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOrder

`func (o *ScenarioHost) SetBuildOrder(v int32)`

SetBuildOrder sets BuildOrder field to given value.

### HasBuildOrder

`func (o *ScenarioHost) HasBuildOrder() bool`

HasBuildOrder returns a boolean if a field has been set.

### GetConfigureScenarioConfiguration

`func (o *ScenarioHost) GetConfigureScenarioConfiguration() Configuration`

GetConfigureScenarioConfiguration returns the ConfigureScenarioConfiguration field if non-nil, zero value otherwise.

### GetConfigureScenarioConfigurationOk

`func (o *ScenarioHost) GetConfigureScenarioConfigurationOk() (*Configuration, bool)`

GetConfigureScenarioConfigurationOk returns a tuple with the ConfigureScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureScenarioConfiguration

`func (o *ScenarioHost) SetConfigureScenarioConfiguration(v Configuration)`

SetConfigureScenarioConfiguration sets ConfigureScenarioConfiguration field to given value.

### HasConfigureScenarioConfiguration

`func (o *ScenarioHost) HasConfigureScenarioConfiguration() bool`

HasConfigureScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *ScenarioHost) GetTeardownScenarioConfiguration() Configuration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *ScenarioHost) GetTeardownScenarioConfigurationOk() (*Configuration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *ScenarioHost) SetTeardownScenarioConfiguration(v Configuration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *ScenarioHost) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.

### GetId

`func (o *ScenarioHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScenarioHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScenarioHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScenarioHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaster

`func (o *ScenarioHost) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ScenarioHost) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ScenarioHost) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ScenarioHost) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetSystemModule

`func (o *ScenarioHost) GetSystemModule() SystemModule`

GetSystemModule returns the SystemModule field if non-nil, zero value otherwise.

### GetSystemModuleOk

`func (o *ScenarioHost) GetSystemModuleOk() (*SystemModule, bool)`

GetSystemModuleOk returns a tuple with the SystemModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModule

`func (o *ScenarioHost) SetSystemModule(v SystemModule)`

SetSystemModule sets SystemModule field to given value.


### GetSystemRole

`func (o *ScenarioHost) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *ScenarioHost) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *ScenarioHost) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


