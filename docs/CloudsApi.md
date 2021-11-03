# \CloudsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProject**](CloudsApi.md#AddProject) | **Put** /api/virtualizationrealms/{id}/projects | Assign Project
[**AllocateVirtualizationRealm**](CloudsApi.md#AllocateVirtualizationRealm) | **Post** /api/clouds/{id}/virtualizationrealms/allocate | Allocate Virtualization Realm
[**AssignManagingTeam**](CloudsApi.md#AssignManagingTeam) | **Put** /api/clouds/{id}/virtualizationrealms/{virtualizationRealmId}/team | Assign Virtualization Realm-managing Team
[**DeallocateVirtRealm**](CloudsApi.md#DeallocateVirtRealm) | **Delete** /api/clouds/{id}/virtualizationrealms/allocate | De-allocate Virtualization Realm
[**DeleteCloud**](CloudsApi.md#DeleteCloud) | **Delete** /api/clouds/{id} | Delete Cloud
[**EnableMaintenceMode**](CloudsApi.md#EnableMaintenceMode) | **Put** /api/clouds/{id}/maintenance | Update Maintenance Mode
[**GetCloud**](CloudsApi.md#GetCloud) | **Get** /api/clouds/{id} | Retrieve Cloud
[**GetCloudResources**](CloudsApi.md#GetCloudResources) | **Get** /api/clouds/{id}/resources | Retrieve Cloud Resources
[**GetClouds**](CloudsApi.md#GetClouds) | **Get** /api/clouds | List Clouds
[**GetDefaultNetwork**](CloudsApi.md#GetDefaultNetwork) | **Get** /api/clouds/defaultnetwork | Retrieve Default Network
[**GetEdgeGatewayIPs**](CloudsApi.md#GetEdgeGatewayIPs) | **Get** /api/clouds/{id}/edgegateways | List Edge Gateways
[**ListVirtRealmsInCloud**](CloudsApi.md#ListVirtRealmsInCloud) | **Get** /api/clouds/{id}/virtualizationrealms | List Virtualization Realms
[**RegisterCloud**](CloudsApi.md#RegisterCloud) | **Post** /api/clouds | Create Cloud
[**RegisterVirtualizationRealm**](CloudsApi.md#RegisterVirtualizationRealm) | **Post** /api/clouds/{id}/virtualizationrealms | Register Virtualization Realm
[**RemoveVirtRealm**](CloudsApi.md#RemoveVirtRealm) | **Delete** /api/clouds/{id}/virtualizationrealms | Unregister Virtualization Realm
[**UnassignManagingTeam**](CloudsApi.md#UnassignManagingTeam) | **Delete** /api/clouds/{id}/virtualizationrealms/{virtualizationRealmId}/team | Unassign Manager from Team
[**UpdateCloud**](CloudsApi.md#UpdateCloud) | **Put** /api/clouds/{id} | Update Cloud Content
[**UpdateVirtRealm**](CloudsApi.md#UpdateVirtRealm) | **Put** /api/clouds/{id}/virtualizationrealms | Update Virtualization Realm
[**UpdateVirtualizationRealm**](CloudsApi.md#UpdateVirtualizationRealm) | **Put** /api/virtualizationrealms/{id} | Update Virtualization Realm
[**UpdateVirtualizationRealmMaximumImpactLevel**](CloudsApi.md#UpdateVirtualizationRealmMaximumImpactLevel) | **Put** /api/clouds/{id}/impactlevel | Update Impact Level



## AddProject

> int32 AddProject(ctx, id).ProjectId(projectId).Execute()

Assign Project



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
    id := "id_example" // string | ID of virtualization realm
    projectId := "projectId_example" // string | ID of project to assign

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.AddProject(context.Background(), id).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.AddProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProject`: int32
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.AddProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | ID of project to assign | 

### Return type

**int32**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllocateVirtualizationRealm

> MinimalVirtualizationRealm AllocateVirtualizationRealm(ctx, id).AbstractCloudSpaceRequest(abstractCloudSpaceRequest).Execute()

Allocate Virtualization Realm



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
    id := "id_example" // string | ID of cloud
    abstractCloudSpaceRequest := *openapiclient.NewAbstractCloudSpaceRequest("CloudSpaceName_example", "Cidr_example", "Subtype_example") // AbstractCloudSpaceRequest | The virtualization realm allocation information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.AllocateVirtualizationRealm(context.Background(), id).AbstractCloudSpaceRequest(abstractCloudSpaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.AllocateVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllocateVirtualizationRealm`: MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.AllocateVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllocateVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **abstractCloudSpaceRequest** | [**AbstractCloudSpaceRequest**](AbstractCloudSpaceRequest.md) | The virtualization realm allocation information | 

### Return type

[**MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignManagingTeam

> bool AssignManagingTeam(ctx, id, virtualizationRealmId).TeamId(teamId).Execute()

Assign Virtualization Realm-managing Team



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
    id := "id_example" // string | ID of cloud
    virtualizationRealmId := "virtualizationRealmId_example" // string | ID of virtualization realm
    teamId := int32(56) // int32 | ID of team to assign

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.AssignManagingTeam(context.Background(), id, virtualizationRealmId).TeamId(teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.AssignManagingTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignManagingTeam`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.AssignManagingTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 
**virtualizationRealmId** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignManagingTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamId** | **int32** | ID of team to assign | 

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


## DeallocateVirtRealm

> bool DeallocateVirtRealm(ctx, id).VirtRealmId(virtRealmId).Execute()

De-allocate Virtualization Realm



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
    id := "id_example" // string | ID of cloud
    virtRealmId := "virtRealmId_example" // string | ID of virtualization realm

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.DeallocateVirtRealm(context.Background(), id).VirtRealmId(virtRealmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.DeallocateVirtRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeallocateVirtRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.DeallocateVirtRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeallocateVirtRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtRealmId** | **string** | ID of virtualization realm | 

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


## DeleteCloud

> bool DeleteCloud(ctx, id).Execute()

Delete Cloud



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
    id := "id_example" // string | ID of cloud

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.DeleteCloud(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.DeleteCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCloud`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.DeleteCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudRequest struct via the builder pattern


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


## EnableMaintenceMode

> bool EnableMaintenceMode(ctx, id).Enable(enable).MaintenanceModeRequest(maintenanceModeRequest).Execute()

Update Maintenance Mode



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
    id := "id_example" // string | ID of cloud
    enable := true // bool | Enable or disable maintenance mode (optional)
    maintenanceModeRequest := *openapiclient.NewMaintenanceModeRequest() // MaintenanceModeRequest | The maintenance mode request, when enabling maintenance mode (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.EnableMaintenceMode(context.Background(), id).Enable(enable).MaintenanceModeRequest(maintenanceModeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.EnableMaintenceMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableMaintenceMode`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.EnableMaintenceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableMaintenceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enable** | **bool** | Enable or disable maintenance mode | 
 **maintenanceModeRequest** | [**MaintenanceModeRequest**](MaintenanceModeRequest.md) | The maintenance mode request, when enabling maintenance mode | 

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


## GetCloud

> FullCloud GetCloud(ctx, id).Execute()

Retrieve Cloud



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
    id := "id_example" // string | ID of cloud

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetCloud(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloud`: FullCloud
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullCloud**](FullCloud.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudResources

> AbstractCloudResources GetCloudResources(ctx, id).Execute()

Retrieve Cloud Resources



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
    id := "id_example" // string | ID of cloud

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetCloudResources(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudResources`: AbstractCloudResources
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetCloudResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbstractCloudResources**](AbstractCloudResources.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClouds

> []MinimalCloud GetClouds(ctx).Maxresults(maxresults).Page(page).Execute()

List Clouds



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
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetClouds(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClouds`: []MinimalCloud
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetClouds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudsRequest struct via the builder pattern


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


## GetDefaultNetwork

> Network GetDefaultNetwork(ctx).Execute()

Retrieve Default Network



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
    resp, r, err := api_client.CloudsApi.GetDefaultNetwork(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetDefaultNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultNetwork`: Network
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetDefaultNetwork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultNetworkRequest struct via the builder pattern


### Return type

[**Network**](Network.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEdgeGatewayIPs

> []string GetEdgeGatewayIPs(ctx, id).Execute()

List Edge Gateways



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
    id := "id_example" // string | ID of cloud

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetEdgeGatewayIPs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetEdgeGatewayIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeGatewayIPs`: []string
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetEdgeGatewayIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeGatewayIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtRealmsInCloud

> []MinimalVirtualizationRealm ListVirtRealmsInCloud(ctx, id).Maxresults(maxresults).Page(page).Execute()

List Virtualization Realms



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
    id := "id_example" // string | ID of cloud
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.ListVirtRealmsInCloud(context.Background(), id).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListVirtRealmsInCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtRealmsInCloud`: []MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListVirtRealmsInCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtRealmsInCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalVirtualizationRealm**](MinimalVirtualizationRealm.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCloud

> string RegisterCloud(ctx).CloudATOConsent(cloudATOConsent).InputCloud(inputCloud).Execute()

Create Cloud



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
    cloudATOConsent := true // bool | Cloud ATO consent (optional) (default to false)
    inputCloud := *openapiclient.NewInputCloud("Name_example", "ExternalIpSource_example", "MaximumImpactLevel_example", *openapiclient.NewInputTeam(int32(123)), "Subtype_example") // InputCloud | The Cloud to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.RegisterCloud(context.Background()).CloudATOConsent(cloudATOConsent).InputCloud(inputCloud).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.RegisterCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterCloud`: string
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.RegisterCloud`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudATOConsent** | **bool** | Cloud ATO consent | [default to false]
 **inputCloud** | [**InputCloud**](InputCloud.md) | The Cloud to create | 

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


## RegisterVirtualizationRealm

> int32 RegisterVirtualizationRealm(ctx, id).AbstractRegisterCloudSpaceRequest(abstractRegisterCloudSpaceRequest).Execute()

Register Virtualization Realm



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
    id := "id_example" // string | ID of cloud
    abstractRegisterCloudSpaceRequest := *openapiclient.NewAbstractRegisterCloudSpaceRequest("VirtualizationRealmType_example", "Name_example", "Cons3rtNetworkName_example", "Password_example", "PrimaryNetworkName_example", "Username_example") // AbstractRegisterCloudSpaceRequest | The virtualization realm registration information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.RegisterVirtualizationRealm(context.Background(), id).AbstractRegisterCloudSpaceRequest(abstractRegisterCloudSpaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.RegisterVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterVirtualizationRealm`: int32
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.RegisterVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **abstractRegisterCloudSpaceRequest** | [**AbstractRegisterCloudSpaceRequest**](AbstractRegisterCloudSpaceRequest.md) | The virtualization realm registration information | 

### Return type

**int32**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVirtRealm

> bool RemoveVirtRealm(ctx, id).VirtRealmId(virtRealmId).Execute()

Unregister Virtualization Realm



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
    id := "id_example" // string | ID of cloud
    virtRealmId := "virtRealmId_example" // string | ID of virtualization realm

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.RemoveVirtRealm(context.Background(), id).VirtRealmId(virtRealmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.RemoveVirtRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVirtRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.RemoveVirtRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVirtRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtRealmId** | **string** | ID of virtualization realm | 

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


## UnassignManagingTeam

> bool UnassignManagingTeam(ctx, id, virtualizationRealmId).TeamId(teamId).Execute()

Unassign Manager from Team



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
    id := "id_example" // string | ID of cloud
    virtualizationRealmId := "virtualizationRealmId_example" // string | ID of virtualization realm
    teamId := int32(56) // int32 | ID of team to unassign

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UnassignManagingTeam(context.Background(), id, virtualizationRealmId).TeamId(teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UnassignManagingTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignManagingTeam`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UnassignManagingTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 
**virtualizationRealmId** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignManagingTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamId** | **int32** | ID of team to unassign | 

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


## UpdateCloud

> bool UpdateCloud(ctx, id).InputCloud(inputCloud).Execute()

Update Cloud Content



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
    id := "id_example" // string | ID of cloud
    inputCloud := *openapiclient.NewInputCloud("Name_example", "ExternalIpSource_example", "MaximumImpactLevel_example", *openapiclient.NewInputTeam(int32(123)), "Subtype_example") // InputCloud | The modified Cloud data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateCloud(context.Background(), id).InputCloud(inputCloud).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloud`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCloud** | [**InputCloud**](InputCloud.md) | The modified Cloud data | 

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


## UpdateVirtRealm

> bool UpdateVirtRealm(ctx, id).VirtRealmId(virtRealmId).InputVirtualizationRealm(inputVirtualizationRealm).Execute()

Update Virtualization Realm



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
    id := "id_example" // string | ID of cloud
    virtRealmId := "virtRealmId_example" // string | ID of virtualization realm
    inputVirtualizationRealm := *openapiclient.NewInputVirtualizationRealm("AccountId_example", "Cidr_example", "Description_example", "Name_example", "Password_example", "Username_example") // InputVirtualizationRealm | The modified virtualization realm information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateVirtRealm(context.Background(), id).VirtRealmId(virtRealmId).InputVirtualizationRealm(inputVirtualizationRealm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateVirtRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateVirtRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtRealmId** | **string** | ID of virtualization realm | 
 **inputVirtualizationRealm** | [**InputVirtualizationRealm**](InputVirtualizationRealm.md) | The modified virtualization realm information | 

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


## UpdateVirtualizationRealm

> bool UpdateVirtualizationRealm(ctx, id).InputVRAdminVirtualizationRealm(inputVRAdminVirtualizationRealm).Execute()

Update Virtualization Realm



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
    id := "id_example" // string | ID of virtualization realm
    inputVRAdminVirtualizationRealm := *openapiclient.NewInputVRAdminVirtualizationRealm("VirtualizationRealmType_example", "Cidr_example", "Description_example", "Name_example") // InputVRAdminVirtualizationRealm | The updated Virtualization Realm data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateVirtualizationRealm(context.Background(), id).InputVRAdminVirtualizationRealm(inputVRAdminVirtualizationRealm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputVRAdminVirtualizationRealm** | [**InputVRAdminVirtualizationRealm**](InputVRAdminVirtualizationRealm.md) | The updated Virtualization Realm data | 

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


## UpdateVirtualizationRealmMaximumImpactLevel

> bool UpdateVirtualizationRealmMaximumImpactLevel(ctx, id).Maximumimpactlevel(maximumimpactlevel).Execute()

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
    id := "id_example" // string | ID of cloud
    maximumimpactlevel := "maximumimpactlevel_example" // string | The maximum impact level type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateVirtualizationRealmMaximumImpactLevel(context.Background(), id).Maximumimpactlevel(maximumimpactlevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateVirtualizationRealmMaximumImpactLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationRealmMaximumImpactLevel`: bool
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateVirtualizationRealmMaximumImpactLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationRealmMaximumImpactLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maximumimpactlevel** | **string** | The maximum impact level type. | 

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

