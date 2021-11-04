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

// Cloud struct for Cloud
type Cloud struct {
	CloudType *string `json:"cloudType,omitempty"`
	Description string `json:"description"`
	Name string `json:"name"`
	ExternalIpAddresses *[]string `json:"externalIpAddresses,omitempty"`
	ExternalIpSource string `json:"externalIpSource"`
	Features *CloudFeatures `json:"features,omitempty"`
	GpuTypes *[]string `json:"gpuTypes,omitempty"`
	Id *int32 `json:"id,omitempty"`
	LinuxRepositoryUrl *string `json:"linuxRepositoryUrl,omitempty"`
	MaximumImpactLevel string `json:"maximumImpactLevel"`
	Networks *[]Network `json:"networks,omitempty"`
	OwningTeam Team `json:"owningTeam"`
	ProviderFeatureConfigurations *map[string]AbstractProviderClientConfiguration `json:"providerFeatureConfigurations,omitempty"`
	State *string `json:"state,omitempty"`
	TemplateVirtualizationRealm *VirtualizationRealm `json:"templateVirtualizationRealm,omitempty"`
	ConnectedVirtualizationRealms *[]VirtualizationRealm `json:"connectedVirtualizationRealms,omitempty"`
	VirtualizationRealms *[]VirtualizationRealm `json:"virtualizationRealms,omitempty"`
	Subtype string `json:"subtype"`
}

// NewCloud instantiates a new Cloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloud(description string, name string, externalIpSource string, maximumImpactLevel string, owningTeam Team, subtype string) *Cloud {
	this := Cloud{}
	this.Description = description
	this.Name = name
	this.ExternalIpSource = externalIpSource
	this.MaximumImpactLevel = maximumImpactLevel
	this.OwningTeam = owningTeam
	this.Subtype = subtype
	return &this
}

// NewCloudWithDefaults instantiates a new Cloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudWithDefaults() *Cloud {
	this := Cloud{}
	return &this
}

// GetCloudType returns the CloudType field value if set, zero value otherwise.
func (o *Cloud) GetCloudType() string {
	if o == nil || o.CloudType == nil {
		var ret string
		return ret
	}
	return *o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetCloudTypeOk() (*string, bool) {
	if o == nil || o.CloudType == nil {
		return nil, false
	}
	return o.CloudType, true
}

// HasCloudType returns a boolean if a field has been set.
func (o *Cloud) HasCloudType() bool {
	if o != nil && o.CloudType != nil {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given string and assigns it to the CloudType field.
func (o *Cloud) SetCloudType(v string) {
	o.CloudType = &v
}

// GetDescription returns the Description field value
func (o *Cloud) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Cloud) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *Cloud) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cloud) SetName(v string) {
	o.Name = v
}

// GetExternalIpAddresses returns the ExternalIpAddresses field value if set, zero value otherwise.
func (o *Cloud) GetExternalIpAddresses() []string {
	if o == nil || o.ExternalIpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.ExternalIpAddresses
}

// GetExternalIpAddressesOk returns a tuple with the ExternalIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetExternalIpAddressesOk() (*[]string, bool) {
	if o == nil || o.ExternalIpAddresses == nil {
		return nil, false
	}
	return o.ExternalIpAddresses, true
}

// HasExternalIpAddresses returns a boolean if a field has been set.
func (o *Cloud) HasExternalIpAddresses() bool {
	if o != nil && o.ExternalIpAddresses != nil {
		return true
	}

	return false
}

// SetExternalIpAddresses gets a reference to the given []string and assigns it to the ExternalIpAddresses field.
func (o *Cloud) SetExternalIpAddresses(v []string) {
	o.ExternalIpAddresses = &v
}

// GetExternalIpSource returns the ExternalIpSource field value
func (o *Cloud) GetExternalIpSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalIpSource
}

// GetExternalIpSourceOk returns a tuple with the ExternalIpSource field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetExternalIpSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExternalIpSource, true
}

// SetExternalIpSource sets field value
func (o *Cloud) SetExternalIpSource(v string) {
	o.ExternalIpSource = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Cloud) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Cloud) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *Cloud) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

// GetGpuTypes returns the GpuTypes field value if set, zero value otherwise.
func (o *Cloud) GetGpuTypes() []string {
	if o == nil || o.GpuTypes == nil {
		var ret []string
		return ret
	}
	return *o.GpuTypes
}

// GetGpuTypesOk returns a tuple with the GpuTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetGpuTypesOk() (*[]string, bool) {
	if o == nil || o.GpuTypes == nil {
		return nil, false
	}
	return o.GpuTypes, true
}

// HasGpuTypes returns a boolean if a field has been set.
func (o *Cloud) HasGpuTypes() bool {
	if o != nil && o.GpuTypes != nil {
		return true
	}

	return false
}

// SetGpuTypes gets a reference to the given []string and assigns it to the GpuTypes field.
func (o *Cloud) SetGpuTypes(v []string) {
	o.GpuTypes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cloud) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cloud) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Cloud) SetId(v int32) {
	o.Id = &v
}

// GetLinuxRepositoryUrl returns the LinuxRepositoryUrl field value if set, zero value otherwise.
func (o *Cloud) GetLinuxRepositoryUrl() string {
	if o == nil || o.LinuxRepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.LinuxRepositoryUrl
}

// GetLinuxRepositoryUrlOk returns a tuple with the LinuxRepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetLinuxRepositoryUrlOk() (*string, bool) {
	if o == nil || o.LinuxRepositoryUrl == nil {
		return nil, false
	}
	return o.LinuxRepositoryUrl, true
}

// HasLinuxRepositoryUrl returns a boolean if a field has been set.
func (o *Cloud) HasLinuxRepositoryUrl() bool {
	if o != nil && o.LinuxRepositoryUrl != nil {
		return true
	}

	return false
}

// SetLinuxRepositoryUrl gets a reference to the given string and assigns it to the LinuxRepositoryUrl field.
func (o *Cloud) SetLinuxRepositoryUrl(v string) {
	o.LinuxRepositoryUrl = &v
}

// GetMaximumImpactLevel returns the MaximumImpactLevel field value
func (o *Cloud) GetMaximumImpactLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaximumImpactLevel
}

// GetMaximumImpactLevelOk returns a tuple with the MaximumImpactLevel field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetMaximumImpactLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaximumImpactLevel, true
}

// SetMaximumImpactLevel sets field value
func (o *Cloud) SetMaximumImpactLevel(v string) {
	o.MaximumImpactLevel = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *Cloud) GetNetworks() []Network {
	if o == nil || o.Networks == nil {
		var ret []Network
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetNetworksOk() (*[]Network, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *Cloud) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []Network and assigns it to the Networks field.
func (o *Cloud) SetNetworks(v []Network) {
	o.Networks = &v
}

// GetOwningTeam returns the OwningTeam field value
func (o *Cloud) GetOwningTeam() Team {
	if o == nil {
		var ret Team
		return ret
	}

	return o.OwningTeam
}

// GetOwningTeamOk returns a tuple with the OwningTeam field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetOwningTeamOk() (*Team, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OwningTeam, true
}

// SetOwningTeam sets field value
func (o *Cloud) SetOwningTeam(v Team) {
	o.OwningTeam = v
}

// GetProviderFeatureConfigurations returns the ProviderFeatureConfigurations field value if set, zero value otherwise.
func (o *Cloud) GetProviderFeatureConfigurations() map[string]AbstractProviderClientConfiguration {
	if o == nil || o.ProviderFeatureConfigurations == nil {
		var ret map[string]AbstractProviderClientConfiguration
		return ret
	}
	return *o.ProviderFeatureConfigurations
}

// GetProviderFeatureConfigurationsOk returns a tuple with the ProviderFeatureConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetProviderFeatureConfigurationsOk() (*map[string]AbstractProviderClientConfiguration, bool) {
	if o == nil || o.ProviderFeatureConfigurations == nil {
		return nil, false
	}
	return o.ProviderFeatureConfigurations, true
}

// HasProviderFeatureConfigurations returns a boolean if a field has been set.
func (o *Cloud) HasProviderFeatureConfigurations() bool {
	if o != nil && o.ProviderFeatureConfigurations != nil {
		return true
	}

	return false
}

// SetProviderFeatureConfigurations gets a reference to the given map[string]AbstractProviderClientConfiguration and assigns it to the ProviderFeatureConfigurations field.
func (o *Cloud) SetProviderFeatureConfigurations(v map[string]AbstractProviderClientConfiguration) {
	o.ProviderFeatureConfigurations = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Cloud) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Cloud) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Cloud) SetState(v string) {
	o.State = &v
}

// GetTemplateVirtualizationRealm returns the TemplateVirtualizationRealm field value if set, zero value otherwise.
func (o *Cloud) GetTemplateVirtualizationRealm() VirtualizationRealm {
	if o == nil || o.TemplateVirtualizationRealm == nil {
		var ret VirtualizationRealm
		return ret
	}
	return *o.TemplateVirtualizationRealm
}

// GetTemplateVirtualizationRealmOk returns a tuple with the TemplateVirtualizationRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetTemplateVirtualizationRealmOk() (*VirtualizationRealm, bool) {
	if o == nil || o.TemplateVirtualizationRealm == nil {
		return nil, false
	}
	return o.TemplateVirtualizationRealm, true
}

// HasTemplateVirtualizationRealm returns a boolean if a field has been set.
func (o *Cloud) HasTemplateVirtualizationRealm() bool {
	if o != nil && o.TemplateVirtualizationRealm != nil {
		return true
	}

	return false
}

// SetTemplateVirtualizationRealm gets a reference to the given VirtualizationRealm and assigns it to the TemplateVirtualizationRealm field.
func (o *Cloud) SetTemplateVirtualizationRealm(v VirtualizationRealm) {
	o.TemplateVirtualizationRealm = &v
}

// GetConnectedVirtualizationRealms returns the ConnectedVirtualizationRealms field value if set, zero value otherwise.
func (o *Cloud) GetConnectedVirtualizationRealms() []VirtualizationRealm {
	if o == nil || o.ConnectedVirtualizationRealms == nil {
		var ret []VirtualizationRealm
		return ret
	}
	return *o.ConnectedVirtualizationRealms
}

// GetConnectedVirtualizationRealmsOk returns a tuple with the ConnectedVirtualizationRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetConnectedVirtualizationRealmsOk() (*[]VirtualizationRealm, bool) {
	if o == nil || o.ConnectedVirtualizationRealms == nil {
		return nil, false
	}
	return o.ConnectedVirtualizationRealms, true
}

// HasConnectedVirtualizationRealms returns a boolean if a field has been set.
func (o *Cloud) HasConnectedVirtualizationRealms() bool {
	if o != nil && o.ConnectedVirtualizationRealms != nil {
		return true
	}

	return false
}

// SetConnectedVirtualizationRealms gets a reference to the given []VirtualizationRealm and assigns it to the ConnectedVirtualizationRealms field.
func (o *Cloud) SetConnectedVirtualizationRealms(v []VirtualizationRealm) {
	o.ConnectedVirtualizationRealms = &v
}

// GetVirtualizationRealms returns the VirtualizationRealms field value if set, zero value otherwise.
func (o *Cloud) GetVirtualizationRealms() []VirtualizationRealm {
	if o == nil || o.VirtualizationRealms == nil {
		var ret []VirtualizationRealm
		return ret
	}
	return *o.VirtualizationRealms
}

// GetVirtualizationRealmsOk returns a tuple with the VirtualizationRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetVirtualizationRealmsOk() (*[]VirtualizationRealm, bool) {
	if o == nil || o.VirtualizationRealms == nil {
		return nil, false
	}
	return o.VirtualizationRealms, true
}

// HasVirtualizationRealms returns a boolean if a field has been set.
func (o *Cloud) HasVirtualizationRealms() bool {
	if o != nil && o.VirtualizationRealms != nil {
		return true
	}

	return false
}

// SetVirtualizationRealms gets a reference to the given []VirtualizationRealm and assigns it to the VirtualizationRealms field.
func (o *Cloud) SetVirtualizationRealms(v []VirtualizationRealm) {
	o.VirtualizationRealms = &v
}

// GetSubtype returns the Subtype field value
func (o *Cloud) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *Cloud) GetSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *Cloud) SetSubtype(v string) {
	o.Subtype = v
}

func (o Cloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudType != nil {
		toSerialize["cloudType"] = o.CloudType
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExternalIpAddresses != nil {
		toSerialize["externalIpAddresses"] = o.ExternalIpAddresses
	}
	if true {
		toSerialize["externalIpSource"] = o.ExternalIpSource
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.GpuTypes != nil {
		toSerialize["gpuTypes"] = o.GpuTypes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LinuxRepositoryUrl != nil {
		toSerialize["linuxRepositoryUrl"] = o.LinuxRepositoryUrl
	}
	if true {
		toSerialize["maximumImpactLevel"] = o.MaximumImpactLevel
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if true {
		toSerialize["owningTeam"] = o.OwningTeam
	}
	if o.ProviderFeatureConfigurations != nil {
		toSerialize["providerFeatureConfigurations"] = o.ProviderFeatureConfigurations
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.TemplateVirtualizationRealm != nil {
		toSerialize["templateVirtualizationRealm"] = o.TemplateVirtualizationRealm
	}
	if o.ConnectedVirtualizationRealms != nil {
		toSerialize["connectedVirtualizationRealms"] = o.ConnectedVirtualizationRealms
	}
	if o.VirtualizationRealms != nil {
		toSerialize["virtualizationRealms"] = o.VirtualizationRealms
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableCloud struct {
	value *Cloud
	isSet bool
}

func (v NullableCloud) Get() *Cloud {
	return v.value
}

func (v *NullableCloud) Set(val *Cloud) {
	v.value = val
	v.isSet = true
}

func (v NullableCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloud(val *Cloud) *NullableCloud {
	return &NullableCloud{value: val, isSet: true}
}

func (v NullableCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


