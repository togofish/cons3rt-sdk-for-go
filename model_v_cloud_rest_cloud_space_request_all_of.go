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

// VCloudRestCloudSpaceRequestAllOf struct for VCloudRestCloudSpaceRequestAllOf
type VCloudRestCloudSpaceRequestAllOf struct {
	ProviderVdcName     *string `json:"providerVdcName,omitempty"`
	VdcNetworkPoolName  *string `json:"vdcNetworkPoolName,omitempty"`
	Username            *string `json:"username,omitempty"`
	Password            *string `json:"password,omitempty"`
	EmailAddress        *string `json:"emailAddress,omitempty"`
	ExternalNetworkName *string `json:"externalNetworkName,omitempty"`
	LocalCatalogName    *string `json:"localCatalogName,omitempty"`
	VdcNetworkQuota     *int32  `json:"vdcNetworkQuota,omitempty"`
}

// NewVCloudRestCloudSpaceRequestAllOf instantiates a new VCloudRestCloudSpaceRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCloudRestCloudSpaceRequestAllOf() *VCloudRestCloudSpaceRequestAllOf {
	this := VCloudRestCloudSpaceRequestAllOf{}
	return &this
}

// NewVCloudRestCloudSpaceRequestAllOfWithDefaults instantiates a new VCloudRestCloudSpaceRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCloudRestCloudSpaceRequestAllOfWithDefaults() *VCloudRestCloudSpaceRequestAllOf {
	this := VCloudRestCloudSpaceRequestAllOf{}
	return &this
}

// GetProviderVdcName returns the ProviderVdcName field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetProviderVdcName() string {
	if o == nil || o.ProviderVdcName == nil {
		var ret string
		return ret
	}
	return *o.ProviderVdcName
}

// GetProviderVdcNameOk returns a tuple with the ProviderVdcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetProviderVdcNameOk() (*string, bool) {
	if o == nil || o.ProviderVdcName == nil {
		return nil, false
	}
	return o.ProviderVdcName, true
}

// HasProviderVdcName returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasProviderVdcName() bool {
	if o != nil && o.ProviderVdcName != nil {
		return true
	}

	return false
}

// SetProviderVdcName gets a reference to the given string and assigns it to the ProviderVdcName field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetProviderVdcName(v string) {
	o.ProviderVdcName = &v
}

// GetVdcNetworkPoolName returns the VdcNetworkPoolName field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetVdcNetworkPoolName() string {
	if o == nil || o.VdcNetworkPoolName == nil {
		var ret string
		return ret
	}
	return *o.VdcNetworkPoolName
}

// GetVdcNetworkPoolNameOk returns a tuple with the VdcNetworkPoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetVdcNetworkPoolNameOk() (*string, bool) {
	if o == nil || o.VdcNetworkPoolName == nil {
		return nil, false
	}
	return o.VdcNetworkPoolName, true
}

// HasVdcNetworkPoolName returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasVdcNetworkPoolName() bool {
	if o != nil && o.VdcNetworkPoolName != nil {
		return true
	}

	return false
}

// SetVdcNetworkPoolName gets a reference to the given string and assigns it to the VdcNetworkPoolName field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetVdcNetworkPoolName(v string) {
	o.VdcNetworkPoolName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetExternalNetworkName returns the ExternalNetworkName field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetExternalNetworkName() string {
	if o == nil || o.ExternalNetworkName == nil {
		var ret string
		return ret
	}
	return *o.ExternalNetworkName
}

// GetExternalNetworkNameOk returns a tuple with the ExternalNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetExternalNetworkNameOk() (*string, bool) {
	if o == nil || o.ExternalNetworkName == nil {
		return nil, false
	}
	return o.ExternalNetworkName, true
}

// HasExternalNetworkName returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasExternalNetworkName() bool {
	if o != nil && o.ExternalNetworkName != nil {
		return true
	}

	return false
}

// SetExternalNetworkName gets a reference to the given string and assigns it to the ExternalNetworkName field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetExternalNetworkName(v string) {
	o.ExternalNetworkName = &v
}

// GetLocalCatalogName returns the LocalCatalogName field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetLocalCatalogName() string {
	if o == nil || o.LocalCatalogName == nil {
		var ret string
		return ret
	}
	return *o.LocalCatalogName
}

// GetLocalCatalogNameOk returns a tuple with the LocalCatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetLocalCatalogNameOk() (*string, bool) {
	if o == nil || o.LocalCatalogName == nil {
		return nil, false
	}
	return o.LocalCatalogName, true
}

// HasLocalCatalogName returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasLocalCatalogName() bool {
	if o != nil && o.LocalCatalogName != nil {
		return true
	}

	return false
}

// SetLocalCatalogName gets a reference to the given string and assigns it to the LocalCatalogName field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetLocalCatalogName(v string) {
	o.LocalCatalogName = &v
}

// GetVdcNetworkQuota returns the VdcNetworkQuota field value if set, zero value otherwise.
func (o *VCloudRestCloudSpaceRequestAllOf) GetVdcNetworkQuota() int32 {
	if o == nil || o.VdcNetworkQuota == nil {
		var ret int32
		return ret
	}
	return *o.VdcNetworkQuota
}

// GetVdcNetworkQuotaOk returns a tuple with the VdcNetworkQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) GetVdcNetworkQuotaOk() (*int32, bool) {
	if o == nil || o.VdcNetworkQuota == nil {
		return nil, false
	}
	return o.VdcNetworkQuota, true
}

// HasVdcNetworkQuota returns a boolean if a field has been set.
func (o *VCloudRestCloudSpaceRequestAllOf) HasVdcNetworkQuota() bool {
	if o != nil && o.VdcNetworkQuota != nil {
		return true
	}

	return false
}

// SetVdcNetworkQuota gets a reference to the given int32 and assigns it to the VdcNetworkQuota field.
func (o *VCloudRestCloudSpaceRequestAllOf) SetVdcNetworkQuota(v int32) {
	o.VdcNetworkQuota = &v
}

func (o VCloudRestCloudSpaceRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProviderVdcName != nil {
		toSerialize["providerVdcName"] = o.ProviderVdcName
	}
	if o.VdcNetworkPoolName != nil {
		toSerialize["vdcNetworkPoolName"] = o.VdcNetworkPoolName
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.ExternalNetworkName != nil {
		toSerialize["externalNetworkName"] = o.ExternalNetworkName
	}
	if o.LocalCatalogName != nil {
		toSerialize["localCatalogName"] = o.LocalCatalogName
	}
	if o.VdcNetworkQuota != nil {
		toSerialize["vdcNetworkQuota"] = o.VdcNetworkQuota
	}
	return json.Marshal(toSerialize)
}

type NullableVCloudRestCloudSpaceRequestAllOf struct {
	value *VCloudRestCloudSpaceRequestAllOf
	isSet bool
}

func (v NullableVCloudRestCloudSpaceRequestAllOf) Get() *VCloudRestCloudSpaceRequestAllOf {
	return v.value
}

func (v *NullableVCloudRestCloudSpaceRequestAllOf) Set(val *VCloudRestCloudSpaceRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVCloudRestCloudSpaceRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVCloudRestCloudSpaceRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCloudRestCloudSpaceRequestAllOf(val *VCloudRestCloudSpaceRequestAllOf) *NullableVCloudRestCloudSpaceRequestAllOf {
	return &NullableVCloudRestCloudSpaceRequestAllOf{value: val, isSet: true}
}

func (v NullableVCloudRestCloudSpaceRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCloudRestCloudSpaceRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
