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

// PowerSchedule struct for PowerSchedule
type PowerSchedule struct {
	Mode                    string `json:"mode"`
	WeekdayEndTimeHour      *int32 `json:"weekdayEndTimeHour,omitempty"`
	WeekdayEndTimeMinutes   *int32 `json:"weekdayEndTimeMinutes,omitempty"`
	WeekdayStartTimeHour    *int32 `json:"weekdayStartTimeHour,omitempty"`
	WeekdayStartTimeMinutes *int32 `json:"weekdayStartTimeMinutes,omitempty"`
	WeekendEndTimeHour      *int32 `json:"weekendEndTimeHour,omitempty"`
	WeekendEndTimeMinutes   *int32 `json:"weekendEndTimeMinutes,omitempty"`
	WeekendStartTimeHour    *int32 `json:"weekendStartTimeHour,omitempty"`
	WeekendStartTimeMinutes *int32 `json:"weekendStartTimeMinutes,omitempty"`
}

// NewPowerSchedule instantiates a new PowerSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerSchedule(mode string) *PowerSchedule {
	this := PowerSchedule{}
	this.Mode = mode
	return &this
}

// NewPowerScheduleWithDefaults instantiates a new PowerSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerScheduleWithDefaults() *PowerSchedule {
	this := PowerSchedule{}
	return &this
}

// GetMode returns the Mode field value
func (o *PowerSchedule) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PowerSchedule) SetMode(v string) {
	o.Mode = v
}

// GetWeekdayEndTimeHour returns the WeekdayEndTimeHour field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekdayEndTimeHour() int32 {
	if o == nil || o.WeekdayEndTimeHour == nil {
		var ret int32
		return ret
	}
	return *o.WeekdayEndTimeHour
}

// GetWeekdayEndTimeHourOk returns a tuple with the WeekdayEndTimeHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekdayEndTimeHourOk() (*int32, bool) {
	if o == nil || o.WeekdayEndTimeHour == nil {
		return nil, false
	}
	return o.WeekdayEndTimeHour, true
}

// HasWeekdayEndTimeHour returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekdayEndTimeHour() bool {
	if o != nil && o.WeekdayEndTimeHour != nil {
		return true
	}

	return false
}

// SetWeekdayEndTimeHour gets a reference to the given int32 and assigns it to the WeekdayEndTimeHour field.
func (o *PowerSchedule) SetWeekdayEndTimeHour(v int32) {
	o.WeekdayEndTimeHour = &v
}

// GetWeekdayEndTimeMinutes returns the WeekdayEndTimeMinutes field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekdayEndTimeMinutes() int32 {
	if o == nil || o.WeekdayEndTimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.WeekdayEndTimeMinutes
}

// GetWeekdayEndTimeMinutesOk returns a tuple with the WeekdayEndTimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekdayEndTimeMinutesOk() (*int32, bool) {
	if o == nil || o.WeekdayEndTimeMinutes == nil {
		return nil, false
	}
	return o.WeekdayEndTimeMinutes, true
}

// HasWeekdayEndTimeMinutes returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekdayEndTimeMinutes() bool {
	if o != nil && o.WeekdayEndTimeMinutes != nil {
		return true
	}

	return false
}

// SetWeekdayEndTimeMinutes gets a reference to the given int32 and assigns it to the WeekdayEndTimeMinutes field.
func (o *PowerSchedule) SetWeekdayEndTimeMinutes(v int32) {
	o.WeekdayEndTimeMinutes = &v
}

// GetWeekdayStartTimeHour returns the WeekdayStartTimeHour field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekdayStartTimeHour() int32 {
	if o == nil || o.WeekdayStartTimeHour == nil {
		var ret int32
		return ret
	}
	return *o.WeekdayStartTimeHour
}

// GetWeekdayStartTimeHourOk returns a tuple with the WeekdayStartTimeHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekdayStartTimeHourOk() (*int32, bool) {
	if o == nil || o.WeekdayStartTimeHour == nil {
		return nil, false
	}
	return o.WeekdayStartTimeHour, true
}

// HasWeekdayStartTimeHour returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekdayStartTimeHour() bool {
	if o != nil && o.WeekdayStartTimeHour != nil {
		return true
	}

	return false
}

// SetWeekdayStartTimeHour gets a reference to the given int32 and assigns it to the WeekdayStartTimeHour field.
func (o *PowerSchedule) SetWeekdayStartTimeHour(v int32) {
	o.WeekdayStartTimeHour = &v
}

// GetWeekdayStartTimeMinutes returns the WeekdayStartTimeMinutes field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekdayStartTimeMinutes() int32 {
	if o == nil || o.WeekdayStartTimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.WeekdayStartTimeMinutes
}

// GetWeekdayStartTimeMinutesOk returns a tuple with the WeekdayStartTimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekdayStartTimeMinutesOk() (*int32, bool) {
	if o == nil || o.WeekdayStartTimeMinutes == nil {
		return nil, false
	}
	return o.WeekdayStartTimeMinutes, true
}

// HasWeekdayStartTimeMinutes returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekdayStartTimeMinutes() bool {
	if o != nil && o.WeekdayStartTimeMinutes != nil {
		return true
	}

	return false
}

// SetWeekdayStartTimeMinutes gets a reference to the given int32 and assigns it to the WeekdayStartTimeMinutes field.
func (o *PowerSchedule) SetWeekdayStartTimeMinutes(v int32) {
	o.WeekdayStartTimeMinutes = &v
}

// GetWeekendEndTimeHour returns the WeekendEndTimeHour field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekendEndTimeHour() int32 {
	if o == nil || o.WeekendEndTimeHour == nil {
		var ret int32
		return ret
	}
	return *o.WeekendEndTimeHour
}

// GetWeekendEndTimeHourOk returns a tuple with the WeekendEndTimeHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekendEndTimeHourOk() (*int32, bool) {
	if o == nil || o.WeekendEndTimeHour == nil {
		return nil, false
	}
	return o.WeekendEndTimeHour, true
}

// HasWeekendEndTimeHour returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekendEndTimeHour() bool {
	if o != nil && o.WeekendEndTimeHour != nil {
		return true
	}

	return false
}

// SetWeekendEndTimeHour gets a reference to the given int32 and assigns it to the WeekendEndTimeHour field.
func (o *PowerSchedule) SetWeekendEndTimeHour(v int32) {
	o.WeekendEndTimeHour = &v
}

// GetWeekendEndTimeMinutes returns the WeekendEndTimeMinutes field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekendEndTimeMinutes() int32 {
	if o == nil || o.WeekendEndTimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.WeekendEndTimeMinutes
}

// GetWeekendEndTimeMinutesOk returns a tuple with the WeekendEndTimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekendEndTimeMinutesOk() (*int32, bool) {
	if o == nil || o.WeekendEndTimeMinutes == nil {
		return nil, false
	}
	return o.WeekendEndTimeMinutes, true
}

// HasWeekendEndTimeMinutes returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekendEndTimeMinutes() bool {
	if o != nil && o.WeekendEndTimeMinutes != nil {
		return true
	}

	return false
}

// SetWeekendEndTimeMinutes gets a reference to the given int32 and assigns it to the WeekendEndTimeMinutes field.
func (o *PowerSchedule) SetWeekendEndTimeMinutes(v int32) {
	o.WeekendEndTimeMinutes = &v
}

// GetWeekendStartTimeHour returns the WeekendStartTimeHour field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekendStartTimeHour() int32 {
	if o == nil || o.WeekendStartTimeHour == nil {
		var ret int32
		return ret
	}
	return *o.WeekendStartTimeHour
}

// GetWeekendStartTimeHourOk returns a tuple with the WeekendStartTimeHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekendStartTimeHourOk() (*int32, bool) {
	if o == nil || o.WeekendStartTimeHour == nil {
		return nil, false
	}
	return o.WeekendStartTimeHour, true
}

// HasWeekendStartTimeHour returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekendStartTimeHour() bool {
	if o != nil && o.WeekendStartTimeHour != nil {
		return true
	}

	return false
}

// SetWeekendStartTimeHour gets a reference to the given int32 and assigns it to the WeekendStartTimeHour field.
func (o *PowerSchedule) SetWeekendStartTimeHour(v int32) {
	o.WeekendStartTimeHour = &v
}

// GetWeekendStartTimeMinutes returns the WeekendStartTimeMinutes field value if set, zero value otherwise.
func (o *PowerSchedule) GetWeekendStartTimeMinutes() int32 {
	if o == nil || o.WeekendStartTimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.WeekendStartTimeMinutes
}

// GetWeekendStartTimeMinutesOk returns a tuple with the WeekendStartTimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerSchedule) GetWeekendStartTimeMinutesOk() (*int32, bool) {
	if o == nil || o.WeekendStartTimeMinutes == nil {
		return nil, false
	}
	return o.WeekendStartTimeMinutes, true
}

// HasWeekendStartTimeMinutes returns a boolean if a field has been set.
func (o *PowerSchedule) HasWeekendStartTimeMinutes() bool {
	if o != nil && o.WeekendStartTimeMinutes != nil {
		return true
	}

	return false
}

// SetWeekendStartTimeMinutes gets a reference to the given int32 and assigns it to the WeekendStartTimeMinutes field.
func (o *PowerSchedule) SetWeekendStartTimeMinutes(v int32) {
	o.WeekendStartTimeMinutes = &v
}

func (o PowerSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.WeekdayEndTimeHour != nil {
		toSerialize["weekdayEndTimeHour"] = o.WeekdayEndTimeHour
	}
	if o.WeekdayEndTimeMinutes != nil {
		toSerialize["weekdayEndTimeMinutes"] = o.WeekdayEndTimeMinutes
	}
	if o.WeekdayStartTimeHour != nil {
		toSerialize["weekdayStartTimeHour"] = o.WeekdayStartTimeHour
	}
	if o.WeekdayStartTimeMinutes != nil {
		toSerialize["weekdayStartTimeMinutes"] = o.WeekdayStartTimeMinutes
	}
	if o.WeekendEndTimeHour != nil {
		toSerialize["weekendEndTimeHour"] = o.WeekendEndTimeHour
	}
	if o.WeekendEndTimeMinutes != nil {
		toSerialize["weekendEndTimeMinutes"] = o.WeekendEndTimeMinutes
	}
	if o.WeekendStartTimeHour != nil {
		toSerialize["weekendStartTimeHour"] = o.WeekendStartTimeHour
	}
	if o.WeekendStartTimeMinutes != nil {
		toSerialize["weekendStartTimeMinutes"] = o.WeekendStartTimeMinutes
	}
	return json.Marshal(toSerialize)
}

type NullablePowerSchedule struct {
	value *PowerSchedule
	isSet bool
}

func (v NullablePowerSchedule) Get() *PowerSchedule {
	return v.value
}

func (v *NullablePowerSchedule) Set(val *PowerSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerSchedule(val *PowerSchedule) *NullablePowerSchedule {
	return &NullablePowerSchedule{value: val, isSet: true}
}

func (v NullablePowerSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
