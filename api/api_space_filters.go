/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.2.0
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

// SpaceFiltersAPIService SpaceFiltersAPI service
type SpaceFiltersAPIService service

type ApiCreateFilterForSpaceRequest struct {
	ctx                       context.Context
	ApiService                *SpaceFiltersAPIService
	createFilterForSpaceInput *CreateFilterForSpaceInput
}

func (r ApiCreateFilterForSpaceRequest) CreateFilterForSpaceInput(
	createFilterForSpaceInput CreateFilterForSpaceInput,
) ApiCreateFilterForSpaceRequest {
	r.createFilterForSpaceInput = &createFilterForSpaceInput
	return r
}

func (r ApiCreateFilterForSpaceRequest) Execute() (*CreateFilterForSpace200Response, *http.Response, error) {
	return r.ApiService.CreateFilterForSpaceExecute(r)
}

/*
CreateFilterForSpace Create Filter for Space

Creates a filter for a space. A space filter applies to events coming from all Sources connected to a space.

• This endpoint is in **Beta** testing.  Please submit any feedback by sending an email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Space Filters feature enabled. Please reach out to your customer success manager for more information.

• When called, this endpoint may generate the `Filter Created` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateFilterForSpaceRequest
*/
func (a *SpaceFiltersAPIService) CreateFilterForSpace(
	ctx context.Context,
) ApiCreateFilterForSpaceRequest {
	return ApiCreateFilterForSpaceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateFilterForSpace200Response
func (a *SpaceFiltersAPIService) CreateFilterForSpaceExecute(
	r ApiCreateFilterForSpaceRequest,
) (*CreateFilterForSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateFilterForSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpaceFiltersAPIService.CreateFilterForSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFilterForSpaceInput == nil {
		return localVarReturnValue, nil, reportError(
			"createFilterForSpaceInput is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{
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
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
		"application/json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createFilterForSpaceInput
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

type ApiDeleteFilterByIdRequest struct {
	ctx        context.Context
	ApiService *SpaceFiltersAPIService
	id         string
}

func (r ApiDeleteFilterByIdRequest) Execute() (*DeleteFilterById200Response, *http.Response, error) {
	return r.ApiService.DeleteFilterByIdExecute(r)
}

/*
DeleteFilterById Delete Filter By Id

Deletes a filter by id.

• This endpoint is in **Beta** testing.  Please submit any feedback by sending an email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Space Filters feature enabled. Please reach out to your customer success manager for more information.

• When called, this endpoint may generate the `Filter Deleted` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiDeleteFilterByIdRequest
*/
func (a *SpaceFiltersAPIService) DeleteFilterById(
	ctx context.Context,
	id string,
) ApiDeleteFilterByIdRequest {
	return ApiDeleteFilterByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DeleteFilterById200Response
func (a *SpaceFiltersAPIService) DeleteFilterByIdExecute(
	r ApiDeleteFilterByIdRequest,
) (*DeleteFilterById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteFilterById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpaceFiltersAPIService.DeleteFilterById",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filters/{id}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"id"+"}",
		url.PathEscape(parameterValueToString(r.id, "id")),
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
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
		"application/json",
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

type ApiGetFilterByIdRequest struct {
	ctx        context.Context
	ApiService *SpaceFiltersAPIService
	id         string
}

func (r ApiGetFilterByIdRequest) Execute() (*GetFilterById200Response, *http.Response, error) {
	return r.ApiService.GetFilterByIdExecute(r)
}

/*
GetFilterById Get Filter By Id

Gets a filter by id.

• This endpoint is in **Beta** testing.  Please submit any feedback by sending an email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Space Filters feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetFilterByIdRequest
*/
func (a *SpaceFiltersAPIService) GetFilterById(
	ctx context.Context,
	id string,
) ApiGetFilterByIdRequest {
	return ApiGetFilterByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetFilterById200Response
func (a *SpaceFiltersAPIService) GetFilterByIdExecute(
	r ApiGetFilterByIdRequest,
) (*GetFilterById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetFilterById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpaceFiltersAPIService.GetFilterById",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filters/{id}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"id"+"}",
		url.PathEscape(parameterValueToString(r.id, "id")),
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
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
		"application/json",
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

type ApiListFiltersForSpaceRequest struct {
	ctx           context.Context
	ApiService    *SpaceFiltersAPIService
	integrationId *string
	pagination    *ListFiltersPaginationInput
}

// The Space Id for which to fetch filters  This parameter exists in beta.
func (r ApiListFiltersForSpaceRequest) IntegrationId(
	integrationId string,
) ApiListFiltersForSpaceRequest {
	r.integrationId = &integrationId
	return r
}

// Pagination parameters.  This parameter exists in beta.
func (r ApiListFiltersForSpaceRequest) Pagination(
	pagination ListFiltersPaginationInput,
) ApiListFiltersForSpaceRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListFiltersForSpaceRequest) Execute() (*ListFiltersForSpace200Response, *http.Response, error) {
	return r.ApiService.ListFiltersForSpaceExecute(r)
}

/*
ListFiltersForSpace List Filters for Space

Lists filters for a space.

• This endpoint is in **Beta** testing.  Please submit any feedback by sending an email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Space Filters feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListFiltersForSpaceRequest
*/
func (a *SpaceFiltersAPIService) ListFiltersForSpace(
	ctx context.Context,
) ApiListFiltersForSpaceRequest {
	return ApiListFiltersForSpaceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListFiltersForSpace200Response
func (a *SpaceFiltersAPIService) ListFiltersForSpaceExecute(
	r ApiListFiltersForSpaceRequest,
) (*ListFiltersForSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListFiltersForSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpaceFiltersAPIService.ListFiltersForSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.integrationId == nil {
		return localVarReturnValue, nil, reportError(
			"integrationId is required and must be specified",
		)
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "integrationId", r.integrationId, "")
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
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
		"application/json",
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

type ApiUpdateFilterByIdRequest struct {
	ctx                   context.Context
	ApiService            *SpaceFiltersAPIService
	id                    string
	updateFilterByIdInput *UpdateFilterByIdInput
}

func (r ApiUpdateFilterByIdRequest) UpdateFilterByIdInput(
	updateFilterByIdInput UpdateFilterByIdInput,
) ApiUpdateFilterByIdRequest {
	r.updateFilterByIdInput = &updateFilterByIdInput
	return r
}

func (r ApiUpdateFilterByIdRequest) Execute() (*UpdateFilterById200Response, *http.Response, error) {
	return r.ApiService.UpdateFilterByIdExecute(r)
}

/*
UpdateFilterById Update Filter By Id

Updates a filter by id and replaces the existing filter.

• This endpoint is in **Beta** testing.  Please submit any feedback by sending an email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Space Filters feature enabled. Please reach out to your customer success manager for more information.

• When called, this endpoint may generate the `Filter Updated` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiUpdateFilterByIdRequest
*/
func (a *SpaceFiltersAPIService) UpdateFilterById(
	ctx context.Context,
	id string,
) ApiUpdateFilterByIdRequest {
	return ApiUpdateFilterByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UpdateFilterById200Response
func (a *SpaceFiltersAPIService) UpdateFilterByIdExecute(
	r ApiUpdateFilterByIdRequest,
) (*UpdateFilterById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateFilterById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpaceFiltersAPIService.UpdateFilterById",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filters/{id}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"id"+"}",
		url.PathEscape(parameterValueToString(r.id, "id")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateFilterByIdInput == nil {
		return localVarReturnValue, nil, reportError(
			"updateFilterByIdInput is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{
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
		"application/vnd.segment.v1beta+json",
		"application/vnd.segment.v1alpha+json",
		"application/json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateFilterByIdInput
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
