# \ClientApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LaunchComposition**](ClientApi.md#LaunchComposition) | **Post** /api/client/{id} | Launch Composition
[**ListCompositionStatus**](ClientApi.md#ListCompositionStatus) | **Get** /api/client | List Composition Statuses



## LaunchComposition

> []int32 LaunchComposition(ctx, id).CompositionLaunchOptions(compositionLaunchOptions).Execute()

Launch Composition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of composition to launch
    compositionLaunchOptions := *openapiclient.NewCompositionLaunchOptions("Username_example", "Password_example") // CompositionLaunchOptions | The Composition launch options (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.LaunchComposition(context.Background(), id).CompositionLaunchOptions(compositionLaunchOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.LaunchComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LaunchComposition`: []int32
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.LaunchComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of composition to launch | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **compositionLaunchOptions** | [**CompositionLaunchOptions**](CompositionLaunchOptions.md) | The Composition launch options | 

### Return type

**[]int32**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompositionStatus

> []AbstractCompositionStatus ListCompositionStatus(ctx).Execute()

List Composition Statuses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.ListCompositionStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.ListCompositionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompositionStatus`: []AbstractCompositionStatus
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.ListCompositionStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCompositionStatusRequest struct via the builder pattern


### Return type

[**[]AbstractCompositionStatus**](AbstractCompositionStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

