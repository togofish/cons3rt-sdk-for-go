# \TemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVirtualizationRealmTemplates**](TemplatesApi.md#ListVirtualizationRealmTemplates) | **Get** /api/templates | List all Templates



## ListVirtualizationRealmTemplates

> []MinimalCons3rtTemplateData ListVirtualizationRealmTemplates(ctx).VirtualizationRealmId(virtualizationRealmId).IncludeRegistrations(includeRegistrations).IncludeSubscriptions(includeSubscriptions).Execute()

List all Templates



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
    virtualizationRealmId := int32(56) // int32 | Virtualization Realm ID to filter by
    includeRegistrations := true // bool | Include templates registered to this virtualization realm (optional) (default to true)
    includeSubscriptions := true // bool | Include templates this virtualization realm is subscribed to (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.ListVirtualizationRealmTemplates(context.Background()).VirtualizationRealmId(virtualizationRealmId).IncludeRegistrations(includeRegistrations).IncludeSubscriptions(includeSubscriptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.ListVirtualizationRealmTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualizationRealmTemplates`: []MinimalCons3rtTemplateData
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.ListVirtualizationRealmTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualizationRealmTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualizationRealmId** | **int32** | Virtualization Realm ID to filter by | 
 **includeRegistrations** | **bool** | Include templates registered to this virtualization realm | [default to true]
 **includeSubscriptions** | **bool** | Include templates this virtualization realm is subscribed to | [default to false]

### Return type

[**[]MinimalCons3rtTemplateData**](MinimalCons3rtTemplateData.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

