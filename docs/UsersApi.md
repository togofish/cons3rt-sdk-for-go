# \UsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/users | Create a new user
[**GetPendingUsers**](UsersApi.md#GetPendingUsers) | **Get** /api/users/pending | List Pending Users
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /api/users | List all Users



## CreateUser

> bool CreateUser(ctx).InputUser(inputUser).Execute()

Create a new user



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
    inputUser := *cons3rtclient.NewInputUser("Username_example", "Email_example", "Firstname_example", "Lastname_example") // InputUser | The User definition (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.CreateUser(context.Background()).InputUser(inputUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: bool
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputUser** | [**InputUser**](InputUser.md) | The User definition | 

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


## GetPendingUsers

> []BasicUser GetPendingUsers(ctx).Execute()

List Pending Users



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetPendingUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetPendingUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingUsers`: []BasicUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetPendingUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingUsersRequest struct via the builder pattern


### Return type

[**[]BasicUser**](BasicUser.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []BasicUser GetUsers(ctx).State(state).Createdbefore(createdbefore).Createdafter(createdafter).Maxresults(maxresults).Page(page).Execute()

List all Users



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
    state := "state_example" // string | The users state, if a state is not provided users in all states will be returned (optional)
    createdbefore := int64(789) // int64 | Date (seconds since epoch) to filter on (optional)
    createdafter := int64(789) // int64 | Date (seconds since epoch) to filter on (optional)
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUsers(context.Background()).State(state).Createdbefore(createdbefore).Createdafter(createdafter).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: []BasicUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | The users state, if a state is not provided users in all states will be returned | 
 **createdbefore** | **int64** | Date (seconds since epoch) to filter on | 
 **createdafter** | **int64** | Date (seconds since epoch) to filter on | 
 **maxresults** | **int64** | Maximum number of results to return | [default to 40]
 **page** | **int64** | Requested page number | [default to 0]

### Return type

[**[]BasicUser**](BasicUser.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

