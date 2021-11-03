# FullSoftwareAssetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**Bits** | Pointer to **string** |  | [optional] 
**Dependencies** | Pointer to **string** |  | [optional] 
**OsFamily** | Pointer to **string** |  | [optional] 
**ProductFamily** | Pointer to **string** |  | [optional] 
**RequiredCpu** | Pointer to **int32** |  | [optional] 
**RequiredCpuSpeed** | Pointer to **int32** |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**VendorMessage** | Pointer to **string** |  | [optional] 
**BuildEngine** | Pointer to **string** |  | [optional] 
**BuildScript** | Pointer to **string** |  | [optional] 
**BuildScriptName** | Pointer to **string** |  | [optional] 
**CheckoutScript** | Pointer to **string** |  | [optional] 
**CheckoutScriptName** | Pointer to **string** |  | [optional] 
**DeployScript** | Pointer to **string** |  | [optional] 
**DeployScriptName** | Pointer to **string** |  | [optional] 
**InstallScript** | Pointer to **string** |  | [optional] 
**InstallScriptName** | Pointer to **string** |  | [optional] 
**RequiredDisk** | Pointer to **int32** |  | [optional] 
**RequiredRam** | Pointer to **int32** |  | [optional] 
**SoftwareType** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**SourceCodeType** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewFullSoftwareAssetAllOf

`func NewFullSoftwareAssetAllOf() *FullSoftwareAssetAllOf`

NewFullSoftwareAssetAllOf instantiates a new FullSoftwareAssetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullSoftwareAssetAllOfWithDefaults

`func NewFullSoftwareAssetAllOfWithDefaults() *FullSoftwareAssetAllOf`

NewFullSoftwareAssetAllOfWithDefaults instantiates a new FullSoftwareAssetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *FullSoftwareAssetAllOf) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *FullSoftwareAssetAllOf) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *FullSoftwareAssetAllOf) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *FullSoftwareAssetAllOf) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBits

`func (o *FullSoftwareAssetAllOf) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *FullSoftwareAssetAllOf) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *FullSoftwareAssetAllOf) SetBits(v string)`

SetBits sets Bits field to given value.

### HasBits

`func (o *FullSoftwareAssetAllOf) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetDependencies

`func (o *FullSoftwareAssetAllOf) GetDependencies() string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *FullSoftwareAssetAllOf) GetDependenciesOk() (*string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *FullSoftwareAssetAllOf) SetDependencies(v string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *FullSoftwareAssetAllOf) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetOsFamily

`func (o *FullSoftwareAssetAllOf) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *FullSoftwareAssetAllOf) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *FullSoftwareAssetAllOf) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *FullSoftwareAssetAllOf) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetProductFamily

`func (o *FullSoftwareAssetAllOf) GetProductFamily() string`

GetProductFamily returns the ProductFamily field if non-nil, zero value otherwise.

### GetProductFamilyOk

`func (o *FullSoftwareAssetAllOf) GetProductFamilyOk() (*string, bool)`

GetProductFamilyOk returns a tuple with the ProductFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFamily

`func (o *FullSoftwareAssetAllOf) SetProductFamily(v string)`

SetProductFamily sets ProductFamily field to given value.

### HasProductFamily

`func (o *FullSoftwareAssetAllOf) HasProductFamily() bool`

HasProductFamily returns a boolean if a field has been set.

### GetRequiredCpu

`func (o *FullSoftwareAssetAllOf) GetRequiredCpu() int32`

GetRequiredCpu returns the RequiredCpu field if non-nil, zero value otherwise.

### GetRequiredCpuOk

`func (o *FullSoftwareAssetAllOf) GetRequiredCpuOk() (*int32, bool)`

GetRequiredCpuOk returns a tuple with the RequiredCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCpu

`func (o *FullSoftwareAssetAllOf) SetRequiredCpu(v int32)`

SetRequiredCpu sets RequiredCpu field to given value.

### HasRequiredCpu

`func (o *FullSoftwareAssetAllOf) HasRequiredCpu() bool`

HasRequiredCpu returns a boolean if a field has been set.

### GetRequiredCpuSpeed

`func (o *FullSoftwareAssetAllOf) GetRequiredCpuSpeed() int32`

GetRequiredCpuSpeed returns the RequiredCpuSpeed field if non-nil, zero value otherwise.

### GetRequiredCpuSpeedOk

`func (o *FullSoftwareAssetAllOf) GetRequiredCpuSpeedOk() (*int32, bool)`

GetRequiredCpuSpeedOk returns a tuple with the RequiredCpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCpuSpeed

`func (o *FullSoftwareAssetAllOf) SetRequiredCpuSpeed(v int32)`

SetRequiredCpuSpeed sets RequiredCpuSpeed field to given value.

### HasRequiredCpuSpeed

`func (o *FullSoftwareAssetAllOf) HasRequiredCpuSpeed() bool`

HasRequiredCpuSpeed returns a boolean if a field has been set.

### GetApplicationType

`func (o *FullSoftwareAssetAllOf) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *FullSoftwareAssetAllOf) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *FullSoftwareAssetAllOf) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *FullSoftwareAssetAllOf) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetVendorMessage

`func (o *FullSoftwareAssetAllOf) GetVendorMessage() string`

GetVendorMessage returns the VendorMessage field if non-nil, zero value otherwise.

### GetVendorMessageOk

`func (o *FullSoftwareAssetAllOf) GetVendorMessageOk() (*string, bool)`

GetVendorMessageOk returns a tuple with the VendorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorMessage

`func (o *FullSoftwareAssetAllOf) SetVendorMessage(v string)`

SetVendorMessage sets VendorMessage field to given value.

### HasVendorMessage

`func (o *FullSoftwareAssetAllOf) HasVendorMessage() bool`

HasVendorMessage returns a boolean if a field has been set.

### GetBuildEngine

`func (o *FullSoftwareAssetAllOf) GetBuildEngine() string`

GetBuildEngine returns the BuildEngine field if non-nil, zero value otherwise.

### GetBuildEngineOk

`func (o *FullSoftwareAssetAllOf) GetBuildEngineOk() (*string, bool)`

GetBuildEngineOk returns a tuple with the BuildEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEngine

`func (o *FullSoftwareAssetAllOf) SetBuildEngine(v string)`

SetBuildEngine sets BuildEngine field to given value.

### HasBuildEngine

`func (o *FullSoftwareAssetAllOf) HasBuildEngine() bool`

HasBuildEngine returns a boolean if a field has been set.

### GetBuildScript

`func (o *FullSoftwareAssetAllOf) GetBuildScript() string`

GetBuildScript returns the BuildScript field if non-nil, zero value otherwise.

### GetBuildScriptOk

`func (o *FullSoftwareAssetAllOf) GetBuildScriptOk() (*string, bool)`

GetBuildScriptOk returns a tuple with the BuildScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildScript

`func (o *FullSoftwareAssetAllOf) SetBuildScript(v string)`

SetBuildScript sets BuildScript field to given value.

### HasBuildScript

`func (o *FullSoftwareAssetAllOf) HasBuildScript() bool`

HasBuildScript returns a boolean if a field has been set.

### GetBuildScriptName

`func (o *FullSoftwareAssetAllOf) GetBuildScriptName() string`

GetBuildScriptName returns the BuildScriptName field if non-nil, zero value otherwise.

### GetBuildScriptNameOk

`func (o *FullSoftwareAssetAllOf) GetBuildScriptNameOk() (*string, bool)`

GetBuildScriptNameOk returns a tuple with the BuildScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildScriptName

`func (o *FullSoftwareAssetAllOf) SetBuildScriptName(v string)`

SetBuildScriptName sets BuildScriptName field to given value.

### HasBuildScriptName

`func (o *FullSoftwareAssetAllOf) HasBuildScriptName() bool`

HasBuildScriptName returns a boolean if a field has been set.

### GetCheckoutScript

`func (o *FullSoftwareAssetAllOf) GetCheckoutScript() string`

GetCheckoutScript returns the CheckoutScript field if non-nil, zero value otherwise.

### GetCheckoutScriptOk

`func (o *FullSoftwareAssetAllOf) GetCheckoutScriptOk() (*string, bool)`

GetCheckoutScriptOk returns a tuple with the CheckoutScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutScript

`func (o *FullSoftwareAssetAllOf) SetCheckoutScript(v string)`

SetCheckoutScript sets CheckoutScript field to given value.

### HasCheckoutScript

`func (o *FullSoftwareAssetAllOf) HasCheckoutScript() bool`

HasCheckoutScript returns a boolean if a field has been set.

### GetCheckoutScriptName

`func (o *FullSoftwareAssetAllOf) GetCheckoutScriptName() string`

GetCheckoutScriptName returns the CheckoutScriptName field if non-nil, zero value otherwise.

### GetCheckoutScriptNameOk

`func (o *FullSoftwareAssetAllOf) GetCheckoutScriptNameOk() (*string, bool)`

GetCheckoutScriptNameOk returns a tuple with the CheckoutScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutScriptName

`func (o *FullSoftwareAssetAllOf) SetCheckoutScriptName(v string)`

SetCheckoutScriptName sets CheckoutScriptName field to given value.

### HasCheckoutScriptName

`func (o *FullSoftwareAssetAllOf) HasCheckoutScriptName() bool`

HasCheckoutScriptName returns a boolean if a field has been set.

### GetDeployScript

`func (o *FullSoftwareAssetAllOf) GetDeployScript() string`

GetDeployScript returns the DeployScript field if non-nil, zero value otherwise.

### GetDeployScriptOk

`func (o *FullSoftwareAssetAllOf) GetDeployScriptOk() (*string, bool)`

GetDeployScriptOk returns a tuple with the DeployScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployScript

`func (o *FullSoftwareAssetAllOf) SetDeployScript(v string)`

SetDeployScript sets DeployScript field to given value.

### HasDeployScript

`func (o *FullSoftwareAssetAllOf) HasDeployScript() bool`

HasDeployScript returns a boolean if a field has been set.

### GetDeployScriptName

`func (o *FullSoftwareAssetAllOf) GetDeployScriptName() string`

GetDeployScriptName returns the DeployScriptName field if non-nil, zero value otherwise.

### GetDeployScriptNameOk

`func (o *FullSoftwareAssetAllOf) GetDeployScriptNameOk() (*string, bool)`

GetDeployScriptNameOk returns a tuple with the DeployScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployScriptName

`func (o *FullSoftwareAssetAllOf) SetDeployScriptName(v string)`

SetDeployScriptName sets DeployScriptName field to given value.

### HasDeployScriptName

`func (o *FullSoftwareAssetAllOf) HasDeployScriptName() bool`

HasDeployScriptName returns a boolean if a field has been set.

### GetInstallScript

`func (o *FullSoftwareAssetAllOf) GetInstallScript() string`

GetInstallScript returns the InstallScript field if non-nil, zero value otherwise.

### GetInstallScriptOk

`func (o *FullSoftwareAssetAllOf) GetInstallScriptOk() (*string, bool)`

GetInstallScriptOk returns a tuple with the InstallScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallScript

`func (o *FullSoftwareAssetAllOf) SetInstallScript(v string)`

SetInstallScript sets InstallScript field to given value.

### HasInstallScript

`func (o *FullSoftwareAssetAllOf) HasInstallScript() bool`

HasInstallScript returns a boolean if a field has been set.

### GetInstallScriptName

`func (o *FullSoftwareAssetAllOf) GetInstallScriptName() string`

GetInstallScriptName returns the InstallScriptName field if non-nil, zero value otherwise.

### GetInstallScriptNameOk

`func (o *FullSoftwareAssetAllOf) GetInstallScriptNameOk() (*string, bool)`

GetInstallScriptNameOk returns a tuple with the InstallScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallScriptName

`func (o *FullSoftwareAssetAllOf) SetInstallScriptName(v string)`

SetInstallScriptName sets InstallScriptName field to given value.

### HasInstallScriptName

`func (o *FullSoftwareAssetAllOf) HasInstallScriptName() bool`

HasInstallScriptName returns a boolean if a field has been set.

### GetRequiredDisk

`func (o *FullSoftwareAssetAllOf) GetRequiredDisk() int32`

GetRequiredDisk returns the RequiredDisk field if non-nil, zero value otherwise.

### GetRequiredDiskOk

`func (o *FullSoftwareAssetAllOf) GetRequiredDiskOk() (*int32, bool)`

GetRequiredDiskOk returns a tuple with the RequiredDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDisk

`func (o *FullSoftwareAssetAllOf) SetRequiredDisk(v int32)`

SetRequiredDisk sets RequiredDisk field to given value.

### HasRequiredDisk

`func (o *FullSoftwareAssetAllOf) HasRequiredDisk() bool`

HasRequiredDisk returns a boolean if a field has been set.

### GetRequiredRam

`func (o *FullSoftwareAssetAllOf) GetRequiredRam() int32`

GetRequiredRam returns the RequiredRam field if non-nil, zero value otherwise.

### GetRequiredRamOk

`func (o *FullSoftwareAssetAllOf) GetRequiredRamOk() (*int32, bool)`

GetRequiredRamOk returns a tuple with the RequiredRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRam

`func (o *FullSoftwareAssetAllOf) SetRequiredRam(v int32)`

SetRequiredRam sets RequiredRam field to given value.

### HasRequiredRam

`func (o *FullSoftwareAssetAllOf) HasRequiredRam() bool`

HasRequiredRam returns a boolean if a field has been set.

### GetSoftwareType

`func (o *FullSoftwareAssetAllOf) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *FullSoftwareAssetAllOf) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *FullSoftwareAssetAllOf) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.

### HasSoftwareType

`func (o *FullSoftwareAssetAllOf) HasSoftwareType() bool`

HasSoftwareType returns a boolean if a field has been set.

### GetVersion

`func (o *FullSoftwareAssetAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FullSoftwareAssetAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FullSoftwareAssetAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FullSoftwareAssetAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSourceCodeType

`func (o *FullSoftwareAssetAllOf) GetSourceCodeType() string`

GetSourceCodeType returns the SourceCodeType field if non-nil, zero value otherwise.

### GetSourceCodeTypeOk

`func (o *FullSoftwareAssetAllOf) GetSourceCodeTypeOk() (*string, bool)`

GetSourceCodeTypeOk returns a tuple with the SourceCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeType

`func (o *FullSoftwareAssetAllOf) SetSourceCodeType(v string)`

SetSourceCodeType sets SourceCodeType field to given value.

### HasSourceCodeType

`func (o *FullSoftwareAssetAllOf) HasSourceCodeType() bool`

HasSourceCodeType returns a boolean if a field has been set.

### GetVendor

`func (o *FullSoftwareAssetAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FullSoftwareAssetAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FullSoftwareAssetAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FullSoftwareAssetAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


