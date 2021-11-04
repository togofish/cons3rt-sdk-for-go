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

// AbstractCloudSpaceRequest struct for AbstractCloudSpaceRequest
type AbstractCloudSpaceRequest struct {
	CloudSpaceName string `json:"cloudSpaceName"`
	MaximumVirtualMachines *int32 `json:"maximumVirtualMachines,omitempty"`
	Cidr string `json:"cidr"`
	NumAvailabilityZones *int32 `json:"numAvailabilityZones,omitempty"`
	PowerOnMaximumDelay *int32 `json:"powerOnMaximumDelay,omitempty"`
	PowerOnMinimumDelay *int32 `json:"powerOnMinimumDelay,omitempty"`
	Subtype string `json:"subtype"`
}

// NewAbstractCloudSpaceRequest instantiates a new AbstractCloudSpaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractCloudSpaceRequest(cloudSpaceName string, cidr string, subtype string) *AbstractCloudSpaceRequest {
	this := AbstractCloudSpaceRequest{}
	this.CloudSpaceName = cloudSpaceName
	this.Cidr = cidr
	this.Subtype = subtype
	return &this
}

// NewAbstractCloudSpaceRequestWithDefaults instantiates a new AbstractCloudSpaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractCloudSpaceRequestWithDefaults() *AbstractCloudSpaceRequest {
	this := AbstractCloudSpaceRequest{}
	return &this
}

// GetCloudSpaceName returns the CloudSpaceName field value
func (o *AbstractCloudSpaceRequest) GetCloudSpaceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudSpaceName
}

// GetCloudSpaceNameOk returns a tuple with the CloudSpaceName field value
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetCloudSpaceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudSpaceName, true
}

// SetCloudSpaceName sets field value
func (o *AbstractCloudSpaceRequest) SetCloudSpaceName(v string) {
	o.CloudSpaceName = v
}

// GetMaximumVirtualMachines returns the MaximumVirtualMachines field value if set, zero value otherwise.
func (o *AbstractCloudSpaceRequest) GetMaximumVirtualMachines() int32 {
	if o == nil || o.MaximumVirtualMachines == nil {
		var ret int32
		return ret
	}
	return *o.MaximumVirtualMachines
}

// GetMaximumVirtualMachinesOk returns a tuple with the MaximumVirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetMaximumVirtualMachinesOk() (*int32, bool) {
	if o == nil || o.MaximumVirtualMachines == nil {
		return nil, false
	}
	return o.MaximumVirtualMachines, true
}

// HasMaximumVirtualMachines returns a boolean if a field has been set.
func (o *AbstractCloudSpaceRequest) HasMaximumVirtualMachines() bool {
	if o != nil && o.MaximumVirtualMachines != nil {
		return true
	}

	return false
}

// SetMaximumVirtualMachines gets a reference to the given int32 and assigns it to the MaximumVirtualMachines field.
func (o *AbstractCloudSpaceRequest) SetMaximumVirtualMachines(v int32) {
	o.MaximumVirtualMachines = &v
}

// GetCidr returns the Cidr field value
func (o *AbstractCloudSpaceRequest) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetCidrOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *AbstractCloudSpaceRequest) SetCidr(v string) {
	o.Cidr = v
}

// GetNumAvailabilityZones returns the NumAvailabilityZones field value if set, zero value otherwise.
func (o *AbstractCloudSpaceRequest) GetNumAvailabilityZones() int32 {
	if o == nil || o.NumAvailabilityZones == nil {
		var ret int32
		return ret
	}
	return *o.NumAvailabilityZones
}

// GetNumAvailabilityZonesOk returns a tuple with the NumAvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetNumAvailabilityZonesOk() (*int32, bool) {
	if o == nil || o.NumAvailabilityZones == nil {
		return nil, false
	}
	return o.NumAvailabilityZones, true
}

// HasNumAvailabilityZones returns a boolean if a field has been set.
func (o *AbstractCloudSpaceRequest) HasNumAvailabilityZones() bool {
	if o != nil && o.NumAvailabilityZones != nil {
		return true
	}

	return false
}

// SetNumAvailabilityZones gets a reference to the given int32 and assigns it to the NumAvailabilityZones field.
func (o *AbstractCloudSpaceRequest) SetNumAvailabilityZones(v int32) {
	o.NumAvailabilityZones = &v
}

// GetPowerOnMaximumDelay returns the PowerOnMaximumDelay field value if set, zero value otherwise.
func (o *AbstractCloudSpaceRequest) GetPowerOnMaximumDelay() int32 {
	if o == nil || o.PowerOnMaximumDelay == nil {
		var ret int32
		return ret
	}
	return *o.PowerOnMaximumDelay
}

// GetPowerOnMaximumDelayOk returns a tuple with the PowerOnMaximumDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetPowerOnMaximumDelayOk() (*int32, bool) {
	if o == nil || o.PowerOnMaximumDelay == nil {
		return nil, false
	}
	return o.PowerOnMaximumDelay, true
}

// HasPowerOnMaximumDelay returns a boolean if a field has been set.
func (o *AbstractCloudSpaceRequest) HasPowerOnMaximumDelay() bool {
	if o != nil && o.PowerOnMaximumDelay != nil {
		return true
	}

	return false
}

// SetPowerOnMaximumDelay gets a reference to the given int32 and assigns it to the PowerOnMaximumDelay field.
func (o *AbstractCloudSpaceRequest) SetPowerOnMaximumDelay(v int32) {
	o.PowerOnMaximumDelay = &v
}

// GetPowerOnMinimumDelay returns the PowerOnMinimumDelay field value if set, zero value otherwise.
func (o *AbstractCloudSpaceRequest) GetPowerOnMinimumDelay() int32 {
	if o == nil || o.PowerOnMinimumDelay == nil {
		var ret int32
		return ret
	}
	return *o.PowerOnMinimumDelay
}

// GetPowerOnMinimumDelayOk returns a tuple with the PowerOnMinimumDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetPowerOnMinimumDelayOk() (*int32, bool) {
	if o == nil || o.PowerOnMinimumDelay == nil {
		return nil, false
	}
	return o.PowerOnMinimumDelay, true
}

// HasPowerOnMinimumDelay returns a boolean if a field has been set.
func (o *AbstractCloudSpaceRequest) HasPowerOnMinimumDelay() bool {
	if o != nil && o.PowerOnMinimumDelay != nil {
		return true
	}

	return false
}

// SetPowerOnMinimumDelay gets a reference to the given int32 and assigns it to the PowerOnMinimumDelay field.
func (o *AbstractCloudSpaceRequest) SetPowerOnMinimumDelay(v int32) {
	o.PowerOnMinimumDelay = &v
}

// GetSubtype returns the Subtype field value
func (o *AbstractCloudSpaceRequest) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *AbstractCloudSpaceRequest) GetSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *AbstractCloudSpaceRequest) SetSubtype(v string) {
	o.Subtype = v
}

func (o AbstractCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudSpaceName"] = o.CloudSpaceName
	}
	if o.MaximumVirtualMachines != nil {
		toSerialize["maximumVirtualMachines"] = o.MaximumVirtualMachines
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if o.NumAvailabilityZones != nil {
		toSerialize["numAvailabilityZones"] = o.NumAvailabilityZones
	}
	if o.PowerOnMaximumDelay != nil {
		toSerialize["powerOnMaximumDelay"] = o.PowerOnMaximumDelay
	}
	if o.PowerOnMinimumDelay != nil {
		toSerialize["powerOnMinimumDelay"] = o.PowerOnMinimumDelay
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableAbstractCloudSpaceRequest struct {
	value *AbstractCloudSpaceRequest
	isSet bool
}

func (v NullableAbstractCloudSpaceRequest) Get() *AbstractCloudSpaceRequest {
	return v.value
}

func (v *NullableAbstractCloudSpaceRequest) Set(val *AbstractCloudSpaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractCloudSpaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractCloudSpaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractCloudSpaceRequest(val *AbstractCloudSpaceRequest) *NullableAbstractCloudSpaceRequest {
	return &NullableAbstractCloudSpaceRequest{value: val, isSet: true}
}

func (v NullableAbstractCloudSpaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractCloudSpaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


