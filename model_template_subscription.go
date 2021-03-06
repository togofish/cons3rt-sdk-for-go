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

// TemplateSubscription struct for TemplateSubscription
type TemplateSubscription struct {
	Id                             *int32                `json:"id,omitempty"`
	MaxNumCpus                     *int32                `json:"maxNumCpus,omitempty"`
	MaxRamInMegabytes              *int32                `json:"maxRamInMegabytes,omitempty"`
	State                          *string               `json:"state,omitempty"`
	SubscribingVirtualizationRealm *VirtualizationRealm  `json:"subscribingVirtualizationRealm,omitempty"`
	TemplateRegistration           *TemplateRegistration `json:"templateRegistration,omitempty"`
	TemplateUuid                   *string               `json:"templateUuid,omitempty"`
}

// NewTemplateSubscription instantiates a new TemplateSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSubscription() *TemplateSubscription {
	this := TemplateSubscription{}
	return &this
}

// NewTemplateSubscriptionWithDefaults instantiates a new TemplateSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSubscriptionWithDefaults() *TemplateSubscription {
	this := TemplateSubscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateSubscription) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateSubscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TemplateSubscription) SetId(v int32) {
	o.Id = &v
}

// GetMaxNumCpus returns the MaxNumCpus field value if set, zero value otherwise.
func (o *TemplateSubscription) GetMaxNumCpus() int32 {
	if o == nil || o.MaxNumCpus == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumCpus
}

// GetMaxNumCpusOk returns a tuple with the MaxNumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetMaxNumCpusOk() (*int32, bool) {
	if o == nil || o.MaxNumCpus == nil {
		return nil, false
	}
	return o.MaxNumCpus, true
}

// HasMaxNumCpus returns a boolean if a field has been set.
func (o *TemplateSubscription) HasMaxNumCpus() bool {
	if o != nil && o.MaxNumCpus != nil {
		return true
	}

	return false
}

// SetMaxNumCpus gets a reference to the given int32 and assigns it to the MaxNumCpus field.
func (o *TemplateSubscription) SetMaxNumCpus(v int32) {
	o.MaxNumCpus = &v
}

// GetMaxRamInMegabytes returns the MaxRamInMegabytes field value if set, zero value otherwise.
func (o *TemplateSubscription) GetMaxRamInMegabytes() int32 {
	if o == nil || o.MaxRamInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRamInMegabytes
}

// GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetMaxRamInMegabytesOk() (*int32, bool) {
	if o == nil || o.MaxRamInMegabytes == nil {
		return nil, false
	}
	return o.MaxRamInMegabytes, true
}

// HasMaxRamInMegabytes returns a boolean if a field has been set.
func (o *TemplateSubscription) HasMaxRamInMegabytes() bool {
	if o != nil && o.MaxRamInMegabytes != nil {
		return true
	}

	return false
}

// SetMaxRamInMegabytes gets a reference to the given int32 and assigns it to the MaxRamInMegabytes field.
func (o *TemplateSubscription) SetMaxRamInMegabytes(v int32) {
	o.MaxRamInMegabytes = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TemplateSubscription) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TemplateSubscription) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TemplateSubscription) SetState(v string) {
	o.State = &v
}

// GetSubscribingVirtualizationRealm returns the SubscribingVirtualizationRealm field value if set, zero value otherwise.
func (o *TemplateSubscription) GetSubscribingVirtualizationRealm() VirtualizationRealm {
	if o == nil || o.SubscribingVirtualizationRealm == nil {
		var ret VirtualizationRealm
		return ret
	}
	return *o.SubscribingVirtualizationRealm
}

// GetSubscribingVirtualizationRealmOk returns a tuple with the SubscribingVirtualizationRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetSubscribingVirtualizationRealmOk() (*VirtualizationRealm, bool) {
	if o == nil || o.SubscribingVirtualizationRealm == nil {
		return nil, false
	}
	return o.SubscribingVirtualizationRealm, true
}

// HasSubscribingVirtualizationRealm returns a boolean if a field has been set.
func (o *TemplateSubscription) HasSubscribingVirtualizationRealm() bool {
	if o != nil && o.SubscribingVirtualizationRealm != nil {
		return true
	}

	return false
}

// SetSubscribingVirtualizationRealm gets a reference to the given VirtualizationRealm and assigns it to the SubscribingVirtualizationRealm field.
func (o *TemplateSubscription) SetSubscribingVirtualizationRealm(v VirtualizationRealm) {
	o.SubscribingVirtualizationRealm = &v
}

// GetTemplateRegistration returns the TemplateRegistration field value if set, zero value otherwise.
func (o *TemplateSubscription) GetTemplateRegistration() TemplateRegistration {
	if o == nil || o.TemplateRegistration == nil {
		var ret TemplateRegistration
		return ret
	}
	return *o.TemplateRegistration
}

// GetTemplateRegistrationOk returns a tuple with the TemplateRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetTemplateRegistrationOk() (*TemplateRegistration, bool) {
	if o == nil || o.TemplateRegistration == nil {
		return nil, false
	}
	return o.TemplateRegistration, true
}

// HasTemplateRegistration returns a boolean if a field has been set.
func (o *TemplateSubscription) HasTemplateRegistration() bool {
	if o != nil && o.TemplateRegistration != nil {
		return true
	}

	return false
}

// SetTemplateRegistration gets a reference to the given TemplateRegistration and assigns it to the TemplateRegistration field.
func (o *TemplateSubscription) SetTemplateRegistration(v TemplateRegistration) {
	o.TemplateRegistration = &v
}

// GetTemplateUuid returns the TemplateUuid field value if set, zero value otherwise.
func (o *TemplateSubscription) GetTemplateUuid() string {
	if o == nil || o.TemplateUuid == nil {
		var ret string
		return ret
	}
	return *o.TemplateUuid
}

// GetTemplateUuidOk returns a tuple with the TemplateUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSubscription) GetTemplateUuidOk() (*string, bool) {
	if o == nil || o.TemplateUuid == nil {
		return nil, false
	}
	return o.TemplateUuid, true
}

// HasTemplateUuid returns a boolean if a field has been set.
func (o *TemplateSubscription) HasTemplateUuid() bool {
	if o != nil && o.TemplateUuid != nil {
		return true
	}

	return false
}

// SetTemplateUuid gets a reference to the given string and assigns it to the TemplateUuid field.
func (o *TemplateSubscription) SetTemplateUuid(v string) {
	o.TemplateUuid = &v
}

func (o TemplateSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MaxNumCpus != nil {
		toSerialize["maxNumCpus"] = o.MaxNumCpus
	}
	if o.MaxRamInMegabytes != nil {
		toSerialize["maxRamInMegabytes"] = o.MaxRamInMegabytes
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SubscribingVirtualizationRealm != nil {
		toSerialize["subscribingVirtualizationRealm"] = o.SubscribingVirtualizationRealm
	}
	if o.TemplateRegistration != nil {
		toSerialize["templateRegistration"] = o.TemplateRegistration
	}
	if o.TemplateUuid != nil {
		toSerialize["templateUuid"] = o.TemplateUuid
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSubscription struct {
	value *TemplateSubscription
	isSet bool
}

func (v NullableTemplateSubscription) Get() *TemplateSubscription {
	return v.value
}

func (v *NullableTemplateSubscription) Set(val *TemplateSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSubscription(val *TemplateSubscription) *NullableTemplateSubscription {
	return &NullableTemplateSubscription{value: val, isSet: true}
}

func (v NullableTemplateSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
