/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// MonthlyTrackedUsersAPIService MonthlyTrackedUsersAPI service
type MonthlyTrackedUsersAPIService service

type ApiGetDailyPerSourceMTUUsageRequest struct {
	ctx        context.Context
	ApiService *MonthlyTrackedUsersAPIService
	period     *string
	pagination *PaginationInput
}

// The start of the usage month, in the ISO-8601 format.  This parameter exists in v1.
func (r ApiGetDailyPerSourceMTUUsageRequest) Period(
	period string,
) ApiGetDailyPerSourceMTUUsageRequest {
	r.period = &period
	return r
}

// Pagination input for per Source MTU counts.  This parameter exists in v1.
func (r ApiGetDailyPerSourceMTUUsageRequest) Pagination(
	pagination PaginationInput,
) ApiGetDailyPerSourceMTUUsageRequest {
	r.pagination = &pagination
	return r
}

func (r ApiGetDailyPerSourceMTUUsageRequest) Execute() (*GetDailyPerSourceMTUUsage200Response, *http.Response, error) {
	return r.ApiService.GetDailyPerSourceMTUUsageExecute(r)
}

/*
GetDailyPerSourceMTUUsage Get Daily Per Source MTU Usage

Provides daily cumulative per-source MTU counts for a usage period.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDailyPerSourceMTUUsageRequest
*/
func (a *MonthlyTrackedUsersAPIService) GetDailyPerSourceMTUUsage(
	ctx context.Context,
) ApiGetDailyPerSourceMTUUsageRequest {
	return ApiGetDailyPerSourceMTUUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDailyPerSourceMTUUsage200Response
func (a *MonthlyTrackedUsersAPIService) GetDailyPerSourceMTUUsageExecute(
	r ApiGetDailyPerSourceMTUUsageRequest,
) (*GetDailyPerSourceMTUUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDailyPerSourceMTUUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"MonthlyTrackedUsersAPIService.GetDailyPerSourceMTUUsage",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/mtu/sources/daily"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.period == nil {
		return localVarReturnValue, nil, reportError("period is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	if r.pagination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination", r.pagination, "")
	}
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiGetDailyWorkspaceMTUUsageRequest struct {
	ctx        context.Context
	ApiService *MonthlyTrackedUsersAPIService
	period     *string
	pagination *PaginationInput
}

// The start of the usage month, in the ISO-8601 format.  This parameter exists in v1.
func (r ApiGetDailyWorkspaceMTUUsageRequest) Period(
	period string,
) ApiGetDailyWorkspaceMTUUsageRequest {
	r.period = &period
	return r
}

// Pagination input for Workspace MTU counts.  This parameter exists in v1.
func (r ApiGetDailyWorkspaceMTUUsageRequest) Pagination(
	pagination PaginationInput,
) ApiGetDailyWorkspaceMTUUsageRequest {
	r.pagination = &pagination
	return r
}

func (r ApiGetDailyWorkspaceMTUUsageRequest) Execute() (*GetDailyWorkspaceMTUUsage200Response, *http.Response, error) {
	return r.ApiService.GetDailyWorkspaceMTUUsageExecute(r)
}

/*
GetDailyWorkspaceMTUUsage Get Daily Workspace MTU Usage

Provides daily cumulative MTU counts for a usage period.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDailyWorkspaceMTUUsageRequest
*/
func (a *MonthlyTrackedUsersAPIService) GetDailyWorkspaceMTUUsage(
	ctx context.Context,
) ApiGetDailyWorkspaceMTUUsageRequest {
	return ApiGetDailyWorkspaceMTUUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDailyWorkspaceMTUUsage200Response
func (a *MonthlyTrackedUsersAPIService) GetDailyWorkspaceMTUUsageExecute(
	r ApiGetDailyWorkspaceMTUUsageRequest,
) (*GetDailyWorkspaceMTUUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDailyWorkspaceMTUUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"MonthlyTrackedUsersAPIService.GetDailyWorkspaceMTUUsage",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/mtu/daily"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.period == nil {
		return localVarReturnValue, nil, reportError("period is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	if r.pagination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination", r.pagination, "")
	}
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
