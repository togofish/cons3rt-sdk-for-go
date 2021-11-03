# InputDiskForTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityInMegabytes** | **int32** |  | 
**IsAdditionalDisk** | Pointer to **bool** |  | [optional] 
**IsBootDisk** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputDiskForTemplate

`func NewInputDiskForTemplate(capacityInMegabytes int32, ) *InputDiskForTemplate`

NewInputDiskForTemplate instantiates a new InputDiskForTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDiskForTemplateWithDefaults

`func NewInputDiskForTemplateWithDefaults() *InputDiskForTemplate`

NewInputDiskForTemplateWithDefaults instantiates a new InputDiskForTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityInMegabytes

`func (o *InputDiskForTemplate) GetCapacityInMegabytes() int32`

GetCapacityInMegabytes returns the CapacityInMegabytes field if non-nil, zero value otherwise.

### GetCapacityInMegabytesOk

`func (o *InputDiskForTemplate) GetCapacityInMegabytesOk() (*int32, bool)`

GetCapacityInMegabytesOk returns a tuple with the CapacityInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInMegabytes

`func (o *InputDiskForTemplate) SetCapacityInMegabytes(v int32)`

SetCapacityInMegabytes sets CapacityInMegabytes field to given value.


### GetIsAdditionalDisk

`func (o *InputDiskForTemplate) GetIsAdditionalDisk() bool`

GetIsAdditionalDisk returns the IsAdditionalDisk field if non-nil, zero value otherwise.

### GetIsAdditionalDiskOk

`func (o *InputDiskForTemplate) GetIsAdditionalDiskOk() (*bool, bool)`

GetIsAdditionalDiskOk returns a tuple with the IsAdditionalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdditionalDisk

`func (o *InputDiskForTemplate) SetIsAdditionalDisk(v bool)`

SetIsAdditionalDisk sets IsAdditionalDisk field to given value.

### HasIsAdditionalDisk

`func (o *InputDiskForTemplate) HasIsAdditionalDisk() bool`

HasIsAdditionalDisk returns a boolean if a field has been set.

### GetIsBootDisk

`func (o *InputDiskForTemplate) GetIsBootDisk() bool`

GetIsBootDisk returns the IsBootDisk field if non-nil, zero value otherwise.

### GetIsBootDiskOk

`func (o *InputDiskForTemplate) GetIsBootDiskOk() (*bool, bool)`

GetIsBootDiskOk returns a tuple with the IsBootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootDisk

`func (o *InputDiskForTemplate) SetIsBootDisk(v bool)`

SetIsBootDisk sets IsBootDisk field to given value.

### HasIsBootDisk

`func (o *InputDiskForTemplate) HasIsBootDisk() bool`

HasIsBootDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


