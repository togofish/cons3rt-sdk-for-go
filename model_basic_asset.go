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

// BasicAsset struct for BasicAsset
type BasicAsset struct {
	Id            *string         `json:"id,omitempty"`
	Name          *string         `json:"name,omitempty"`
	Description   *string         `json:"description,omitempty"`
	Offline       *bool           `json:"offline,omitempty"`
	State         *string         `json:"state,omitempty"`
	Visibility    *string         `json:"visibility,omitempty"`
	Creator       *MinimalUser    `json:"creator,omitempty"`
	OwningProject *MinimalProject `json:"owningProject,omitempty"`
	Subtype       string          `json:"subtype"`
}

// NewBasicAsset instantiates a new BasicAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAsset(subtype string) *BasicAsset {
	this := BasicAsset{}
	this.Subtype = subtype
	return &this
}

// NewBasicAssetWithDefaults instantiates a new BasicAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAssetWithDefaults() *BasicAsset {
	this := BasicAsset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BasicAsset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BasicAsset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BasicAsset) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BasicAsset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BasicAsset) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BasicAsset) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BasicAsset) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BasicAsset) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BasicAsset) SetDescription(v string) {
	o.Description = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *BasicAsset) GetOffline() bool {
	if o == nil || o.Offline == nil {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetOfflineOk() (*bool, bool) {
	if o == nil || o.Offline == nil {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *BasicAsset) HasOffline() bool {
	if o != nil && o.Offline != nil {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *BasicAsset) SetOffline(v bool) {
	o.Offline = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BasicAsset) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BasicAsset) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BasicAsset) SetState(v string) {
	o.State = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BasicAsset) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BasicAsset) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BasicAsset) SetVisibility(v string) {
	o.Visibility = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *BasicAsset) GetCreator() MinimalUser {
	if o == nil || o.Creator == nil {
		var ret MinimalUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetCreatorOk() (*MinimalUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *BasicAsset) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given MinimalUser and assigns it to the Creator field.
func (o *BasicAsset) SetCreator(v MinimalUser) {
	o.Creator = &v
}

// GetOwningProject returns the OwningProject field value if set, zero value otherwise.
func (o *BasicAsset) GetOwningProject() MinimalProject {
	if o == nil || o.OwningProject == nil {
		var ret MinimalProject
		return ret
	}
	return *o.OwningProject
}

// GetOwningProjectOk returns a tuple with the OwningProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetOwningProjectOk() (*MinimalProject, bool) {
	if o == nil || o.OwningProject == nil {
		return nil, false
	}
	return o.OwningProject, true
}

// HasOwningProject returns a boolean if a field has been set.
func (o *BasicAsset) HasOwningProject() bool {
	if o != nil && o.OwningProject != nil {
		return true
	}

	return false
}

// SetOwningProject gets a reference to the given MinimalProject and assigns it to the OwningProject field.
func (o *BasicAsset) SetOwningProject(v MinimalProject) {
	o.OwningProject = &v
}

// GetSubtype returns the Subtype field value
func (o *BasicAsset) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *BasicAsset) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *BasicAsset) SetSubtype(v string) {
	o.Subtype = v
}

func (o BasicAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.OwningProject != nil {
		toSerialize["owningProject"] = o.OwningProject
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableBasicAsset struct {
	value *BasicAsset
	isSet bool
}

func (v NullableBasicAsset) Get() *BasicAsset {
	return v.value
}

func (v *NullableBasicAsset) Set(val *BasicAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicAsset(val *BasicAsset) *NullableBasicAsset {
	return &NullableBasicAsset{value: val, isSet: true}
}

func (v NullableBasicAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
