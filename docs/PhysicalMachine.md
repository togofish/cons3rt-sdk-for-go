# PhysicalMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TrustedProjects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] 
**DependentAssetCount** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**OwningProject** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ImpactLevel** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**Bits** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**BaseDisks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**RemoteAccessTemplates** | Pointer to [**[]RemoteAccessTemplate**](RemoteAccessTemplate.md) |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**CpuCount** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **int32** |  | [optional] 

## Methods

### NewPhysicalMachine

`func NewPhysicalMachine() *PhysicalMachine`

NewPhysicalMachine instantiates a new PhysicalMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalMachineWithDefaults

`func NewPhysicalMachineWithDefaults() *PhysicalMachine`

NewPhysicalMachineWithDefaults instantiates a new PhysicalMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhysicalMachine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalMachine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalMachine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhysicalMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedProjects

`func (o *PhysicalMachine) GetTrustedProjects() []Project`

GetTrustedProjects returns the TrustedProjects field if non-nil, zero value otherwise.

### GetTrustedProjectsOk

`func (o *PhysicalMachine) GetTrustedProjectsOk() (*[]Project, bool)`

GetTrustedProjectsOk returns a tuple with the TrustedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProjects

`func (o *PhysicalMachine) SetTrustedProjects(v []Project)`

SetTrustedProjects sets TrustedProjects field to given value.

### HasTrustedProjects

`func (o *PhysicalMachine) HasTrustedProjects() bool`

HasTrustedProjects returns a boolean if a field has been set.

### GetCreator

`func (o *PhysicalMachine) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *PhysicalMachine) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *PhysicalMachine) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *PhysicalMachine) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDependentAssetCount

`func (o *PhysicalMachine) GetDependentAssetCount() int32`

GetDependentAssetCount returns the DependentAssetCount field if non-nil, zero value otherwise.

### GetDependentAssetCountOk

`func (o *PhysicalMachine) GetDependentAssetCountOk() (*int32, bool)`

GetDependentAssetCountOk returns a tuple with the DependentAssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentAssetCount

`func (o *PhysicalMachine) SetDependentAssetCount(v int32)`

SetDependentAssetCount sets DependentAssetCount field to given value.

### HasDependentAssetCount

`func (o *PhysicalMachine) HasDependentAssetCount() bool`

HasDependentAssetCount returns a boolean if a field has been set.

### GetDescription

`func (o *PhysicalMachine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhysicalMachine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhysicalMachine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhysicalMachine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *PhysicalMachine) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PhysicalMachine) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PhysicalMachine) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PhysicalMachine) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *PhysicalMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhysicalMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhysicalMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhysicalMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffline

`func (o *PhysicalMachine) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *PhysicalMachine) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *PhysicalMachine) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *PhysicalMachine) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOwningProject

`func (o *PhysicalMachine) GetOwningProject() Project`

GetOwningProject returns the OwningProject field if non-nil, zero value otherwise.

### GetOwningProjectOk

`func (o *PhysicalMachine) GetOwningProjectOk() (*Project, bool)`

GetOwningProjectOk returns a tuple with the OwningProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningProject

`func (o *PhysicalMachine) SetOwningProject(v Project)`

SetOwningProject sets OwningProject field to given value.

### HasOwningProject

`func (o *PhysicalMachine) HasOwningProject() bool`

HasOwningProject returns a boolean if a field has been set.

### GetState

`func (o *PhysicalMachine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PhysicalMachine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PhysicalMachine) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PhysicalMachine) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVisibility

`func (o *PhysicalMachine) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PhysicalMachine) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PhysicalMachine) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PhysicalMachine) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetImpactLevel

`func (o *PhysicalMachine) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *PhysicalMachine) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *PhysicalMachine) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *PhysicalMachine) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetCategories

`func (o *PhysicalMachine) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PhysicalMachine) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PhysicalMachine) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PhysicalMachine) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetArchitecture

`func (o *PhysicalMachine) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *PhysicalMachine) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *PhysicalMachine) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *PhysicalMachine) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBits

`func (o *PhysicalMachine) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *PhysicalMachine) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *PhysicalMachine) SetBits(v string)`

SetBits sets Bits field to given value.

### HasBits

`func (o *PhysicalMachine) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *PhysicalMachine) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *PhysicalMachine) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *PhysicalMachine) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *PhysicalMachine) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetBaseDisks

`func (o *PhysicalMachine) GetBaseDisks() []Disk`

GetBaseDisks returns the BaseDisks field if non-nil, zero value otherwise.

### GetBaseDisksOk

`func (o *PhysicalMachine) GetBaseDisksOk() (*[]Disk, bool)`

GetBaseDisksOk returns a tuple with the BaseDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDisks

`func (o *PhysicalMachine) SetBaseDisks(v []Disk)`

SetBaseDisks sets BaseDisks field to given value.

### HasBaseDisks

`func (o *PhysicalMachine) HasBaseDisks() bool`

HasBaseDisks returns a boolean if a field has been set.

### GetRemoteAccessTemplates

`func (o *PhysicalMachine) GetRemoteAccessTemplates() []RemoteAccessTemplate`

GetRemoteAccessTemplates returns the RemoteAccessTemplates field if non-nil, zero value otherwise.

### GetRemoteAccessTemplatesOk

`func (o *PhysicalMachine) GetRemoteAccessTemplatesOk() (*[]RemoteAccessTemplate, bool)`

GetRemoteAccessTemplatesOk returns a tuple with the RemoteAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessTemplates

`func (o *PhysicalMachine) SetRemoteAccessTemplates(v []RemoteAccessTemplate)`

SetRemoteAccessTemplates sets RemoteAccessTemplates field to given value.

### HasRemoteAccessTemplates

`func (o *PhysicalMachine) HasRemoteAccessTemplates() bool`

HasRemoteAccessTemplates returns a boolean if a field has been set.

### GetHostName

`func (o *PhysicalMachine) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *PhysicalMachine) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *PhysicalMachine) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *PhysicalMachine) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpAddress

`func (o *PhysicalMachine) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PhysicalMachine) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PhysicalMachine) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PhysicalMachine) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *PhysicalMachine) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PhysicalMachine) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PhysicalMachine) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PhysicalMachine) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetCpuCount

`func (o *PhysicalMachine) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *PhysicalMachine) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *PhysicalMachine) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *PhysicalMachine) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetRam

`func (o *PhysicalMachine) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *PhysicalMachine) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *PhysicalMachine) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *PhysicalMachine) HasRam() bool`

HasRam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


