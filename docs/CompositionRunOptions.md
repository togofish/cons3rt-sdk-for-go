# CompositionRunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**HostOptions** | Pointer to [**[]CompositionHostOption**](CompositionHostOption.md) |  | [optional] 
**VirtualizationRealmId** | **int32** |  | 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**WindowsDomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewCompositionRunOptions

`func NewCompositionRunOptions(virtualizationRealmId int32, ) *CompositionRunOptions`

NewCompositionRunOptions instantiates a new CompositionRunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositionRunOptionsWithDefaults

`func NewCompositionRunOptionsWithDefaults() *CompositionRunOptions`

NewCompositionRunOptionsWithDefaults instantiates a new CompositionRunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CompositionRunOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompositionRunOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompositionRunOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompositionRunOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostOptions

`func (o *CompositionRunOptions) GetHostOptions() []CompositionHostOption`

GetHostOptions returns the HostOptions field if non-nil, zero value otherwise.

### GetHostOptionsOk

`func (o *CompositionRunOptions) GetHostOptionsOk() (*[]CompositionHostOption, bool)`

GetHostOptionsOk returns a tuple with the HostOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOptions

`func (o *CompositionRunOptions) SetHostOptions(v []CompositionHostOption)`

SetHostOptions sets HostOptions field to given value.

### HasHostOptions

`func (o *CompositionRunOptions) HasHostOptions() bool`

HasHostOptions returns a boolean if a field has been set.

### GetVirtualizationRealmId

`func (o *CompositionRunOptions) GetVirtualizationRealmId() int32`

GetVirtualizationRealmId returns the VirtualizationRealmId field if non-nil, zero value otherwise.

### GetVirtualizationRealmIdOk

`func (o *CompositionRunOptions) GetVirtualizationRealmIdOk() (*int32, bool)`

GetVirtualizationRealmIdOk returns a tuple with the VirtualizationRealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationRealmId

`func (o *CompositionRunOptions) SetVirtualizationRealmId(v int32)`

SetVirtualizationRealmId sets VirtualizationRealmId field to given value.


### GetProperties

`func (o *CompositionRunOptions) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CompositionRunOptions) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CompositionRunOptions) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CompositionRunOptions) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetWindowsDomainName

`func (o *CompositionRunOptions) GetWindowsDomainName() string`

GetWindowsDomainName returns the WindowsDomainName field if non-nil, zero value otherwise.

### GetWindowsDomainNameOk

`func (o *CompositionRunOptions) GetWindowsDomainNameOk() (*string, bool)`

GetWindowsDomainNameOk returns a tuple with the WindowsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomainName

`func (o *CompositionRunOptions) SetWindowsDomainName(v string)`

SetWindowsDomainName sets WindowsDomainName field to given value.

### HasWindowsDomainName

`func (o *CompositionRunOptions) HasWindowsDomainName() bool`

HasWindowsDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


