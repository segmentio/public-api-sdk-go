/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 47.0.0
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

// ProfilesSyncAPIService ProfilesSyncAPI service
type ProfilesSyncAPIService service

type ApiCreateProfilesWarehouseRequest struct {
	ctx                               context.Context
	ApiService                        *ProfilesSyncAPIService
	spaceId                           string
	createProfilesWarehouseAlphaInput *CreateProfilesWarehouseAlphaInput
}

func (r ApiCreateProfilesWarehouseRequest) CreateProfilesWarehouseAlphaInput(
	createProfilesWarehouseAlphaInput CreateProfilesWarehouseAlphaInput,
) ApiCreateProfilesWarehouseRequest {
	r.createProfilesWarehouseAlphaInput = &createProfilesWarehouseAlphaInput
	return r
}

func (r ApiCreateProfilesWarehouseRequest) Execute() (*CreateProfilesWarehouse201Response, *http.Response, error) {
	return r.ApiService.CreateProfilesWarehouseExecute(r)
}

/*
CreateProfilesWarehouse Create Profiles Warehouse

Creates a new Profiles Warehouse.

• When called, this endpoint may generate the `Profiles Sync Warehouse Created` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@return ApiCreateProfilesWarehouseRequest
*/
func (a *ProfilesSyncAPIService) CreateProfilesWarehouse(
	ctx context.Context,
	spaceId string,
) ApiCreateProfilesWarehouseRequest {
	return ApiCreateProfilesWarehouseRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return CreateProfilesWarehouse201Response
func (a *ProfilesSyncAPIService) CreateProfilesWarehouseExecute(
	r ApiCreateProfilesWarehouseRequest,
) (*CreateProfilesWarehouse201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateProfilesWarehouse201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ProfilesSyncAPIService.CreateProfilesWarehouse",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/profiles-warehouses"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterValueToString(r.spaceId, "spaceId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createProfilesWarehouseAlphaInput == nil {
		return localVarReturnValue, nil, reportError(
			"createProfilesWarehouseAlphaInput is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.segment.v1alpha+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/vnd.segment.v1alpha+json",
		"application/json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createProfilesWarehouseAlphaInput
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

type ApiListProfilesWarehouseInSpaceRequest struct {
	ctx        context.Context
	ApiService *ProfilesSyncAPIService
	spaceId    string
	pagination *PaginationInput
}

// Defines the pagination parameters.  This parameter exists in alpha.
func (r ApiListProfilesWarehouseInSpaceRequest) Pagination(
	pagination PaginationInput,
) ApiListProfilesWarehouseInSpaceRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListProfilesWarehouseInSpaceRequest) Execute() (*ListProfilesWarehouseInSpace200Response, *http.Response, error) {
	return r.ApiService.ListProfilesWarehouseInSpaceExecute(r)
}

/*
ListProfilesWarehouseInSpace List Profiles Warehouse in Space

Lists all Profile Warehouses for a given space id.

• When called, this endpoint may generate the `Profiles Sync Warehouse Retrieved` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@return ApiListProfilesWarehouseInSpaceRequest
*/
func (a *ProfilesSyncAPIService) ListProfilesWarehouseInSpace(
	ctx context.Context,
	spaceId string,
) ApiListProfilesWarehouseInSpaceRequest {
	return ApiListProfilesWarehouseInSpaceRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return ListProfilesWarehouseInSpace200Response
func (a *ProfilesSyncAPIService) ListProfilesWarehouseInSpaceExecute(
	r ApiListProfilesWarehouseInSpaceRequest,
) (*ListProfilesWarehouseInSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListProfilesWarehouseInSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ProfilesSyncAPIService.ListProfilesWarehouseInSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/profiles-warehouses"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterValueToString(r.spaceId, "spaceId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListSelectiveSyncsFromWarehouseAndSpaceRequest struct {
	ctx         context.Context
	ApiService  *ProfilesSyncAPIService
	spaceId     string
	warehouseId string
	pagination  *PaginationInput
}

// Defines the pagination parameters.  This parameter exists in alpha.
func (r ApiListSelectiveSyncsFromWarehouseAndSpaceRequest) Pagination(
	pagination PaginationInput,
) ApiListSelectiveSyncsFromWarehouseAndSpaceRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListSelectiveSyncsFromWarehouseAndSpaceRequest) Execute() (*ListSelectiveSyncsFromWarehouseAndSpace200Response, *http.Response, error) {
	return r.ApiService.ListSelectiveSyncsFromWarehouseAndSpaceExecute(r)
}

/*
ListSelectiveSyncsFromWarehouseAndSpace List Selective Syncs from Warehouse And Space

Returns the schema for a Space Warehouse connection, including Collections and Properties.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@param warehouseId
	@return ApiListSelectiveSyncsFromWarehouseAndSpaceRequest
*/
func (a *ProfilesSyncAPIService) ListSelectiveSyncsFromWarehouseAndSpace(
	ctx context.Context,
	spaceId string,
	warehouseId string,
) ApiListSelectiveSyncsFromWarehouseAndSpaceRequest {
	return ApiListSelectiveSyncsFromWarehouseAndSpaceRequest{
		ApiService:  a,
		ctx:         ctx,
		spaceId:     spaceId,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return ListSelectiveSyncsFromWarehouseAndSpace200Response
func (a *ProfilesSyncAPIService) ListSelectiveSyncsFromWarehouseAndSpaceExecute(
	r ApiListSelectiveSyncsFromWarehouseAndSpaceRequest,
) (*ListSelectiveSyncsFromWarehouseAndSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSelectiveSyncsFromWarehouseAndSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ProfilesSyncAPIService.ListSelectiveSyncsFromWarehouseAndSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/profiles-warehouses/{warehouseId}/selective-syncs"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterValueToString(r.spaceId, "spaceId")),
		-1,
	)
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

type ApiRemoveProfilesWarehouseFromSpaceRequest struct {
	ctx         context.Context
	ApiService  *ProfilesSyncAPIService
	spaceId     string
	warehouseId string
}

func (r ApiRemoveProfilesWarehouseFromSpaceRequest) Execute() (*RemoveProfilesWarehouseFromSpace200Response, *http.Response, error) {
	return r.ApiService.RemoveProfilesWarehouseFromSpaceExecute(r)
}

/*
RemoveProfilesWarehouseFromSpace Remove Profiles Warehouse from Space

Deletes an existing Profiles Warehouse.

• When called, this endpoint may generate the `Profiles Sync Warehouse Deleted` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@param warehouseId
	@return ApiRemoveProfilesWarehouseFromSpaceRequest
*/
func (a *ProfilesSyncAPIService) RemoveProfilesWarehouseFromSpace(
	ctx context.Context,
	spaceId string,
	warehouseId string,
) ApiRemoveProfilesWarehouseFromSpaceRequest {
	return ApiRemoveProfilesWarehouseFromSpaceRequest{
		ApiService:  a,
		ctx:         ctx,
		spaceId:     spaceId,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return RemoveProfilesWarehouseFromSpace200Response
func (a *ProfilesSyncAPIService) RemoveProfilesWarehouseFromSpaceExecute(
	r ApiRemoveProfilesWarehouseFromSpaceRequest,
) (*RemoveProfilesWarehouseFromSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RemoveProfilesWarehouseFromSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ProfilesSyncAPIService.RemoveProfilesWarehouseFromSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/profiles-warehouses/{warehouseId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterValueToString(r.spaceId, "spaceId")),
		-1,
	)
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

type ApiUpdateProfilesWarehouseForSpaceWarehouseRequest struct {
	ctx                                                context.Context
	ApiService                                         *ProfilesSyncAPIService
	spaceId                                            string
	warehouseId                                        string
	updateProfilesWarehouseForSpaceWarehouseAlphaInput *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput
}

func (r ApiUpdateProfilesWarehouseForSpaceWarehouseRequest) UpdateProfilesWarehouseForSpaceWarehouseAlphaInput(
	updateProfilesWarehouseForSpaceWarehouseAlphaInput UpdateProfilesWarehouseForSpaceWarehouseAlphaInput,
) ApiUpdateProfilesWarehouseForSpaceWarehouseRequest {
	r.updateProfilesWarehouseForSpaceWarehouseAlphaInput = &updateProfilesWarehouseForSpaceWarehouseAlphaInput
	return r
}

func (r ApiUpdateProfilesWarehouseForSpaceWarehouseRequest) Execute() (*UpdateProfilesWarehouseForSpaceWarehouse200Response, *http.Response, error) {
	return r.ApiService.UpdateProfilesWarehouseForSpaceWarehouseExecute(r)
}

/*
UpdateProfilesWarehouseForSpaceWarehouse Update Profiles Warehouse for Space Warehouse

Updates an existing Profiles Warehouse.

• When called, this endpoint may generate the `Profiles Sync Warehouse Updated` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@param warehouseId
	@return ApiUpdateProfilesWarehouseForSpaceWarehouseRequest
*/
func (a *ProfilesSyncAPIService) UpdateProfilesWarehouseForSpaceWarehouse(
	ctx context.Context,
	spaceId string,
	warehouseId string,
) ApiUpdateProfilesWarehouseForSpaceWarehouseRequest {
	return ApiUpdateProfilesWarehouseForSpaceWarehouseRequest{
		ApiService:  a,
		ctx:         ctx,
		spaceId:     spaceId,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return UpdateProfilesWarehouseForSpaceWarehouse200Response
func (a *ProfilesSyncAPIService) UpdateProfilesWarehouseForSpaceWarehouseExecute(
	r ApiUpdateProfilesWarehouseForSpaceWarehouseRequest,
) (*UpdateProfilesWarehouseForSpaceWarehouse200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateProfilesWarehouseForSpaceWarehouse200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ProfilesSyncAPIService.UpdateProfilesWarehouseForSpaceWarehouse",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/profiles-warehouses/{warehouseId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterValueToString(r.spaceId, "spaceId")),
		-1,
	)
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateProfilesWarehouseForSpaceWarehouseAlphaInput == nil {
		return localVarReturnValue, nil, reportError(
			"updateProfilesWarehouseForSpaceWarehouseAlphaInput is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.segment.v1alpha+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/vnd.segment.v1alpha+json",
		"application/json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateProfilesWarehouseForSpaceWarehouseAlphaInput
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

type ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest struct {
	ctx                                               context.Context
	ApiService                                        *ProfilesSyncAPIService
	spaceId                                           string
	warehouseId                                       string
	updateSelectiveSyncForWarehouseAndSpaceAlphaInput *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput
}

func (r ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest) UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput(
	updateSelectiveSyncForWarehouseAndSpaceAlphaInput UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput,
) ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest {
	r.updateSelectiveSyncForWarehouseAndSpaceAlphaInput = &updateSelectiveSyncForWarehouseAndSpaceAlphaInput
	return r
}

func (r ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest) Execute() (*UpdateSelectiveSyncForWarehouseAndSpace200Response, *http.Response, error) {
	return r.ApiService.UpdateSelectiveSyncForWarehouseAndSpaceExecute(r)
}

/*
UpdateSelectiveSyncForWarehouseAndSpace Update Selective Sync for Warehouse And Space

Updates the schema for a Space Warehouse connection, including Collections and Properties.

• When called, this endpoint may generate the `Profiles Sync Warehouse Modified` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@param warehouseId
	@return ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest
*/
func (a *ProfilesSyncAPIService) UpdateSelectiveSyncForWarehouseAndSpace(
	ctx context.Context,
	spaceId string,
	warehouseId string,
) ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest {
	return ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest{
		ApiService:  a,
		ctx:         ctx,
		spaceId:     spaceId,
		warehouseId: warehouseId,
	}
}

// Execute executes the request
//
//	@return UpdateSelectiveSyncForWarehouseAndSpace200Response
func (a *ProfilesSyncAPIService) UpdateSelectiveSyncForWarehouseAndSpaceExecute(
	r ApiUpdateSelectiveSyncForWarehouseAndSpaceRequest,
) (*UpdateSelectiveSyncForWarehouseAndSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSelectiveSyncForWarehouseAndSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ProfilesSyncAPIService.UpdateSelectiveSyncForWarehouseAndSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/profiles-warehouses/{warehouseId}/selective-syncs"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterValueToString(r.spaceId, "spaceId")),
		-1,
	)
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"warehouseId"+"}",
		url.PathEscape(parameterValueToString(r.warehouseId, "warehouseId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSelectiveSyncForWarehouseAndSpaceAlphaInput == nil {
		return localVarReturnValue, nil, reportError(
			"updateSelectiveSyncForWarehouseAndSpaceAlphaInput is required and must be specified",
		)
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.segment.v1alpha+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/vnd.segment.v1alpha+json",
		"application/json",
	}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateSelectiveSyncForWarehouseAndSpaceAlphaInput
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
