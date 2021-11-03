# InputPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string][]string** |  | [optional] 
**BodyAsString** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to [**InputPartMediaType**](InputPartMediaType.md) |  | [optional] 
**ContentTypeFromMessage** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputPart

`func NewInputPart() *InputPart`

NewInputPart instantiates a new InputPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputPartWithDefaults

`func NewInputPartWithDefaults() *InputPart`

NewInputPartWithDefaults instantiates a new InputPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *InputPart) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InputPart) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InputPart) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InputPart) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBodyAsString

`func (o *InputPart) GetBodyAsString() string`

GetBodyAsString returns the BodyAsString field if non-nil, zero value otherwise.

### GetBodyAsStringOk

`func (o *InputPart) GetBodyAsStringOk() (*string, bool)`

GetBodyAsStringOk returns a tuple with the BodyAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyAsString

`func (o *InputPart) SetBodyAsString(v string)`

SetBodyAsString sets BodyAsString field to given value.

### HasBodyAsString

`func (o *InputPart) HasBodyAsString() bool`

HasBodyAsString returns a boolean if a field has been set.

### GetMediaType

`func (o *InputPart) GetMediaType() InputPartMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *InputPart) GetMediaTypeOk() (*InputPartMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *InputPart) SetMediaType(v InputPartMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *InputPart) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetContentTypeFromMessage

`func (o *InputPart) GetContentTypeFromMessage() bool`

GetContentTypeFromMessage returns the ContentTypeFromMessage field if non-nil, zero value otherwise.

### GetContentTypeFromMessageOk

`func (o *InputPart) GetContentTypeFromMessageOk() (*bool, bool)`

GetContentTypeFromMessageOk returns a tuple with the ContentTypeFromMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeFromMessage

`func (o *InputPart) SetContentTypeFromMessage(v bool)`

SetContentTypeFromMessage sets ContentTypeFromMessage field to given value.

### HasContentTypeFromMessage

`func (o *InputPart) HasContentTypeFromMessage() bool`

HasContentTypeFromMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


