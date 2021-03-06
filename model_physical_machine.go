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

// PhysicalMachine struct for PhysicalMachine
type PhysicalMachine struct {
	Id                    *string                 `json:"id,omitempty"`
	TrustedProjects       *[]Project              `json:"trustedProjects,omitempty"`
	Creator               *User                   `json:"creator,omitempty"`
	DependentAssetCount   *int32                  `json:"dependentAssetCount,omitempty"`
	Description           *string                 `json:"description,omitempty"`
	Metadata              *Metadata               `json:"metadata,omitempty"`
	Name                  *string                 `json:"name,omitempty"`
	Offline               *bool                   `json:"offline,omitempty"`
	OwningProject         *Project                `json:"owningProject,omitempty"`
	State                 *string                 `json:"state,omitempty"`
	Visibility            *string                 `json:"visibility,omitempty"`
	ImpactLevel           *string                 `json:"impactLevel,omitempty"`
	Categories            *[]Category             `json:"categories,omitempty"`
	Architecture          *string                 `json:"architecture,omitempty"`
	Bits                  *string                 `json:"bits,omitempty"`
	OperatingSystem       *string                 `json:"operatingSystem,omitempty"`
	BaseDisks             *[]Disk                 `json:"baseDisks,omitempty"`
	RemoteAccessTemplates *[]RemoteAccessTemplate `json:"remoteAccessTemplates,omitempty"`
	HostName              *string                 `json:"hostName,omitempty"`
	IpAddress             *string                 `json:"ipAddress,omitempty"`
	MacAddress            *string                 `json:"macAddress,omitempty"`
	CpuCount              *int32                  `json:"cpuCount,omitempty"`
	Ram                   *int32                  `json:"ram,omitempty"`
}

// NewPhysicalMachine instantiates a new PhysicalMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalMachine() *PhysicalMachine {
	this := PhysicalMachine{}
	return &this
}

// NewPhysicalMachineWithDefaults instantiates a new PhysicalMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalMachineWithDefaults() *PhysicalMachine {
	this := PhysicalMachine{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PhysicalMachine) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PhysicalMachine) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PhysicalMachine) SetId(v string) {
	o.Id = &v
}

// GetTrustedProjects returns the TrustedProjects field value if set, zero value otherwise.
func (o *PhysicalMachine) GetTrustedProjects() []Project {
	if o == nil || o.TrustedProjects == nil {
		var ret []Project
		return ret
	}
	return *o.TrustedProjects
}

// GetTrustedProjectsOk returns a tuple with the TrustedProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetTrustedProjectsOk() (*[]Project, bool) {
	if o == nil || o.TrustedProjects == nil {
		return nil, false
	}
	return o.TrustedProjects, true
}

// HasTrustedProjects returns a boolean if a field has been set.
func (o *PhysicalMachine) HasTrustedProjects() bool {
	if o != nil && o.TrustedProjects != nil {
		return true
	}

	return false
}

// SetTrustedProjects gets a reference to the given []Project and assigns it to the TrustedProjects field.
func (o *PhysicalMachine) SetTrustedProjects(v []Project) {
	o.TrustedProjects = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *PhysicalMachine) GetCreator() User {
	if o == nil || o.Creator == nil {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetCreatorOk() (*User, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *PhysicalMachine) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *PhysicalMachine) SetCreator(v User) {
	o.Creator = &v
}

// GetDependentAssetCount returns the DependentAssetCount field value if set, zero value otherwise.
func (o *PhysicalMachine) GetDependentAssetCount() int32 {
	if o == nil || o.DependentAssetCount == nil {
		var ret int32
		return ret
	}
	return *o.DependentAssetCount
}

// GetDependentAssetCountOk returns a tuple with the DependentAssetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetDependentAssetCountOk() (*int32, bool) {
	if o == nil || o.DependentAssetCount == nil {
		return nil, false
	}
	return o.DependentAssetCount, true
}

// HasDependentAssetCount returns a boolean if a field has been set.
func (o *PhysicalMachine) HasDependentAssetCount() bool {
	if o != nil && o.DependentAssetCount != nil {
		return true
	}

	return false
}

// SetDependentAssetCount gets a reference to the given int32 and assigns it to the DependentAssetCount field.
func (o *PhysicalMachine) SetDependentAssetCount(v int32) {
	o.DependentAssetCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PhysicalMachine) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PhysicalMachine) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PhysicalMachine) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PhysicalMachine) GetMetadata() Metadata {
	if o == nil || o.Metadata == nil {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetMetadataOk() (*Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PhysicalMachine) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *PhysicalMachine) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PhysicalMachine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PhysicalMachine) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PhysicalMachine) SetName(v string) {
	o.Name = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *PhysicalMachine) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *PhysicalMachine) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *PhysicalMachine) SetOffline(v bool) {
	o.Offline = &v
}

// GetOwningProject returns the OwningProject field value if set, zero value otherwise.
func (o *PhysicalMachine) GetOwningProject() Project {
	if o == nil || o.OwningProject == nil {
		var ret Project
		return ret
	}
	return *o.OwningProject
}

// GetOwningProjectOk returns a tuple with the OwningProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetOwningProjectOk() (*Project, bool) {
	if o == nil || o.OwningProject == nil {
		return nil, false
	}
	return o.OwningProject, true
}

// HasOwningProject returns a boolean if a field has been set.
func (o *PhysicalMachine) HasOwningProject() bool {
	if o != nil && o.OwningProject != nil {
		return true
	}

	return false
}

// SetOwningProject gets a reference to the given Project and assigns it to the OwningProject field.
func (o *PhysicalMachine) SetOwningProject(v Project) {
	o.OwningProject = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PhysicalMachine) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PhysicalMachine) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PhysicalMachine) SetState(v string) {
	o.State = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *PhysicalMachine) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *PhysicalMachine) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *PhysicalMachine) SetVisibility(v string) {
	o.Visibility = &v
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise.
func (o *PhysicalMachine) GetImpactLevel() string {
	if o == nil || o.ImpactLevel == nil {
		var ret string
		return ret
	}
	return *o.ImpactLevel
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetImpactLevelOk() (*string, bool) {
	if o == nil || o.ImpactLevel == nil {
		return nil, false
	}
	return o.ImpactLevel, true
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *PhysicalMachine) HasImpactLevel() bool {
	if o != nil && o.ImpactLevel != nil {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given string and assigns it to the ImpactLevel field.
func (o *PhysicalMachine) SetImpactLevel(v string) {
	o.ImpactLevel = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *PhysicalMachine) GetCategories() []Category {
	if o == nil || o.Categories == nil {
		var ret []Category
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetCategoriesOk() (*[]Category, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *PhysicalMachine) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *PhysicalMachine) SetCategories(v []Category) {
	o.Categories = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *PhysicalMachine) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *PhysicalMachine) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *PhysicalMachine) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *PhysicalMachine) GetBits() string {
	if o == nil || o.Bits == nil {
		var ret string
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetBitsOk() (*string, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *PhysicalMachine) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given string and assigns it to the Bits field.
func (o *PhysicalMachine) SetBits(v string) {
	o.Bits = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *PhysicalMachine) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *PhysicalMachine) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *PhysicalMachine) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetBaseDisks returns the BaseDisks field value if set, zero value otherwise.
func (o *PhysicalMachine) GetBaseDisks() []Disk {
	if o == nil || o.BaseDisks == nil {
		var ret []Disk
		return ret
	}
	return *o.BaseDisks
}

// GetBaseDisksOk returns a tuple with the BaseDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetBaseDisksOk() (*[]Disk, bool) {
	if o == nil || o.BaseDisks == nil {
		return nil, false
	}
	return o.BaseDisks, true
}

// HasBaseDisks returns a boolean if a field has been set.
func (o *PhysicalMachine) HasBaseDisks() bool {
	if o != nil && o.BaseDisks != nil {
		return true
	}

	return false
}

// SetBaseDisks gets a reference to the given []Disk and assigns it to the BaseDisks field.
func (o *PhysicalMachine) SetBaseDisks(v []Disk) {
	o.BaseDisks = &v
}

// GetRemoteAccessTemplates returns the RemoteAccessTemplates field value if set, zero value otherwise.
func (o *PhysicalMachine) GetRemoteAccessTemplates() []RemoteAccessTemplate {
	if o == nil || o.RemoteAccessTemplates == nil {
		var ret []RemoteAccessTemplate
		return ret
	}
	return *o.RemoteAccessTemplates
}

// GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetRemoteAccessTemplatesOk() (*[]RemoteAccessTemplate, bool) {
	if o == nil || o.RemoteAccessTemplates == nil {
		return nil, false
	}
	return o.RemoteAccessTemplates, true
}

// HasRemoteAccessTemplates returns a boolean if a field has been set.
func (o *PhysicalMachine) HasRemoteAccessTemplates() bool {
	if o != nil && o.RemoteAccessTemplates != nil {
		return true
	}

	return false
}

// SetRemoteAccessTemplates gets a reference to the given []RemoteAccessTemplate and assigns it to the RemoteAccessTemplates field.
func (o *PhysicalMachine) SetRemoteAccessTemplates(v []RemoteAccessTemplate) {
	o.RemoteAccessTemplates = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *PhysicalMachine) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *PhysicalMachine) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *PhysicalMachine) SetHostName(v string) {
	o.HostName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *PhysicalMachine) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *PhysicalMachine) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *PhysicalMachine) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *PhysicalMachine) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *PhysicalMachine) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *PhysicalMachine) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *PhysicalMachine) GetCpuCount() int32 {
	if o == nil || o.CpuCount == nil {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetCpuCountOk() (*int32, bool) {
	if o == nil || o.CpuCount == nil {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *PhysicalMachine) HasCpuCount() bool {
	if o != nil && o.CpuCount != nil {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *PhysicalMachine) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *PhysicalMachine) GetRam() int32 {
	if o == nil || o.Ram == nil {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachine) GetRamOk() (*int32, bool) {
	if o == nil || o.Ram == nil {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *PhysicalMachine) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *PhysicalMachine) SetRam(v int32) {
	o.Ram = &v
}

func (o PhysicalMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TrustedProjects != nil {
		toSerialize["trustedProjects"] = o.TrustedProjects
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.DependentAssetCount != nil {
		toSerialize["dependentAssetCount"] = o.DependentAssetCount
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	if o.OwningProject != nil {
		toSerialize["owningProject"] = o.OwningProject
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.ImpactLevel != nil {
		toSerialize["impactLevel"] = o.ImpactLevel
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
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
	if o.BaseDisks != nil {
		toSerialize["baseDisks"] = o.BaseDisks
	}
	if o.RemoteAccessTemplates != nil {
		toSerialize["remoteAccessTemplates"] = o.RemoteAccessTemplates
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
	return json.Marshal(toSerialize)
}

type NullablePhysicalMachine struct {
	value *PhysicalMachine
	isSet bool
}

func (v NullablePhysicalMachine) Get() *PhysicalMachine {
	return v.value
}

func (v *NullablePhysicalMachine) Set(val *PhysicalMachine) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalMachine) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalMachine(val *PhysicalMachine) *NullablePhysicalMachine {
	return &NullablePhysicalMachine{value: val, isSet: true}
}

func (v NullablePhysicalMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
