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

// PocInfo struct for PocInfo
type PocInfo struct {
	Email *string `json:"email,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// NewPocInfo instantiates a new PocInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPocInfo() *PocInfo {
	this := PocInfo{}
	return &this
}

// NewPocInfoWithDefaults instantiates a new PocInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPocInfoWithDefaults() *PocInfo {
	this := PocInfo{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PocInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PocInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PocInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PocInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *PocInfo) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PocInfo) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *PocInfo) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *PocInfo) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *PocInfo) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PocInfo) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *PocInfo) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *PocInfo) SetLastname(v string) {
	o.Lastname = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PocInfo) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PocInfo) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PocInfo) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *PocInfo) SetOrganization(v string) {
	o.Organization = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PocInfo) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PocInfo) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PocInfo) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *PocInfo) SetPhone(v string) {
	o.Phone = &v
}

func (o PocInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullablePocInfo struct {
	value *PocInfo
	isSet bool
}

func (v NullablePocInfo) Get() *PocInfo {
	return v.value
}

func (v *NullablePocInfo) Set(val *PocInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePocInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePocInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePocInfo(val *PocInfo) *NullablePocInfo {
	return &NullablePocInfo{value: val, isSet: true}
}

func (v NullablePocInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePocInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


