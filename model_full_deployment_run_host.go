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

// FullDeploymentRunHost struct for FullDeploymentRunHost
type FullDeploymentRunHost struct {
	BuildOrder                        *int32                  `json:"buildOrder,omitempty"`
	CreatedPassword                   *string                 `json:"createdPassword,omitempty"`
	CreatedUsername                   *string                 `json:"createdUsername,omitempty"`
	DefaultPassword                   *string                 `json:"defaultPassword,omitempty"`
	DefaultUsername                   *string                 `json:"defaultUsername,omitempty"`
	Disks                             *[]Disk                 `json:"disks,omitempty"`
	FapStatus                         *string                 `json:"fapStatus,omitempty"`
	GpuProfile                        *string                 `json:"gpuProfile,omitempty"`
	GpuType                           *string                 `json:"gpuType,omitempty"`
	HasGpu                            *bool                   `json:"hasGpu,omitempty"`
	HostActionInProcess               *bool                   `json:"hostActionInProcess,omitempty"`
	Hostname                          *string                 `json:"hostname,omitempty"`
	Id                                *int32                  `json:"id,omitempty"`
	Installations                     *[]AbstractInstallation `json:"installations,omitempty"`
	InstanceTypeName                  *string                 `json:"instanceTypeName,omitempty"`
	NestedVirtualization              *bool                   `json:"nestedVirtualization,omitempty"`
	NetworkInterfaces                 *[]NetworkInterface     `json:"networkInterfaces,omitempty"`
	NumCpus                           *int32                  `json:"numCpus,omitempty"`
	PhysicalMachineDataOrTemplateUuid *string                 `json:"physicalMachineDataOrTemplateUuid,omitempty"`
	PhysicalMachineOrTemplateName     *string                 `json:"physicalMachineOrTemplateName,omitempty"`
	Published                         *bool                   `json:"published,omitempty"`
	Ram                               *int32                  `json:"ram,omitempty"`
	SnapshotAvailable                 *bool                   `json:"snapshotAvailable,omitempty"`
	SnapshotDate                      *int32                  `json:"snapshotDate,omitempty"`
	SystemModuleId                    *int32                  `json:"systemModuleId,omitempty"`
	SystemModuleType                  *string                 `json:"systemModuleType,omitempty"`
	SystemRole                        *string                 `json:"systemRole,omitempty"`
	VirtualizationRealmStatus         *string                 `json:"VirtualizationRealmStatus,omitempty"`
	Virtual                           *bool                   `json:"virtual,omitempty"`
	Provisionable                     *bool                   `json:"provisionable,omitempty"`
}

// NewFullDeploymentRunHost instantiates a new FullDeploymentRunHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDeploymentRunHost() *FullDeploymentRunHost {
	this := FullDeploymentRunHost{}
	return &this
}

// NewFullDeploymentRunHostWithDefaults instantiates a new FullDeploymentRunHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDeploymentRunHostWithDefaults() *FullDeploymentRunHost {
	this := FullDeploymentRunHost{}
	return &this
}

// GetBuildOrder returns the BuildOrder field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetBuildOrder() int32 {
	if o == nil || o.BuildOrder == nil {
		var ret int32
		return ret
	}
	return *o.BuildOrder
}

// GetBuildOrderOk returns a tuple with the BuildOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetBuildOrderOk() (*int32, bool) {
	if o == nil || o.BuildOrder == nil {
		return nil, false
	}
	return o.BuildOrder, true
}

// HasBuildOrder returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasBuildOrder() bool {
	if o != nil && o.BuildOrder != nil {
		return true
	}

	return false
}

// SetBuildOrder gets a reference to the given int32 and assigns it to the BuildOrder field.
func (o *FullDeploymentRunHost) SetBuildOrder(v int32) {
	o.BuildOrder = &v
}

// GetCreatedPassword returns the CreatedPassword field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetCreatedPassword() string {
	if o == nil || o.CreatedPassword == nil {
		var ret string
		return ret
	}
	return *o.CreatedPassword
}

// GetCreatedPasswordOk returns a tuple with the CreatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetCreatedPasswordOk() (*string, bool) {
	if o == nil || o.CreatedPassword == nil {
		return nil, false
	}
	return o.CreatedPassword, true
}

// HasCreatedPassword returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasCreatedPassword() bool {
	if o != nil && o.CreatedPassword != nil {
		return true
	}

	return false
}

// SetCreatedPassword gets a reference to the given string and assigns it to the CreatedPassword field.
func (o *FullDeploymentRunHost) SetCreatedPassword(v string) {
	o.CreatedPassword = &v
}

// GetCreatedUsername returns the CreatedUsername field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetCreatedUsername() string {
	if o == nil || o.CreatedUsername == nil {
		var ret string
		return ret
	}
	return *o.CreatedUsername
}

// GetCreatedUsernameOk returns a tuple with the CreatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetCreatedUsernameOk() (*string, bool) {
	if o == nil || o.CreatedUsername == nil {
		return nil, false
	}
	return o.CreatedUsername, true
}

// HasCreatedUsername returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasCreatedUsername() bool {
	if o != nil && o.CreatedUsername != nil {
		return true
	}

	return false
}

// SetCreatedUsername gets a reference to the given string and assigns it to the CreatedUsername field.
func (o *FullDeploymentRunHost) SetCreatedUsername(v string) {
	o.CreatedUsername = &v
}

// GetDefaultPassword returns the DefaultPassword field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetDefaultPassword() string {
	if o == nil || o.DefaultPassword == nil {
		var ret string
		return ret
	}
	return *o.DefaultPassword
}

// GetDefaultPasswordOk returns a tuple with the DefaultPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetDefaultPasswordOk() (*string, bool) {
	if o == nil || o.DefaultPassword == nil {
		return nil, false
	}
	return o.DefaultPassword, true
}

// HasDefaultPassword returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasDefaultPassword() bool {
	if o != nil && o.DefaultPassword != nil {
		return true
	}

	return false
}

// SetDefaultPassword gets a reference to the given string and assigns it to the DefaultPassword field.
func (o *FullDeploymentRunHost) SetDefaultPassword(v string) {
	o.DefaultPassword = &v
}

// GetDefaultUsername returns the DefaultUsername field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetDefaultUsername() string {
	if o == nil || o.DefaultUsername == nil {
		var ret string
		return ret
	}
	return *o.DefaultUsername
}

// GetDefaultUsernameOk returns a tuple with the DefaultUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetDefaultUsernameOk() (*string, bool) {
	if o == nil || o.DefaultUsername == nil {
		return nil, false
	}
	return o.DefaultUsername, true
}

// HasDefaultUsername returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasDefaultUsername() bool {
	if o != nil && o.DefaultUsername != nil {
		return true
	}

	return false
}

// SetDefaultUsername gets a reference to the given string and assigns it to the DefaultUsername field.
func (o *FullDeploymentRunHost) SetDefaultUsername(v string) {
	o.DefaultUsername = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetDisks() []Disk {
	if o == nil || o.Disks == nil {
		var ret []Disk
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetDisksOk() (*[]Disk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []Disk and assigns it to the Disks field.
func (o *FullDeploymentRunHost) SetDisks(v []Disk) {
	o.Disks = &v
}

// GetFapStatus returns the FapStatus field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetFapStatus() string {
	if o == nil || o.FapStatus == nil {
		var ret string
		return ret
	}
	return *o.FapStatus
}

// GetFapStatusOk returns a tuple with the FapStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetFapStatusOk() (*string, bool) {
	if o == nil || o.FapStatus == nil {
		return nil, false
	}
	return o.FapStatus, true
}

// HasFapStatus returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasFapStatus() bool {
	if o != nil && o.FapStatus != nil {
		return true
	}

	return false
}

// SetFapStatus gets a reference to the given string and assigns it to the FapStatus field.
func (o *FullDeploymentRunHost) SetFapStatus(v string) {
	o.FapStatus = &v
}

// GetGpuProfile returns the GpuProfile field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetGpuProfile() string {
	if o == nil || o.GpuProfile == nil {
		var ret string
		return ret
	}
	return *o.GpuProfile
}

// GetGpuProfileOk returns a tuple with the GpuProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetGpuProfileOk() (*string, bool) {
	if o == nil || o.GpuProfile == nil {
		return nil, false
	}
	return o.GpuProfile, true
}

// HasGpuProfile returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasGpuProfile() bool {
	if o != nil && o.GpuProfile != nil {
		return true
	}

	return false
}

// SetGpuProfile gets a reference to the given string and assigns it to the GpuProfile field.
func (o *FullDeploymentRunHost) SetGpuProfile(v string) {
	o.GpuProfile = &v
}

// GetGpuType returns the GpuType field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetGpuType() string {
	if o == nil || o.GpuType == nil {
		var ret string
		return ret
	}
	return *o.GpuType
}

// GetGpuTypeOk returns a tuple with the GpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetGpuTypeOk() (*string, bool) {
	if o == nil || o.GpuType == nil {
		return nil, false
	}
	return o.GpuType, true
}

// HasGpuType returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasGpuType() bool {
	if o != nil && o.GpuType != nil {
		return true
	}

	return false
}

// SetGpuType gets a reference to the given string and assigns it to the GpuType field.
func (o *FullDeploymentRunHost) SetGpuType(v string) {
	o.GpuType = &v
}

// GetHasGpu returns the HasGpu field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetHasGpu() bool {
	if o == nil || o.HasGpu == nil {
		var ret bool
		return ret
	}
	return *o.HasGpu
}

// GetHasGpuOk returns a tuple with the HasGpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetHasGpuOk() (*bool, bool) {
	if o == nil || o.HasGpu == nil {
		return nil, false
	}
	return o.HasGpu, true
}

// HasHasGpu returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasHasGpu() bool {
	if o != nil && o.HasGpu != nil {
		return true
	}

	return false
}

// SetHasGpu gets a reference to the given bool and assigns it to the HasGpu field.
func (o *FullDeploymentRunHost) SetHasGpu(v bool) {
	o.HasGpu = &v
}

// GetHostActionInProcess returns the HostActionInProcess field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetHostActionInProcess() bool {
	if o == nil || o.HostActionInProcess == nil {
		var ret bool
		return ret
	}
	return *o.HostActionInProcess
}

// GetHostActionInProcessOk returns a tuple with the HostActionInProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetHostActionInProcessOk() (*bool, bool) {
	if o == nil || o.HostActionInProcess == nil {
		return nil, false
	}
	return o.HostActionInProcess, true
}

// HasHostActionInProcess returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasHostActionInProcess() bool {
	if o != nil && o.HostActionInProcess != nil {
		return true
	}

	return false
}

// SetHostActionInProcess gets a reference to the given bool and assigns it to the HostActionInProcess field.
func (o *FullDeploymentRunHost) SetHostActionInProcess(v bool) {
	o.HostActionInProcess = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FullDeploymentRunHost) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FullDeploymentRunHost) SetId(v int32) {
	o.Id = &v
}

// GetInstallations returns the Installations field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetInstallations() []AbstractInstallation {
	if o == nil || o.Installations == nil {
		var ret []AbstractInstallation
		return ret
	}
	return *o.Installations
}

// GetInstallationsOk returns a tuple with the Installations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetInstallationsOk() (*[]AbstractInstallation, bool) {
	if o == nil || o.Installations == nil {
		return nil, false
	}
	return o.Installations, true
}

// HasInstallations returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasInstallations() bool {
	if o != nil && o.Installations != nil {
		return true
	}

	return false
}

// SetInstallations gets a reference to the given []AbstractInstallation and assigns it to the Installations field.
func (o *FullDeploymentRunHost) SetInstallations(v []AbstractInstallation) {
	o.Installations = &v
}

// GetInstanceTypeName returns the InstanceTypeName field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetInstanceTypeName() string {
	if o == nil || o.InstanceTypeName == nil {
		var ret string
		return ret
	}
	return *o.InstanceTypeName
}

// GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetInstanceTypeNameOk() (*string, bool) {
	if o == nil || o.InstanceTypeName == nil {
		return nil, false
	}
	return o.InstanceTypeName, true
}

// HasInstanceTypeName returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasInstanceTypeName() bool {
	if o != nil && o.InstanceTypeName != nil {
		return true
	}

	return false
}

// SetInstanceTypeName gets a reference to the given string and assigns it to the InstanceTypeName field.
func (o *FullDeploymentRunHost) SetInstanceTypeName(v string) {
	o.InstanceTypeName = &v
}

// GetNestedVirtualization returns the NestedVirtualization field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetNestedVirtualization() bool {
	if o == nil || o.NestedVirtualization == nil {
		var ret bool
		return ret
	}
	return *o.NestedVirtualization
}

// GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetNestedVirtualizationOk() (*bool, bool) {
	if o == nil || o.NestedVirtualization == nil {
		return nil, false
	}
	return o.NestedVirtualization, true
}

// HasNestedVirtualization returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasNestedVirtualization() bool {
	if o != nil && o.NestedVirtualization != nil {
		return true
	}

	return false
}

// SetNestedVirtualization gets a reference to the given bool and assigns it to the NestedVirtualization field.
func (o *FullDeploymentRunHost) SetNestedVirtualization(v bool) {
	o.NestedVirtualization = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetNetworkInterfaces() []NetworkInterface {
	if o == nil || o.NetworkInterfaces == nil {
		var ret []NetworkInterface
		return ret
	}
	return *o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetNetworkInterfacesOk() (*[]NetworkInterface, bool) {
	if o == nil || o.NetworkInterfaces == nil {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasNetworkInterfaces() bool {
	if o != nil && o.NetworkInterfaces != nil {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []NetworkInterface and assigns it to the NetworkInterfaces field.
func (o *FullDeploymentRunHost) SetNetworkInterfaces(v []NetworkInterface) {
	o.NetworkInterfaces = &v
}

// GetNumCpus returns the NumCpus field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetNumCpus() int32 {
	if o == nil || o.NumCpus == nil {
		var ret int32
		return ret
	}
	return *o.NumCpus
}

// GetNumCpusOk returns a tuple with the NumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetNumCpusOk() (*int32, bool) {
	if o == nil || o.NumCpus == nil {
		return nil, false
	}
	return o.NumCpus, true
}

// HasNumCpus returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasNumCpus() bool {
	if o != nil && o.NumCpus != nil {
		return true
	}

	return false
}

// SetNumCpus gets a reference to the given int32 and assigns it to the NumCpus field.
func (o *FullDeploymentRunHost) SetNumCpus(v int32) {
	o.NumCpus = &v
}

// GetPhysicalMachineDataOrTemplateUuid returns the PhysicalMachineDataOrTemplateUuid field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetPhysicalMachineDataOrTemplateUuid() string {
	if o == nil || o.PhysicalMachineDataOrTemplateUuid == nil {
		var ret string
		return ret
	}
	return *o.PhysicalMachineDataOrTemplateUuid
}

// GetPhysicalMachineDataOrTemplateUuidOk returns a tuple with the PhysicalMachineDataOrTemplateUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetPhysicalMachineDataOrTemplateUuidOk() (*string, bool) {
	if o == nil || o.PhysicalMachineDataOrTemplateUuid == nil {
		return nil, false
	}
	return o.PhysicalMachineDataOrTemplateUuid, true
}

// HasPhysicalMachineDataOrTemplateUuid returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasPhysicalMachineDataOrTemplateUuid() bool {
	if o != nil && o.PhysicalMachineDataOrTemplateUuid != nil {
		return true
	}

	return false
}

// SetPhysicalMachineDataOrTemplateUuid gets a reference to the given string and assigns it to the PhysicalMachineDataOrTemplateUuid field.
func (o *FullDeploymentRunHost) SetPhysicalMachineDataOrTemplateUuid(v string) {
	o.PhysicalMachineDataOrTemplateUuid = &v
}

// GetPhysicalMachineOrTemplateName returns the PhysicalMachineOrTemplateName field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetPhysicalMachineOrTemplateName() string {
	if o == nil || o.PhysicalMachineOrTemplateName == nil {
		var ret string
		return ret
	}
	return *o.PhysicalMachineOrTemplateName
}

// GetPhysicalMachineOrTemplateNameOk returns a tuple with the PhysicalMachineOrTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetPhysicalMachineOrTemplateNameOk() (*string, bool) {
	if o == nil || o.PhysicalMachineOrTemplateName == nil {
		return nil, false
	}
	return o.PhysicalMachineOrTemplateName, true
}

// HasPhysicalMachineOrTemplateName returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasPhysicalMachineOrTemplateName() bool {
	if o != nil && o.PhysicalMachineOrTemplateName != nil {
		return true
	}

	return false
}

// SetPhysicalMachineOrTemplateName gets a reference to the given string and assigns it to the PhysicalMachineOrTemplateName field.
func (o *FullDeploymentRunHost) SetPhysicalMachineOrTemplateName(v string) {
	o.PhysicalMachineOrTemplateName = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetPublished() bool {
	if o == nil || o.Published == nil {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetPublishedOk() (*bool, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasPublished() bool {
	if o != nil && o.Published != nil {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *FullDeploymentRunHost) SetPublished(v bool) {
	o.Published = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetRam() int32 {
	if o == nil || o.Ram == nil {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetRamOk() (*int32, bool) {
	if o == nil || o.Ram == nil {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *FullDeploymentRunHost) SetRam(v int32) {
	o.Ram = &v
}

// GetSnapshotAvailable returns the SnapshotAvailable field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetSnapshotAvailable() bool {
	if o == nil || o.SnapshotAvailable == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotAvailable
}

// GetSnapshotAvailableOk returns a tuple with the SnapshotAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetSnapshotAvailableOk() (*bool, bool) {
	if o == nil || o.SnapshotAvailable == nil {
		return nil, false
	}
	return o.SnapshotAvailable, true
}

// HasSnapshotAvailable returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasSnapshotAvailable() bool {
	if o != nil && o.SnapshotAvailable != nil {
		return true
	}

	return false
}

// SetSnapshotAvailable gets a reference to the given bool and assigns it to the SnapshotAvailable field.
func (o *FullDeploymentRunHost) SetSnapshotAvailable(v bool) {
	o.SnapshotAvailable = &v
}

// GetSnapshotDate returns the SnapshotDate field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetSnapshotDate() int32 {
	if o == nil || o.SnapshotDate == nil {
		var ret int32
		return ret
	}
	return *o.SnapshotDate
}

// GetSnapshotDateOk returns a tuple with the SnapshotDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetSnapshotDateOk() (*int32, bool) {
	if o == nil || o.SnapshotDate == nil {
		return nil, false
	}
	return o.SnapshotDate, true
}

// HasSnapshotDate returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasSnapshotDate() bool {
	if o != nil && o.SnapshotDate != nil {
		return true
	}

	return false
}

// SetSnapshotDate gets a reference to the given int32 and assigns it to the SnapshotDate field.
func (o *FullDeploymentRunHost) SetSnapshotDate(v int32) {
	o.SnapshotDate = &v
}

// GetSystemModuleId returns the SystemModuleId field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetSystemModuleId() int32 {
	if o == nil || o.SystemModuleId == nil {
		var ret int32
		return ret
	}
	return *o.SystemModuleId
}

// GetSystemModuleIdOk returns a tuple with the SystemModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetSystemModuleIdOk() (*int32, bool) {
	if o == nil || o.SystemModuleId == nil {
		return nil, false
	}
	return o.SystemModuleId, true
}

// HasSystemModuleId returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasSystemModuleId() bool {
	if o != nil && o.SystemModuleId != nil {
		return true
	}

	return false
}

// SetSystemModuleId gets a reference to the given int32 and assigns it to the SystemModuleId field.
func (o *FullDeploymentRunHost) SetSystemModuleId(v int32) {
	o.SystemModuleId = &v
}

// GetSystemModuleType returns the SystemModuleType field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetSystemModuleType() string {
	if o == nil || o.SystemModuleType == nil {
		var ret string
		return ret
	}
	return *o.SystemModuleType
}

// GetSystemModuleTypeOk returns a tuple with the SystemModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetSystemModuleTypeOk() (*string, bool) {
	if o == nil || o.SystemModuleType == nil {
		return nil, false
	}
	return o.SystemModuleType, true
}

// HasSystemModuleType returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasSystemModuleType() bool {
	if o != nil && o.SystemModuleType != nil {
		return true
	}

	return false
}

// SetSystemModuleType gets a reference to the given string and assigns it to the SystemModuleType field.
func (o *FullDeploymentRunHost) SetSystemModuleType(v string) {
	o.SystemModuleType = &v
}

// GetSystemRole returns the SystemRole field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetSystemRole() string {
	if o == nil || o.SystemRole == nil {
		var ret string
		return ret
	}
	return *o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetSystemRoleOk() (*string, bool) {
	if o == nil || o.SystemRole == nil {
		return nil, false
	}
	return o.SystemRole, true
}

// HasSystemRole returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasSystemRole() bool {
	if o != nil && o.SystemRole != nil {
		return true
	}

	return false
}

// SetSystemRole gets a reference to the given string and assigns it to the SystemRole field.
func (o *FullDeploymentRunHost) SetSystemRole(v string) {
	o.SystemRole = &v
}

// GetVirtualizationRealmStatus returns the VirtualizationRealmStatus field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetVirtualizationRealmStatus() string {
	if o == nil || o.VirtualizationRealmStatus == nil {
		var ret string
		return ret
	}
	return *o.VirtualizationRealmStatus
}

// GetVirtualizationRealmStatusOk returns a tuple with the VirtualizationRealmStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetVirtualizationRealmStatusOk() (*string, bool) {
	if o == nil || o.VirtualizationRealmStatus == nil {
		return nil, false
	}
	return o.VirtualizationRealmStatus, true
}

// HasVirtualizationRealmStatus returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasVirtualizationRealmStatus() bool {
	if o != nil && o.VirtualizationRealmStatus != nil {
		return true
	}

	return false
}

// SetVirtualizationRealmStatus gets a reference to the given string and assigns it to the VirtualizationRealmStatus field.
func (o *FullDeploymentRunHost) SetVirtualizationRealmStatus(v string) {
	o.VirtualizationRealmStatus = &v
}

// GetVirtual returns the Virtual field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetVirtual() bool {
	if o == nil || o.Virtual == nil {
		var ret bool
		return ret
	}
	return *o.Virtual
}

// GetVirtualOk returns a tuple with the Virtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetVirtualOk() (*bool, bool) {
	if o == nil || o.Virtual == nil {
		return nil, false
	}
	return o.Virtual, true
}

// HasVirtual returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasVirtual() bool {
	if o != nil && o.Virtual != nil {
		return true
	}

	return false
}

// SetVirtual gets a reference to the given bool and assigns it to the Virtual field.
func (o *FullDeploymentRunHost) SetVirtual(v bool) {
	o.Virtual = &v
}

// GetProvisionable returns the Provisionable field value if set, zero value otherwise.
func (o *FullDeploymentRunHost) GetProvisionable() bool {
	if o == nil || o.Provisionable == nil {
		var ret bool
		return ret
	}
	return *o.Provisionable
}

// GetProvisionableOk returns a tuple with the Provisionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDeploymentRunHost) GetProvisionableOk() (*bool, bool) {
	if o == nil || o.Provisionable == nil {
		return nil, false
	}
	return o.Provisionable, true
}

// HasProvisionable returns a boolean if a field has been set.
func (o *FullDeploymentRunHost) HasProvisionable() bool {
	if o != nil && o.Provisionable != nil {
		return true
	}

	return false
}

// SetProvisionable gets a reference to the given bool and assigns it to the Provisionable field.
func (o *FullDeploymentRunHost) SetProvisionable(v bool) {
	o.Provisionable = &v
}

func (o FullDeploymentRunHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildOrder != nil {
		toSerialize["buildOrder"] = o.BuildOrder
	}
	if o.CreatedPassword != nil {
		toSerialize["createdPassword"] = o.CreatedPassword
	}
	if o.CreatedUsername != nil {
		toSerialize["createdUsername"] = o.CreatedUsername
	}
	if o.DefaultPassword != nil {
		toSerialize["defaultPassword"] = o.DefaultPassword
	}
	if o.DefaultUsername != nil {
		toSerialize["defaultUsername"] = o.DefaultUsername
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	if o.FapStatus != nil {
		toSerialize["fapStatus"] = o.FapStatus
	}
	if o.GpuProfile != nil {
		toSerialize["gpuProfile"] = o.GpuProfile
	}
	if o.GpuType != nil {
		toSerialize["gpuType"] = o.GpuType
	}
	if o.HasGpu != nil {
		toSerialize["hasGpu"] = o.HasGpu
	}
	if o.HostActionInProcess != nil {
		toSerialize["hostActionInProcess"] = o.HostActionInProcess
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Installations != nil {
		toSerialize["installations"] = o.Installations
	}
	if o.InstanceTypeName != nil {
		toSerialize["instanceTypeName"] = o.InstanceTypeName
	}
	if o.NestedVirtualization != nil {
		toSerialize["nestedVirtualization"] = o.NestedVirtualization
	}
	if o.NetworkInterfaces != nil {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if o.NumCpus != nil {
		toSerialize["numCpus"] = o.NumCpus
	}
	if o.PhysicalMachineDataOrTemplateUuid != nil {
		toSerialize["physicalMachineDataOrTemplateUuid"] = o.PhysicalMachineDataOrTemplateUuid
	}
	if o.PhysicalMachineOrTemplateName != nil {
		toSerialize["physicalMachineOrTemplateName"] = o.PhysicalMachineOrTemplateName
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}
	if o.SnapshotAvailable != nil {
		toSerialize["snapshotAvailable"] = o.SnapshotAvailable
	}
	if o.SnapshotDate != nil {
		toSerialize["snapshotDate"] = o.SnapshotDate
	}
	if o.SystemModuleId != nil {
		toSerialize["systemModuleId"] = o.SystemModuleId
	}
	if o.SystemModuleType != nil {
		toSerialize["systemModuleType"] = o.SystemModuleType
	}
	if o.SystemRole != nil {
		toSerialize["systemRole"] = o.SystemRole
	}
	if o.VirtualizationRealmStatus != nil {
		toSerialize["VirtualizationRealmStatus"] = o.VirtualizationRealmStatus
	}
	if o.Virtual != nil {
		toSerialize["virtual"] = o.Virtual
	}
	if o.Provisionable != nil {
		toSerialize["provisionable"] = o.Provisionable
	}
	return json.Marshal(toSerialize)
}

type NullableFullDeploymentRunHost struct {
	value *FullDeploymentRunHost
	isSet bool
}

func (v NullableFullDeploymentRunHost) Get() *FullDeploymentRunHost {
	return v.value
}

func (v *NullableFullDeploymentRunHost) Set(val *FullDeploymentRunHost) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDeploymentRunHost) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDeploymentRunHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDeploymentRunHost(val *FullDeploymentRunHost) *NullableFullDeploymentRunHost {
	return &NullableFullDeploymentRunHost{value: val, isSet: true}
}

func (v NullableFullDeploymentRunHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDeploymentRunHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
