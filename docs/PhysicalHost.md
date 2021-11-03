# PhysicalHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhysicalMachine** | Pointer to [**PhysicalMachine**](PhysicalMachine.md) |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 

## Methods

### NewPhysicalHost

`func NewPhysicalHost() *PhysicalHost`

NewPhysicalHost instantiates a new PhysicalHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalHostWithDefaults

`func NewPhysicalHostWithDefaults() *PhysicalHost`

NewPhysicalHostWithDefaults instantiates a new PhysicalHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhysicalMachine

`func (o *PhysicalHost) GetPhysicalMachine() PhysicalMachine`

GetPhysicalMachine returns the PhysicalMachine field if non-nil, zero value otherwise.

### GetPhysicalMachineOk

`func (o *PhysicalHost) GetPhysicalMachineOk() (*PhysicalMachine, bool)`

GetPhysicalMachineOk returns a tuple with the PhysicalMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMachine

`func (o *PhysicalHost) SetPhysicalMachine(v PhysicalMachine)`

SetPhysicalMachine sets PhysicalMachine field to given value.

### HasPhysicalMachine

`func (o *PhysicalHost) HasPhysicalMachine() bool`

HasPhysicalMachine returns a boolean if a field has been set.

### GetIpAddress

`func (o *PhysicalHost) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PhysicalHost) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PhysicalHost) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PhysicalHost) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *PhysicalHost) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PhysicalHost) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PhysicalHost) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PhysicalHost) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetHostname

`func (o *PhysicalHost) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PhysicalHost) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PhysicalHost) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PhysicalHost) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


