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

// UserProject struct for UserProject
type UserProject struct {
	Id *int32 `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Email *string `json:"email,omitempty"`
	MembershipState *string `json:"membershipState,omitempty"`
	Roles *[]string `json:"roles,omitempty"`
}

// NewUserProject instantiates a new UserProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProject() *UserProject {
	this := UserProject{}
	return &this
}

// NewUserProjectWithDefaults instantiates a new UserProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProjectWithDefaults() *UserProject {
	this := UserProject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserProject) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProject) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserProject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserProject) SetId(v int32) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserProject) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProject) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserProject) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserProject) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserProject) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProject) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserProject) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserProject) SetEmail(v string) {
	o.Email = &v
}

// GetMembershipState returns the MembershipState field value if set, zero value otherwise.
func (o *UserProject) GetMembershipState() string {
	if o == nil || o.MembershipState == nil {
		var ret string
		return ret
	}
	return *o.MembershipState
}

// GetMembershipStateOk returns a tuple with the MembershipState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProject) GetMembershipStateOk() (*string, bool) {
	if o == nil || o.MembershipState == nil {
		return nil, false
	}
	return o.MembershipState, true
}

// HasMembershipState returns a boolean if a field has been set.
func (o *UserProject) HasMembershipState() bool {
	if o != nil && o.MembershipState != nil {
		return true
	}

	return false
}

// SetMembershipState gets a reference to the given string and assigns it to the MembershipState field.
func (o *UserProject) SetMembershipState(v string) {
	o.MembershipState = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserProject) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProject) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserProject) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserProject) SetRoles(v []string) {
	o.Roles = &v
}

func (o UserProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.MembershipState != nil {
		toSerialize["membershipState"] = o.MembershipState
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableUserProject struct {
	value *UserProject
	isSet bool
}

func (v NullableUserProject) Get() *UserProject {
	return v.value
}

func (v *NullableUserProject) Set(val *UserProject) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProject) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProject(val *UserProject) *NullableUserProject {
	return &NullableUserProject{value: val, isSet: true}
}

func (v NullableUserProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


