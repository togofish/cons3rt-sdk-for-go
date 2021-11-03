# InputVirtualHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]InputAbstractComponent**](InputAbstractComponent.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemplateProfile** | Pointer to [**InputTemplateProfile**](InputTemplateProfile.md) |  | [optional] 

## Methods

### NewInputVirtualHostAllOf

`func NewInputVirtualHostAllOf() *InputVirtualHostAllOf`

NewInputVirtualHostAllOf instantiates a new InputVirtualHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVirtualHostAllOfWithDefaults

`func NewInputVirtualHostAllOfWithDefaults() *InputVirtualHostAllOf`

NewInputVirtualHostAllOfWithDefaults instantiates a new InputVirtualHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *InputVirtualHostAllOf) GetComponents() []InputAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *InputVirtualHostAllOf) GetComponentsOk() (*[]InputAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *InputVirtualHostAllOf) SetComponents(v []InputAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *InputVirtualHostAllOf) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetName

`func (o *InputVirtualHostAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputVirtualHostAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputVirtualHostAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputVirtualHostAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *InputVirtualHostAllOf) GetTemplateProfile() InputTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *InputVirtualHostAllOf) GetTemplateProfileOk() (*InputTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *InputVirtualHostAllOf) SetTemplateProfile(v InputTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *InputVirtualHostAllOf) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


