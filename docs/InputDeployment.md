# InputDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**InputMetadata**](InputMetadata.md) |  | [optional] 
**Scenarios** | Pointer to [**[]InputScenarioFull**](InputScenarioFull.md) |  | [optional] 
**TestBundles** | Pointer to [**[]InputTestBundle**](InputTestBundle.md) |  | [optional] 

## Methods

### NewInputDeployment

`func NewInputDeployment() *InputDeployment`

NewInputDeployment instantiates a new InputDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDeploymentWithDefaults

`func NewInputDeploymentWithDefaults() *InputDeployment`

NewInputDeploymentWithDefaults instantiates a new InputDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *InputDeployment) GetMetadata() InputMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InputDeployment) GetMetadataOk() (*InputMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InputDeployment) SetMetadata(v InputMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InputDeployment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScenarios

`func (o *InputDeployment) GetScenarios() []InputScenarioFull`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *InputDeployment) GetScenariosOk() (*[]InputScenarioFull, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *InputDeployment) SetScenarios(v []InputScenarioFull)`

SetScenarios sets Scenarios field to given value.

### HasScenarios

`func (o *InputDeployment) HasScenarios() bool`

HasScenarios returns a boolean if a field has been set.

### GetTestBundles

`func (o *InputDeployment) GetTestBundles() []InputTestBundle`

GetTestBundles returns the TestBundles field if non-nil, zero value otherwise.

### GetTestBundlesOk

`func (o *InputDeployment) GetTestBundlesOk() (*[]InputTestBundle, bool)`

GetTestBundlesOk returns a tuple with the TestBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBundles

`func (o *InputDeployment) SetTestBundles(v []InputTestBundle)`

SetTestBundles sets TestBundles field to given value.

### HasTestBundles

`func (o *InputDeployment) HasTestBundles() bool`

HasTestBundles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


