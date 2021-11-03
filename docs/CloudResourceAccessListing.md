# CloudResourceAccessListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**CloudResourceObject**](CloudResourceObject.md) |  | [optional] 
**Access** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudResourceAccessListing

`func NewCloudResourceAccessListing() *CloudResourceAccessListing`

NewCloudResourceAccessListing instantiates a new CloudResourceAccessListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudResourceAccessListingWithDefaults

`func NewCloudResourceAccessListingWithDefaults() *CloudResourceAccessListing`

NewCloudResourceAccessListingWithDefaults instantiates a new CloudResourceAccessListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *CloudResourceAccessListing) GetResource() CloudResourceObject`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CloudResourceAccessListing) GetResourceOk() (*CloudResourceObject, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CloudResourceAccessListing) SetResource(v CloudResourceObject)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CloudResourceAccessListing) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAccess

`func (o *CloudResourceAccessListing) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CloudResourceAccessListing) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CloudResourceAccessListing) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CloudResourceAccessListing) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


