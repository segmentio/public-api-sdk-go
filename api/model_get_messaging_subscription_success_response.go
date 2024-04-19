/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetMessagingSubscriptionSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagingSubscriptionSuccessResponse{}

// GetMessagingSubscriptionSuccessResponse struct for GetMessagingSubscriptionSuccessResponse
type GetMessagingSubscriptionSuccessResponse struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used.
	Type string `json:"type"`
	// The user subscribed, unsubscribed, or on initial status. This is absent if the phone number or email is not found.
	Status *string `json:"status,omitempty"`
	// Optional subscription groups.
	Groups []GroupSubscriptionStatusResponse `json:"groups,omitempty"`
	// The timestamp of this subscription status's last change.
	UpdatedAt *string `json:"updatedAt,omitempty"`
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
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionSuccessResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetMessagingSubscriptionSuccessResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetMessagingSubscriptionSuccessResponse) SetStatus(v string) {
	o.Status = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *GetMessagingSubscriptionSuccessResponse) GetGroups() []GroupSubscriptionStatusResponse {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupSubscriptionStatusResponse
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionSuccessResponse) GetGroupsOk() ([]GroupSubscriptionStatusResponse, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GetMessagingSubscriptionSuccessResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupSubscriptionStatusResponse and assigns it to the Groups field.
func (o *GetMessagingSubscriptionSuccessResponse) SetGroups(v []GroupSubscriptionStatusResponse) {
	o.Groups = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetMessagingSubscriptionSuccessResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingSubscriptionSuccessResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetMessagingSubscriptionSuccessResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetMessagingSubscriptionSuccessResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o GetMessagingSubscriptionSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMessagingSubscriptionSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
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
