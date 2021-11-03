# FullAppliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]MinimalAbstractComponent**](MinimalAbstractComponent.md) |  | [optional] 
**TemplateProfile** | Pointer to [**MinimalTemplateProfile**](MinimalTemplateProfile.md) |  | [optional] 

## Methods

### NewFullAppliance

`func NewFullAppliance() *FullAppliance`

NewFullAppliance instantiates a new FullAppliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullApplianceWithDefaults

`func NewFullApplianceWithDefaults() *FullAppliance`

NewFullApplianceWithDefaults instantiates a new FullAppliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *FullAppliance) GetComponents() []MinimalAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FullAppliance) GetComponentsOk() (*[]MinimalAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FullAppliance) SetComponents(v []MinimalAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FullAppliance) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *FullAppliance) GetTemplateProfile() MinimalTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *FullAppliance) GetTemplateProfileOk() (*MinimalTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *FullAppliance) SetTemplateProfile(v MinimalTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *FullAppliance) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


