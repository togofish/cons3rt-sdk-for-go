# OpenStackCloudSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatFlavorType** | **string** |  | 
**NatImageId** | **string** |  | 
**DnsServerIpAddresses** | Pointer to **string** |  | [optional] 
**ExternalNetworkName** | **string** |  | 

## Methods

### NewOpenStackCloudSpaceRequest

`func NewOpenStackCloudSpaceRequest(natFlavorType string, natImageId string, externalNetworkName string, ) *OpenStackCloudSpaceRequest`

NewOpenStackCloudSpaceRequest instantiates a new OpenStackCloudSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenStackCloudSpaceRequestWithDefaults

`func NewOpenStackCloudSpaceRequestWithDefaults() *OpenStackCloudSpaceRequest`

NewOpenStackCloudSpaceRequestWithDefaults instantiates a new OpenStackCloudSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatFlavorType

`func (o *OpenStackCloudSpaceRequest) GetNatFlavorType() string`

GetNatFlavorType returns the NatFlavorType field if non-nil, zero value otherwise.

### GetNatFlavorTypeOk

`func (o *OpenStackCloudSpaceRequest) GetNatFlavorTypeOk() (*string, bool)`

GetNatFlavorTypeOk returns a tuple with the NatFlavorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatFlavorType

`func (o *OpenStackCloudSpaceRequest) SetNatFlavorType(v string)`

SetNatFlavorType sets NatFlavorType field to given value.


### GetNatImageId

`func (o *OpenStackCloudSpaceRequest) GetNatImageId() string`

GetNatImageId returns the NatImageId field if non-nil, zero value otherwise.

### GetNatImageIdOk

`func (o *OpenStackCloudSpaceRequest) GetNatImageIdOk() (*string, bool)`

GetNatImageIdOk returns a tuple with the NatImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatImageId

`func (o *OpenStackCloudSpaceRequest) SetNatImageId(v string)`

SetNatImageId sets NatImageId field to given value.


### GetDnsServerIpAddresses

`func (o *OpenStackCloudSpaceRequest) GetDnsServerIpAddresses() string`

GetDnsServerIpAddresses returns the DnsServerIpAddresses field if non-nil, zero value otherwise.

### GetDnsServerIpAddressesOk

`func (o *OpenStackCloudSpaceRequest) GetDnsServerIpAddressesOk() (*string, bool)`

GetDnsServerIpAddressesOk returns a tuple with the DnsServerIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerIpAddresses

`func (o *OpenStackCloudSpaceRequest) SetDnsServerIpAddresses(v string)`

SetDnsServerIpAddresses sets DnsServerIpAddresses field to given value.

### HasDnsServerIpAddresses

`func (o *OpenStackCloudSpaceRequest) HasDnsServerIpAddresses() bool`

HasDnsServerIpAddresses returns a boolean if a field has been set.

### GetExternalNetworkName

`func (o *OpenStackCloudSpaceRequest) GetExternalNetworkName() string`

GetExternalNetworkName returns the ExternalNetworkName field if non-nil, zero value otherwise.

### GetExternalNetworkNameOk

`func (o *OpenStackCloudSpaceRequest) GetExternalNetworkNameOk() (*string, bool)`

GetExternalNetworkNameOk returns a tuple with the ExternalNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetworkName

`func (o *OpenStackCloudSpaceRequest) SetExternalNetworkName(v string)`

SetExternalNetworkName sets ExternalNetworkName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


