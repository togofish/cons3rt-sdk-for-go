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

// OpenStackClient struct for OpenStackClient
type OpenStackClient struct {
	KeystoneVersion *string `json:"keystoneVersion,omitempty"`
	TenantName      *string `json:"tenantName,omitempty"`
	Username        string
	Password        string
	Subtype         string
}

// NewOpenStackClient instantiates a new OpenStackClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackClient(username string, password string, subtype string) *OpenStackClient {
	this := OpenStackClient{}
	this.Username = username
	this.Password = password
	this.Subtype = subtype
	return &this
}

// NewOpenStackClientWithDefaults instantiates a new OpenStackClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackClientWithDefaults() *OpenStackClient {
	this := OpenStackClient{}
	return &this
}

// GetKeystoneVersion returns the KeystoneVersion field value if set, zero value otherwise.
func (o *OpenStackClient) GetKeystoneVersion() string {
	if o == nil || o.KeystoneVersion == nil {
		var ret string
		return ret
	}
	return *o.KeystoneVersion
}

// GetKeystoneVersionOk returns a tuple with the KeystoneVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackClient) GetKeystoneVersionOk() (*string, bool) {
	if o == nil || o.KeystoneVersion == nil {
		return nil, false
	}
	return o.KeystoneVersion, true
}

// HasKeystoneVersion returns a boolean if a field has been set.
func (o *OpenStackClient) HasKeystoneVersion() bool {
	if o != nil && o.KeystoneVersion != nil {
		return true
	}

	return false
}

// SetKeystoneVersion gets a reference to the given string and assigns it to the KeystoneVersion field.
func (o *OpenStackClient) SetKeystoneVersion(v string) {
	o.KeystoneVersion = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *OpenStackClient) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackClient) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *OpenStackClient) HasTenantName() bool {
	if o != nil && o.TenantName != nil {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *OpenStackClient) SetTenantName(v string) {
	o.TenantName = &v
}

func (o OpenStackClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeystoneVersion != nil {
		toSerialize["keystoneVersion"] = o.KeystoneVersion
	}
	if o.TenantName != nil {
		toSerialize["tenantName"] = o.TenantName
	}
	return json.Marshal(toSerialize)
}

type NullableOpenStackClient struct {
	value *OpenStackClient
	isSet bool
}

func (v NullableOpenStackClient) Get() *OpenStackClient {
	return v.value
}

func (v *NullableOpenStackClient) Set(val *OpenStackClient) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackClient) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackClient(val *OpenStackClient) *NullableOpenStackClient {
	return &NullableOpenStackClient{value: val, isSet: true}
}

func (v NullableOpenStackClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
