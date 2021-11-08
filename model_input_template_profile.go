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

// InputTemplateProfile struct for InputTemplateProfile
type InputTemplateProfile struct {
	VirtRealmTemplateName        *string      `json:"virtRealmTemplateName,omitempty"`
	OperatingSystem              *string      `json:"operatingSystem,omitempty"`
	MinNumCpus                   int32        `json:"minNumCpus"`
	MinRam                       int32        `json:"minRam"`
	VgpuRequired                 *bool        `json:"vgpuRequired,omitempty"`
	RequiresNestedVirtualization *bool        `json:"requiresNestedVirtualization,omitempty"`
	AdditionalDisks              *[]InputDisk `json:"additionalDisks,omitempty"`
	MinBootDiskCapacity          *int32       `json:"minBootDiskCapacity,omitempty"`
	RemoteAccessRequired         *bool        `json:"remoteAccessRequired,omitempty"`
	VirtRealmId                  *int32       `json:"virtRealmId,omitempty"`
}

// NewInputTemplateProfile instantiates a new InputTemplateProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputTemplateProfile(minNumCpus int32, minRam int32) *InputTemplateProfile {
	this := InputTemplateProfile{}
	this.MinNumCpus = minNumCpus
	this.MinRam = minRam
	return &this
}

// NewInputTemplateProfileWithDefaults instantiates a new InputTemplateProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputTemplateProfileWithDefaults() *InputTemplateProfile {
	this := InputTemplateProfile{}
	return &this
}

// GetVirtRealmTemplateName returns the VirtRealmTemplateName field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetVirtRealmTemplateName() string {
	if o == nil || o.VirtRealmTemplateName == nil {
		var ret string
		return ret
	}
	return *o.VirtRealmTemplateName
}

// GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetVirtRealmTemplateNameOk() (*string, bool) {
	if o == nil || o.VirtRealmTemplateName == nil {
		return nil, false
	}
	return o.VirtRealmTemplateName, true
}

// HasVirtRealmTemplateName returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasVirtRealmTemplateName() bool {
	if o != nil && o.VirtRealmTemplateName != nil {
		return true
	}

	return false
}

// SetVirtRealmTemplateName gets a reference to the given string and assigns it to the VirtRealmTemplateName field.
func (o *InputTemplateProfile) SetVirtRealmTemplateName(v string) {
	o.VirtRealmTemplateName = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *InputTemplateProfile) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetMinNumCpus returns the MinNumCpus field value
func (o *InputTemplateProfile) GetMinNumCpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinNumCpus
}

// GetMinNumCpusOk returns a tuple with the MinNumCpus field value
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetMinNumCpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinNumCpus, true
}

// SetMinNumCpus sets field value
func (o *InputTemplateProfile) SetMinNumCpus(v int32) {
	o.MinNumCpus = v
}

// GetMinRam returns the MinRam field value
func (o *InputTemplateProfile) GetMinRam() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinRam
}

// GetMinRamOk returns a tuple with the MinRam field value
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetMinRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinRam, true
}

// SetMinRam sets field value
func (o *InputTemplateProfile) SetMinRam(v int32) {
	o.MinRam = v
}

// GetVgpuRequired returns the VgpuRequired field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetVgpuRequired() bool {
	if o == nil || o.VgpuRequired == nil {
		var ret bool
		return ret
	}
	return *o.VgpuRequired
}

// GetVgpuRequiredOk returns a tuple with the VgpuRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetVgpuRequiredOk() (*bool, bool) {
	if o == nil || o.VgpuRequired == nil {
		return nil, false
	}
	return o.VgpuRequired, true
}

// HasVgpuRequired returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasVgpuRequired() bool {
	if o != nil && o.VgpuRequired != nil {
		return true
	}

	return false
}

// SetVgpuRequired gets a reference to the given bool and assigns it to the VgpuRequired field.
func (o *InputTemplateProfile) SetVgpuRequired(v bool) {
	o.VgpuRequired = &v
}

// GetRequiresNestedVirtualization returns the RequiresNestedVirtualization field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetRequiresNestedVirtualization() bool {
	if o == nil || o.RequiresNestedVirtualization == nil {
		var ret bool
		return ret
	}
	return *o.RequiresNestedVirtualization
}

// GetRequiresNestedVirtualizationOk returns a tuple with the RequiresNestedVirtualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetRequiresNestedVirtualizationOk() (*bool, bool) {
	if o == nil || o.RequiresNestedVirtualization == nil {
		return nil, false
	}
	return o.RequiresNestedVirtualization, true
}

// HasRequiresNestedVirtualization returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasRequiresNestedVirtualization() bool {
	if o != nil && o.RequiresNestedVirtualization != nil {
		return true
	}

	return false
}

// SetRequiresNestedVirtualization gets a reference to the given bool and assigns it to the RequiresNestedVirtualization field.
func (o *InputTemplateProfile) SetRequiresNestedVirtualization(v bool) {
	o.RequiresNestedVirtualization = &v
}

// GetAdditionalDisks returns the AdditionalDisks field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetAdditionalDisks() []InputDisk {
	if o == nil || o.AdditionalDisks == nil {
		var ret []InputDisk
		return ret
	}
	return *o.AdditionalDisks
}

// GetAdditionalDisksOk returns a tuple with the AdditionalDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetAdditionalDisksOk() (*[]InputDisk, bool) {
	if o == nil || o.AdditionalDisks == nil {
		return nil, false
	}
	return o.AdditionalDisks, true
}

// HasAdditionalDisks returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasAdditionalDisks() bool {
	if o != nil && o.AdditionalDisks != nil {
		return true
	}

	return false
}

// SetAdditionalDisks gets a reference to the given []InputDisk and assigns it to the AdditionalDisks field.
func (o *InputTemplateProfile) SetAdditionalDisks(v []InputDisk) {
	o.AdditionalDisks = &v
}

// GetMinBootDiskCapacity returns the MinBootDiskCapacity field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetMinBootDiskCapacity() int32 {
	if o == nil || o.MinBootDiskCapacity == nil {
		var ret int32
		return ret
	}
	return *o.MinBootDiskCapacity
}

// GetMinBootDiskCapacityOk returns a tuple with the MinBootDiskCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetMinBootDiskCapacityOk() (*int32, bool) {
	if o == nil || o.MinBootDiskCapacity == nil {
		return nil, false
	}
	return o.MinBootDiskCapacity, true
}

// HasMinBootDiskCapacity returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasMinBootDiskCapacity() bool {
	if o != nil && o.MinBootDiskCapacity != nil {
		return true
	}

	return false
}

// SetMinBootDiskCapacity gets a reference to the given int32 and assigns it to the MinBootDiskCapacity field.
func (o *InputTemplateProfile) SetMinBootDiskCapacity(v int32) {
	o.MinBootDiskCapacity = &v
}

// GetRemoteAccessRequired returns the RemoteAccessRequired field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetRemoteAccessRequired() bool {
	if o == nil || o.RemoteAccessRequired == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessRequired
}

// GetRemoteAccessRequiredOk returns a tuple with the RemoteAccessRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetRemoteAccessRequiredOk() (*bool, bool) {
	if o == nil || o.RemoteAccessRequired == nil {
		return nil, false
	}
	return o.RemoteAccessRequired, true
}

// HasRemoteAccessRequired returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasRemoteAccessRequired() bool {
	if o != nil && o.RemoteAccessRequired != nil {
		return true
	}

	return false
}

// SetRemoteAccessRequired gets a reference to the given bool and assigns it to the RemoteAccessRequired field.
func (o *InputTemplateProfile) SetRemoteAccessRequired(v bool) {
	o.RemoteAccessRequired = &v
}

// GetVirtRealmId returns the VirtRealmId field value if set, zero value otherwise.
func (o *InputTemplateProfile) GetVirtRealmId() int32 {
	if o == nil || o.VirtRealmId == nil {
		var ret int32
		return ret
	}
	return *o.VirtRealmId
}

// GetVirtRealmIdOk returns a tuple with the VirtRealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateProfile) GetVirtRealmIdOk() (*int32, bool) {
	if o == nil || o.VirtRealmId == nil {
		return nil, false
	}
	return o.VirtRealmId, true
}

// HasVirtRealmId returns a boolean if a field has been set.
func (o *InputTemplateProfile) HasVirtRealmId() bool {
	if o != nil && o.VirtRealmId != nil {
		return true
	}

	return false
}

// SetVirtRealmId gets a reference to the given int32 and assigns it to the VirtRealmId field.
func (o *InputTemplateProfile) SetVirtRealmId(v int32) {
	o.VirtRealmId = &v
}

func (o InputTemplateProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtRealmTemplateName != nil {
		toSerialize["virtRealmTemplateName"] = o.VirtRealmTemplateName
	}
	if o.OperatingSystem != nil {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if true {
		toSerialize["minNumCpus"] = o.MinNumCpus
	}
	if true {
		toSerialize["minRam"] = o.MinRam
	}
	if o.VgpuRequired != nil {
		toSerialize["vgpuRequired"] = o.VgpuRequired
	}
	if o.RequiresNestedVirtualization != nil {
		toSerialize["requiresNestedVirtualization"] = o.RequiresNestedVirtualization
	}
	if o.AdditionalDisks != nil {
		toSerialize["additionalDisks"] = o.AdditionalDisks
	}
	if o.MinBootDiskCapacity != nil {
		toSerialize["minBootDiskCapacity"] = o.MinBootDiskCapacity
	}
	if o.RemoteAccessRequired != nil {
		toSerialize["remoteAccessRequired"] = o.RemoteAccessRequired
	}
	if o.VirtRealmId != nil {
		toSerialize["virtRealmId"] = o.VirtRealmId
	}
	return json.Marshal(toSerialize)
}

type NullableInputTemplateProfile struct {
	value *InputTemplateProfile
	isSet bool
}

func (v NullableInputTemplateProfile) Get() *InputTemplateProfile {
	return v.value
}

func (v *NullableInputTemplateProfile) Set(val *InputTemplateProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTemplateProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTemplateProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTemplateProfile(val *InputTemplateProfile) *NullableInputTemplateProfile {
	return &NullableInputTemplateProfile{value: val, isSet: true}
}

func (v NullableInputTemplateProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputTemplateProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
