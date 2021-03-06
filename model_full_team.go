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

// FullTeam struct for FullTeam
type FullTeam struct {
	AssetBundleInstallerEnabled    *bool                         `json:"assetBundleInstallerEnabled,omitempty"`
	AssetBypassScanningEnabled     *bool                         `json:"assetBypassScanningEnabled,omitempty"`
	AvailabilityZoneEnabled        *bool                         `json:"availabilityZoneEnabled,omitempty"`
	BypassScanningEnabled          *bool                         `json:"bypassScanningEnabled,omitempty"`
	LeadUser                       *MinimalUser                  `json:"leadUser,omitempty"`
	GpuTypeMaximums                *map[string]int32             `json:"gpuTypeMaximums,omitempty"`
	Icon                           *string                       `json:"icon,omitempty"`
	Id                             *int32                        `json:"id,omitempty"`
	ManagedVirtualizationRealms    *[]MinimalVirtualizationRealm `json:"managedVirtualizationRealms,omitempty"`
	TeamManagers                   *[]MinimalUser                `json:"teamManagers,omitempty"`
	MaxAssets                      *int32                        `json:"maxAssets,omitempty"`
	MaxManagedVirtualizationRealms *int32                        `json:"maxManagedVirtualizationRealms,omitempty"`
	MaxNumCpus                     *int32                        `json:"maxNumCpus,omitempty"`
	MaxNumGpus                     *int32                        `json:"maxNumGpus,omitempty"`
	MaxProjects                    *int32                        `json:"maxProjects,omitempty"`
	MaxRamInMegabytes              *int32                        `json:"maxRamInMegabytes,omitempty"`
	MaxSharedRemoteAccessSessions  *int32                        `json:"maxSharedRemoteAccessSessions,omitempty"`
	MaxStorageInMegabytes          *int32                        `json:"maxStorageInMegabytes,omitempty"`
	MaxUsers                       *int32                        `json:"maxUsers,omitempty"`
	MaxVirtualMachines             *int32                        `json:"maxVirtualMachines,omitempty"`
	Name                           *string                       `json:"name,omitempty"`
	OrderNumber                    *string                       `json:"orderNumber,omitempty"`
	OwnedClouds                    *[]MinimalCloud               `json:"ownedClouds,omitempty"`
	OwnedProjects                  *[]MinimalProject             `json:"ownedProjects,omitempty"`
	ContactInfo                    PocInfo                       `json:"contactInfo"`
	PowerScheduleEnabled           *bool                         `json:"powerScheduleEnabled,omitempty"`
	Private                        *bool                         `json:"private,omitempty"`
	RdpClientProxyEnabled          *bool                         `json:"rdpClientProxyEnabled,omitempty"`
	RdpClientSessionDuration       *int32                        `json:"rdpClientSessionDuration,omitempty"`
	SnapshotEnabled                *bool                         `json:"snapshotEnabled,omitempty"`
	State                          string                        `json:"state"`
	ValidUtil                      *int32                        `json:"validUtil,omitempty"`
}

// NewFullTeam instantiates a new FullTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullTeam(contactInfo PocInfo, state string) *FullTeam {
	this := FullTeam{}
	this.ContactInfo = contactInfo
	this.State = state
	return &this
}

// NewFullTeamWithDefaults instantiates a new FullTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullTeamWithDefaults() *FullTeam {
	this := FullTeam{}
	return &this
}

// GetAssetBundleInstallerEnabled returns the AssetBundleInstallerEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetAssetBundleInstallerEnabled() bool {
	if o == nil || o.AssetBundleInstallerEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AssetBundleInstallerEnabled
}

// GetAssetBundleInstallerEnabledOk returns a tuple with the AssetBundleInstallerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetAssetBundleInstallerEnabledOk() (*bool, bool) {
	if o == nil || o.AssetBundleInstallerEnabled == nil {
		return nil, false
	}
	return o.AssetBundleInstallerEnabled, true
}

// HasAssetBundleInstallerEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasAssetBundleInstallerEnabled() bool {
	if o != nil && o.AssetBundleInstallerEnabled != nil {
		return true
	}

	return false
}

// SetAssetBundleInstallerEnabled gets a reference to the given bool and assigns it to the AssetBundleInstallerEnabled field.
func (o *FullTeam) SetAssetBundleInstallerEnabled(v bool) {
	o.AssetBundleInstallerEnabled = &v
}

// GetAssetBypassScanningEnabled returns the AssetBypassScanningEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetAssetBypassScanningEnabled() bool {
	if o == nil || o.AssetBypassScanningEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AssetBypassScanningEnabled
}

// GetAssetBypassScanningEnabledOk returns a tuple with the AssetBypassScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetAssetBypassScanningEnabledOk() (*bool, bool) {
	if o == nil || o.AssetBypassScanningEnabled == nil {
		return nil, false
	}
	return o.AssetBypassScanningEnabled, true
}

// HasAssetBypassScanningEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasAssetBypassScanningEnabled() bool {
	if o != nil && o.AssetBypassScanningEnabled != nil {
		return true
	}

	return false
}

// SetAssetBypassScanningEnabled gets a reference to the given bool and assigns it to the AssetBypassScanningEnabled field.
func (o *FullTeam) SetAssetBypassScanningEnabled(v bool) {
	o.AssetBypassScanningEnabled = &v
}

// GetAvailabilityZoneEnabled returns the AvailabilityZoneEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetAvailabilityZoneEnabled() bool {
	if o == nil || o.AvailabilityZoneEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityZoneEnabled
}

// GetAvailabilityZoneEnabledOk returns a tuple with the AvailabilityZoneEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetAvailabilityZoneEnabledOk() (*bool, bool) {
	if o == nil || o.AvailabilityZoneEnabled == nil {
		return nil, false
	}
	return o.AvailabilityZoneEnabled, true
}

// HasAvailabilityZoneEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasAvailabilityZoneEnabled() bool {
	if o != nil && o.AvailabilityZoneEnabled != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneEnabled gets a reference to the given bool and assigns it to the AvailabilityZoneEnabled field.
func (o *FullTeam) SetAvailabilityZoneEnabled(v bool) {
	o.AvailabilityZoneEnabled = &v
}

// GetBypassScanningEnabled returns the BypassScanningEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetBypassScanningEnabled() bool {
	if o == nil || o.BypassScanningEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BypassScanningEnabled
}

// GetBypassScanningEnabledOk returns a tuple with the BypassScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetBypassScanningEnabledOk() (*bool, bool) {
	if o == nil || o.BypassScanningEnabled == nil {
		return nil, false
	}
	return o.BypassScanningEnabled, true
}

// HasBypassScanningEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasBypassScanningEnabled() bool {
	if o != nil && o.BypassScanningEnabled != nil {
		return true
	}

	return false
}

// SetBypassScanningEnabled gets a reference to the given bool and assigns it to the BypassScanningEnabled field.
func (o *FullTeam) SetBypassScanningEnabled(v bool) {
	o.BypassScanningEnabled = &v
}

// GetLeadUser returns the LeadUser field value if set, zero value otherwise.
func (o *FullTeam) GetLeadUser() MinimalUser {
	if o == nil || o.LeadUser == nil {
		var ret MinimalUser
		return ret
	}
	return *o.LeadUser
}

// GetLeadUserOk returns a tuple with the LeadUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetLeadUserOk() (*MinimalUser, bool) {
	if o == nil || o.LeadUser == nil {
		return nil, false
	}
	return o.LeadUser, true
}

// HasLeadUser returns a boolean if a field has been set.
func (o *FullTeam) HasLeadUser() bool {
	if o != nil && o.LeadUser != nil {
		return true
	}

	return false
}

// SetLeadUser gets a reference to the given MinimalUser and assigns it to the LeadUser field.
func (o *FullTeam) SetLeadUser(v MinimalUser) {
	o.LeadUser = &v
}

// GetGpuTypeMaximums returns the GpuTypeMaximums field value if set, zero value otherwise.
func (o *FullTeam) GetGpuTypeMaximums() map[string]int32 {
	if o == nil || o.GpuTypeMaximums == nil {
		var ret map[string]int32
		return ret
	}
	return *o.GpuTypeMaximums
}

// GetGpuTypeMaximumsOk returns a tuple with the GpuTypeMaximums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetGpuTypeMaximumsOk() (*map[string]int32, bool) {
	if o == nil || o.GpuTypeMaximums == nil {
		return nil, false
	}
	return o.GpuTypeMaximums, true
}

// HasGpuTypeMaximums returns a boolean if a field has been set.
func (o *FullTeam) HasGpuTypeMaximums() bool {
	if o != nil && o.GpuTypeMaximums != nil {
		return true
	}

	return false
}

// SetGpuTypeMaximums gets a reference to the given map[string]int32 and assigns it to the GpuTypeMaximums field.
func (o *FullTeam) SetGpuTypeMaximums(v map[string]int32) {
	o.GpuTypeMaximums = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *FullTeam) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *FullTeam) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *FullTeam) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullTeam) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullTeam) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FullTeam) SetId(v int32) {
	o.Id = &v
}

// GetManagedVirtualizationRealms returns the ManagedVirtualizationRealms field value if set, zero value otherwise.
func (o *FullTeam) GetManagedVirtualizationRealms() []MinimalVirtualizationRealm {
	if o == nil || o.ManagedVirtualizationRealms == nil {
		var ret []MinimalVirtualizationRealm
		return ret
	}
	return *o.ManagedVirtualizationRealms
}

// GetManagedVirtualizationRealmsOk returns a tuple with the ManagedVirtualizationRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetManagedVirtualizationRealmsOk() (*[]MinimalVirtualizationRealm, bool) {
	if o == nil || o.ManagedVirtualizationRealms == nil {
		return nil, false
	}
	return o.ManagedVirtualizationRealms, true
}

// HasManagedVirtualizationRealms returns a boolean if a field has been set.
func (o *FullTeam) HasManagedVirtualizationRealms() bool {
	if o != nil && o.ManagedVirtualizationRealms != nil {
		return true
	}

	return false
}

// SetManagedVirtualizationRealms gets a reference to the given []MinimalVirtualizationRealm and assigns it to the ManagedVirtualizationRealms field.
func (o *FullTeam) SetManagedVirtualizationRealms(v []MinimalVirtualizationRealm) {
	o.ManagedVirtualizationRealms = &v
}

// GetTeamManagers returns the TeamManagers field value if set, zero value otherwise.
func (o *FullTeam) GetTeamManagers() []MinimalUser {
	if o == nil || o.TeamManagers == nil {
		var ret []MinimalUser
		return ret
	}
	return *o.TeamManagers
}

// GetTeamManagersOk returns a tuple with the TeamManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetTeamManagersOk() (*[]MinimalUser, bool) {
	if o == nil || o.TeamManagers == nil {
		return nil, false
	}
	return o.TeamManagers, true
}

// HasTeamManagers returns a boolean if a field has been set.
func (o *FullTeam) HasTeamManagers() bool {
	if o != nil && o.TeamManagers != nil {
		return true
	}

	return false
}

// SetTeamManagers gets a reference to the given []MinimalUser and assigns it to the TeamManagers field.
func (o *FullTeam) SetTeamManagers(v []MinimalUser) {
	o.TeamManagers = &v
}

// GetMaxAssets returns the MaxAssets field value if set, zero value otherwise.
func (o *FullTeam) GetMaxAssets() int32 {
	if o == nil || o.MaxAssets == nil {
		var ret int32
		return ret
	}
	return *o.MaxAssets
}

// GetMaxAssetsOk returns a tuple with the MaxAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxAssetsOk() (*int32, bool) {
	if o == nil || o.MaxAssets == nil {
		return nil, false
	}
	return o.MaxAssets, true
}

// HasMaxAssets returns a boolean if a field has been set.
func (o *FullTeam) HasMaxAssets() bool {
	if o != nil && o.MaxAssets != nil {
		return true
	}

	return false
}

// SetMaxAssets gets a reference to the given int32 and assigns it to the MaxAssets field.
func (o *FullTeam) SetMaxAssets(v int32) {
	o.MaxAssets = &v
}

// GetMaxManagedVirtualizationRealms returns the MaxManagedVirtualizationRealms field value if set, zero value otherwise.
func (o *FullTeam) GetMaxManagedVirtualizationRealms() int32 {
	if o == nil || o.MaxManagedVirtualizationRealms == nil {
		var ret int32
		return ret
	}
	return *o.MaxManagedVirtualizationRealms
}

// GetMaxManagedVirtualizationRealmsOk returns a tuple with the MaxManagedVirtualizationRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxManagedVirtualizationRealmsOk() (*int32, bool) {
	if o == nil || o.MaxManagedVirtualizationRealms == nil {
		return nil, false
	}
	return o.MaxManagedVirtualizationRealms, true
}

// HasMaxManagedVirtualizationRealms returns a boolean if a field has been set.
func (o *FullTeam) HasMaxManagedVirtualizationRealms() bool {
	if o != nil && o.MaxManagedVirtualizationRealms != nil {
		return true
	}

	return false
}

// SetMaxManagedVirtualizationRealms gets a reference to the given int32 and assigns it to the MaxManagedVirtualizationRealms field.
func (o *FullTeam) SetMaxManagedVirtualizationRealms(v int32) {
	o.MaxManagedVirtualizationRealms = &v
}

// GetMaxNumCpus returns the MaxNumCpus field value if set, zero value otherwise.
func (o *FullTeam) GetMaxNumCpus() int32 {
	if o == nil || o.MaxNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumCpus
}

// GetMaxNumCpusOk returns a tuple with the MaxNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxNumCpusOk() (*int32, bool) {
	if o == nil || o.MaxNumCpus == nil {
		return nil, false
	}
	return o.MaxNumCpus, true
}

// HasMaxNumCpus returns a boolean if a field has been set.
func (o *FullTeam) HasMaxNumCpus() bool {
	if o != nil && o.MaxNumCpus != nil {
		return true
	}

	return false
}

// SetMaxNumCpus gets a reference to the given int32 and assigns it to the MaxNumCpus field.
func (o *FullTeam) SetMaxNumCpus(v int32) {
	o.MaxNumCpus = &v
}

// GetMaxNumGpus returns the MaxNumGpus field value if set, zero value otherwise.
func (o *FullTeam) GetMaxNumGpus() int32 {
	if o == nil || o.MaxNumGpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumGpus
}

// GetMaxNumGpusOk returns a tuple with the MaxNumGpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxNumGpusOk() (*int32, bool) {
	if o == nil || o.MaxNumGpus == nil {
		return nil, false
	}
	return o.MaxNumGpus, true
}

// HasMaxNumGpus returns a boolean if a field has been set.
func (o *FullTeam) HasMaxNumGpus() bool {
	if o != nil && o.MaxNumGpus != nil {
		return true
	}

	return false
}

// SetMaxNumGpus gets a reference to the given int32 and assigns it to the MaxNumGpus field.
func (o *FullTeam) SetMaxNumGpus(v int32) {
	o.MaxNumGpus = &v
}

// GetMaxProjects returns the MaxProjects field value if set, zero value otherwise.
func (o *FullTeam) GetMaxProjects() int32 {
	if o == nil || o.MaxProjects == nil {
		var ret int32
		return ret
	}
	return *o.MaxProjects
}

// GetMaxProjectsOk returns a tuple with the MaxProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxProjectsOk() (*int32, bool) {
	if o == nil || o.MaxProjects == nil {
		return nil, false
	}
	return o.MaxProjects, true
}

// HasMaxProjects returns a boolean if a field has been set.
func (o *FullTeam) HasMaxProjects() bool {
	if o != nil && o.MaxProjects != nil {
		return true
	}

	return false
}

// SetMaxProjects gets a reference to the given int32 and assigns it to the MaxProjects field.
func (o *FullTeam) SetMaxProjects(v int32) {
	o.MaxProjects = &v
}

// GetMaxRamInMegabytes returns the MaxRamInMegabytes field value if set, zero value otherwise.
func (o *FullTeam) GetMaxRamInMegabytes() int32 {
	if o == nil || o.MaxRamInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRamInMegabytes
}

// GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxRamInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaxRamInMegabytes == nil {
		return nil, false
	}
	return o.MaxRamInMegabytes, true
}

// HasMaxRamInMegabytes returns a boolean if a field has been set.
func (o *FullTeam) HasMaxRamInMegabytes() bool {
	if o != nil && o.MaxRamInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxRamInMegabytes gets a reference to the given int32 and assigns it to the MaxRamInMegabytes field.
func (o *FullTeam) SetMaxRamInMegabytes(v int32) {
	o.MaxRamInMegabytes = &v
}

// GetMaxSharedRemoteAccessSessions returns the MaxSharedRemoteAccessSessions field value if set, zero value otherwise.
func (o *FullTeam) GetMaxSharedRemoteAccessSessions() int32 {
	if o == nil || o.MaxSharedRemoteAccessSessions == nil {
		var ret int32
		return ret
	}
	return *o.MaxSharedRemoteAccessSessions
}

// GetMaxSharedRemoteAccessSessionsOk returns a tuple with the MaxSharedRemoteAccessSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxSharedRemoteAccessSessionsOk() (*int32, bool) {
	if o == nil || o.MaxSharedRemoteAccessSessions == nil {
		return nil, false
	}
	return o.MaxSharedRemoteAccessSessions, true
}

// HasMaxSharedRemoteAccessSessions returns a boolean if a field has been set.
func (o *FullTeam) HasMaxSharedRemoteAccessSessions() bool {
	if o != nil && o.MaxSharedRemoteAccessSessions != nil {
		return true
	}

	return false
}

// SetMaxSharedRemoteAccessSessions gets a reference to the given int32 and assigns it to the MaxSharedRemoteAccessSessions field.
func (o *FullTeam) SetMaxSharedRemoteAccessSessions(v int32) {
	o.MaxSharedRemoteAccessSessions = &v
}

// GetMaxStorageInMegabytes returns the MaxStorageInMegabytes field value if set, zero value otherwise.
func (o *FullTeam) GetMaxStorageInMegabytes() int32 {
	if o == nil || o.MaxStorageInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxStorageInMegabytes
}

// GetMaxStorageInMegabytesOk returns a tuple with the MaxStorageInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxStorageInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaxStorageInMegabytes == nil {
		return nil, false
	}
	return o.MaxStorageInMegabytes, true
}

// HasMaxStorageInMegabytes returns a boolean if a field has been set.
func (o *FullTeam) HasMaxStorageInMegabytes() bool {
	if o != nil && o.MaxStorageInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxStorageInMegabytes gets a reference to the given int32 and assigns it to the MaxStorageInMegabytes field.
func (o *FullTeam) SetMaxStorageInMegabytes(v int32) {
	o.MaxStorageInMegabytes = &v
}

// GetMaxUsers returns the MaxUsers field value if set, zero value otherwise.
func (o *FullTeam) GetMaxUsers() int32 {
	if o == nil || o.MaxUsers == nil {
		var ret int32
		return ret
	}
	return *o.MaxUsers
}

// GetMaxUsersOk returns a tuple with the MaxUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxUsersOk() (*int32, bool) {
	if o == nil || o.MaxUsers == nil {
		return nil, false
	}
	return o.MaxUsers, true
}

// HasMaxUsers returns a boolean if a field has been set.
func (o *FullTeam) HasMaxUsers() bool {
	if o != nil && o.MaxUsers != nil {
		return true
	}

	return false
}

// SetMaxUsers gets a reference to the given int32 and assigns it to the MaxUsers field.
func (o *FullTeam) SetMaxUsers(v int32) {
	o.MaxUsers = &v
}

// GetMaxVirtualMachines returns the MaxVirtualMachines field value if set, zero value otherwise.
func (o *FullTeam) GetMaxVirtualMachines() int32 {
	if o == nil || o.MaxVirtualMachines == nil {
		var ret int32
		return ret
	}
	return *o.MaxVirtualMachines
}

// GetMaxVirtualMachinesOk returns a tuple with the MaxVirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetMaxVirtualMachinesOk() (*int32, bool) {
	if o == nil || o.MaxVirtualMachines == nil {
		return nil, false
	}
	return o.MaxVirtualMachines, true
}

// HasMaxVirtualMachines returns a boolean if a field has been set.
func (o *FullTeam) HasMaxVirtualMachines() bool {
	if o != nil && o.MaxVirtualMachines != nil {
		return true
	}

	return false
}

// SetMaxVirtualMachines gets a reference to the given int32 and assigns it to the MaxVirtualMachines field.
func (o *FullTeam) SetMaxVirtualMachines(v int32) {
	o.MaxVirtualMachines = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullTeam) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullTeam) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullTeam) SetName(v string) {
	o.Name = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *FullTeam) GetOrderNumber() string {
	if o == nil || o.OrderNumber == nil {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetOrderNumberOk() (*string, bool) {
	if o == nil || o.OrderNumber == nil {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *FullTeam) HasOrderNumber() bool {
	if o != nil && o.OrderNumber != nil {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *FullTeam) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetOwnedClouds returns the OwnedClouds field value if set, zero value otherwise.
func (o *FullTeam) GetOwnedClouds() []MinimalCloud {
	if o == nil || o.OwnedClouds == nil {
		var ret []MinimalCloud
		return ret
	}
	return *o.OwnedClouds
}

// GetOwnedCloudsOk returns a tuple with the OwnedClouds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetOwnedCloudsOk() (*[]MinimalCloud, bool) {
	if o == nil || o.OwnedClouds == nil {
		return nil, false
	}
	return o.OwnedClouds, true
}

// HasOwnedClouds returns a boolean if a field has been set.
func (o *FullTeam) HasOwnedClouds() bool {
	if o != nil && o.OwnedClouds != nil {
		return true
	}

	return false
}

// SetOwnedClouds gets a reference to the given []MinimalCloud and assigns it to the OwnedClouds field.
func (o *FullTeam) SetOwnedClouds(v []MinimalCloud) {
	o.OwnedClouds = &v
}

// GetOwnedProjects returns the OwnedProjects field value if set, zero value otherwise.
func (o *FullTeam) GetOwnedProjects() []MinimalProject {
	if o == nil || o.OwnedProjects == nil {
		var ret []MinimalProject
		return ret
	}
	return *o.OwnedProjects
}

// GetOwnedProjectsOk returns a tuple with the OwnedProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetOwnedProjectsOk() (*[]MinimalProject, bool) {
	if o == nil || o.OwnedProjects == nil {
		return nil, false
	}
	return o.OwnedProjects, true
}

// HasOwnedProjects returns a boolean if a field has been set.
func (o *FullTeam) HasOwnedProjects() bool {
	if o != nil && o.OwnedProjects != nil {
		return true
	}

	return false
}

// SetOwnedProjects gets a reference to the given []MinimalProject and assigns it to the OwnedProjects field.
func (o *FullTeam) SetOwnedProjects(v []MinimalProject) {
	o.OwnedProjects = &v
}

// GetContactInfo returns the ContactInfo field value
func (o *FullTeam) GetContactInfo() PocInfo {
	if o == nil {
		var ret PocInfo
		return ret
	}

	return o.ContactInfo
}

// GetContactInfoOk returns a tuple with the ContactInfo field value
// and a boolean to check if the value has been set.
func (o *FullTeam) GetContactInfoOk() (*PocInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactInfo, true
}

// SetContactInfo sets field value
func (o *FullTeam) SetContactInfo(v PocInfo) {
	o.ContactInfo = v
}

// GetPowerScheduleEnabled returns the PowerScheduleEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetPowerScheduleEnabled() bool {
	if o == nil || o.PowerScheduleEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PowerScheduleEnabled
}

// GetPowerScheduleEnabledOk returns a tuple with the PowerScheduleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetPowerScheduleEnabledOk() (*bool, bool) {
	if o == nil || o.PowerScheduleEnabled == nil {
		return nil, false
	}
	return o.PowerScheduleEnabled, true
}

// HasPowerScheduleEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasPowerScheduleEnabled() bool {
	if o != nil && o.PowerScheduleEnabled != nil {
		return true
	}

	return false
}

// SetPowerScheduleEnabled gets a reference to the given bool and assigns it to the PowerScheduleEnabled field.
func (o *FullTeam) SetPowerScheduleEnabled(v bool) {
	o.PowerScheduleEnabled = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *FullTeam) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *FullTeam) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *FullTeam) SetPrivate(v bool) {
	o.Private = &v
}

// GetRdpClientProxyEnabled returns the RdpClientProxyEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetRdpClientProxyEnabled() bool {
	if o == nil || o.RdpClientProxyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RdpClientProxyEnabled
}

// GetRdpClientProxyEnabledOk returns a tuple with the RdpClientProxyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetRdpClientProxyEnabledOk() (*bool, bool) {
	if o == nil || o.RdpClientProxyEnabled == nil {
		return nil, false
	}
	return o.RdpClientProxyEnabled, true
}

// HasRdpClientProxyEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasRdpClientProxyEnabled() bool {
	if o != nil && o.RdpClientProxyEnabled != nil {
		return true
	}

	return false
}

// SetRdpClientProxyEnabled gets a reference to the given bool and assigns it to the RdpClientProxyEnabled field.
func (o *FullTeam) SetRdpClientProxyEnabled(v bool) {
	o.RdpClientProxyEnabled = &v
}

// GetRdpClientSessionDuration returns the RdpClientSessionDuration field value if set, zero value otherwise.
func (o *FullTeam) GetRdpClientSessionDuration() int32 {
	if o == nil || o.RdpClientSessionDuration == nil {
		var ret int32
		return ret
	}
	return *o.RdpClientSessionDuration
}

// GetRdpClientSessionDurationOk returns a tuple with the RdpClientSessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetRdpClientSessionDurationOk() (*int32, bool) {
	if o == nil || o.RdpClientSessionDuration == nil {
		return nil, false
	}
	return o.RdpClientSessionDuration, true
}

// HasRdpClientSessionDuration returns a boolean if a field has been set.
func (o *FullTeam) HasRdpClientSessionDuration() bool {
	if o != nil && o.RdpClientSessionDuration != nil {
		return true
	}

	return false
}

// SetRdpClientSessionDuration gets a reference to the given int32 and assigns it to the RdpClientSessionDuration field.
func (o *FullTeam) SetRdpClientSessionDuration(v int32) {
	o.RdpClientSessionDuration = &v
}

// GetSnapshotEnabled returns the SnapshotEnabled field value if set, zero value otherwise.
func (o *FullTeam) GetSnapshotEnabled() bool {
	if o == nil || o.SnapshotEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotEnabled
}

// GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetSnapshotEnabledOk() (*bool, bool) {
	if o == nil || o.SnapshotEnabled == nil {
		return nil, false
	}
	return o.SnapshotEnabled, true
}

// HasSnapshotEnabled returns a boolean if a field has been set.
func (o *FullTeam) HasSnapshotEnabled() bool {
	if o != nil && o.SnapshotEnabled != nil {
		return true
	}

	return false
}

// SetSnapshotEnabled gets a reference to the given bool and assigns it to the SnapshotEnabled field.
func (o *FullTeam) SetSnapshotEnabled(v bool) {
	o.SnapshotEnabled = &v
}

// GetState returns the State field value
func (o *FullTeam) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FullTeam) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *FullTeam) SetState(v string) {
	o.State = v
}

// GetValidUtil returns the ValidUtil field value if set, zero value otherwise.
func (o *FullTeam) GetValidUtil() int32 {
	if o == nil || o.ValidUtil == nil {
		var ret int32
		return ret
	}
	return *o.ValidUtil
}

// GetValidUtilOk returns a tuple with the ValidUtil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTeam) GetValidUtilOk() (*int32, bool) {
	if o == nil || o.ValidUtil == nil {
		return nil, false
	}
	return o.ValidUtil, true
}

// HasValidUtil returns a boolean if a field has been set.
func (o *FullTeam) HasValidUtil() bool {
	if o != nil && o.ValidUtil != nil {
		return true
	}

	return false
}

// SetValidUtil gets a reference to the given int32 and assigns it to the ValidUtil field.
func (o *FullTeam) SetValidUtil(v int32) {
	o.ValidUtil = &v
}

func (o FullTeam) MarshalJSON() ([]byte, error) {
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
	if o.LeadUser != nil {
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
	if o.ManagedVirtualizationRealms != nil {
		toSerialize["managedVirtualizationRealms"] = o.ManagedVirtualizationRealms
	}
	if o.TeamManagers != nil {
		toSerialize["teamManagers"] = o.TeamManagers
	}
	if o.MaxAssets != nil {
		toSerialize["maxAssets"] = o.MaxAssets
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrderNumber != nil {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if o.OwnedClouds != nil {
		toSerialize["ownedClouds"] = o.OwnedClouds
	}
	if o.OwnedProjects != nil {
		toSerialize["ownedProjects"] = o.OwnedProjects
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
	if o.ValidUtil != nil {
		toSerialize["validUtil"] = o.ValidUtil
	}
	return json.Marshal(toSerialize)
}

type NullableFullTeam struct {
	value *FullTeam
	isSet bool
}

func (v NullableFullTeam) Get() *FullTeam {
	return v.value
}

func (v *NullableFullTeam) Set(val *FullTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableFullTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableFullTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullTeam(val *FullTeam) *NullableFullTeam {
	return &NullableFullTeam{value: val, isSet: true}
}

func (v NullableFullTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
