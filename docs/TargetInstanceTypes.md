# TargetInstanceTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BestMatches** | Pointer to [**[]InstanceType**](InstanceType.md) |  | [optional] 
**Matches** | Pointer to [**[]InstanceTypeFamily**](InstanceTypeFamily.md) |  | [optional] 

## Methods

### NewTargetInstanceTypes

`func NewTargetInstanceTypes() *TargetInstanceTypes`

NewTargetInstanceTypes instantiates a new TargetInstanceTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetInstanceTypesWithDefaults

`func NewTargetInstanceTypesWithDefaults() *TargetInstanceTypes`

NewTargetInstanceTypesWithDefaults instantiates a new TargetInstanceTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBestMatches

`func (o *TargetInstanceTypes) GetBestMatches() []InstanceType`

GetBestMatches returns the BestMatches field if non-nil, zero value otherwise.

### GetBestMatchesOk

`func (o *TargetInstanceTypes) GetBestMatchesOk() (*[]InstanceType, bool)`

GetBestMatchesOk returns a tuple with the BestMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestMatches

`func (o *TargetInstanceTypes) SetBestMatches(v []InstanceType)`

SetBestMatches sets BestMatches field to given value.

### HasBestMatches

`func (o *TargetInstanceTypes) HasBestMatches() bool`

HasBestMatches returns a boolean if a field has been set.

### GetMatches

`func (o *TargetInstanceTypes) GetMatches() []InstanceTypeFamily`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *TargetInstanceTypes) GetMatchesOk() (*[]InstanceTypeFamily, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *TargetInstanceTypes) SetMatches(v []InstanceTypeFamily)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *TargetInstanceTypes) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


