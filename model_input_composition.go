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

// InputComposition struct for InputComposition
type InputComposition struct {
	Name                 string                     `json:"name"`
	Description          *string                    `json:"description,omitempty"`
	DeploymentRunOptions InputCompositionRunOptions `json:"deploymentRunOptions"`
}

// NewInputComposition instantiates a new InputComposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputComposition(name string, deploymentRunOptions InputCompositionRunOptions) *InputComposition {
	this := InputComposition{}
	this.Name = name
	this.DeploymentRunOptions = deploymentRunOptions
	return &this
}

// NewInputCompositionWithDefaults instantiates a new InputComposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCompositionWithDefaults() *InputComposition {
	this := InputComposition{}
	return &this
}

// GetName returns the Name field value
func (o *InputComposition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputComposition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputComposition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputComposition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputComposition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputComposition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputComposition) SetDescription(v string) {
	o.Description = &v
}

// GetDeploymentRunOptions returns the DeploymentRunOptions field value
func (o *InputComposition) GetDeploymentRunOptions() InputCompositionRunOptions {
	if o == nil {
		var ret InputCompositionRunOptions
		return ret
	}

	return o.DeploymentRunOptions
}

// GetDeploymentRunOptionsOk returns a tuple with the DeploymentRunOptions field value
// and a boolean to check if the value has been set.
func (o *InputComposition) GetDeploymentRunOptionsOk() (*InputCompositionRunOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentRunOptions, true
}

// SetDeploymentRunOptions sets field value
func (o *InputComposition) SetDeploymentRunOptions(v InputCompositionRunOptions) {
	o.DeploymentRunOptions = v
}

func (o InputComposition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["deploymentRunOptions"] = o.DeploymentRunOptions
	}
	return json.Marshal(toSerialize)
}

type NullableInputComposition struct {
	value *InputComposition
	isSet bool
}

func (v NullableInputComposition) Get() *InputComposition {
	return v.value
}

func (v *NullableInputComposition) Set(val *InputComposition) {
	v.value = val
	v.isSet = true
}

func (v NullableInputComposition) IsSet() bool {
	return v.isSet
}

func (v *NullableInputComposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputComposition(val *InputComposition) *NullableInputComposition {
	return &NullableInputComposition{value: val, isSet: true}
}

func (v NullableInputComposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputComposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
