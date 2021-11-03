# AzureCloudResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]ImageReferenceDTO**](ImageReferenceDTO.md) |  | [optional] 

## Methods

### NewAzureCloudResources

`func NewAzureCloudResources() *AzureCloudResources`

NewAzureCloudResources instantiates a new AzureCloudResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudResourcesWithDefaults

`func NewAzureCloudResourcesWithDefaults() *AzureCloudResources`

NewAzureCloudResourcesWithDefaults instantiates a new AzureCloudResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *AzureCloudResources) GetImages() []ImageReferenceDTO`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AzureCloudResources) GetImagesOk() (*[]ImageReferenceDTO, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AzureCloudResources) SetImages(v []ImageReferenceDTO)`

SetImages sets Images field to given value.

### HasImages

`func (o *AzureCloudResources) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


