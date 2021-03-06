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

// OpenStackCloudResourcesAllOf struct for OpenStackCloudResourcesAllOf
type OpenStackCloudResourcesAllOf struct {
	FlavorTypeNames *[]string `json:"flavorTypeNames,omitempty"`
	NatImageNames   *[]string `json:"natImageNames,omitempty"`
}

// NewOpenStackCloudResourcesAllOf instantiates a new OpenStackCloudResourcesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackCloudResourcesAllOf() *OpenStackCloudResourcesAllOf {
	this := OpenStackCloudResourcesAllOf{}
	return &this
}

// NewOpenStackCloudResourcesAllOfWithDefaults instantiates a new OpenStackCloudResourcesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackCloudResourcesAllOfWithDefaults() *OpenStackCloudResourcesAllOf {
	this := OpenStackCloudResourcesAllOf{}
	return &this
}

// GetFlavorTypeNames returns the FlavorTypeNames field value if set, zero value otherwise.
func (o *OpenStackCloudResourcesAllOf) GetFlavorTypeNames() []string {
	if o == nil || o.FlavorTypeNames == nil {
		var ret []string
		return ret
	}
	return *o.FlavorTypeNames
}

// GetFlavorTypeNamesOk returns a tuple with the FlavorTypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackCloudResourcesAllOf) GetFlavorTypeNamesOk() (*[]string, bool) {
	if o == nil || o.FlavorTypeNames == nil {
		return nil, false
	}
	return o.FlavorTypeNames, true
}

// HasFlavorTypeNames returns a boolean if a field has been set.
func (o *OpenStackCloudResourcesAllOf) HasFlavorTypeNames() bool {
	if o != nil && o.FlavorTypeNames != nil {
		return true
	}

	return false
}

// SetFlavorTypeNames gets a reference to the given []string and assigns it to the FlavorTypeNames field.
func (o *OpenStackCloudResourcesAllOf) SetFlavorTypeNames(v []string) {
	o.FlavorTypeNames = &v
}

// GetNatImageNames returns the NatImageNames field value if set, zero value otherwise.
func (o *OpenStackCloudResourcesAllOf) GetNatImageNames() []string {
	if o == nil || o.NatImageNames == nil {
		var ret []string
		return ret
	}
	return *o.NatImageNames
}

// GetNatImageNamesOk returns a tuple with the NatImageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackCloudResourcesAllOf) GetNatImageNamesOk() (*[]string, bool) {
	if o == nil || o.NatImageNames == nil {
		return nil, false
	}
	return o.NatImageNames, true
}

// HasNatImageNames returns a boolean if a field has been set.
func (o *OpenStackCloudResourcesAllOf) HasNatImageNames() bool {
	if o != nil && o.NatImageNames != nil {
		return true
	}

	return false
}

// SetNatImageNames gets a reference to the given []string and assigns it to the NatImageNames field.
func (o *OpenStackCloudResourcesAllOf) SetNatImageNames(v []string) {
	o.NatImageNames = &v
}

func (o OpenStackCloudResourcesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlavorTypeNames != nil {
		toSerialize["flavorTypeNames"] = o.FlavorTypeNames
	}
	if o.NatImageNames != nil {
		toSerialize["natImageNames"] = o.NatImageNames
	}
	return json.Marshal(toSerialize)
}

type NullableOpenStackCloudResourcesAllOf struct {
	value *OpenStackCloudResourcesAllOf
	isSet bool
}

func (v NullableOpenStackCloudResourcesAllOf) Get() *OpenStackCloudResourcesAllOf {
	return v.value
}

func (v *NullableOpenStackCloudResourcesAllOf) Set(val *OpenStackCloudResourcesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackCloudResourcesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackCloudResourcesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackCloudResourcesAllOf(val *OpenStackCloudResourcesAllOf) *NullableOpenStackCloudResourcesAllOf {
	return &NullableOpenStackCloudResourcesAllOf{value: val, isSet: true}
}

func (v NullableOpenStackCloudResourcesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackCloudResourcesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
