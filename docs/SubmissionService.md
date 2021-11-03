# SubmissionService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostLoggingString** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**SubmissionEndpoint** | [**SubmissionEndpoint**](SubmissionEndpoint.md) |  | 

## Methods

### NewSubmissionService

`func NewSubmissionService(name string, submissionEndpoint SubmissionEndpoint, ) *SubmissionService`

NewSubmissionService instantiates a new SubmissionService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionServiceWithDefaults

`func NewSubmissionServiceWithDefaults() *SubmissionService`

NewSubmissionServiceWithDefaults instantiates a new SubmissionService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostLoggingString

`func (o *SubmissionService) GetHostLoggingString() string`

GetHostLoggingString returns the HostLoggingString field if non-nil, zero value otherwise.

### GetHostLoggingStringOk

`func (o *SubmissionService) GetHostLoggingStringOk() (*string, bool)`

GetHostLoggingStringOk returns a tuple with the HostLoggingString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostLoggingString

`func (o *SubmissionService) SetHostLoggingString(v string)`

SetHostLoggingString sets HostLoggingString field to given value.

### HasHostLoggingString

`func (o *SubmissionService) HasHostLoggingString() bool`

HasHostLoggingString returns a boolean if a field has been set.

### GetCredentials

`func (o *SubmissionService) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SubmissionService) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SubmissionService) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SubmissionService) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEnabled

`func (o *SubmissionService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubmissionService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubmissionService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SubmissionService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *SubmissionService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionService) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SubmissionService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubmissionService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubmissionService) SetName(v string)`

SetName sets Name field to given value.


### GetSubmissionEndpoint

`func (o *SubmissionService) GetSubmissionEndpoint() SubmissionEndpoint`

GetSubmissionEndpoint returns the SubmissionEndpoint field if non-nil, zero value otherwise.

### GetSubmissionEndpointOk

`func (o *SubmissionService) GetSubmissionEndpointOk() (*SubmissionEndpoint, bool)`

GetSubmissionEndpointOk returns a tuple with the SubmissionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionEndpoint

`func (o *SubmissionService) SetSubmissionEndpoint(v SubmissionEndpoint)`

SetSubmissionEndpoint sets SubmissionEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


