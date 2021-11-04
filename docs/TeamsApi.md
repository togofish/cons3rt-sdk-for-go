# \TeamsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamManagerToTeam**](TeamsApi.md#AddTeamManagerToTeam) | **Put** /api/teams/{id}/managers | Assign Manager
[**CreateTeam**](TeamsApi.md#CreateTeam) | **Post** /api/teams | Create Team
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /api/teams/{id} | Delete Team
[**GetTeamOwnedClouds**](TeamsApi.md#GetTeamOwnedClouds) | **Get** /api/teams/{id}/clouds | List Clouds
[**GetTeamOwnedOrManagedVirtualizationRealms**](TeamsApi.md#GetTeamOwnedOrManagedVirtualizationRealms) | **Get** /api/teams/{id}/virtualizationrealms | List Virtualization Realms
[**GetTeams**](TeamsApi.md#GetTeams) | **Get** /api/teams | List Teams
[**RemoveTeamManagerFromTeam**](TeamsApi.md#RemoveTeamManagerFromTeam) | **Delete** /api/teams/{id}/managers | Unassign Manager from Team
[**RetrieveTeam**](TeamsApi.md#RetrieveTeam) | **Get** /api/teams/{id} | Retrieve Team
[**SetProjectLimits**](TeamsApi.md#SetProjectLimits) | **Put** /api/teams/{id}/project/{project_id}/limits | Update Project Limits
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Put** /api/teams/{id} | Update Team
[**UpdateTeamState**](TeamsApi.md#UpdateTeamState) | **Put** /api/teams/{id}/updatestate | Update State



## AddTeamManagerToTeam

> bool AddTeamManagerToTeam(ctx, id).Username(username).Execute()

Assign Manager



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
    id := "id_example" // string | ID of team
    username := "username_example" // string | Username of desired manager

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.AddTeamManagerToTeam(context.Background(), id).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.AddTeamManagerToTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamManagerToTeam`: bool
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.AddTeamManagerToTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamManagerToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | Username of desired manager | 

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


## CreateTeam

> string CreateTeam(ctx).InputTeamFull(inputTeamFull).Execute()

Create Team



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
    inputTeamFull := *cons3rtclient.NewInputTeamFull(*cons3rtclient.NewInputUser("Username_example", "Email_example", "Firstname_example", "Lastname_example"), []cons3rtclient.InputUser{*cons3rtclient.NewInputUser("Username_example", "Email_example", "Firstname_example", "Lastname_example")}, "Name_example", *cons3rtclient.NewPocInfo(), "State_example", int32(123)) // InputTeamFull | The Team to create (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.CreateTeam(context.Background()).InputTeamFull(inputTeamFull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeam`: string
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputTeamFull** | [**InputTeamFull**](InputTeamFull.md) | The Team to create | 

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


## DeleteTeam

> bool DeleteTeam(ctx, id).Execute()

Delete Team



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
    id := "id_example" // string | ID of team

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.DeleteTeam(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.DeleteTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTeam`: bool
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.DeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## GetTeamOwnedClouds

> []MinimalCloud GetTeamOwnedClouds(ctx, id).Maxresults(maxresults).Page(page).Execute()

List Clouds



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
    id := "id_example" // string | ID of team
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.GetTeamOwnedClouds(context.Background(), id).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeamOwnedClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamOwnedClouds`: []MinimalCloud
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetTeamOwnedClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamOwnedCloudsRequest struct via the builder pattern


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


## GetTeamOwnedOrManagedVirtualizationRealms

> []MinimalVirtualizationRealm GetTeamOwnedOrManagedVirtualizationRealms(ctx, id).Maxresults(maxresults).Page(page).Execute()

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
    id := "id_example" // string | ID of team
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.GetTeamOwnedOrManagedVirtualizationRealms(context.Background(), id).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeamOwnedOrManagedVirtualizationRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamOwnedOrManagedVirtualizationRealms`: []MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetTeamOwnedOrManagedVirtualizationRealms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamOwnedOrManagedVirtualizationRealmsRequest struct via the builder pattern


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


## GetTeams

> []MinimalTeam GetTeams(ctx).Maxresults(maxresults).Page(page).Execute()

List Teams



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
    resp, r, err := api_client.TeamsApi.GetTeams(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeams`: []MinimalTeam
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]MinimalTeam**](MinimalTeam.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTeamManagerFromTeam

> bool RemoveTeamManagerFromTeam(ctx, id).Username(username).Execute()

Unassign Manager from Team



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
    id := "id_example" // string | ID of team
    username := "username_example" // string | Username of manager to remove

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.RemoveTeamManagerFromTeam(context.Background(), id).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RemoveTeamManagerFromTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTeamManagerFromTeam`: bool
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.RemoveTeamManagerFromTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTeamManagerFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | Username of manager to remove | 

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


## RetrieveTeam

> FullTeam RetrieveTeam(ctx, id).Execute()

Retrieve Team



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
    id := "id_example" // string | ID of team

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.RetrieveTeam(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RetrieveTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTeam`: FullTeam
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.RetrieveTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullTeam**](FullTeam.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProjectLimits

> bool SetProjectLimits(ctx, id, projectId).ProjectLimits(projectLimits).Execute()

Update Project Limits



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
    id := "id_example" // string | ID of team
    projectId := "projectId_example" // string | ID of project
    projectLimits := *cons3rtclient.NewProjectLimits(int32(123), int32(123), int32(123), int32(123), int32(123)) // ProjectLimits | The desired project limits

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.SetProjectLimits(context.Background(), id, projectId).ProjectLimits(projectLimits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.SetProjectLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetProjectLimits`: bool
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.SetProjectLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 
**projectId** | **string** | ID of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectLimits** | [**ProjectLimits**](ProjectLimits.md) | The desired project limits | 

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


## UpdateTeam

> bool UpdateTeam(ctx, id).InputTeamFull(inputTeamFull).Execute()

Update Team



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
    id := "id_example" // string | ID of team
    inputTeamFull := *cons3rtclient.NewInputTeamFull(*cons3rtclient.NewInputUser("Username_example", "Email_example", "Firstname_example", "Lastname_example"), []cons3rtclient.InputUser{*cons3rtclient.NewInputUser("Username_example", "Email_example", "Firstname_example", "Lastname_example")}, "Name_example", *cons3rtclient.NewPocInfo(), "State_example", int32(123)) // InputTeamFull | The modified team definition (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.UpdateTeam(context.Background(), id).InputTeamFull(inputTeamFull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeam`: bool
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputTeamFull** | [**InputTeamFull**](InputTeamFull.md) | The modified team definition | 

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


## UpdateTeamState

> bool UpdateTeamState(ctx, id).State(state).Execute()

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
    id := "id_example" // string | ID of team
    state := "state_example" // string | Updated team state type

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.UpdateTeamState(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamState`: bool
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Updated team state type | 

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

