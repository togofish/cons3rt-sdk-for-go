# AwsClientConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionToken** | Pointer to **string** |  | [optional] 
**RegionName** | **string** |  | 

## Methods

### NewAwsClientConfiguration

`func NewAwsClientConfiguration(regionName string, ) *AwsClientConfiguration`

NewAwsClientConfiguration instantiates a new AwsClientConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsClientConfigurationWithDefaults

`func NewAwsClientConfigurationWithDefaults() *AwsClientConfiguration`

NewAwsClientConfigurationWithDefaults instantiates a new AwsClientConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionToken

`func (o *AwsClientConfiguration) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *AwsClientConfiguration) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *AwsClientConfiguration) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *AwsClientConfiguration) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetRegionName

`func (o *AwsClientConfiguration) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *AwsClientConfiguration) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *AwsClientConfiguration) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


