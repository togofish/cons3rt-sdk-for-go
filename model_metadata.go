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

// Metadata struct for Metadata
type Metadata struct {
	AssetDirectory  *string     `json:"assetDirectory,omitempty"`
	Version         *int32      `json:"version,omitempty"`
	Cloud           *Cloud      `json:"cloud,omitempty"`
	CreationDate    *int32      `json:"creationDate,omitempty"`
	Documentation   *string     `json:"documentation,omitempty"`
	Id              *int32      `json:"id,omitempty"`
	InstanceLimit   *int32      `json:"instanceLimit,omitempty"`
	ItarRestricted  *bool       `json:"itarRestricted,omitempty"`
	License         *string     `json:"license,omitempty"`
	Modifier        *User       `json:"modifier,omitempty"`
	ModifierDate    *int32      `json:"modifierDate,omitempty"`
	Properties      *[]Property `json:"properties,omitempty"`
	Uri             *string     `json:"uri,omitempty"`
	Validated       *bool       `json:"validated,omitempty"`
	SizeInMegabytes *int32      `json:"sizeInMegabytes,omitempty"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetAssetDirectory returns the AssetDirectory field value if set, zero value otherwise.
func (o *Metadata) GetAssetDirectory() string {
	if o == nil || o.AssetDirectory == nil {
		var ret string
		return ret
	}
	return *o.AssetDirectory
}

// GetAssetDirectoryOk returns a tuple with the AssetDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetAssetDirectoryOk() (*string, bool) {
	if o == nil || o.AssetDirectory == nil {
		return nil, false
	}
	return o.AssetDirectory, true
}

// HasAssetDirectory returns a boolean if a field has been set.
func (o *Metadata) HasAssetDirectory() bool {
	if o != nil && o.AssetDirectory != nil {
		return true
	}

	return false
}

// SetAssetDirectory gets a reference to the given string and assigns it to the AssetDirectory field.
func (o *Metadata) SetAssetDirectory(v string) {
	o.AssetDirectory = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Metadata) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Metadata) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Metadata) SetVersion(v int32) {
	o.Version = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *Metadata) GetCloud() Cloud {
	if o == nil || o.Cloud == nil {
		var ret Cloud
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetCloudOk() (*Cloud, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *Metadata) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given Cloud and assigns it to the Cloud field.
func (o *Metadata) SetCloud(v Cloud) {
	o.Cloud = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Metadata) GetCreationDate() int32 {
	if o == nil || o.CreationDate == nil {
		var ret int32
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetCreationDateOk() (*int32, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Metadata) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int32 and assigns it to the CreationDate field.
func (o *Metadata) SetCreationDate(v int32) {
	o.CreationDate = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *Metadata) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *Metadata) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *Metadata) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Metadata) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Metadata) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Metadata) SetId(v int32) {
	o.Id = &v
}

// GetInstanceLimit returns the InstanceLimit field value if set, zero value otherwise.
func (o *Metadata) GetInstanceLimit() int32 {
	if o == nil || o.InstanceLimit == nil {
		var ret int32
		return ret
	}
	return *o.InstanceLimit
}

// GetInstanceLimitOk returns a tuple with the InstanceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetInstanceLimitOk() (*int32, bool) {
	if o == nil || o.InstanceLimit == nil {
		return nil, false
	}
	return o.InstanceLimit, true
}

// HasInstanceLimit returns a boolean if a field has been set.
func (o *Metadata) HasInstanceLimit() bool {
	if o != nil && o.InstanceLimit != nil {
		return true
	}

	return false
}

// SetInstanceLimit gets a reference to the given int32 and assigns it to the InstanceLimit field.
func (o *Metadata) SetInstanceLimit(v int32) {
	o.InstanceLimit = &v
}

// GetItarRestricted returns the ItarRestricted field value if set, zero value otherwise.
func (o *Metadata) GetItarRestricted() bool {
	if o == nil || o.ItarRestricted == nil {
		var ret bool
		return ret
	}
	return *o.ItarRestricted
}

// GetItarRestrictedOk returns a tuple with the ItarRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetItarRestrictedOk() (*bool, bool) {
	if o == nil || o.ItarRestricted == nil {
		return nil, false
	}
	return o.ItarRestricted, true
}

// HasItarRestricted returns a boolean if a field has been set.
func (o *Metadata) HasItarRestricted() bool {
	if o != nil && o.ItarRestricted != nil {
		return true
	}

	return false
}

// SetItarRestricted gets a reference to the given bool and assigns it to the ItarRestricted field.
func (o *Metadata) SetItarRestricted(v bool) {
	o.ItarRestricted = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *Metadata) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *Metadata) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *Metadata) SetLicense(v string) {
	o.License = &v
}

// GetModifier returns the Modifier field value if set, zero value otherwise.
func (o *Metadata) GetModifier() User {
	if o == nil || o.Modifier == nil {
		var ret User
		return ret
	}
	return *o.Modifier
}

// GetModifierOk returns a tuple with the Modifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetModifierOk() (*User, bool) {
	if o == nil || o.Modifier == nil {
		return nil, false
	}
	return o.Modifier, true
}

// HasModifier returns a boolean if a field has been set.
func (o *Metadata) HasModifier() bool {
	if o != nil && o.Modifier != nil {
		return true
	}

	return false
}

// SetModifier gets a reference to the given User and assigns it to the Modifier field.
func (o *Metadata) SetModifier(v User) {
	o.Modifier = &v
}

// GetModifierDate returns the ModifierDate field value if set, zero value otherwise.
func (o *Metadata) GetModifierDate() int32 {
	if o == nil || o.ModifierDate == nil {
		var ret int32
		return ret
	}
	return *o.ModifierDate
}

// GetModifierDateOk returns a tuple with the ModifierDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetModifierDateOk() (*int32, bool) {
	if o == nil || o.ModifierDate == nil {
		return nil, false
	}
	return o.ModifierDate, true
}

// HasModifierDate returns a boolean if a field has been set.
func (o *Metadata) HasModifierDate() bool {
	if o != nil && o.ModifierDate != nil {
		return true
	}

	return false
}

// SetModifierDate gets a reference to the given int32 and assigns it to the ModifierDate field.
func (o *Metadata) SetModifierDate(v int32) {
	o.ModifierDate = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Metadata) GetProperties() []Property {
	if o == nil || o.Properties == nil {
		var ret []Property
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetPropertiesOk() (*[]Property, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Metadata) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *Metadata) SetProperties(v []Property) {
	o.Properties = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Metadata) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Metadata) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Metadata) SetUri(v string) {
	o.Uri = &v
}

// GetValidated returns the Validated field value if set, zero value otherwise.
func (o *Metadata) GetValidated() bool {
	if o == nil || o.Validated == nil {
		var ret bool
		return ret
	}
	return *o.Validated
}

// GetValidatedOk returns a tuple with the Validated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetValidatedOk() (*bool, bool) {
	if o == nil || o.Validated == nil {
		return nil, false
	}
	return o.Validated, true
}

// HasValidated returns a boolean if a field has been set.
func (o *Metadata) HasValidated() bool {
	if o != nil && o.Validated != nil {
		return true
	}

	return false
}

// SetValidated gets a reference to the given bool and assigns it to the Validated field.
func (o *Metadata) SetValidated(v bool) {
	o.Validated = &v
}

// GetSizeInMegabytes returns the SizeInMegabytes field value if set, zero value otherwise.
func (o *Metadata) GetSizeInMegabytes() int32 {
	if o == nil || o.SizeInMegabytes == nil {
		var ret int32
		return ret
	}
	return *o.SizeInMegabytes
}

// GetSizeInMegabytesOk returns a tuple with the SizeInMegabytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetSizeInMegabytesOk() (*int32, bool) {
	if o == nil || o.SizeInMegabytes == nil {
		return nil, false
	}
	return o.SizeInMegabytes, true
}

// HasSizeInMegabytes returns a boolean if a field has been set.
func (o *Metadata) HasSizeInMegabytes() bool {
	if o != nil && o.SizeInMegabytes != nil {
		return true
	}

	return false
}

// SetSizeInMegabytes gets a reference to the given int32 and assigns it to the SizeInMegabytes field.
func (o *Metadata) SetSizeInMegabytes(v int32) {
	o.SizeInMegabytes = &v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetDirectory != nil {
		toSerialize["assetDirectory"] = o.AssetDirectory
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InstanceLimit != nil {
		toSerialize["instanceLimit"] = o.InstanceLimit
	}
	if o.ItarRestricted != nil {
		toSerialize["itarRestricted"] = o.ItarRestricted
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.Modifier != nil {
		toSerialize["modifier"] = o.Modifier
	}
	if o.ModifierDate != nil {
		toSerialize["modifierDate"] = o.ModifierDate
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Validated != nil {
		toSerialize["validated"] = o.Validated
	}
	if o.SizeInMegabytes != nil {
		toSerialize["sizeInMegabytes"] = o.SizeInMegabytes
	}
	return json.Marshal(toSerialize)
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
