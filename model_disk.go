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

// Disk struct for Disk
type Disk struct {
	AdditionalDisk      *bool  `json:"additionalDisk,omitempty"`
	BootDisk            *bool  `json:"bootDisk,omitempty"`
	CapacityInMegabytes int32  `json:"capacityInMegabytes"`
	CreateOrder         *int32 `json:"createOrder,omitempty"`
	IsAdditionalDisk    *bool  `json:"isAdditionalDisk,omitempty"`
	IsBootDisk          *bool  `json:"isBootDisk,omitempty"`
	Id                  *int32 `json:"id,omitempty"`
}

// NewDisk instantiates a new Disk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisk(capacityInMegabytes int32) *Disk {
	this := Disk{}
	this.CapacityInMegabytes = capacityInMegabytes
	return &this
}

// NewDiskWithDefaults instantiates a new Disk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskWithDefaults() *Disk {
	this := Disk{}
	return &this
}

// GetAdditionalDisk returns the AdditionalDisk field value if set, zero value otherwise.
func (o *Disk) GetAdditionalDisk() bool {
	if o == nil || o.AdditionalDisk == nil {
		var ret bool
		return ret
	}
	return *o.AdditionalDisk
}

// GetAdditionalDiskOk returns a tuple with the AdditionalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetAdditionalDiskOk() (*bool, bool) {
	if o == nil || o.AdditionalDisk == nil {
		return nil, false
	}
	return o.AdditionalDisk, true
}

// HasAdditionalDisk returns a boolean if a field has been set.
func (o *Disk) HasAdditionalDisk() bool {
	if o != nil && o.AdditionalDisk != nil {
		return true
	}

	return false
}

// SetAdditionalDisk gets a reference to the given bool and assigns it to the AdditionalDisk field.
func (o *Disk) SetAdditionalDisk(v bool) {
	o.AdditionalDisk = &v
}

// GetBootDisk returns the BootDisk field value if set, zero value otherwise.
func (o *Disk) GetBootDisk() bool {
	if o == nil || o.BootDisk == nil {
		var ret bool
		return ret
	}
	return *o.BootDisk
}

// GetBootDiskOk returns a tuple with the BootDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetBootDiskOk() (*bool, bool) {
	if o == nil || o.BootDisk == nil {
		return nil, false
	}
	return o.BootDisk, true
}

// HasBootDisk returns a boolean if a field has been set.
func (o *Disk) HasBootDisk() bool {
	if o != nil && o.BootDisk != nil {
		return true
	}

	return false
}

// SetBootDisk gets a reference to the given bool and assigns it to the BootDisk field.
func (o *Disk) SetBootDisk(v bool) {
	o.BootDisk = &v
}

// GetCapacityInMegabytes returns the CapacityInMegabytes field value
func (o *Disk) GetCapacityInMegabytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CapacityInMegabytes
}

// GetCapacityInMegabytesOk returns a tuple with the CapacityInMegabytes field value
// and a boolean to check if the value has been set.
func (o *Disk) GetCapacityInMegabytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityInMegabytes, true
}

// SetCapacityInMegabytes sets field value
func (o *Disk) SetCapacityInMegabytes(v int32) {
	o.CapacityInMegabytes = v
}

// GetCreateOrder returns the CreateOrder field value if set, zero value otherwise.
func (o *Disk) GetCreateOrder() int32 {
	if o == nil || o.CreateOrder == nil {
		var ret int32
		return ret
	}
	return *o.CreateOrder
}

// GetCreateOrderOk returns a tuple with the CreateOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetCreateOrderOk() (*int32, bool) {
	if o == nil || o.CreateOrder == nil {
		return nil, false
	}
	return o.CreateOrder, true
}

// HasCreateOrder returns a boolean if a field has been set.
func (o *Disk) HasCreateOrder() bool {
	if o != nil && o.CreateOrder != nil {
		return true
	}

	return false
}

// SetCreateOrder gets a reference to the given int32 and assigns it to the CreateOrder field.
func (o *Disk) SetCreateOrder(v int32) {
	o.CreateOrder = &v
}

// GetIsAdditionalDisk returns the IsAdditionalDisk field value if set, zero value otherwise.
func (o *Disk) GetIsAdditionalDisk() bool {
	if o == nil || o.IsAdditionalDisk == nil {
		var ret bool
		return ret
	}
	return *o.IsAdditionalDisk
}

// GetIsAdditionalDiskOk returns a tuple with the IsAdditionalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetIsAdditionalDiskOk() (*bool, bool) {
	if o == nil || o.IsAdditionalDisk == nil {
		return nil, false
	}
	return o.IsAdditionalDisk, true
}

// HasIsAdditionalDisk returns a boolean if a field has been set.
func (o *Disk) HasIsAdditionalDisk() bool {
	if o != nil && o.IsAdditionalDisk != nil {
		return true
	}

	return false
}

// SetIsAdditionalDisk gets a reference to the given bool and assigns it to the IsAdditionalDisk field.
func (o *Disk) SetIsAdditionalDisk(v bool) {
	o.IsAdditionalDisk = &v
}

// GetIsBootDisk returns the IsBootDisk field value if set, zero value otherwise.
func (o *Disk) GetIsBootDisk() bool {
	if o == nil || o.IsBootDisk == nil {
		var ret bool
		return ret
	}
	return *o.IsBootDisk
}

// GetIsBootDiskOk returns a tuple with the IsBootDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetIsBootDiskOk() (*bool, bool) {
	if o == nil || o.IsBootDisk == nil {
		return nil, false
	}
	return o.IsBootDisk, true
}

// HasIsBootDisk returns a boolean if a field has been set.
func (o *Disk) HasIsBootDisk() bool {
	if o != nil && o.IsBootDisk != nil {
		return true
	}

	return false
}

// SetIsBootDisk gets a reference to the given bool and assigns it to the IsBootDisk field.
func (o *Disk) SetIsBootDisk(v bool) {
	o.IsBootDisk = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Disk) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Disk) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Disk) SetId(v int32) {
	o.Id = &v
}

func (o Disk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalDisk != nil {
		toSerialize["additionalDisk"] = o.AdditionalDisk
	}
	if o.BootDisk != nil {
		toSerialize["bootDisk"] = o.BootDisk
	}
	if true {
		toSerialize["capacityInMegabytes"] = o.CapacityInMegabytes
	}
	if o.CreateOrder != nil {
		toSerialize["createOrder"] = o.CreateOrder
	}
	if o.IsAdditionalDisk != nil {
		toSerialize["isAdditionalDisk"] = o.IsAdditionalDisk
	}
	if o.IsBootDisk != nil {
		toSerialize["isBootDisk"] = o.IsBootDisk
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDisk struct {
	value *Disk
	isSet bool
}

func (v NullableDisk) Get() *Disk {
	return v.value
}

func (v *NullableDisk) Set(val *Disk) {
	v.value = val
	v.isSet = true
}

func (v NullableDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisk(val *Disk) *NullableDisk {
	return &NullableDisk{value: val, isSet: true}
}

func (v NullableDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
