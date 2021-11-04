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

// AwsRegisterCloudSpaceRequest struct for AwsRegisterCloudSpaceRequest
type AwsRegisterCloudSpaceRequest struct {
	AccountId string `json:"accountId"`
	Region string `json:"region"`
	VpcId string `json:"vpcId"`
	UserKeyName *string `json:"userKeyName,omitempty"`
	NetworkSecurityGroupMap map[string]string `json:"networkSecurityGroupMap"`
}

// NewAwsRegisterCloudSpaceRequest instantiates a new AwsRegisterCloudSpaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsRegisterCloudSpaceRequest(accountId string, region string, vpcId string, networkSecurityGroupMap map[string]string, virtualizationRealmType string, name string, cons3rtNetworkName string, password string, primaryNetworkName string, username string) *AwsRegisterCloudSpaceRequest {
	this := AwsRegisterCloudSpaceRequest{}
	this.VirtualizationRealmType = virtualizationRealmType
	this.Name = name
	this.Cons3rtNetworkName = cons3rtNetworkName
	this.Password = password
	this.PrimaryNetworkName = primaryNetworkName
	this.Username = username
	this.AccountId = accountId
	this.Region = region
	this.VpcId = vpcId
	this.NetworkSecurityGroupMap = networkSecurityGroupMap
	return &this
}

// NewAwsRegisterCloudSpaceRequestWithDefaults instantiates a new AwsRegisterCloudSpaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsRegisterCloudSpaceRequestWithDefaults() *AwsRegisterCloudSpaceRequest {
	this := AwsRegisterCloudSpaceRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AwsRegisterCloudSpaceRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AwsRegisterCloudSpaceRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetRegion returns the Region field value
func (o *AwsRegisterCloudSpaceRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequest) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AwsRegisterCloudSpaceRequest) SetRegion(v string) {
	o.Region = v
}

// GetVpcId returns the VpcId field value
func (o *AwsRegisterCloudSpaceRequest) GetVpcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequest) GetVpcIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VpcId, true
}

// SetVpcId sets field value
func (o *AwsRegisterCloudSpaceRequest) SetVpcId(v string) {
	o.VpcId = v
}

// GetUserKeyName returns the UserKeyName field value if set, zero value otherwise.
func (o *AwsRegisterCloudSpaceRequest) GetUserKeyName() string {
	if o == nil || o.UserKeyName == nil {
		var ret string
		return ret
	}
	return *o.UserKeyName
}

// GetUserKeyNameOk returns a tuple with the UserKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequest) GetUserKeyNameOk() (*string, bool) {
	if o == nil || o.UserKeyName == nil {
		return nil, false
	}
	return o.UserKeyName, true
}

// HasUserKeyName returns a boolean if a field has been set.
func (o *AwsRegisterCloudSpaceRequest) HasUserKeyName() bool {
	if o != nil && o.UserKeyName != nil {
		return true
	}

	return false
}

// SetUserKeyName gets a reference to the given string and assigns it to the UserKeyName field.
func (o *AwsRegisterCloudSpaceRequest) SetUserKeyName(v string) {
	o.UserKeyName = &v
}

// GetNetworkSecurityGroupMap returns the NetworkSecurityGroupMap field value
func (o *AwsRegisterCloudSpaceRequest) GetNetworkSecurityGroupMap() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.NetworkSecurityGroupMap
}

// GetNetworkSecurityGroupMapOk returns a tuple with the NetworkSecurityGroupMap field value
// and a boolean to check if the value has been set.
func (o *AwsRegisterCloudSpaceRequest) GetNetworkSecurityGroupMapOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkSecurityGroupMap, true
}

// SetNetworkSecurityGroupMap sets field value
func (o *AwsRegisterCloudSpaceRequest) SetNetworkSecurityGroupMap(v map[string]string) {
	o.NetworkSecurityGroupMap = v
}

func (o AwsRegisterCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["vpcId"] = o.VpcId
	}
	if o.UserKeyName != nil {
		toSerialize["userKeyName"] = o.UserKeyName
	}
	if true {
		toSerialize["networkSecurityGroupMap"] = o.NetworkSecurityGroupMap
	}
	return json.Marshal(toSerialize)
}

type NullableAwsRegisterCloudSpaceRequest struct {
	value *AwsRegisterCloudSpaceRequest
	isSet bool
}

func (v NullableAwsRegisterCloudSpaceRequest) Get() *AwsRegisterCloudSpaceRequest {
	return v.value
}

func (v *NullableAwsRegisterCloudSpaceRequest) Set(val *AwsRegisterCloudSpaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsRegisterCloudSpaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsRegisterCloudSpaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsRegisterCloudSpaceRequest(val *AwsRegisterCloudSpaceRequest) *NullableAwsRegisterCloudSpaceRequest {
	return &NullableAwsRegisterCloudSpaceRequest{value: val, isSet: true}
}

func (v NullableAwsRegisterCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsRegisterCloudSpaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


