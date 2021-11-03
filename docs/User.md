# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**AdministeredClouds** | Pointer to [**[]Cloud**](Cloud.md) |  | [optional] 
**AdministeredVirtRealms** | Pointer to [**[]VirtualizationRealm**](VirtualizationRealm.md) |  | [optional] 
**Certificates** | Pointer to [**[]Certificate**](Certificate.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DefaultProject** | Pointer to [**Project**](Project.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**LogEntries** | Pointer to [**[]LogEntry**](LogEntry.md) |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**ProjectCount** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TermsOfServiceAccepted** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *User) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAdministeredClouds

`func (o *User) GetAdministeredClouds() []Cloud`

GetAdministeredClouds returns the AdministeredClouds field if non-nil, zero value otherwise.

### GetAdministeredCloudsOk

`func (o *User) GetAdministeredCloudsOk() (*[]Cloud, bool)`

GetAdministeredCloudsOk returns a tuple with the AdministeredClouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministeredClouds

`func (o *User) SetAdministeredClouds(v []Cloud)`

SetAdministeredClouds sets AdministeredClouds field to given value.

### HasAdministeredClouds

`func (o *User) HasAdministeredClouds() bool`

HasAdministeredClouds returns a boolean if a field has been set.

### GetAdministeredVirtRealms

`func (o *User) GetAdministeredVirtRealms() []VirtualizationRealm`

GetAdministeredVirtRealms returns the AdministeredVirtRealms field if non-nil, zero value otherwise.

### GetAdministeredVirtRealmsOk

`func (o *User) GetAdministeredVirtRealmsOk() (*[]VirtualizationRealm, bool)`

GetAdministeredVirtRealmsOk returns a tuple with the AdministeredVirtRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministeredVirtRealms

`func (o *User) SetAdministeredVirtRealms(v []VirtualizationRealm)`

SetAdministeredVirtRealms sets AdministeredVirtRealms field to given value.

### HasAdministeredVirtRealms

`func (o *User) HasAdministeredVirtRealms() bool`

HasAdministeredVirtRealms returns a boolean if a field has been set.

### GetCertificates

`func (o *User) GetCertificates() []Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *User) GetCertificatesOk() (*[]Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *User) SetCertificates(v []Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *User) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetComment

`func (o *User) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *User) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *User) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *User) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultProject

`func (o *User) GetDefaultProject() Project`

GetDefaultProject returns the DefaultProject field if non-nil, zero value otherwise.

### GetDefaultProjectOk

`func (o *User) GetDefaultProjectOk() (*Project, bool)`

GetDefaultProjectOk returns a tuple with the DefaultProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProject

`func (o *User) SetDefaultProject(v Project)`

SetDefaultProject sets DefaultProject field to given value.

### HasDefaultProject

`func (o *User) HasDefaultProject() bool`

HasDefaultProject returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *User) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *User) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *User) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *User) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastname

`func (o *User) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *User) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *User) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *User) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLogEntries

`func (o *User) GetLogEntries() []LogEntry`

GetLogEntries returns the LogEntries field if non-nil, zero value otherwise.

### GetLogEntriesOk

`func (o *User) GetLogEntriesOk() (*[]LogEntry, bool)`

GetLogEntriesOk returns a tuple with the LogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEntries

`func (o *User) SetLogEntries(v []LogEntry)`

SetLogEntries sets LogEntries field to given value.

### HasLogEntries

`func (o *User) HasLogEntries() bool`

HasLogEntries returns a boolean if a field has been set.

### GetOrganization

`func (o *User) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *User) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *User) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *User) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjectCount

`func (o *User) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *User) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *User) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.

### HasProjectCount

`func (o *User) HasProjectCount() bool`

HasProjectCount returns a boolean if a field has been set.

### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTermsOfServiceAccepted

`func (o *User) GetTermsOfServiceAccepted() bool`

GetTermsOfServiceAccepted returns the TermsOfServiceAccepted field if non-nil, zero value otherwise.

### GetTermsOfServiceAcceptedOk

`func (o *User) GetTermsOfServiceAcceptedOk() (*bool, bool)`

GetTermsOfServiceAcceptedOk returns a tuple with the TermsOfServiceAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceAccepted

`func (o *User) SetTermsOfServiceAccepted(v bool)`

SetTermsOfServiceAccepted sets TermsOfServiceAccepted field to given value.

### HasTermsOfServiceAccepted

`func (o *User) HasTermsOfServiceAccepted() bool`

HasTermsOfServiceAccepted returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


