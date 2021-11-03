# InputSubmissionServiceForProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**SubmissionEndpoint** | [**InputSubmissionEndpointForProject**](InputSubmissionEndpointForProject.md) |  | 

## Methods

### NewInputSubmissionServiceForProject

`func NewInputSubmissionServiceForProject(submissionEndpoint InputSubmissionEndpointForProject, ) *InputSubmissionServiceForProject`

NewInputSubmissionServiceForProject instantiates a new InputSubmissionServiceForProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSubmissionServiceForProjectWithDefaults

`func NewInputSubmissionServiceForProjectWithDefaults() *InputSubmissionServiceForProject`

NewInputSubmissionServiceForProjectWithDefaults instantiates a new InputSubmissionServiceForProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InputSubmissionServiceForProject) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InputSubmissionServiceForProject) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InputSubmissionServiceForProject) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InputSubmissionServiceForProject) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCredentials

`func (o *InputSubmissionServiceForProject) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *InputSubmissionServiceForProject) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *InputSubmissionServiceForProject) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *InputSubmissionServiceForProject) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetSubmissionEndpoint

`func (o *InputSubmissionServiceForProject) GetSubmissionEndpoint() InputSubmissionEndpointForProject`

GetSubmissionEndpoint returns the SubmissionEndpoint field if non-nil, zero value otherwise.

### GetSubmissionEndpointOk

`func (o *InputSubmissionServiceForProject) GetSubmissionEndpointOk() (*InputSubmissionEndpointForProject, bool)`

GetSubmissionEndpointOk returns a tuple with the SubmissionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionEndpoint

`func (o *InputSubmissionServiceForProject) SetSubmissionEndpoint(v InputSubmissionEndpointForProject)`

SetSubmissionEndpoint sets SubmissionEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


