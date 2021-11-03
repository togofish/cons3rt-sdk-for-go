# RemoteAccessTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Port** | **int32** |  | 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoteAccessTemplate

`func NewRemoteAccessTemplate(name string, type_ string, port int32, ) *RemoteAccessTemplate`

NewRemoteAccessTemplate instantiates a new RemoteAccessTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAccessTemplateWithDefaults

`func NewRemoteAccessTemplateWithDefaults() *RemoteAccessTemplate`

NewRemoteAccessTemplateWithDefaults instantiates a new RemoteAccessTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteAccessTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteAccessTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteAccessTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteAccessTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RemoteAccessTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteAccessTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteAccessTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RemoteAccessTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemoteAccessTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemoteAccessTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetPort

`func (o *RemoteAccessTemplate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemoteAccessTemplate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemoteAccessTemplate) SetPort(v int32)`

SetPort sets Port field to given value.


### GetPassword

`func (o *RemoteAccessTemplate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RemoteAccessTemplate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RemoteAccessTemplate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RemoteAccessTemplate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


