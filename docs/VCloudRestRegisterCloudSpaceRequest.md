# VCloudRestRegisterCloudSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStorageName** | Pointer to **string** |  | [optional] 
**Organization** | **string** |  | 
**OrganizationVdc** | **string** |  | 

## Methods

### NewVCloudRestRegisterCloudSpaceRequest

`func NewVCloudRestRegisterCloudSpaceRequest(organization string, organizationVdc string, ) *VCloudRestRegisterCloudSpaceRequest`

NewVCloudRestRegisterCloudSpaceRequest instantiates a new VCloudRestRegisterCloudSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCloudRestRegisterCloudSpaceRequestWithDefaults

`func NewVCloudRestRegisterCloudSpaceRequestWithDefaults() *VCloudRestRegisterCloudSpaceRequest`

NewVCloudRestRegisterCloudSpaceRequestWithDefaults instantiates a new VCloudRestRegisterCloudSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStorageName

`func (o *VCloudRestRegisterCloudSpaceRequest) GetLocalStorageName() string`

GetLocalStorageName returns the LocalStorageName field if non-nil, zero value otherwise.

### GetLocalStorageNameOk

`func (o *VCloudRestRegisterCloudSpaceRequest) GetLocalStorageNameOk() (*string, bool)`

GetLocalStorageNameOk returns a tuple with the LocalStorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageName

`func (o *VCloudRestRegisterCloudSpaceRequest) SetLocalStorageName(v string)`

SetLocalStorageName sets LocalStorageName field to given value.

### HasLocalStorageName

`func (o *VCloudRestRegisterCloudSpaceRequest) HasLocalStorageName() bool`

HasLocalStorageName returns a boolean if a field has been set.

### GetOrganization

`func (o *VCloudRestRegisterCloudSpaceRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VCloudRestRegisterCloudSpaceRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VCloudRestRegisterCloudSpaceRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationVdc

`func (o *VCloudRestRegisterCloudSpaceRequest) GetOrganizationVdc() string`

GetOrganizationVdc returns the OrganizationVdc field if non-nil, zero value otherwise.

### GetOrganizationVdcOk

`func (o *VCloudRestRegisterCloudSpaceRequest) GetOrganizationVdcOk() (*string, bool)`

GetOrganizationVdcOk returns a tuple with the OrganizationVdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationVdc

`func (o *VCloudRestRegisterCloudSpaceRequest) SetOrganizationVdc(v string)`

SetOrganizationVdc sets OrganizationVdc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


