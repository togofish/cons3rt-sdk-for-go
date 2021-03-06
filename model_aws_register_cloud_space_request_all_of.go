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

// AwsRegisterCloudSpaceRequestAllOf struct for AwsRegisterCloudSpaceRequestAllOf
type AwsRegisterCloudSpaceRequestAllOf struct {
	AccountId               *string            `json:"accountId,omitempty"`
	Region                  *string            `json:"region,omitempty"`
	VpcId                   *string            `json:"vpcId,omitempty"`
	UserKeyName             *string            `json:"userKeyName,omitempty"`
	NetworkSecurityGroupMap *map[string]string `json:"networkSecurityGroupMap,omitempty"`
}

// NewAwsRegisterCloudSpaceRequestAllOf instantiates a new AwsRegisterCloudSpaceRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsRegisterCloudSpaceRequestAllOf() *AwsRegisterCloudSpaceRequestAllOf {
	this := AwsRegisterCloudSpaceRequestAllOf{}
	return &this
}

// NewAwsRegisterCloudSpaceRequestAllOfWithDefaults instantiates a new AwsRegisterCloudSpaceRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsRegisterCloudSpaceRequestAllOfWithDefaults() *AwsRegisterCloudSpaceRequestAllOf {
	this := AwsRegisterCloudSpaceRequestAllOf{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AwsRegisterCloudSpaceRequestAllOf) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AwsRegisterCloudSpaceRequestAllOf) SetRegion(v string) {
	o.Region = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *AwsRegisterCloudSpaceRequestAllOf) SetVpcId(v string) {
	o.VpcId = &v
}

// GetUserKeyName returns the UserKeyName field value if set, zero value otherwise.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetUserKeyName() string {
	if o == nil || o.UserKeyName == nil {
		var ret string
		return ret
	}
	return *o.UserKeyName
}

// GetUserKeyNameOk returns a tuple with the UserKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetUserKeyNameOk() (*string, bool) {
	if o == nil || o.UserKeyName == nil {
		return nil, false
	}
	return o.UserKeyName, true
}

// HasUserKeyName returns a boolean if a field has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) HasUserKeyName() bool {
	if o != nil && o.UserKeyName != nil {
		return true
	}

	return false
}

// SetUserKeyName gets a reference to the given string and assigns it to the UserKeyName field.
func (o *AwsRegisterCloudSpaceRequestAllOf) SetUserKeyName(v string) {
	o.UserKeyName = &v
}

// GetNetworkSecurityGroupMap returns the NetworkSecurityGroupMap field value if set, zero value otherwise.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetNetworkSecurityGroupMap() map[string]string {
	if o == nil || o.NetworkSecurityGroupMap == nil {
		var ret map[string]string
		return ret
	}
	return *o.NetworkSecurityGroupMap
}

// GetNetworkSecurityGroupMapOk returns a tuple with the NetworkSecurityGroupMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) GetNetworkSecurityGroupMapOk() (*map[string]string, bool) {
	if o == nil || o.NetworkSecurityGroupMap == nil {
		return nil, false
	}
	return o.NetworkSecurityGroupMap, true
}

// HasNetworkSecurityGroupMap returns a boolean if a field has been set.
func (o *AwsRegisterCloudSpaceRequestAllOf) HasNetworkSecurityGroupMap() bool {
	if o != nil && o.NetworkSecurityGroupMap != nil {
		return true
	}

	return false
}

// SetNetworkSecurityGroupMap gets a reference to the given map[string]string and assigns it to the NetworkSecurityGroupMap field.
func (o *AwsRegisterCloudSpaceRequestAllOf) SetNetworkSecurityGroupMap(v map[string]string) {
	o.NetworkSecurityGroupMap = &v
}

func (o AwsRegisterCloudSpaceRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.VpcId != nil {
		toSerialize["vpcId"] = o.VpcId
	}
	if o.UserKeyName != nil {
		toSerialize["userKeyName"] = o.UserKeyName
	}
	if o.NetworkSecurityGroupMap != nil {
		toSerialize["networkSecurityGroupMap"] = o.NetworkSecurityGroupMap
	}
	return json.Marshal(toSerialize)
}

type NullableAwsRegisterCloudSpaceRequestAllOf struct {
	value *AwsRegisterCloudSpaceRequestAllOf
	isSet bool
}

func (v NullableAwsRegisterCloudSpaceRequestAllOf) Get() *AwsRegisterCloudSpaceRequestAllOf {
	return v.value
}

func (v *NullableAwsRegisterCloudSpaceRequestAllOf) Set(val *AwsRegisterCloudSpaceRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsRegisterCloudSpaceRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsRegisterCloudSpaceRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsRegisterCloudSpaceRequestAllOf(val *AwsRegisterCloudSpaceRequestAllOf) *NullableAwsRegisterCloudSpaceRequestAllOf {
	return &NullableAwsRegisterCloudSpaceRequestAllOf{value: val, isSet: true}
}

func (v NullableAwsRegisterCloudSpaceRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsRegisterCloudSpaceRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
