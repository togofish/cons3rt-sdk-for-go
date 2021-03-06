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

// InputSFTPSubmissionEndpointForProject struct for InputSFTPSubmissionEndpointForProject
type InputSFTPSubmissionEndpointForProject struct {
	RemotePath             *string `json:"remotePath,omitempty"`
	UseUserDirectoryAsRoot *bool   `json:"useUserDirectoryAsRoot,omitempty"`
	Host                   string
	Subtype                string
}

// NewInputSFTPSubmissionEndpointForProject instantiates a new InputSFTPSubmissionEndpointForProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputSFTPSubmissionEndpointForProject(host string, subtype string) *InputSFTPSubmissionEndpointForProject {
	this := InputSFTPSubmissionEndpointForProject{}
	this.Host = host
	this.Subtype = subtype
	return &this
}

// NewInputSFTPSubmissionEndpointForProjectWithDefaults instantiates a new InputSFTPSubmissionEndpointForProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputSFTPSubmissionEndpointForProjectWithDefaults() *InputSFTPSubmissionEndpointForProject {
	this := InputSFTPSubmissionEndpointForProject{}
	return &this
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise.
func (o *InputSFTPSubmissionEndpointForProject) GetRemotePath() string {
	if o == nil || o.RemotePath == nil {
		var ret string
		return ret
	}
	return *o.RemotePath
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSFTPSubmissionEndpointForProject) GetRemotePathOk() (*string, bool) {
	if o == nil || o.RemotePath == nil {
		return nil, false
	}
	return o.RemotePath, true
}

// HasRemotePath returns a boolean if a field has been set.
func (o *InputSFTPSubmissionEndpointForProject) HasRemotePath() bool {
	if o != nil && o.RemotePath != nil {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given string and assigns it to the RemotePath field.
func (o *InputSFTPSubmissionEndpointForProject) SetRemotePath(v string) {
	o.RemotePath = &v
}

// GetUseUserDirectoryAsRoot returns the UseUserDirectoryAsRoot field value if set, zero value otherwise.
func (o *InputSFTPSubmissionEndpointForProject) GetUseUserDirectoryAsRoot() bool {
	if o == nil || o.UseUserDirectoryAsRoot == nil {
		var ret bool
		return ret
	}
	return *o.UseUserDirectoryAsRoot
}

// GetUseUserDirectoryAsRootOk returns a tuple with the UseUserDirectoryAsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSFTPSubmissionEndpointForProject) GetUseUserDirectoryAsRootOk() (*bool, bool) {
	if o == nil || o.UseUserDirectoryAsRoot == nil {
		return nil, false
	}
	return o.UseUserDirectoryAsRoot, true
}

// HasUseUserDirectoryAsRoot returns a boolean if a field has been set.
func (o *InputSFTPSubmissionEndpointForProject) HasUseUserDirectoryAsRoot() bool {
	if o != nil && o.UseUserDirectoryAsRoot != nil {
		return true
	}

	return false
}

// SetUseUserDirectoryAsRoot gets a reference to the given bool and assigns it to the UseUserDirectoryAsRoot field.
func (o *InputSFTPSubmissionEndpointForProject) SetUseUserDirectoryAsRoot(v bool) {
	o.UseUserDirectoryAsRoot = &v
}

func (o InputSFTPSubmissionEndpointForProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemotePath != nil {
		toSerialize["remotePath"] = o.RemotePath
	}
	if o.UseUserDirectoryAsRoot != nil {
		toSerialize["useUserDirectoryAsRoot"] = o.UseUserDirectoryAsRoot
	}
	return json.Marshal(toSerialize)
}

type NullableInputSFTPSubmissionEndpointForProject struct {
	value *InputSFTPSubmissionEndpointForProject
	isSet bool
}

func (v NullableInputSFTPSubmissionEndpointForProject) Get() *InputSFTPSubmissionEndpointForProject {
	return v.value
}

func (v *NullableInputSFTPSubmissionEndpointForProject) Set(val *InputSFTPSubmissionEndpointForProject) {
	v.value = val
	v.isSet = true
}

func (v NullableInputSFTPSubmissionEndpointForProject) IsSet() bool {
	return v.isSet
}

func (v *NullableInputSFTPSubmissionEndpointForProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputSFTPSubmissionEndpointForProject(val *InputSFTPSubmissionEndpointForProject) *NullableInputSFTPSubmissionEndpointForProject {
	return &NullableInputSFTPSubmissionEndpointForProject{value: val, isSet: true}
}

func (v NullableInputSFTPSubmissionEndpointForProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputSFTPSubmissionEndpointForProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
