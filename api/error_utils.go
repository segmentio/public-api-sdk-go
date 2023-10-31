/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"errors"
)

// Unwrap the full errors if it available
func UnwrapFullErrors(err error) *RequestErrorEnvelope {
	if err != nil {
		openApiError := new(GenericOpenAPIError)

		if errors.As(err, &openApiError) {
			model, ok := openApiError.Model().(RequestErrorEnvelope)
			if ok {
				return &model
			}
		}
	}

	return nil
}
