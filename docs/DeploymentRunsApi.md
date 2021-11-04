# \DeploymentRunsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCategoryToDeploymentRun**](DeploymentRunsApi.md#AddCategoryToDeploymentRun) | **Put** /api/categories/{id}/run | Assign Category to Run
[**CreateIdentity**](DeploymentRunsApi.md#CreateIdentity) | **Post** /api/drs/{id}/host/{hostid}/identity | Create a host identity
[**DeleteDeploymentRun**](DeploymentRunsApi.md#DeleteDeploymentRun) | **Delete** /api/drs/{id} | Delete Deployment Run
[**DeleteIdentity**](DeploymentRunsApi.md#DeleteIdentity) | **Delete** /api/drs/{id}/host/{hostid}/identity | Delete host identity
[**DeleteIdentityById**](DeploymentRunsApi.md#DeleteIdentityById) | **Delete** /api/drs/{id}/host/{hostid}/identity/{username} | Deletes identity for specified user
[**DownloadDeploymentRunTestReport**](DeploymentRunsApi.md#DownloadDeploymentRunTestReport) | **Get** /api/drs/{id}/downloadreport | Download Report
[**DownloadHost**](DeploymentRunsApi.md#DownloadHost) | **Get** /api/drs/{id}/downloadhost | Download Host
[**GetDeploymentRun**](DeploymentRunsApi.md#GetDeploymentRun) | **Get** /api/drs/{id} | Retrieve Deployment Run
[**GetDeploymentRunReports**](DeploymentRunsApi.md#GetDeploymentRunReports) | **Get** /api/drs/{id}/reports | List Reports
[**GetDeploymentRuns**](DeploymentRunsApi.md#GetDeploymentRuns) | **Get** /api/deployments/{id}/runs | List Deployment Runs
[**GetDeploymentRuns1**](DeploymentRunsApi.md#GetDeploymentRuns1) | **Get** /api/drs | List Deployment Runs
[**GetHost**](DeploymentRunsApi.md#GetHost) | **Get** /api/drs/{id}/host/{hostid} | Retrieve Host
[**GetHostAccess**](DeploymentRunsApi.md#GetHostAccess) | **Get** /api/drs/{id}/host/{hostid}/access | List Host Access Logs
[**GetHostConfigurationMetrics**](DeploymentRunsApi.md#GetHostConfigurationMetrics) | **Get** /api/projects/{id}/metrics/hostconfiguration | Retrieve Metrics
[**GetHostInstanceTypes**](DeploymentRunsApi.md#GetHostInstanceTypes) | **Get** /api/drs/{id}/host/{hostid}/resize | List available instance types for host
[**GetIdentities**](DeploymentRunsApi.md#GetIdentities) | **Get** /api/drs/{id}/host/{hostid}/identities | Get Host Identities
[**GetIdentity**](DeploymentRunsApi.md#GetIdentity) | **Get** /api/drs/{id}/host/{hostid}/identity | Get Host Identity For User
[**PerformHostAction**](DeploymentRunsApi.md#PerformHostAction) | **Put** /api/drs/{id}/hostaction | Execute Host Action
[**PublishDeploymentRun**](DeploymentRunsApi.md#PublishDeploymentRun) | **Post** /api/drs/{id}/publish | Publish Deployment Run
[**RedeployContainerAsset**](DeploymentRunsApi.md#RedeployContainerAsset) | **Put** /api/drs/{id}/host/{hostid}/container | Re-deploy Container Asset
[**RedeployDeploymentRunHosts**](DeploymentRunsApi.md#RedeployDeploymentRunHosts) | **Put** /api/drs/{id}/redeployhosts | Redeploy Deployment Run Hosts
[**RelaunchDeploymentRun**](DeploymentRunsApi.md#RelaunchDeploymentRun) | **Put** /api/drs/{id}/rerun | Relaunch Deployment Run
[**ReleaseDeploymentRun**](DeploymentRunsApi.md#ReleaseDeploymentRun) | **Put** /api/drs/{id}/release | Release Deployment Run
[**RemoveCategoryFromDeploymentRun**](DeploymentRunsApi.md#RemoveCategoryFromDeploymentRun) | **Delete** /api/categories/{id}/run | Unassign Category from deployment run
[**RetestDeploymentRun**](DeploymentRunsApi.md#RetestDeploymentRun) | **Put** /api/drs/{id}/retest | Re-test Deployment Run
[**SetDeploymentRunLock**](DeploymentRunsApi.md#SetDeploymentRunLock) | **Put** /api/drs/{id}/setlock | Update Lock
[**SetPowerScheduleForDeploymentRun**](DeploymentRunsApi.md#SetPowerScheduleForDeploymentRun) | **Put** /api/drs/{id}/powerschedule | Update Power Schedule
[**UnpublishDeploymentRun**](DeploymentRunsApi.md#UnpublishDeploymentRun) | **Delete** /api/drs/{id}/publish | Unpublish Deployment Run



## AddCategoryToDeploymentRun

> bool AddCategoryToDeploymentRun(ctx, id).Runid(runid).Execute()

Assign Category to Run



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
    id := "id_example" // string | ID of category
    runid := "runid_example" // string | ID of run to assign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.AddCategoryToDeploymentRun(context.Background(), id).Runid(runid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.AddCategoryToDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCategoryToDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.AddCategoryToDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCategoryToDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runid** | **string** | ID of run to assign | 

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


## CreateIdentity

> []BaseIdentity CreateIdentity(ctx, id, hostid).CloudResourceObject(cloudResourceObject).Execute()

Create a host identity



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host
    cloudResourceObject := []cons3rtclient.CloudResourceObject{*cons3rtclient.NewCloudResourceObject("Type_example", "Identifier_example")} // []CloudResourceObject | The cloud resources to be accessed by the host identity

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.CreateIdentity(context.Background(), id, hostid).CloudResourceObject(cloudResourceObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.CreateIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentity`: []BaseIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.CreateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudResourceObject** | [**[]CloudResourceObject**](CloudResourceObject.md) | The cloud resources to be accessed by the host identity | 

### Return type

[**[]BaseIdentity**](BaseIdentity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentRun

> bool DeleteDeploymentRun(ctx, id).Purge(purge).Execute()

Delete Deployment Run



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
    id := "id_example" // string | ID of deployment run
    purge := true // bool | Delete all dependencies of the deployment run (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.DeleteDeploymentRun(context.Background(), id).Purge(purge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.DeleteDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.DeleteDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | Delete all dependencies of the deployment run | [default to false]

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


## DeleteIdentity

> bool DeleteIdentity(ctx, id, hostid).Execute()

Delete host identity



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.DeleteIdentity(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.DeleteIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentity`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.DeleteIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityRequest struct via the builder pattern


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


## DeleteIdentityById

> []BaseIdentity DeleteIdentityById(ctx, id, hostid, username).Execute()

Deletes identity for specified user



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host
    username := "username_example" // string | Username of the identity to be deleted

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.DeleteIdentityById(context.Background(), id, hostid, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.DeleteIdentityById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdentityById`: []BaseIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.DeleteIdentityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 
**username** | **string** | Username of the identity to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BaseIdentity**](BaseIdentity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDeploymentRunTestReport

> DownloadDeploymentRunTestReport(ctx, id).File(file).Number(number).Execute()

Download Report



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
    id := "id_example" // string | ID of deployment run
    file := "file_example" // string | Report file name (optional)
    number := "number_example" // string | Report number (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.DownloadDeploymentRunTestReport(context.Background(), id).File(file).Number(number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.DownloadDeploymentRunTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDeploymentRunTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **string** | Report file name | 
 **number** | **string** | Report number | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadHost

> DownloadHost(ctx, id).Role(role).Background(background).Execute()

Download Host



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
    id := "id_example" // string | ID of deployment run
    role := "role_example" // string | Name of host to bundle for download
    background := true // bool | Force the download to happen in the background (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.DownloadHost(context.Background(), id).Role(role).Background(background).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.DownloadHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | Name of host to bundle for download | 
 **background** | **bool** | Force the download to happen in the background | [default to false]

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRun

> FullDeploymentRun GetDeploymentRun(ctx, id).Execute()

Retrieve Deployment Run



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
    id := "id_example" // string | ID of deployment run

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRun`: FullDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullDeploymentRun**](FullDeploymentRun.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRunReports

> []string GetDeploymentRunReports(ctx, id).Execute()

List Reports



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
    id := "id_example" // string | ID of deployment run

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetDeploymentRunReports(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetDeploymentRunReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRunReports`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetDeploymentRunReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRunReportsRequest struct via the builder pattern


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
    resp, r, err := api_client.DeploymentRunsApi.GetDeploymentRuns(context.Background(), id).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetDeploymentRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRuns`: []MinimalDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetDeploymentRuns`: %v\n", resp)
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
    resp, r, err := api_client.DeploymentRunsApi.GetDeploymentRuns1(context.Background()).SearchType(searchType).InProject(inProject).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetDeploymentRuns1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRuns1`: []MinimalDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetDeploymentRuns1`: %v\n", resp)
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


## GetHost

> FullDeploymentRunHost GetHost(ctx, id, hostid).Execute()

Retrieve Host



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetHost(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHost`: FullDeploymentRunHost
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FullDeploymentRunHost**](FullDeploymentRunHost.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostAccess

> []RemoteAccessSession GetHostAccess(ctx, id, hostid).Maxresults(maxresults).Page(page).Execute()

List Host Access Logs



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetHostAccess(context.Background(), id, hostid).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetHostAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostAccess`: []RemoteAccessSession
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetHostAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]RemoteAccessSession**](RemoteAccessSession.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostConfigurationMetrics

> string GetHostConfigurationMetrics(ctx, id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()

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
    id := "id_example" // string | ID of project
    start := int64(789) // int64 | Interval start time, specified in seconds since epoch
    end := int64(789) // int64 | Interval end time, specified in seconds since epoch
    interval := int64(789) // int64 | Number of intervals (optional) (default to 1)
    intervalUnit := "intervalUnit_example" // string | Interval unit (optional) (default to "HOURS")

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetHostConfigurationMetrics(context.Background(), id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetHostConfigurationMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostConfigurationMetrics`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetHostConfigurationMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostConfigurationMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int64** | Interval start time, specified in seconds since epoch | 
 **end** | **int64** | Interval end time, specified in seconds since epoch | 
 **interval** | **int64** | Number of intervals | [default to 1]
 **intervalUnit** | **string** | Interval unit | [default to &quot;HOURS&quot;]

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


## GetHostInstanceTypes

> TargetInstanceTypes GetHostInstanceTypes(ctx, id, hostid).Execute()

List available instance types for host



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetHostInstanceTypes(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetHostInstanceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostInstanceTypes`: TargetInstanceTypes
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetHostInstanceTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TargetInstanceTypes**](TargetInstanceTypes.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentities

> []BaseIdentity GetIdentities(ctx, id, hostid).Execute()

Get Host Identities



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetIdentities(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentities`: []BaseIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BaseIdentity**](BaseIdentity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> []CloudResourceAccessListing GetIdentity(ctx, id, hostid).Execute()

Get Host Identity For User



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.GetIdentity(context.Background(), id, hostid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.GetIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentity`: []CloudResourceAccessListing
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CloudResourceAccessListing**](CloudResourceAccessListing.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformHostAction

> bool PerformHostAction(ctx, id).Deploymentrunhostid(deploymentrunhostid).Action(action).Cpu(cpu).Ram(ram).InstanceTypeName(instanceTypeName).Execute()

Execute Host Action



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
    id := "id_example" // string | ID of deployment run
    deploymentrunhostid := "deploymentrunhostid_example" // string | ID of host
    action := "action_example" // string | Action to perform
    cpu := int32(56) // int32 | Desired number of CPUs, if resizing host in a non instance type based virtualization realm (optional)
    ram := int32(56) // int32 | Desired amount of RAM in Mebibytes, if resizing host in a non instance type based virtualization realm (optional)
    instanceTypeName := "instanceTypeName_example" // string | The instance type name to resize to, if resizing host in an instance type based virtualization realm (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.PerformHostAction(context.Background(), id).Deploymentrunhostid(deploymentrunhostid).Action(action).Cpu(cpu).Ram(ram).InstanceTypeName(instanceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.PerformHostAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformHostAction`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.PerformHostAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformHostActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentrunhostid** | **string** | ID of host | 
 **action** | **string** | Action to perform | 
 **cpu** | **int32** | Desired number of CPUs, if resizing host in a non instance type based virtualization realm | 
 **ram** | **int32** | Desired amount of RAM in Mebibytes, if resizing host in a non instance type based virtualization realm | 
 **instanceTypeName** | **string** | The instance type name to resize to, if resizing host in an instance type based virtualization realm | 

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


## PublishDeploymentRun

> bool PublishDeploymentRun(ctx, id).Execute()

Publish Deployment Run



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
    id := "id_example" // string | ID of deployment run

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.PublishDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.PublishDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.PublishDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishDeploymentRunRequest struct via the builder pattern


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


## RedeployContainerAsset

> bool RedeployContainerAsset(ctx, id, hostid).Installationid(installationid).InputContainerComponent(inputContainerComponent).Execute()

Re-deploy Container Asset



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
    id := "id_example" // string | ID of deployment run
    hostid := "hostid_example" // string | ID of host
    installationid := "installationid_example" // string | ID of container asset installation
    inputContainerComponent := *cons3rtclient.NewInputContainerComponent(*cons3rtclient.NewInputContainerConfiguration("ContainerName_example"), *cons3rtclient.NewInputAsset(int32(123)), "Subtype_example") // InputContainerComponent | The updated Container Component definition (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.RedeployContainerAsset(context.Background(), id, hostid).Installationid(installationid).InputContainerComponent(inputContainerComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.RedeployContainerAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeployContainerAsset`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.RedeployContainerAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 
**hostid** | **string** | ID of host | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployContainerAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installationid** | **string** | ID of container asset installation | 
 **inputContainerComponent** | [**InputContainerComponent**](InputContainerComponent.md) | The updated Container Component definition | 

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


## RedeployDeploymentRunHosts

> bool RedeployDeploymentRunHosts(ctx, id).RestIdObject(restIdObject).Execute()

Redeploy Deployment Run Hosts



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
    id := "id_example" // string | ID of deployment run
    restIdObject := []cons3rtclient.RestIdObject{*cons3rtclient.NewRestIdObject("Id_example")} // []RestIdObject | The collection of deployment run host ids to redeploy (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.RedeployDeploymentRunHosts(context.Background(), id).RestIdObject(restIdObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.RedeployDeploymentRunHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeployDeploymentRunHosts`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.RedeployDeploymentRunHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployDeploymentRunHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restIdObject** | [**[]RestIdObject**](RestIdObject.md) | The collection of deployment run host ids to redeploy | 

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


## RelaunchDeploymentRun

> string RelaunchDeploymentRun(ctx, id).Execute()

Relaunch Deployment Run



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
    id := "id_example" // string | ID of deployment run

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.RelaunchDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.RelaunchDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RelaunchDeploymentRun`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.RelaunchDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelaunchDeploymentRunRequest struct via the builder pattern


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


## ReleaseDeploymentRun

> bool ReleaseDeploymentRun(ctx, id).Force(force).Execute()

Release Deployment Run



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
    id := "id_example" // string | ID of deployment run
    force := true // bool | Force the release of this run (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.ReleaseDeploymentRun(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.ReleaseDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.ReleaseDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force the release of this run | [default to false]

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


## RemoveCategoryFromDeploymentRun

> bool RemoveCategoryFromDeploymentRun(ctx, id).Runid(runid).Execute()

Unassign Category from deployment run



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
    id := "id_example" // string | ID of category
    runid := "runid_example" // string | ID of run to unassign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.RemoveCategoryFromDeploymentRun(context.Background(), id).Runid(runid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.RemoveCategoryFromDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCategoryFromDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.RemoveCategoryFromDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCategoryFromDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runid** | **string** | ID of run to unassign | 

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


## RetestDeploymentRun

> bool RetestDeploymentRun(ctx, id).Execute()

Re-test Deployment Run



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
    id := "id_example" // string | ID of deployment run

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.RetestDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.RetestDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetestDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.RetestDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetestDeploymentRunRequest struct via the builder pattern


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


## SetDeploymentRunLock

> bool SetDeploymentRunLock(ctx, id).Lock(lock).Execute()

Update Lock



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
    id := "id_example" // string | ID of deployment run
    lock := true // bool | The desired lock state

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.SetDeploymentRunLock(context.Background(), id).Lock(lock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.SetDeploymentRunLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDeploymentRunLock`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.SetDeploymentRunLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeploymentRunLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lock** | **bool** | The desired lock state | 

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


## SetPowerScheduleForDeploymentRun

> bool SetPowerScheduleForDeploymentRun(ctx, id).PowerSchedule(powerSchedule).Execute()

Update Power Schedule



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
    id := "id_example" // string | ID of deployment run
    powerSchedule := *cons3rtclient.NewPowerSchedule("Mode_example") // PowerSchedule | The desired power schedule (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.SetPowerScheduleForDeploymentRun(context.Background(), id).PowerSchedule(powerSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.SetPowerScheduleForDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPowerScheduleForDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.SetPowerScheduleForDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPowerScheduleForDeploymentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **powerSchedule** | [**PowerSchedule**](PowerSchedule.md) | The desired power schedule | 

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


## UnpublishDeploymentRun

> bool UnpublishDeploymentRun(ctx, id).Execute()

Unpublish Deployment Run



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
    id := "id_example" // string | ID of deployment run

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentRunsApi.UnpublishDeploymentRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentRunsApi.UnpublishDeploymentRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnpublishDeploymentRun`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeploymentRunsApi.UnpublishDeploymentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of deployment run | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishDeploymentRunRequest struct via the builder pattern


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

