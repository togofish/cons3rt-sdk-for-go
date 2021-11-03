# OpenStackCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**KeystoneHostname** | **string** |  | 
**KeystonePassword** | **string** |  | 
**KeystonePort** | **int32** |  | 
**KeystoneProtocol** | **string** |  | 
**KeystoneUsername** | **string** |  | 
**KeystoneVersion** | **string** |  | 
**NatImageId** | **string** |  | 
**NatInstanceType** | **string** |  | 
**StorageServiceHostname** | Pointer to **string** |  | [optional] 
**StorageServicePort** | Pointer to **int32** |  | [optional] 
**StorageServiceProtocol** | Pointer to **string** |  | [optional] 
**StorageServiceUsername** | Pointer to **string** |  | [optional] 
**StorageServicePassword** | Pointer to **string** |  | [optional] 
**Tenant** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewOpenStackCloud

`func NewOpenStackCloud(keystoneHostname string, keystonePassword string, keystonePort int32, keystoneProtocol string, keystoneUsername string, keystoneVersion string, natImageId string, natInstanceType string, tenant string, tenantId string, ) *OpenStackCloud`

NewOpenStackCloud instantiates a new OpenStackCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenStackCloudWithDefaults

`func NewOpenStackCloudWithDefaults() *OpenStackCloud`

NewOpenStackCloudWithDefaults instantiates a new OpenStackCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *OpenStackCloud) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *OpenStackCloud) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *OpenStackCloud) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *OpenStackCloud) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetKeystoneHostname

`func (o *OpenStackCloud) GetKeystoneHostname() string`

GetKeystoneHostname returns the KeystoneHostname field if non-nil, zero value otherwise.

### GetKeystoneHostnameOk

`func (o *OpenStackCloud) GetKeystoneHostnameOk() (*string, bool)`

GetKeystoneHostnameOk returns a tuple with the KeystoneHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneHostname

`func (o *OpenStackCloud) SetKeystoneHostname(v string)`

SetKeystoneHostname sets KeystoneHostname field to given value.


### GetKeystonePassword

`func (o *OpenStackCloud) GetKeystonePassword() string`

GetKeystonePassword returns the KeystonePassword field if non-nil, zero value otherwise.

### GetKeystonePasswordOk

`func (o *OpenStackCloud) GetKeystonePasswordOk() (*string, bool)`

GetKeystonePasswordOk returns a tuple with the KeystonePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystonePassword

`func (o *OpenStackCloud) SetKeystonePassword(v string)`

SetKeystonePassword sets KeystonePassword field to given value.


### GetKeystonePort

`func (o *OpenStackCloud) GetKeystonePort() int32`

GetKeystonePort returns the KeystonePort field if non-nil, zero value otherwise.

### GetKeystonePortOk

`func (o *OpenStackCloud) GetKeystonePortOk() (*int32, bool)`

GetKeystonePortOk returns a tuple with the KeystonePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystonePort

`func (o *OpenStackCloud) SetKeystonePort(v int32)`

SetKeystonePort sets KeystonePort field to given value.


### GetKeystoneProtocol

`func (o *OpenStackCloud) GetKeystoneProtocol() string`

GetKeystoneProtocol returns the KeystoneProtocol field if non-nil, zero value otherwise.

### GetKeystoneProtocolOk

`func (o *OpenStackCloud) GetKeystoneProtocolOk() (*string, bool)`

GetKeystoneProtocolOk returns a tuple with the KeystoneProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneProtocol

`func (o *OpenStackCloud) SetKeystoneProtocol(v string)`

SetKeystoneProtocol sets KeystoneProtocol field to given value.


### GetKeystoneUsername

`func (o *OpenStackCloud) GetKeystoneUsername() string`

GetKeystoneUsername returns the KeystoneUsername field if non-nil, zero value otherwise.

### GetKeystoneUsernameOk

`func (o *OpenStackCloud) GetKeystoneUsernameOk() (*string, bool)`

GetKeystoneUsernameOk returns a tuple with the KeystoneUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneUsername

`func (o *OpenStackCloud) SetKeystoneUsername(v string)`

SetKeystoneUsername sets KeystoneUsername field to given value.


### GetKeystoneVersion

`func (o *OpenStackCloud) GetKeystoneVersion() string`

GetKeystoneVersion returns the KeystoneVersion field if non-nil, zero value otherwise.

### GetKeystoneVersionOk

`func (o *OpenStackCloud) GetKeystoneVersionOk() (*string, bool)`

GetKeystoneVersionOk returns a tuple with the KeystoneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneVersion

`func (o *OpenStackCloud) SetKeystoneVersion(v string)`

SetKeystoneVersion sets KeystoneVersion field to given value.


### GetNatImageId

`func (o *OpenStackCloud) GetNatImageId() string`

GetNatImageId returns the NatImageId field if non-nil, zero value otherwise.

### GetNatImageIdOk

`func (o *OpenStackCloud) GetNatImageIdOk() (*string, bool)`

GetNatImageIdOk returns a tuple with the NatImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatImageId

`func (o *OpenStackCloud) SetNatImageId(v string)`

SetNatImageId sets NatImageId field to given value.


### GetNatInstanceType

`func (o *OpenStackCloud) GetNatInstanceType() string`

GetNatInstanceType returns the NatInstanceType field if non-nil, zero value otherwise.

### GetNatInstanceTypeOk

`func (o *OpenStackCloud) GetNatInstanceTypeOk() (*string, bool)`

GetNatInstanceTypeOk returns a tuple with the NatInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstanceType

`func (o *OpenStackCloud) SetNatInstanceType(v string)`

SetNatInstanceType sets NatInstanceType field to given value.


### GetStorageServiceHostname

`func (o *OpenStackCloud) GetStorageServiceHostname() string`

GetStorageServiceHostname returns the StorageServiceHostname field if non-nil, zero value otherwise.

### GetStorageServiceHostnameOk

`func (o *OpenStackCloud) GetStorageServiceHostnameOk() (*string, bool)`

GetStorageServiceHostnameOk returns a tuple with the StorageServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceHostname

`func (o *OpenStackCloud) SetStorageServiceHostname(v string)`

SetStorageServiceHostname sets StorageServiceHostname field to given value.

### HasStorageServiceHostname

`func (o *OpenStackCloud) HasStorageServiceHostname() bool`

HasStorageServiceHostname returns a boolean if a field has been set.

### GetStorageServicePort

`func (o *OpenStackCloud) GetStorageServicePort() int32`

GetStorageServicePort returns the StorageServicePort field if non-nil, zero value otherwise.

### GetStorageServicePortOk

`func (o *OpenStackCloud) GetStorageServicePortOk() (*int32, bool)`

GetStorageServicePortOk returns a tuple with the StorageServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServicePort

`func (o *OpenStackCloud) SetStorageServicePort(v int32)`

SetStorageServicePort sets StorageServicePort field to given value.

### HasStorageServicePort

`func (o *OpenStackCloud) HasStorageServicePort() bool`

HasStorageServicePort returns a boolean if a field has been set.

### GetStorageServiceProtocol

`func (o *OpenStackCloud) GetStorageServiceProtocol() string`

GetStorageServiceProtocol returns the StorageServiceProtocol field if non-nil, zero value otherwise.

### GetStorageServiceProtocolOk

`func (o *OpenStackCloud) GetStorageServiceProtocolOk() (*string, bool)`

GetStorageServiceProtocolOk returns a tuple with the StorageServiceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceProtocol

`func (o *OpenStackCloud) SetStorageServiceProtocol(v string)`

SetStorageServiceProtocol sets StorageServiceProtocol field to given value.

### HasStorageServiceProtocol

`func (o *OpenStackCloud) HasStorageServiceProtocol() bool`

HasStorageServiceProtocol returns a boolean if a field has been set.

### GetStorageServiceUsername

`func (o *OpenStackCloud) GetStorageServiceUsername() string`

GetStorageServiceUsername returns the StorageServiceUsername field if non-nil, zero value otherwise.

### GetStorageServiceUsernameOk

`func (o *OpenStackCloud) GetStorageServiceUsernameOk() (*string, bool)`

GetStorageServiceUsernameOk returns a tuple with the StorageServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceUsername

`func (o *OpenStackCloud) SetStorageServiceUsername(v string)`

SetStorageServiceUsername sets StorageServiceUsername field to given value.

### HasStorageServiceUsername

`func (o *OpenStackCloud) HasStorageServiceUsername() bool`

HasStorageServiceUsername returns a boolean if a field has been set.

### GetStorageServicePassword

`func (o *OpenStackCloud) GetStorageServicePassword() string`

GetStorageServicePassword returns the StorageServicePassword field if non-nil, zero value otherwise.

### GetStorageServicePasswordOk

`func (o *OpenStackCloud) GetStorageServicePasswordOk() (*string, bool)`

GetStorageServicePasswordOk returns a tuple with the StorageServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServicePassword

`func (o *OpenStackCloud) SetStorageServicePassword(v string)`

SetStorageServicePassword sets StorageServicePassword field to given value.

### HasStorageServicePassword

`func (o *OpenStackCloud) HasStorageServicePassword() bool`

HasStorageServicePassword returns a boolean if a field has been set.

### GetTenant

`func (o *OpenStackCloud) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *OpenStackCloud) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *OpenStackCloud) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTenantId

`func (o *OpenStackCloud) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OpenStackCloud) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OpenStackCloud) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


