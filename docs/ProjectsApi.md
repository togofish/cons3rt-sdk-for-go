# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectMember**](ProjectsApi.md#AddProjectMember) | **Put** /api/projects/{id}/members | Assign Project Member
[**AddRoleToProjectMember**](ProjectsApi.md#AddRoleToProjectMember) | **Put** /api/projects/{id}/members/roles | Assign Role to Member
[**AddSubmissionServiceToProject**](ProjectsApi.md#AddSubmissionServiceToProject) | **Post** /api/projects/{id}/submission | Add Submission Service
[**AddTrustedProject1**](ProjectsApi.md#AddTrustedProject1) | **Put** /api/projects/{id}/addtrustedproject | Assign Trusted Project to Project
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /api/projects | Create a Project
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /api/projects/{id} | Delete Project
[**GetHostConfigurationMetrics**](ProjectsApi.md#GetHostConfigurationMetrics) | **Get** /api/projects/{id}/metrics/hostconfiguration | Retrieve Metrics
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /api/projects/{id} | Retrieve Project
[**GetProjectVirtRealms**](ProjectsApi.md#GetProjectVirtRealms) | **Get** /api/projects/{id}/virtualizationrealms | List Virtualization Realms
[**GetProjects**](ProjectsApi.md#GetProjects) | **Get** /api/projects | List Joined Projects
[**GetProjectsExpanded**](ProjectsApi.md#GetProjectsExpanded) | **Get** /api/projects/expanded | List Unjoined Projects
[**GetVirtualMachineCountMetrics**](ProjectsApi.md#GetVirtualMachineCountMetrics) | **Get** /api/projects/{id}/metrics/virtualmachinecount | Retrieve Virtual Machine Metrics
[**ListMembers**](ProjectsApi.md#ListMembers) | **Get** /api/projects/{id}/members | List Members
[**ListSubmissionSerivcesForProject**](ProjectsApi.md#ListSubmissionSerivcesForProject) | **Get** /api/projects/{id}/submission | List Submission Services
[**RemoveProjectMember**](ProjectsApi.md#RemoveProjectMember) | **Delete** /api/projects/{id}/members | Unassign Member from Project
[**RemoveRoleFromProjectMember**](ProjectsApi.md#RemoveRoleFromProjectMember) | **Delete** /api/projects/{id}/members/roles | Unassign Role from Member
[**RemoveSubmissionServiceFromProject**](ProjectsApi.md#RemoveSubmissionServiceFromProject) | **Delete** /api/projects/{id}/submission/{submission_service_id} | Remove Submission Service
[**RemoveTrustedProject1**](ProjectsApi.md#RemoveTrustedProject1) | **Put** /api/projects/{id}/removetrustedproject | Unassign Trusted Project from Project
[**RequestProjectInvitation**](ProjectsApi.md#RequestProjectInvitation) | **Post** /api/projects/{id}/invitation | Create Invitation Code
[**SetProjectDefaultPowerSchedule**](ProjectsApi.md#SetProjectDefaultPowerSchedule) | **Put** /api/projects/{id}/powerschedule | Update Default Power Schedule
[**SetProjectDefaultVirtualizationRealm**](ProjectsApi.md#SetProjectDefaultVirtualizationRealm) | **Put** /api/projects/{id}/virtualizationrealms/default | Update Default Virtualization Realm
[**SetProjectItarInformation**](ProjectsApi.md#SetProjectItarInformation) | **Put** /api/projects/{id}/itar | Set Asset Export Restriction
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /api/projects/{id} | Update Project
[**UpdateSubmissionService**](ProjectsApi.md#UpdateSubmissionService) | **Put** /api/projects/{id}/submission/{submission_service_id} | Update Submission Service



## AddProjectMember

> bool AddProjectMember(ctx, id).Username(username).Execute()

Assign Project Member



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
    username := "username_example" // string | Username of desired member
    id := "id_example" // string | ID of project

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.AddProjectMember(context.Background(), id).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddProjectMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProjectMember`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.AddProjectMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | Username of desired member | 


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


## AddRoleToProjectMember

> bool AddRoleToProjectMember(ctx, id).Username(username).Role(role).Execute()

Assign Role to Member



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
    username := "username_example" // string | Username of project member
    role := "role_example" // string | Project role to add

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.AddRoleToProjectMember(context.Background(), id).Username(username).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddRoleToProjectMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleToProjectMember`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.AddRoleToProjectMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleToProjectMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | Username of project member | 
 **role** | **string** | Project role to add | 

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


## AddSubmissionServiceToProject

> bool AddSubmissionServiceToProject(ctx, id).InputSubmissionServiceForProject(inputSubmissionServiceForProject).Execute()

Add Submission Service



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
    inputSubmissionServiceForProject := *cons3rtclient.NewInputSubmissionServiceForProject(*cons3rtclient.NewInputSubmissionEndpointForProject("Host_example", "Subtype_example")) // InputSubmissionServiceForProject | The submission service

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.AddSubmissionServiceToProject(context.Background(), id).InputSubmissionServiceForProject(inputSubmissionServiceForProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddSubmissionServiceToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSubmissionServiceToProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.AddSubmissionServiceToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSubmissionServiceToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputSubmissionServiceForProject** | [**InputSubmissionServiceForProject**](InputSubmissionServiceForProject.md) | The submission service | 

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


## AddTrustedProject1

> bool AddTrustedProject1(ctx, id).Trustedid(trustedid).Execute()

Assign Trusted Project to Project



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
    trustedid := "trustedid_example" // string | ID of trusted project

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.AddTrustedProject1(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddTrustedProject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedProject1`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.AddTrustedProject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTrustedProject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedid** | **string** | ID of trusted project | 

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


## CreateProject

> string CreateProject(ctx).InputProjectFull(inputProjectFull).Execute()

Create a Project



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
    inputProjectFull := *cons3rtclient.NewInputProjectFull("Name_example", *cons3rtclient.NewProjectLimits(int32(123), int32(123), int32(123), int32(123), int32(123))) // InputProjectFull | The Project to create (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.CreateProject(context.Background()).InputProjectFull(inputProjectFull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputProjectFull** | [**InputProjectFull**](InputProjectFull.md) | The Project to create | 

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


## DeleteProject

> bool DeleteProject(ctx, id).Force(force).Execute()

Delete Project



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
    force := true // bool | Delete all dependencies (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.DeleteProject(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.DeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Delete all dependencies | [default to false]

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
    resp, r, err := api_client.ProjectsApi.GetHostConfigurationMetrics(context.Background(), id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetHostConfigurationMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostConfigurationMetrics`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetHostConfigurationMetrics`: %v\n", resp)
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


## GetProject

> FullProject GetProject(ctx, id).Execute()

Retrieve Project



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GetProject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: FullProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullProject**](FullProject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectVirtRealms

> []MinimalVirtualizationRealm GetProjectVirtRealms(ctx, id).Execute()

List Virtualization Realms



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GetProjectVirtRealms(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectVirtRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectVirtRealms`: []MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectVirtRealms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectVirtRealmsRequest struct via the builder pattern


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


## GetProjects

> []MinimalProject GetProjects(ctx).Maxresults(maxresults).Page(page).Execute()

List Joined Projects



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
    resp, r, err := api_client.ProjectsApi.GetProjects(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: []MinimalProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalProject**](MinimalProject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsExpanded

> []MinimalProject GetProjectsExpanded(ctx).Maxresults(maxresults).Page(page).Execute()

List Unjoined Projects



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
    resp, r, err := api_client.ProjectsApi.GetProjectsExpanded(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectsExpanded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectsExpanded`: []MinimalProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectsExpanded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalProject**](MinimalProject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineCountMetrics

> string GetVirtualMachineCountMetrics(ctx, id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()

Retrieve Virtual Machine Metrics



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
    resp, r, err := api_client.ProjectsApi.GetVirtualMachineCountMetrics(context.Background(), id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetVirtualMachineCountMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineCountMetrics`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetVirtualMachineCountMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineCountMetricsRequest struct via the builder pattern


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


## ListMembers

> []UserProject ListMembers(ctx, id).MembershipState(membershipState).Role(role).Name(name).Maxresults(maxresults).Page(page).Execute()

List Members



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
    membershipState := "membershipState_example" // string | Project membership state type (optional)
    role := "role_example" // string | Project member role type (optional)
    name := "name_example" // string | User name to search for (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ListMembers(context.Background(), id).MembershipState(membershipState).Role(role).Name(name).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembers`: []UserProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipState** | **string** | Project membership state type | 
 **role** | **string** | Project member role type | 
 **name** | **string** | User name to search for | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]UserProject**](UserProject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubmissionSerivcesForProject

> []SubmissionService ListSubmissionSerivcesForProject(ctx, id).Execute()

List Submission Services



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ListSubmissionSerivcesForProject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ListSubmissionSerivcesForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubmissionSerivcesForProject`: []SubmissionService
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ListSubmissionSerivcesForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubmissionSerivcesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SubmissionService**](SubmissionService.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProjectMember

> bool RemoveProjectMember(ctx, id).Username(username).Execute()

Unassign Member from Project



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
    username := "username_example" // string | Username of desired member
    id := "id_example" // string | ID of project

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.RemoveProjectMember(context.Background(), id).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RemoveProjectMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveProjectMember`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RemoveProjectMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | Username of desired member | 


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


## RemoveRoleFromProjectMember

> bool RemoveRoleFromProjectMember(ctx, id).Username(username).Role(role).Execute()

Unassign Role from Member



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
    username := "username_example" // string | Username of project member
    role := "role_example" // string | Project role to remove

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.RemoveRoleFromProjectMember(context.Background(), id).Username(username).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RemoveRoleFromProjectMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveRoleFromProjectMember`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RemoveRoleFromProjectMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleFromProjectMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | Username of project member | 
 **role** | **string** | Project role to remove | 

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


## RemoveSubmissionServiceFromProject

> bool RemoveSubmissionServiceFromProject(ctx, id, submissionServiceId).Execute()

Remove Submission Service



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
    submissionServiceId := "submissionServiceId_example" // string | ID of submission service

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.RemoveSubmissionServiceFromProject(context.Background(), id, submissionServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RemoveSubmissionServiceFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSubmissionServiceFromProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RemoveSubmissionServiceFromProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 
**submissionServiceId** | **string** | ID of submission service | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubmissionServiceFromProjectRequest struct via the builder pattern


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


## RemoveTrustedProject1

> bool RemoveTrustedProject1(ctx, id).Trustedid(trustedid).Execute()

Unassign Trusted Project from Project



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
    trustedid := "trustedid_example" // string | ID of trusted project

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.RemoveTrustedProject1(context.Background(), id).Trustedid(trustedid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RemoveTrustedProject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTrustedProject1`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RemoveTrustedProject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTrustedProject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedid** | **string** | ID of trusted project | 

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


## RequestProjectInvitation

> string RequestProjectInvitation(ctx, id).Email(email).Execute()

Create Invitation Code



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
    email := "email_example" // string | Email address of invitee

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.RequestProjectInvitation(context.Background(), id).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RequestProjectInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestProjectInvitation`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RequestProjectInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestProjectInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **string** | Email address of invitee | 

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


## SetProjectDefaultPowerSchedule

> bool SetProjectDefaultPowerSchedule(ctx, id).PowerSchedule(powerSchedule).Execute()

Update Default Power Schedule



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
    powerSchedule := *cons3rtclient.NewPowerSchedule("Mode_example") // PowerSchedule | The desired power schedule (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.SetProjectDefaultPowerSchedule(context.Background(), id).PowerSchedule(powerSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SetProjectDefaultPowerSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProjectDefaultPowerSchedule`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SetProjectDefaultPowerSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectDefaultPowerScheduleRequest struct via the builder pattern


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


## SetProjectDefaultVirtualizationRealm

> bool SetProjectDefaultVirtualizationRealm(ctx, id).Virtualizationrealmid(virtualizationrealmid).Execute()

Update Default Virtualization Realm



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
    virtualizationrealmid := "virtualizationrealmid_example" // string | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.SetProjectDefaultVirtualizationRealm(context.Background(), id).Virtualizationrealmid(virtualizationrealmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SetProjectDefaultVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProjectDefaultVirtualizationRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SetProjectDefaultVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectDefaultVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationrealmid** | **string** | ID of virtualization realm | 

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


## SetProjectItarInformation

> bool SetProjectItarInformation(ctx, id).Message(message).Execute()

Set Asset Export Restriction



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
    message := "message_example" // string | Additional information about the export restriction (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.SetProjectItarInformation(context.Background(), id).Message(message).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SetProjectItarInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProjectItarInformation`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SetProjectItarInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectItarInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **string** | Additional information about the export restriction | 

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


## UpdateProject

> bool UpdateProject(ctx, id).InputProjectUpdate(inputProjectUpdate).Execute()

Update Project



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
    inputProjectUpdate := *cons3rtclient.NewInputProjectUpdate("Name_example") // InputProjectUpdate | The modified Project (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.UpdateProject(context.Background(), id).InputProjectUpdate(inputProjectUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputProjectUpdate** | [**InputProjectUpdate**](InputProjectUpdate.md) | The modified Project | 

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


## UpdateSubmissionService

> bool UpdateSubmissionService(ctx, id, submissionServiceId).InputSubmissionServiceForProject(inputSubmissionServiceForProject).Execute()

Update Submission Service



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
    submissionServiceId := "submissionServiceId_example" // string | ID of submission service
    inputSubmissionServiceForProject := *cons3rtclient.NewInputSubmissionServiceForProject(*cons3rtclient.NewInputSubmissionEndpointForProject("Host_example", "Subtype_example")) // InputSubmissionServiceForProject | The submission service

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.UpdateSubmissionService(context.Background(), id, submissionServiceId).InputSubmissionServiceForProject(inputSubmissionServiceForProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateSubmissionService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubmissionService`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateSubmissionService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of project | 
**submissionServiceId** | **string** | ID of submission service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubmissionServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputSubmissionServiceForProject** | [**InputSubmissionServiceForProject**](InputSubmissionServiceForProject.md) | The submission service | 

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

