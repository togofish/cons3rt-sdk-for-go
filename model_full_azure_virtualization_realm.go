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

// FullAzureVirtualizationRealm struct for FullAzureVirtualizationRealm
type FullAzureVirtualizationRealm struct {
	Environment        string  `json:"environment"`
	PublicContainerUrl *string `json:"publicContainerUrl,omitempty"`
	Region             string  `json:"region"`
	ResourceGroupName  string  `json:"resourceGroupName"`
	TenantId           string  `json:"tenantId"`
	VirtualNetworkName string  `json:"virtualNetworkName"`
}

// NewFullAzureVirtualizationRealm instantiates a new FullAzureVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullAzureVirtualizationRealm(environment string, region string, resourceGroupName string, tenantId string, virtualNetworkName string, name string, accountId string, networks []Network, description string, username string) *FullAzureVirtualizationRealm {
	this := FullAzureVirtualizationRealm{}
	this.Name = name
	this.AccountId = accountId
	this.Networks = networks
	this.Description = description
	this.Username = username
	this.Environment = environment
	this.Region = region
	this.ResourceGroupName = resourceGroupName
	this.TenantId = tenantId
	this.VirtualNetworkName = virtualNetworkName
	return &this
}

// NewFullAzureVirtualizationRealmWithDefaults instantiates a new FullAzureVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullAzureVirtualizationRealmWithDefaults() *FullAzureVirtualizationRealm {
	this := FullAzureVirtualizationRealm{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *FullAzureVirtualizationRealm) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *FullAzureVirtualizationRealm) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *FullAzureVirtualizationRealm) SetEnvironment(v string) {
	o.Environment = v
}

// GetPublicContainerUrl returns the PublicContainerUrl field value if set, zero value otherwise.
func (o *FullAzureVirtualizationRealm) GetPublicContainerUrl() string {
	if o == nil || o.PublicContainerUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicContainerUrl
}

// GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureVirtualizationRealm) GetPublicContainerUrlOk() (*string, bool) {
	if o == nil || o.PublicContainerUrl == nil {
		return nil, false
	}
	return o.PublicContainerUrl, true
}

// HasPublicContainerUrl returns a boolean if a field has been set.
func (o *FullAzureVirtualizationRealm) HasPublicContainerUrl() bool {
	if o != nil && o.PublicContainerUrl != nil {
		return true
	}

	return false
}

// SetPublicContainerUrl gets a reference to the given string and assigns it to the PublicContainerUrl field.
func (o *FullAzureVirtualizationRealm) SetPublicContainerUrl(v string) {
	o.PublicContainerUrl = &v
}

// GetRegion returns the Region field value
func (o *FullAzureVirtualizationRealm) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *FullAzureVirtualizationRealm) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *FullAzureVirtualizationRealm) SetRegion(v string) {
	o.Region = v
}

// GetResourceGroupName returns the ResourceGroupName field value
func (o *FullAzureVirtualizationRealm) GetResourceGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value
// and a boolean to check if the value has been set.
func (o *FullAzureVirtualizationRealm) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroupName, true
}

// SetResourceGroupName sets field value
func (o *FullAzureVirtualizationRealm) SetResourceGroupName(v string) {
	o.ResourceGroupName = v
}

// GetTenantId returns the TenantId field value
func (o *FullAzureVirtualizationRealm) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *FullAzureVirtualizationRealm) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *FullAzureVirtualizationRealm) SetTenantId(v string) {
	o.TenantId = v
}

// GetVirtualNetworkName returns the VirtualNetworkName field value
func (o *FullAzureVirtualizationRealm) GetVirtualNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualNetworkName
}

// GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field value
// and a boolean to check if the value has been set.
func (o *FullAzureVirtualizationRealm) GetVirtualNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetworkName, true
}

// SetVirtualNetworkName sets field value
func (o *FullAzureVirtualizationRealm) SetVirtualNetworkName(v string) {
	o.VirtualNetworkName = v
}

func (o FullAzureVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if o.PublicContainerUrl != nil {
		toSerialize["publicContainerUrl"] = o.PublicContainerUrl
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["resourceGroupName"] = o.ResourceGroupName
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["virtualNetworkName"] = o.VirtualNetworkName
	}
	return json.Marshal(toSerialize)
}

type NullableFullAzureVirtualizationRealm struct {
	value *FullAzureVirtualizationRealm
	isSet bool
}

func (v NullableFullAzureVirtualizationRealm) Get() *FullAzureVirtualizationRealm {
	return v.value
}

func (v *NullableFullAzureVirtualizationRealm) Set(val *FullAzureVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableFullAzureVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableFullAzureVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullAzureVirtualizationRealm(val *FullAzureVirtualizationRealm) *NullableFullAzureVirtualizationRealm {
	return &NullableFullAzureVirtualizationRealm{value: val, isSet: true}
}

func (v NullableFullAzureVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullAzureVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
