/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cons3rt

import (
	"encoding/json"
)

// ProjectLimits struct for ProjectLimits
type ProjectLimits struct {
	GpuTypeMaximums       *map[string]int32 `json:"gpuTypeMaximums,omitempty"`
	MaxNumCpus            int32             `json:"maxNumCpus"`
	MaxNumGpus            int32             `json:"maxNumGpus"`
	MaxRamInMegabytes     int32             `json:"maxRamInMegabytes"`
	MaxStorageInMegabytes int32             `json:"maxStorageInMegabytes"`
	MaxVirtualMachines    int32             `json:"maxVirtualMachines"`
	ValidUntil            *int32            `json:"validUntil,omitempty"`
}

// NewProjectLimits instantiates a new ProjectLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectLimits(maxNumCpus int32, maxNumGpus int32, maxRamInMegabytes int32, maxStorageInMegabytes int32, maxVirtualMachines int32) *ProjectLimits {
	this := ProjectLimits{}
	this.MaxNumCpus = maxNumCpus
	this.MaxNumGpus = maxNumGpus
	this.MaxRamInMegabytes = maxRamInMegabytes
	this.MaxStorageInMegabytes = maxStorageInMegabytes
	this.MaxVirtualMachines = maxVirtualMachines
	return &this
}

// NewProjectLimitsWithDefaults instantiates a new ProjectLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectLimitsWithDefaults() *ProjectLimits {
	this := ProjectLimits{}
	return &this
}

// GetGpuTypeMaximums returns the GpuTypeMaximums field value if set, zero value otherwise.
func (o *ProjectLimits) GetGpuTypeMaximums() map[string]int32 {
	if o == nil || o.GpuTypeMaximums == nil {
		var ret map[string]int32
		return ret
	}
	return *o.GpuTypeMaximums
}

// GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetGpuTypeMaximumsOk() (*map[string]int32, bool) {
	if o == nil || o.GpuTypeMaximums == nil {
		return nil, false
	}
	return o.GpuTypeMaximums, true
}

// HasGpuTypeMaximums returns a boolean if a field has been set.
func (o *ProjectLimits) HasGpuTypeMaximums() bool {
	if o != nil && o.GpuTypeMaximums != nil {
		return true
	}

	return false
}

// SetGpuTypeMaximums gets a reference to the given map[string]int32 and assigns it to the GpuTypeMaximums field.
func (o *ProjectLimits) SetGpuTypeMaximums(v map[string]int32) {
	o.GpuTypeMaximums = &v
}

// GetMaxNumCpus returns the MaxNumCpus field value
func (o *ProjectLimits) GetMaxNumCpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxNumCpus
}

// GetMaxNumCpusOk returns a tuple with the MaxNumCpus field value
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetMaxNumCpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumCpus, true
}

// SetMaxNumCpus sets field value
func (o *ProjectLimits) SetMaxNumCpus(v int32) {
	o.MaxNumCpus = v
}

// GetMaxNumGpus returns the MaxNumGpus field value
func (o *ProjectLimits) GetMaxNumGpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxNumGpus
}

// GetMaxNumGpusOk returns a tuple with the MaxNumGpus field value
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetMaxNumGpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumGpus, true
}

// SetMaxNumGpus sets field value
func (o *ProjectLimits) SetMaxNumGpus(v int32) {
	o.MaxNumGpus = v
}

// GetMaxRamInMegabytes returns the MaxRamInMegabytes field value
func (o *ProjectLimits) GetMaxRamInMegabytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxRamInMegabytes
}

// GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field value
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetMaxRamInMegabytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRamInMegabytes, true
}

// SetMaxRamInMegabytes sets field value
func (o *ProjectLimits) SetMaxRamInMegabytes(v int32) {
	o.MaxRamInMegabytes = v
}

// GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field value
func (o *ProjectLimits) GetMaxStorageInMegabytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxStorageInMegabytes
}

// GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field value
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetMaxStorageInMegabytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStorageInMegabytes, true
}

// SetMaxStorageInMegabytes sets field value
func (o *ProjectLimits) SetMaxStorageInMegabytes(v int32) {
	o.MaxStorageInMegabytes = v
}

// GetMaxVirtualMachines returns the MaxVirtualMachines field value
func (o *ProjectLimits) GetMaxVirtualMachines() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxVirtualMachines
}

// GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field value
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetMaxVirtualMachinesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxVirtualMachines, true
}

// SetMaxVirtualMachines sets field value
func (o *ProjectLimits) SetMaxVirtualMachines(v int32) {
	o.MaxVirtualMachines = v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *ProjectLimits) GetValidUntil() int32 {
	if o == nil || o.ValidUntil == nil {
		var ret int32
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLimits) GetValidUntilOk() (*int32, bool) {
	if o == nil || o.ValidUntil == nil {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *ProjectLimits) HasValidUntil() bool {
	if o != nil && o.ValidUntil != nil {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given int32 and assigns it to the ValidUntil field.
func (o *ProjectLimits) SetValidUntil(v int32) {
	o.ValidUntil = &v
}

func (o ProjectLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GpuTypeMaximums != nil {
		toSerialize["gpuTypeMaximums"] = o.GpuTypeMaximums
	}
	if true {
		toSerialize["maxNumCpus"] = o.MaxNumCpus
	}
	if true {
		toSerialize["maxNumGpus"] = o.MaxNumGpus
	}
	if true {
		toSerialize["maxRamInMegabytes"] = o.MaxRamInMegabytes
	}
	if true {
		toSerialize["maxStorageInMegabytes"] = o.MaxStorageInMegabytes
	}
	if true {
		toSerialize["maxVirtualMachines"] = o.MaxVirtualMachines
	}
	if o.ValidUntil != nil {
		toSerialize["validUntil"] = o.ValidUntil
	}
	return json.Marshal(toSerialize)
}

type NullableProjectLimits struct {
	value *ProjectLimits
	isSet bool
}

func (v NullableProjectLimits) Get() *ProjectLimits {
	return v.value
}

func (v *NullableProjectLimits) Set(val *ProjectLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectLimits(val *ProjectLimits) *NullableProjectLimits {
	return &NullableProjectLimits{value: val, isSet: true}
}

func (v NullableProjectLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
