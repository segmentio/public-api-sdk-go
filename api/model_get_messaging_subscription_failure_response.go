/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetMessagingSubscriptionFailureResponse struct for GetMessagingSubscriptionFailureResponse
type GetMessagingSubscriptionFailureResponse struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// This will be the exact type as given in the request.
	Type string `json:"type"`
	// Per key errors, such as validation errors.
	Errors []MessageSubscriptionResponseError `json:"errors"`
}

// NewGetMessagingSubscriptionFailureResponse instantiates a new GetMessagingSubscriptionFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingSubscriptionFailureResponse(
	key string,
	type_ string,
	errors []MessageSubscriptionResponseError,
) *GetMessagingSubscriptionFailureResponse {
	this := GetMessagingSubscriptionFailureResponse{}
	this.Key = key
	this.Type = type_
	this.Errors = errors
	return &this
}

// NewGetMessagingSubscriptionFailureResponseWithDefaults instantiates a new GetMessagingSubscriptionFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingSubscriptionFailureResponseWithDefaults() *GetMessagingSubscriptionFailureResponse {
	this := GetMessagingSubscriptionFailureResponse{}
	return &this
}

// GetKey returns the Key field value
func (o *GetMessagingSubscriptionFailureResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionFailureResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetMessagingSubscriptionFailureResponse) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value
func (o *GetMessagingSubscriptionFailureResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionFailureResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetMessagingSubscriptionFailureResponse) SetType(v string) {
	o.Type = v
}

// GetErrors returns the Errors field value
func (o *GetMessagingSubscriptionFailureResponse) GetErrors() []MessageSubscriptionResponseError {
	if o == nil {
		var ret []MessageSubscriptionResponseError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionFailureResponse) GetErrorsOk() ([]MessageSubscriptionResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *GetMessagingSubscriptionFailureResponse) SetErrors(v []MessageSubscriptionResponseError) {
	o.Errors = v
}

func (o GetMessagingSubscriptionFailureResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableGetMessagingSubscriptionFailureResponse struct {
	value *GetMessagingSubscriptionFailureResponse
	isSet bool
}

func (v NullableGetMessagingSubscriptionFailureResponse) Get() *GetMessagingSubscriptionFailureResponse {
	return v.value
}

func (v *NullableGetMessagingSubscriptionFailureResponse) Set(
	val *GetMessagingSubscriptionFailureResponse,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingSubscriptionFailureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingSubscriptionFailureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingSubscriptionFailureResponse(
	val *GetMessagingSubscriptionFailureResponse,
) *NullableGetMessagingSubscriptionFailureResponse {
	return &NullableGetMessagingSubscriptionFailureResponse{value: val, isSet: true}
}

func (v NullableGetMessagingSubscriptionFailureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMessagingSubscriptionFailureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
