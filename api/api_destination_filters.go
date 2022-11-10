/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
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
	"strings"
)


// DestinationFiltersApiService DestinationFiltersApi service
type DestinationFiltersApiService service

type ApiCreateFilterForDestinationRequest struct {
	ctx context.Context
	ApiService *DestinationFiltersApiService
	destinationId string
	createFilterForDestinationV1Input *CreateFilterForDestinationV1Input
}

func (r ApiCreateFilterForDestinationRequest) CreateFilterForDestinationV1Input(createFilterForDestinationV1Input CreateFilterForDestinationV1Input) ApiCreateFilterForDestinationRequest {
	r.createFilterForDestinationV1Input = &createFilterForDestinationV1Input
	return r
}

func (r ApiCreateFilterForDestinationRequest) Execute() (*CreateFilterForDestination200Response, *http.Response, error) {
	return r.ApiService.CreateFilterForDestinationExecute(r)
}

/*
CreateFilterForDestination Create Filter for Destination

Creates a filter in a Destination.



When called, this endpoint may generate the `Destination Filter Created` [Audit Trail](/tag/Audit-Trail) event.
      

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param destinationId
 @return ApiCreateFilterForDestinationRequest
*/
func (a *DestinationFiltersApiService) CreateFilterForDestination(ctx context.Context, destinationId string) ApiCreateFilterForDestinationRequest {
	return ApiCreateFilterForDestinationRequest{
		ApiService: a,
		ctx: ctx,
		destinationId: destinationId,
	}
}

// Execute executes the request
//  @return CreateFilterForDestination200Response
func (a *DestinationFiltersApiService) CreateFilterForDestinationExecute(r ApiCreateFilterForDestinationRequest) (*CreateFilterForDestination200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFilterForDestination200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DestinationFiltersApiService.CreateFilterForDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/destination/{destinationId}/filters"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterToString(r.destinationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFilterForDestinationV1Input == nil {
		return localVarReturnValue, nil, reportError("createFilterForDestinationV1Input is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createFilterForDestinationV1Input
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFilterInDestinationRequest struct {
	ctx context.Context
	ApiService *DestinationFiltersApiService
	destinationId string
	filterId string
}

func (r ApiGetFilterInDestinationRequest) Execute() (*GetFilterInDestination200Response, *http.Response, error) {
	return r.ApiService.GetFilterInDestinationExecute(r)
}

/*
GetFilterInDestination Get Filter in Destination

Gets a Destination filter by id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param destinationId
 @param filterId
 @return ApiGetFilterInDestinationRequest
*/
func (a *DestinationFiltersApiService) GetFilterInDestination(ctx context.Context, destinationId string, filterId string) ApiGetFilterInDestinationRequest {
	return ApiGetFilterInDestinationRequest{
		ApiService: a,
		ctx: ctx,
		destinationId: destinationId,
		filterId: filterId,
	}
}

// Execute executes the request
//  @return GetFilterInDestination200Response
func (a *DestinationFiltersApiService) GetFilterInDestinationExecute(r ApiGetFilterInDestinationRequest) (*GetFilterInDestination200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFilterInDestination200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DestinationFiltersApiService.GetFilterInDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/destination/{destinationId}/filters/{filterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterToString(r.destinationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filterId"+"}", url.PathEscape(parameterToString(r.filterId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListFiltersFromDestinationRequest struct {
	ctx context.Context
	ApiService *DestinationFiltersApiService
	destinationId string
	pagination *PaginationInput
}

// Pagination options.  This parameter exists in alpha.
func (r ApiListFiltersFromDestinationRequest) Pagination(pagination PaginationInput) ApiListFiltersFromDestinationRequest {
	r.pagination = &pagination
	return r
}

func (r ApiListFiltersFromDestinationRequest) Execute() (*ListFiltersFromDestination200Response, *http.Response, error) {
	return r.ApiService.ListFiltersFromDestinationExecute(r)
}

/*
ListFiltersFromDestination List Filters from Destination

Lists filters for a Destination.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param destinationId
 @return ApiListFiltersFromDestinationRequest
*/
func (a *DestinationFiltersApiService) ListFiltersFromDestination(ctx context.Context, destinationId string) ApiListFiltersFromDestinationRequest {
	return ApiListFiltersFromDestinationRequest{
		ApiService: a,
		ctx: ctx,
		destinationId: destinationId,
	}
}

// Execute executes the request
//  @return ListFiltersFromDestination200Response
func (a *DestinationFiltersApiService) ListFiltersFromDestinationExecute(r ApiListFiltersFromDestinationRequest) (*ListFiltersFromDestination200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListFiltersFromDestination200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DestinationFiltersApiService.ListFiltersFromDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/destination/{destinationId}/filters"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterToString(r.destinationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	localVarQueryParams.Add("pagination", parameterToString(*r.pagination, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPreviewDestinationFilterRequest struct {
	ctx context.Context
	ApiService *DestinationFiltersApiService
	previewDestinationFilterV1Input *PreviewDestinationFilterV1Input
}

func (r ApiPreviewDestinationFilterRequest) PreviewDestinationFilterV1Input(previewDestinationFilterV1Input PreviewDestinationFilterV1Input) ApiPreviewDestinationFilterRequest {
	r.previewDestinationFilterV1Input = &previewDestinationFilterV1Input
	return r
}

func (r ApiPreviewDestinationFilterRequest) Execute() (*PreviewDestinationFilter200Response, *http.Response, error) {
	return r.ApiService.PreviewDestinationFilterExecute(r)
}

/*
PreviewDestinationFilter Preview Destination Filter

Simulates the application of a Destination filter to a provided JSON payload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPreviewDestinationFilterRequest
*/
func (a *DestinationFiltersApiService) PreviewDestinationFilter(ctx context.Context) ApiPreviewDestinationFilterRequest {
	return ApiPreviewDestinationFilterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PreviewDestinationFilter200Response
func (a *DestinationFiltersApiService) PreviewDestinationFilterExecute(r ApiPreviewDestinationFilterRequest) (*PreviewDestinationFilter200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PreviewDestinationFilter200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DestinationFiltersApiService.PreviewDestinationFilter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/destination/filters/preview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.previewDestinationFilterV1Input == nil {
		return localVarReturnValue, nil, reportError("previewDestinationFilterV1Input is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.previewDestinationFilterV1Input
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveFilterFromDestinationRequest struct {
	ctx context.Context
	ApiService *DestinationFiltersApiService
	destinationId string
	filterId string
}

func (r ApiRemoveFilterFromDestinationRequest) Execute() (*RemoveFilterFromDestination200Response, *http.Response, error) {
	return r.ApiService.RemoveFilterFromDestinationExecute(r)
}

/*
RemoveFilterFromDestination Remove Filter from Destination

Deletes a Destination filter.



When called, this endpoint may generate the `Destination Filter Deleted` [Audit Trail](/tag/Audit-Trail) event.
      

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param destinationId
 @param filterId
 @return ApiRemoveFilterFromDestinationRequest
*/
func (a *DestinationFiltersApiService) RemoveFilterFromDestination(ctx context.Context, destinationId string, filterId string) ApiRemoveFilterFromDestinationRequest {
	return ApiRemoveFilterFromDestinationRequest{
		ApiService: a,
		ctx: ctx,
		destinationId: destinationId,
		filterId: filterId,
	}
}

// Execute executes the request
//  @return RemoveFilterFromDestination200Response
func (a *DestinationFiltersApiService) RemoveFilterFromDestinationExecute(r ApiRemoveFilterFromDestinationRequest) (*RemoveFilterFromDestination200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RemoveFilterFromDestination200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DestinationFiltersApiService.RemoveFilterFromDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/destination/{destinationId}/filters/{filterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterToString(r.destinationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filterId"+"}", url.PathEscape(parameterToString(r.filterId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateFilterForDestinationRequest struct {
	ctx context.Context
	ApiService *DestinationFiltersApiService
	destinationId string
	filterId string
	updateFilterForDestinationV1Input *UpdateFilterForDestinationV1Input
}

func (r ApiUpdateFilterForDestinationRequest) UpdateFilterForDestinationV1Input(updateFilterForDestinationV1Input UpdateFilterForDestinationV1Input) ApiUpdateFilterForDestinationRequest {
	r.updateFilterForDestinationV1Input = &updateFilterForDestinationV1Input
	return r
}

func (r ApiUpdateFilterForDestinationRequest) Execute() (*UpdateFilterForDestination200Response, *http.Response, error) {
	return r.ApiService.UpdateFilterForDestinationExecute(r)
}

/*
UpdateFilterForDestination Update Filter for Destination

Updates a filter in a Destination.

When called, this endpoint may generate one or more of the following [Audit Trail](/tag/Audit-Trail) events:
* Destination Filter Enabled
* Destination Filter Disabled
      

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param destinationId
 @param filterId
 @return ApiUpdateFilterForDestinationRequest
*/
func (a *DestinationFiltersApiService) UpdateFilterForDestination(ctx context.Context, destinationId string, filterId string) ApiUpdateFilterForDestinationRequest {
	return ApiUpdateFilterForDestinationRequest{
		ApiService: a,
		ctx: ctx,
		destinationId: destinationId,
		filterId: filterId,
	}
}

// Execute executes the request
//  @return UpdateFilterForDestination200Response
func (a *DestinationFiltersApiService) UpdateFilterForDestinationExecute(r ApiUpdateFilterForDestinationRequest) (*UpdateFilterForDestination200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateFilterForDestination200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DestinationFiltersApiService.UpdateFilterForDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/destination/{destinationId}/filters/{filterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterToString(r.destinationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filterId"+"}", url.PathEscape(parameterToString(r.filterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateFilterForDestinationV1Input == nil {
		return localVarReturnValue, nil, reportError("updateFilterForDestinationV1Input is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.segment.v1alpha+json", "application/vnd.segment.v1beta+json", "application/vnd.segment.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateFilterForDestinationV1Input
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
