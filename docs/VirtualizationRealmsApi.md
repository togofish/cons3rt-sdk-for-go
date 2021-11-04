# \VirtualizationRealmsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetwork**](VirtualizationRealmsApi.md#AddNetwork) | **Post** /api/virtualizationrealms/{id}/networks | Allocate Network
[**AddProject**](VirtualizationRealmsApi.md#AddProject) | **Put** /api/virtualizationrealms/{id}/projects | Assign Project
[**CreateTemplateSubsciption**](VirtualizationRealmsApi.md#CreateTemplateSubsciption) | **Post** /api/virtualizationrealms/{id}/templates/subscriptions | Create Template Subscription
[**DeleteNetwork**](VirtualizationRealmsApi.md#DeleteNetwork) | **Delete** /api/virtualizationrealms/{id}/networks/{networkId} | Delete Network
[**DeleteTemplateSubscription**](VirtualizationRealmsApi.md#DeleteTemplateSubscription) | **Delete** /api/virtualizationrealms/{id}/templates/subscriptions/{subscription_id} | Delete Template Subscription
[**DisableVirtRealmRemoteAccess**](VirtualizationRealmsApi.md#DisableVirtRealmRemoteAccess) | **Delete** /api/virtualizationrealms/{id}/remoteaccess | Disable Remote Access
[**EnableMaintenceMode1**](VirtualizationRealmsApi.md#EnableMaintenceMode1) | **Put** /api/virtualizationrealms/{id}/maintenance | Update Maintenance Mode
[**EnableVirtRealmRemoteAccess**](VirtualizationRealmsApi.md#EnableVirtRealmRemoteAccess) | **Post** /api/virtualizationrealms/{id}/remoteaccess | Enable Remote Access
[**GetDeploymentRunsInVirtualizationRealm**](VirtualizationRealmsApi.md#GetDeploymentRunsInVirtualizationRealm) | **Get** /api/virtualizationrealms/{id}/deploymentruns | List Deployment Runs
[**GetHostConfigurationMetrics1**](VirtualizationRealmsApi.md#GetHostConfigurationMetrics1) | **Get** /api/virtualizationrealms/{id}/metrics/hostconfiguration | Retrieve Metrics
[**GetNetwork**](VirtualizationRealmsApi.md#GetNetwork) | **Get** /api/virtualizationrealms/{id}/networks/{networkId} | Retrieve Network
[**GetNetworks**](VirtualizationRealmsApi.md#GetNetworks) | **Get** /api/virtualizationrealms/{id}/networks | List Networks
[**GetTemplateSubscription**](VirtualizationRealmsApi.md#GetTemplateSubscription) | **Get** /api/virtualizationrealms/{id}/templates/subscriptions/{subscription_id} | Retrieve Template Subscription
[**GetTemplatesInVirtualizationRealm**](VirtualizationRealmsApi.md#GetTemplatesInVirtualizationRealm) | **Get** /api/virtualizationrealms/{id}/templates | List Templates
[**GetUnregisteredNetworks**](VirtualizationRealmsApi.md#GetUnregisteredNetworks) | **Get** /api/virtualizationrealms/{id}/networks/unregistered | List Unregistered Networks
[**GetVirtualMachineCountMetrics1**](VirtualizationRealmsApi.md#GetVirtualMachineCountMetrics1) | **Get** /api/virtualizationrealms/{id}/metrics/virtualmachinecount | Retrieve Virtual Machine Metrics
[**GetVirtualizationRealm**](VirtualizationRealmsApi.md#GetVirtualizationRealm) | **Get** /api/virtualizationrealms/{id} | Retrieve Virtualization Realm
[**GetVirtualizationRealmResources**](VirtualizationRealmsApi.md#GetVirtualizationRealmResources) | **Get** /api/virtualizationrealms/{id}/resources | Retrieve Virtualization Realm Resources
[**GetVirtualizationRealms**](VirtualizationRealmsApi.md#GetVirtualizationRealms) | **Get** /api/virtualizationrealms | List Virtualization Realms
[**InvalidateTemplateCacheInVirtualizationRealm**](VirtualizationRealmsApi.md#InvalidateTemplateCacheInVirtualizationRealm) | **Put** /api/virtualizationrealms/{id}/templates/registrations | Refresh Template Cache
[**ListPendingTemplateSubscriptions**](VirtualizationRealmsApi.md#ListPendingTemplateSubscriptions) | **Get** /api/virtualizationrealms/{id}/templates/subscriptions/pending | List Pending Subscriptions
[**ListProjects**](VirtualizationRealmsApi.md#ListProjects) | **Get** /api/virtualizationrealms/{id}/projects | List Projects
[**ListTemplateRegistrations**](VirtualizationRealmsApi.md#ListTemplateRegistrations) | **Get** /api/virtualizationrealms/{id}/templates/registrations | List Template Registrations
[**ListTemplateSubscriptions**](VirtualizationRealmsApi.md#ListTemplateSubscriptions) | **Get** /api/virtualizationrealms/{id}/templates/subscriptions | List Template Subscriptions
[**ListUnregisteredTemplates**](VirtualizationRealmsApi.md#ListUnregisteredTemplates) | **Get** /api/virtualizationrealms/{id}/templates/registrations/pending | List Unregistered Templates
[**RegisterNetwork**](VirtualizationRealmsApi.md#RegisterNetwork) | **Put** /api/virtualizationrealms/{id}/networks/{networkIdentifier} | Register Network
[**RegisterTemplate**](VirtualizationRealmsApi.md#RegisterTemplate) | **Post** /api/virtualizationrealms/{id}/templates/registrations | Register Template
[**RemoveProject**](VirtualizationRealmsApi.md#RemoveProject) | **Delete** /api/virtualizationrealms/{id}/projects | Unassign Project
[**RetrieveTemplateRegistration**](VirtualizationRealmsApi.md#RetrieveTemplateRegistration) | **Get** /api/virtualizationrealms/{id}/templates/registrations/{registration_id} | Retrieve Template Registration
[**SetVirtualizationRealmActive**](VirtualizationRealmsApi.md#SetVirtualizationRealmActive) | **Put** /api/virtualizationrealms/{id}/activate | Update State
[**ShareTemplateRegistration**](VirtualizationRealmsApi.md#ShareTemplateRegistration) | **Post** /api/virtualizationrealms/{id}/templates/registrations/{registration_id}/share | Share Template
[**UnregisterTemplate**](VirtualizationRealmsApi.md#UnregisterTemplate) | **Delete** /api/virtualizationrealms/{id}/templates/registrations/{registration_id} | Delete Template Registration
[**UnshareTemplateRegistration**](VirtualizationRealmsApi.md#UnshareTemplateRegistration) | **Delete** /api/virtualizationrealms/{id}/templates/registrations/{registration_id}/share | Unshare Template
[**UpdateTemplateRegistration**](VirtualizationRealmsApi.md#UpdateTemplateRegistration) | **Put** /api/virtualizationrealms/{id}/templates/registrations/{registration_id} | Update Template Registration
[**UpdateTemplateSubscription**](VirtualizationRealmsApi.md#UpdateTemplateSubscription) | **Put** /api/virtualizationrealms/{id}/templates/subscriptions/{subscription_id} | Update Template Subscription
[**UpdateVirtRealmRemoteAccessConfig**](VirtualizationRealmsApi.md#UpdateVirtRealmRemoteAccessConfig) | **Put** /api/virtualizationrealms/{id}/remoteaccess | Update Remote Access
[**UpdateVirtualizationRealm**](VirtualizationRealmsApi.md#UpdateVirtualizationRealm) | **Put** /api/virtualizationrealms/{id} | Update Virtualization Realm
[**UpdateVirtualizationRealmReachability**](VirtualizationRealmsApi.md#UpdateVirtualizationRealmReachability) | **Put** /api/virtualizationrealms/{id}/updatereachability | 



## AddNetwork

> MinimalVirtualizationRealm AddNetwork(ctx, id).AbstractAddNetworkCloudSpaceRequest(abstractAddNetworkCloudSpaceRequest).Execute()

Allocate Network



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
    abstractAddNetworkCloudSpaceRequest := *cons3rtclient.NewAbstractAddNetworkCloudSpaceRequest(*cons3rtclient.NewNetwork("Cidr_example", []cons3rtclient.DnatRule{*cons3rtclient.NewDnatRule(false, "DnatPort_example", "DnatProtocol_example", "DnatTargetIp_example", "DnatTargetPort_example")}, []cons3rtclient.FirewallRule{*cons3rtclient.NewFirewallRule(int32(123), "Protocol_example", "RuleAction_example", "RuleDestination_example", false, "RuleSource_example")}, "Name_example"), "Subtype_example") // AbstractAddNetworkCloudSpaceRequest | The network allocation information

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.AddNetwork(context.Background(), id).AbstractAddNetworkCloudSpaceRequest(abstractAddNetworkCloudSpaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.AddNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetwork`: MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.AddNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **abstractAddNetworkCloudSpaceRequest** | [**AbstractAddNetworkCloudSpaceRequest**](AbstractAddNetworkCloudSpaceRequest.md) | The network allocation information | 

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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of virtualization realm
    projectId := "projectId_example" // string | ID of project to assign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.AddProject(context.Background(), id).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.AddProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProject`: int32
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.AddProject`: %v\n", resp)
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


## CreateTemplateSubsciption

> FullTemplateSubscription CreateTemplateSubsciption(ctx, id).RegistrationId(registrationId).Execute()

Create Template Subscription



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
    id := int32(56) // int32 | ID of virtualization realm
    registrationId := int32(56) // int32 | ID of template registration

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.CreateTemplateSubsciption(context.Background(), id).RegistrationId(registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.CreateTemplateSubsciption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplateSubsciption`: FullTemplateSubscription
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.CreateTemplateSubsciption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateSubsciptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registrationId** | **int32** | ID of template registration | 

### Return type

[**FullTemplateSubscription**](FullTemplateSubscription.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> bool DeleteNetwork(ctx, id, networkId).Deallocate(deallocate).Execute()

Delete Network



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
    networkId := int32(56) // int32 | ID of network
    deallocate := true // bool | Attempt to delete all back-end resources (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.DeleteNetwork(context.Background(), id, networkId).Deallocate(deallocate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.DeleteNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetwork`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 
**networkId** | **int32** | ID of network | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deallocate** | **bool** | Attempt to delete all back-end resources | [default to false]

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


## DeleteTemplateSubscription

> bool DeleteTemplateSubscription(ctx, id, subscriptionId).Execute()

Delete Template Subscription



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
    id := int32(56) // int32 | ID of virtualization realm
    subscriptionId := int32(56) // int32 | ID of template subscription

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.DeleteTemplateSubscription(context.Background(), id, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.DeleteTemplateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTemplateSubscription`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.DeleteTemplateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**subscriptionId** | **int32** | ID of template subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateSubscriptionRequest struct via the builder pattern


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


## DisableVirtRealmRemoteAccess

> bool DisableVirtRealmRemoteAccess(ctx, id).Execute()

Disable Remote Access



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.DisableVirtRealmRemoteAccess(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.DisableVirtRealmRemoteAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableVirtRealmRemoteAccess`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.DisableVirtRealmRemoteAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableVirtRealmRemoteAccessRequest struct via the builder pattern


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


## EnableMaintenceMode1

> bool EnableMaintenceMode1(ctx, id).Enable(enable).MaintenanceModeRequest(maintenanceModeRequest).Execute()

Update Maintenance Mode



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
    enable := true // bool | Enable or disable maintenance mode (optional)
    maintenanceModeRequest := *cons3rtclient.NewMaintenanceModeRequest() // MaintenanceModeRequest | The maintenance mode request, when enabling maintenance mode (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.EnableMaintenceMode1(context.Background(), id).Enable(enable).MaintenanceModeRequest(maintenanceModeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.EnableMaintenceMode1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableMaintenceMode1`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.EnableMaintenceMode1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableMaintenceMode1Request struct via the builder pattern


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


## EnableVirtRealmRemoteAccess

> bool EnableVirtRealmRemoteAccess(ctx, id).InstanceType(instanceType).Execute()

Enable Remote Access



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
    instanceType := "instanceType_example" // string | Remote access server instance type (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.EnableVirtRealmRemoteAccess(context.Background(), id).InstanceType(instanceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.EnableVirtRealmRemoteAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableVirtRealmRemoteAccess`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.EnableVirtRealmRemoteAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableVirtRealmRemoteAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceType** | **string** | Remote access server instance type | 

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
    resp, r, err := api_client.VirtualizationRealmsApi.GetDeploymentRunsInVirtualizationRealm(context.Background(), id).SearchType(searchType).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetDeploymentRunsInVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRunsInVirtualizationRealm`: []MinimalDeploymentRun
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetDeploymentRunsInVirtualizationRealm`: %v\n", resp)
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


## GetHostConfigurationMetrics1

> string GetHostConfigurationMetrics1(ctx, id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()

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
    id := "id_example" // string | ID of virtualization realm
    start := int64(789) // int64 | Interval start time, specified in seconds since epoch
    end := int64(789) // int64 | Interval end time, specified in seconds since epoch
    interval := int64(789) // int64 | Number of intervals (optional) (default to 1)
    intervalUnit := "intervalUnit_example" // string | Interval unit (optional) (default to "HOURS")

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetHostConfigurationMetrics1(context.Background(), id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetHostConfigurationMetrics1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostConfigurationMetrics1`: string
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetHostConfigurationMetrics1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostConfigurationMetrics1Request struct via the builder pattern


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


## GetNetwork

> Network GetNetwork(ctx, id, networkId).Execute()

Retrieve Network



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
    networkId := int32(56) // int32 | ID of network

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetNetwork(context.Background(), id, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetwork`: Network
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 
**networkId** | **int32** | ID of network | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetNetworks

> []MinimalNetwork GetNetworks(ctx, id).Execute()

List Networks



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetNetworks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworks`: []MinimalNetwork
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksRequest struct via the builder pattern


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


## GetTemplateSubscription

> FullTemplateSubscription GetTemplateSubscription(ctx, id, subscriptionId).Execute()

Retrieve Template Subscription



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
    id := int32(56) // int32 | ID of virtualization realm
    subscriptionId := int32(56) // int32 | ID of template subscription

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetTemplateSubscription(context.Background(), id, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetTemplateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateSubscription`: FullTemplateSubscription
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetTemplateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**subscriptionId** | **int32** | ID of template subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FullTemplateSubscription**](FullTemplateSubscription.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesInVirtualizationRealm

> []MinimalCons3rtTemplateData GetTemplatesInVirtualizationRealm(ctx, id).IncludeRegistrations(includeRegistrations).IncludeSubscriptions(includeSubscriptions).Execute()

List Templates



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
    id := int32(56) // int32 | ID of virtualization realm
    includeRegistrations := true // bool | Include templates registered to the virtualization realm (optional) (default to true)
    includeSubscriptions := true // bool | Include templates subscribed to by the virtualization realm (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetTemplatesInVirtualizationRealm(context.Background(), id).IncludeRegistrations(includeRegistrations).IncludeSubscriptions(includeSubscriptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetTemplatesInVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesInVirtualizationRealm`: []MinimalCons3rtTemplateData
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetTemplatesInVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesInVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRegistrations** | **bool** | Include templates registered to the virtualization realm | [default to true]
 **includeSubscriptions** | **bool** | Include templates subscribed to by the virtualization realm | [default to false]

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


## GetUnregisteredNetworks

> []Subnet GetUnregisteredNetworks(ctx, id).Execute()

List Unregistered Networks



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetUnregisteredNetworks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetUnregisteredNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnregisteredNetworks`: []Subnet
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetUnregisteredNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnregisteredNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Subnet**](Subnet.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineCountMetrics1

> string GetVirtualMachineCountMetrics1(ctx, id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()

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
    id := "id_example" // string | ID of virtualization realm
    start := int64(789) // int64 | Interval start time, specified in seconds since epoch
    end := int64(789) // int64 | Interval end time, specified in seconds since epoch
    interval := int64(789) // int64 | Number of intervals (optional) (default to 1)
    intervalUnit := "intervalUnit_example" // string | Interval unit (optional) (default to "HOURS")

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetVirtualMachineCountMetrics1(context.Background(), id).Start(start).End(end).Interval(interval).IntervalUnit(intervalUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetVirtualMachineCountMetrics1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineCountMetrics1`: string
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetVirtualMachineCountMetrics1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineCountMetrics1Request struct via the builder pattern


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


## GetVirtualizationRealm

> FullVirtualizationRealm GetVirtualizationRealm(ctx, id).Execute()

Retrieve Virtualization Realm



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetVirtualizationRealm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationRealm`: FullVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullVirtualizationRealm**](FullVirtualizationRealm.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationRealmResources

> AbstractCloudResources GetVirtualizationRealmResources(ctx, id).Execute()

Retrieve Virtualization Realm Resources



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetVirtualizationRealmResources(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetVirtualizationRealmResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationRealmResources`: AbstractCloudResources
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetVirtualizationRealmResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationRealmResourcesRequest struct via the builder pattern


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


## GetVirtualizationRealms

> []MinimalVirtualizationRealm GetVirtualizationRealms(ctx).Maxresults(maxresults).Page(page).Execute()

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
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.GetVirtualizationRealms(context.Background()).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.GetVirtualizationRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationRealms`: []MinimalVirtualizationRealm
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.GetVirtualizationRealms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationRealmsRequest struct via the builder pattern


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


## InvalidateTemplateCacheInVirtualizationRealm

> bool InvalidateTemplateCacheInVirtualizationRealm(ctx, id).Execute()

Refresh Template Cache



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
    id := int32(56) // int32 | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.InvalidateTemplateCacheInVirtualizationRealm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.InvalidateTemplateCacheInVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateTemplateCacheInVirtualizationRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.InvalidateTemplateCacheInVirtualizationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateTemplateCacheInVirtualizationRealmRequest struct via the builder pattern


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


## ListPendingTemplateSubscriptions

> []FullTemplateRegistrationForSubscription ListPendingTemplateSubscriptions(ctx, id).Execute()

List Pending Subscriptions



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
    id := int32(56) // int32 | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.ListPendingTemplateSubscriptions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.ListPendingTemplateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPendingTemplateSubscriptions`: []FullTemplateRegistrationForSubscription
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.ListPendingTemplateSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPendingTemplateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FullTemplateRegistrationForSubscription**](FullTemplateRegistrationForSubscription.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> []MinimalProject ListProjects(ctx, id).Maxresults(maxresults).Page(page).Execute()

List Projects



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
    maxresults := int64(789) // int64 | Maximum number of results to return (optional) (default to 40)
    page := int64(789) // int64 | Requested page number (optional) (default to 0)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.ListProjects(context.Background(), id).Maxresults(maxresults).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.ListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjects`: []MinimalProject
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.ListProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


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


## ListTemplateRegistrations

> []MinimalTemplateRegistration ListTemplateRegistrations(ctx, id).Execute()

List Template Registrations



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
    id := int32(56) // int32 | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.ListTemplateRegistrations(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.ListTemplateRegistrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplateRegistrations`: []MinimalTemplateRegistration
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.ListTemplateRegistrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MinimalTemplateRegistration**](MinimalTemplateRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateSubscriptions

> []MinimalTemplateSubscription ListTemplateSubscriptions(ctx, id).Execute()

List Template Subscriptions



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
    id := int32(56) // int32 | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.ListTemplateSubscriptions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.ListTemplateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplateSubscriptions`: []MinimalTemplateSubscription
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.ListTemplateSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MinimalTemplateSubscription**](MinimalTemplateSubscription.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUnregisteredTemplates

> []Cons3rtTemplateTagData ListUnregisteredTemplates(ctx, id).Execute()

List Unregistered Templates



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
    id := int32(56) // int32 | ID of virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.ListUnregisteredTemplates(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.ListUnregisteredTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUnregisteredTemplates`: []Cons3rtTemplateTagData
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.ListUnregisteredTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUnregisteredTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Cons3rtTemplateTagData**](Cons3rtTemplateTagData.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNetwork

> bool RegisterNetwork(ctx, id, networkIdentifier).Execute()

Register Network



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
    networkIdentifier := "networkIdentifier_example" // string | Back-end network identifier

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.RegisterNetwork(context.Background(), id, networkIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.RegisterNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterNetwork`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.RegisterNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 
**networkIdentifier** | **string** | Back-end network identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNetworkRequest struct via the builder pattern


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


## RegisterTemplate

> FullTemplateRegistration RegisterTemplate(ctx, id).InputRegisterTemplateObject(inputRegisterTemplateObject).Execute()

Register Template



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
    id := int32(56) // int32 | ID of virtualization realm
    inputRegisterTemplateObject := *cons3rtclient.NewInputRegisterTemplateObject(*cons3rtclient.NewInputCons3rtTemplateData("VirtRealmTemplateName_example", "OperatingSystem_example")) // InputRegisterTemplateObject | The template registration data (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.RegisterTemplate(context.Background(), id).InputRegisterTemplateObject(inputRegisterTemplateObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.RegisterTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterTemplate`: FullTemplateRegistration
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.RegisterTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputRegisterTemplateObject** | [**InputRegisterTemplateObject**](InputRegisterTemplateObject.md) | The template registration data | 

### Return type

[**FullTemplateRegistration**](FullTemplateRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProject

> bool RemoveProject(ctx, id).ProjectId(projectId).Execute()

Unassign Project



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
    projectId := "projectId_example" // string | ID of project to unassign

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.RemoveProject(context.Background(), id).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.RemoveProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveProject`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.RemoveProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | ID of project to unassign | 

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


## RetrieveTemplateRegistration

> FullTemplateRegistration RetrieveTemplateRegistration(ctx, id, registrationId).Execute()

Retrieve Template Registration



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
    id := int32(56) // int32 | ID of virtualization realm
    registrationId := int32(56) // int32 | ID of template registration

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.RetrieveTemplateRegistration(context.Background(), id, registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.RetrieveTemplateRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTemplateRegistration`: FullTemplateRegistration
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.RetrieveTemplateRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**registrationId** | **int32** | ID of template registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTemplateRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FullTemplateRegistration**](FullTemplateRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVirtualizationRealmActive

> bool SetVirtualizationRealmActive(ctx, id).Activate(activate).Execute()

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
    id := "id_example" // string | ID of virtualization realm
    activate := true // bool | Activate or deactivate virtualization realm

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.SetVirtualizationRealmActive(context.Background(), id).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.SetVirtualizationRealmActive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetVirtualizationRealmActive`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.SetVirtualizationRealmActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVirtualizationRealmActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activate** | **bool** | Activate or deactivate virtualization realm | 

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


## ShareTemplateRegistration

> bool ShareTemplateRegistration(ctx, id, registrationId).TargetRealmIds(targetRealmIds).Execute()

Share Template



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
    id := int32(56) // int32 | ID of virtualization realm
    registrationId := int32(56) // int32 | ID of template registration
    targetRealmIds := []int32{int32(123)} // []int32 | ID(s) of external virtualization realms to share template registration with

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.ShareTemplateRegistration(context.Background(), id, registrationId).TargetRealmIds(targetRealmIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.ShareTemplateRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShareTemplateRegistration`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.ShareTemplateRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**registrationId** | **int32** | ID of template registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareTemplateRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetRealmIds** | **[]int32** | ID(s) of external virtualization realms to share template registration with | 

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


## UnregisterTemplate

> bool UnregisterTemplate(ctx, id, registrationId).InputUnregisterTemplateObject(inputUnregisterTemplateObject).Execute()

Delete Template Registration



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
    id := int32(56) // int32 | ID of virtualization realm
    registrationId := int32(56) // int32 | ID of template registration
    inputUnregisterTemplateObject := *cons3rtclient.NewInputUnregisterTemplateObject(false) // InputUnregisterTemplateObject | The deletion settings

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UnregisterTemplate(context.Background(), id, registrationId).InputUnregisterTemplateObject(inputUnregisterTemplateObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UnregisterTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterTemplate`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UnregisterTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**registrationId** | **int32** | ID of template registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputUnregisterTemplateObject** | [**InputUnregisterTemplateObject**](InputUnregisterTemplateObject.md) | The deletion settings | 

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


## UnshareTemplateRegistration

> bool UnshareTemplateRegistration(ctx, id, registrationId).TargetRealmId(targetRealmId).Execute()

Unshare Template



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
    id := int32(56) // int32 | ID of virtualization realm
    registrationId := int32(56) // int32 | ID of template registration
    targetRealmId := int32(56) // int32 | ID of external virtualization realm to revoke template registration access for

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UnshareTemplateRegistration(context.Background(), id, registrationId).TargetRealmId(targetRealmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UnshareTemplateRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnshareTemplateRegistration`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UnshareTemplateRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**registrationId** | **int32** | ID of template registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareTemplateRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetRealmId** | **int32** | ID of external virtualization realm to revoke template registration access for | 

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


## UpdateTemplateRegistration

> bool UpdateTemplateRegistration(ctx, id, registrationId).InputCons3rtTemplateData(inputCons3rtTemplateData).Offline(offline).Execute()

Update Template Registration



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
    id := int32(56) // int32 | ID of virtualization realm
    registrationId := int32(56) // int32 | ID of template registration
    inputCons3rtTemplateData := *cons3rtclient.NewInputCons3rtTemplateData("VirtRealmTemplateName_example", "OperatingSystem_example") // InputCons3rtTemplateData | The modified template registration data
    offline := true // bool | The desired template registration state (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UpdateTemplateRegistration(context.Background(), id, registrationId).InputCons3rtTemplateData(inputCons3rtTemplateData).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UpdateTemplateRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplateRegistration`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UpdateTemplateRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**registrationId** | **int32** | ID of template registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputCons3rtTemplateData** | [**InputCons3rtTemplateData**](InputCons3rtTemplateData.md) | The modified template registration data | 
 **offline** | **bool** | The desired template registration state | [default to false]

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


## UpdateTemplateSubscription

> bool UpdateTemplateSubscription(ctx, id, subscriptionId).InputTemplateSubscription(inputTemplateSubscription).Offline(offline).Execute()

Update Template Subscription



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
    id := int32(56) // int32 | ID of virtualization realm
    subscriptionId := int32(56) // int32 | ID of template subscription
    inputTemplateSubscription := *cons3rtclient.NewInputTemplateSubscription() // InputTemplateSubscription | The modified template subscription data
    offline := true // bool | The desired template subscription state (optional) (default to false)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UpdateTemplateSubscription(context.Background(), id, subscriptionId).InputTemplateSubscription(inputTemplateSubscription).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UpdateTemplateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplateSubscription`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UpdateTemplateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of virtualization realm | 
**subscriptionId** | **int32** | ID of template subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputTemplateSubscription** | [**InputTemplateSubscription**](InputTemplateSubscription.md) | The modified template subscription data | 
 **offline** | **bool** | The desired template subscription state | [default to false]

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


## UpdateVirtRealmRemoteAccessConfig

> bool UpdateVirtRealmRemoteAccessConfig(ctx, id).RemoteAccessConfig(remoteAccessConfig).Execute()

Update Remote Access



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
    remoteAccessConfig := *cons3rtclient.NewRemoteAccessConfig("RemoteAccessIpAddress_example", int32(123), "InstanceType_example") // RemoteAccessConfig | The updated remote access configuration

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UpdateVirtRealmRemoteAccessConfig(context.Background(), id).RemoteAccessConfig(remoteAccessConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UpdateVirtRealmRemoteAccessConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtRealmRemoteAccessConfig`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UpdateVirtRealmRemoteAccessConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtRealmRemoteAccessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteAccessConfig** | [**RemoteAccessConfig**](RemoteAccessConfig.md) | The updated remote access configuration | 

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
    cons3rtclient "./gocons3rt"
)

func main() {
    id := "id_example" // string | ID of virtualization realm
    inputVRAdminVirtualizationRealm := *cons3rtclient.NewInputVRAdminVirtualizationRealm("VirtualizationRealmType_example", "Cidr_example", "Description_example", "Name_example") // InputVRAdminVirtualizationRealm | The updated Virtualization Realm data (optional)

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UpdateVirtualizationRealm(context.Background(), id).InputVRAdminVirtualizationRealm(inputVRAdminVirtualizationRealm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UpdateVirtualizationRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationRealm`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UpdateVirtualizationRealm`: %v\n", resp)
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


## UpdateVirtualizationRealmReachability

> bool UpdateVirtualizationRealmReachability(ctx, id).Execute()



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

    configuration := cons3rtclient.NewConfiguration()
    api_client := cons3rtclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationRealmsApi.UpdateVirtualizationRealmReachability(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationRealmsApi.UpdateVirtualizationRealmReachability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationRealmReachability`: bool
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationRealmsApi.UpdateVirtualizationRealmReachability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of virtualization realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationRealmReachabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [Username](../README.md#Username)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

