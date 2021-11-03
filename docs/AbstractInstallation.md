# AbstractInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **int32** |  | [optional] 
**AssetName** | Pointer to **string** |  | [optional] 
**AverageInstallationTime** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LoadOrder** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewAbstractInstallation

`func NewAbstractInstallation(subtype string, ) *AbstractInstallation`

NewAbstractInstallation instantiates a new AbstractInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractInstallationWithDefaults

`func NewAbstractInstallationWithDefaults() *AbstractInstallation`

NewAbstractInstallationWithDefaults instantiates a new AbstractInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AbstractInstallation) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AbstractInstallation) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AbstractInstallation) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AbstractInstallation) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetName

`func (o *AbstractInstallation) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AbstractInstallation) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AbstractInstallation) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *AbstractInstallation) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### GetAverageInstallationTime

`func (o *AbstractInstallation) GetAverageInstallationTime() int64`

GetAverageInstallationTime returns the AverageInstallationTime field if non-nil, zero value otherwise.

### GetAverageInstallationTimeOk

`func (o *AbstractInstallation) GetAverageInstallationTimeOk() (*int64, bool)`

GetAverageInstallationTimeOk returns a tuple with the AverageInstallationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInstallationTime

`func (o *AbstractInstallation) SetAverageInstallationTime(v int64)`

SetAverageInstallationTime sets AverageInstallationTime field to given value.

### HasAverageInstallationTime

`func (o *AbstractInstallation) HasAverageInstallationTime() bool`

HasAverageInstallationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *AbstractInstallation) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AbstractInstallation) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AbstractInstallation) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AbstractInstallation) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetError

`func (o *AbstractInstallation) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AbstractInstallation) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AbstractInstallation) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AbstractInstallation) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *AbstractInstallation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractInstallation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractInstallation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractInstallation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadOrder

`func (o *AbstractInstallation) GetLoadOrder() int32`

GetLoadOrder returns the LoadOrder field if non-nil, zero value otherwise.

### GetLoadOrderOk

`func (o *AbstractInstallation) GetLoadOrderOk() (*int32, bool)`

GetLoadOrderOk returns a tuple with the LoadOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOrder

`func (o *AbstractInstallation) SetLoadOrder(v int32)`

SetLoadOrder sets LoadOrder field to given value.

### HasLoadOrder

`func (o *AbstractInstallation) HasLoadOrder() bool`

HasLoadOrder returns a boolean if a field has been set.

### GetStartDate

`func (o *AbstractInstallation) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AbstractInstallation) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AbstractInstallation) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AbstractInstallation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *AbstractInstallation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AbstractInstallation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AbstractInstallation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AbstractInstallation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *AbstractInstallation) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AbstractInstallation) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AbstractInstallation) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


