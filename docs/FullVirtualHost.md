# FullVirtualHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]MinimalAbstractComponent**](MinimalAbstractComponent.md) |  | [optional] 
**TemplateProfile** | Pointer to [**MinimalTemplateProfile**](MinimalTemplateProfile.md) |  | [optional] 

## Methods

### NewFullVirtualHost

`func NewFullVirtualHost() *FullVirtualHost`

NewFullVirtualHost instantiates a new FullVirtualHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullVirtualHostWithDefaults

`func NewFullVirtualHostWithDefaults() *FullVirtualHost`

NewFullVirtualHostWithDefaults instantiates a new FullVirtualHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *FullVirtualHost) GetComponents() []MinimalAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FullVirtualHost) GetComponentsOk() (*[]MinimalAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FullVirtualHost) SetComponents(v []MinimalAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FullVirtualHost) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *FullVirtualHost) GetTemplateProfile() MinimalTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *FullVirtualHost) GetTemplateProfileOk() (*MinimalTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *FullVirtualHost) SetTemplateProfile(v MinimalTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *FullVirtualHost) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


