# AzureCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**Environment** | **string** |  | 
**RegionName** | **string** |  | 
**SecretAccessKey** | **string** |  | 
**SubscriptionId** | **string** |  | 
**Tenant** | **string** |  | 
**PublicContainerUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureCloud

`func NewAzureCloud(clientId string, environment string, regionName string, secretAccessKey string, subscriptionId string, tenant string, ) *AzureCloud`

NewAzureCloud instantiates a new AzureCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudWithDefaults

`func NewAzureCloudWithDefaults() *AzureCloud`

NewAzureCloudWithDefaults instantiates a new AzureCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureCloud) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureCloud) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureCloud) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetEnvironment

`func (o *AzureCloud) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AzureCloud) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AzureCloud) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetRegionName

`func (o *AzureCloud) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AzureCloud) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AzureCloud) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetSecretAccessKey

`func (o *AzureCloud) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AzureCloud) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AzureCloud) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### GetSubscriptionId

`func (o *AzureCloud) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureCloud) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureCloud) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenant

`func (o *AzureCloud) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AzureCloud) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AzureCloud) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetPublicContainerUrl

`func (o *AzureCloud) GetPublicContainerUrl() string`

GetPublicContainerUrl returns the PublicContainerUrl field if non-nil, zero value otherwise.

### GetPublicContainerUrlOk

`func (o *AzureCloud) GetPublicContainerUrlOk() (*string, bool)`

GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicContainerUrl

`func (o *AzureCloud) SetPublicContainerUrl(v string)`

SetPublicContainerUrl sets PublicContainerUrl field to given value.

### HasPublicContainerUrl

`func (o *AzureCloud) HasPublicContainerUrl() bool`

HasPublicContainerUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


