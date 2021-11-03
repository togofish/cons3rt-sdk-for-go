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

// FullAzureCloud struct for FullAzureCloud
type FullAzureCloud struct {
	ClientId           string  `json:"clientId"`
	Environment        string  `json:"environment"`
	RegionName         string  `json:"regionName"`
	SubscriptionId     string  `json:"subscriptionId"`
	Tenant             string  `json:"tenant"`
	PublicContainerUrl *string `json:"publicContainerUrl,omitempty"`
}

// NewFullAzureCloud instantiates a new FullAzureCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullAzureCloud(clientId string, environment string, regionName string, subscriptionId string, tenant string, name string, externalIpSource string, maximumImpactLevel string, subtype string) *FullAzureCloud {
	this := FullAzureCloud{}
	this.Name = name
	this.ExternalIpSource = externalIpSource
	this.MaximumImpactLevel = maximumImpactLevel
	this.Subtype = subtype
	this.ClientId = clientId
	this.Environment = environment
	this.RegionName = regionName
	this.SubscriptionId = subscriptionId
	this.Tenant = tenant
	return &this
}

// NewFullAzureCloudWithDefaults instantiates a new FullAzureCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullAzureCloudWithDefaults() *FullAzureCloud {
	this := FullAzureCloud{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *FullAzureCloud) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *FullAzureCloud) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *FullAzureCloud) SetClientId(v string) {
	o.ClientId = v
}

// GetEnvironment returns the Environment field value
func (o *FullAzureCloud) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *FullAzureCloud) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *FullAzureCloud) SetEnvironment(v string) {
	o.Environment = v
}

// GetRegionName returns the RegionName field value
func (o *FullAzureCloud) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *FullAzureCloud) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *FullAzureCloud) SetRegionName(v string) {
	o.RegionName = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *FullAzureCloud) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *FullAzureCloud) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *FullAzureCloud) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetTenant returns the Tenant field value
func (o *FullAzureCloud) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *FullAzureCloud) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *FullAzureCloud) SetTenant(v string) {
	o.Tenant = v
}

// GetPublicContainerUrl returns the PublicContainerUrl field value if set, zero value otherwise.
func (o *FullAzureCloud) GetPublicContainerUrl() string {
	if o == nil || o.PublicContainerUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicContainerUrl
}

// GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloud) GetPublicContainerUrlOk() (*string, bool) {
	if o == nil || o.PublicContainerUrl == nil {
		return nil, false
	}
	return o.PublicContainerUrl, true
}

// HasPublicContainerUrl returns a boolean if a field has been set.
func (o *FullAzureCloud) HasPublicContainerUrl() bool {
	if o != nil && o.PublicContainerUrl != nil {
		return true
	}

	return false
}

// SetPublicContainerUrl gets a reference to the given string and assigns it to the PublicContainerUrl field.
func (o *FullAzureCloud) SetPublicContainerUrl(v string) {
	o.PublicContainerUrl = &v
}

func (o FullAzureCloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if o.PublicContainerUrl != nil {
		toSerialize["publicContainerUrl"] = o.PublicContainerUrl
	}
	return json.Marshal(toSerialize)
}

type NullableFullAzureCloud struct {
	value *FullAzureCloud
	isSet bool
}

func (v NullableFullAzureCloud) Get() *FullAzureCloud {
	return v.value
}

func (v *NullableFullAzureCloud) Set(val *FullAzureCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableFullAzureCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableFullAzureCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullAzureCloud(val *FullAzureCloud) *NullableFullAzureCloud {
	return &NullableFullAzureCloud{value: val, isSet: true}
}

func (v NullableFullAzureCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullAzureCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
