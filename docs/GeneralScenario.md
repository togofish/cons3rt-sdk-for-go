# GeneralScenario

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ScenarioBuildOrder** | Pointer to **int32** |  | [optional] 
**ScenarioHosts** | Pointer to [**[]MinimalScenarioHost**](MinimalScenarioHost.md) |  | [optional] 
**PrepareScenarioConfiguration** | Pointer to [**MinimalConfiguration**](MinimalConfiguration.md) |  | [optional] 
**TeardownScenarioConfiguration** | Pointer to [**MinimalConfiguration**](MinimalConfiguration.md) |  | [optional] 

## Methods

### NewGeneralScenario

`func NewGeneralScenario() *GeneralScenario`

NewGeneralScenario instantiates a new GeneralScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralScenarioWithDefaults

`func NewGeneralScenarioWithDefaults() *GeneralScenario`

NewGeneralScenarioWithDefaults instantiates a new GeneralScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GeneralScenario) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneralScenario) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneralScenario) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GeneralScenario) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GeneralScenario) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeneralScenario) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeneralScenario) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GeneralScenario) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GeneralScenario) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeneralScenario) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeneralScenario) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeneralScenario) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffline

`func (o *GeneralScenario) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *GeneralScenario) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *GeneralScenario) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *GeneralScenario) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetState

`func (o *GeneralScenario) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeneralScenario) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeneralScenario) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GeneralScenario) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *GeneralScenario) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GeneralScenario) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GeneralScenario) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GeneralScenario) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetScenarioBuildOrder

`func (o *GeneralScenario) GetScenarioBuildOrder() int32`

GetScenarioBuildOrder returns the ScenarioBuildOrder field if non-nil, zero value otherwise.

### GetScenarioBuildOrderOk

`func (o *GeneralScenario) GetScenarioBuildOrderOk() (*int32, bool)`

GetScenarioBuildOrderOk returns a tuple with the ScenarioBuildOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioBuildOrder

`func (o *GeneralScenario) SetScenarioBuildOrder(v int32)`

SetScenarioBuildOrder sets ScenarioBuildOrder field to given value.

### HasScenarioBuildOrder

`func (o *GeneralScenario) HasScenarioBuildOrder() bool`

HasScenarioBuildOrder returns a boolean if a field has been set.

### GetScenarioHosts

`func (o *GeneralScenario) GetScenarioHosts() []MinimalScenarioHost`

GetScenarioHosts returns the ScenarioHosts field if non-nil, zero value otherwise.

### GetScenarioHostsOk

`func (o *GeneralScenario) GetScenarioHostsOk() (*[]MinimalScenarioHost, bool)`

GetScenarioHostsOk returns a tuple with the ScenarioHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioHosts

`func (o *GeneralScenario) SetScenarioHosts(v []MinimalScenarioHost)`

SetScenarioHosts sets ScenarioHosts field to given value.

### HasScenarioHosts

`func (o *GeneralScenario) HasScenarioHosts() bool`

HasScenarioHosts returns a boolean if a field has been set.

### GetPrepareScenarioConfiguration

`func (o *GeneralScenario) GetPrepareScenarioConfiguration() MinimalConfiguration`

GetPrepareScenarioConfiguration returns the PrepareScenarioConfiguration field if non-nil, zero value otherwise.

### GetPrepareScenarioConfigurationOk

`func (o *GeneralScenario) GetPrepareScenarioConfigurationOk() (*MinimalConfiguration, bool)`

GetPrepareScenarioConfigurationOk returns a tuple with the PrepareScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareScenarioConfiguration

`func (o *GeneralScenario) SetPrepareScenarioConfiguration(v MinimalConfiguration)`

SetPrepareScenarioConfiguration sets PrepareScenarioConfiguration field to given value.

### HasPrepareScenarioConfiguration

`func (o *GeneralScenario) HasPrepareScenarioConfiguration() bool`

HasPrepareScenarioConfiguration returns a boolean if a field has been set.

### GetTeardownScenarioConfiguration

`func (o *GeneralScenario) GetTeardownScenarioConfiguration() MinimalConfiguration`

GetTeardownScenarioConfiguration returns the TeardownScenarioConfiguration field if non-nil, zero value otherwise.

### GetTeardownScenarioConfigurationOk

`func (o *GeneralScenario) GetTeardownScenarioConfigurationOk() (*MinimalConfiguration, bool)`

GetTeardownScenarioConfigurationOk returns a tuple with the TeardownScenarioConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownScenarioConfiguration

`func (o *GeneralScenario) SetTeardownScenarioConfiguration(v MinimalConfiguration)`

SetTeardownScenarioConfiguration sets TeardownScenarioConfiguration field to given value.

### HasTeardownScenarioConfiguration

`func (o *GeneralScenario) HasTeardownScenarioConfiguration() bool`

HasTeardownScenarioConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


