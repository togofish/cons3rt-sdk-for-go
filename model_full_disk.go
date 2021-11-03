/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FullDisk struct for FullDisk
type FullDisk struct {
	CreateOrder         *int32  `json:"createOrder,omitempty"`
	CapacityInMegabytes *int32  `json:"capacityInMegabytes,omitempty"`
	FileSystemType      *string `json:"fileSystemType,omitempty"`
	MountPoint          *string `json:"mountPoint,omitempty"`
	Name                *string `json:"name,omitempty"`
	Id                  *int32  `json:"id,omitempty"`
	IsAdditionalDisk    *bool   `json:"isAdditionalDisk,omitempty"`
	IsBootDisk          *bool   `json:"isBootDisk,omitempty"`
}

// NewFullDisk instantiates a new FullDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDisk() *FullDisk {
	this := FullDisk{}
	return &this
}

// NewFullDiskWithDefaults instantiates a new FullDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDiskWithDefaults() *FullDisk {
	this := FullDisk{}
	return &this
}

// GetCreateOrder returns the CreateOrder field value if set, zero value otherwise.
func (o *FullDisk) GetCreateOrder() int32 {
	if o == nil || o.CreateOrder == nil {
		var ret int32
		return ret
	}
	return *o.CreateOrder
}

// GetCreateOrderOk returns a tuple with the CreateOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetCreateOrderOk() (*int32, bool) {
	if o == nil || o.CreateOrder == nil {
		return nil, false
	}
	return o.CreateOrder, true
}

// HasCreateOrder returns a boolean if a field has been set.
func (o *FullDisk) HasCreateOrder() bool {
	if o != nil && o.CreateOrder != nil {
		return true
	}

	return false
}

// SetCreateOrder gets a reference to the given int32 and assigns it to the CreateOrder field.
func (o *FullDisk) SetCreateOrder(v int32) {
	o.CreateOrder = &v
}

// GetCapacityInMegabytes returns the CapacityInMegabytes field value if set, zero value otherwise.
func (o *FullDisk) GetCapacityInMegabytes() int32 {
	if o == nil || o.CapacityInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.CapacityInMegabytes
}

// GetCapacityInMegabytesOk returns a tuple with the CapacityInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetCapacityInMegabytesOk() (*int32, bool) {
	if o == nil || o.CapacityInMegabytes == nil {
		return nil, false
	}
	return o.CapacityInMegabytes, true
}

// HasCapacityInMegabytes returns a boolean if a field has been set.
func (o *FullDisk) HasCapacityInMegabytes() bool {
	if o != nil && o.CapacityInMegabytes != nil {
		return true
	}

	return false
}

// SetCapacityInMegabytes gets a reference to the given int32 and assigns it to the CapacityInMegabytes field.
func (o *FullDisk) SetCapacityInMegabytes(v int32) {
	o.CapacityInMegabytes = &v
}

// GetFileSystemType returns the FileSystemType field value if set, zero value otherwise.
func (o *FullDisk) GetFileSystemType() string {
	if o == nil || o.FileSystemType == nil {
		var ret string
		return ret
	}
	return *o.FileSystemType
}

// GetFileSystemTypeOk returns a tuple with the FileSystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetFileSystemTypeOk() (*string, bool) {
	if o == nil || o.FileSystemType == nil {
		return nil, false
	}
	return o.FileSystemType, true
}

// HasFileSystemType returns a boolean if a field has been set.
func (o *FullDisk) HasFileSystemType() bool {
	if o != nil && o.FileSystemType != nil {
		return true
	}

	return false
}

// SetFileSystemType gets a reference to the given string and assigns it to the FileSystemType field.
func (o *FullDisk) SetFileSystemType(v string) {
	o.FileSystemType = &v
}

// GetMountPoint returns the MountPoint field value if set, zero value otherwise.
func (o *FullDisk) GetMountPoint() string {
	if o == nil || o.MountPoint == nil {
		var ret string
		return ret
	}
	return *o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetMountPointOk() (*string, bool) {
	if o == nil || o.MountPoint == nil {
		return nil, false
	}
	return o.MountPoint, true
}

// HasMountPoint returns a boolean if a field has been set.
func (o *FullDisk) HasMountPoint() bool {
	if o != nil && o.MountPoint != nil {
		return true
	}

	return false
}

// SetMountPoint gets a reference to the given string and assigns it to the MountPoint field.
func (o *FullDisk) SetMountPoint(v string) {
	o.MountPoint = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullDisk) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullDisk) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullDisk) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullDisk) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullDisk) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FullDisk) SetId(v int32) {
	o.Id = &v
}

// GetIsAdditionalDisk returns the IsAdditionalDisk field value if set, zero value otherwise.
func (o *FullDisk) GetIsAdditionalDisk() bool {
	if o == nil || o.IsAdditionalDisk == nil {
		var ret bool
		return ret
	}
	return *o.IsAdditionalDisk
}

// GetIsAdditionalDiskOk returns a tuple with the IsAdditionalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetIsAdditionalDiskOk() (*bool, bool) {
	if o == nil || o.IsAdditionalDisk == nil {
		return nil, false
	}
	return o.IsAdditionalDisk, true
}

// HasIsAdditionalDisk returns a boolean if a field has been set.
func (o *FullDisk) HasIsAdditionalDisk() bool {
	if o != nil && o.IsAdditionalDisk != nil {
		return true
	}

	return false
}

// SetIsAdditionalDisk gets a reference to the given bool and assigns it to the IsAdditionalDisk field.
func (o *FullDisk) SetIsAdditionalDisk(v bool) {
	o.IsAdditionalDisk = &v
}

// GetIsBootDisk returns the IsBootDisk field value if set, zero value otherwise.
func (o *FullDisk) GetIsBootDisk() bool {
	if o == nil || o.IsBootDisk == nil {
		var ret bool
		return ret
	}
	return *o.IsBootDisk
}

// GetIsBootDiskOk returns a tuple with the IsBootDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDisk) GetIsBootDiskOk() (*bool, bool) {
	if o == nil || o.IsBootDisk == nil {
		return nil, false
	}
	return o.IsBootDisk, true
}

// HasIsBootDisk returns a boolean if a field has been set.
func (o *FullDisk) HasIsBootDisk() bool {
	if o != nil && o.IsBootDisk != nil {
		return true
	}

	return false
}

// SetIsBootDisk gets a reference to the given bool and assigns it to the IsBootDisk field.
func (o *FullDisk) SetIsBootDisk(v bool) {
	o.IsBootDisk = &v
}

func (o FullDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateOrder != nil {
		toSerialize["createOrder"] = o.CreateOrder
	}
	if o.CapacityInMegabytes != nil {
		toSerialize["capacityInMegabytes"] = o.CapacityInMegabytes
	}
	if o.FileSystemType != nil {
		toSerialize["fileSystemType"] = o.FileSystemType
	}
	if o.MountPoint != nil {
		toSerialize["mountPoint"] = o.MountPoint
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsAdditionalDisk != nil {
		toSerialize["isAdditionalDisk"] = o.IsAdditionalDisk
	}
	if o.IsBootDisk != nil {
		toSerialize["isBootDisk"] = o.IsBootDisk
	}
	return json.Marshal(toSerialize)
}

type NullableFullDisk struct {
	value *FullDisk
	isSet bool
}

func (v NullableFullDisk) Get() *FullDisk {
	return v.value
}

func (v *NullableFullDisk) Set(val *FullDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDisk(val *FullDisk) *NullableFullDisk {
	return &NullableFullDisk{value: val, isSet: true}
}

func (v NullableFullDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
