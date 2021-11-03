# VCloudClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgName** | Pointer to **string** |  | [optional] 
**VdcName** | Pointer to **string** |  | [optional] 

## Methods

### NewVCloudClient

`func NewVCloudClient() *VCloudClient`

NewVCloudClient instantiates a new VCloudClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCloudClientWithDefaults

`func NewVCloudClientWithDefaults() *VCloudClient`

NewVCloudClientWithDefaults instantiates a new VCloudClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgName

`func (o *VCloudClient) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *VCloudClient) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *VCloudClient) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *VCloudClient) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetVdcName

`func (o *VCloudClient) GetVdcName() string`

GetVdcName returns the VdcName field if non-nil, zero value otherwise.

### GetVdcNameOk

`func (o *VCloudClient) GetVdcNameOk() (*string, bool)`

GetVdcNameOk returns a tuple with the VdcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcName

`func (o *VCloudClient) SetVdcName(v string)`

SetVdcName sets VdcName field to given value.

### HasVdcName

`func (o *VCloudClient) HasVdcName() bool`

HasVdcName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


