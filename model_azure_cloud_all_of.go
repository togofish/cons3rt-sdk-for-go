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

// AzureCloudAllOf struct for AzureCloudAllOf
type AzureCloudAllOf struct {
	ClientId           *string `json:"clientId,omitempty"`
	Environment        *string `json:"environment,omitempty"`
	RegionName         *string `json:"regionName,omitempty"`
	SecretAccessKey    *string `json:"secretAccessKey,omitempty"`
	SubscriptionId     *string `json:"subscriptionId,omitempty"`
	Tenant             *string `json:"tenant,omitempty"`
	PublicContainerUrl *string `json:"publicContainerUrl,omitempty"`
}

// NewAzureCloudAllOf instantiates a new AzureCloudAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCloudAllOf() *AzureCloudAllOf {
	this := AzureCloudAllOf{}
	return &this
}

// NewAzureCloudAllOfWithDefaults instantiates a new AzureCloudAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCloudAllOfWithDefaults() *AzureCloudAllOf {
	this := AzureCloudAllOf{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AzureCloudAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AzureCloudAllOf) SetEnvironment(v string) {
	o.Environment = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AzureCloudAllOf) SetRegionName(v string) {
	o.RegionName = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AzureCloudAllOf) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AzureCloudAllOf) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AzureCloudAllOf) SetTenant(v string) {
	o.Tenant = &v
}

// GetPublicContainerUrl returns the PublicContainerUrl field value if set, zero value otherwise.
func (o *AzureCloudAllOf) GetPublicContainerUrl() string {
	if o == nil || o.PublicContainerUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicContainerUrl
}

// GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudAllOf) GetPublicContainerUrlOk() (*string, bool) {
	if o == nil || o.PublicContainerUrl == nil {
		return nil, false
	}
	return o.PublicContainerUrl, true
}

// HasPublicContainerUrl returns a boolean if a field has been set.
func (o *AzureCloudAllOf) HasPublicContainerUrl() bool {
	if o != nil && o.PublicContainerUrl != nil {
		return true
	}

	return false
}

// SetPublicContainerUrl gets a reference to the given string and assigns it to the PublicContainerUrl field.
func (o *AzureCloudAllOf) SetPublicContainerUrl(v string) {
	o.PublicContainerUrl = &v
}

func (o AzureCloudAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	if o.SecretAccessKey != nil {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.PublicContainerUrl != nil {
		toSerialize["publicContainerUrl"] = o.PublicContainerUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAzureCloudAllOf struct {
	value *AzureCloudAllOf
	isSet bool
}

func (v NullableAzureCloudAllOf) Get() *AzureCloudAllOf {
	return v.value
}

func (v *NullableAzureCloudAllOf) Set(val *AzureCloudAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCloudAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCloudAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCloudAllOf(val *AzureCloudAllOf) *NullableAzureCloudAllOf {
	return &NullableAzureCloudAllOf{value: val, isSet: true}
}

func (v NullableAzureCloudAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCloudAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
