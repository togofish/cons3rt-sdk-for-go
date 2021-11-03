# AbstractCloudResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatInstanceTypes** | Pointer to [**[]InstanceType**](InstanceType.md) |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewAbstractCloudResources

`func NewAbstractCloudResources(subtype string, ) *AbstractCloudResources`

NewAbstractCloudResources instantiates a new AbstractCloudResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractCloudResourcesWithDefaults

`func NewAbstractCloudResourcesWithDefaults() *AbstractCloudResources`

NewAbstractCloudResourcesWithDefaults instantiates a new AbstractCloudResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatInstanceTypes

`func (o *AbstractCloudResources) GetNatInstanceTypes() []InstanceType`

GetNatInstanceTypes returns the NatInstanceTypes field if non-nil, zero value otherwise.

### GetNatInstanceTypesOk

`func (o *AbstractCloudResources) GetNatInstanceTypesOk() (*[]InstanceType, bool)`

GetNatInstanceTypesOk returns a tuple with the NatInstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstanceTypes

`func (o *AbstractCloudResources) SetNatInstanceTypes(v []InstanceType)`

SetNatInstanceTypes sets NatInstanceTypes field to given value.

### HasNatInstanceTypes

`func (o *AbstractCloudResources) HasNatInstanceTypes() bool`

HasNatInstanceTypes returns a boolean if a field has been set.

### GetSubtype

`func (o *AbstractCloudResources) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AbstractCloudResources) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AbstractCloudResources) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


