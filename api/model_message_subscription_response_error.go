/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 35.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MessageSubscriptionResponseError struct for MessageSubscriptionResponseError
type MessageSubscriptionResponseError struct {
	// Error code.
	Code string `json:"code"`
	// Error message.
	Message string `json:"message"`
}

// NewMessageSubscriptionResponseError instantiates a new MessageSubscriptionResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSubscriptionResponseError(
	code string,
	message string,
) *MessageSubscriptionResponseError {
	this := MessageSubscriptionResponseError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewMessageSubscriptionResponseErrorWithDefaults instantiates a new MessageSubscriptionResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSubscriptionResponseErrorWithDefaults() *MessageSubscriptionResponseError {
	this := MessageSubscriptionResponseError{}
	return &this
}

// GetCode returns the Code field value
func (o *MessageSubscriptionResponseError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponseError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *MessageSubscriptionResponseError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *MessageSubscriptionResponseError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *MessageSubscriptionResponseError) SetMessage(v string) {
	o.Message = v
}

func (o MessageSubscriptionResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableMessageSubscriptionResponseError struct {
	value *MessageSubscriptionResponseError
	isSet bool
}

func (v NullableMessageSubscriptionResponseError) Get() *MessageSubscriptionResponseError {
	return v.value
}

func (v *NullableMessageSubscriptionResponseError) Set(val *MessageSubscriptionResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSubscriptionResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSubscriptionResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSubscriptionResponseError(
	val *MessageSubscriptionResponseError,
) *NullableMessageSubscriptionResponseError {
	return &NullableMessageSubscriptionResponseError{value: val, isSet: true}
}

func (v NullableMessageSubscriptionResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSubscriptionResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
