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

// InputOpenStackCloud struct for InputOpenStackCloud
type InputOpenStackCloud struct {
	DomainName         *string `json:"domainName,omitempty"`
	KeystoneHostname   string  `json:"keystoneHostname"`
	KeystonePassword   string  `json:"keystonePassword"`
	KeystonePort       int32   `json:"keystonePort"`
	KeystoneProtocol   string  `json:"keystoneProtocol"`
	KeystoneUsername   string  `json:"keystoneUsername"`
	KeystoneVersion    string  `json:"keystoneVersion"`
	NatImageId         string  `json:"natImageId"`
	NatInstanceType    string  `json:"natInstanceType"`
	Tenant             string  `json:"tenant"`
	TenantId           string  `json:"tenantId"`
	Name               string
	ExternalIpSource   string
	MaximumImpactLevel string
	OwningTeam         InputTeam
	Subtype            string
}

// NewInputOpenStackCloud instantiates a new InputOpenStackCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputOpenStackCloud(keystoneHostname string, keystonePassword string, keystonePort int32, keystoneProtocol string, keystoneUsername string, keystoneVersion string, natImageId string, natInstanceType string, tenant string, tenantId string, name string, externalIpSource string, maximumImpactLevel string, owningTeam InputTeam, subtype string) *InputOpenStackCloud {
	this := InputOpenStackCloud{}
	this.Name = name
	this.ExternalIpSource = externalIpSource
	this.MaximumImpactLevel = maximumImpactLevel
	this.OwningTeam = owningTeam
	this.Subtype = subtype
	this.KeystoneHostname = keystoneHostname
	this.KeystonePassword = keystonePassword
	this.KeystonePort = keystonePort
	this.KeystoneProtocol = keystoneProtocol
	this.KeystoneUsername = keystoneUsername
	this.KeystoneVersion = keystoneVersion
	this.NatImageId = natImageId
	this.NatInstanceType = natInstanceType
	this.Tenant = tenant
	this.TenantId = tenantId
	return &this
}

// NewInputOpenStackCloudWithDefaults instantiates a new InputOpenStackCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputOpenStackCloudWithDefaults() *InputOpenStackCloud {
	this := InputOpenStackCloud{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *InputOpenStackCloud) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *InputOpenStackCloud) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *InputOpenStackCloud) SetDomainName(v string) {
	o.DomainName = &v
}

// GetKeystoneHostname returns the KeystoneHostname field value
func (o *InputOpenStackCloud) GetKeystoneHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeystoneHostname
}

// GetKeystoneHostnameOk returns a tuple with the KeystoneHostname field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetKeystoneHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeystoneHostname, true
}

// SetKeystoneHostname sets field value
func (o *InputOpenStackCloud) SetKeystoneHostname(v string) {
	o.KeystoneHostname = v
}

// GetKeystonePassword returns the KeystonePassword field value
func (o *InputOpenStackCloud) GetKeystonePassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeystonePassword
}

// GetKeystonePasswordOk returns a tuple with the KeystonePassword field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetKeystonePasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeystonePassword, true
}

// SetKeystonePassword sets field value
func (o *InputOpenStackCloud) SetKeystonePassword(v string) {
	o.KeystonePassword = v
}

// GetKeystonePort returns the KeystonePort field value
func (o *InputOpenStackCloud) GetKeystonePort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KeystonePort
}

// GetKeystonePortOk returns a tuple with the KeystonePort field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetKeystonePortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeystonePort, true
}

// SetKeystonePort sets field value
func (o *InputOpenStackCloud) SetKeystonePort(v int32) {
	o.KeystonePort = v
}

// GetKeystoneProtocol returns the KeystoneProtocol field value
func (o *InputOpenStackCloud) GetKeystoneProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeystoneProtocol
}

// GetKeystoneProtocolOk returns a tuple with the KeystoneProtocol field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetKeystoneProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeystoneProtocol, true
}

// SetKeystoneProtocol sets field value
func (o *InputOpenStackCloud) SetKeystoneProtocol(v string) {
	o.KeystoneProtocol = v
}

// GetKeystoneUsername returns the KeystoneUsername field value
func (o *InputOpenStackCloud) GetKeystoneUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeystoneUsername
}

// GetKeystoneUsernameOk returns a tuple with the KeystoneUsername field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetKeystoneUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeystoneUsername, true
}

// SetKeystoneUsername sets field value
func (o *InputOpenStackCloud) SetKeystoneUsername(v string) {
	o.KeystoneUsername = v
}

// GetKeystoneVersion returns the KeystoneVersion field value
func (o *InputOpenStackCloud) GetKeystoneVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeystoneVersion
}

// GetKeystoneVersionOk returns a tuple with the KeystoneVersion field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetKeystoneVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeystoneVersion, true
}

// SetKeystoneVersion sets field value
func (o *InputOpenStackCloud) SetKeystoneVersion(v string) {
	o.KeystoneVersion = v
}

// GetNatImageId returns the NatImageId field value
func (o *InputOpenStackCloud) GetNatImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NatImageId
}

// GetNatImageIdOk returns a tuple with the NatImageId field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetNatImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NatImageId, true
}

// SetNatImageId sets field value
func (o *InputOpenStackCloud) SetNatImageId(v string) {
	o.NatImageId = v
}

// GetNatInstanceType returns the NatInstanceType field value
func (o *InputOpenStackCloud) GetNatInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NatInstanceType
}

// GetNatInstanceTypeOk returns a tuple with the NatInstanceType field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetNatInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NatInstanceType, true
}

// SetNatInstanceType sets field value
func (o *InputOpenStackCloud) SetNatInstanceType(v string) {
	o.NatInstanceType = v
}

// GetTenant returns the Tenant field value
func (o *InputOpenStackCloud) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *InputOpenStackCloud) SetTenant(v string) {
	o.Tenant = v
}

// GetTenantId returns the TenantId field value
func (o *InputOpenStackCloud) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *InputOpenStackCloud) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *InputOpenStackCloud) SetTenantId(v string) {
	o.TenantId = v
}

func (o InputOpenStackCloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainName != nil {
		toSerialize["domainName"] = o.DomainName
	}
	if true {
		toSerialize["keystoneHostname"] = o.KeystoneHostname
	}
	if true {
		toSerialize["keystonePassword"] = o.KeystonePassword
	}
	if true {
		toSerialize["keystonePort"] = o.KeystonePort
	}
	if true {
		toSerialize["keystoneProtocol"] = o.KeystoneProtocol
	}
	if true {
		toSerialize["keystoneUsername"] = o.KeystoneUsername
	}
	if true {
		toSerialize["keystoneVersion"] = o.KeystoneVersion
	}
	if true {
		toSerialize["natImageId"] = o.NatImageId
	}
	if true {
		toSerialize["natInstanceType"] = o.NatInstanceType
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	return json.Marshal(toSerialize)
}

type NullableInputOpenStackCloud struct {
	value *InputOpenStackCloud
	isSet bool
}

func (v NullableInputOpenStackCloud) Get() *InputOpenStackCloud {
	return v.value
}

func (v *NullableInputOpenStackCloud) Set(val *InputOpenStackCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableInputOpenStackCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableInputOpenStackCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputOpenStackCloud(val *InputOpenStackCloud) *NullableInputOpenStackCloud {
	return &NullableInputOpenStackCloud{value: val, isSet: true}
}

func (v NullableInputOpenStackCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputOpenStackCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
