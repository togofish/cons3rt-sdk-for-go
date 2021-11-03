# MultipartFormDataInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormData** | Pointer to [**map[string]InputPart**](InputPart.md) |  | [optional] 
**FormDataMap** | Pointer to [**map[string][]InputPart**](array.md) |  | [optional] 
**Preamble** | Pointer to **string** |  | [optional] 
**Parts** | Pointer to [**[]InputPart**](InputPart.md) |  | [optional] 

## Methods

### NewMultipartFormDataInput

`func NewMultipartFormDataInput() *MultipartFormDataInput`

NewMultipartFormDataInput instantiates a new MultipartFormDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipartFormDataInputWithDefaults

`func NewMultipartFormDataInputWithDefaults() *MultipartFormDataInput`

NewMultipartFormDataInputWithDefaults instantiates a new MultipartFormDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormData

`func (o *MultipartFormDataInput) GetFormData() map[string]InputPart`

GetFormData returns the FormData field if non-nil, zero value otherwise.

### GetFormDataOk

`func (o *MultipartFormDataInput) GetFormDataOk() (*map[string]InputPart, bool)`

GetFormDataOk returns a tuple with the FormData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormData

`func (o *MultipartFormDataInput) SetFormData(v map[string]InputPart)`

SetFormData sets FormData field to given value.

### HasFormData

`func (o *MultipartFormDataInput) HasFormData() bool`

HasFormData returns a boolean if a field has been set.

### GetFormDataMap

`func (o *MultipartFormDataInput) GetFormDataMap() map[string][]InputPart`

GetFormDataMap returns the FormDataMap field if non-nil, zero value otherwise.

### GetFormDataMapOk

`func (o *MultipartFormDataInput) GetFormDataMapOk() (*map[string][]InputPart, bool)`

GetFormDataMapOk returns a tuple with the FormDataMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormDataMap

`func (o *MultipartFormDataInput) SetFormDataMap(v map[string][]InputPart)`

SetFormDataMap sets FormDataMap field to given value.

### HasFormDataMap

`func (o *MultipartFormDataInput) HasFormDataMap() bool`

HasFormDataMap returns a boolean if a field has been set.

### GetPreamble

`func (o *MultipartFormDataInput) GetPreamble() string`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *MultipartFormDataInput) GetPreambleOk() (*string, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *MultipartFormDataInput) SetPreamble(v string)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *MultipartFormDataInput) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.

### GetParts

`func (o *MultipartFormDataInput) GetParts() []InputPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MultipartFormDataInput) GetPartsOk() (*[]InputPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MultipartFormDataInput) SetParts(v []InputPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MultipartFormDataInput) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


