# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  | 
**Name** | **string** |  | 
**NatInstance** | [**NatInstance**](NatInstance.md) |  | 
**SecurityGroup** | [**SecurityGroup**](SecurityGroup.md) |  | 

## Methods

### NewSubnet

`func NewSubnet(cidr string, identifier string, name string, natInstance NatInstance, securityGroup SecurityGroup, ) *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *Subnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Subnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Subnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetId

`func (o *Subnet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Subnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *Subnet) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Subnet) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Subnet) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *Subnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnet) SetName(v string)`

SetName sets Name field to given value.


### GetNatInstance

`func (o *Subnet) GetNatInstance() NatInstance`

GetNatInstance returns the NatInstance field if non-nil, zero value otherwise.

### GetNatInstanceOk

`func (o *Subnet) GetNatInstanceOk() (*NatInstance, bool)`

GetNatInstanceOk returns a tuple with the NatInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstance

`func (o *Subnet) SetNatInstance(v NatInstance)`

SetNatInstance sets NatInstance field to given value.


### GetSecurityGroup

`func (o *Subnet) GetSecurityGroup() SecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *Subnet) GetSecurityGroupOk() (*SecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *Subnet) SetSecurityGroup(v SecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


