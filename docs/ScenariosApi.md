# \ScenariosApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustedProject**](ScenariosApi.md#AddTrustedProject) | **Put** /api/assets/{id}/addtrustedproject | Assign Trusted Project to Asset
[**CloneScenario**](ScenariosApi.md#CloneScenario) | **Put** /api/scenarios/{id}/clone | Clone Scenario
[**CreateScenario**](ScenariosApi.md#CreateScenario) | **Put** /api/scenarios/createscenario | Create Scenario
[**CreateSystemEntire**](ScenariosApi.md#CreateSystemEntire) | **Put** /api/systems/createsystem | Create System
[**DeleteAsset**](ScenariosApi.md#DeleteAsset) | **Delete** /api/assets/{id} | Delete asset
[**GetBindingsForDeployment1**](ScenariosApi.md#GetBindingsForDeployment1) | **Get** /api/scenarios/{id}/bindings | List Bindings
[**GetScenario**](ScenariosApi.md#GetScenario) | **Get** /api/scenarios/{id} | Retrieve Scenario
[**GetScenarios**](ScenariosApi.md#GetScenarios) | **Get** /api/scenarios | List Scenarios
[**GetScenariosExpanded**](ScenariosApi.md#GetScenariosExpanded) | **Get** /api/scenarios/expanded | List all Scenarios, including Project Assets
[**ListDependentAssets**](ScenariosApi.md#ListDependentAssets) | **Get** /api/assets/{id}/dependent | List all Dependent Assets
[**PublishScenarioToComposition**](ScenariosApi.md#PublishScenarioToComposition) | **Post** /api/scenarios/{id}/publish | Publish Scenario
[**QuickBuild**](ScenariosApi.md#QuickBuild) | **Put** /api/scenarios/{id}/launch | Launch Scenario
[**QuickBuild1**](ScenariosApi.md#QuickBuild1) | **Put** /api/systems/{id}/launch | Launch System
[**RemoveTrustedProject**](ScenariosApi.md#RemoveTrustedProject) | **Put** /api/assets/{id}/removetrustedproject | Unassign Trusted Project from Asset
[**UpdateAsset**](ScenariosApi.md#UpdateAsset) | **Put** /api/assets/{id}/update | Update Asset
[**UpdateAssetState**](ScenariosApi.md#UpdateAssetState) | **Put** /api/assets/{id}/updatestate | Update State
[**UpdateAssetVisibilityQuery**](ScenariosApi.md#UpdateAssetVisibilityQuery) | **Put** /api/assets/{id}/updatevisibility | Update Visibility
[**UpdateImpactLevel**](ScenariosApi.md#UpdateImpactLevel) | **Put** /api/assets/{id}/impactlevel | Update Impact Level
[**UpdateInstanceLimit**](ScenariosApi.md#UpdateInstanceLimit) | **Put** /api/assets/{id}/limit | Update Instance Limit
[**UpdateOfflineStatus**](ScenariosApi.md#UpdateOfflineStatus) | **Put** /api/assets/{id}/offline | Update Offline Status



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
    resp, r, err := api_client.ScenariosApi.AddTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.AddTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.AddTrustedProject`: %v\n", resp)
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


## CloneScenario

> string CloneScenario(ctx, id).Name(name).Execute()

Clone Scenario



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
    id := "id_example" // string | ID of scenario
    name := "name_example" // string | Name of the new scenario

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScenariosApi.CloneScenario(context.Background(), id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.CloneScenario``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneScenario`: string
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.CloneScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of scenario | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the new scenario | 

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


## CreateScenario

> string CreateScenario(ctx).InputScenario(inputScenario).Execute()

Create Scenario



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
    inputScenario := *openapiclient.NewInputScenario() // InputScenario | The Scenario to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScenariosApi.CreateScenario(context.Background()).InputScenario(inputScenario).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.CreateScenario``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScenario`: string
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.CreateScenario`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputScenario** | [**InputScenario**](InputScenario.md) | The Scenario to create | 

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
    resp, r, err := api_client.ScenariosApi.CreateSystemEntire(context.Background()).InputSystemModule(inputSystemModule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.CreateSystemEntire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystemEntire`: string
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.CreateSystemEntire`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.DeleteAsset(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.DeleteAsset`: %v\n", resp)
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


## GetBindingsForDeployment1

> []VirtualizationRealmBinding GetBindingsForDeployment1(ctx, id).VirtualizationRealmId(virtualizationRealmId).Execute()

List Bindings



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
    id := "id_example" // string | ID of scenario
    virtualizationRealmId := int32(56) // int32 | ID of preferred virtualization realm (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScenariosApi.GetBindingsForDeployment1(context.Background(), id).VirtualizationRealmId(virtualizationRealmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.GetBindingsForDeployment1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBindingsForDeployment1`: []VirtualizationRealmBinding
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.GetBindingsForDeployment1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of scenario | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBindingsForDeployment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationRealmId** | **int32** | ID of preferred virtualization realm | 

### Return type

[**[]VirtualizationRealmBinding**](VirtualizationRealmBinding.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenario

> FullScenario GetScenario(ctx, id).Execute()

Retrieve Scenario



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
    id := "id_example" // string | ID of scenario

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScenariosApi.GetScenario(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.GetScenario``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScenario`: FullScenario
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.GetScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of scenario | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullScenario**](FullScenario.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenarios

> []MinimalScenario GetScenarios(ctx).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List Scenarios



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
    resp, r, err := api_client.ScenariosApi.GetScenarios(context.Background()).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.GetScenarios``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScenarios`: []MinimalScenario
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.GetScenarios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScenariosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalScenario**](MinimalScenario.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenariosExpanded

> []BasicScenario GetScenariosExpanded(ctx).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Scenarios, including Project Assets



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
    resp, r, err := api_client.ScenariosApi.GetScenariosExpanded(context.Background()).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.GetScenariosExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScenariosExpanded`: []BasicScenario
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.GetScenariosExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScenariosExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **community** | **bool** | Include community assets | [default to false]
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicScenario**](BasicScenario.md)

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
    resp, r, err := api_client.ScenariosApi.ListDependentAssets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.ListDependentAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentAssets`: []BasicAsset
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.ListDependentAssets`: %v\n", resp)
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


## PublishScenarioToComposition

> string PublishScenarioToComposition(ctx, id).InputComposition(inputComposition).Execute()

Publish Scenario



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
    inputComposition := *openapiclient.NewInputComposition("Name_example", *openapiclient.NewInputCompositionRunOptions(int32(123))) // InputComposition | The composition definition used when launching the deployment (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScenariosApi.PublishScenarioToComposition(context.Background(), id).InputComposition(inputComposition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.PublishScenarioToComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishScenarioToComposition`: string
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.PublishScenarioToComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishScenarioToCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputComposition** | [**InputComposition**](InputComposition.md) | The composition definition used when launching the deployment | 

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


## QuickBuild

> string QuickBuild(ctx, id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()

Launch Scenario



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
    id := "id_example" // string | ID of scenario
    inputDeploymentRunOptions := *openapiclient.NewInputDeploymentRunOptions("EndState_example", "Username_example", "Password_example") // InputDeploymentRunOptions | The deployment run options to use when launching the deployment (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScenariosApi.QuickBuild(context.Background(), id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.QuickBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuickBuild`: string
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.QuickBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of scenario | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuickBuildRequest struct via the builder pattern


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
    resp, r, err := api_client.ScenariosApi.QuickBuild1(context.Background(), id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.QuickBuild1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuickBuild1`: string
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.QuickBuild1`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.RemoveTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.RemoveTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.RemoveTrustedProject`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.UpdateAsset(context.Background(), id).InputAssetForUpdate(inputAssetForUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.UpdateAsset`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.UpdateAssetState(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.UpdateAssetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetState`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.UpdateAssetState`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.UpdateAssetVisibilityQuery(context.Background(), id).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.UpdateAssetVisibilityQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetVisibilityQuery`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.UpdateAssetVisibilityQuery`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.UpdateImpactLevel(context.Background(), id).Impactlevel(impactlevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.UpdateImpactLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImpactLevel`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.UpdateImpactLevel`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.UpdateInstanceLimit(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.UpdateInstanceLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceLimit`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.UpdateInstanceLimit`: %v\n", resp)
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
    resp, r, err := api_client.ScenariosApi.UpdateOfflineStatus(context.Background(), id).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScenariosApi.UpdateOfflineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfflineStatus`: bool
    fmt.Fprintf(os.Stdout, "Response from `ScenariosApi.UpdateOfflineStatus`: %v\n", resp)
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

