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

// AwsCloudAllOf struct for AwsCloudAllOf
type AwsCloudAllOf struct {
	AccessKey       *string `json:"accessKey,omitempty"`
	OwnerId         *string `json:"ownerId,omitempty"`
	RegionName      *string `json:"regionName,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
}

// NewAwsCloudAllOf instantiates a new AwsCloudAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCloudAllOf() *AwsCloudAllOf {
	this := AwsCloudAllOf{}
	return &this
}

// NewAwsCloudAllOfWithDefaults instantiates a new AwsCloudAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCloudAllOfWithDefaults() *AwsCloudAllOf {
	this := AwsCloudAllOf{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *AwsCloudAllOf) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCloudAllOf) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *AwsCloudAllOf) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *AwsCloudAllOf) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AwsCloudAllOf) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCloudAllOf) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AwsCloudAllOf) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AwsCloudAllOf) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AwsCloudAllOf) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCloudAllOf) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AwsCloudAllOf) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AwsCloudAllOf) SetRegionName(v string) {
	o.RegionName = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AwsCloudAllOf) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCloudAllOf) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AwsCloudAllOf) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AwsCloudAllOf) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

func (o AwsCloudAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey != nil {
		toSerialize["accessKey"] = o.AccessKey
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	if o.SecretAccessKey != nil {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	return json.Marshal(toSerialize)
}

type NullableAwsCloudAllOf struct {
	value *AwsCloudAllOf
	isSet bool
}

func (v NullableAwsCloudAllOf) Get() *AwsCloudAllOf {
	return v.value
}

func (v *NullableAwsCloudAllOf) Set(val *AwsCloudAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCloudAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCloudAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCloudAllOf(val *AwsCloudAllOf) *NullableAwsCloudAllOf {
	return &NullableAwsCloudAllOf{value: val, isSet: true}
}

func (v NullableAwsCloudAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCloudAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
