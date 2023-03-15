/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EchoAlphaOutput Echo response.
type EchoAlphaOutput struct {
	// The HTTP method used for this round-trip.  Currently, this endpoint supports only `get` and `post` methods.
	Method string `json:"method"`
	// The string passed in the `message` input field.
	Message string `json:"message"`
	// The request's HTTP headers.
	Headers map[string]interface{} `json:"headers"`
}

// NewEchoAlphaOutput instantiates a new EchoAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEchoAlphaOutput(
	method string,
	message string,
	headers map[string]interface{},
) *EchoAlphaOutput {
	this := EchoAlphaOutput{}
	this.Method = method
	this.Message = message
	this.Headers = headers
	return &this
}

// NewEchoAlphaOutputWithDefaults instantiates a new EchoAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEchoAlphaOutputWithDefaults() *EchoAlphaOutput {
	this := EchoAlphaOutput{}
	return &this
}

// GetMethod returns the Method field value
func (o *EchoAlphaOutput) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *EchoAlphaOutput) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *EchoAlphaOutput) SetMethod(v string) {
	o.Method = v
}

// GetMessage returns the Message field value
func (o *EchoAlphaOutput) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *EchoAlphaOutput) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *EchoAlphaOutput) SetMessage(v string) {
	o.Message = v
}

// GetHeaders returns the Headers field value
func (o *EchoAlphaOutput) GetHeaders() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *EchoAlphaOutput) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *EchoAlphaOutput) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

func (o EchoAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["headers"] = o.Headers
	}
	return json.Marshal(toSerialize)
}

type NullableEchoAlphaOutput struct {
	value *EchoAlphaOutput
	isSet bool
}

func (v NullableEchoAlphaOutput) Get() *EchoAlphaOutput {
	return v.value
}

func (v *NullableEchoAlphaOutput) Set(val *EchoAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEchoAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEchoAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEchoAlphaOutput(val *EchoAlphaOutput) *NullableEchoAlphaOutput {
	return &NullableEchoAlphaOutput{value: val, isSet: true}
}

func (v NullableEchoAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEchoAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
