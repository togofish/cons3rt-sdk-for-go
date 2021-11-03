# ProjectLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuTypeMaximums** | Pointer to **map[string]int32** |  | [optional] 
**MaxNumCpus** | **int32** |  | 
**MaxNumGpus** | **int32** |  | 
**MaxRamInMegabytes** | **int32** |  | 
**MaxStorageInMegabytes** | **int32** |  | 
**MaxVirtualMachines** | **int32** |  | 
**ValidUntil** | Pointer to **int32** |  | [optional] 

## Methods

### NewProjectLimits

`func NewProjectLimits(maxNumCpus int32, maxNumGpus int32, maxRamInMegabytes int32, maxStorageInMegabytes int32, maxVirtualMachines int32, ) *ProjectLimits`

NewProjectLimits instantiates a new ProjectLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectLimitsWithDefaults

`func NewProjectLimitsWithDefaults() *ProjectLimits`

NewProjectLimitsWithDefaults instantiates a new ProjectLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuTypeMaximums

`func (o *ProjectLimits) GetGpuTypeMaximums() map[string]int32`

GetGpuTypeMaximums returns the GpuTypeMaximums field if non-nil, zero value otherwise.

### GetGpuTypeMaximumsOk

`func (o *ProjectLimits) GetGpuTypeMaximumsOk() (*map[string]int32, bool)`

GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypeMaximums

`func (o *ProjectLimits) SetGpuTypeMaximums(v map[string]int32)`

SetGpuTypeMaximums sets GpuTypeMaximums field to given value.

### HasGpuTypeMaximums

`func (o *ProjectLimits) HasGpuTypeMaximums() bool`

HasGpuTypeMaximums returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *ProjectLimits) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *ProjectLimits) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *ProjectLimits) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.


### GetMaxNumGpus

`func (o *ProjectLimits) GetMaxNumGpus() int32`

GetMaxNumGpus returns the MaxNumGpus field if non-nil, zero value otherwise.

### GetMaxNumGpusOk

`func (o *ProjectLimits) GetMaxNumGpusOk() (*int32, bool)`

GetMaxNumGpusOk returns a tuple with the MaxNumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumGpus

`func (o *ProjectLimits) SetMaxNumGpus(v int32)`

SetMaxNumGpus sets MaxNumGpus field to given value.


### GetMaxRamInMegabytes

`func (o *ProjectLimits) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *ProjectLimits) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *ProjectLimits) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.


### GetMaxStorageInMegabytes

`func (o *ProjectLimits) GetMaxStorageInMegabytes() int32`

GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field if non-nil, zero value otherwise.

### GetMaxStorageInMegabytesOk

`func (o *ProjectLimits) GetMaxStorageInMegabytesOk() (*int32, bool)`

GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorageInMegabytes

`func (o *ProjectLimits) SetMaxStorageInMegabytes(v int32)`

SetMaxStorageInMegabytes sets MaxStorageInMegabytes field to given value.


### GetMaxVirtualMachines

`func (o *ProjectLimits) GetMaxVirtualMachines() int32`

GetMaxVirtualMachines returns the MaxVirtualMachines field if non-nil, zero value otherwise.

### GetMaxVirtualMachinesOk

`func (o *ProjectLimits) GetMaxVirtualMachinesOk() (*int32, bool)`

GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualMachines

`func (o *ProjectLimits) SetMaxVirtualMachines(v int32)`

SetMaxVirtualMachines sets MaxVirtualMachines field to given value.


### GetValidUntil

`func (o *ProjectLimits) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ProjectLimits) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ProjectLimits) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *ProjectLimits) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


