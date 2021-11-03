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

// InputCons3rtTemplateData struct for InputCons3rtTemplateData
type InputCons3rtTemplateData struct {
	DisplayName           *string                      `json:"displayName,omitempty"`
	VirtRealmTemplateName string                       `json:"virtRealmTemplateName"`
	OperatingSystem       string                       `json:"operatingSystem"`
	Cons3rtAgentInstalled *bool                        `json:"cons3rtAgentInstalled,omitempty"`
	ContainerCapable      *bool                        `json:"containerCapable,omitempty"`
	Disks                 *[]InputDiskForTemplate      `json:"disks,omitempty"`
	HasGpu                *bool                        `json:"hasGpu,omitempty"`
	License               *string                      `json:"license,omitempty"`
	MaxNumCpus            *int32                       `json:"maxNumCpus,omitempty"`
	MaxRamInMegabytes     *int64                       `json:"maxRamInMegabytes,omitempty"`
	Note                  *string                      `json:"note,omitempty"`
	PackageManagementType *string                      `json:"packageManagementType,omitempty"`
	PowerOnDelayOverride  *int32                       `json:"powerOnDelayOverride,omitempty"`
	PowerShellVersion     *string                      `json:"powerShellVersion,omitempty"`
	ServiceManagementType *string                      `json:"serviceManagementType,omitempty"`
	RemoteAccessTemplates *[]InputRemoteAccessTemplate `json:"remoteAccessTemplates,omitempty"`
}

// NewInputCons3rtTemplateData instantiates a new InputCons3rtTemplateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCons3rtTemplateData(virtRealmTemplateName string, operatingSystem string) *InputCons3rtTemplateData {
	this := InputCons3rtTemplateData{}
	this.VirtRealmTemplateName = virtRealmTemplateName
	this.OperatingSystem = operatingSystem
	return &this
}

// NewInputCons3rtTemplateDataWithDefaults instantiates a new InputCons3rtTemplateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCons3rtTemplateDataWithDefaults() *InputCons3rtTemplateData {
	this := InputCons3rtTemplateData{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InputCons3rtTemplateData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVirtRealmTemplateName returns the VirtRealmTemplateName field value
func (o *InputCons3rtTemplateData) GetVirtRealmTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtRealmTemplateName
}

// GetVirtRealmTemplateNameOk returns a tuple with the VirtRealmTemplateName field value
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetVirtRealmTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtRealmTemplateName, true
}

// SetVirtRealmTemplateName sets field value
func (o *InputCons3rtTemplateData) SetVirtRealmTemplateName(v string) {
	o.VirtRealmTemplateName = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *InputCons3rtTemplateData) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *InputCons3rtTemplateData) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetCons3rtAgentInstalled returns the Cons3rtAgentInstalled field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetCons3rtAgentInstalled() bool {
	if o == nil || o.Cons3rtAgentInstalled == nil {
		var ret bool
		return ret
	}
	return *o.Cons3rtAgentInstalled
}

// GetCons3rtAgentInstalledOk returns a tuple with the Cons3rtAgentInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetCons3rtAgentInstalledOk() (*bool, bool) {
	if o == nil || o.Cons3rtAgentInstalled == nil {
		return nil, false
	}
	return o.Cons3rtAgentInstalled, true
}

// HasCons3rtAgentInstalled returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasCons3rtAgentInstalled() bool {
	if o != nil && o.Cons3rtAgentInstalled != nil {
		return true
	}

	return false
}

// SetCons3rtAgentInstalled gets a reference to the given bool and assigns it to the Cons3rtAgentInstalled field.
func (o *InputCons3rtTemplateData) SetCons3rtAgentInstalled(v bool) {
	o.Cons3rtAgentInstalled = &v
}

// GetContainerCapable returns the ContainerCapable field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetContainerCapable() bool {
	if o == nil || o.ContainerCapable == nil {
		var ret bool
		return ret
	}
	return *o.ContainerCapable
}

// GetContainerCapableOk returns a tuple with the ContainerCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetContainerCapableOk() (*bool, bool) {
	if o == nil || o.ContainerCapable == nil {
		return nil, false
	}
	return o.ContainerCapable, true
}

// HasContainerCapable returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasContainerCapable() bool {
	if o != nil && o.ContainerCapable != nil {
		return true
	}

	return false
}

// SetContainerCapable gets a reference to the given bool and assigns it to the ContainerCapable field.
func (o *InputCons3rtTemplateData) SetContainerCapable(v bool) {
	o.ContainerCapable = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetDisks() []InputDiskForTemplate {
	if o == nil || o.Disks == nil {
		var ret []InputDiskForTemplate
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetDisksOk() (*[]InputDiskForTemplate, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []InputDiskForTemplate and assigns it to the Disks field.
func (o *InputCons3rtTemplateData) SetDisks(v []InputDiskForTemplate) {
	o.Disks = &v
}

// GetHasGpu returns the HasGpu field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetHasGpu() bool {
	if o == nil || o.HasGpu == nil {
		var ret bool
		return ret
	}
	return *o.HasGpu
}

// GetHasGpuOk returns a tuple with the HasGpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetHasGpuOk() (*bool, bool) {
	if o == nil || o.HasGpu == nil {
		return nil, false
	}
	return o.HasGpu, true
}

// HasHasGpu returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasHasGpu() bool {
	if o != nil && o.HasGpu != nil {
		return true
	}

	return false
}

// SetHasGpu gets a reference to the given bool and assigns it to the HasGpu field.
func (o *InputCons3rtTemplateData) SetHasGpu(v bool) {
	o.HasGpu = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *InputCons3rtTemplateData) SetLicense(v string) {
	o.License = &v
}

// GetMaxNumCpus returns the MaxNumCpus field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetMaxNumCpus() int32 {
	if o == nil || o.MaxNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumCpus
}

// GetMaxNumCpusOk returns a tuple with the MaxNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetMaxNumCpusOk() (*int32, bool) {
	if o == nil || o.MaxNumCpus == nil {
		return nil, false
	}
	return o.MaxNumCpus, true
}

// HasMaxNumCpus returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasMaxNumCpus() bool {
	if o != nil && o.MaxNumCpus != nil {
		return true
	}

	return false
}

// SetMaxNumCpus gets a reference to the given int32 and assigns it to the MaxNumCpus field.
func (o *InputCons3rtTemplateData) SetMaxNumCpus(v int32) {
	o.MaxNumCpus = &v
}

// GetMaxRamInMegabytes returns the MaxRamInMegabytes field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetMaxRamInMegabytes() int64 {
	if o == nil || o.MaxRamInMegabytes == nil {
		var ret int64
		return ret
	}
	return *o.MaxRamInMegabytes
}

// GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetMaxRamInMegabytesOk() (*int64, bool) {
	if o == nil || o.MaxRamInMegabytes == nil {
		return nil, false
	}
	return o.MaxRamInMegabytes, true
}

// HasMaxRamInMegabytes returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasMaxRamInMegabytes() bool {
	if o != nil && o.MaxRamInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxRamInMegabytes gets a reference to the given int64 and assigns it to the MaxRamInMegabytes field.
func (o *InputCons3rtTemplateData) SetMaxRamInMegabytes(v int64) {
	o.MaxRamInMegabytes = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *InputCons3rtTemplateData) SetNote(v string) {
	o.Note = &v
}

// GetPackageManagementType returns the PackageManagementType field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetPackageManagementType() string {
	if o == nil || o.PackageManagementType == nil {
		var ret string
		return ret
	}
	return *o.PackageManagementType
}

// GetPackageManagementTypeOk returns a tuple with the PackageManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetPackageManagementTypeOk() (*string, bool) {
	if o == nil || o.PackageManagementType == nil {
		return nil, false
	}
	return o.PackageManagementType, true
}

// HasPackageManagementType returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasPackageManagementType() bool {
	if o != nil && o.PackageManagementType != nil {
		return true
	}

	return false
}

// SetPackageManagementType gets a reference to the given string and assigns it to the PackageManagementType field.
func (o *InputCons3rtTemplateData) SetPackageManagementType(v string) {
	o.PackageManagementType = &v
}

// GetPowerOnDelayOverride returns the PowerOnDelayOverride field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetPowerOnDelayOverride() int32 {
	if o == nil || o.PowerOnDelayOverride == nil {
		var ret int32
		return ret
	}
	return *o.PowerOnDelayOverride
}

// GetPowerOnDelayOverrideOk returns a tuple with the PowerOnDelayOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetPowerOnDelayOverrideOk() (*int32, bool) {
	if o == nil || o.PowerOnDelayOverride == nil {
		return nil, false
	}
	return o.PowerOnDelayOverride, true
}

// HasPowerOnDelayOverride returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasPowerOnDelayOverride() bool {
	if o != nil && o.PowerOnDelayOverride != nil {
		return true
	}

	return false
}

// SetPowerOnDelayOverride gets a reference to the given int32 and assigns it to the PowerOnDelayOverride field.
func (o *InputCons3rtTemplateData) SetPowerOnDelayOverride(v int32) {
	o.PowerOnDelayOverride = &v
}

// GetPowerShellVersion returns the PowerShellVersion field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetPowerShellVersion() string {
	if o == nil || o.PowerShellVersion == nil {
		var ret string
		return ret
	}
	return *o.PowerShellVersion
}

// GetPowerShellVersionOk returns a tuple with the PowerShellVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetPowerShellVersionOk() (*string, bool) {
	if o == nil || o.PowerShellVersion == nil {
		return nil, false
	}
	return o.PowerShellVersion, true
}

// HasPowerShellVersion returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasPowerShellVersion() bool {
	if o != nil && o.PowerShellVersion != nil {
		return true
	}

	return false
}

// SetPowerShellVersion gets a reference to the given string and assigns it to the PowerShellVersion field.
func (o *InputCons3rtTemplateData) SetPowerShellVersion(v string) {
	o.PowerShellVersion = &v
}

// GetServiceManagementType returns the ServiceManagementType field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetServiceManagementType() string {
	if o == nil || o.ServiceManagementType == nil {
		var ret string
		return ret
	}
	return *o.ServiceManagementType
}

// GetServiceManagementTypeOk returns a tuple with the ServiceManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetServiceManagementTypeOk() (*string, bool) {
	if o == nil || o.ServiceManagementType == nil {
		return nil, false
	}
	return o.ServiceManagementType, true
}

// HasServiceManagementType returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasServiceManagementType() bool {
	if o != nil && o.ServiceManagementType != nil {
		return true
	}

	return false
}

// SetServiceManagementType gets a reference to the given string and assigns it to the ServiceManagementType field.
func (o *InputCons3rtTemplateData) SetServiceManagementType(v string) {
	o.ServiceManagementType = &v
}

// GetRemoteAccessTemplates returns the RemoteAccessTemplates field value if set, zero value otherwise.
func (o *InputCons3rtTemplateData) GetRemoteAccessTemplates() []InputRemoteAccessTemplate {
	if o == nil || o.RemoteAccessTemplates == nil {
		var ret []InputRemoteAccessTemplate
		return ret
	}
	return *o.RemoteAccessTemplates
}

// GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCons3rtTemplateData) GetRemoteAccessTemplatesOk() (*[]InputRemoteAccessTemplate, bool) {
	if o == nil || o.RemoteAccessTemplates == nil {
		return nil, false
	}
	return o.RemoteAccessTemplates, true
}

// HasRemoteAccessTemplates returns a boolean if a field has been set.
func (o *InputCons3rtTemplateData) HasRemoteAccessTemplates() bool {
	if o != nil && o.RemoteAccessTemplates != nil {
		return true
	}

	return false
}

// SetRemoteAccessTemplates gets a reference to the given []InputRemoteAccessTemplate and assigns it to the RemoteAccessTemplates field.
func (o *InputCons3rtTemplateData) SetRemoteAccessTemplates(v []InputRemoteAccessTemplate) {
	o.RemoteAccessTemplates = &v
}

func (o InputCons3rtTemplateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["virtRealmTemplateName"] = o.VirtRealmTemplateName
	}
	if true {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if o.Cons3rtAgentInstalled != nil {
		toSerialize["cons3rtAgentInstalled"] = o.Cons3rtAgentInstalled
	}
	if o.ContainerCapable != nil {
		toSerialize["containerCapable"] = o.ContainerCapable
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	if o.HasGpu != nil {
		toSerialize["hasGpu"] = o.HasGpu
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.MaxNumCpus != nil {
		toSerialize["maxNumCpus"] = o.MaxNumCpus
	}
	if o.MaxRamInMegabytes != nil {
		toSerialize["maxRamInMegabytes"] = o.MaxRamInMegabytes
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.PackageManagementType != nil {
		toSerialize["packageManagementType"] = o.PackageManagementType
	}
	if o.PowerOnDelayOverride != nil {
		toSerialize["powerOnDelayOverride"] = o.PowerOnDelayOverride
	}
	if o.PowerShellVersion != nil {
		toSerialize["powerShellVersion"] = o.PowerShellVersion
	}
	if o.ServiceManagementType != nil {
		toSerialize["serviceManagementType"] = o.ServiceManagementType
	}
	if o.RemoteAccessTemplates != nil {
		toSerialize["remoteAccessTemplates"] = o.RemoteAccessTemplates
	}
	return json.Marshal(toSerialize)
}

type NullableInputCons3rtTemplateData struct {
	value *InputCons3rtTemplateData
	isSet bool
}

func (v NullableInputCons3rtTemplateData) Get() *InputCons3rtTemplateData {
	return v.value
}

func (v *NullableInputCons3rtTemplateData) Set(val *InputCons3rtTemplateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCons3rtTemplateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCons3rtTemplateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCons3rtTemplateData(val *InputCons3rtTemplateData) *NullableInputCons3rtTemplateData {
	return &NullableInputCons3rtTemplateData{value: val, isSet: true}
}

func (v NullableInputCons3rtTemplateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCons3rtTemplateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
