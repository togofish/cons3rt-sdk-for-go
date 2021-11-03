# InputVirtualHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]InputAbstractComponent**](InputAbstractComponent.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemplateProfile** | Pointer to [**InputTemplateProfile**](InputTemplateProfile.md) |  | [optional] 

## Methods

### NewInputVirtualHost

`func NewInputVirtualHost() *InputVirtualHost`

NewInputVirtualHost instantiates a new InputVirtualHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVirtualHostWithDefaults

`func NewInputVirtualHostWithDefaults() *InputVirtualHost`

NewInputVirtualHostWithDefaults instantiates a new InputVirtualHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *InputVirtualHost) GetComponents() []InputAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *InputVirtualHost) GetComponentsOk() (*[]InputAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *InputVirtualHost) SetComponents(v []InputAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *InputVirtualHost) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetName

`func (o *InputVirtualHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputVirtualHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputVirtualHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputVirtualHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *InputVirtualHost) GetTemplateProfile() InputTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *InputVirtualHost) GetTemplateProfileOk() (*InputTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *InputVirtualHost) SetTemplateProfile(v InputTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *InputVirtualHost) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


