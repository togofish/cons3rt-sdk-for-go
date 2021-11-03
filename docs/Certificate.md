# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pem** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificate

`func NewCertificate() *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPem

`func (o *Certificate) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *Certificate) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *Certificate) SetPem(v string)`

SetPem sets Pem field to given value.

### HasPem

`func (o *Certificate) HasPem() bool`

HasPem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


