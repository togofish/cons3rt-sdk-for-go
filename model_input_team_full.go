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

// InputTeamFull struct for InputTeamFull
type InputTeamFull struct {
	AssetBundleInstallerEnabled *bool `json:"assetBundleInstallerEnabled,omitempty"`
	AssetBypassScanningEnabled *bool `json:"assetBypassScanningEnabled,omitempty"`
	AvailabilityZoneEnabled *bool `json:"availabilityZoneEnabled,omitempty"`
	BypassScanningEnabled *bool `json:"bypassScanningEnabled,omitempty"`
	LeadUser InputUser `json:"leadUser"`
	GpuTypeMaximums *map[string]int32 `json:"gpuTypeMaximums,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Id *int32 `json:"id,omitempty"`
	TeamManagers []InputUser `json:"teamManagers"`
	MaxManagedVirtualizationRealms *int32 `json:"maxManagedVirtualizationRealms,omitempty"`
	MaxNumCpus *int32 `json:"maxNumCpus,omitempty"`
	MaxNumGpus *int32 `json:"maxNumGpus,omitempty"`
	MaxProjects *int32 `json:"maxProjects,omitempty"`
	MaxRamInMegabytes *int32 `json:"maxRamInMegabytes,omitempty"`
	MaxSharedRemoteAccessSessions *int32 `json:"maxSharedRemoteAccessSessions,omitempty"`
	MaxStorageInMegabytes *int32 `json:"maxStorageInMegabytes,omitempty"`
	MaxUsers *int32 `json:"maxUsers,omitempty"`
	MaxVirtualMachines *int32 `json:"maxVirtualMachines,omitempty"`
	Name string `json:"name"`
	OrderNumber *string `json:"orderNumber,omitempty"`
	ContactInfo PocInfo `json:"contactInfo"`
	PowerScheduleEnabled *bool `json:"powerScheduleEnabled,omitempty"`
	Private *bool `json:"private,omitempty"`
	RdpClientProxyEnabled *bool `json:"rdpClientProxyEnabled,omitempty"`
	RdpClientSessionDuration *int32 `json:"rdpClientSessionDuration,omitempty"`
	SnapshotEnabled *bool `json:"snapshotEnabled,omitempty"`
	State string `json:"state"`
	ValidUntil int32 `json:"validUntil"`
}

// NewInputTeamFull instantiates a new InputTeamFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputTeamFull(leadUser InputUser, teamManagers []InputUser, name string, contactInfo PocInfo, state string, validUntil int32) *InputTeamFull {
	this := InputTeamFull{}
	this.LeadUser = leadUser
	this.TeamManagers = teamManagers
	this.Name = name
	this.ContactInfo = contactInfo
	this.State = state
	this.ValidUntil = validUntil
	return &this
}

// NewInputTeamFullWithDefaults instantiates a new InputTeamFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputTeamFullWithDefaults() *InputTeamFull {
	this := InputTeamFull{}
	return &this
}

// GetAssetBundleInstallerEnabled returns the AssetBundleInstallerEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetAssetBundleInstallerEnabled() bool {
	if o == nil || o.AssetBundleInstallerEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AssetBundleInstallerEnabled
}

// GetAssetBundleInstallerEnabledOk returns a tuple with the AssetBundleInstallerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetAssetBundleInstallerEnabledOk() (*bool, bool) {
	if o == nil || o.AssetBundleInstallerEnabled == nil {
		return nil, false
	}
	return o.AssetBundleInstallerEnabled, true
}

// HasAssetBundleInstallerEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasAssetBundleInstallerEnabled() bool {
	if o != nil && o.AssetBundleInstallerEnabled != nil {
		return true
	}

	return false
}

// SetAssetBundleInstallerEnabled gets a reference to the given bool and assigns it to the AssetBundleInstallerEnabled field.
func (o *InputTeamFull) SetAssetBundleInstallerEnabled(v bool) {
	o.AssetBundleInstallerEnabled = &v
}

// GetAssetBypassScanningEnabled returns the AssetBypassScanningEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetAssetBypassScanningEnabled() bool {
	if o == nil || o.AssetBypassScanningEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AssetBypassScanningEnabled
}

// GetAssetBypassScanningEnabledOk returns a tuple with the AssetBypassScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetAssetBypassScanningEnabledOk() (*bool, bool) {
	if o == nil || o.AssetBypassScanningEnabled == nil {
		return nil, false
	}
	return o.AssetBypassScanningEnabled, true
}

// HasAssetBypassScanningEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasAssetBypassScanningEnabled() bool {
	if o != nil && o.AssetBypassScanningEnabled != nil {
		return true
	}

	return false
}

// SetAssetBypassScanningEnabled gets a reference to the given bool and assigns it to the AssetBypassScanningEnabled field.
func (o *InputTeamFull) SetAssetBypassScanningEnabled(v bool) {
	o.AssetBypassScanningEnabled = &v
}

// GetAvailabilityZoneEnabled returns the AvailabilityZoneEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetAvailabilityZoneEnabled() bool {
	if o == nil || o.AvailabilityZoneEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityZoneEnabled
}

// GetAvailabilityZoneEnabledOk returns a tuple with the AvailabilityZoneEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetAvailabilityZoneEnabledOk() (*bool, bool) {
	if o == nil || o.AvailabilityZoneEnabled == nil {
		return nil, false
	}
	return o.AvailabilityZoneEnabled, true
}

// HasAvailabilityZoneEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasAvailabilityZoneEnabled() bool {
	if o != nil && o.AvailabilityZoneEnabled != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneEnabled gets a reference to the given bool and assigns it to the AvailabilityZoneEnabled field.
func (o *InputTeamFull) SetAvailabilityZoneEnabled(v bool) {
	o.AvailabilityZoneEnabled = &v
}

// GetBypassScanningEnabled returns the BypassScanningEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetBypassScanningEnabled() bool {
	if o == nil || o.BypassScanningEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BypassScanningEnabled
}

// GetBypassScanningEnabledOk returns a tuple with the BypassScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetBypassScanningEnabledOk() (*bool, bool) {
	if o == nil || o.BypassScanningEnabled == nil {
		return nil, false
	}
	return o.BypassScanningEnabled, true
}

// HasBypassScanningEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasBypassScanningEnabled() bool {
	if o != nil && o.BypassScanningEnabled != nil {
		return true
	}

	return false
}

// SetBypassScanningEnabled gets a reference to the given bool and assigns it to the BypassScanningEnabled field.
func (o *InputTeamFull) SetBypassScanningEnabled(v bool) {
	o.BypassScanningEnabled = &v
}

// GetLeadUser returns the LeadUser field value
func (o *InputTeamFull) GetLeadUser() InputUser {
	if o == nil {
		var ret InputUser
		return ret
	}

	return o.LeadUser
}

// GetLeadUserOk returns a tuple with the LeadUser field value
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetLeadUserOk() (*InputUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LeadUser, true
}

// SetLeadUser sets field value
func (o *InputTeamFull) SetLeadUser(v InputUser) {
	o.LeadUser = v
}

// GetGpuTypeMaximums returns the GpuTypeMaximums field value if set, zero value otherwise.
func (o *InputTeamFull) GetGpuTypeMaximums() map[string]int32 {
	if o == nil || o.GpuTypeMaximums == nil {
		var ret map[string]int32
		return ret
	}
	return *o.GpuTypeMaximums
}

// GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetGpuTypeMaximumsOk() (*map[string]int32, bool) {
	if o == nil || o.GpuTypeMaximums == nil {
		return nil, false
	}
	return o.GpuTypeMaximums, true
}

// HasGpuTypeMaximums returns a boolean if a field has been set.
func (o *InputTeamFull) HasGpuTypeMaximums() bool {
	if o != nil && o.GpuTypeMaximums != nil {
		return true
	}

	return false
}

// SetGpuTypeMaximums gets a reference to the given map[string]int32 and assigns it to the GpuTypeMaximums field.
func (o *InputTeamFull) SetGpuTypeMaximums(v map[string]int32) {
	o.GpuTypeMaximums = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *InputTeamFull) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *InputTeamFull) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *InputTeamFull) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InputTeamFull) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InputTeamFull) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InputTeamFull) SetId(v int32) {
	o.Id = &v
}

// GetTeamManagers returns the TeamManagers field value
func (o *InputTeamFull) GetTeamManagers() []InputUser {
	if o == nil {
		var ret []InputUser
		return ret
	}

	return o.TeamManagers
}

// GetTeamManagersOk returns a tuple with the TeamManagers field value
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetTeamManagersOk() (*[]InputUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TeamManagers, true
}

// SetTeamManagers sets field value
func (o *InputTeamFull) SetTeamManagers(v []InputUser) {
	o.TeamManagers = v
}

// GetMaxManagedVirtualizationRealms returns the MaxManagedVirtualizationRealms field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxManagedVirtualizationRealms() int32 {
	if o == nil || o.MaxManagedVirtualizationRealms == nil {
		var ret int32
		return ret
	}
	return *o.MaxManagedVirtualizationRealms
}

// GetMaxManagedVirtualizationRealmsOk returns a tuple with the MaxManagedVirtualizationRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxManagedVirtualizationRealmsOk() (*int32, bool) {
	if o == nil || o.MaxManagedVirtualizationRealms == nil {
		return nil, false
	}
	return o.MaxManagedVirtualizationRealms, true
}

// HasMaxManagedVirtualizationRealms returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxManagedVirtualizationRealms() bool {
	if o != nil && o.MaxManagedVirtualizationRealms != nil {
		return true
	}

	return false
}

// SetMaxManagedVirtualizationRealms gets a reference to the given int32 and assigns it to the MaxManagedVirtualizationRealms field.
func (o *InputTeamFull) SetMaxManagedVirtualizationRealms(v int32) {
	o.MaxManagedVirtualizationRealms = &v
}

// GetMaxNumCpus returns the MaxNumCpus field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxNumCpus() int32 {
	if o == nil || o.MaxNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumCpus
}

// GetMaxNumCpusOk returns a tuple with the MaxNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxNumCpusOk() (*int32, bool) {
	if o == nil || o.MaxNumCpus == nil {
		return nil, false
	}
	return o.MaxNumCpus, true
}

// HasMaxNumCpus returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxNumCpus() bool {
	if o != nil && o.MaxNumCpus != nil {
		return true
	}

	return false
}

// SetMaxNumCpus gets a reference to the given int32 and assigns it to the MaxNumCpus field.
func (o *InputTeamFull) SetMaxNumCpus(v int32) {
	o.MaxNumCpus = &v
}

// GetMaxNumGpus returns the MaxNumGpus field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxNumGpus() int32 {
	if o == nil || o.MaxNumGpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumGpus
}

// GetMaxNumGpusOk returns a tuple with the MaxNumGpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxNumGpusOk() (*int32, bool) {
	if o == nil || o.MaxNumGpus == nil {
		return nil, false
	}
	return o.MaxNumGpus, true
}

// HasMaxNumGpus returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxNumGpus() bool {
	if o != nil && o.MaxNumGpus != nil {
		return true
	}

	return false
}

// SetMaxNumGpus gets a reference to the given int32 and assigns it to the MaxNumGpus field.
func (o *InputTeamFull) SetMaxNumGpus(v int32) {
	o.MaxNumGpus = &v
}

// GetMaxProjects returns the MaxProjects field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxProjects() int32 {
	if o == nil || o.MaxProjects == nil {
		var ret int32
		return ret
	}
	return *o.MaxProjects
}

// GetMaxProjectsOk returns a tuple with the MaxProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxProjectsOk() (*int32, bool) {
	if o == nil || o.MaxProjects == nil {
		return nil, false
	}
	return o.MaxProjects, true
}

// HasMaxProjects returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxProjects() bool {
	if o != nil && o.MaxProjects != nil {
		return true
	}

	return false
}

// SetMaxProjects gets a reference to the given int32 and assigns it to the MaxProjects field.
func (o *InputTeamFull) SetMaxProjects(v int32) {
	o.MaxProjects = &v
}

// GetMaxRamInMegabytes returns the MaxRamInMegabytes field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxRamInMegabytes() int32 {
	if o == nil || o.MaxRamInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRamInMegabytes
}

// GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxRamInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaxRamInMegabytes == nil {
		return nil, false
	}
	return o.MaxRamInMegabytes, true
}

// HasMaxRamInMegabytes returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxRamInMegabytes() bool {
	if o != nil && o.MaxRamInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxRamInMegabytes gets a reference to the given int32 and assigns it to the MaxRamInMegabytes field.
func (o *InputTeamFull) SetMaxRamInMegabytes(v int32) {
	o.MaxRamInMegabytes = &v
}

// GetMaxSharedRemoteAccessSessions returns the MaxSharedRemoteAccessSessions field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxSharedRemoteAccessSessions() int32 {
	if o == nil || o.MaxSharedRemoteAccessSessions == nil {
		var ret int32
		return ret
	}
	return *o.MaxSharedRemoteAccessSessions
}

// GetMaxSharedRemoteAccessSessionsOk returns a tuple with the MaxSharedRemoteAccessSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxSharedRemoteAccessSessionsOk() (*int32, bool) {
	if o == nil || o.MaxSharedRemoteAccessSessions == nil {
		return nil, false
	}
	return o.MaxSharedRemoteAccessSessions, true
}

// HasMaxSharedRemoteAccessSessions returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxSharedRemoteAccessSessions() bool {
	if o != nil && o.MaxSharedRemoteAccessSessions != nil {
		return true
	}

	return false
}

// SetMaxSharedRemoteAccessSessions gets a reference to the given int32 and assigns it to the MaxSharedRemoteAccessSessions field.
func (o *InputTeamFull) SetMaxSharedRemoteAccessSessions(v int32) {
	o.MaxSharedRemoteAccessSessions = &v
}

// GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxStorageInMegabytes() int32 {
	if o == nil || o.MaxStorageInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxStorageInMegabytes
}

// GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxStorageInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaxStorageInMegabytes == nil {
		return nil, false
	}
	return o.MaxStorageInMegabytes, true
}

// HasMaxStorageInMegabytes returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxStorageInMegabytes() bool {
	if o != nil && o.MaxStorageInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxStorageInMegabytes gets a reference to the given int32 and assigns it to the MaxStorageInMegabytes field.
func (o *InputTeamFull) SetMaxStorageInMegabytes(v int32) {
	o.MaxStorageInMegabytes = &v
}

// GetMaxUsers returns the MaxUsers field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxUsers() int32 {
	if o == nil || o.MaxUsers == nil {
		var ret int32
		return ret
	}
	return *o.MaxUsers
}

// GetMaxUsersOk returns a tuple with the MaxUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxUsersOk() (*int32, bool) {
	if o == nil || o.MaxUsers == nil {
		return nil, false
	}
	return o.MaxUsers, true
}

// HasMaxUsers returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxUsers() bool {
	if o != nil && o.MaxUsers != nil {
		return true
	}

	return false
}

// SetMaxUsers gets a reference to the given int32 and assigns it to the MaxUsers field.
func (o *InputTeamFull) SetMaxUsers(v int32) {
	o.MaxUsers = &v
}

// GetMaxVirtualMachines returns the MaxVirtualMachines field value if set, zero value otherwise.
func (o *InputTeamFull) GetMaxVirtualMachines() int32 {
	if o == nil || o.MaxVirtualMachines == nil {
		var ret int32
		return ret
	}
	return *o.MaxVirtualMachines
}

// GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetMaxVirtualMachinesOk() (*int32, bool) {
	if o == nil || o.MaxVirtualMachines == nil {
		return nil, false
	}
	return o.MaxVirtualMachines, true
}

// HasMaxVirtualMachines returns a boolean if a field has been set.
func (o *InputTeamFull) HasMaxVirtualMachines() bool {
	if o != nil && o.MaxVirtualMachines != nil {
		return true
	}

	return false
}

// SetMaxVirtualMachines gets a reference to the given int32 and assigns it to the MaxVirtualMachines field.
func (o *InputTeamFull) SetMaxVirtualMachines(v int32) {
	o.MaxVirtualMachines = &v
}

// GetName returns the Name field value
func (o *InputTeamFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputTeamFull) SetName(v string) {
	o.Name = v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *InputTeamFull) GetOrderNumber() string {
	if o == nil || o.OrderNumber == nil {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetOrderNumberOk() (*string, bool) {
	if o == nil || o.OrderNumber == nil {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *InputTeamFull) HasOrderNumber() bool {
	if o != nil && o.OrderNumber != nil {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *InputTeamFull) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetContactInfo returns the ContactInfo field value
func (o *InputTeamFull) GetContactInfo() PocInfo {
	if o == nil {
		var ret PocInfo
		return ret
	}

	return o.ContactInfo
}

// GetContactInfoOk returns a tuple with the ContactInfo field value
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetContactInfoOk() (*PocInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContactInfo, true
}

// SetContactInfo sets field value
func (o *InputTeamFull) SetContactInfo(v PocInfo) {
	o.ContactInfo = v
}

// GetPowerScheduleEnabled returns the PowerScheduleEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetPowerScheduleEnabled() bool {
	if o == nil || o.PowerScheduleEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PowerScheduleEnabled
}

// GetPowerScheduleEnabledOk returns a tuple with the PowerScheduleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetPowerScheduleEnabledOk() (*bool, bool) {
	if o == nil || o.PowerScheduleEnabled == nil {
		return nil, false
	}
	return o.PowerScheduleEnabled, true
}

// HasPowerScheduleEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasPowerScheduleEnabled() bool {
	if o != nil && o.PowerScheduleEnabled != nil {
		return true
	}

	return false
}

// SetPowerScheduleEnabled gets a reference to the given bool and assigns it to the PowerScheduleEnabled field.
func (o *InputTeamFull) SetPowerScheduleEnabled(v bool) {
	o.PowerScheduleEnabled = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *InputTeamFull) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *InputTeamFull) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *InputTeamFull) SetPrivate(v bool) {
	o.Private = &v
}

// GetRdpClientProxyEnabled returns the RdpClientProxyEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetRdpClientProxyEnabled() bool {
	if o == nil || o.RdpClientProxyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RdpClientProxyEnabled
}

// GetRdpClientProxyEnabledOk returns a tuple with the RdpClientProxyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetRdpClientProxyEnabledOk() (*bool, bool) {
	if o == nil || o.RdpClientProxyEnabled == nil {
		return nil, false
	}
	return o.RdpClientProxyEnabled, true
}

// HasRdpClientProxyEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasRdpClientProxyEnabled() bool {
	if o != nil && o.RdpClientProxyEnabled != nil {
		return true
	}

	return false
}

// SetRdpClientProxyEnabled gets a reference to the given bool and assigns it to the RdpClientProxyEnabled field.
func (o *InputTeamFull) SetRdpClientProxyEnabled(v bool) {
	o.RdpClientProxyEnabled = &v
}

// GetRdpClientSessionDuration returns the RdpClientSessionDuration field value if set, zero value otherwise.
func (o *InputTeamFull) GetRdpClientSessionDuration() int32 {
	if o == nil || o.RdpClientSessionDuration == nil {
		var ret int32
		return ret
	}
	return *o.RdpClientSessionDuration
}

// GetRdpClientSessionDurationOk returns a tuple with the RdpClientSessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetRdpClientSessionDurationOk() (*int32, bool) {
	if o == nil || o.RdpClientSessionDuration == nil {
		return nil, false
	}
	return o.RdpClientSessionDuration, true
}

// HasRdpClientSessionDuration returns a boolean if a field has been set.
func (o *InputTeamFull) HasRdpClientSessionDuration() bool {
	if o != nil && o.RdpClientSessionDuration != nil {
		return true
	}

	return false
}

// SetRdpClientSessionDuration gets a reference to the given int32 and assigns it to the RdpClientSessionDuration field.
func (o *InputTeamFull) SetRdpClientSessionDuration(v int32) {
	o.RdpClientSessionDuration = &v
}

// GetSnapshotEnabled returns the SnapshotEnabled field value if set, zero value otherwise.
func (o *InputTeamFull) GetSnapshotEnabled() bool {
	if o == nil || o.SnapshotEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotEnabled
}

// GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetSnapshotEnabledOk() (*bool, bool) {
	if o == nil || o.SnapshotEnabled == nil {
		return nil, false
	}
	return o.SnapshotEnabled, true
}

// HasSnapshotEnabled returns a boolean if a field has been set.
func (o *InputTeamFull) HasSnapshotEnabled() bool {
	if o != nil && o.SnapshotEnabled != nil {
		return true
	}

	return false
}

// SetSnapshotEnabled gets a reference to the given bool and assigns it to the SnapshotEnabled field.
func (o *InputTeamFull) SetSnapshotEnabled(v bool) {
	o.SnapshotEnabled = &v
}

// GetState returns the State field value
func (o *InputTeamFull) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InputTeamFull) SetState(v string) {
	o.State = v
}

// GetValidUntil returns the ValidUntil field value
func (o *InputTeamFull) GetValidUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *InputTeamFull) GetValidUntilOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *InputTeamFull) SetValidUntil(v int32) {
	o.ValidUntil = v
}

func (o InputTeamFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetBundleInstallerEnabled != nil {
		toSerialize["assetBundleInstallerEnabled"] = o.AssetBundleInstallerEnabled
	}
	if o.AssetBypassScanningEnabled != nil {
		toSerialize["assetBypassScanningEnabled"] = o.AssetBypassScanningEnabled
	}
	if o.AvailabilityZoneEnabled != nil {
		toSerialize["availabilityZoneEnabled"] = o.AvailabilityZoneEnabled
	}
	if o.BypassScanningEnabled != nil {
		toSerialize["bypassScanningEnabled"] = o.BypassScanningEnabled
	}
	if true {
		toSerialize["leadUser"] = o.LeadUser
	}
	if o.GpuTypeMaximums != nil {
		toSerialize["gpuTypeMaximums"] = o.GpuTypeMaximums
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["teamManagers"] = o.TeamManagers
	}
	if o.MaxManagedVirtualizationRealms != nil {
		toSerialize["maxManagedVirtualizationRealms"] = o.MaxManagedVirtualizationRealms
	}
	if o.MaxNumCpus != nil {
		toSerialize["maxNumCpus"] = o.MaxNumCpus
	}
	if o.MaxNumGpus != nil {
		toSerialize["maxNumGpus"] = o.MaxNumGpus
	}
	if o.MaxProjects != nil {
		toSerialize["maxProjects"] = o.MaxProjects
	}
	if o.MaxRamInMegabytes != nil {
		toSerialize["maxRamInMegabytes"] = o.MaxRamInMegabytes
	}
	if o.MaxSharedRemoteAccessSessions != nil {
		toSerialize["maxSharedRemoteAccessSessions"] = o.MaxSharedRemoteAccessSessions
	}
	if o.MaxStorageInMegabytes != nil {
		toSerialize["maxStorageInMegabytes"] = o.MaxStorageInMegabytes
	}
	if o.MaxUsers != nil {
		toSerialize["maxUsers"] = o.MaxUsers
	}
	if o.MaxVirtualMachines != nil {
		toSerialize["maxVirtualMachines"] = o.MaxVirtualMachines
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OrderNumber != nil {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if true {
		toSerialize["contactInfo"] = o.ContactInfo
	}
	if o.PowerScheduleEnabled != nil {
		toSerialize["powerScheduleEnabled"] = o.PowerScheduleEnabled
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.RdpClientProxyEnabled != nil {
		toSerialize["rdpClientProxyEnabled"] = o.RdpClientProxyEnabled
	}
	if o.RdpClientSessionDuration != nil {
		toSerialize["rdpClientSessionDuration"] = o.RdpClientSessionDuration
	}
	if o.SnapshotEnabled != nil {
		toSerialize["snapshotEnabled"] = o.SnapshotEnabled
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["validUntil"] = o.ValidUntil
	}
	return json.Marshal(toSerialize)
}

type NullableInputTeamFull struct {
	value *InputTeamFull
	isSet bool
}

func (v NullableInputTeamFull) Get() *InputTeamFull {
	return v.value
}

func (v *NullableInputTeamFull) Set(val *InputTeamFull) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTeamFull) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTeamFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTeamFull(val *InputTeamFull) *NullableInputTeamFull {
	return &NullableInputTeamFull{value: val, isSet: true}
}

func (v NullableInputTeamFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputTeamFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


