# BaseIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**MinimalUser**](MinimalUser.md) |  | [optional] 
**Project** | Pointer to [**MinimalProject**](MinimalProject.md) |  | [optional] 
**AccessListing** | Pointer to [**[]CloudResourceAccessListing**](CloudResourceAccessListing.md) |  | [optional] 

## Methods

### NewBaseIdentity

`func NewBaseIdentity() *BaseIdentity`

NewBaseIdentity instantiates a new BaseIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseIdentityWithDefaults

`func NewBaseIdentityWithDefaults() *BaseIdentity`

NewBaseIdentityWithDefaults instantiates a new BaseIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *BaseIdentity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BaseIdentity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BaseIdentity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BaseIdentity) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetUser

`func (o *BaseIdentity) GetUser() MinimalUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BaseIdentity) GetUserOk() (*MinimalUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BaseIdentity) SetUser(v MinimalUser)`

SetUser sets User field to given value.

### HasUser

`func (o *BaseIdentity) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProject

`func (o *BaseIdentity) GetProject() MinimalProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BaseIdentity) GetProjectOk() (*MinimalProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BaseIdentity) SetProject(v MinimalProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *BaseIdentity) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccessListing

`func (o *BaseIdentity) GetAccessListing() []CloudResourceAccessListing`

GetAccessListing returns the AccessListing field if non-nil, zero value otherwise.

### GetAccessListingOk

`func (o *BaseIdentity) GetAccessListingOk() (*[]CloudResourceAccessListing, bool)`

GetAccessListingOk returns a tuple with the AccessListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListing

`func (o *BaseIdentity) SetAccessListing(v []CloudResourceAccessListing)`

SetAccessListing sets AccessListing field to given value.

### HasAccessListing

`func (o *BaseIdentity) HasAccessListing() bool`

HasAccessListing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


