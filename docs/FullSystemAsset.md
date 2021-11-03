# FullSystemAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemAssetVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**CurrentAsset** | Pointer to [**MinimalAsset**](MinimalAsset.md) |  | [optional] 
**TemplateProfile** | Pointer to [**MinimalTemplateProfile**](MinimalTemplateProfile.md) |  | [optional] 
**Type** | **string** |  | 
**Softwareversion** | Pointer to **string** |  | [optional] 

## Methods

### NewFullSystemAsset

`func NewFullSystemAsset(type_ string, ) *FullSystemAsset`

NewFullSystemAsset instantiates a new FullSystemAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullSystemAssetWithDefaults

`func NewFullSystemAssetWithDefaults() *FullSystemAsset`

NewFullSystemAssetWithDefaults instantiates a new FullSystemAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemAssetVersion

`func (o *FullSystemAsset) GetSystemAssetVersion() string`

GetSystemAssetVersion returns the SystemAssetVersion field if non-nil, zero value otherwise.

### GetSystemAssetVersionOk

`func (o *FullSystemAsset) GetSystemAssetVersionOk() (*string, bool)`

GetSystemAssetVersionOk returns a tuple with the SystemAssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAssetVersion

`func (o *FullSystemAsset) SetSystemAssetVersion(v string)`

SetSystemAssetVersion sets SystemAssetVersion field to given value.

### HasSystemAssetVersion

`func (o *FullSystemAsset) HasSystemAssetVersion() bool`

HasSystemAssetVersion returns a boolean if a field has been set.

### GetStatus

`func (o *FullSystemAsset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FullSystemAsset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FullSystemAsset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FullSystemAsset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssetVersion

`func (o *FullSystemAsset) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *FullSystemAsset) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *FullSystemAsset) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *FullSystemAsset) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetError

`func (o *FullSystemAsset) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FullSystemAsset) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FullSystemAsset) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FullSystemAsset) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *FullSystemAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullSystemAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullSystemAsset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FullSystemAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *FullSystemAsset) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FullSystemAsset) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FullSystemAsset) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FullSystemAsset) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *FullSystemAsset) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FullSystemAsset) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FullSystemAsset) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FullSystemAsset) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *FullSystemAsset) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FullSystemAsset) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FullSystemAsset) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FullSystemAsset) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCurrentAsset

`func (o *FullSystemAsset) GetCurrentAsset() MinimalAsset`

GetCurrentAsset returns the CurrentAsset field if non-nil, zero value otherwise.

### GetCurrentAssetOk

`func (o *FullSystemAsset) GetCurrentAssetOk() (*MinimalAsset, bool)`

GetCurrentAssetOk returns a tuple with the CurrentAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAsset

`func (o *FullSystemAsset) SetCurrentAsset(v MinimalAsset)`

SetCurrentAsset sets CurrentAsset field to given value.

### HasCurrentAsset

`func (o *FullSystemAsset) HasCurrentAsset() bool`

HasCurrentAsset returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *FullSystemAsset) GetTemplateProfile() MinimalTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *FullSystemAsset) GetTemplateProfileOk() (*MinimalTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *FullSystemAsset) SetTemplateProfile(v MinimalTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *FullSystemAsset) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.

### GetType

`func (o *FullSystemAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullSystemAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullSystemAsset) SetType(v string)`

SetType sets Type field to given value.


### GetSoftwareversion

`func (o *FullSystemAsset) GetSoftwareversion() string`

GetSoftwareversion returns the Softwareversion field if non-nil, zero value otherwise.

### GetSoftwareversionOk

`func (o *FullSystemAsset) GetSoftwareversionOk() (*string, bool)`

GetSoftwareversionOk returns a tuple with the Softwareversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareversion

`func (o *FullSystemAsset) SetSoftwareversion(v string)`

SetSoftwareversion sets Softwareversion field to given value.

### HasSoftwareversion

`func (o *FullSystemAsset) HasSoftwareversion() bool`

HasSoftwareversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


