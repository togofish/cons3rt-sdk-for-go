# CloudResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Identifier** | **string** |  | 

## Methods

### NewCloudResourceObject

`func NewCloudResourceObject(type_ string, identifier string, ) *CloudResourceObject`

NewCloudResourceObject instantiates a new CloudResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudResourceObjectWithDefaults

`func NewCloudResourceObjectWithDefaults() *CloudResourceObject`

NewCloudResourceObjectWithDefaults instantiates a new CloudResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudResourceObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudResourceObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudResourceObject) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CloudResourceObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudResourceObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudResourceObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudResourceObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *CloudResourceObject) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CloudResourceObject) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CloudResourceObject) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


