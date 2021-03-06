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

// FullVirtualServiceAllOf struct for FullVirtualServiceAllOf
type FullVirtualServiceAllOf struct {
	BasePath         *string `json:"basePath,omitempty"`
	PortString       *string `json:"portString,omitempty"`
	ServiceImageFile *string `json:"serviceImageFile,omitempty"`
	ServiceModelFile *string `json:"serviceModelFile,omitempty"`
}

// NewFullVirtualServiceAllOf instantiates a new FullVirtualServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullVirtualServiceAllOf() *FullVirtualServiceAllOf {
	this := FullVirtualServiceAllOf{}
	return &this
}

// NewFullVirtualServiceAllOfWithDefaults instantiates a new FullVirtualServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullVirtualServiceAllOfWithDefaults() *FullVirtualServiceAllOf {
	this := FullVirtualServiceAllOf{}
	return &this
}

// GetBasePath returns the BasePath field value if set, zero value otherwise.
func (o *FullVirtualServiceAllOf) GetBasePath() string {
	if o == nil || o.BasePath == nil {
		var ret string
		return ret
	}
	return *o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVirtualServiceAllOf) GetBasePathOk() (*string, bool) {
	if o == nil || o.BasePath == nil {
		return nil, false
	}
	return o.BasePath, true
}

// HasBasePath returns a boolean if a field has been set.
func (o *FullVirtualServiceAllOf) HasBasePath() bool {
	if o != nil && o.BasePath != nil {
		return true
	}

	return false
}

// SetBasePath gets a reference to the given string and assigns it to the BasePath field.
func (o *FullVirtualServiceAllOf) SetBasePath(v string) {
	o.BasePath = &v
}

// GetPortString returns the PortString field value if set, zero value otherwise.
func (o *FullVirtualServiceAllOf) GetPortString() string {
	if o == nil || o.PortString == nil {
		var ret string
		return ret
	}
	return *o.PortString
}

// GetPortStringOk returns a tuple with the PortString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVirtualServiceAllOf) GetPortStringOk() (*string, bool) {
	if o == nil || o.PortString == nil {
		return nil, false
	}
	return o.PortString, true
}

// HasPortString returns a boolean if a field has been set.
func (o *FullVirtualServiceAllOf) HasPortString() bool {
	if o != nil && o.PortString != nil {
		return true
	}

	return false
}

// SetPortString gets a reference to the given string and assigns it to the PortString field.
func (o *FullVirtualServiceAllOf) SetPortString(v string) {
	o.PortString = &v
}

// GetServiceImageFile returns the ServiceImageFile field value if set, zero value otherwise.
func (o *FullVirtualServiceAllOf) GetServiceImageFile() string {
	if o == nil || o.ServiceImageFile == nil {
		var ret string
		return ret
	}
	return *o.ServiceImageFile
}

// GetServiceImageFileOk returns a tuple with the ServiceImageFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVirtualServiceAllOf) GetServiceImageFileOk() (*string, bool) {
	if o == nil || o.ServiceImageFile == nil {
		return nil, false
	}
	return o.ServiceImageFile, true
}

// HasServiceImageFile returns a boolean if a field has been set.
func (o *FullVirtualServiceAllOf) HasServiceImageFile() bool {
	if o != nil && o.ServiceImageFile != nil {
		return true
	}

	return false
}

// SetServiceImageFile gets a reference to the given string and assigns it to the ServiceImageFile field.
func (o *FullVirtualServiceAllOf) SetServiceImageFile(v string) {
	o.ServiceImageFile = &v
}

// GetServiceModelFile returns the ServiceModelFile field value if set, zero value otherwise.
func (o *FullVirtualServiceAllOf) GetServiceModelFile() string {
	if o == nil || o.ServiceModelFile == nil {
		var ret string
		return ret
	}
	return *o.ServiceModelFile
}

// GetServiceModelFileOk returns a tuple with the ServiceModelFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullVirtualServiceAllOf) GetServiceModelFileOk() (*string, bool) {
	if o == nil || o.ServiceModelFile == nil {
		return nil, false
	}
	return o.ServiceModelFile, true
}

// HasServiceModelFile returns a boolean if a field has been set.
func (o *FullVirtualServiceAllOf) HasServiceModelFile() bool {
	if o != nil && o.ServiceModelFile != nil {
		return true
	}

	return false
}

// SetServiceModelFile gets a reference to the given string and assigns it to the ServiceModelFile field.
func (o *FullVirtualServiceAllOf) SetServiceModelFile(v string) {
	o.ServiceModelFile = &v
}

func (o FullVirtualServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BasePath != nil {
		toSerialize["basePath"] = o.BasePath
	}
	if o.PortString != nil {
		toSerialize["portString"] = o.PortString
	}
	if o.ServiceImageFile != nil {
		toSerialize["serviceImageFile"] = o.ServiceImageFile
	}
	if o.ServiceModelFile != nil {
		toSerialize["serviceModelFile"] = o.ServiceModelFile
	}
	return json.Marshal(toSerialize)
}

type NullableFullVirtualServiceAllOf struct {
	value *FullVirtualServiceAllOf
	isSet bool
}

func (v NullableFullVirtualServiceAllOf) Get() *FullVirtualServiceAllOf {
	return v.value
}

func (v *NullableFullVirtualServiceAllOf) Set(val *FullVirtualServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFullVirtualServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFullVirtualServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullVirtualServiceAllOf(val *FullVirtualServiceAllOf) *NullableFullVirtualServiceAllOf {
	return &NullableFullVirtualServiceAllOf{value: val, isSet: true}
}

func (v NullableFullVirtualServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullVirtualServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
