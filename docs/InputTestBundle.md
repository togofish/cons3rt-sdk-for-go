# InputTestBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestAsset** | Pointer to [**InputTestAsset**](InputTestAsset.md) |  | [optional] 
**TestToolPoolType** | **string** |  | 

## Methods

### NewInputTestBundle

`func NewInputTestBundle(testToolPoolType string, ) *InputTestBundle`

NewInputTestBundle instantiates a new InputTestBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTestBundleWithDefaults

`func NewInputTestBundleWithDefaults() *InputTestBundle`

NewInputTestBundleWithDefaults instantiates a new InputTestBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestAsset

`func (o *InputTestBundle) GetTestAsset() InputTestAsset`

GetTestAsset returns the TestAsset field if non-nil, zero value otherwise.

### GetTestAssetOk

`func (o *InputTestBundle) GetTestAssetOk() (*InputTestAsset, bool)`

GetTestAssetOk returns a tuple with the TestAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAsset

`func (o *InputTestBundle) SetTestAsset(v InputTestAsset)`

SetTestAsset sets TestAsset field to given value.

### HasTestAsset

`func (o *InputTestBundle) HasTestAsset() bool`

HasTestAsset returns a boolean if a field has been set.

### GetTestToolPoolType

`func (o *InputTestBundle) GetTestToolPoolType() string`

GetTestToolPoolType returns the TestToolPoolType field if non-nil, zero value otherwise.

### GetTestToolPoolTypeOk

`func (o *InputTestBundle) GetTestToolPoolTypeOk() (*string, bool)`

GetTestToolPoolTypeOk returns a tuple with the TestToolPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestToolPoolType

`func (o *InputTestBundle) SetTestToolPoolType(v string)`

SetTestToolPoolType sets TestToolPoolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


