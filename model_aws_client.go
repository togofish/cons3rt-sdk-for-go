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

// AwsClient struct for AwsClient
type AwsClient struct {
	Region   *string `json:"region,omitempty"`
	Username string
	Password string
	Subtype  string
}

// NewAwsClient instantiates a new AwsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsClient(username string, password string, subtype string) *AwsClient {
	this := AwsClient{}
	this.Username = username
	this.Password = password
	this.Subtype = subtype
	return &this
}

// NewAwsClientWithDefaults instantiates a new AwsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsClientWithDefaults() *AwsClient {
	this := AwsClient{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AwsClient) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsClient) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AwsClient) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AwsClient) SetRegion(v string) {
	o.Region = &v
}

func (o AwsClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableAwsClient struct {
	value *AwsClient
	isSet bool
}

func (v NullableAwsClient) Get() *AwsClient {
	return v.value
}

func (v *NullableAwsClient) Set(val *AwsClient) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsClient) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsClient(val *AwsClient) *NullableAwsClient {
	return &NullableAwsClient{value: val, isSet: true}
}

func (v NullableAwsClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
