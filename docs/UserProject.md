# UserProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**MembershipState** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserProject

`func NewUserProject() *UserProject`

NewUserProject instantiates a new UserProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProjectWithDefaults

`func NewUserProjectWithDefaults() *UserProject`

NewUserProjectWithDefaults instantiates a new UserProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UserProject) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserProject) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserProject) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserProject) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserProject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMembershipState

`func (o *UserProject) GetMembershipState() string`

GetMembershipState returns the MembershipState field if non-nil, zero value otherwise.

### GetMembershipStateOk

`func (o *UserProject) GetMembershipStateOk() (*string, bool)`

GetMembershipStateOk returns a tuple with the MembershipState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipState

`func (o *UserProject) SetMembershipState(v string)`

SetMembershipState sets MembershipState field to given value.

### HasMembershipState

`func (o *UserProject) HasMembershipState() bool`

HasMembershipState returns a boolean if a field has been set.

### GetRoles

`func (o *UserProject) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserProject) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserProject) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserProject) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


