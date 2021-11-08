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

// FullAzureCloudAllOf struct for FullAzureCloudAllOf
type FullAzureCloudAllOf struct {
	ClientId           *string `json:"clientId,omitempty"`
	Environment        *string `json:"environment,omitempty"`
	RegionName         *string `json:"regionName,omitempty"`
	SubscriptionId     *string `json:"subscriptionId,omitempty"`
	Tenant             *string `json:"tenant,omitempty"`
	PublicContainerUrl *string `json:"publicContainerUrl,omitempty"`
}

// NewFullAzureCloudAllOf instantiates a new FullAzureCloudAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullAzureCloudAllOf() *FullAzureCloudAllOf {
	this := FullAzureCloudAllOf{}
	return &this
}

// NewFullAzureCloudAllOfWithDefaults instantiates a new FullAzureCloudAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullAzureCloudAllOfWithDefaults() *FullAzureCloudAllOf {
	this := FullAzureCloudAllOf{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *FullAzureCloudAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloudAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *FullAzureCloudAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *FullAzureCloudAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FullAzureCloudAllOf) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloudAllOf) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FullAzureCloudAllOf) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *FullAzureCloudAllOf) SetEnvironment(v string) {
	o.Environment = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *FullAzureCloudAllOf) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloudAllOf) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *FullAzureCloudAllOf) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *FullAzureCloudAllOf) SetRegionName(v string) {
	o.RegionName = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *FullAzureCloudAllOf) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloudAllOf) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *FullAzureCloudAllOf) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *FullAzureCloudAllOf) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *FullAzureCloudAllOf) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloudAllOf) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *FullAzureCloudAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *FullAzureCloudAllOf) SetTenant(v string) {
	o.Tenant = &v
}

// GetPublicContainerUrl returns the PublicContainerUrl field value if set, zero value otherwise.
func (o *FullAzureCloudAllOf) GetPublicContainerUrl() string {
	if o == nil || o.PublicContainerUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicContainerUrl
}

// GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAzureCloudAllOf) GetPublicContainerUrlOk() (*string, bool) {
	if o == nil || o.PublicContainerUrl == nil {
		return nil, false
	}
	return o.PublicContainerUrl, true
}

// HasPublicContainerUrl returns a boolean if a field has been set.
func (o *FullAzureCloudAllOf) HasPublicContainerUrl() bool {
	if o != nil && o.PublicContainerUrl != nil {
		return true
	}

	return false
}

// SetPublicContainerUrl gets a reference to the given string and assigns it to the PublicContainerUrl field.
func (o *FullAzureCloudAllOf) SetPublicContainerUrl(v string) {
	o.PublicContainerUrl = &v
}

func (o FullAzureCloudAllOf) MarshalJSON() ([]byte, error) {
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

type NullableFullAzureCloudAllOf struct {
	value *FullAzureCloudAllOf
	isSet bool
}

func (v NullableFullAzureCloudAllOf) Get() *FullAzureCloudAllOf {
	return v.value
}

func (v *NullableFullAzureCloudAllOf) Set(val *FullAzureCloudAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullAzureCloudAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullAzureCloudAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullAzureCloudAllOf(val *FullAzureCloudAllOf) *NullableFullAzureCloudAllOf {
	return &NullableFullAzureCloudAllOf{value: val, isSet: true}
}

func (v NullableFullAzureCloudAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullAzureCloudAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
