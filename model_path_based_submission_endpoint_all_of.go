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

// PathBasedSubmissionEndpointAllOf struct for PathBasedSubmissionEndpointAllOf
type PathBasedSubmissionEndpointAllOf struct {
	RemotePath *string `json:"remotePath,omitempty"`
	Subtype    *string `json:"subtype,omitempty"`
}

// NewPathBasedSubmissionEndpointAllOf instantiates a new PathBasedSubmissionEndpointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathBasedSubmissionEndpointAllOf() *PathBasedSubmissionEndpointAllOf {
	this := PathBasedSubmissionEndpointAllOf{}
	return &this
}

// NewPathBasedSubmissionEndpointAllOfWithDefaults instantiates a new PathBasedSubmissionEndpointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathBasedSubmissionEndpointAllOfWithDefaults() *PathBasedSubmissionEndpointAllOf {
	this := PathBasedSubmissionEndpointAllOf{}
	return &this
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise.
func (o *PathBasedSubmissionEndpointAllOf) GetRemotePath() string {
	if o == nil || o.RemotePath == nil {
		var ret string
		return ret
	}
	return *o.RemotePath
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathBasedSubmissionEndpointAllOf) GetRemotePathOk() (*string, bool) {
	if o == nil || o.RemotePath == nil {
		return nil, false
	}
	return o.RemotePath, true
}

// HasRemotePath returns a boolean if a field has been set.
func (o *PathBasedSubmissionEndpointAllOf) HasRemotePath() bool {
	if o != nil && o.RemotePath != nil {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given string and assigns it to the RemotePath field.
func (o *PathBasedSubmissionEndpointAllOf) SetRemotePath(v string) {
	o.RemotePath = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *PathBasedSubmissionEndpointAllOf) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathBasedSubmissionEndpointAllOf) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *PathBasedSubmissionEndpointAllOf) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *PathBasedSubmissionEndpointAllOf) SetSubtype(v string) {
	o.Subtype = &v
}

func (o PathBasedSubmissionEndpointAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemotePath != nil {
		toSerialize["remotePath"] = o.RemotePath
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullablePathBasedSubmissionEndpointAllOf struct {
	value *PathBasedSubmissionEndpointAllOf
	isSet bool
}

func (v NullablePathBasedSubmissionEndpointAllOf) Get() *PathBasedSubmissionEndpointAllOf {
	return v.value
}

func (v *NullablePathBasedSubmissionEndpointAllOf) Set(val *PathBasedSubmissionEndpointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePathBasedSubmissionEndpointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePathBasedSubmissionEndpointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathBasedSubmissionEndpointAllOf(val *PathBasedSubmissionEndpointAllOf) *NullablePathBasedSubmissionEndpointAllOf {
	return &NullablePathBasedSubmissionEndpointAllOf{value: val, isSet: true}
}

func (v NullablePathBasedSubmissionEndpointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathBasedSubmissionEndpointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
