# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetDirectory** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Cloud** | Pointer to [**Cloud**](Cloud.md) |  | [optional] 
**CreationDate** | Pointer to **int32** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InstanceLimit** | Pointer to **int32** |  | [optional] 
**ItarRestricted** | Pointer to **bool** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Modifier** | Pointer to [**User**](User.md) |  | [optional] 
**ModifierDate** | Pointer to **int32** |  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Validated** | Pointer to **bool** |  | [optional] 
**SizeInMegabytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetDirectory

`func (o *Metadata) GetAssetDirectory() string`

GetAssetDirectory returns the AssetDirectory field if non-nil, zero value otherwise.

### GetAssetDirectoryOk

`func (o *Metadata) GetAssetDirectoryOk() (*string, bool)`

GetAssetDirectoryOk returns a tuple with the AssetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDirectory

`func (o *Metadata) SetAssetDirectory(v string)`

SetAssetDirectory sets AssetDirectory field to given value.

### HasAssetDirectory

`func (o *Metadata) HasAssetDirectory() bool`

HasAssetDirectory returns a boolean if a field has been set.

### GetVersion

`func (o *Metadata) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Metadata) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Metadata) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Metadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCloud

`func (o *Metadata) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *Metadata) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *Metadata) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *Metadata) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCreationDate

`func (o *Metadata) GetCreationDate() int32`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Metadata) GetCreationDateOk() (*int32, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Metadata) SetCreationDate(v int32)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Metadata) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDocumentation

`func (o *Metadata) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Metadata) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Metadata) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *Metadata) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetId

`func (o *Metadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Metadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Metadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Metadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceLimit

`func (o *Metadata) GetInstanceLimit() int32`

GetInstanceLimit returns the InstanceLimit field if non-nil, zero value otherwise.

### GetInstanceLimitOk

`func (o *Metadata) GetInstanceLimitOk() (*int32, bool)`

GetInstanceLimitOk returns a tuple with the InstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLimit

`func (o *Metadata) SetInstanceLimit(v int32)`

SetInstanceLimit sets InstanceLimit field to given value.

### HasInstanceLimit

`func (o *Metadata) HasInstanceLimit() bool`

HasInstanceLimit returns a boolean if a field has been set.

### GetItarRestricted

`func (o *Metadata) GetItarRestricted() bool`

GetItarRestricted returns the ItarRestricted field if non-nil, zero value otherwise.

### GetItarRestrictedOk

`func (o *Metadata) GetItarRestrictedOk() (*bool, bool)`

GetItarRestrictedOk returns a tuple with the ItarRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItarRestricted

`func (o *Metadata) SetItarRestricted(v bool)`

SetItarRestricted sets ItarRestricted field to given value.

### HasItarRestricted

`func (o *Metadata) HasItarRestricted() bool`

HasItarRestricted returns a boolean if a field has been set.

### GetLicense

`func (o *Metadata) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Metadata) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Metadata) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Metadata) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetModifier

`func (o *Metadata) GetModifier() User`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *Metadata) GetModifierOk() (*User, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *Metadata) SetModifier(v User)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *Metadata) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetModifierDate

`func (o *Metadata) GetModifierDate() int32`

GetModifierDate returns the ModifierDate field if non-nil, zero value otherwise.

### GetModifierDateOk

`func (o *Metadata) GetModifierDateOk() (*int32, bool)`

GetModifierDateOk returns a tuple with the ModifierDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierDate

`func (o *Metadata) SetModifierDate(v int32)`

SetModifierDate sets ModifierDate field to given value.

### HasModifierDate

`func (o *Metadata) HasModifierDate() bool`

HasModifierDate returns a boolean if a field has been set.

### GetProperties

`func (o *Metadata) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Metadata) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Metadata) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Metadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUri

`func (o *Metadata) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Metadata) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Metadata) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Metadata) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetValidated

`func (o *Metadata) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *Metadata) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *Metadata) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *Metadata) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetSizeInMegabytes

`func (o *Metadata) GetSizeInMegabytes() int32`

GetSizeInMegabytes returns the SizeInMegabytes field if non-nil, zero value otherwise.

### GetSizeInMegabytesOk

`func (o *Metadata) GetSizeInMegabytesOk() (*int32, bool)`

GetSizeInMegabytesOk returns a tuple with the SizeInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInMegabytes

`func (o *Metadata) SetSizeInMegabytes(v int32)`

SetSizeInMegabytes sets SizeInMegabytes field to given value.

### HasSizeInMegabytes

`func (o *Metadata) HasSizeInMegabytes() bool`

HasSizeInMegabytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


