# InputScenarioFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScenarioHosts** | Pointer to [**[]InputScenarioHost**](InputScenarioHost.md) |  | [optional] 
**ScenarioBuildOrder** | Pointer to **int32** |  | [optional] 
**PrepareScenarioConfiguration** | Pointer to [**InputConfiguration**](InputConfiguration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**InputConfiguration**](InputConfiguration.md) |  | [optional] 

## Methods

### NewInputScenarioFull

`func NewInputScenarioFull() *InputScenarioFull`

NewInputScenarioFull instantiates a new InputScenarioFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputScenarioFullWithDefaults

`func NewInputScenarioFullWithDefaults() *InputScenarioFull`

NewInputScenarioFullWithDefaults instantiates a new InputScenarioFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputScenarioFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputScenarioFull) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputScenarioFull) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputScenarioFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InputScenarioFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputScenarioFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputScenarioFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputScenarioFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScenarioHosts

`func (o *InputScenarioFull) GetScenarioHosts() []InputScenarioHost`

GetScenarioHosts returns the ScenarioHosts field if non-nil, zero value otherwise.

### GetScenarioHostsOk

`func (o *InputScenarioFull) GetScenarioHostsOk() (*[]InputScenarioHost, bool)`

GetScenarioHostsOk returns a tuple with the ScenarioHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioHosts

`func (o *InputScenarioFull) SetScenarioHosts(v []InputScenarioHost)`

SetScenarioHosts sets ScenarioHosts field to given value.

### HasScenarioHosts

`func (o *InputScenarioFull) HasScenarioHosts() bool`

HasScenarioHosts returns a boolean if a field has been set.

### GetScenarioBuildOrder

`func (o *InputScenarioFull) GetScenarioBuildOrder() int32`

GetScenarioBuildOrder returns the ScenarioBuildOrder field if non-nil, zero value otherwise.

### GetScenarioBuildOrderOk

`func (o *InputScenarioFull) GetScenarioBuildOrderOk() (*int32, bool)`

GetScenarioBuildOrderOk returns a tuple with the ScenarioBuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioBuildOrder

`func (o *InputScenarioFull) SetScenarioBuildOrder(v int32)`

SetScenarioBuildOrder sets ScenarioBuildOrder field to given value.

### HasScenarioBuildOrder

`func (o *InputScenarioFull) HasScenarioBuildOrder() bool`

HasScenarioBuildOrder returns a boolean if a field has been set.

### GetPrepareScenarioConfiguration

`func (o *InputScenarioFull) GetPrepareScenarioConfiguration() InputConfiguration`

GetPrepareScenarioConfiguration returns the PrepareScenarioConfiguration field if non-nil, zero value otherwise.

### GetPrepareScenarioConfigurationOk

`func (o *InputScenarioFull) GetPrepareScenarioConfigurationOk() (*InputConfiguration, bool)`

GetPrepareScenarioConfigurationOk returns a tuple with the PrepareScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareScenarioConfiguration

`func (o *InputScenarioFull) SetPrepareScenarioConfiguration(v InputConfiguration)`

SetPrepareScenarioConfiguration sets PrepareScenarioConfiguration field to given value.

### HasPrepareScenarioConfiguration

`func (o *InputScenarioFull) HasPrepareScenarioConfiguration() bool`

HasPrepareScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *InputScenarioFull) GetTeardownScenarioConfiguration() InputConfiguration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *InputScenarioFull) GetTeardownScenarioConfigurationOk() (*InputConfiguration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *InputScenarioFull) SetTeardownScenarioConfiguration(v InputConfiguration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *InputScenarioFull) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


