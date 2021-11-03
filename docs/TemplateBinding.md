# TemplateBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**MinimalCons3rtTemplateData**](MinimalCons3rtTemplateData.md) |  | [optional] 
**BestMatches** | Pointer to [**[]InstanceType**](InstanceType.md) |  | [optional] 
**Matches** | Pointer to [**[]InstanceTypeFamily**](InstanceTypeFamily.md) |  | [optional] 

## Methods

### NewTemplateBinding

`func NewTemplateBinding() *TemplateBinding`

NewTemplateBinding instantiates a new TemplateBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateBindingWithDefaults

`func NewTemplateBindingWithDefaults() *TemplateBinding`

NewTemplateBindingWithDefaults instantiates a new TemplateBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *TemplateBinding) GetTemplate() MinimalCons3rtTemplateData`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *TemplateBinding) GetTemplateOk() (*MinimalCons3rtTemplateData, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *TemplateBinding) SetTemplate(v MinimalCons3rtTemplateData)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *TemplateBinding) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetBestMatches

`func (o *TemplateBinding) GetBestMatches() []InstanceType`

GetBestMatches returns the BestMatches field if non-nil, zero value otherwise.

### GetBestMatchesOk

`func (o *TemplateBinding) GetBestMatchesOk() (*[]InstanceType, bool)`

GetBestMatchesOk returns a tuple with the BestMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestMatches

`func (o *TemplateBinding) SetBestMatches(v []InstanceType)`

SetBestMatches sets BestMatches field to given value.

### HasBestMatches

`func (o *TemplateBinding) HasBestMatches() bool`

HasBestMatches returns a boolean if a field has been set.

### GetMatches

`func (o *TemplateBinding) GetMatches() []InstanceTypeFamily`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *TemplateBinding) GetMatchesOk() (*[]InstanceTypeFamily, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *TemplateBinding) SetMatches(v []InstanceTypeFamily)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *TemplateBinding) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


