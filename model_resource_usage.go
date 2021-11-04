/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gocons3rt

import (
	"encoding/json"
)

// ResourceUsage struct for ResourceUsage
type ResourceUsage struct {
	NumCpus *int32 `json:"numCpus,omitempty"`
	NumGpus *int32 `json:"numGpus,omitempty"`
	RamInMegabytes *int32 `json:"ramInMegabytes,omitempty"`
	StorageInMegabytes *int32 `json:"storageInMegabytes,omitempty"`
	VirtualMachines *int32 `json:"virtualMachines,omitempty"`
}

// NewResourceUsage instantiates a new ResourceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceUsage() *ResourceUsage {
	this := ResourceUsage{}
	return &this
}

// NewResourceUsageWithDefaults instantiates a new ResourceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceUsageWithDefaults() *ResourceUsage {
	this := ResourceUsage{}
	return &this
}

// GetNumCpus returns the NumCpus field value if set, zero value otherwise.
func (o *ResourceUsage) GetNumCpus() int32 {
	if o == nil || o.NumCpus == nil {
		var ret int32
		return ret
	}
	return *o.NumCpus
}

// GetNumCpusOk returns a tuple with the NumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetNumCpusOk() (*int32, bool) {
	if o == nil || o.NumCpus == nil {
		return nil, false
	}
	return o.NumCpus, true
}

// HasNumCpus returns a boolean if a field has been set.
func (o *ResourceUsage) HasNumCpus() bool {
	if o != nil && o.NumCpus != nil {
		return true
	}

	return false
}

// SetNumCpus gets a reference to the given int32 and assigns it to the NumCpus field.
func (o *ResourceUsage) SetNumCpus(v int32) {
	o.NumCpus = &v
}

// GetNumGpus returns the NumGpus field value if set, zero value otherwise.
func (o *ResourceUsage) GetNumGpus() int32 {
	if o == nil || o.NumGpus == nil {
		var ret int32
		return ret
	}
	return *o.NumGpus
}

// GetNumGpusOk returns a tuple with the NumGpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetNumGpusOk() (*int32, bool) {
	if o == nil || o.NumGpus == nil {
		return nil, false
	}
	return o.NumGpus, true
}

// HasNumGpus returns a boolean if a field has been set.
func (o *ResourceUsage) HasNumGpus() bool {
	if o != nil && o.NumGpus != nil {
		return true
	}

	return false
}

// SetNumGpus gets a reference to the given int32 and assigns it to the NumGpus field.
func (o *ResourceUsage) SetNumGpus(v int32) {
	o.NumGpus = &v
}

// GetRamInMegabytes returns the RamInMegabytes field value if set, zero value otherwise.
func (o *ResourceUsage) GetRamInMegabytes() int32 {
	if o == nil || o.RamInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.RamInMegabytes
}

// GetRamInMegabytesOk returns a tuple with the RamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetRamInMegabytesOk() (*int32, bool) {
	if o == nil || o.RamInMegabytes == nil {
		return nil, false
	}
	return o.RamInMegabytes, true
}

// HasRamInMegabytes returns a boolean if a field has been set.
func (o *ResourceUsage) HasRamInMegabytes() bool {
	if o != nil && o.RamInMegabytes != nil {
		return true
	}

	return false
}

// SetRamInMegabytes gets a reference to the given int32 and assigns it to the RamInMegabytes field.
func (o *ResourceUsage) SetRamInMegabytes(v int32) {
	o.RamInMegabytes = &v
}

// GetStorageInMegabytes returns the StorageInMegabytes field value if set, zero value otherwise.
func (o *ResourceUsage) GetStorageInMegabytes() int32 {
	if o == nil || o.StorageInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.StorageInMegabytes
}

// GetStorageInMegabytesOk returns a tuple with the StorageInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetStorageInMegabytesOk() (*int32, bool) {
	if o == nil || o.StorageInMegabytes == nil {
		return nil, false
	}
	return o.StorageInMegabytes, true
}

// HasStorageInMegabytes returns a boolean if a field has been set.
func (o *ResourceUsage) HasStorageInMegabytes() bool {
	if o != nil && o.StorageInMegabytes != nil {
		return true
	}

	return false
}

// SetStorageInMegabytes gets a reference to the given int32 and assigns it to the StorageInMegabytes field.
func (o *ResourceUsage) SetStorageInMegabytes(v int32) {
	o.StorageInMegabytes = &v
}

// GetVirtualMachines returns the VirtualMachines field value if set, zero value otherwise.
func (o *ResourceUsage) GetVirtualMachines() int32 {
	if o == nil || o.VirtualMachines == nil {
		var ret int32
		return ret
	}
	return *o.VirtualMachines
}

// GetVirtualMachinesOk returns a tuple with the VirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetVirtualMachinesOk() (*int32, bool) {
	if o == nil || o.VirtualMachines == nil {
		return nil, false
	}
	return o.VirtualMachines, true
}

// HasVirtualMachines returns a boolean if a field has been set.
func (o *ResourceUsage) HasVirtualMachines() bool {
	if o != nil && o.VirtualMachines != nil {
		return true
	}

	return false
}

// SetVirtualMachines gets a reference to the given int32 and assigns it to the VirtualMachines field.
func (o *ResourceUsage) SetVirtualMachines(v int32) {
	o.VirtualMachines = &v
}

func (o ResourceUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumCpus != nil {
		toSerialize["numCpus"] = o.NumCpus
	}
	if o.NumGpus != nil {
		toSerialize["numGpus"] = o.NumGpus
	}
	if o.RamInMegabytes != nil {
		toSerialize["ramInMegabytes"] = o.RamInMegabytes
	}
	if o.StorageInMegabytes != nil {
		toSerialize["storageInMegabytes"] = o.StorageInMegabytes
	}
	if o.VirtualMachines != nil {
		toSerialize["virtualMachines"] = o.VirtualMachines
	}
	return json.Marshal(toSerialize)
}

type NullableResourceUsage struct {
	value *ResourceUsage
	isSet bool
}

func (v NullableResourceUsage) Get() *ResourceUsage {
	return v.value
}

func (v *NullableResourceUsage) Set(val *ResourceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceUsage(val *ResourceUsage) *NullableResourceUsage {
	return &NullableResourceUsage{value: val, isSet: true}
}

func (v NullableResourceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


