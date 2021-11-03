# \SystemsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustedProject**](SystemsApi.md#AddTrustedProject) | **Put** /api/assets/{id}/addtrustedproject | Assign Trusted Project to Asset
[**CloneSystem**](SystemsApi.md#CloneSystem) | **Put** /api/systems/{id}/clone | Clone System
[**CreateScenario1**](SystemsApi.md#CreateScenario1) | **Put** /api/systems/{id}/createscenario | Create Scenario from System
[**CreateSoftwareBundleFromSystemModule**](SystemsApi.md#CreateSoftwareBundleFromSystemModule) | **Post** /api/systems/{id}/softwarebundle | Create Software Bundle
[**CreateSystemEntire**](SystemsApi.md#CreateSystemEntire) | **Put** /api/systems/createsystem | Create System
[**DeleteAsset**](SystemsApi.md#DeleteAsset) | **Delete** /api/assets/{id} | Delete asset
[**GetSystem**](SystemsApi.md#GetSystem) | **Get** /api/systems/{id} | Retrieve System
[**GetSystems**](SystemsApi.md#GetSystems) | **Get** /api/systems | List Systems
[**GetSystemsExpanded**](SystemsApi.md#GetSystemsExpanded) | **Get** /api/systems/expanded | List all Systems, including Project Assets
[**ListDependentAssets**](SystemsApi.md#ListDependentAssets) | **Get** /api/assets/{id}/dependent | List all Dependent Assets
[**QuickBuild1**](SystemsApi.md#QuickBuild1) | **Put** /api/systems/{id}/launch | Launch System
[**RemoveTrustedProject**](SystemsApi.md#RemoveTrustedProject) | **Put** /api/assets/{id}/removetrustedproject | Unassign Trusted Project from Asset
[**ReplaceSoftwareAssetForSoftwareComponent**](SystemsApi.md#ReplaceSoftwareAssetForSoftwareComponent) | **Put** /api/systems/{id}/replacesoftware | Replace Software Asset
[**UpdateAsset**](SystemsApi.md#UpdateAsset) | **Put** /api/assets/{id}/update | Update Asset
[**UpdateAssetState**](SystemsApi.md#UpdateAssetState) | **Put** /api/assets/{id}/updatestate | Update State
[**UpdateAssetVisibilityQuery**](SystemsApi.md#UpdateAssetVisibilityQuery) | **Put** /api/assets/{id}/updatevisibility | Update Visibility
[**UpdateImpactLevel**](SystemsApi.md#UpdateImpactLevel) | **Put** /api/assets/{id}/impactlevel | Update Impact Level
[**UpdateInstanceLimit**](SystemsApi.md#UpdateInstanceLimit) | **Put** /api/assets/{id}/limit | Update Instance Limit
[**UpdateOfflineStatus**](SystemsApi.md#UpdateOfflineStatus) | **Put** /api/assets/{id}/offline | Update Offline Status
[**UpdateSoftwareComponentForSystemModule**](SystemsApi.md#UpdateSoftwareComponentForSystemModule) | **Put** /api/systems/{id}/softwarecomponent/{componentid} | Update Software Component
[**UpdateSoftwareComponentsForSystemModule**](SystemsApi.md#UpdateSoftwareComponentsForSystemModule) | **Put** /api/systems/{id}/softwarecomponents | Update Software Components
[**UpdateSoftwareConfigurationForSystemModule**](SystemsApi.md#UpdateSoftwareConfigurationForSystemModule) | **Put** /api/systems/{id}/softwareconfiguration | Update Software Configuration
[**UpdateTemplateProfileForSystemModule**](SystemsApi.md#UpdateTemplateProfileForSystemModule) | **Put** /api/systems/{id}/templateprofile | Updates Template Profile



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
    resp, r, err := api_client.SystemsApi.AddTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.AddTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.AddTrustedProject`: %v\n", resp)
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


## CloneSystem

> string CloneSystem(ctx, id).Name(name).Execute()

Clone System



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
    id := "id_example" // string | ID of system
    name := "name_example" // string | Name of the new system

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.CloneSystem(context.Background(), id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.CloneSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.CloneSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the new system | 

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


## CreateScenario1

> bool CreateScenario1(ctx, id).Name(name).Role(role).Execute()

Create Scenario from System



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
    id := "id_example" // string | ID of system
    name := "name_example" // string | Name of the new scenario
    role := "role_example" // string | System role name in the new scenario

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.CreateScenario1(context.Background(), id).Name(name).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.CreateScenario1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScenario1`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.CreateScenario1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScenario1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the new scenario | 
 **role** | **string** | System role name in the new scenario | 

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


## CreateSoftwareBundleFromSystemModule

> string CreateSoftwareBundleFromSystemModule(ctx, id).Name(name).Body(body).Execute()

Create Software Bundle



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
    id := "id_example" // string | ID of system module
    name := "name_example" // string | Name of the new software bundle
    body := "body_example" // string | Description of the new software bundle (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.CreateSoftwareBundleFromSystemModule(context.Background(), id).Name(name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.CreateSoftwareBundleFromSystemModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSoftwareBundleFromSystemModule`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.CreateSoftwareBundleFromSystemModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system module | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoftwareBundleFromSystemModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## CreateSystemEntire

> string CreateSystemEntire(ctx).InputSystemModule(inputSystemModule).Execute()

Create System



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
    inputSystemModule := *openapiclient.NewInputSystemModule("Subtype_example") // InputSystemModule | The System to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.CreateSystemEntire(context.Background()).InputSystemModule(inputSystemModule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.CreateSystemEntire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystemEntire`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.CreateSystemEntire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemEntireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputSystemModule** | [**InputSystemModule**](InputSystemModule.md) | The System to create | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/xml, application/json
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
    resp, r, err := api_client.SystemsApi.DeleteAsset(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.DeleteAsset`: %v\n", resp)
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


## GetSystem

> FullSystemModule GetSystem(ctx, id).Execute()

Retrieve System



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
    id := "id_example" // string | ID of system

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.GetSystem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystem`: FullSystemModule
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullSystemModule**](FullSystemModule.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystems

> []MinimalSystemModule GetSystems(ctx).Type_(type_).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List Systems



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
    type_ := "type__example" // string | System module type (optional)
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.GetSystems(context.Background()).Type_(type_).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystems`: []MinimalSystemModule
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | System module type | 
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalSystemModule**](MinimalSystemModule.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemsExpanded

> []BasicSystemModule GetSystemsExpanded(ctx).Type_(type_).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Systems, including Project Assets



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
    type_ := "type__example" // string | System module type (optional)
    community := true // bool | Include community assets (optional) (default to false)
    categoryids := []int32{int32(123)} // []int32 | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.GetSystemsExpanded(context.Background()).Type_(type_).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetSystemsExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemsExpanded`: []BasicSystemModule
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetSystemsExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemsExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | System module type | 
 **community** | **bool** | Include community assets | [default to false]
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicSystemModule**](BasicSystemModule.md)

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
    resp, r, err := api_client.SystemsApi.ListDependentAssets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.ListDependentAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentAssets`: []BasicAsset
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.ListDependentAssets`: %v\n", resp)
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


## QuickBuild1

> string QuickBuild1(ctx, id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()

Launch System



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
    id := "id_example" // string | ID of system
    inputDeploymentRunOptions := *openapiclient.NewInputDeploymentRunOptions("EndState_example", "Username_example", "Password_example") // InputDeploymentRunOptions | The deployment run options to use when launching the deployment (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.QuickBuild1(context.Background(), id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.QuickBuild1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuickBuild1`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.QuickBuild1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuickBuild1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputDeploymentRunOptions** | [**InputDeploymentRunOptions**](InputDeploymentRunOptions.md) | The deployment run options to use when launching the deployment | 

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
    resp, r, err := api_client.SystemsApi.RemoveTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.RemoveTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.RemoveTrustedProject`: %v\n", resp)
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


## ReplaceSoftwareAssetForSoftwareComponent

> bool ReplaceSoftwareAssetForSoftwareComponent(ctx, id).Componentid(componentid).Assetid(assetid).Execute()

Replace Software Asset



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
    id := "id_example" // string | ID of system module
    componentid := "componentid_example" // string | ID of software component to modify
    assetid := "assetid_example" // string | ID of desired software asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.ReplaceSoftwareAssetForSoftwareComponent(context.Background(), id).Componentid(componentid).Assetid(assetid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.ReplaceSoftwareAssetForSoftwareComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSoftwareAssetForSoftwareComponent`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.ReplaceSoftwareAssetForSoftwareComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system module | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSoftwareAssetForSoftwareComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentid** | **string** | ID of software component to modify | 
 **assetid** | **string** | ID of desired software asset | 

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
    resp, r, err := api_client.SystemsApi.UpdateAsset(context.Background(), id).InputAssetForUpdate(inputAssetForUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateAsset`: %v\n", resp)
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
    resp, r, err := api_client.SystemsApi.UpdateAssetState(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateAssetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetState`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateAssetState`: %v\n", resp)
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
    resp, r, err := api_client.SystemsApi.UpdateAssetVisibilityQuery(context.Background(), id).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateAssetVisibilityQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetVisibilityQuery`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateAssetVisibilityQuery`: %v\n", resp)
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
    resp, r, err := api_client.SystemsApi.UpdateImpactLevel(context.Background(), id).Impactlevel(impactlevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateImpactLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImpactLevel`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateImpactLevel`: %v\n", resp)
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
    resp, r, err := api_client.SystemsApi.UpdateInstanceLimit(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateInstanceLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceLimit`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateInstanceLimit`: %v\n", resp)
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
    resp, r, err := api_client.SystemsApi.UpdateOfflineStatus(context.Background(), id).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateOfflineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfflineStatus`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateOfflineStatus`: %v\n", resp)
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


## UpdateSoftwareComponentForSystemModule

> bool UpdateSoftwareComponentForSystemModule(ctx, id, componentid).InputSoftwareComponent(inputSoftwareComponent).Execute()

Update Software Component



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
    id := "id_example" // string | ID of system module
    componentid := "componentid_example" // string | ID of software component to modify
    inputSoftwareComponent := *openapiclient.NewInputSoftwareComponent(*openapiclient.NewInputAsset(int32(123))) // InputSoftwareComponent | The modified Software Component

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.UpdateSoftwareComponentForSystemModule(context.Background(), id, componentid).InputSoftwareComponent(inputSoftwareComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateSoftwareComponentForSystemModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftwareComponentForSystemModule`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateSoftwareComponentForSystemModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system module | 
**componentid** | **string** | ID of software component to modify | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftwareComponentForSystemModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputSoftwareComponent** | [**InputSoftwareComponent**](InputSoftwareComponent.md) | The modified Software Component | 

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


## UpdateSoftwareComponentsForSystemModule

> bool UpdateSoftwareComponentsForSystemModule(ctx, id).InputSoftwareComponent(inputSoftwareComponent).Execute()

Update Software Components



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
    id := "id_example" // string | ID of system module
    inputSoftwareComponent := []openapiclient.InputSoftwareComponent{*openapiclient.NewInputSoftwareComponent(*openapiclient.NewInputAsset(int32(123)))} // []InputSoftwareComponent | The updated software components

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.UpdateSoftwareComponentsForSystemModule(context.Background(), id).InputSoftwareComponent(inputSoftwareComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateSoftwareComponentsForSystemModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftwareComponentsForSystemModule`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateSoftwareComponentsForSystemModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system module | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftwareComponentsForSystemModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputSoftwareComponent** | [**[]InputSoftwareComponent**](InputSoftwareComponent.md) | The updated software components | 

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


## UpdateSoftwareConfigurationForSystemModule

> bool UpdateSoftwareConfigurationForSystemModule(ctx, id).InputSoftwareComponent(inputSoftwareComponent).Execute()

Update Software Configuration



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
    id := "id_example" // string | ID of system module
    inputSoftwareComponent := []openapiclient.InputSoftwareComponent{*openapiclient.NewInputSoftwareComponent(*openapiclient.NewInputAsset(int32(123)))} // []InputSoftwareComponent | The updated software components

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.UpdateSoftwareConfigurationForSystemModule(context.Background(), id).InputSoftwareComponent(inputSoftwareComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateSoftwareConfigurationForSystemModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftwareConfigurationForSystemModule`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateSoftwareConfigurationForSystemModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system module | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftwareConfigurationForSystemModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputSoftwareComponent** | [**[]InputSoftwareComponent**](InputSoftwareComponent.md) | The updated software components | 

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


## UpdateTemplateProfileForSystemModule

> bool UpdateTemplateProfileForSystemModule(ctx, id).InputTemplateProfile(inputTemplateProfile).Execute()

Updates Template Profile



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
    id := "id_example" // string | ID of system module
    inputTemplateProfile := *openapiclient.NewInputTemplateProfile(int32(123), int32(123)) // InputTemplateProfile | The modified Template Profile

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemsApi.UpdateTemplateProfileForSystemModule(context.Background(), id).InputTemplateProfile(inputTemplateProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.UpdateTemplateProfileForSystemModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplateProfileForSystemModule`: bool
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.UpdateTemplateProfileForSystemModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of system module | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateProfileForSystemModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputTemplateProfile** | [**InputTemplateProfile**](InputTemplateProfile.md) | The modified Template Profile | 

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

