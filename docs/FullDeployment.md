# FullDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurringSchedules** | Pointer to [**[]MinimalRecurringSchedule**](MinimalRecurringSchedule.md) |  | [optional] 
**Scenario** | Pointer to [**[]GeneralScenario**](GeneralScenario.md) |  | [optional] 
**TestBundles** | Pointer to [**[]MinimalTestBundle**](MinimalTestBundle.md) |  | [optional] 
**DeploymentHosts** | Pointer to [**[]MinimalDeploymentHost**](MinimalDeploymentHost.md) |  | [optional] 

## Methods

### NewFullDeployment

`func NewFullDeployment() *FullDeployment`

NewFullDeployment instantiates a new FullDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDeploymentWithDefaults

`func NewFullDeploymentWithDefaults() *FullDeployment`

NewFullDeploymentWithDefaults instantiates a new FullDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurringSchedules

`func (o *FullDeployment) GetRecurringSchedules() []MinimalRecurringSchedule`

GetRecurringSchedules returns the RecurringSchedules field if non-nil, zero value otherwise.

### GetRecurringSchedulesOk

`func (o *FullDeployment) GetRecurringSchedulesOk() (*[]MinimalRecurringSchedule, bool)`

GetRecurringSchedulesOk returns a tuple with the RecurringSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSchedules

`func (o *FullDeployment) SetRecurringSchedules(v []MinimalRecurringSchedule)`

SetRecurringSchedules sets RecurringSchedules field to given value.

### HasRecurringSchedules

`func (o *FullDeployment) HasRecurringSchedules() bool`

HasRecurringSchedules returns a boolean if a field has been set.

### GetScenario

`func (o *FullDeployment) GetScenario() []GeneralScenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *FullDeployment) GetScenarioOk() (*[]GeneralScenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *FullDeployment) SetScenario(v []GeneralScenario)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *FullDeployment) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetTestBundles

`func (o *FullDeployment) GetTestBundles() []MinimalTestBundle`

GetTestBundles returns the TestBundles field if non-nil, zero value otherwise.

### GetTestBundlesOk

`func (o *FullDeployment) GetTestBundlesOk() (*[]MinimalTestBundle, bool)`

GetTestBundlesOk returns a tuple with the TestBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBundles

`func (o *FullDeployment) SetTestBundles(v []MinimalTestBundle)`

SetTestBundles sets TestBundles field to given value.

### HasTestBundles

`func (o *FullDeployment) HasTestBundles() bool`

HasTestBundles returns a boolean if a field has been set.

### GetDeploymentHosts

`func (o *FullDeployment) GetDeploymentHosts() []MinimalDeploymentHost`

GetDeploymentHosts returns the DeploymentHosts field if non-nil, zero value otherwise.

### GetDeploymentHostsOk

`func (o *FullDeployment) GetDeploymentHostsOk() (*[]MinimalDeploymentHost, bool)`

GetDeploymentHostsOk returns a tuple with the DeploymentHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentHosts

`func (o *FullDeployment) SetDeploymentHosts(v []MinimalDeploymentHost)`

SetDeploymentHosts sets DeploymentHosts field to given value.

### HasDeploymentHosts

`func (o *FullDeployment) HasDeploymentHosts() bool`

HasDeploymentHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


