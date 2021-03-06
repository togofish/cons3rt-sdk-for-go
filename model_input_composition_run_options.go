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

// InputCompositionRunOptions struct for InputCompositionRunOptions
type InputCompositionRunOptions struct {
	Description           *string                       `json:"description,omitempty"`
	VirtualizationRealmId int32                         `json:"virtualizationRealmId"`
	Properties            *[]InputProperty              `json:"properties,omitempty"`
	HostOptions           *[]InputCompositionHostOption `json:"hostOptions,omitempty"`
	WindowsDomainName     *string                       `json:"windowsDomainName,omitempty"`
}

// NewInputCompositionRunOptions instantiates a new InputCompositionRunOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCompositionRunOptions(virtualizationRealmId int32) *InputCompositionRunOptions {
	this := InputCompositionRunOptions{}
	this.VirtualizationRealmId = virtualizationRealmId
	return &this
}

// NewInputCompositionRunOptionsWithDefaults instantiates a new InputCompositionRunOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCompositionRunOptionsWithDefaults() *InputCompositionRunOptions {
	this := InputCompositionRunOptions{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputCompositionRunOptions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionRunOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputCompositionRunOptions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputCompositionRunOptions) SetDescription(v string) {
	o.Description = &v
}

// GetVirtualizationRealmId returns the VirtualizationRealmId field value
func (o *InputCompositionRunOptions) GetVirtualizationRealmId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualizationRealmId
}

// GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field value
// and a boolean to check if the value has been set.
func (o *InputCompositionRunOptions) GetVirtualizationRealmIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualizationRealmId, true
}

// SetVirtualizationRealmId sets field value
func (o *InputCompositionRunOptions) SetVirtualizationRealmId(v int32) {
	o.VirtualizationRealmId = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *InputCompositionRunOptions) GetProperties() []InputProperty {
	if o == nil || o.Properties == nil {
		var ret []InputProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionRunOptions) GetPropertiesOk() (*[]InputProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *InputCompositionRunOptions) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []InputProperty and assigns it to the Properties field.
func (o *InputCompositionRunOptions) SetProperties(v []InputProperty) {
	o.Properties = &v
}

// GetHostOptions returns the HostOptions field value if set, zero value otherwise.
func (o *InputCompositionRunOptions) GetHostOptions() []InputCompositionHostOption {
	if o == nil || o.HostOptions == nil {
		var ret []InputCompositionHostOption
		return ret
	}
	return *o.HostOptions
}

// GetHostOptionsOk returns a tuple with the HostOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionRunOptions) GetHostOptionsOk() (*[]InputCompositionHostOption, bool) {
	if o == nil || o.HostOptions == nil {
		return nil, false
	}
	return o.HostOptions, true
}

// HasHostOptions returns a boolean if a field has been set.
func (o *InputCompositionRunOptions) HasHostOptions() bool {
	if o != nil && o.HostOptions != nil {
		return true
	}

	return false
}

// SetHostOptions gets a reference to the given []InputCompositionHostOption and assigns it to the HostOptions field.
func (o *InputCompositionRunOptions) SetHostOptions(v []InputCompositionHostOption) {
	o.HostOptions = &v
}

// GetWindowsDomainName returns the WindowsDomainName field value if set, zero value otherwise.
func (o *InputCompositionRunOptions) GetWindowsDomainName() string {
	if o == nil || o.WindowsDomainName == nil {
		var ret string
		return ret
	}
	return *o.WindowsDomainName
}

// GetWindowsDomainNameOk returns a tuple with the WindowsDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionRunOptions) GetWindowsDomainNameOk() (*string, bool) {
	if o == nil || o.WindowsDomainName == nil {
		return nil, false
	}
	return o.WindowsDomainName, true
}

// HasWindowsDomainName returns a boolean if a field has been set.
func (o *InputCompositionRunOptions) HasWindowsDomainName() bool {
	if o != nil && o.WindowsDomainName != nil {
		return true
	}

	return false
}

// SetWindowsDomainName gets a reference to the given string and assigns it to the WindowsDomainName field.
func (o *InputCompositionRunOptions) SetWindowsDomainName(v string) {
	o.WindowsDomainName = &v
}

func (o InputCompositionRunOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["virtualizationRealmId"] = o.VirtualizationRealmId
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.HostOptions != nil {
		toSerialize["hostOptions"] = o.HostOptions
	}
	if o.WindowsDomainName != nil {
		toSerialize["windowsDomainName"] = o.WindowsDomainName
	}
	return json.Marshal(toSerialize)
}

type NullableInputCompositionRunOptions struct {
	value *InputCompositionRunOptions
	isSet bool
}

func (v NullableInputCompositionRunOptions) Get() *InputCompositionRunOptions {
	return v.value
}

func (v *NullableInputCompositionRunOptions) Set(val *InputCompositionRunOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCompositionRunOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCompositionRunOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCompositionRunOptions(val *InputCompositionRunOptions) *NullableInputCompositionRunOptions {
	return &NullableInputCompositionRunOptions{value: val, isSet: true}
}

func (v NullableInputCompositionRunOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCompositionRunOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
