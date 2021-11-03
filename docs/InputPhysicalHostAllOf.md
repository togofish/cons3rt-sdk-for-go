# InputPhysicalHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PhysicalMachine** | Pointer to [**InputPhysicalMachine**](InputPhysicalMachine.md) |  | [optional] 
**Components** | Pointer to [**[]InputAbstractComponent**](InputAbstractComponent.md) |  | [optional] 

## Methods

### NewInputPhysicalHostAllOf

`func NewInputPhysicalHostAllOf() *InputPhysicalHostAllOf`

NewInputPhysicalHostAllOf instantiates a new InputPhysicalHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputPhysicalHostAllOfWithDefaults

`func NewInputPhysicalHostAllOfWithDefaults() *InputPhysicalHostAllOf`

NewInputPhysicalHostAllOfWithDefaults instantiates a new InputPhysicalHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputPhysicalHostAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputPhysicalHostAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputPhysicalHostAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputPhysicalHostAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhysicalMachine

`func (o *InputPhysicalHostAllOf) GetPhysicalMachine() InputPhysicalMachine`

GetPhysicalMachine returns the PhysicalMachine field if non-nil, zero value otherwise.

### GetPhysicalMachineOk

`func (o *InputPhysicalHostAllOf) GetPhysicalMachineOk() (*InputPhysicalMachine, bool)`

GetPhysicalMachineOk returns a tuple with the PhysicalMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachine

`func (o *InputPhysicalHostAllOf) SetPhysicalMachine(v InputPhysicalMachine)`

SetPhysicalMachine sets PhysicalMachine field to given value.

### HasPhysicalMachine

`func (o *InputPhysicalHostAllOf) HasPhysicalMachine() bool`

HasPhysicalMachine returns a boolean if a field has been set.

### GetComponents

`func (o *InputPhysicalHostAllOf) GetComponents() []InputAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *InputPhysicalHostAllOf) GetComponentsOk() (*[]InputAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *InputPhysicalHostAllOf) SetComponents(v []InputAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *InputPhysicalHostAllOf) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


