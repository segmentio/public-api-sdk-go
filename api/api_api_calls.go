/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// APICallsApiService APICallsApi service
type APICallsApiService service

type ApiGetDailyPerSourceAPICallsUsageRequest struct {
	ctx        context.Context
	ApiService *APICallsApiService
	period     *string
	pagination *PaginationInput
}

// The start of the usage month in the ISO-8601 format.  This parameter exists in v1.
func (r ApiGetDailyPerSourceAPICallsUsageRequest) Period(
	period string,
) ApiGetDailyPerSourceAPICallsUsageRequest {
	r.period = &period
	return r
}

// Pagination input for per Source API calls counts.  This parameter exists in v1.
func (r ApiGetDailyPerSourceAPICallsUsageRequest) Pagination(
	pagination PaginationInput,
) ApiGetDailyPerSourceAPICallsUsageRequest {
	r.pagination = &pagination
	return r
}

func (r ApiGetDailyPerSourceAPICallsUsageRequest) Execute() (*GetDailyPerSourceAPICallsUsage200Response, *http.Response, error) {
	return r.ApiService.GetDailyPerSourceAPICallsUsageExecute(r)
}

/*
GetDailyPerSourceAPICallsUsage Get Daily Per Source API Calls Usage

Provides daily cumulative per-source API call counts for a usage period.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDailyPerSourceAPICallsUsageRequest
*/
func (a *APICallsApiService) GetDailyPerSourceAPICallsUsage(
	ctx context.Context,
) ApiGetDailyPerSourceAPICallsUsageRequest {
	return ApiGetDailyPerSourceAPICallsUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDailyPerSourceAPICallsUsage200Response
func (a *APICallsApiService) GetDailyPerSourceAPICallsUsageExecute(
	r ApiGetDailyPerSourceAPICallsUsageRequest,
) (*GetDailyPerSourceAPICallsUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDailyPerSourceAPICallsUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"APICallsApiService.GetDailyPerSourceAPICallsUsage",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/api-calls/sources/daily"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.period == nil {
		return localVarReturnValue, nil, reportError("period is required and must be specified")
	}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	localVarQueryParams.Add("pagination", parameterToString(*r.pagination, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/vnd.segment.v1+json",
		"application/json",
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		formFiles,
	)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RequestErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RequestErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RequestErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(
		&localVarReturnValue,
		localVarBody,
		localVarHTTPResponse.Header.Get("Content-Type"),
	)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDailyWorkspaceAPICallsUsageRequest struct {
	ctx        context.Context
	ApiService *APICallsApiService
	period     *string
	pagination *PaginationInput
}

// The start of the usage month in the ISO-8601 format.  This parameter exists in v1.
func (r ApiGetDailyWorkspaceAPICallsUsageRequest) Period(
	period string,
) ApiGetDailyWorkspaceAPICallsUsageRequest {
	r.period = &period
	return r
}

// Pagination input for Workspace API call counts.  This parameter exists in v1.
func (r ApiGetDailyWorkspaceAPICallsUsageRequest) Pagination(
	pagination PaginationInput,
) ApiGetDailyWorkspaceAPICallsUsageRequest {
	r.pagination = &pagination
	return r
}

func (r ApiGetDailyWorkspaceAPICallsUsageRequest) Execute() (*GetDailyWorkspaceAPICallsUsage200Response, *http.Response, error) {
	return r.ApiService.GetDailyWorkspaceAPICallsUsageExecute(r)
}

/*
GetDailyWorkspaceAPICallsUsage Get Daily Workspace API Calls Usage

Provides daily cumulative API call counts for a usage period.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDailyWorkspaceAPICallsUsageRequest
*/
func (a *APICallsApiService) GetDailyWorkspaceAPICallsUsage(
	ctx context.Context,
) ApiGetDailyWorkspaceAPICallsUsageRequest {
	return ApiGetDailyWorkspaceAPICallsUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDailyWorkspaceAPICallsUsage200Response
func (a *APICallsApiService) GetDailyWorkspaceAPICallsUsageExecute(
	r ApiGetDailyWorkspaceAPICallsUsageRequest,
) (*GetDailyWorkspaceAPICallsUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDailyWorkspaceAPICallsUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"APICallsApiService.GetDailyWorkspaceAPICallsUsage",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/api-calls/daily"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.period == nil {
		return localVarReturnValue, nil, reportError("period is required and must be specified")
	}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	localVarQueryParams.Add("pagination", parameterToString(*r.pagination, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/vnd.segment.v1+json",
		"application/json",
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		formFiles,
	)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RequestErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RequestErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RequestErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(
		&localVarReturnValue,
		localVarBody,
		localVarHTTPResponse.Header.Get("Content-Type"),
	)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
