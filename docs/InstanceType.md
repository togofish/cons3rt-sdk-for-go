# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceType

`func NewInstanceType() *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *InstanceType) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *InstanceType) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *InstanceType) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *InstanceType) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetType

`func (o *InstanceType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


