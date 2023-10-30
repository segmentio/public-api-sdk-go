/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EchoV1Output Echo response.
type EchoV1Output struct {
	// The HTTP method used for this round-trip.  Currently, this endpoint supports only `get` and `post` methods.
	Method string `json:"method"`
	// The string passed in the `message` input field.
	Message string `json:"message"`
	// The request's HTTP headers.
	Headers map[string]interface{} `json:"headers"`
}

// NewEchoV1Output instantiates a new EchoV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEchoV1Output(method string, message string, headers map[string]interface{}) *EchoV1Output {
	this := EchoV1Output{}
	this.Method = method
	this.Message = message
	this.Headers = headers
	return &this
}

// NewEchoV1OutputWithDefaults instantiates a new EchoV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEchoV1OutputWithDefaults() *EchoV1Output {
	this := EchoV1Output{}
	return &this
}

// GetMethod returns the Method field value
func (o *EchoV1Output) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *EchoV1Output) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *EchoV1Output) SetMethod(v string) {
	o.Method = v
}

// GetMessage returns the Message field value
func (o *EchoV1Output) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *EchoV1Output) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *EchoV1Output) SetMessage(v string) {
	o.Message = v
}

// GetHeaders returns the Headers field value
func (o *EchoV1Output) GetHeaders() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *EchoV1Output) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *EchoV1Output) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

func (o EchoV1Output) MarshalJSON() ([]byte, error) {
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

type NullableEchoV1Output struct {
	value *EchoV1Output
	isSet bool
}

func (v NullableEchoV1Output) Get() *EchoV1Output {
	return v.value
}

func (v *NullableEchoV1Output) Set(val *EchoV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableEchoV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableEchoV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEchoV1Output(val *EchoV1Output) *NullableEchoV1Output {
	return &NullableEchoV1Output{value: val, isSet: true}
}

func (v NullableEchoV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEchoV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
