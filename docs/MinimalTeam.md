# MinimalTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Private** | Pointer to **bool** |  | [optional] 
**State** | **string** |  | 
**ValidUtil** | Pointer to **int32** |  | [optional] 
**ContactInfo** | [**PocInfo**](PocInfo.md) |  | 

## Methods

### NewMinimalTeam

`func NewMinimalTeam(name string, state string, contactInfo PocInfo, ) *MinimalTeam`

NewMinimalTeam instantiates a new MinimalTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalTeamWithDefaults

`func NewMinimalTeamWithDefaults() *MinimalTeam`

NewMinimalTeamWithDefaults instantiates a new MinimalTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MinimalTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalTeam) SetName(v string)`

SetName sets Name field to given value.


### GetPrivate

`func (o *MinimalTeam) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *MinimalTeam) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *MinimalTeam) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *MinimalTeam) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetState

`func (o *MinimalTeam) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MinimalTeam) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MinimalTeam) SetState(v string)`

SetState sets State field to given value.


### GetValidUtil

`func (o *MinimalTeam) GetValidUtil() int32`

GetValidUtil returns the ValidUtil field if non-nil, zero value otherwise.

### GetValidUtilOk

`func (o *MinimalTeam) GetValidUtilOk() (*int32, bool)`

GetValidUtilOk returns a tuple with the ValidUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUtil

`func (o *MinimalTeam) SetValidUtil(v int32)`

SetValidUtil sets ValidUtil field to given value.

### HasValidUtil

`func (o *MinimalTeam) HasValidUtil() bool`

HasValidUtil returns a boolean if a field has been set.

### GetContactInfo

`func (o *MinimalTeam) GetContactInfo() PocInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *MinimalTeam) GetContactInfoOk() (*PocInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *MinimalTeam) SetContactInfo(v PocInfo)`

SetContactInfo sets ContactInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


