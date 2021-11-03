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

// AwsVirtualizationRealm struct for AwsVirtualizationRealm
type AwsVirtualizationRealm struct {
	NatImageId         *string `json:"natImageId,omitempty"`
	NatInstanceType    *string `json:"natInstanceType,omitempty"`
	NatKeyName         *string `json:"natKeyName,omitempty"`
	NatKeyPem          *string `json:"natKeyPem,omitempty"`
	Region             *string `json:"region,omitempty"`
	SecurityGroupId    *string `json:"securityGroupId,omitempty"`
	UserKeyName        *string `json:"userKeyName,omitempty"`
	UserKeyPem         *string `json:"userKeyPem,omitempty"`
	VirtualNetworkName *string `json:"virtualNetworkName,omitempty"`
	VpcId              *string `json:"vpcId,omitempty"`
	VpcSubnetName      *string `json:"vpcSubnetName,omitempty"`
}

// NewAwsVirtualizationRealm instantiates a new AwsVirtualizationRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsVirtualizationRealm(accountId string, cidr string, description string, name string, password string, username string) *AwsVirtualizationRealm {
	this := AwsVirtualizationRealm{}
	this.AccountId = accountId
	this.Cidr = cidr
	this.Description = description
	this.Name = name
	this.Password = password
	this.Username = username
	return &this
}

// NewAwsVirtualizationRealmWithDefaults instantiates a new AwsVirtualizationRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsVirtualizationRealmWithDefaults() *AwsVirtualizationRealm {
	this := AwsVirtualizationRealm{}
	return &this
}

// GetNatImageId returns the NatImageId field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetNatImageId() string {
	if o == nil || o.NatImageId == nil {
		var ret string
		return ret
	}
	return *o.NatImageId
}

// GetNatImageIdOk returns a tuple with the NatImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetNatImageIdOk() (*string, bool) {
	if o == nil || o.NatImageId == nil {
		return nil, false
	}
	return o.NatImageId, true
}

// HasNatImageId returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasNatImageId() bool {
	if o != nil && o.NatImageId != nil {
		return true
	}

	return false
}

// SetNatImageId gets a reference to the given string and assigns it to the NatImageId field.
func (o *AwsVirtualizationRealm) SetNatImageId(v string) {
	o.NatImageId = &v
}

// GetNatInstanceType returns the NatInstanceType field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetNatInstanceType() string {
	if o == nil || o.NatInstanceType == nil {
		var ret string
		return ret
	}
	return *o.NatInstanceType
}

// GetNatInstanceTypeOk returns a tuple with the NatInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetNatInstanceTypeOk() (*string, bool) {
	if o == nil || o.NatInstanceType == nil {
		return nil, false
	}
	return o.NatInstanceType, true
}

// HasNatInstanceType returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasNatInstanceType() bool {
	if o != nil && o.NatInstanceType != nil {
		return true
	}

	return false
}

// SetNatInstanceType gets a reference to the given string and assigns it to the NatInstanceType field.
func (o *AwsVirtualizationRealm) SetNatInstanceType(v string) {
	o.NatInstanceType = &v
}

// GetNatKeyName returns the NatKeyName field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetNatKeyName() string {
	if o == nil || o.NatKeyName == nil {
		var ret string
		return ret
	}
	return *o.NatKeyName
}

// GetNatKeyNameOk returns a tuple with the NatKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetNatKeyNameOk() (*string, bool) {
	if o == nil || o.NatKeyName == nil {
		return nil, false
	}
	return o.NatKeyName, true
}

// HasNatKeyName returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasNatKeyName() bool {
	if o != nil && o.NatKeyName != nil {
		return true
	}

	return false
}

// SetNatKeyName gets a reference to the given string and assigns it to the NatKeyName field.
func (o *AwsVirtualizationRealm) SetNatKeyName(v string) {
	o.NatKeyName = &v
}

// GetNatKeyPem returns the NatKeyPem field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetNatKeyPem() string {
	if o == nil || o.NatKeyPem == nil {
		var ret string
		return ret
	}
	return *o.NatKeyPem
}

// GetNatKeyPemOk returns a tuple with the NatKeyPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetNatKeyPemOk() (*string, bool) {
	if o == nil || o.NatKeyPem == nil {
		return nil, false
	}
	return o.NatKeyPem, true
}

// HasNatKeyPem returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasNatKeyPem() bool {
	if o != nil && o.NatKeyPem != nil {
		return true
	}

	return false
}

// SetNatKeyPem gets a reference to the given string and assigns it to the NatKeyPem field.
func (o *AwsVirtualizationRealm) SetNatKeyPem(v string) {
	o.NatKeyPem = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AwsVirtualizationRealm) SetRegion(v string) {
	o.Region = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *AwsVirtualizationRealm) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetUserKeyName returns the UserKeyName field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetUserKeyName() string {
	if o == nil || o.UserKeyName == nil {
		var ret string
		return ret
	}
	return *o.UserKeyName
}

// GetUserKeyNameOk returns a tuple with the UserKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetUserKeyNameOk() (*string, bool) {
	if o == nil || o.UserKeyName == nil {
		return nil, false
	}
	return o.UserKeyName, true
}

// HasUserKeyName returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasUserKeyName() bool {
	if o != nil && o.UserKeyName != nil {
		return true
	}

	return false
}

// SetUserKeyName gets a reference to the given string and assigns it to the UserKeyName field.
func (o *AwsVirtualizationRealm) SetUserKeyName(v string) {
	o.UserKeyName = &v
}

// GetUserKeyPem returns the UserKeyPem field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetUserKeyPem() string {
	if o == nil || o.UserKeyPem == nil {
		var ret string
		return ret
	}
	return *o.UserKeyPem
}

// GetUserKeyPemOk returns a tuple with the UserKeyPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetUserKeyPemOk() (*string, bool) {
	if o == nil || o.UserKeyPem == nil {
		return nil, false
	}
	return o.UserKeyPem, true
}

// HasUserKeyPem returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasUserKeyPem() bool {
	if o != nil && o.UserKeyPem != nil {
		return true
	}

	return false
}

// SetUserKeyPem gets a reference to the given string and assigns it to the UserKeyPem field.
func (o *AwsVirtualizationRealm) SetUserKeyPem(v string) {
	o.UserKeyPem = &v
}

// GetVirtualNetworkName returns the VirtualNetworkName field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetVirtualNetworkName() string {
	if o == nil || o.VirtualNetworkName == nil {
		var ret string
		return ret
	}
	return *o.VirtualNetworkName
}

// GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetVirtualNetworkNameOk() (*string, bool) {
	if o == nil || o.VirtualNetworkName == nil {
		return nil, false
	}
	return o.VirtualNetworkName, true
}

// HasVirtualNetworkName returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasVirtualNetworkName() bool {
	if o != nil && o.VirtualNetworkName != nil {
		return true
	}

	return false
}

// SetVirtualNetworkName gets a reference to the given string and assigns it to the VirtualNetworkName field.
func (o *AwsVirtualizationRealm) SetVirtualNetworkName(v string) {
	o.VirtualNetworkName = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *AwsVirtualizationRealm) SetVpcId(v string) {
	o.VpcId = &v
}

// GetVpcSubnetName returns the VpcSubnetName field value if set, zero value otherwise.
func (o *AwsVirtualizationRealm) GetVpcSubnetName() string {
	if o == nil || o.VpcSubnetName == nil {
		var ret string
		return ret
	}
	return *o.VpcSubnetName
}

// GetVpcSubnetNameOk returns a tuple with the VpcSubnetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsVirtualizationRealm) GetVpcSubnetNameOk() (*string, bool) {
	if o == nil || o.VpcSubnetName == nil {
		return nil, false
	}
	return o.VpcSubnetName, true
}

// HasVpcSubnetName returns a boolean if a field has been set.
func (o *AwsVirtualizationRealm) HasVpcSubnetName() bool {
	if o != nil && o.VpcSubnetName != nil {
		return true
	}

	return false
}

// SetVpcSubnetName gets a reference to the given string and assigns it to the VpcSubnetName field.
func (o *AwsVirtualizationRealm) SetVpcSubnetName(v string) {
	o.VpcSubnetName = &v
}

func (o AwsVirtualizationRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NatImageId != nil {
		toSerialize["natImageId"] = o.NatImageId
	}
	if o.NatInstanceType != nil {
		toSerialize["natInstanceType"] = o.NatInstanceType
	}
	if o.NatKeyName != nil {
		toSerialize["natKeyName"] = o.NatKeyName
	}
	if o.NatKeyPem != nil {
		toSerialize["natKeyPem"] = o.NatKeyPem
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.SecurityGroupId != nil {
		toSerialize["securityGroupId"] = o.SecurityGroupId
	}
	if o.UserKeyName != nil {
		toSerialize["userKeyName"] = o.UserKeyName
	}
	if o.UserKeyPem != nil {
		toSerialize["userKeyPem"] = o.UserKeyPem
	}
	if o.VirtualNetworkName != nil {
		toSerialize["virtualNetworkName"] = o.VirtualNetworkName
	}
	if o.VpcId != nil {
		toSerialize["vpcId"] = o.VpcId
	}
	if o.VpcSubnetName != nil {
		toSerialize["vpcSubnetName"] = o.VpcSubnetName
	}
	return json.Marshal(toSerialize)
}

type NullableAwsVirtualizationRealm struct {
	value *AwsVirtualizationRealm
	isSet bool
}

func (v NullableAwsVirtualizationRealm) Get() *AwsVirtualizationRealm {
	return v.value
}

func (v *NullableAwsVirtualizationRealm) Set(val *AwsVirtualizationRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsVirtualizationRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsVirtualizationRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsVirtualizationRealm(val *AwsVirtualizationRealm) *NullableAwsVirtualizationRealm {
	return &NullableAwsVirtualizationRealm{value: val, isSet: true}
}

func (v NullableAwsVirtualizationRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsVirtualizationRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
