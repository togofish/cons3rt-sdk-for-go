# NatInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIp** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**NatSecurityGroupIdentifier** | Pointer to **string** |  | [optional] 
**SshKeyName** | Pointer to **string** |  | [optional] 
**SshKeyPem** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewNatInstance

`func NewNatInstance(externalIp string, ) *NatInstance`

NewNatInstance instantiates a new NatInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatInstanceWithDefaults

`func NewNatInstanceWithDefaults() *NatInstance`

NewNatInstanceWithDefaults instantiates a new NatInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIp

`func (o *NatInstance) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *NatInstance) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *NatInstance) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.


### GetIdentifier

`func (o *NatInstance) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NatInstance) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NatInstance) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *NatInstance) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInternalIp

`func (o *NatInstance) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *NatInstance) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *NatInstance) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *NatInstance) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetPassword

`func (o *NatInstance) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NatInstance) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NatInstance) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NatInstance) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNatSecurityGroupIdentifier

`func (o *NatInstance) GetNatSecurityGroupIdentifier() string`

GetNatSecurityGroupIdentifier returns the NatSecurityGroupIdentifier field if non-nil, zero value otherwise.

### GetNatSecurityGroupIdentifierOk

`func (o *NatInstance) GetNatSecurityGroupIdentifierOk() (*string, bool)`

GetNatSecurityGroupIdentifierOk returns a tuple with the NatSecurityGroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSecurityGroupIdentifier

`func (o *NatInstance) SetNatSecurityGroupIdentifier(v string)`

SetNatSecurityGroupIdentifier sets NatSecurityGroupIdentifier field to given value.

### HasNatSecurityGroupIdentifier

`func (o *NatInstance) HasNatSecurityGroupIdentifier() bool`

HasNatSecurityGroupIdentifier returns a boolean if a field has been set.

### GetSshKeyName

`func (o *NatInstance) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *NatInstance) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *NatInstance) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.

### HasSshKeyName

`func (o *NatInstance) HasSshKeyName() bool`

HasSshKeyName returns a boolean if a field has been set.

### GetSshKeyPem

`func (o *NatInstance) GetSshKeyPem() string`

GetSshKeyPem returns the SshKeyPem field if non-nil, zero value otherwise.

### GetSshKeyPemOk

`func (o *NatInstance) GetSshKeyPemOk() (*string, bool)`

GetSshKeyPemOk returns a tuple with the SshKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPem

`func (o *NatInstance) SetSshKeyPem(v string)`

SetSshKeyPem sets SshKeyPem field to given value.

### HasSshKeyPem

`func (o *NatInstance) HasSshKeyPem() bool`

HasSshKeyPem returns a boolean if a field has been set.

### GetUsername

`func (o *NatInstance) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NatInstance) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NatInstance) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NatInstance) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


