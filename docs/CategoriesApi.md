# \CategoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCategoryToAsset**](CategoriesApi.md#AddCategoryToAsset) | **Put** /api/categories/{id}/asset | Assign Category to Asset
[**AddCategoryToDeploymentRun**](CategoriesApi.md#AddCategoryToDeploymentRun) | **Put** /api/categories/{id}/run | Assign Category to Run
[**CreateCategory**](CategoriesApi.md#CreateCategory) | **Post** /api/categories | Create a Category
[**DeleteCategory**](CategoriesApi.md#DeleteCategory) | **Delete** /api/categories/{id} | Delete Category
[**GetCategories**](CategoriesApi.md#GetCategories) | **Get** /api/categories | List all Categories
[**RemoveCategoryFromAsset**](CategoriesApi.md#RemoveCategoryFromAsset) | **Delete** /api/categories/{id}/asset | Unassign Category from Asset
[**RemoveCategoryFromDeploymentRun**](CategoriesApi.md#RemoveCategoryFromDeploymentRun) | **Delete** /api/categories/{id}/run | Unassign Category from deployment run
[**SetCategoryParent**](CategoriesApi.md#SetCategoryParent) | **Put** /api/categories/{id}/parent | Set Parent Category
[**UpdateCategory**](CategoriesApi.md#UpdateCategory) | **Put** /api/categories/{id} | Update Category



## AddCategoryToAsset

> bool AddCategoryToAsset(ctx, id).Assetid(assetid).Execute()

Assign Category to Asset



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
    id := "id_example" // string | ID of category
    assetid := "assetid_example" // string | ID of asset to assign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.AddCategoryToAsset(context.Background(), id).Assetid(assetid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.AddCategoryToAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCategoryToAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.AddCategoryToAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCategoryToAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetid** | **string** | ID of asset to assign | 

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


## AddCategoryToDeploymentRun

> bool AddCategoryToDeploymentRun(ctx, id).Runid(runid).Execute()

Assign Category to Run



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
    id := "id_example" // string | ID of category
    runid := "runid_example" // string | ID of run to assign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.AddCategoryToDeploymentRun(context.Background(), id).Runid(runid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.AddCategoryToDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCategoryToDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.AddCategoryToDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCategoryToDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runid** | **string** | ID of run to assign | 

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


## CreateCategory

> FullCategory CreateCategory(ctx).InputCategory(inputCategory).Execute()

Create a Category



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
    inputCategory := *cons3rtclient.NewInputCategory("Name_example") // InputCategory | The Category to create (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.CreateCategory(context.Background()).InputCategory(inputCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: FullCategory
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.CreateCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCategory** | [**InputCategory**](InputCategory.md) | The Category to create | 

### Return type

[**FullCategory**](FullCategory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> bool DeleteCategory(ctx, id).Execute()

Delete Category



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
    id := "id_example" // string | ID of category to delete

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.DeleteCategory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.DeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCategory`: bool
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.DeleteCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## GetCategories

> []MinimalCategory GetCategories(ctx).Execute()

List all Categories



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.GetCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategories`: []MinimalCategory
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesRequest struct via the builder pattern


### Return type

[**[]MinimalCategory**](MinimalCategory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCategoryFromAsset

> bool RemoveCategoryFromAsset(ctx, id).Assetid(assetid).Execute()

Unassign Category from Asset



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
    id := "id_example" // string | ID of category
    assetid := "assetid_example" // string | ID of asset to unassign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.RemoveCategoryFromAsset(context.Background(), id).Assetid(assetid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.RemoveCategoryFromAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCategoryFromAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.RemoveCategoryFromAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCategoryFromAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetid** | **string** | ID of asset to unassign | 

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


## RemoveCategoryFromDeploymentRun

> bool RemoveCategoryFromDeploymentRun(ctx, id).Runid(runid).Execute()

Unassign Category from deployment run



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
    id := "id_example" // string | ID of category
    runid := "runid_example" // string | ID of run to unassign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.RemoveCategoryFromDeploymentRun(context.Background(), id).Runid(runid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.RemoveCategoryFromDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCategoryFromDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.RemoveCategoryFromDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCategoryFromDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runid** | **string** | ID of run to unassign | 

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


## SetCategoryParent

> FullCategory SetCategoryParent(ctx, id).Parentid(parentid).Execute()

Set Parent Category



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
    id := "id_example" // string | ID of category to modify
    parentid := "parentid_example" // string | ID of desired parent category

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.SetCategoryParent(context.Background(), id).Parentid(parentid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.SetCategoryParent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCategoryParent`: FullCategory
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.SetCategoryParent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category to modify | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCategoryParentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentid** | **string** | ID of desired parent category | 

### Return type

[**FullCategory**](FullCategory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCategory

> FullCategory UpdateCategory(ctx, id).InputCategory(inputCategory).Execute()

Update Category



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
    id := "id_example" // string | ID of Category to update
    inputCategory := *cons3rtclient.NewInputCategory("Name_example") // InputCategory | The modified Category definition (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.UpdateCategory(context.Background(), id).InputCategory(inputCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: FullCategory
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Category to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCategory** | [**InputCategory**](InputCategory.md) | The modified Category definition | 

### Return type

[**FullCategory**](FullCategory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

