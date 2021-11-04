# \ContainerAssetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustedProject**](ContainerAssetsApi.md#AddTrustedProject) | **Put** /api/assets/{id}/addtrustedproject | Assign Trusted Project to Asset
[**DeleteAsset**](ContainerAssetsApi.md#DeleteAsset) | **Delete** /api/assets/{id} | Delete asset
[**Download**](ContainerAssetsApi.md#Download) | **Get** /api/assets/{id}/download | Download Asset
[**GetContainer**](ContainerAssetsApi.md#GetContainer) | **Get** /api/containers/{id} | Retrieve Container Asset
[**GetContainerAssets**](ContainerAssetsApi.md#GetContainerAssets) | **Get** /api/containers | List all Containers
[**GetContainerAssetsExpanded**](ContainerAssetsApi.md#GetContainerAssetsExpanded) | **Get** /api/containers/expanded | List all Containers, including Project Assets
[**ItarRestrictAsset**](ContainerAssetsApi.md#ItarRestrictAsset) | **Put** /api/assets/{id}/setitar | Set Asset Export Restriction
[**ListDependentAssets**](ContainerAssetsApi.md#ListDependentAssets) | **Get** /api/assets/{id}/dependent | List all Dependent Assets
[**RemoveTrustedProject**](ContainerAssetsApi.md#RemoveTrustedProject) | **Put** /api/assets/{id}/removetrustedproject | Unassign Trusted Project from Asset
[**SubmitAssetToSubmissionService**](ContainerAssetsApi.md#SubmitAssetToSubmissionService) | **Post** /api/containers/{id}/submit/{submission_service_id} | Submit Container to the Project&#39;s Submission Service
[**SubmitBucketResourceToSubmissionService**](ContainerAssetsApi.md#SubmitBucketResourceToSubmissionService) | **Post** /api/buckets/{id}/submit/{submission_service_id} | Submit Bucket Resource to the Project&#39;s Submission Service
[**UpdateAsset**](ContainerAssetsApi.md#UpdateAsset) | **Put** /api/assets/{id}/update | Update Asset
[**UpdateAssetContent**](ContainerAssetsApi.md#UpdateAssetContent) | **Put** /api/assets/{id}/updatecontent | Update Asset Content
[**UpdateAssetState**](ContainerAssetsApi.md#UpdateAssetState) | **Put** /api/assets/{id}/updatestate | Update State
[**UpdateAssetVisibilityQuery**](ContainerAssetsApi.md#UpdateAssetVisibilityQuery) | **Put** /api/assets/{id}/updatevisibility | Update Visibility
[**UpdateImpactLevel**](ContainerAssetsApi.md#UpdateImpactLevel) | **Put** /api/assets/{id}/impactlevel | Update Impact Level
[**UpdateInstanceLimit**](ContainerAssetsApi.md#UpdateInstanceLimit) | **Put** /api/assets/{id}/limit | Update Instance Limit
[**UpdateOfflineStatus**](ContainerAssetsApi.md#UpdateOfflineStatus) | **Put** /api/assets/{id}/offline | Update Offline Status



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
    resp, r, err := api_client.ContainerAssetsApi.AddTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.AddTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.AddTrustedProject`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.DeleteAsset(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.DeleteAsset`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.Download(context.Background(), id).Background(background).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.Download`: %v\n", resp)
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


## GetContainer

> FullContainerAsset GetContainer(ctx, id).Execute()

Retrieve Container Asset



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
    id := "id_example" // string | ID of container to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerAssetsApi.GetContainer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.GetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainer`: FullContainerAsset
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of container to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullContainerAsset**](FullContainerAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerAssets

> []MinimalContainerAsset GetContainerAssets(ctx).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Containers



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
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerAssetsApi.GetContainerAssets(context.Background()).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.GetContainerAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerAssets`: []MinimalContainerAsset
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.GetContainerAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalContainerAsset**](MinimalContainerAsset.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerAssetsExpanded

> []BasicContainerAsset GetContainerAssetsExpanded(ctx).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Containers, including Project Assets



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
    community := true // bool | Include community assets (optional) (default to false)
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerAssetsApi.GetContainerAssetsExpanded(context.Background()).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.GetContainerAssetsExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerAssetsExpanded`: []BasicContainerAsset
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.GetContainerAssetsExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerAssetsExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **community** | **bool** | Include community assets | [default to false]
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicContainerAsset**](BasicContainerAsset.md)

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
    resp, r, err := api_client.ContainerAssetsApi.ItarRestrictAsset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.ItarRestrictAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItarRestrictAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.ItarRestrictAsset`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.ListDependentAssets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.ListDependentAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentAssets`: []BasicAsset
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.ListDependentAssets`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.RemoveTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.RemoveTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.RemoveTrustedProject`: %v\n", resp)
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


## SubmitAssetToSubmissionService

> bool SubmitAssetToSubmissionService(ctx, id, submissionServiceId).InputSubmissionServiceForAssetSubmission(inputSubmissionServiceForAssetSubmission).Execute()

Submit Container to the Project's Submission Service



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
    id := "id_example" // string | ID of Container Asset to publish
    submissionServiceId := "submissionServiceId_example" // string | ID of project submission service
    inputSubmissionServiceForAssetSubmission := *openapiclient.NewInputSubmissionServiceForAssetSubmission(*openapiclient.NewInputSubmissionEndpointForAssetSubmission("Subtype_example")) // InputSubmissionServiceForAssetSubmission | Submission service override values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerAssetsApi.SubmitAssetToSubmissionService(context.Background(), id, submissionServiceId).InputSubmissionServiceForAssetSubmission(inputSubmissionServiceForAssetSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.SubmitAssetToSubmissionService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAssetToSubmissionService`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.SubmitAssetToSubmissionService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Container Asset to publish | 
**submissionServiceId** | **string** | ID of project submission service | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAssetToSubmissionServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputSubmissionServiceForAssetSubmission** | [**InputSubmissionServiceForAssetSubmission**](InputSubmissionServiceForAssetSubmission.md) | Submission service override values | 

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


## SubmitBucketResourceToSubmissionService

> bool SubmitBucketResourceToSubmissionService(ctx, id, submissionServiceId).FileName(fileName).InputSubmissionServiceForAssetSubmission(inputSubmissionServiceForAssetSubmission).Execute()

Submit Bucket Resource to the Project's Submission Service



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
    id := "id_example" // string | ID of the bucket to publish a resource from
    submissionServiceId := "submissionServiceId_example" // string | ID of project submission service
    fileName := "fileName_example" // string | The filename within the bucket to download
    inputSubmissionServiceForAssetSubmission := *openapiclient.NewInputSubmissionServiceForAssetSubmission(*openapiclient.NewInputSubmissionEndpointForAssetSubmission("Subtype_example")) // InputSubmissionServiceForAssetSubmission | Submission service override values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerAssetsApi.SubmitBucketResourceToSubmissionService(context.Background(), id, submissionServiceId).FileName(fileName).InputSubmissionServiceForAssetSubmission(inputSubmissionServiceForAssetSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.SubmitBucketResourceToSubmissionService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitBucketResourceToSubmissionService`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.SubmitBucketResourceToSubmissionService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the bucket to publish a resource from | 
**submissionServiceId** | **string** | ID of project submission service | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBucketResourceToSubmissionServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileName** | **string** | The filename within the bucket to download | 
 **inputSubmissionServiceForAssetSubmission** | [**InputSubmissionServiceForAssetSubmission**](InputSubmissionServiceForAssetSubmission.md) | Submission service override values | 

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
    resp, r, err := api_client.ContainerAssetsApi.UpdateAsset(context.Background(), id).InputAssetForUpdate(inputAssetForUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateAsset`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.UpdateAssetContent(context.Background(), id).File(file).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateAssetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetContent`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateAssetContent`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.UpdateAssetState(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateAssetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetState`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateAssetState`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.UpdateAssetVisibilityQuery(context.Background(), id).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateAssetVisibilityQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetVisibilityQuery`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateAssetVisibilityQuery`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.UpdateImpactLevel(context.Background(), id).Impactlevel(impactlevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateImpactLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImpactLevel`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateImpactLevel`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.UpdateInstanceLimit(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateInstanceLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceLimit`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateInstanceLimit`: %v\n", resp)
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
    resp, r, err := api_client.ContainerAssetsApi.UpdateOfflineStatus(context.Background(), id).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAssetsApi.UpdateOfflineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfflineStatus`: bool
    fmt.Fprintf(os.Stdout, "Response from `ContainerAssetsApi.UpdateOfflineStatus`: %v\n", resp)
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

