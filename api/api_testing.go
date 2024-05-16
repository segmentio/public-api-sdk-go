/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
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

// TestingAPIService TestingAPI service
type TestingAPIService service

type ApiEchoRequest struct {
	ctx                    context.Context
	ApiService             *TestingAPIService
	message                *string
	delay                  *float32
	triggerError           *bool
	triggerMultipleErrors  *bool
	triggerUnexpectedError *bool
	statusCode             *float32
}

// Sets the response &#x60;message&#x60; field. The response contains this field&#39;s entry.  This parameter exists in alpha.
func (r ApiEchoRequest) Message(message string) ApiEchoRequest {
	r.message = &message
	return r
}

// The desired response delay, in milliseconds.  This parameter exists in alpha.
func (r ApiEchoRequest) Delay(delay float32) ApiEchoRequest {
	r.delay = &delay
	return r
}

// If &#x60;true&#x60;, returns an HTTP &#x60;4xx&#x60; error that contains the string in &#x60;message&#x60;.  This parameter exists in alpha.
func (r ApiEchoRequest) TriggerError(triggerError bool) ApiEchoRequest {
	r.triggerError = &triggerError
	return r
}

// If &#x60;true&#x60;, returns an HTTP &#x60;4xx&#x60; error that contains the value of the &#x60;message&#x60; field in the error message array.  This has no effect if the request sets &#x60;triggerError&#x60;.  This parameter exists in alpha.
func (r ApiEchoRequest) TriggerMultipleErrors(triggerMultipleErrors bool) ApiEchoRequest {
	r.triggerMultipleErrors = &triggerMultipleErrors
	return r
}

// If &#x60;true&#x60;, triggers a &#x60;500&#x60; error.  This has no effect if the request sets either &#x60;triggerError&#x60; or &#x60;triggerMultipleErrors&#x60;.  This parameter exists in alpha.
func (r ApiEchoRequest) TriggerUnexpectedError(triggerUnexpectedError bool) ApiEchoRequest {
	r.triggerUnexpectedError = &triggerUnexpectedError
	return r
}

// Sets the HTTP status code to return.  This parameter exists in alpha.
func (r ApiEchoRequest) StatusCode(statusCode float32) ApiEchoRequest {
	r.statusCode = &statusCode
	return r
}

func (r ApiEchoRequest) Execute() (*Echo200Response, *http.Response, error) {
	return r.ApiService.EchoExecute(r)
}

/*
Echo Echo

Public Echo endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEchoRequest
*/
func (a *TestingAPIService) Echo(ctx context.Context) ApiEchoRequest {
	return ApiEchoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Echo200Response
func (a *TestingAPIService) EchoExecute(
	r ApiEchoRequest,
) (*Echo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Echo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestingAPIService.Echo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/echo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.message == nil {
		return localVarReturnValue, nil, reportError("message is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "message", r.message, "")
	if r.delay != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delay", r.delay, "")
	}
	if r.triggerError != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "triggerError", r.triggerError, "")
	}
	if r.triggerMultipleErrors != nil {
		parameterAddToHeaderOrQuery(
			localVarQueryParams,
			"triggerMultipleErrors",
			r.triggerMultipleErrors,
			"",
		)
	}
	if r.triggerUnexpectedError != nil {
		parameterAddToHeaderOrQuery(
			localVarQueryParams,
			"triggerUnexpectedError",
			r.triggerUnexpectedError,
			"",
		)
	}
	if r.statusCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "statusCode", r.statusCode, "")
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
