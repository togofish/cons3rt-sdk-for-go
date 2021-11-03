# InputPhysicalHostFullAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PhysicalMachine** | Pointer to [**InputPhysicalMachine**](InputPhysicalMachine.md) |  | [optional] 
**Components** | Pointer to [**[]InputAbstractComponent**](InputAbstractComponent.md) |  | [optional] 

## Methods

### NewInputPhysicalHostFullAllOf

`func NewInputPhysicalHostFullAllOf() *InputPhysicalHostFullAllOf`

NewInputPhysicalHostFullAllOf instantiates a new InputPhysicalHostFullAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputPhysicalHostFullAllOfWithDefaults

`func NewInputPhysicalHostFullAllOfWithDefaults() *InputPhysicalHostFullAllOf`

NewInputPhysicalHostFullAllOfWithDefaults instantiates a new InputPhysicalHostFullAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputPhysicalHostFullAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputPhysicalHostFullAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputPhysicalHostFullAllOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InputPhysicalHostFullAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InputPhysicalHostFullAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputPhysicalHostFullAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputPhysicalHostFullAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputPhysicalHostFullAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhysicalMachine

`func (o *InputPhysicalHostFullAllOf) GetPhysicalMachine() InputPhysicalMachine`

GetPhysicalMachine returns the PhysicalMachine field if non-nil, zero value otherwise.

### GetPhysicalMachineOk

`func (o *InputPhysicalHostFullAllOf) GetPhysicalMachineOk() (*InputPhysicalMachine, bool)`

GetPhysicalMachineOk returns a tuple with the PhysicalMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachine

`func (o *InputPhysicalHostFullAllOf) SetPhysicalMachine(v InputPhysicalMachine)`

SetPhysicalMachine sets PhysicalMachine field to given value.

### HasPhysicalMachine

`func (o *InputPhysicalHostFullAllOf) HasPhysicalMachine() bool`

HasPhysicalMachine returns a boolean if a field has been set.

### GetComponents

`func (o *InputPhysicalHostFullAllOf) GetComponents() []InputAbstractComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *InputPhysicalHostFullAllOf) GetComponentsOk() (*[]InputAbstractComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *InputPhysicalHostFullAllOf) SetComponents(v []InputAbstractComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *InputPhysicalHostFullAllOf) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


