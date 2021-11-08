/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cons3rt

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// TemplatesApiService TemplatesApi service
type TemplatesApiService service

type ApiListVirtualizationRealmTemplatesRequest struct {
	ctx                   _context.Context
	ApiService            *TemplatesApiService
	virtualizationRealmId *int32
	includeRegistrations  *bool
	includeSubscriptions  *bool
}

// Virtualization Realm ID to filter by
func (r ApiListVirtualizationRealmTemplatesRequest) VirtualizationRealmId(virtualizationRealmId int32) ApiListVirtualizationRealmTemplatesRequest {
	r.virtualizationRealmId = &virtualizationRealmId
	return r
}

// Include templates registered to this virtualization realm
func (r ApiListVirtualizationRealmTemplatesRequest) IncludeRegistrations(includeRegistrations bool) ApiListVirtualizationRealmTemplatesRequest {
	r.includeRegistrations = &includeRegistrations
	return r
}

// Include templates this virtualization realm is subscribed to
func (r ApiListVirtualizationRealmTemplatesRequest) IncludeSubscriptions(includeSubscriptions bool) ApiListVirtualizationRealmTemplatesRequest {
	r.includeSubscriptions = &includeSubscriptions
	return r
}

func (r ApiListVirtualizationRealmTemplatesRequest) Execute() ([]MinimalCons3rtTemplateData, *_nethttp.Response, error) {
	return r.ApiService.ListVirtualizationRealmTemplatesExecute(r)
}

/*
ListVirtualizationRealmTemplates List all Templates

Returns a collection of the user's available Templates matching a specified query.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListVirtualizationRealmTemplatesRequest
*/
func (a *TemplatesApiService) ListVirtualizationRealmTemplates(ctx _context.Context) ApiListVirtualizationRealmTemplatesRequest {
	return ApiListVirtualizationRealmTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []MinimalCons3rtTemplateData
func (a *TemplatesApiService) ListVirtualizationRealmTemplatesExecute(r ApiListVirtualizationRealmTemplatesRequest) ([]MinimalCons3rtTemplateData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MinimalCons3rtTemplateData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesApiService.ListVirtualizationRealmTemplates")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.virtualizationRealmId == nil {
		return localVarReturnValue, nil, reportError("virtualizationRealmId is required and must be specified")
	}

	localVarQueryParams.Add("virtualization_realm_id", parameterToString(*r.virtualizationRealmId, ""))
	if r.includeRegistrations != nil {
		localVarQueryParams.Add("include_registrations", parameterToString(*r.includeRegistrations, ""))
	}
	if r.includeSubscriptions != nil {
		localVarQueryParams.Add("include_subscriptions", parameterToString(*r.includeSubscriptions, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Username"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["username"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
