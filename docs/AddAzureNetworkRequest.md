# AddAzureNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatInstanceType** | **string** |  | 
**NatImageReference** | [**ImageReferenceDTO**](ImageReferenceDTO.md) |  | 

## Methods

### NewAddAzureNetworkRequest

`func NewAddAzureNetworkRequest(natInstanceType string, natImageReference ImageReferenceDTO, ) *AddAzureNetworkRequest`

NewAddAzureNetworkRequest instantiates a new AddAzureNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAzureNetworkRequestWithDefaults

`func NewAddAzureNetworkRequestWithDefaults() *AddAzureNetworkRequest`

NewAddAzureNetworkRequestWithDefaults instantiates a new AddAzureNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatInstanceType

`func (o *AddAzureNetworkRequest) GetNatInstanceType() string`

GetNatInstanceType returns the NatInstanceType field if non-nil, zero value otherwise.

### GetNatInstanceTypeOk

`func (o *AddAzureNetworkRequest) GetNatInstanceTypeOk() (*string, bool)`

GetNatInstanceTypeOk returns a tuple with the NatInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInstanceType

`func (o *AddAzureNetworkRequest) SetNatInstanceType(v string)`

SetNatInstanceType sets NatInstanceType field to given value.


### GetNatImageReference

`func (o *AddAzureNetworkRequest) GetNatImageReference() ImageReferenceDTO`

GetNatImageReference returns the NatImageReference field if non-nil, zero value otherwise.

### GetNatImageReferenceOk

`func (o *AddAzureNetworkRequest) GetNatImageReferenceOk() (*ImageReferenceDTO, bool)`

GetNatImageReferenceOk returns a tuple with the NatImageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatImageReference

`func (o *AddAzureNetworkRequest) SetNatImageReference(v ImageReferenceDTO)`

SetNatImageReference sets NatImageReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


