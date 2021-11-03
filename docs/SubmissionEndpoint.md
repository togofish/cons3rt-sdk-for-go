# SubmissionEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewSubmissionEndpoint

`func NewSubmissionEndpoint(host string, subtype string, ) *SubmissionEndpoint`

NewSubmissionEndpoint instantiates a new SubmissionEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionEndpointWithDefaults

`func NewSubmissionEndpointWithDefaults() *SubmissionEndpoint`

NewSubmissionEndpointWithDefaults instantiates a new SubmissionEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SubmissionEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SubmissionEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SubmissionEndpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetId

`func (o *SubmissionEndpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionEndpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionEndpoint) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *SubmissionEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SubmissionEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SubmissionEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SubmissionEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetType

`func (o *SubmissionEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubmissionEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubmissionEndpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubmissionEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *SubmissionEndpoint) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SubmissionEndpoint) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SubmissionEndpoint) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


