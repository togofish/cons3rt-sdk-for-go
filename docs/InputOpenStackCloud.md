# InputOpenStackCloud

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
**Tenant** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewInputOpenStackCloud

`func NewInputOpenStackCloud(keystoneHostname string, keystonePassword string, keystonePort int32, keystoneProtocol string, keystoneUsername string, keystoneVersion string, natImageId string, natInstanceType string, tenant string, tenantId string, ) *InputOpenStackCloud`

NewInputOpenStackCloud instantiates a new InputOpenStackCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputOpenStackCloudWithDefaults

`func NewInputOpenStackCloudWithDefaults() *InputOpenStackCloud`

NewInputOpenStackCloudWithDefaults instantiates a new InputOpenStackCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *InputOpenStackCloud) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *InputOpenStackCloud) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *InputOpenStackCloud) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *InputOpenStackCloud) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetKeystoneHostname

`func (o *InputOpenStackCloud) GetKeystoneHostname() string`

GetKeystoneHostname returns the KeystoneHostname field if non-nil, zero value otherwise.

### GetKeystoneHostnameOk

`func (o *InputOpenStackCloud) GetKeystoneHostnameOk() (*string, bool)`

GetKeystoneHostnameOk returns a tuple with the KeystoneHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneHostname

`func (o *InputOpenStackCloud) SetKeystoneHostname(v string)`

SetKeystoneHostname sets KeystoneHostname field to given value.


### GetKeystonePassword

`func (o *InputOpenStackCloud) GetKeystonePassword() string`

GetKeystonePassword returns the KeystonePassword field if non-nil, zero value otherwise.

### GetKeystonePasswordOk

`func (o *InputOpenStackCloud) GetKeystonePasswordOk() (*string, bool)`

GetKeystonePasswordOk returns a tuple with the KeystonePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystonePassword

`func (o *InputOpenStackCloud) SetKeystonePassword(v string)`

SetKeystonePassword sets KeystonePassword field to given value.


### GetKeystonePort

`func (o *InputOpenStackCloud) GetKeystonePort() int32`

GetKeystonePort returns the KeystonePort field if non-nil, zero value otherwise.

### GetKeystonePortOk

`func (o *InputOpenStackCloud) GetKeystonePortOk() (*int32, bool)`

GetKeystonePortOk returns a tuple with the KeystonePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystonePort

`func (o *InputOpenStackCloud) SetKeystonePort(v int32)`

SetKeystonePort sets KeystonePort field to given value.


### GetKeystoneProtocol

`func (o *InputOpenStackCloud) GetKeystoneProtocol() string`

GetKeystoneProtocol returns the KeystoneProtocol field if non-nil, zero value otherwise.

### GetKeystoneProtocolOk

`func (o *InputOpenStackCloud) GetKeystoneProtocolOk() (*string, bool)`

GetKeystoneProtocolOk returns a tuple with the KeystoneProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneProtocol

`func (o *InputOpenStackCloud) SetKeystoneProtocol(v string)`

SetKeystoneProtocol sets KeystoneProtocol field to given value.


### GetKeystoneUsername

`func (o *InputOpenStackCloud) GetKeystoneUsername() string`

GetKeystoneUsername returns the KeystoneUsername field if non-nil, zero value otherwise.

### GetKeystoneUsernameOk

`func (o *InputOpenStackCloud) GetKeystoneUsernameOk() (*string, bool)`

GetKeystoneUsernameOk returns a tuple with the KeystoneUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneUsername

`func (o *InputOpenStackCloud) SetKeystoneUsername(v string)`

SetKeystoneUsername sets KeystoneUsername field to given value.


### GetKeystoneVersion

`func (o *InputOpenStackCloud) GetKeystoneVersion() string`

GetKeystoneVersion returns the KeystoneVersion field if non-nil, zero value otherwise.

### GetKeystoneVersionOk

`func (o *InputOpenStackCloud) GetKeystoneVersionOk() (*string, bool)`

GetKeystoneVersionOk returns a tuple with the KeystoneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneVersion

`func (o *InputOpenStackCloud) SetKeystoneVersion(v string)`

SetKeystoneVersion sets KeystoneVersion field to given value.


### GetNatImageId

`func (o *InputOpenStackCloud) GetNatImageId() string`

GetNatImageId returns the NatImageId field if non-nil, zero value otherwise.

### GetNatImageIdOk

`func (o *InputOpenStackCloud) GetNatImageIdOk() (*string, bool)`

GetNatImageIdOk returns a tuple with the NatImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatImageId

`func (o *InputOpenStackCloud) SetNatImageId(v string)`

SetNatImageId sets NatImageId field to given value.


### GetNatInstanceType

`func (o *InputOpenStackCloud) GetNatInstanceType() string`

GetNatInstanceType returns the NatInstanceType field if non-nil, zero value otherwise.

### GetNatInstanceTypeOk

`func (o *InputOpenStackCloud) GetNatInstanceTypeOk() (*string, bool)`

GetNatInstanceTypeOk returns a tuple with the NatInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstanceType

`func (o *InputOpenStackCloud) SetNatInstanceType(v string)`

SetNatInstanceType sets NatInstanceType field to given value.


### GetTenant

`func (o *InputOpenStackCloud) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *InputOpenStackCloud) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *InputOpenStackCloud) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTenantId

`func (o *InputOpenStackCloud) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InputOpenStackCloud) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InputOpenStackCloud) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


