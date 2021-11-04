/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gocons3rt

import (
	"encoding/json"
)

// InputAzureVirtualizationRealmAllOf struct for InputAzureVirtualizationRealmAllOf
type InputAzureVirtualizationRealmAllOf struct {
	PublicContainerUrl *string `json:"publicContainerUrl,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	VirtualNetworkName *string `json:"virtualNetworkName,omitempty"`
}

// NewInputAzureVirtualizationRealmAllOf instantiates a new InputAzureVirtualizationRealmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputAzureVirtualizationRealmAllOf() *InputAzureVirtualizationRealmAllOf {
	this := InputAzureVirtualizationRealmAllOf{}
	return &this
}

// NewInputAzureVirtualizationRealmAllOfWithDefaults instantiates a new InputAzureVirtualizationRealmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputAzureVirtualizationRealmAllOfWithDefaults() *InputAzureVirtualizationRealmAllOf {
	this := InputAzureVirtualizationRealmAllOf{}
	return &this
}

// GetPublicContainerUrl returns the PublicContainerUrl field value if set, zero value otherwise.
func (o *InputAzureVirtualizationRealmAllOf) GetPublicContainerUrl() string {
	if o == nil || o.PublicContainerUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicContainerUrl
}

// GetPublicContainerUrlOk returns a tuple with the PublicContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAzureVirtualizationRealmAllOf) GetPublicContainerUrlOk() (*string, bool) {
	if o == nil || o.PublicContainerUrl == nil {
		return nil, false
	}
	return o.PublicContainerUrl, true
}

// HasPublicContainerUrl returns a boolean if a field has been set.
func (o *InputAzureVirtualizationRealmAllOf) HasPublicContainerUrl() bool {
	if o != nil && o.PublicContainerUrl != nil {
		return true
	}

	return false
}

// SetPublicContainerUrl gets a reference to the given string and assigns it to the PublicContainerUrl field.
func (o *InputAzureVirtualizationRealmAllOf) SetPublicContainerUrl(v string) {
	o.PublicContainerUrl = &v
}

// GetResourceGroupName returns the ResourceGroupName field value if set, zero value otherwise.
func (o *InputAzureVirtualizationRealmAllOf) GetResourceGroupName() string {
	if o == nil || o.ResourceGroupName == nil {
		var ret string
		return ret
	}
	return *o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAzureVirtualizationRealmAllOf) GetResourceGroupNameOk() (*string, bool) {
	if o == nil || o.ResourceGroupName == nil {
		return nil, false
	}
	return o.ResourceGroupName, true
}

// HasResourceGroupName returns a boolean if a field has been set.
func (o *InputAzureVirtualizationRealmAllOf) HasResourceGroupName() bool {
	if o != nil && o.ResourceGroupName != nil {
		return true
	}

	return false
}

// SetResourceGroupName gets a reference to the given string and assigns it to the ResourceGroupName field.
func (o *InputAzureVirtualizationRealmAllOf) SetResourceGroupName(v string) {
	o.ResourceGroupName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InputAzureVirtualizationRealmAllOf) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAzureVirtualizationRealmAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InputAzureVirtualizationRealmAllOf) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InputAzureVirtualizationRealmAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

// GetVirtualNetworkName returns the VirtualNetworkName field value if set, zero value otherwise.
func (o *InputAzureVirtualizationRealmAllOf) GetVirtualNetworkName() string {
	if o == nil || o.VirtualNetworkName == nil {
		var ret string
		return ret
	}
	return *o.VirtualNetworkName
}

// GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputAzureVirtualizationRealmAllOf) GetVirtualNetworkNameOk() (*string, bool) {
	if o == nil || o.VirtualNetworkName == nil {
		return nil, false
	}
	return o.VirtualNetworkName, true
}

// HasVirtualNetworkName returns a boolean if a field has been set.
func (o *InputAzureVirtualizationRealmAllOf) HasVirtualNetworkName() bool {
	if o != nil && o.VirtualNetworkName != nil {
		return true
	}

	return false
}

// SetVirtualNetworkName gets a reference to the given string and assigns it to the VirtualNetworkName field.
func (o *InputAzureVirtualizationRealmAllOf) SetVirtualNetworkName(v string) {
	o.VirtualNetworkName = &v
}

func (o InputAzureVirtualizationRealmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicContainerUrl != nil {
		toSerialize["publicContainerUrl"] = o.PublicContainerUrl
	}
	if o.ResourceGroupName != nil {
		toSerialize["resourceGroupName"] = o.ResourceGroupName
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.VirtualNetworkName != nil {
		toSerialize["virtualNetworkName"] = o.VirtualNetworkName
	}
	return json.Marshal(toSerialize)
}

type NullableInputAzureVirtualizationRealmAllOf struct {
	value *InputAzureVirtualizationRealmAllOf
	isSet bool
}

func (v NullableInputAzureVirtualizationRealmAllOf) Get() *InputAzureVirtualizationRealmAllOf {
	return v.value
}

func (v *NullableInputAzureVirtualizationRealmAllOf) Set(val *InputAzureVirtualizationRealmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInputAzureVirtualizationRealmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInputAzureVirtualizationRealmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputAzureVirtualizationRealmAllOf(val *InputAzureVirtualizationRealmAllOf) *NullableInputAzureVirtualizationRealmAllOf {
	return &NullableInputAzureVirtualizationRealmAllOf{value: val, isSet: true}
}

func (v NullableInputAzureVirtualizationRealmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputAzureVirtualizationRealmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


