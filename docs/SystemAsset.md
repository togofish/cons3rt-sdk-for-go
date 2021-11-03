# SystemAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemAssetVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Password** | **string** |  | 
**Url** | **string** |  | 
**Username** | **string** |  | 
**CurrentAsset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**TemplateProfile** | Pointer to [**TemplateProfile**](TemplateProfile.md) |  | [optional] 
**Type** | **string** |  | 
**Softwareversion** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemAsset

`func NewSystemAsset(password string, url string, username string, type_ string, ) *SystemAsset`

NewSystemAsset instantiates a new SystemAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAssetWithDefaults

`func NewSystemAssetWithDefaults() *SystemAsset`

NewSystemAssetWithDefaults instantiates a new SystemAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemAssetVersion

`func (o *SystemAsset) GetSystemAssetVersion() string`

GetSystemAssetVersion returns the SystemAssetVersion field if non-nil, zero value otherwise.

### GetSystemAssetVersionOk

`func (o *SystemAsset) GetSystemAssetVersionOk() (*string, bool)`

GetSystemAssetVersionOk returns a tuple with the SystemAssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAssetVersion

`func (o *SystemAsset) SetSystemAssetVersion(v string)`

SetSystemAssetVersion sets SystemAssetVersion field to given value.

### HasSystemAssetVersion

`func (o *SystemAsset) HasSystemAssetVersion() bool`

HasSystemAssetVersion returns a boolean if a field has been set.

### GetStatus

`func (o *SystemAsset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemAsset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemAsset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemAsset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssetVersion

`func (o *SystemAsset) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *SystemAsset) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *SystemAsset) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *SystemAsset) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetError

`func (o *SystemAsset) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SystemAsset) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SystemAsset) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SystemAsset) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *SystemAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemAsset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SystemAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *SystemAsset) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SystemAsset) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SystemAsset) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUrl

`func (o *SystemAsset) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SystemAsset) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SystemAsset) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *SystemAsset) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SystemAsset) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SystemAsset) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCurrentAsset

`func (o *SystemAsset) GetCurrentAsset() Asset`

GetCurrentAsset returns the CurrentAsset field if non-nil, zero value otherwise.

### GetCurrentAssetOk

`func (o *SystemAsset) GetCurrentAssetOk() (*Asset, bool)`

GetCurrentAssetOk returns a tuple with the CurrentAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAsset

`func (o *SystemAsset) SetCurrentAsset(v Asset)`

SetCurrentAsset sets CurrentAsset field to given value.

### HasCurrentAsset

`func (o *SystemAsset) HasCurrentAsset() bool`

HasCurrentAsset returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *SystemAsset) GetTemplateProfile() TemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *SystemAsset) GetTemplateProfileOk() (*TemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *SystemAsset) SetTemplateProfile(v TemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *SystemAsset) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.

### GetType

`func (o *SystemAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemAsset) SetType(v string)`

SetType sets Type field to given value.


### GetSoftwareversion

`func (o *SystemAsset) GetSoftwareversion() string`

GetSoftwareversion returns the Softwareversion field if non-nil, zero value otherwise.

### GetSoftwareversionOk

`func (o *SystemAsset) GetSoftwareversionOk() (*string, bool)`

GetSoftwareversionOk returns a tuple with the Softwareversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareversion

`func (o *SystemAsset) SetSoftwareversion(v string)`

SetSoftwareversion sets Softwareversion field to given value.

### HasSoftwareversion

`func (o *SystemAsset) HasSoftwareversion() bool`

HasSoftwareversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


