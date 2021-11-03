# CompositionLaunchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**TimeToLive** | Pointer to **int32** |  | [optional] 

## Methods

### NewCompositionLaunchOptions

`func NewCompositionLaunchOptions(username string, password string, ) *CompositionLaunchOptions`

NewCompositionLaunchOptions instantiates a new CompositionLaunchOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositionLaunchOptionsWithDefaults

`func NewCompositionLaunchOptionsWithDefaults() *CompositionLaunchOptions`

NewCompositionLaunchOptionsWithDefaults instantiates a new CompositionLaunchOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CompositionLaunchOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CompositionLaunchOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CompositionLaunchOptions) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CompositionLaunchOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CompositionLaunchOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CompositionLaunchOptions) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTimeToLive

`func (o *CompositionLaunchOptions) GetTimeToLive() int32`

GetTimeToLive returns the TimeToLive field if non-nil, zero value otherwise.

### GetTimeToLiveOk

`func (o *CompositionLaunchOptions) GetTimeToLiveOk() (*int32, bool)`

GetTimeToLiveOk returns a tuple with the TimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLive

`func (o *CompositionLaunchOptions) SetTimeToLive(v int32)`

SetTimeToLive sets TimeToLive field to given value.

### HasTimeToLive

`func (o *CompositionLaunchOptions) HasTimeToLive() bool`

HasTimeToLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


