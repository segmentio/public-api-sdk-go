/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
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

// SpacesApiService SpacesApi service
type SpacesApiService service

type ApiBatchQueryMessagingSubscriptionsForSpaceRequest struct {
	ctx                                                context.Context
	ApiService                                         *SpacesApiService
	spaceId                                            string
	batchQueryMessagingSubscriptionsForSpaceAlphaInput *BatchQueryMessagingSubscriptionsForSpaceAlphaInput
}

func (r ApiBatchQueryMessagingSubscriptionsForSpaceRequest) BatchQueryMessagingSubscriptionsForSpaceAlphaInput(
	batchQueryMessagingSubscriptionsForSpaceAlphaInput BatchQueryMessagingSubscriptionsForSpaceAlphaInput,
) ApiBatchQueryMessagingSubscriptionsForSpaceRequest {
	r.batchQueryMessagingSubscriptionsForSpaceAlphaInput = &batchQueryMessagingSubscriptionsForSpaceAlphaInput
	return r
}

func (r ApiBatchQueryMessagingSubscriptionsForSpaceRequest) Execute() (*BatchQueryMessagingSubscriptionsForSpace200Response, *http.Response, error) {
	return r.ApiService.BatchQueryMessagingSubscriptionsForSpaceExecute(r)
}

/*
BatchQueryMessagingSubscriptionsForSpace Batch Query Messaging Subscriptions for Space

Get Messaging Subscriptions for space.

• This endpoint is in **Alpha** testing.  Please submit any feedback by sending email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Spaces feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@return ApiBatchQueryMessagingSubscriptionsForSpaceRequest
*/
func (a *SpacesApiService) BatchQueryMessagingSubscriptionsForSpace(
	ctx context.Context,
	spaceId string,
) ApiBatchQueryMessagingSubscriptionsForSpaceRequest {
	return ApiBatchQueryMessagingSubscriptionsForSpaceRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return BatchQueryMessagingSubscriptionsForSpace200Response
func (a *SpacesApiService) BatchQueryMessagingSubscriptionsForSpaceExecute(
	r ApiBatchQueryMessagingSubscriptionsForSpaceRequest,
) (*BatchQueryMessagingSubscriptionsForSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchQueryMessagingSubscriptionsForSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpacesApiService.BatchQueryMessagingSubscriptionsForSpace",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/messaging-subscriptions/batch"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterToString(r.spaceId, "")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchQueryMessagingSubscriptionsForSpaceAlphaInput == nil {
		return localVarReturnValue, nil, reportError(
			"batchQueryMessagingSubscriptionsForSpaceAlphaInput is required and must be specified",
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
	localVarPostBody = r.batchQueryMessagingSubscriptionsForSpaceAlphaInput
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

type ApiGetSpaceRequest struct {
	ctx        context.Context
	ApiService *SpacesApiService
	spaceId    string
}

func (r ApiGetSpaceRequest) Execute() (*GetSpace200Response, *http.Response, error) {
	return r.ApiService.GetSpaceExecute(r)
}

/*
GetSpace Get Space

Returns the Space by id.

• This endpoint is in **Alpha** testing.  Please submit any feedback by sending email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Spaces feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@return ApiGetSpaceRequest
*/
func (a *SpacesApiService) GetSpace(ctx context.Context, spaceId string) ApiGetSpaceRequest {
	return ApiGetSpaceRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return GetSpace200Response
func (a *SpacesApiService) GetSpaceExecute(
	r ApiGetSpaceRequest,
) (*GetSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpacesApiService.GetSpace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterToString(r.spaceId, "")),
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

type ApiReplaceMessagingSubscriptionsInSpacesRequest struct {
	ctx                                             context.Context
	ApiService                                      *SpacesApiService
	spaceId                                         string
	replaceMessagingSubscriptionsInSpacesAlphaInput *ReplaceMessagingSubscriptionsInSpacesAlphaInput
}

func (r ApiReplaceMessagingSubscriptionsInSpacesRequest) ReplaceMessagingSubscriptionsInSpacesAlphaInput(
	replaceMessagingSubscriptionsInSpacesAlphaInput ReplaceMessagingSubscriptionsInSpacesAlphaInput,
) ApiReplaceMessagingSubscriptionsInSpacesRequest {
	r.replaceMessagingSubscriptionsInSpacesAlphaInput = &replaceMessagingSubscriptionsInSpacesAlphaInput
	return r
}

func (r ApiReplaceMessagingSubscriptionsInSpacesRequest) Execute() (*ReplaceMessagingSubscriptionsInSpaces200Response, *http.Response, error) {
	return r.ApiService.ReplaceMessagingSubscriptionsInSpacesExecute(r)
}

/*
ReplaceMessagingSubscriptionsInSpaces Replace Messaging Subscriptions in Spaces

Replace Messaging Subscriptions in Spaces.

• This endpoint is in **Alpha** testing.  Please submit any feedback by sending email to friends@segment.com.

• In order to successfully call this endpoint, the specified Workspace needs to have the Spaces feature enabled. Please reach out to your customer success manager for more information.

The rate limit for this endpoint is 60 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId
	@return ApiReplaceMessagingSubscriptionsInSpacesRequest
*/
func (a *SpacesApiService) ReplaceMessagingSubscriptionsInSpaces(
	ctx context.Context,
	spaceId string,
) ApiReplaceMessagingSubscriptionsInSpacesRequest {
	return ApiReplaceMessagingSubscriptionsInSpacesRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return ReplaceMessagingSubscriptionsInSpaces200Response
func (a *SpacesApiService) ReplaceMessagingSubscriptionsInSpacesExecute(
	r ApiReplaceMessagingSubscriptionsInSpacesRequest,
) (*ReplaceMessagingSubscriptionsInSpaces200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplaceMessagingSubscriptionsInSpaces200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"SpacesApiService.ReplaceMessagingSubscriptionsInSpaces",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/messaging-subscriptions"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"spaceId"+"}",
		url.PathEscape(parameterToString(r.spaceId, "")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.replaceMessagingSubscriptionsInSpacesAlphaInput == nil {
		return localVarReturnValue, nil, reportError(
			"replaceMessagingSubscriptionsInSpacesAlphaInput is required and must be specified",
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
	localVarPostBody = r.replaceMessagingSubscriptionsInSpacesAlphaInput
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
