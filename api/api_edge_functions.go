/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
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

// EdgeFunctionsApiService EdgeFunctionsApi service
type EdgeFunctionsApiService service

type ApiCreateEdgeFunctionsRequest struct {
	ctx                           context.Context
	ApiService                    *EdgeFunctionsApiService
	sourceId                      string
	createEdgeFunctionsAlphaInput *CreateEdgeFunctionsAlphaInput
}

func (r ApiCreateEdgeFunctionsRequest) CreateEdgeFunctionsAlphaInput(
	createEdgeFunctionsAlphaInput CreateEdgeFunctionsAlphaInput,
) ApiCreateEdgeFunctionsRequest {
	r.createEdgeFunctionsAlphaInput = &createEdgeFunctionsAlphaInput
	return r
}

func (r ApiCreateEdgeFunctionsRequest) Execute() (*CreateEdgeFunctions200Response, *http.Response, error) {
	return r.ApiService.CreateEdgeFunctionsExecute(r)
}

/*
CreateEdgeFunctions Create Edge Functions

Create EdgeFunctions for your Source, given a valid upload URL for an Edge Functions bundle.

**Note**: In order to successfully call this endpoint, the specified Workspace needs to have the Edge Functions feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceId
	@return ApiCreateEdgeFunctionsRequest
*/
func (a *EdgeFunctionsApiService) CreateEdgeFunctions(
	ctx context.Context,
	sourceId string,
) ApiCreateEdgeFunctionsRequest {
	return ApiCreateEdgeFunctionsRequest{
		ApiService: a,
		ctx:        ctx,
		sourceId:   sourceId,
	}
}

// Execute executes the request
//
//	@return CreateEdgeFunctions200Response
func (a *EdgeFunctionsApiService) CreateEdgeFunctionsExecute(
	r ApiCreateEdgeFunctionsRequest,
) (*CreateEdgeFunctions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateEdgeFunctions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"EdgeFunctionsApiService.CreateEdgeFunctions",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/{sourceId}/edge-functions"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"sourceId"+"}",
		url.PathEscape(parameterToString(r.sourceId, "")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createEdgeFunctionsAlphaInput == nil {
		return localVarReturnValue, nil, reportError(
			"createEdgeFunctionsAlphaInput is required and must be specified",
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
	localVarPostBody = r.createEdgeFunctionsAlphaInput
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

type ApiDisableEdgeFunctionsRequest struct {
	ctx        context.Context
	ApiService *EdgeFunctionsApiService
	sourceId   string
}

func (r ApiDisableEdgeFunctionsRequest) Execute() (*DisableEdgeFunctions200Response, *http.Response, error) {
	return r.ApiService.DisableEdgeFunctionsExecute(r)
}

/*
DisableEdgeFunctions Disable Edge Functions

Disable Edge Functions for your Source.

**Note**: In order to successfully call this endpoint, the specified Workspace needs to have the Edge Functions feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceId
	@return ApiDisableEdgeFunctionsRequest
*/
func (a *EdgeFunctionsApiService) DisableEdgeFunctions(
	ctx context.Context,
	sourceId string,
) ApiDisableEdgeFunctionsRequest {
	return ApiDisableEdgeFunctionsRequest{
		ApiService: a,
		ctx:        ctx,
		sourceId:   sourceId,
	}
}

// Execute executes the request
//
//	@return DisableEdgeFunctions200Response
func (a *EdgeFunctionsApiService) DisableEdgeFunctionsExecute(
	r ApiDisableEdgeFunctionsRequest,
) (*DisableEdgeFunctions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DisableEdgeFunctions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"EdgeFunctionsApiService.DisableEdgeFunctions",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/{sourceId}/edge-functions/disable"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"sourceId"+"}",
		url.PathEscape(parameterToString(r.sourceId, "")),
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

type ApiGenerateUploadURLForEdgeFunctionsRequest struct {
	ctx        context.Context
	ApiService *EdgeFunctionsApiService
	sourceId   string
}

func (r ApiGenerateUploadURLForEdgeFunctionsRequest) Execute() (*GenerateUploadURLForEdgeFunctions200Response, *http.Response, error) {
	return r.ApiService.GenerateUploadURLForEdgeFunctionsExecute(r)
}

/*
GenerateUploadURLForEdgeFunctions Generate Upload URL for Edge Functions

Generate a temporary upload URL, that can be used to upload an Edge Functions bundle.

**Note**: In order to successfully call this endpoint, the specified Workspace needs to have the Edge Functions feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceId
	@return ApiGenerateUploadURLForEdgeFunctionsRequest
*/
func (a *EdgeFunctionsApiService) GenerateUploadURLForEdgeFunctions(
	ctx context.Context,
	sourceId string,
) ApiGenerateUploadURLForEdgeFunctionsRequest {
	return ApiGenerateUploadURLForEdgeFunctionsRequest{
		ApiService: a,
		ctx:        ctx,
		sourceId:   sourceId,
	}
}

// Execute executes the request
//
//	@return GenerateUploadURLForEdgeFunctions200Response
func (a *EdgeFunctionsApiService) GenerateUploadURLForEdgeFunctionsExecute(
	r ApiGenerateUploadURLForEdgeFunctionsRequest,
) (*GenerateUploadURLForEdgeFunctions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GenerateUploadURLForEdgeFunctions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"EdgeFunctionsApiService.GenerateUploadURLForEdgeFunctions",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/{sourceId}/edge-functions/upload-url"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"sourceId"+"}",
		url.PathEscape(parameterToString(r.sourceId, "")),
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

type ApiGetLatestFromEdgeFunctionsRequest struct {
	ctx        context.Context
	ApiService *EdgeFunctionsApiService
	sourceId   string
}

func (r ApiGetLatestFromEdgeFunctionsRequest) Execute() (*GetLatestFromEdgeFunctions200Response, *http.Response, error) {
	return r.ApiService.GetLatestFromEdgeFunctionsExecute(r)
}

/*
GetLatestFromEdgeFunctions Get Latest from Edge Functions

Get the latest Edge Functions for your Source.

**Note**: In order to successfully call this endpoint, the specified Workspace needs to have the Edge Functions feature enabled. Please reach out to your customer success manager for more information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceId
	@return ApiGetLatestFromEdgeFunctionsRequest
*/
func (a *EdgeFunctionsApiService) GetLatestFromEdgeFunctions(
	ctx context.Context,
	sourceId string,
) ApiGetLatestFromEdgeFunctionsRequest {
	return ApiGetLatestFromEdgeFunctionsRequest{
		ApiService: a,
		ctx:        ctx,
		sourceId:   sourceId,
	}
}

// Execute executes the request
//
//	@return GetLatestFromEdgeFunctions200Response
func (a *EdgeFunctionsApiService) GetLatestFromEdgeFunctionsExecute(
	r ApiGetLatestFromEdgeFunctionsRequest,
) (*GetLatestFromEdgeFunctions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLatestFromEdgeFunctions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(
		r.ctx,
		"EdgeFunctionsApiService.GetLatestFromEdgeFunctions",
	)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sources/{sourceId}/edge-functions/latest"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"sourceId"+"}",
		url.PathEscape(parameterToString(r.sourceId, "")),
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
