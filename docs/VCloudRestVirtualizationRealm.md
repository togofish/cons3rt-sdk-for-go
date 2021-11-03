# VCloudRestVirtualizationRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationVdc** | Pointer to **string** |  | [optional] 
**VdcNetworkQuota** | Pointer to **int32** |  | [optional] 

## Methods

### NewVCloudRestVirtualizationRealm

`func NewVCloudRestVirtualizationRealm() *VCloudRestVirtualizationRealm`

NewVCloudRestVirtualizationRealm instantiates a new VCloudRestVirtualizationRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCloudRestVirtualizationRealmWithDefaults

`func NewVCloudRestVirtualizationRealmWithDefaults() *VCloudRestVirtualizationRealm`

NewVCloudRestVirtualizationRealmWithDefaults instantiates a new VCloudRestVirtualizationRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *VCloudRestVirtualizationRealm) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VCloudRestVirtualizationRealm) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VCloudRestVirtualizationRealm) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VCloudRestVirtualizationRealm) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationVdc

`func (o *VCloudRestVirtualizationRealm) GetOrganizationVdc() string`

GetOrganizationVdc returns the OrganizationVdc field if non-nil, zero value otherwise.

### GetOrganizationVdcOk

`func (o *VCloudRestVirtualizationRealm) GetOrganizationVdcOk() (*string, bool)`

GetOrganizationVdcOk returns a tuple with the OrganizationVdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationVdc

`func (o *VCloudRestVirtualizationRealm) SetOrganizationVdc(v string)`

SetOrganizationVdc sets OrganizationVdc field to given value.

### HasOrganizationVdc

`func (o *VCloudRestVirtualizationRealm) HasOrganizationVdc() bool`

HasOrganizationVdc returns a boolean if a field has been set.

### GetVdcNetworkQuota

`func (o *VCloudRestVirtualizationRealm) GetVdcNetworkQuota() int32`

GetVdcNetworkQuota returns the VdcNetworkQuota field if non-nil, zero value otherwise.

### GetVdcNetworkQuotaOk

`func (o *VCloudRestVirtualizationRealm) GetVdcNetworkQuotaOk() (*int32, bool)`

GetVdcNetworkQuotaOk returns a tuple with the VdcNetworkQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcNetworkQuota

`func (o *VCloudRestVirtualizationRealm) SetVdcNetworkQuota(v int32)`

SetVdcNetworkQuota sets VdcNetworkQuota field to given value.

### HasVdcNetworkQuota

`func (o *VCloudRestVirtualizationRealm) HasVdcNetworkQuota() bool`

HasVdcNetworkQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


