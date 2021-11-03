# Cloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Name** | **string** |  | 
**ExternalIpAddresses** | Pointer to **[]string** |  | [optional] 
**ExternalIpSource** | **string** |  | 
**Features** | Pointer to [**CloudFeatures**](CloudFeatures.md) |  | [optional] 
**GpuTypes** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LinuxRepositoryUrl** | Pointer to **string** |  | [optional] 
**MaximumImpactLevel** | **string** |  | 
**Networks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**OwningTeam** | [**Team**](Team.md) |  | 
**ProviderFeatureConfigurations** | Pointer to [**map[string]AbstractProviderClientConfiguration**](AbstractProviderClientConfiguration.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TemplateVirtualizationRealm** | Pointer to [**VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**ConnectedVirtualizationRealms** | Pointer to [**[]VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**VirtualizationRealms** | Pointer to [**[]VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewCloud

`func NewCloud(description string, name string, externalIpSource string, maximumImpactLevel string, owningTeam Team, subtype string, ) *Cloud`

NewCloud instantiates a new Cloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWithDefaults

`func NewCloudWithDefaults() *Cloud`

NewCloudWithDefaults instantiates a new Cloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *Cloud) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *Cloud) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *Cloud) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *Cloud) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetDescription

`func (o *Cloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cloud) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *Cloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cloud) SetName(v string)`

SetName sets Name field to given value.


### GetExternalIpAddresses

`func (o *Cloud) GetExternalIpAddresses() []string`

GetExternalIpAddresses returns the ExternalIpAddresses field if non-nil, zero value otherwise.

### GetExternalIpAddressesOk

`func (o *Cloud) GetExternalIpAddressesOk() (*[]string, bool)`

GetExternalIpAddressesOk returns a tuple with the ExternalIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddresses

`func (o *Cloud) SetExternalIpAddresses(v []string)`

SetExternalIpAddresses sets ExternalIpAddresses field to given value.

### HasExternalIpAddresses

`func (o *Cloud) HasExternalIpAddresses() bool`

HasExternalIpAddresses returns a boolean if a field has been set.

### GetExternalIpSource

`func (o *Cloud) GetExternalIpSource() string`

GetExternalIpSource returns the ExternalIpSource field if non-nil, zero value otherwise.

### GetExternalIpSourceOk

`func (o *Cloud) GetExternalIpSourceOk() (*string, bool)`

GetExternalIpSourceOk returns a tuple with the ExternalIpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpSource

`func (o *Cloud) SetExternalIpSource(v string)`

SetExternalIpSource sets ExternalIpSource field to given value.


### GetFeatures

`func (o *Cloud) GetFeatures() CloudFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Cloud) GetFeaturesOk() (*CloudFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Cloud) SetFeatures(v CloudFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Cloud) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGpuTypes

`func (o *Cloud) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *Cloud) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *Cloud) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *Cloud) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetId

`func (o *Cloud) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cloud) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cloud) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Cloud) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinuxRepositoryUrl

`func (o *Cloud) GetLinuxRepositoryUrl() string`

GetLinuxRepositoryUrl returns the LinuxRepositoryUrl field if non-nil, zero value otherwise.

### GetLinuxRepositoryUrlOk

`func (o *Cloud) GetLinuxRepositoryUrlOk() (*string, bool)`

GetLinuxRepositoryUrlOk returns a tuple with the LinuxRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxRepositoryUrl

`func (o *Cloud) SetLinuxRepositoryUrl(v string)`

SetLinuxRepositoryUrl sets LinuxRepositoryUrl field to given value.

### HasLinuxRepositoryUrl

`func (o *Cloud) HasLinuxRepositoryUrl() bool`

HasLinuxRepositoryUrl returns a boolean if a field has been set.

### GetMaximumImpactLevel

`func (o *Cloud) GetMaximumImpactLevel() string`

GetMaximumImpactLevel returns the MaximumImpactLevel field if non-nil, zero value otherwise.

### GetMaximumImpactLevelOk

`func (o *Cloud) GetMaximumImpactLevelOk() (*string, bool)`

GetMaximumImpactLevelOk returns a tuple with the MaximumImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumImpactLevel

`func (o *Cloud) SetMaximumImpactLevel(v string)`

SetMaximumImpactLevel sets MaximumImpactLevel field to given value.


### GetNetworks

`func (o *Cloud) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Cloud) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Cloud) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *Cloud) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOwningTeam

`func (o *Cloud) GetOwningTeam() Team`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *Cloud) GetOwningTeamOk() (*Team, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *Cloud) SetOwningTeam(v Team)`

SetOwningTeam sets OwningTeam field to given value.


### GetProviderFeatureConfigurations

`func (o *Cloud) GetProviderFeatureConfigurations() map[string]AbstractProviderClientConfiguration`

GetProviderFeatureConfigurations returns the ProviderFeatureConfigurations field if non-nil, zero value otherwise.

### GetProviderFeatureConfigurationsOk

`func (o *Cloud) GetProviderFeatureConfigurationsOk() (*map[string]AbstractProviderClientConfiguration, bool)`

GetProviderFeatureConfigurationsOk returns a tuple with the ProviderFeatureConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderFeatureConfigurations

`func (o *Cloud) SetProviderFeatureConfigurations(v map[string]AbstractProviderClientConfiguration)`

SetProviderFeatureConfigurations sets ProviderFeatureConfigurations field to given value.

### HasProviderFeatureConfigurations

`func (o *Cloud) HasProviderFeatureConfigurations() bool`

HasProviderFeatureConfigurations returns a boolean if a field has been set.

### GetState

`func (o *Cloud) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Cloud) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Cloud) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Cloud) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTemplateVirtualizationRealm

`func (o *Cloud) GetTemplateVirtualizationRealm() VirtualizationRealm`

GetTemplateVirtualizationRealm returns the TemplateVirtualizationRealm field if non-nil, zero value otherwise.

### GetTemplateVirtualizationRealmOk

`func (o *Cloud) GetTemplateVirtualizationRealmOk() (*VirtualizationRealm, bool)`

GetTemplateVirtualizationRealmOk returns a tuple with the TemplateVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVirtualizationRealm

`func (o *Cloud) SetTemplateVirtualizationRealm(v VirtualizationRealm)`

SetTemplateVirtualizationRealm sets TemplateVirtualizationRealm field to given value.

### HasTemplateVirtualizationRealm

`func (o *Cloud) HasTemplateVirtualizationRealm() bool`

HasTemplateVirtualizationRealm returns a boolean if a field has been set.

### GetConnectedVirtualizationRealms

`func (o *Cloud) GetConnectedVirtualizationRealms() []VirtualizationRealm`

GetConnectedVirtualizationRealms returns the ConnectedVirtualizationRealms field if non-nil, zero value otherwise.

### GetConnectedVirtualizationRealmsOk

`func (o *Cloud) GetConnectedVirtualizationRealmsOk() (*[]VirtualizationRealm, bool)`

GetConnectedVirtualizationRealmsOk returns a tuple with the ConnectedVirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVirtualizationRealms

`func (o *Cloud) SetConnectedVirtualizationRealms(v []VirtualizationRealm)`

SetConnectedVirtualizationRealms sets ConnectedVirtualizationRealms field to given value.

### HasConnectedVirtualizationRealms

`func (o *Cloud) HasConnectedVirtualizationRealms() bool`

HasConnectedVirtualizationRealms returns a boolean if a field has been set.

### GetVirtualizationRealms

`func (o *Cloud) GetVirtualizationRealms() []VirtualizationRealm`

GetVirtualizationRealms returns the VirtualizationRealms field if non-nil, zero value otherwise.

### GetVirtualizationRealmsOk

`func (o *Cloud) GetVirtualizationRealmsOk() (*[]VirtualizationRealm, bool)`

GetVirtualizationRealmsOk returns a tuple with the VirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealms

`func (o *Cloud) SetVirtualizationRealms(v []VirtualizationRealm)`

SetVirtualizationRealms sets VirtualizationRealms field to given value.

### HasVirtualizationRealms

`func (o *Cloud) HasVirtualizationRealms() bool`

HasVirtualizationRealms returns a boolean if a field has been set.

### GetSubtype

`func (o *Cloud) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Cloud) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Cloud) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


