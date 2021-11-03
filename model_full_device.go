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

// FullDevice struct for FullDevice
type FullDevice struct {
	Components *[]MinimalAbstractComponent `json:"components,omitempty"`
	Hostname   *string                     `json:"hostname,omitempty"`
	IpAddress  *string                     `json:"ipAddress,omitempty"`
	MacAddress *string                     `json:"macAddress,omitempty"`
}

// NewFullDevice instantiates a new FullDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDevice(subtype string) *FullDevice {
	this := FullDevice{}
	this.Subtype = subtype
	return &this
}

// NewFullDeviceWithDefaults instantiates a new FullDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDeviceWithDefaults() *FullDevice {
	this := FullDevice{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *FullDevice) GetComponents() []MinimalAbstractComponent {
	if o == nil || o.Components == nil {
		var ret []MinimalAbstractComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDevice) GetComponentsOk() (*[]MinimalAbstractComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *FullDevice) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []MinimalAbstractComponent and assigns it to the Components field.
func (o *FullDevice) SetComponents(v []MinimalAbstractComponent) {
	o.Components = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FullDevice) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDevice) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FullDevice) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FullDevice) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *FullDevice) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDevice) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *FullDevice) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *FullDevice) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *FullDevice) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDevice) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *FullDevice) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *FullDevice) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o FullDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
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

type NullableFullDevice struct {
	value *FullDevice
	isSet bool
}

func (v NullableFullDevice) Get() *FullDevice {
	return v.value
}

func (v *NullableFullDevice) Set(val *FullDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDevice(val *FullDevice) *NullableFullDevice {
	return &NullableFullDevice{value: val, isSet: true}
}

func (v NullableFullDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
