# \DeploymentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecurringSchedule**](DeploymentsApi.md#AddRecurringSchedule) | **Put** /api/deployments/{id}/schedules | Add Recurring Schedule
[**AddTrustedProject**](DeploymentsApi.md#AddTrustedProject) | **Put** /api/assets/{id}/addtrustedproject | Assign Trusted Project to Asset
[**CloneDeployment**](DeploymentsApi.md#CloneDeployment) | **Put** /api/deployments/{id}/clone | Clone Deployment
[**CloneSystem**](DeploymentsApi.md#CloneSystem) | **Put** /api/systems/{id}/clone | Clone System
[**CreateDeploymentEntire**](DeploymentsApi.md#CreateDeploymentEntire) | **Put** /api/deployments/createdeployment | Create Deployment
[**DeleteAsset**](DeploymentsApi.md#DeleteAsset) | **Delete** /api/assets/{id} | Delete asset
[**DetermineValidVirtualizationRealms**](DeploymentsApi.md#DetermineValidVirtualizationRealms) | **Get** /api/deployments/{id}/validrealms | List Valid Virtualization Realms
[**GetBindingsForDeployment**](DeploymentsApi.md#GetBindingsForDeployment) | **Post** /api/deployments/{id}/bindings | List Bindings
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /api/deployments/{id} | Retrieve Deployment
[**GetDeploymentMetric**](DeploymentsApi.md#GetDeploymentMetric) | **Get** /api/deployments/{id}/metrics | Retrieve Metrics
[**GetDeploymentRuns**](DeploymentsApi.md#GetDeploymentRuns) | **Get** /api/deployments/{id}/runs | List Deployment Runs
[**GetDeploymentRuns1**](DeploymentsApi.md#GetDeploymentRuns1) | **Get** /api/drs | List Deployment Runs
[**GetDeploymentRunsInVirtualizationRealm**](DeploymentsApi.md#GetDeploymentRunsInVirtualizationRealm) | **Get** /api/virtualizationrealms/{id}/deploymentruns | List Deployment Runs
[**GetDeployments**](DeploymentsApi.md#GetDeployments) | **Get** /api/deployments | List Deployments
[**GetDeploymentsExpanded**](DeploymentsApi.md#GetDeploymentsExpanded) | **Get** /api/deployments/expanded | List all Deployments, including Project Assets
[**GetValidNetworksForVirtualizationRealm**](DeploymentsApi.md#GetValidNetworksForVirtualizationRealm) | **Get** /api/deployments/{id}/networks/{virtualizationRealmId} | List Virtualization Realm Networks
[**LaunchDeployment**](DeploymentsApi.md#LaunchDeployment) | **Post** /api/deployments/{id}/launch | Launch Deployment
[**ListDependentAssets**](DeploymentsApi.md#ListDependentAssets) | **Get** /api/assets/{id}/dependent | List all Dependent Assets
[**ListOptions**](DeploymentsApi.md#ListOptions) | **Get** /api/deployments/{id}/options | List Run Options
[**ListProperties**](DeploymentsApi.md#ListProperties) | **Get** /api/deployments/{id}/properties | List Properties
[**ListSchedules**](DeploymentsApi.md#ListSchedules) | **Get** /api/deployments/{id}/schedules | List Recurring Schedules
[**RemoveRecurringSchedule**](DeploymentsApi.md#RemoveRecurringSchedule) | **Delete** /api/deployments/{id}/schedules | Remove Recurring Schedule
[**RemoveTrustedProject**](DeploymentsApi.md#RemoveTrustedProject) | **Put** /api/assets/{id}/removetrustedproject | Unassign Trusted Project from Asset
[**UpdateAsset**](DeploymentsApi.md#UpdateAsset) | **Put** /api/assets/{id}/update | Update Asset
[**UpdateAssetState**](DeploymentsApi.md#UpdateAssetState) | **Put** /api/assets/{id}/updatestate | Update State
[**UpdateAssetVisibilityQuery**](DeploymentsApi.md#UpdateAssetVisibilityQuery) | **Put** /api/assets/{id}/updatevisibility | Update Visibility
[**UpdateImpactLevel**](DeploymentsApi.md#UpdateImpactLevel) | **Put** /api/assets/{id}/impactlevel | Update Impact Level
[**UpdateInstanceLimit**](DeploymentsApi.md#UpdateInstanceLimit) | **Put** /api/assets/{id}/limit | Update Instance Limit
[**UpdateOfflineStatus**](DeploymentsApi.md#UpdateOfflineStatus) | **Put** /api/assets/{id}/offline | Update Offline Status



## AddRecurringSchedule

> bool AddRecurringSchedule(ctx, id).InputRecurringSchedule(inputRecurringSchedule).Execute()

Add Recurring Schedule



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
    id := "id_example" // string | ID of deployment
    inputRecurringSchedule := *cons3rtclient.NewInputRecurringSchedule("Timezone_example", "Schedule_example", *cons3rtclient.NewInputDeploymentRunOptions("EndState_example", "Username_example", "Password_example")) // InputRecurringSchedule | The Recurring Schedule definition (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.AddRecurringSchedule(context.Background(), id).InputRecurringSchedule(inputRecurringSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.AddRecurringSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRecurringSchedule`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.AddRecurringSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRecurringScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputRecurringSchedule** | [**InputRecurringSchedule**](InputRecurringSchedule.md) | The Recurring Schedule definition | 

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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset
    trustedid := "trustedid_example" // string | ID of project to trust

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.AddTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.AddTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.AddTrustedProject`: %v\n", resp)
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


## CloneDeployment

> string CloneDeployment(ctx, id).Name(name).Execute()

Clone Deployment



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
    id := "id_example" // string | ID of deployment
    name := "name_example" // string | Name of the new deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.CloneDeployment(context.Background(), id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CloneDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDeployment`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CloneDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the new deployment | 

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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of system
    name := "name_example" // string | Name of the new system

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.CloneSystem(context.Background(), id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CloneSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneSystem`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CloneSystem`: %v\n", resp)
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


## CreateDeploymentEntire

> string CreateDeploymentEntire(ctx).InputDeployment(inputDeployment).Execute()

Create Deployment



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
    inputDeployment := *cons3rtclient.NewInputDeployment() // InputDeployment | The Deployment to create (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.CreateDeploymentEntire(context.Background()).InputDeployment(inputDeployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CreateDeploymentEntire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeploymentEntire`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CreateDeploymentEntire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentEntireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputDeployment** | [**InputDeployment**](InputDeployment.md) | The Deployment to create | 

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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset to delete
    force := true // bool | Allow delete if there are dependent assets (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.DeleteAsset(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.DeleteAsset`: %v\n", resp)
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


## DetermineValidVirtualizationRealms

> []MinimalVirtualizationRealm DetermineValidVirtualizationRealms(ctx, id).Execute()

List Valid Virtualization Realms



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.DetermineValidVirtualizationRealms(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DetermineValidVirtualizationRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetermineValidVirtualizationRealms`: []MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.DetermineValidVirtualizationRealms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetermineValidVirtualizationRealmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetBindingsForDeployment

> []VirtualizationRealmBinding GetBindingsForDeployment(ctx, id).InputDeploymentRunOptionsForBindings(inputDeploymentRunOptionsForBindings).Execute()

List Bindings



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
    id := "id_example" // string | ID of deployment
    inputDeploymentRunOptionsForBindings := *cons3rtclient.NewInputDeploymentRunOptionsForBindings() // InputDeploymentRunOptionsForBindings | The deployment run options to use when launching the deployment (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetBindingsForDeployment(context.Background(), id).InputDeploymentRunOptionsForBindings(inputDeploymentRunOptionsForBindings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetBindingsForDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBindingsForDeployment`: []VirtualizationRealmBinding
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetBindingsForDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBindingsForDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputDeploymentRunOptionsForBindings** | [**InputDeploymentRunOptionsForBindings**](InputDeploymentRunOptionsForBindings.md) | The deployment run options to use when launching the deployment | 

### Return type

[**[]VirtualizationRealmBinding**](VirtualizationRealmBinding.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployment

> FullDeployment GetDeployment(ctx, id).Execute()

Retrieve Deployment



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeployment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployment`: FullDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullDeployment**](FullDeployment.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentMetric

> DeploymentAssetMetric GetDeploymentMetric(ctx, id).Execute()

Retrieve Metrics



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentMetric(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentMetric`: DeploymentAssetMetric
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentAssetMetric**](DeploymentAssetMetric.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRuns

> []MinimalDeploymentRun GetDeploymentRuns(ctx, id).Maxresults(maxresults).Page(page).Execute()

List Deployment Runs



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
    id := "id_example" // string | ID of deployment
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentRuns(context.Background(), id).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRuns`: []MinimalDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalDeploymentRun**](MinimalDeploymentRun.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRuns1

> []MinimalDeploymentRun GetDeploymentRuns1(ctx).SearchType(searchType).InProject(inProject).Maxresults(maxresults).Page(page).Execute()

List Deployment Runs



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
    searchType := "searchType_example" // string | Deployment run status
    inProject := true // bool | Include project runs (optional) (default to false)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentRuns1(context.Background()).SearchType(searchType).InProject(inProject).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentRuns1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRuns1`: []MinimalDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentRuns1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRuns1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchType** | **string** | Deployment run status | 
 **inProject** | **bool** | Include project runs | [default to false]
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalDeploymentRun**](MinimalDeploymentRun.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRunsInVirtualizationRealm

> []MinimalDeploymentRun GetDeploymentRunsInVirtualizationRealm(ctx, id).SearchType(searchType).Maxresults(maxresults).Page(page).Execute()

List Deployment Runs



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
    id := "id_example" // string | ID of virtualization realm
    searchType := "searchType_example" // string | Deployment run status type (default to "SEARCH_ALL")
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentRunsInVirtualizationRealm(context.Background(), id).SearchType(searchType).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentRunsInVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRunsInVirtualizationRealm`: []MinimalDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentRunsInVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRunsInVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchType** | **string** | Deployment run status type | [default to &quot;SEARCH_ALL&quot;]
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalDeploymentRun**](MinimalDeploymentRun.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployments

> []MinimalDeployment GetDeployments(ctx).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List Deployments



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
    resp, r, err := api_client.DeploymentsApi.GetDeployments(context.Background()).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployments`: []MinimalDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalDeployment**](MinimalDeployment.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentsExpanded

> []BasicDeployment GetDeploymentsExpanded(ctx).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()

List all Deployments, including Project Assets



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
    resp, r, err := api_client.DeploymentsApi.GetDeploymentsExpanded(context.Background()).Community(community).Categoryids(categoryids).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentsExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentsExpanded`: []BasicDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentsExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **community** | **bool** | Include community assets | [default to false]
 **categoryids** | **[]int32** | Category ID(s) to filter by. Multiple categories can be provided with comma-separated strings. | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicDeployment**](BasicDeployment.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidNetworksForVirtualizationRealm

> []MinimalNetwork GetValidNetworksForVirtualizationRealm(ctx, id, virtualizationRealmId).Execute()

List Virtualization Realm Networks



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
    id := "id_example" // string | ID of deployment
    virtualizationRealmId := "virtualizationRealmId_example" // string | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetValidNetworksForVirtualizationRealm(context.Background(), id, virtualizationRealmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetValidNetworksForVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidNetworksForVirtualizationRealm`: []MinimalNetwork
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetValidNetworksForVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 
**virtualizationRealmId** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidNetworksForVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MinimalNetwork**](MinimalNetwork.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchDeployment

> SubmitDeploymentRunRequestReturnType LaunchDeployment(ctx, id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()

Launch Deployment



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
    id := "id_example" // string | ID of deployment
    inputDeploymentRunOptions := *cons3rtclient.NewInputDeploymentRunOptions("EndState_example", "Username_example", "Password_example") // InputDeploymentRunOptions | The deployment run options to use when launching the deployment (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.LaunchDeployment(context.Background(), id).InputDeploymentRunOptions(inputDeploymentRunOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.LaunchDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LaunchDeployment`: SubmitDeploymentRunRequestReturnType
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.LaunchDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputDeploymentRunOptions** | [**InputDeploymentRunOptions**](InputDeploymentRunOptions.md) | The deployment run options to use when launching the deployment | 

### Return type

[**SubmitDeploymentRunRequestReturnType**](SubmitDeploymentRunRequestReturnType.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/xml, application/json
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListDependentAssets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListDependentAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentAssets`: []BasicAsset
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListDependentAssets`: %v\n", resp)
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


## ListOptions

> []MinimalDeploymentRunOptions ListOptions(ctx, id).Execute()

List Run Options



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListOptions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOptions`: []MinimalDeploymentRunOptions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MinimalDeploymentRunOptions**](MinimalDeploymentRunOptions.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProperties

> []Property ListProperties(ctx, id).Execute()

List Properties



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListProperties(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProperties`: []Property
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Property**](Property.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> []MinimalRecurringSchedule ListSchedules(ctx, id).Execute()

List Recurring Schedules



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListSchedules(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedules`: []MinimalRecurringSchedule
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MinimalRecurringSchedule**](MinimalRecurringSchedule.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRecurringSchedule

> bool RemoveRecurringSchedule(ctx, id).Execute()

Remove Recurring Schedule



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
    id := "id_example" // string | ID of deployment

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.RemoveRecurringSchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RemoveRecurringSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveRecurringSchedule`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.RemoveRecurringSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRecurringScheduleRequest struct via the builder pattern


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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset
    trustedid := "trustedid_example" // string | ID of project to untrust

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.RemoveTrustedProject(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.RemoveTrustedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTrustedProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.RemoveTrustedProject`: %v\n", resp)
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset
    inputAssetForUpdate := *cons3rtclient.NewInputAssetForUpdate() // InputAssetForUpdate | The modified Asset metadata (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateAsset(context.Background(), id).InputAssetForUpdate(inputAssetForUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateAsset`: %v\n", resp)
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset to delete
    state := "state_example" // string | The new asset state type

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateAssetState(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateAssetState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetState`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateAssetState`: %v\n", resp)
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset to update
    visibility := "visibility_example" // string | The new asset visibility type

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateAssetVisibilityQuery(context.Background(), id).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateAssetVisibilityQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetVisibilityQuery`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateAssetVisibilityQuery`: %v\n", resp)
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset
    impactlevel := "impactlevel_example" // string | The new asset impact level type.

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateImpactLevel(context.Background(), id).Impactlevel(impactlevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateImpactLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImpactLevel`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateImpactLevel`: %v\n", resp)
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset
    limit := int64(789) // int64 | The new asset instance limit

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateInstanceLimit(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateInstanceLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceLimit`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateInstanceLimit`: %v\n", resp)
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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of asset to update
    offline := true // bool | Set the asset status to offline (optional) (default to true)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateOfflineStatus(context.Background(), id).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateOfflineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfflineStatus`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateOfflineStatus`: %v\n", resp)
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

