/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetMessagingSubscriptionSuccessResponse struct for GetMessagingSubscriptionSuccessResponse
type GetMessagingSubscriptionSuccessResponse struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used. Either EMAIL or SMS.
	Type string `json:"type"`
	// The user subscribed, unsubscribed, or on initial status. This is absent if the phone number or email is not found.
	Status *string `json:"status,omitempty"`
}

// NewGetMessagingSubscriptionSuccessResponse instantiates a new GetMessagingSubscriptionSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingSubscriptionSuccessResponse(
	key string,
	type_ string,
) *GetMessagingSubscriptionSuccessResponse {
	this := GetMessagingSubscriptionSuccessResponse{}
	this.Key = key
	this.Type = type_
	return &this
}

// NewGetMessagingSubscriptionSuccessResponseWithDefaults instantiates a new GetMessagingSubscriptionSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingSubscriptionSuccessResponseWithDefaults() *GetMessagingSubscriptionSuccessResponse {
	this := GetMessagingSubscriptionSuccessResponse{}
	return &this
}

// GetKey returns the Key field value
func (o *GetMessagingSubscriptionSuccessResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionSuccessResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetMessagingSubscriptionSuccessResponse) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value
func (o *GetMessagingSubscriptionSuccessResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionSuccessResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetMessagingSubscriptionSuccessResponse) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetMessagingSubscriptionSuccessResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionSuccessResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetMessagingSubscriptionSuccessResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetMessagingSubscriptionSuccessResponse) SetStatus(v string) {
	o.Status = &v
}

func (o GetMessagingSubscriptionSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetMessagingSubscriptionSuccessResponse struct {
	value *GetMessagingSubscriptionSuccessResponse
	isSet bool
}

func (v NullableGetMessagingSubscriptionSuccessResponse) Get() *GetMessagingSubscriptionSuccessResponse {
	return v.value
}

func (v *NullableGetMessagingSubscriptionSuccessResponse) Set(
	val *GetMessagingSubscriptionSuccessResponse,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingSubscriptionSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingSubscriptionSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingSubscriptionSuccessResponse(
	val *GetMessagingSubscriptionSuccessResponse,
) *NullableGetMessagingSubscriptionSuccessResponse {
	return &NullableGetMessagingSubscriptionSuccessResponse{value: val, isSet: true}
}

func (v NullableGetMessagingSubscriptionSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMessagingSubscriptionSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
