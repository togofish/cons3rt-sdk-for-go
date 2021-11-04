# \ImportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadFile**](ImportApi.md#UploadFile) | **Post** /api/import | Import a New Asset



## UploadFile

> int32 UploadFile(ctx).File(file).Filename(filename).Execute()

Import a New Asset



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
    file := []*os.File{"TODO"} // []*os.File |  (optional)
    filename := "filename_example" // string |  (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportApi.UploadFile(context.Background()).File(file).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: int32
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


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

