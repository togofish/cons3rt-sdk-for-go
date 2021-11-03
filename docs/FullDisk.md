# FullDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateOrder** | Pointer to **int32** |  | [optional] 
**CapacityInMegabytes** | Pointer to **int32** |  | [optional] 
**FileSystemType** | Pointer to **string** |  | [optional] 
**MountPoint** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsAdditionalDisk** | Pointer to **bool** |  | [optional] 
**IsBootDisk** | Pointer to **bool** |  | [optional] 

## Methods

### NewFullDisk

`func NewFullDisk() *FullDisk`

NewFullDisk instantiates a new FullDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDiskWithDefaults

`func NewFullDiskWithDefaults() *FullDisk`

NewFullDiskWithDefaults instantiates a new FullDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateOrder

`func (o *FullDisk) GetCreateOrder() int32`

GetCreateOrder returns the CreateOrder field if non-nil, zero value otherwise.

### GetCreateOrderOk

`func (o *FullDisk) GetCreateOrderOk() (*int32, bool)`

GetCreateOrderOk returns a tuple with the CreateOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOrder

`func (o *FullDisk) SetCreateOrder(v int32)`

SetCreateOrder sets CreateOrder field to given value.

### HasCreateOrder

`func (o *FullDisk) HasCreateOrder() bool`

HasCreateOrder returns a boolean if a field has been set.

### GetCapacityInMegabytes

`func (o *FullDisk) GetCapacityInMegabytes() int32`

GetCapacityInMegabytes returns the CapacityInMegabytes field if non-nil, zero value otherwise.

### GetCapacityInMegabytesOk

`func (o *FullDisk) GetCapacityInMegabytesOk() (*int32, bool)`

GetCapacityInMegabytesOk returns a tuple with the CapacityInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInMegabytes

`func (o *FullDisk) SetCapacityInMegabytes(v int32)`

SetCapacityInMegabytes sets CapacityInMegabytes field to given value.

### HasCapacityInMegabytes

`func (o *FullDisk) HasCapacityInMegabytes() bool`

HasCapacityInMegabytes returns a boolean if a field has been set.

### GetFileSystemType

`func (o *FullDisk) GetFileSystemType() string`

GetFileSystemType returns the FileSystemType field if non-nil, zero value otherwise.

### GetFileSystemTypeOk

`func (o *FullDisk) GetFileSystemTypeOk() (*string, bool)`

GetFileSystemTypeOk returns a tuple with the FileSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemType

`func (o *FullDisk) SetFileSystemType(v string)`

SetFileSystemType sets FileSystemType field to given value.

### HasFileSystemType

`func (o *FullDisk) HasFileSystemType() bool`

HasFileSystemType returns a boolean if a field has been set.

### GetMountPoint

`func (o *FullDisk) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *FullDisk) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *FullDisk) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *FullDisk) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetName

`func (o *FullDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *FullDisk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullDisk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullDisk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullDisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAdditionalDisk

`func (o *FullDisk) GetIsAdditionalDisk() bool`

GetIsAdditionalDisk returns the IsAdditionalDisk field if non-nil, zero value otherwise.

### GetIsAdditionalDiskOk

`func (o *FullDisk) GetIsAdditionalDiskOk() (*bool, bool)`

GetIsAdditionalDiskOk returns a tuple with the IsAdditionalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdditionalDisk

`func (o *FullDisk) SetIsAdditionalDisk(v bool)`

SetIsAdditionalDisk sets IsAdditionalDisk field to given value.

### HasIsAdditionalDisk

`func (o *FullDisk) HasIsAdditionalDisk() bool`

HasIsAdditionalDisk returns a boolean if a field has been set.

### GetIsBootDisk

`func (o *FullDisk) GetIsBootDisk() bool`

GetIsBootDisk returns the IsBootDisk field if non-nil, zero value otherwise.

### GetIsBootDiskOk

`func (o *FullDisk) GetIsBootDiskOk() (*bool, bool)`

GetIsBootDiskOk returns a tuple with the IsBootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootDisk

`func (o *FullDisk) SetIsBootDisk(v bool)`

SetIsBootDisk sets IsBootDisk field to given value.

### HasIsBootDisk

`func (o *FullDisk) HasIsBootDisk() bool`

HasIsBootDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


