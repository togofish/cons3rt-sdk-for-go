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

// DeviceAllOf struct for DeviceAllOf
type DeviceAllOf struct {
	Hostname   *string `json:"hostname,omitempty"`
	IpAddress  *string `json:"ipAddress,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
}

// NewDeviceAllOf instantiates a new DeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAllOf() *DeviceAllOf {
	this := DeviceAllOf{}
	return &this
}

// NewDeviceAllOfWithDefaults instantiates a new DeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAllOfWithDefaults() *DeviceAllOf {
	this := DeviceAllOf{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DeviceAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DeviceAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DeviceAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *DeviceAllOf) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *DeviceAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *DeviceAllOf) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *DeviceAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *DeviceAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *DeviceAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o DeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.MacAddress != nil {
		toSerialize["macAddress"] = o.MacAddress
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAllOf struct {
	value *DeviceAllOf
	isSet bool
}

func (v NullableDeviceAllOf) Get() *DeviceAllOf {
	return v.value
}

func (v *NullableDeviceAllOf) Set(val *DeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAllOf(val *DeviceAllOf) *NullableDeviceAllOf {
	return &NullableDeviceAllOf{value: val, isSet: true}
}

func (v NullableDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
