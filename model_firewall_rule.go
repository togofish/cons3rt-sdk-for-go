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

// FirewallRule struct for FirewallRule
type FirewallRule struct {
	Id *int32 `json:"id,omitempty"`
	RuleOrder int32 `json:"ruleOrder"`
	Protocol string `json:"protocol"`
	RuleAction string `json:"ruleAction"`
	RuleDestination string `json:"ruleDestination"`
	RuleDestPort *string `json:"ruleDestPort,omitempty"`
	RuleEnabled bool `json:"ruleEnabled"`
	RuleSource string `json:"ruleSource"`
	RuleSourcePort *string `json:"ruleSourcePort,omitempty"`
}

// NewFirewallRule instantiates a new FirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRule(ruleOrder int32, protocol string, ruleAction string, ruleDestination string, ruleEnabled bool, ruleSource string) *FirewallRule {
	this := FirewallRule{}
	this.RuleOrder = ruleOrder
	this.Protocol = protocol
	this.RuleAction = ruleAction
	this.RuleDestination = ruleDestination
	this.RuleEnabled = ruleEnabled
	this.RuleSource = ruleSource
	return &this
}

// NewFirewallRuleWithDefaults instantiates a new FirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleWithDefaults() *FirewallRule {
	this := FirewallRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirewallRule) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirewallRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FirewallRule) SetId(v int32) {
	o.Id = &v
}

// GetRuleOrder returns the RuleOrder field value
func (o *FirewallRule) GetRuleOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleOrder
}

// GetRuleOrderOk returns a tuple with the RuleOrder field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleOrder, true
}

// SetRuleOrder sets field value
func (o *FirewallRule) SetRuleOrder(v int32) {
	o.RuleOrder = v
}

// GetProtocol returns the Protocol field value
func (o *FirewallRule) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetProtocolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *FirewallRule) SetProtocol(v string) {
	o.Protocol = v
}

// GetRuleAction returns the RuleAction field value
func (o *FirewallRule) GetRuleAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleAction
}

// GetRuleActionOk returns a tuple with the RuleAction field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleAction, true
}

// SetRuleAction sets field value
func (o *FirewallRule) SetRuleAction(v string) {
	o.RuleAction = v
}

// GetRuleDestination returns the RuleDestination field value
func (o *FirewallRule) GetRuleDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleDestination
}

// GetRuleDestinationOk returns a tuple with the RuleDestination field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleDestination, true
}

// SetRuleDestination sets field value
func (o *FirewallRule) SetRuleDestination(v string) {
	o.RuleDestination = v
}

// GetRuleDestPort returns the RuleDestPort field value if set, zero value otherwise.
func (o *FirewallRule) GetRuleDestPort() string {
	if o == nil || o.RuleDestPort == nil {
		var ret string
		return ret
	}
	return *o.RuleDestPort
}

// GetRuleDestPortOk returns a tuple with the RuleDestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleDestPortOk() (*string, bool) {
	if o == nil || o.RuleDestPort == nil {
		return nil, false
	}
	return o.RuleDestPort, true
}

// HasRuleDestPort returns a boolean if a field has been set.
func (o *FirewallRule) HasRuleDestPort() bool {
	if o != nil && o.RuleDestPort != nil {
		return true
	}

	return false
}

// SetRuleDestPort gets a reference to the given string and assigns it to the RuleDestPort field.
func (o *FirewallRule) SetRuleDestPort(v string) {
	o.RuleDestPort = &v
}

// GetRuleEnabled returns the RuleEnabled field value
func (o *FirewallRule) GetRuleEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RuleEnabled
}

// GetRuleEnabledOk returns a tuple with the RuleEnabled field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleEnabled, true
}

// SetRuleEnabled sets field value
func (o *FirewallRule) SetRuleEnabled(v bool) {
	o.RuleEnabled = v
}

// GetRuleSource returns the RuleSource field value
func (o *FirewallRule) GetRuleSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleSource
}

// GetRuleSourceOk returns a tuple with the RuleSource field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleSource, true
}

// SetRuleSource sets field value
func (o *FirewallRule) SetRuleSource(v string) {
	o.RuleSource = v
}

// GetRuleSourcePort returns the RuleSourcePort field value if set, zero value otherwise.
func (o *FirewallRule) GetRuleSourcePort() string {
	if o == nil || o.RuleSourcePort == nil {
		var ret string
		return ret
	}
	return *o.RuleSourcePort
}

// GetRuleSourcePortOk returns a tuple with the RuleSourcePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetRuleSourcePortOk() (*string, bool) {
	if o == nil || o.RuleSourcePort == nil {
		return nil, false
	}
	return o.RuleSourcePort, true
}

// HasRuleSourcePort returns a boolean if a field has been set.
func (o *FirewallRule) HasRuleSourcePort() bool {
	if o != nil && o.RuleSourcePort != nil {
		return true
	}

	return false
}

// SetRuleSourcePort gets a reference to the given string and assigns it to the RuleSourcePort field.
func (o *FirewallRule) SetRuleSourcePort(v string) {
	o.RuleSourcePort = &v
}

func (o FirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ruleOrder"] = o.RuleOrder
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["ruleAction"] = o.RuleAction
	}
	if true {
		toSerialize["ruleDestination"] = o.RuleDestination
	}
	if o.RuleDestPort != nil {
		toSerialize["ruleDestPort"] = o.RuleDestPort
	}
	if true {
		toSerialize["ruleEnabled"] = o.RuleEnabled
	}
	if true {
		toSerialize["ruleSource"] = o.RuleSource
	}
	if o.RuleSourcePort != nil {
		toSerialize["ruleSourcePort"] = o.RuleSourcePort
	}
	return json.Marshal(toSerialize)
}

type NullableFirewallRule struct {
	value *FirewallRule
	isSet bool
}

func (v NullableFirewallRule) Get() *FirewallRule {
	return v.value
}

func (v *NullableFirewallRule) Set(val *FirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRule(val *FirewallRule) *NullableFirewallRule {
	return &NullableFirewallRule{value: val, isSet: true}
}

func (v NullableFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


