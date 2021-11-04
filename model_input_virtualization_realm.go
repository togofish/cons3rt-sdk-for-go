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

// InputVirtualizationRealm struct for InputVirtualizationRealm
type InputVirtualizationRealm struct {
	VirtualizationRealmType *string `json:"virtualizationRealmType,omitempty"`
	AccessPoint *string `json:"accessPoint,omitempty"`
	AccountId string `json:"accountId"`
	Cidr string `json:"cidr"`
	DefaultWindowsDomainName *string `json:"defaultWindowsDomainName,omitempty"`
	Description string `json:"description"`
	Id *int32 `json:"id,omitempty"`
	LocalStorageName *string `json:"localStorageName,omitempty"`
	MaximumNumCpus *int32 `json:"maximumNumCpus,omitempty"`
	MaximumNumGpus *int32 `json:"maximumNumGpus,omitempty"`
	MaximumRamInMegabytes *int32 `json:"maximumRamInMegabytes,omitempty"`
	MaximumStorageInMegabytes *int32 `json:"maximumStorageInMegabytes,omitempty"`
	MaximumVirtualMachines *int32 `json:"maximumVirtualMachines,omitempty"`
	Name string `json:"name"`
	Password string `json:"password"`
	PowerOnDelayBase *int64 `json:"powerOnDelayBase,omitempty"`
	PowerOnInitialDelayBase *int32 `json:"powerOnInitialDelayBase,omitempty"`
	PowerOnMaximumDelay *int32 `json:"powerOnMaximumDelay,omitempty"`
	PowerOnMinimumDelay *int32 `json:"powerOnMinimumDelay,omitempty"`
	RemoteAccessConfig *RemoteAccessConfig `json:"remoteAccessConfig,omitempty"`
	State *string `json:"state,omitempty"`
	Username string `json:"username"`
	ZoneCount *int32 `json:"zoneCount,omitempty"`
}

// NewInputVirtualizationRealm instantiates a new InputVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVirtualizationRealm(accountId string, cidr string, description string, name string, password string, username string) *InputVirtualizationRealm {
	this := InputVirtualizationRealm{}
	this.AccountId = accountId
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	this.Password = password
	this.Username = username
	return &this
}

// NewInputVirtualizationRealmWithDefaults instantiates a new InputVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVirtualizationRealmWithDefaults() *InputVirtualizationRealm {
	this := InputVirtualizationRealm{}
	return &this
}

// GetVirtualizationRealmType returns the VirtualizationRealmType field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetVirtualizationRealmType() string {
	if o == nil || o.VirtualizationRealmType == nil {
		var ret string
		return ret
	}
	return *o.VirtualizationRealmType
}

// GetVirtualizationRealmTypeOk returns a tuple with the VirtualizationRealmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetVirtualizationRealmTypeOk() (*string, bool) {
	if o == nil || o.VirtualizationRealmType == nil {
		return nil, false
	}
	return o.VirtualizationRealmType, true
}

// HasVirtualizationRealmType returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasVirtualizationRealmType() bool {
	if o != nil && o.VirtualizationRealmType != nil {
		return true
	}

	return false
}

// SetVirtualizationRealmType gets a reference to the given string and assigns it to the VirtualizationRealmType field.
func (o *InputVirtualizationRealm) SetVirtualizationRealmType(v string) {
	o.VirtualizationRealmType = &v
}

// GetAccessPoint returns the AccessPoint field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetAccessPoint() string {
	if o == nil || o.AccessPoint == nil {
		var ret string
		return ret
	}
	return *o.AccessPoint
}

// GetAccessPointOk returns a tuple with the AccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetAccessPointOk() (*string, bool) {
	if o == nil || o.AccessPoint == nil {
		return nil, false
	}
	return o.AccessPoint, true
}

// HasAccessPoint returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasAccessPoint() bool {
	if o != nil && o.AccessPoint != nil {
		return true
	}

	return false
}

// SetAccessPoint gets a reference to the given string and assigns it to the AccessPoint field.
func (o *InputVirtualizationRealm) SetAccessPoint(v string) {
	o.AccessPoint = &v
}

// GetAccountId returns the AccountId field value
func (o *InputVirtualizationRealm) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *InputVirtualizationRealm) SetAccountId(v string) {
	o.AccountId = v
}

// GetCidr returns the Cidr field value
func (o *InputVirtualizationRealm) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetCidrOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *InputVirtualizationRealm) SetCidr(v string) {
	o.Cidr = v
}

// GetDefaultWindowsDomainName returns the DefaultWindowsDomainName field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetDefaultWindowsDomainName() string {
	if o == nil || o.DefaultWindowsDomainName == nil {
		var ret string
		return ret
	}
	return *o.DefaultWindowsDomainName
}

// GetDefaultWindowsDomainNameOk returns a tuple with the DefaultWindowsDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetDefaultWindowsDomainNameOk() (*string, bool) {
	if o == nil || o.DefaultWindowsDomainName == nil {
		return nil, false
	}
	return o.DefaultWindowsDomainName, true
}

// HasDefaultWindowsDomainName returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasDefaultWindowsDomainName() bool {
	if o != nil && o.DefaultWindowsDomainName != nil {
		return true
	}

	return false
}

// SetDefaultWindowsDomainName gets a reference to the given string and assigns it to the DefaultWindowsDomainName field.
func (o *InputVirtualizationRealm) SetDefaultWindowsDomainName(v string) {
	o.DefaultWindowsDomainName = &v
}

// GetDescription returns the Description field value
func (o *InputVirtualizationRealm) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InputVirtualizationRealm) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InputVirtualizationRealm) SetId(v int32) {
	o.Id = &v
}

// GetLocalStorageName returns the LocalStorageName field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetLocalStorageName() string {
	if o == nil || o.LocalStorageName == nil {
		var ret string
		return ret
	}
	return *o.LocalStorageName
}

// GetLocalStorageNameOk returns a tuple with the LocalStorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetLocalStorageNameOk() (*string, bool) {
	if o == nil || o.LocalStorageName == nil {
		return nil, false
	}
	return o.LocalStorageName, true
}

// HasLocalStorageName returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasLocalStorageName() bool {
	if o != nil && o.LocalStorageName != nil {
		return true
	}

	return false
}

// SetLocalStorageName gets a reference to the given string and assigns it to the LocalStorageName field.
func (o *InputVirtualizationRealm) SetLocalStorageName(v string) {
	o.LocalStorageName = &v
}

// GetMaximumNumCpus returns the MaximumNumCpus field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetMaximumNumCpus() int32 {
	if o == nil || o.MaximumNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MaximumNumCpus
}

// GetMaximumNumCpusOk returns a tuple with the MaximumNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetMaximumNumCpusOk() (*int32, bool) {
	if o == nil || o.MaximumNumCpus == nil {
		return nil, false
	}
	return o.MaximumNumCpus, true
}

// HasMaximumNumCpus returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasMaximumNumCpus() bool {
	if o != nil && o.MaximumNumCpus != nil {
		return true
	}

	return false
}

// SetMaximumNumCpus gets a reference to the given int32 and assigns it to the MaximumNumCpus field.
func (o *InputVirtualizationRealm) SetMaximumNumCpus(v int32) {
	o.MaximumNumCpus = &v
}

// GetMaximumNumGpus returns the MaximumNumGpus field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetMaximumNumGpus() int32 {
	if o == nil || o.MaximumNumGpus == nil {
		var ret int32
		return ret
	}
	return *o.MaximumNumGpus
}

// GetMaximumNumGpusOk returns a tuple with the MaximumNumGpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetMaximumNumGpusOk() (*int32, bool) {
	if o == nil || o.MaximumNumGpus == nil {
		return nil, false
	}
	return o.MaximumNumGpus, true
}

// HasMaximumNumGpus returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasMaximumNumGpus() bool {
	if o != nil && o.MaximumNumGpus != nil {
		return true
	}

	return false
}

// SetMaximumNumGpus gets a reference to the given int32 and assigns it to the MaximumNumGpus field.
func (o *InputVirtualizationRealm) SetMaximumNumGpus(v int32) {
	o.MaximumNumGpus = &v
}

// GetMaximumRamInMegabytes returns the MaximumRamInMegabytes field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetMaximumRamInMegabytes() int32 {
	if o == nil || o.MaximumRamInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaximumRamInMegabytes
}

// GetMaximumRamInMegabytesOk returns a tuple with the MaximumRamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetMaximumRamInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaximumRamInMegabytes == nil {
		return nil, false
	}
	return o.MaximumRamInMegabytes, true
}

// HasMaximumRamInMegabytes returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasMaximumRamInMegabytes() bool {
	if o != nil && o.MaximumRamInMegabytes != nil {
		return true
	}

	return false
}

// SetMaximumRamInMegabytes gets a reference to the given int32 and assigns it to the MaximumRamInMegabytes field.
func (o *InputVirtualizationRealm) SetMaximumRamInMegabytes(v int32) {
	o.MaximumRamInMegabytes = &v
}

// GetMaximumStorageInMegabytes returns the MaximumStorageInMegabytes field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetMaximumStorageInMegabytes() int32 {
	if o == nil || o.MaximumStorageInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaximumStorageInMegabytes
}

// GetMaximumStorageInMegabytesOk returns a tuple with the MaximumStorageInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetMaximumStorageInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaximumStorageInMegabytes == nil {
		return nil, false
	}
	return o.MaximumStorageInMegabytes, true
}

// HasMaximumStorageInMegabytes returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasMaximumStorageInMegabytes() bool {
	if o != nil && o.MaximumStorageInMegabytes != nil {
		return true
	}

	return false
}

// SetMaximumStorageInMegabytes gets a reference to the given int32 and assigns it to the MaximumStorageInMegabytes field.
func (o *InputVirtualizationRealm) SetMaximumStorageInMegabytes(v int32) {
	o.MaximumStorageInMegabytes = &v
}

// GetMaximumVirtualMachines returns the MaximumVirtualMachines field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetMaximumVirtualMachines() int32 {
	if o == nil || o.MaximumVirtualMachines == nil {
		var ret int32
		return ret
	}
	return *o.MaximumVirtualMachines
}

// GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetMaximumVirtualMachinesOk() (*int32, bool) {
	if o == nil || o.MaximumVirtualMachines == nil {
		return nil, false
	}
	return o.MaximumVirtualMachines, true
}

// HasMaximumVirtualMachines returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasMaximumVirtualMachines() bool {
	if o != nil && o.MaximumVirtualMachines != nil {
		return true
	}

	return false
}

// SetMaximumVirtualMachines gets a reference to the given int32 and assigns it to the MaximumVirtualMachines field.
func (o *InputVirtualizationRealm) SetMaximumVirtualMachines(v int32) {
	o.MaximumVirtualMachines = &v
}

// GetName returns the Name field value
func (o *InputVirtualizationRealm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputVirtualizationRealm) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value
func (o *InputVirtualizationRealm) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *InputVirtualizationRealm) SetPassword(v string) {
	o.Password = v
}

// GetPowerOnDelayBase returns the PowerOnDelayBase field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetPowerOnDelayBase() int64 {
	if o == nil || o.PowerOnDelayBase == nil {
		var ret int64
		return ret
	}
	return *o.PowerOnDelayBase
}

// GetPowerOnDelayBaseOk returns a tuple with the PowerOnDelayBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetPowerOnDelayBaseOk() (*int64, bool) {
	if o == nil || o.PowerOnDelayBase == nil {
		return nil, false
	}
	return o.PowerOnDelayBase, true
}

// HasPowerOnDelayBase returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasPowerOnDelayBase() bool {
	if o != nil && o.PowerOnDelayBase != nil {
		return true
	}

	return false
}

// SetPowerOnDelayBase gets a reference to the given int64 and assigns it to the PowerOnDelayBase field.
func (o *InputVirtualizationRealm) SetPowerOnDelayBase(v int64) {
	o.PowerOnDelayBase = &v
}

// GetPowerOnInitialDelayBase returns the PowerOnInitialDelayBase field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetPowerOnInitialDelayBase() int32 {
	if o == nil || o.PowerOnInitialDelayBase == nil {
		var ret int32
		return ret
	}
	return *o.PowerOnInitialDelayBase
}

// GetPowerOnInitialDelayBaseOk returns a tuple with the PowerOnInitialDelayBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetPowerOnInitialDelayBaseOk() (*int32, bool) {
	if o == nil || o.PowerOnInitialDelayBase == nil {
		return nil, false
	}
	return o.PowerOnInitialDelayBase, true
}

// HasPowerOnInitialDelayBase returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasPowerOnInitialDelayBase() bool {
	if o != nil && o.PowerOnInitialDelayBase != nil {
		return true
	}

	return false
}

// SetPowerOnInitialDelayBase gets a reference to the given int32 and assigns it to the PowerOnInitialDelayBase field.
func (o *InputVirtualizationRealm) SetPowerOnInitialDelayBase(v int32) {
	o.PowerOnInitialDelayBase = &v
}

// GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetPowerOnMaximumDelay() int32 {
	if o == nil || o.PowerOnMaximumDelay == nil {
		var ret int32
		return ret
	}
	return *o.PowerOnMaximumDelay
}

// GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetPowerOnMaximumDelayOk() (*int32, bool) {
	if o == nil || o.PowerOnMaximumDelay == nil {
		return nil, false
	}
	return o.PowerOnMaximumDelay, true
}

// HasPowerOnMaximumDelay returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasPowerOnMaximumDelay() bool {
	if o != nil && o.PowerOnMaximumDelay != nil {
		return true
	}

	return false
}

// SetPowerOnMaximumDelay gets a reference to the given int32 and assigns it to the PowerOnMaximumDelay field.
func (o *InputVirtualizationRealm) SetPowerOnMaximumDelay(v int32) {
	o.PowerOnMaximumDelay = &v
}

// GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetPowerOnMinimumDelay() int32 {
	if o == nil || o.PowerOnMinimumDelay == nil {
		var ret int32
		return ret
	}
	return *o.PowerOnMinimumDelay
}

// GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetPowerOnMinimumDelayOk() (*int32, bool) {
	if o == nil || o.PowerOnMinimumDelay == nil {
		return nil, false
	}
	return o.PowerOnMinimumDelay, true
}

// HasPowerOnMinimumDelay returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasPowerOnMinimumDelay() bool {
	if o != nil && o.PowerOnMinimumDelay != nil {
		return true
	}

	return false
}

// SetPowerOnMinimumDelay gets a reference to the given int32 and assigns it to the PowerOnMinimumDelay field.
func (o *InputVirtualizationRealm) SetPowerOnMinimumDelay(v int32) {
	o.PowerOnMinimumDelay = &v
}

// GetRemoteAccessConfig returns the RemoteAccessConfig field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetRemoteAccessConfig() RemoteAccessConfig {
	if o == nil || o.RemoteAccessConfig == nil {
		var ret RemoteAccessConfig
		return ret
	}
	return *o.RemoteAccessConfig
}

// GetRemoteAccessConfigOk returns a tuple with the RemoteAccessConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetRemoteAccessConfigOk() (*RemoteAccessConfig, bool) {
	if o == nil || o.RemoteAccessConfig == nil {
		return nil, false
	}
	return o.RemoteAccessConfig, true
}

// HasRemoteAccessConfig returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasRemoteAccessConfig() bool {
	if o != nil && o.RemoteAccessConfig != nil {
		return true
	}

	return false
}

// SetRemoteAccessConfig gets a reference to the given RemoteAccessConfig and assigns it to the RemoteAccessConfig field.
func (o *InputVirtualizationRealm) SetRemoteAccessConfig(v RemoteAccessConfig) {
	o.RemoteAccessConfig = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *InputVirtualizationRealm) SetState(v string) {
	o.State = &v
}

// GetUsername returns the Username field value
func (o *InputVirtualizationRealm) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InputVirtualizationRealm) SetUsername(v string) {
	o.Username = v
}

// GetZoneCount returns the ZoneCount field value if set, zero value otherwise.
func (o *InputVirtualizationRealm) GetZoneCount() int32 {
	if o == nil || o.ZoneCount == nil {
		var ret int32
		return ret
	}
	return *o.ZoneCount
}

// GetZoneCountOk returns a tuple with the ZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputVirtualizationRealm) GetZoneCountOk() (*int32, bool) {
	if o == nil || o.ZoneCount == nil {
		return nil, false
	}
	return o.ZoneCount, true
}

// HasZoneCount returns a boolean if a field has been set.
func (o *InputVirtualizationRealm) HasZoneCount() bool {
	if o != nil && o.ZoneCount != nil {
		return true
	}

	return false
}

// SetZoneCount gets a reference to the given int32 and assigns it to the ZoneCount field.
func (o *InputVirtualizationRealm) SetZoneCount(v int32) {
	o.ZoneCount = &v
}

func (o InputVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualizationRealmType != nil {
		toSerialize["virtualizationRealmType"] = o.VirtualizationRealmType
	}
	if o.AccessPoint != nil {
		toSerialize["accessPoint"] = o.AccessPoint
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if o.DefaultWindowsDomainName != nil {
		toSerialize["defaultWindowsDomainName"] = o.DefaultWindowsDomainName
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LocalStorageName != nil {
		toSerialize["localStorageName"] = o.LocalStorageName
	}
	if o.MaximumNumCpus != nil {
		toSerialize["maximumNumCpus"] = o.MaximumNumCpus
	}
	if o.MaximumNumGpus != nil {
		toSerialize["maximumNumGpus"] = o.MaximumNumGpus
	}
	if o.MaximumRamInMegabytes != nil {
		toSerialize["maximumRamInMegabytes"] = o.MaximumRamInMegabytes
	}
	if o.MaximumStorageInMegabytes != nil {
		toSerialize["maximumStorageInMegabytes"] = o.MaximumStorageInMegabytes
	}
	if o.MaximumVirtualMachines != nil {
		toSerialize["maximumVirtualMachines"] = o.MaximumVirtualMachines
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.PowerOnDelayBase != nil {
		toSerialize["powerOnDelayBase"] = o.PowerOnDelayBase
	}
	if o.PowerOnInitialDelayBase != nil {
		toSerialize["powerOnInitialDelayBase"] = o.PowerOnInitialDelayBase
	}
	if o.PowerOnMaximumDelay != nil {
		toSerialize["powerOnMaximumDelay"] = o.PowerOnMaximumDelay
	}
	if o.PowerOnMinimumDelay != nil {
		toSerialize["powerOnMinimumDelay"] = o.PowerOnMinimumDelay
	}
	if o.RemoteAccessConfig != nil {
		toSerialize["remoteAccessConfig"] = o.RemoteAccessConfig
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.ZoneCount != nil {
		toSerialize["zoneCount"] = o.ZoneCount
	}
	return json.Marshal(toSerialize)
}

type NullableInputVirtualizationRealm struct {
	value *InputVirtualizationRealm
	isSet bool
}

func (v NullableInputVirtualizationRealm) Get() *InputVirtualizationRealm {
	return v.value
}

func (v *NullableInputVirtualizationRealm) Set(val *InputVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVirtualizationRealm(val *InputVirtualizationRealm) *NullableInputVirtualizationRealm {
	return &NullableInputVirtualizationRealm{value: val, isSet: true}
}

func (v NullableInputVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


