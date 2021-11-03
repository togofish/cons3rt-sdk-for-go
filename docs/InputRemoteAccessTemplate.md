# InputRemoteAccessTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Port** | **int32** |  | 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewInputRemoteAccessTemplate

`func NewInputRemoteAccessTemplate(name string, type_ string, port int32, ) *InputRemoteAccessTemplate`

NewInputRemoteAccessTemplate instantiates a new InputRemoteAccessTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRemoteAccessTemplateWithDefaults

`func NewInputRemoteAccessTemplateWithDefaults() *InputRemoteAccessTemplate`

NewInputRemoteAccessTemplateWithDefaults instantiates a new InputRemoteAccessTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputRemoteAccessTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputRemoteAccessTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputRemoteAccessTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InputRemoteAccessTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputRemoteAccessTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputRemoteAccessTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetPort

`func (o *InputRemoteAccessTemplate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InputRemoteAccessTemplate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InputRemoteAccessTemplate) SetPort(v int32)`

SetPort sets Port field to given value.


### GetPassword

`func (o *InputRemoteAccessTemplate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputRemoteAccessTemplate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputRemoteAccessTemplate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InputRemoteAccessTemplate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


