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

// InputAzureCloud struct for InputAzureCloud
type InputAzureCloud struct {
	ClientId           string  `json:"clientId"`
	Environment        string  `json:"environment"`
	RegionName         string  `json:"regionName"`
	SecretAccessKey    string  `json:"secretAccessKey"`
	SubscriptionId     string  `json:"subscriptionId"`
	Tenant             string  `json:"tenant"`
	PublicContainerUrl *string `json:"publicContainerUrl,omitempty"`
}

// NewInputAzureCloud instantiates a new InputAzureCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputAzureCloud(clientId string, environment string, regionName string, secretAccessKey string, subscriptionId string, tenant string, name string, externalIpSource string, maximumImpactLevel string, owningTeam InputTeam, subtype string) *InputAzureCloud {
	this := InputAzureCloud{}
	this.Name = name
	this.ExternalIpSource = externalIpSource
	this.MaximumImpactLevel = maximumImpactLevel
	this.OwningTeam = owningTeam
	this.Subtype = subtype
	this.ClientId = clientId
	this.Environment = environment
	this.RegionName = regionName
	this.SecretAccessKey = secretAccessKey
	this.SubscriptionId = subscriptionId
	this.Tenant = tenant
	return &this
}

// NewInputAzureCloudWithDefaults instantiates a new InputAzureCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputAzureCloudWithDefaults() *InputAzureCloud {
	this := InputAzureCloud{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *InputAzureCloud) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *InputAzureCloud) SetClientId(v string) {
	o.ClientId = v
}

// GetEnvironment returns the Environment field value
func (o *InputAzureCloud) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *InputAzureCloud) SetEnvironment(v string) {
	o.Environment = v
}

// GetRegionName returns the RegionName field value
func (o *InputAzureCloud) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *InputAzureCloud) SetRegionName(v string) {
	o.RegionName = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *InputAzureCloud) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *InputAzureCloud) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *InputAzureCloud) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *InputAzureCloud) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetTenant returns the Tenant field value
func (o *InputAzureCloud) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *InputAzureCloud) SetTenant(v string) {
	o.Tenant = v
}

// GetPublicContainerUrl returns the PublicContainerUrl field value if set, zero value otherwise.
func (o *InputAzureCloud) GetPublicContainerUrl() string {
	if o == nil || o.PublicContainerUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicContainerUrl
}

// GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAzureCloud) GetPublicContainerUrlOk() (*string, bool) {
	if o == nil || o.PublicContainerUrl == nil {
		return nil, false
	}
	return o.PublicContainerUrl, true
}

// HasPublicContainerUrl returns a boolean if a field has been set.
func (o *InputAzureCloud) HasPublicContainerUrl() bool {
	if o != nil && o.PublicContainerUrl != nil {
		return true
	}

	return false
}

// SetPublicContainerUrl gets a reference to the given string and assigns it to the PublicContainerUrl field.
func (o *InputAzureCloud) SetPublicContainerUrl(v string) {
	o.PublicContainerUrl = &v
}

func (o InputAzureCloud) MarshalJSON() ([]byte, error) {
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
		toSerialize["secretAccessKey"] = o.SecretAccessKey
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

type NullableInputAzureCloud struct {
	value *InputAzureCloud
	isSet bool
}

func (v NullableInputAzureCloud) Get() *InputAzureCloud {
	return v.value
}

func (v *NullableInputAzureCloud) Set(val *InputAzureCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableInputAzureCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableInputAzureCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputAzureCloud(val *InputAzureCloud) *NullableInputAzureCloud {
	return &NullableInputAzureCloud{value: val, isSet: true}
}

func (v NullableInputAzureCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputAzureCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
