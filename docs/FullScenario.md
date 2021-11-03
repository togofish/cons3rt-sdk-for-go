# FullScenario

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScenarioHosts** | Pointer to [**[]MinimalScenarioHost**](MinimalScenarioHost.md) |  | [optional] 
**PrepareScenarioConfiguration** | Pointer to [**MinimalConfiguration**](MinimalConfiguration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**MinimalConfiguration**](MinimalConfiguration.md) |  | [optional] 

## Methods

### NewFullScenario

`func NewFullScenario() *FullScenario`

NewFullScenario instantiates a new FullScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullScenarioWithDefaults

`func NewFullScenarioWithDefaults() *FullScenario`

NewFullScenarioWithDefaults instantiates a new FullScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenarioHosts

`func (o *FullScenario) GetScenarioHosts() []MinimalScenarioHost`

GetScenarioHosts returns the ScenarioHosts field if non-nil, zero value otherwise.

### GetScenarioHostsOk

`func (o *FullScenario) GetScenarioHostsOk() (*[]MinimalScenarioHost, bool)`

GetScenarioHostsOk returns a tuple with the ScenarioHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioHosts

`func (o *FullScenario) SetScenarioHosts(v []MinimalScenarioHost)`

SetScenarioHosts sets ScenarioHosts field to given value.

### HasScenarioHosts

`func (o *FullScenario) HasScenarioHosts() bool`

HasScenarioHosts returns a boolean if a field has been set.

### GetPrepareScenarioConfiguration

`func (o *FullScenario) GetPrepareScenarioConfiguration() MinimalConfiguration`

GetPrepareScenarioConfiguration returns the PrepareScenarioConfiguration field if non-nil, zero value otherwise.

### GetPrepareScenarioConfigurationOk

`func (o *FullScenario) GetPrepareScenarioConfigurationOk() (*MinimalConfiguration, bool)`

GetPrepareScenarioConfigurationOk returns a tuple with the PrepareScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareScenarioConfiguration

`func (o *FullScenario) SetPrepareScenarioConfiguration(v MinimalConfiguration)`

SetPrepareScenarioConfiguration sets PrepareScenarioConfiguration field to given value.

### HasPrepareScenarioConfiguration

`func (o *FullScenario) HasPrepareScenarioConfiguration() bool`

HasPrepareScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *FullScenario) GetTeardownScenarioConfiguration() MinimalConfiguration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *FullScenario) GetTeardownScenarioConfigurationOk() (*MinimalConfiguration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *FullScenario) SetTeardownScenarioConfiguration(v MinimalConfiguration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *FullScenario) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


