# InputTemplateSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**AllowGpuUsage** | Pointer to **bool** |  | [optional] 
**MaxNumCpus** | Pointer to **int32** |  | [optional] 
**MaxRamInMegabytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputTemplateSubscription

`func NewInputTemplateSubscription() *InputTemplateSubscription`

NewInputTemplateSubscription instantiates a new InputTemplateSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTemplateSubscriptionWithDefaults

`func NewInputTemplateSubscriptionWithDefaults() *InputTemplateSubscription`

NewInputTemplateSubscriptionWithDefaults instantiates a new InputTemplateSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *InputTemplateSubscription) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InputTemplateSubscription) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InputTemplateSubscription) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InputTemplateSubscription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAllowGpuUsage

`func (o *InputTemplateSubscription) GetAllowGpuUsage() bool`

GetAllowGpuUsage returns the AllowGpuUsage field if non-nil, zero value otherwise.

### GetAllowGpuUsageOk

`func (o *InputTemplateSubscription) GetAllowGpuUsageOk() (*bool, bool)`

GetAllowGpuUsageOk returns a tuple with the AllowGpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGpuUsage

`func (o *InputTemplateSubscription) SetAllowGpuUsage(v bool)`

SetAllowGpuUsage sets AllowGpuUsage field to given value.

### HasAllowGpuUsage

`func (o *InputTemplateSubscription) HasAllowGpuUsage() bool`

HasAllowGpuUsage returns a boolean if a field has been set.

### GetMaxNumCpus

`func (o *InputTemplateSubscription) GetMaxNumCpus() int32`

GetMaxNumCpus returns the MaxNumCpus field if non-nil, zero value otherwise.

### GetMaxNumCpusOk

`func (o *InputTemplateSubscription) GetMaxNumCpusOk() (*int32, bool)`

GetMaxNumCpusOk returns a tuple with the MaxNumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumCpus

`func (o *InputTemplateSubscription) SetMaxNumCpus(v int32)`

SetMaxNumCpus sets MaxNumCpus field to given value.

### HasMaxNumCpus

`func (o *InputTemplateSubscription) HasMaxNumCpus() bool`

HasMaxNumCpus returns a boolean if a field has been set.

### GetMaxRamInMegabytes

`func (o *InputTemplateSubscription) GetMaxRamInMegabytes() int32`

GetMaxRamInMegabytes returns the MaxRamInMegabytes field if non-nil, zero value otherwise.

### GetMaxRamInMegabytesOk

`func (o *InputTemplateSubscription) GetMaxRamInMegabytesOk() (*int32, bool)`

GetMaxRamInMegabytesOk returns a tuple with the MaxRamInMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamInMegabytes

`func (o *InputTemplateSubscription) SetMaxRamInMegabytes(v int32)`

SetMaxRamInMegabytes sets MaxRamInMegabytes field to given value.

### HasMaxRamInMegabytes

`func (o *InputTemplateSubscription) HasMaxRamInMegabytes() bool`

HasMaxRamInMegabytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


