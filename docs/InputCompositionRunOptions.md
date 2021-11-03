# InputCompositionRunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**VirtualizationRealmId** | **int32** |  | 
**Properties** | Pointer to [**[]InputProperty**](InputProperty.md) |  | [optional] 
**HostOptions** | Pointer to [**[]InputCompositionHostOption**](InputCompositionHostOption.md) |  | [optional] 
**WindowsDomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewInputCompositionRunOptions

`func NewInputCompositionRunOptions(virtualizationRealmId int32, ) *InputCompositionRunOptions`

NewInputCompositionRunOptions instantiates a new InputCompositionRunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCompositionRunOptionsWithDefaults

`func NewInputCompositionRunOptionsWithDefaults() *InputCompositionRunOptions`

NewInputCompositionRunOptionsWithDefaults instantiates a new InputCompositionRunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InputCompositionRunOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCompositionRunOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCompositionRunOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCompositionRunOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVirtualizationRealmId

`func (o *InputCompositionRunOptions) GetVirtualizationRealmId() int32`

GetVirtualizationRealmId returns the VirtualizationRealmId field if non-nil, zero value otherwise.

### GetVirtualizationRealmIdOk

`func (o *InputCompositionRunOptions) GetVirtualizationRealmIdOk() (*int32, bool)`

GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmId

`func (o *InputCompositionRunOptions) SetVirtualizationRealmId(v int32)`

SetVirtualizationRealmId sets VirtualizationRealmId field to given value.


### GetProperties

`func (o *InputCompositionRunOptions) GetProperties() []InputProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InputCompositionRunOptions) GetPropertiesOk() (*[]InputProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InputCompositionRunOptions) SetProperties(v []InputProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InputCompositionRunOptions) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetHostOptions

`func (o *InputCompositionRunOptions) GetHostOptions() []InputCompositionHostOption`

GetHostOptions returns the HostOptions field if non-nil, zero value otherwise.

### GetHostOptionsOk

`func (o *InputCompositionRunOptions) GetHostOptionsOk() (*[]InputCompositionHostOption, bool)`

GetHostOptionsOk returns a tuple with the HostOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOptions

`func (o *InputCompositionRunOptions) SetHostOptions(v []InputCompositionHostOption)`

SetHostOptions sets HostOptions field to given value.

### HasHostOptions

`func (o *InputCompositionRunOptions) HasHostOptions() bool`

HasHostOptions returns a boolean if a field has been set.

### GetWindowsDomainName

`func (o *InputCompositionRunOptions) GetWindowsDomainName() string`

GetWindowsDomainName returns the WindowsDomainName field if non-nil, zero value otherwise.

### GetWindowsDomainNameOk

`func (o *InputCompositionRunOptions) GetWindowsDomainNameOk() (*string, bool)`

GetWindowsDomainNameOk returns a tuple with the WindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomainName

`func (o *InputCompositionRunOptions) SetWindowsDomainName(v string)`

SetWindowsDomainName sets WindowsDomainName field to given value.

### HasWindowsDomainName

`func (o *InputCompositionRunOptions) HasWindowsDomainName() bool`

HasWindowsDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


