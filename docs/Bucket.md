# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudResourceVisibility** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**CloudId** | Pointer to **int32** |  | [optional] 
**ScanningDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewBucket

`func NewBucket(cloudResourceVisibility string, name string, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudResourceVisibility

`func (o *Bucket) GetCloudResourceVisibility() string`

GetCloudResourceVisibility returns the CloudResourceVisibility field if non-nil, zero value otherwise.

### GetCloudResourceVisibilityOk

`func (o *Bucket) GetCloudResourceVisibilityOk() (*string, bool)`

GetCloudResourceVisibilityOk returns a tuple with the CloudResourceVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResourceVisibility

`func (o *Bucket) SetCloudResourceVisibility(v string)`

SetCloudResourceVisibility sets CloudResourceVisibility field to given value.


### GetIdentifier

`func (o *Bucket) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Bucket) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Bucket) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Bucket) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.


### GetCloudId

`func (o *Bucket) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *Bucket) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *Bucket) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *Bucket) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetScanningDisabled

`func (o *Bucket) GetScanningDisabled() bool`

GetScanningDisabled returns the ScanningDisabled field if non-nil, zero value otherwise.

### GetScanningDisabledOk

`func (o *Bucket) GetScanningDisabledOk() (*bool, bool)`

GetScanningDisabledOk returns a tuple with the ScanningDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanningDisabled

`func (o *Bucket) SetScanningDisabled(v bool)`

SetScanningDisabled sets ScanningDisabled field to given value.

### HasScanningDisabled

`func (o *Bucket) HasScanningDisabled() bool`

HasScanningDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


