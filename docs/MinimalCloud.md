# MinimalCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 

## Methods

### NewMinimalCloud

`func NewMinimalCloud(name string, ) *MinimalCloud`

NewMinimalCloud instantiates a new MinimalCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalCloudWithDefaults

`func NewMinimalCloudWithDefaults() *MinimalCloud`

NewMinimalCloudWithDefaults instantiates a new MinimalCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalCloud) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalCloud) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalCloud) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalCloud) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MinimalCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalCloud) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *MinimalCloud) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MinimalCloud) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MinimalCloud) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MinimalCloud) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCloudType

`func (o *MinimalCloud) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *MinimalCloud) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *MinimalCloud) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *MinimalCloud) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


