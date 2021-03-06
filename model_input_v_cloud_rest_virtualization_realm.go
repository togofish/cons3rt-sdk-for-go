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

// InputVCloudRestVirtualizationRealm struct for InputVCloudRestVirtualizationRealm
type InputVCloudRestVirtualizationRealm struct {
	Organization    string `json:"organization"`
	OrganizationVdc string `json:"organizationVdc"`
	VdcNetworkQuota int32  `json:"vdcNetworkQuota"`
	AccountId       string
	Cidr            string
	Description     string
	Name            string
	Password        string
	Username        string
}

// NewInputVCloudRestVirtualizationRealm instantiates a new InputVCloudRestVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputVCloudRestVirtualizationRealm(organization string, organizationVdc string, vdcNetworkQuota int32, accountId string, cidr string, description string, name string, password string, username string) *InputVCloudRestVirtualizationRealm {
	this := InputVCloudRestVirtualizationRealm{}
	this.AccountId = accountId
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	this.Password = password
	this.Username = username
	this.Organization = organization
	this.OrganizationVdc = organizationVdc
	this.VdcNetworkQuota = vdcNetworkQuota
	return &this
}

// NewInputVCloudRestVirtualizationRealmWithDefaults instantiates a new InputVCloudRestVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputVCloudRestVirtualizationRealmWithDefaults() *InputVCloudRestVirtualizationRealm {
	this := InputVCloudRestVirtualizationRealm{}
	return &this
}

// GetOrganization returns the Organization field value
func (o *InputVCloudRestVirtualizationRealm) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *InputVCloudRestVirtualizationRealm) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *InputVCloudRestVirtualizationRealm) SetOrganization(v string) {
	o.Organization = v
}

// GetOrganizationVdc returns the OrganizationVdc field value
func (o *InputVCloudRestVirtualizationRealm) GetOrganizationVdc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationVdc
}

// GetOrganizationVdcOk returns a tuple with the OrganizationVdc field value
// and a boolean to check if the value has been set.
func (o *InputVCloudRestVirtualizationRealm) GetOrganizationVdcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationVdc, true
}

// SetOrganizationVdc sets field value
func (o *InputVCloudRestVirtualizationRealm) SetOrganizationVdc(v string) {
	o.OrganizationVdc = v
}

// GetVdcNetworkQuota returns the VdcNetworkQuota field value
func (o *InputVCloudRestVirtualizationRealm) GetVdcNetworkQuota() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VdcNetworkQuota
}

// GetVdcNetworkQuotaOk returns a tuple with the VdcNetworkQuota field value
// and a boolean to check if the value has been set.
func (o *InputVCloudRestVirtualizationRealm) GetVdcNetworkQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VdcNetworkQuota, true
}

// SetVdcNetworkQuota sets field value
func (o *InputVCloudRestVirtualizationRealm) SetVdcNetworkQuota(v int32) {
	o.VdcNetworkQuota = v
}

func (o InputVCloudRestVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["organizationVdc"] = o.OrganizationVdc
	}
	if true {
		toSerialize["vdcNetworkQuota"] = o.VdcNetworkQuota
	}
	return json.Marshal(toSerialize)
}

type NullableInputVCloudRestVirtualizationRealm struct {
	value *InputVCloudRestVirtualizationRealm
	isSet bool
}

func (v NullableInputVCloudRestVirtualizationRealm) Get() *InputVCloudRestVirtualizationRealm {
	return v.value
}

func (v *NullableInputVCloudRestVirtualizationRealm) Set(val *InputVCloudRestVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableInputVCloudRestVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableInputVCloudRestVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputVCloudRestVirtualizationRealm(val *InputVCloudRestVirtualizationRealm) *NullableInputVCloudRestVirtualizationRealm {
	return &NullableInputVCloudRestVirtualizationRealm{value: val, isSet: true}
}

func (v NullableInputVCloudRestVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputVCloudRestVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
