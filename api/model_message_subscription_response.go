/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MessageSubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSubscriptionResponse{}

// MessageSubscriptionResponse struct for MessageSubscriptionResponse
type MessageSubscriptionResponse struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used. Either SMS, EMAIL or WHATSAPP.
	Type string `json:"type"`
	// The user subscribed, unsubscribed, or on initial status.
	Status *string `json:"status,omitempty"`
	// Error messages.
	Errors []MessageSubscriptionResponseError `json:"errors,omitempty"`
	// Optional subscription groups.
	Groups []UpdateGroupSubscriptionStatusResponse `json:"groups,omitempty"`
}

// NewMessageSubscriptionResponse instantiates a new MessageSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSubscriptionResponse(key string, type_ string) *MessageSubscriptionResponse {
	this := MessageSubscriptionResponse{}
	this.Key = key
	this.Type = type_
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

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessageSubscriptionResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessageSubscriptionResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MessageSubscriptionResponse) SetStatus(v string) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *MessageSubscriptionResponse) GetErrors() []MessageSubscriptionResponseError {
	if o == nil || IsNil(o.Errors) {
		var ret []MessageSubscriptionResponseError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetErrorsOk() ([]MessageSubscriptionResponseError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MessageSubscriptionResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []MessageSubscriptionResponseError and assigns it to the Errors field.
func (o *MessageSubscriptionResponse) SetErrors(v []MessageSubscriptionResponseError) {
	o.Errors = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *MessageSubscriptionResponse) GetGroups() []UpdateGroupSubscriptionStatusResponse {
	if o == nil || IsNil(o.Groups) {
		var ret []UpdateGroupSubscriptionStatusResponse
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSubscriptionResponse) GetGroupsOk() ([]UpdateGroupSubscriptionStatusResponse, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MessageSubscriptionResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []UpdateGroupSubscriptionStatusResponse and assigns it to the Groups field.
func (o *MessageSubscriptionResponse) SetGroups(v []UpdateGroupSubscriptionStatusResponse) {
	o.Groups = v
}

func (o MessageSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
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
