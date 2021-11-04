# \SystemAssetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemAsset**](SystemAssetsApi.md#CreateSystemAsset) | **Post** /api/systemassets | Create System Asset
[**DeleteSystemAsset**](SystemAssetsApi.md#DeleteSystemAsset) | **Delete** /api/systemassets/{id} | Delete System Asset
[**ImportSystemAsset**](SystemAssetsApi.md#ImportSystemAsset) | **Post** /api/systemassets/{id}/import | Import System Asset
[**ListSystemAssets**](SystemAssetsApi.md#ListSystemAssets) | **Get** /api/systemassets | List System Assets
[**RetrieveSystemAsset**](SystemAssetsApi.md#RetrieveSystemAsset) | **Get** /api/systemassets/{id} | Retrieve System Asset
[**UpdateSystemAsset**](SystemAssetsApi.md#UpdateSystemAsset) | **Put** /api/systemassets/{id} | Update System Asset



## CreateSystemAsset

> []FullSystemAsset CreateSystemAsset(ctx).InputSystemAsset(inputSystemAsset).Execute()

Create System Asset



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
    inputSystemAsset := *openapiclient.NewInputSystemAsset("Type_example") // InputSystemAsset | The System Asset definition (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemAssetsApi.CreateSystemAsset(context.Background()).InputSystemAsset(inputSystemAsset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAssetsApi.CreateSystemAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystemAsset`: []FullSystemAsset
    fmt.Fprintf(os.Stdout, "Response from `SystemAssetsApi.CreateSystemAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputSystemAsset** | [**InputSystemAsset**](InputSystemAsset.md) | The System Asset definition | 

### Return type

[**[]FullSystemAsset**](FullSystemAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSystemAsset

> bool DeleteSystemAsset(ctx, id).Execute()

Delete System Asset



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
    id := "id_example" // string | ID of system asset to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemAssetsApi.DeleteSystemAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAssetsApi.DeleteSystemAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSystemAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemAssetsApi.DeleteSystemAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system asset to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemAssetRequest struct via the builder pattern


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


## ImportSystemAsset

> bool ImportSystemAsset(ctx, id).Execute()

Import System Asset



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
    id := "id_example" // string | ID of system asset to import

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemAssetsApi.ImportSystemAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAssetsApi.ImportSystemAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSystemAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemAssetsApi.ImportSystemAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system asset to import | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSystemAssetRequest struct via the builder pattern


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


## ListSystemAssets

> []MinimalSystemAsset ListSystemAssets(ctx).Execute()

List System Assets



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
    resp, r, err := api_client.SystemAssetsApi.ListSystemAssets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAssetsApi.ListSystemAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemAssets`: []MinimalSystemAsset
    fmt.Fprintf(os.Stdout, "Response from `SystemAssetsApi.ListSystemAssets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemAssetsRequest struct via the builder pattern


### Return type

[**[]MinimalSystemAsset**](MinimalSystemAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSystemAsset

> []FullSystemAsset RetrieveSystemAsset(ctx, id).Execute()

Retrieve System Asset



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
    id := "id_example" // string | ID of System Asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemAssetsApi.RetrieveSystemAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAssetsApi.RetrieveSystemAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSystemAsset`: []FullSystemAsset
    fmt.Fprintf(os.Stdout, "Response from `SystemAssetsApi.RetrieveSystemAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of System Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSystemAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FullSystemAsset**](FullSystemAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemAsset

> FullSystemAsset UpdateSystemAsset(ctx, id).InputSystemAsset(inputSystemAsset).Execute()

Update System Asset



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
    id := "id_example" // string | ID of System Asset to update
    inputSystemAsset := *openapiclient.NewInputSystemAsset("Type_example") // InputSystemAsset | The modified System Asset definition (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemAssetsApi.UpdateSystemAsset(context.Background(), id).InputSystemAsset(inputSystemAsset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAssetsApi.UpdateSystemAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSystemAsset`: FullSystemAsset
    fmt.Fprintf(os.Stdout, "Response from `SystemAssetsApi.UpdateSystemAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of System Asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputSystemAsset** | [**InputSystemAsset**](InputSystemAsset.md) | The modified System Asset definition | 

### Return type

[**FullSystemAsset**](FullSystemAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

