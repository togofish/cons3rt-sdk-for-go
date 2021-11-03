# MinimalUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewMinimalUser

`func NewMinimalUser() *MinimalUser`

NewMinimalUser instantiates a new MinimalUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalUserWithDefaults

`func NewMinimalUserWithDefaults() *MinimalUser`

NewMinimalUserWithDefaults instantiates a new MinimalUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *MinimalUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MinimalUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MinimalUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MinimalUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *MinimalUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MinimalUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MinimalUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MinimalUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


