# InputScenario

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ScenarioHosts** | Pointer to [**[]InputScenarioHost**](InputScenarioHost.md) |  | [optional] 
**PrepareScenarioConfiguration** | Pointer to [**InputConfiguration**](InputConfiguration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**InputConfiguration**](InputConfiguration.md) |  | [optional] 

## Methods

### NewInputScenario

`func NewInputScenario() *InputScenario`

NewInputScenario instantiates a new InputScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputScenarioWithDefaults

`func NewInputScenarioWithDefaults() *InputScenario`

NewInputScenarioWithDefaults instantiates a new InputScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputScenario) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputScenario) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputScenario) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputScenario) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScenarioHosts

`func (o *InputScenario) GetScenarioHosts() []InputScenarioHost`

GetScenarioHosts returns the ScenarioHosts field if non-nil, zero value otherwise.

### GetScenarioHostsOk

`func (o *InputScenario) GetScenarioHostsOk() (*[]InputScenarioHost, bool)`

GetScenarioHostsOk returns a tuple with the ScenarioHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioHosts

`func (o *InputScenario) SetScenarioHosts(v []InputScenarioHost)`

SetScenarioHosts sets ScenarioHosts field to given value.

### HasScenarioHosts

`func (o *InputScenario) HasScenarioHosts() bool`

HasScenarioHosts returns a boolean if a field has been set.

### GetPrepareScenarioConfiguration

`func (o *InputScenario) GetPrepareScenarioConfiguration() InputConfiguration`

GetPrepareScenarioConfiguration returns the PrepareScenarioConfiguration field if non-nil, zero value otherwise.

### GetPrepareScenarioConfigurationOk

`func (o *InputScenario) GetPrepareScenarioConfigurationOk() (*InputConfiguration, bool)`

GetPrepareScenarioConfigurationOk returns a tuple with the PrepareScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareScenarioConfiguration

`func (o *InputScenario) SetPrepareScenarioConfiguration(v InputConfiguration)`

SetPrepareScenarioConfiguration sets PrepareScenarioConfiguration field to given value.

### HasPrepareScenarioConfiguration

`func (o *InputScenario) HasPrepareScenarioConfiguration() bool`

HasPrepareScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *InputScenario) GetTeardownScenarioConfiguration() InputConfiguration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *InputScenario) GetTeardownScenarioConfigurationOk() (*InputConfiguration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *InputScenario) SetTeardownScenarioConfiguration(v InputConfiguration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *InputScenario) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


