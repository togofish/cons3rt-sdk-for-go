# AzureRegisterCloudSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Environment** | **string** |  | 
**LocalStorageResourceGroupName** | Pointer to **string** |  | [optional] 
**LocalStorageKey** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**ResourceGroupName** | **string** |  | 
**TenantId** | **string** |  | 
**VirtualNetworkName** | **string** |  | 
**LocalStorageName** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureRegisterCloudSpaceRequest

`func NewAzureRegisterCloudSpaceRequest(accountId string, environment string, region string, resourceGroupName string, tenantId string, virtualNetworkName string, ) *AzureRegisterCloudSpaceRequest`

NewAzureRegisterCloudSpaceRequest instantiates a new AzureRegisterCloudSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRegisterCloudSpaceRequestWithDefaults

`func NewAzureRegisterCloudSpaceRequestWithDefaults() *AzureRegisterCloudSpaceRequest`

NewAzureRegisterCloudSpaceRequestWithDefaults instantiates a new AzureRegisterCloudSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AzureRegisterCloudSpaceRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AzureRegisterCloudSpaceRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AzureRegisterCloudSpaceRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEnvironment

`func (o *AzureRegisterCloudSpaceRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AzureRegisterCloudSpaceRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AzureRegisterCloudSpaceRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetLocalStorageResourceGroupName

`func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageResourceGroupName() string`

GetLocalStorageResourceGroupName returns the LocalStorageResourceGroupName field if non-nil, zero value otherwise.

### GetLocalStorageResourceGroupNameOk

`func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageResourceGroupNameOk() (*string, bool)`

GetLocalStorageResourceGroupNameOk returns a tuple with the LocalStorageResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageResourceGroupName

`func (o *AzureRegisterCloudSpaceRequest) SetLocalStorageResourceGroupName(v string)`

SetLocalStorageResourceGroupName sets LocalStorageResourceGroupName field to given value.

### HasLocalStorageResourceGroupName

`func (o *AzureRegisterCloudSpaceRequest) HasLocalStorageResourceGroupName() bool`

HasLocalStorageResourceGroupName returns a boolean if a field has been set.

### GetLocalStorageKey

`func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageKey() string`

GetLocalStorageKey returns the LocalStorageKey field if non-nil, zero value otherwise.

### GetLocalStorageKeyOk

`func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageKeyOk() (*string, bool)`

GetLocalStorageKeyOk returns a tuple with the LocalStorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageKey

`func (o *AzureRegisterCloudSpaceRequest) SetLocalStorageKey(v string)`

SetLocalStorageKey sets LocalStorageKey field to given value.

### HasLocalStorageKey

`func (o *AzureRegisterCloudSpaceRequest) HasLocalStorageKey() bool`

HasLocalStorageKey returns a boolean if a field has been set.

### GetRegion

`func (o *AzureRegisterCloudSpaceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureRegisterCloudSpaceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureRegisterCloudSpaceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceGroupName

`func (o *AzureRegisterCloudSpaceRequest) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureRegisterCloudSpaceRequest) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureRegisterCloudSpaceRequest) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetTenantId

`func (o *AzureRegisterCloudSpaceRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureRegisterCloudSpaceRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureRegisterCloudSpaceRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetVirtualNetworkName

`func (o *AzureRegisterCloudSpaceRequest) GetVirtualNetworkName() string`

GetVirtualNetworkName returns the VirtualNetworkName field if non-nil, zero value otherwise.

### GetVirtualNetworkNameOk

`func (o *AzureRegisterCloudSpaceRequest) GetVirtualNetworkNameOk() (*string, bool)`

GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkName

`func (o *AzureRegisterCloudSpaceRequest) SetVirtualNetworkName(v string)`

SetVirtualNetworkName sets VirtualNetworkName field to given value.


### GetLocalStorageName

`func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageName() string`

GetLocalStorageName returns the LocalStorageName field if non-nil, zero value otherwise.

### GetLocalStorageNameOk

`func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageNameOk() (*string, bool)`

GetLocalStorageNameOk returns a tuple with the LocalStorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageName

`func (o *AzureRegisterCloudSpaceRequest) SetLocalStorageName(v string)`

SetLocalStorageName sets LocalStorageName field to given value.

### HasLocalStorageName

`func (o *AzureRegisterCloudSpaceRequest) HasLocalStorageName() bool`

HasLocalStorageName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


