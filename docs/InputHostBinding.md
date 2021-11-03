# InputHostBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRole** | **string** |  | 
**TemplateName** | **string** |  | 
**InstanceType** | Pointer to **string** | Required for Azure and AWS EC2 Cloudspaces | [optional] 

## Methods

### NewInputHostBinding

`func NewInputHostBinding(systemRole string, templateName string, ) *InputHostBinding`

NewInputHostBinding instantiates a new InputHostBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputHostBindingWithDefaults

`func NewInputHostBindingWithDefaults() *InputHostBinding`

NewInputHostBindingWithDefaults instantiates a new InputHostBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRole

`func (o *InputHostBinding) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *InputHostBinding) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *InputHostBinding) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.


### GetTemplateName

`func (o *InputHostBinding) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *InputHostBinding) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *InputHostBinding) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetInstanceType

`func (o *InputHostBinding) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InputHostBinding) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InputHostBinding) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *InputHostBinding) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


