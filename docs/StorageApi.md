# \StorageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](StorageApi.md#CreateBucket) | **Post** /api/buckets | Create Storage Bucket
[**DeleteBucket**](StorageApi.md#DeleteBucket) | **Delete** /api/buckets/{id} | Delete Storage Buckets
[**DownloadFileFromBucket**](StorageApi.md#DownloadFileFromBucket) | **Get** /api/buckets/{id}/download | Download File From Bucket
[**GetBucket**](StorageApi.md#GetBucket) | **Get** /api/buckets/{id} | Retrieve Storage Buckets
[**GetBucketListing**](StorageApi.md#GetBucketListing) | **Get** /api/buckets/{id}/listing | List Bucket Contents
[**ListAvailableCloudsForBuckets**](StorageApi.md#ListAvailableCloudsForBuckets) | **Get** /api/buckets/clouds | List Clouds Available For Bucket Creation
[**ListBuckets**](StorageApi.md#ListBuckets) | **Get** /api/buckets | List Storage Buckets
[**SubmitBucketResourceToSubmissionService**](StorageApi.md#SubmitBucketResourceToSubmissionService) | **Post** /api/buckets/{id}/submit/{submission_service_id} | Submit Bucket Resource to the Project&#39;s Submission Service
[**UpdateBucket**](StorageApi.md#UpdateBucket) | **Put** /api/buckets/{id} | Update Storage Buckets
[**UploadFileToBucket**](StorageApi.md#UploadFileToBucket) | **Post** /api/buckets/{id} | Upload File to Bucket



## CreateBucket

> string CreateBucket(ctx).Bucket(bucket).Execute()

Create Storage Bucket



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
    bucket := *cons3rtclient.NewBucket("CloudResourceVisibility_example", "Name_example") // Bucket | The bucket creation information

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.CreateBucket(context.Background()).Bucket(bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBucket`: string
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | [**Bucket**](Bucket.md) | The bucket creation information | 

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


## DeleteBucket

> bool DeleteBucket(ctx, id).Execute()

Delete Storage Buckets



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
    id := "id_example" // string | ID of bucket

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.DeleteBucket(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBucket`: bool
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.DeleteBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


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


## DownloadFileFromBucket

> bool DownloadFileFromBucket(ctx, id).FileName(fileName).Background(background).Execute()

Download File From Bucket



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
    id := "id_example" // string | ID of bucket
    fileName := "fileName_example" // string | The filename within the bucket to download
    background := true // bool | Force the download to happen in the background (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.DownloadFileFromBucket(context.Background(), id).FileName(fileName).Background(background).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.DownloadFileFromBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFileFromBucket`: bool
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.DownloadFileFromBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileFromBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileName** | **string** | The filename within the bucket to download | 
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


## GetBucket

> Bucket GetBucket(ctx, id).Execute()

Retrieve Storage Buckets



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
    id := "id_example" // string | ID of bucket

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.GetBucket(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucket`: Bucket
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bucket**](Bucket.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketListing

> string GetBucketListing(ctx, id).Execute()

List Bucket Contents



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
    id := "id_example" // string | ID of bucket

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.GetBucketListing(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetBucketListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketListing`: string
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetBucketListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListAvailableCloudsForBuckets

> []MinimalCloud ListAvailableCloudsForBuckets(ctx).Maxresults(maxresults).Page(page).Execute()

List Clouds Available For Bucket Creation



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
    resp, r, err := api_client.StorageApi.ListAvailableCloudsForBuckets(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.ListAvailableCloudsForBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableCloudsForBuckets`: []MinimalCloud
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.ListAvailableCloudsForBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableCloudsForBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalCloud**](MinimalCloud.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> []Bucket ListBuckets(ctx).Cloud(cloud).Project(project).Maxresults(maxresults).Page(page).Execute()

List Storage Buckets



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
    cloud := "cloud_example" // string | ID of the cloud (optional)
    project := true // bool | Include project buckets (optional) (default to false)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.ListBuckets(context.Background()).Cloud(cloud).Project(project).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.ListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuckets`: []Bucket
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.ListBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | ID of the cloud | 
 **project** | **bool** | Include project buckets | [default to false]
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]Bucket**](Bucket.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of the bucket to publish a resource from
    submissionServiceId := "submissionServiceId_example" // string | ID of project submission service
    fileName := "fileName_example" // string | The filename within the bucket to download
    inputSubmissionServiceForAssetSubmission := *cons3rtclient.NewInputSubmissionServiceForAssetSubmission(*cons3rtclient.NewInputSubmissionEndpointForAssetSubmission("Subtype_example")) // InputSubmissionServiceForAssetSubmission | Submission service override values (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.SubmitBucketResourceToSubmissionService(context.Background(), id, submissionServiceId).FileName(fileName).InputSubmissionServiceForAssetSubmission(inputSubmissionServiceForAssetSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.SubmitBucketResourceToSubmissionService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitBucketResourceToSubmissionService`: bool
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.SubmitBucketResourceToSubmissionService`: %v\n", resp)
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


## UpdateBucket

> Bucket UpdateBucket(ctx, id).Bucket(bucket).Execute()

Update Storage Buckets



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
    id := "id_example" // string | ID of bucket
    bucket := *cons3rtclient.NewBucket("CloudResourceVisibility_example", "Name_example") // Bucket | The bucket creation information

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.UpdateBucket(context.Background(), id).Bucket(bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.UpdateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBucket`: Bucket
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucket** | [**Bucket**](Bucket.md) | The bucket creation information | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToBucket

> int32 UploadFileToBucket(ctx, id).File(file).Filename(filename).Execute()

Upload File to Bucket



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
    id := "id_example" // string | ID of bucket
    file := []*os.File{"TODO"} // []*os.File |  (optional)
    filename := "filename_example" // string |  (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageApi.UploadFileToBucket(context.Background(), id).File(file).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.UploadFileToBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileToBucket`: int32
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.UploadFileToBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileToBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **[]*os.File** |  | 
 **filename** | **string** |  | 

### Return type

**int32**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

