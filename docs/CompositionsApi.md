# \CompositionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComposition**](CompositionsApi.md#DeleteComposition) | **Delete** /api/compositions/{id} | Delete Composition
[**GetComposition**](CompositionsApi.md#GetComposition) | **Get** /api/compositions/{id} | Retrieve Composition
[**ListCompositions**](CompositionsApi.md#ListCompositions) | **Get** /api/compositions | List all Compositions
[**PublishDeploymentRun**](CompositionsApi.md#PublishDeploymentRun) | **Post** /api/drs/{id}/publish | Publish Deployment Run
[**PublishScenarioToComposition**](CompositionsApi.md#PublishScenarioToComposition) | **Post** /api/scenarios/{id}/publish | Publish Scenario
[**UnpublishDeploymentRun**](CompositionsApi.md#UnpublishDeploymentRun) | **Delete** /api/drs/{id}/publish | Unpublish Deployment Run
[**UpdateComposition**](CompositionsApi.md#UpdateComposition) | **Put** /api/compositions/{id} | Update Composition



## DeleteComposition

> bool DeleteComposition(ctx, id).Execute()

Delete Composition



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
    id := "id_example" // string | ID of composition to delete

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.DeleteComposition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.DeleteComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteComposition`: bool
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.DeleteComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of composition to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompositionRequest struct via the builder pattern


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


## GetComposition

> FullComposition GetComposition(ctx, id).Execute()

Retrieve Composition



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
    id := "id_example" // string | ID of composition to return

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.GetComposition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.GetComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComposition`: FullComposition
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.GetComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of composition to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullComposition**](FullComposition.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompositions

> []FullComposition ListCompositions(ctx).Maxresults(maxresults).Page(page).Execute()

List all Compositions



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
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.ListCompositions(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.ListCompositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompositions`: []FullComposition
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.ListCompositions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]FullComposition**](FullComposition.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishDeploymentRun

> bool PublishDeploymentRun(ctx, id).Execute()

Publish Deployment Run



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.PublishDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.PublishDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.PublishDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishDeploymentRunRequest struct via the builder pattern


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


## PublishScenarioToComposition

> string PublishScenarioToComposition(ctx, id).InputComposition(inputComposition).Execute()

Publish Scenario



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
    id := "id_example" // string | 
    inputComposition := *cons3rtclient.NewInputComposition("Name_example", *cons3rtclient.NewInputCompositionRunOptions(int32(123))) // InputComposition | The composition definition used when launching the deployment (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.PublishScenarioToComposition(context.Background(), id).InputComposition(inputComposition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.PublishScenarioToComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishScenarioToComposition`: string
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.PublishScenarioToComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishScenarioToCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputComposition** | [**InputComposition**](InputComposition.md) | The composition definition used when launching the deployment | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishDeploymentRun

> bool UnpublishDeploymentRun(ctx, id).Execute()

Unpublish Deployment Run



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.UnpublishDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.UnpublishDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnpublishDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.UnpublishDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishDeploymentRunRequest struct via the builder pattern


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


## UpdateComposition

> FullComposition UpdateComposition(ctx, id).InputComposition(inputComposition).Execute()

Update Composition



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
    id := "id_example" // string | ID of Composition to update
    inputComposition := *cons3rtclient.NewInputComposition("Name_example", *cons3rtclient.NewInputCompositionRunOptions(int32(123))) // InputComposition | The modified Composition definition (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompositionsApi.UpdateComposition(context.Background(), id).InputComposition(inputComposition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositionsApi.UpdateComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComposition`: FullComposition
    fmt.Fprintf(os.Stdout, "Response from `CompositionsApi.UpdateComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Composition to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputComposition** | [**InputComposition**](InputComposition.md) | The modified Composition definition | 

### Return type

[**FullComposition**](FullComposition.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

