# SoftwareInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to **string** |  | [optional] 
**RebootDelaySeconds** | Pointer to **int32** |  | [optional] 
**RebootRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewSoftwareInstallation

`func NewSoftwareInstallation() *SoftwareInstallation`

NewSoftwareInstallation instantiates a new SoftwareInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareInstallationWithDefaults

`func NewSoftwareInstallationWithDefaults() *SoftwareInstallation`

NewSoftwareInstallationWithDefaults instantiates a new SoftwareInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *SoftwareInstallation) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *SoftwareInstallation) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *SoftwareInstallation) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *SoftwareInstallation) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetRebootDelaySeconds

`func (o *SoftwareInstallation) GetRebootDelaySeconds() int32`

GetRebootDelaySeconds returns the RebootDelaySeconds field if non-nil, zero value otherwise.

### GetRebootDelaySecondsOk

`func (o *SoftwareInstallation) GetRebootDelaySecondsOk() (*int32, bool)`

GetRebootDelaySecondsOk returns a tuple with the RebootDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDelaySeconds

`func (o *SoftwareInstallation) SetRebootDelaySeconds(v int32)`

SetRebootDelaySeconds sets RebootDelaySeconds field to given value.

### HasRebootDelaySeconds

`func (o *SoftwareInstallation) HasRebootDelaySeconds() bool`

HasRebootDelaySeconds returns a boolean if a field has been set.

### GetRebootRequired

`func (o *SoftwareInstallation) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *SoftwareInstallation) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *SoftwareInstallation) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.

### HasRebootRequired

`func (o *SoftwareInstallation) HasRebootRequired() bool`

HasRebootRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


