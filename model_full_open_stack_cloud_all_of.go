/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FullOpenStackCloudAllOf struct for FullOpenStackCloudAllOf
type FullOpenStackCloudAllOf struct {
	DomainName       *string `json:"domainName,omitempty"`
	KeystoneHostname *string `json:"keystoneHostname,omitempty"`
	KeystonePort     *int32  `json:"keystonePort,omitempty"`
	KeystoneProtocol *string `json:"keystoneProtocol,omitempty"`
	KeystoneUsername *string `json:"keystoneUsername,omitempty"`
	KeystoneVersion  *string `json:"keystoneVersion,omitempty"`
	NatImageId       *string `json:"natImageId,omitempty"`
	NatInstanceType  *string `json:"natInstanceType,omitempty"`
	Tenant           *string `json:"tenant,omitempty"`
	TenantId         *string `json:"tenantId,omitempty"`
}

// NewFullOpenStackCloudAllOf instantiates a new FullOpenStackCloudAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullOpenStackCloudAllOf() *FullOpenStackCloudAllOf {
	this := FullOpenStackCloudAllOf{}
	return &this
}

// NewFullOpenStackCloudAllOfWithDefaults instantiates a new FullOpenStackCloudAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullOpenStackCloudAllOfWithDefaults() *FullOpenStackCloudAllOf {
	this := FullOpenStackCloudAllOf{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *FullOpenStackCloudAllOf) SetDomainName(v string) {
	o.DomainName = &v
}

// GetKeystoneHostname returns the KeystoneHostname field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetKeystoneHostname() string {
	if o == nil || o.KeystoneHostname == nil {
		var ret string
		return ret
	}
	return *o.KeystoneHostname
}

// GetKeystoneHostnameOk returns a tuple with the KeystoneHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetKeystoneHostnameOk() (*string, bool) {
	if o == nil || o.KeystoneHostname == nil {
		return nil, false
	}
	return o.KeystoneHostname, true
}

// HasKeystoneHostname returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasKeystoneHostname() bool {
	if o != nil && o.KeystoneHostname != nil {
		return true
	}

	return false
}

// SetKeystoneHostname gets a reference to the given string and assigns it to the KeystoneHostname field.
func (o *FullOpenStackCloudAllOf) SetKeystoneHostname(v string) {
	o.KeystoneHostname = &v
}

// GetKeystonePort returns the KeystonePort field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetKeystonePort() int32 {
	if o == nil || o.KeystonePort == nil {
		var ret int32
		return ret
	}
	return *o.KeystonePort
}

// GetKeystonePortOk returns a tuple with the KeystonePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetKeystonePortOk() (*int32, bool) {
	if o == nil || o.KeystonePort == nil {
		return nil, false
	}
	return o.KeystonePort, true
}

// HasKeystonePort returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasKeystonePort() bool {
	if o != nil && o.KeystonePort != nil {
		return true
	}

	return false
}

// SetKeystonePort gets a reference to the given int32 and assigns it to the KeystonePort field.
func (o *FullOpenStackCloudAllOf) SetKeystonePort(v int32) {
	o.KeystonePort = &v
}

// GetKeystoneProtocol returns the KeystoneProtocol field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetKeystoneProtocol() string {
	if o == nil || o.KeystoneProtocol == nil {
		var ret string
		return ret
	}
	return *o.KeystoneProtocol
}

// GetKeystoneProtocolOk returns a tuple with the KeystoneProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetKeystoneProtocolOk() (*string, bool) {
	if o == nil || o.KeystoneProtocol == nil {
		return nil, false
	}
	return o.KeystoneProtocol, true
}

// HasKeystoneProtocol returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasKeystoneProtocol() bool {
	if o != nil && o.KeystoneProtocol != nil {
		return true
	}

	return false
}

// SetKeystoneProtocol gets a reference to the given string and assigns it to the KeystoneProtocol field.
func (o *FullOpenStackCloudAllOf) SetKeystoneProtocol(v string) {
	o.KeystoneProtocol = &v
}

// GetKeystoneUsername returns the KeystoneUsername field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetKeystoneUsername() string {
	if o == nil || o.KeystoneUsername == nil {
		var ret string
		return ret
	}
	return *o.KeystoneUsername
}

// GetKeystoneUsernameOk returns a tuple with the KeystoneUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetKeystoneUsernameOk() (*string, bool) {
	if o == nil || o.KeystoneUsername == nil {
		return nil, false
	}
	return o.KeystoneUsername, true
}

// HasKeystoneUsername returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasKeystoneUsername() bool {
	if o != nil && o.KeystoneUsername != nil {
		return true
	}

	return false
}

// SetKeystoneUsername gets a reference to the given string and assigns it to the KeystoneUsername field.
func (o *FullOpenStackCloudAllOf) SetKeystoneUsername(v string) {
	o.KeystoneUsername = &v
}

// GetKeystoneVersion returns the KeystoneVersion field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetKeystoneVersion() string {
	if o == nil || o.KeystoneVersion == nil {
		var ret string
		return ret
	}
	return *o.KeystoneVersion
}

// GetKeystoneVersionOk returns a tuple with the KeystoneVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetKeystoneVersionOk() (*string, bool) {
	if o == nil || o.KeystoneVersion == nil {
		return nil, false
	}
	return o.KeystoneVersion, true
}

// HasKeystoneVersion returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasKeystoneVersion() bool {
	if o != nil && o.KeystoneVersion != nil {
		return true
	}

	return false
}

// SetKeystoneVersion gets a reference to the given string and assigns it to the KeystoneVersion field.
func (o *FullOpenStackCloudAllOf) SetKeystoneVersion(v string) {
	o.KeystoneVersion = &v
}

// GetNatImageId returns the NatImageId field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetNatImageId() string {
	if o == nil || o.NatImageId == nil {
		var ret string
		return ret
	}
	return *o.NatImageId
}

// GetNatImageIdOk returns a tuple with the NatImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetNatImageIdOk() (*string, bool) {
	if o == nil || o.NatImageId == nil {
		return nil, false
	}
	return o.NatImageId, true
}

// HasNatImageId returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasNatImageId() bool {
	if o != nil && o.NatImageId != nil {
		return true
	}

	return false
}

// SetNatImageId gets a reference to the given string and assigns it to the NatImageId field.
func (o *FullOpenStackCloudAllOf) SetNatImageId(v string) {
	o.NatImageId = &v
}

// GetNatInstanceType returns the NatInstanceType field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetNatInstanceType() string {
	if o == nil || o.NatInstanceType == nil {
		var ret string
		return ret
	}
	return *o.NatInstanceType
}

// GetNatInstanceTypeOk returns a tuple with the NatInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetNatInstanceTypeOk() (*string, bool) {
	if o == nil || o.NatInstanceType == nil {
		return nil, false
	}
	return o.NatInstanceType, true
}

// HasNatInstanceType returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasNatInstanceType() bool {
	if o != nil && o.NatInstanceType != nil {
		return true
	}

	return false
}

// SetNatInstanceType gets a reference to the given string and assigns it to the NatInstanceType field.
func (o *FullOpenStackCloudAllOf) SetNatInstanceType(v string) {
	o.NatInstanceType = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *FullOpenStackCloudAllOf) SetTenant(v string) {
	o.Tenant = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *FullOpenStackCloudAllOf) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOpenStackCloudAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *FullOpenStackCloudAllOf) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *FullOpenStackCloudAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

func (o FullOpenStackCloudAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainName != nil {
		toSerialize["domainName"] = o.DomainName
	}
	if o.KeystoneHostname != nil {
		toSerialize["keystoneHostname"] = o.KeystoneHostname
	}
	if o.KeystonePort != nil {
		toSerialize["keystonePort"] = o.KeystonePort
	}
	if o.KeystoneProtocol != nil {
		toSerialize["keystoneProtocol"] = o.KeystoneProtocol
	}
	if o.KeystoneUsername != nil {
		toSerialize["keystoneUsername"] = o.KeystoneUsername
	}
	if o.KeystoneVersion != nil {
		toSerialize["keystoneVersion"] = o.KeystoneVersion
	}
	if o.NatImageId != nil {
		toSerialize["natImageId"] = o.NatImageId
	}
	if o.NatInstanceType != nil {
		toSerialize["natInstanceType"] = o.NatInstanceType
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	return json.Marshal(toSerialize)
}

type NullableFullOpenStackCloudAllOf struct {
	value *FullOpenStackCloudAllOf
	isSet bool
}

func (v NullableFullOpenStackCloudAllOf) Get() *FullOpenStackCloudAllOf {
	return v.value
}

func (v *NullableFullOpenStackCloudAllOf) Set(val *FullOpenStackCloudAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullOpenStackCloudAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullOpenStackCloudAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullOpenStackCloudAllOf(val *FullOpenStackCloudAllOf) *NullableFullOpenStackCloudAllOf {
	return &NullableFullOpenStackCloudAllOf{value: val, isSet: true}
}

func (v NullableFullOpenStackCloudAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullOpenStackCloudAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
