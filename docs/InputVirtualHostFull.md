# InputVirtualHostFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TemplateProfile** | Pointer to [**InputTemplateProfile**](InputTemplateProfile.md) |  | [optional] 
**Components** | Pointer to [**[]InputAbstractComponent**](InputAbstractComponent.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewInputVirtualHostFull

`func NewInputVirtualHostFull() *InputVirtualHostFull`

NewInputVirtualHostFull instantiates a new InputVirtualHostFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVirtualHostFullWithDefaults

`func NewInputVirtualHostFullWithDefaults() *InputVirtualHostFull`

NewInputVirtualHostFullWithDefaults instantiates a new InputVirtualHostFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputVirtualHostFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputVirtualHostFull) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputVirtualHostFull) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputVirtualHostFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *InputVirtualHostFull) GetTemplateProfile() InputTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *InputVirtualHostFull) GetTemplateProfileOk() (*InputTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *InputVirtualHostFull) SetTemplateProfile(v InputTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *InputVirtualHostFull) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.

### GetComponents

`func (o *InputVirtualHostFull) GetComponents() []InputAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *InputVirtualHostFull) GetComponentsOk() (*[]InputAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *InputVirtualHostFull) SetComponents(v []InputAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *InputVirtualHostFull) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetName

`func (o *InputVirtualHostFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputVirtualHostFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputVirtualHostFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputVirtualHostFull) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


