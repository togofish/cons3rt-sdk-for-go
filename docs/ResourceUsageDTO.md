# ResourceUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuTypeUsage** | Pointer to **map[string]int32** |  | [optional] 
**NumCpus** | Pointer to **int32** |  | [optional] 
**RamInMegabytes** | Pointer to **int32** |  | [optional] 
**StorageInMegabytes** | Pointer to **int32** |  | [optional] 
**VirtualMachines** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourceUsageDTO

`func NewResourceUsageDTO() *ResourceUsageDTO`

NewResourceUsageDTO instantiates a new ResourceUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageDTOWithDefaults

`func NewResourceUsageDTOWithDefaults() *ResourceUsageDTO`

NewResourceUsageDTOWithDefaults instantiates a new ResourceUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuTypeUsage

`func (o *ResourceUsageDTO) GetGpuTypeUsage() map[string]int32`

GetGpuTypeUsage returns the GpuTypeUsage field if non-nil, zero value otherwise.

### GetGpuTypeUsageOk

`func (o *ResourceUsageDTO) GetGpuTypeUsageOk() (*map[string]int32, bool)`

GetGpuTypeUsageOk returns a tuple with the GpuTypeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypeUsage

`func (o *ResourceUsageDTO) SetGpuTypeUsage(v map[string]int32)`

SetGpuTypeUsage sets GpuTypeUsage field to given value.

### HasGpuTypeUsage

`func (o *ResourceUsageDTO) HasGpuTypeUsage() bool`

HasGpuTypeUsage returns a boolean if a field has been set.

### GetNumCpus

`func (o *ResourceUsageDTO) GetNumCpus() int32`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ResourceUsageDTO) GetNumCpusOk() (*int32, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ResourceUsageDTO) SetNumCpus(v int32)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ResourceUsageDTO) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetRamInMegabytes

`func (o *ResourceUsageDTO) GetRamInMegabytes() int32`

GetRamInMegabytes returns the RamInMegabytes field if non-nil, zero value otherwise.

### GetRamInMegabytesOk

`func (o *ResourceUsageDTO) GetRamInMegabytesOk() (*int32, bool)`

GetRamInMegabytesOk returns a tuple with the RamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamInMegabytes

`func (o *ResourceUsageDTO) SetRamInMegabytes(v int32)`

SetRamInMegabytes sets RamInMegabytes field to given value.

### HasRamInMegabytes

`func (o *ResourceUsageDTO) HasRamInMegabytes() bool`

HasRamInMegabytes returns a boolean if a field has been set.

### GetStorageInMegabytes

`func (o *ResourceUsageDTO) GetStorageInMegabytes() int32`

GetStorageInMegabytes returns the StorageInMegabytes field if non-nil, zero value otherwise.

### GetStorageInMegabytesOk

`func (o *ResourceUsageDTO) GetStorageInMegabytesOk() (*int32, bool)`

GetStorageInMegabytesOk returns a tuple with the StorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInMegabytes

`func (o *ResourceUsageDTO) SetStorageInMegabytes(v int32)`

SetStorageInMegabytes sets StorageInMegabytes field to given value.

### HasStorageInMegabytes

`func (o *ResourceUsageDTO) HasStorageInMegabytes() bool`

HasStorageInMegabytes returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *ResourceUsageDTO) GetVirtualMachines() int32`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *ResourceUsageDTO) GetVirtualMachinesOk() (*int32, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *ResourceUsageDTO) SetVirtualMachines(v int32)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *ResourceUsageDTO) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


