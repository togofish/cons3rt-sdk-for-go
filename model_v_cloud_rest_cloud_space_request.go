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

// VCloudRestCloudSpaceRequest struct for VCloudRestCloudSpaceRequest
type VCloudRestCloudSpaceRequest struct {
	ProviderVdcName     string  `json:"providerVdcName"`
	VdcNetworkPoolName  string  `json:"vdcNetworkPoolName"`
	Username            string  `json:"username"`
	Password            string  `json:"password"`
	EmailAddress        string  `json:"emailAddress"`
	ExternalNetworkName string  `json:"externalNetworkName"`
	LocalCatalogName    *string `json:"localCatalogName,omitempty"`
	VdcNetworkQuota     int32   `json:"vdcNetworkQuota"`
	CloudSpaceName      string
	Cidr                string
	Subtype             string
}

// NewVCloudRestCloudSpaceRequest instantiates a new VCloudRestCloudSpaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCloudRestCloudSpaceRequest(providerVdcName string, vdcNetworkPoolName string, username string, password string, emailAddress string, externalNetworkName string, vdcNetworkQuota int32, cloudSpaceName string, cidr string, subtype string) *VCloudRestCloudSpaceRequest {
	this := VCloudRestCloudSpaceRequest{}
	this.CloudSpaceName = cloudSpaceName
	this.Cidr = cidr
	this.Subtype = subtype
	this.ProviderVdcName = providerVdcName
	this.VdcNetworkPoolName = vdcNetworkPoolName
	this.Username = username
	this.Password = password
	this.EmailAddress = emailAddress
	this.ExternalNetworkName = externalNetworkName
	this.VdcNetworkQuota = vdcNetworkQuota
	return &this
}

// NewVCloudRestCloudSpaceRequestWithDefaults instantiates a new VCloudRestCloudSpaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCloudRestCloudSpaceRequestWithDefaults() *VCloudRestCloudSpaceRequest {
	this := VCloudRestCloudSpaceRequest{}
	return &this
}

// GetProviderVdcName returns the ProviderVdcName field value
func (o *VCloudRestCloudSpaceRequest) GetProviderVdcName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderVdcName
}

// GetProviderVdcNameOk returns a tuple with the ProviderVdcName field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetProviderVdcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderVdcName, true
}

// SetProviderVdcName sets field value
func (o *VCloudRestCloudSpaceRequest) SetProviderVdcName(v string) {
	o.ProviderVdcName = v
}

// GetVdcNetworkPoolName returns the VdcNetworkPoolName field value
func (o *VCloudRestCloudSpaceRequest) GetVdcNetworkPoolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VdcNetworkPoolName
}

// GetVdcNetworkPoolNameOk returns a tuple with the VdcNetworkPoolName field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetVdcNetworkPoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VdcNetworkPoolName, true
}

// SetVdcNetworkPoolName sets field value
func (o *VCloudRestCloudSpaceRequest) SetVdcNetworkPoolName(v string) {
	o.VdcNetworkPoolName = v
}

// GetUsername returns the Username field value
func (o *VCloudRestCloudSpaceRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VCloudRestCloudSpaceRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *VCloudRestCloudSpaceRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VCloudRestCloudSpaceRequest) SetPassword(v string) {
	o.Password = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *VCloudRestCloudSpaceRequest) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *VCloudRestCloudSpaceRequest) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetExternalNetworkName returns the ExternalNetworkName field value
func (o *VCloudRestCloudSpaceRequest) GetExternalNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalNetworkName
}

// GetExternalNetworkNameOk returns a tuple with the ExternalNetworkName field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetExternalNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalNetworkName, true
}

// SetExternalNetworkName sets field value
func (o *VCloudRestCloudSpaceRequest) SetExternalNetworkName(v string) {
	o.ExternalNetworkName = v
}

// GetLocalCatalogName returns the LocalCatalogName field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequest) GetLocalCatalogName() string {
	if o == nil || o.LocalCatalogName == nil {
		var ret string
		return ret
	}
	return *o.LocalCatalogName
}

// GetLocalCatalogNameOk returns a tuple with the LocalCatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetLocalCatalogNameOk() (*string, bool) {
	if o == nil || o.LocalCatalogName == nil {
		return nil, false
	}
	return o.LocalCatalogName, true
}

// HasLocalCatalogName returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequest) HasLocalCatalogName() bool {
	if o != nil && o.LocalCatalogName != nil {
		return true
	}

	return false
}

// SetLocalCatalogName gets a reference to the given string and assigns it to the LocalCatalogName field.
func (o *VCloudRestCloudSpaceRequest) SetLocalCatalogName(v string) {
	o.LocalCatalogName = &v
}

// GetVdcNetworkQuota returns the VdcNetworkQuota field value
func (o *VCloudRestCloudSpaceRequest) GetVdcNetworkQuota() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VdcNetworkQuota
}

// GetVdcNetworkQuotaOk returns a tuple with the VdcNetworkQuota field value
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequest) GetVdcNetworkQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VdcNetworkQuota, true
}

// SetVdcNetworkQuota sets field value
func (o *VCloudRestCloudSpaceRequest) SetVdcNetworkQuota(v int32) {
	o.VdcNetworkQuota = v
}

func (o VCloudRestCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerVdcName"] = o.ProviderVdcName
	}
	if true {
		toSerialize["vdcNetworkPoolName"] = o.VdcNetworkPoolName
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if true {
		toSerialize["externalNetworkName"] = o.ExternalNetworkName
	}
	if o.LocalCatalogName != nil {
		toSerialize["localCatalogName"] = o.LocalCatalogName
	}
	if true {
		toSerialize["vdcNetworkQuota"] = o.VdcNetworkQuota
	}
	return json.Marshal(toSerialize)
}

type NullableVCloudRestCloudSpaceRequest struct {
	value *VCloudRestCloudSpaceRequest
	isSet bool
}

func (v NullableVCloudRestCloudSpaceRequest) Get() *VCloudRestCloudSpaceRequest {
	return v.value
}

func (v *NullableVCloudRestCloudSpaceRequest) Set(val *VCloudRestCloudSpaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVCloudRestCloudSpaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVCloudRestCloudSpaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCloudRestCloudSpaceRequest(val *VCloudRestCloudSpaceRequest) *NullableVCloudRestCloudSpaceRequest {
	return &NullableVCloudRestCloudSpaceRequest{value: val, isSet: true}
}

func (v NullableVCloudRestCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCloudRestCloudSpaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
