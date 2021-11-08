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

// MinimalTemplateProfile struct for MinimalTemplateProfile
type MinimalTemplateProfile struct {
	VirtRealmTemplateName        *string        `json:"virtRealmTemplateName,omitempty"`
	OperatingSystem              *string        `json:"operatingSystem,omitempty"`
	MinBootDiskCapacity          *int32         `json:"minBootDiskCapacity,omitempty"`
	MinNumCpus                   *int32         `json:"minNumCpus,omitempty"`
	MinRam                       *int32         `json:"minRam,omitempty"`
	AdditionalDisks              *[]MinimalDisk `json:"additionalDisks,omitempty"`
	Cons3rtAgentRequired         *bool          `json:"cons3rtAgentRequired,omitempty"`
	VgpuRequired                 *bool          `json:"vgpuRequired,omitempty"`
	GpuType                      *string        `json:"gpuType,omitempty"`
	RequiresNestedVirtualization *bool          `json:"requiresNestedVirtualization,omitempty"`
	RemoteAccessRequired         *bool          `json:"remoteAccessRequired,omitempty"`
}

// NewMinimalTemplateProfile instantiates a new MinimalTemplateProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalTemplateProfile() *MinimalTemplateProfile {
	this := MinimalTemplateProfile{}
	return &this
}

// NewMinimalTemplateProfileWithDefaults instantiates a new MinimalTemplateProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalTemplateProfileWithDefaults() *MinimalTemplateProfile {
	this := MinimalTemplateProfile{}
	return &this
}

// GetVirtRealmTemplateName returns the VirtRealmTemplateName field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetVirtRealmTemplateName() string {
	if o == nil || o.VirtRealmTemplateName == nil {
		var ret string
		return ret
	}
	return *o.VirtRealmTemplateName
}

// GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetVirtRealmTemplateNameOk() (*string, bool) {
	if o == nil || o.VirtRealmTemplateName == nil {
		return nil, false
	}
	return o.VirtRealmTemplateName, true
}

// HasVirtRealmTemplateName returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasVirtRealmTemplateName() bool {
	if o != nil && o.VirtRealmTemplateName != nil {
		return true
	}

	return false
}

// SetVirtRealmTemplateName gets a reference to the given string and assigns it to the VirtRealmTemplateName field.
func (o *MinimalTemplateProfile) SetVirtRealmTemplateName(v string) {
	o.VirtRealmTemplateName = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *MinimalTemplateProfile) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetMinBootDiskCapacity returns the MinBootDiskCapacity field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetMinBootDiskCapacity() int32 {
	if o == nil || o.MinBootDiskCapacity == nil {
		var ret int32
		return ret
	}
	return *o.MinBootDiskCapacity
}

// GetMinBootDiskCapacityOk returns a tuple with the MinBootDiskCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetMinBootDiskCapacityOk() (*int32, bool) {
	if o == nil || o.MinBootDiskCapacity == nil {
		return nil, false
	}
	return o.MinBootDiskCapacity, true
}

// HasMinBootDiskCapacity returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasMinBootDiskCapacity() bool {
	if o != nil && o.MinBootDiskCapacity != nil {
		return true
	}

	return false
}

// SetMinBootDiskCapacity gets a reference to the given int32 and assigns it to the MinBootDiskCapacity field.
func (o *MinimalTemplateProfile) SetMinBootDiskCapacity(v int32) {
	o.MinBootDiskCapacity = &v
}

// GetMinNumCpus returns the MinNumCpus field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetMinNumCpus() int32 {
	if o == nil || o.MinNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MinNumCpus
}

// GetMinNumCpusOk returns a tuple with the MinNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetMinNumCpusOk() (*int32, bool) {
	if o == nil || o.MinNumCpus == nil {
		return nil, false
	}
	return o.MinNumCpus, true
}

// HasMinNumCpus returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasMinNumCpus() bool {
	if o != nil && o.MinNumCpus != nil {
		return true
	}

	return false
}

// SetMinNumCpus gets a reference to the given int32 and assigns it to the MinNumCpus field.
func (o *MinimalTemplateProfile) SetMinNumCpus(v int32) {
	o.MinNumCpus = &v
}

// GetMinRam returns the MinRam field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetMinRam() int32 {
	if o == nil || o.MinRam == nil {
		var ret int32
		return ret
	}
	return *o.MinRam
}

// GetMinRamOk returns a tuple with the MinRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetMinRamOk() (*int32, bool) {
	if o == nil || o.MinRam == nil {
		return nil, false
	}
	return o.MinRam, true
}

// HasMinRam returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasMinRam() bool {
	if o != nil && o.MinRam != nil {
		return true
	}

	return false
}

// SetMinRam gets a reference to the given int32 and assigns it to the MinRam field.
func (o *MinimalTemplateProfile) SetMinRam(v int32) {
	o.MinRam = &v
}

// GetAdditionalDisks returns the AdditionalDisks field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetAdditionalDisks() []MinimalDisk {
	if o == nil || o.AdditionalDisks == nil {
		var ret []MinimalDisk
		return ret
	}
	return *o.AdditionalDisks
}

// GetAdditionalDisksOk returns a tuple with the AdditionalDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetAdditionalDisksOk() (*[]MinimalDisk, bool) {
	if o == nil || o.AdditionalDisks == nil {
		return nil, false
	}
	return o.AdditionalDisks, true
}

// HasAdditionalDisks returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasAdditionalDisks() bool {
	if o != nil && o.AdditionalDisks != nil {
		return true
	}

	return false
}

// SetAdditionalDisks gets a reference to the given []MinimalDisk and assigns it to the AdditionalDisks field.
func (o *MinimalTemplateProfile) SetAdditionalDisks(v []MinimalDisk) {
	o.AdditionalDisks = &v
}

// GetCons3rtAgentRequired returns the Cons3rtAgentRequired field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetCons3rtAgentRequired() bool {
	if o == nil || o.Cons3rtAgentRequired == nil {
		var ret bool
		return ret
	}
	return *o.Cons3rtAgentRequired
}

// GetCons3rtAgentRequiredOk returns a tuple with the Cons3rtAgentRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetCons3rtAgentRequiredOk() (*bool, bool) {
	if o == nil || o.Cons3rtAgentRequired == nil {
		return nil, false
	}
	return o.Cons3rtAgentRequired, true
}

// HasCons3rtAgentRequired returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasCons3rtAgentRequired() bool {
	if o != nil && o.Cons3rtAgentRequired != nil {
		return true
	}

	return false
}

// SetCons3rtAgentRequired gets a reference to the given bool and assigns it to the Cons3rtAgentRequired field.
func (o *MinimalTemplateProfile) SetCons3rtAgentRequired(v bool) {
	o.Cons3rtAgentRequired = &v
}

// GetVgpuRequired returns the VgpuRequired field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetVgpuRequired() bool {
	if o == nil || o.VgpuRequired == nil {
		var ret bool
		return ret
	}
	return *o.VgpuRequired
}

// GetVgpuRequiredOk returns a tuple with the VgpuRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetVgpuRequiredOk() (*bool, bool) {
	if o == nil || o.VgpuRequired == nil {
		return nil, false
	}
	return o.VgpuRequired, true
}

// HasVgpuRequired returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasVgpuRequired() bool {
	if o != nil && o.VgpuRequired != nil {
		return true
	}

	return false
}

// SetVgpuRequired gets a reference to the given bool and assigns it to the VgpuRequired field.
func (o *MinimalTemplateProfile) SetVgpuRequired(v bool) {
	o.VgpuRequired = &v
}

// GetGpuType returns the GpuType field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetGpuType() string {
	if o == nil || o.GpuType == nil {
		var ret string
		return ret
	}
	return *o.GpuType
}

// GetGpuTypeOk returns a tuple with the GpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetGpuTypeOk() (*string, bool) {
	if o == nil || o.GpuType == nil {
		return nil, false
	}
	return o.GpuType, true
}

// HasGpuType returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasGpuType() bool {
	if o != nil && o.GpuType != nil {
		return true
	}

	return false
}

// SetGpuType gets a reference to the given string and assigns it to the GpuType field.
func (o *MinimalTemplateProfile) SetGpuType(v string) {
	o.GpuType = &v
}

// GetRequiresNestedVirtualization returns the RequiresNestedVirtualization field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetRequiresNestedVirtualization() bool {
	if o == nil || o.RequiresNestedVirtualization == nil {
		var ret bool
		return ret
	}
	return *o.RequiresNestedVirtualization
}

// GetRequiresNestedVirtualizationOk returns a tuple with the RequiresNestedVirtualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetRequiresNestedVirtualizationOk() (*bool, bool) {
	if o == nil || o.RequiresNestedVirtualization == nil {
		return nil, false
	}
	return o.RequiresNestedVirtualization, true
}

// HasRequiresNestedVirtualization returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasRequiresNestedVirtualization() bool {
	if o != nil && o.RequiresNestedVirtualization != nil {
		return true
	}

	return false
}

// SetRequiresNestedVirtualization gets a reference to the given bool and assigns it to the RequiresNestedVirtualization field.
func (o *MinimalTemplateProfile) SetRequiresNestedVirtualization(v bool) {
	o.RequiresNestedVirtualization = &v
}

// GetRemoteAccessRequired returns the RemoteAccessRequired field value if set, zero value otherwise.
func (o *MinimalTemplateProfile) GetRemoteAccessRequired() bool {
	if o == nil || o.RemoteAccessRequired == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessRequired
}

// GetRemoteAccessRequiredOk returns a tuple with the RemoteAccessRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalTemplateProfile) GetRemoteAccessRequiredOk() (*bool, bool) {
	if o == nil || o.RemoteAccessRequired == nil {
		return nil, false
	}
	return o.RemoteAccessRequired, true
}

// HasRemoteAccessRequired returns a boolean if a field has been set.
func (o *MinimalTemplateProfile) HasRemoteAccessRequired() bool {
	if o != nil && o.RemoteAccessRequired != nil {
		return true
	}

	return false
}

// SetRemoteAccessRequired gets a reference to the given bool and assigns it to the RemoteAccessRequired field.
func (o *MinimalTemplateProfile) SetRemoteAccessRequired(v bool) {
	o.RemoteAccessRequired = &v
}

func (o MinimalTemplateProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtRealmTemplateName != nil {
		toSerialize["virtRealmTemplateName"] = o.VirtRealmTemplateName
	}
	if o.OperatingSystem != nil {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if o.MinBootDiskCapacity != nil {
		toSerialize["minBootDiskCapacity"] = o.MinBootDiskCapacity
	}
	if o.MinNumCpus != nil {
		toSerialize["minNumCpus"] = o.MinNumCpus
	}
	if o.MinRam != nil {
		toSerialize["minRam"] = o.MinRam
	}
	if o.AdditionalDisks != nil {
		toSerialize["additionalDisks"] = o.AdditionalDisks
	}
	if o.Cons3rtAgentRequired != nil {
		toSerialize["cons3rtAgentRequired"] = o.Cons3rtAgentRequired
	}
	if o.VgpuRequired != nil {
		toSerialize["vgpuRequired"] = o.VgpuRequired
	}
	if o.GpuType != nil {
		toSerialize["gpuType"] = o.GpuType
	}
	if o.RequiresNestedVirtualization != nil {
		toSerialize["requiresNestedVirtualization"] = o.RequiresNestedVirtualization
	}
	if o.RemoteAccessRequired != nil {
		toSerialize["remoteAccessRequired"] = o.RemoteAccessRequired
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalTemplateProfile struct {
	value *MinimalTemplateProfile
	isSet bool
}

func (v NullableMinimalTemplateProfile) Get() *MinimalTemplateProfile {
	return v.value
}

func (v *NullableMinimalTemplateProfile) Set(val *MinimalTemplateProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalTemplateProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalTemplateProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalTemplateProfile(val *MinimalTemplateProfile) *NullableMinimalTemplateProfile {
	return &NullableMinimalTemplateProfile{value: val, isSet: true}
}

func (v NullableMinimalTemplateProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalTemplateProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
