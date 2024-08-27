/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MessagesSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagesSubscriptionRequest{}

// MessagesSubscriptionRequest struct for MessagesSubscriptionRequest
type MessagesSubscriptionRequest struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used.
	Type string `json:"type"`
	// The user subscribed, unsubscribed, or on initial status globally.
	Status *string `json:"status,omitempty"`
	// Optional groups subscription status.
	Groups []GroupSubscriptionStatus `json:"groups,omitempty"`
}

// NewMessagesSubscriptionRequest instantiates a new MessagesSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesSubscriptionRequest(key string, type_ string) *MessagesSubscriptionRequest {
	this := MessagesSubscriptionRequest{}
	this.Key = key
	this.Type = type_
	return &this
}

// NewMessagesSubscriptionRequestWithDefaults instantiates a new MessagesSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesSubscriptionRequestWithDefaults() *MessagesSubscriptionRequest {
	this := MessagesSubscriptionRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *MessagesSubscriptionRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscriptionRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MessagesSubscriptionRequest) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value
func (o *MessagesSubscriptionRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessagesSubscriptionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessagesSubscriptionRequest) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessagesSubscriptionRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesSubscriptionRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessagesSubscriptionRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MessagesSubscriptionRequest) SetStatus(v string) {
	o.Status = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *MessagesSubscriptionRequest) GetGroups() []GroupSubscriptionStatus {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupSubscriptionStatus
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesSubscriptionRequest) GetGroupsOk() ([]GroupSubscriptionStatus, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MessagesSubscriptionRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupSubscriptionStatus and assigns it to the Groups field.
func (o *MessagesSubscriptionRequest) SetGroups(v []GroupSubscriptionStatus) {
	o.Groups = v
}

func (o MessagesSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagesSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableMessagesSubscriptionRequest struct {
	value *MessagesSubscriptionRequest
	isSet bool
}

func (v NullableMessagesSubscriptionRequest) Get() *MessagesSubscriptionRequest {
	return v.value
}

func (v *NullableMessagesSubscriptionRequest) Set(val *MessagesSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesSubscriptionRequest(
	val *MessagesSubscriptionRequest,
) *NullableMessagesSubscriptionRequest {
	return &NullableMessagesSubscriptionRequest{value: val, isSet: true}
}

func (v NullableMessagesSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagesSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
