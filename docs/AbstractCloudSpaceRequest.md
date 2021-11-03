# AbstractCloudSpaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudSpaceName** | **string** |  | 
**MaximumVirtualMachines** | Pointer to **int32** |  | [optional] 
**Cidr** | **string** |  | 
**NumAvailabilityZones** | Pointer to **int32** |  | [optional] 
**PowerOnMaximumDelay** | Pointer to **int32** |  | [optional] 
**PowerOnMinimumDelay** | Pointer to **int32** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewAbstractCloudSpaceRequest

`func NewAbstractCloudSpaceRequest(cloudSpaceName string, cidr string, subtype string, ) *AbstractCloudSpaceRequest`

NewAbstractCloudSpaceRequest instantiates a new AbstractCloudSpaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractCloudSpaceRequestWithDefaults

`func NewAbstractCloudSpaceRequestWithDefaults() *AbstractCloudSpaceRequest`

NewAbstractCloudSpaceRequestWithDefaults instantiates a new AbstractCloudSpaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudSpaceName

`func (o *AbstractCloudSpaceRequest) GetCloudSpaceName() string`

GetCloudSpaceName returns the CloudSpaceName field if non-nil, zero value otherwise.

### GetCloudSpaceNameOk

`func (o *AbstractCloudSpaceRequest) GetCloudSpaceNameOk() (*string, bool)`

GetCloudSpaceNameOk returns a tuple with the CloudSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpaceName

`func (o *AbstractCloudSpaceRequest) SetCloudSpaceName(v string)`

SetCloudSpaceName sets CloudSpaceName field to given value.


### GetMaximumVirtualMachines

`func (o *AbstractCloudSpaceRequest) GetMaximumVirtualMachines() int32`

GetMaximumVirtualMachines returns the MaximumVirtualMachines field if non-nil, zero value otherwise.

### GetMaximumVirtualMachinesOk

`func (o *AbstractCloudSpaceRequest) GetMaximumVirtualMachinesOk() (*int32, bool)`

GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualMachines

`func (o *AbstractCloudSpaceRequest) SetMaximumVirtualMachines(v int32)`

SetMaximumVirtualMachines sets MaximumVirtualMachines field to given value.

### HasMaximumVirtualMachines

`func (o *AbstractCloudSpaceRequest) HasMaximumVirtualMachines() bool`

HasMaximumVirtualMachines returns a boolean if a field has been set.

### GetCidr

`func (o *AbstractCloudSpaceRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *AbstractCloudSpaceRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *AbstractCloudSpaceRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetNumAvailabilityZones

`func (o *AbstractCloudSpaceRequest) GetNumAvailabilityZones() int32`

GetNumAvailabilityZones returns the NumAvailabilityZones field if non-nil, zero value otherwise.

### GetNumAvailabilityZonesOk

`func (o *AbstractCloudSpaceRequest) GetNumAvailabilityZonesOk() (*int32, bool)`

GetNumAvailabilityZonesOk returns a tuple with the NumAvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAvailabilityZones

`func (o *AbstractCloudSpaceRequest) SetNumAvailabilityZones(v int32)`

SetNumAvailabilityZones sets NumAvailabilityZones field to given value.

### HasNumAvailabilityZones

`func (o *AbstractCloudSpaceRequest) HasNumAvailabilityZones() bool`

HasNumAvailabilityZones returns a boolean if a field has been set.

### GetPowerOnMaximumDelay

`func (o *AbstractCloudSpaceRequest) GetPowerOnMaximumDelay() int32`

GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field if non-nil, zero value otherwise.

### GetPowerOnMaximumDelayOk

`func (o *AbstractCloudSpaceRequest) GetPowerOnMaximumDelayOk() (*int32, bool)`

GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMaximumDelay

`func (o *AbstractCloudSpaceRequest) SetPowerOnMaximumDelay(v int32)`

SetPowerOnMaximumDelay sets PowerOnMaximumDelay field to given value.

### HasPowerOnMaximumDelay

`func (o *AbstractCloudSpaceRequest) HasPowerOnMaximumDelay() bool`

HasPowerOnMaximumDelay returns a boolean if a field has been set.

### GetPowerOnMinimumDelay

`func (o *AbstractCloudSpaceRequest) GetPowerOnMinimumDelay() int32`

GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field if non-nil, zero value otherwise.

### GetPowerOnMinimumDelayOk

`func (o *AbstractCloudSpaceRequest) GetPowerOnMinimumDelayOk() (*int32, bool)`

GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnMinimumDelay

`func (o *AbstractCloudSpaceRequest) SetPowerOnMinimumDelay(v int32)`

SetPowerOnMinimumDelay sets PowerOnMinimumDelay field to given value.

### HasPowerOnMinimumDelay

`func (o *AbstractCloudSpaceRequest) HasPowerOnMinimumDelay() bool`

HasPowerOnMinimumDelay returns a boolean if a field has been set.

### GetSubtype

`func (o *AbstractCloudSpaceRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AbstractCloudSpaceRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AbstractCloudSpaceRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


