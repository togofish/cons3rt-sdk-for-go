# AwsRegisterCloudSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Region** | **string** |  | 
**VpcId** | **string** |  | 
**UserKeyName** | Pointer to **string** |  | [optional] 
**NetworkSecurityGroupMap** | **map[string]string** |  | 

## Methods

### NewAwsRegisterCloudSpaceRequest

`func NewAwsRegisterCloudSpaceRequest(accountId string, region string, vpcId string, networkSecurityGroupMap map[string]string, ) *AwsRegisterCloudSpaceRequest`

NewAwsRegisterCloudSpaceRequest instantiates a new AwsRegisterCloudSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRegisterCloudSpaceRequestWithDefaults

`func NewAwsRegisterCloudSpaceRequestWithDefaults() *AwsRegisterCloudSpaceRequest`

NewAwsRegisterCloudSpaceRequestWithDefaults instantiates a new AwsRegisterCloudSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AwsRegisterCloudSpaceRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsRegisterCloudSpaceRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsRegisterCloudSpaceRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetRegion

`func (o *AwsRegisterCloudSpaceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsRegisterCloudSpaceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsRegisterCloudSpaceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVpcId

`func (o *AwsRegisterCloudSpaceRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsRegisterCloudSpaceRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsRegisterCloudSpaceRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetUserKeyName

`func (o *AwsRegisterCloudSpaceRequest) GetUserKeyName() string`

GetUserKeyName returns the UserKeyName field if non-nil, zero value otherwise.

### GetUserKeyNameOk

`func (o *AwsRegisterCloudSpaceRequest) GetUserKeyNameOk() (*string, bool)`

GetUserKeyNameOk returns a tuple with the UserKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyName

`func (o *AwsRegisterCloudSpaceRequest) SetUserKeyName(v string)`

SetUserKeyName sets UserKeyName field to given value.

### HasUserKeyName

`func (o *AwsRegisterCloudSpaceRequest) HasUserKeyName() bool`

HasUserKeyName returns a boolean if a field has been set.

### GetNetworkSecurityGroupMap

`func (o *AwsRegisterCloudSpaceRequest) GetNetworkSecurityGroupMap() map[string]string`

GetNetworkSecurityGroupMap returns the NetworkSecurityGroupMap field if non-nil, zero value otherwise.

### GetNetworkSecurityGroupMapOk

`func (o *AwsRegisterCloudSpaceRequest) GetNetworkSecurityGroupMapOk() (*map[string]string, bool)`

GetNetworkSecurityGroupMapOk returns a tuple with the NetworkSecurityGroupMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurityGroupMap

`func (o *AwsRegisterCloudSpaceRequest) SetNetworkSecurityGroupMap(v map[string]string)`

SetNetworkSecurityGroupMap sets NetworkSecurityGroupMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


