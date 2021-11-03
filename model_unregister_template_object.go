/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UnregisterTemplateObject struct for UnregisterTemplateObject
type UnregisterTemplateObject struct {
	Client              *VirtualizationRealmClient `json:"client,omitempty"`
	RemoveSubscriptions bool                       `json:"removeSubscriptions"`
}

// NewUnregisterTemplateObject instantiates a new UnregisterTemplateObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnregisterTemplateObject(removeSubscriptions bool) *UnregisterTemplateObject {
	this := UnregisterTemplateObject{}
	this.RemoveSubscriptions = removeSubscriptions
	return &this
}

// NewUnregisterTemplateObjectWithDefaults instantiates a new UnregisterTemplateObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnregisterTemplateObjectWithDefaults() *UnregisterTemplateObject {
	this := UnregisterTemplateObject{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *UnregisterTemplateObject) GetClient() VirtualizationRealmClient {
	if o == nil || o.Client == nil {
		var ret VirtualizationRealmClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnregisterTemplateObject) GetClientOk() (*VirtualizationRealmClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *UnregisterTemplateObject) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given VirtualizationRealmClient and assigns it to the Client field.
func (o *UnregisterTemplateObject) SetClient(v VirtualizationRealmClient) {
	o.Client = &v
}

// GetRemoveSubscriptions returns the RemoveSubscriptions field value
func (o *UnregisterTemplateObject) GetRemoveSubscriptions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoveSubscriptions
}

// GetRemoveSubscriptionsOk returns a tuple with the RemoveSubscriptions field value
// and a boolean to check if the value has been set.
func (o *UnregisterTemplateObject) GetRemoveSubscriptionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveSubscriptions, true
}

// SetRemoveSubscriptions sets field value
func (o *UnregisterTemplateObject) SetRemoveSubscriptions(v bool) {
	o.RemoveSubscriptions = v
}

func (o UnregisterTemplateObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if true {
		toSerialize["removeSubscriptions"] = o.RemoveSubscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableUnregisterTemplateObject struct {
	value *UnregisterTemplateObject
	isSet bool
}

func (v NullableUnregisterTemplateObject) Get() *UnregisterTemplateObject {
	return v.value
}

func (v *NullableUnregisterTemplateObject) Set(val *UnregisterTemplateObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUnregisterTemplateObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUnregisterTemplateObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnregisterTemplateObject(val *UnregisterTemplateObject) *NullableUnregisterTemplateObject {
	return &NullableUnregisterTemplateObject{value: val, isSet: true}
}

func (v NullableUnregisterTemplateObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnregisterTemplateObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
