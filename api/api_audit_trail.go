/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
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

// AuditTrailApiService AuditTrailApi service
type AuditTrailApiService service

type ApiListAuditEventsRequest struct {
	ctx          context.Context
	ApiService   *AuditTrailApiService
	pagination   *PaginationInput
	startTime    *string
	endTime      *string
	resourceId   *string
	resourceType *string
}

// Defines the pagination parameters.  This parameter exists in v1.
func (r ApiListAuditEventsRequest) Pagination(
	pagination PaginationInput,
) ApiListAuditEventsRequest {
	r.pagination = &pagination
	return r
}

// Filter response to events that happened after this time.  This parameter exists in v1.
func (r ApiListAuditEventsRequest) StartTime(startTime string) ApiListAuditEventsRequest {
	r.startTime = &startTime
	return r
}

// Filter response to events that happened before this time. Defaults to the current time, or the end time from the pagination cursor.  This parameter exists in v1.
func (r ApiListAuditEventsRequest) EndTime(endTime string) ApiListAuditEventsRequest {
	r.endTime = &endTime
	return r
}

// Filter response to events that affect a specific resource, for example, a single Source.  This parameter exists in v1.
func (r ApiListAuditEventsRequest) ResourceId(resourceId string) ApiListAuditEventsRequest {
	r.resourceId = &resourceId
	return r
}

// Filter response to events that affect a specific type, for example, Sources, Warehouses, and Tracking Plans.  This parameter exists in v1.
func (r ApiListAuditEventsRequest) ResourceType(resourceType string) ApiListAuditEventsRequest {
	r.resourceType = &resourceType
	return r
}

func (r ApiListAuditEventsRequest) Execute() (*ListAuditEvents200Response, *http.Response, error) {
	return r.ApiService.ListAuditEventsExecute(r)
}

/*
ListAuditEvents List Audit Events

Returns a list of Audit Trail events.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAuditEventsRequest
*/
func (a *AuditTrailApiService) ListAuditEvents(ctx context.Context) ApiListAuditEventsRequest {
	return ApiListAuditEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListAuditEvents200Response
func (a *AuditTrailApiService) ListAuditEventsExecute(
	r ApiListAuditEventsRequest,
) (*ListAuditEvents200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListAuditEvents200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"AuditTrailApiService.ListAuditEvents",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audit-events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pagination == nil {
		return localVarReturnValue, nil, reportError("pagination is required and must be specified")
	}

	if r.startTime != nil {
		localVarQueryParams.Add("startTime", parameterToString(*r.startTime, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.resourceId != nil {
		localVarQueryParams.Add("resourceId", parameterToString(*r.resourceId, ""))
	}
	if r.resourceType != nil {
		localVarQueryParams.Add("resourceType", parameterToString(*r.resourceType, ""))
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
