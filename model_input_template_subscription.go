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

// InputTemplateSubscription struct for InputTemplateSubscription
type InputTemplateSubscription struct {
	State *string `json:"state,omitempty"`
	AllowGpuUsage *bool `json:"allowGpuUsage,omitempty"`
	MaxNumCpus *int32 `json:"maxNumCpus,omitempty"`
	MaxRamInMegabytes *int32 `json:"maxRamInMegabytes,omitempty"`
}

// NewInputTemplateSubscription instantiates a new InputTemplateSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputTemplateSubscription() *InputTemplateSubscription {
	this := InputTemplateSubscription{}
	return &this
}

// NewInputTemplateSubscriptionWithDefaults instantiates a new InputTemplateSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputTemplateSubscriptionWithDefaults() *InputTemplateSubscription {
	this := InputTemplateSubscription{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InputTemplateSubscription) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateSubscription) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InputTemplateSubscription) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *InputTemplateSubscription) SetState(v string) {
	o.State = &v
}

// GetAllowGpuUsage returns the AllowGpuUsage field value if set, zero value otherwise.
func (o *InputTemplateSubscription) GetAllowGpuUsage() bool {
	if o == nil || o.AllowGpuUsage == nil {
		var ret bool
		return ret
	}
	return *o.AllowGpuUsage
}

// GetAllowGpuUsageOk returns a tuple with the AllowGpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateSubscription) GetAllowGpuUsageOk() (*bool, bool) {
	if o == nil || o.AllowGpuUsage == nil {
		return nil, false
	}
	return o.AllowGpuUsage, true
}

// HasAllowGpuUsage returns a boolean if a field has been set.
func (o *InputTemplateSubscription) HasAllowGpuUsage() bool {
	if o != nil && o.AllowGpuUsage != nil {
		return true
	}

	return false
}

// SetAllowGpuUsage gets a reference to the given bool and assigns it to the AllowGpuUsage field.
func (o *InputTemplateSubscription) SetAllowGpuUsage(v bool) {
	o.AllowGpuUsage = &v
}

// GetMaxNumCpus returns the MaxNumCpus field value if set, zero value otherwise.
func (o *InputTemplateSubscription) GetMaxNumCpus() int32 {
	if o == nil || o.MaxNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumCpus
}

// GetMaxNumCpusOk returns a tuple with the MaxNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateSubscription) GetMaxNumCpusOk() (*int32, bool) {
	if o == nil || o.MaxNumCpus == nil {
		return nil, false
	}
	return o.MaxNumCpus, true
}

// HasMaxNumCpus returns a boolean if a field has been set.
func (o *InputTemplateSubscription) HasMaxNumCpus() bool {
	if o != nil && o.MaxNumCpus != nil {
		return true
	}

	return false
}

// SetMaxNumCpus gets a reference to the given int32 and assigns it to the MaxNumCpus field.
func (o *InputTemplateSubscription) SetMaxNumCpus(v int32) {
	o.MaxNumCpus = &v
}

// GetMaxRamInMegabytes returns the MaxRamInMegabytes field value if set, zero value otherwise.
func (o *InputTemplateSubscription) GetMaxRamInMegabytes() int32 {
	if o == nil || o.MaxRamInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRamInMegabytes
}

// GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTemplateSubscription) GetMaxRamInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaxRamInMegabytes == nil {
		return nil, false
	}
	return o.MaxRamInMegabytes, true
}

// HasMaxRamInMegabytes returns a boolean if a field has been set.
func (o *InputTemplateSubscription) HasMaxRamInMegabytes() bool {
	if o != nil && o.MaxRamInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxRamInMegabytes gets a reference to the given int32 and assigns it to the MaxRamInMegabytes field.
func (o *InputTemplateSubscription) SetMaxRamInMegabytes(v int32) {
	o.MaxRamInMegabytes = &v
}

func (o InputTemplateSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.AllowGpuUsage != nil {
		toSerialize["allowGpuUsage"] = o.AllowGpuUsage
	}
	if o.MaxNumCpus != nil {
		toSerialize["maxNumCpus"] = o.MaxNumCpus
	}
	if o.MaxRamInMegabytes != nil {
		toSerialize["maxRamInMegabytes"] = o.MaxRamInMegabytes
	}
	return json.Marshal(toSerialize)
}

type NullableInputTemplateSubscription struct {
	value *InputTemplateSubscription
	isSet bool
}

func (v NullableInputTemplateSubscription) Get() *InputTemplateSubscription {
	return v.value
}

func (v *NullableInputTemplateSubscription) Set(val *InputTemplateSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTemplateSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTemplateSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTemplateSubscription(val *InputTemplateSubscription) *NullableInputTemplateSubscription {
	return &NullableInputTemplateSubscription{value: val, isSet: true}
}

func (v NullableInputTemplateSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputTemplateSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


