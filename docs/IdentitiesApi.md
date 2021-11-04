# \IdentitiesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentity**](IdentitiesApi.md#CreateIdentity) | **Post** /api/drs/{id}/host/{hostid}/identity | Create a host identity
[**DeleteIdentity**](IdentitiesApi.md#DeleteIdentity) | **Delete** /api/drs/{id}/host/{hostid}/identity | Delete host identity
[**DeleteIdentityById**](IdentitiesApi.md#DeleteIdentityById) | **Delete** /api/drs/{id}/host/{hostid}/identity/{username} | Deletes identity for specified user
[**GetIdentities**](IdentitiesApi.md#GetIdentities) | **Get** /api/drs/{id}/host/{hostid}/identities | Get Host Identities
[**GetIdentity**](IdentitiesApi.md#GetIdentity) | **Get** /api/drs/{id}/host/{hostid}/identity | Get Host Identity For User



## CreateIdentity

> []BaseIdentity CreateIdentity(ctx, id, hostid).CloudResourceObject(cloudResourceObject).Execute()

Create a host identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host
    cloudResourceObject := []cons3rtclient.CloudResourceObject{*cons3rtclient.NewCloudResourceObject("Type_example", "Identifier_example")} // []CloudResourceObject | The cloud resources to be accessed by the host identity

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitiesApi.CreateIdentity(context.Background(), id, hostid).CloudResourceObject(cloudResourceObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.CreateIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentity`: []BaseIdentity
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.CreateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudResourceObject** | [**[]CloudResourceObject**](CloudResourceObject.md) | The cloud resources to be accessed by the host identity | 

### Return type

[**[]BaseIdentity**](BaseIdentity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentity

> bool DeleteIdentity(ctx, id, hostid).Execute()

Delete host identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitiesApi.DeleteIdentity(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.DeleteIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentity`: bool
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.DeleteIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityById

> []BaseIdentity DeleteIdentityById(ctx, id, hostid, username).Execute()

Deletes identity for specified user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host
    username := "username_example" // string | Username of the identity to be deleted

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitiesApi.DeleteIdentityById(context.Background(), id, hostid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.DeleteIdentityById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentityById`: []BaseIdentity
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.DeleteIdentityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 
**username** | **string** | Username of the identity to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BaseIdentity**](BaseIdentity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentities

> []BaseIdentity GetIdentities(ctx, id, hostid).Execute()

Get Host Identities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitiesApi.GetIdentities(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.GetIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentities`: []BaseIdentity
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.GetIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BaseIdentity**](BaseIdentity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> []CloudResourceAccessListing GetIdentity(ctx, id, hostid).Execute()

Get Host Identity For User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitiesApi.GetIdentity(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.GetIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentity`: []CloudResourceAccessListing
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CloudResourceAccessListing**](CloudResourceAccessListing.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

