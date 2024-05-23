/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.2.0
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

// ReverseETLAPIService ReverseETLAPI service
type ReverseETLAPIService service

type ApiCreateReverseETLManualSyncRequest struct {
	ctx                             context.Context
	ApiService                      *ReverseETLAPIService
	createReverseETLManualSyncInput *CreateReverseETLManualSyncInput
}

func (r ApiCreateReverseETLManualSyncRequest) CreateReverseETLManualSyncInput(
	createReverseETLManualSyncInput CreateReverseETLManualSyncInput,
) ApiCreateReverseETLManualSyncRequest {
	r.createReverseETLManualSyncInput = &createReverseETLManualSyncInput
	return r
}

func (r ApiCreateReverseETLManualSyncRequest) Execute() (*CreateReverseETLManualSync200Response, *http.Response, error) {
	return r.ApiService.CreateReverseETLManualSyncExecute(r)
}

/*
CreateReverseETLManualSync Create Reverse ETL Manual Sync

Triggers a manual sync for a Reverse ETL Connection.
The `subscriptionId` is the id that follows after /mappings/ in the URL of the mapping page.
For example, the `subscriptionId` would be '2' for this sync: https://app.segment.com/example-workspace/reverse-etl/destinations/example-destination/sources/example-source/instances/1/mappings/2/source-id/3/model-id/4/sync-details

The rate limit for this endpoint is 20 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateReverseETLManualSyncRequest
*/
func (a *ReverseETLAPIService) CreateReverseETLManualSync(
	ctx context.Context,
) ApiCreateReverseETLManualSyncRequest {
	return ApiCreateReverseETLManualSyncRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateReverseETLManualSync200Response
func (a *ReverseETLAPIService) CreateReverseETLManualSyncExecute(
	r ApiCreateReverseETLManualSyncRequest,
) (*CreateReverseETLManualSync200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateReverseETLManualSync200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.CreateReverseETLManualSync",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-syncs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createReverseETLManualSyncInput == nil {
		return localVarReturnValue, nil, reportError(
			"createReverseETLManualSyncInput is required and must be specified",
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
	localVarPostBody = r.createReverseETLManualSyncInput
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

type ApiCreateReverseEtlModelRequest struct {
	ctx                        context.Context
	ApiService                 *ReverseETLAPIService
	createReverseEtlModelInput *CreateReverseEtlModelInput
}

func (r ApiCreateReverseEtlModelRequest) CreateReverseEtlModelInput(
	createReverseEtlModelInput CreateReverseEtlModelInput,
) ApiCreateReverseEtlModelRequest {
	r.createReverseEtlModelInput = &createReverseEtlModelInput
	return r
}

func (r ApiCreateReverseEtlModelRequest) Execute() (*CreateReverseEtlModel201Response, *http.Response, error) {
	return r.ApiService.CreateReverseEtlModelExecute(r)
}

/*
CreateReverseEtlModel Create Reverse Etl Model

Creates a new Reverse ETL Model.

• When called, this endpoint may generate the `Model Created` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateReverseEtlModelRequest
*/
func (a *ReverseETLAPIService) CreateReverseEtlModel(
	ctx context.Context,
) ApiCreateReverseEtlModelRequest {
	return ApiCreateReverseEtlModelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateReverseEtlModel201Response
func (a *ReverseETLAPIService) CreateReverseEtlModelExecute(
	r ApiCreateReverseEtlModelRequest,
) (*CreateReverseEtlModel201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateReverseEtlModel201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.CreateReverseEtlModel",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createReverseEtlModelInput == nil {
		return localVarReturnValue, nil, reportError(
			"createReverseEtlModelInput is required and must be specified",
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
	localVarPostBody = r.createReverseEtlModelInput
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

type ApiDeleteReverseEtlModelRequest struct {
	ctx        context.Context
	ApiService *ReverseETLAPIService
	modelId    string
}

func (r ApiDeleteReverseEtlModelRequest) Execute() (*DeleteReverseEtlModel200Response, *http.Response, error) {
	return r.ApiService.DeleteReverseEtlModelExecute(r)
}

/*
DeleteReverseEtlModel Delete Reverse Etl Model

Deletes an existing Model.

• When called, this endpoint may generate the `Model Deleted` event in the [audit trail](/tag/Audit-Trail).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param modelId
	@return ApiDeleteReverseEtlModelRequest
*/
func (a *ReverseETLAPIService) DeleteReverseEtlModel(
	ctx context.Context,
	modelId string,
) ApiDeleteReverseEtlModelRequest {
	return ApiDeleteReverseEtlModelRequest{
		ApiService: a,
		ctx:        ctx,
		modelId:    modelId,
	}
}

// Execute executes the request
//
//	@return DeleteReverseEtlModel200Response
func (a *ReverseETLAPIService) DeleteReverseEtlModelExecute(
	r ApiDeleteReverseEtlModelRequest,
) (*DeleteReverseEtlModel200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteReverseEtlModel200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.DeleteReverseEtlModel",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-models/{modelId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"modelId"+"}",
		url.PathEscape(parameterValueToString(r.modelId, "modelId")),
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

type ApiGetReverseETLSyncStatusRequest struct {
	ctx        context.Context
	ApiService *ReverseETLAPIService
	modelId    string
	syncId     string
}

func (r ApiGetReverseETLSyncStatusRequest) Execute() (*GetReverseETLSyncStatus200Response, *http.Response, error) {
	return r.ApiService.GetReverseETLSyncStatusExecute(r)
}

/*
GetReverseETLSyncStatus Get Reverse ETL Sync Status

Get the sync status for a Reverse ETL sync.
The sync status includes all detailed information about the sync - sync status, duration, details about the extract and load phase if applicable, etc.

The rate limit for this endpoint is 250 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param modelId
	@param syncId
	@return ApiGetReverseETLSyncStatusRequest
*/
func (a *ReverseETLAPIService) GetReverseETLSyncStatus(
	ctx context.Context,
	modelId string,
	syncId string,
) ApiGetReverseETLSyncStatusRequest {
	return ApiGetReverseETLSyncStatusRequest{
		ApiService: a,
		ctx:        ctx,
		modelId:    modelId,
		syncId:     syncId,
	}
}

// Execute executes the request
//
//	@return GetReverseETLSyncStatus200Response
func (a *ReverseETLAPIService) GetReverseETLSyncStatusExecute(
	r ApiGetReverseETLSyncStatusRequest,
) (*GetReverseETLSyncStatus200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetReverseETLSyncStatus200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.GetReverseETLSyncStatus",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-models/{modelId}/syncs/{syncId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"modelId"+"}",
		url.PathEscape(parameterValueToString(r.modelId, "modelId")),
		-1,
	)
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"syncId"+"}",
		url.PathEscape(parameterValueToString(r.syncId, "syncId")),
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

type ApiGetReverseEtlModelRequest struct {
	ctx        context.Context
	ApiService *ReverseETLAPIService
	modelId    string
}

func (r ApiGetReverseEtlModelRequest) Execute() (*GetReverseEtlModel200Response, *http.Response, error) {
	return r.ApiService.GetReverseEtlModelExecute(r)
}

/*
GetReverseEtlModel Get Reverse Etl Model

Returns a Reverse ETL Model by its id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param modelId
	@return ApiGetReverseEtlModelRequest
*/
func (a *ReverseETLAPIService) GetReverseEtlModel(
	ctx context.Context,
	modelId string,
) ApiGetReverseEtlModelRequest {
	return ApiGetReverseEtlModelRequest{
		ApiService: a,
		ctx:        ctx,
		modelId:    modelId,
	}
}

// Execute executes the request
//
//	@return GetReverseEtlModel200Response
func (a *ReverseETLAPIService) GetReverseEtlModelExecute(
	r ApiGetReverseEtlModelRequest,
) (*GetReverseEtlModel200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetReverseEtlModel200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.GetReverseEtlModel",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-models/{modelId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"modelId"+"}",
		url.PathEscape(parameterValueToString(r.modelId, "modelId")),
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

type ApiListReverseEtlModelsRequest struct {
	ctx        context.Context
	ApiService *ReverseETLAPIService
	pagination *PaginationInput
}

// Defines the pagination parameters.  This parameter exists in alpha.
func (r ApiListReverseEtlModelsRequest) Pagination(
	pagination PaginationInput,
) ApiListReverseEtlModelsRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListReverseEtlModelsRequest) Execute() (*ListReverseEtlModels200Response, *http.Response, error) {
	return r.ApiService.ListReverseEtlModelsExecute(r)
}

/*
ListReverseEtlModels List Reverse Etl Models

Returns a list of Reverse ETL Models.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListReverseEtlModelsRequest
*/
func (a *ReverseETLAPIService) ListReverseEtlModels(
	ctx context.Context,
) ApiListReverseEtlModelsRequest {
	return ApiListReverseEtlModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListReverseEtlModels200Response
func (a *ReverseETLAPIService) ListReverseEtlModelsExecute(
	r ApiListReverseEtlModelsRequest,
) (*ListReverseEtlModels200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListReverseEtlModels200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.ListReverseEtlModels",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-models"

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

type ApiUpdateReverseEtlModelRequest struct {
	ctx                        context.Context
	ApiService                 *ReverseETLAPIService
	modelId                    string
	updateReverseEtlModelInput *UpdateReverseEtlModelInput
}

func (r ApiUpdateReverseEtlModelRequest) UpdateReverseEtlModelInput(
	updateReverseEtlModelInput UpdateReverseEtlModelInput,
) ApiUpdateReverseEtlModelRequest {
	r.updateReverseEtlModelInput = &updateReverseEtlModelInput
	return r
}

func (r ApiUpdateReverseEtlModelRequest) Execute() (*UpdateReverseEtlModel200Response, *http.Response, error) {
	return r.ApiService.UpdateReverseEtlModelExecute(r)
}

/*
UpdateReverseEtlModel Update Reverse Etl Model

Updates an existing Reverse ETL Model.

• When called, this endpoint may generate one or more of the following [audit trail](/tag/Audit-Trail) events:* Model Settings Saved
* Model State Change Toggled

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param modelId
	@return ApiUpdateReverseEtlModelRequest
*/
func (a *ReverseETLAPIService) UpdateReverseEtlModel(
	ctx context.Context,
	modelId string,
) ApiUpdateReverseEtlModelRequest {
	return ApiUpdateReverseEtlModelRequest{
		ApiService: a,
		ctx:        ctx,
		modelId:    modelId,
	}
}

// Execute executes the request
//
//	@return UpdateReverseEtlModel200Response
func (a *ReverseETLAPIService) UpdateReverseEtlModelExecute(
	r ApiUpdateReverseEtlModelRequest,
) (*UpdateReverseEtlModel200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateReverseEtlModel200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"ReverseETLAPIService.UpdateReverseEtlModel",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reverse-etl-models/{modelId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"modelId"+"}",
		url.PathEscape(parameterValueToString(r.modelId, "modelId")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateReverseEtlModelInput == nil {
		return localVarReturnValue, nil, reportError(
			"updateReverseEtlModelInput is required and must be specified",
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
	localVarPostBody = r.updateReverseEtlModelInput
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
