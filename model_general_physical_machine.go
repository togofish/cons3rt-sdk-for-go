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

// GeneralPhysicalMachine struct for GeneralPhysicalMachine
type GeneralPhysicalMachine struct {
	BaseDisks       *[]FullDisk `json:"baseDisks,omitempty"`
	Id              *string     `json:"id,omitempty"`
	Name            *string     `json:"name,omitempty"`
	Description     *string     `json:"description,omitempty"`
	Offline         *bool       `json:"offline,omitempty"`
	State           *string     `json:"state,omitempty"`
	Visibility      *string     `json:"visibility,omitempty"`
	HostName        *string     `json:"hostName,omitempty"`
	IpAddress       *string     `json:"ipAddress,omitempty"`
	MacAddress      *string     `json:"macAddress,omitempty"`
	CpuCount        *int32      `json:"cpuCount,omitempty"`
	Ram             *int32      `json:"ram,omitempty"`
	Architecture    *string     `json:"architecture,omitempty"`
	Bits            *string     `json:"bits,omitempty"`
	OperatingSystem *string     `json:"operatingSystem,omitempty"`
}

// NewGeneralPhysicalMachine instantiates a new GeneralPhysicalMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralPhysicalMachine() *GeneralPhysicalMachine {
	this := GeneralPhysicalMachine{}
	return &this
}

// NewGeneralPhysicalMachineWithDefaults instantiates a new GeneralPhysicalMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralPhysicalMachineWithDefaults() *GeneralPhysicalMachine {
	this := GeneralPhysicalMachine{}
	return &this
}

// GetBaseDisks returns the BaseDisks field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetBaseDisks() []FullDisk {
	if o == nil || o.BaseDisks == nil {
		var ret []FullDisk
		return ret
	}
	return *o.BaseDisks
}

// GetBaseDisksOk returns a tuple with the BaseDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetBaseDisksOk() (*[]FullDisk, bool) {
	if o == nil || o.BaseDisks == nil {
		return nil, false
	}
	return o.BaseDisks, true
}

// HasBaseDisks returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasBaseDisks() bool {
	if o != nil && o.BaseDisks != nil {
		return true
	}

	return false
}

// SetBaseDisks gets a reference to the given []FullDisk and assigns it to the BaseDisks field.
func (o *GeneralPhysicalMachine) SetBaseDisks(v []FullDisk) {
	o.BaseDisks = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GeneralPhysicalMachine) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GeneralPhysicalMachine) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GeneralPhysicalMachine) SetDescription(v string) {
	o.Description = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *GeneralPhysicalMachine) SetOffline(v bool) {
	o.Offline = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GeneralPhysicalMachine) SetState(v string) {
	o.State = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *GeneralPhysicalMachine) SetVisibility(v string) {
	o.Visibility = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *GeneralPhysicalMachine) SetHostName(v string) {
	o.HostName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *GeneralPhysicalMachine) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *GeneralPhysicalMachine) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetCpuCount() int32 {
	if o == nil || o.CpuCount == nil {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetCpuCountOk() (*int32, bool) {
	if o == nil || o.CpuCount == nil {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasCpuCount() bool {
	if o != nil && o.CpuCount != nil {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *GeneralPhysicalMachine) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetRam() int32 {
	if o == nil || o.Ram == nil {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetRamOk() (*int32, bool) {
	if o == nil || o.Ram == nil {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *GeneralPhysicalMachine) SetRam(v int32) {
	o.Ram = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *GeneralPhysicalMachine) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetBits() string {
	if o == nil || o.Bits == nil {
		var ret string
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetBitsOk() (*string, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given string and assigns it to the Bits field.
func (o *GeneralPhysicalMachine) SetBits(v string) {
	o.Bits = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *GeneralPhysicalMachine) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralPhysicalMachine) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *GeneralPhysicalMachine) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *GeneralPhysicalMachine) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

func (o GeneralPhysicalMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseDisks != nil {
		toSerialize["baseDisks"] = o.BaseDisks
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
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
	return json.Marshal(toSerialize)
}

type NullableGeneralPhysicalMachine struct {
	value *GeneralPhysicalMachine
	isSet bool
}

func (v NullableGeneralPhysicalMachine) Get() *GeneralPhysicalMachine {
	return v.value
}

func (v *NullableGeneralPhysicalMachine) Set(val *GeneralPhysicalMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralPhysicalMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralPhysicalMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralPhysicalMachine(val *GeneralPhysicalMachine) *NullableGeneralPhysicalMachine {
	return &NullableGeneralPhysicalMachine{value: val, isSet: true}
}

func (v NullableGeneralPhysicalMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralPhysicalMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
