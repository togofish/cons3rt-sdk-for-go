# AwsVirtualizationRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatImageId** | Pointer to **string** |  | [optional] 
**NatInstanceType** | Pointer to **string** |  | [optional] 
**NatKeyName** | Pointer to **string** |  | [optional] 
**NatKeyPem** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecurityGroupId** | Pointer to **string** |  | [optional] 
**UserKeyName** | Pointer to **string** |  | [optional] 
**UserKeyPem** | Pointer to **string** |  | [optional] 
**VirtualNetworkName** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**VpcSubnetName** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsVirtualizationRealm

`func NewAwsVirtualizationRealm() *AwsVirtualizationRealm`

NewAwsVirtualizationRealm instantiates a new AwsVirtualizationRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVirtualizationRealmWithDefaults

`func NewAwsVirtualizationRealmWithDefaults() *AwsVirtualizationRealm`

NewAwsVirtualizationRealmWithDefaults instantiates a new AwsVirtualizationRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatImageId

`func (o *AwsVirtualizationRealm) GetNatImageId() string`

GetNatImageId returns the NatImageId field if non-nil, zero value otherwise.

### GetNatImageIdOk

`func (o *AwsVirtualizationRealm) GetNatImageIdOk() (*string, bool)`

GetNatImageIdOk returns a tuple with the NatImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatImageId

`func (o *AwsVirtualizationRealm) SetNatImageId(v string)`

SetNatImageId sets NatImageId field to given value.

### HasNatImageId

`func (o *AwsVirtualizationRealm) HasNatImageId() bool`

HasNatImageId returns a boolean if a field has been set.

### GetNatInstanceType

`func (o *AwsVirtualizationRealm) GetNatInstanceType() string`

GetNatInstanceType returns the NatInstanceType field if non-nil, zero value otherwise.

### GetNatInstanceTypeOk

`func (o *AwsVirtualizationRealm) GetNatInstanceTypeOk() (*string, bool)`

GetNatInstanceTypeOk returns a tuple with the NatInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstanceType

`func (o *AwsVirtualizationRealm) SetNatInstanceType(v string)`

SetNatInstanceType sets NatInstanceType field to given value.

### HasNatInstanceType

`func (o *AwsVirtualizationRealm) HasNatInstanceType() bool`

HasNatInstanceType returns a boolean if a field has been set.

### GetNatKeyName

`func (o *AwsVirtualizationRealm) GetNatKeyName() string`

GetNatKeyName returns the NatKeyName field if non-nil, zero value otherwise.

### GetNatKeyNameOk

`func (o *AwsVirtualizationRealm) GetNatKeyNameOk() (*string, bool)`

GetNatKeyNameOk returns a tuple with the NatKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatKeyName

`func (o *AwsVirtualizationRealm) SetNatKeyName(v string)`

SetNatKeyName sets NatKeyName field to given value.

### HasNatKeyName

`func (o *AwsVirtualizationRealm) HasNatKeyName() bool`

HasNatKeyName returns a boolean if a field has been set.

### GetNatKeyPem

`func (o *AwsVirtualizationRealm) GetNatKeyPem() string`

GetNatKeyPem returns the NatKeyPem field if non-nil, zero value otherwise.

### GetNatKeyPemOk

`func (o *AwsVirtualizationRealm) GetNatKeyPemOk() (*string, bool)`

GetNatKeyPemOk returns a tuple with the NatKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatKeyPem

`func (o *AwsVirtualizationRealm) SetNatKeyPem(v string)`

SetNatKeyPem sets NatKeyPem field to given value.

### HasNatKeyPem

`func (o *AwsVirtualizationRealm) HasNatKeyPem() bool`

HasNatKeyPem returns a boolean if a field has been set.

### GetRegion

`func (o *AwsVirtualizationRealm) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsVirtualizationRealm) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsVirtualizationRealm) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsVirtualizationRealm) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *AwsVirtualizationRealm) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AwsVirtualizationRealm) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AwsVirtualizationRealm) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AwsVirtualizationRealm) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetUserKeyName

`func (o *AwsVirtualizationRealm) GetUserKeyName() string`

GetUserKeyName returns the UserKeyName field if non-nil, zero value otherwise.

### GetUserKeyNameOk

`func (o *AwsVirtualizationRealm) GetUserKeyNameOk() (*string, bool)`

GetUserKeyNameOk returns a tuple with the UserKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyName

`func (o *AwsVirtualizationRealm) SetUserKeyName(v string)`

SetUserKeyName sets UserKeyName field to given value.

### HasUserKeyName

`func (o *AwsVirtualizationRealm) HasUserKeyName() bool`

HasUserKeyName returns a boolean if a field has been set.

### GetUserKeyPem

`func (o *AwsVirtualizationRealm) GetUserKeyPem() string`

GetUserKeyPem returns the UserKeyPem field if non-nil, zero value otherwise.

### GetUserKeyPemOk

`func (o *AwsVirtualizationRealm) GetUserKeyPemOk() (*string, bool)`

GetUserKeyPemOk returns a tuple with the UserKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyPem

`func (o *AwsVirtualizationRealm) SetUserKeyPem(v string)`

SetUserKeyPem sets UserKeyPem field to given value.

### HasUserKeyPem

`func (o *AwsVirtualizationRealm) HasUserKeyPem() bool`

HasUserKeyPem returns a boolean if a field has been set.

### GetVirtualNetworkName

`func (o *AwsVirtualizationRealm) GetVirtualNetworkName() string`

GetVirtualNetworkName returns the VirtualNetworkName field if non-nil, zero value otherwise.

### GetVirtualNetworkNameOk

`func (o *AwsVirtualizationRealm) GetVirtualNetworkNameOk() (*string, bool)`

GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkName

`func (o *AwsVirtualizationRealm) SetVirtualNetworkName(v string)`

SetVirtualNetworkName sets VirtualNetworkName field to given value.

### HasVirtualNetworkName

`func (o *AwsVirtualizationRealm) HasVirtualNetworkName() bool`

HasVirtualNetworkName returns a boolean if a field has been set.

### GetVpcId

`func (o *AwsVirtualizationRealm) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsVirtualizationRealm) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsVirtualizationRealm) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsVirtualizationRealm) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetVpcSubnetName

`func (o *AwsVirtualizationRealm) GetVpcSubnetName() string`

GetVpcSubnetName returns the VpcSubnetName field if non-nil, zero value otherwise.

### GetVpcSubnetNameOk

`func (o *AwsVirtualizationRealm) GetVpcSubnetNameOk() (*string, bool)`

GetVpcSubnetNameOk returns a tuple with the VpcSubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcSubnetName

`func (o *AwsVirtualizationRealm) SetVpcSubnetName(v string)`

SetVpcSubnetName sets VpcSubnetName field to given value.

### HasVpcSubnetName

`func (o *AwsVirtualizationRealm) HasVpcSubnetName() bool`

HasVpcSubnetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


