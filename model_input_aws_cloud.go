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

// InputAwsCloud struct for InputAwsCloud
type InputAwsCloud struct {
	AccessKey          string `json:"accessKey"`
	OwnerId            string `json:"ownerId"`
	RegionName         string `json:"regionName"`
	SecretAccessKey    string `json:"secretAccessKey"`
	Name               string
	ExternalIpSource   string
	MaximumImpactLevel string
	OwningTeam         InputTeam
	Subtype            string
}

// NewInputAwsCloud instantiates a new InputAwsCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputAwsCloud(accessKey string, ownerId string, regionName string, secretAccessKey string, name string, externalIpSource string, maximumImpactLevel string, owningTeam InputTeam, subtype string) *InputAwsCloud {
	this := InputAwsCloud{}
	this.Name = name
	this.ExternalIpSource = externalIpSource
	this.MaximumImpactLevel = maximumImpactLevel
	this.OwningTeam = owningTeam
	this.Subtype = subtype
	this.AccessKey = accessKey
	this.OwnerId = ownerId
	this.RegionName = regionName
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewInputAwsCloudWithDefaults instantiates a new InputAwsCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputAwsCloudWithDefaults() *InputAwsCloud {
	this := InputAwsCloud{}
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *InputAwsCloud) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *InputAwsCloud) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *InputAwsCloud) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetOwnerId returns the OwnerId field value
func (o *InputAwsCloud) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *InputAwsCloud) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *InputAwsCloud) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetRegionName returns the RegionName field value
func (o *InputAwsCloud) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *InputAwsCloud) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *InputAwsCloud) SetRegionName(v string) {
	o.RegionName = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *InputAwsCloud) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *InputAwsCloud) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *InputAwsCloud) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

func (o InputAwsCloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessKey"] = o.AccessKey
	}
	if true {
		toSerialize["ownerId"] = o.OwnerId
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if true {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	return json.Marshal(toSerialize)
}

type NullableInputAwsCloud struct {
	value *InputAwsCloud
	isSet bool
}

func (v NullableInputAwsCloud) Get() *InputAwsCloud {
	return v.value
}

func (v *NullableInputAwsCloud) Set(val *InputAwsCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableInputAwsCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableInputAwsCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputAwsCloud(val *InputAwsCloud) *NullableInputAwsCloud {
	return &NullableInputAwsCloud{value: val, isSet: true}
}

func (v NullableInputAwsCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputAwsCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
