# VirtualizationRealmClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**AcceptAllCerts** | Pointer to **bool** |  | [optional] 
**AcceptSelfSignedCerts** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewVirtualizationRealmClient

`func NewVirtualizationRealmClient(username string, password string, subtype string, ) *VirtualizationRealmClient`

NewVirtualizationRealmClient instantiates a new VirtualizationRealmClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationRealmClientWithDefaults

`func NewVirtualizationRealmClientWithDefaults() *VirtualizationRealmClient`

NewVirtualizationRealmClientWithDefaults instantiates a new VirtualizationRealmClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *VirtualizationRealmClient) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VirtualizationRealmClient) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VirtualizationRealmClient) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VirtualizationRealmClient) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VirtualizationRealmClient) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VirtualizationRealmClient) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAcceptAllCerts

`func (o *VirtualizationRealmClient) GetAcceptAllCerts() bool`

GetAcceptAllCerts returns the AcceptAllCerts field if non-nil, zero value otherwise.

### GetAcceptAllCertsOk

`func (o *VirtualizationRealmClient) GetAcceptAllCertsOk() (*bool, bool)`

GetAcceptAllCertsOk returns a tuple with the AcceptAllCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAllCerts

`func (o *VirtualizationRealmClient) SetAcceptAllCerts(v bool)`

SetAcceptAllCerts sets AcceptAllCerts field to given value.

### HasAcceptAllCerts

`func (o *VirtualizationRealmClient) HasAcceptAllCerts() bool`

HasAcceptAllCerts returns a boolean if a field has been set.

### GetAcceptSelfSignedCerts

`func (o *VirtualizationRealmClient) GetAcceptSelfSignedCerts() bool`

GetAcceptSelfSignedCerts returns the AcceptSelfSignedCerts field if non-nil, zero value otherwise.

### GetAcceptSelfSignedCertsOk

`func (o *VirtualizationRealmClient) GetAcceptSelfSignedCertsOk() (*bool, bool)`

GetAcceptSelfSignedCertsOk returns a tuple with the AcceptSelfSignedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptSelfSignedCerts

`func (o *VirtualizationRealmClient) SetAcceptSelfSignedCerts(v bool)`

SetAcceptSelfSignedCerts sets AcceptSelfSignedCerts field to given value.

### HasAcceptSelfSignedCerts

`func (o *VirtualizationRealmClient) HasAcceptSelfSignedCerts() bool`

HasAcceptSelfSignedCerts returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationRealmClient) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationRealmClient) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationRealmClient) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationRealmClient) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *VirtualizationRealmClient) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualizationRealmClient) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualizationRealmClient) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VirtualizationRealmClient) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *VirtualizationRealmClient) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VirtualizationRealmClient) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VirtualizationRealmClient) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VirtualizationRealmClient) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSubtype

`func (o *VirtualizationRealmClient) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *VirtualizationRealmClient) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *VirtualizationRealmClient) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


