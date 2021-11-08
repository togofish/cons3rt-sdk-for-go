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

// FullVCloudRestCloudAllOf struct for FullVCloudRestCloudAllOf
type FullVCloudRestCloudAllOf struct {
	GpuProfileTypes        *[]string `json:"gpuProfileTypes,omitempty"`
	Hostname               *string   `json:"hostname,omitempty"`
	Port                   *int32    `json:"port,omitempty"`
	Protocol               *string   `json:"protocol,omitempty"`
	RemoteAccessInternalIp *string   `json:"remoteAccessInternalIp,omitempty"`
	RemoteAccessPort       *int32    `json:"remoteAccessPort,omitempty"`
	Username               *string   `json:"username,omitempty"`
	VsphereApiVersion      *string   `json:"vsphereApiVersion,omitempty"`
	VsphereHost            *string   `json:"vsphereHost,omitempty"`
	VspherePort            *int32    `json:"vspherePort,omitempty"`
}

// NewFullVCloudRestCloudAllOf instantiates a new FullVCloudRestCloudAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullVCloudRestCloudAllOf() *FullVCloudRestCloudAllOf {
	this := FullVCloudRestCloudAllOf{}
	return &this
}

// NewFullVCloudRestCloudAllOfWithDefaults instantiates a new FullVCloudRestCloudAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullVCloudRestCloudAllOfWithDefaults() *FullVCloudRestCloudAllOf {
	this := FullVCloudRestCloudAllOf{}
	return &this
}

// GetGpuProfileTypes returns the GpuProfileTypes field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetGpuProfileTypes() []string {
	if o == nil || o.GpuProfileTypes == nil {
		var ret []string
		return ret
	}
	return *o.GpuProfileTypes
}

// GetGpuProfileTypesOk returns a tuple with the GpuProfileTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetGpuProfileTypesOk() (*[]string, bool) {
	if o == nil || o.GpuProfileTypes == nil {
		return nil, false
	}
	return o.GpuProfileTypes, true
}

// HasGpuProfileTypes returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasGpuProfileTypes() bool {
	if o != nil && o.GpuProfileTypes != nil {
		return true
	}

	return false
}

// SetGpuProfileTypes gets a reference to the given []string and assigns it to the GpuProfileTypes field.
func (o *FullVCloudRestCloudAllOf) SetGpuProfileTypes(v []string) {
	o.GpuProfileTypes = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FullVCloudRestCloudAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *FullVCloudRestCloudAllOf) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *FullVCloudRestCloudAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRemoteAccessInternalIp returns the RemoteAccessInternalIp field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetRemoteAccessInternalIp() string {
	if o == nil || o.RemoteAccessInternalIp == nil {
		var ret string
		return ret
	}
	return *o.RemoteAccessInternalIp
}

// GetRemoteAccessInternalIpOk returns a tuple with the RemoteAccessInternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetRemoteAccessInternalIpOk() (*string, bool) {
	if o == nil || o.RemoteAccessInternalIp == nil {
		return nil, false
	}
	return o.RemoteAccessInternalIp, true
}

// HasRemoteAccessInternalIp returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasRemoteAccessInternalIp() bool {
	if o != nil && o.RemoteAccessInternalIp != nil {
		return true
	}

	return false
}

// SetRemoteAccessInternalIp gets a reference to the given string and assigns it to the RemoteAccessInternalIp field.
func (o *FullVCloudRestCloudAllOf) SetRemoteAccessInternalIp(v string) {
	o.RemoteAccessInternalIp = &v
}

// GetRemoteAccessPort returns the RemoteAccessPort field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetRemoteAccessPort() int32 {
	if o == nil || o.RemoteAccessPort == nil {
		var ret int32
		return ret
	}
	return *o.RemoteAccessPort
}

// GetRemoteAccessPortOk returns a tuple with the RemoteAccessPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetRemoteAccessPortOk() (*int32, bool) {
	if o == nil || o.RemoteAccessPort == nil {
		return nil, false
	}
	return o.RemoteAccessPort, true
}

// HasRemoteAccessPort returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasRemoteAccessPort() bool {
	if o != nil && o.RemoteAccessPort != nil {
		return true
	}

	return false
}

// SetRemoteAccessPort gets a reference to the given int32 and assigns it to the RemoteAccessPort field.
func (o *FullVCloudRestCloudAllOf) SetRemoteAccessPort(v int32) {
	o.RemoteAccessPort = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *FullVCloudRestCloudAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetVsphereApiVersion returns the VsphereApiVersion field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetVsphereApiVersion() string {
	if o == nil || o.VsphereApiVersion == nil {
		var ret string
		return ret
	}
	return *o.VsphereApiVersion
}

// GetVsphereApiVersionOk returns a tuple with the VsphereApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetVsphereApiVersionOk() (*string, bool) {
	if o == nil || o.VsphereApiVersion == nil {
		return nil, false
	}
	return o.VsphereApiVersion, true
}

// HasVsphereApiVersion returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasVsphereApiVersion() bool {
	if o != nil && o.VsphereApiVersion != nil {
		return true
	}

	return false
}

// SetVsphereApiVersion gets a reference to the given string and assigns it to the VsphereApiVersion field.
func (o *FullVCloudRestCloudAllOf) SetVsphereApiVersion(v string) {
	o.VsphereApiVersion = &v
}

// GetVsphereHost returns the VsphereHost field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetVsphereHost() string {
	if o == nil || o.VsphereHost == nil {
		var ret string
		return ret
	}
	return *o.VsphereHost
}

// GetVsphereHostOk returns a tuple with the VsphereHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetVsphereHostOk() (*string, bool) {
	if o == nil || o.VsphereHost == nil {
		return nil, false
	}
	return o.VsphereHost, true
}

// HasVsphereHost returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasVsphereHost() bool {
	if o != nil && o.VsphereHost != nil {
		return true
	}

	return false
}

// SetVsphereHost gets a reference to the given string and assigns it to the VsphereHost field.
func (o *FullVCloudRestCloudAllOf) SetVsphereHost(v string) {
	o.VsphereHost = &v
}

// GetVspherePort returns the VspherePort field value if set, zero value otherwise.
func (o *FullVCloudRestCloudAllOf) GetVspherePort() int32 {
	if o == nil || o.VspherePort == nil {
		var ret int32
		return ret
	}
	return *o.VspherePort
}

// GetVspherePortOk returns a tuple with the VspherePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVCloudRestCloudAllOf) GetVspherePortOk() (*int32, bool) {
	if o == nil || o.VspherePort == nil {
		return nil, false
	}
	return o.VspherePort, true
}

// HasVspherePort returns a boolean if a field has been set.
func (o *FullVCloudRestCloudAllOf) HasVspherePort() bool {
	if o != nil && o.VspherePort != nil {
		return true
	}

	return false
}

// SetVspherePort gets a reference to the given int32 and assigns it to the VspherePort field.
func (o *FullVCloudRestCloudAllOf) SetVspherePort(v int32) {
	o.VspherePort = &v
}

func (o FullVCloudRestCloudAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GpuProfileTypes != nil {
		toSerialize["gpuProfileTypes"] = o.GpuProfileTypes
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.RemoteAccessInternalIp != nil {
		toSerialize["remoteAccessInternalIp"] = o.RemoteAccessInternalIp
	}
	if o.RemoteAccessPort != nil {
		toSerialize["remoteAccessPort"] = o.RemoteAccessPort
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.VsphereApiVersion != nil {
		toSerialize["vsphereApiVersion"] = o.VsphereApiVersion
	}
	if o.VsphereHost != nil {
		toSerialize["vsphereHost"] = o.VsphereHost
	}
	if o.VspherePort != nil {
		toSerialize["vspherePort"] = o.VspherePort
	}
	return json.Marshal(toSerialize)
}

type NullableFullVCloudRestCloudAllOf struct {
	value *FullVCloudRestCloudAllOf
	isSet bool
}

func (v NullableFullVCloudRestCloudAllOf) Get() *FullVCloudRestCloudAllOf {
	return v.value
}

func (v *NullableFullVCloudRestCloudAllOf) Set(val *FullVCloudRestCloudAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullVCloudRestCloudAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullVCloudRestCloudAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullVCloudRestCloudAllOf(val *FullVCloudRestCloudAllOf) *NullableFullVCloudRestCloudAllOf {
	return &NullableFullVCloudRestCloudAllOf{value: val, isSet: true}
}

func (v NullableFullVCloudRestCloudAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullVCloudRestCloudAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
