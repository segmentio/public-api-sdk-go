/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.5.0
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

// EventsAPIService EventsAPI service
type EventsAPIService service

type ApiGetEventsVolumeFromWorkspaceRequest struct {
	ctx         context.Context
	ApiService  *EventsAPIService
	granularity *string
	startTime   *string
	endTime     *string
	groupBy     *[]string
	sourceId    *[]string
	eventName   *[]string
	eventType   *[]string
	appVersion  *[]string
	pagination  *PaginationInput
}

// The size of each bucket in the requested window.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) Granularity(
	granularity string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.granularity = &granularity
	return r
}

// The ISO8601 formatted timestamp that corresponds to the beginning of the requested time frame, inclusive.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) StartTime(
	startTime string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.startTime = &startTime
	return r
}

// The ISO8601 formatted timestamp that corresponds to the end of the requested time frame, noninclusive. Segment recommends that you lag queries 1 minute behind clock time to reduce the risk for latency to impact the counts.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) EndTime(
	endTime string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.endTime = &endTime
	return r
}

// A comma-delimited list of strings that represents the dimensions to group the result by. The options are: &#x60;eventName&#x60;, &#x60;eventType&#x60; and &#x60;source&#x60;.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) GroupBy(
	groupBy []string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.groupBy = &groupBy
	return r
}

// A list of strings which filters the results to the given SourceIds.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) SourceId(
	sourceId []string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.sourceId = &sourceId
	return r
}

// A list of strings which filters the results to the given EventNames.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) EventName(
	eventName []string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.eventName = &eventName
	return r
}

// A list of strings which filters the results to the given EventTypes.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) EventType(
	eventType []string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.eventType = &eventType
	return r
}

// A list of strings which filters the results to the given AppVersions.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) AppVersion(
	appVersion []string,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.appVersion = &appVersion
	return r
}

// Pagination input for event volume by Workspace.  This parameter exists in v1.
func (r ApiGetEventsVolumeFromWorkspaceRequest) Pagination(
	pagination PaginationInput,
) ApiGetEventsVolumeFromWorkspaceRequest {
	r.pagination = &pagination
	return r
}

func (r ApiGetEventsVolumeFromWorkspaceRequest) Execute() (*GetEventsVolumeFromWorkspace200Response, *http.Response, error) {
	return r.ApiService.GetEventsVolumeFromWorkspaceExecute(r)
}

/*
GetEventsVolumeFromWorkspace Get Events Volume from Workspace

Enumerates the Workspace event volumes over time in minute increments.

The rate limit for this endpoint is 60 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEventsVolumeFromWorkspaceRequest
*/
func (a *EventsAPIService) GetEventsVolumeFromWorkspace(
	ctx context.Context,
) ApiGetEventsVolumeFromWorkspaceRequest {
	return ApiGetEventsVolumeFromWorkspaceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEventsVolumeFromWorkspace200Response
func (a *EventsAPIService) GetEventsVolumeFromWorkspaceExecute(
	r ApiGetEventsVolumeFromWorkspaceRequest,
) (*GetEventsVolumeFromWorkspace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEventsVolumeFromWorkspace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"EventsAPIService.GetEventsVolumeFromWorkspace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/volume"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.granularity == nil {
		return localVarReturnValue, nil, reportError(
			"granularity is required and must be specified",
		)
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, reportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, reportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "")
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "groupBy", r.groupBy, "csv")
	}
	if r.sourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceId", r.sourceId, "csv")
	}
	if r.eventName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventName", r.eventName, "csv")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "csv")
	}
	if r.appVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "appVersion", r.appVersion, "csv")
	}
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
