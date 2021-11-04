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

// ImageReferenceDTO struct for ImageReferenceDTO
type ImageReferenceDTO struct {
	Publisher string `json:"publisher"`
	Type string `json:"type"`
	ImageVersion string `json:"imageVersion"`
}

// NewImageReferenceDTO instantiates a new ImageReferenceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageReferenceDTO(publisher string, type_ string, imageVersion string) *ImageReferenceDTO {
	this := ImageReferenceDTO{}
	this.Publisher = publisher
	this.Type = type_
	this.ImageVersion = imageVersion
	return &this
}

// NewImageReferenceDTOWithDefaults instantiates a new ImageReferenceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageReferenceDTOWithDefaults() *ImageReferenceDTO {
	this := ImageReferenceDTO{}
	return &this
}

// GetPublisher returns the Publisher field value
func (o *ImageReferenceDTO) GetPublisher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value
// and a boolean to check if the value has been set.
func (o *ImageReferenceDTO) GetPublisherOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Publisher, true
}

// SetPublisher sets field value
func (o *ImageReferenceDTO) SetPublisher(v string) {
	o.Publisher = v
}

// GetType returns the Type field value
func (o *ImageReferenceDTO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImageReferenceDTO) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImageReferenceDTO) SetType(v string) {
	o.Type = v
}

// GetImageVersion returns the ImageVersion field value
func (o *ImageReferenceDTO) GetImageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageVersion
}

// GetImageVersionOk returns a tuple with the ImageVersion field value
// and a boolean to check if the value has been set.
func (o *ImageReferenceDTO) GetImageVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageVersion, true
}

// SetImageVersion sets field value
func (o *ImageReferenceDTO) SetImageVersion(v string) {
	o.ImageVersion = v
}

func (o ImageReferenceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisher"] = o.Publisher
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["imageVersion"] = o.ImageVersion
	}
	return json.Marshal(toSerialize)
}

type NullableImageReferenceDTO struct {
	value *ImageReferenceDTO
	isSet bool
}

func (v NullableImageReferenceDTO) Get() *ImageReferenceDTO {
	return v.value
}

func (v *NullableImageReferenceDTO) Set(val *ImageReferenceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableImageReferenceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableImageReferenceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageReferenceDTO(val *ImageReferenceDTO) *NullableImageReferenceDTO {
	return &NullableImageReferenceDTO{value: val, isSet: true}
}

func (v NullableImageReferenceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageReferenceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


