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

// MetadataDocsLicense struct for MetadataDocsLicense
type MetadataDocsLicense struct {
	License       *string `json:"license,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
}

// NewMetadataDocsLicense instantiates a new MetadataDocsLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataDocsLicense() *MetadataDocsLicense {
	this := MetadataDocsLicense{}
	return &this
}

// NewMetadataDocsLicenseWithDefaults instantiates a new MetadataDocsLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataDocsLicenseWithDefaults() *MetadataDocsLicense {
	this := MetadataDocsLicense{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *MetadataDocsLicense) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDocsLicense) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *MetadataDocsLicense) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *MetadataDocsLicense) SetLicense(v string) {
	o.License = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *MetadataDocsLicense) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataDocsLicense) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *MetadataDocsLicense) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *MetadataDocsLicense) SetDocumentation(v string) {
	o.Documentation = &v
}

func (o MetadataDocsLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataDocsLicense struct {
	value *MetadataDocsLicense
	isSet bool
}

func (v NullableMetadataDocsLicense) Get() *MetadataDocsLicense {
	return v.value
}

func (v *NullableMetadataDocsLicense) Set(val *MetadataDocsLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataDocsLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataDocsLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataDocsLicense(val *MetadataDocsLicense) *NullableMetadataDocsLicense {
	return &NullableMetadataDocsLicense{value: val, isSet: true}
}

func (v NullableMetadataDocsLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataDocsLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
