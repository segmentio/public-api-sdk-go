/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GroupSubscriptionStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSubscriptionStatusResponse{}

// GroupSubscriptionStatusResponse struct for GroupSubscriptionStatusResponse
type GroupSubscriptionStatusResponse struct {
	// Name of the group.
	Name string `json:"name"`
	// The user subscribed, unsubscribed, or on initial status.
	Status string `json:"status"`
	// The group id.
	Id string `json:"id"`
	// The timestamp of this subscription status's last change.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewGroupSubscriptionStatusResponse instantiates a new GroupSubscriptionStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSubscriptionStatusResponse(
	name string,
	status string,
	id string,
) *GroupSubscriptionStatusResponse {
	this := GroupSubscriptionStatusResponse{}
	this.Name = name
	this.Status = status
	this.Id = id
	return &this
}

// NewGroupSubscriptionStatusResponseWithDefaults instantiates a new GroupSubscriptionStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSubscriptionStatusResponseWithDefaults() *GroupSubscriptionStatusResponse {
	this := GroupSubscriptionStatusResponse{}
	return &this
}

// GetName returns the Name field value
func (o *GroupSubscriptionStatusResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupSubscriptionStatusResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupSubscriptionStatusResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *GroupSubscriptionStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GroupSubscriptionStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GroupSubscriptionStatusResponse) SetStatus(v string) {
	o.Status = v
}

// GetId returns the Id field value
func (o *GroupSubscriptionStatusResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupSubscriptionStatusResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupSubscriptionStatusResponse) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GroupSubscriptionStatusResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSubscriptionStatusResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupSubscriptionStatusResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GroupSubscriptionStatusResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o GroupSubscriptionStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSubscriptionStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["id"] = o.Id
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGroupSubscriptionStatusResponse struct {
	value *GroupSubscriptionStatusResponse
	isSet bool
}

func (v NullableGroupSubscriptionStatusResponse) Get() *GroupSubscriptionStatusResponse {
	return v.value
}

func (v *NullableGroupSubscriptionStatusResponse) Set(val *GroupSubscriptionStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSubscriptionStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSubscriptionStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSubscriptionStatusResponse(
	val *GroupSubscriptionStatusResponse,
) *NullableGroupSubscriptionStatusResponse {
	return &NullableGroupSubscriptionStatusResponse{value: val, isSet: true}
}

func (v NullableGroupSubscriptionStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSubscriptionStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
