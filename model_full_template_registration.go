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

// FullTemplateRegistration struct for FullTemplateRegistration
type FullTemplateRegistration struct {
	Id *int32 `json:"id,omitempty"`
	RegisteringVirtualizationRealm *MinimalVirtualizationRealm `json:"registeringVirtualizationRealm,omitempty"`
	State *string `json:"state,omitempty"`
	Subscriptions *[]MinimalTemplateSubscription `json:"subscriptions,omitempty"`
	TemplateData *FullCons3rtTemplateData `json:"templateData,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty"`
	VirtRealmsSharedTo *[]MinimalVirtualizationRealm `json:"virtRealmsSharedTo,omitempty"`
}

// NewFullTemplateRegistration instantiates a new FullTemplateRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullTemplateRegistration() *FullTemplateRegistration {
	this := FullTemplateRegistration{}
	return &this
}

// NewFullTemplateRegistrationWithDefaults instantiates a new FullTemplateRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullTemplateRegistrationWithDefaults() *FullTemplateRegistration {
	this := FullTemplateRegistration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FullTemplateRegistration) SetId(v int32) {
	o.Id = &v
}

// GetRegisteringVirtualizationRealm returns the RegisteringVirtualizationRealm field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetRegisteringVirtualizationRealm() MinimalVirtualizationRealm {
	if o == nil || o.RegisteringVirtualizationRealm == nil {
		var ret MinimalVirtualizationRealm
		return ret
	}
	return *o.RegisteringVirtualizationRealm
}

// GetRegisteringVirtualizationRealmOk returns a tuple with the RegisteringVirtualizationRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetRegisteringVirtualizationRealmOk() (*MinimalVirtualizationRealm, bool) {
	if o == nil || o.RegisteringVirtualizationRealm == nil {
		return nil, false
	}
	return o.RegisteringVirtualizationRealm, true
}

// HasRegisteringVirtualizationRealm returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasRegisteringVirtualizationRealm() bool {
	if o != nil && o.RegisteringVirtualizationRealm != nil {
		return true
	}

	return false
}

// SetRegisteringVirtualizationRealm gets a reference to the given MinimalVirtualizationRealm and assigns it to the RegisteringVirtualizationRealm field.
func (o *FullTemplateRegistration) SetRegisteringVirtualizationRealm(v MinimalVirtualizationRealm) {
	o.RegisteringVirtualizationRealm = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FullTemplateRegistration) SetState(v string) {
	o.State = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetSubscriptions() []MinimalTemplateSubscription {
	if o == nil || o.Subscriptions == nil {
		var ret []MinimalTemplateSubscription
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetSubscriptionsOk() (*[]MinimalTemplateSubscription, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []MinimalTemplateSubscription and assigns it to the Subscriptions field.
func (o *FullTemplateRegistration) SetSubscriptions(v []MinimalTemplateSubscription) {
	o.Subscriptions = &v
}

// GetTemplateData returns the TemplateData field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetTemplateData() FullCons3rtTemplateData {
	if o == nil || o.TemplateData == nil {
		var ret FullCons3rtTemplateData
		return ret
	}
	return *o.TemplateData
}

// GetTemplateDataOk returns a tuple with the TemplateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetTemplateDataOk() (*FullCons3rtTemplateData, bool) {
	if o == nil || o.TemplateData == nil {
		return nil, false
	}
	return o.TemplateData, true
}

// HasTemplateData returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasTemplateData() bool {
	if o != nil && o.TemplateData != nil {
		return true
	}

	return false
}

// SetTemplateData gets a reference to the given FullCons3rtTemplateData and assigns it to the TemplateData field.
func (o *FullTemplateRegistration) SetTemplateData(v FullCons3rtTemplateData) {
	o.TemplateData = &v
}

// GetTemplateUuid returns the TemplateUuid field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetTemplateUuid() string {
	if o == nil || o.TemplateUuid == nil {
		var ret string
		return ret
	}
	return *o.TemplateUuid
}

// GetTemplateUuidOk returns a tuple with the TemplateUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetTemplateUuidOk() (*string, bool) {
	if o == nil || o.TemplateUuid == nil {
		return nil, false
	}
	return o.TemplateUuid, true
}

// HasTemplateUuid returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasTemplateUuid() bool {
	if o != nil && o.TemplateUuid != nil {
		return true
	}

	return false
}

// SetTemplateUuid gets a reference to the given string and assigns it to the TemplateUuid field.
func (o *FullTemplateRegistration) SetTemplateUuid(v string) {
	o.TemplateUuid = &v
}

// GetVirtRealmsSharedTo returns the VirtRealmsSharedTo field value if set, zero value otherwise.
func (o *FullTemplateRegistration) GetVirtRealmsSharedTo() []MinimalVirtualizationRealm {
	if o == nil || o.VirtRealmsSharedTo == nil {
		var ret []MinimalVirtualizationRealm
		return ret
	}
	return *o.VirtRealmsSharedTo
}

// GetVirtRealmsSharedToOk returns a tuple with the VirtRealmsSharedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullTemplateRegistration) GetVirtRealmsSharedToOk() (*[]MinimalVirtualizationRealm, bool) {
	if o == nil || o.VirtRealmsSharedTo == nil {
		return nil, false
	}
	return o.VirtRealmsSharedTo, true
}

// HasVirtRealmsSharedTo returns a boolean if a field has been set.
func (o *FullTemplateRegistration) HasVirtRealmsSharedTo() bool {
	if o != nil && o.VirtRealmsSharedTo != nil {
		return true
	}

	return false
}

// SetVirtRealmsSharedTo gets a reference to the given []MinimalVirtualizationRealm and assigns it to the VirtRealmsSharedTo field.
func (o *FullTemplateRegistration) SetVirtRealmsSharedTo(v []MinimalVirtualizationRealm) {
	o.VirtRealmsSharedTo = &v
}

func (o FullTemplateRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RegisteringVirtualizationRealm != nil {
		toSerialize["registeringVirtualizationRealm"] = o.RegisteringVirtualizationRealm
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.TemplateData != nil {
		toSerialize["templateData"] = o.TemplateData
	}
	if o.TemplateUuid != nil {
		toSerialize["templateUuid"] = o.TemplateUuid
	}
	if o.VirtRealmsSharedTo != nil {
		toSerialize["virtRealmsSharedTo"] = o.VirtRealmsSharedTo
	}
	return json.Marshal(toSerialize)
}

type NullableFullTemplateRegistration struct {
	value *FullTemplateRegistration
	isSet bool
}

func (v NullableFullTemplateRegistration) Get() *FullTemplateRegistration {
	return v.value
}

func (v *NullableFullTemplateRegistration) Set(val *FullTemplateRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableFullTemplateRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableFullTemplateRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullTemplateRegistration(val *FullTemplateRegistration) *NullableFullTemplateRegistration {
	return &NullableFullTemplateRegistration{value: val, isSet: true}
}

func (v NullableFullTemplateRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullTemplateRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


