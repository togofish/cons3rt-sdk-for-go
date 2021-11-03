# ContainerMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 
**Source** | **string** |  | 
**Target** | **string** |  | 
**Options** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerMount

`func NewContainerMount(type_ string, source string, target string, ) *ContainerMount`

NewContainerMount instantiates a new ContainerMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerMountWithDefaults

`func NewContainerMountWithDefaults() *ContainerMount`

NewContainerMountWithDefaults instantiates a new ContainerMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerMount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerMount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerMount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerMount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ContainerMount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerMount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerMount) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *ContainerMount) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContainerMount) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContainerMount) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *ContainerMount) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContainerMount) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContainerMount) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetOptions

`func (o *ContainerMount) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ContainerMount) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ContainerMount) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ContainerMount) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


