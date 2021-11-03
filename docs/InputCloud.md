# InputCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExternalIpAddresses** | Pointer to **[]string** |  | [optional] 
**AdditionalNetworks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**Cons3rtNetwork** | Pointer to [**Network**](Network.md) |  | [optional] 
**ExternalIpSource** | **string** |  | 
**Features** | Pointer to [**CloudFeatures**](CloudFeatures.md) |  | [optional] 
**GpuTypes** | Pointer to **[]string** |  | [optional] 
**LinuxRepositoryUrl** | Pointer to **string** |  | [optional] 
**MaximumImpactLevel** | **string** |  | 
**OwningTeam** | [**InputTeam**](InputTeam.md) |  | 
**State** | Pointer to **string** |  | [optional] 
**TemplateVirtualizationRealm** | Pointer to [**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewInputCloud

`func NewInputCloud(name string, externalIpSource string, maximumImpactLevel string, owningTeam InputTeam, subtype string, ) *InputCloud`

NewInputCloud instantiates a new InputCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCloudWithDefaults

`func NewInputCloudWithDefaults() *InputCloud`

NewInputCloudWithDefaults instantiates a new InputCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCloud) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InputCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalIpAddresses

`func (o *InputCloud) GetExternalIpAddresses() []string`

GetExternalIpAddresses returns the ExternalIpAddresses field if non-nil, zero value otherwise.

### GetExternalIpAddressesOk

`func (o *InputCloud) GetExternalIpAddressesOk() (*[]string, bool)`

GetExternalIpAddressesOk returns a tuple with the ExternalIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddresses

`func (o *InputCloud) SetExternalIpAddresses(v []string)`

SetExternalIpAddresses sets ExternalIpAddresses field to given value.

### HasExternalIpAddresses

`func (o *InputCloud) HasExternalIpAddresses() bool`

HasExternalIpAddresses returns a boolean if a field has been set.

### GetAdditionalNetworks

`func (o *InputCloud) GetAdditionalNetworks() []Network`

GetAdditionalNetworks returns the AdditionalNetworks field if non-nil, zero value otherwise.

### GetAdditionalNetworksOk

`func (o *InputCloud) GetAdditionalNetworksOk() (*[]Network, bool)`

GetAdditionalNetworksOk returns a tuple with the AdditionalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNetworks

`func (o *InputCloud) SetAdditionalNetworks(v []Network)`

SetAdditionalNetworks sets AdditionalNetworks field to given value.

### HasAdditionalNetworks

`func (o *InputCloud) HasAdditionalNetworks() bool`

HasAdditionalNetworks returns a boolean if a field has been set.

### GetCons3rtNetwork

`func (o *InputCloud) GetCons3rtNetwork() Network`

GetCons3rtNetwork returns the Cons3rtNetwork field if non-nil, zero value otherwise.

### GetCons3rtNetworkOk

`func (o *InputCloud) GetCons3rtNetworkOk() (*Network, bool)`

GetCons3rtNetworkOk returns a tuple with the Cons3rtNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCons3rtNetwork

`func (o *InputCloud) SetCons3rtNetwork(v Network)`

SetCons3rtNetwork sets Cons3rtNetwork field to given value.

### HasCons3rtNetwork

`func (o *InputCloud) HasCons3rtNetwork() bool`

HasCons3rtNetwork returns a boolean if a field has been set.

### GetExternalIpSource

`func (o *InputCloud) GetExternalIpSource() string`

GetExternalIpSource returns the ExternalIpSource field if non-nil, zero value otherwise.

### GetExternalIpSourceOk

`func (o *InputCloud) GetExternalIpSourceOk() (*string, bool)`

GetExternalIpSourceOk returns a tuple with the ExternalIpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpSource

`func (o *InputCloud) SetExternalIpSource(v string)`

SetExternalIpSource sets ExternalIpSource field to given value.


### GetFeatures

`func (o *InputCloud) GetFeatures() CloudFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InputCloud) GetFeaturesOk() (*CloudFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InputCloud) SetFeatures(v CloudFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InputCloud) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGpuTypes

`func (o *InputCloud) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *InputCloud) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *InputCloud) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *InputCloud) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetLinuxRepositoryUrl

`func (o *InputCloud) GetLinuxRepositoryUrl() string`

GetLinuxRepositoryUrl returns the LinuxRepositoryUrl field if non-nil, zero value otherwise.

### GetLinuxRepositoryUrlOk

`func (o *InputCloud) GetLinuxRepositoryUrlOk() (*string, bool)`

GetLinuxRepositoryUrlOk returns a tuple with the LinuxRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxRepositoryUrl

`func (o *InputCloud) SetLinuxRepositoryUrl(v string)`

SetLinuxRepositoryUrl sets LinuxRepositoryUrl field to given value.

### HasLinuxRepositoryUrl

`func (o *InputCloud) HasLinuxRepositoryUrl() bool`

HasLinuxRepositoryUrl returns a boolean if a field has been set.

### GetMaximumImpactLevel

`func (o *InputCloud) GetMaximumImpactLevel() string`

GetMaximumImpactLevel returns the MaximumImpactLevel field if non-nil, zero value otherwise.

### GetMaximumImpactLevelOk

`func (o *InputCloud) GetMaximumImpactLevelOk() (*string, bool)`

GetMaximumImpactLevelOk returns a tuple with the MaximumImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumImpactLevel

`func (o *InputCloud) SetMaximumImpactLevel(v string)`

SetMaximumImpactLevel sets MaximumImpactLevel field to given value.


### GetOwningTeam

`func (o *InputCloud) GetOwningTeam() InputTeam`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *InputCloud) GetOwningTeamOk() (*InputTeam, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *InputCloud) SetOwningTeam(v InputTeam)`

SetOwningTeam sets OwningTeam field to given value.


### GetState

`func (o *InputCloud) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InputCloud) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InputCloud) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InputCloud) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTemplateVirtualizationRealm

`func (o *InputCloud) GetTemplateVirtualizationRealm() MinimalVirtualizationRealm`

GetTemplateVirtualizationRealm returns the TemplateVirtualizationRealm field if non-nil, zero value otherwise.

### GetTemplateVirtualizationRealmOk

`func (o *InputCloud) GetTemplateVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool)`

GetTemplateVirtualizationRealmOk returns a tuple with the TemplateVirtualizationRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVirtualizationRealm

`func (o *InputCloud) SetTemplateVirtualizationRealm(v MinimalVirtualizationRealm)`

SetTemplateVirtualizationRealm sets TemplateVirtualizationRealm field to given value.

### HasTemplateVirtualizationRealm

`func (o *InputCloud) HasTemplateVirtualizationRealm() bool`

HasTemplateVirtualizationRealm returns a boolean if a field has been set.

### GetSubtype

`func (o *InputCloud) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputCloud) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputCloud) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


