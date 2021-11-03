# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDisk** | Pointer to **bool** |  | [optional] 
**BootDisk** | Pointer to **bool** |  | [optional] 
**CapacityInMegabytes** | **int32** |  | 
**CreateOrder** | Pointer to **int32** |  | [optional] 
**IsAdditionalDisk** | Pointer to **bool** |  | [optional] 
**IsBootDisk** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewDisk

`func NewDisk(capacityInMegabytes int32, ) *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDisk

`func (o *Disk) GetAdditionalDisk() bool`

GetAdditionalDisk returns the AdditionalDisk field if non-nil, zero value otherwise.

### GetAdditionalDiskOk

`func (o *Disk) GetAdditionalDiskOk() (*bool, bool)`

GetAdditionalDiskOk returns a tuple with the AdditionalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDisk

`func (o *Disk) SetAdditionalDisk(v bool)`

SetAdditionalDisk sets AdditionalDisk field to given value.

### HasAdditionalDisk

`func (o *Disk) HasAdditionalDisk() bool`

HasAdditionalDisk returns a boolean if a field has been set.

### GetBootDisk

`func (o *Disk) GetBootDisk() bool`

GetBootDisk returns the BootDisk field if non-nil, zero value otherwise.

### GetBootDiskOk

`func (o *Disk) GetBootDiskOk() (*bool, bool)`

GetBootDiskOk returns a tuple with the BootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDisk

`func (o *Disk) SetBootDisk(v bool)`

SetBootDisk sets BootDisk field to given value.

### HasBootDisk

`func (o *Disk) HasBootDisk() bool`

HasBootDisk returns a boolean if a field has been set.

### GetCapacityInMegabytes

`func (o *Disk) GetCapacityInMegabytes() int32`

GetCapacityInMegabytes returns the CapacityInMegabytes field if non-nil, zero value otherwise.

### GetCapacityInMegabytesOk

`func (o *Disk) GetCapacityInMegabytesOk() (*int32, bool)`

GetCapacityInMegabytesOk returns a tuple with the CapacityInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInMegabytes

`func (o *Disk) SetCapacityInMegabytes(v int32)`

SetCapacityInMegabytes sets CapacityInMegabytes field to given value.


### GetCreateOrder

`func (o *Disk) GetCreateOrder() int32`

GetCreateOrder returns the CreateOrder field if non-nil, zero value otherwise.

### GetCreateOrderOk

`func (o *Disk) GetCreateOrderOk() (*int32, bool)`

GetCreateOrderOk returns a tuple with the CreateOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOrder

`func (o *Disk) SetCreateOrder(v int32)`

SetCreateOrder sets CreateOrder field to given value.

### HasCreateOrder

`func (o *Disk) HasCreateOrder() bool`

HasCreateOrder returns a boolean if a field has been set.

### GetIsAdditionalDisk

`func (o *Disk) GetIsAdditionalDisk() bool`

GetIsAdditionalDisk returns the IsAdditionalDisk field if non-nil, zero value otherwise.

### GetIsAdditionalDiskOk

`func (o *Disk) GetIsAdditionalDiskOk() (*bool, bool)`

GetIsAdditionalDiskOk returns a tuple with the IsAdditionalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdditionalDisk

`func (o *Disk) SetIsAdditionalDisk(v bool)`

SetIsAdditionalDisk sets IsAdditionalDisk field to given value.

### HasIsAdditionalDisk

`func (o *Disk) HasIsAdditionalDisk() bool`

HasIsAdditionalDisk returns a boolean if a field has been set.

### GetIsBootDisk

`func (o *Disk) GetIsBootDisk() bool`

GetIsBootDisk returns the IsBootDisk field if non-nil, zero value otherwise.

### GetIsBootDiskOk

`func (o *Disk) GetIsBootDiskOk() (*bool, bool)`

GetIsBootDiskOk returns a tuple with the IsBootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootDisk

`func (o *Disk) SetIsBootDisk(v bool)`

SetIsBootDisk sets IsBootDisk field to given value.

### HasIsBootDisk

`func (o *Disk) HasIsBootDisk() bool`

HasIsBootDisk returns a boolean if a field has been set.

### GetId

`func (o *Disk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Disk) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


