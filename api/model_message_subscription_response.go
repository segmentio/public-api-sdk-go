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

// MessageSubscriptionResponse struct for MessageSubscriptionResponse
type MessageSubscriptionResponse struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used. Either EMAIL or SMS.
	Type string `json:"type"`
	// The user subscribed, unsubscribed, or on initial status.
	Status string `json:"status"`
	// Error messages.
	Errors []MessageSubscriptionResponseError `json:"errors,omitempty"`
}

// NewMessageSubscriptionResponse instantiates a new MessageSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSubscriptionResponse(
	key string,
	type_ string,
	status string,
) *MessageSubscriptionResponse {
	this := MessageSubscriptionResponse{}
	this.Key = key
	this.Type = type_
	this.Status = status
	return &this
}

// NewMessageSubscriptionResponseWithDefaults instantiates a new MessageSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSubscriptionResponseWithDefaults() *MessageSubscriptionResponse {
	this := MessageSubscriptionResponse{}
	return &this
}

// GetKey returns the Key field value
func (o *MessageSubscriptionResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MessageSubscriptionResponse) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value
func (o *MessageSubscriptionResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageSubscriptionResponse) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *MessageSubscriptionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MessageSubscriptionResponse) SetStatus(v string) {
	o.Status = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *MessageSubscriptionResponse) GetErrors() []MessageSubscriptionResponseError {
	if o == nil || o.Errors == nil {
		var ret []MessageSubscriptionResponseError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetErrorsOk() ([]MessageSubscriptionResponseError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MessageSubscriptionResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []MessageSubscriptionResponseError and assigns it to the Errors field.
func (o *MessageSubscriptionResponse) SetErrors(v []MessageSubscriptionResponseError) {
	o.Errors = v
}

func (o MessageSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableMessageSubscriptionResponse struct {
	value *MessageSubscriptionResponse
	isSet bool
}

func (v NullableMessageSubscriptionResponse) Get() *MessageSubscriptionResponse {
	return v.value
}

func (v *NullableMessageSubscriptionResponse) Set(val *MessageSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSubscriptionResponse(
	val *MessageSubscriptionResponse,
) *NullableMessageSubscriptionResponse {
	return &NullableMessageSubscriptionResponse{value: val, isSet: true}
}

func (v NullableMessageSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
