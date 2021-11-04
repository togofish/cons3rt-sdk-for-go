# \SoftwareBundlesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSoftwareAssetBundleFromSystemModule**](SoftwareBundlesApi.md#CreateSoftwareAssetBundleFromSystemModule) | **Post** /api/software/bundles | Create Software Bundle
[**DeleteSoftwareAssetBundle**](SoftwareBundlesApi.md#DeleteSoftwareAssetBundle) | **Delete** /api/software/bundles/{id} | Delete Software Bundle
[**GetSoftwareAssetBundleExpanded**](SoftwareBundlesApi.md#GetSoftwareAssetBundleExpanded) | **Get** /api/software/bundles/expanded | List all Software Bundles, including Project Assets
[**GetSoftwareBundle**](SoftwareBundlesApi.md#GetSoftwareBundle) | **Get** /api/software/bundles/{id} | Retrieve Software Bundle
[**GetSoftwareBundles**](SoftwareBundlesApi.md#GetSoftwareBundles) | **Get** /api/software/bundles | List Software Bundles



## CreateSoftwareAssetBundleFromSystemModule

> string CreateSoftwareAssetBundleFromSystemModule(ctx).SystemId(systemId).Name(name).Body(body).Execute()

Create Software Bundle



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
    systemId := "systemId_example" // string | ID of system module
    name := "name_example" // string | Name of the new software bundle
    body := "body_example" // string | Description of the new software bundle (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareBundlesApi.CreateSoftwareAssetBundleFromSystemModule(context.Background()).SystemId(systemId).Name(name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareBundlesApi.CreateSoftwareAssetBundleFromSystemModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSoftwareAssetBundleFromSystemModule`: string
    fmt.Fprintf(os.Stdout, "Response from `SoftwareBundlesApi.CreateSoftwareAssetBundleFromSystemModule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoftwareAssetBundleFromSystemModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string** | ID of system module | 
 **name** | **string** | Name of the new software bundle | 
 **body** | **string** | Description of the new software bundle | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSoftwareAssetBundle

> bool DeleteSoftwareAssetBundle(ctx, id).Execute()

Delete Software Bundle



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
    id := "id_example" // string | ID of software bundle

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareBundlesApi.DeleteSoftwareAssetBundle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareBundlesApi.DeleteSoftwareAssetBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSoftwareAssetBundle`: bool
    fmt.Fprintf(os.Stdout, "Response from `SoftwareBundlesApi.DeleteSoftwareAssetBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of software bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSoftwareAssetBundleRequest struct via the builder pattern


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


## GetSoftwareAssetBundleExpanded

> []BasicSoftwareAssetBundle GetSoftwareAssetBundleExpanded(ctx).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Software Bundles, including Project Assets



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
    community := true // bool | Include community assets (optional) (default to false)
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareBundlesApi.GetSoftwareAssetBundleExpanded(context.Background()).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareBundlesApi.GetSoftwareAssetBundleExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareAssetBundleExpanded`: []BasicSoftwareAssetBundle
    fmt.Fprintf(os.Stdout, "Response from `SoftwareBundlesApi.GetSoftwareAssetBundleExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareAssetBundleExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **community** | **bool** | Include community assets | [default to false]
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicSoftwareAssetBundle**](BasicSoftwareAssetBundle.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareBundle

> FullSoftwareAssetBundle GetSoftwareBundle(ctx, id).Execute()

Retrieve Software Bundle



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
    id := "id_example" // string | ID of software bundle

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareBundlesApi.GetSoftwareBundle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareBundlesApi.GetSoftwareBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareBundle`: FullSoftwareAssetBundle
    fmt.Fprintf(os.Stdout, "Response from `SoftwareBundlesApi.GetSoftwareBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of software bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullSoftwareAssetBundle**](FullSoftwareAssetBundle.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareBundles

> []MinimalSoftwareAssetBundle GetSoftwareBundles(ctx).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List Software Bundles



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
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.SoftwareBundlesApi.GetSoftwareBundles(context.Background()).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareBundlesApi.GetSoftwareBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftwareBundles`: []MinimalSoftwareAssetBundle
    fmt.Fprintf(os.Stdout, "Response from `SoftwareBundlesApi.GetSoftwareBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalSoftwareAssetBundle**](MinimalSoftwareAssetBundle.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

