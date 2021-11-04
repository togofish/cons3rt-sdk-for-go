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

// Subnet struct for Subnet
type Subnet struct {
	Cidr string `json:"cidr"`
	Id *int32 `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	Name string `json:"name"`
	NatInstance NatInstance `json:"natInstance"`
	SecurityGroup SecurityGroup `json:"securityGroup"`
}

// NewSubnet instantiates a new Subnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnet(cidr string, identifier string, name string, natInstance NatInstance, securityGroup SecurityGroup) *Subnet {
	this := Subnet{}
	this.Cidr = cidr
	this.Identifier = identifier
	this.Name = name
	this.NatInstance = natInstance
	this.SecurityGroup = securityGroup
	return &this
}

// NewSubnetWithDefaults instantiates a new Subnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetWithDefaults() *Subnet {
	this := Subnet{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *Subnet) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetCidrOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *Subnet) SetCidr(v string) {
	o.Cidr = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subnet) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subnet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Subnet) SetId(v int32) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value
func (o *Subnet) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Subnet) SetIdentifier(v string) {
	o.Identifier = v
}

// GetName returns the Name field value
func (o *Subnet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subnet) SetName(v string) {
	o.Name = v
}

// GetNatInstance returns the NatInstance field value
func (o *Subnet) GetNatInstance() NatInstance {
	if o == nil {
		var ret NatInstance
		return ret
	}

	return o.NatInstance
}

// GetNatInstanceOk returns a tuple with the NatInstance field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetNatInstanceOk() (*NatInstance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NatInstance, true
}

// SetNatInstance sets field value
func (o *Subnet) SetNatInstance(v NatInstance) {
	o.NatInstance = v
}

// GetSecurityGroup returns the SecurityGroup field value
func (o *Subnet) GetSecurityGroup() SecurityGroup {
	if o == nil {
		var ret SecurityGroup
		return ret
	}

	return o.SecurityGroup
}

// GetSecurityGroupOk returns a tuple with the SecurityGroup field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetSecurityGroupOk() (*SecurityGroup, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityGroup, true
}

// SetSecurityGroup sets field value
func (o *Subnet) SetSecurityGroup(v SecurityGroup) {
	o.SecurityGroup = v
}

func (o Subnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["natInstance"] = o.NatInstance
	}
	if true {
		toSerialize["securityGroup"] = o.SecurityGroup
	}
	return json.Marshal(toSerialize)
}

type NullableSubnet struct {
	value *Subnet
	isSet bool
}

func (v NullableSubnet) Get() *Subnet {
	return v.value
}

func (v *NullableSubnet) Set(val *Subnet) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnet(val *Subnet) *NullableSubnet {
	return &NullableSubnet{value: val, isSet: true}
}

func (v NullableSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


