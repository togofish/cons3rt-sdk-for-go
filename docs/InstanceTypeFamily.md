# InstanceTypeFamily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** |  | [optional] 
**InstanceTypes** | Pointer to [**[]InstanceType**](InstanceType.md) |  | [optional] 

## Methods

### NewInstanceTypeFamily

`func NewInstanceTypeFamily() *InstanceTypeFamily`

NewInstanceTypeFamily instantiates a new InstanceTypeFamily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeFamilyWithDefaults

`func NewInstanceTypeFamilyWithDefaults() *InstanceTypeFamily`

NewInstanceTypeFamilyWithDefaults instantiates a new InstanceTypeFamily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *InstanceTypeFamily) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *InstanceTypeFamily) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *InstanceTypeFamily) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *InstanceTypeFamily) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *InstanceTypeFamily) GetInstanceTypes() []InstanceType`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *InstanceTypeFamily) GetInstanceTypesOk() (*[]InstanceType, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *InstanceTypeFamily) SetInstanceTypes(v []InstanceType)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *InstanceTypeFamily) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


