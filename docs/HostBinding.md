# HostBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualizationRealmTemplates** | [**[]TemplateBinding**](TemplateBinding.md) |  | 
**SystemRole** | **string** |  | 

## Methods

### NewHostBinding

`func NewHostBinding(virtualizationRealmTemplates []TemplateBinding, systemRole string, ) *HostBinding`

NewHostBinding instantiates a new HostBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostBindingWithDefaults

`func NewHostBindingWithDefaults() *HostBinding`

NewHostBindingWithDefaults instantiates a new HostBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualizationRealmTemplates

`func (o *HostBinding) GetVirtualizationRealmTemplates() []TemplateBinding`

GetVirtualizationRealmTemplates returns the VirtualizationRealmTemplates field if non-nil, zero value otherwise.

### GetVirtualizationRealmTemplatesOk

`func (o *HostBinding) GetVirtualizationRealmTemplatesOk() (*[]TemplateBinding, bool)`

GetVirtualizationRealmTemplatesOk returns a tuple with the VirtualizationRealmTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmTemplates

`func (o *HostBinding) SetVirtualizationRealmTemplates(v []TemplateBinding)`

SetVirtualizationRealmTemplates sets VirtualizationRealmTemplates field to given value.


### GetSystemRole

`func (o *HostBinding) GetSystemRole() string`

GetSystemRole returns the SystemRole field if non-nil, zero value otherwise.

### GetSystemRoleOk

`func (o *HostBinding) GetSystemRoleOk() (*string, bool)`

GetSystemRoleOk returns a tuple with the SystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRole

`func (o *HostBinding) SetSystemRole(v string)`

SetSystemRole sets SystemRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


