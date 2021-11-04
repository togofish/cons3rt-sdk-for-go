# \DownloadApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadFileFromBucket**](DownloadApi.md#DownloadFileFromBucket) | **Get** /api/buckets/{id}/download | Download File From Bucket



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
    resp, r, err := api_client.DownloadApi.DownloadFileFromBucket(context.Background(), id).FileName(fileName).Background(background).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadApi.DownloadFileFromBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFileFromBucket`: bool
    fmt.Fprintf(os.Stdout, "Response from `DownloadApi.DownloadFileFromBucket`: %v\n", resp)
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

