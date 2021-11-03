# InputAppliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TemplateProfile** | Pointer to [**InputTemplateProfile**](InputTemplateProfile.md) |  | [optional] 

## Methods

### NewInputAppliance

`func NewInputAppliance() *InputAppliance`

NewInputAppliance instantiates a new InputAppliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputApplianceWithDefaults

`func NewInputApplianceWithDefaults() *InputAppliance`

NewInputApplianceWithDefaults instantiates a new InputAppliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputAppliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputAppliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputAppliance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputAppliance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateProfile

`func (o *InputAppliance) GetTemplateProfile() InputTemplateProfile`

GetTemplateProfile returns the TemplateProfile field if non-nil, zero value otherwise.

### GetTemplateProfileOk

`func (o *InputAppliance) GetTemplateProfileOk() (*InputTemplateProfile, bool)`

GetTemplateProfileOk returns a tuple with the TemplateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProfile

`func (o *InputAppliance) SetTemplateProfile(v InputTemplateProfile)`

SetTemplateProfile sets TemplateProfile field to given value.

### HasTemplateProfile

`func (o *InputAppliance) HasTemplateProfile() bool`

HasTemplateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


