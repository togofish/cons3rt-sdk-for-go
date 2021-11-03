# InputProjectFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ItarRestricted** | Pointer to **bool** |  | [optional] 
**OwningTeam** | Pointer to [**InputTeam**](InputTeam.md) |  | [optional] 
**Limits** | [**ProjectLimits**](ProjectLimits.md) |  | 
**DefaultRole** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**ProjectFeatures**](ProjectFeatures.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputProjectFull

`func NewInputProjectFull(name string, limits ProjectLimits, ) *InputProjectFull`

NewInputProjectFull instantiates a new InputProjectFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProjectFullWithDefaults

`func NewInputProjectFullWithDefaults() *InputProjectFull`

NewInputProjectFullWithDefaults instantiates a new InputProjectFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputProjectFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputProjectFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputProjectFull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InputProjectFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputProjectFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputProjectFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputProjectFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItarRestricted

`func (o *InputProjectFull) GetItarRestricted() bool`

GetItarRestricted returns the ItarRestricted field if non-nil, zero value otherwise.

### GetItarRestrictedOk

`func (o *InputProjectFull) GetItarRestrictedOk() (*bool, bool)`

GetItarRestrictedOk returns a tuple with the ItarRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItarRestricted

`func (o *InputProjectFull) SetItarRestricted(v bool)`

SetItarRestricted sets ItarRestricted field to given value.

### HasItarRestricted

`func (o *InputProjectFull) HasItarRestricted() bool`

HasItarRestricted returns a boolean if a field has been set.

### GetOwningTeam

`func (o *InputProjectFull) GetOwningTeam() InputTeam`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *InputProjectFull) GetOwningTeamOk() (*InputTeam, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *InputProjectFull) SetOwningTeam(v InputTeam)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *InputProjectFull) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetLimits

`func (o *InputProjectFull) GetLimits() ProjectLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *InputProjectFull) GetLimitsOk() (*ProjectLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *InputProjectFull) SetLimits(v ProjectLimits)`

SetLimits sets Limits field to given value.


### GetDefaultRole

`func (o *InputProjectFull) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *InputProjectFull) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *InputProjectFull) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *InputProjectFull) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetFeatures

`func (o *InputProjectFull) GetFeatures() ProjectFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InputProjectFull) GetFeaturesOk() (*ProjectFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InputProjectFull) SetFeatures(v ProjectFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InputProjectFull) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIsPrivate

`func (o *InputProjectFull) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *InputProjectFull) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *InputProjectFull) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *InputProjectFull) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


