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

// OpenStackCloudResources struct for OpenStackCloudResources
type OpenStackCloudResources struct {
	FlavorTypeNames *[]string `json:"flavorTypeNames,omitempty"`
	NatImageNames   *[]string `json:"natImageNames,omitempty"`
	Subtype         string
}

// NewOpenStackCloudResources instantiates a new OpenStackCloudResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackCloudResources(subtype string) *OpenStackCloudResources {
	this := OpenStackCloudResources{}
	this.Subtype = subtype
	return &this
}

// NewOpenStackCloudResourcesWithDefaults instantiates a new OpenStackCloudResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackCloudResourcesWithDefaults() *OpenStackCloudResources {
	this := OpenStackCloudResources{}
	return &this
}

// GetFlavorTypeNames returns the FlavorTypeNames field value if set, zero value otherwise.
func (o *OpenStackCloudResources) GetFlavorTypeNames() []string {
	if o == nil || o.FlavorTypeNames == nil {
		var ret []string
		return ret
	}
	return *o.FlavorTypeNames
}

// GetFlavorTypeNamesOk returns a tuple with the FlavorTypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackCloudResources) GetFlavorTypeNamesOk() (*[]string, bool) {
	if o == nil || o.FlavorTypeNames == nil {
		return nil, false
	}
	return o.FlavorTypeNames, true
}

// HasFlavorTypeNames returns a boolean if a field has been set.
func (o *OpenStackCloudResources) HasFlavorTypeNames() bool {
	if o != nil && o.FlavorTypeNames != nil {
		return true
	}

	return false
}

// SetFlavorTypeNames gets a reference to the given []string and assigns it to the FlavorTypeNames field.
func (o *OpenStackCloudResources) SetFlavorTypeNames(v []string) {
	o.FlavorTypeNames = &v
}

// GetNatImageNames returns the NatImageNames field value if set, zero value otherwise.
func (o *OpenStackCloudResources) GetNatImageNames() []string {
	if o == nil || o.NatImageNames == nil {
		var ret []string
		return ret
	}
	return *o.NatImageNames
}

// GetNatImageNamesOk returns a tuple with the NatImageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackCloudResources) GetNatImageNamesOk() (*[]string, bool) {
	if o == nil || o.NatImageNames == nil {
		return nil, false
	}
	return o.NatImageNames, true
}

// HasNatImageNames returns a boolean if a field has been set.
func (o *OpenStackCloudResources) HasNatImageNames() bool {
	if o != nil && o.NatImageNames != nil {
		return true
	}

	return false
}

// SetNatImageNames gets a reference to the given []string and assigns it to the NatImageNames field.
func (o *OpenStackCloudResources) SetNatImageNames(v []string) {
	o.NatImageNames = &v
}

func (o OpenStackCloudResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlavorTypeNames != nil {
		toSerialize["flavorTypeNames"] = o.FlavorTypeNames
	}
	if o.NatImageNames != nil {
		toSerialize["natImageNames"] = o.NatImageNames
	}
	return json.Marshal(toSerialize)
}

type NullableOpenStackCloudResources struct {
	value *OpenStackCloudResources
	isSet bool
}

func (v NullableOpenStackCloudResources) Get() *OpenStackCloudResources {
	return v.value
}

func (v *NullableOpenStackCloudResources) Set(val *OpenStackCloudResources) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackCloudResources) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackCloudResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackCloudResources(val *OpenStackCloudResources) *NullableOpenStackCloudResources {
	return &NullableOpenStackCloudResources{value: val, isSet: true}
}

func (v NullableOpenStackCloudResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackCloudResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
