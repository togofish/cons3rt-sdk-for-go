# BaseCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **map[string]string** |  | [optional] 
**Context** | Pointer to **map[string]string** |  | [optional] 
**Resources** | Pointer to [**[]CloudResourceAccessListing**](CloudResourceAccessListing.md) |  | [optional] 

## Methods

### NewBaseCredentials

`func NewBaseCredentials() *BaseCredentials`

NewBaseCredentials instantiates a new BaseCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCredentialsWithDefaults

`func NewBaseCredentialsWithDefaults() *BaseCredentials`

NewBaseCredentialsWithDefaults instantiates a new BaseCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *BaseCredentials) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BaseCredentials) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BaseCredentials) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BaseCredentials) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetContext

`func (o *BaseCredentials) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BaseCredentials) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BaseCredentials) SetContext(v map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *BaseCredentials) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResources

`func (o *BaseCredentials) GetResources() []CloudResourceAccessListing`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BaseCredentials) GetResourcesOk() (*[]CloudResourceAccessListing, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BaseCredentials) SetResources(v []CloudResourceAccessListing)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BaseCredentials) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


