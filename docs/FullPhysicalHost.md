# FullPhysicalHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]MinimalAbstractComponent**](MinimalAbstractComponent.md) |  | [optional] 
**PhysicalMachine** | Pointer to [**GeneralPhysicalMachine**](GeneralPhysicalMachine.md) |  | [optional] 

## Methods

### NewFullPhysicalHost

`func NewFullPhysicalHost() *FullPhysicalHost`

NewFullPhysicalHost instantiates a new FullPhysicalHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullPhysicalHostWithDefaults

`func NewFullPhysicalHostWithDefaults() *FullPhysicalHost`

NewFullPhysicalHostWithDefaults instantiates a new FullPhysicalHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *FullPhysicalHost) GetComponents() []MinimalAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FullPhysicalHost) GetComponentsOk() (*[]MinimalAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FullPhysicalHost) SetComponents(v []MinimalAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FullPhysicalHost) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetPhysicalMachine

`func (o *FullPhysicalHost) GetPhysicalMachine() GeneralPhysicalMachine`

GetPhysicalMachine returns the PhysicalMachine field if non-nil, zero value otherwise.

### GetPhysicalMachineOk

`func (o *FullPhysicalHost) GetPhysicalMachineOk() (*GeneralPhysicalMachine, bool)`

GetPhysicalMachineOk returns a tuple with the PhysicalMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachine

`func (o *FullPhysicalHost) SetPhysicalMachine(v GeneralPhysicalMachine)`

SetPhysicalMachine sets PhysicalMachine field to given value.

### HasPhysicalMachine

`func (o *FullPhysicalHost) HasPhysicalMachine() bool`

HasPhysicalMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


