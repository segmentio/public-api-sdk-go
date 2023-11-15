/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.1.0
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
	"strings"
)

// SelectiveSyncAPIService SelectiveSyncAPI service
type SelectiveSyncAPIService service

type ApiGetAdvancedSyncScheduleFromWarehouseRequest struct {
	ctx         context.Context
	ApiService  *SelectiveSyncAPIService
	warehouseId string
}

func (r ApiGetAdvancedSyncScheduleFromWarehouseRequest) Execute() (*GetAdvancedSyncScheduleFromWarehouse200Response, *http.Response, error) {
	return r.ApiService.GetAdvancedSyncScheduleFromWarehouseExecute(r)
}

/*
GetAdvancedSyncScheduleFromWarehouse Get Advanced Sync Schedule from Warehouse

Returns the advanced sync schedule for a Warehouse.

The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param warehouseId
	@return ApiGetAdvancedSyncScheduleFromWarehouseRequest
*/
func (a *SelectiveSyncAPIService) GetAdvancedSyncScheduleFromWarehouse(
	ctx context.Context,
	warehouseId string,
) ApiGetAdvancedSyncScheduleFromWarehouseRequest {
	return ApiGetAdvancedSyncScheduleFromWarehouseRequest{
		ApiService:  a,
		ctx:         ctx,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return GetAdvancedSyncScheduleFromWarehouse200Response
func (a *SelectiveSyncAPIService) GetAdvancedSyncScheduleFromWarehouseExecute(
	r ApiGetAdvancedSyncScheduleFromWarehouseRequest,
) (*GetAdvancedSyncScheduleFromWarehouse200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAdvancedSyncScheduleFromWarehouse200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SelectiveSyncAPIService.GetAdvancedSyncScheduleFromWarehouse",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/warehouses/{warehouseId}/advanced-sync-schedule"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListSelectiveSyncsFromWarehouseAndSourceRequest struct {
	ctx         context.Context
	ApiService  *SelectiveSyncAPIService
	warehouseId string
	sourceId    string
	pagination  *PaginationInput
}

// Defines the pagination parameters.  This parameter exists in v1.
func (r ApiListSelectiveSyncsFromWarehouseAndSourceRequest) Pagination(
	pagination PaginationInput,
) ApiListSelectiveSyncsFromWarehouseAndSourceRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListSelectiveSyncsFromWarehouseAndSourceRequest) Execute() (*ListSelectiveSyncsFromWarehouseAndSource200Response, *http.Response, error) {
	return r.ApiService.ListSelectiveSyncsFromWarehouseAndSourceExecute(r)
}

/*
ListSelectiveSyncsFromWarehouseAndSource List Selective Syncs from Warehouse And Source

Returns the schema for a Warehouse, including Sources, Collections, and Properties.

The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param warehouseId
	@param sourceId
	@return ApiListSelectiveSyncsFromWarehouseAndSourceRequest
*/
func (a *SelectiveSyncAPIService) ListSelectiveSyncsFromWarehouseAndSource(
	ctx context.Context,
	warehouseId string,
	sourceId string,
) ApiListSelectiveSyncsFromWarehouseAndSourceRequest {
	return ApiListSelectiveSyncsFromWarehouseAndSourceRequest{
		ApiService:  a,
		ctx:         ctx,
		warehouseId: warehouseId,
		sourceId:    sourceId,
	}
}

// Execute executes the request
//
//	@return ListSelectiveSyncsFromWarehouseAndSource200Response
func (a *SelectiveSyncAPIService) ListSelectiveSyncsFromWarehouseAndSourceExecute(
	r ApiListSelectiveSyncsFromWarehouseAndSourceRequest,
) (*ListSelectiveSyncsFromWarehouseAndSource200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSelectiveSyncsFromWarehouseAndSource200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SelectiveSyncAPIService.ListSelectiveSyncsFromWarehouseAndSource",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/warehouses/{warehouseId}/connected-sources/{sourceId}/selective-syncs"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"sourceId"+"}",
		url.PathEscape(parameterValueToString(r.sourceId, "sourceId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "pagination", r.pagination, "")
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

type ApiListSyncsFromWarehouseRequest struct {
	ctx         context.Context
	ApiService  *SelectiveSyncAPIService
	warehouseId string
	pagination  *PaginationInput
}

// Defines the pagination parameters.  This parameter exists in v1.
func (r ApiListSyncsFromWarehouseRequest) Pagination(
	pagination PaginationInput,
) ApiListSyncsFromWarehouseRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListSyncsFromWarehouseRequest) Execute() (*ListSyncsFromWarehouse200Response, *http.Response, error) {
	return r.ApiService.ListSyncsFromWarehouseExecute(r)
}

/*
ListSyncsFromWarehouse List Syncs from Warehouse

Returns the overview of syncs for every Source connected to a Warehouse.

The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param warehouseId
	@return ApiListSyncsFromWarehouseRequest
*/
func (a *SelectiveSyncAPIService) ListSyncsFromWarehouse(
	ctx context.Context,
	warehouseId string,
) ApiListSyncsFromWarehouseRequest {
	return ApiListSyncsFromWarehouseRequest{
		ApiService:  a,
		ctx:         ctx,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return ListSyncsFromWarehouse200Response
func (a *SelectiveSyncAPIService) ListSyncsFromWarehouseExecute(
	r ApiListSyncsFromWarehouseRequest,
) (*ListSyncsFromWarehouse200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSyncsFromWarehouse200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SelectiveSyncAPIService.ListSyncsFromWarehouse",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/warehouses/{warehouseId}/syncs"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "pagination", r.pagination, "")
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

type ApiListSyncsFromWarehouseAndSourceRequest struct {
	ctx         context.Context
	ApiService  *SelectiveSyncAPIService
	warehouseId string
	sourceId    string
	pagination  *PaginationInput
}

// Defines the pagination parameters.  This parameter exists in v1.
func (r ApiListSyncsFromWarehouseAndSourceRequest) Pagination(
	pagination PaginationInput,
) ApiListSyncsFromWarehouseAndSourceRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListSyncsFromWarehouseAndSourceRequest) Execute() (*ListSyncsFromWarehouseAndSource200Response, *http.Response, error) {
	return r.ApiService.ListSyncsFromWarehouseAndSourceExecute(r)
}

/*
ListSyncsFromWarehouseAndSource List Syncs from Warehouse And Source

Returns the overview of syncs for a Source connected to a Warehouse.

The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param warehouseId
	@param sourceId
	@return ApiListSyncsFromWarehouseAndSourceRequest
*/
func (a *SelectiveSyncAPIService) ListSyncsFromWarehouseAndSource(
	ctx context.Context,
	warehouseId string,
	sourceId string,
) ApiListSyncsFromWarehouseAndSourceRequest {
	return ApiListSyncsFromWarehouseAndSourceRequest{
		ApiService:  a,
		ctx:         ctx,
		warehouseId: warehouseId,
		sourceId:    sourceId,
	}
}

// Execute executes the request
//
//	@return ListSyncsFromWarehouseAndSource200Response
func (a *SelectiveSyncAPIService) ListSyncsFromWarehouseAndSourceExecute(
	r ApiListSyncsFromWarehouseAndSourceRequest,
) (*ListSyncsFromWarehouseAndSource200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSyncsFromWarehouseAndSource200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SelectiveSyncAPIService.ListSyncsFromWarehouseAndSource",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/warehouses/{warehouseId}/connected-sources/{sourceId}/syncs"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"sourceId"+"}",
		url.PathEscape(parameterValueToString(r.sourceId, "sourceId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "pagination", r.pagination, "")
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

type ApiReplaceAdvancedSyncScheduleForWarehouseRequest struct {
	ctx                                            context.Context
	ApiService                                     *SelectiveSyncAPIService
	warehouseId                                    string
	replaceAdvancedSyncScheduleForWarehouseV1Input *ReplaceAdvancedSyncScheduleForWarehouseV1Input
}

func (r ApiReplaceAdvancedSyncScheduleForWarehouseRequest) ReplaceAdvancedSyncScheduleForWarehouseV1Input(
	replaceAdvancedSyncScheduleForWarehouseV1Input ReplaceAdvancedSyncScheduleForWarehouseV1Input,
) ApiReplaceAdvancedSyncScheduleForWarehouseRequest {
	r.replaceAdvancedSyncScheduleForWarehouseV1Input = &replaceAdvancedSyncScheduleForWarehouseV1Input
	return r
}

func (r ApiReplaceAdvancedSyncScheduleForWarehouseRequest) Execute() (*ReplaceAdvancedSyncScheduleForWarehouse200Response, *http.Response, error) {
	return r.ApiService.ReplaceAdvancedSyncScheduleForWarehouseExecute(r)
}

/*
ReplaceAdvancedSyncScheduleForWarehouse Replace Advanced Sync Schedule for Warehouse

Updates the advanced sync schedule for a Warehouse, replacing the sync schedule with a new schedule.

The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param warehouseId
	@return ApiReplaceAdvancedSyncScheduleForWarehouseRequest
*/
func (a *SelectiveSyncAPIService) ReplaceAdvancedSyncScheduleForWarehouse(
	ctx context.Context,
	warehouseId string,
) ApiReplaceAdvancedSyncScheduleForWarehouseRequest {
	return ApiReplaceAdvancedSyncScheduleForWarehouseRequest{
		ApiService:  a,
		ctx:         ctx,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return ReplaceAdvancedSyncScheduleForWarehouse200Response
func (a *SelectiveSyncAPIService) ReplaceAdvancedSyncScheduleForWarehouseExecute(
	r ApiReplaceAdvancedSyncScheduleForWarehouseRequest,
) (*ReplaceAdvancedSyncScheduleForWarehouse200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplaceAdvancedSyncScheduleForWarehouse200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SelectiveSyncAPIService.ReplaceAdvancedSyncScheduleForWarehouse",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/warehouses/{warehouseId}/advanced-sync-schedule"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.replaceAdvancedSyncScheduleForWarehouseV1Input == nil {
		return localVarReturnValue, nil, reportError(
			"replaceAdvancedSyncScheduleForWarehouseV1Input is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{
		"application/vnd.segment.v1+json",
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
	}

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
	// body params
	localVarPostBody = r.replaceAdvancedSyncScheduleForWarehouseV1Input
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

type ApiUpdateSelectiveSyncForWarehouseRequest struct {
	ctx                                    context.Context
	ApiService                             *SelectiveSyncAPIService
	warehouseId                            string
	updateSelectiveSyncForWarehouseV1Input *UpdateSelectiveSyncForWarehouseV1Input
}

func (r ApiUpdateSelectiveSyncForWarehouseRequest) UpdateSelectiveSyncForWarehouseV1Input(
	updateSelectiveSyncForWarehouseV1Input UpdateSelectiveSyncForWarehouseV1Input,
) ApiUpdateSelectiveSyncForWarehouseRequest {
	r.updateSelectiveSyncForWarehouseV1Input = &updateSelectiveSyncForWarehouseV1Input
	return r
}

func (r ApiUpdateSelectiveSyncForWarehouseRequest) Execute() (*UpdateSelectiveSyncForWarehouse200Response, *http.Response, error) {
	return r.ApiService.UpdateSelectiveSyncForWarehouseExecute(r)
}

/*
UpdateSelectiveSyncForWarehouse Update Selective Sync for Warehouse

Configures the schema for a Warehouse, including Sources, Collections, and Properties.

• When called, this endpoint may generate the `Storage Destination Modified` event in the [audit trail](/tag/Audit-Trail).

The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param warehouseId
	@return ApiUpdateSelectiveSyncForWarehouseRequest
*/
func (a *SelectiveSyncAPIService) UpdateSelectiveSyncForWarehouse(
	ctx context.Context,
	warehouseId string,
) ApiUpdateSelectiveSyncForWarehouseRequest {
	return ApiUpdateSelectiveSyncForWarehouseRequest{
		ApiService:  a,
		ctx:         ctx,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return UpdateSelectiveSyncForWarehouse200Response
func (a *SelectiveSyncAPIService) UpdateSelectiveSyncForWarehouseExecute(
	r ApiUpdateSelectiveSyncForWarehouseRequest,
) (*UpdateSelectiveSyncForWarehouse200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSelectiveSyncForWarehouse200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SelectiveSyncAPIService.UpdateSelectiveSyncForWarehouse",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/warehouses/{warehouseId}/selective-sync"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSelectiveSyncForWarehouseV1Input == nil {
		return localVarReturnValue, nil, reportError(
			"updateSelectiveSyncForWarehouseV1Input is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{
		"application/vnd.segment.v1+json",
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
	}

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
	// body params
	localVarPostBody = r.updateSelectiveSyncForWarehouseV1Input
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
