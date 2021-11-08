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

// InputCompositionHostOption struct for InputCompositionHostOption
type InputCompositionHostOption struct {
	SystemRole      string       `json:"systemRole"`
	Cpus            *int32       `json:"cpus,omitempty"`
	Ram             *int32       `json:"ram,omitempty"`
	AdditionalDisks *[]InputDisk `json:"additionalDisks,omitempty"`
	// Required for Azure and AWS EC2 Cloudspaces
	InstanceTypeName  *string                  `json:"instanceTypeName,omitempty"`
	NetworkInterfaces *[]InputNetworkInterface `json:"networkInterfaces,omitempty"`
	TemplateName      string                   `json:"templateName"`
}

// NewInputCompositionHostOption instantiates a new InputCompositionHostOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCompositionHostOption(systemRole string, templateName string) *InputCompositionHostOption {
	this := InputCompositionHostOption{}
	this.SystemRole = systemRole
	this.TemplateName = templateName
	return &this
}

// NewInputCompositionHostOptionWithDefaults instantiates a new InputCompositionHostOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCompositionHostOptionWithDefaults() *InputCompositionHostOption {
	this := InputCompositionHostOption{}
	return &this
}

// GetSystemRole returns the SystemRole field value
func (o *InputCompositionHostOption) GetSystemRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetSystemRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemRole, true
}

// SetSystemRole sets field value
func (o *InputCompositionHostOption) SetSystemRole(v string) {
	o.SystemRole = v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *InputCompositionHostOption) GetCpus() int32 {
	if o == nil || o.Cpus == nil {
		var ret int32
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetCpusOk() (*int32, bool) {
	if o == nil || o.Cpus == nil {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *InputCompositionHostOption) HasCpus() bool {
	if o != nil && o.Cpus != nil {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int32 and assigns it to the Cpus field.
func (o *InputCompositionHostOption) SetCpus(v int32) {
	o.Cpus = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *InputCompositionHostOption) GetRam() int32 {
	if o == nil || o.Ram == nil {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetRamOk() (*int32, bool) {
	if o == nil || o.Ram == nil {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *InputCompositionHostOption) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *InputCompositionHostOption) SetRam(v int32) {
	o.Ram = &v
}

// GetAdditionalDisks returns the AdditionalDisks field value if set, zero value otherwise.
func (o *InputCompositionHostOption) GetAdditionalDisks() []InputDisk {
	if o == nil || o.AdditionalDisks == nil {
		var ret []InputDisk
		return ret
	}
	return *o.AdditionalDisks
}

// GetAdditionalDisksOk returns a tuple with the AdditionalDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetAdditionalDisksOk() (*[]InputDisk, bool) {
	if o == nil || o.AdditionalDisks == nil {
		return nil, false
	}
	return o.AdditionalDisks, true
}

// HasAdditionalDisks returns a boolean if a field has been set.
func (o *InputCompositionHostOption) HasAdditionalDisks() bool {
	if o != nil && o.AdditionalDisks != nil {
		return true
	}

	return false
}

// SetAdditionalDisks gets a reference to the given []InputDisk and assigns it to the AdditionalDisks field.
func (o *InputCompositionHostOption) SetAdditionalDisks(v []InputDisk) {
	o.AdditionalDisks = &v
}

// GetInstanceTypeName returns the InstanceTypeName field value if set, zero value otherwise.
func (o *InputCompositionHostOption) GetInstanceTypeName() string {
	if o == nil || o.InstanceTypeName == nil {
		var ret string
		return ret
	}
	return *o.InstanceTypeName
}

// GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetInstanceTypeNameOk() (*string, bool) {
	if o == nil || o.InstanceTypeName == nil {
		return nil, false
	}
	return o.InstanceTypeName, true
}

// HasInstanceTypeName returns a boolean if a field has been set.
func (o *InputCompositionHostOption) HasInstanceTypeName() bool {
	if o != nil && o.InstanceTypeName != nil {
		return true
	}

	return false
}

// SetInstanceTypeName gets a reference to the given string and assigns it to the InstanceTypeName field.
func (o *InputCompositionHostOption) SetInstanceTypeName(v string) {
	o.InstanceTypeName = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *InputCompositionHostOption) GetNetworkInterfaces() []InputNetworkInterface {
	if o == nil || o.NetworkInterfaces == nil {
		var ret []InputNetworkInterface
		return ret
	}
	return *o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetNetworkInterfacesOk() (*[]InputNetworkInterface, bool) {
	if o == nil || o.NetworkInterfaces == nil {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *InputCompositionHostOption) HasNetworkInterfaces() bool {
	if o != nil && o.NetworkInterfaces != nil {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []InputNetworkInterface and assigns it to the NetworkInterfaces field.
func (o *InputCompositionHostOption) SetNetworkInterfaces(v []InputNetworkInterface) {
	o.NetworkInterfaces = &v
}

// GetTemplateName returns the TemplateName field value
func (o *InputCompositionHostOption) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *InputCompositionHostOption) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *InputCompositionHostOption) SetTemplateName(v string) {
	o.TemplateName = v
}

func (o InputCompositionHostOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["systemRole"] = o.SystemRole
	}
	if o.Cpus != nil {
		toSerialize["cpus"] = o.Cpus
	}
	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}
	if o.AdditionalDisks != nil {
		toSerialize["additionalDisks"] = o.AdditionalDisks
	}
	if o.InstanceTypeName != nil {
		toSerialize["instanceTypeName"] = o.InstanceTypeName
	}
	if o.NetworkInterfaces != nil {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if true {
		toSerialize["templateName"] = o.TemplateName
	}
	return json.Marshal(toSerialize)
}

type NullableInputCompositionHostOption struct {
	value *InputCompositionHostOption
	isSet bool
}

func (v NullableInputCompositionHostOption) Get() *InputCompositionHostOption {
	return v.value
}

func (v *NullableInputCompositionHostOption) Set(val *InputCompositionHostOption) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCompositionHostOption) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCompositionHostOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCompositionHostOption(val *InputCompositionHostOption) *NullableInputCompositionHostOption {
	return &NullableInputCompositionHostOption{value: val, isSet: true}
}

func (v NullableInputCompositionHostOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCompositionHostOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
