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

// FullSoftwareAssetAllOf struct for FullSoftwareAssetAllOf
type FullSoftwareAssetAllOf struct {
	Architecture *string `json:"architecture,omitempty"`
	Bits *string `json:"bits,omitempty"`
	Dependencies *string `json:"dependencies,omitempty"`
	OsFamily *string `json:"osFamily,omitempty"`
	ProductFamily *string `json:"productFamily,omitempty"`
	RequiredCpu *int32 `json:"requiredCpu,omitempty"`
	RequiredCpuSpeed *int32 `json:"requiredCpuSpeed,omitempty"`
	ApplicationType *string `json:"applicationType,omitempty"`
	VendorMessage *string `json:"vendorMessage,omitempty"`
	BuildEngine *string `json:"buildEngine,omitempty"`
	BuildScript *string `json:"buildScript,omitempty"`
	BuildScriptName *string `json:"buildScriptName,omitempty"`
	CheckoutScript *string `json:"checkoutScript,omitempty"`
	CheckoutScriptName *string `json:"checkoutScriptName,omitempty"`
	DeployScript *string `json:"deployScript,omitempty"`
	DeployScriptName *string `json:"deployScriptName,omitempty"`
	InstallScript *string `json:"installScript,omitempty"`
	InstallScriptName *string `json:"installScriptName,omitempty"`
	RequiredDisk *int32 `json:"requiredDisk,omitempty"`
	RequiredRam *int32 `json:"requiredRam,omitempty"`
	SoftwareType *string `json:"softwareType,omitempty"`
	Version *string `json:"version,omitempty"`
	SourceCodeType *string `json:"sourceCodeType,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewFullSoftwareAssetAllOf instantiates a new FullSoftwareAssetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullSoftwareAssetAllOf() *FullSoftwareAssetAllOf {
	this := FullSoftwareAssetAllOf{}
	return &this
}

// NewFullSoftwareAssetAllOfWithDefaults instantiates a new FullSoftwareAssetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullSoftwareAssetAllOfWithDefaults() *FullSoftwareAssetAllOf {
	this := FullSoftwareAssetAllOf{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *FullSoftwareAssetAllOf) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetBits() string {
	if o == nil || o.Bits == nil {
		var ret string
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetBitsOk() (*string, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given string and assigns it to the Bits field.
func (o *FullSoftwareAssetAllOf) SetBits(v string) {
	o.Bits = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetDependencies() string {
	if o == nil || o.Dependencies == nil {
		var ret string
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetDependenciesOk() (*string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given string and assigns it to the Dependencies field.
func (o *FullSoftwareAssetAllOf) SetDependencies(v string) {
	o.Dependencies = &v
}

// GetOsFamily returns the OsFamily field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetOsFamily() string {
	if o == nil || o.OsFamily == nil {
		var ret string
		return ret
	}
	return *o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetOsFamilyOk() (*string, bool) {
	if o == nil || o.OsFamily == nil {
		return nil, false
	}
	return o.OsFamily, true
}

// HasOsFamily returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasOsFamily() bool {
	if o != nil && o.OsFamily != nil {
		return true
	}

	return false
}

// SetOsFamily gets a reference to the given string and assigns it to the OsFamily field.
func (o *FullSoftwareAssetAllOf) SetOsFamily(v string) {
	o.OsFamily = &v
}

// GetProductFamily returns the ProductFamily field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetProductFamily() string {
	if o == nil || o.ProductFamily == nil {
		var ret string
		return ret
	}
	return *o.ProductFamily
}

// GetProductFamilyOk returns a tuple with the ProductFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetProductFamilyOk() (*string, bool) {
	if o == nil || o.ProductFamily == nil {
		return nil, false
	}
	return o.ProductFamily, true
}

// HasProductFamily returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasProductFamily() bool {
	if o != nil && o.ProductFamily != nil {
		return true
	}

	return false
}

// SetProductFamily gets a reference to the given string and assigns it to the ProductFamily field.
func (o *FullSoftwareAssetAllOf) SetProductFamily(v string) {
	o.ProductFamily = &v
}

// GetRequiredCpu returns the RequiredCpu field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetRequiredCpu() int32 {
	if o == nil || o.RequiredCpu == nil {
		var ret int32
		return ret
	}
	return *o.RequiredCpu
}

// GetRequiredCpuOk returns a tuple with the RequiredCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetRequiredCpuOk() (*int32, bool) {
	if o == nil || o.RequiredCpu == nil {
		return nil, false
	}
	return o.RequiredCpu, true
}

// HasRequiredCpu returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasRequiredCpu() bool {
	if o != nil && o.RequiredCpu != nil {
		return true
	}

	return false
}

// SetRequiredCpu gets a reference to the given int32 and assigns it to the RequiredCpu field.
func (o *FullSoftwareAssetAllOf) SetRequiredCpu(v int32) {
	o.RequiredCpu = &v
}

// GetRequiredCpuSpeed returns the RequiredCpuSpeed field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetRequiredCpuSpeed() int32 {
	if o == nil || o.RequiredCpuSpeed == nil {
		var ret int32
		return ret
	}
	return *o.RequiredCpuSpeed
}

// GetRequiredCpuSpeedOk returns a tuple with the RequiredCpuSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetRequiredCpuSpeedOk() (*int32, bool) {
	if o == nil || o.RequiredCpuSpeed == nil {
		return nil, false
	}
	return o.RequiredCpuSpeed, true
}

// HasRequiredCpuSpeed returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasRequiredCpuSpeed() bool {
	if o != nil && o.RequiredCpuSpeed != nil {
		return true
	}

	return false
}

// SetRequiredCpuSpeed gets a reference to the given int32 and assigns it to the RequiredCpuSpeed field.
func (o *FullSoftwareAssetAllOf) SetRequiredCpuSpeed(v int32) {
	o.RequiredCpuSpeed = &v
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetApplicationType() string {
	if o == nil || o.ApplicationType == nil {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetApplicationTypeOk() (*string, bool) {
	if o == nil || o.ApplicationType == nil {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasApplicationType() bool {
	if o != nil && o.ApplicationType != nil {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *FullSoftwareAssetAllOf) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetVendorMessage returns the VendorMessage field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetVendorMessage() string {
	if o == nil || o.VendorMessage == nil {
		var ret string
		return ret
	}
	return *o.VendorMessage
}

// GetVendorMessageOk returns a tuple with the VendorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetVendorMessageOk() (*string, bool) {
	if o == nil || o.VendorMessage == nil {
		return nil, false
	}
	return o.VendorMessage, true
}

// HasVendorMessage returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasVendorMessage() bool {
	if o != nil && o.VendorMessage != nil {
		return true
	}

	return false
}

// SetVendorMessage gets a reference to the given string and assigns it to the VendorMessage field.
func (o *FullSoftwareAssetAllOf) SetVendorMessage(v string) {
	o.VendorMessage = &v
}

// GetBuildEngine returns the BuildEngine field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetBuildEngine() string {
	if o == nil || o.BuildEngine == nil {
		var ret string
		return ret
	}
	return *o.BuildEngine
}

// GetBuildEngineOk returns a tuple with the BuildEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetBuildEngineOk() (*string, bool) {
	if o == nil || o.BuildEngine == nil {
		return nil, false
	}
	return o.BuildEngine, true
}

// HasBuildEngine returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasBuildEngine() bool {
	if o != nil && o.BuildEngine != nil {
		return true
	}

	return false
}

// SetBuildEngine gets a reference to the given string and assigns it to the BuildEngine field.
func (o *FullSoftwareAssetAllOf) SetBuildEngine(v string) {
	o.BuildEngine = &v
}

// GetBuildScript returns the BuildScript field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetBuildScript() string {
	if o == nil || o.BuildScript == nil {
		var ret string
		return ret
	}
	return *o.BuildScript
}

// GetBuildScriptOk returns a tuple with the BuildScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetBuildScriptOk() (*string, bool) {
	if o == nil || o.BuildScript == nil {
		return nil, false
	}
	return o.BuildScript, true
}

// HasBuildScript returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasBuildScript() bool {
	if o != nil && o.BuildScript != nil {
		return true
	}

	return false
}

// SetBuildScript gets a reference to the given string and assigns it to the BuildScript field.
func (o *FullSoftwareAssetAllOf) SetBuildScript(v string) {
	o.BuildScript = &v
}

// GetBuildScriptName returns the BuildScriptName field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetBuildScriptName() string {
	if o == nil || o.BuildScriptName == nil {
		var ret string
		return ret
	}
	return *o.BuildScriptName
}

// GetBuildScriptNameOk returns a tuple with the BuildScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetBuildScriptNameOk() (*string, bool) {
	if o == nil || o.BuildScriptName == nil {
		return nil, false
	}
	return o.BuildScriptName, true
}

// HasBuildScriptName returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasBuildScriptName() bool {
	if o != nil && o.BuildScriptName != nil {
		return true
	}

	return false
}

// SetBuildScriptName gets a reference to the given string and assigns it to the BuildScriptName field.
func (o *FullSoftwareAssetAllOf) SetBuildScriptName(v string) {
	o.BuildScriptName = &v
}

// GetCheckoutScript returns the CheckoutScript field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetCheckoutScript() string {
	if o == nil || o.CheckoutScript == nil {
		var ret string
		return ret
	}
	return *o.CheckoutScript
}

// GetCheckoutScriptOk returns a tuple with the CheckoutScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetCheckoutScriptOk() (*string, bool) {
	if o == nil || o.CheckoutScript == nil {
		return nil, false
	}
	return o.CheckoutScript, true
}

// HasCheckoutScript returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasCheckoutScript() bool {
	if o != nil && o.CheckoutScript != nil {
		return true
	}

	return false
}

// SetCheckoutScript gets a reference to the given string and assigns it to the CheckoutScript field.
func (o *FullSoftwareAssetAllOf) SetCheckoutScript(v string) {
	o.CheckoutScript = &v
}

// GetCheckoutScriptName returns the CheckoutScriptName field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetCheckoutScriptName() string {
	if o == nil || o.CheckoutScriptName == nil {
		var ret string
		return ret
	}
	return *o.CheckoutScriptName
}

// GetCheckoutScriptNameOk returns a tuple with the CheckoutScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetCheckoutScriptNameOk() (*string, bool) {
	if o == nil || o.CheckoutScriptName == nil {
		return nil, false
	}
	return o.CheckoutScriptName, true
}

// HasCheckoutScriptName returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasCheckoutScriptName() bool {
	if o != nil && o.CheckoutScriptName != nil {
		return true
	}

	return false
}

// SetCheckoutScriptName gets a reference to the given string and assigns it to the CheckoutScriptName field.
func (o *FullSoftwareAssetAllOf) SetCheckoutScriptName(v string) {
	o.CheckoutScriptName = &v
}

// GetDeployScript returns the DeployScript field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetDeployScript() string {
	if o == nil || o.DeployScript == nil {
		var ret string
		return ret
	}
	return *o.DeployScript
}

// GetDeployScriptOk returns a tuple with the DeployScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetDeployScriptOk() (*string, bool) {
	if o == nil || o.DeployScript == nil {
		return nil, false
	}
	return o.DeployScript, true
}

// HasDeployScript returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasDeployScript() bool {
	if o != nil && o.DeployScript != nil {
		return true
	}

	return false
}

// SetDeployScript gets a reference to the given string and assigns it to the DeployScript field.
func (o *FullSoftwareAssetAllOf) SetDeployScript(v string) {
	o.DeployScript = &v
}

// GetDeployScriptName returns the DeployScriptName field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetDeployScriptName() string {
	if o == nil || o.DeployScriptName == nil {
		var ret string
		return ret
	}
	return *o.DeployScriptName
}

// GetDeployScriptNameOk returns a tuple with the DeployScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetDeployScriptNameOk() (*string, bool) {
	if o == nil || o.DeployScriptName == nil {
		return nil, false
	}
	return o.DeployScriptName, true
}

// HasDeployScriptName returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasDeployScriptName() bool {
	if o != nil && o.DeployScriptName != nil {
		return true
	}

	return false
}

// SetDeployScriptName gets a reference to the given string and assigns it to the DeployScriptName field.
func (o *FullSoftwareAssetAllOf) SetDeployScriptName(v string) {
	o.DeployScriptName = &v
}

// GetInstallScript returns the InstallScript field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetInstallScript() string {
	if o == nil || o.InstallScript == nil {
		var ret string
		return ret
	}
	return *o.InstallScript
}

// GetInstallScriptOk returns a tuple with the InstallScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetInstallScriptOk() (*string, bool) {
	if o == nil || o.InstallScript == nil {
		return nil, false
	}
	return o.InstallScript, true
}

// HasInstallScript returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasInstallScript() bool {
	if o != nil && o.InstallScript != nil {
		return true
	}

	return false
}

// SetInstallScript gets a reference to the given string and assigns it to the InstallScript field.
func (o *FullSoftwareAssetAllOf) SetInstallScript(v string) {
	o.InstallScript = &v
}

// GetInstallScriptName returns the InstallScriptName field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetInstallScriptName() string {
	if o == nil || o.InstallScriptName == nil {
		var ret string
		return ret
	}
	return *o.InstallScriptName
}

// GetInstallScriptNameOk returns a tuple with the InstallScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetInstallScriptNameOk() (*string, bool) {
	if o == nil || o.InstallScriptName == nil {
		return nil, false
	}
	return o.InstallScriptName, true
}

// HasInstallScriptName returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasInstallScriptName() bool {
	if o != nil && o.InstallScriptName != nil {
		return true
	}

	return false
}

// SetInstallScriptName gets a reference to the given string and assigns it to the InstallScriptName field.
func (o *FullSoftwareAssetAllOf) SetInstallScriptName(v string) {
	o.InstallScriptName = &v
}

// GetRequiredDisk returns the RequiredDisk field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetRequiredDisk() int32 {
	if o == nil || o.RequiredDisk == nil {
		var ret int32
		return ret
	}
	return *o.RequiredDisk
}

// GetRequiredDiskOk returns a tuple with the RequiredDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetRequiredDiskOk() (*int32, bool) {
	if o == nil || o.RequiredDisk == nil {
		return nil, false
	}
	return o.RequiredDisk, true
}

// HasRequiredDisk returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasRequiredDisk() bool {
	if o != nil && o.RequiredDisk != nil {
		return true
	}

	return false
}

// SetRequiredDisk gets a reference to the given int32 and assigns it to the RequiredDisk field.
func (o *FullSoftwareAssetAllOf) SetRequiredDisk(v int32) {
	o.RequiredDisk = &v
}

// GetRequiredRam returns the RequiredRam field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetRequiredRam() int32 {
	if o == nil || o.RequiredRam == nil {
		var ret int32
		return ret
	}
	return *o.RequiredRam
}

// GetRequiredRamOk returns a tuple with the RequiredRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetRequiredRamOk() (*int32, bool) {
	if o == nil || o.RequiredRam == nil {
		return nil, false
	}
	return o.RequiredRam, true
}

// HasRequiredRam returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasRequiredRam() bool {
	if o != nil && o.RequiredRam != nil {
		return true
	}

	return false
}

// SetRequiredRam gets a reference to the given int32 and assigns it to the RequiredRam field.
func (o *FullSoftwareAssetAllOf) SetRequiredRam(v int32) {
	o.RequiredRam = &v
}

// GetSoftwareType returns the SoftwareType field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetSoftwareType() string {
	if o == nil || o.SoftwareType == nil {
		var ret string
		return ret
	}
	return *o.SoftwareType
}

// GetSoftwareTypeOk returns a tuple with the SoftwareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetSoftwareTypeOk() (*string, bool) {
	if o == nil || o.SoftwareType == nil {
		return nil, false
	}
	return o.SoftwareType, true
}

// HasSoftwareType returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasSoftwareType() bool {
	if o != nil && o.SoftwareType != nil {
		return true
	}

	return false
}

// SetSoftwareType gets a reference to the given string and assigns it to the SoftwareType field.
func (o *FullSoftwareAssetAllOf) SetSoftwareType(v string) {
	o.SoftwareType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FullSoftwareAssetAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetSourceCodeType returns the SourceCodeType field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetSourceCodeType() string {
	if o == nil || o.SourceCodeType == nil {
		var ret string
		return ret
	}
	return *o.SourceCodeType
}

// GetSourceCodeTypeOk returns a tuple with the SourceCodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetSourceCodeTypeOk() (*string, bool) {
	if o == nil || o.SourceCodeType == nil {
		return nil, false
	}
	return o.SourceCodeType, true
}

// HasSourceCodeType returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasSourceCodeType() bool {
	if o != nil && o.SourceCodeType != nil {
		return true
	}

	return false
}

// SetSourceCodeType gets a reference to the given string and assigns it to the SourceCodeType field.
func (o *FullSoftwareAssetAllOf) SetSourceCodeType(v string) {
	o.SourceCodeType = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *FullSoftwareAssetAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullSoftwareAssetAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *FullSoftwareAssetAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *FullSoftwareAssetAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o FullSoftwareAssetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Architecture != nil {
		toSerialize["architecture"] = o.Architecture
	}
	if o.Bits != nil {
		toSerialize["bits"] = o.Bits
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.OsFamily != nil {
		toSerialize["osFamily"] = o.OsFamily
	}
	if o.ProductFamily != nil {
		toSerialize["productFamily"] = o.ProductFamily
	}
	if o.RequiredCpu != nil {
		toSerialize["requiredCpu"] = o.RequiredCpu
	}
	if o.RequiredCpuSpeed != nil {
		toSerialize["requiredCpuSpeed"] = o.RequiredCpuSpeed
	}
	if o.ApplicationType != nil {
		toSerialize["applicationType"] = o.ApplicationType
	}
	if o.VendorMessage != nil {
		toSerialize["vendorMessage"] = o.VendorMessage
	}
	if o.BuildEngine != nil {
		toSerialize["buildEngine"] = o.BuildEngine
	}
	if o.BuildScript != nil {
		toSerialize["buildScript"] = o.BuildScript
	}
	if o.BuildScriptName != nil {
		toSerialize["buildScriptName"] = o.BuildScriptName
	}
	if o.CheckoutScript != nil {
		toSerialize["checkoutScript"] = o.CheckoutScript
	}
	if o.CheckoutScriptName != nil {
		toSerialize["checkoutScriptName"] = o.CheckoutScriptName
	}
	if o.DeployScript != nil {
		toSerialize["deployScript"] = o.DeployScript
	}
	if o.DeployScriptName != nil {
		toSerialize["deployScriptName"] = o.DeployScriptName
	}
	if o.InstallScript != nil {
		toSerialize["installScript"] = o.InstallScript
	}
	if o.InstallScriptName != nil {
		toSerialize["installScriptName"] = o.InstallScriptName
	}
	if o.RequiredDisk != nil {
		toSerialize["requiredDisk"] = o.RequiredDisk
	}
	if o.RequiredRam != nil {
		toSerialize["requiredRam"] = o.RequiredRam
	}
	if o.SoftwareType != nil {
		toSerialize["softwareType"] = o.SoftwareType
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.SourceCodeType != nil {
		toSerialize["sourceCodeType"] = o.SourceCodeType
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableFullSoftwareAssetAllOf struct {
	value *FullSoftwareAssetAllOf
	isSet bool
}

func (v NullableFullSoftwareAssetAllOf) Get() *FullSoftwareAssetAllOf {
	return v.value
}

func (v *NullableFullSoftwareAssetAllOf) Set(val *FullSoftwareAssetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullSoftwareAssetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullSoftwareAssetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullSoftwareAssetAllOf(val *FullSoftwareAssetAllOf) *NullableFullSoftwareAssetAllOf {
	return &NullableFullSoftwareAssetAllOf{value: val, isSet: true}
}

func (v NullableFullSoftwareAssetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullSoftwareAssetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


