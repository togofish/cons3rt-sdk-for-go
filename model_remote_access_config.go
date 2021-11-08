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

// RemoteAccessConfig struct for RemoteAccessConfig
type RemoteAccessConfig struct {
	Id                    *int32  `json:"id,omitempty"`
	RemoteAccessIpAddress string  `json:"remoteAccessIpAddress"`
	RemoteAccessPort      int32   `json:"remoteAccessPort"`
	InstanceType          string  `json:"instanceType"`
	TemplateName          *string `json:"templateName,omitempty"`
}

// NewRemoteAccessConfig instantiates a new RemoteAccessConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteAccessConfig(remoteAccessIpAddress string, remoteAccessPort int32, instanceType string) *RemoteAccessConfig {
	this := RemoteAccessConfig{}
	this.RemoteAccessIpAddress = remoteAccessIpAddress
	this.RemoteAccessPort = remoteAccessPort
	this.InstanceType = instanceType
	return &this
}

// NewRemoteAccessConfigWithDefaults instantiates a new RemoteAccessConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteAccessConfigWithDefaults() *RemoteAccessConfig {
	this := RemoteAccessConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteAccessConfig) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessConfig) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteAccessConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RemoteAccessConfig) SetId(v int32) {
	o.Id = &v
}

// GetRemoteAccessIpAddress returns the RemoteAccessIpAddress field value
func (o *RemoteAccessConfig) GetRemoteAccessIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteAccessIpAddress
}

// GetRemoteAccessIpAddressOk returns a tuple with the RemoteAccessIpAddress field value
// and a boolean to check if the value has been set.
func (o *RemoteAccessConfig) GetRemoteAccessIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAccessIpAddress, true
}

// SetRemoteAccessIpAddress sets field value
func (o *RemoteAccessConfig) SetRemoteAccessIpAddress(v string) {
	o.RemoteAccessIpAddress = v
}

// GetRemoteAccessPort returns the RemoteAccessPort field value
func (o *RemoteAccessConfig) GetRemoteAccessPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemoteAccessPort
}

// GetRemoteAccessPortOk returns a tuple with the RemoteAccessPort field value
// and a boolean to check if the value has been set.
func (o *RemoteAccessConfig) GetRemoteAccessPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAccessPort, true
}

// SetRemoteAccessPort sets field value
func (o *RemoteAccessConfig) SetRemoteAccessPort(v int32) {
	o.RemoteAccessPort = v
}

// GetInstanceType returns the InstanceType field value
func (o *RemoteAccessConfig) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *RemoteAccessConfig) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *RemoteAccessConfig) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *RemoteAccessConfig) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAccessConfig) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *RemoteAccessConfig) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *RemoteAccessConfig) SetTemplateName(v string) {
	o.TemplateName = &v
}

func (o RemoteAccessConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["remoteAccessIpAddress"] = o.RemoteAccessIpAddress
	}
	if true {
		toSerialize["remoteAccessPort"] = o.RemoteAccessPort
	}
	if true {
		toSerialize["instanceType"] = o.InstanceType
	}
	if o.TemplateName != nil {
		toSerialize["templateName"] = o.TemplateName
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteAccessConfig struct {
	value *RemoteAccessConfig
	isSet bool
}

func (v NullableRemoteAccessConfig) Get() *RemoteAccessConfig {
	return v.value
}

func (v *NullableRemoteAccessConfig) Set(val *RemoteAccessConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteAccessConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteAccessConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteAccessConfig(val *RemoteAccessConfig) *NullableRemoteAccessConfig {
	return &NullableRemoteAccessConfig{value: val, isSet: true}
}

func (v NullableRemoteAccessConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteAccessConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
