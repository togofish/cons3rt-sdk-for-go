# \TestAssetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTestAssetTrustedProject**](TestAssetsApi.md#AddTestAssetTrustedProject) | **Put** /api/testassets/{id}/addtrustedproject | 
[**AddTrustedProject**](TestAssetsApi.md#AddTrustedProject) | **Put** /api/assets/{id}/addtrustedproject | Assign Trusted Project to Asset
[**DeleteAsset**](TestAssetsApi.md#DeleteAsset) | **Delete** /api/assets/{id} | Delete asset
[**Download**](TestAssetsApi.md#Download) | **Get** /api/assets/{id}/download | Download Asset
[**GetTestAsset**](TestAssetsApi.md#GetTestAsset) | **Get** /api/testassets/{id} | Retrieve Test Asset
[**GetTestAssets**](TestAssetsApi.md#GetTestAssets) | **Get** /api/testassets | List Test Assets
[**GetTestAssetsExpanded**](TestAssetsApi.md#GetTestAssetsExpanded) | **Get** /api/testassets/expanded | List all Test Assets, including Project Assets
[**ItarRestrictAsset**](TestAssetsApi.md#ItarRestrictAsset) | **Put** /api/assets/{id}/setitar | Set Asset Export Restriction
[**ListDependentAssets**](TestAssetsApi.md#ListDependentAssets) | **Get** /api/assets/{id}/dependent | List all Dependent Assets
[**RemoveTrustedProject**](TestAssetsApi.md#RemoveTrustedProject) | **Put** /api/assets/{id}/removetrustedproject | Unassign Trusted Project from Asset
[**UpdateAsset**](TestAssetsApi.md#UpdateAsset) | **Put** /api/assets/{id}/update | Update Asset
[**UpdateAssetContent**](TestAssetsApi.md#UpdateAssetContent) | **Put** /api/assets/{id}/updatecontent | Update Asset Content
[**UpdateAssetState**](TestAssetsApi.md#UpdateAssetState) | **Put** /api/assets/{id}/updatestate | Update State
[**UpdateAssetVisibilityQuery**](TestAssetsApi.md#UpdateAssetVisibilityQuery) | **Put** /api/assets/{id}/updatevisibility | Update Visibility
[**UpdateImpactLevel**](TestAssetsApi.md#UpdateImpactLevel) | **Put** /api/assets/{id}/impactlevel | Update Impact Level
[**UpdateInstanceLimit**](TestAssetsApi.md#UpdateInstanceLimit) | **Put** /api/assets/{id}/limit | Update Instance Limit
[**UpdateOfflineStatus**](TestAssetsApi.md#UpdateOfflineStatus) | **Put** /api/assets/{id}/offline | Update Offline Status



## AddTestAssetTrustedProject

> string AddTestAssetTrustedProject(ctx, id).Trustedid(trustedid).Execute()



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
    id := "id_example" // string | 
    trustedid := "trustedid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.AddTestAssetTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.AddTestAssetTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTestAssetTrustedProject`: string
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.AddTestAssetTrustedProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTestAssetTrustedProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedid** | **string** |  | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTrustedProject

> bool AddTrustedProject(ctx, id).Trustedid(trustedid).Execute()

Assign Trusted Project to Asset



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
    id := "id_example" // string | ID of asset
    trustedid := "trustedid_example" // string | ID of project to trust

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.AddTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.AddTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.AddTrustedProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTrustedProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedid** | **string** | ID of project to trust | 

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


## DeleteAsset

> bool DeleteAsset(ctx, id).Force(force).Execute()

Delete asset



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
    id := "id_example" // string | ID of asset to delete
    force := true // bool | Allow delete if there are dependent assets (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.DeleteAsset(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.DeleteAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Allow delete if there are dependent assets | [default to false]

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


## Download

> bool Download(ctx, id).Background(background).Execute()

Download Asset



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
    id := "id_example" // string | ID of asset
    background := true // bool | Force the download to happen in the background (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.Download(context.Background(), id).Background(background).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.Download`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **background** | **bool** | Force the download to happen in the background | [default to false]

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


## GetTestAsset

> FullTestAsset GetTestAsset(ctx, id).Execute()

Retrieve Test Asset



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
    id := "id_example" // string | ID of test asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.GetTestAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.GetTestAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestAsset`: FullTestAsset
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.GetTestAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of test asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullTestAsset**](FullTestAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestAssets

> []MinimalTestAsset GetTestAssets(ctx).Type_(type_).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List Test Assets



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
    type_ := "type__example" // string | Test Asset type (optional)
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.GetTestAssets(context.Background()).Type_(type_).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.GetTestAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestAssets`: []MinimalTestAsset
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.GetTestAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTestAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Test Asset type | 
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalTestAsset**](MinimalTestAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestAssetsExpanded

> []BasicTestAsset GetTestAssetsExpanded(ctx).Type_(type_).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Test Assets, including Project Assets



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
    type_ := "type__example" // string | Test Asset type (optional)
    community := true // bool | Include community assets (optional) (default to false)
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.GetTestAssetsExpanded(context.Background()).Type_(type_).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.GetTestAssetsExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestAssetsExpanded`: []BasicTestAsset
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.GetTestAssetsExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTestAssetsExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Test Asset type | 
 **community** | **bool** | Include community assets | [default to false]
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicTestAsset**](BasicTestAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItarRestrictAsset

> bool ItarRestrictAsset(ctx, id).Execute()

Set Asset Export Restriction



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
    id := "id_example" // string | ID of asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.ItarRestrictAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.ItarRestrictAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItarRestrictAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.ItarRestrictAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiItarRestrictAssetRequest struct via the builder pattern


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


## ListDependentAssets

> []BasicAsset ListDependentAssets(ctx, id).Execute()

List all Dependent Assets



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
    id := "id_example" // string | ID of asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.ListDependentAssets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.ListDependentAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentAssets`: []BasicAsset
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.ListDependentAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDependentAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BasicAsset**](BasicAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTrustedProject

> bool RemoveTrustedProject(ctx, id).Trustedid(trustedid).Execute()

Unassign Trusted Project from Asset



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
    id := "id_example" // string | ID of asset
    trustedid := "trustedid_example" // string | ID of project to untrust

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.RemoveTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.RemoveTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.RemoveTrustedProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTrustedProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedid** | **string** | ID of project to untrust | 

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


## UpdateAsset

> bool UpdateAsset(ctx, id).InputAssetForUpdate(inputAssetForUpdate).Execute()

Update Asset



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
    id := "id_example" // string | ID of asset
    inputAssetForUpdate := *openapiclient.NewInputAssetForUpdate() // InputAssetForUpdate | The modified Asset metadata (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateAsset(context.Background(), id).InputAssetForUpdate(inputAssetForUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputAssetForUpdate** | [**InputAssetForUpdate**](InputAssetForUpdate.md) | The modified Asset metadata | 

### Return type

**bool**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetContent

> bool UpdateAssetContent(ctx, id).File(file).Filename(filename).Execute()

Update Asset Content



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
    id := "id_example" // string | ID of asset
    file := []*os.File{"TODO"} // []*os.File |  (optional)
    filename := "filename_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateAssetContent(context.Background(), id).File(file).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateAssetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetContent`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateAssetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **[]*os.File** |  | 
 **filename** | **string** |  | 

### Return type

**bool**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetState

> bool UpdateAssetState(ctx, id).State(state).Execute()

Update State



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
    id := "id_example" // string | ID of asset to delete
    state := "state_example" // string | The new asset state type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateAssetState(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateAssetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetState`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateAssetState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | The new asset state type | 

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


## UpdateAssetVisibilityQuery

> bool UpdateAssetVisibilityQuery(ctx, id).Visibility(visibility).Execute()

Update Visibility



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
    id := "id_example" // string | ID of asset to update
    visibility := "visibility_example" // string | The new asset visibility type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateAssetVisibilityQuery(context.Background(), id).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateAssetVisibilityQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetVisibilityQuery`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateAssetVisibilityQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetVisibilityQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibility** | **string** | The new asset visibility type | 

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


## UpdateImpactLevel

> bool UpdateImpactLevel(ctx, id).Impactlevel(impactlevel).Execute()

Update Impact Level



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
    id := "id_example" // string | ID of asset
    impactlevel := "impactlevel_example" // string | The new asset impact level type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateImpactLevel(context.Background(), id).Impactlevel(impactlevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateImpactLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImpactLevel`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateImpactLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImpactLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **impactlevel** | **string** | The new asset impact level type. | 

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


## UpdateInstanceLimit

> bool UpdateInstanceLimit(ctx, id).Limit(limit).Execute()

Update Instance Limit



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
    id := "id_example" // string | ID of asset
    limit := int64(789) // int64 | The new asset instance limit

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateInstanceLimit(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateInstanceLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceLimit`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateInstanceLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The new asset instance limit | 

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


## UpdateOfflineStatus

> bool UpdateOfflineStatus(ctx, id).Offline(offline).Execute()

Update Offline Status



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
    id := "id_example" // string | ID of asset to update
    offline := true // bool | Set the asset status to offline (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TestAssetsApi.UpdateOfflineStatus(context.Background(), id).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestAssetsApi.UpdateOfflineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfflineStatus`: bool
    fmt.Fprintf(os.Stdout, "Response from `TestAssetsApi.UpdateOfflineStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfflineStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offline** | **bool** | Set the asset status to offline | [default to true]

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

