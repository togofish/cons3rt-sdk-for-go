# FullCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalIpAddresses** | Pointer to **[]string** |  | [optional] 
**ExternalIpSource** | **string** |  | 
**Features** | Pointer to [**CloudFeatures**](CloudFeatures.md) |  | [optional] 
**GpuTypes** | Pointer to **[]string** |  | [optional] 
**LinuxRepositoryUrl** | Pointer to **string** |  | [optional] 
**MaximumImpactLevel** | **string** |  | 
**Networks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**OwningTeam** | Pointer to [**MinimalTeam**](MinimalTeam.md) |  | [optional] 
**TemplateVirtualizationRealm** | Pointer to [**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**VirtualizationRealms** | Pointer to [**[]MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewFullCloud

`func NewFullCloud(name string, externalIpSource string, maximumImpactLevel string, subtype string, ) *FullCloud`

NewFullCloud instantiates a new FullCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullCloudWithDefaults

`func NewFullCloudWithDefaults() *FullCloud`

NewFullCloudWithDefaults instantiates a new FullCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *FullCloud) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *FullCloud) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *FullCloud) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *FullCloud) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetId

`func (o *FullCloud) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullCloud) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullCloud) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullCloud) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullCloud) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *FullCloud) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FullCloud) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FullCloud) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FullCloud) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDescription

`func (o *FullCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalIpAddresses

`func (o *FullCloud) GetExternalIpAddresses() []string`

GetExternalIpAddresses returns the ExternalIpAddresses field if non-nil, zero value otherwise.

### GetExternalIpAddressesOk

`func (o *FullCloud) GetExternalIpAddressesOk() (*[]string, bool)`

GetExternalIpAddressesOk returns a tuple with the ExternalIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddresses

`func (o *FullCloud) SetExternalIpAddresses(v []string)`

SetExternalIpAddresses sets ExternalIpAddresses field to given value.

### HasExternalIpAddresses

`func (o *FullCloud) HasExternalIpAddresses() bool`

HasExternalIpAddresses returns a boolean if a field has been set.

### GetExternalIpSource

`func (o *FullCloud) GetExternalIpSource() string`

GetExternalIpSource returns the ExternalIpSource field if non-nil, zero value otherwise.

### GetExternalIpSourceOk

`func (o *FullCloud) GetExternalIpSourceOk() (*string, bool)`

GetExternalIpSourceOk returns a tuple with the ExternalIpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpSource

`func (o *FullCloud) SetExternalIpSource(v string)`

SetExternalIpSource sets ExternalIpSource field to given value.


### GetFeatures

`func (o *FullCloud) GetFeatures() CloudFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FullCloud) GetFeaturesOk() (*CloudFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FullCloud) SetFeatures(v CloudFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FullCloud) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGpuTypes

`func (o *FullCloud) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *FullCloud) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *FullCloud) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *FullCloud) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetLinuxRepositoryUrl

`func (o *FullCloud) GetLinuxRepositoryUrl() string`

GetLinuxRepositoryUrl returns the LinuxRepositoryUrl field if non-nil, zero value otherwise.

### GetLinuxRepositoryUrlOk

`func (o *FullCloud) GetLinuxRepositoryUrlOk() (*string, bool)`

GetLinuxRepositoryUrlOk returns a tuple with the LinuxRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxRepositoryUrl

`func (o *FullCloud) SetLinuxRepositoryUrl(v string)`

SetLinuxRepositoryUrl sets LinuxRepositoryUrl field to given value.

### HasLinuxRepositoryUrl

`func (o *FullCloud) HasLinuxRepositoryUrl() bool`

HasLinuxRepositoryUrl returns a boolean if a field has been set.

### GetMaximumImpactLevel

`func (o *FullCloud) GetMaximumImpactLevel() string`

GetMaximumImpactLevel returns the MaximumImpactLevel field if non-nil, zero value otherwise.

### GetMaximumImpactLevelOk

`func (o *FullCloud) GetMaximumImpactLevelOk() (*string, bool)`

GetMaximumImpactLevelOk returns a tuple with the MaximumImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumImpactLevel

`func (o *FullCloud) SetMaximumImpactLevel(v string)`

SetMaximumImpactLevel sets MaximumImpactLevel field to given value.


### GetNetworks

`func (o *FullCloud) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *FullCloud) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *FullCloud) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *FullCloud) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOwningTeam

`func (o *FullCloud) GetOwningTeam() MinimalTeam`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *FullCloud) GetOwningTeamOk() (*MinimalTeam, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *FullCloud) SetOwningTeam(v MinimalTeam)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *FullCloud) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetTemplateVirtualizationRealm

`func (o *FullCloud) GetTemplateVirtualizationRealm() MinimalVirtualizationRealm`

GetTemplateVirtualizationRealm returns the TemplateVirtualizationRealm field if non-nil, zero value otherwise.

### GetTemplateVirtualizationRealmOk

`func (o *FullCloud) GetTemplateVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool)`

GetTemplateVirtualizationRealmOk returns a tuple with the TemplateVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVirtualizationRealm

`func (o *FullCloud) SetTemplateVirtualizationRealm(v MinimalVirtualizationRealm)`

SetTemplateVirtualizationRealm sets TemplateVirtualizationRealm field to given value.

### HasTemplateVirtualizationRealm

`func (o *FullCloud) HasTemplateVirtualizationRealm() bool`

HasTemplateVirtualizationRealm returns a boolean if a field has been set.

### GetVirtualizationRealms

`func (o *FullCloud) GetVirtualizationRealms() []MinimalVirtualizationRealm`

GetVirtualizationRealms returns the VirtualizationRealms field if non-nil, zero value otherwise.

### GetVirtualizationRealmsOk

`func (o *FullCloud) GetVirtualizationRealmsOk() (*[]MinimalVirtualizationRealm, bool)`

GetVirtualizationRealmsOk returns a tuple with the VirtualizationRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealms

`func (o *FullCloud) SetVirtualizationRealms(v []MinimalVirtualizationRealm)`

SetVirtualizationRealms sets VirtualizationRealms field to given value.

### HasVirtualizationRealms

`func (o *FullCloud) HasVirtualizationRealms() bool`

HasVirtualizationRealms returns a boolean if a field has been set.

### GetSubtype

`func (o *FullCloud) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FullCloud) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FullCloud) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


