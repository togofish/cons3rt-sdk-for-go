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

// FullPhysicalMachine struct for FullPhysicalMachine
type FullPhysicalMachine struct {
	HostName              *string                        `json:"hostName,omitempty"`
	IpAddress             *string                        `json:"ipAddress,omitempty"`
	MacAddress            *string                        `json:"macAddress,omitempty"`
	CpuCount              *int32                         `json:"cpuCount,omitempty"`
	Ram                   *int32                         `json:"ram,omitempty"`
	Architecture          *string                        `json:"architecture,omitempty"`
	Bits                  *string                        `json:"bits,omitempty"`
	OperatingSystem       *string                        `json:"operatingSystem,omitempty"`
	RemoteAccessCapable   *bool                          `json:"remoteAccessCapable,omitempty"`
	RemoteAccessTemplates *[]MinimalRemoteAccessTemplate `json:"remoteAccessTemplates,omitempty"`
	BaseDisks             *[]FullDisk                    `json:"baseDisks,omitempty"`
}

// NewFullPhysicalMachine instantiates a new FullPhysicalMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullPhysicalMachine(subtype string) *FullPhysicalMachine {
	this := FullPhysicalMachine{}
	this.Subtype = subtype
	return &this
}

// NewFullPhysicalMachineWithDefaults instantiates a new FullPhysicalMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullPhysicalMachineWithDefaults() *FullPhysicalMachine {
	this := FullPhysicalMachine{}
	return &this
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *FullPhysicalMachine) SetHostName(v string) {
	o.HostName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *FullPhysicalMachine) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *FullPhysicalMachine) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetCpuCount() int32 {
	if o == nil || o.CpuCount == nil {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetCpuCountOk() (*int32, bool) {
	if o == nil || o.CpuCount == nil {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasCpuCount() bool {
	if o != nil && o.CpuCount != nil {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *FullPhysicalMachine) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetRam() int32 {
	if o == nil || o.Ram == nil {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetRamOk() (*int32, bool) {
	if o == nil || o.Ram == nil {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *FullPhysicalMachine) SetRam(v int32) {
	o.Ram = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *FullPhysicalMachine) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetBits() string {
	if o == nil || o.Bits == nil {
		var ret string
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetBitsOk() (*string, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given string and assigns it to the Bits field.
func (o *FullPhysicalMachine) SetBits(v string) {
	o.Bits = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *FullPhysicalMachine) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetRemoteAccessCapable returns the RemoteAccessCapable field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetRemoteAccessCapable() bool {
	if o == nil || o.RemoteAccessCapable == nil {
		var ret bool
		return ret
	}
	return *o.RemoteAccessCapable
}

// GetRemoteAccessCapableOk returns a tuple with the RemoteAccessCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetRemoteAccessCapableOk() (*bool, bool) {
	if o == nil || o.RemoteAccessCapable == nil {
		return nil, false
	}
	return o.RemoteAccessCapable, true
}

// HasRemoteAccessCapable returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasRemoteAccessCapable() bool {
	if o != nil && o.RemoteAccessCapable != nil {
		return true
	}

	return false
}

// SetRemoteAccessCapable gets a reference to the given bool and assigns it to the RemoteAccessCapable field.
func (o *FullPhysicalMachine) SetRemoteAccessCapable(v bool) {
	o.RemoteAccessCapable = &v
}

// GetRemoteAccessTemplates returns the RemoteAccessTemplates field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetRemoteAccessTemplates() []MinimalRemoteAccessTemplate {
	if o == nil || o.RemoteAccessTemplates == nil {
		var ret []MinimalRemoteAccessTemplate
		return ret
	}
	return *o.RemoteAccessTemplates
}

// GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetRemoteAccessTemplatesOk() (*[]MinimalRemoteAccessTemplate, bool) {
	if o == nil || o.RemoteAccessTemplates == nil {
		return nil, false
	}
	return o.RemoteAccessTemplates, true
}

// HasRemoteAccessTemplates returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasRemoteAccessTemplates() bool {
	if o != nil && o.RemoteAccessTemplates != nil {
		return true
	}

	return false
}

// SetRemoteAccessTemplates gets a reference to the given []MinimalRemoteAccessTemplate and assigns it to the RemoteAccessTemplates field.
func (o *FullPhysicalMachine) SetRemoteAccessTemplates(v []MinimalRemoteAccessTemplate) {
	o.RemoteAccessTemplates = &v
}

// GetBaseDisks returns the BaseDisks field value if set, zero value otherwise.
func (o *FullPhysicalMachine) GetBaseDisks() []FullDisk {
	if o == nil || o.BaseDisks == nil {
		var ret []FullDisk
		return ret
	}
	return *o.BaseDisks
}

// GetBaseDisksOk returns a tuple with the BaseDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPhysicalMachine) GetBaseDisksOk() (*[]FullDisk, bool) {
	if o == nil || o.BaseDisks == nil {
		return nil, false
	}
	return o.BaseDisks, true
}

// HasBaseDisks returns a boolean if a field has been set.
func (o *FullPhysicalMachine) HasBaseDisks() bool {
	if o != nil && o.BaseDisks != nil {
		return true
	}

	return false
}

// SetBaseDisks gets a reference to the given []FullDisk and assigns it to the BaseDisks field.
func (o *FullPhysicalMachine) SetBaseDisks(v []FullDisk) {
	o.BaseDisks = &v
}

func (o FullPhysicalMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.MacAddress != nil {
		toSerialize["macAddress"] = o.MacAddress
	}
	if o.CpuCount != nil {
		toSerialize["cpuCount"] = o.CpuCount
	}
	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}
	if o.Architecture != nil {
		toSerialize["architecture"] = o.Architecture
	}
	if o.Bits != nil {
		toSerialize["bits"] = o.Bits
	}
	if o.OperatingSystem != nil {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if o.RemoteAccessCapable != nil {
		toSerialize["remoteAccessCapable"] = o.RemoteAccessCapable
	}
	if o.RemoteAccessTemplates != nil {
		toSerialize["remoteAccessTemplates"] = o.RemoteAccessTemplates
	}
	if o.BaseDisks != nil {
		toSerialize["baseDisks"] = o.BaseDisks
	}
	return json.Marshal(toSerialize)
}

type NullableFullPhysicalMachine struct {
	value *FullPhysicalMachine
	isSet bool
}

func (v NullableFullPhysicalMachine) Get() *FullPhysicalMachine {
	return v.value
}

func (v *NullableFullPhysicalMachine) Set(val *FullPhysicalMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableFullPhysicalMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableFullPhysicalMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullPhysicalMachine(val *FullPhysicalMachine) *NullableFullPhysicalMachine {
	return &NullableFullPhysicalMachine{value: val, isSet: true}
}

func (v NullableFullPhysicalMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullPhysicalMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
