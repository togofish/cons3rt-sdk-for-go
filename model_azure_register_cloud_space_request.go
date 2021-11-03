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

// AzureRegisterCloudSpaceRequest struct for AzureRegisterCloudSpaceRequest
type AzureRegisterCloudSpaceRequest struct {
	AccountId                     string  `json:"accountId"`
	Environment                   string  `json:"environment"`
	LocalStorageResourceGroupName *string `json:"localStorageResourceGroupName,omitempty"`
	LocalStorageKey               *string `json:"localStorageKey,omitempty"`
	Region                        string  `json:"region"`
	ResourceGroupName             string  `json:"resourceGroupName"`
	TenantId                      string  `json:"tenantId"`
	VirtualNetworkName            string  `json:"virtualNetworkName"`
	LocalStorageName              *string `json:"localStorageName,omitempty"`
}

// NewAzureRegisterCloudSpaceRequest instantiates a new AzureRegisterCloudSpaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureRegisterCloudSpaceRequest(accountId string, environment string, region string, resourceGroupName string, tenantId string, virtualNetworkName string, virtualizationRealmType string, name string, cons3rtNetworkName string, password string, primaryNetworkName string, username string) *AzureRegisterCloudSpaceRequest {
	this := AzureRegisterCloudSpaceRequest{}
	this.VirtualizationRealmType = virtualizationRealmType
	this.Name = name
	this.Cons3rtNetworkName = cons3rtNetworkName
	this.Password = password
	this.PrimaryNetworkName = primaryNetworkName
	this.Username = username
	this.AccountId = accountId
	this.Environment = environment
	this.Region = region
	this.ResourceGroupName = resourceGroupName
	this.TenantId = tenantId
	this.VirtualNetworkName = virtualNetworkName
	return &this
}

// NewAzureRegisterCloudSpaceRequestWithDefaults instantiates a new AzureRegisterCloudSpaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureRegisterCloudSpaceRequestWithDefaults() *AzureRegisterCloudSpaceRequest {
	this := AzureRegisterCloudSpaceRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AzureRegisterCloudSpaceRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AzureRegisterCloudSpaceRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetEnvironment returns the Environment field value
func (o *AzureRegisterCloudSpaceRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *AzureRegisterCloudSpaceRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetLocalStorageResourceGroupName returns the LocalStorageResourceGroupName field value if set, zero value otherwise.
func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageResourceGroupName() string {
	if o == nil || o.LocalStorageResourceGroupName == nil {
		var ret string
		return ret
	}
	return *o.LocalStorageResourceGroupName
}

// GetLocalStorageResourceGroupNameOk returns a tuple with the LocalStorageResourceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageResourceGroupNameOk() (*string, bool) {
	if o == nil || o.LocalStorageResourceGroupName == nil {
		return nil, false
	}
	return o.LocalStorageResourceGroupName, true
}

// HasLocalStorageResourceGroupName returns a boolean if a field has been set.
func (o *AzureRegisterCloudSpaceRequest) HasLocalStorageResourceGroupName() bool {
	if o != nil && o.LocalStorageResourceGroupName != nil {
		return true
	}

	return false
}

// SetLocalStorageResourceGroupName gets a reference to the given string and assigns it to the LocalStorageResourceGroupName field.
func (o *AzureRegisterCloudSpaceRequest) SetLocalStorageResourceGroupName(v string) {
	o.LocalStorageResourceGroupName = &v
}

// GetLocalStorageKey returns the LocalStorageKey field value if set, zero value otherwise.
func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageKey() string {
	if o == nil || o.LocalStorageKey == nil {
		var ret string
		return ret
	}
	return *o.LocalStorageKey
}

// GetLocalStorageKeyOk returns a tuple with the LocalStorageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageKeyOk() (*string, bool) {
	if o == nil || o.LocalStorageKey == nil {
		return nil, false
	}
	return o.LocalStorageKey, true
}

// HasLocalStorageKey returns a boolean if a field has been set.
func (o *AzureRegisterCloudSpaceRequest) HasLocalStorageKey() bool {
	if o != nil && o.LocalStorageKey != nil {
		return true
	}

	return false
}

// SetLocalStorageKey gets a reference to the given string and assigns it to the LocalStorageKey field.
func (o *AzureRegisterCloudSpaceRequest) SetLocalStorageKey(v string) {
	o.LocalStorageKey = &v
}

// GetRegion returns the Region field value
func (o *AzureRegisterCloudSpaceRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AzureRegisterCloudSpaceRequest) SetRegion(v string) {
	o.Region = v
}

// GetResourceGroupName returns the ResourceGroupName field value
func (o *AzureRegisterCloudSpaceRequest) GetResourceGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroupName, true
}

// SetResourceGroupName sets field value
func (o *AzureRegisterCloudSpaceRequest) SetResourceGroupName(v string) {
	o.ResourceGroupName = v
}

// GetTenantId returns the TenantId field value
func (o *AzureRegisterCloudSpaceRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *AzureRegisterCloudSpaceRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetVirtualNetworkName returns the VirtualNetworkName field value
func (o *AzureRegisterCloudSpaceRequest) GetVirtualNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualNetworkName
}

// GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field value
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetVirtualNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetworkName, true
}

// SetVirtualNetworkName sets field value
func (o *AzureRegisterCloudSpaceRequest) SetVirtualNetworkName(v string) {
	o.VirtualNetworkName = v
}

// GetLocalStorageName returns the LocalStorageName field value if set, zero value otherwise.
func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageName() string {
	if o == nil || o.LocalStorageName == nil {
		var ret string
		return ret
	}
	return *o.LocalStorageName
}

// GetLocalStorageNameOk returns a tuple with the LocalStorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRegisterCloudSpaceRequest) GetLocalStorageNameOk() (*string, bool) {
	if o == nil || o.LocalStorageName == nil {
		return nil, false
	}
	return o.LocalStorageName, true
}

// HasLocalStorageName returns a boolean if a field has been set.
func (o *AzureRegisterCloudSpaceRequest) HasLocalStorageName() bool {
	if o != nil && o.LocalStorageName != nil {
		return true
	}

	return false
}

// SetLocalStorageName gets a reference to the given string and assigns it to the LocalStorageName field.
func (o *AzureRegisterCloudSpaceRequest) SetLocalStorageName(v string) {
	o.LocalStorageName = &v
}

func (o AzureRegisterCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if o.LocalStorageResourceGroupName != nil {
		toSerialize["localStorageResourceGroupName"] = o.LocalStorageResourceGroupName
	}
	if o.LocalStorageKey != nil {
		toSerialize["localStorageKey"] = o.LocalStorageKey
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
	if o.LocalStorageName != nil {
		toSerialize["localStorageName"] = o.LocalStorageName
	}
	return json.Marshal(toSerialize)
}

type NullableAzureRegisterCloudSpaceRequest struct {
	value *AzureRegisterCloudSpaceRequest
	isSet bool
}

func (v NullableAzureRegisterCloudSpaceRequest) Get() *AzureRegisterCloudSpaceRequest {
	return v.value
}

func (v *NullableAzureRegisterCloudSpaceRequest) Set(val *AzureRegisterCloudSpaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureRegisterCloudSpaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureRegisterCloudSpaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureRegisterCloudSpaceRequest(val *AzureRegisterCloudSpaceRequest) *NullableAzureRegisterCloudSpaceRequest {
	return &NullableAzureRegisterCloudSpaceRequest{value: val, isSet: true}
}

func (v NullableAzureRegisterCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureRegisterCloudSpaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
