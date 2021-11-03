# FullOpenStackCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**KeystoneHostname** | Pointer to **string** |  | [optional] 
**KeystonePort** | Pointer to **int32** |  | [optional] 
**KeystoneProtocol** | Pointer to **string** |  | [optional] 
**KeystoneUsername** | Pointer to **string** |  | [optional] 
**KeystoneVersion** | Pointer to **string** |  | [optional] 
**NatImageId** | **string** |  | 
**NatInstanceType** | **string** |  | 
**Tenant** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewFullOpenStackCloud

`func NewFullOpenStackCloud(natImageId string, natInstanceType string, tenant string, tenantId string, ) *FullOpenStackCloud`

NewFullOpenStackCloud instantiates a new FullOpenStackCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullOpenStackCloudWithDefaults

`func NewFullOpenStackCloudWithDefaults() *FullOpenStackCloud`

NewFullOpenStackCloudWithDefaults instantiates a new FullOpenStackCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *FullOpenStackCloud) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *FullOpenStackCloud) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *FullOpenStackCloud) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *FullOpenStackCloud) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetKeystoneHostname

`func (o *FullOpenStackCloud) GetKeystoneHostname() string`

GetKeystoneHostname returns the KeystoneHostname field if non-nil, zero value otherwise.

### GetKeystoneHostnameOk

`func (o *FullOpenStackCloud) GetKeystoneHostnameOk() (*string, bool)`

GetKeystoneHostnameOk returns a tuple with the KeystoneHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneHostname

`func (o *FullOpenStackCloud) SetKeystoneHostname(v string)`

SetKeystoneHostname sets KeystoneHostname field to given value.

### HasKeystoneHostname

`func (o *FullOpenStackCloud) HasKeystoneHostname() bool`

HasKeystoneHostname returns a boolean if a field has been set.

### GetKeystonePort

`func (o *FullOpenStackCloud) GetKeystonePort() int32`

GetKeystonePort returns the KeystonePort field if non-nil, zero value otherwise.

### GetKeystonePortOk

`func (o *FullOpenStackCloud) GetKeystonePortOk() (*int32, bool)`

GetKeystonePortOk returns a tuple with the KeystonePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystonePort

`func (o *FullOpenStackCloud) SetKeystonePort(v int32)`

SetKeystonePort sets KeystonePort field to given value.

### HasKeystonePort

`func (o *FullOpenStackCloud) HasKeystonePort() bool`

HasKeystonePort returns a boolean if a field has been set.

### GetKeystoneProtocol

`func (o *FullOpenStackCloud) GetKeystoneProtocol() string`

GetKeystoneProtocol returns the KeystoneProtocol field if non-nil, zero value otherwise.

### GetKeystoneProtocolOk

`func (o *FullOpenStackCloud) GetKeystoneProtocolOk() (*string, bool)`

GetKeystoneProtocolOk returns a tuple with the KeystoneProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneProtocol

`func (o *FullOpenStackCloud) SetKeystoneProtocol(v string)`

SetKeystoneProtocol sets KeystoneProtocol field to given value.

### HasKeystoneProtocol

`func (o *FullOpenStackCloud) HasKeystoneProtocol() bool`

HasKeystoneProtocol returns a boolean if a field has been set.

### GetKeystoneUsername

`func (o *FullOpenStackCloud) GetKeystoneUsername() string`

GetKeystoneUsername returns the KeystoneUsername field if non-nil, zero value otherwise.

### GetKeystoneUsernameOk

`func (o *FullOpenStackCloud) GetKeystoneUsernameOk() (*string, bool)`

GetKeystoneUsernameOk returns a tuple with the KeystoneUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneUsername

`func (o *FullOpenStackCloud) SetKeystoneUsername(v string)`

SetKeystoneUsername sets KeystoneUsername field to given value.

### HasKeystoneUsername

`func (o *FullOpenStackCloud) HasKeystoneUsername() bool`

HasKeystoneUsername returns a boolean if a field has been set.

### GetKeystoneVersion

`func (o *FullOpenStackCloud) GetKeystoneVersion() string`

GetKeystoneVersion returns the KeystoneVersion field if non-nil, zero value otherwise.

### GetKeystoneVersionOk

`func (o *FullOpenStackCloud) GetKeystoneVersionOk() (*string, bool)`

GetKeystoneVersionOk returns a tuple with the KeystoneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneVersion

`func (o *FullOpenStackCloud) SetKeystoneVersion(v string)`

SetKeystoneVersion sets KeystoneVersion field to given value.

### HasKeystoneVersion

`func (o *FullOpenStackCloud) HasKeystoneVersion() bool`

HasKeystoneVersion returns a boolean if a field has been set.

### GetNatImageId

`func (o *FullOpenStackCloud) GetNatImageId() string`

GetNatImageId returns the NatImageId field if non-nil, zero value otherwise.

### GetNatImageIdOk

`func (o *FullOpenStackCloud) GetNatImageIdOk() (*string, bool)`

GetNatImageIdOk returns a tuple with the NatImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatImageId

`func (o *FullOpenStackCloud) SetNatImageId(v string)`

SetNatImageId sets NatImageId field to given value.


### GetNatInstanceType

`func (o *FullOpenStackCloud) GetNatInstanceType() string`

GetNatInstanceType returns the NatInstanceType field if non-nil, zero value otherwise.

### GetNatInstanceTypeOk

`func (o *FullOpenStackCloud) GetNatInstanceTypeOk() (*string, bool)`

GetNatInstanceTypeOk returns a tuple with the NatInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstanceType

`func (o *FullOpenStackCloud) SetNatInstanceType(v string)`

SetNatInstanceType sets NatInstanceType field to given value.


### GetTenant

`func (o *FullOpenStackCloud) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FullOpenStackCloud) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FullOpenStackCloud) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTenantId

`func (o *FullOpenStackCloud) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FullOpenStackCloud) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FullOpenStackCloud) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


