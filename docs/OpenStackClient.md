# OpenStackClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeystoneVersion** | Pointer to **string** |  | [optional] 
**TenantName** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenStackClient

`func NewOpenStackClient() *OpenStackClient`

NewOpenStackClient instantiates a new OpenStackClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenStackClientWithDefaults

`func NewOpenStackClientWithDefaults() *OpenStackClient`

NewOpenStackClientWithDefaults instantiates a new OpenStackClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeystoneVersion

`func (o *OpenStackClient) GetKeystoneVersion() string`

GetKeystoneVersion returns the KeystoneVersion field if non-nil, zero value otherwise.

### GetKeystoneVersionOk

`func (o *OpenStackClient) GetKeystoneVersionOk() (*string, bool)`

GetKeystoneVersionOk returns a tuple with the KeystoneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneVersion

`func (o *OpenStackClient) SetKeystoneVersion(v string)`

SetKeystoneVersion sets KeystoneVersion field to given value.

### HasKeystoneVersion

`func (o *OpenStackClient) HasKeystoneVersion() bool`

HasKeystoneVersion returns a boolean if a field has been set.

### GetTenantName

`func (o *OpenStackClient) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *OpenStackClient) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *OpenStackClient) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *OpenStackClient) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


